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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
	"net/http"
	"net/url"
	"regexp"
)

// ProtectedItemsServer is a fake server for instances of the armrecoveryservicesbackup.ProtectedItemsClient type.
type ProtectedItemsServer struct {
	// CreateOrUpdate is the fake for method ProtectedItemsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	CreateOrUpdate func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, parameters armrecoveryservicesbackup.ProtectedItemResource, options *armrecoveryservicesbackup.ProtectedItemsClientCreateOrUpdateOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectedItemsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ProtectedItemsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	Delete func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, options *armrecoveryservicesbackup.ProtectedItemsClientDeleteOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectedItemsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ProtectedItemsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, options *armrecoveryservicesbackup.ProtectedItemsClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.ProtectedItemsClientGetResponse], errResp azfake.ErrorResponder)
}

// NewProtectedItemsServerTransport creates a new instance of ProtectedItemsServerTransport with the provided implementation.
// The returned ProtectedItemsServerTransport instance is connected to an instance of armrecoveryservicesbackup.ProtectedItemsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProtectedItemsServerTransport(srv *ProtectedItemsServer) *ProtectedItemsServerTransport {
	return &ProtectedItemsServerTransport{srv: srv}
}

// ProtectedItemsServerTransport connects instances of armrecoveryservicesbackup.ProtectedItemsClient to instances of ProtectedItemsServer.
// Don't use this type directly, use NewProtectedItemsServerTransport instead.
type ProtectedItemsServerTransport struct {
	srv *ProtectedItemsServer
}

// Do implements the policy.Transporter interface for ProtectedItemsServerTransport.
func (p *ProtectedItemsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProtectedItemsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if protectedItemsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = protectedItemsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProtectedItemsClient.CreateOrUpdate":
				res.resp, res.err = p.dispatchCreateOrUpdate(req)
			case "ProtectedItemsClient.Delete":
				res.resp, res.err = p.dispatchDelete(req)
			case "ProtectedItemsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
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

func (p *ProtectedItemsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.ProtectedItemResource](req)
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
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
	if err != nil {
		return nil, err
	}
	xMSAuthorizationAuxiliaryParam := getOptional(getHeaderValue(req.Header, "x-ms-authorization-auxiliary"))
	var options *armrecoveryservicesbackup.ProtectedItemsClientCreateOrUpdateOptions
	if xMSAuthorizationAuxiliaryParam != nil {
		options = &armrecoveryservicesbackup.ProtectedItemsClientCreateOrUpdateOptions{
			XMSAuthorizationAuxiliary: xMSAuthorizationAuxiliaryParam,
		}
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, protectedItemNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectedItemResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectedItemsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
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
	fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
	if err != nil {
		return nil, err
	}
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, protectedItemNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProtectedItemsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectionContainers/(?P<containerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/protectedItems/(?P<protectedItemName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
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
	containerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerName")])
	if err != nil {
		return nil, err
	}
	protectedItemNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectedItemName")])
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	var options *armrecoveryservicesbackup.ProtectedItemsClientGetOptions
	if filterParam != nil {
		options = &armrecoveryservicesbackup.ProtectedItemsClientGetOptions{
			Filter: filterParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, fabricNameParam, containerNameParam, protectedItemNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProtectedItemResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProtectedItemsServerTransport
var protectedItemsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
