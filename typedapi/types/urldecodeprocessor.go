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

// UrlDecodeProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ingest/_types/Processors.ts#L361-L365
type UrlDecodeProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// UrlDecodeProcessorBuilder holds UrlDecodeProcessor struct and provides a builder API.
type UrlDecodeProcessorBuilder struct {
	v *UrlDecodeProcessor
}

// NewUrlDecodeProcessor provides a builder for the UrlDecodeProcessor struct.
func NewUrlDecodeProcessorBuilder() *UrlDecodeProcessorBuilder {
	r := UrlDecodeProcessorBuilder{
		&UrlDecodeProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the UrlDecodeProcessor struct
func (rb *UrlDecodeProcessorBuilder) Build() UrlDecodeProcessor {
	return *rb.v
}

func (rb *UrlDecodeProcessorBuilder) Field(field Field) *UrlDecodeProcessorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *UrlDecodeProcessorBuilder) If_(if_ string) *UrlDecodeProcessorBuilder {
	rb.v.If = &if_
	return rb
}

func (rb *UrlDecodeProcessorBuilder) IgnoreFailure(ignorefailure bool) *UrlDecodeProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

func (rb *UrlDecodeProcessorBuilder) IgnoreMissing(ignoremissing bool) *UrlDecodeProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

func (rb *UrlDecodeProcessorBuilder) OnFailure(on_failure []ProcessorContainerBuilder) *UrlDecodeProcessorBuilder {
	tmp := make([]ProcessorContainer, len(on_failure))
	for _, value := range on_failure {
		tmp = append(tmp, value.Build())
	}
	rb.v.OnFailure = tmp
	return rb
}

func (rb *UrlDecodeProcessorBuilder) Tag(tag string) *UrlDecodeProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

func (rb *UrlDecodeProcessorBuilder) TargetField(targetfield Field) *UrlDecodeProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
