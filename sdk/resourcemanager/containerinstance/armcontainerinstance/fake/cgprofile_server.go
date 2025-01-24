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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance/v2"
	"net/http"
	"net/url"
	"regexp"
)

// CGProfileServer is a fake server for instances of the armcontainerinstance.CGProfileClient type.
type CGProfileServer struct {
	// CreateOrUpdate is the fake for method CGProfileClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, containerGroupProfileName string, containerGroupProfile armcontainerinstance.ContainerGroupProfile, options *armcontainerinstance.CGProfileClientCreateOrUpdateOptions) (resp azfake.Responder[armcontainerinstance.CGProfileClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CGProfileClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *armcontainerinstance.CGProfileClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerinstance.CGProfileClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CGProfileClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, containerGroupProfileName string, options *armcontainerinstance.CGProfileClientGetOptions) (resp azfake.Responder[armcontainerinstance.CGProfileClientGetResponse], errResp azfake.ErrorResponder)

	// GetByRevisionNumber is the fake for method CGProfileClient.GetByRevisionNumber
	// HTTP status codes to indicate success: http.StatusOK
	GetByRevisionNumber func(ctx context.Context, resourceGroupName string, containerGroupProfileName string, revisionNumber string, options *armcontainerinstance.CGProfileClientGetByRevisionNumberOptions) (resp azfake.Responder[armcontainerinstance.CGProfileClientGetByRevisionNumberResponse], errResp azfake.ErrorResponder)

	// NewListAllRevisionsPager is the fake for method CGProfileClient.NewListAllRevisionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllRevisionsPager func(resourceGroupName string, containerGroupProfileName string, options *armcontainerinstance.CGProfileClientListAllRevisionsOptions) (resp azfake.PagerResponder[armcontainerinstance.CGProfileClientListAllRevisionsResponse])

	// Update is the fake for method CGProfileClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, containerGroupProfileName string, properties armcontainerinstance.ContainerGroupProfilePatch, options *armcontainerinstance.CGProfileClientUpdateOptions) (resp azfake.Responder[armcontainerinstance.CGProfileClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCGProfileServerTransport creates a new instance of CGProfileServerTransport with the provided implementation.
// The returned CGProfileServerTransport instance is connected to an instance of armcontainerinstance.CGProfileClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCGProfileServerTransport(srv *CGProfileServer) *CGProfileServerTransport {
	return &CGProfileServerTransport{
		srv:                      srv,
		beginDelete:              newTracker[azfake.PollerResponder[armcontainerinstance.CGProfileClientDeleteResponse]](),
		newListAllRevisionsPager: newTracker[azfake.PagerResponder[armcontainerinstance.CGProfileClientListAllRevisionsResponse]](),
	}
}

// CGProfileServerTransport connects instances of armcontainerinstance.CGProfileClient to instances of CGProfileServer.
// Don't use this type directly, use NewCGProfileServerTransport instead.
type CGProfileServerTransport struct {
	srv                      *CGProfileServer
	beginDelete              *tracker[azfake.PollerResponder[armcontainerinstance.CGProfileClientDeleteResponse]]
	newListAllRevisionsPager *tracker[azfake.PagerResponder[armcontainerinstance.CGProfileClientListAllRevisionsResponse]]
}

// Do implements the policy.Transporter interface for CGProfileServerTransport.
func (c *CGProfileServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CGProfileClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "CGProfileClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CGProfileClient.Get":
		resp, err = c.dispatchGet(req)
	case "CGProfileClient.GetByRevisionNumber":
		resp, err = c.dispatchGetByRevisionNumber(req)
	case "CGProfileClient.NewListAllRevisionsPager":
		resp, err = c.dispatchNewListAllRevisionsPager(req)
	case "CGProfileClient.Update":
		resp, err = c.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CGProfileServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcontainerinstance.ContainerGroupProfile](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, containerGroupProfileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroupProfile, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSCorrelationRequestID; val != nil {
		resp.Header.Set("x-ms-correlation-request-id", *val)
	}
	return resp, nil
}

func (c *CGProfileServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, containerGroupProfileNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *CGProfileServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, containerGroupProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroupProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CGProfileServerTransport) dispatchGetByRevisionNumber(req *http.Request) (*http.Response, error) {
	if c.srv.GetByRevisionNumber == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByRevisionNumber not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<revisionNumber>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
	if err != nil {
		return nil, err
	}
	revisionNumberParam, err := url.PathUnescape(matches[regex.SubexpIndex("revisionNumber")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetByRevisionNumber(req.Context(), resourceGroupNameParam, containerGroupProfileNameParam, revisionNumberParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroupProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CGProfileServerTransport) dispatchNewListAllRevisionsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListAllRevisionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllRevisionsPager not implemented")}
	}
	newListAllRevisionsPager := c.newListAllRevisionsPager.get(req)
	if newListAllRevisionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListAllRevisionsPager(resourceGroupNameParam, containerGroupProfileNameParam, nil)
		newListAllRevisionsPager = &resp
		c.newListAllRevisionsPager.add(req, newListAllRevisionsPager)
		server.PagerResponderInjectNextLinks(newListAllRevisionsPager, req, func(page *armcontainerinstance.CGProfileClientListAllRevisionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllRevisionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListAllRevisionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllRevisionsPager) {
		c.newListAllRevisionsPager.remove(req)
	}
	return resp, nil
}

func (c *CGProfileServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerInstance/containerGroupProfiles/(?P<containerGroupProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcontainerinstance.ContainerGroupProfilePatch](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerGroupProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerGroupProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, containerGroupProfileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerGroupProfile, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSCorrelationRequestID; val != nil {
		resp.Header.Set("x-ms-correlation-request-id", *val)
	}
	return resp, nil
}
