package bootstrap

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/hathawway/back-gloss/internal/config"
	"github.com/hathawway/back-gloss/internal/transport"
	"github.com/hathawway/back-gloss/internal/transport/rest_api"
)

func ApiEntryPoint(ctx context.Context, cfg *config.Config) (func(context.Context) error, error) {
	mngr := transport.NewManager()

	mngr.AddServer(rest_api.NewServer(cfg))

	go func() {
		err := mngr.Start(ctx)
		if err != nil {
			logrus.Fatalf("error starting server %s", err.Error())
		}
	}()

	return mngr.Stop, nil
}
