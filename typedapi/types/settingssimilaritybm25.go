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

// SettingsSimilarityBm25 type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/indices/_types/IndexSettings.ts#L180-L185
type SettingsSimilarityBm25 struct {
	B                float64 `json:"b"`
	DiscountOverlaps bool    `json:"discount_overlaps"`
	K1               float64 `json:"k1"`
	Type             string  `json:"type,omitempty"`
}

// SettingsSimilarityBm25Builder holds SettingsSimilarityBm25 struct and provides a builder API.
type SettingsSimilarityBm25Builder struct {
	v *SettingsSimilarityBm25
}

// NewSettingsSimilarityBm25 provides a builder for the SettingsSimilarityBm25 struct.
func NewSettingsSimilarityBm25Builder() *SettingsSimilarityBm25Builder {
	r := SettingsSimilarityBm25Builder{
		&SettingsSimilarityBm25{},
	}

	r.v.Type = "BM25"

	return &r
}

// Build finalize the chain and returns the SettingsSimilarityBm25 struct
func (rb *SettingsSimilarityBm25Builder) Build() SettingsSimilarityBm25 {
	return *rb.v
}

func (rb *SettingsSimilarityBm25Builder) B(b float64) *SettingsSimilarityBm25Builder {
	rb.v.B = b
	return rb
}

func (rb *SettingsSimilarityBm25Builder) DiscountOverlaps(discountoverlaps bool) *SettingsSimilarityBm25Builder {
	rb.v.DiscountOverlaps = discountoverlaps
	return rb
}

func (rb *SettingsSimilarityBm25Builder) K1(k1 float64) *SettingsSimilarityBm25Builder {
	rb.v.K1 = k1
	return rb
}
