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

// BucketsLongTermsBucket holds the union for the following types:
//
//	[]LongTermsBucket
//	map[string]LongTermsBucket
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/Aggregate.ts#L307-L316
type BucketsLongTermsBucket interface{}

// BucketsLongTermsBucketBuilder holds BucketsLongTermsBucket struct and provides a builder API.
type BucketsLongTermsBucketBuilder struct {
	v BucketsLongTermsBucket
}

// NewBucketsLongTermsBucket provides a builder for the BucketsLongTermsBucket struct.
func NewBucketsLongTermsBucketBuilder() *BucketsLongTermsBucketBuilder {
	return &BucketsLongTermsBucketBuilder{}
}

// Build finalize the chain and returns the BucketsLongTermsBucket struct
func (u *BucketsLongTermsBucketBuilder) Build() BucketsLongTermsBucket {
	return u.v
}

func (u *BucketsLongTermsBucketBuilder) LongTermsBuckets(longtermsbuckets []LongTermsBucketBuilder) *BucketsLongTermsBucketBuilder {
	tmp := make([]LongTermsBucket, len(longtermsbuckets))
	for _, value := range longtermsbuckets {
		tmp = append(tmp, value.Build())
	}
	u.v = tmp
	return u
}

func (u *BucketsLongTermsBucketBuilder) Map(values map[string]*LongTermsBucketBuilder) *BucketsLongTermsBucketBuilder {
	tmp := make(map[string]LongTermsBucket, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	u.v = tmp
	return u
}
