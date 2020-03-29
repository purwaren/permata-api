package permata

import (
	"context"

	"github.com/juju/errors"
	cmnCtx "github.com/kedifeles/go-my-commons/context"
)

// DoAuthentication authenticate using OAuth2
func (b *Permata) DoAuthentication(ctx context.Context) (*OAuth2Resp, error) {
	ctx = cmnCtx.With(ctx, cmnCtx.SessID(b.api.sessID))

	b.log(ctx).Info("=== START DO_AUTH ===")

	dtoResp, err := b.api.postGetToken(ctx)
	if err != nil {
		b.log(ctx).Error(errors.Details(err))
		return nil, errors.Trace(err)
	}

	b.api.setAccessToken(dtoResp.AccessToken)

	b.log(ctx).Info("=== END DO_AUTH ===")

	return dtoResp, nil
}
