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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresql/armpostgresqlflexibleservers/v5"
	"net/http"
	"net/url"
	"regexp"
)

// CheckNameAvailabilityWithLocationServer is a fake server for instances of the armpostgresqlflexibleservers.CheckNameAvailabilityWithLocationClient type.
type CheckNameAvailabilityWithLocationServer struct {
	// Execute is the fake for method CheckNameAvailabilityWithLocationClient.Execute
	// HTTP status codes to indicate success: http.StatusOK
	Execute func(ctx context.Context, locationName string, nameAvailabilityRequest armpostgresqlflexibleservers.CheckNameAvailabilityRequest, options *armpostgresqlflexibleservers.CheckNameAvailabilityWithLocationClientExecuteOptions) (resp azfake.Responder[armpostgresqlflexibleservers.CheckNameAvailabilityWithLocationClientExecuteResponse], errResp azfake.ErrorResponder)
}

// NewCheckNameAvailabilityWithLocationServerTransport creates a new instance of CheckNameAvailabilityWithLocationServerTransport with the provided implementation.
// The returned CheckNameAvailabilityWithLocationServerTransport instance is connected to an instance of armpostgresqlflexibleservers.CheckNameAvailabilityWithLocationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCheckNameAvailabilityWithLocationServerTransport(srv *CheckNameAvailabilityWithLocationServer) *CheckNameAvailabilityWithLocationServerTransport {
	return &CheckNameAvailabilityWithLocationServerTransport{srv: srv}
}

// CheckNameAvailabilityWithLocationServerTransport connects instances of armpostgresqlflexibleservers.CheckNameAvailabilityWithLocationClient to instances of CheckNameAvailabilityWithLocationServer.
// Don't use this type directly, use NewCheckNameAvailabilityWithLocationServerTransport instead.
type CheckNameAvailabilityWithLocationServerTransport struct {
	srv *CheckNameAvailabilityWithLocationServer
}

// Do implements the policy.Transporter interface for CheckNameAvailabilityWithLocationServerTransport.
func (c *CheckNameAvailabilityWithLocationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CheckNameAvailabilityWithLocationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if checkNameAvailabilityWithLocationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = checkNameAvailabilityWithLocationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CheckNameAvailabilityWithLocationClient.Execute":
				res.resp, res.err = c.dispatchExecute(req)
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

func (c *CheckNameAvailabilityWithLocationServerTransport) dispatchExecute(req *http.Request) (*http.Response, error) {
	if c.srv.Execute == nil {
		return nil, &nonRetriableError{errors.New("fake for method Execute not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpostgresqlflexibleservers.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Execute(req.Context(), locationNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NameAvailability, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CheckNameAvailabilityWithLocationServerTransport
var checkNameAvailabilityWithLocationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
