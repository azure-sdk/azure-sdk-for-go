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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosdbmongocluster/armcosmosdbmongocluster"
	"net/http"
	"net/url"
	"regexp"
)

// FirewallRulesServer is a fake server for instances of the armcosmosdbmongocluster.FirewallRulesClient type.
type FirewallRulesServer struct {
	// BeginCreateOrUpdate is the fake for method FirewallRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, mongoClusterName string, firewallRuleName string, resource armcosmosdbmongocluster.FirewallRule, options *armcosmosdbmongocluster.FirewallRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcosmosdbmongocluster.FirewallRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FirewallRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, mongoClusterName string, firewallRuleName string, options *armcosmosdbmongocluster.FirewallRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcosmosdbmongocluster.FirewallRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FirewallRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, mongoClusterName string, firewallRuleName string, options *armcosmosdbmongocluster.FirewallRulesClientGetOptions) (resp azfake.Responder[armcosmosdbmongocluster.FirewallRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMongoClusterPager is the fake for method FirewallRulesClient.NewListByMongoClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMongoClusterPager func(resourceGroupName string, mongoClusterName string, options *armcosmosdbmongocluster.FirewallRulesClientListByMongoClusterOptions) (resp azfake.PagerResponder[armcosmosdbmongocluster.FirewallRulesClientListByMongoClusterResponse])
}

// NewFirewallRulesServerTransport creates a new instance of FirewallRulesServerTransport with the provided implementation.
// The returned FirewallRulesServerTransport instance is connected to an instance of armcosmosdbmongocluster.FirewallRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFirewallRulesServerTransport(srv *FirewallRulesServer) *FirewallRulesServerTransport {
	return &FirewallRulesServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armcosmosdbmongocluster.FirewallRulesClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armcosmosdbmongocluster.FirewallRulesClientDeleteResponse]](),
		newListByMongoClusterPager: newTracker[azfake.PagerResponder[armcosmosdbmongocluster.FirewallRulesClientListByMongoClusterResponse]](),
	}
}

// FirewallRulesServerTransport connects instances of armcosmosdbmongocluster.FirewallRulesClient to instances of FirewallRulesServer.
// Don't use this type directly, use NewFirewallRulesServerTransport instead.
type FirewallRulesServerTransport struct {
	srv                        *FirewallRulesServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armcosmosdbmongocluster.FirewallRulesClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armcosmosdbmongocluster.FirewallRulesClientDeleteResponse]]
	newListByMongoClusterPager *tracker[azfake.PagerResponder[armcosmosdbmongocluster.FirewallRulesClientListByMongoClusterResponse]]
}

// Do implements the policy.Transporter interface for FirewallRulesServerTransport.
func (f *FirewallRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FirewallRulesClient.BeginCreateOrUpdate":
		resp, err = f.dispatchBeginCreateOrUpdate(req)
	case "FirewallRulesClient.BeginDelete":
		resp, err = f.dispatchBeginDelete(req)
	case "FirewallRulesClient.Get":
		resp, err = f.dispatchGet(req)
	case "FirewallRulesClient.NewListByMongoClusterPager":
		resp, err = f.dispatchNewListByMongoClusterPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FirewallRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/firewallRules/(?P<firewallRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcosmosdbmongocluster.FirewallRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		firewallRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, mongoClusterNameParam, firewallRuleNameParam, body, nil)
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

func (f *FirewallRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/firewallRules/(?P<firewallRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		firewallRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallRuleName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, mongoClusterNameParam, firewallRuleNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FirewallRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/firewallRules/(?P<firewallRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
	if err != nil {
		return nil, err
	}
	firewallRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("firewallRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, mongoClusterNameParam, firewallRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FirewallRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FirewallRulesServerTransport) dispatchNewListByMongoClusterPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByMongoClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMongoClusterPager not implemented")}
	}
	newListByMongoClusterPager := f.newListByMongoClusterPager.get(req)
	if newListByMongoClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/firewallRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListByMongoClusterPager(resourceGroupNameParam, mongoClusterNameParam, nil)
		newListByMongoClusterPager = &resp
		f.newListByMongoClusterPager.add(req, newListByMongoClusterPager)
		server.PagerResponderInjectNextLinks(newListByMongoClusterPager, req, func(page *armcosmosdbmongocluster.FirewallRulesClientListByMongoClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMongoClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByMongoClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMongoClusterPager) {
		f.newListByMongoClusterPager.remove(req)
	}
	return resp, nil
}
