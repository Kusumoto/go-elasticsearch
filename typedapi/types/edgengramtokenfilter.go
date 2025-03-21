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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/edgengramside"
)

// EdgeNGramTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/analysis/token_filters.ts#L78-L84
type EdgeNGramTokenFilter struct {
	MaxGram          *int                         `json:"max_gram,omitempty"`
	MinGram          *int                         `json:"min_gram,omitempty"`
	PreserveOriginal *bool                        `json:"preserve_original,omitempty"`
	Side             *edgengramside.EdgeNGramSide `json:"side,omitempty"`
	Type             string                       `json:"type,omitempty"`
	Version          *VersionString               `json:"version,omitempty"`
}

// EdgeNGramTokenFilterBuilder holds EdgeNGramTokenFilter struct and provides a builder API.
type EdgeNGramTokenFilterBuilder struct {
	v *EdgeNGramTokenFilter
}

// NewEdgeNGramTokenFilter provides a builder for the EdgeNGramTokenFilter struct.
func NewEdgeNGramTokenFilterBuilder() *EdgeNGramTokenFilterBuilder {
	r := EdgeNGramTokenFilterBuilder{
		&EdgeNGramTokenFilter{},
	}

	r.v.Type = "edge_ngram"

	return &r
}

// Build finalize the chain and returns the EdgeNGramTokenFilter struct
func (rb *EdgeNGramTokenFilterBuilder) Build() EdgeNGramTokenFilter {
	return *rb.v
}

func (rb *EdgeNGramTokenFilterBuilder) MaxGram(maxgram int) *EdgeNGramTokenFilterBuilder {
	rb.v.MaxGram = &maxgram
	return rb
}

func (rb *EdgeNGramTokenFilterBuilder) MinGram(mingram int) *EdgeNGramTokenFilterBuilder {
	rb.v.MinGram = &mingram
	return rb
}

func (rb *EdgeNGramTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *EdgeNGramTokenFilterBuilder {
	rb.v.PreserveOriginal = &preserveoriginal
	return rb
}

func (rb *EdgeNGramTokenFilterBuilder) Side(side edgengramside.EdgeNGramSide) *EdgeNGramTokenFilterBuilder {
	rb.v.Side = &side
	return rb
}

func (rb *EdgeNGramTokenFilterBuilder) Version(version VersionString) *EdgeNGramTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
