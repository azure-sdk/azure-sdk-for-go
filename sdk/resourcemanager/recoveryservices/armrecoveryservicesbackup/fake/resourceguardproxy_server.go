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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ResourceGuardProxyServer is a fake server for instances of the armrecoveryservicesbackup.ResourceGuardProxyClient type.
type ResourceGuardProxyServer struct {
	// Delete is the fake for method ResourceGuardProxyClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *armrecoveryservicesbackup.ResourceGuardProxyClientDeleteOptions) (resp azfake.Responder[armrecoveryservicesbackup.ResourceGuardProxyClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ResourceGuardProxyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *armrecoveryservicesbackup.ResourceGuardProxyClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ResourceGuardProxyClientGetResponse], errResp azfake.ErrorResponder)

	// Put is the fake for method ResourceGuardProxyClient.Put
	// HTTP status codes to indicate success: http.StatusOK
	Put func(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters armrecoveryservicesbackup.ResourceGuardProxyBaseResource, options *armrecoveryservicesbackup.ResourceGuardProxyClientPutOptions) (resp azfake.Responder[armrecoveryservicesbackup.ResourceGuardProxyClientPutResponse], errResp azfake.ErrorResponder)

	// UnlockDelete is the fake for method ResourceGuardProxyClient.UnlockDelete
	// HTTP status codes to indicate success: http.StatusOK
	UnlockDelete func(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters armrecoveryservicesbackup.UnlockDeleteRequest, options *armrecoveryservicesbackup.ResourceGuardProxyClientUnlockDeleteOptions) (resp azfake.Responder[armrecoveryservicesbackup.ResourceGuardProxyClientUnlockDeleteResponse], errResp azfake.ErrorResponder)
}

// NewResourceGuardProxyServerTransport creates a new instance of ResourceGuardProxyServerTransport with the provided implementation.
// The returned ResourceGuardProxyServerTransport instance is connected to an instance of armrecoveryservicesbackup.ResourceGuardProxyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceGuardProxyServerTransport(srv *ResourceGuardProxyServer) *ResourceGuardProxyServerTransport {
	return &ResourceGuardProxyServerTransport{srv: srv}
}

// ResourceGuardProxyServerTransport connects instances of armrecoveryservicesbackup.ResourceGuardProxyClient to instances of ResourceGuardProxyServer.
// Don't use this type directly, use NewResourceGuardProxyServerTransport instead.
type ResourceGuardProxyServerTransport struct {
	srv *ResourceGuardProxyServer
}

// Do implements the policy.Transporter interface for ResourceGuardProxyServerTransport.
func (r *ResourceGuardProxyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ResourceGuardProxyServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if resourceGuardProxyServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = resourceGuardProxyServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ResourceGuardProxyClient.Delete":
				res.resp, res.err = r.dispatchDelete(req)
			case "ResourceGuardProxyClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ResourceGuardProxyClient.Put":
				res.resp, res.err = r.dispatchPut(req)
			case "ResourceGuardProxyClient.UnlockDelete":
				res.resp, res.err = r.dispatchUnlockDelete(req)
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

func (r *ResourceGuardProxyServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupResourceGuardProxies/(?P<resourceGuardProxyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceGuardProxyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGuardProxyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Delete(req.Context(), vaultNameParam, resourceGroupNameParam, resourceGuardProxyNameParam, nil)
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

func (r *ResourceGuardProxyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupResourceGuardProxies/(?P<resourceGuardProxyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceGuardProxyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGuardProxyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, resourceGuardProxyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceGuardProxyBaseResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceGuardProxyServerTransport) dispatchPut(req *http.Request) (*http.Response, error) {
	if r.srv.Put == nil {
		return nil, &nonRetriableError{errors.New("fake for method Put not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupResourceGuardProxies/(?P<resourceGuardProxyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.ResourceGuardProxyBaseResource](req)
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceGuardProxyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGuardProxyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Put(req.Context(), vaultNameParam, resourceGroupNameParam, resourceGuardProxyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceGuardProxyBaseResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceGuardProxyServerTransport) dispatchUnlockDelete(req *http.Request) (*http.Response, error) {
	if r.srv.UnlockDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method UnlockDelete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupResourceGuardProxies/(?P<resourceGuardProxyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unlockDelete`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.UnlockDeleteRequest](req)
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceGuardProxyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGuardProxyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.UnlockDelete(req.Context(), vaultNameParam, resourceGroupNameParam, resourceGuardProxyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).UnlockDeleteResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ResourceGuardProxyServerTransport
var resourceGuardProxyServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
