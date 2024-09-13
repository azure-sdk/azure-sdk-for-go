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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// AdminRulesServer is a fake server for instances of the armnetwork.AdminRulesClient type.
type AdminRulesServer struct {
	// CreateOrUpdate is the fake for method AdminRulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, adminRule armnetwork.AdminRule, options *armnetwork.AdminRulesClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.AdminRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AdminRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *armnetwork.AdminRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.AdminRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AdminRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *armnetwork.AdminRulesClientGetOptions) (resp azfake.Responder[armnetwork.AdminRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AdminRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *armnetwork.AdminRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.AdminRulesClientListResponse])
}

// NewAdminRulesServerTransport creates a new instance of AdminRulesServerTransport with the provided implementation.
// The returned AdminRulesServerTransport instance is connected to an instance of armnetwork.AdminRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAdminRulesServerTransport(srv *AdminRulesServer) *AdminRulesServerTransport {
	return &AdminRulesServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.AdminRulesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.AdminRulesClientListResponse]](),
	}
}

// AdminRulesServerTransport connects instances of armnetwork.AdminRulesClient to instances of AdminRulesServer.
// Don't use this type directly, use NewAdminRulesServerTransport instead.
type AdminRulesServerTransport struct {
	srv          *AdminRulesServer
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.AdminRulesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.AdminRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for AdminRulesServerTransport.
func (a *AdminRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AdminRulesClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "AdminRulesClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AdminRulesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AdminRulesClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AdminRulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAdminConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.AdminRule](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
	if err != nil {
		return nil, err
	}
	ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
	if err != nil {
		return nil, err
	}
	ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AdminRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdminRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAdminConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
		if err != nil {
			return nil, err
		}
		ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
		if err != nil {
			return nil, err
		}
		forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
		if err != nil {
			return nil, err
		}
		forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armnetwork.AdminRulesClientBeginDeleteOptions
		if forceParam != nil {
			options = &armnetwork.AdminRulesClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AdminRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAdminConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
	if err != nil {
		return nil, err
	}
	ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
	if err != nil {
		return nil, err
	}
	ruleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AdminRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdminRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityAdminConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		configurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationName")])
		if err != nil {
			return nil, err
		}
		ruleCollectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleCollectionName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.AdminRulesClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.AdminRulesClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := a.srv.NewListPager(resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, options)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.AdminRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		a.newListPager.remove(req)
	}
	return resp, nil
}
