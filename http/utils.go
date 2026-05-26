// Copyright (c) 2023 IBM Corp.
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package http

import (
	"io"
	H "net/http"
	"net/url"
	"regexp"

	A "github.com/IBM/fp-go/array"
	E "github.com/IBM/fp-go/either"
	"github.com/IBM/fp-go/errors"
	F "github.com/IBM/fp-go/function"
	O "github.com/IBM/fp-go/option"
	P "github.com/IBM/fp-go/pair"
	R "github.com/IBM/fp-go/record/generic"
)

type (
	ParsedMediaType = P.Pair[string, map[string]string]

	HttpError struct {
		statusCode int
		headers    H.Header
		body       []byte
		url        *url.URL
	}
)

var (
	// mime type to check if a media type matches
	isJSONMimeType = regexp.MustCompile(`application/(?:\w+\+)?json`).MatchString
	// ValidateResponse validates an HTTP response and returns an [E.Either] if the response is not a success
	ValidateResponse = E.FromPredicate(isValidStatus, StatusCodeError)
	// alidateJsonContentTypeString parses a content type a validates that it is valid JSON
	validateJSONContentTypeString = F.Flow2(
		ParseMediaType,
		E.ChainFirst(F.Flow2(
			P.Head[string, map[string]string],
			E.FromPredicate(isJSONMimeType, errors.OnSome[string]("mimetype [%s] is not a valid JSON content type")),
		)),
	)
	// ValidateJSONResponse checks if an HTTP response is a valid JSON response
	ValidateJSONResponse = F.Flow2(
		E.Of[error, *H.Response],
		E.ChainFirst(F.Flow5(
			GetHeader,
			R.Lookup[H.Header](HeaderContentType),
			O.Chain(A.First[string]),
			E.FromOption[string](errors.OnNone("unable to access the [%s] header", HeaderContentType)),
			E.ChainFirst(validateJSONContentTypeString),
		)))
	// ValidateJsonResponse checks if an HTTP response is a valid JSON response
	//
	// Deprecated: use [ValidateJSONResponse] instead
	ValidateJsonResponse = ValidateJSONResponse
)

const (
	HeaderContentType = "Content-Type"
)

// ParseMediaType parses a media type into a tuple
func ParseMediaType(mediaType string) E.Either[error, ParsedMediaType] {
	_ = "STUB: not implemented"
	return nil
}

// Error fulfills the error interface
func (r *HttpError) Error() string { _ = "STUB: not implemented"; return "" }

func (r *HttpError) String() string { _ = "STUB: not implemented"; return "" }

func (r *HttpError) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (r *HttpError) Headers() H.Header { _ = "STUB: not implemented"; return *new(H.Header) }

func (r *HttpError) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (r *HttpError) Body() []byte { _ = "STUB: not implemented"; return nil }

func GetHeader(resp *H.Response) H.Header { _ = "STUB: not implemented"; return *new(H.Header) }

func GetBody(resp *H.Response) io.ReadCloser { _ = "STUB: not implemented"; return *new(io.ReadCloser) }

func isValidStatus(resp *H.Response) bool { _ = "STUB: not implemented"; return false }

// StatusCodeError creates an instance of [HttpError] filled with information from the response
func StatusCodeError(resp *H.Response) error {
	_ = "STUB: not implemented"
	// read the body
	return nil
}

// try to access body content

// return an error with comprehensive information
