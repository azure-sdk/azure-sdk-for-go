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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork"
	"net/http"
	"net/url"
	"regexp"
)

// FirewallPolicyRuleCollectionGroupsServer is a fake server for instances of the armnetwork.FirewallPolicyRuleCollectionGroupsClient type.
type FirewallPolicyRuleCollectionGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters armnetwork.FirewallPolicyRuleCollectionGroup, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FirewallPolicyRuleCollectionGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FirewallPolicyRuleCollectionGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientGetOptions) (resp azfake.Responder[armnetwork.FirewallPolicyRuleCollectionGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FirewallPolicyRuleCollectionGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientListOptions) (resp azfake.PagerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse])
}

// NewFirewallPolicyRuleCollectionGroupsServerTransport creates a new instance of FirewallPolicyRuleCollectionGroupsServerTransport with the provided implementation.
// The returned FirewallPolicyRuleCollectionGroupsServerTransport instance is connected to an instance of armnetwork.FirewallPolicyRuleCollectionGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFirewallPolicyRuleCollectionGroupsServerTransport(srv *FirewallPolicyRuleCollectionGroupsServer) *FirewallPolicyRuleCollectionGroupsServerTransport {
	return &FirewallPolicyRuleCollectionGroupsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse]](),
	}
}

// FirewallPolicyRuleCollectionGroupsServerTransport connects instances of armnetwork.FirewallPolicyRuleCollectionGroupsClient to instances of FirewallPolicyRuleCollectionGroupsServer.
// Don't use this type directly, use NewFirewallPolicyRuleCollectionGroupsServerTransport instead.
type FirewallPolicyRuleCollectionGroupsServerTransport struct {
	srv                 *FirewallPolicyRuleCollectionGroupsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse]]
}

// Do implements the policy.Transporter interface for FirewallPolicyRuleCollectionGroupsServerTransport.
func (f *FirewallPolicyRuleCollectionGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FirewallPolicyRuleCollectionGroupsClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FirewallPolicyRuleCollectionGroupsClient.Get":
		resp, err = f.dispatchGet(req)
	case "FirewallPolicyRuleCollectionGroupsClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups/(?P<ruleCollectionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.FirewallPolicyRuleCollectionGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, ruleCollectionGroupNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups/(?P<ruleCollectionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, ruleCollectionGroupNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups/(?P<ruleCollectionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
	if err != nil {
		return nil, err
	}
	ruleCollectionGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, ruleCollectionGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallPolicyRuleCollectionGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		firewallPolicyNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("firewallPolicyName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListPager(resourceGroupNameUnescaped, firewallPolicyNameUnescaped, nil)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}
