// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap/confmaptest"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(loadConfig(t, "default"), settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	enabledMetrics := make(map[string]bool)

	enabledMetrics["redis.clients.blocked"] = true
	mb.RecordRedisClientsBlockedDataPoint(ts, 1)

	enabledMetrics["redis.clients.connected"] = true
	mb.RecordRedisClientsConnectedDataPoint(ts, 1)

	enabledMetrics["redis.clients.max_input_buffer"] = true
	mb.RecordRedisClientsMaxInputBufferDataPoint(ts, 1)

	enabledMetrics["redis.clients.max_output_buffer"] = true
	mb.RecordRedisClientsMaxOutputBufferDataPoint(ts, 1)

	mb.RecordRedisCmdCallsDataPoint(ts, 1, "attr-val")

	mb.RecordRedisCmdUsecDataPoint(ts, 1, "attr-val")

	enabledMetrics["redis.commands"] = true
	mb.RecordRedisCommandsDataPoint(ts, 1)

	enabledMetrics["redis.commands.processed"] = true
	mb.RecordRedisCommandsProcessedDataPoint(ts, 1)

	enabledMetrics["redis.connections.received"] = true
	mb.RecordRedisConnectionsReceivedDataPoint(ts, 1)

	enabledMetrics["redis.connections.rejected"] = true
	mb.RecordRedisConnectionsRejectedDataPoint(ts, 1)

	enabledMetrics["redis.cpu.time"] = true
	mb.RecordRedisCPUTimeDataPoint(ts, 1, AttributeState(1))

	enabledMetrics["redis.db.avg_ttl"] = true
	mb.RecordRedisDbAvgTTLDataPoint(ts, 1, "attr-val")

	enabledMetrics["redis.db.expires"] = true
	mb.RecordRedisDbExpiresDataPoint(ts, 1, "attr-val")

	enabledMetrics["redis.db.keys"] = true
	mb.RecordRedisDbKeysDataPoint(ts, 1, "attr-val")

	enabledMetrics["redis.keys.evicted"] = true
	mb.RecordRedisKeysEvictedDataPoint(ts, 1)

	enabledMetrics["redis.keys.expired"] = true
	mb.RecordRedisKeysExpiredDataPoint(ts, 1)

	enabledMetrics["redis.keyspace.hits"] = true
	mb.RecordRedisKeyspaceHitsDataPoint(ts, 1)

	enabledMetrics["redis.keyspace.misses"] = true
	mb.RecordRedisKeyspaceMissesDataPoint(ts, 1)

	enabledMetrics["redis.latest_fork"] = true
	mb.RecordRedisLatestForkDataPoint(ts, 1)

	mb.RecordRedisMaxmemoryDataPoint(ts, 1)

	enabledMetrics["redis.memory.fragmentation_ratio"] = true
	mb.RecordRedisMemoryFragmentationRatioDataPoint(ts, 1)

	enabledMetrics["redis.memory.lua"] = true
	mb.RecordRedisMemoryLuaDataPoint(ts, 1)

	enabledMetrics["redis.memory.peak"] = true
	mb.RecordRedisMemoryPeakDataPoint(ts, 1)

	enabledMetrics["redis.memory.rss"] = true
	mb.RecordRedisMemoryRssDataPoint(ts, 1)

	enabledMetrics["redis.memory.used"] = true
	mb.RecordRedisMemoryUsedDataPoint(ts, 1)

	enabledMetrics["redis.net.input"] = true
	mb.RecordRedisNetInputDataPoint(ts, 1)

	enabledMetrics["redis.net.output"] = true
	mb.RecordRedisNetOutputDataPoint(ts, 1)

	enabledMetrics["redis.rdb.changes_since_last_save"] = true
	mb.RecordRedisRdbChangesSinceLastSaveDataPoint(ts, 1)

	enabledMetrics["redis.replication.backlog_first_byte_offset"] = true
	mb.RecordRedisReplicationBacklogFirstByteOffsetDataPoint(ts, 1)

	enabledMetrics["redis.replication.offset"] = true
	mb.RecordRedisReplicationOffsetDataPoint(ts, 1)

	mb.RecordRedisRoleDataPoint(ts, 1, AttributeRole(1))

	enabledMetrics["redis.slaves.connected"] = true
	mb.RecordRedisSlavesConnectedDataPoint(ts, 1)

	enabledMetrics["redis.uptime"] = true
	mb.RecordRedisUptimeDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(loadConfig(t, "all_metrics"), settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordRedisClientsBlockedDataPoint(ts, 1)
	mb.RecordRedisClientsConnectedDataPoint(ts, 1)
	mb.RecordRedisClientsMaxInputBufferDataPoint(ts, 1)
	mb.RecordRedisClientsMaxOutputBufferDataPoint(ts, 1)
	mb.RecordRedisCmdCallsDataPoint(ts, 1, "attr-val")
	mb.RecordRedisCmdUsecDataPoint(ts, 1, "attr-val")
	mb.RecordRedisCommandsDataPoint(ts, 1)
	mb.RecordRedisCommandsProcessedDataPoint(ts, 1)
	mb.RecordRedisConnectionsReceivedDataPoint(ts, 1)
	mb.RecordRedisConnectionsRejectedDataPoint(ts, 1)
	mb.RecordRedisCPUTimeDataPoint(ts, 1, AttributeState(1))
	mb.RecordRedisDbAvgTTLDataPoint(ts, 1, "attr-val")
	mb.RecordRedisDbExpiresDataPoint(ts, 1, "attr-val")
	mb.RecordRedisDbKeysDataPoint(ts, 1, "attr-val")
	mb.RecordRedisKeysEvictedDataPoint(ts, 1)
	mb.RecordRedisKeysExpiredDataPoint(ts, 1)
	mb.RecordRedisKeyspaceHitsDataPoint(ts, 1)
	mb.RecordRedisKeyspaceMissesDataPoint(ts, 1)
	mb.RecordRedisLatestForkDataPoint(ts, 1)
	mb.RecordRedisMaxmemoryDataPoint(ts, 1)
	mb.RecordRedisMemoryFragmentationRatioDataPoint(ts, 1)
	mb.RecordRedisMemoryLuaDataPoint(ts, 1)
	mb.RecordRedisMemoryPeakDataPoint(ts, 1)
	mb.RecordRedisMemoryRssDataPoint(ts, 1)
	mb.RecordRedisMemoryUsedDataPoint(ts, 1)
	mb.RecordRedisNetInputDataPoint(ts, 1)
	mb.RecordRedisNetOutputDataPoint(ts, 1)
	mb.RecordRedisRdbChangesSinceLastSaveDataPoint(ts, 1)
	mb.RecordRedisReplicationBacklogFirstByteOffsetDataPoint(ts, 1)
	mb.RecordRedisReplicationOffsetDataPoint(ts, 1)
	mb.RecordRedisRoleDataPoint(ts, 1, AttributeRole(1))
	mb.RecordRedisSlavesConnectedDataPoint(ts, 1)
	mb.RecordRedisUptimeDataPoint(ts, 1)

	metrics := mb.Emit(WithRedisVersion("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("redis.version")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "redis.clients.blocked":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of clients pending on a blocking call", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.clients.blocked"] = struct{}{}
		case "redis.clients.connected":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of client connections (excluding connections from replicas)", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.clients.connected"] = struct{}{}
		case "redis.clients.max_input_buffer":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Biggest input buffer among current client connections", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.clients.max_input_buffer"] = struct{}{}
		case "redis.clients.max_output_buffer":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Longest output list among current client connections", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.clients.max_output_buffer"] = struct{}{}
		case "redis.cmd.calls":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of calls for a command", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("cmd")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["redis.cmd.calls"] = struct{}{}
		case "redis.cmd.usec":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total time for all executions of this command", ms.At(i).Description())
			assert.Equal(t, "us", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("cmd")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["redis.cmd.usec"] = struct{}{}
		case "redis.commands":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of commands processed per second", ms.At(i).Description())
			assert.Equal(t, "{ops}/s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.commands"] = struct{}{}
		case "redis.commands.processed":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of commands processed by the server", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.commands.processed"] = struct{}{}
		case "redis.connections.received":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of connections accepted by the server", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.connections.received"] = struct{}{}
		case "redis.connections.rejected":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of connections rejected because of maxclients limit", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.connections.rejected"] = struct{}{}
		case "redis.cpu.time":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "System CPU consumed by the Redis server in seconds since server start", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			attrVal, ok := dp.Attributes().Get("state")
			assert.True(t, ok)
			assert.Equal(t, "sys", attrVal.Str())
			validatedMetrics["redis.cpu.time"] = struct{}{}
		case "redis.db.avg_ttl":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Average keyspace keys TTL", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("db")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["redis.db.avg_ttl"] = struct{}{}
		case "redis.db.expires":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of keyspace keys with an expiration", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("db")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["redis.db.expires"] = struct{}{}
		case "redis.db.keys":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of keyspace keys", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("db")
			assert.True(t, ok)
			assert.EqualValues(t, "attr-val", attrVal.Str())
			validatedMetrics["redis.db.keys"] = struct{}{}
		case "redis.keys.evicted":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of evicted keys due to maxmemory limit", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.keys.evicted"] = struct{}{}
		case "redis.keys.expired":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of key expiration events", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.keys.expired"] = struct{}{}
		case "redis.keyspace.hits":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of successful lookup of keys in the main dictionary", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.keyspace.hits"] = struct{}{}
		case "redis.keyspace.misses":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of failed lookup of keys in the main dictionary", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.keyspace.misses"] = struct{}{}
		case "redis.latest_fork":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Duration of the latest fork operation in microseconds", ms.At(i).Description())
			assert.Equal(t, "us", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.latest_fork"] = struct{}{}
		case "redis.maxmemory":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The value of the maxmemory configuration directive", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.maxmemory"] = struct{}{}
		case "redis.memory.fragmentation_ratio":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Ratio between used_memory_rss and used_memory", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
			assert.Equal(t, float64(1), dp.DoubleValue())
			validatedMetrics["redis.memory.fragmentation_ratio"] = struct{}{}
		case "redis.memory.lua":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of bytes used by the Lua engine", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.memory.lua"] = struct{}{}
		case "redis.memory.peak":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Peak memory consumed by Redis (in bytes)", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.memory.peak"] = struct{}{}
		case "redis.memory.rss":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Number of bytes that Redis allocated as seen by the operating system", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.memory.rss"] = struct{}{}
		case "redis.memory.used":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Total number of bytes allocated by Redis using its allocator", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.memory.used"] = struct{}{}
		case "redis.net.input":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of bytes read from the network", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.net.input"] = struct{}{}
		case "redis.net.output":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "The total number of bytes written to the network", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.net.output"] = struct{}{}
		case "redis.rdb.changes_since_last_save":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of changes since the last dump", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.rdb.changes_since_last_save"] = struct{}{}
		case "redis.replication.backlog_first_byte_offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The master offset of the replication backlog buffer", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.replication.backlog_first_byte_offset"] = struct{}{}
		case "redis.replication.offset":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The server's current replication offset", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.replication.offset"] = struct{}{}
		case "redis.role":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Redis node's role", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("role")
			assert.True(t, ok)
			assert.Equal(t, "replica", attrVal.Str())
			validatedMetrics["redis.role"] = struct{}{}
		case "redis.slaves.connected":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of connected replicas", ms.At(i).Description())
			assert.Equal(t, "", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.slaves.connected"] = struct{}{}
		case "redis.uptime":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of seconds since Redis server start", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["redis.uptime"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	observedZapCore, observedLogs := observer.New(zap.WarnLevel)
	settings := receivertest.NewNopCreateSettings()
	settings.Logger = zap.New(observedZapCore)
	mb := NewMetricsBuilder(loadConfig(t, "no_metrics"), settings, WithStartTime(start))

	assert.Equal(t, 0, observedLogs.Len())

	mb.RecordRedisClientsBlockedDataPoint(ts, 1)
	mb.RecordRedisClientsConnectedDataPoint(ts, 1)
	mb.RecordRedisClientsMaxInputBufferDataPoint(ts, 1)
	mb.RecordRedisClientsMaxOutputBufferDataPoint(ts, 1)
	mb.RecordRedisCmdCallsDataPoint(ts, 1, "attr-val")
	mb.RecordRedisCmdUsecDataPoint(ts, 1, "attr-val")
	mb.RecordRedisCommandsDataPoint(ts, 1)
	mb.RecordRedisCommandsProcessedDataPoint(ts, 1)
	mb.RecordRedisConnectionsReceivedDataPoint(ts, 1)
	mb.RecordRedisConnectionsRejectedDataPoint(ts, 1)
	mb.RecordRedisCPUTimeDataPoint(ts, 1, AttributeState(1))
	mb.RecordRedisDbAvgTTLDataPoint(ts, 1, "attr-val")
	mb.RecordRedisDbExpiresDataPoint(ts, 1, "attr-val")
	mb.RecordRedisDbKeysDataPoint(ts, 1, "attr-val")
	mb.RecordRedisKeysEvictedDataPoint(ts, 1)
	mb.RecordRedisKeysExpiredDataPoint(ts, 1)
	mb.RecordRedisKeyspaceHitsDataPoint(ts, 1)
	mb.RecordRedisKeyspaceMissesDataPoint(ts, 1)
	mb.RecordRedisLatestForkDataPoint(ts, 1)
	mb.RecordRedisMaxmemoryDataPoint(ts, 1)
	mb.RecordRedisMemoryFragmentationRatioDataPoint(ts, 1)
	mb.RecordRedisMemoryLuaDataPoint(ts, 1)
	mb.RecordRedisMemoryPeakDataPoint(ts, 1)
	mb.RecordRedisMemoryRssDataPoint(ts, 1)
	mb.RecordRedisMemoryUsedDataPoint(ts, 1)
	mb.RecordRedisNetInputDataPoint(ts, 1)
	mb.RecordRedisNetOutputDataPoint(ts, 1)
	mb.RecordRedisRdbChangesSinceLastSaveDataPoint(ts, 1)
	mb.RecordRedisReplicationBacklogFirstByteOffsetDataPoint(ts, 1)
	mb.RecordRedisReplicationOffsetDataPoint(ts, 1)
	mb.RecordRedisRoleDataPoint(ts, 1, AttributeRole(1))
	mb.RecordRedisSlavesConnectedDataPoint(ts, 1)
	mb.RecordRedisUptimeDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}

func loadConfig(t *testing.T, name string) MetricsSettings {
	cm, err := confmaptest.LoadConf(filepath.Join("testdata", "config.yaml"))
	require.NoError(t, err)
	sub, err := cm.Sub(name)
	require.NoError(t, err)
	cfg := DefaultMetricsSettings()
	require.NoError(t, component.UnmarshalConfig(sub, &cfg))
	return cfg
}
