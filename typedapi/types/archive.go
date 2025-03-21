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

// Archive type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/xpack/usage/types.ts#L44-L46
type Archive struct {
	Available    bool  `json:"available"`
	Enabled      bool  `json:"enabled"`
	IndicesCount int64 `json:"indices_count"`
}

// ArchiveBuilder holds Archive struct and provides a builder API.
type ArchiveBuilder struct {
	v *Archive
}

// NewArchive provides a builder for the Archive struct.
func NewArchiveBuilder() *ArchiveBuilder {
	r := ArchiveBuilder{
		&Archive{},
	}

	return &r
}

// Build finalize the chain and returns the Archive struct
func (rb *ArchiveBuilder) Build() Archive {
	return *rb.v
}

func (rb *ArchiveBuilder) Available(available bool) *ArchiveBuilder {
	rb.v.Available = available
	return rb
}

func (rb *ArchiveBuilder) Enabled(enabled bool) *ArchiveBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *ArchiveBuilder) IndicesCount(indicescount int64) *ArchiveBuilder {
	rb.v.IndicesCount = indicescount
	return rb
}
