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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos/v2"
	"net/http"
	"net/url"
	"regexp"
)

// TargetsServer is a fake server for instances of the armchaos.TargetsClient type.
type TargetsServer struct {
	// CreateOrUpdate is the fake for method TargetsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, target armchaos.Target, options *armchaos.TargetsClientCreateOrUpdateOptions) (resp azfake.Responder[armchaos.TargetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TargetsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *armchaos.TargetsClientDeleteOptions) (resp azfake.Responder[armchaos.TargetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TargetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, targetName string, options *armchaos.TargetsClientGetOptions) (resp azfake.Responder[armchaos.TargetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TargetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, parentProviderNamespace string, parentResourceType string, parentResourceName string, options *armchaos.TargetsClientListOptions) (resp azfake.PagerResponder[armchaos.TargetsClientListResponse])
}

// NewTargetsServerTransport creates a new instance of TargetsServerTransport with the provided implementation.
// The returned TargetsServerTransport instance is connected to an instance of armchaos.TargetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTargetsServerTransport(srv *TargetsServer) *TargetsServerTransport {
	return &TargetsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armchaos.TargetsClientListResponse]](),
	}
}

// TargetsServerTransport connects instances of armchaos.TargetsClient to instances of TargetsServer.
// Don't use this type directly, use NewTargetsServerTransport instead.
type TargetsServerTransport struct {
	srv          *TargetsServer
	newListPager *tracker[azfake.PagerResponder[armchaos.TargetsClientListResponse]]
}

// Do implements the policy.Transporter interface for TargetsServerTransport.
func (t *TargetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TargetsClient.CreateOrUpdate":
		resp, err = t.dispatchCreateOrUpdate(req)
	case "TargetsClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TargetsClient.Get":
		resp, err = t.dispatchGet(req)
	case "TargetsClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TargetsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<parentProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/targets/(?P<targetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armchaos.Target](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	parentProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentProviderNamespace")])
	if err != nil {
		return nil, err
	}
	parentResourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceType")])
	if err != nil {
		return nil, err
	}
	parentResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceName")])
	if err != nil {
		return nil, err
	}
	targetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, parentProviderNamespaceParam, parentResourceTypeParam, parentResourceNameParam, targetNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Target, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TargetsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<parentProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/targets/(?P<targetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	parentProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentProviderNamespace")])
	if err != nil {
		return nil, err
	}
	parentResourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceType")])
	if err != nil {
		return nil, err
	}
	parentResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceName")])
	if err != nil {
		return nil, err
	}
	targetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), resourceGroupNameParam, parentProviderNamespaceParam, parentResourceTypeParam, parentResourceNameParam, targetNameParam, nil)
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

func (t *TargetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<parentProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/targets/(?P<targetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	parentProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentProviderNamespace")])
	if err != nil {
		return nil, err
	}
	parentResourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceType")])
	if err != nil {
		return nil, err
	}
	parentResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceName")])
	if err != nil {
		return nil, err
	}
	targetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, parentProviderNamespaceParam, parentResourceTypeParam, parentResourceNameParam, targetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Target, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TargetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/(?P<parentProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<parentResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/targets`
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
		parentProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentProviderNamespace")])
		if err != nil {
			return nil, err
		}
		parentResourceTypeParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceType")])
		if err != nil {
			return nil, err
		}
		parentResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("parentResourceName")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armchaos.TargetsClientListOptions
		if continuationTokenParam != nil {
			options = &armchaos.TargetsClientListOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := t.srv.NewListPager(resourceGroupNameParam, parentProviderNamespaceParam, parentResourceTypeParam, parentResourceNameParam, options)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armchaos.TargetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}
