//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
)

// FirewallPolicyDeploymentsServer is a fake server for instances of the armnetwork.FirewallPolicyDeploymentsClient type.
type FirewallPolicyDeploymentsServer struct {
	// BeginDeploy is the fake for method FirewallPolicyDeploymentsClient.BeginDeploy
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginDeploy func(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPolicyDeploymentsClientBeginDeployOptions) (resp azfake.PollerResponder[armnetwork.FirewallPolicyDeploymentsClientDeployResponse], errResp azfake.ErrorResponder)
}

// NewFirewallPolicyDeploymentsServerTransport creates a new instance of FirewallPolicyDeploymentsServerTransport with the provided implementation.
// The returned FirewallPolicyDeploymentsServerTransport instance is connected to an instance of armnetwork.FirewallPolicyDeploymentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFirewallPolicyDeploymentsServerTransport(srv *FirewallPolicyDeploymentsServer) *FirewallPolicyDeploymentsServerTransport {
	return &FirewallPolicyDeploymentsServerTransport{
		srv:         srv,
		beginDeploy: newTracker[azfake.PollerResponder[armnetwork.FirewallPolicyDeploymentsClientDeployResponse]](),
	}
}

// FirewallPolicyDeploymentsServerTransport connects instances of armnetwork.FirewallPolicyDeploymentsClient to instances of FirewallPolicyDeploymentsServer.
// Don't use this type directly, use NewFirewallPolicyDeploymentsServerTransport instead.
type FirewallPolicyDeploymentsServerTransport struct {
	srv         *FirewallPolicyDeploymentsServer
	beginDeploy *tracker[azfake.PollerResponder[armnetwork.FirewallPolicyDeploymentsClientDeployResponse]]
}

// Do implements the policy.Transporter interface for FirewallPolicyDeploymentsServerTransport.
func (f *FirewallPolicyDeploymentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FirewallPolicyDeploymentsClient.BeginDeploy":
		resp, err = f.dispatchBeginDeploy(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FirewallPolicyDeploymentsServerTransport) dispatchBeginDeploy(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDeploy == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeploy not implemented")}
	}
	beginDeploy := f.beginDeploy.get(req)
	if beginDeploy == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deploy`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDeploy(req.Context(), resourceGroupNameParam, firewallPolicyNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeploy = &respr
		f.beginDeploy.add(req, beginDeploy)
	}

	resp, err := server.PollerResponderNext(beginDeploy, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		f.beginDeploy.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeploy) {
		f.beginDeploy.remove(req)
	}

	return resp, nil
}
