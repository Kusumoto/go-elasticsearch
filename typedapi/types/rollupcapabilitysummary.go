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

// RollupCapabilitySummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/rollup/get_rollup_caps/types.ts#L28-L33
type RollupCapabilitySummary struct {
	Fields       map[Field]map[string]interface{} `json:"fields"`
	IndexPattern string                           `json:"index_pattern"`
	JobId        string                           `json:"job_id"`
	RollupIndex  string                           `json:"rollup_index"`
}

// RollupCapabilitySummaryBuilder holds RollupCapabilitySummary struct and provides a builder API.
type RollupCapabilitySummaryBuilder struct {
	v *RollupCapabilitySummary
}

// NewRollupCapabilitySummary provides a builder for the RollupCapabilitySummary struct.
func NewRollupCapabilitySummaryBuilder() *RollupCapabilitySummaryBuilder {
	r := RollupCapabilitySummaryBuilder{
		&RollupCapabilitySummary{
			Fields: make(map[Field]map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RollupCapabilitySummary struct
func (rb *RollupCapabilitySummaryBuilder) Build() RollupCapabilitySummary {
	return *rb.v
}

func (rb *RollupCapabilitySummaryBuilder) Fields(value map[Field]map[string]interface{}) *RollupCapabilitySummaryBuilder {
	rb.v.Fields = value
	return rb
}

func (rb *RollupCapabilitySummaryBuilder) IndexPattern(indexpattern string) *RollupCapabilitySummaryBuilder {
	rb.v.IndexPattern = indexpattern
	return rb
}

func (rb *RollupCapabilitySummaryBuilder) JobId(jobid string) *RollupCapabilitySummaryBuilder {
	rb.v.JobId = jobid
	return rb
}

func (rb *RollupCapabilitySummaryBuilder) RollupIndex(rollupindex string) *RollupCapabilitySummaryBuilder {
	rb.v.RollupIndex = rollupindex
	return rb
}
