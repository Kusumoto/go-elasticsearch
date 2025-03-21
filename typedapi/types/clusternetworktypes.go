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

// ClusterNetworkTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/cluster/stats/types.ts#L171-L174
type ClusterNetworkTypes struct {
	HttpTypes      map[string]int `json:"http_types"`
	TransportTypes map[string]int `json:"transport_types"`
}

// ClusterNetworkTypesBuilder holds ClusterNetworkTypes struct and provides a builder API.
type ClusterNetworkTypesBuilder struct {
	v *ClusterNetworkTypes
}

// NewClusterNetworkTypes provides a builder for the ClusterNetworkTypes struct.
func NewClusterNetworkTypesBuilder() *ClusterNetworkTypesBuilder {
	r := ClusterNetworkTypesBuilder{
		&ClusterNetworkTypes{
			HttpTypes:      make(map[string]int, 0),
			TransportTypes: make(map[string]int, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ClusterNetworkTypes struct
func (rb *ClusterNetworkTypesBuilder) Build() ClusterNetworkTypes {
	return *rb.v
}

func (rb *ClusterNetworkTypesBuilder) HttpTypes(value map[string]int) *ClusterNetworkTypesBuilder {
	rb.v.HttpTypes = value
	return rb
}

func (rb *ClusterNetworkTypesBuilder) TransportTypes(value map[string]int) *ClusterNetworkTypesBuilder {
	rb.v.TransportTypes = value
	return rb
}
