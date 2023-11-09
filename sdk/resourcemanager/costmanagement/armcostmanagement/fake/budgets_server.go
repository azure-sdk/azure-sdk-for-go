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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement/v2"
	"net/http"
	"net/url"
	"regexp"
)

// BudgetsServer is a fake server for instances of the armcostmanagement.BudgetsClient type.
type BudgetsServer struct {
	// CreateOrUpdate is the fake for method BudgetsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, scope string, budgetName string, parameters armcostmanagement.Budget, options *armcostmanagement.BudgetsClientCreateOrUpdateOptions) (resp azfake.Responder[armcostmanagement.BudgetsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BudgetsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, scope string, budgetName string, options *armcostmanagement.BudgetsClientDeleteOptions) (resp azfake.Responder[armcostmanagement.BudgetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BudgetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, scope string, budgetName string, options *armcostmanagement.BudgetsClientGetOptions) (resp azfake.Responder[armcostmanagement.BudgetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BudgetsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(scope string, options *armcostmanagement.BudgetsClientListOptions) (resp azfake.PagerResponder[armcostmanagement.BudgetsClientListResponse])
}

// NewBudgetsServerTransport creates a new instance of BudgetsServerTransport with the provided implementation.
// The returned BudgetsServerTransport instance is connected to an instance of armcostmanagement.BudgetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBudgetsServerTransport(srv *BudgetsServer) *BudgetsServerTransport {
	return &BudgetsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcostmanagement.BudgetsClientListResponse]](),
	}
}

// BudgetsServerTransport connects instances of armcostmanagement.BudgetsClient to instances of BudgetsServer.
// Don't use this type directly, use NewBudgetsServerTransport instead.
type BudgetsServerTransport struct {
	srv          *BudgetsServer
	newListPager *tracker[azfake.PagerResponder[armcostmanagement.BudgetsClientListResponse]]
}

// Do implements the policy.Transporter interface for BudgetsServerTransport.
func (b *BudgetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BudgetsClient.CreateOrUpdate":
		resp, err = b.dispatchCreateOrUpdate(req)
	case "BudgetsClient.Delete":
		resp, err = b.dispatchDelete(req)
	case "BudgetsClient.Get":
		resp, err = b.dispatchGet(req)
	case "BudgetsClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BudgetsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/budgets/(?P<budgetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcostmanagement.Budget](req)
	if err != nil {
		return nil, err
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	budgetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("budgetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrUpdate(req.Context(), scopeParam, budgetNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Budget, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BudgetsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/budgets/(?P<budgetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	budgetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("budgetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), scopeParam, budgetNameParam, nil)
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

func (b *BudgetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/budgets/(?P<budgetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	scopeParam, err := url.PathUnescape(matches[regex.SubexpIndex("scope")])
	if err != nil {
		return nil, err
	}
	budgetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("budgetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), scopeParam, budgetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Budget, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BudgetsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/(?P<scope>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CostManagement/budgets`
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
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armcostmanagement.BudgetsClientListOptions
		if filterParam != nil {
			options = &armcostmanagement.BudgetsClientListOptions{
				Filter: filterParam,
			}
		}
		resp := b.srv.NewListPager(scopeParam, options)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcostmanagement.BudgetsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}
