package permata

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSignature(t *testing.T) {
	type args struct {
		staticKey     string
		apiKeyOrToken string
		timestamp     string
		requestBody   string
	}
	tests := []struct {
		name          string
		args          args
		wantSign      string
		wantStrToSign string
		wantErr       bool
	}{
		{name: "from doc: GetToken", args: args{
			staticKey:     "WD62811305f7f796a480bb2c53d76099",
			apiKeyOrToken: "cd86fe4b-bc4f-4200-8ad4-d04971c65ac6",
			timestamp:     "2017-12-09T03:52:01.000+07:00",
			requestBody:   "grant_type=client_credentials",
		},
			wantStrToSign: "cd86fe4b-bc4f-4200-8ad4-d04971c65ac6:2017-12-09T03:52:01.000+07:00:grant_type=client_credentials",
			wantSign:      "YcKit0AOOJns0z/DWS+VJDHPJr/MWGiqZzqzFHIkBgo=",
			wantErr:       false,
		},
		// {name: "from doc: BalanceInquiry", args: args{
		// 	staticKey:     "WD62811305f7f796a480bb2c53d76099",
		// 	apiKeyOrToken: "15hBgzeXpK6M8WZJqPwL5z615iFxFS2OW1hKTnUV6c17OEuJKSXsKy",
		// 	timestamp:     "2017-11-07T10:22:57.000",
		// 	requestBody:   `{"BalInqRq":{"MsgRqHdr":{"RequestTimestamp":"2017-07-21T14:32:01.000+07:00","CustRefID":"0878987654321"},"InqInfo":{"AccountNumber":"701075323"}}}`,
		// },
		// 	wantStrToSign: `15hBgzeXpK6M8WZJqPwL5z615iFxFS2OW1hKTnUV6c17OEuJKSXsKy:2017-11-07T10:22:57.000:{"BalInqRq":{"MsgRqHdr":{"RequestTimestamp":"2017-07-21T14:32:01.000+07:00","CustRefID":"0878987654321"},"InqInfo":{"AccountNumber":"701075323"}}}`,
		// 	wantSign:      "YSHxWxEc2rohVDCI8f/H1S3oNBn7l1wf6hNcAscAdb4=",
		// 	wantErr:       false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSign, gotStrToSign, err := GenerateSignature(tt.args.staticKey, tt.args.apiKeyOrToken, tt.args.timestamp, tt.args.requestBody)
			require.NoError(t, err)
			require.Equal(t, tt.wantStrToSign, gotStrToSign)
			require.Equal(t, tt.wantSign, gotSign)
		})
	}
}
