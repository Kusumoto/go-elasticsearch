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


// Deletes expired and unused machine learning data.
package deleteexpireddata

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
)

const (
	jobidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteExpiredData struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	jobid string
}

// NewDeleteExpiredData type alias for index.
type NewDeleteExpiredData func() *DeleteExpiredData

// NewDeleteExpiredDataFunc returns a new instance of DeleteExpiredData with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteExpiredDataFunc(tp elastictransport.Interface) NewDeleteExpiredData {
	return func() *DeleteExpiredData {
		n := New(tp)

		return n
	}
}

// Deletes expired and unused machine learning data.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/ml-delete-expired-data.html
func New(tp elastictransport.Interface) *DeleteExpiredData {
	r := &DeleteExpiredData{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *DeleteExpiredData) Raw(raw json.RawMessage) *DeleteExpiredData {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *DeleteExpiredData) Request(req *Request) *DeleteExpiredData {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteExpiredData) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for DeleteExpiredData: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == jobidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("_delete_expired_data")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.jobid))

		method = http.MethodDelete
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("_delete_expired_data")

		method = http.MethodDelete
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
func (r DeleteExpiredData) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteExpiredData query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the DeleteExpiredData headers map.
func (r *DeleteExpiredData) Header(key, value string) *DeleteExpiredData {
	r.headers.Set(key, value)

	return r
}

// JobId Identifier for an anomaly detection job. It can be a job identifier, a
// group name, or a wildcard expression.
// API Name: jobid
func (r *DeleteExpiredData) JobId(v string) *DeleteExpiredData {
	r.paramSet |= jobidMask
	r.jobid = v

	return r
}

// RequestsPerSecond The desired requests per second for the deletion processes. The default
// behavior is no throttling.
// API name: requests_per_second
func (r *DeleteExpiredData) RequestsPerSecond(value string) *DeleteExpiredData {
	r.values.Set("requests_per_second", value)

	return r
}

// Timeout How long can the underlying delete processes run until they are canceled.
// API name: timeout
func (r *DeleteExpiredData) Timeout(value string) *DeleteExpiredData {
	r.values.Set("timeout", value)

	return r
}
