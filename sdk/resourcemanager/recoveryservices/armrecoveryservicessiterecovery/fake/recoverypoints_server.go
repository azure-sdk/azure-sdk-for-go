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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v3"
	"net/http"
	"net/url"
	"regexp"
)

// RecoveryPointsServer is a fake server for instances of the armrecoveryservicessiterecovery.RecoveryPointsClient type.
type RecoveryPointsServer struct {
	// Get is the fake for method RecoveryPointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string, options *armrecoveryservicessiterecovery.RecoveryPointsClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.RecoveryPointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByReplicationProtectedItemsPager is the fake for method RecoveryPointsClient.NewListByReplicationProtectedItemsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReplicationProtectedItemsPager func(resourceName string, resourceGroupName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, options *armrecoveryservicessiterecovery.RecoveryPointsClientListByReplicationProtectedItemsOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.RecoveryPointsClientListByReplicationProtectedItemsResponse])
}

// NewRecoveryPointsServerTransport creates a new instance of RecoveryPointsServerTransport with the provided implementation.
// The returned RecoveryPointsServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.RecoveryPointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRecoveryPointsServerTransport(srv *RecoveryPointsServer) *RecoveryPointsServerTransport {
	return &RecoveryPointsServerTransport{
		srv:                                     srv,
		newListByReplicationProtectedItemsPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.RecoveryPointsClientListByReplicationProtectedItemsResponse]](),
	}
}

// RecoveryPointsServerTransport connects instances of armrecoveryservicessiterecovery.RecoveryPointsClient to instances of RecoveryPointsServer.
// Don't use this type directly, use NewRecoveryPointsServerTransport instead.
type RecoveryPointsServerTransport struct {
	srv                                     *RecoveryPointsServer
	newListByReplicationProtectedItemsPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.RecoveryPointsClientListByReplicationProtectedItemsResponse]]
}

// Do implements the policy.Transporter interface for RecoveryPointsServerTransport.
func (r *RecoveryPointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RecoveryPointsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RecoveryPointsClient.NewListByReplicationProtectedItemsPager":
		resp, err = r.dispatchNewListByReplicationProtectedItemsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RecoveryPointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionContainers/(?P<protectionContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectedItems/(?P<replicatedProtectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints/(?P<recoveryPointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	protectionContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectionContainerName")])
	if err != nil {
		return nil, err
	}
	replicatedProtectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicatedProtectedItemName")])
	if err != nil {
		return nil, err
	}
	recoveryPointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("recoveryPointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceNameParam, resourceGroupNameParam, fabricNameParam, protectionContainerNameParam, replicatedProtectedItemNameParam, recoveryPointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RecoveryPoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RecoveryPointsServerTransport) dispatchNewListByReplicationProtectedItemsPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByReplicationProtectedItemsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReplicationProtectedItemsPager not implemented")}
	}
	newListByReplicationProtectedItemsPager := r.newListByReplicationProtectedItemsPager.get(req)
	if newListByReplicationProtectedItemsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionContainers/(?P<protectionContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectedItems/(?P<replicatedProtectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		protectionContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectionContainerName")])
		if err != nil {
			return nil, err
		}
		replicatedProtectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicatedProtectedItemName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByReplicationProtectedItemsPager(resourceNameParam, resourceGroupNameParam, fabricNameParam, protectionContainerNameParam, replicatedProtectedItemNameParam, nil)
		newListByReplicationProtectedItemsPager = &resp
		r.newListByReplicationProtectedItemsPager.add(req, newListByReplicationProtectedItemsPager)
		server.PagerResponderInjectNextLinks(newListByReplicationProtectedItemsPager, req, func(page *armrecoveryservicessiterecovery.RecoveryPointsClientListByReplicationProtectedItemsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReplicationProtectedItemsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByReplicationProtectedItemsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReplicationProtectedItemsPager) {
		r.newListByReplicationProtectedItemsPager.remove(req)
	}
	return resp, nil
}
