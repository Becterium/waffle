package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"time"

	mediav1 "waffle/api/media/service/v1"
	userv1 "waffle/api/user/service/v1"
	"waffle/app/waffle/interface/internal/conf"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"

	consulAPI "github.com/hashicorp/consul/api"
	grpcx "google.golang.org/grpc"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewMediaRepo, NewDiscovery, NewUserServiceClient, NewMediaServiceClient, NewRegistrar)

// Data .
type Data struct {
	log *log.Helper
	uc  userv1.UserClient
	mc  mediav1.MediaClient
}

func NewData(uc userv1.UserClient, mc mediav1.MediaClient, logger log.Logger) (*Data, error) {
	err := initTracer("192.168.37.134:4318")
	if err != nil {
		panic(err)
	}
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{
		log: l,
		uc:  uc,
		mc:  mc,
	}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewUserServiceClient(r registry.Discovery) userv1.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///waffle.user.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			metadata.Client(),
		),
		grpc.WithTimeout(3*time.Second),
		// 设置空闲连接超时时间
		grpc.WithOptions(grpcx.WithIdleTimeout(0)),
	)
	if err != nil {
		panic(err)
	}
	c := userv1.NewUserClient(conn)
	return c
}

func NewMediaServiceClient(r registry.Discovery) mediav1.MediaClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///waffle.media.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			tracing.Client(),
			metadata.Client(),
		),
		grpc.WithTimeout(3*time.Second),
		// 设置空闲连接超时时间
		grpc.WithOptions(grpcx.WithIdleTimeout(0)),
	)
	if err != nil {
		panic(err)
	}
	c := mediav1.NewMediaClient(conn)
	return c
}

func initTracer(endpoint string) error {
	// 创建 exporter
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpoint(endpoint),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return err
	}
	tp := tracesdk.NewTracerProvider(
		// 将基于父span的采样率设置为100%
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(1.0))),
		// 始终确保在生产中批量处理
		tracesdk.WithBatcher(exporter),
		// 在资源中记录有关此应用程序的信息
		tracesdk.WithResource(resource.NewSchemaless(
			semconv.ServiceNameKey.String("waffle-trace"),
			attribute.String("exporter", "otlp"),
			attribute.Float64("float", 312.23),
		)),
	)
	otel.SetTracerProvider(tp)
	return nil
}
