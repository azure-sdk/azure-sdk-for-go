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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth"
	"net/http"
)

// OperationsServer is a fake server for instances of the armresourcehealth.OperationsClient type.
type OperationsServer struct {
	// List is the fake for method OperationsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, options *armresourcehealth.OperationsClientListOptions) (resp azfake.Responder[armresourcehealth.OperationsClientListResponse], errResp azfake.ErrorResponder)
}

// NewOperationsServerTransport creates a new instance of OperationsServerTransport with the provided implementation.
// The returned OperationsServerTransport instance is connected to an instance of armresourcehealth.OperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOperationsServerTransport(srv *OperationsServer) *OperationsServerTransport {
	return &OperationsServerTransport{srv: srv}
}

// OperationsServerTransport connects instances of armresourcehealth.OperationsClient to instances of OperationsServer.
// Don't use this type directly, use NewOperationsServerTransport instead.
type OperationsServerTransport struct {
	srv *OperationsServer
}

// Do implements the policy.Transporter interface for OperationsServerTransport.
func (o *OperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return o.dispatchToMethodFake(req, method)
}

func (o *OperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if operationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = operationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "OperationsClient.List":
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

func (o *OperationsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if o.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	respr, errRespr := o.srv.List(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OperationListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to OperationsServerTransport
var operationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
