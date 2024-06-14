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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/monitor/armmonitor"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// DataCollectionEndpointsServer is a fake server for instances of the armmonitor.DataCollectionEndpointsClient type.
type DataCollectionEndpointsServer struct {
	// Create is the fake for method DataCollectionEndpointsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *armmonitor.DataCollectionEndpointsClientCreateOptions) (resp azfake.Responder[armmonitor.DataCollectionEndpointsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DataCollectionEndpointsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *armmonitor.DataCollectionEndpointsClientDeleteOptions) (resp azfake.Responder[armmonitor.DataCollectionEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DataCollectionEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *armmonitor.DataCollectionEndpointsClientGetOptions) (resp azfake.Responder[armmonitor.DataCollectionEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// GetNSP is the fake for method DataCollectionEndpointsClient.GetNSP
	// HTTP status codes to indicate success: http.StatusOK
	GetNSP func(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, networkSecurityPerimeterConfigurationName string, options *armmonitor.DataCollectionEndpointsClientGetNSPOptions) (resp azfake.Responder[armmonitor.DataCollectionEndpointsClientGetNSPResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method DataCollectionEndpointsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmonitor.DataCollectionEndpointsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method DataCollectionEndpointsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmonitor.DataCollectionEndpointsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListBySubscriptionResponse])

	// NewListNSPPager is the fake for method DataCollectionEndpointsClient.NewListNSPPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListNSPPager func(resourceGroupName string, dataCollectionEndpointName string, options *armmonitor.DataCollectionEndpointsClientListNSPOptions) (resp azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListNSPResponse])

	// BeginReconcileNSP is the fake for method DataCollectionEndpointsClient.BeginReconcileNSP
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginReconcileNSP func(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, networkSecurityPerimeterConfigurationName string, options *armmonitor.DataCollectionEndpointsClientBeginReconcileNSPOptions) (resp azfake.PollerResponder[armmonitor.DataCollectionEndpointsClientReconcileNSPResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method DataCollectionEndpointsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *armmonitor.DataCollectionEndpointsClientUpdateOptions) (resp azfake.Responder[armmonitor.DataCollectionEndpointsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDataCollectionEndpointsServerTransport creates a new instance of DataCollectionEndpointsServerTransport with the provided implementation.
// The returned DataCollectionEndpointsServerTransport instance is connected to an instance of armmonitor.DataCollectionEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDataCollectionEndpointsServerTransport(srv *DataCollectionEndpointsServer) *DataCollectionEndpointsServerTransport {
	return &DataCollectionEndpointsServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListBySubscriptionResponse]](),
		newListNSPPager:             newTracker[azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListNSPResponse]](),
		beginReconcileNSP:           newTracker[azfake.PollerResponder[armmonitor.DataCollectionEndpointsClientReconcileNSPResponse]](),
	}
}

// DataCollectionEndpointsServerTransport connects instances of armmonitor.DataCollectionEndpointsClient to instances of DataCollectionEndpointsServer.
// Don't use this type directly, use NewDataCollectionEndpointsServerTransport instead.
type DataCollectionEndpointsServerTransport struct {
	srv                         *DataCollectionEndpointsServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListBySubscriptionResponse]]
	newListNSPPager             *tracker[azfake.PagerResponder[armmonitor.DataCollectionEndpointsClientListNSPResponse]]
	beginReconcileNSP           *tracker[azfake.PollerResponder[armmonitor.DataCollectionEndpointsClientReconcileNSPResponse]]
}

// Do implements the policy.Transporter interface for DataCollectionEndpointsServerTransport.
func (d *DataCollectionEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DataCollectionEndpointsClient.Create":
		resp, err = d.dispatchCreate(req)
	case "DataCollectionEndpointsClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DataCollectionEndpointsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DataCollectionEndpointsClient.GetNSP":
		resp, err = d.dispatchGetNSP(req)
	case "DataCollectionEndpointsClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DataCollectionEndpointsClient.NewListBySubscriptionPager":
		resp, err = d.dispatchNewListBySubscriptionPager(req)
	case "DataCollectionEndpointsClient.NewListNSPPager":
		resp, err = d.dispatchNewListNSPPager(req)
	case "DataCollectionEndpointsClient.BeginReconcileNSP":
		resp, err = d.dispatchBeginReconcileNSP(req)
	case "DataCollectionEndpointsClient.Update":
		resp, err = d.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if d.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.DataCollectionEndpointResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
	if err != nil {
		return nil, err
	}
	var options *armmonitor.DataCollectionEndpointsClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmonitor.DataCollectionEndpointsClientCreateOptions{
			Body: &body,
		}
	}
	respr, errRespr := d.srv.Create(req.Context(), resourceGroupNameParam, dataCollectionEndpointNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataCollectionEndpointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, dataCollectionEndpointNameParam, nil)
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

func (d *DataCollectionEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, dataCollectionEndpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataCollectionEndpointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchGetNSP(req *http.Request) (*http.Response, error) {
	if d.srv.GetNSP == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetNSP not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<networkSecurityPerimeterConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.GetNSP(req.Context(), resourceGroupNameParam, dataCollectionEndpointNameParam, networkSecurityPerimeterConfigurationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NetworkSecurityPerimeterConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := d.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		d.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmonitor.DataCollectionEndpointsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		d.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := d.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		d.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmonitor.DataCollectionEndpointsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		d.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchNewListNSPPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListNSPPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListNSPPager not implemented")}
	}
	newListNSPPager := d.newListNSPPager.get(req)
	if newListNSPPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListNSPPager(resourceGroupNameParam, dataCollectionEndpointNameParam, nil)
		newListNSPPager = &resp
		d.newListNSPPager.add(req, newListNSPPager)
		server.PagerResponderInjectNextLinks(newListNSPPager, req, func(page *armmonitor.DataCollectionEndpointsClientListNSPResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListNSPPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListNSPPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListNSPPager) {
		d.newListNSPPager.remove(req)
	}
	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchBeginReconcileNSP(req *http.Request) (*http.Response, error) {
	if d.srv.BeginReconcileNSP == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReconcileNSP not implemented")}
	}
	beginReconcileNSP := d.beginReconcileNSP.get(req)
	if beginReconcileNSP == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/networkSecurityPerimeterConfigurations/(?P<networkSecurityPerimeterConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reconcile`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
		if err != nil {
			return nil, err
		}
		networkSecurityPerimeterConfigurationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginReconcileNSP(req.Context(), resourceGroupNameParam, dataCollectionEndpointNameParam, networkSecurityPerimeterConfigurationNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReconcileNSP = &respr
		d.beginReconcileNSP.add(req, beginReconcileNSP)
	}

	resp, err := server.PollerResponderNext(beginReconcileNSP, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		d.beginReconcileNSP.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReconcileNSP) {
		d.beginReconcileNSP.remove(req)
	}

	return resp, nil
}

func (d *DataCollectionEndpointsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Insights/dataCollectionEndpoints/(?P<dataCollectionEndpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmonitor.ResourceForUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataCollectionEndpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataCollectionEndpointName")])
	if err != nil {
		return nil, err
	}
	var options *armmonitor.DataCollectionEndpointsClientUpdateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armmonitor.DataCollectionEndpointsClientUpdateOptions{
			Body: &body,
		}
	}
	respr, errRespr := d.srv.Update(req.Context(), resourceGroupNameParam, dataCollectionEndpointNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataCollectionEndpointResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
