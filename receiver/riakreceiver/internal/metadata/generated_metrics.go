// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricSettings) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsSettings provides settings for riakreceiver metrics.
type MetricsSettings struct {
	RiakMemoryLimit              MetricSettings `mapstructure:"riak.memory.limit"`
	RiakNodeOperationCount       MetricSettings `mapstructure:"riak.node.operation.count"`
	RiakNodeOperationTimeMean    MetricSettings `mapstructure:"riak.node.operation.time.mean"`
	RiakNodeReadRepairCount      MetricSettings `mapstructure:"riak.node.read_repair.count"`
	RiakVnodeIndexOperationCount MetricSettings `mapstructure:"riak.vnode.index.operation.count"`
	RiakVnodeOperationCount      MetricSettings `mapstructure:"riak.vnode.operation.count"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		RiakMemoryLimit: MetricSettings{
			Enabled: true,
		},
		RiakNodeOperationCount: MetricSettings{
			Enabled: true,
		},
		RiakNodeOperationTimeMean: MetricSettings{
			Enabled: true,
		},
		RiakNodeReadRepairCount: MetricSettings{
			Enabled: true,
		},
		RiakVnodeIndexOperationCount: MetricSettings{
			Enabled: true,
		},
		RiakVnodeOperationCount: MetricSettings{
			Enabled: true,
		},
	}
}

// AttributeOperation specifies the a value operation attribute.
type AttributeOperation int

const (
	_ AttributeOperation = iota
	AttributeOperationRead
	AttributeOperationWrite
	AttributeOperationDelete
)

// String returns the string representation of the AttributeOperation.
func (av AttributeOperation) String() string {
	switch av {
	case AttributeOperationRead:
		return "read"
	case AttributeOperationWrite:
		return "write"
	case AttributeOperationDelete:
		return "delete"
	}
	return ""
}

// MapAttributeOperation is a helper map of string to AttributeOperation attribute value.
var MapAttributeOperation = map[string]AttributeOperation{
	"read":   AttributeOperationRead,
	"write":  AttributeOperationWrite,
	"delete": AttributeOperationDelete,
}

// AttributeRequest specifies the a value request attribute.
type AttributeRequest int

const (
	_ AttributeRequest = iota
	AttributeRequestPut
	AttributeRequestGet
)

// String returns the string representation of the AttributeRequest.
func (av AttributeRequest) String() string {
	switch av {
	case AttributeRequestPut:
		return "put"
	case AttributeRequestGet:
		return "get"
	}
	return ""
}

// MapAttributeRequest is a helper map of string to AttributeRequest attribute value.
var MapAttributeRequest = map[string]AttributeRequest{
	"put": AttributeRequestPut,
	"get": AttributeRequestGet,
}

type metricRiakMemoryLimit struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.memory.limit metric with initial data.
func (m *metricRiakMemoryLimit) init() {
	m.data.SetName("riak.memory.limit")
	m.data.SetDescription("The amount of memory allocated to the node.")
	m.data.SetUnit("By")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRiakMemoryLimit) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakMemoryLimit) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakMemoryLimit) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakMemoryLimit(settings MetricSettings) metricRiakMemoryLimit {
	m := metricRiakMemoryLimit{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRiakNodeOperationCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.node.operation.count metric with initial data.
func (m *metricRiakNodeOperationCount) init() {
	m.data.SetName("riak.node.operation.count")
	m.data.SetDescription("The number of operations performed by the node.")
	m.data.SetUnit("{operation}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakNodeOperationCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("request", requestAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakNodeOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakNodeOperationCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakNodeOperationCount(settings MetricSettings) metricRiakNodeOperationCount {
	m := metricRiakNodeOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRiakNodeOperationTimeMean struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.node.operation.time.mean metric with initial data.
func (m *metricRiakNodeOperationTimeMean) init() {
	m.data.SetName("riak.node.operation.time.mean")
	m.data.SetDescription("The mean time between request and response for operations performed by the node over the last minute.")
	m.data.SetUnit("us")
	m.data.SetEmptyGauge()
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakNodeOperationTimeMean) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("request", requestAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakNodeOperationTimeMean) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakNodeOperationTimeMean) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakNodeOperationTimeMean(settings MetricSettings) metricRiakNodeOperationTimeMean {
	m := metricRiakNodeOperationTimeMean{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRiakNodeReadRepairCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.node.read_repair.count metric with initial data.
func (m *metricRiakNodeReadRepairCount) init() {
	m.data.SetName("riak.node.read_repair.count")
	m.data.SetDescription("The number of read repairs performed by the node.")
	m.data.SetUnit("{read_repair}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
}

func (m *metricRiakNodeReadRepairCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakNodeReadRepairCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakNodeReadRepairCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakNodeReadRepairCount(settings MetricSettings) metricRiakNodeReadRepairCount {
	m := metricRiakNodeReadRepairCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRiakVnodeIndexOperationCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.vnode.index.operation.count metric with initial data.
func (m *metricRiakVnodeIndexOperationCount) init() {
	m.data.SetName("riak.vnode.index.operation.count")
	m.data.SetDescription("The number of index operations performed by vnodes on the node.")
	m.data.SetUnit("{operation}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakVnodeIndexOperationCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, operationAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("operation", operationAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakVnodeIndexOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakVnodeIndexOperationCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakVnodeIndexOperationCount(settings MetricSettings) metricRiakVnodeIndexOperationCount {
	m := metricRiakVnodeIndexOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

type metricRiakVnodeOperationCount struct {
	data     pmetric.Metric // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills riak.vnode.operation.count metric with initial data.
func (m *metricRiakVnodeOperationCount) init() {
	m.data.SetName("riak.vnode.operation.count")
	m.data.SetDescription("The number of operations performed by vnodes on the node.")
	m.data.SetUnit("{operation}")
	m.data.SetEmptySum()
	m.data.Sum().SetIsMonotonic(true)
	m.data.Sum().SetAggregationTemporality(pmetric.AggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricRiakVnodeOperationCount) recordDataPoint(start pcommon.Timestamp, ts pcommon.Timestamp, val int64, requestAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntValue(val)
	dp.Attributes().PutStr("request", requestAttributeValue)
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricRiakVnodeOperationCount) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricRiakVnodeOperationCount) emit(metrics pmetric.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricRiakVnodeOperationCount(settings MetricSettings) metricRiakVnodeOperationCount {
	m := metricRiakVnodeOperationCount{settings: settings}
	if settings.Enabled {
		m.data = pmetric.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                          pcommon.Timestamp   // start time that will be applied to all recorded data points.
	metricsCapacity                    int                 // maximum observed number of metrics per resource.
	resourceCapacity                   int                 // maximum observed number of resource attributes.
	metricsBuffer                      pmetric.Metrics     // accumulates metrics data before emitting.
	buildInfo                          component.BuildInfo // contains version information
	metricRiakMemoryLimit              metricRiakMemoryLimit
	metricRiakNodeOperationCount       metricRiakNodeOperationCount
	metricRiakNodeOperationTimeMean    metricRiakNodeOperationTimeMean
	metricRiakNodeReadRepairCount      metricRiakNodeReadRepairCount
	metricRiakVnodeIndexOperationCount metricRiakVnodeIndexOperationCount
	metricRiakVnodeOperationCount      metricRiakVnodeOperationCount
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pcommon.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(ms MetricsSettings, settings receiver.CreateSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                          pcommon.NewTimestampFromTime(time.Now()),
		metricsBuffer:                      pmetric.NewMetrics(),
		buildInfo:                          settings.BuildInfo,
		metricRiakMemoryLimit:              newMetricRiakMemoryLimit(ms.RiakMemoryLimit),
		metricRiakNodeOperationCount:       newMetricRiakNodeOperationCount(ms.RiakNodeOperationCount),
		metricRiakNodeOperationTimeMean:    newMetricRiakNodeOperationTimeMean(ms.RiakNodeOperationTimeMean),
		metricRiakNodeReadRepairCount:      newMetricRiakNodeReadRepairCount(ms.RiakNodeReadRepairCount),
		metricRiakVnodeIndexOperationCount: newMetricRiakVnodeIndexOperationCount(ms.RiakVnodeIndexOperationCount),
		metricRiakVnodeOperationCount:      newMetricRiakVnodeOperationCount(ms.RiakVnodeOperationCount),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// updateCapacity updates max length of metrics and resource attributes that will be used for the slice capacity.
func (mb *MetricsBuilder) updateCapacity(rm pmetric.ResourceMetrics) {
	if mb.metricsCapacity < rm.ScopeMetrics().At(0).Metrics().Len() {
		mb.metricsCapacity = rm.ScopeMetrics().At(0).Metrics().Len()
	}
	if mb.resourceCapacity < rm.Resource().Attributes().Len() {
		mb.resourceCapacity = rm.Resource().Attributes().Len()
	}
}

// ResourceMetricsOption applies changes to provided resource metrics.
type ResourceMetricsOption func(pmetric.ResourceMetrics)

// WithRiakNodeName sets provided value as "riak.node.name" attribute for current resource.
func WithRiakNodeName(val string) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		rm.Resource().Attributes().PutStr("riak.node.name", val)
	}
}

// WithStartTimeOverride overrides start time for all the resource metrics data points.
// This option should be only used if different start time has to be set on metrics coming from different resources.
func WithStartTimeOverride(start pcommon.Timestamp) ResourceMetricsOption {
	return func(rm pmetric.ResourceMetrics) {
		var dps pmetric.NumberDataPointSlice
		metrics := rm.ScopeMetrics().At(0).Metrics()
		for i := 0; i < metrics.Len(); i++ {
			switch metrics.At(i).Type() {
			case pmetric.MetricTypeGauge:
				dps = metrics.At(i).Gauge().DataPoints()
			case pmetric.MetricTypeSum:
				dps = metrics.At(i).Sum().DataPoints()
			}
			for j := 0; j < dps.Len(); j++ {
				dps.At(j).SetStartTimestamp(start)
			}
		}
	}
}

// EmitForResource saves all the generated metrics under a new resource and updates the internal state to be ready for
// recording another set of data points as part of another resource. This function can be helpful when one scraper
// needs to emit metrics from several resources. Otherwise calling this function is not required,
// just `Emit` function can be called instead.
// Resource attributes should be provided as ResourceMetricsOption arguments.
func (mb *MetricsBuilder) EmitForResource(rmo ...ResourceMetricsOption) {
	rm := pmetric.NewResourceMetrics()
	rm.Resource().Attributes().EnsureCapacity(mb.resourceCapacity)
	ils := rm.ScopeMetrics().AppendEmpty()
	ils.Scope().SetName("otelcol/riakreceiver")
	ils.Scope().SetVersion(mb.buildInfo.Version)
	ils.Metrics().EnsureCapacity(mb.metricsCapacity)
	mb.metricRiakMemoryLimit.emit(ils.Metrics())
	mb.metricRiakNodeOperationCount.emit(ils.Metrics())
	mb.metricRiakNodeOperationTimeMean.emit(ils.Metrics())
	mb.metricRiakNodeReadRepairCount.emit(ils.Metrics())
	mb.metricRiakVnodeIndexOperationCount.emit(ils.Metrics())
	mb.metricRiakVnodeOperationCount.emit(ils.Metrics())
	for _, op := range rmo {
		op(rm)
	}
	if ils.Metrics().Len() > 0 {
		mb.updateCapacity(rm)
		rm.MoveTo(mb.metricsBuffer.ResourceMetrics().AppendEmpty())
	}
}

// Emit returns all the metrics accumulated by the metrics builder and updates the internal state to be ready for
// recording another set of metrics. This function will be responsible for applying all the transformations required to
// produce metric representation defined in metadata and user settings, e.g. delta or cumulative.
func (mb *MetricsBuilder) Emit(rmo ...ResourceMetricsOption) pmetric.Metrics {
	mb.EmitForResource(rmo...)
	metrics := pmetric.NewMetrics()
	mb.metricsBuffer.MoveTo(metrics)
	return metrics
}

// RecordRiakMemoryLimitDataPoint adds a data point to riak.memory.limit metric.
func (mb *MetricsBuilder) RecordRiakMemoryLimitDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRiakMemoryLimit.recordDataPoint(mb.startTime, ts, val)
}

// RecordRiakNodeOperationCountDataPoint adds a data point to riak.node.operation.count metric.
func (mb *MetricsBuilder) RecordRiakNodeOperationCountDataPoint(ts pcommon.Timestamp, val int64, requestAttributeValue AttributeRequest) {
	mb.metricRiakNodeOperationCount.recordDataPoint(mb.startTime, ts, val, requestAttributeValue.String())
}

// RecordRiakNodeOperationTimeMeanDataPoint adds a data point to riak.node.operation.time.mean metric.
func (mb *MetricsBuilder) RecordRiakNodeOperationTimeMeanDataPoint(ts pcommon.Timestamp, val int64, requestAttributeValue AttributeRequest) {
	mb.metricRiakNodeOperationTimeMean.recordDataPoint(mb.startTime, ts, val, requestAttributeValue.String())
}

// RecordRiakNodeReadRepairCountDataPoint adds a data point to riak.node.read_repair.count metric.
func (mb *MetricsBuilder) RecordRiakNodeReadRepairCountDataPoint(ts pcommon.Timestamp, val int64) {
	mb.metricRiakNodeReadRepairCount.recordDataPoint(mb.startTime, ts, val)
}

// RecordRiakVnodeIndexOperationCountDataPoint adds a data point to riak.vnode.index.operation.count metric.
func (mb *MetricsBuilder) RecordRiakVnodeIndexOperationCountDataPoint(ts pcommon.Timestamp, val int64, operationAttributeValue AttributeOperation) {
	mb.metricRiakVnodeIndexOperationCount.recordDataPoint(mb.startTime, ts, val, operationAttributeValue.String())
}

// RecordRiakVnodeOperationCountDataPoint adds a data point to riak.vnode.operation.count metric.
func (mb *MetricsBuilder) RecordRiakVnodeOperationCountDataPoint(ts pcommon.Timestamp, val int64, requestAttributeValue AttributeRequest) {
	mb.metricRiakVnodeOperationCount.recordDataPoint(mb.startTime, ts, val, requestAttributeValue.String())
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pcommon.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}
