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

// AggregationProfileDelegateDebugFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/search/_types/profile.ts#L70-L75
type AggregationProfileDelegateDebugFilter struct {
	Query                         *string `json:"query,omitempty"`
	ResultsFromMetadata           *int    `json:"results_from_metadata,omitempty"`
	SegmentsCountedInConstantTime *int    `json:"segments_counted_in_constant_time,omitempty"`
	SpecializedFor                *string `json:"specialized_for,omitempty"`
}

// AggregationProfileDelegateDebugFilterBuilder holds AggregationProfileDelegateDebugFilter struct and provides a builder API.
type AggregationProfileDelegateDebugFilterBuilder struct {
	v *AggregationProfileDelegateDebugFilter
}

// NewAggregationProfileDelegateDebugFilter provides a builder for the AggregationProfileDelegateDebugFilter struct.
func NewAggregationProfileDelegateDebugFilterBuilder() *AggregationProfileDelegateDebugFilterBuilder {
	r := AggregationProfileDelegateDebugFilterBuilder{
		&AggregationProfileDelegateDebugFilter{},
	}

	return &r
}

// Build finalize the chain and returns the AggregationProfileDelegateDebugFilter struct
func (rb *AggregationProfileDelegateDebugFilterBuilder) Build() AggregationProfileDelegateDebugFilter {
	return *rb.v
}

func (rb *AggregationProfileDelegateDebugFilterBuilder) Query(query string) *AggregationProfileDelegateDebugFilterBuilder {
	rb.v.Query = &query
	return rb
}

func (rb *AggregationProfileDelegateDebugFilterBuilder) ResultsFromMetadata(resultsfrommetadata int) *AggregationProfileDelegateDebugFilterBuilder {
	rb.v.ResultsFromMetadata = &resultsfrommetadata
	return rb
}

func (rb *AggregationProfileDelegateDebugFilterBuilder) SegmentsCountedInConstantTime(segmentscountedinconstanttime int) *AggregationProfileDelegateDebugFilterBuilder {
	rb.v.SegmentsCountedInConstantTime = &segmentscountedinconstanttime
	return rb
}

func (rb *AggregationProfileDelegateDebugFilterBuilder) SpecializedFor(specializedfor string) *AggregationProfileDelegateDebugFilterBuilder {
	rb.v.SpecializedFor = &specializedfor
	return rb
}
