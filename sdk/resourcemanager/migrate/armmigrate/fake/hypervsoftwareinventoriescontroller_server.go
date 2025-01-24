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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// HypervSoftwareInventoriesControllerServer is a fake server for instances of the armmigrate.HypervSoftwareInventoriesControllerClient type.
type HypervSoftwareInventoriesControllerServer struct {
	// GetMachineSoftwareInventory is the fake for method HypervSoftwareInventoriesControllerClient.GetMachineSoftwareInventory
	// HTTP status codes to indicate success: http.StatusOK
	GetMachineSoftwareInventory func(ctx context.Context, resourceGroupName string, siteName string, machineName string, defaultParam armmigrate.Default, options *armmigrate.HypervSoftwareInventoriesControllerClientGetMachineSoftwareInventoryOptions) (resp azfake.Responder[armmigrate.HypervSoftwareInventoriesControllerClientGetMachineSoftwareInventoryResponse], errResp azfake.ErrorResponder)

	// NewListByHypervMachinePager is the fake for method HypervSoftwareInventoriesControllerClient.NewListByHypervMachinePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHypervMachinePager func(resourceGroupName string, siteName string, machineName string, options *armmigrate.HypervSoftwareInventoriesControllerClientListByHypervMachineOptions) (resp azfake.PagerResponder[armmigrate.HypervSoftwareInventoriesControllerClientListByHypervMachineResponse])
}

// NewHypervSoftwareInventoriesControllerServerTransport creates a new instance of HypervSoftwareInventoriesControllerServerTransport with the provided implementation.
// The returned HypervSoftwareInventoriesControllerServerTransport instance is connected to an instance of armmigrate.HypervSoftwareInventoriesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervSoftwareInventoriesControllerServerTransport(srv *HypervSoftwareInventoriesControllerServer) *HypervSoftwareInventoriesControllerServerTransport {
	return &HypervSoftwareInventoriesControllerServerTransport{
		srv:                         srv,
		newListByHypervMachinePager: newTracker[azfake.PagerResponder[armmigrate.HypervSoftwareInventoriesControllerClientListByHypervMachineResponse]](),
	}
}

// HypervSoftwareInventoriesControllerServerTransport connects instances of armmigrate.HypervSoftwareInventoriesControllerClient to instances of HypervSoftwareInventoriesControllerServer.
// Don't use this type directly, use NewHypervSoftwareInventoriesControllerServerTransport instead.
type HypervSoftwareInventoriesControllerServerTransport struct {
	srv                         *HypervSoftwareInventoriesControllerServer
	newListByHypervMachinePager *tracker[azfake.PagerResponder[armmigrate.HypervSoftwareInventoriesControllerClientListByHypervMachineResponse]]
}

// Do implements the policy.Transporter interface for HypervSoftwareInventoriesControllerServerTransport.
func (h *HypervSoftwareInventoriesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HypervSoftwareInventoriesControllerClient.GetMachineSoftwareInventory":
		resp, err = h.dispatchGetMachineSoftwareInventory(req)
	case "HypervSoftwareInventoriesControllerClient.NewListByHypervMachinePager":
		resp, err = h.dispatchNewListByHypervMachinePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HypervSoftwareInventoriesControllerServerTransport) dispatchGetMachineSoftwareInventory(req *http.Request) (*http.Response, error) {
	if h.srv.GetMachineSoftwareInventory == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetMachineSoftwareInventory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/softwareInventories/(?P<default>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := h.srv.GetMachineSoftwareInventory(req.Context(), resourceGroupNameParam, siteNameParam, machineNameParam, defaultParamParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervVMSoftwareInventory, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervSoftwareInventoriesControllerServerTransport) dispatchNewListByHypervMachinePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByHypervMachinePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHypervMachinePager not implemented")}
	}
	newListByHypervMachinePager := h.newListByHypervMachinePager.get(req)
	if newListByHypervMachinePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/softwareinventories`
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
		resp := h.srv.NewListByHypervMachinePager(resourceGroupNameParam, siteNameParam, machineNameParam, nil)
		newListByHypervMachinePager = &resp
		h.newListByHypervMachinePager.add(req, newListByHypervMachinePager)
		server.PagerResponderInjectNextLinks(newListByHypervMachinePager, req, func(page *armmigrate.HypervSoftwareInventoriesControllerClientListByHypervMachineResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHypervMachinePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByHypervMachinePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHypervMachinePager) {
		h.newListByHypervMachinePager.remove(req)
	}
	return resp, nil
}
