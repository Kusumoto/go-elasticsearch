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


package getrecords

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package getrecords
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/get_records/MlGetAnomalyRecordsRequest.ts#L26-L127
type Request struct {

	// Desc Refer to the description for the `desc` query parameter.
	Desc *bool `json:"desc,omitempty"`

	// End Refer to the description for the `end` query parameter.
	End *types.DateTime `json:"end,omitempty"`

	// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
	ExcludeInterim *bool `json:"exclude_interim,omitempty"`

	Page *types.Page `json:"page,omitempty"`

	// RecordScore Refer to the description for the `record_score` query parameter.
	RecordScore *float64 `json:"record_score,omitempty"`

	// Sort Refer to the description for the `sort` query parameter.
	Sort *types.Field `json:"sort,omitempty"`

	// Start Refer to the description for the `start` query parameter.
	Start *types.DateTime `json:"start,omitempty"`
}

// RequestBuilder is the builder API for the getrecords.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Getrecords request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Desc(desc bool) *RequestBuilder {
	rb.v.Desc = &desc
	return rb
}

func (rb *RequestBuilder) End(end *types.DateTimeBuilder) *RequestBuilder {
	v := end.Build()
	rb.v.End = &v
	return rb
}

func (rb *RequestBuilder) ExcludeInterim(excludeinterim bool) *RequestBuilder {
	rb.v.ExcludeInterim = &excludeinterim
	return rb
}

func (rb *RequestBuilder) Page(page *types.PageBuilder) *RequestBuilder {
	v := page.Build()
	rb.v.Page = &v
	return rb
}

func (rb *RequestBuilder) RecordScore(recordscore float64) *RequestBuilder {
	rb.v.RecordScore = &recordscore
	return rb
}

func (rb *RequestBuilder) Sort(sort types.Field) *RequestBuilder {
	rb.v.Sort = &sort
	return rb
}

func (rb *RequestBuilder) Start(start *types.DateTimeBuilder) *RequestBuilder {
	v := start.Build()
	rb.v.Start = &v
	return rb
}
