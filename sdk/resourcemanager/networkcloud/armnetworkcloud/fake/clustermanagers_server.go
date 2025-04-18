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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkcloud/armnetworkcloud"
	"net/http"
	"net/url"
	"regexp"
)

// ClusterManagersServer is a fake server for instances of the armnetworkcloud.ClusterManagersClient type.
type ClusterManagersServer struct {
	// BeginCreateOrUpdate is the fake for method ClusterManagersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerParameters armnetworkcloud.ClusterManager, options *armnetworkcloud.ClusterManagersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.ClusterManagersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ClusterManagersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterManagerName string, options *armnetworkcloud.ClusterManagersClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkcloud.ClusterManagersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ClusterManagersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterManagerName string, options *armnetworkcloud.ClusterManagersClientGetOptions) (resp azfake.Responder[armnetworkcloud.ClusterManagersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ClusterManagersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetworkcloud.ClusterManagersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetworkcloud.ClusterManagersClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ClusterManagersClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetworkcloud.ClusterManagersClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetworkcloud.ClusterManagersClientListBySubscriptionResponse])

	// Update is the fake for method ClusterManagersClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, clusterManagerName string, clusterManagerUpdateParameters armnetworkcloud.ClusterManagerPatchParameters, options *armnetworkcloud.ClusterManagersClientUpdateOptions) (resp azfake.Responder[armnetworkcloud.ClusterManagersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewClusterManagersServerTransport creates a new instance of ClusterManagersServerTransport with the provided implementation.
// The returned ClusterManagersServerTransport instance is connected to an instance of armnetworkcloud.ClusterManagersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClusterManagersServerTransport(srv *ClusterManagersServer) *ClusterManagersServerTransport {
	return &ClusterManagersServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetworkcloud.ClusterManagersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetworkcloud.ClusterManagersClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetworkcloud.ClusterManagersClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armnetworkcloud.ClusterManagersClientListBySubscriptionResponse]](),
	}
}

// ClusterManagersServerTransport connects instances of armnetworkcloud.ClusterManagersClient to instances of ClusterManagersServer.
// Don't use this type directly, use NewClusterManagersServerTransport instead.
type ClusterManagersServerTransport struct {
	srv                         *ClusterManagersServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetworkcloud.ClusterManagersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetworkcloud.ClusterManagersClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetworkcloud.ClusterManagersClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armnetworkcloud.ClusterManagersClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for ClusterManagersServerTransport.
func (c *ClusterManagersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ClusterManagersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if clusterManagersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = clusterManagersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ClusterManagersClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "ClusterManagersClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ClusterManagersClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ClusterManagersClient.NewListByResourceGroupPager":
				res.resp, res.err = c.dispatchNewListByResourceGroupPager(req)
			case "ClusterManagersClient.NewListBySubscriptionPager":
				res.resp, res.err = c.dispatchNewListBySubscriptionPager(req)
			case "ClusterManagersClient.Update":
				res.resp, res.err = c.dispatchUpdate(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (c *ClusterManagersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusterManagers/(?P<clusterManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterManager](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterManagerName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ClusterManagersClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ClusterManagersClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, clusterManagerNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ClusterManagersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusterManagers/(?P<clusterManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterManagerName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ClusterManagersClientBeginDeleteOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ClusterManagersClientBeginDeleteOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterManagerNameParam, options)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ClusterManagersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusterManagers/(?P<clusterManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterManagerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, clusterManagerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ClusterManager, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ClusterManagersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusterManagers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		c.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetworkcloud.ClusterManagersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		c.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (c *ClusterManagersServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusterManagers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armnetworkcloud.ClusterManagersClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *ClusterManagersServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusterManagers/(?P<clusterManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterManagerPatchParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterManagerName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armnetworkcloud.ClusterManagersClientUpdateOptions
	if ifMatchParam != nil || ifNoneMatchParam != nil {
		options = &armnetworkcloud.ClusterManagersClientUpdateOptions{
			IfMatch:     ifMatchParam,
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, clusterManagerNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ClusterManager, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ClusterManagersServerTransport
var clusterManagersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
