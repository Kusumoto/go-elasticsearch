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

// GeoLineSort type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/metric.ts#L89-L91
type GeoLineSort struct {
	Field Field `json:"field"`
}

// GeoLineSortBuilder holds GeoLineSort struct and provides a builder API.
type GeoLineSortBuilder struct {
	v *GeoLineSort
}

// NewGeoLineSort provides a builder for the GeoLineSort struct.
func NewGeoLineSortBuilder() *GeoLineSortBuilder {
	r := GeoLineSortBuilder{
		&GeoLineSort{},
	}

	return &r
}

// Build finalize the chain and returns the GeoLineSort struct
func (rb *GeoLineSortBuilder) Build() GeoLineSort {
	return *rb.v
}

func (rb *GeoLineSortBuilder) Field(field Field) *GeoLineSortBuilder {
	rb.v.Field = field
	return rb
}
