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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armpolicy"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DefinitionsServer is a fake server for instances of the armpolicy.DefinitionsClient type.
type DefinitionsServer struct {
	// CreateOrUpdate is the fake for method DefinitionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusCreated
	CreateOrUpdate func(ctx context.Context, policyDefinitionName string, parameters armpolicy.Definition, options *armpolicy.DefinitionsClientCreateOrUpdateOptions) (resp azfake.Responder[armpolicy.DefinitionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAtManagementGroup is the fake for method DefinitionsClient.CreateOrUpdateAtManagementGroup
	// HTTP status codes to indicate success: http.StatusCreated
	CreateOrUpdateAtManagementGroup func(ctx context.Context, managementGroupID string, policyDefinitionName string, parameters armpolicy.Definition, options *armpolicy.DefinitionsClientCreateOrUpdateAtManagementGroupOptions) (resp azfake.Responder[armpolicy.DefinitionsClientCreateOrUpdateAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DefinitionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, policyDefinitionName string, options *armpolicy.DefinitionsClientDeleteOptions) (resp azfake.Responder[armpolicy.DefinitionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAtManagementGroup is the fake for method DefinitionsClient.DeleteAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAtManagementGroup func(ctx context.Context, managementGroupID string, policyDefinitionName string, options *armpolicy.DefinitionsClientDeleteAtManagementGroupOptions) (resp azfake.Responder[armpolicy.DefinitionsClientDeleteAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DefinitionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, policyDefinitionName string, options *armpolicy.DefinitionsClientGetOptions) (resp azfake.Responder[armpolicy.DefinitionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAtManagementGroup is the fake for method DefinitionsClient.GetAtManagementGroup
	// HTTP status codes to indicate success: http.StatusOK
	GetAtManagementGroup func(ctx context.Context, managementGroupID string, policyDefinitionName string, options *armpolicy.DefinitionsClientGetAtManagementGroupOptions) (resp azfake.Responder[armpolicy.DefinitionsClientGetAtManagementGroupResponse], errResp azfake.ErrorResponder)

	// GetBuiltIn is the fake for method DefinitionsClient.GetBuiltIn
	// HTTP status codes to indicate success: http.StatusOK
	GetBuiltIn func(ctx context.Context, policyDefinitionName string, options *armpolicy.DefinitionsClientGetBuiltInOptions) (resp azfake.Responder[armpolicy.DefinitionsClientGetBuiltInResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DefinitionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armpolicy.DefinitionsClientListOptions) (resp azfake.PagerResponder[armpolicy.DefinitionsClientListResponse])

	// NewListBuiltInPager is the fake for method DefinitionsClient.NewListBuiltInPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBuiltInPager func(options *armpolicy.DefinitionsClientListBuiltInOptions) (resp azfake.PagerResponder[armpolicy.DefinitionsClientListBuiltInResponse])

	// NewListByManagementGroupPager is the fake for method DefinitionsClient.NewListByManagementGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByManagementGroupPager func(managementGroupID string, options *armpolicy.DefinitionsClientListByManagementGroupOptions) (resp azfake.PagerResponder[armpolicy.DefinitionsClientListByManagementGroupResponse])
}

// NewDefinitionsServerTransport creates a new instance of DefinitionsServerTransport with the provided implementation.
// The returned DefinitionsServerTransport instance is connected to an instance of armpolicy.DefinitionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDefinitionsServerTransport(srv *DefinitionsServer) *DefinitionsServerTransport {
	return &DefinitionsServerTransport{
		srv:                           srv,
		newListPager:                  newTracker[azfake.PagerResponder[armpolicy.DefinitionsClientListResponse]](),
		newListBuiltInPager:           newTracker[azfake.PagerResponder[armpolicy.DefinitionsClientListBuiltInResponse]](),
		newListByManagementGroupPager: newTracker[azfake.PagerResponder[armpolicy.DefinitionsClientListByManagementGroupResponse]](),
	}
}

// DefinitionsServerTransport connects instances of armpolicy.DefinitionsClient to instances of DefinitionsServer.
// Don't use this type directly, use NewDefinitionsServerTransport instead.
type DefinitionsServerTransport struct {
	srv                           *DefinitionsServer
	newListPager                  *tracker[azfake.PagerResponder[armpolicy.DefinitionsClientListResponse]]
	newListBuiltInPager           *tracker[azfake.PagerResponder[armpolicy.DefinitionsClientListBuiltInResponse]]
	newListByManagementGroupPager *tracker[azfake.PagerResponder[armpolicy.DefinitionsClientListByManagementGroupResponse]]
}

// Do implements the policy.Transporter interface for DefinitionsServerTransport.
func (d *DefinitionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DefinitionsClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DefinitionsClient.CreateOrUpdateAtManagementGroup":
		resp, err = d.dispatchCreateOrUpdateAtManagementGroup(req)
	case "DefinitionsClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DefinitionsClient.DeleteAtManagementGroup":
		resp, err = d.dispatchDeleteAtManagementGroup(req)
	case "DefinitionsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DefinitionsClient.GetAtManagementGroup":
		resp, err = d.dispatchGetAtManagementGroup(req)
	case "DefinitionsClient.GetBuiltIn":
		resp, err = d.dispatchGetBuiltIn(req)
	case "DefinitionsClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DefinitionsClient.NewListBuiltInPager":
		resp, err = d.dispatchNewListBuiltInPager(req)
	case "DefinitionsClient.NewListByManagementGroupPager":
		resp, err = d.dispatchNewListByManagementGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpolicy.Definition](req)
	if err != nil {
		return nil, err
	}
	policyDefinitionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("policyDefinitionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), policyDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Definition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchCreateOrUpdateAtManagementGroup(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdateAtManagementGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAtManagementGroup not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Management/managementGroups/(?P<managementGroupId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Authorization/policyDefinitions/(?P<policyDefinitionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armpolicy.Definition](req)
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
	respr, errRespr := d.srv.CreateOrUpdateAtManagementGroup(req.Context(), managementGroupIDParam, policyDefinitionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Definition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
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
	respr, errRespr := d.srv.Delete(req.Context(), policyDefinitionNameParam, nil)
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

func (d *DefinitionsServerTransport) dispatchDeleteAtManagementGroup(req *http.Request) (*http.Response, error) {
	if d.srv.DeleteAtManagementGroup == nil {
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
	respr, errRespr := d.srv.DeleteAtManagementGroup(req.Context(), managementGroupIDParam, policyDefinitionNameParam, nil)
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

func (d *DefinitionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
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
	respr, errRespr := d.srv.Get(req.Context(), policyDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Definition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchGetAtManagementGroup(req *http.Request) (*http.Response, error) {
	if d.srv.GetAtManagementGroup == nil {
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
	respr, errRespr := d.srv.GetAtManagementGroup(req.Context(), managementGroupIDParam, policyDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Definition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchGetBuiltIn(req *http.Request) (*http.Response, error) {
	if d.srv.GetBuiltIn == nil {
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
	respr, errRespr := d.srv.GetBuiltIn(req.Context(), policyDefinitionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Definition, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
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
		var options *armpolicy.DefinitionsClientListOptions
		if filterParam != nil || topParam != nil {
			options = &armpolicy.DefinitionsClientListOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := d.srv.NewListPager(options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpolicy.DefinitionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchNewListBuiltInPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBuiltInPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBuiltInPager not implemented")}
	}
	newListBuiltInPager := d.newListBuiltInPager.get(req)
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
		var options *armpolicy.DefinitionsClientListBuiltInOptions
		if filterParam != nil || topParam != nil {
			options = &armpolicy.DefinitionsClientListBuiltInOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := d.srv.NewListBuiltInPager(options)
		newListBuiltInPager = &resp
		d.newListBuiltInPager.add(req, newListBuiltInPager)
		server.PagerResponderInjectNextLinks(newListBuiltInPager, req, func(page *armpolicy.DefinitionsClientListBuiltInResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBuiltInPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListBuiltInPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBuiltInPager) {
		d.newListBuiltInPager.remove(req)
	}
	return resp, nil
}

func (d *DefinitionsServerTransport) dispatchNewListByManagementGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByManagementGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByManagementGroupPager not implemented")}
	}
	newListByManagementGroupPager := d.newListByManagementGroupPager.get(req)
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
		var options *armpolicy.DefinitionsClientListByManagementGroupOptions
		if filterParam != nil || topParam != nil {
			options = &armpolicy.DefinitionsClientListByManagementGroupOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := d.srv.NewListByManagementGroupPager(managementGroupIDParam, options)
		newListByManagementGroupPager = &resp
		d.newListByManagementGroupPager.add(req, newListByManagementGroupPager)
		server.PagerResponderInjectNextLinks(newListByManagementGroupPager, req, func(page *armpolicy.DefinitionsClientListByManagementGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByManagementGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByManagementGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByManagementGroupPager) {
		d.newListByManagementGroupPager.remove(req)
	}
	return resp, nil
}
