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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectedVMwarevSphereAssignmentsReportsServer is a fake server for instances of the armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClient type.
type ConnectedVMwarevSphereAssignmentsReportsServer struct {
	// Get is the fake for method ConnectedVMwarevSphereAssignmentsReportsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vmName string, guestConfigurationAssignmentName string, reportID string, options *armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClientGetOptions) (resp azfake.Responder[armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ConnectedVMwarevSphereAssignmentsReportsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, vmName string, guestConfigurationAssignmentName string, options *armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClientListOptions) (resp azfake.Responder[armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClientListResponse], errResp azfake.ErrorResponder)
}

// NewConnectedVMwarevSphereAssignmentsReportsServerTransport creates a new instance of ConnectedVMwarevSphereAssignmentsReportsServerTransport with the provided implementation.
// The returned ConnectedVMwarevSphereAssignmentsReportsServerTransport instance is connected to an instance of armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectedVMwarevSphereAssignmentsReportsServerTransport(srv *ConnectedVMwarevSphereAssignmentsReportsServer) *ConnectedVMwarevSphereAssignmentsReportsServerTransport {
	return &ConnectedVMwarevSphereAssignmentsReportsServerTransport{srv: srv}
}

// ConnectedVMwarevSphereAssignmentsReportsServerTransport connects instances of armguestconfiguration.ConnectedVMwarevSphereAssignmentsReportsClient to instances of ConnectedVMwarevSphereAssignmentsReportsServer.
// Don't use this type directly, use NewConnectedVMwarevSphereAssignmentsReportsServerTransport instead.
type ConnectedVMwarevSphereAssignmentsReportsServerTransport struct {
	srv *ConnectedVMwarevSphereAssignmentsReportsServer
}

// Do implements the policy.Transporter interface for ConnectedVMwarevSphereAssignmentsReportsServerTransport.
func (c *ConnectedVMwarevSphereAssignmentsReportsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectedVMwarevSphereAssignmentsReportsClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectedVMwarevSphereAssignmentsReportsClient.List":
		resp, err = c.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectedVMwarevSphereAssignmentsReportsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualmachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports/(?P<reportId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
	if err != nil {
		return nil, err
	}
	guestConfigurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("guestConfigurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	reportIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, vmNameParam, guestConfigurationAssignmentNameParam, reportIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssignmentReport, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectedVMwarevSphereAssignmentsReportsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if c.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ConnectedVMwarevSphere/virtualmachines/(?P<vmName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.GuestConfiguration/guestConfigurationAssignments/(?P<guestConfigurationAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reports`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vmNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vmName")])
	if err != nil {
		return nil, err
	}
	guestConfigurationAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("guestConfigurationAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.List(req.Context(), resourceGroupNameParam, vmNameParam, guestConfigurationAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssignmentReportList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}