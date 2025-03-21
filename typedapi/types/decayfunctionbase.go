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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
)

// DecayFunctionBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_types/query_dsl/compound.ts#L84-L86
type DecayFunctionBase struct {
	MultiValueMode *multivaluemode.MultiValueMode `json:"multi_value_mode,omitempty"`
}

// DecayFunctionBaseBuilder holds DecayFunctionBase struct and provides a builder API.
type DecayFunctionBaseBuilder struct {
	v *DecayFunctionBase
}

// NewDecayFunctionBase provides a builder for the DecayFunctionBase struct.
func NewDecayFunctionBaseBuilder() *DecayFunctionBaseBuilder {
	r := DecayFunctionBaseBuilder{
		&DecayFunctionBase{},
	}

	return &r
}

// Build finalize the chain and returns the DecayFunctionBase struct
func (rb *DecayFunctionBaseBuilder) Build() DecayFunctionBase {
	return *rb.v
}

func (rb *DecayFunctionBaseBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *DecayFunctionBaseBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}
