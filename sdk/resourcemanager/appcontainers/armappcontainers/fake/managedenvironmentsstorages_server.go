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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ManagedEnvironmentsStoragesServer is a fake server for instances of the armappcontainers.ManagedEnvironmentsStoragesClient type.
type ManagedEnvironmentsStoragesServer struct {
	// CreateOrUpdate is the fake for method ManagedEnvironmentsStoragesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, environmentName string, storageName string, storageEnvelope armappcontainers.ManagedEnvironmentStorage, options *armappcontainers.ManagedEnvironmentsStoragesClientCreateOrUpdateOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsStoragesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ManagedEnvironmentsStoragesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, environmentName string, storageName string, options *armappcontainers.ManagedEnvironmentsStoragesClientDeleteOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsStoragesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ManagedEnvironmentsStoragesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, environmentName string, storageName string, options *armappcontainers.ManagedEnvironmentsStoragesClientGetOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsStoragesClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ManagedEnvironmentsStoragesClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, environmentName string, options *armappcontainers.ManagedEnvironmentsStoragesClientListOptions) (resp azfake.Responder[armappcontainers.ManagedEnvironmentsStoragesClientListResponse], errResp azfake.ErrorResponder)
}

// NewManagedEnvironmentsStoragesServerTransport creates a new instance of ManagedEnvironmentsStoragesServerTransport with the provided implementation.
// The returned ManagedEnvironmentsStoragesServerTransport instance is connected to an instance of armappcontainers.ManagedEnvironmentsStoragesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewManagedEnvironmentsStoragesServerTransport(srv *ManagedEnvironmentsStoragesServer) *ManagedEnvironmentsStoragesServerTransport {
	return &ManagedEnvironmentsStoragesServerTransport{srv: srv}
}

// ManagedEnvironmentsStoragesServerTransport connects instances of armappcontainers.ManagedEnvironmentsStoragesClient to instances of ManagedEnvironmentsStoragesServer.
// Don't use this type directly, use NewManagedEnvironmentsStoragesServerTransport instead.
type ManagedEnvironmentsStoragesServerTransport struct {
	srv *ManagedEnvironmentsStoragesServer
}

// Do implements the policy.Transporter interface for ManagedEnvironmentsStoragesServerTransport.
func (m *ManagedEnvironmentsStoragesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ManagedEnvironmentsStoragesClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ManagedEnvironmentsStoragesClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ManagedEnvironmentsStoragesClient.Get":
		resp, err = m.dispatchGet(req)
	case "ManagedEnvironmentsStoragesClient.List":
		resp, err = m.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ManagedEnvironmentsStoragesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages/(?P<storageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcontainers.ManagedEnvironmentStorage](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	environmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("environmentName")])
	if err != nil {
		return nil, err
	}
	storageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, environmentNameParam, storageNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedEnvironmentStorage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedEnvironmentsStoragesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages/(?P<storageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	storageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, environmentNameParam, storageNameParam, nil)
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

func (m *ManagedEnvironmentsStoragesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages/(?P<storageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	storageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("storageName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, environmentNameParam, storageNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedEnvironmentStorage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ManagedEnvironmentsStoragesServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if m.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/managedEnvironments/(?P<environmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/storages`
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
	respr, errRespr := m.srv.List(req.Context(), resourceGroupNameParam, environmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ManagedEnvironmentStoragesCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
