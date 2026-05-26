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
	"net/http"

	RIOE "github.com/IBM/fp-go/context/readerioeither"
	F "github.com/IBM/fp-go/function"
	H "github.com/IBM/fp-go/http"
	IOE "github.com/IBM/fp-go/ioeither"
)

type (
	// Requester is a reader that constructs a request
	Requester = RIOE.ReaderIOEither[*http.Request]

	Client interface {
		// Do can send an HTTP request considering a context
		Do(Requester) RIOE.ReaderIOEither[*http.Response]
	}

	client struct {
		delegate *http.Client
		doIOE    func(*http.Request) IOE.IOEither[error, *http.Response]
	}
)

var (
	// MakeRequest is an eitherized version of [http.NewRequestWithContext]
	MakeRequest = RIOE.Eitherize3(http.NewRequestWithContext)
	makeRequest = F.Bind13of3(MakeRequest)

	// specialize
	MakeGetRequest = makeRequest("GET", nil)
)

func (client client) Do(req Requester) RIOE.ReaderIOEither[*http.Response] {
	_ = "STUB: not implemented"
	return nil
}

// MakeClient creates an HTTP client proxy
func MakeClient(httpClient *http.Client) Client { _ = "STUB: not implemented"; return *new(Client) }

// ReadFullResponse sends a request,  reads the response as a byte array and represents the result as a tuple
func ReadFullResponse(client Client) func(Requester) RIOE.ReaderIOEither[H.FullResponse] {
	_ = "STUB: not implemented"
	return nil
}

// ReadAll sends a request and reads the response as bytes
func ReadAll(client Client) func(Requester) RIOE.ReaderIOEither[[]byte] {
	_ = "STUB: not implemented"
	return nil
}

// ReadText sends a request, reads the response and represents the response as a text string
func ReadText(client Client) func(Requester) RIOE.ReaderIOEither[string] {
	_ = "STUB: not implemented"
	return nil
}

// ReadJson sends a request, reads the response and parses the response as JSON
//
// Deprecated: use [ReadJSON] instead
func ReadJson[A any](client Client) func(Requester) RIOE.ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}

func readJSON(client Client) func(Requester) RIOE.ReaderIOEither[[]byte] {
	_ = "STUB: not implemented"
	return nil
}

// ReadJSON sends a request, reads the response and parses the response as JSON
func ReadJSON[A any](client Client) func(Requester) RIOE.ReaderIOEither[A] {
	_ = "STUB: not implemented"
	return nil
}
