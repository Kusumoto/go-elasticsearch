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

// DateHistogramBucket type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L342-L345
type DateHistogramBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"-"`
	DocCount     int64                       `json:"doc_count"`
	Key          EpochTimeUnitMillis         `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s DateHistogramBucket) MarshalJSON() ([]byte, error) {
	type opt DateHistogramBucket
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

// DateHistogramBucketBuilder holds DateHistogramBucket struct and provides a builder API.
type DateHistogramBucketBuilder struct {
	v *DateHistogramBucket
}

// NewDateHistogramBucket provides a builder for the DateHistogramBucket struct.
func NewDateHistogramBucketBuilder() *DateHistogramBucketBuilder {
	r := DateHistogramBucketBuilder{
		&DateHistogramBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DateHistogramBucket struct
func (rb *DateHistogramBucketBuilder) Build() DateHistogramBucket {
	return *rb.v
}

func (rb *DateHistogramBucketBuilder) Aggregations(values map[AggregateName]*AggregateBuilder) *DateHistogramBucketBuilder {
	tmp := make(map[AggregateName]Aggregate, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aggregations = tmp
	return rb
}

func (rb *DateHistogramBucketBuilder) DocCount(doccount int64) *DateHistogramBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

func (rb *DateHistogramBucketBuilder) Key(key *EpochTimeUnitMillisBuilder) *DateHistogramBucketBuilder {
	v := key.Build()
	rb.v.Key = v
	return rb
}

func (rb *DateHistogramBucketBuilder) KeyAsString(keyasstring string) *DateHistogramBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}
