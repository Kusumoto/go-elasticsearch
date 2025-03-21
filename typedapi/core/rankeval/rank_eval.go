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


// Allows to evaluate the quality of ranked search results over a set of typical
// search queries
package rankeval

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RankEval struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req *Request
	raw json.RawMessage

	paramSet int

	index string
}

// NewRankEval type alias for index.
type NewRankEval func(index string) *RankEval

// NewRankEvalFunc returns a new instance of RankEval with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRankEvalFunc(tp elastictransport.Interface) NewRankEval {
	return func(index string) *RankEval {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Allows to evaluate the quality of ranked search results over a set of typical
// search queries
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-rank-eval.html
func New(tp elastictransport.Interface) *RankEval {
	r := &RankEval{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *RankEval) Raw(raw json.RawMessage) *RankEval {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *RankEval) Request(req *Request) *RankEval {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *RankEval) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if r.raw != nil {
		r.buf.Write(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for RankEval: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_rank_eval")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_rank_eval")

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
func (r RankEval) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the RankEval query execution: %w", err)
	}

	return res, nil
}

// Header set a key, value pair in the RankEval headers map.
func (r *RankEval) Header(key, value string) *RankEval {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and index aliases used to
// limit the request. Wildcard (`*`) expressions are supported.
// To target all data streams and indices in a cluster, omit this parameter or
// use `_all` or `*`.
// API Name: index
func (r *RankEval) Index(v string) *RankEval {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices. This behavior
// applies even if the request targets other open indices. For example, a
// request targeting `foo*,bar*` returns an error if an index starts with `foo`
// but no index starts with `bar`.
// API name: allow_no_indices
func (r *RankEval) AllowNoIndices(b bool) *RankEval {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *RankEval) ExpandWildcards(value string) *RankEval {
	r.values.Set("expand_wildcards", value)

	return r
}

// IgnoreUnavailable If `true`, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *RankEval) IgnoreUnavailable(b bool) *RankEval {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// SearchType Search operation type
// API name: search_type
func (r *RankEval) SearchType(value string) *RankEval {
	r.values.Set("search_type", value)

	return r
}
