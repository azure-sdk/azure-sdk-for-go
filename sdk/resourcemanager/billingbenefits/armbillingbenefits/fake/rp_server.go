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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits/v3"
	"net/http"
)

// RPServer is a fake server for instances of the armbillingbenefits.RPClient type.
type RPServer struct {
	// ValidatePurchase is the fake for method RPClient.ValidatePurchase
	// HTTP status codes to indicate success: http.StatusOK
	ValidatePurchase func(ctx context.Context, body armbillingbenefits.SavingsPlanPurchaseValidateRequest, options *armbillingbenefits.RPClientValidatePurchaseOptions) (resp azfake.Responder[armbillingbenefits.RPClientValidatePurchaseResponse], errResp azfake.ErrorResponder)
}

// NewRPServerTransport creates a new instance of RPServerTransport with the provided implementation.
// The returned RPServerTransport instance is connected to an instance of armbillingbenefits.RPClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRPServerTransport(srv *RPServer) *RPServerTransport {
	return &RPServerTransport{srv: srv}
}

// RPServerTransport connects instances of armbillingbenefits.RPClient to instances of RPServer.
// Don't use this type directly, use NewRPServerTransport instead.
type RPServerTransport struct {
	srv *RPServer
}

// Do implements the policy.Transporter interface for RPServerTransport.
func (r *RPServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *RPServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if rpServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = rpServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "RPClient.ValidatePurchase":
				res.resp, res.err = r.dispatchValidatePurchase(req)
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

func (r *RPServerTransport) dispatchValidatePurchase(req *http.Request) (*http.Response, error) {
	if r.srv.ValidatePurchase == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidatePurchase not implemented")}
	}
	body, err := server.UnmarshalRequestAsJSON[armbillingbenefits.SavingsPlanPurchaseValidateRequest](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.ValidatePurchase(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SavingsPlanValidateResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to RPServerTransport
var rpServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
