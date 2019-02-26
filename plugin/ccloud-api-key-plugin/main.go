package main

import (
	"context"
	golog "log"
	"os"

	plugin "github.com/hashicorp/go-plugin"
	"github.com/sirupsen/logrus"

	chttp "github.com/confluentinc/ccloud-sdk-go"
	authv1 "github.com/confluentinc/ccloudapis/auth/v1"
	log "github.com/confluentinc/cli/log"
	metric "github.com/confluentinc/cli/metric"
	"github.com/confluentinc/cli/shared"
	"github.com/confluentinc/cli/shared/api-key"
)

// Compile-time check for Interface adherence
var _ chttp.APIKey = (*ApiKey)(nil)

func main() {
	var logger *log.Logger
	{
		logger = log.New()
		logger.Log("msg", "hello")
		defer logger.Log("msg", "goodbye")

		f, err := os.OpenFile("/tmp/confluent-api-key-plugin.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
		check(err)
		logger.SetLevel(logrus.DebugLevel)
		logger.Logger.Out = f
	}

	var metricSink shared.MetricSink
	{
		metricSink = metric.NewSink()
	}

	var config *shared.Config
	{
		config = shared.NewConfig(&shared.Config{
			MetricSink: metricSink,
			Logger:     logger,
		})
		err := config.Load()
		if err != nil && err != shared.ErrNoConfig {
			logger.WithError(err).Errorf("unable to load config")
		}
	}

	var impl *ApiKey
	{
		client := chttp.NewClientWithJWT(context.Background(), config.AuthToken, config.AuthURL, config.Logger)
		impl = &ApiKey{Logger: logger, Client: client}
	}

	shared.PluginMap[api_key.Name] = &api_key.Plugin{Impl: impl}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins:         shared.PluginMap,
		GRPCServer:      plugin.DefaultGRPCServer,
	})
}

type ApiKey struct {
	Logger *log.Logger
	Client *chttp.Client
}

func (c *ApiKey) Create(ctx context.Context, key *authv1.ApiKey) (*authv1.ApiKey, error) {
	c.Logger.Log("msg", "apiKey.Create()")
	ret, err := c.Client.APIKey.Create(ctx, key)
	return ret, shared.ConvertAPIError(err)
}

func (c *ApiKey) Delete(ctx context.Context, key *authv1.ApiKey) error {
	c.Logger.Log("msg", "apiKey.Delete()")
	err := c.Client.APIKey.Delete(ctx, key)
	return shared.ConvertAPIError(err)
}

func (c *ApiKey) List(ctx context.Context, key *authv1.ApiKey) ([]*authv1.ApiKey, error) {
	c.Logger.Log("msg", "apiKey.List()")
	ret, err := c.Client.APIKey.List(ctx, key)
	return ret, shared.ConvertAPIError(err)
}

func check(err error) {
	if err != nil {
		golog.Fatal(err)
	}
}
