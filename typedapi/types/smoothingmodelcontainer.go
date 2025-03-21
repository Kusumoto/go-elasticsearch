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

// SmoothingModelContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/_global/search/_types/suggester.ts#L224-L231
type SmoothingModelContainer struct {
	Laplace             *LaplaceSmoothingModel             `json:"laplace,omitempty"`
	LinearInterpolation *LinearInterpolationSmoothingModel `json:"linear_interpolation,omitempty"`
	StupidBackoff       *StupidBackoffSmoothingModel       `json:"stupid_backoff,omitempty"`
}

// SmoothingModelContainerBuilder holds SmoothingModelContainer struct and provides a builder API.
type SmoothingModelContainerBuilder struct {
	v *SmoothingModelContainer
}

// NewSmoothingModelContainer provides a builder for the SmoothingModelContainer struct.
func NewSmoothingModelContainerBuilder() *SmoothingModelContainerBuilder {
	r := SmoothingModelContainerBuilder{
		&SmoothingModelContainer{},
	}

	return &r
}

// Build finalize the chain and returns the SmoothingModelContainer struct
func (rb *SmoothingModelContainerBuilder) Build() SmoothingModelContainer {
	return *rb.v
}

func (rb *SmoothingModelContainerBuilder) Laplace(laplace *LaplaceSmoothingModelBuilder) *SmoothingModelContainerBuilder {
	v := laplace.Build()
	rb.v.Laplace = &v
	return rb
}

func (rb *SmoothingModelContainerBuilder) LinearInterpolation(linearinterpolation *LinearInterpolationSmoothingModelBuilder) *SmoothingModelContainerBuilder {
	v := linearinterpolation.Build()
	rb.v.LinearInterpolation = &v
	return rb
}

func (rb *SmoothingModelContainerBuilder) StupidBackoff(stupidbackoff *StupidBackoffSmoothingModelBuilder) *SmoothingModelContainerBuilder {
	v := stupidbackoff.Build()
	rb.v.StupidBackoff = &v
	return rb
}
