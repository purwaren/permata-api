package permata

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/juju/errors"
	"github.com/kedifeles/go-my-commons/logger"
	"github.com/lithammer/shortuuid"
	"go.uber.org/zap"

	cmnCtx "github.com/kedifeles/go-my-commons/context"
)

var (
	httpHeaderChannelID    string = "ChannelID"
	httpHeaderCredentialID string = "CredentialID"
)

type api struct {
	config     Config
	httpClient *http.Client // for postGetToken only

	mutex       sync.Mutex
	accessToken string
	sessID      string
}

func newAPI(config Config) *api {

	httpClient := cleanhttp.DefaultPooledClient()

	api := api{config: config,
		httpClient: httpClient,
	}

	return &api
}

func (api *api) setAccessToken(accessToken string) {
	api.mutex.Lock()
	defer api.mutex.Unlock()

	newSessID := shortuuid.New()

	api.accessToken = accessToken
	api.sessID = newSessID
}

// === AUTH ===

func (api *api) postGetToken(ctx context.Context) (*OAuth2Resp, error) {
	urlTarget, err := buildURL(api.config.URL, "/apiservice/oauth/token", url.Values{})
	if err != nil {
		return nil, errors.Trace(err)
	}

	form := url.Values{"grant_type": []string{"client_credentials"}}
	bodyReq := strings.NewReader(form.Encode())

	req, err := http.NewRequest(http.MethodPost, urlTarget, bodyReq)
	if err != nil {
		return nil, errors.Trace(err)
	}
	req = req.WithContext(ctx)

	timestamp := time.Now().Format("2006-01-02T15:04:05.999Z07:00")
	sign, _, err := GenerateSignature(api.config.StaticKey, api.config.APIKey, timestamp, form.Encode())
	if err != nil {
		return nil, errors.Trace(err)
	}

	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("API-Key", api.config.APIKey)
	req.Header.Set("OAUTH-Signature", sign)
	req.Header.Set("OAUTH-Timestamp", timestamp)
	req.SetBasicAuth(api.config.ClientID, api.config.ClientSecret)

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer resp.Body.Close()

	bodyRespBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Trace(err)
	}

	api.log(ctx).Info(resp.StatusCode)
	api.log(ctx).Info(string(bodyRespBytes))

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyRespBytes))

	var dtoResp OAuth2Resp
	err = json.NewDecoder(resp.Body).Decode(&dtoResp)

	if err != nil {
		return nil, errors.Trace(err)
	}

	return &dtoResp, nil
}

// === misc func ===
func (api *api) log(ctx context.Context) *zap.SugaredLogger {
	return logger.Logger(cmnCtx.With(ctx, cmnCtx.SessID(api.sessID)))
}

func buildURL(baseURL, paths string, query url.Values) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", errors.Trace(err)
	}

	u.Path = path.Join(u.Path, paths)
	u.RawQuery = query.Encode()

	return u.String(), nil
}
