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

// RankEvalMetricDiscountedCumulativeGain type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/rank_eval/types.ts#L66-L77
type RankEvalMetricDiscountedCumulativeGain struct {
	// K Sets the maximum number of documents retrieved per query. This value will act
	// in place of the usual size parameter in the query.
	K *int `json:"k,omitempty"`
	// Normalize If set to true, this metric will calculate the Normalized DCG.
	Normalize *bool `json:"normalize,omitempty"`
}

// RankEvalMetricDiscountedCumulativeGainBuilder holds RankEvalMetricDiscountedCumulativeGain struct and provides a builder API.
type RankEvalMetricDiscountedCumulativeGainBuilder struct {
	v *RankEvalMetricDiscountedCumulativeGain
}

// NewRankEvalMetricDiscountedCumulativeGain provides a builder for the RankEvalMetricDiscountedCumulativeGain struct.
func NewRankEvalMetricDiscountedCumulativeGainBuilder() *RankEvalMetricDiscountedCumulativeGainBuilder {
	r := RankEvalMetricDiscountedCumulativeGainBuilder{
		&RankEvalMetricDiscountedCumulativeGain{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetricDiscountedCumulativeGain struct
func (rb *RankEvalMetricDiscountedCumulativeGainBuilder) Build() RankEvalMetricDiscountedCumulativeGain {
	return *rb.v
}

// K Sets the maximum number of documents retrieved per query. This value will act
// in place of the usual size parameter in the query.

func (rb *RankEvalMetricDiscountedCumulativeGainBuilder) K(k int) *RankEvalMetricDiscountedCumulativeGainBuilder {
	rb.v.K = &k
	return rb
}

// Normalize If set to true, this metric will calculate the Normalized DCG.

func (rb *RankEvalMetricDiscountedCumulativeGainBuilder) Normalize(normalize bool) *RankEvalMetricDiscountedCumulativeGainBuilder {
	rb.v.Normalize = &normalize
	return rb
}
