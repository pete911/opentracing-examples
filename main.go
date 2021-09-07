package main

import (
	"github.com/opentracing/opentracing-go"
	"github.com/pete911/opentracing-examples/internal/handler"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log/zap"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
)

func main() {

	logger, err := NewZapConfig(zapcore.DebugLevel).Build()
	if err != nil {
		log.Fatalf("zap build: %v", err)
	}

	jaegerLogger := jaegercfg.Logger(jaegerlog.NewLogger(logger))
	jaegerMetrics := jaegercfg.Metrics(metrics.NullFactory)
	tracer, closer, err := NewJaegerConfig("opentracing-examples").NewTracer(jaegerLogger, jaegerMetrics)
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	http.HandleFunc("/", handler.GetDashboard)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func NewZapConfig(level zapcore.Level) zap.Config {

	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.Level.SetLevel(level)
	return config
}

func NewJaegerConfig(serviceName string) jaegercfg.Configuration {

	return jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}
}
