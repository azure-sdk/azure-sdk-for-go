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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// ManagedNetworkProvisionsServer is a fake server for instances of the armmachinelearning.ManagedNetworkProvisionsClient type.
type ManagedNetworkProvisionsServer struct {
	// BeginProvisionManagedNetwork is the fake for method ManagedNetworkProvisionsClient.BeginProvisionManagedNetwork
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginProvisionManagedNetwork func(ctx context.Context, resourceGroupName string, workspaceName string, options *armmachinelearning.ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions) (resp azfake.PollerResponder[armmachinelearning.ManagedNetworkProvisionsClientProvisionManagedNetworkResponse], errResp azfake.ErrorResponder)
}

// NewManagedNetworkProvisionsServerTransport creates a new instance of ManagedNetworkProvisionsServerTransport with the provided implementation.
// The returned ManagedNetworkProvisionsServerTransport instance is connected to an instance of armmachinelearning.ManagedNetworkProvisionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedNetworkProvisionsServerTransport(srv *ManagedNetworkProvisionsServer) *ManagedNetworkProvisionsServerTransport {
	return &ManagedNetworkProvisionsServerTransport{
		srv:                          srv,
		beginProvisionManagedNetwork: newTracker[azfake.PollerResponder[armmachinelearning.ManagedNetworkProvisionsClientProvisionManagedNetworkResponse]](),
	}
}

// ManagedNetworkProvisionsServerTransport connects instances of armmachinelearning.ManagedNetworkProvisionsClient to instances of ManagedNetworkProvisionsServer.
// Don't use this type directly, use NewManagedNetworkProvisionsServerTransport instead.
type ManagedNetworkProvisionsServerTransport struct {
	srv                          *ManagedNetworkProvisionsServer
	beginProvisionManagedNetwork *tracker[azfake.PollerResponder[armmachinelearning.ManagedNetworkProvisionsClientProvisionManagedNetworkResponse]]
}

// Do implements the policy.Transporter interface for ManagedNetworkProvisionsServerTransport.
func (m *ManagedNetworkProvisionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedNetworkProvisionsClient.BeginProvisionManagedNetwork":
		resp, err = m.dispatchBeginProvisionManagedNetwork(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedNetworkProvisionsServerTransport) dispatchBeginProvisionManagedNetwork(req *http.Request) (*http.Response, error) {
	if m.srv.BeginProvisionManagedNetwork == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginProvisionManagedNetwork not implemented")}
	}
	beginProvisionManagedNetwork := m.beginProvisionManagedNetwork.get(req)
	if beginProvisionManagedNetwork == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/provisionManagedNetwork`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.ManagedNetworkProvisionOptions](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		var options *armmachinelearning.ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armmachinelearning.ManagedNetworkProvisionsClientBeginProvisionManagedNetworkOptions{
				Body: &body,
			}
		}
		respr, errRespr := m.srv.BeginProvisionManagedNetwork(req.Context(), resourceGroupNameParam, workspaceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginProvisionManagedNetwork = &respr
		m.beginProvisionManagedNetwork.add(req, beginProvisionManagedNetwork)
	}

	resp, err := server.PollerResponderNext(beginProvisionManagedNetwork, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginProvisionManagedNetwork.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginProvisionManagedNetwork) {
		m.beginProvisionManagedNetwork.remove(req)
	}

	return resp, nil
}
