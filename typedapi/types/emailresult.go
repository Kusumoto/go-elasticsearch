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

// EmailResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/watcher/_types/Actions.ts#L205-L209
type EmailResult struct {
	Account *string `json:"account,omitempty"`
	Message Email   `json:"message"`
	Reason  *string `json:"reason,omitempty"`
}

// EmailResultBuilder holds EmailResult struct and provides a builder API.
type EmailResultBuilder struct {
	v *EmailResult
}

// NewEmailResult provides a builder for the EmailResult struct.
func NewEmailResultBuilder() *EmailResultBuilder {
	r := EmailResultBuilder{
		&EmailResult{},
	}

	return &r
}

// Build finalize the chain and returns the EmailResult struct
func (rb *EmailResultBuilder) Build() EmailResult {
	return *rb.v
}

func (rb *EmailResultBuilder) Account(account string) *EmailResultBuilder {
	rb.v.Account = &account
	return rb
}

func (rb *EmailResultBuilder) Message(message *EmailBuilder) *EmailResultBuilder {
	v := message.Build()
	rb.v.Message = v
	return rb
}

func (rb *EmailResultBuilder) Reason(reason string) *EmailResultBuilder {
	rb.v.Reason = &reason
	return rb
}
