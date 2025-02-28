// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elastic/armelastic/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AllTrafficFiltersServer is a fake server for instances of the armelastic.AllTrafficFiltersClient type.
type AllTrafficFiltersServer struct {
	// List is the fake for method AllTrafficFiltersClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, monitorName string, options *armelastic.AllTrafficFiltersClientListOptions) (resp azfake.Responder[armelastic.AllTrafficFiltersClientListResponse], errResp azfake.ErrorResponder)
}

// NewAllTrafficFiltersServerTransport creates a new instance of AllTrafficFiltersServerTransport with the provided implementation.
// The returned AllTrafficFiltersServerTransport instance is connected to an instance of armelastic.AllTrafficFiltersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAllTrafficFiltersServerTransport(srv *AllTrafficFiltersServer) *AllTrafficFiltersServerTransport {
	return &AllTrafficFiltersServerTransport{srv: srv}
}

// AllTrafficFiltersServerTransport connects instances of armelastic.AllTrafficFiltersClient to instances of AllTrafficFiltersServer.
// Don't use this type directly, use NewAllTrafficFiltersServerTransport instead.
type AllTrafficFiltersServerTransport struct {
	srv *AllTrafficFiltersServer
}

// Do implements the policy.Transporter interface for AllTrafficFiltersServerTransport.
func (a *AllTrafficFiltersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AllTrafficFiltersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if allTrafficFiltersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = allTrafficFiltersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AllTrafficFiltersClient.List":
				res.resp, res.err = a.dispatchList(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (a *AllTrafficFiltersServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if a.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Elastic/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listAllTrafficFilters`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.List(req.Context(), resourceGroupNameParam, monitorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TrafficFilterResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AllTrafficFiltersServerTransport
var allTrafficFiltersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
