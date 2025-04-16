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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription"
	"net/http"
	"net/url"
	"regexp"
)

// BillingAccountServer is a fake server for instances of the armsubscription.BillingAccountClient type.
type BillingAccountServer struct {
	// GetPolicy is the fake for method BillingAccountClient.GetPolicy
	// HTTP status codes to indicate success: http.StatusOK
	GetPolicy func(ctx context.Context, billingAccountID string, options *armsubscription.BillingAccountClientGetPolicyOptions) (resp azfake.Responder[armsubscription.BillingAccountClientGetPolicyResponse], errResp azfake.ErrorResponder)
}

// NewBillingAccountServerTransport creates a new instance of BillingAccountServerTransport with the provided implementation.
// The returned BillingAccountServerTransport instance is connected to an instance of armsubscription.BillingAccountClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBillingAccountServerTransport(srv *BillingAccountServer) *BillingAccountServerTransport {
	return &BillingAccountServerTransport{srv: srv}
}

// BillingAccountServerTransport connects instances of armsubscription.BillingAccountClient to instances of BillingAccountServer.
// Don't use this type directly, use NewBillingAccountServerTransport instead.
type BillingAccountServerTransport struct {
	srv *BillingAccountServer
}

// Do implements the policy.Transporter interface for BillingAccountServerTransport.
func (b *BillingAccountServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BillingAccountServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if billingAccountServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = billingAccountServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BillingAccountClient.GetPolicy":
				res.resp, res.err = b.dispatchGetPolicy(req)
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

func (b *BillingAccountServerTransport) dispatchGetPolicy(req *http.Request) (*http.Response, error) {
	if b.srv.GetPolicy == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPolicy not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Billing/billingAccounts/(?P<billingAccountId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Subscription/policies/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	billingAccountIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("billingAccountId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.GetPolicy(req.Context(), billingAccountIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BillingAccountPoliciesResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BillingAccountServerTransport
var billingAccountServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
