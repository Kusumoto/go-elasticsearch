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

// PluginsRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/cat/plugins/types.ts#L22-L52
type PluginsRecord struct {
	// Component component
	Component *string `json:"component,omitempty"`
	// Description plugin details
	Description *string `json:"description,omitempty"`
	// Id unique node id
	Id *NodeId `json:"id,omitempty"`
	// Name node name
	Name *Name `json:"name,omitempty"`
	// Type plugin type
	Type *string `json:"type,omitempty"`
	// Version component version
	Version *VersionString `json:"version,omitempty"`
}

// PluginsRecordBuilder holds PluginsRecord struct and provides a builder API.
type PluginsRecordBuilder struct {
	v *PluginsRecord
}

// NewPluginsRecord provides a builder for the PluginsRecord struct.
func NewPluginsRecordBuilder() *PluginsRecordBuilder {
	r := PluginsRecordBuilder{
		&PluginsRecord{},
	}

	return &r
}

// Build finalize the chain and returns the PluginsRecord struct
func (rb *PluginsRecordBuilder) Build() PluginsRecord {
	return *rb.v
}

// Component component

func (rb *PluginsRecordBuilder) Component(component string) *PluginsRecordBuilder {
	rb.v.Component = &component
	return rb
}

// Description plugin details

func (rb *PluginsRecordBuilder) Description(description string) *PluginsRecordBuilder {
	rb.v.Description = &description
	return rb
}

// Id unique node id

func (rb *PluginsRecordBuilder) Id(id NodeId) *PluginsRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Name node name

func (rb *PluginsRecordBuilder) Name(name Name) *PluginsRecordBuilder {
	rb.v.Name = &name
	return rb
}

// Type plugin type

func (rb *PluginsRecordBuilder) Type_(type_ string) *PluginsRecordBuilder {
	rb.v.Type = &type_
	return rb
}

// Version component version

func (rb *PluginsRecordBuilder) Version(version VersionString) *PluginsRecordBuilder {
	rb.v.Version = &version
	return rb
}
