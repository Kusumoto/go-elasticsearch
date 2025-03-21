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


package puttemplate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package puttemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/put_template/IndicesPutTemplateRequest.ts#L29-L93
type Request struct {

	// Aliases Aliases for the index.
	Aliases map[types.IndexName]types.Alias `json:"aliases,omitempty"`

	// IndexPatterns Array of wildcard expressions used to match the names
	// of indices during creation.
	IndexPatterns []string `json:"index_patterns,omitempty"`

	// Mappings Mapping for fields in the index.
	Mappings *types.TypeMapping `json:"mappings,omitempty"`

	// Order Order in which Elasticsearch applies this template if index
	// matches multiple templates.
	//
	// Templates with lower 'order' values are merged first. Templates with higher
	// 'order' values are merged later, overriding templates with lower values.
	Order *int `json:"order,omitempty"`

	// Settings Configuration options for the index.
	Settings map[string]interface{} `json:"settings,omitempty"`

	// Version Version number used to manage index templates externally. This number
	// is not automatically generated by Elasticsearch.
	Version *types.VersionNumber `json:"version,omitempty"`
}

// RequestBuilder is the builder API for the puttemplate.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Aliases:  make(map[types.IndexName]types.Alias, 0),
			Settings: make(map[string]interface{}, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Puttemplate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Aliases(values map[types.IndexName]*types.AliasBuilder) *RequestBuilder {
	tmp := make(map[types.IndexName]types.Alias, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.Aliases = tmp
	return rb
}

func (rb *RequestBuilder) IndexPatterns(arg []string) *RequestBuilder {
	rb.v.IndexPatterns = arg
	return rb
}

func (rb *RequestBuilder) Mappings(mappings *types.TypeMappingBuilder) *RequestBuilder {
	v := mappings.Build()
	rb.v.Mappings = &v
	return rb
}

func (rb *RequestBuilder) Order(order int) *RequestBuilder {
	rb.v.Order = &order
	return rb
}

func (rb *RequestBuilder) Settings(value map[string]interface{}) *RequestBuilder {
	rb.v.Settings = value
	return rb
}

func (rb *RequestBuilder) Version(version types.VersionNumber) *RequestBuilder {
	rb.v.Version = &version
	return rb
}
