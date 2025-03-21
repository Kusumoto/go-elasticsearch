// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/excludefrequent"
)

// Detector type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/Detector.ts#L25-L67
type Detector struct {
	// ByFieldName The field used to split the data. In particular, this property is used for
	// analyzing the splits with respect to their own history. It is used for
	// finding unusual values in the context of the split.
	ByFieldName *Field `json:"by_field_name,omitempty"`
	// CustomRules Custom rules enable you to customize the way detectors operate. For example,
	// a rule may dictate conditions under which results should be skipped. Kibana
	// refers to custom rules as job rules.
	CustomRules []DetectionRule `json:"custom_rules,omitempty"`
	// DetectorDescription A description of the detector.
	DetectorDescription *string `json:"detector_description,omitempty"`
	// DetectorIndex A unique identifier for the detector. This identifier is based on the order
	// of the detectors in the `analysis_config`, starting at zero. If you specify a
	// value for this property, it is ignored.
	DetectorIndex *int `json:"detector_index,omitempty"`
	// ExcludeFrequent If set, frequent entities are excluded from influencing the anomaly results.
	// Entities can be considered frequent over time or frequent in a population. If
	// you are working with both over and by fields, you can set `exclude_frequent`
	// to `all` for both fields, or to `by` or `over` for those specific fields.
	ExcludeFrequent *excludefrequent.ExcludeFrequent `json:"exclude_frequent,omitempty"`
	// FieldName The field that the detector uses in the function. If you use an event rate
	// function such as count or rare, do not specify this field. The `field_name`
	// cannot contain double quotes or backslashes.
	FieldName *Field `json:"field_name,omitempty"`
	// Function The analysis function that is used. For example, `count`, `rare`, `mean`,
	// `min`, `max`, or `sum`.
	Function string `json:"function"`
	// OverFieldName The field used to split the data. In particular, this property is used for
	// analyzing the splits with respect to the history of all splits. It is used
	// for finding unusual values in the population of all splits.
	OverFieldName *Field `json:"over_field_name,omitempty"`
	// PartitionFieldName The field used to segment the analysis. When you use this property, you have
	// completely independent baselines for each value of this field.
	PartitionFieldName *Field `json:"partition_field_name,omitempty"`
	// UseNull Defines whether a new series is used as the null series when there is no
	// value for the by or partition fields.
	UseNull *bool `json:"use_null,omitempty"`
}

// DetectorBuilder holds Detector struct and provides a builder API.
type DetectorBuilder struct {
	v *Detector
}

// NewDetector provides a builder for the Detector struct.
func NewDetectorBuilder() *DetectorBuilder {
	r := DetectorBuilder{
		&Detector{},
	}

	return &r
}

// Build finalize the chain and returns the Detector struct
func (rb *DetectorBuilder) Build() Detector {
	return *rb.v
}

// ByFieldName The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to their own history. It is used for
// finding unusual values in the context of the split.

func (rb *DetectorBuilder) ByFieldName(byfieldname Field) *DetectorBuilder {
	rb.v.ByFieldName = &byfieldname
	return rb
}

// CustomRules Custom rules enable you to customize the way detectors operate. For example,
// a rule may dictate conditions under which results should be skipped. Kibana
// refers to custom rules as job rules.

func (rb *DetectorBuilder) CustomRules(custom_rules []DetectionRuleBuilder) *DetectorBuilder {
	tmp := make([]DetectionRule, len(custom_rules))
	for _, value := range custom_rules {
		tmp = append(tmp, value.Build())
	}
	rb.v.CustomRules = tmp
	return rb
}

// DetectorDescription A description of the detector.

func (rb *DetectorBuilder) DetectorDescription(detectordescription string) *DetectorBuilder {
	rb.v.DetectorDescription = &detectordescription
	return rb
}

// DetectorIndex A unique identifier for the detector. This identifier is based on the order
// of the detectors in the `analysis_config`, starting at zero. If you specify a
// value for this property, it is ignored.

func (rb *DetectorBuilder) DetectorIndex(detectorindex int) *DetectorBuilder {
	rb.v.DetectorIndex = &detectorindex
	return rb
}

// ExcludeFrequent If set, frequent entities are excluded from influencing the anomaly results.
// Entities can be considered frequent over time or frequent in a population. If
// you are working with both over and by fields, you can set `exclude_frequent`
// to `all` for both fields, or to `by` or `over` for those specific fields.

func (rb *DetectorBuilder) ExcludeFrequent(excludefrequent excludefrequent.ExcludeFrequent) *DetectorBuilder {
	rb.v.ExcludeFrequent = &excludefrequent
	return rb
}

// FieldName The field that the detector uses in the function. If you use an event rate
// function such as count or rare, do not specify this field. The `field_name`
// cannot contain double quotes or backslashes.

func (rb *DetectorBuilder) FieldName(fieldname Field) *DetectorBuilder {
	rb.v.FieldName = &fieldname
	return rb
}

// Function The analysis function that is used. For example, `count`, `rare`, `mean`,
// `min`, `max`, or `sum`.

func (rb *DetectorBuilder) Function(function string) *DetectorBuilder {
	rb.v.Function = function
	return rb
}

// OverFieldName The field used to split the data. In particular, this property is used for
// analyzing the splits with respect to the history of all splits. It is used
// for finding unusual values in the population of all splits.

func (rb *DetectorBuilder) OverFieldName(overfieldname Field) *DetectorBuilder {
	rb.v.OverFieldName = &overfieldname
	return rb
}

// PartitionFieldName The field used to segment the analysis. When you use this property, you have
// completely independent baselines for each value of this field.

func (rb *DetectorBuilder) PartitionFieldName(partitionfieldname Field) *DetectorBuilder {
	rb.v.PartitionFieldName = &partitionfieldname
	return rb
}

// UseNull Defines whether a new series is used as the null series when there is no
// value for the by or partition fields.

func (rb *DetectorBuilder) UseNull(usenull bool) *DetectorBuilder {
	rb.v.UseNull = &usenull
	return rb
}
