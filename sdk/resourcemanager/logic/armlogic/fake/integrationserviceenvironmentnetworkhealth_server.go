// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic/v2"
	"net/http"
	"net/url"
	"regexp"
)

// IntegrationServiceEnvironmentNetworkHealthServer is a fake server for instances of the armlogic.IntegrationServiceEnvironmentNetworkHealthClient type.
type IntegrationServiceEnvironmentNetworkHealthServer struct {
	// Get is the fake for method IntegrationServiceEnvironmentNetworkHealthClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, integrationServiceEnvironmentName string, options *armlogic.IntegrationServiceEnvironmentNetworkHealthClientGetOptions) (resp azfake.Responder[armlogic.IntegrationServiceEnvironmentNetworkHealthClientGetResponse], errResp azfake.ErrorResponder)
}

// NewIntegrationServiceEnvironmentNetworkHealthServerTransport creates a new instance of IntegrationServiceEnvironmentNetworkHealthServerTransport with the provided implementation.
// The returned IntegrationServiceEnvironmentNetworkHealthServerTransport instance is connected to an instance of armlogic.IntegrationServiceEnvironmentNetworkHealthClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationServiceEnvironmentNetworkHealthServerTransport(srv *IntegrationServiceEnvironmentNetworkHealthServer) *IntegrationServiceEnvironmentNetworkHealthServerTransport {
	return &IntegrationServiceEnvironmentNetworkHealthServerTransport{srv: srv}
}

// IntegrationServiceEnvironmentNetworkHealthServerTransport connects instances of armlogic.IntegrationServiceEnvironmentNetworkHealthClient to instances of IntegrationServiceEnvironmentNetworkHealthServer.
// Don't use this type directly, use NewIntegrationServiceEnvironmentNetworkHealthServerTransport instead.
type IntegrationServiceEnvironmentNetworkHealthServerTransport struct {
	srv *IntegrationServiceEnvironmentNetworkHealthServer
}

// Do implements the policy.Transporter interface for IntegrationServiceEnvironmentNetworkHealthServerTransport.
func (i *IntegrationServiceEnvironmentNetworkHealthServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IntegrationServiceEnvironmentNetworkHealthServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if integrationServiceEnvironmentNetworkHealthServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = integrationServiceEnvironmentNetworkHealthServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IntegrationServiceEnvironmentNetworkHealthClient.Get":
				res.resp, res.err = i.dispatchGet(req)
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

func (i *IntegrationServiceEnvironmentNetworkHealthServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationServiceEnvironments/(?P<integrationServiceEnvironmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/network`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	integrationServiceEnvironmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationServiceEnvironmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), integrationServiceEnvironmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationServiceEnvironmentNetworkHealth, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to IntegrationServiceEnvironmentNetworkHealthServerTransport
var integrationServiceEnvironmentNetworkHealthServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
