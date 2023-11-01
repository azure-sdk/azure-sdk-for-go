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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineScaleSetApplicationsServer is a fake server for instances of the armcompute.VirtualMachineScaleSetApplicationsClient type.
type VirtualMachineScaleSetApplicationsServer struct {
	// BeginDelete is the fake for method VirtualMachineScaleSetApplicationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *armcompute.VirtualMachineScaleSetApplicationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetApplicationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineScaleSetApplicationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, options *armcompute.VirtualMachineScaleSetApplicationsClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetApplicationsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method VirtualMachineScaleSetApplicationsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, vmScaleSetName string, options *armcompute.VirtualMachineScaleSetApplicationsClientListOptions) (resp azfake.Responder[armcompute.VirtualMachineScaleSetApplicationsClientListResponse], errResp azfake.ErrorResponder)

	// BeginPut is the fake for method VirtualMachineScaleSetApplicationsClient.BeginPut
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginPut func(ctx context.Context, resourceGroupName string, vmScaleSetName string, applicationName string, application armcompute.VMApplicationProxyResource, options *armcompute.VirtualMachineScaleSetApplicationsClientBeginPutOptions) (resp azfake.PollerResponder[armcompute.VirtualMachineScaleSetApplicationsClientPutResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineScaleSetApplicationsServerTransport creates a new instance of VirtualMachineScaleSetApplicationsServerTransport with the provided implementation.
// The returned VirtualMachineScaleSetApplicationsServerTransport instance is connected to an instance of armcompute.VirtualMachineScaleSetApplicationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineScaleSetApplicationsServerTransport(srv *VirtualMachineScaleSetApplicationsServer) *VirtualMachineScaleSetApplicationsServerTransport {
	return &VirtualMachineScaleSetApplicationsServerTransport{
		srv:         srv,
		beginDelete: newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetApplicationsClientDeleteResponse]](),
		beginPut:    newTracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetApplicationsClientPutResponse]](),
	}
}

// VirtualMachineScaleSetApplicationsServerTransport connects instances of armcompute.VirtualMachineScaleSetApplicationsClient to instances of VirtualMachineScaleSetApplicationsServer.
// Don't use this type directly, use NewVirtualMachineScaleSetApplicationsServerTransport instead.
type VirtualMachineScaleSetApplicationsServerTransport struct {
	srv         *VirtualMachineScaleSetApplicationsServer
	beginDelete *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetApplicationsClientDeleteResponse]]
	beginPut    *tracker[azfake.PollerResponder[armcompute.VirtualMachineScaleSetApplicationsClientPutResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineScaleSetApplicationsServerTransport.
func (v *VirtualMachineScaleSetApplicationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineScaleSetApplicationsClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualMachineScaleSetApplicationsClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineScaleSetApplicationsClient.List":
		resp, err = v.dispatchList(req)
	case "VirtualMachineScaleSetApplicationsClient.BeginPut":
		resp, err = v.dispatchBeginPut(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetApplicationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, applicationNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineScaleSetApplicationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, applicationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VMApplicationProxyResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetApplicationsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.List(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineApplicationsProxyResourceListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineScaleSetApplicationsServerTransport) dispatchBeginPut(req *http.Request) (*http.Response, error) {
	if v.srv.BeginPut == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPut not implemented")}
	}
	beginPut := v.beginPut.get(req)
	if beginPut == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/virtualMachineScaleSets/(?P<vmScaleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applications/(?P<applicationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.VMApplicationProxyResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vmScaleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmScaleSetName")])
		if err != nil {
			return nil, err
		}
		applicationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginPut(req.Context(), resourceGroupNameParam, vmScaleSetNameParam, applicationNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPut = &respr
		v.beginPut.add(req, beginPut)
	}

	resp, err := server.PollerResponderNext(beginPut, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginPut.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPut) {
		v.beginPut.remove(req)
	}

	return resp, nil
}
