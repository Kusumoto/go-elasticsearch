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


package stopdatafeed

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package stopdatafeed
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/stop_datafeed/MlStopDatafeedRequest.ts#L24-L78
type Request struct {

	// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`

	// Force Refer to the description for the `force` query parameter.
	Force *bool `json:"force,omitempty"`

	// Timeout Refer to the description for the `timeout` query parameter.
	Timeout *types.Duration `json:"timeout,omitempty"`
}

// RequestBuilder is the builder API for the stopdatafeed.Request
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
		return nil, fmt.Errorf("could not deserialise json into Stopdatafeed request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AllowNoMatch(allownomatch bool) *RequestBuilder {
	rb.v.AllowNoMatch = &allownomatch
	return rb
}

func (rb *RequestBuilder) Force(force bool) *RequestBuilder {
	rb.v.Force = &force
	return rb
}

func (rb *RequestBuilder) Timeout(timeout *types.DurationBuilder) *RequestBuilder {
	v := timeout.Build()
	rb.v.Timeout = &v
	return rb
}
