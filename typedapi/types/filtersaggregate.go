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

// FiltersAggregate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L552-L553
type FiltersAggregate struct {
	Buckets BucketsFiltersBucket `json:"buckets"`
	Meta    *Metadata            `json:"meta,omitempty"`
}

// FiltersAggregateBuilder holds FiltersAggregate struct and provides a builder API.
type FiltersAggregateBuilder struct {
	v *FiltersAggregate
}

// NewFiltersAggregate provides a builder for the FiltersAggregate struct.
func NewFiltersAggregateBuilder() *FiltersAggregateBuilder {
	r := FiltersAggregateBuilder{
		&FiltersAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the FiltersAggregate struct
func (rb *FiltersAggregateBuilder) Build() FiltersAggregate {
	return *rb.v
}

func (rb *FiltersAggregateBuilder) Buckets(buckets *BucketsFiltersBucketBuilder) *FiltersAggregateBuilder {
	v := buckets.Build()
	rb.v.Buckets = v
	return rb
}

func (rb *FiltersAggregateBuilder) Meta(meta *MetadataBuilder) *FiltersAggregateBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}
