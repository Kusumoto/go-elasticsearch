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


package samlauthenticate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package samlauthenticate
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/security/saml_authenticate/Request.ts#L23-L38
type Request struct {

	// Content The SAML response as it was sent by the user’s browser, usually a Base64
	// encoded XML document.
	Content string `json:"content"`

	// Ids A json array with all the valid SAML Request Ids that the caller of the API
	// has for the current user.
	Ids types.Ids `json:"ids"`

	// Realm The name of the realm that should authenticate the SAML response. Useful in
	// cases where many SAML realms are defined.
	Realm *string `json:"realm,omitempty"`
}

// RequestBuilder is the builder API for the samlauthenticate.Request
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
		return nil, fmt.Errorf("could not deserialise json into Samlauthenticate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Content(content string) *RequestBuilder {
	rb.v.Content = content
	return rb
}

func (rb *RequestBuilder) Ids(ids *types.IdsBuilder) *RequestBuilder {
	v := ids.Build()
	rb.v.Ids = v
	return rb
}

func (rb *RequestBuilder) Realm(realm string) *RequestBuilder {
	rb.v.Realm = &realm
	return rb
}
