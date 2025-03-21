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

// ScheduleTimeOfDay holds the union for the following types:
//
//	HourAndMinute
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/watcher/_types/Schedule.ts#L103-L108
type ScheduleTimeOfDay interface{}

// ScheduleTimeOfDayBuilder holds ScheduleTimeOfDay struct and provides a builder API.
type ScheduleTimeOfDayBuilder struct {
	v ScheduleTimeOfDay
}

// NewScheduleTimeOfDay provides a builder for the ScheduleTimeOfDay struct.
func NewScheduleTimeOfDayBuilder() *ScheduleTimeOfDayBuilder {
	return &ScheduleTimeOfDayBuilder{}
}

// Build finalize the chain and returns the ScheduleTimeOfDay struct
func (u *ScheduleTimeOfDayBuilder) Build() ScheduleTimeOfDay {
	return u.v
}

func (u *ScheduleTimeOfDayBuilder) HourAndMinute(hourandminute *HourAndMinuteBuilder) *ScheduleTimeOfDayBuilder {
	v := hourandminute.Build()
	u.v = &v
	return u
}

func (u *ScheduleTimeOfDayBuilder) String(string string) *ScheduleTimeOfDayBuilder {
	u.v = &string
	return u
}
