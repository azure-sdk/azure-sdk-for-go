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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"net/http"
	"net/url"
	"regexp"
)

// CloudVMClustersServer is a fake server for instances of the armoracledatabase.CloudVMClustersClient type.
type CloudVMClustersServer struct {
	// BeginAddVMs is the fake for method CloudVMClustersClient.BeginAddVMs
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginAddVMs func(ctx context.Context, resourceGroupName string, cloudvmclustername string, body armoracledatabase.AddRemoveDbNode, options *armoracledatabase.CloudVMClustersClientBeginAddVMsOptions) (resp azfake.PollerResponder[armoracledatabase.CloudVMClustersClientAddVMsResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method CloudVMClustersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, cloudvmclustername string, resource armoracledatabase.CloudVMCluster, options *armoracledatabase.CloudVMClustersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armoracledatabase.CloudVMClustersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CloudVMClustersClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, cloudvmclustername string, options *armoracledatabase.CloudVMClustersClientBeginDeleteOptions) (resp azfake.PollerResponder[armoracledatabase.CloudVMClustersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CloudVMClustersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cloudvmclustername string, options *armoracledatabase.CloudVMClustersClientGetOptions) (resp azfake.Responder[armoracledatabase.CloudVMClustersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method CloudVMClustersClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armoracledatabase.CloudVMClustersClientListByResourceGroupOptions) (resp azfake.PagerResponder[armoracledatabase.CloudVMClustersClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method CloudVMClustersClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armoracledatabase.CloudVMClustersClientListBySubscriptionOptions) (resp azfake.PagerResponder[armoracledatabase.CloudVMClustersClientListBySubscriptionResponse])

	// ListPrivateIPAddresses is the fake for method CloudVMClustersClient.ListPrivateIPAddresses
	// HTTP status codes to indicate success: http.StatusOK
	ListPrivateIPAddresses func(ctx context.Context, resourceGroupName string, cloudvmclustername string, body armoracledatabase.PrivateIPAddressesFilter, options *armoracledatabase.CloudVMClustersClientListPrivateIPAddressesOptions) (resp azfake.Responder[armoracledatabase.CloudVMClustersClientListPrivateIPAddressesResponse], errResp azfake.ErrorResponder)

	// BeginRemoveVMs is the fake for method CloudVMClustersClient.BeginRemoveVMs
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRemoveVMs func(ctx context.Context, resourceGroupName string, cloudvmclustername string, body armoracledatabase.AddRemoveDbNode, options *armoracledatabase.CloudVMClustersClientBeginRemoveVMsOptions) (resp azfake.PollerResponder[armoracledatabase.CloudVMClustersClientRemoveVMsResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method CloudVMClustersClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, cloudvmclustername string, properties armoracledatabase.CloudVMClusterUpdate, options *armoracledatabase.CloudVMClustersClientBeginUpdateOptions) (resp azfake.PollerResponder[armoracledatabase.CloudVMClustersClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCloudVMClustersServerTransport creates a new instance of CloudVMClustersServerTransport with the provided implementation.
// The returned CloudVMClustersServerTransport instance is connected to an instance of armoracledatabase.CloudVMClustersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudVMClustersServerTransport(srv *CloudVMClustersServer) *CloudVMClustersServerTransport {
	return &CloudVMClustersServerTransport{
		srv:                         srv,
		beginAddVMs:                 newTracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientAddVMsResponse]](),
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armoracledatabase.CloudVMClustersClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armoracledatabase.CloudVMClustersClientListBySubscriptionResponse]](),
		beginRemoveVMs:              newTracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientRemoveVMsResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientUpdateResponse]](),
	}
}

// CloudVMClustersServerTransport connects instances of armoracledatabase.CloudVMClustersClient to instances of CloudVMClustersServer.
// Don't use this type directly, use NewCloudVMClustersServerTransport instead.
type CloudVMClustersServerTransport struct {
	srv                         *CloudVMClustersServer
	beginAddVMs                 *tracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientAddVMsResponse]]
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armoracledatabase.CloudVMClustersClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armoracledatabase.CloudVMClustersClientListBySubscriptionResponse]]
	beginRemoveVMs              *tracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientRemoveVMsResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armoracledatabase.CloudVMClustersClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for CloudVMClustersServerTransport.
func (c *CloudVMClustersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudVMClustersClient.BeginAddVMs":
		resp, err = c.dispatchBeginAddVMs(req)
	case "CloudVMClustersClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CloudVMClustersClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudVMClustersClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudVMClustersClient.NewListByResourceGroupPager":
		resp, err = c.dispatchNewListByResourceGroupPager(req)
	case "CloudVMClustersClient.NewListBySubscriptionPager":
		resp, err = c.dispatchNewListBySubscriptionPager(req)
	case "CloudVMClustersClient.ListPrivateIPAddresses":
		resp, err = c.dispatchListPrivateIPAddresses(req)
	case "CloudVMClustersClient.BeginRemoveVMs":
		resp, err = c.dispatchBeginRemoveVMs(req)
	case "CloudVMClustersClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudVMClustersServerTransport) dispatchBeginAddVMs(req *http.Request) (*http.Response, error) {
	if c.srv.BeginAddVMs == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginAddVMs not implemented")}
	}
	beginAddVMs := c.beginAddVMs.get(req)
	if beginAddVMs == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addVms`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.AddRemoveDbNode](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginAddVMs(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginAddVMs = &respr
		c.beginAddVMs.add(req, beginAddVMs)
	}

	resp, err := server.PollerResponderNext(beginAddVMs, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginAddVMs.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginAddVMs) {
		c.beginAddVMs.remove(req)
	}

	return resp, nil
}

func (c *CloudVMClustersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.CloudVMCluster](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, body, nil)
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

func (c *CloudVMClustersServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, nil)
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

func (c *CloudVMClustersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudVMCluster, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudVMClustersServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := c.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters`
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
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armoracledatabase.CloudVMClustersClientListByResourceGroupResponse, createLink func() string) {
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

func (c *CloudVMClustersServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armoracledatabase.CloudVMClustersClientListBySubscriptionResponse, createLink func() string) {
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

func (c *CloudVMClustersServerTransport) dispatchListPrivateIPAddresses(req *http.Request) (*http.Response, error) {
	if c.srv.ListPrivateIPAddresses == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListPrivateIPAddresses not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listPrivateIpAddresses`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoracledatabase.PrivateIPAddressesFilter](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.ListPrivateIPAddresses(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateIPAddressPropertiesArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudVMClustersServerTransport) dispatchBeginRemoveVMs(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRemoveVMs == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRemoveVMs not implemented")}
	}
	beginRemoveVMs := c.beginRemoveVMs.get(req)
	if beginRemoveVMs == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/removeVms`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.AddRemoveDbNode](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginRemoveVMs(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRemoveVMs = &respr
		c.beginRemoveVMs.add(req, beginRemoveVMs)
	}

	resp, err := server.PollerResponderNext(beginRemoveVMs, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginRemoveVMs.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRemoveVMs) {
		c.beginRemoveVMs.remove(req)
	}

	return resp, nil
}

func (c *CloudVMClustersServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/cloudVmClusters/(?P<cloudvmclustername>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.CloudVMClusterUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudvmclusternameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudvmclustername")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameParam, cloudvmclusternameParam, body, nil)
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
