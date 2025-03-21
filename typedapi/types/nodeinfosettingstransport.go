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

// NodeInfoSettingsTransport type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/nodes/info/types.ts#L197-L201
type NodeInfoSettingsTransport struct {
	Features    *NodeInfoSettingsTransportFeatures `json:"features,omitempty"`
	Type        NodeInfoSettingsTransportType      `json:"type"`
	TypeDefault *string                            `json:"type.default,omitempty"`
}

// NodeInfoSettingsTransportBuilder holds NodeInfoSettingsTransport struct and provides a builder API.
type NodeInfoSettingsTransportBuilder struct {
	v *NodeInfoSettingsTransport
}

// NewNodeInfoSettingsTransport provides a builder for the NodeInfoSettingsTransport struct.
func NewNodeInfoSettingsTransportBuilder() *NodeInfoSettingsTransportBuilder {
	r := NodeInfoSettingsTransportBuilder{
		&NodeInfoSettingsTransport{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsTransport struct
func (rb *NodeInfoSettingsTransportBuilder) Build() NodeInfoSettingsTransport {
	return *rb.v
}

func (rb *NodeInfoSettingsTransportBuilder) Features(features *NodeInfoSettingsTransportFeaturesBuilder) *NodeInfoSettingsTransportBuilder {
	v := features.Build()
	rb.v.Features = &v
	return rb
}

func (rb *NodeInfoSettingsTransportBuilder) Type_(type_ *NodeInfoSettingsTransportTypeBuilder) *NodeInfoSettingsTransportBuilder {
	v := type_.Build()
	rb.v.Type = v
	return rb
}

func (rb *NodeInfoSettingsTransportBuilder) TypeDefault(typedefault string) *NodeInfoSettingsTransportBuilder {
	rb.v.TypeDefault = &typedefault
	return rb
}
