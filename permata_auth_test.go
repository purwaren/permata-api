package permata_test

import (
	"context"
	"os"
	"testing"

	"github.com/purwaren/permata-api"
	"github.com/stretchr/testify/require"
)

func TestPermata_Auth_integration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	t.Run("DoAuthentication", func(t *testing.T) {
		givenConfig := permata.Config{
			URL:          os.Getenv("URL"),
			APIKey:       os.Getenv("API_KEY"),
			StaticKey:    os.Getenv("STATIC_KEY"),
			ClientID:     os.Getenv("CLIENT_ID"),
			ClientSecret: os.Getenv("CLIENT_SECRET"),
		}

		b := permata.New(givenConfig)

		// resp based on sandbox resp
		dtoResp, err := b.DoAuthentication(context.Background())
		require.NoError(t, err)
		// require.Empty(t, dtoResp.Error)
		require.NotEmpty(t, dtoResp)
	})
}
