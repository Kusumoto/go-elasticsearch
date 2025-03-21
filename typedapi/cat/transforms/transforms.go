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


// Gets configuration and usage information about transforms.
package transforms

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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

const (
	transformidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Transforms struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	transformid string
}

// NewTransforms type alias for index.
type NewTransforms func() *Transforms

// NewTransformsFunc returns a new instance of Transforms with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTransformsFunc(tp elastictransport.Interface) NewTransforms {
	return func() *Transforms {
		n := New(tp)

		return n
	}
}

// Gets configuration and usage information about transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/cat-transforms.html
func New(tp elastictransport.Interface) *Transforms {
	r := &Transforms{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Transforms) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("transforms")

		method = http.MethodGet
	case r.paramSet == transformidMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("transforms")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.transformid))

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
func (r Transforms) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Transforms query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Transforms) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the Transforms headers map.
func (r *Transforms) Header(key, value string) *Transforms {
	r.headers.Set(key, value)

	return r
}

// TransformId The id of the transform for which to get stats. '_all' or '*' implies all
// transforms
// API Name: transformid
func (r *Transforms) TransformId(v string) *Transforms {
	r.paramSet |= transformidMask
	r.transformid = v

	return r
}

// AllowNoMatch Whether to ignore if a wildcard expression matches no transforms. (This
// includes `_all` string or when no transforms have been specified)
// API name: allow_no_match
func (r *Transforms) AllowNoMatch(b bool) *Transforms {
	r.values.Set("allow_no_match", strconv.FormatBool(b))

	return r
}

// From skips a number of transform configs, defaults to 0
// API name: from
func (r *Transforms) From(i int) *Transforms {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// H Comma-separated list of column names to display.
// API name: h
func (r *Transforms) H(value string) *Transforms {
	r.values.Set("h", value)

	return r
}

// S Comma-separated list of column names or column aliases used to sort the
// response.
// API name: s
func (r *Transforms) S(value string) *Transforms {
	r.values.Set("s", value)

	return r
}

// Time Unit used to display time values.
// API name: time
func (r *Transforms) Time(enum timeunit.TimeUnit) *Transforms {
	r.values.Set("time", enum.String())

	return r
}

// Size specifies a max number of transforms to get, defaults to 100
// API name: size
func (r *Transforms) Size(i int) *Transforms {
	r.values.Set("size", strconv.Itoa(i))

	return r
}
