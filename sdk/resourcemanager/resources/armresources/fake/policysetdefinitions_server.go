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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PolicySetDefinitionsServer is a fake server for instances of the armresources.PolicySetDefinitionsClient type.
type PolicySetDefinitionsServer struct {
	// CreateOrUpdate is the fake for method PolicySetDefinitionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, policySetDefinitionName string, parameters armresources.PolicySetDefinition, options *armresources.PolicySetDefinitionsClientCreateOrUpdateOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtManagementGroup is the fake for method PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdateAtManagementGroup func(ctx context.Context, managementGroupID string, policySetDefinitionName string, parameters armresources.PolicySetDefinition, options *armresources.PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PolicySetDefinitionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, policySetDefinitionName string, options *armresources.PolicySetDefinitionsClientDeleteOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtManagementGroup is the fake for method PolicySetDefinitionsClient.DeleteAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAtManagementGroup func(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *armresources.PolicySetDefinitionsClientDeleteAtManagementGroupOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientDeleteAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PolicySetDefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, policySetDefinitionName string, options *armresources.PolicySetDefinitionsClientGetOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtManagementGroup is the fake for method PolicySetDefinitionsClient.GetAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	GetAtManagementGroup func(ctx context.Context, managementGroupID string, policySetDefinitionName string, options *armresources.PolicySetDefinitionsClientGetAtManagementGroupOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientGetAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// GetBuiltIn is the fake for method PolicySetDefinitionsClient.GetBuiltIn
	// HTTP status codes to indicate success: http.StatusOK
	GetBuiltIn func(ctx context.Context, policySetDefinitionName string, options *armresources.PolicySetDefinitionsClientGetBuiltInOptions) (resp azfake.Responder[armresources.PolicySetDefinitionsClientGetBuiltInResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PolicySetDefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armresources.PolicySetDefinitionsClientListOptions) (resp azfake.PagerResponder[armresources.PolicySetDefinitionsClientListResponse])

	// NewListBuiltInPager is the fake for method PolicySetDefinitionsClient.NewListBuiltInPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBuiltInPager func(options *armresources.PolicySetDefinitionsClientListBuiltInOptions) (resp azfake.PagerResponder[armresources.PolicySetDefinitionsClientListBuiltInResponse])

	// NewListByManagementGroupPager is the fake for method PolicySetDefinitionsClient.NewListByManagementGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagementGroupPager func(managementGroupID string, options *armresources.PolicySetDefinitionsClientListByManagementGroupOptions) (resp azfake.PagerResponder[armresources.PolicySetDefinitionsClientListByManagementGroupResponse])
}

// NewPolicySetDefinitionsServerTransport creates a new instance of PolicySetDefinitionsServerTransport with the provided implementation.
// The returned PolicySetDefinitionsServerTransport instance is connected to an instance of armresources.PolicySetDefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPolicySetDefinitionsServerTransport(srv *PolicySetDefinitionsServer) *PolicySetDefinitionsServerTransport {
	return &PolicySetDefinitionsServerTransport{
		srv:                           srv,
		newListPager:                  newTracker[azfake.PagerResponder[armresources.PolicySetDefinitionsClientListResponse]](),
		newListBuiltInPager:           newTracker[azfake.PagerResponder[armresources.PolicySetDefinitionsClientListBuiltInResponse]](),
		newListByManagementGroupPager: newTracker[azfake.PagerResponder[armresources.PolicySetDefinitionsClientListByManagementGroupResponse]](),
	}
}

// PolicySetDefinitionsServerTransport connects instances of armresources.PolicySetDefinitionsClient to instances of PolicySetDefinitionsServer.
// Don't use this type directly, use NewPolicySetDefinitionsServerTransport instead.
type PolicySetDefinitionsServerTransport struct {
	srv                           *PolicySetDefinitionsServer
	newListPager                  *tracker[azfake.PagerResponder[armresources.PolicySetDefinitionsClientListResponse]]
	newListBuiltInPager           *tracker[azfake.PagerResponder[armresources.PolicySetDefinitionsClientListBuiltInResponse]]
	newListByManagementGroupPager *tracker[azfake.PagerResponder[armresources.PolicySetDefinitionsClientListByManagementGroupResponse]]
}

// Do implements the policy.Transporter interface for PolicySetDefinitionsServerTransport.
func (p *PolicySetDefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PolicySetDefinitionsClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup":
		resp, err = p.dispatchCreateOrUpdateAtManagementGroup(req)
	case "PolicySetDefinitionsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PolicySetDefinitionsClient.DeleteAtManagementGroup":
		resp, err = p.dispatchDeleteAtManagementGroup(req)
	case "PolicySetDefinitionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PolicySetDefinitionsClient.GetAtManagementGroup":
		resp, err = p.dispatchGetAtManagementGroup(req)
	case "PolicySetDefinitionsClient.GetBuiltIn":
		resp, err = p.dispatchGetBuiltIn(req)
	case "PolicySetDefinitionsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PolicySetDefinitionsClient.NewListBuiltInPager":
		resp, err = p.dispatchNewListBuiltInPager(req)
	case "PolicySetDefinitionsClient.NewListByManagementGroupPager":
		resp, err = p.dispatchNewListByManagementGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PolicySetDefinitionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.PolicySetDefinition](req)
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), policySetDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionsServerTransport) dispatchCreateOrUpdateAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdateAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.PolicySetDefinition](req)
	if err != nil {
		return nil, err
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdateAtManagementGroup(req.Context(), managementGroupIDParam, policySetDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), policySetDefinitionNameParam, nil)
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

func (p *PolicySetDefinitionsServerTransport) dispatchDeleteAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.DeleteAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DeleteAtManagementGroup(req.Context(), managementGroupIDParam, policySetDefinitionNameParam, nil)
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

func (p *PolicySetDefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), policySetDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionsServerTransport) dispatchGetAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.GetAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupId")])
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetAtManagementGroup(req.Context(), managementGroupIDParam, policySetDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionsServerTransport) dispatchGetBuiltIn(req *http.Request) (*http.Response, error) {
	if p.srv.GetBuiltIn == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBuiltIn not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetBuiltIn(req.Context(), policySetDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions`
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
		var options *armresources.PolicySetDefinitionsClientListOptions
		if filterParam != nil || topParam != nil {
			options = &armresources.PolicySetDefinitionsClientListOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListPager(options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armresources.PolicySetDefinitionsClientListResponse, createLink func() string) {
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

func (p *PolicySetDefinitionsServerTransport) dispatchNewListBuiltInPager(req *http.Request) (*http.Response, error) {
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
		var options *armresources.PolicySetDefinitionsClientListBuiltInOptions
		if filterParam != nil || topParam != nil {
			options = &armresources.PolicySetDefinitionsClientListBuiltInOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListBuiltInPager(options)
		newListBuiltInPager = &resp
		p.newListBuiltInPager.add(req, newListBuiltInPager)
		server.PagerResponderInjectNextLinks(newListBuiltInPager, req, func(page *armresources.PolicySetDefinitionsClientListBuiltInResponse, createLink func() string) {
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

func (p *PolicySetDefinitionsServerTransport) dispatchNewListByManagementGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByManagementGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagementGroupPager not implemented")}
	}
	newListByManagementGroupPager := p.newListByManagementGroupPager.get(req)
	if newListByManagementGroupPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions`
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
		var options *armresources.PolicySetDefinitionsClientListByManagementGroupOptions
		if filterParam != nil || topParam != nil {
			options = &armresources.PolicySetDefinitionsClientListByManagementGroupOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := p.srv.NewListByManagementGroupPager(managementGroupIDParam, options)
		newListByManagementGroupPager = &resp
		p.newListByManagementGroupPager.add(req, newListByManagementGroupPager)
		server.PagerResponderInjectNextLinks(newListByManagementGroupPager, req, func(page *armresources.PolicySetDefinitionsClientListByManagementGroupResponse, createLink func() string) {
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
