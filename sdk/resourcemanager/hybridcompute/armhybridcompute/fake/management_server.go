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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridcompute/armhybridcompute/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ManagementServer is a fake server for instances of the armhybridcompute.ManagementClient type.
type ManagementServer struct {
	// BeginSetupExtensions is the fake for method ManagementClient.BeginSetupExtensions
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSetupExtensions func(ctx context.Context, resourceGroupName string, machineName string, extensionSetupBody armhybridcompute.MachineExtensionSetup, options *armhybridcompute.ManagementClientBeginSetupExtensionsOptions) (resp azfake.PollerResponder[armhybridcompute.ManagementClientSetupExtensionsResponse], errResp azfake.ErrorResponder)

	// BeginUpgradeExtensions is the fake for method ManagementClient.BeginUpgradeExtensions
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpgradeExtensions func(ctx context.Context, resourceGroupName string, machineName string, extensionUpgradeParameters armhybridcompute.MachineExtensionUpgrade, options *armhybridcompute.ManagementClientBeginUpgradeExtensionsOptions) (resp azfake.PollerResponder[armhybridcompute.ManagementClientUpgradeExtensionsResponse], errResp azfake.ErrorResponder)
}

// NewManagementServerTransport creates a new instance of ManagementServerTransport with the provided implementation.
// The returned ManagementServerTransport instance is connected to an instance of armhybridcompute.ManagementClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagementServerTransport(srv *ManagementServer) *ManagementServerTransport {
	return &ManagementServerTransport{
		srv:                    srv,
		beginSetupExtensions:   newTracker[azfake.PollerResponder[armhybridcompute.ManagementClientSetupExtensionsResponse]](),
		beginUpgradeExtensions: newTracker[azfake.PollerResponder[armhybridcompute.ManagementClientUpgradeExtensionsResponse]](),
	}
}

// ManagementServerTransport connects instances of armhybridcompute.ManagementClient to instances of ManagementServer.
// Don't use this type directly, use NewManagementServerTransport instead.
type ManagementServerTransport struct {
	srv                    *ManagementServer
	beginSetupExtensions   *tracker[azfake.PollerResponder[armhybridcompute.ManagementClientSetupExtensionsResponse]]
	beginUpgradeExtensions *tracker[azfake.PollerResponder[armhybridcompute.ManagementClientUpgradeExtensionsResponse]]
}

// Do implements the policy.Transporter interface for ManagementServerTransport.
func (m *ManagementServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagementClient.BeginSetupExtensions":
		resp, err = m.dispatchBeginSetupExtensions(req)
	case "ManagementClient.BeginUpgradeExtensions":
		resp, err = m.dispatchBeginUpgradeExtensions(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagementServerTransport) dispatchBeginSetupExtensions(req *http.Request) (*http.Response, error) {
	if m.srv.BeginSetupExtensions == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSetupExtensions not implemented")}
	}
	beginSetupExtensions := m.beginSetupExtensions.get(req)
	if beginSetupExtensions == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/setupExtensions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcompute.MachineExtensionSetup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginSetupExtensions(req.Context(), resourceGroupNameParam, machineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSetupExtensions = &respr
		m.beginSetupExtensions.add(req, beginSetupExtensions)
	}

	resp, err := server.PollerResponderNext(beginSetupExtensions, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginSetupExtensions.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSetupExtensions) {
		m.beginSetupExtensions.remove(req)
	}

	return resp, nil
}

func (m *ManagementServerTransport) dispatchBeginUpgradeExtensions(req *http.Request) (*http.Response, error) {
	if m.srv.BeginUpgradeExtensions == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpgradeExtensions not implemented")}
	}
	beginUpgradeExtensions := m.beginUpgradeExtensions.get(req)
	if beginUpgradeExtensions == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HybridCompute/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upgradeExtensions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhybridcompute.MachineExtensionUpgrade](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginUpgradeExtensions(req.Context(), resourceGroupNameParam, machineNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpgradeExtensions = &respr
		m.beginUpgradeExtensions.add(req, beginUpgradeExtensions)
	}

	resp, err := server.PollerResponderNext(beginUpgradeExtensions, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginUpgradeExtensions.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpgradeExtensions) {
		m.beginUpgradeExtensions.remove(req)
	}

	return resp, nil
}
