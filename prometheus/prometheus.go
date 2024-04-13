package prometheus

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

type Metrics struct {
	ApiReqCnt     *prometheus.CounterVec
	ApiReqSize    *prometheus.SummaryVec
	ApiRespSize   *prometheus.SummaryVec
	ApiReqDurHist *prometheus.HistogramVec
}

func New(cfg Config) Metrics {
	var m Metrics

	m.ApiReqCnt = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      "requests_total",
		Help:      "The API request processed count",
	}, []string{"method", "code"})

	m.ApiReqDurHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      "request_processing_time_histogram_ms",
		Help:      "The API request duration (ms)",
		Buckets:   []float64{5, 10, 25, 50, 100, 250, 500, 1000},
	}, []string{"method", "code"})

	m.ApiReqSize = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      "request_size_bytes",
		Help:      "The HTTP request sizes in bytes",
	}, []string{"method", "code"})

	m.ApiRespSize = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: cfg.Namespace,
		Subsystem: cfg.Subsystem,
		Name:      "response_size_bytes",
		Help:      "The HTTP response sizes in bytes",
	}, []string{"method", "code"})

	prometheus.MustRegister(
		m.ApiReqCnt,
		m.ApiReqDurHist,
		m.ApiReqSize,
		m.ApiRespSize,
	)

	return m
}

func (m Metrics) UnaryServerInterceptor(ctx context.Context) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		startTime := time.Now()

		resp, err := handler(ctx, req)

		elapsed := time.Since(startTime).Seconds() * 1000

		st, _ := status.FromError(err)
		statusCode := st.Code().String()

		promLabels := prometheus.Labels{"method": info.FullMethod, "code": statusCode}

		m.ApiReqCnt.With(promLabels).Inc()
		m.ApiReqDurHist.With(promLabels).Observe(elapsed)
		m.ApiReqSize.With(promLabels).Observe(elapsed)
		m.ApiRespSize.With(promLabels).Observe(elapsed)

		return resp, err
	}
}
