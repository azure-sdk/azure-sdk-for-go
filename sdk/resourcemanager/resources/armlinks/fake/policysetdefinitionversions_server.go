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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PolicySetDefinitionVersionsServer is a fake server for instances of the armlinks.PolicySetDefinitionVersionsClient type.
type PolicySetDefinitionVersionsServer struct {
	// CreateOrUpdate is the fake for method PolicySetDefinitionVersionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, policySetDefinitionName string, policyDefinitionVersion string, parameters armlinks.PolicySetDefinitionVersion, options *armlinks.PolicySetDefinitionVersionsClientCreateOrUpdateOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtManagementGroup is the fake for method PolicySetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdateAtManagementGroup func(ctx context.Context, managementGroupName string, policySetDefinitionName string, policyDefinitionVersion string, parameters armlinks.PolicySetDefinitionVersion, options *armlinks.PolicySetDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientCreateOrUpdateAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PolicySetDefinitionVersionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, policySetDefinitionName string, policyDefinitionVersion string, options *armlinks.PolicySetDefinitionVersionsClientDeleteOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtManagementGroup is the fake for method PolicySetDefinitionVersionsClient.DeleteAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAtManagementGroup func(ctx context.Context, managementGroupName string, policySetDefinitionName string, policyDefinitionVersion string, options *armlinks.PolicySetDefinitionVersionsClientDeleteAtManagementGroupOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientDeleteAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PolicySetDefinitionVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, policySetDefinitionName string, policyDefinitionVersion string, options *armlinks.PolicySetDefinitionVersionsClientGetOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtManagementGroup is the fake for method PolicySetDefinitionVersionsClient.GetAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	GetAtManagementGroup func(ctx context.Context, managementGroupName string, policySetDefinitionName string, policyDefinitionVersion string, options *armlinks.PolicySetDefinitionVersionsClientGetAtManagementGroupOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientGetAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// GetBuiltIn is the fake for method PolicySetDefinitionVersionsClient.GetBuiltIn
	// HTTP status codes to indicate success: http.StatusOK
	GetBuiltIn func(ctx context.Context, policySetDefinitionName string, policyDefinitionVersion string, options *armlinks.PolicySetDefinitionVersionsClientGetBuiltInOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientGetBuiltInResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PolicySetDefinitionVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(policySetDefinitionName string, options *armlinks.PolicySetDefinitionVersionsClientListOptions) (resp azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListResponse])

	// ListAll is the fake for method PolicySetDefinitionVersionsClient.ListAll
	// HTTP status codes to indicate success: http.StatusOK
	ListAll func(ctx context.Context, options *armlinks.PolicySetDefinitionVersionsClientListAllOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientListAllResponse], errResp azfake.ErrorResponder)

	// ListAllAtManagementGroup is the fake for method PolicySetDefinitionVersionsClient.ListAllAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	ListAllAtManagementGroup func(ctx context.Context, managementGroupName string, options *armlinks.PolicySetDefinitionVersionsClientListAllAtManagementGroupOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientListAllAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// ListAllBuiltins is the fake for method PolicySetDefinitionVersionsClient.ListAllBuiltins
	// HTTP status codes to indicate success: http.StatusOK
	ListAllBuiltins func(ctx context.Context, options *armlinks.PolicySetDefinitionVersionsClientListAllBuiltinsOptions) (resp azfake.Responder[armlinks.PolicySetDefinitionVersionsClientListAllBuiltinsResponse], errResp azfake.ErrorResponder)

	// NewListBuiltInPager is the fake for method PolicySetDefinitionVersionsClient.NewListBuiltInPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBuiltInPager func(policySetDefinitionName string, options *armlinks.PolicySetDefinitionVersionsClientListBuiltInOptions) (resp azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListBuiltInResponse])

	// NewListByManagementGroupPager is the fake for method PolicySetDefinitionVersionsClient.NewListByManagementGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagementGroupPager func(managementGroupName string, policySetDefinitionName string, options *armlinks.PolicySetDefinitionVersionsClientListByManagementGroupOptions) (resp azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListByManagementGroupResponse])
}

// NewPolicySetDefinitionVersionsServerTransport creates a new instance of PolicySetDefinitionVersionsServerTransport with the provided implementation.
// The returned PolicySetDefinitionVersionsServerTransport instance is connected to an instance of armlinks.PolicySetDefinitionVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPolicySetDefinitionVersionsServerTransport(srv *PolicySetDefinitionVersionsServer) *PolicySetDefinitionVersionsServerTransport {
	return &PolicySetDefinitionVersionsServerTransport{
		srv:                           srv,
		newListPager:                  newTracker[azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListResponse]](),
		newListBuiltInPager:           newTracker[azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListBuiltInResponse]](),
		newListByManagementGroupPager: newTracker[azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListByManagementGroupResponse]](),
	}
}

// PolicySetDefinitionVersionsServerTransport connects instances of armlinks.PolicySetDefinitionVersionsClient to instances of PolicySetDefinitionVersionsServer.
// Don't use this type directly, use NewPolicySetDefinitionVersionsServerTransport instead.
type PolicySetDefinitionVersionsServerTransport struct {
	srv                           *PolicySetDefinitionVersionsServer
	newListPager                  *tracker[azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListResponse]]
	newListBuiltInPager           *tracker[azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListBuiltInResponse]]
	newListByManagementGroupPager *tracker[azfake.PagerResponder[armlinks.PolicySetDefinitionVersionsClientListByManagementGroupResponse]]
}

// Do implements the policy.Transporter interface for PolicySetDefinitionVersionsServerTransport.
func (p *PolicySetDefinitionVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PolicySetDefinitionVersionsClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PolicySetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup":
		resp, err = p.dispatchCreateOrUpdateAtManagementGroup(req)
	case "PolicySetDefinitionVersionsClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PolicySetDefinitionVersionsClient.DeleteAtManagementGroup":
		resp, err = p.dispatchDeleteAtManagementGroup(req)
	case "PolicySetDefinitionVersionsClient.Get":
		resp, err = p.dispatchGet(req)
	case "PolicySetDefinitionVersionsClient.GetAtManagementGroup":
		resp, err = p.dispatchGetAtManagementGroup(req)
	case "PolicySetDefinitionVersionsClient.GetBuiltIn":
		resp, err = p.dispatchGetBuiltIn(req)
	case "PolicySetDefinitionVersionsClient.NewListPager":
		resp, err = p.dispatchNewListPager(req)
	case "PolicySetDefinitionVersionsClient.ListAll":
		resp, err = p.dispatchListAll(req)
	case "PolicySetDefinitionVersionsClient.ListAllAtManagementGroup":
		resp, err = p.dispatchListAllAtManagementGroup(req)
	case "PolicySetDefinitionVersionsClient.ListAllBuiltins":
		resp, err = p.dispatchListAllBuiltins(req)
	case "PolicySetDefinitionVersionsClient.NewListBuiltInPager":
		resp, err = p.dispatchNewListBuiltInPager(req)
	case "PolicySetDefinitionVersionsClient.NewListByManagementGroupPager":
		resp, err = p.dispatchNewListByManagementGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlinks.PolicySetDefinitionVersion](req)
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), policySetDefinitionNameParam, policyDefinitionVersionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchCreateOrUpdateAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdateAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlinks.PolicySetDefinitionVersion](req)
	if err != nil {
		return nil, err
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.CreateOrUpdateAtManagementGroup(req.Context(), managementGroupNameParam, policySetDefinitionNameParam, policyDefinitionVersionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), policySetDefinitionNameParam, policyDefinitionVersionParam, nil)
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

func (p *PolicySetDefinitionVersionsServerTransport) dispatchDeleteAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.DeleteAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.DeleteAtManagementGroup(req.Context(), managementGroupNameParam, policySetDefinitionNameParam, policyDefinitionVersionParam, nil)
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

func (p *PolicySetDefinitionVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), policySetDefinitionNameParam, policyDefinitionVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchGetAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.GetAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	managementGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managementGroupName")])
	if err != nil {
		return nil, err
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetAtManagementGroup(req.Context(), managementGroupNameParam, policySetDefinitionNameParam, policyDefinitionVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchGetBuiltIn(req *http.Request) (*http.Response, error) {
	if p.srv.GetBuiltIn == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetBuiltIn not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<policyDefinitionVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
	if err != nil {
		return nil, err
	}
	policyDefinitionVersionParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionVersion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetBuiltIn(req.Context(), policySetDefinitionNameParam, policyDefinitionVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
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
		var options *armlinks.PolicySetDefinitionVersionsClientListOptions
		if topParam != nil {
			options = &armlinks.PolicySetDefinitionVersionsClientListOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListPager(policySetDefinitionNameParam, options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlinks.PolicySetDefinitionVersionsClientListResponse, createLink func() string) {
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

func (p *PolicySetDefinitionVersionsServerTransport) dispatchListAll(req *http.Request) (*http.Response, error) {
	if p.srv.ListAll == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAll not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/listPolicySetDefinitionVersions`
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchListAllAtManagementGroup(req *http.Request) (*http.Response, error) {
	if p.srv.ListAllAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAllAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/listPolicySetDefinitionVersions`
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchListAllBuiltins(req *http.Request) (*http.Response, error) {
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicySetDefinitionVersionListResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PolicySetDefinitionVersionsServerTransport) dispatchNewListBuiltInPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListBuiltInPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBuiltInPager not implemented")}
	}
	newListBuiltInPager := p.newListBuiltInPager.get(req)
	if newListBuiltInPager == nil {
		const regexStr = `/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
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
		var options *armlinks.PolicySetDefinitionVersionsClientListBuiltInOptions
		if topParam != nil {
			options = &armlinks.PolicySetDefinitionVersionsClientListBuiltInOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListBuiltInPager(policySetDefinitionNameParam, options)
		newListBuiltInPager = &resp
		p.newListBuiltInPager.add(req, newListBuiltInPager)
		server.PagerResponderInjectNextLinks(newListBuiltInPager, req, func(page *armlinks.PolicySetDefinitionVersionsClientListBuiltInResponse, createLink func() string) {
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

func (p *PolicySetDefinitionVersionsServerTransport) dispatchNewListByManagementGroupPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByManagementGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagementGroupPager not implemented")}
	}
	newListByManagementGroupPager := p.newListByManagementGroupPager.get(req)
	if newListByManagementGroupPager == nil {
		const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policySetDefinitions/(?P<policySetDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
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
		policySetDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policySetDefinitionName")])
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
		var options *armlinks.PolicySetDefinitionVersionsClientListByManagementGroupOptions
		if topParam != nil {
			options = &armlinks.PolicySetDefinitionVersionsClientListByManagementGroupOptions{
				Top: topParam,
			}
		}
		resp := p.srv.NewListByManagementGroupPager(managementGroupNameParam, policySetDefinitionNameParam, options)
		newListByManagementGroupPager = &resp
		p.newListByManagementGroupPager.add(req, newListByManagementGroupPager)
		server.PagerResponderInjectNextLinks(newListByManagementGroupPager, req, func(page *armlinks.PolicySetDefinitionVersionsClientListByManagementGroupResponse, createLink func() string) {
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
