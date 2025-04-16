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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// FirewallPolicyIdpsSignaturesFilterValuesServer is a fake server for instances of the armnetwork.FirewallPolicyIdpsSignaturesFilterValuesClient type.
type FirewallPolicyIdpsSignaturesFilterValuesServer struct {
	// List is the fake for method FirewallPolicyIdpsSignaturesFilterValuesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters armnetwork.SignatureOverridesFilterValuesQuery, options *armnetwork.FirewallPolicyIdpsSignaturesFilterValuesClientListOptions) (resp azfake.Responder[armnetwork.FirewallPolicyIdpsSignaturesFilterValuesClientListResponse], errResp azfake.ErrorResponder)
}

// NewFirewallPolicyIdpsSignaturesFilterValuesServerTransport creates a new instance of FirewallPolicyIdpsSignaturesFilterValuesServerTransport with the provided implementation.
// The returned FirewallPolicyIdpsSignaturesFilterValuesServerTransport instance is connected to an instance of armnetwork.FirewallPolicyIdpsSignaturesFilterValuesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFirewallPolicyIdpsSignaturesFilterValuesServerTransport(srv *FirewallPolicyIdpsSignaturesFilterValuesServer) *FirewallPolicyIdpsSignaturesFilterValuesServerTransport {
	return &FirewallPolicyIdpsSignaturesFilterValuesServerTransport{srv: srv}
}

// FirewallPolicyIdpsSignaturesFilterValuesServerTransport connects instances of armnetwork.FirewallPolicyIdpsSignaturesFilterValuesClient to instances of FirewallPolicyIdpsSignaturesFilterValuesServer.
// Don't use this type directly, use NewFirewallPolicyIdpsSignaturesFilterValuesServerTransport instead.
type FirewallPolicyIdpsSignaturesFilterValuesServerTransport struct {
	srv *FirewallPolicyIdpsSignaturesFilterValuesServer
}

// Do implements the policy.Transporter interface for FirewallPolicyIdpsSignaturesFilterValuesServerTransport.
func (f *FirewallPolicyIdpsSignaturesFilterValuesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FirewallPolicyIdpsSignaturesFilterValuesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if firewallPolicyIdpsSignaturesFilterValuesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = firewallPolicyIdpsSignaturesFilterValuesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FirewallPolicyIdpsSignaturesFilterValuesClient.List":
				res.resp, res.err = f.dispatchList(req)
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

func (f *FirewallPolicyIdpsSignaturesFilterValuesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if f.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listIdpsFilterOptions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.SignatureOverridesFilterValuesQuery](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.List(req.Context(), resourceGroupNameParam, firewallPolicyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SignatureOverridesFilterValuesResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to FirewallPolicyIdpsSignaturesFilterValuesServerTransport
var firewallPolicyIdpsSignaturesFilterValuesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
