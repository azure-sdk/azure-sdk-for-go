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

// SecurityUserRulesServer is a fake server for instances of the armnetwork.SecurityUserRulesClient type.
type SecurityUserRulesServer struct {
	// CreateOrUpdate is the fake for method SecurityUserRulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, securityUserRule armnetwork.SecurityUserRule, options *armnetwork.SecurityUserRulesClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.SecurityUserRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SecurityUserRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *armnetwork.SecurityUserRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.SecurityUserRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SecurityUserRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, ruleName string, options *armnetwork.SecurityUserRulesClientGetOptions) (resp azfake.Responder[armnetwork.SecurityUserRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SecurityUserRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, configurationName string, ruleCollectionName string, options *armnetwork.SecurityUserRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.SecurityUserRulesClientListResponse])
}

// NewSecurityUserRulesServerTransport creates a new instance of SecurityUserRulesServerTransport with the provided implementation.
// The returned SecurityUserRulesServerTransport instance is connected to an instance of armnetwork.SecurityUserRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSecurityUserRulesServerTransport(srv *SecurityUserRulesServer) *SecurityUserRulesServerTransport {
	return &SecurityUserRulesServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.SecurityUserRulesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.SecurityUserRulesClientListResponse]](),
	}
}

// SecurityUserRulesServerTransport connects instances of armnetwork.SecurityUserRulesClient to instances of SecurityUserRulesServer.
// Don't use this type directly, use NewSecurityUserRulesServerTransport instead.
type SecurityUserRulesServerTransport struct {
	srv          *SecurityUserRulesServer
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.SecurityUserRulesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.SecurityUserRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for SecurityUserRulesServerTransport.
func (s *SecurityUserRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SecurityUserRulesClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SecurityUserRulesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SecurityUserRulesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SecurityUserRulesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SecurityUserRulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityUserConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.SecurityUserRule](req)
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
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityUserRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityUserRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityUserConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		var options *armnetwork.SecurityUserRulesClientBeginDeleteOptions
		if forceParam != nil {
			options = &armnetwork.SecurityUserRulesClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SecurityUserRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityUserConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules/(?P<ruleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, ruleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityUserRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityUserRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/securityUserConfigurations/(?P<configurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleCollections/(?P<ruleCollectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rules`
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
		var options *armnetwork.SecurityUserRulesClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.SecurityUserRulesClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, networkManagerNameParam, configurationNameParam, ruleCollectionNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.SecurityUserRulesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}
