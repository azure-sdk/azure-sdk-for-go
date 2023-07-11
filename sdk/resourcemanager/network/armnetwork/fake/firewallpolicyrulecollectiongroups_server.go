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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// FirewallPolicyRuleCollectionGroupsServer is a fake server for instances of the armnetwork.FirewallPolicyRuleCollectionGroupsClient type.
type FirewallPolicyRuleCollectionGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters armnetwork.FirewallPolicyRuleCollectionGroup, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdateDraft is the fake for method FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdateDraft
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdateDraft func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters armnetwork.FirewallPolicyRuleCollectionGroupDraft, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientBeginCreateOrUpdateDraftOptions) (resp azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateDraftResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FirewallPolicyRuleCollectionGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteDraft is the fake for method FirewallPolicyRuleCollectionGroupsClient.DeleteDraft
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteDraft func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters armnetwork.FirewallPolicyRuleCollectionGroup, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteDraftOptions) (resp azfake.Responder[armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteDraftResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FirewallPolicyRuleCollectionGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientGetOptions) (resp azfake.Responder[armnetwork.FirewallPolicyRuleCollectionGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// GetDraft is the fake for method FirewallPolicyRuleCollectionGroupsClient.GetDraft
	// HTTP status codes to indicate success: http.StatusOK
	GetDraft func(ctx context.Context, resourceGroupName string, firewallPolicyName string, ruleCollectionGroupName string, parameters armnetwork.FirewallPolicyRuleCollectionGroup, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientGetDraftOptions) (resp azfake.Responder[armnetwork.FirewallPolicyRuleCollectionGroupsClientGetDraftResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FirewallPolicyRuleCollectionGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, firewallPolicyName string, options *armnetwork.FirewallPolicyRuleCollectionGroupsClientListOptions) (resp azfake.PagerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse])
}

// NewFirewallPolicyRuleCollectionGroupsServerTransport creates a new instance of FirewallPolicyRuleCollectionGroupsServerTransport with the provided implementation.
// The returned FirewallPolicyRuleCollectionGroupsServerTransport instance is connected to an instance of armnetwork.FirewallPolicyRuleCollectionGroupsClient by way of the
// undefined.Transporter field.
func NewFirewallPolicyRuleCollectionGroupsServerTransport(srv *FirewallPolicyRuleCollectionGroupsServer) *FirewallPolicyRuleCollectionGroupsServerTransport {
	return &FirewallPolicyRuleCollectionGroupsServerTransport{srv: srv}
}

// FirewallPolicyRuleCollectionGroupsServerTransport connects instances of armnetwork.FirewallPolicyRuleCollectionGroupsClient to instances of FirewallPolicyRuleCollectionGroupsServer.
// Don't use this type directly, use NewFirewallPolicyRuleCollectionGroupsServerTransport instead.
type FirewallPolicyRuleCollectionGroupsServerTransport struct {
	srv                      *FirewallPolicyRuleCollectionGroupsServer
	beginCreateOrUpdate      *azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateResponse]
	beginCreateOrUpdateDraft *azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientCreateOrUpdateDraftResponse]
	beginDelete              *azfake.PollerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientDeleteResponse]
	newListPager             *azfake.PagerResponder[armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse]
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
	case "FirewallPolicyRuleCollectionGroupsClient.BeginCreateOrUpdateDraft":
		resp, err = f.dispatchBeginCreateOrUpdateDraft(req)
	case "FirewallPolicyRuleCollectionGroupsClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FirewallPolicyRuleCollectionGroupsClient.DeleteDraft":
		resp, err = f.dispatchDeleteDraft(req)
	case "FirewallPolicyRuleCollectionGroupsClient.Get":
		resp, err = f.dispatchGet(req)
	case "FirewallPolicyRuleCollectionGroupsClient.GetDraft":
		resp, err = f.dispatchGetDraft(req)
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
	if f.beginCreateOrUpdate == nil {
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
		f.beginCreateOrUpdate = &respr
	}

	resp, err := server.PollerResponderNext(f.beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginCreateOrUpdate) {
		f.beginCreateOrUpdate = nil
	}

	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchBeginCreateOrUpdateDraft(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdateDraft == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdateDraft not implemented")}
	}
	if f.beginCreateOrUpdateDraft == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups/(?P<ruleCollectionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/putDraft`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.FirewallPolicyRuleCollectionGroupDraft](req)
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
		respr, errRespr := f.srv.BeginCreateOrUpdateDraft(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, ruleCollectionGroupNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		f.beginCreateOrUpdateDraft = &respr
	}

	resp, err := server.PollerResponderNext(f.beginCreateOrUpdateDraft, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginCreateOrUpdateDraft) {
		f.beginCreateOrUpdateDraft = nil
	}

	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	if f.beginDelete == nil {
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
		f.beginDelete = &respr
	}

	resp, err := server.PollerResponderNext(f.beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(f.beginDelete) {
		f.beginDelete = nil
	}

	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchDeleteDraft(req *http.Request) (*http.Response, error) {
	if f.srv.DeleteDraft == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteDraft not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups/(?P<ruleCollectionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deleteDraft`
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
	respr, errRespr := f.srv.DeleteDraft(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, ruleCollectionGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallPolicyRuleCollectionGroup, req)
	if err != nil {
		return nil, err
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

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchGetDraft(req *http.Request) (*http.Response, error) {
	if f.srv.GetDraft == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDraft not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/firewallPolicies/(?P<firewallPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollectionGroups/(?P<ruleCollectionGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getDraft`
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
	respr, errRespr := f.srv.GetDraft(req.Context(), resourceGroupNameUnescaped, firewallPolicyNameUnescaped, ruleCollectionGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallPolicyRuleCollectionGroupDraft, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallPolicyRuleCollectionGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if f.newListPager == nil {
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
		f.newListPager = &resp
		server.PagerResponderInjectNextLinks(f.newListPager, req, func(page *armnetwork.FirewallPolicyRuleCollectionGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(f.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(f.newListPager) {
		f.newListPager = nil
	}
	return resp, nil
}
