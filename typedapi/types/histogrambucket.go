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
	"encoding/json"
	"fmt"
)

// HistogramBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L334-L337
type HistogramBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	Key          float64                     `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s HistogramBucket) MarshalJSON() ([]byte, error) {
	type opt HistogramBucket
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.Aggregations {
		tmp[fmt.Sprintf("%s", key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// HistogramBucketBuilder holds HistogramBucket struct and provides a builder API.
type HistogramBucketBuilder struct {
	v *HistogramBucket
}

// NewHistogramBucket provides a builder for the HistogramBucket struct.
func NewHistogramBucketBuilder() *HistogramBucketBuilder {
	r := HistogramBucketBuilder{
		&HistogramBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the HistogramBucket struct
func (rb *HistogramBucketBuilder) Build() HistogramBucket {
	return *rb.v
}

func (rb *HistogramBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *HistogramBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *HistogramBucketBuilder) DocCount(doccount int64) *HistogramBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *HistogramBucketBuilder) Key(key float64) *HistogramBucketBuilder {
	rb.v.Key = key
	return rb
}

func (rb *HistogramBucketBuilder) KeyAsString(keyasstring string) *HistogramBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}
