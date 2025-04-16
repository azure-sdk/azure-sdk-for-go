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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
	"net/http"
	"regexp"
)

// CheckNameAvailabilityServer is a fake server for instances of the armdevcenter.CheckNameAvailabilityClient type.
type CheckNameAvailabilityServer struct {
	// Execute is the fake for method CheckNameAvailabilityClient.Execute
	// HTTP status codes to indicate success: http.StatusOK
	Execute func(ctx context.Context, nameAvailabilityRequest armdevcenter.CheckNameAvailabilityRequest, options *armdevcenter.CheckNameAvailabilityClientExecuteOptions) (resp azfake.Responder[armdevcenter.CheckNameAvailabilityClientExecuteResponse], errResp azfake.ErrorResponder)
}

// NewCheckNameAvailabilityServerTransport creates a new instance of CheckNameAvailabilityServerTransport with the provided implementation.
// The returned CheckNameAvailabilityServerTransport instance is connected to an instance of armdevcenter.CheckNameAvailabilityClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCheckNameAvailabilityServerTransport(srv *CheckNameAvailabilityServer) *CheckNameAvailabilityServerTransport {
	return &CheckNameAvailabilityServerTransport{srv: srv}
}

// CheckNameAvailabilityServerTransport connects instances of armdevcenter.CheckNameAvailabilityClient to instances of CheckNameAvailabilityServer.
// Don't use this type directly, use NewCheckNameAvailabilityServerTransport instead.
type CheckNameAvailabilityServerTransport struct {
	srv *CheckNameAvailabilityServer
}

// Do implements the policy.Transporter interface for CheckNameAvailabilityServerTransport.
func (c *CheckNameAvailabilityServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CheckNameAvailabilityServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if checkNameAvailabilityServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = checkNameAvailabilityServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CheckNameAvailabilityClient.Execute":
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

func (c *CheckNameAvailabilityServerTransport) dispatchExecute(req *http.Request) (*http.Response, error) {
	if c.srv.Execute == nil {
		return nil, &nonRetriableError{errors.New("fake for method Execute not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevcenter.CheckNameAvailabilityRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Execute(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckNameAvailabilityResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CheckNameAvailabilityServerTransport
var checkNameAvailabilityServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
