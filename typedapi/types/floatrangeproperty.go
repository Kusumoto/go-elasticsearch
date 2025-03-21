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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// FloatRangeProperty type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/mapping/range.ts#L38-L40
type FloatRangeProperty struct {
	Boost         *float64                       `json:"boost,omitempty"`
	Coerce        *bool                          `json:"coerce,omitempty"`
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	Index         *bool                          `json:"index,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	// Meta Metadata about the field.
	Meta       map[string]string         `json:"meta,omitempty"`
	Properties map[PropertyName]Property `json:"properties,omitempty"`
	Similarity *string                   `json:"similarity,omitempty"`
	Store      *bool                     `json:"store,omitempty"`
	Type       string                    `json:"type,omitempty"`
}

// FloatRangePropertyBuilder holds FloatRangeProperty struct and provides a builder API.
type FloatRangePropertyBuilder struct {
	v *FloatRangeProperty
}

// NewFloatRangeProperty provides a builder for the FloatRangeProperty struct.
func NewFloatRangePropertyBuilder() *FloatRangePropertyBuilder {
	r := FloatRangePropertyBuilder{
		&FloatRangeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "float_range"

	return &r
}

// Build finalize the chain and returns the FloatRangeProperty struct
func (rb *FloatRangePropertyBuilder) Build() FloatRangeProperty {
	return *rb.v
}

func (rb *FloatRangePropertyBuilder) Boost(boost float64) *FloatRangePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *FloatRangePropertyBuilder) Coerce(coerce bool) *FloatRangePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

func (rb *FloatRangePropertyBuilder) CopyTo(copyto *FieldsBuilder) *FloatRangePropertyBuilder {
	v := copyto.Build()
	rb.v.CopyTo = &v
	return rb
}

func (rb *FloatRangePropertyBuilder) DocValues(docvalues bool) *FloatRangePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

func (rb *FloatRangePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *FloatRangePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

func (rb *FloatRangePropertyBuilder) Fields(values map[PropertyName]*PropertyBuilder) *FloatRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Fields = tmp
	return rb
}

func (rb *FloatRangePropertyBuilder) IgnoreAbove(ignoreabove int) *FloatRangePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

func (rb *FloatRangePropertyBuilder) Index(index bool) *FloatRangePropertyBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *FloatRangePropertyBuilder) LocalMetadata(localmetadata *MetadataBuilder) *FloatRangePropertyBuilder {
	v := localmetadata.Build()
	rb.v.LocalMetadata = &v
	return rb
}

// Meta Metadata about the field.

func (rb *FloatRangePropertyBuilder) Meta(value map[string]string) *FloatRangePropertyBuilder {
	rb.v.Meta = value
	return rb
}

func (rb *FloatRangePropertyBuilder) Properties(values map[PropertyName]*PropertyBuilder) *FloatRangePropertyBuilder {
	tmp := make(map[PropertyName]Property, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Properties = tmp
	return rb
}

func (rb *FloatRangePropertyBuilder) Similarity(similarity string) *FloatRangePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

func (rb *FloatRangePropertyBuilder) Store(store bool) *FloatRangePropertyBuilder {
	rb.v.Store = &store
	return rb
}
