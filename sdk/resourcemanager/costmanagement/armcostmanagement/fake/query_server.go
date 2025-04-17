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

// QueryServer is a fake server for instances of the armcostmanagement.QueryClient type.
type QueryServer struct {
	// Usage is the fake for method QueryClient.Usage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Usage func(ctx context.Context, scope string, parameters armcostmanagement.QueryDefinition, options *armcostmanagement.QueryClientUsageOptions) (resp azfake.Responder[armcostmanagement.QueryClientUsageResponse], errResp azfake.ErrorResponder)

	// UsageByExternalCloudProviderType is the fake for method QueryClient.UsageByExternalCloudProviderType
	// HTTP status codes to indicate success: http.StatusOK
	UsageByExternalCloudProviderType func(ctx context.Context, externalCloudProviderType armcostmanagement.ExternalCloudProviderType, externalCloudProviderID string, parameters armcostmanagement.QueryDefinition, options *armcostmanagement.QueryClientUsageByExternalCloudProviderTypeOptions) (resp azfake.Responder[armcostmanagement.QueryClientUsageByExternalCloudProviderTypeResponse], errResp azfake.ErrorResponder)
}

// NewQueryServerTransport creates a new instance of QueryServerTransport with the provided implementation.
// The returned QueryServerTransport instance is connected to an instance of armcostmanagement.QueryClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewQueryServerTransport(srv *QueryServer) *QueryServerTransport {
	return &QueryServerTransport{srv: srv}
}

// QueryServerTransport connects instances of armcostmanagement.QueryClient to instances of QueryServer.
// Don't use this type directly, use NewQueryServerTransport instead.
type QueryServerTransport struct {
	srv *QueryServer
}

// Do implements the policy.Transporter interface for QueryServerTransport.
func (q *QueryServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return q.dispatchToMethodFake(req, method)
}

func (q *QueryServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if queryServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = queryServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "QueryClient.Usage":
				res.resp, res.err = q.dispatchUsage(req)
			case "QueryClient.UsageByExternalCloudProviderType":
				res.resp, res.err = q.dispatchUsageByExternalCloudProviderType(req)
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

func (q *QueryServerTransport) dispatchUsage(req *http.Request) (*http.Response, error) {
	if q.srv.Usage == nil {
		return nil, &nonRetriableError{errors.New("fake for method Usage not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/query`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcostmanagement.QueryDefinition](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := q.srv.Usage(req.Context(), scopeParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueryResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (q *QueryServerTransport) dispatchUsageByExternalCloudProviderType(req *http.Request) (*http.Response, error) {
	if q.srv.UsageByExternalCloudProviderType == nil {
		return nil, &nonRetriableError{errors.New("fake for method UsageByExternalCloudProviderType not implemented")}
	}
	const regexStr = `/providers/Microsoft\.CostManagement/(?P<externalCloudProviderType>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/(?P<externalCloudProviderId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/query`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcostmanagement.QueryDefinition](req)
	if err != nil {
		return nil, err
	}
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
	respr, errRespr := q.srv.UsageByExternalCloudProviderType(req.Context(), externalCloudProviderTypeParam, externalCloudProviderIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).QueryResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to QueryServerTransport
var queryServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
