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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ForecastServer is a fake server for instances of the armcostmanagement.ForecastClient type.
type ForecastServer struct {
	// ExternalCloudProviderUsage is the fake for method ForecastClient.ExternalCloudProviderUsage
	// HTTP status codes to indicate success: http.StatusOK
	ExternalCloudProviderUsage func(ctx context.Context, externalCloudProviderType armcostmanagement.ExternalCloudProviderType, externalCloudProviderID string, parameters armcostmanagement.ForecastDefinition, options *armcostmanagement.ForecastClientExternalCloudProviderUsageOptions) (resp azfake.Responder[armcostmanagement.ForecastClientExternalCloudProviderUsageResponse], errResp azfake.ErrorResponder)

	// Usage is the fake for method ForecastClient.Usage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Usage func(ctx context.Context, scope string, parameters armcostmanagement.ForecastDefinition, options *armcostmanagement.ForecastClientUsageOptions) (resp azfake.Responder[armcostmanagement.ForecastClientUsageResponse], errResp azfake.ErrorResponder)
}

// NewForecastServerTransport creates a new instance of ForecastServerTransport with the provided implementation.
// The returned ForecastServerTransport instance is connected to an instance of armcostmanagement.ForecastClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewForecastServerTransport(srv *ForecastServer) *ForecastServerTransport {
	return &ForecastServerTransport{srv: srv}
}

// ForecastServerTransport connects instances of armcostmanagement.ForecastClient to instances of ForecastServer.
// Don't use this type directly, use NewForecastServerTransport instead.
type ForecastServerTransport struct {
	srv *ForecastServer
}

// Do implements the policy.Transporter interface for ForecastServerTransport.
func (f *ForecastServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *ForecastServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if forecastServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = forecastServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ForecastClient.ExternalCloudProviderUsage":
				res.resp, res.err = f.dispatchExternalCloudProviderUsage(req)
			case "ForecastClient.Usage":
				res.resp, res.err = f.dispatchUsage(req)
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

func (f *ForecastServerTransport) dispatchExternalCloudProviderUsage(req *http.Request) (*http.Response, error) {
	if f.srv.ExternalCloudProviderUsage == nil {
		return nil, &nonRetriableError{errors.New("fake for method ExternalCloudProviderUsage not implemented")}
	}
	const regexStr = `/providers/Microsoft\.CostManagement/(?P<externalCloudProviderType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<externalCloudProviderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/forecast`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armcostmanagement.ForecastDefinition](req)
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	externalCloudProviderTypeParam, err := parseWithCast(matches[regex.SubexpIndex("externalCloudProviderType")], func(v string) (armcostmanagement.ExternalCloudProviderType, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armcostmanagement.ExternalCloudProviderType(p), nil
	})
	if err != nil {
		return nil, err
	}
	externalCloudProviderIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("externalCloudProviderId")])
	if err != nil {
		return nil, err
	}
	var options *armcostmanagement.ForecastClientExternalCloudProviderUsageOptions
	if filterParam != nil {
		options = &armcostmanagement.ForecastClientExternalCloudProviderUsageOptions{
			Filter: filterParam,
		}
	}
	respr, errRespr := f.srv.ExternalCloudProviderUsage(req.Context(), externalCloudProviderTypeParam, externalCloudProviderIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ForecastResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *ForecastServerTransport) dispatchUsage(req *http.Request) (*http.Response, error) {
	if f.srv.Usage == nil {
		return nil, &nonRetriableError{errors.New("fake for method Usage not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/forecast`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armcostmanagement.ForecastDefinition](req)
	if err != nil {
		return nil, err
	}
	filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
	if err != nil {
		return nil, err
	}
	filterParam := getOptional(filterUnescaped)
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	var options *armcostmanagement.ForecastClientUsageOptions
	if filterParam != nil {
		options = &armcostmanagement.ForecastClientUsageOptions{
			Filter: filterParam,
		}
	}
	respr, errRespr := f.srv.Usage(req.Context(), scopeParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ForecastResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ForecastServerTransport
var forecastServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
