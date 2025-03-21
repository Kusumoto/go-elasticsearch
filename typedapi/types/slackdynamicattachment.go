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

// SlackDynamicAttachment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/watcher/_types/Actions.ts#L125-L128
type SlackDynamicAttachment struct {
	AttachmentTemplate SlackAttachment `json:"attachment_template"`
	ListPath           string          `json:"list_path"`
}

// SlackDynamicAttachmentBuilder holds SlackDynamicAttachment struct and provides a builder API.
type SlackDynamicAttachmentBuilder struct {
	v *SlackDynamicAttachment
}

// NewSlackDynamicAttachment provides a builder for the SlackDynamicAttachment struct.
func NewSlackDynamicAttachmentBuilder() *SlackDynamicAttachmentBuilder {
	r := SlackDynamicAttachmentBuilder{
		&SlackDynamicAttachment{},
	}

	return &r
}

// Build finalize the chain and returns the SlackDynamicAttachment struct
func (rb *SlackDynamicAttachmentBuilder) Build() SlackDynamicAttachment {
	return *rb.v
}

func (rb *SlackDynamicAttachmentBuilder) AttachmentTemplate(attachmenttemplate *SlackAttachmentBuilder) *SlackDynamicAttachmentBuilder {
	v := attachmenttemplate.Build()
	rb.v.AttachmentTemplate = v
	return rb
}

func (rb *SlackDynamicAttachmentBuilder) ListPath(listpath string) *SlackDynamicAttachmentBuilder {
	rb.v.ListPath = listpath
	return rb
}
