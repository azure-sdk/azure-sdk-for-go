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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"net/http"
	"net/url"
	"regexp"
)

// AutonomousDatabasesServer is a fake server for instances of the armoracledatabase.AutonomousDatabasesClient type.
type AutonomousDatabasesServer struct {
	// BeginCreateOrUpdate is the fake for method AutonomousDatabasesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, resource armoracledatabase.AutonomousDatabase, options *armoracledatabase.AutonomousDatabasesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AutonomousDatabasesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, options *armoracledatabase.AutonomousDatabasesClientBeginDeleteOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFailover is the fake for method AutonomousDatabasesClient.BeginFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailover func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, body armoracledatabase.PeerDbDetails, options *armoracledatabase.AutonomousDatabasesClientBeginFailoverOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientFailoverResponse], errResp azfake.ErrorResponder)

	// GenerateWallet is the fake for method AutonomousDatabasesClient.GenerateWallet
	// HTTP status codes to indicate success: http.StatusOK
	GenerateWallet func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, body armoracledatabase.GenerateAutonomousDatabaseWalletDetails, options *armoracledatabase.AutonomousDatabasesClientGenerateWalletOptions) (resp azfake.Responder[armoracledatabase.AutonomousDatabasesClientGenerateWalletResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AutonomousDatabasesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, options *armoracledatabase.AutonomousDatabasesClientGetOptions) (resp azfake.Responder[armoracledatabase.AutonomousDatabasesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AutonomousDatabasesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armoracledatabase.AutonomousDatabasesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armoracledatabase.AutonomousDatabasesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method AutonomousDatabasesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armoracledatabase.AutonomousDatabasesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armoracledatabase.AutonomousDatabasesClientListBySubscriptionResponse])

	// BeginRestore is the fake for method AutonomousDatabasesClient.BeginRestore
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestore func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, body armoracledatabase.RestoreAutonomousDatabaseDetails, options *armoracledatabase.AutonomousDatabasesClientBeginRestoreOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientRestoreResponse], errResp azfake.ErrorResponder)

	// BeginShrink is the fake for method AutonomousDatabasesClient.BeginShrink
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginShrink func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, options *armoracledatabase.AutonomousDatabasesClientBeginShrinkOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientShrinkResponse], errResp azfake.ErrorResponder)

	// BeginSwitchover is the fake for method AutonomousDatabasesClient.BeginSwitchover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSwitchover func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, body armoracledatabase.PeerDbDetails, options *armoracledatabase.AutonomousDatabasesClientBeginSwitchoverOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientSwitchoverResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method AutonomousDatabasesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, autonomousdatabasename string, properties armoracledatabase.AutonomousDatabaseUpdate, options *armoracledatabase.AutonomousDatabasesClientBeginUpdateOptions) (resp azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAutonomousDatabasesServerTransport creates a new instance of AutonomousDatabasesServerTransport with the provided implementation.
// The returned AutonomousDatabasesServerTransport instance is connected to an instance of armoracledatabase.AutonomousDatabasesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAutonomousDatabasesServerTransport(srv *AutonomousDatabasesServer) *AutonomousDatabasesServerTransport {
	return &AutonomousDatabasesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientDeleteResponse]](),
		beginFailover:               newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientFailoverResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armoracledatabase.AutonomousDatabasesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armoracledatabase.AutonomousDatabasesClientListBySubscriptionResponse]](),
		beginRestore:                newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientRestoreResponse]](),
		beginShrink:                 newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientShrinkResponse]](),
		beginSwitchover:             newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientSwitchoverResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientUpdateResponse]](),
	}
}

// AutonomousDatabasesServerTransport connects instances of armoracledatabase.AutonomousDatabasesClient to instances of AutonomousDatabasesServer.
// Don't use this type directly, use NewAutonomousDatabasesServerTransport instead.
type AutonomousDatabasesServerTransport struct {
	srv                         *AutonomousDatabasesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientDeleteResponse]]
	beginFailover               *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientFailoverResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armoracledatabase.AutonomousDatabasesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armoracledatabase.AutonomousDatabasesClientListBySubscriptionResponse]]
	beginRestore                *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientRestoreResponse]]
	beginShrink                 *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientShrinkResponse]]
	beginSwitchover             *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientSwitchoverResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armoracledatabase.AutonomousDatabasesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for AutonomousDatabasesServerTransport.
func (a *AutonomousDatabasesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AutonomousDatabasesClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AutonomousDatabasesClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AutonomousDatabasesClient.BeginFailover":
		resp, err = a.dispatchBeginFailover(req)
	case "AutonomousDatabasesClient.GenerateWallet":
		resp, err = a.dispatchGenerateWallet(req)
	case "AutonomousDatabasesClient.Get":
		resp, err = a.dispatchGet(req)
	case "AutonomousDatabasesClient.NewListByResourceGroupPager":
		resp, err = a.dispatchNewListByResourceGroupPager(req)
	case "AutonomousDatabasesClient.NewListBySubscriptionPager":
		resp, err = a.dispatchNewListBySubscriptionPager(req)
	case "AutonomousDatabasesClient.BeginRestore":
		resp, err = a.dispatchBeginRestore(req)
	case "AutonomousDatabasesClient.BeginShrink":
		resp, err = a.dispatchBeginShrink(req)
	case "AutonomousDatabasesClient.BeginSwitchover":
		resp, err = a.dispatchBeginSwitchover(req)
	case "AutonomousDatabasesClient.BeginUpdate":
		resp, err = a.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.AutonomousDatabase](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginFailover(req *http.Request) (*http.Response, error) {
	if a.srv.BeginFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFailover not implemented")}
	}
	beginFailover := a.beginFailover.get(req)
	if beginFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.PeerDbDetails](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginFailover(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFailover = &respr
		a.beginFailover.add(req, beginFailover)
	}

	resp, err := server.PollerResponderNext(beginFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFailover) {
		a.beginFailover.remove(req)
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchGenerateWallet(req *http.Request) (*http.Response, error) {
	if a.srv.GenerateWallet == nil {
		return nil, &nonRetriableError{errors.New("fake for method GenerateWallet not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/generateWallet`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armoracledatabase.GenerateAutonomousDatabaseWalletDetails](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GenerateWallet(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AutonomousDatabaseWalletFile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AutonomousDatabase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armoracledatabase.AutonomousDatabasesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armoracledatabase.AutonomousDatabasesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		a.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginRestore(req *http.Request) (*http.Response, error) {
	if a.srv.BeginRestore == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestore not implemented")}
	}
	beginRestore := a.beginRestore.get(req)
	if beginRestore == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restore`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.RestoreAutonomousDatabaseDetails](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginRestore(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestore = &respr
		a.beginRestore.add(req, beginRestore)
	}

	resp, err := server.PollerResponderNext(beginRestore, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginRestore.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestore) {
		a.beginRestore.remove(req)
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginShrink(req *http.Request) (*http.Response, error) {
	if a.srv.BeginShrink == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginShrink not implemented")}
	}
	beginShrink := a.beginShrink.get(req)
	if beginShrink == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shrink`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginShrink(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginShrink = &respr
		a.beginShrink.add(req, beginShrink)
	}

	resp, err := server.PollerResponderNext(beginShrink, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginShrink.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginShrink) {
		a.beginShrink.remove(req)
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginSwitchover(req *http.Request) (*http.Response, error) {
	if a.srv.BeginSwitchover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSwitchover not implemented")}
	}
	beginSwitchover := a.beginSwitchover.get(req)
	if beginSwitchover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/switchover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.PeerDbDetails](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginSwitchover(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSwitchover = &respr
		a.beginSwitchover.add(req, beginSwitchover)
	}

	resp, err := server.PollerResponderNext(beginSwitchover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginSwitchover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSwitchover) {
		a.beginSwitchover.remove(req)
	}

	return resp, nil
}

func (a *AutonomousDatabasesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := a.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/autonomousDatabases/(?P<autonomousdatabasename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armoracledatabase.AutonomousDatabaseUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		autonomousdatabasenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("autonomousdatabasename")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginUpdate(req.Context(), resourceGroupNameParam, autonomousdatabasenameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		a.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		a.beginUpdate.remove(req)
	}

	return resp, nil
}
