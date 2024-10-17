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

// ManagedEnvironmentDiagnosticsServer is a fake server for instances of the armappcontainers.ManagedEnvironmentDiagnosticsClient type.
type ManagedEnvironmentDiagnosticsServer struct {
	// GetDetector is the fake for method ManagedEnvironmentDiagnosticsClient.GetDetector
	// HTTP status codes to indicate success: http.StatusOK
	GetDetector func(ctx context.Context, resourceGroupName string, environmentName string, detectorName string, options *armappcontainers.ManagedEnvironmentDiagnosticsClientGetDetectorOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentDiagnosticsClientGetDetectorResponse], errResp azfake.ErrorResponder)

	// ListDetectors is the fake for method ManagedEnvironmentDiagnosticsClient.ListDetectors
	// HTTP status codes to indicate success: http.StatusOK
	ListDetectors func(ctx context.Context, resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentDiagnosticsClientListDetectorsOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentDiagnosticsClientListDetectorsResponse], errResp azfake.ErrorResponder)
}

// NewManagedEnvironmentDiagnosticsServerTransport creates a new instance of ManagedEnvironmentDiagnosticsServerTransport with the provided implementation.
// The returned ManagedEnvironmentDiagnosticsServerTransport instance is connected to an instance of armappcontainers.ManagedEnvironmentDiagnosticsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedEnvironmentDiagnosticsServerTransport(srv *ManagedEnvironmentDiagnosticsServer) *ManagedEnvironmentDiagnosticsServerTransport {
	return &ManagedEnvironmentDiagnosticsServerTransport{srv: srv}
}

// ManagedEnvironmentDiagnosticsServerTransport connects instances of armappcontainers.ManagedEnvironmentDiagnosticsClient to instances of ManagedEnvironmentDiagnosticsServer.
// Don't use this type directly, use NewManagedEnvironmentDiagnosticsServerTransport instead.
type ManagedEnvironmentDiagnosticsServerTransport struct {
	srv *ManagedEnvironmentDiagnosticsServer
}

// Do implements the policy.Transporter interface for ManagedEnvironmentDiagnosticsServerTransport.
func (m *ManagedEnvironmentDiagnosticsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedEnvironmentDiagnosticsClient.GetDetector":
		resp, err = m.dispatchGetDetector(req)
	case "ManagedEnvironmentDiagnosticsClient.ListDetectors":
		resp, err = m.dispatchListDetectors(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedEnvironmentDiagnosticsServerTransport) dispatchGetDetector(req *http.Request) (*http.Response, error) {
	if m.srv.GetDetector == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDetector not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectors/(?P<detectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	detectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("detectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetDetector(req.Context(), resourceGroupNameParam, environmentNameParam, detectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Diagnostics, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedEnvironmentDiagnosticsServerTransport) dispatchListDetectors(req *http.Request) (*http.Response, error) {
	if m.srv.ListDetectors == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListDetectors not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detectors`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.ListDetectors(req.Context(), resourceGroupNameParam, environmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiagnosticsCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
