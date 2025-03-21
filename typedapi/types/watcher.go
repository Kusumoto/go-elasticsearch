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

// Watcher type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/xpack/usage/types.ts#L442-L446
type Watcher struct {
	Available bool           `json:"available"`
	Count     Counter        `json:"count"`
	Enabled   bool           `json:"enabled"`
	Execution WatcherActions `json:"execution"`
	Watch     WatcherWatch   `json:"watch"`
}

// WatcherBuilder holds Watcher struct and provides a builder API.
type WatcherBuilder struct {
	v *Watcher
}

// NewWatcher provides a builder for the Watcher struct.
func NewWatcherBuilder() *WatcherBuilder {
	r := WatcherBuilder{
		&Watcher{},
	}

	return &r
}

// Build finalize the chain and returns the Watcher struct
func (rb *WatcherBuilder) Build() Watcher {
	return *rb.v
}

func (rb *WatcherBuilder) Available(available bool) *WatcherBuilder {
	rb.v.Available = available
	return rb
}

func (rb *WatcherBuilder) Count(count *CounterBuilder) *WatcherBuilder {
	v := count.Build()
	rb.v.Count = v
	return rb
}

func (rb *WatcherBuilder) Enabled(enabled bool) *WatcherBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *WatcherBuilder) Execution(execution *WatcherActionsBuilder) *WatcherBuilder {
	v := execution.Build()
	rb.v.Execution = v
	return rb
}

func (rb *WatcherBuilder) Watch(watch *WatcherWatchBuilder) *WatcherBuilder {
	v := watch.Build()
	rb.v.Watch = v
	return rb
}
