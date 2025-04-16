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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights/v2"
	"net/http"
	"net/url"
	"regexp"
)

// IntelligencePacksServer is a fake server for instances of the armoperationalinsights.IntelligencePacksClient type.
type IntelligencePacksServer struct {
	// Disable is the fake for method IntelligencePacksClient.Disable
	// HTTP status codes to indicate success: http.StatusOK
	Disable func(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string, options *armoperationalinsights.IntelligencePacksClientDisableOptions) (resp azfake.Responder[armoperationalinsights.IntelligencePacksClientDisableResponse], errResp azfake.ErrorResponder)

	// Enable is the fake for method IntelligencePacksClient.Enable
	// HTTP status codes to indicate success: http.StatusOK
	Enable func(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string, options *armoperationalinsights.IntelligencePacksClientEnableOptions) (resp azfake.Responder[armoperationalinsights.IntelligencePacksClientEnableResponse], errResp azfake.ErrorResponder)

	// List is the fake for method IntelligencePacksClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, workspaceName string, options *armoperationalinsights.IntelligencePacksClientListOptions) (resp azfake.Responder[armoperationalinsights.IntelligencePacksClientListResponse], errResp azfake.ErrorResponder)
}

// NewIntelligencePacksServerTransport creates a new instance of IntelligencePacksServerTransport with the provided implementation.
// The returned IntelligencePacksServerTransport instance is connected to an instance of armoperationalinsights.IntelligencePacksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntelligencePacksServerTransport(srv *IntelligencePacksServer) *IntelligencePacksServerTransport {
	return &IntelligencePacksServerTransport{srv: srv}
}

// IntelligencePacksServerTransport connects instances of armoperationalinsights.IntelligencePacksClient to instances of IntelligencePacksServer.
// Don't use this type directly, use NewIntelligencePacksServerTransport instead.
type IntelligencePacksServerTransport struct {
	srv *IntelligencePacksServer
}

// Do implements the policy.Transporter interface for IntelligencePacksServerTransport.
func (i *IntelligencePacksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IntelligencePacksServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if intelligencePacksServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = intelligencePacksServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IntelligencePacksClient.Disable":
				res.resp, res.err = i.dispatchDisable(req)
			case "IntelligencePacksClient.Enable":
				res.resp, res.err = i.dispatchEnable(req)
			case "IntelligencePacksClient.List":
				res.resp, res.err = i.dispatchList(req)
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

func (i *IntelligencePacksServerTransport) dispatchDisable(req *http.Request) (*http.Response, error) {
	if i.srv.Disable == nil {
		return nil, &nonRetriableError{errors.New("fake for method Disable not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/intelligencePacks/(?P<intelligencePackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Disable`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	intelligencePackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("intelligencePackName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Disable(req.Context(), resourceGroupNameParam, workspaceNameParam, intelligencePackNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntelligencePacksServerTransport) dispatchEnable(req *http.Request) (*http.Response, error) {
	if i.srv.Enable == nil {
		return nil, &nonRetriableError{errors.New("fake for method Enable not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/intelligencePacks/(?P<intelligencePackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/Enable`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	intelligencePackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("intelligencePackName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Enable(req.Context(), resourceGroupNameParam, workspaceNameParam, intelligencePackNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntelligencePacksServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if i.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/intelligencePacks`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.List(req.Context(), resourceGroupNameParam, workspaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntelligencePackArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to IntelligencePacksServerTransport
var intelligencePacksServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
