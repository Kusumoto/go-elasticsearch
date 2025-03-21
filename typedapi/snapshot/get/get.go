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


// Returns information about a snapshot.
package get

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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/snapshotsort"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

const (
	repositoryMask = iota + 1

	snapshotMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Get struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	repository string
	snapshot   string
}

// NewGet type alias for index.
type NewGet func(repository, snapshot string) *Get

// NewGetFunc returns a new instance of Get with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetFunc(tp elastictransport.Interface) NewGet {
	return func(repository, snapshot string) *Get {
		n := New(tp)

		n.Repository(repository)

		n.Snapshot(snapshot)

		return n
	}
}

// Returns information about a snapshot.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/modules-snapshots.html
func New(tp elastictransport.Interface) *Get {
	r := &Get{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Get) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == repositoryMask|snapshotMask:
		path.WriteString("/")
		path.WriteString("_snapshot")
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.repository))
		path.WriteString("/")
		path.WriteString(url.PathEscape(r.snapshot))

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
func (r Get) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Get query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Get) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the Get headers map.
func (r *Get) Header(key, value string) *Get {
	r.headers.Set(key, value)

	return r
}

// Repository Comma-separated list of snapshot repository names used to limit the request.
// Wildcard (*) expressions are supported.
// API Name: repository
func (r *Get) Repository(v string) *Get {
	r.paramSet |= repositoryMask
	r.repository = v

	return r
}

// Snapshot Comma-separated list of snapshot names to retrieve. Also accepts wildcards
// (*).
// - To get information about all snapshots in a registered repository, use a
// wildcard (*) or _all.
// - To get information about any snapshots that are currently running, use
// _current.
// API Name: snapshot
func (r *Get) Snapshot(v string) *Get {
	r.paramSet |= snapshotMask
	r.snapshot = v

	return r
}

// IgnoreUnavailable If false, the request returns an error for any snapshots that are
// unavailable.
// API name: ignore_unavailable
func (r *Get) IgnoreUnavailable(b bool) *Get {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *Get) MasterTimeout(value string) *Get {
	r.values.Set("master_timeout", value)

	return r
}

// Verbose If true, returns additional information about each snapshot such as the
// version of Elasticsearch which took the snapshot, the start and end times of
// the snapshot, and the number of shards snapshotted.
// API name: verbose
func (r *Get) Verbose(b bool) *Get {
	r.values.Set("verbose", strconv.FormatBool(b))

	return r
}

// IndexDetails If true, returns additional information about each index in the snapshot
// comprising the number of shards in the index, the total size of the index in
// bytes, and the maximum number of segments per shard in the index. Defaults to
// false, meaning that this information is omitted.
// API name: index_details
func (r *Get) IndexDetails(b bool) *Get {
	r.values.Set("index_details", strconv.FormatBool(b))

	return r
}

// IndexNames If true, returns the name of each index in each snapshot.
// API name: index_names
func (r *Get) IndexNames(b bool) *Get {
	r.values.Set("index_names", strconv.FormatBool(b))

	return r
}

// IncludeRepository If true, returns the repository name in each snapshot.
// API name: include_repository
func (r *Get) IncludeRepository(b bool) *Get {
	r.values.Set("include_repository", strconv.FormatBool(b))

	return r
}

// Sort Allows setting a sort order for the result. Defaults to start_time, i.e.
// sorting by snapshot start time stamp.
// API name: sort
func (r *Get) Sort(enum snapshotsort.SnapshotSort) *Get {
	r.values.Set("sort", enum.String())

	return r
}

// Size Maximum number of snapshots to return. Defaults to 0 which means return all
// that match the request without limit.
// API name: size
func (r *Get) Size(i int) *Get {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// Order Sort order. Valid values are asc for ascending and desc for descending order.
// Defaults to asc, meaning ascending order.
// API name: order
func (r *Get) Order(enum sortorder.SortOrder) *Get {
	r.values.Set("order", enum.String())

	return r
}

// After Offset identifier to start pagination from as returned by the next field in
// the response body.
// API name: after
func (r *Get) After(value string) *Get {
	r.values.Set("after", value)

	return r
}

// Offset Numeric offset to start pagination from based on the snapshots matching this
// request. Using a non-zero value for this parameter is mutually exclusive with
// using the after parameter. Defaults to 0.
// API name: offset
func (r *Get) Offset(i int) *Get {
	r.values.Set("offset", strconv.Itoa(i))

	return r
}

// FromSortValue Value of the current sort column at which to start retrieval. Can either be a
// string snapshot- or repository name when sorting by snapshot or repository
// name, a millisecond time value or a number when sorting by index- or shard
// count.
// API name: from_sort_value
func (r *Get) FromSortValue(value string) *Get {
	r.values.Set("from_sort_value", value)

	return r
}

// SlmPolicyFilter Filter snapshots by a comma-separated list of SLM policy names that snapshots
// belong to. Also accepts wildcards (*) and combinations of wildcards followed
// by exclude patterns starting with -. To include snapshots not created by an
// SLM policy you can use the special pattern _none that will match all
// snapshots without an SLM policy.
// API name: slm_policy_filter
func (r *Get) SlmPolicyFilter(value string) *Get {
	r.values.Set("slm_policy_filter", value)

	return r
}
