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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PolicyDefinitionsServer is a fake server for instances of the armsubscriptions.PolicyDefinitionsClient type.
type PolicyDefinitionsServer struct {
	// CreateOrUpdate is the fake for method PolicyDefinitionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	CreateOrUpdate func(ctx context.Context, policyDefinitionName string, parameters armsubscriptions.PolicyDefinition, options *armsubscriptions.PolicyDefinitionsClientCreateOrUpdateOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtManagementGroup is the fake for method PolicyDefinitionsClient.CreateOrUpdateAtManagementGroup
	// HTTP status codes to indicate success: http.StatusCreated
	CreateOrUpdateAtManagementGroup func(ctx context.Context, managementGroupID string, policyDefinitionName string, parameters armsubscriptions.PolicyDefinition, options *armsubscriptions.PolicyDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientCreateOrUpdateAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PolicyDefinitionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, policyDefinitionName string, options *armsubscriptions.PolicyDefinitionsClientDeleteOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtManagementGroup is the fake for method PolicyDefinitionsClient.DeleteAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAtManagementGroup func(ctx context.Context, managementGroupID string, policyDefinitionName string, options *armsubscriptions.PolicyDefinitionsClientDeleteAtManagementGroupOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientDeleteAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PolicyDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, policyDefinitionName string, options *armsubscriptions.PolicyDefinitionsClientGetOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtManagementGroup is the fake for method PolicyDefinitionsClient.GetAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	GetAtManagementGroup func(ctx context.Context, managementGroupID string, policyDefinitionName string, options *armsubscriptions.PolicyDefinitionsClientGetAtManagementGroupOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientGetAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// GetBuiltIn is the fake for method PolicyDefinitionsClient.GetBuiltIn
	// HTTP status codes to indicate success: http.StatusOK
	GetBuiltIn func(ctx context.Context, policyDefinitionName string, options *armsubscriptions.PolicyDefinitionsClientGetBuiltInOptions) (resp azfake.Responder[armsubscriptions.PolicyDefinitionsClientGetBuiltInResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PolicyDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armsubscriptions.PolicyDefinitionsClientListOptions) (resp azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListResponse])

	// NewListBuiltInPager is the fake for method PolicyDefinitionsClient.NewListBuiltInPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBuiltInPager func(options *armsubscriptions.PolicyDefinitionsClientListBuiltInOptions) (resp azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListBuiltInResponse])

	// NewListByManagementGroupPager is the fake for method PolicyDefinitionsClient.NewListByManagementGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagementGroupPager func(managementGroupID string, options *armsubscriptions.PolicyDefinitionsClientListByManagementGroupOptions) (resp azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListByManagementGroupResponse])
}

// NewPolicyDefinitionsServerTransport creates a new instance of PolicyDefinitionsServerTransport with the provided implementation.
// The returned PolicyDefinitionsServerTransport instance is connected to an instance of armsubscriptions.PolicyDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPolicyDefinitionsServerTransport(srv *PolicyDefinitionsServer) *PolicyDefinitionsServerTransport {
	return &PolicyDefinitionsServerTransport{
		srv:                           srv,
		newListPager:                  newTracker[azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListResponse]](),
		newListBuiltInPager:           newTracker[azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListBuiltInResponse]](),
		newListByManagementGroupPager: newTracker[azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListByManagementGroupResponse]](),
	}
}

// PolicyDefinitionsServerTransport connects instances of armsubscriptions.PolicyDefinitionsClient to instances of PolicyDefinitionsServer.
// Don't use this type directly, use NewPolicyDefinitionsServerTransport instead.
type PolicyDefinitionsServerTransport struct {
	srv                           *PolicyDefinitionsServer
	newListPager                  *tracker[azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListResponse]]
	newListBuiltInPager           *tracker[azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListBuiltInResponse]]
	newListByManagementGroupPager *tracker[azfake.PagerResponder[armsubscriptions.PolicyDefinitionsClientListByManagementGroupResponse]]
}

// Do implements the policy.Transporter interface for PolicyDefinitionsServerTransport.
func (p *PolicyDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PolicyDefinitionsClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PolicyDefinitionsClient.CreateOrUpdateAtManagementGroup":
		resp, err = p.dispatchCreateOrUpdateAtManagementGroup(req)
	case "PolicyDefinitionsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PolicyDefinitionsClient.DeleteAtManagementGroup":
		resp, err = p.dispatchDeleteAtManagementGroup(req)
	case "PolicyDefinitionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PolicyDefinitionsClient.GetAtManagementGroup":
		resp, err = p.dispatchGetAtManagementGroup(req)
	case "PolicyDefinitionsClient.GetBuiltIn":
		resp, err = p.dispatchGetBuiltIn(req)
	case "PolicyDefinitionsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PolicyDefinitionsClient.NewListBuiltInPager":
		resp, err = p.dispatchNewListBuiltInPager(req)
	case "PolicyDefinitionsClient.NewListByManagementGroupPager":
		resp, err = p.dispatchNewListByManagementGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsubscriptions.PolicyDefinition](req)
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), policyDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchCreateOrUpdateAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdateAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsubscriptions.PolicyDefinition](req)
	if err != nil {
		return nil, err
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdateAtManagementGroup(req.Context(), managementGroupIDParam, policyDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), policyDefinitionNameParam, nil)
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

func (p *PolicyDefinitionsServerTransport) dispatchDeleteAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.DeleteAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DeleteAtManagementGroup(req.Context(), managementGroupIDParam, policyDefinitionNameParam, nil)
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

func (p *PolicyDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), policyDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchGetAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.GetAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetAtManagementGroup(req.Context(), managementGroupIDParam, policyDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchGetBuiltIn(req *http.Request) (*http.Response, error) {
	if p.srv.GetBuiltIn == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBuiltIn not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetBuiltIn(req.Context(), policyDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsubscriptions.PolicyDefinitionsClientListOptions
		if filterParam != nil || topParam != nil {
			options = &armsubscriptions.PolicyDefinitionsClientListOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListPager(options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsubscriptions.PolicyDefinitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchNewListBuiltInPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBuiltInPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBuiltInPager not implemented")}
	}
	newListBuiltInPager := p.newListBuiltInPager.get(req)
	if newListBuiltInPager == nil {
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsubscriptions.PolicyDefinitionsClientListBuiltInOptions
		if filterParam != nil || topParam != nil {
			options = &armsubscriptions.PolicyDefinitionsClientListBuiltInOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListBuiltInPager(options)
		newListBuiltInPager = &resp
		p.newListBuiltInPager.add(req, newListBuiltInPager)
		server.PagerResponderInjectNextLinks(newListBuiltInPager, req, func(page *armsubscriptions.PolicyDefinitionsClientListBuiltInResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBuiltInPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListBuiltInPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBuiltInPager) {
		p.newListBuiltInPager.remove(req)
	}
	return resp, nil
}

func (p *PolicyDefinitionsServerTransport) dispatchNewListByManagementGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByManagementGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagementGroupPager not implemented")}
	}
	newListByManagementGroupPager := p.newListByManagementGroupPager.get(req)
	if newListByManagementGroupPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armsubscriptions.PolicyDefinitionsClientListByManagementGroupOptions
		if filterParam != nil || topParam != nil {
			options = &armsubscriptions.PolicyDefinitionsClientListByManagementGroupOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListByManagementGroupPager(managementGroupIDParam, options)
		newListByManagementGroupPager = &resp
		p.newListByManagementGroupPager.add(req, newListByManagementGroupPager)
		server.PagerResponderInjectNextLinks(newListByManagementGroupPager, req, func(page *armsubscriptions.PolicyDefinitionsClientListByManagementGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByManagementGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByManagementGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByManagementGroupPager) {
		p.newListByManagementGroupPager.remove(req)
	}
	return resp, nil
}
