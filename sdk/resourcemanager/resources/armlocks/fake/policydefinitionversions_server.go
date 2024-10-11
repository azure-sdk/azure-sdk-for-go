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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlocks/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PolicyDefinitionVersionsServer is a fake server for instances of the armlocks.PolicyDefinitionVersionsClient type.
type PolicyDefinitionVersionsServer struct {
	// CreateOrUpdate is the fake for method PolicyDefinitionVersionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, policyDefinitionName string, policyDefinitionVersion string, parameters armlocks.PolicyDefinitionVersion, options *armlocks.PolicyDefinitionVersionsClientCreateOrUpdateOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtManagementGroup is the fake for method PolicyDefinitionVersionsClient.CreateOrUpdateAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdateAtManagementGroup func(ctx context.Context, managementGroupName string, policyDefinitionName string, policyDefinitionVersion string, parameters armlocks.PolicyDefinitionVersion, options *armlocks.PolicyDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PolicyDefinitionVersionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, policyDefinitionName string, policyDefinitionVersion string, options *armlocks.PolicyDefinitionVersionsClientDeleteOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtManagementGroup is the fake for method PolicyDefinitionVersionsClient.DeleteAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAtManagementGroup func(ctx context.Context, managementGroupName string, policyDefinitionName string, policyDefinitionVersion string, options *armlocks.PolicyDefinitionVersionsClientDeleteAtManagementGroupOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientDeleteAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PolicyDefinitionVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, policyDefinitionName string, policyDefinitionVersion string, options *armlocks.PolicyDefinitionVersionsClientGetOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtManagementGroup is the fake for method PolicyDefinitionVersionsClient.GetAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	GetAtManagementGroup func(ctx context.Context, managementGroupName string, policyDefinitionName string, policyDefinitionVersion string, options *armlocks.PolicyDefinitionVersionsClientGetAtManagementGroupOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientGetAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// GetBuiltIn is the fake for method PolicyDefinitionVersionsClient.GetBuiltIn
	// HTTP status codes to indicate success: http.StatusOK
	GetBuiltIn func(ctx context.Context, policyDefinitionName string, policyDefinitionVersion string, options *armlocks.PolicyDefinitionVersionsClientGetBuiltInOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientGetBuiltInResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PolicyDefinitionVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(policyDefinitionName string, options *armlocks.PolicyDefinitionVersionsClientListOptions) (resp azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListResponse])

	// ListAll is the fake for method PolicyDefinitionVersionsClient.ListAll
	// HTTP status codes to indicate success: http.StatusOK
	ListAll func(ctx context.Context, options *armlocks.PolicyDefinitionVersionsClientListAllOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientListAllResponse], errResp azfake.ErrorResponder)

	// ListAllAtManagementGroup is the fake for method PolicyDefinitionVersionsClient.ListAllAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	ListAllAtManagementGroup func(ctx context.Context, managementGroupName string, options *armlocks.PolicyDefinitionVersionsClientListAllAtManagementGroupOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientListAllAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// ListAllBuiltins is the fake for method PolicyDefinitionVersionsClient.ListAllBuiltins
	// HTTP status codes to indicate success: http.StatusOK
	ListAllBuiltins func(ctx context.Context, options *armlocks.PolicyDefinitionVersionsClientListAllBuiltinsOptions) (resp azfake.Responder[armlocks.PolicyDefinitionVersionsClientListAllBuiltinsResponse], errResp azfake.ErrorResponder)

	// NewListBuiltInPager is the fake for method PolicyDefinitionVersionsClient.NewListBuiltInPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBuiltInPager func(policyDefinitionName string, options *armlocks.PolicyDefinitionVersionsClientListBuiltInOptions) (resp azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListBuiltInResponse])

	// NewListByManagementGroupPager is the fake for method PolicyDefinitionVersionsClient.NewListByManagementGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagementGroupPager func(managementGroupName string, policyDefinitionName string, options *armlocks.PolicyDefinitionVersionsClientListByManagementGroupOptions) (resp azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListByManagementGroupResponse])
}

// NewPolicyDefinitionVersionsServerTransport creates a new instance of PolicyDefinitionVersionsServerTransport with the provided implementation.
// The returned PolicyDefinitionVersionsServerTransport instance is connected to an instance of armlocks.PolicyDefinitionVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPolicyDefinitionVersionsServerTransport(srv *PolicyDefinitionVersionsServer) *PolicyDefinitionVersionsServerTransport {
	return &PolicyDefinitionVersionsServerTransport{
		srv:                           srv,
		newListPager:                  newTracker[azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListResponse]](),
		newListBuiltInPager:           newTracker[azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListBuiltInResponse]](),
		newListByManagementGroupPager: newTracker[azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListByManagementGroupResponse]](),
	}
}

// PolicyDefinitionVersionsServerTransport connects instances of armlocks.PolicyDefinitionVersionsClient to instances of PolicyDefinitionVersionsServer.
// Don't use this type directly, use NewPolicyDefinitionVersionsServerTransport instead.
type PolicyDefinitionVersionsServerTransport struct {
	srv                           *PolicyDefinitionVersionsServer
	newListPager                  *tracker[azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListResponse]]
	newListBuiltInPager           *tracker[azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListBuiltInResponse]]
	newListByManagementGroupPager *tracker[azfake.PagerResponder[armlocks.PolicyDefinitionVersionsClientListByManagementGroupResponse]]
}

// Do implements the policy.Transporter interface for PolicyDefinitionVersionsServerTransport.
func (p *PolicyDefinitionVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PolicyDefinitionVersionsClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PolicyDefinitionVersionsClient.CreateOrUpdateAtManagementGroup":
		resp, err = p.dispatchCreateOrUpdateAtManagementGroup(req)
	case "PolicyDefinitionVersionsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PolicyDefinitionVersionsClient.DeleteAtManagementGroup":
		resp, err = p.dispatchDeleteAtManagementGroup(req)
	case "PolicyDefinitionVersionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PolicyDefinitionVersionsClient.GetAtManagementGroup":
		resp, err = p.dispatchGetAtManagementGroup(req)
	case "PolicyDefinitionVersionsClient.GetBuiltIn":
		resp, err = p.dispatchGetBuiltIn(req)
	case "PolicyDefinitionVersionsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PolicyDefinitionVersionsClient.ListAll":
		resp, err = p.dispatchListAll(req)
	case "PolicyDefinitionVersionsClient.ListAllAtManagementGroup":
		resp, err = p.dispatchListAllAtManagementGroup(req)
	case "PolicyDefinitionVersionsClient.ListAllBuiltins":
		resp, err = p.dispatchListAllBuiltins(req)
	case "PolicyDefinitionVersionsClient.NewListBuiltInPager":
		resp, err = p.dispatchNewListBuiltInPager(req)
	case "PolicyDefinitionVersionsClient.NewListByManagementGroupPager":
		resp, err = p.dispatchNewListByManagementGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlocks.PolicyDefinitionVersion](req)
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), policyDefinitionNameParam, policyDefinitionVersionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchCreateOrUpdateAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdateAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlocks.PolicyDefinitionVersion](req)
	if err != nil {
		return nil, err
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdateAtManagementGroup(req.Context(), managementGroupNameParam, policyDefinitionNameParam, policyDefinitionVersionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), policyDefinitionNameParam, policyDefinitionVersionParam, nil)
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

func (p *PolicyDefinitionVersionsServerTransport) dispatchDeleteAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.DeleteAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DeleteAtManagementGroup(req.Context(), managementGroupNameParam, policyDefinitionNameParam, policyDefinitionVersionParam, nil)
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

func (p *PolicyDefinitionVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), policyDefinitionNameParam, policyDefinitionVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchGetAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.GetAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetAtManagementGroup(req.Context(), managementGroupNameParam, policyDefinitionNameParam, policyDefinitionVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchGetBuiltIn(req *http.Request) (*http.Response, error) {
	if p.srv.GetBuiltIn == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBuiltIn not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetBuiltIn(req.Context(), policyDefinitionNameParam, policyDefinitionVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
		if err != nil {
			return nil, err
		}
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
		var options *armlocks.PolicyDefinitionVersionsClientListOptions
		if topParam != nil {
			options = &armlocks.PolicyDefinitionVersionsClientListOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListPager(policyDefinitionNameParam, options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlocks.PolicyDefinitionVersionsClientListResponse, createLink func() string) {
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

func (p *PolicyDefinitionVersionsServerTransport) dispatchListAll(req *http.Request) (*http.Response, error) {
	if p.srv.ListAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAll not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/listPolicyDefinitionVersions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := p.srv.ListAll(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchListAllAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.ListAllAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAllAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/listPolicyDefinitionVersions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.ListAllAtManagementGroup(req.Context(), managementGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchListAllBuiltins(req *http.Request) (*http.Response, error) {
	if p.srv.ListAllBuiltins == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAllBuiltins not implemented")}
	}
	respr, errRespr := p.srv.ListAllBuiltins(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyDefinitionVersionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicyDefinitionVersionsServerTransport) dispatchNewListBuiltInPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBuiltInPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBuiltInPager not implemented")}
	}
	newListBuiltInPager := p.newListBuiltInPager.get(req)
	if newListBuiltInPager == nil {
		const regexStr = `/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
		if err != nil {
			return nil, err
		}
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
		var options *armlocks.PolicyDefinitionVersionsClientListBuiltInOptions
		if topParam != nil {
			options = &armlocks.PolicyDefinitionVersionsClientListBuiltInOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListBuiltInPager(policyDefinitionNameParam, options)
		newListBuiltInPager = &resp
		p.newListBuiltInPager.add(req, newListBuiltInPager)
		server.PagerResponderInjectNextLinks(newListBuiltInPager, req, func(page *armlocks.PolicyDefinitionVersionsClientListBuiltInResponse, createLink func() string) {
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

func (p *PolicyDefinitionVersionsServerTransport) dispatchNewListByManagementGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByManagementGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagementGroupPager not implemented")}
	}
	newListByManagementGroupPager := p.newListByManagementGroupPager.get(req)
	if newListByManagementGroupPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
		if err != nil {
			return nil, err
		}
		policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
		if err != nil {
			return nil, err
		}
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
		var options *armlocks.PolicyDefinitionVersionsClientListByManagementGroupOptions
		if topParam != nil {
			options = &armlocks.PolicyDefinitionVersionsClientListByManagementGroupOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListByManagementGroupPager(managementGroupNameParam, policyDefinitionNameParam, options)
		newListByManagementGroupPager = &resp
		p.newListByManagementGroupPager.add(req, newListByManagementGroupPager)
		server.PagerResponderInjectNextLinks(newListByManagementGroupPager, req, func(page *armlocks.PolicyDefinitionVersionsClientListByManagementGroupResponse, createLink func() string) {
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
