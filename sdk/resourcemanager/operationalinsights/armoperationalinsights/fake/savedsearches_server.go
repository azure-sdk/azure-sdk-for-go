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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// SavedSearchesServer is a fake server for instances of the armoperationalinsights.SavedSearchesClient type.
type SavedSearchesServer struct {
	// CreateOrUpdate is the fake for method SavedSearchesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, parameters armoperationalinsights.SavedSearch, options *armoperationalinsights.SavedSearchesClientCreateOrUpdateOptions) (resp azfake.Responder[armoperationalinsights.SavedSearchesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SavedSearchesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *armoperationalinsights.SavedSearchesClientDeleteOptions) (resp azfake.Responder[armoperationalinsights.SavedSearchesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SavedSearchesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, savedSearchID string, options *armoperationalinsights.SavedSearchesClientGetOptions) (resp azfake.Responder[armoperationalinsights.SavedSearchesClientGetResponse], errResp azfake.ErrorResponder)

	// ListByWorkspace is the fake for method SavedSearchesClient.ListByWorkspace
	// HTTP status codes to indicate success: http.StatusOK
	ListByWorkspace func(ctx context.Context, resourceGroupName string, workspaceName string, options *armoperationalinsights.SavedSearchesClientListByWorkspaceOptions) (resp azfake.Responder[armoperationalinsights.SavedSearchesClientListByWorkspaceResponse], errResp azfake.ErrorResponder)
}

// NewSavedSearchesServerTransport creates a new instance of SavedSearchesServerTransport with the provided implementation.
// The returned SavedSearchesServerTransport instance is connected to an instance of armoperationalinsights.SavedSearchesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSavedSearchesServerTransport(srv *SavedSearchesServer) *SavedSearchesServerTransport {
	return &SavedSearchesServerTransport{srv: srv}
}

// SavedSearchesServerTransport connects instances of armoperationalinsights.SavedSearchesClient to instances of SavedSearchesServer.
// Don't use this type directly, use NewSavedSearchesServerTransport instead.
type SavedSearchesServerTransport struct {
	srv *SavedSearchesServer
}

// Do implements the policy.Transporter interface for SavedSearchesServerTransport.
func (s *SavedSearchesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SavedSearchesClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "SavedSearchesClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "SavedSearchesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SavedSearchesClient.ListByWorkspace":
		resp, err = s.dispatchListByWorkspace(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SavedSearchesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/savedSearches/(?P<savedSearchId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoperationalinsights.SavedSearch](req)
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
	savedSearchIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savedSearchId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, savedSearchIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SavedSearch, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SavedSearchesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/savedSearches/(?P<savedSearchId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	savedSearchIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savedSearchId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, savedSearchIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SavedSearchesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/savedSearches/(?P<savedSearchId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	savedSearchIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("savedSearchId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, savedSearchIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SavedSearch, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SavedSearchesServerTransport) dispatchListByWorkspace(req *http.Request) (*http.Response, error) {
	if s.srv.ListByWorkspace == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListByWorkspace not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/savedSearches`
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
	respr, errRespr := s.srv.ListByWorkspace(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SavedSearchesListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
