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

// TrainedModelInferenceFeatureImportance type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/ml/_types/inference.ts#L404-L408
type TrainedModelInferenceFeatureImportance struct {
	Classes     []TrainedModelInferenceClassImportance `json:"classes,omitempty"`
	FeatureName string                                 `json:"feature_name"`
	Importance  *float64                               `json:"importance,omitempty"`
}

// TrainedModelInferenceFeatureImportanceBuilder holds TrainedModelInferenceFeatureImportance struct and provides a builder API.
type TrainedModelInferenceFeatureImportanceBuilder struct {
	v *TrainedModelInferenceFeatureImportance
}

// NewTrainedModelInferenceFeatureImportance provides a builder for the TrainedModelInferenceFeatureImportance struct.
func NewTrainedModelInferenceFeatureImportanceBuilder() *TrainedModelInferenceFeatureImportanceBuilder {
	r := TrainedModelInferenceFeatureImportanceBuilder{
		&TrainedModelInferenceFeatureImportance{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelInferenceFeatureImportance struct
func (rb *TrainedModelInferenceFeatureImportanceBuilder) Build() TrainedModelInferenceFeatureImportance {
	return *rb.v
}

func (rb *TrainedModelInferenceFeatureImportanceBuilder) Classes(classes []TrainedModelInferenceClassImportanceBuilder) *TrainedModelInferenceFeatureImportanceBuilder {
	tmp := make([]TrainedModelInferenceClassImportance, len(classes))
	for _, value := range classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.Classes = tmp
	return rb
}

func (rb *TrainedModelInferenceFeatureImportanceBuilder) FeatureName(featurename string) *TrainedModelInferenceFeatureImportanceBuilder {
	rb.v.FeatureName = featurename
	return rb
}

func (rb *TrainedModelInferenceFeatureImportanceBuilder) Importance(importance float64) *TrainedModelInferenceFeatureImportanceBuilder {
	rb.v.Importance = &importance
	return rb
}
