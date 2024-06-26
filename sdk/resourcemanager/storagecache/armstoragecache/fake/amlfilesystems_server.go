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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagecache/armstoragecache/v4"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// AmlFilesystemsServer is a fake server for instances of the armstoragecache.AmlFilesystemsClient type.
type AmlFilesystemsServer struct {
	// Archive is the fake for method AmlFilesystemsClient.Archive
	// HTTP status codes to indicate success: http.StatusOK
	Archive func(ctx context.Context, resourceGroupName string, amlFilesystemName string, options *armstoragecache.AmlFilesystemsClientArchiveOptions) (resp azfake.Responder[armstoragecache.AmlFilesystemsClientArchiveResponse], errResp azfake.ErrorResponder)

	// CancelArchive is the fake for method AmlFilesystemsClient.CancelArchive
	// HTTP status codes to indicate success: http.StatusOK
	CancelArchive func(ctx context.Context, resourceGroupName string, amlFilesystemName string, options *armstoragecache.AmlFilesystemsClientCancelArchiveOptions) (resp azfake.Responder[armstoragecache.AmlFilesystemsClientCancelArchiveResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method AmlFilesystemsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, amlFilesystemName string, amlFilesystem armstoragecache.AmlFilesystem, options *armstoragecache.AmlFilesystemsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armstoragecache.AmlFilesystemsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AmlFilesystemsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, amlFilesystemName string, options *armstoragecache.AmlFilesystemsClientBeginDeleteOptions) (resp azfake.PollerResponder[armstoragecache.AmlFilesystemsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AmlFilesystemsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, amlFilesystemName string, options *armstoragecache.AmlFilesystemsClientGetOptions) (resp azfake.Responder[armstoragecache.AmlFilesystemsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method AmlFilesystemsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armstoragecache.AmlFilesystemsClientListOptions) (resp azfake.PagerResponder[armstoragecache.AmlFilesystemsClientListResponse])

	// NewListByResourceGroupPager is the fake for method AmlFilesystemsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armstoragecache.AmlFilesystemsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armstoragecache.AmlFilesystemsClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method AmlFilesystemsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, amlFilesystemName string, amlFilesystem armstoragecache.AmlFilesystemUpdate, options *armstoragecache.AmlFilesystemsClientBeginUpdateOptions) (resp azfake.PollerResponder[armstoragecache.AmlFilesystemsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAmlFilesystemsServerTransport creates a new instance of AmlFilesystemsServerTransport with the provided implementation.
// The returned AmlFilesystemsServerTransport instance is connected to an instance of armstoragecache.AmlFilesystemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAmlFilesystemsServerTransport(srv *AmlFilesystemsServer) *AmlFilesystemsServerTransport {
	return &AmlFilesystemsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armstoragecache.AmlFilesystemsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armstoragecache.AmlFilesystemsClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armstoragecache.AmlFilesystemsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armstoragecache.AmlFilesystemsClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armstoragecache.AmlFilesystemsClientUpdateResponse]](),
	}
}

// AmlFilesystemsServerTransport connects instances of armstoragecache.AmlFilesystemsClient to instances of AmlFilesystemsServer.
// Don't use this type directly, use NewAmlFilesystemsServerTransport instead.
type AmlFilesystemsServerTransport struct {
	srv                         *AmlFilesystemsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armstoragecache.AmlFilesystemsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armstoragecache.AmlFilesystemsClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armstoragecache.AmlFilesystemsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armstoragecache.AmlFilesystemsClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armstoragecache.AmlFilesystemsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AmlFilesystemsServerTransport.
func (a *AmlFilesystemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AmlFilesystemsClient.Archive":
		resp, err = a.dispatchArchive(req)
	case "AmlFilesystemsClient.CancelArchive":
		resp, err = a.dispatchCancelArchive(req)
	case "AmlFilesystemsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AmlFilesystemsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AmlFilesystemsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AmlFilesystemsClient.NewListPager":
		resp, err = a.dispatchNewListPager(req)
	case "AmlFilesystemsClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AmlFilesystemsClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AmlFilesystemsServerTransport) dispatchArchive(req *http.Request) (*http.Response, error) {
	if a.srv.Archive == nil {
		return nil, &nonRetriableError{errors.New("fake for method Archive not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems/(?P<amlFilesystemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/archive`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armstoragecache.AmlFilesystemArchiveInfo](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	amlFilesystemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("amlFilesystemName")])
	if err != nil {
		return nil, err
	}
	var options *armstoragecache.AmlFilesystemsClientArchiveOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armstoragecache.AmlFilesystemsClientArchiveOptions{
			ArchiveInfo: &body,
		}
	}
	respr, errRespr := a.srv.Archive(req.Context(), resourceGroupNameParam, amlFilesystemNameParam, options)
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

func (a *AmlFilesystemsServerTransport) dispatchCancelArchive(req *http.Request) (*http.Response, error) {
	if a.srv.CancelArchive == nil {
		return nil, &nonRetriableError{errors.New("fake for method CancelArchive not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems/(?P<amlFilesystemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancelArchive`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	amlFilesystemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("amlFilesystemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CancelArchive(req.Context(), resourceGroupNameParam, amlFilesystemNameParam, nil)
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

func (a *AmlFilesystemsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems/(?P<amlFilesystemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragecache.AmlFilesystem](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		amlFilesystemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("amlFilesystemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, amlFilesystemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AmlFilesystemsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems/(?P<amlFilesystemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		amlFilesystemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("amlFilesystemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, amlFilesystemNameParam, nil)
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

func (a *AmlFilesystemsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems/(?P<amlFilesystemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	amlFilesystemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("amlFilesystemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, amlFilesystemNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AmlFilesystem, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AmlFilesystemsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := a.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListPager(nil)
		newListPager = &resp
		a.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armstoragecache.AmlFilesystemsClientListResponse, createLink func() string) {
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

func (a *AmlFilesystemsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armstoragecache.AmlFilesystemsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *AmlFilesystemsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.StorageCache/amlFilesystems/(?P<amlFilesystemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armstoragecache.AmlFilesystemUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		amlFilesystemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("amlFilesystemName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, amlFilesystemNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}
