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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// NatRulesServer is a fake server for instances of the armnetwork.NatRulesClient type.
type NatRulesServer struct {
	// BeginCreateOrUpdate is the fake for method NatRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, natRuleParameters armnetwork.VPNGatewayNatRule, options *armnetwork.NatRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.NatRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method NatRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *armnetwork.NatRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.NatRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NatRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, gatewayName string, natRuleName string, options *armnetwork.NatRulesClientGetOptions) (resp azfake.Responder[armnetwork.NatRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByVPNGatewayPager is the fake for method NatRulesClient.NewListByVPNGatewayPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByVPNGatewayPager func(resourceGroupName string, gatewayName string, options *armnetwork.NatRulesClientListByVPNGatewayOptions) (resp azfake.PagerResponder[armnetwork.NatRulesClientListByVPNGatewayResponse])
}

// NewNatRulesServerTransport creates a new instance of NatRulesServerTransport with the provided implementation.
// The returned NatRulesServerTransport instance is connected to an instance of armnetwork.NatRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNatRulesServerTransport(srv *NatRulesServer) *NatRulesServerTransport {
	return &NatRulesServerTransport{
		srv:                      srv,
		beginCreateOrUpdate:      newTracker[azfake.PollerResponder[armnetwork.NatRulesClientCreateOrUpdateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armnetwork.NatRulesClientDeleteResponse]](),
		newListByVPNGatewayPager: newTracker[azfake.PagerResponder[armnetwork.NatRulesClientListByVPNGatewayResponse]](),
	}
}

// NatRulesServerTransport connects instances of armnetwork.NatRulesClient to instances of NatRulesServer.
// Don't use this type directly, use NewNatRulesServerTransport instead.
type NatRulesServerTransport struct {
	srv                      *NatRulesServer
	beginCreateOrUpdate      *tracker[azfake.PollerResponder[armnetwork.NatRulesClientCreateOrUpdateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armnetwork.NatRulesClientDeleteResponse]]
	newListByVPNGatewayPager *tracker[azfake.PagerResponder[armnetwork.NatRulesClientListByVPNGatewayResponse]]
}

// Do implements the policy.Transporter interface for NatRulesServerTransport.
func (n *NatRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NatRulesClient.BeginCreateOrUpdate":
		resp, err = n.dispatchBeginCreateOrUpdate(req)
	case "NatRulesClient.BeginDelete":
		resp, err = n.dispatchBeginDelete(req)
	case "NatRulesClient.Get":
		resp, err = n.dispatchGet(req)
	case "NatRulesClient.NewListByVPNGatewayPager":
		resp, err = n.dispatchNewListByVPNGatewayPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NatRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules/(?P<natRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VPNGatewayNatRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		natRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("natRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, gatewayNameParam, natRuleNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NatRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if n.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := n.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules/(?P<natRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		natRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("natRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := n.srv.BeginDelete(req.Context(), resourceGroupNameParam, gatewayNameParam, natRuleNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		n.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		n.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		n.beginDelete.remove(req)
	}

	return resp, nil
}

func (n *NatRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules/(?P<natRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
	if err != nil {
		return nil, err
	}
	natRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("natRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, gatewayNameParam, natRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VPNGatewayNatRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NatRulesServerTransport) dispatchNewListByVPNGatewayPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListByVPNGatewayPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByVPNGatewayPager not implemented")}
	}
	newListByVPNGatewayPager := n.newListByVPNGatewayPager.get(req)
	if newListByVPNGatewayPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/vpnGateways/(?P<gatewayName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/natRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		gatewayNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("gatewayName")])
		if err != nil {
			return nil, err
		}
		resp := n.srv.NewListByVPNGatewayPager(resourceGroupNameParam, gatewayNameParam, nil)
		newListByVPNGatewayPager = &resp
		n.newListByVPNGatewayPager.add(req, newListByVPNGatewayPager)
		server.PagerResponderInjectNextLinks(newListByVPNGatewayPager, req, func(page *armnetwork.NatRulesClientListByVPNGatewayResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByVPNGatewayPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListByVPNGatewayPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByVPNGatewayPager) {
		n.newListByVPNGatewayPager.remove(req)
	}
	return resp, nil
}
