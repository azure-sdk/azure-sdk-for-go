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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automanage/armautomanage/v2"
	"net/http"
	"net/url"
	"regexp"
)

// HCRPReportsServer is a fake server for instances of the armautomanage.HCRPReportsClient type.
type HCRPReportsServer struct {
	// Get is the fake for method HCRPReportsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, machineName string, configurationProfileAssignmentName string, reportName string, options *armautomanage.HCRPReportsClientGetOptions) (resp azfake.Responder[armautomanage.HCRPReportsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByConfigurationProfileAssignmentsPager is the fake for method HCRPReportsClient.NewListByConfigurationProfileAssignmentsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByConfigurationProfileAssignmentsPager func(resourceGroupName string, machineName string, configurationProfileAssignmentName string, options *armautomanage.HCRPReportsClientListByConfigurationProfileAssignmentsOptions) (resp azfake.PagerResponder[armautomanage.HCRPReportsClientListByConfigurationProfileAssignmentsResponse])
}

// NewHCRPReportsServerTransport creates a new instance of HCRPReportsServerTransport with the provided implementation.
// The returned HCRPReportsServerTransport instance is connected to an instance of armautomanage.HCRPReportsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHCRPReportsServerTransport(srv *HCRPReportsServer) *HCRPReportsServerTransport {
	return &HCRPReportsServerTransport{
		srv: srv,
		newListByConfigurationProfileAssignmentsPager: newTracker[azfake.PagerResponder[armautomanage.HCRPReportsClientListByConfigurationProfileAssignmentsResponse]](),
	}
}

// HCRPReportsServerTransport connects instances of armautomanage.HCRPReportsClient to instances of HCRPReportsServer.
// Don't use this type directly, use NewHCRPReportsServerTransport instead.
type HCRPReportsServerTransport struct {
	srv                                           *HCRPReportsServer
	newListByConfigurationProfileAssignmentsPager *tracker[azfake.PagerResponder[armautomanage.HCRPReportsClientListByConfigurationProfileAssignmentsResponse]]
}

// Do implements the policy.Transporter interface for HCRPReportsServerTransport.
func (h *HCRPReportsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HCRPReportsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HCRPReportsClient.NewListByConfigurationProfileAssignmentsPager":
		resp, err = h.dispatchNewListByConfigurationProfileAssignmentsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HCRPReportsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
	if err != nil {
		return nil, err
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, machineNameParam, configurationProfileAssignmentNameParam, reportNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Report, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HCRPReportsServerTransport) dispatchNewListByConfigurationProfileAssignmentsPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByConfigurationProfileAssignmentsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByConfigurationProfileAssignmentsPager not implemented")}
	}
	newListByConfigurationProfileAssignmentsPager := h.newListByConfigurationProfileAssignmentsPager.get(req)
	if newListByConfigurationProfileAssignmentsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListByConfigurationProfileAssignmentsPager(resourceGroupNameParam, machineNameParam, configurationProfileAssignmentNameParam, nil)
		newListByConfigurationProfileAssignmentsPager = &resp
		h.newListByConfigurationProfileAssignmentsPager.add(req, newListByConfigurationProfileAssignmentsPager)
	}
	resp, err := server.PagerResponderNext(newListByConfigurationProfileAssignmentsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByConfigurationProfileAssignmentsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByConfigurationProfileAssignmentsPager) {
		h.newListByConfigurationProfileAssignmentsPager.remove(req)
	}
	return resp, nil
}
