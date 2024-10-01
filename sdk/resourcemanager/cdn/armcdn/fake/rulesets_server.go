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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
	"net/http"
	"net/url"
	"regexp"
)

// RuleSetsServer is a fake server for instances of the armcdn.RuleSetsClient type.
type RuleSetsServer struct {
	// Create is the fake for method RuleSetsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *armcdn.RuleSetsClientCreateOptions) (resp azfake.Responder[armcdn.RuleSetsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RuleSetsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *armcdn.RuleSetsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcdn.RuleSetsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RuleSetsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, ruleSetName string, options *armcdn.RuleSetsClientGetOptions) (resp azfake.Responder[armcdn.RuleSetsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProfilePager is the fake for method RuleSetsClient.NewListByProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProfilePager func(resourceGroupName string, profileName string, options *armcdn.RuleSetsClientListByProfileOptions) (resp azfake.PagerResponder[armcdn.RuleSetsClientListByProfileResponse])

	// NewListResourceUsagePager is the fake for method RuleSetsClient.NewListResourceUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListResourceUsagePager func(resourceGroupName string, profileName string, ruleSetName string, options *armcdn.RuleSetsClientListResourceUsageOptions) (resp azfake.PagerResponder[armcdn.RuleSetsClientListResourceUsageResponse])
}

// NewRuleSetsServerTransport creates a new instance of RuleSetsServerTransport with the provided implementation.
// The returned RuleSetsServerTransport instance is connected to an instance of armcdn.RuleSetsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRuleSetsServerTransport(srv *RuleSetsServer) *RuleSetsServerTransport {
	return &RuleSetsServerTransport{
		srv:                       srv,
		beginDelete:               newTracker[azfake.PollerResponder[armcdn.RuleSetsClientDeleteResponse]](),
		newListByProfilePager:     newTracker[azfake.PagerResponder[armcdn.RuleSetsClientListByProfileResponse]](),
		newListResourceUsagePager: newTracker[azfake.PagerResponder[armcdn.RuleSetsClientListResourceUsageResponse]](),
	}
}

// RuleSetsServerTransport connects instances of armcdn.RuleSetsClient to instances of RuleSetsServer.
// Don't use this type directly, use NewRuleSetsServerTransport instead.
type RuleSetsServerTransport struct {
	srv                       *RuleSetsServer
	beginDelete               *tracker[azfake.PollerResponder[armcdn.RuleSetsClientDeleteResponse]]
	newListByProfilePager     *tracker[azfake.PagerResponder[armcdn.RuleSetsClientListByProfileResponse]]
	newListResourceUsagePager *tracker[azfake.PagerResponder[armcdn.RuleSetsClientListResourceUsageResponse]]
}

// Do implements the policy.Transporter interface for RuleSetsServerTransport.
func (r *RuleSetsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RuleSetsClient.Create":
		resp, err = r.dispatchCreate(req)
	case "RuleSetsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RuleSetsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RuleSetsClient.NewListByProfilePager":
		resp, err = r.dispatchNewListByProfilePager(req)
	case "RuleSetsClient.NewListResourceUsagePager":
		resp, err = r.dispatchNewListResourceUsagePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RuleSetsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if r.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleSets/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Create(req.Context(), resourceGroupNameParam, profileNameParam, ruleSetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RuleSet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RuleSetsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleSets/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, profileNameParam, ruleSetNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RuleSetsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleSets/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, ruleSetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RuleSet, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RuleSetsServerTransport) dispatchNewListByProfilePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProfilePager not implemented")}
	}
	newListByProfilePager := r.newListByProfilePager.get(req)
	if newListByProfilePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleSets`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByProfilePager(resourceGroupNameParam, profileNameParam, nil)
		newListByProfilePager = &resp
		r.newListByProfilePager.add(req, newListByProfilePager)
		server.PagerResponderInjectNextLinks(newListByProfilePager, req, func(page *armcdn.RuleSetsClientListByProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProfilePager) {
		r.newListByProfilePager.remove(req)
	}
	return resp, nil
}

func (r *RuleSetsServerTransport) dispatchNewListResourceUsagePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListResourceUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListResourceUsagePager not implemented")}
	}
	newListResourceUsagePager := r.newListResourceUsagePager.get(req)
	if newListResourceUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ruleSets/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListResourceUsagePager(resourceGroupNameParam, profileNameParam, ruleSetNameParam, nil)
		newListResourceUsagePager = &resp
		r.newListResourceUsagePager.add(req, newListResourceUsagePager)
		server.PagerResponderInjectNextLinks(newListResourceUsagePager, req, func(page *armcdn.RuleSetsClientListResourceUsageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListResourceUsagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListResourceUsagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListResourceUsagePager) {
		r.newListResourceUsagePager.remove(req)
	}
	return resp, nil
}
