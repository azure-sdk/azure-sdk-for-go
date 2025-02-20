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

// SecurityPerimeterAccessRulesServer is a fake server for instances of the armnetwork.SecurityPerimeterAccessRulesClient type.
type SecurityPerimeterAccessRulesServer struct {
	// CreateOrUpdate is the fake for method SecurityPerimeterAccessRulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string, parameters armnetwork.NspAccessRule, options *armnetwork.SecurityPerimeterAccessRulesClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.SecurityPerimeterAccessRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SecurityPerimeterAccessRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string, options *armnetwork.SecurityPerimeterAccessRulesClientGetOptions) (resp azfake.Responder[armnetwork.SecurityPerimeterAccessRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SecurityPerimeterAccessRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkSecurityPerimeterName string, profileName string, options *armnetwork.SecurityPerimeterAccessRulesClientListOptions) (resp azfake.PagerResponder[armnetwork.SecurityPerimeterAccessRulesClientListResponse])

	// Reconcile is the fake for method SecurityPerimeterAccessRulesClient.Reconcile
	// HTTP status codes to indicate success: http.StatusOK
	Reconcile func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, accessRuleName string, parameters any, options *armnetwork.SecurityPerimeterAccessRulesClientReconcileOptions) (resp azfake.Responder[armnetwork.SecurityPerimeterAccessRulesClientReconcileResponse], errResp azfake.ErrorResponder)
}

// NewSecurityPerimeterAccessRulesServerTransport creates a new instance of SecurityPerimeterAccessRulesServerTransport with the provided implementation.
// The returned SecurityPerimeterAccessRulesServerTransport instance is connected to an instance of armnetwork.SecurityPerimeterAccessRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSecurityPerimeterAccessRulesServerTransport(srv *SecurityPerimeterAccessRulesServer) *SecurityPerimeterAccessRulesServerTransport {
	return &SecurityPerimeterAccessRulesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.SecurityPerimeterAccessRulesClientListResponse]](),
	}
}

// SecurityPerimeterAccessRulesServerTransport connects instances of armnetwork.SecurityPerimeterAccessRulesClient to instances of SecurityPerimeterAccessRulesServer.
// Don't use this type directly, use NewSecurityPerimeterAccessRulesServerTransport instead.
type SecurityPerimeterAccessRulesServerTransport struct {
	srv          *SecurityPerimeterAccessRulesServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.SecurityPerimeterAccessRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for SecurityPerimeterAccessRulesServerTransport.
func (s *SecurityPerimeterAccessRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SecurityPerimeterAccessRulesClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SecurityPerimeterAccessRulesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SecurityPerimeterAccessRulesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SecurityPerimeterAccessRulesClient.Reconcile":
		resp, err = s.dispatchReconcile(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SecurityPerimeterAccessRulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessRules/(?P<accessRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.NspAccessRule](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	accessRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, accessRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NspAccessRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityPerimeterAccessRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessRules/(?P<accessRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	accessRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, accessRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NspAccessRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityPerimeterAccessRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
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
		var options *armnetwork.SecurityPerimeterAccessRulesClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.SecurityPerimeterAccessRulesClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.SecurityPerimeterAccessRulesClientListResponse, createLink func() string) {
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

func (s *SecurityPerimeterAccessRulesServerTransport) dispatchReconcile(req *http.Request) (*http.Response, error) {
	if s.srv.Reconcile == nil {
		return nil, &nonRetriableError{errors.New("fake for method Reconcile not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/accessRules/(?P<accessRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reconcile`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[any](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	accessRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accessRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Reconcile(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, accessRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Interface, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
