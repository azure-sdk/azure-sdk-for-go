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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationVaultHealthServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationVaultHealthClient type.
type ReplicationVaultHealthServer struct {
	// Get is the fake for method ReplicationVaultHealthClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armrecoveryservicessiterecovery.ReplicationVaultHealthClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationVaultHealthClientGetResponse], errResp azfake.ErrorResponder)

	// BeginRefresh is the fake for method ReplicationVaultHealthClient.BeginRefresh
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefresh func(ctx context.Context, resourceGroupName string, resourceName string, options *armrecoveryservicessiterecovery.ReplicationVaultHealthClientBeginRefreshOptions) (resp azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationVaultHealthClientRefreshResponse], errResp azfake.ErrorResponder)
}

// NewReplicationVaultHealthServerTransport creates a new instance of ReplicationVaultHealthServerTransport with the provided implementation.
// The returned ReplicationVaultHealthServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationVaultHealthClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationVaultHealthServerTransport(srv *ReplicationVaultHealthServer) *ReplicationVaultHealthServerTransport {
	return &ReplicationVaultHealthServerTransport{
		srv:          srv,
		beginRefresh: newTracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationVaultHealthClientRefreshResponse]](),
	}
}

// ReplicationVaultHealthServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationVaultHealthClient to instances of ReplicationVaultHealthServer.
// Don't use this type directly, use NewReplicationVaultHealthServerTransport instead.
type ReplicationVaultHealthServerTransport struct {
	srv          *ReplicationVaultHealthServer
	beginRefresh *tracker[azfake.PollerResponder[armrecoveryservicessiterecovery.ReplicationVaultHealthClientRefreshResponse]]
}

// Do implements the policy.Transporter interface for ReplicationVaultHealthServerTransport.
func (r *ReplicationVaultHealthServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReplicationVaultHealthClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReplicationVaultHealthClient.BeginRefresh":
		resp, err = r.dispatchBeginRefresh(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReplicationVaultHealthServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationVaultHealth`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VaultHealthDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationVaultHealthServerTransport) dispatchBeginRefresh(req *http.Request) (*http.Response, error) {
	if r.srv.BeginRefresh == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefresh not implemented")}
	}
	beginRefresh := r.beginRefresh.get(req)
	if beginRefresh == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationVaultHealth/default/refresh`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginRefresh(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefresh = &respr
		r.beginRefresh.add(req, beginRefresh)
	}

	resp, err := server.PollerResponderNext(beginRefresh, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginRefresh.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefresh) {
		r.beginRefresh.remove(req)
	}

	return resp, nil
}
