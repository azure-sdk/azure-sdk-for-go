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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift/v2"
	"net/http"
	"net/url"
	"regexp"
)

// OpenShiftClustersServer is a fake server for instances of the armredhatopenshift.OpenShiftClustersClient type.
type OpenShiftClustersServer struct {
	// BeginCreateOrUpdate is the fake for method OpenShiftClustersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, parameters armredhatopenshift.OpenShiftCluster, options *armredhatopenshift.OpenShiftClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method OpenShiftClustersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, options *armredhatopenshift.OpenShiftClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method OpenShiftClustersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armredhatopenshift.OpenShiftClustersClientGetOptions) (resp azfake.Responder[armredhatopenshift.OpenShiftClustersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method OpenShiftClustersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armredhatopenshift.OpenShiftClustersClientListOptions) (resp azfake.PagerResponder[armredhatopenshift.OpenShiftClustersClientListResponse])

	// ListAdminCredentials is the fake for method OpenShiftClustersClient.ListAdminCredentials
	// HTTP status codes to indicate success: http.StatusOK
	ListAdminCredentials func(ctx context.Context, resourceGroupName string, resourceName string, options *armredhatopenshift.OpenShiftClustersClientListAdminCredentialsOptions) (resp azfake.Responder[armredhatopenshift.OpenShiftClustersClientListAdminCredentialsResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method OpenShiftClustersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armredhatopenshift.OpenShiftClustersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armredhatopenshift.OpenShiftClustersClientListByResourceGroupResponse])

	// ListCredentials is the fake for method OpenShiftClustersClient.ListCredentials
	// HTTP status codes to indicate success: http.StatusOK
	ListCredentials func(ctx context.Context, resourceGroupName string, resourceName string, options *armredhatopenshift.OpenShiftClustersClientListCredentialsOptions) (resp azfake.Responder[armredhatopenshift.OpenShiftClustersClientListCredentialsResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method OpenShiftClustersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, resourceName string, parameters armredhatopenshift.OpenShiftClusterUpdate, options *armredhatopenshift.OpenShiftClustersClientBeginUpdateOptions) (resp azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewOpenShiftClustersServerTransport creates a new instance of OpenShiftClustersServerTransport with the provided implementation.
// The returned OpenShiftClustersServerTransport instance is connected to an instance of armredhatopenshift.OpenShiftClustersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewOpenShiftClustersServerTransport(srv *OpenShiftClustersServer) *OpenShiftClustersServerTransport {
	return &OpenShiftClustersServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armredhatopenshift.OpenShiftClustersClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armredhatopenshift.OpenShiftClustersClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientUpdateResponse]](),
	}
}

// OpenShiftClustersServerTransport connects instances of armredhatopenshift.OpenShiftClustersClient to instances of OpenShiftClustersServer.
// Don't use this type directly, use NewOpenShiftClustersServerTransport instead.
type OpenShiftClustersServerTransport struct {
	srv                         *OpenShiftClustersServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armredhatopenshift.OpenShiftClustersClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armredhatopenshift.OpenShiftClustersClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armredhatopenshift.OpenShiftClustersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for OpenShiftClustersServerTransport.
func (o *OpenShiftClustersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "OpenShiftClustersClient.BeginCreateOrUpdate":
		resp, err = o.dispatchBeginCreateOrUpdate(req)
	case "OpenShiftClustersClient.BeginDelete":
		resp, err = o.dispatchBeginDelete(req)
	case "OpenShiftClustersClient.Get":
		resp, err = o.dispatchGet(req)
	case "OpenShiftClustersClient.NewListPager":
		resp, err = o.dispatchNewListPager(req)
	case "OpenShiftClustersClient.ListAdminCredentials":
		resp, err = o.dispatchListAdminCredentials(req)
	case "OpenShiftClustersClient.NewListByResourceGroupPager":
		resp, err = o.dispatchNewListByResourceGroupPager(req)
	case "OpenShiftClustersClient.ListCredentials":
		resp, err = o.dispatchListCredentials(req)
	case "OpenShiftClustersClient.BeginUpdate":
		resp, err = o.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := o.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armredhatopenshift.OpenShiftCluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		o.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		o.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		o.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if o.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := o.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		o.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		o.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		o.beginDelete.remove(req)
	}

	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if o.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OpenShiftCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := o.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := o.srv.NewListPager(nil)
		newListPager = &resp
		o.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armredhatopenshift.OpenShiftClustersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		o.newListPager.remove(req)
	}
	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchListAdminCredentials(req *http.Request) (*http.Response, error) {
	if o.srv.ListAdminCredentials == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAdminCredentials not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listAdminCredentials`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.ListAdminCredentials(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OpenShiftClusterAdminKubeconfig, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if o.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := o.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := o.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		o.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armredhatopenshift.OpenShiftClustersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		o.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		o.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchListCredentials(req *http.Request) (*http.Response, error) {
	if o.srv.ListCredentials == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListCredentials not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listCredentials`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := o.srv.ListCredentials(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).OpenShiftClusterCredentials, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *OpenShiftClustersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if o.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := o.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RedHatOpenShift/openShiftClusters/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armredhatopenshift.OpenShiftClusterUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := o.srv.BeginUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		o.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		o.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		o.beginUpdate.remove(req)
	}

	return resp, nil
}
