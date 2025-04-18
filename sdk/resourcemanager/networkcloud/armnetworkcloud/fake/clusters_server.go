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
	"reflect"
	"regexp"
)

// ClustersServer is a fake server for instances of the armnetworkcloud.ClustersClient type.
type ClustersServer struct {
	// BeginContinueUpdateVersion is the fake for method ClustersClient.BeginContinueUpdateVersion
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginContinueUpdateVersion func(ctx context.Context, resourceGroupName string, clusterName string, clusterContinueUpdateVersionParameters armnetworkcloud.ClusterContinueUpdateVersionParameters, options *armnetworkcloud.ClustersClientBeginContinueUpdateVersionOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientContinueUpdateVersionResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ClustersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterName string, clusterParameters armnetworkcloud.Cluster, options *armnetworkcloud.ClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ClustersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, clusterName string, options *armnetworkcloud.ClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeploy is the fake for method ClustersClient.BeginDeploy
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeploy func(ctx context.Context, resourceGroupName string, clusterName string, options *armnetworkcloud.ClustersClientBeginDeployOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientDeployResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ClustersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, options *armnetworkcloud.ClustersClientGetOptions) (resp azfake.Responder[armnetworkcloud.ClustersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ClustersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetworkcloud.ClustersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetworkcloud.ClustersClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ClustersClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armnetworkcloud.ClustersClientListBySubscriptionOptions) (resp azfake.PagerResponder[armnetworkcloud.ClustersClientListBySubscriptionResponse])

	// BeginScanRuntime is the fake for method ClustersClient.BeginScanRuntime
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginScanRuntime func(ctx context.Context, resourceGroupName string, clusterName string, options *armnetworkcloud.ClustersClientBeginScanRuntimeOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientScanRuntimeResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ClustersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateParameters armnetworkcloud.ClusterPatchParameters, options *armnetworkcloud.ClustersClientBeginUpdateOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateVersion is the fake for method ClustersClient.BeginUpdateVersion
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateVersion func(ctx context.Context, resourceGroupName string, clusterName string, clusterUpdateVersionParameters armnetworkcloud.ClusterUpdateVersionParameters, options *armnetworkcloud.ClustersClientBeginUpdateVersionOptions) (resp azfake.PollerResponder[armnetworkcloud.ClustersClientUpdateVersionResponse], errResp azfake.ErrorResponder)
}

// NewClustersServerTransport creates a new instance of ClustersServerTransport with the provided implementation.
// The returned ClustersServerTransport instance is connected to an instance of armnetworkcloud.ClustersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClustersServerTransport(srv *ClustersServer) *ClustersServerTransport {
	return &ClustersServerTransport{
		srv:                         srv,
		beginContinueUpdateVersion:  newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientContinueUpdateVersionResponse]](),
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientDeleteResponse]](),
		beginDeploy:                 newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientDeployResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetworkcloud.ClustersClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armnetworkcloud.ClustersClientListBySubscriptionResponse]](),
		beginScanRuntime:            newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientScanRuntimeResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientUpdateResponse]](),
		beginUpdateVersion:          newTracker[azfake.PollerResponder[armnetworkcloud.ClustersClientUpdateVersionResponse]](),
	}
}

// ClustersServerTransport connects instances of armnetworkcloud.ClustersClient to instances of ClustersServer.
// Don't use this type directly, use NewClustersServerTransport instead.
type ClustersServerTransport struct {
	srv                         *ClustersServer
	beginContinueUpdateVersion  *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientContinueUpdateVersionResponse]]
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientDeleteResponse]]
	beginDeploy                 *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientDeployResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetworkcloud.ClustersClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armnetworkcloud.ClustersClientListBySubscriptionResponse]]
	beginScanRuntime            *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientScanRuntimeResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientUpdateResponse]]
	beginUpdateVersion          *tracker[azfake.PollerResponder[armnetworkcloud.ClustersClientUpdateVersionResponse]]
}

// Do implements the policy.Transporter interface for ClustersServerTransport.
func (c *ClustersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ClustersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if clustersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = clustersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ClustersClient.BeginContinueUpdateVersion":
				res.resp, res.err = c.dispatchBeginContinueUpdateVersion(req)
			case "ClustersClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "ClustersClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ClustersClient.BeginDeploy":
				res.resp, res.err = c.dispatchBeginDeploy(req)
			case "ClustersClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ClustersClient.NewListByResourceGroupPager":
				res.resp, res.err = c.dispatchNewListByResourceGroupPager(req)
			case "ClustersClient.NewListBySubscriptionPager":
				res.resp, res.err = c.dispatchNewListBySubscriptionPager(req)
			case "ClustersClient.BeginScanRuntime":
				res.resp, res.err = c.dispatchBeginScanRuntime(req)
			case "ClustersClient.BeginUpdate":
				res.resp, res.err = c.dispatchBeginUpdate(req)
			case "ClustersClient.BeginUpdateVersion":
				res.resp, res.err = c.dispatchBeginUpdateVersion(req)
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

func (c *ClustersServerTransport) dispatchBeginContinueUpdateVersion(req *http.Request) (*http.Response, error) {
	if c.srv.BeginContinueUpdateVersion == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginContinueUpdateVersion not implemented")}
	}
	beginContinueUpdateVersion := c.beginContinueUpdateVersion.get(req)
	if beginContinueUpdateVersion == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/continueUpdateVersion`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterContinueUpdateVersionParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginContinueUpdateVersion(req.Context(), resourceGroupNameParam, clusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginContinueUpdateVersion = &respr
		c.beginContinueUpdateVersion.add(req, beginContinueUpdateVersion)
	}

	resp, err := server.PollerResponderNext(beginContinueUpdateVersion, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginContinueUpdateVersion.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginContinueUpdateVersion) {
		c.beginContinueUpdateVersion.remove(req)
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.Cluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ClustersClientBeginCreateOrUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ClustersClientBeginCreateOrUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, body, options)
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

func (c *ClustersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ClustersClientBeginDeleteOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ClustersClientBeginDeleteOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, clusterNameParam, options)
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

func (c *ClustersServerTransport) dispatchBeginDeploy(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDeploy == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeploy not implemented")}
	}
	beginDeploy := c.beginDeploy.get(req)
	if beginDeploy == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deploy`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterDeployParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		var options *armnetworkcloud.ClustersClientBeginDeployOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetworkcloud.ClustersClientBeginDeployOptions{
				ClusterDeployParameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginDeploy(req.Context(), resourceGroupNameParam, clusterNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeploy = &respr
		c.beginDeploy.add(req, beginDeploy)
	}

	resp, err := server.PollerResponderNext(beginDeploy, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginDeploy.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeploy) {
		c.beginDeploy.remove(req)
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Cluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ClustersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters`
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
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetworkcloud.ClustersClientListByResourceGroupResponse, createLink func() string) {
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

func (c *ClustersServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armnetworkcloud.ClustersClientListBySubscriptionResponse, createLink func() string) {
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

func (c *ClustersServerTransport) dispatchBeginScanRuntime(req *http.Request) (*http.Response, error) {
	if c.srv.BeginScanRuntime == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginScanRuntime not implemented")}
	}
	beginScanRuntime := c.beginScanRuntime.get(req)
	if beginScanRuntime == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/scanRuntime`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterScanRuntimeParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		var options *armnetworkcloud.ClustersClientBeginScanRuntimeOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetworkcloud.ClustersClientBeginScanRuntimeOptions{
				ClusterScanRuntimeParameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginScanRuntime(req.Context(), resourceGroupNameParam, clusterNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginScanRuntime = &respr
		c.beginScanRuntime.add(req, beginScanRuntime)
	}

	resp, err := server.PollerResponderNext(beginScanRuntime, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginScanRuntime.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginScanRuntime) {
		c.beginScanRuntime.remove(req)
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterPatchParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
		var options *armnetworkcloud.ClustersClientBeginUpdateOptions
		if ifMatchParam != nil || ifNoneMatchParam != nil {
			options = &armnetworkcloud.ClustersClientBeginUpdateOptions{
				IfMatch:     ifMatchParam,
				IfNoneMatch: ifNoneMatchParam,
			}
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		c.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}

func (c *ClustersServerTransport) dispatchBeginUpdateVersion(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdateVersion == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateVersion not implemented")}
	}
	beginUpdateVersion := c.beginUpdateVersion.get(req)
	if beginUpdateVersion == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkCloud/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateVersion`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkcloud.ClusterUpdateVersionParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdateVersion(req.Context(), resourceGroupNameParam, clusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateVersion = &respr
		c.beginUpdateVersion.add(req, beginUpdateVersion)
	}

	resp, err := server.PollerResponderNext(beginUpdateVersion, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdateVersion.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateVersion) {
		c.beginUpdateVersion.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ClustersServerTransport
var clustersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
