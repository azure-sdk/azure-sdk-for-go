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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// CloudServicesServer is a fake server for instances of the armcompute.CloudServicesClient type.
type CloudServicesServer struct {
	// BeginCreateOrUpdate is the fake for method CloudServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, cloudServiceName string, parameters armcompute.CloudService, options *armcompute.CloudServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method CloudServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDeleteInstances is the fake for method CloudServicesClient.BeginDeleteInstances
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDeleteInstances func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginDeleteInstancesOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method CloudServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientGetOptions) (resp azfake.Responder[armcompute.CloudServicesClientGetResponse], errResp azfake.ErrorResponder)

	// GetInstanceView is the fake for method CloudServicesClient.GetInstanceView
	// HTTP status codes to indicate success: http.StatusOK
	GetInstanceView func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientGetInstanceViewOptions) (resp azfake.Responder[armcompute.CloudServicesClientGetInstanceViewResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method CloudServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armcompute.CloudServicesClientListOptions) (resp azfake.PagerResponder[armcompute.CloudServicesClientListResponse])

	// NewListAllPager is the fake for method CloudServicesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armcompute.CloudServicesClientListAllOptions) (resp azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse])

	// BeginPowerOff is the fake for method CloudServicesClient.BeginPowerOff
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPowerOff func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginPowerOffOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse], errResp azfake.ErrorResponder)

	// BeginRebuild is the fake for method CloudServicesClient.BeginRebuild
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRebuild func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginRebuildOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse], errResp azfake.ErrorResponder)

	// BeginReimage is the fake for method CloudServicesClient.BeginReimage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReimage func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginReimageOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse], errResp azfake.ErrorResponder)

	// BeginRestart is the fake for method CloudServicesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestart func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginRestartOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method CloudServicesClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, cloudServiceName string, options *armcompute.CloudServicesClientBeginStartOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientStartResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method CloudServicesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, cloudServiceName string, parameters armcompute.CloudServiceUpdate, options *armcompute.CloudServicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewCloudServicesServerTransport creates a new instance of CloudServicesServerTransport with the provided implementation.
// The returned CloudServicesServerTransport instance is connected to an instance of armcompute.CloudServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudServicesServerTransport(srv *CloudServicesServer) *CloudServicesServerTransport {
	return &CloudServicesServerTransport{
		srv:                  srv,
		beginCreateOrUpdate:  newTracker[azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse]](),
		beginDelete:          newTracker[azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse]](),
		beginDeleteInstances: newTracker[azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse]](),
		newListPager:         newTracker[azfake.PagerResponder[armcompute.CloudServicesClientListResponse]](),
		newListAllPager:      newTracker[azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse]](),
		beginPowerOff:        newTracker[azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse]](),
		beginRebuild:         newTracker[azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse]](),
		beginReimage:         newTracker[azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse]](),
		beginRestart:         newTracker[azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse]](),
		beginStart:           newTracker[azfake.PollerResponder[armcompute.CloudServicesClientStartResponse]](),
		beginUpdate:          newTracker[azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse]](),
	}
}

// CloudServicesServerTransport connects instances of armcompute.CloudServicesClient to instances of CloudServicesServer.
// Don't use this type directly, use NewCloudServicesServerTransport instead.
type CloudServicesServerTransport struct {
	srv                  *CloudServicesServer
	beginCreateOrUpdate  *tracker[azfake.PollerResponder[armcompute.CloudServicesClientCreateOrUpdateResponse]]
	beginDelete          *tracker[azfake.PollerResponder[armcompute.CloudServicesClientDeleteResponse]]
	beginDeleteInstances *tracker[azfake.PollerResponder[armcompute.CloudServicesClientDeleteInstancesResponse]]
	newListPager         *tracker[azfake.PagerResponder[armcompute.CloudServicesClientListResponse]]
	newListAllPager      *tracker[azfake.PagerResponder[armcompute.CloudServicesClientListAllResponse]]
	beginPowerOff        *tracker[azfake.PollerResponder[armcompute.CloudServicesClientPowerOffResponse]]
	beginRebuild         *tracker[azfake.PollerResponder[armcompute.CloudServicesClientRebuildResponse]]
	beginReimage         *tracker[azfake.PollerResponder[armcompute.CloudServicesClientReimageResponse]]
	beginRestart         *tracker[azfake.PollerResponder[armcompute.CloudServicesClientRestartResponse]]
	beginStart           *tracker[azfake.PollerResponder[armcompute.CloudServicesClientStartResponse]]
	beginUpdate          *tracker[azfake.PollerResponder[armcompute.CloudServicesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for CloudServicesServerTransport.
func (c *CloudServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CloudServicesClient.BeginCreateOrUpdate":
		resp, err = c.dispatchBeginCreateOrUpdate(req)
	case "CloudServicesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "CloudServicesClient.BeginDeleteInstances":
		resp, err = c.dispatchBeginDeleteInstances(req)
	case "CloudServicesClient.Get":
		resp, err = c.dispatchGet(req)
	case "CloudServicesClient.GetInstanceView":
		resp, err = c.dispatchGetInstanceView(req)
	case "CloudServicesClient.NewListPager":
		resp, err = c.dispatchNewListPager(req)
	case "CloudServicesClient.NewListAllPager":
		resp, err = c.dispatchNewListAllPager(req)
	case "CloudServicesClient.BeginPowerOff":
		resp, err = c.dispatchBeginPowerOff(req)
	case "CloudServicesClient.BeginRebuild":
		resp, err = c.dispatchBeginRebuild(req)
	case "CloudServicesClient.BeginReimage":
		resp, err = c.dispatchBeginReimage(req)
	case "CloudServicesClient.BeginRestart":
		resp, err = c.dispatchBeginRestart(req)
	case "CloudServicesClient.BeginStart":
		resp, err = c.dispatchBeginStart(req)
	case "CloudServicesClient.BeginUpdate":
		resp, err = c.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.CloudService](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, body, nil)
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

func (c *CloudServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
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

func (c *CloudServicesServerTransport) dispatchBeginDeleteInstances(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDeleteInstances == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteInstances not implemented")}
	}
	beginDeleteInstances := c.beginDeleteInstances.get(req)
	if beginDeleteInstances == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/delete`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginDeleteInstancesOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginDeleteInstancesOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginDeleteInstances(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteInstances = &respr
		c.beginDeleteInstances.add(req, beginDeleteInstances)
	}

	resp, err := server.PollerResponderNext(beginDeleteInstances, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginDeleteInstances.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteInstances) {
		c.beginDeleteInstances.remove(req)
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudService, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchGetInstanceView(req *http.Request) (*http.Response, error) {
	if c.srv.GetInstanceView == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetInstanceView not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/instanceView`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetInstanceView(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CloudServiceInstanceView, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListPager(resourceGroupNameUnescaped, nil)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.CloudServicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := c.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := c.srv.NewListAllPager(nil)
		newListAllPager = &resp
		c.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armcompute.CloudServicesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		c.newListAllPager.remove(req)
	}
	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginPowerOff(req *http.Request) (*http.Response, error) {
	if c.srv.BeginPowerOff == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPowerOff not implemented")}
	}
	beginPowerOff := c.beginPowerOff.get(req)
	if beginPowerOff == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/poweroff`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginPowerOff(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPowerOff = &respr
		c.beginPowerOff.add(req, beginPowerOff)
	}

	resp, err := server.PollerResponderNext(beginPowerOff, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginPowerOff.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPowerOff) {
		c.beginPowerOff.remove(req)
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginRebuild(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRebuild == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRebuild not implemented")}
	}
	beginRebuild := c.beginRebuild.get(req)
	if beginRebuild == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rebuild`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginRebuildOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginRebuildOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginRebuild(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRebuild = &respr
		c.beginRebuild.add(req, beginRebuild)
	}

	resp, err := server.PollerResponderNext(beginRebuild, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginRebuild.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRebuild) {
		c.beginRebuild.remove(req)
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginReimage(req *http.Request) (*http.Response, error) {
	if c.srv.BeginReimage == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReimage not implemented")}
	}
	beginReimage := c.beginReimage.get(req)
	if beginReimage == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reimage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginReimageOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginReimageOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginReimage(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReimage = &respr
		c.beginReimage.add(req, beginReimage)
	}

	resp, err := server.PollerResponderNext(beginReimage, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginReimage.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReimage) {
		c.beginReimage.remove(req)
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := c.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.RoleInstances](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		var options *armcompute.CloudServicesClientBeginRestartOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armcompute.CloudServicesClientBeginRestartOptions{
				Parameters: &body,
			}
		}
		respr, errRespr := c.srv.BeginRestart(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		c.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		c.beginRestart.remove(req)
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if c.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := c.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginStart(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		c.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		c.beginStart.remove(req)
	}

	return resp, nil
}

func (c *CloudServicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := c.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/cloudServices/(?P<cloudServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.CloudServiceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		cloudServiceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("cloudServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, cloudServiceNameUnescaped, body, nil)
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

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		c.beginUpdate.remove(req)
	}

	return resp, nil
}
