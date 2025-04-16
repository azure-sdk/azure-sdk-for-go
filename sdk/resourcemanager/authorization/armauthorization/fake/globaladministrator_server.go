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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
)

// GlobalAdministratorServer is a fake server for instances of the armauthorization.GlobalAdministratorClient type.
type GlobalAdministratorServer struct {
	// ElevateAccess is the fake for method GlobalAdministratorClient.ElevateAccess
	// HTTP status codes to indicate success: http.StatusOK
	ElevateAccess func(ctx context.Context, options *armauthorization.GlobalAdministratorClientElevateAccessOptions) (resp azfake.Responder[armauthorization.GlobalAdministratorClientElevateAccessResponse], errResp azfake.ErrorResponder)
}

// NewGlobalAdministratorServerTransport creates a new instance of GlobalAdministratorServerTransport with the provided implementation.
// The returned GlobalAdministratorServerTransport instance is connected to an instance of armauthorization.GlobalAdministratorClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGlobalAdministratorServerTransport(srv *GlobalAdministratorServer) *GlobalAdministratorServerTransport {
	return &GlobalAdministratorServerTransport{srv: srv}
}

// GlobalAdministratorServerTransport connects instances of armauthorization.GlobalAdministratorClient to instances of GlobalAdministratorServer.
// Don't use this type directly, use NewGlobalAdministratorServerTransport instead.
type GlobalAdministratorServerTransport struct {
	srv *GlobalAdministratorServer
}

// Do implements the policy.Transporter interface for GlobalAdministratorServerTransport.
func (g *GlobalAdministratorServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GlobalAdministratorServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if globalAdministratorServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = globalAdministratorServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GlobalAdministratorClient.ElevateAccess":
				res.resp, res.err = g.dispatchElevateAccess(req)
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

func (g *GlobalAdministratorServerTransport) dispatchElevateAccess(req *http.Request) (*http.Response, error) {
	if g.srv.ElevateAccess == nil {
		return nil, &nonRetriableError{errors.New("fake for method ElevateAccess not implemented")}
	}
	respr, errRespr := g.srv.ElevateAccess(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to GlobalAdministratorServerTransport
var globalAdministratorServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
