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


// Retrieves usage information for transforms.
package gettransformstats

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	transformidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetTransformStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	transformid string
}

// NewGetTransformStats type alias for index.
type NewGetTransformStats func(transformid string) *GetTransformStats

// NewGetTransformStatsFunc returns a new instance of GetTransformStats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetTransformStatsFunc(tp elastictransport.Interface) NewGetTransformStats {
	return func(transformid string) *GetTransformStats {
		n := New(tp)

		n.TransformId(transformid)

		return n
	}
}

// Retrieves usage information for transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/get-transform-stats.html
func New(tp elastictransport.Interface) *GetTransformStats {
	r := &GetTransformStats{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetTransformStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == transformidMask:
		path.WriteString("/")
		path.WriteString("_transform")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.transformid))
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r GetTransformStats) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetTransformStats query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetTransformStats) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the GetTransformStats headers map.
func (r *GetTransformStats) Header(key, value string) *GetTransformStats {
	r.headers.Set(key, value)

	return r
}

// TransformId Identifier for the transform. It can be a transform identifier or a
// wildcard expression. You can get information for all transforms by using
// `_all`, by specifying `*` as the `<transform_id>`, or by omitting the
// `<transform_id>`.
// API Name: transformid
func (r *GetTransformStats) TransformId(v string) *GetTransformStats {
	r.paramSet |= transformidMask
	r.transformid = v

	return r
}

// AllowNoMatch Specifies what to do when the request:
//
// 1. Contains wildcard expressions and there are no transforms that match.
// 2. Contains the _all string or no identifiers and there are no matches.
// 3. Contains wildcard expressions and there are only partial matches.
//
// If this parameter is false, the request returns a 404 status code when
// there are no matches or only partial matches.
// API name: allow_no_match
func (r *GetTransformStats) AllowNoMatch(b bool) *GetTransformStats {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// From Skips the specified number of transforms.
// API name: from
func (r *GetTransformStats) From(value string) *GetTransformStats {
	r.values.Set("from", value)

	return r
}

// Size Specifies the maximum number of transforms to obtain.
// API name: size
func (r *GetTransformStats) Size(value string) *GetTransformStats {
	r.values.Set("size", value)

	return r
}
