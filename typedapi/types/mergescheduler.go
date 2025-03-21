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

// MergeScheduler type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/_types/IndexSettings.ts#L327-L330
type MergeScheduler struct {
	MaxMergeCount  *int `json:"max_merge_count,omitempty"`
	MaxThreadCount *int `json:"max_thread_count,omitempty"`
}

// MergeSchedulerBuilder holds MergeScheduler struct and provides a builder API.
type MergeSchedulerBuilder struct {
	v *MergeScheduler
}

// NewMergeScheduler provides a builder for the MergeScheduler struct.
func NewMergeSchedulerBuilder() *MergeSchedulerBuilder {
	r := MergeSchedulerBuilder{
		&MergeScheduler{},
	}

	return &r
}

// Build finalize the chain and returns the MergeScheduler struct
func (rb *MergeSchedulerBuilder) Build() MergeScheduler {
	return *rb.v
}

func (rb *MergeSchedulerBuilder) MaxMergeCount(maxmergecount int) *MergeSchedulerBuilder {
	rb.v.MaxMergeCount = &maxmergecount
	return rb
}

func (rb *MergeSchedulerBuilder) MaxThreadCount(maxthreadcount int) *MergeSchedulerBuilder {
	rb.v.MaxThreadCount = &maxthreadcount
	return rb
}
