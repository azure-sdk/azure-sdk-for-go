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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerAppsAPIServer is a fake server for instances of the armappcontainers.ContainerAppsAPIClient type.
type ContainerAppsAPIServer struct {
	// GetCustomDomainVerificationID is the fake for method ContainerAppsAPIClient.GetCustomDomainVerificationID
	// HTTP status codes to indicate success: http.StatusOK
	GetCustomDomainVerificationID func(ctx context.Context, options *armappcontainers.ContainerAppsAPIClientGetCustomDomainVerificationIDOptions) (resp azfake.Responder[armappcontainers.ContainerAppsAPIClientGetCustomDomainVerificationIDResponse], errResp azfake.ErrorResponder)

	// JobExecution is the fake for method ContainerAppsAPIClient.JobExecution
	// HTTP status codes to indicate success: http.StatusOK
	JobExecution func(ctx context.Context, resourceGroupName string, jobName string, jobExecutionName string, options *armappcontainers.ContainerAppsAPIClientJobExecutionOptions) (resp azfake.Responder[armappcontainers.ContainerAppsAPIClientJobExecutionResponse], errResp azfake.ErrorResponder)
}

// NewContainerAppsAPIServerTransport creates a new instance of ContainerAppsAPIServerTransport with the provided implementation.
// The returned ContainerAppsAPIServerTransport instance is connected to an instance of armappcontainers.ContainerAppsAPIClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsAPIServerTransport(srv *ContainerAppsAPIServer) *ContainerAppsAPIServerTransport {
	return &ContainerAppsAPIServerTransport{srv: srv}
}

// ContainerAppsAPIServerTransport connects instances of armappcontainers.ContainerAppsAPIClient to instances of ContainerAppsAPIServer.
// Don't use this type directly, use NewContainerAppsAPIServerTransport instead.
type ContainerAppsAPIServerTransport struct {
	srv *ContainerAppsAPIServer
}

// Do implements the policy.Transporter interface for ContainerAppsAPIServerTransport.
func (c *ContainerAppsAPIServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerAppsAPIClient.GetCustomDomainVerificationID":
		resp, err = c.dispatchGetCustomDomainVerificationID(req)
	case "ContainerAppsAPIClient.JobExecution":
		resp, err = c.dispatchJobExecution(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerAppsAPIServerTransport) dispatchGetCustomDomainVerificationID(req *http.Request) (*http.Response, error) {
	if c.srv.GetCustomDomainVerificationID == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetCustomDomainVerificationID not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/getCustomDomainVerificationId`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := c.srv.GetCustomDomainVerificationID(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsAPIServerTransport) dispatchJobExecution(req *http.Request) (*http.Response, error) {
	if c.srv.JobExecution == nil {
		return nil, &nonRetriableError{errors.New("fake for method JobExecution not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executions/(?P<jobExecutionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	jobExecutionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobExecutionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.JobExecution(req.Context(), resourceGroupNameParam, jobNameParam, jobExecutionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobExecution, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
