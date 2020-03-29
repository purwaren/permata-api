package permata

import (
	"context"
	"os"

	"github.com/avast/retry-go"
	"github.com/juju/errors"
	cmnCtx "github.com/kedifeles/go-my-commons/context"
	"github.com/kedifeles/go-my-commons/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Permata provide access to Permata API
type Permata struct {
	api    *api
	config Config
}

const maxRetryAttempts uint = 2

// New return new instance of BCA
func New(config Config) *Permata {
	bca := Permata{
		config: config,
		api:    newAPI(config),
	}

	logger.SetOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {

		fileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
			Filename: config.LogPath,
			MaxSize:  500, // megabytes
			// MaxBackups: 3,
			// MaxAge:     28, // days
		})
		stdoutWriteSyncer := zapcore.AddSync(os.Stdout)

		return zapcore.NewCore(
			zapcore.NewJSONEncoder(logger.DefaultEncoderConfig),
			zapcore.NewMultiWriteSyncer(fileWriteSyncer, stdoutWriteSyncer),
			zap.InfoLevel,
		)

		// return core
	}))

	return &bca
}

var errESB14009 = errors.New("Auth err from BCA API (ESB-14-009)")

// func errorIfErrCodeESB14009(dtoError Error) error {
// 	if dtoError.ErrorCode == "ESB-14-009" {
// 		return errESB14009
// 	}
// 	return nil
// }

func (b *Permata) retryDecision(ctx context.Context) func(err error) bool {
	return func(err error) bool {
		return err == errESB14009
	}
}

func (b *Permata) retryOptions(ctx context.Context) []retry.Option {
	return []retry.Option{
		retry.Attempts(maxRetryAttempts),
		retry.RetryIf(b.retryDecision(ctx)),
		retry.OnRetry(func(n uint, err error) {
			b.log(ctx).Infof("=== START ON RETRY === [Attempts: %d Err: %+v]", n, err)
			_, err = b.DoAuthentication(ctx)
			if err != nil {
				b.log(ctx).Error(err)
			}
			b.log(ctx).Infof("=== END ON RETRY ===")
		}),
	}
}

// === misc func ===

func (b *Permata) log(ctx context.Context) *zap.SugaredLogger {
	return logger.Logger(cmnCtx.With(ctx, cmnCtx.SessID(b.api.sessID)))
}
