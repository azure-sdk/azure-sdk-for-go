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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase"
	"net/http"
	"net/url"
	"regexp"
)

// ImageDefinitionsServer is a fake server for instances of the armtestbase.ImageDefinitionsClient type.
type ImageDefinitionsServer struct {
	// Create is the fake for method ImageDefinitionsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, testBaseAccountName string, imageDefinitionName string, parameters armtestbase.ImageDefinitionResource, options *armtestbase.ImageDefinitionsClientCreateOptions) (resp azfake.Responder[armtestbase.ImageDefinitionsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ImageDefinitionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, testBaseAccountName string, imageDefinitionName string, options *armtestbase.ImageDefinitionsClientDeleteOptions) (resp azfake.Responder[armtestbase.ImageDefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ImageDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, testBaseAccountName string, imageDefinitionName string, options *armtestbase.ImageDefinitionsClientGetOptions) (resp azfake.Responder[armtestbase.ImageDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByTestBaseAccountPager is the fake for method ImageDefinitionsClient.NewListByTestBaseAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByTestBaseAccountPager func(resourceGroupName string, testBaseAccountName string, options *armtestbase.ImageDefinitionsClientListByTestBaseAccountOptions) (resp azfake.PagerResponder[armtestbase.ImageDefinitionsClientListByTestBaseAccountResponse])
}

// NewImageDefinitionsServerTransport creates a new instance of ImageDefinitionsServerTransport with the provided implementation.
// The returned ImageDefinitionsServerTransport instance is connected to an instance of armtestbase.ImageDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImageDefinitionsServerTransport(srv *ImageDefinitionsServer) *ImageDefinitionsServerTransport {
	return &ImageDefinitionsServerTransport{
		srv:                           srv,
		newListByTestBaseAccountPager: newTracker[azfake.PagerResponder[armtestbase.ImageDefinitionsClientListByTestBaseAccountResponse]](),
	}
}

// ImageDefinitionsServerTransport connects instances of armtestbase.ImageDefinitionsClient to instances of ImageDefinitionsServer.
// Don't use this type directly, use NewImageDefinitionsServerTransport instead.
type ImageDefinitionsServerTransport struct {
	srv                           *ImageDefinitionsServer
	newListByTestBaseAccountPager *tracker[azfake.PagerResponder[armtestbase.ImageDefinitionsClientListByTestBaseAccountResponse]]
}

// Do implements the policy.Transporter interface for ImageDefinitionsServerTransport.
func (i *ImageDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ImageDefinitionsClient.Create":
		resp, err = i.dispatchCreate(req)
	case "ImageDefinitionsClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "ImageDefinitionsClient.Get":
		resp, err = i.dispatchGet(req)
	case "ImageDefinitionsClient.NewListByTestBaseAccountPager":
		resp, err = i.dispatchNewListByTestBaseAccountPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *ImageDefinitionsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if i.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/imageDefinitions/(?P<imageDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armtestbase.ImageDefinitionResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
	if err != nil {
		return nil, err
	}
	imageDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Create(req.Context(), resourceGroupNameParam, testBaseAccountNameParam, imageDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImageDefinitionResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImageDefinitionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/imageDefinitions/(?P<imageDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
	if err != nil {
		return nil, err
	}
	imageDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, testBaseAccountNameParam, imageDefinitionNameParam, nil)
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

func (i *ImageDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/imageDefinitions/(?P<imageDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
	if err != nil {
		return nil, err
	}
	imageDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, testBaseAccountNameParam, imageDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImageDefinitionResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImageDefinitionsServerTransport) dispatchNewListByTestBaseAccountPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByTestBaseAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByTestBaseAccountPager not implemented")}
	}
	newListByTestBaseAccountPager := i.newListByTestBaseAccountPager.get(req)
	if newListByTestBaseAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.TestBase/testBaseAccounts/(?P<testBaseAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/imageDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		testBaseAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("testBaseAccountName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByTestBaseAccountPager(resourceGroupNameParam, testBaseAccountNameParam, nil)
		newListByTestBaseAccountPager = &resp
		i.newListByTestBaseAccountPager.add(req, newListByTestBaseAccountPager)
		server.PagerResponderInjectNextLinks(newListByTestBaseAccountPager, req, func(page *armtestbase.ImageDefinitionsClientListByTestBaseAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByTestBaseAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByTestBaseAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByTestBaseAccountPager) {
		i.newListByTestBaseAccountPager.remove(req)
	}
	return resp, nil
}
