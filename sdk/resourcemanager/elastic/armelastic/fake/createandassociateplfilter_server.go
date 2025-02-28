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

// CreateAndAssociatePLFilterServer is a fake server for instances of the armelastic.CreateAndAssociatePLFilterClient type.
type CreateAndAssociatePLFilterServer struct {
	// BeginCreate is the fake for method CreateAndAssociatePLFilterClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginCreate func(ctx context.Context, resourceGroupName string, monitorName string, options *armelastic.CreateAndAssociatePLFilterClientBeginCreateOptions) (resp azfake.PollerResponder[armelastic.CreateAndAssociatePLFilterClientCreateResponse], errResp azfake.ErrorResponder)
}

// NewCreateAndAssociatePLFilterServerTransport creates a new instance of CreateAndAssociatePLFilterServerTransport with the provided implementation.
// The returned CreateAndAssociatePLFilterServerTransport instance is connected to an instance of armelastic.CreateAndAssociatePLFilterClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCreateAndAssociatePLFilterServerTransport(srv *CreateAndAssociatePLFilterServer) *CreateAndAssociatePLFilterServerTransport {
	return &CreateAndAssociatePLFilterServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armelastic.CreateAndAssociatePLFilterClientCreateResponse]](),
	}
}

// CreateAndAssociatePLFilterServerTransport connects instances of armelastic.CreateAndAssociatePLFilterClient to instances of CreateAndAssociatePLFilterServer.
// Don't use this type directly, use NewCreateAndAssociatePLFilterServerTransport instead.
type CreateAndAssociatePLFilterServerTransport struct {
	srv         *CreateAndAssociatePLFilterServer
	beginCreate *tracker[azfake.PollerResponder[armelastic.CreateAndAssociatePLFilterClientCreateResponse]]
}

// Do implements the policy.Transporter interface for CreateAndAssociatePLFilterServerTransport.
func (c *CreateAndAssociatePLFilterServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CreateAndAssociatePLFilterServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if createAndAssociatePlFilterServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = createAndAssociatePlFilterServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CreateAndAssociatePLFilterClient.BeginCreate":
				res.resp, res.err = c.dispatchBeginCreate(req)
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

func (c *CreateAndAssociatePLFilterServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Elastic/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createAndAssociatePLFilter`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		nameUnescaped, err := url.QueryUnescape(qp.Get("name"))
		if err != nil {
			return nil, err
		}
		nameParam := getOptional(nameUnescaped)
		privateEndpointGUIDUnescaped, err := url.QueryUnescape(qp.Get("privateEndpointGuid"))
		if err != nil {
			return nil, err
		}
		privateEndpointGUIDParam := getOptional(privateEndpointGUIDUnescaped)
		privateEndpointNameUnescaped, err := url.QueryUnescape(qp.Get("privateEndpointName"))
		if err != nil {
			return nil, err
		}
		privateEndpointNameParam := getOptional(privateEndpointNameUnescaped)
		var options *armelastic.CreateAndAssociatePLFilterClientBeginCreateOptions
		if nameParam != nil || privateEndpointGUIDParam != nil || privateEndpointNameParam != nil {
			options = &armelastic.CreateAndAssociatePLFilterClientBeginCreateOptions{
				Name:                nameParam,
				PrivateEndpointGUID: privateEndpointGUIDParam,
				PrivateEndpointName: privateEndpointNameParam,
			}
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), resourceGroupNameParam, monitorNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to CreateAndAssociatePLFilterServerTransport
var createAndAssociatePlFilterServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
