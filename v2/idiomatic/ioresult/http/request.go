// Copyright (c) 2023 - 2025 IBM Corp.
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
	"net/http"

	F "github.com/IBM/fp-go/v2/function"
	H "github.com/IBM/fp-go/v2/http"
	"github.com/IBM/fp-go/v2/idiomatic/ioresult"
)

type (
	client struct {
		delegate *http.Client
		doIOE    Kleisli[*http.Request, *http.Response]
	}
)

var (
	// MakeRequest is an eitherized version of [http.NewRequest]
	MakeRequest = ioresult.Eitherize3(http.NewRequest)
	makeRequest = F.Bind13of3(MakeRequest)

	// specialize
	MakeGetRequest = makeRequest("GET", nil)
)

// MakeBodyRequest creates a request that carries a body
func MakeBodyRequest(method string, body IOResult[[]byte]) Kleisli[string, *http.Request] {
	_ = "STUB: not implemented"
	return nil
}

func (client client) Do(req Requester) IOResult[*http.Response] {
	_ = "STUB: not implemented"
	return nil
}

func MakeClient(httpClient *http.Client) Client { _ = "STUB: not implemented"; return *new(Client) }

// ReadFullResponse sends a request,  reads the response as a byte array and represents the result as a tuple
func ReadFullResponse(client Client) Operator[*http.Request, H.FullResponse] {
	_ = "STUB: not implemented"
	return nil
}

// var x R.Reader[*http.Response, IOResult[[]byte]] = F.Flow3(
// 	H.GetBody,
// 	ioresult.Of,
// 	IOEF.ReadAll,
// )

// ReadAll sends a request and reads the response as bytes
func ReadAll(client Client) Operator[*http.Request, []byte] { _ = "STUB: not implemented"; return nil }

// ReadText sends a request, reads the response and represents the response as a text string
func ReadText(client Client) Operator[*http.Request, string] { _ = "STUB: not implemented"; return nil }

// readJSON sends a request, reads the response and parses the response as a []byte
func readJSON(client Client) Operator[*http.Request, []byte] { _ = "STUB: not implemented"; return nil }

// ReadJSON sends a request, reads the response and parses the response as JSON
func ReadJSON[A any](client Client) Operator[*http.Request, A] {
	_ = "STUB: not implemented"
	return nil
}
