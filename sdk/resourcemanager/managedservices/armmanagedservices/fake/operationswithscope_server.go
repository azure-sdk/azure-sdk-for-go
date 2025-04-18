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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managedservices/armmanagedservices"
	"net/http"
	"net/url"
	"regexp"
)

// OperationsWithScopeServer is a fake server for instances of the armmanagedservices.OperationsWithScopeClient type.
type OperationsWithScopeServer struct {
	// List is the fake for method OperationsWithScopeClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, scope string, options *armmanagedservices.OperationsWithScopeClientListOptions) (resp azfake.Responder[armmanagedservices.OperationsWithScopeClientListResponse], errResp azfake.ErrorResponder)
}

// NewOperationsWithScopeServerTransport creates a new instance of OperationsWithScopeServerTransport with the provided implementation.
// The returned OperationsWithScopeServerTransport instance is connected to an instance of armmanagedservices.OperationsWithScopeClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationsWithScopeServerTransport(srv *OperationsWithScopeServer) *OperationsWithScopeServerTransport {
	return &OperationsWithScopeServerTransport{srv: srv}
}

// OperationsWithScopeServerTransport connects instances of armmanagedservices.OperationsWithScopeClient to instances of OperationsWithScopeServer.
// Don't use this type directly, use NewOperationsWithScopeServerTransport instead.
type OperationsWithScopeServerTransport struct {
	srv *OperationsWithScopeServer
}

// Do implements the policy.Transporter interface for OperationsWithScopeServerTransport.
func (o *OperationsWithScopeServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OperationsWithScopeServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if operationsWithScopeServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = operationsWithScopeServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OperationsWithScopeClient.List":
				res.resp, res.err = o.dispatchList(req)
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

func (o *OperationsWithScopeServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if o.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ManagedServices/operations`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.List(req.Context(), scopeParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to OperationsWithScopeServerTransport
var operationsWithScopeServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
