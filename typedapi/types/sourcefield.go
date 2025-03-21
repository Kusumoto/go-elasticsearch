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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sourcefieldmode"
)

// SourceField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/mapping/meta-fields.ts#L58-L65
type SourceField struct {
	Compress          *bool                            `json:"compress,omitempty"`
	CompressThreshold *string                          `json:"compress_threshold,omitempty"`
	Enabled           *bool                            `json:"enabled,omitempty"`
	Excludes          []string                         `json:"excludes,omitempty"`
	Includes          []string                         `json:"includes,omitempty"`
	Mode              *sourcefieldmode.SourceFieldMode `json:"mode,omitempty"`
}

// SourceFieldBuilder holds SourceField struct and provides a builder API.
type SourceFieldBuilder struct {
	v *SourceField
}

// NewSourceField provides a builder for the SourceField struct.
func NewSourceFieldBuilder() *SourceFieldBuilder {
	r := SourceFieldBuilder{
		&SourceField{},
	}

	return &r
}

// Build finalize the chain and returns the SourceField struct
func (rb *SourceFieldBuilder) Build() SourceField {
	return *rb.v
}

func (rb *SourceFieldBuilder) Compress(compress bool) *SourceFieldBuilder {
	rb.v.Compress = &compress
	return rb
}

func (rb *SourceFieldBuilder) CompressThreshold(compressthreshold string) *SourceFieldBuilder {
	rb.v.CompressThreshold = &compressthreshold
	return rb
}

func (rb *SourceFieldBuilder) Enabled(enabled bool) *SourceFieldBuilder {
	rb.v.Enabled = &enabled
	return rb
}

func (rb *SourceFieldBuilder) Excludes(excludes ...string) *SourceFieldBuilder {
	rb.v.Excludes = excludes
	return rb
}

func (rb *SourceFieldBuilder) Includes(includes ...string) *SourceFieldBuilder {
	rb.v.Includes = includes
	return rb
}

func (rb *SourceFieldBuilder) Mode(mode sourcefieldmode.SourceFieldMode) *SourceFieldBuilder {
	rb.v.Mode = &mode
	return rb
}
