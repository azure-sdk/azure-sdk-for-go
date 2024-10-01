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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate/v2"
	"net/http"
	"net/url"
	"regexp"
)

// VmwareSoftwareInventoriesControllerServer is a fake server for instances of the armmigrate.VmwareSoftwareInventoriesControllerClient type.
type VmwareSoftwareInventoriesControllerServer struct {
	// GetMachineSoftwareInventory is the fake for method VmwareSoftwareInventoriesControllerClient.GetMachineSoftwareInventory
	// HTTP status codes to indicate success: http.StatusOK
	GetMachineSoftwareInventory func(ctx context.Context, resourceGroupName string, siteName string, machineName string, defaultParam armmigrate.Default, options *armmigrate.VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryOptions) (resp azfake.Responder[armmigrate.VmwareSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse], errResp azfake.ErrorResponder)

	// NewListByMachineResourcePager is the fake for method VmwareSoftwareInventoriesControllerClient.NewListByMachineResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMachineResourcePager func(resourceGroupName string, siteName string, machineName string, options *armmigrate.VmwareSoftwareInventoriesControllerClientListByMachineResourceOptions) (resp azfake.PagerResponder[armmigrate.VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse])
}

// NewVmwareSoftwareInventoriesControllerServerTransport creates a new instance of VmwareSoftwareInventoriesControllerServerTransport with the provided implementation.
// The returned VmwareSoftwareInventoriesControllerServerTransport instance is connected to an instance of armmigrate.VmwareSoftwareInventoriesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVmwareSoftwareInventoriesControllerServerTransport(srv *VmwareSoftwareInventoriesControllerServer) *VmwareSoftwareInventoriesControllerServerTransport {
	return &VmwareSoftwareInventoriesControllerServerTransport{
		srv:                           srv,
		newListByMachineResourcePager: newTracker[azfake.PagerResponder[armmigrate.VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse]](),
	}
}

// VmwareSoftwareInventoriesControllerServerTransport connects instances of armmigrate.VmwareSoftwareInventoriesControllerClient to instances of VmwareSoftwareInventoriesControllerServer.
// Don't use this type directly, use NewVmwareSoftwareInventoriesControllerServerTransport instead.
type VmwareSoftwareInventoriesControllerServerTransport struct {
	srv                           *VmwareSoftwareInventoriesControllerServer
	newListByMachineResourcePager *tracker[azfake.PagerResponder[armmigrate.VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse]]
}

// Do implements the policy.Transporter interface for VmwareSoftwareInventoriesControllerServerTransport.
func (v *VmwareSoftwareInventoriesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VmwareSoftwareInventoriesControllerClient.GetMachineSoftwareInventory":
		resp, err = v.dispatchGetMachineSoftwareInventory(req)
	case "VmwareSoftwareInventoriesControllerClient.NewListByMachineResourcePager":
		resp, err = v.dispatchNewListByMachineResourcePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VmwareSoftwareInventoriesControllerServerTransport) dispatchGetMachineSoftwareInventory(req *http.Request) (*http.Response, error) {
	if v.srv.GetMachineSoftwareInventory == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMachineSoftwareInventory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/softwareInventories/(?P<default>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	defaultParamParam, err := parseWithCast(matches[regex.SubexpIndex("default")], func(v string) (armmigrate.Default, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armmigrate.Default(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetMachineSoftwareInventory(req.Context(), resourceGroupNameParam, siteNameParam, machineNameParam, defaultParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareMachineSoftwareInventory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VmwareSoftwareInventoriesControllerServerTransport) dispatchNewListByMachineResourcePager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByMachineResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMachineResourcePager not implemented")}
	}
	newListByMachineResourcePager := v.newListByMachineResourcePager.get(req)
	if newListByMachineResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/softwareinventories`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByMachineResourcePager(resourceGroupNameParam, siteNameParam, machineNameParam, nil)
		newListByMachineResourcePager = &resp
		v.newListByMachineResourcePager.add(req, newListByMachineResourcePager)
		server.PagerResponderInjectNextLinks(newListByMachineResourcePager, req, func(page *armmigrate.VmwareSoftwareInventoriesControllerClientListByMachineResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMachineResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByMachineResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMachineResourcePager) {
		v.newListByMachineResourcePager.remove(req)
	}
	return resp, nil
}
