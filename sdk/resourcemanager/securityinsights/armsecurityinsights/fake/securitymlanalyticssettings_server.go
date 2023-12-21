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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights"
	"net/http"
	"net/url"
	"regexp"
)

// SecurityMLAnalyticsSettingsServer is a fake server for instances of the armsecurityinsights.SecurityMLAnalyticsSettingsClient type.
type SecurityMLAnalyticsSettingsServer struct {
	// CreateOrUpdate is the fake for method SecurityMLAnalyticsSettingsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, securityMLAnalyticsSetting armsecurityinsights.SecurityMLAnalyticsSettingClassification, options *armsecurityinsights.SecurityMLAnalyticsSettingsClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurityinsights.SecurityMLAnalyticsSettingsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SecurityMLAnalyticsSettingsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, options *armsecurityinsights.SecurityMLAnalyticsSettingsClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.SecurityMLAnalyticsSettingsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SecurityMLAnalyticsSettingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, settingsResourceName string, options *armsecurityinsights.SecurityMLAnalyticsSettingsClientGetOptions) (resp azfake.Responder[armsecurityinsights.SecurityMLAnalyticsSettingsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SecurityMLAnalyticsSettingsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.SecurityMLAnalyticsSettingsClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.SecurityMLAnalyticsSettingsClientListResponse])
}

// NewSecurityMLAnalyticsSettingsServerTransport creates a new instance of SecurityMLAnalyticsSettingsServerTransport with the provided implementation.
// The returned SecurityMLAnalyticsSettingsServerTransport instance is connected to an instance of armsecurityinsights.SecurityMLAnalyticsSettingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSecurityMLAnalyticsSettingsServerTransport(srv *SecurityMLAnalyticsSettingsServer) *SecurityMLAnalyticsSettingsServerTransport {
	return &SecurityMLAnalyticsSettingsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.SecurityMLAnalyticsSettingsClientListResponse]](),
	}
}

// SecurityMLAnalyticsSettingsServerTransport connects instances of armsecurityinsights.SecurityMLAnalyticsSettingsClient to instances of SecurityMLAnalyticsSettingsServer.
// Don't use this type directly, use NewSecurityMLAnalyticsSettingsServerTransport instead.
type SecurityMLAnalyticsSettingsServerTransport struct {
	srv          *SecurityMLAnalyticsSettingsServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.SecurityMLAnalyticsSettingsClientListResponse]]
}

// Do implements the policy.Transporter interface for SecurityMLAnalyticsSettingsServerTransport.
func (s *SecurityMLAnalyticsSettingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SecurityMLAnalyticsSettingsClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SecurityMLAnalyticsSettingsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "SecurityMLAnalyticsSettingsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SecurityMLAnalyticsSettingsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SecurityMLAnalyticsSettingsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/securityMLAnalyticsSettings/(?P<settingsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	raw, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}
	body, err := unmarshalSecurityMLAnalyticsSettingClassification(raw)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	settingsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("settingsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, settingsResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityMLAnalyticsSettingClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityMLAnalyticsSettingsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/securityMLAnalyticsSettings/(?P<settingsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	settingsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("settingsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, settingsResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityMLAnalyticsSettingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/securityMLAnalyticsSettings/(?P<settingsResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	settingsResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("settingsResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, settingsResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SecurityMLAnalyticsSettingClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityMLAnalyticsSettingsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/securityMLAnalyticsSettings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.SecurityMLAnalyticsSettingsClientListResponse, createLink func() string) {
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
