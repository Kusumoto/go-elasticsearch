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

// AddAction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/update_aliases/types.ts#L30-L44
type AddAction struct {
	Alias         *IndexAlias     `json:"alias,omitempty"`
	Aliases       []IndexAlias    `json:"aliases,omitempty"`
	Filter        *QueryContainer `json:"filter,omitempty"`
	Index         *IndexName      `json:"index,omitempty"`
	IndexRouting  *Routing        `json:"index_routing,omitempty"`
	Indices       *Indices        `json:"indices,omitempty"`
	IsHidden      *bool           `json:"is_hidden,omitempty"`
	IsWriteIndex  *bool           `json:"is_write_index,omitempty"`
	MustExist     *bool           `json:"must_exist,omitempty"`
	Routing       *Routing        `json:"routing,omitempty"`
	SearchRouting *Routing        `json:"search_routing,omitempty"`
}

// AddActionBuilder holds AddAction struct and provides a builder API.
type AddActionBuilder struct {
	v *AddAction
}

// NewAddAction provides a builder for the AddAction struct.
func NewAddActionBuilder() *AddActionBuilder {
	r := AddActionBuilder{
		&AddAction{},
	}

	return &r
}

// Build finalize the chain and returns the AddAction struct
func (rb *AddActionBuilder) Build() AddAction {
	return *rb.v
}

func (rb *AddActionBuilder) Alias(alias IndexAlias) *AddActionBuilder {
	rb.v.Alias = &alias
	return rb
}

func (rb *AddActionBuilder) Aliases(arg []IndexAlias) *AddActionBuilder {
	rb.v.Aliases = arg
	return rb
}

func (rb *AddActionBuilder) Filter(filter *QueryContainerBuilder) *AddActionBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *AddActionBuilder) Index(index IndexName) *AddActionBuilder {
	rb.v.Index = &index
	return rb
}

func (rb *AddActionBuilder) IndexRouting(indexrouting Routing) *AddActionBuilder {
	rb.v.IndexRouting = &indexrouting
	return rb
}

func (rb *AddActionBuilder) Indices(indices *IndicesBuilder) *AddActionBuilder {
	v := indices.Build()
	rb.v.Indices = &v
	return rb
}

func (rb *AddActionBuilder) IsHidden(ishidden bool) *AddActionBuilder {
	rb.v.IsHidden = &ishidden
	return rb
}

func (rb *AddActionBuilder) IsWriteIndex(iswriteindex bool) *AddActionBuilder {
	rb.v.IsWriteIndex = &iswriteindex
	return rb
}

func (rb *AddActionBuilder) MustExist(mustexist bool) *AddActionBuilder {
	rb.v.MustExist = &mustexist
	return rb
}

func (rb *AddActionBuilder) Routing(routing Routing) *AddActionBuilder {
	rb.v.Routing = &routing
	return rb
}

func (rb *AddActionBuilder) SearchRouting(searchrouting Routing) *AddActionBuilder {
	rb.v.SearchRouting = &searchrouting
	return rb
}
