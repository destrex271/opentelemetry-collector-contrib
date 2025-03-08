// Code generated by mdatagen. DO NOT EDIT.

package googlemanagedprometheusexporter

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/exporter/exportertest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

var typ = component.MustNewType("googlemanagedprometheus")

func TestComponentFactoryType(t *testing.T) {
	require.Equal(t, typ, NewFactory().Type())
}

func TestComponentConfigStruct(t *testing.T) {
	require.NoError(t, componenttest.CheckConfigStruct(NewFactory().CreateDefaultConfig()))
}

func TestComponentLifecycle(t *testing.T) {
	factory := NewFactory()

	tests := []struct {
		createFn func(ctx context.Context, set exporter.Settings, cfg component.Config) (component.Component, error)
		name     string
	}{

		{
			name: "metrics",
			createFn: func(ctx context.Context, set exporter.Settings, cfg component.Config) (component.Component, error) {
				return factory.CreateMetrics(ctx, set, cfg)
			},
		},
	}

	cm, err := confmaptest.LoadConf("metadata.yaml")
	require.NoError(t, err)
	cfg := factory.CreateDefaultConfig()
	sub, err := cm.Sub("tests::config")
	require.NoError(t, err)
	require.NoError(t, sub.Unmarshal(&cfg))

	for _, tt := range tests {
		t.Run(tt.name+"-shutdown", func(t *testing.T) {
			c, err := tt.createFn(context.Background(), exportertest.NewNopSettings(typ), cfg)
			require.NoError(t, err)
			err = c.Shutdown(context.Background())
			require.NoError(t, err)
		})
	}
}

func generateLifecycleTestLogs() plog.Logs {
	logs := plog.NewLogs()
	rl := logs.ResourceLogs().AppendEmpty()
	rl.Resource().Attributes().PutStr("resource", "R1")
	l := rl.ScopeLogs().AppendEmpty().LogRecords().AppendEmpty()
	l.Body().SetStr("test log message")
	l.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	return logs
}

func generateLifecycleTestMetrics() pmetric.Metrics {
	metrics := pmetric.NewMetrics()
	rm := metrics.ResourceMetrics().AppendEmpty()
	rm.Resource().Attributes().PutStr("resource", "R1")
	m := rm.ScopeMetrics().AppendEmpty().Metrics().AppendEmpty()
	m.SetName("test_metric")
	dp := m.SetEmptyGauge().DataPoints().AppendEmpty()
	dp.Attributes().PutStr("test_attr", "value_1")
	dp.SetIntValue(123)
	dp.SetTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	return metrics
}

func generateLifecycleTestTraces() ptrace.Traces {
	traces := ptrace.NewTraces()
	rs := traces.ResourceSpans().AppendEmpty()
	rs.Resource().Attributes().PutStr("resource", "R1")
	span := rs.ScopeSpans().AppendEmpty().Spans().AppendEmpty()
	span.Attributes().PutStr("test_attr", "value_1")
	span.SetName("test_span")
	span.SetStartTimestamp(pcommon.NewTimestampFromTime(time.Now().Add(-1 * time.Second)))
	span.SetEndTimestamp(pcommon.NewTimestampFromTime(time.Now()))
	return traces
}
