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

// FetchProfileBreakdown type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/search/_types/profile.ts#L148-L157
type FetchProfileBreakdown struct {
	LoadSource            *int `json:"load_source,omitempty"`
	LoadSourceCount       *int `json:"load_source_count,omitempty"`
	LoadStoredFields      *int `json:"load_stored_fields,omitempty"`
	LoadStoredFieldsCount *int `json:"load_stored_fields_count,omitempty"`
	NextReader            *int `json:"next_reader,omitempty"`
	NextReaderCount       *int `json:"next_reader_count,omitempty"`
	Process               *int `json:"process,omitempty"`
	ProcessCount          *int `json:"process_count,omitempty"`
}

// FetchProfileBreakdownBuilder holds FetchProfileBreakdown struct and provides a builder API.
type FetchProfileBreakdownBuilder struct {
	v *FetchProfileBreakdown
}

// NewFetchProfileBreakdown provides a builder for the FetchProfileBreakdown struct.
func NewFetchProfileBreakdownBuilder() *FetchProfileBreakdownBuilder {
	r := FetchProfileBreakdownBuilder{
		&FetchProfileBreakdown{},
	}

	return &r
}

// Build finalize the chain and returns the FetchProfileBreakdown struct
func (rb *FetchProfileBreakdownBuilder) Build() FetchProfileBreakdown {
	return *rb.v
}

func (rb *FetchProfileBreakdownBuilder) LoadSource(loadsource int) *FetchProfileBreakdownBuilder {
	rb.v.LoadSource = &loadsource
	return rb
}

func (rb *FetchProfileBreakdownBuilder) LoadSourceCount(loadsourcecount int) *FetchProfileBreakdownBuilder {
	rb.v.LoadSourceCount = &loadsourcecount
	return rb
}

func (rb *FetchProfileBreakdownBuilder) LoadStoredFields(loadstoredfields int) *FetchProfileBreakdownBuilder {
	rb.v.LoadStoredFields = &loadstoredfields
	return rb
}

func (rb *FetchProfileBreakdownBuilder) LoadStoredFieldsCount(loadstoredfieldscount int) *FetchProfileBreakdownBuilder {
	rb.v.LoadStoredFieldsCount = &loadstoredfieldscount
	return rb
}

func (rb *FetchProfileBreakdownBuilder) NextReader(nextreader int) *FetchProfileBreakdownBuilder {
	rb.v.NextReader = &nextreader
	return rb
}

func (rb *FetchProfileBreakdownBuilder) NextReaderCount(nextreadercount int) *FetchProfileBreakdownBuilder {
	rb.v.NextReaderCount = &nextreadercount
	return rb
}

func (rb *FetchProfileBreakdownBuilder) Process(process int) *FetchProfileBreakdownBuilder {
	rb.v.Process = &process
	return rb
}

func (rb *FetchProfileBreakdownBuilder) ProcessCount(processcount int) *FetchProfileBreakdownBuilder {
	rb.v.ProcessCount = &processcount
	return rb
}
