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


// Downsample an index
package downsample

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexMask = iota + 1

	targetindexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Downsample struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *types.DownsampleConfig
	raw json.RawMessage

	paramSet int

	index       string
	targetindex string
}

// NewDownsample type alias for index.
type NewDownsample func(index, targetindex string) *Downsample

// NewDownsampleFunc returns a new instance of Downsample with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDownsampleFunc(tp elastictransport.Interface) NewDownsample {
	return func(index, targetindex string) *Downsample {
		n := New(tp)

		n.Index(index)

		n.TargetIndex(targetindex)

		return n
	}
}

// Downsample an index
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/xpack-rollup.html
func New(tp elastictransport.Interface) *Downsample {
	r := &Downsample{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Downsample) Raw(raw json.RawMessage) *Downsample {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Downsample) Request(req *types.DownsampleConfig) *Downsample {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Downsample) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Downsample: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|targetindexMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_downsample")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.targetindex))

		method = http.MethodPost
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
func (r Downsample) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Downsample query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the Downsample headers map.
func (r *Downsample) Header(key, value string) *Downsample {
	r.headers.Set(key, value)

	return r
}

// Index The index to downsample
// API Name: index
func (r *Downsample) Index(v string) *Downsample {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// TargetIndex The name of the target index to store downsampled data
// API Name: targetindex
func (r *Downsample) TargetIndex(v string) *Downsample {
	r.paramSet |= targetindexMask
	r.targetindex = v

	return r
}
