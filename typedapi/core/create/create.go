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


// Creates a new document in the index.
//
// Returns a 409 response when a document with a same ID already exists in the
// index.
package create

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

const (
	idMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Create struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req interface{}
	raw json.RawMessage

	paramSet int

	id    string
	index string
}

// NewCreate type alias for index.
type NewCreate func(index, id string) *Create

// NewCreateFunc returns a new instance of Create with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	return func(index, id string) *Create {
		n := New(tp)

		n.Id(id)

		n.Index(index)

		return n
	}
}

// Creates a new document in the index.
//
// Returns a 409 response when a document with a same ID already exists in the
// index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-index_.html
func New(tp elastictransport.Interface) *Create {
	r := &Create{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Create) Raw(raw json.RawMessage) *Create {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Create) Request(req interface{}) *Create {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Create: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|idMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_create")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.id))

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r Create) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Create query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Create headers map.
func (r *Create) Header(key, value string) *Create {
	r.headers.Set(key, value)

	return r
}

// Id Document ID
// API Name: id
func (r *Create) Id(v string) *Create {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Index The name of the index
// API Name: index
func (r *Create) Index(v string) *Create {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Pipeline The pipeline id to preprocess incoming documents with
// API name: pipeline
func (r *Create) Pipeline(value string) *Create {
	r.values.Set("pipeline", value)

	return r
}

// Refresh If `true` then refresh the affected shards to make this operation visible to
// search, if `wait_for` then wait for a refresh to make this operation visible
// to search, if `false` (the default) then do nothing with refreshes.
// API name: refresh
func (r *Create) Refresh(enum refresh.Refresh) *Create {
	r.values.Set("refresh", enum.String())

	return r
}

// Routing Specific routing value
// API name: routing
func (r *Create) Routing(value string) *Create {
	r.values.Set("routing", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Create) Timeout(value string) *Create {
	r.values.Set("timeout", value)

	return r
}

// Version Explicit version number for concurrency control
// API name: version
func (r *Create) Version(value string) *Create {
	r.values.Set("version", value)

	return r
}

// VersionType Specific version type
// API name: version_type
func (r *Create) VersionType(enum versiontype.VersionType) *Create {
	r.values.Set("version_type", enum.String())

	return r
}

// WaitForActiveShards Sets the number of shard copies that must be active before proceeding with
// the index operation. Defaults to 1, meaning the primary shard only. Set to
// `all` for all shard copies, otherwise set to any non-negative value less than
// or equal to the total number of copies for the shard (number of replicas + 1)
// API name: wait_for_active_shards
func (r *Create) WaitForActiveShards(value string) *Create {
	r.values.Set("wait_for_active_shards", value)

	return r
}
