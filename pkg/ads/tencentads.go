package ads

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"net/http"
	"strconv"
	"time"
)

type SDKClient struct {
	http.RoundTripper
	Config         *config.SDKConfig
	Client         *api.APIClient
	Ctx            *context.Context
	Version        string
	ApiVersion     string
	middlewareList []Middleware
}

func Init(cfg *config.SDKConfig) *SDKClient {
	version := "2.0.0"
	apiVersion := "v1.1"
	ctx := context.Background()
	nonce := uuid.NewV4().String()
	apiKey := config.APIKey{
		AccessToken: cfg.AccessToken,
		Timestamp:   strconv.FormatInt(time.Now().Unix(), 10),
		Nonce:       nonce[0:8] + nonce[9:13] + nonce[14:18],
	}
	ctx = context.WithValue(ctx, config.ContextAPIKey, apiKey)
	client := api.NewAPIClient(&cfg.Configuration)
	sdkClient := &SDKClient{
		Config:       cfg,
		Ctx:          &ctx,
		Client:       client,
		RoundTripper: http.DefaultTransport,
		Version:      version,
		ApiVersion:   apiVersion,
	}
	sdkClient.Client.Cfg.HTTPClient.Transport = sdkClient
	sdkClient.UseSandbox()
	sdkClient.middlewareList = []Middleware{
		&AuthMiddleware{sdkClient},
		&HeaderMiddleware{sdkClient},
		&DiffHostMiddleware{sdkClient},
		&LogMiddleware{sdkClient},
	}
	return sdkClient
}

func (tads *SDKClient) SetHost(host string, schema string) *SDKClient {
	modified := false
	if host != "" {
		tads.Client.Cfg.Host = host
		modified = true
	}
	if schema != "" {
		tads.Client.Cfg.Scheme = schema
		modified = true
	}
	if modified {
		tads.Client.Cfg.BasePath = tads.Client.Cfg.Scheme + "://" + tads.Client.Cfg.Host + "/" + tads.ApiVersion
	}
	return tads
}

func (tads *SDKClient) UseSandbox() *SDKClient {
	return tads.SetHost("sandbox-api.e.qq.com", "https")
}

func (tads *SDKClient) UseProduction() *SDKClient {
	return tads.SetHost("api.e.qq.com", "https")
}

func (tads *SDKClient) GetAccessToken() string {
	return tads.Config.AccessToken
}

func (tads *SDKClient) SetAccessToken(accessToken string) *SDKClient {
	tads.Config.AccessToken = accessToken
	return tads
}

func (tads *SDKClient) GetVersion() string {
	return tads.Version
}

func (tads *SDKClient) IsDebug() bool {
	return tads.Config.IsDebug
}

func (tads *SDKClient) SetDebug(debug bool) *SDKClient {
	tads.Config.IsDebug = debug
	return tads
}

func (tads *SDKClient) GetDebugFile() string {
	return tads.Config.DebugFile
}

func (tads *SDKClient) SetDebugFile(debugFile string) *SDKClient {
	tads.Config.DebugFile = debugFile
	return tads
}

func (tads *SDKClient) IsMonitor() bool {
	return !tads.Config.SkipMonitor
}

func (tads *SDKClient) SetHeaders(header http.Header) *SDKClient {
	tads.Config.GlobalConfig.HttpOption.Header = header
	return tads
}

func (tads *SDKClient) SetHeader(key string, value string) *SDKClient {
	if tads.Config.GlobalConfig.HttpOption.Header == nil {
		tads.Config.GlobalConfig.HttpOption.Header = http.Header{}
	}
	tads.Config.GlobalConfig.HttpOption.Header.Set(key, value)
	return tads
}

func (tads *SDKClient) SetMonitor(monitor bool) *SDKClient {
	tads.Config.SkipMonitor = !monitor
	return tads
}

func (tads *SDKClient) AppendMiddleware(middleware Middleware) *SDKClient {
	tads.middlewareList = append(tads.middlewareList, middleware)
	return tads
}

func (tads *SDKClient) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	beforeFunc := func(req *http.Request) (rsp *http.Response, err error) {
		return tads.RoundTripper.RoundTrip(req)
	}
	middlewareCount := len(tads.middlewareList)
	// 逆序遍历
	for i := middlewareCount - 1; i >= 0; i-- {
		beforeFunc = tads.GenMiddlewareHandleFunc(tads.middlewareList[i], beforeFunc)
	}
	return beforeFunc(req)
}

func (tads *SDKClient) GenMiddlewareHandleFunc(middleware Middleware, beforeFunc func(req *http.Request) (rsp *http.Response, err error)) func(req *http.Request) (rsp *http.Response, err error) {
	return func(req *http.Request) (rsp *http.Response, err error) {
		return middleware.Handle(req, beforeFunc)
	}
}
