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


// Provides statistics on operations happening in a data stream.
package datastreamsstats

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DataStreamsStats struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	name string
}

// NewDataStreamsStats type alias for index.
type NewDataStreamsStats func() *DataStreamsStats

// NewDataStreamsStatsFunc returns a new instance of DataStreamsStats with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDataStreamsStatsFunc(tp elastictransport.Interface) NewDataStreamsStats {
	return func() *DataStreamsStats {
		n := New(tp)

		return n
	}
}

// Provides statistics on operations happening in a data stream.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/data-streams.html
func New(tp elastictransport.Interface) *DataStreamsStats {
	r := &DataStreamsStats{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DataStreamsStats) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")
		path.WriteString("_stats")

		method = http.MethodGet
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.name))
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
func (r DataStreamsStats) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DataStreamsStats query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r DataStreamsStats) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DataStreamsStats headers map.
func (r *DataStreamsStats) Header(key, value string) *DataStreamsStats {
	r.headers.Set(key, value)

	return r
}

// Name A comma-separated list of data stream names; use `_all` or empty string to
// perform the operation on all data streams
// API Name: name
func (r *DataStreamsStats) Name(v string) *DataStreamsStats {
	r.paramSet |= nameMask
	r.name = v

	return r
}

// API name: expand_wildcards
func (r *DataStreamsStats) ExpandWildcards(value string) *DataStreamsStats {
	r.values.Set("expand_wildcards", value)

	return r
}
