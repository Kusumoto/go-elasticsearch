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

// ValueCountAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/aggregations/metric.ts#L196-L196
type ValueCountAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// ValueCountAggregationBuilder holds ValueCountAggregation struct and provides a builder API.
type ValueCountAggregationBuilder struct {
	v *ValueCountAggregation
}

// NewValueCountAggregation provides a builder for the ValueCountAggregation struct.
func NewValueCountAggregationBuilder() *ValueCountAggregationBuilder {
	r := ValueCountAggregationBuilder{
		&ValueCountAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the ValueCountAggregation struct
func (rb *ValueCountAggregationBuilder) Build() ValueCountAggregation {
	return *rb.v
}

func (rb *ValueCountAggregationBuilder) Field(field Field) *ValueCountAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *ValueCountAggregationBuilder) Format(format string) *ValueCountAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *ValueCountAggregationBuilder) Missing(missing *MissingBuilder) *ValueCountAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *ValueCountAggregationBuilder) Script(script *ScriptBuilder) *ValueCountAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
