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

// ConfigurationProfileHCIAssignmentsServer is a fake server for instances of the armautomanage.ConfigurationProfileHCIAssignmentsClient type.
type ConfigurationProfileHCIAssignmentsServer struct {
	// CreateOrUpdate is the fake for method ConfigurationProfileHCIAssignmentsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, parameters armautomanage.ConfigurationProfileAssignment, options *armautomanage.ConfigurationProfileHCIAssignmentsClientCreateOrUpdateOptions) (resp azfake.Responder[armautomanage.ConfigurationProfileHCIAssignmentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ConfigurationProfileHCIAssignmentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *armautomanage.ConfigurationProfileHCIAssignmentsClientDeleteOptions) (resp azfake.Responder[armautomanage.ConfigurationProfileHCIAssignmentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConfigurationProfileHCIAssignmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, configurationProfileAssignmentName string, options *armautomanage.ConfigurationProfileHCIAssignmentsClientGetOptions) (resp azfake.Responder[armautomanage.ConfigurationProfileHCIAssignmentsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewConfigurationProfileHCIAssignmentsServerTransport creates a new instance of ConfigurationProfileHCIAssignmentsServerTransport with the provided implementation.
// The returned ConfigurationProfileHCIAssignmentsServerTransport instance is connected to an instance of armautomanage.ConfigurationProfileHCIAssignmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConfigurationProfileHCIAssignmentsServerTransport(srv *ConfigurationProfileHCIAssignmentsServer) *ConfigurationProfileHCIAssignmentsServerTransport {
	return &ConfigurationProfileHCIAssignmentsServerTransport{srv: srv}
}

// ConfigurationProfileHCIAssignmentsServerTransport connects instances of armautomanage.ConfigurationProfileHCIAssignmentsClient to instances of ConfigurationProfileHCIAssignmentsServer.
// Don't use this type directly, use NewConfigurationProfileHCIAssignmentsServerTransport instead.
type ConfigurationProfileHCIAssignmentsServerTransport struct {
	srv *ConfigurationProfileHCIAssignmentsServer
}

// Do implements the policy.Transporter interface for ConfigurationProfileHCIAssignmentsServerTransport.
func (c *ConfigurationProfileHCIAssignmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConfigurationProfileHCIAssignmentsClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ConfigurationProfileHCIAssignmentsClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ConfigurationProfileHCIAssignmentsClient.Get":
		resp, err = c.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConfigurationProfileHCIAssignmentsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHci/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armautomanage.ConfigurationProfileAssignment](req)
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
	configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, clusterNameParam, configurationProfileAssignmentNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationProfileAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConfigurationProfileHCIAssignmentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHci/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, clusterNameParam, configurationProfileAssignmentNameParam, nil)
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

func (c *ConfigurationProfileHCIAssignmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureStackHci/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automanage/configurationProfileAssignments/(?P<configurationProfileAssignmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	configurationProfileAssignmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("configurationProfileAssignmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, configurationProfileAssignmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ConfigurationProfileAssignment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
