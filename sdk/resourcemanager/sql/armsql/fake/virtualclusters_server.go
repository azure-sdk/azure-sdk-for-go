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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualClustersServer is a fake server for instances of the armsql.VirtualClustersClient type.
type VirtualClustersServer struct {
	// BeginDelete is the fake for method VirtualClustersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, virtualClusterName string, options *armsql.VirtualClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.VirtualClustersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualClustersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, virtualClusterName string, options *armsql.VirtualClustersClientGetOptions) (resp azfake.Responder[armsql.VirtualClustersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualClustersClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsql.VirtualClustersClientListOptions) (resp azfake.PagerResponder[armsql.VirtualClustersClientListResponse])

	// NewListByResourceGroupPager is the fake for method VirtualClustersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsql.VirtualClustersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsql.VirtualClustersClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method VirtualClustersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, virtualClusterName string, parameters armsql.VirtualClusterUpdate, options *armsql.VirtualClustersClientBeginUpdateOptions) (resp azfake.PollerResponder[armsql.VirtualClustersClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDNSServers is the fake for method VirtualClustersClient.BeginUpdateDNSServers
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateDNSServers func(ctx context.Context, resourceGroupName string, virtualClusterName string, options *armsql.VirtualClustersClientBeginUpdateDNSServersOptions) (resp azfake.PollerResponder[armsql.VirtualClustersClientUpdateDNSServersResponse], errResp azfake.ErrorResponder)
}

// NewVirtualClustersServerTransport creates a new instance of VirtualClustersServerTransport with the provided implementation.
// The returned VirtualClustersServerTransport instance is connected to an instance of armsql.VirtualClustersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualClustersServerTransport(srv *VirtualClustersServer) *VirtualClustersServerTransport {
	return &VirtualClustersServerTransport{
		srv:                         srv,
		beginDelete:                 newTracker[azfake.PollerResponder[armsql.VirtualClustersClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armsql.VirtualClustersClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armsql.VirtualClustersClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armsql.VirtualClustersClientUpdateResponse]](),
		beginUpdateDNSServers:       newTracker[azfake.PollerResponder[armsql.VirtualClustersClientUpdateDNSServersResponse]](),
	}
}

// VirtualClustersServerTransport connects instances of armsql.VirtualClustersClient to instances of VirtualClustersServer.
// Don't use this type directly, use NewVirtualClustersServerTransport instead.
type VirtualClustersServerTransport struct {
	srv                         *VirtualClustersServer
	beginDelete                 *tracker[azfake.PollerResponder[armsql.VirtualClustersClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armsql.VirtualClustersClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armsql.VirtualClustersClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armsql.VirtualClustersClientUpdateResponse]]
	beginUpdateDNSServers       *tracker[azfake.PollerResponder[armsql.VirtualClustersClientUpdateDNSServersResponse]]
}

// Do implements the policy.Transporter interface for VirtualClustersServerTransport.
func (v *VirtualClustersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualClustersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if virtualClustersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = virtualClustersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VirtualClustersClient.BeginDelete":
				res.resp, res.err = v.dispatchBeginDelete(req)
			case "VirtualClustersClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VirtualClustersClient.NewListPager":
				res.resp, res.err = v.dispatchNewListPager(req)
			case "VirtualClustersClient.NewListByResourceGroupPager":
				res.resp, res.err = v.dispatchNewListByResourceGroupPager(req)
			case "VirtualClustersClient.BeginUpdate":
				res.resp, res.err = v.dispatchBeginUpdate(req)
			case "VirtualClustersClient.BeginUpdateDNSServers":
				res.resp, res.err = v.dispatchBeginUpdateDNSServers(req)
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

func (v *VirtualClustersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/virtualClusters/(?P<virtualClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, virtualClusterNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualClustersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/virtualClusters/(?P<virtualClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	virtualClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualClusterName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, virtualClusterNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualClustersServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/virtualClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsql.VirtualClustersClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualClustersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := v.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/virtualClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		v.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsql.VirtualClustersClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		v.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualClustersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/virtualClusters/(?P<virtualClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.VirtualClusterUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, virtualClusterNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualClustersServerTransport) dispatchBeginUpdateDNSServers(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdateDNSServers == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateDNSServers not implemented")}
	}
	beginUpdateDNSServers := v.beginUpdateDNSServers.get(req)
	if beginUpdateDNSServers == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/virtualClusters/(?P<virtualClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateManagedInstanceDnsServers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualClusterName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdateDNSServers(req.Context(), resourceGroupNameParam, virtualClusterNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateDNSServers = &respr
		v.beginUpdateDNSServers.add(req, beginUpdateDNSServers)
	}

	resp, err := server.PollerResponderNext(beginUpdateDNSServers, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdateDNSServers.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateDNSServers) {
		v.beginUpdateDNSServers.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to VirtualClustersServerTransport
var virtualClustersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
