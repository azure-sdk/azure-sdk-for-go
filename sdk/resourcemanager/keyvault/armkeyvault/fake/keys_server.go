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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"net/http"
	"net/url"
	"regexp"
)

// KeysServer is a fake server for instances of the armkeyvault.KeysClient type.
type KeysServer struct {
	// CreateIfNotExist is the fake for method KeysClient.CreateIfNotExist
	// HTTP status codes to indicate success: http.StatusOK
	CreateIfNotExist func(ctx context.Context, resourceGroupName string, vaultName string, keyName string, parameters armkeyvault.KeyCreateParameters, options *armkeyvault.KeysClientCreateIfNotExistOptions) (resp azfake.Responder[armkeyvault.KeysClientCreateIfNotExistResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method KeysClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, keyName string, options *armkeyvault.KeysClientGetOptions) (resp azfake.Responder[armkeyvault.KeysClientGetResponse], errResp azfake.ErrorResponder)

	// GetVersion is the fake for method KeysClient.GetVersion
	// HTTP status codes to indicate success: http.StatusOK
	GetVersion func(ctx context.Context, resourceGroupName string, vaultName string, keyName string, keyVersion string, options *armkeyvault.KeysClientGetVersionOptions) (resp azfake.Responder[armkeyvault.KeysClientGetVersionResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method KeysClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vaultName string, options *armkeyvault.KeysClientListOptions) (resp azfake.PagerResponder[armkeyvault.KeysClientListResponse])

	// NewListVersionsPager is the fake for method KeysClient.NewListVersionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListVersionsPager func(resourceGroupName string, vaultName string, keyName string, options *armkeyvault.KeysClientListVersionsOptions) (resp azfake.PagerResponder[armkeyvault.KeysClientListVersionsResponse])
}

// NewKeysServerTransport creates a new instance of KeysServerTransport with the provided implementation.
// The returned KeysServerTransport instance is connected to an instance of armkeyvault.KeysClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewKeysServerTransport(srv *KeysServer) *KeysServerTransport {
	return &KeysServerTransport{
		srv:                  srv,
		newListPager:         newTracker[azfake.PagerResponder[armkeyvault.KeysClientListResponse]](),
		newListVersionsPager: newTracker[azfake.PagerResponder[armkeyvault.KeysClientListVersionsResponse]](),
	}
}

// KeysServerTransport connects instances of armkeyvault.KeysClient to instances of KeysServer.
// Don't use this type directly, use NewKeysServerTransport instead.
type KeysServerTransport struct {
	srv                  *KeysServer
	newListPager         *tracker[azfake.PagerResponder[armkeyvault.KeysClientListResponse]]
	newListVersionsPager *tracker[azfake.PagerResponder[armkeyvault.KeysClientListVersionsResponse]]
}

// Do implements the policy.Transporter interface for KeysServerTransport.
func (k *KeysServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return k.dispatchToMethodFake(req, method)
}

func (k *KeysServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if keysServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = keysServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "KeysClient.CreateIfNotExist":
				res.resp, res.err = k.dispatchCreateIfNotExist(req)
			case "KeysClient.Get":
				res.resp, res.err = k.dispatchGet(req)
			case "KeysClient.GetVersion":
				res.resp, res.err = k.dispatchGetVersion(req)
			case "KeysClient.NewListPager":
				res.resp, res.err = k.dispatchNewListPager(req)
			case "KeysClient.NewListVersionsPager":
				res.resp, res.err = k.dispatchNewListVersionsPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (k *KeysServerTransport) dispatchCreateIfNotExist(req *http.Request) (*http.Response, error) {
	if k.srv.CreateIfNotExist == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateIfNotExist not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armkeyvault.KeyCreateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := k.srv.CreateIfNotExist(req.Context(), resourceGroupNameParam, vaultNameParam, keyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Key, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (k *KeysServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if k.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := k.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, keyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Key, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (k *KeysServerTransport) dispatchGetVersion(req *http.Request) (*http.Response, error) {
	if k.srv.GetVersion == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetVersion not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<keyVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
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
	respr, errRespr := k.srv.GetVersion(req.Context(), resourceGroupNameParam, vaultNameParam, keyNameParam, keyVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Key, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (k *KeysServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if k.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := k.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resp := k.srv.NewListPager(resourceGroupNameParam, vaultNameParam, nil)
		newListPager = &resp
		k.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armkeyvault.KeysClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		k.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		k.newListPager.remove(req)
	}
	return resp, nil
}

func (k *KeysServerTransport) dispatchNewListVersionsPager(req *http.Request) (*http.Response, error) {
	if k.srv.NewListVersionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListVersionsPager not implemented")}
	}
	newListVersionsPager := k.newListVersionsPager.get(req)
	if newListVersionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KeyVault/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/keys/(?P<keyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		keyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("keyName")])
		if err != nil {
			return nil, err
		}
		resp := k.srv.NewListVersionsPager(resourceGroupNameParam, vaultNameParam, keyNameParam, nil)
		newListVersionsPager = &resp
		k.newListVersionsPager.add(req, newListVersionsPager)
		server.PagerResponderInjectNextLinks(newListVersionsPager, req, func(page *armkeyvault.KeysClientListVersionsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListVersionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		k.newListVersionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListVersionsPager) {
		k.newListVersionsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to KeysServerTransport
var keysServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
