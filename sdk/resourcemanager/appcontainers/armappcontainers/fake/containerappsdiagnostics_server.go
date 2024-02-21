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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerAppsDiagnosticsServer is a fake server for instances of the armappcontainers.ContainerAppsDiagnosticsClient type.
type ContainerAppsDiagnosticsServer struct {
	// GetDetector is the fake for method ContainerAppsDiagnosticsClient.GetDetector
	// HTTP status codes to indicate success: http.StatusOK
	GetDetector func(ctx context.Context, resourceGroupName string, containerAppName string, detectorName string, options *armappcontainers.ContainerAppsDiagnosticsClientGetDetectorOptions) (resp azfake.Responder[armappcontainers.ContainerAppsDiagnosticsClientGetDetectorResponse], errResp azfake.ErrorResponder)

	// GetRevision is the fake for method ContainerAppsDiagnosticsClient.GetRevision
	// HTTP status codes to indicate success: http.StatusOK
	GetRevision func(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, options *armappcontainers.ContainerAppsDiagnosticsClientGetRevisionOptions) (resp azfake.Responder[armappcontainers.ContainerAppsDiagnosticsClientGetRevisionResponse], errResp azfake.ErrorResponder)

	// GetRoot is the fake for method ContainerAppsDiagnosticsClient.GetRoot
	// HTTP status codes to indicate success: http.StatusOK
	GetRoot func(ctx context.Context, resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsDiagnosticsClientGetRootOptions) (resp azfake.Responder[armappcontainers.ContainerAppsDiagnosticsClientGetRootResponse], errResp azfake.ErrorResponder)

	// NewListDetectorsPager is the fake for method ContainerAppsDiagnosticsClient.NewListDetectorsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListDetectorsPager func(resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsDiagnosticsClientListDetectorsOptions) (resp azfake.PagerResponder[armappcontainers.ContainerAppsDiagnosticsClientListDetectorsResponse])

	// NewListRevisionsPager is the fake for method ContainerAppsDiagnosticsClient.NewListRevisionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListRevisionsPager func(resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsDiagnosticsClientListRevisionsOptions) (resp azfake.PagerResponder[armappcontainers.ContainerAppsDiagnosticsClientListRevisionsResponse])
}

// NewContainerAppsDiagnosticsServerTransport creates a new instance of ContainerAppsDiagnosticsServerTransport with the provided implementation.
// The returned ContainerAppsDiagnosticsServerTransport instance is connected to an instance of armappcontainers.ContainerAppsDiagnosticsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsDiagnosticsServerTransport(srv *ContainerAppsDiagnosticsServer) *ContainerAppsDiagnosticsServerTransport {
	return &ContainerAppsDiagnosticsServerTransport{
		srv:                   srv,
		newListDetectorsPager: newTracker[azfake.PagerResponder[armappcontainers.ContainerAppsDiagnosticsClientListDetectorsResponse]](),
		newListRevisionsPager: newTracker[azfake.PagerResponder[armappcontainers.ContainerAppsDiagnosticsClientListRevisionsResponse]](),
	}
}

// ContainerAppsDiagnosticsServerTransport connects instances of armappcontainers.ContainerAppsDiagnosticsClient to instances of ContainerAppsDiagnosticsServer.
// Don't use this type directly, use NewContainerAppsDiagnosticsServerTransport instead.
type ContainerAppsDiagnosticsServerTransport struct {
	srv                   *ContainerAppsDiagnosticsServer
	newListDetectorsPager *tracker[azfake.PagerResponder[armappcontainers.ContainerAppsDiagnosticsClientListDetectorsResponse]]
	newListRevisionsPager *tracker[azfake.PagerResponder[armappcontainers.ContainerAppsDiagnosticsClientListRevisionsResponse]]
}

// Do implements the policy.Transporter interface for ContainerAppsDiagnosticsServerTransport.
func (c *ContainerAppsDiagnosticsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerAppsDiagnosticsClient.GetDetector":
		resp, err = c.dispatchGetDetector(req)
	case "ContainerAppsDiagnosticsClient.GetRevision":
		resp, err = c.dispatchGetRevision(req)
	case "ContainerAppsDiagnosticsClient.GetRoot":
		resp, err = c.dispatchGetRoot(req)
	case "ContainerAppsDiagnosticsClient.NewListDetectorsPager":
		resp, err = c.dispatchNewListDetectorsPager(req)
	case "ContainerAppsDiagnosticsClient.NewListRevisionsPager":
		resp, err = c.dispatchNewListRevisionsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerAppsDiagnosticsServerTransport) dispatchGetDetector(req *http.Request) (*http.Response, error) {
	if c.srv.GetDetector == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDetector not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectors/(?P<detectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	detectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("detectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetDetector(req.Context(), resourceGroupNameParam, containerAppNameParam, detectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Diagnostics, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsDiagnosticsServerTransport) dispatchGetRevision(req *http.Request) (*http.Response, error) {
	if c.srv.GetRevision == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRevision not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectorProperties/revisionsApi/revisions/(?P<revisionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	revisionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("revisionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetRevision(req.Context(), resourceGroupNameParam, containerAppNameParam, revisionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Revision, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsDiagnosticsServerTransport) dispatchGetRoot(req *http.Request) (*http.Response, error) {
	if c.srv.GetRoot == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRoot not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectorProperties/rootApi/`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetRoot(req.Context(), resourceGroupNameParam, containerAppNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerApp, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsDiagnosticsServerTransport) dispatchNewListDetectorsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListDetectorsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListDetectorsPager not implemented")}
	}
	newListDetectorsPager := c.newListDetectorsPager.get(req)
	if newListDetectorsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListDetectorsPager(resourceGroupNameParam, containerAppNameParam, nil)
		newListDetectorsPager = &resp
		c.newListDetectorsPager.add(req, newListDetectorsPager)
		server.PagerResponderInjectNextLinks(newListDetectorsPager, req, func(page *armappcontainers.ContainerAppsDiagnosticsClientListDetectorsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListDetectorsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListDetectorsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListDetectorsPager) {
		c.newListDetectorsPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerAppsDiagnosticsServerTransport) dispatchNewListRevisionsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListRevisionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListRevisionsPager not implemented")}
	}
	newListRevisionsPager := c.newListRevisionsPager.get(req)
	if newListRevisionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectorProperties/revisionsApi/revisions/`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armappcontainers.ContainerAppsDiagnosticsClientListRevisionsOptions
		if filterParam != nil {
			options = &armappcontainers.ContainerAppsDiagnosticsClientListRevisionsOptions{
				Filter: filterParam,
			}
		}
		resp := c.srv.NewListRevisionsPager(resourceGroupNameParam, containerAppNameParam, options)
		newListRevisionsPager = &resp
		c.newListRevisionsPager.add(req, newListRevisionsPager)
		server.PagerResponderInjectNextLinks(newListRevisionsPager, req, func(page *armappcontainers.ContainerAppsDiagnosticsClientListRevisionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListRevisionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListRevisionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListRevisionsPager) {
		c.newListRevisionsPager.remove(req)
	}
	return resp, nil
}
