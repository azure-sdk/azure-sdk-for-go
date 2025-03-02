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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/security/armsecurity"
	"net/http"
	"net/url"
	"regexp"
)

// AdvancedThreatProtectionServer is a fake server for instances of the armsecurity.AdvancedThreatProtectionClient type.
type AdvancedThreatProtectionServer struct {
	// Create is the fake for method AdvancedThreatProtectionClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceID string, advancedThreatProtectionSetting armsecurity.AdvancedThreatProtectionSetting, options *armsecurity.AdvancedThreatProtectionClientCreateOptions) (resp azfake.Responder[armsecurity.AdvancedThreatProtectionClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AdvancedThreatProtectionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceID string, options *armsecurity.AdvancedThreatProtectionClientGetOptions) (resp azfake.Responder[armsecurity.AdvancedThreatProtectionClientGetResponse], errResp azfake.ErrorResponder)
}

// NewAdvancedThreatProtectionServerTransport creates a new instance of AdvancedThreatProtectionServerTransport with the provided implementation.
// The returned AdvancedThreatProtectionServerTransport instance is connected to an instance of armsecurity.AdvancedThreatProtectionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAdvancedThreatProtectionServerTransport(srv *AdvancedThreatProtectionServer) *AdvancedThreatProtectionServerTransport {
	return &AdvancedThreatProtectionServerTransport{srv: srv}
}

// AdvancedThreatProtectionServerTransport connects instances of armsecurity.AdvancedThreatProtectionClient to instances of AdvancedThreatProtectionServer.
// Don't use this type directly, use NewAdvancedThreatProtectionServerTransport instead.
type AdvancedThreatProtectionServerTransport struct {
	srv *AdvancedThreatProtectionServer
}

// Do implements the policy.Transporter interface for AdvancedThreatProtectionServerTransport.
func (a *AdvancedThreatProtectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AdvancedThreatProtectionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if advancedThreatProtectionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = advancedThreatProtectionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AdvancedThreatProtectionClient.Create":
				res.resp, res.err = a.dispatchCreate(req)
			case "AdvancedThreatProtectionClient.Get":
				res.resp, res.err = a.dispatchGet(req)
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

func (a *AdvancedThreatProtectionServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if a.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/advancedThreatProtectionSettings/(?P<settingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurity.AdvancedThreatProtectionSetting](req)
	if err != nil {
		return nil, err
	}
	resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Create(req.Context(), resourceIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AdvancedThreatProtectionSetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AdvancedThreatProtectionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Security/advancedThreatProtectionSettings/(?P<settingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AdvancedThreatProtectionSetting, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AdvancedThreatProtectionServerTransport
var advancedThreatProtectionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
