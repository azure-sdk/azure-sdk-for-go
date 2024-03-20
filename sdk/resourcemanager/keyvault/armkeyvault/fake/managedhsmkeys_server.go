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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedHsmKeysServer is a fake server for instances of the armkeyvault.ManagedHsmKeysClient type.
type ManagedHsmKeysServer struct {
	// CreateIfNotExist is the fake for method ManagedHsmKeysClient.CreateIfNotExist
	// HTTP status codes to indicate success: http.StatusOK
	CreateIfNotExist func(ctx context.Context, resourceGroupName string, name string, keyName string, parameters armkeyvault.ManagedHsmKeyCreateParameters, options *armkeyvault.ManagedHsmKeysClientCreateIfNotExistOptions) (resp azfake.Responder[armkeyvault.ManagedHsmKeysClientCreateIfNotExistResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedHsmKeysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, name string, keyName string, options *armkeyvault.ManagedHsmKeysClientGetOptions) (resp azfake.Responder[armkeyvault.ManagedHsmKeysClientGetResponse], errResp azfake.ErrorResponder)

	// GetVersion is the fake for method ManagedHsmKeysClient.GetVersion
	// HTTP status codes to indicate success: http.StatusOK
	GetVersion func(ctx context.Context, resourceGroupName string, name string, keyName string, keyVersion string, options *armkeyvault.ManagedHsmKeysClientGetVersionOptions) (resp azfake.Responder[armkeyvault.ManagedHsmKeysClientGetVersionResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ManagedHsmKeysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, name string, options *armkeyvault.ManagedHsmKeysClientListOptions) (resp azfake.PagerResponder[armkeyvault.ManagedHsmKeysClientListResponse])

	// NewListVersionsPager is the fake for method ManagedHsmKeysClient.NewListVersionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListVersionsPager func(resourceGroupName string, name string, keyName string, options *armkeyvault.ManagedHsmKeysClientListVersionsOptions) (resp azfake.PagerResponder[armkeyvault.ManagedHsmKeysClientListVersionsResponse])
}

// NewManagedHsmKeysServerTransport creates a new instance of ManagedHsmKeysServerTransport with the provided implementation.
// The returned ManagedHsmKeysServerTransport instance is connected to an instance of armkeyvault.ManagedHsmKeysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedHsmKeysServerTransport(srv *ManagedHsmKeysServer) *ManagedHsmKeysServerTransport {
	return &ManagedHsmKeysServerTransport{
		srv:                  srv,
		newListPager:         newTracker[azfake.PagerResponder[armkeyvault.ManagedHsmKeysClientListResponse]](),
		newListVersionsPager: newTracker[azfake.PagerResponder[armkeyvault.ManagedHsmKeysClientListVersionsResponse]](),
	}
}

// ManagedHsmKeysServerTransport connects instances of armkeyvault.ManagedHsmKeysClient to instances of ManagedHsmKeysServer.
// Don't use this type directly, use NewManagedHsmKeysServerTransport instead.
type ManagedHsmKeysServerTransport struct {
	srv                  *ManagedHsmKeysServer
	newListPager         *tracker[azfake.PagerResponder[armkeyvault.ManagedHsmKeysClientListResponse]]
	newListVersionsPager *tracker[azfake.PagerResponder[armkeyvault.ManagedHsmKeysClientListVersionsResponse]]
}

// Do implements the policy.Transporter interface for ManagedHsmKeysServerTransport.
func (m *ManagedHsmKeysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedHsmKeysClient.CreateIfNotExist":
		resp, err = m.dispatchCreateIfNotExist(req)
	case "ManagedHsmKeysClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedHsmKeysClient.GetVersion":
		resp, err = m.dispatchGetVersion(req)
	case "ManagedHsmKeysClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	case "ManagedHsmKeysClient.NewListVersionsPager":
		resp, err = m.dispatchNewListVersionsPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedHsmKeysServerTransport) dispatchCreateIfNotExist(req *http.Request) (*http.Response, error) {
	if m.srv.CreateIfNotExist == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateIfNotExist not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armkeyvault.ManagedHsmKeyCreateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateIfNotExist(req.Context(), resourceGroupNameParam, nameParam, keyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedHsmKey, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedHsmKeysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, nameParam, keyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedHsmKey, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedHsmKeysServerTransport) dispatchGetVersion(req *http.Request) (*http.Response, error) {
	if m.srv.GetVersion == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetVersion not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<keyVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
	if err != nil {
		return nil, err
	}
	keyVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.GetVersion(req.Context(), resourceGroupNameParam, nameParam, keyNameParam, keyVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedHsmKey, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedHsmKeysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListPager(resourceGroupNameParam, nameParam, nil)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armkeyvault.ManagedHsmKeysClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

func (m *ManagedHsmKeysServerTransport) dispatchNewListVersionsPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListVersionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListVersionsPager not implemented")}
	}
	newListVersionsPager := m.newListVersionsPager.get(req)
	if newListVersionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/managedHSMs/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
		if err != nil {
			return nil, err
		}
		resp := m.srv.NewListVersionsPager(resourceGroupNameParam, nameParam, keyNameParam, nil)
		newListVersionsPager = &resp
		m.newListVersionsPager.add(req, newListVersionsPager)
		server.PagerResponderInjectNextLinks(newListVersionsPager, req, func(page *armkeyvault.ManagedHsmKeysClientListVersionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListVersionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListVersionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListVersionsPager) {
		m.newListVersionsPager.remove(req)
	}
	return resp, nil
}
