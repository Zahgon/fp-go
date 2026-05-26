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
	"github.com/IBM/fp-go/v2/ioeither"
)

type (
	// Requester is a reader that constructs a request
	Requester = ioeither.IOEither[error, *http.Request]

	Client interface {
		Do(Requester) ioeither.IOEither[error, *http.Response]
	}

	client struct {
		delegate *http.Client
		doIOE    Kleisli[error, *http.Request, *http.Response]
	}
)

var (
	// MakeRequest is an eitherized version of [http.NewRequest]
	MakeRequest = ioeither.Eitherize3(http.NewRequest)
	makeRequest = F.Bind13of3(MakeRequest)

	// specialize
	MakeGetRequest = makeRequest("GET", nil)
)

// MakeBodyRequest creates a request that carries a body
func MakeBodyRequest(method string, body ioeither.IOEither[error, []byte]) Kleisli[error, string, *http.Request] {
	_ = "STUB: not implemented"
	return nil
}

func (client client) Do(req Requester) ioeither.IOEither[error, *http.Response] {
	_ = "STUB: not implemented"
	return nil
}

func MakeClient(httpClient *http.Client) Client { _ = "STUB: not implemented"; return *new(Client) }

// ReadFullResponse sends a request,  reads the response as a byte array and represents the result as a tuple
func ReadFullResponse(client Client) Kleisli[error, Requester, H.FullResponse] {
	_ = "STUB: not implemented"
	return nil
}

// ReadAll sends a request and reads the response as bytes
func ReadAll(client Client) Kleisli[error, Requester, []byte] {
	_ = "STUB: not implemented"
	return nil
}

// ReadText sends a request, reads the response and represents the response as a text string
func ReadText(client Client) Kleisli[error, Requester, string] {
	_ = "STUB: not implemented"
	return nil
}

// ReadJson sends a request, reads the response and parses the response as JSON
//
// Deprecated: use [ReadJSON] instead
func ReadJson[A any](client Client) Kleisli[error, Requester, A] {
	_ = "STUB: not implemented"
	return nil

	// readJSON sends a request, reads the response and parses the response as a []byte
}

func readJSON(client Client) Kleisli[error, Requester, []byte] {
	_ = "STUB: not implemented"
	return nil
}

// ReadJSON sends a request, reads the response and parses the response as JSON
func ReadJSON[A any](client Client) Kleisli[error, Requester, A] {
	_ = "STUB: not implemented"
	return nil
}
