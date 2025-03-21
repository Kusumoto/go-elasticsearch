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

// FillMaskInferenceUpdateOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/inference.ts#L364-L371
type FillMaskInferenceUpdateOptions struct {
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// Tokenization The tokenization options to update when inferring
	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

// FillMaskInferenceUpdateOptionsBuilder holds FillMaskInferenceUpdateOptions struct and provides a builder API.
type FillMaskInferenceUpdateOptionsBuilder struct {
	v *FillMaskInferenceUpdateOptions
}

// NewFillMaskInferenceUpdateOptions provides a builder for the FillMaskInferenceUpdateOptions struct.
func NewFillMaskInferenceUpdateOptionsBuilder() *FillMaskInferenceUpdateOptionsBuilder {
	r := FillMaskInferenceUpdateOptionsBuilder{
		&FillMaskInferenceUpdateOptions{},
	}

	return &r
}

// Build finalize the chain and returns the FillMaskInferenceUpdateOptions struct
func (rb *FillMaskInferenceUpdateOptionsBuilder) Build() FillMaskInferenceUpdateOptions {
	return *rb.v
}

// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.

func (rb *FillMaskInferenceUpdateOptionsBuilder) NumTopClasses(numtopclasses int) *FillMaskInferenceUpdateOptionsBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *FillMaskInferenceUpdateOptionsBuilder) ResultsField(resultsfield string) *FillMaskInferenceUpdateOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// Tokenization The tokenization options to update when inferring

func (rb *FillMaskInferenceUpdateOptionsBuilder) Tokenization(tokenization *NlpTokenizationUpdateOptionsBuilder) *FillMaskInferenceUpdateOptionsBuilder {
	v := tokenization.Build()
	rb.v.Tokenization = &v
	return rb
}
