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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ExportsServer is a fake server for instances of the armcostmanagement.ExportsClient type.
type ExportsServer struct {
	// CreateOrUpdate is the fake for method ExportsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, scope string, exportName string, parameters armcostmanagement.Export, options *armcostmanagement.ExportsClientCreateOrUpdateOptions) (resp azfake.Responder[armcostmanagement.ExportsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ExportsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, scope string, exportName string, options *armcostmanagement.ExportsClientDeleteOptions) (resp azfake.Responder[armcostmanagement.ExportsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Execute is the fake for method ExportsClient.Execute
	// HTTP status codes to indicate success: http.StatusOK
	Execute func(ctx context.Context, scope string, exportName string, options *armcostmanagement.ExportsClientExecuteOptions) (resp azfake.Responder[armcostmanagement.ExportsClientExecuteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ExportsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, exportName string, options *armcostmanagement.ExportsClientGetOptions) (resp azfake.Responder[armcostmanagement.ExportsClientGetResponse], errResp azfake.ErrorResponder)

	// GetExecutionHistory is the fake for method ExportsClient.GetExecutionHistory
	// HTTP status codes to indicate success: http.StatusOK
	GetExecutionHistory func(ctx context.Context, scope string, exportName string, options *armcostmanagement.ExportsClientGetExecutionHistoryOptions) (resp azfake.Responder[armcostmanagement.ExportsClientGetExecutionHistoryResponse], errResp azfake.ErrorResponder)

	// List is the fake for method ExportsClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, scope string, options *armcostmanagement.ExportsClientListOptions) (resp azfake.Responder[armcostmanagement.ExportsClientListResponse], errResp azfake.ErrorResponder)
}

// NewExportsServerTransport creates a new instance of ExportsServerTransport with the provided implementation.
// The returned ExportsServerTransport instance is connected to an instance of armcostmanagement.ExportsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExportsServerTransport(srv *ExportsServer) *ExportsServerTransport {
	return &ExportsServerTransport{srv: srv}
}

// ExportsServerTransport connects instances of armcostmanagement.ExportsClient to instances of ExportsServer.
// Don't use this type directly, use NewExportsServerTransport instead.
type ExportsServerTransport struct {
	srv *ExportsServer
}

// Do implements the policy.Transporter interface for ExportsServerTransport.
func (e *ExportsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExportsClient.CreateOrUpdate":
		resp, err = e.dispatchCreateOrUpdate(req)
	case "ExportsClient.Delete":
		resp, err = e.dispatchDelete(req)
	case "ExportsClient.Execute":
		resp, err = e.dispatchExecute(req)
	case "ExportsClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExportsClient.GetExecutionHistory":
		resp, err = e.dispatchGetExecutionHistory(req)
	case "ExportsClient.List":
		resp, err = e.dispatchList(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExportsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/exports/(?P<exportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcostmanagement.Export](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	exportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.CreateOrUpdate(req.Context(), scopeParam, exportNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Export, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExportsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/exports/(?P<exportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	exportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), scopeParam, exportNameParam, nil)
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

func (e *ExportsServerTransport) dispatchExecute(req *http.Request) (*http.Response, error) {
	if e.srv.Execute == nil {
		return nil, &nonRetriableError{errors.New("fake for method Execute not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/exports/(?P<exportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/run`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	exportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Execute(req.Context(), scopeParam, exportNameParam, nil)
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

func (e *ExportsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/exports/(?P<exportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	exportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exportName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcostmanagement.ExportsClientGetOptions
	if expandParam != nil {
		options = &armcostmanagement.ExportsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := e.srv.Get(req.Context(), scopeParam, exportNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Export, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExportsServerTransport) dispatchGetExecutionHistory(req *http.Request) (*http.Response, error) {
	if e.srv.GetExecutionHistory == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetExecutionHistory not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/exports/(?P<exportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runHistory`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	exportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("exportName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetExecutionHistory(req.Context(), scopeParam, exportNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExportExecutionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExportsServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if e.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/exports`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armcostmanagement.ExportsClientListOptions
	if expandParam != nil {
		options = &armcostmanagement.ExportsClientListOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := e.srv.List(req.Context(), scopeParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExportListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
