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


// Deletes an alias.
package deletealias

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
	indexMask = iota + 1

	nameMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteAlias struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	index string
	name  string
}

// NewDeleteAlias type alias for index.
type NewDeleteAlias func(index, name string) *DeleteAlias

// NewDeleteAliasFunc returns a new instance of DeleteAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteAliasFunc(tp elastictransport.Interface) NewDeleteAlias {
	return func(index, name string) *DeleteAlias {
		n := New(tp)

		n.Index(index)

		n.Name(name)

		return n
	}
}

// Deletes an alias.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html
func New(tp elastictransport.Interface) *DeleteAlias {
	r := &DeleteAlias{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_alias")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.name))

		method = http.MethodDelete
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.index))
		path.WriteString("/")
		path.WriteString("_aliases")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.name))

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r DeleteAlias) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteAlias query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r DeleteAlias) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the DeleteAlias headers map.
func (r *DeleteAlias) Header(key, value string) *DeleteAlias {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names (supports wildcards); use `_all` for
// all indices
// API Name: index
func (r *DeleteAlias) Index(v string) *DeleteAlias {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// Name A comma-separated list of aliases to delete (supports wildcards); use `_all`
// to delete all aliases for the specified indices.
// API Name: name
func (r *DeleteAlias) Name(v string) *DeleteAlias {
	r.paramSet |= nameMask
	r.name = v

	return r
}

// MasterTimeout Specify timeout for connection to master
// API name: master_timeout
func (r *DeleteAlias) MasterTimeout(value string) *DeleteAlias {
	r.values.Set("master_timeout", value)

	return r
}

// Timeout Explicit timestamp for the document
// API name: timeout
func (r *DeleteAlias) Timeout(value string) *DeleteAlias {
	r.values.Set("timeout", value)

	return r
}
