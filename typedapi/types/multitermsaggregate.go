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

// MultiTermsAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L452-L454
type MultiTermsAggregate struct {
	Buckets                 BucketsMultiTermsBucket `json:"buckets"`
	DocCountErrorUpperBound *int64                  `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata               `json:"meta,omitempty"`
	SumOtherDocCount        *int64                  `json:"sum_other_doc_count,omitempty"`
}

// MultiTermsAggregateBuilder holds MultiTermsAggregate struct and provides a builder API.
type MultiTermsAggregateBuilder struct {
	v *MultiTermsAggregate
}

// NewMultiTermsAggregate provides a builder for the MultiTermsAggregate struct.
func NewMultiTermsAggregateBuilder() *MultiTermsAggregateBuilder {
	r := MultiTermsAggregateBuilder{
		&MultiTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the MultiTermsAggregate struct
func (rb *MultiTermsAggregateBuilder) Build() MultiTermsAggregate {
	return *rb.v
}

func (rb *MultiTermsAggregateBuilder) Buckets(buckets *BucketsMultiTermsBucketBuilder) *MultiTermsAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *MultiTermsAggregateBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *MultiTermsAggregateBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

func (rb *MultiTermsAggregateBuilder) Meta(meta *MetadataBuilder) *MultiTermsAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *MultiTermsAggregateBuilder) SumOtherDocCount(sumotherdoccount int64) *MultiTermsAggregateBuilder {
	rb.v.SumOtherDocCount = &sumotherdoccount
	return rb
}
