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

// Duration holds the union for the following types:
//
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/Time.ts#L52-L58
type Duration interface{}

// DurationBuilder holds Duration struct and provides a builder API.
type DurationBuilder struct {
	v Duration
}

// NewDuration provides a builder for the Duration struct.
func NewDurationBuilder() *DurationBuilder {
	return &DurationBuilder{}
}

// Build finalize the chain and returns the Duration struct
func (u *DurationBuilder) Build() Duration {
	return u.v
}

func (u *DurationBuilder) String(string string) *DurationBuilder {
	u.v = &string
	return u
}
