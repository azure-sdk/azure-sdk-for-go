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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate/v2"
	"net/http"
	"net/url"
	"regexp"
)

// HypervCollectorsOperationsServer is a fake server for instances of the armmigrate.HypervCollectorsOperationsClient type.
type HypervCollectorsOperationsServer struct {
	// BeginCreate is the fake for method HypervCollectorsOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, projectName string, hypervCollectorName string, resource armmigrate.HypervCollector, options *armmigrate.HypervCollectorsOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrate.HypervCollectorsOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HypervCollectorsOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, hypervCollectorName string, options *armmigrate.HypervCollectorsOperationsClientDeleteOptions) (resp azfake.Responder[armmigrate.HypervCollectorsOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HypervCollectorsOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, hypervCollectorName string, options *armmigrate.HypervCollectorsOperationsClientGetOptions) (resp azfake.Responder[armmigrate.HypervCollectorsOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAssessmentProjectPager is the fake for method HypervCollectorsOperationsClient.NewListByAssessmentProjectPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAssessmentProjectPager func(resourceGroupName string, projectName string, options *armmigrate.HypervCollectorsOperationsClientListByAssessmentProjectOptions) (resp azfake.PagerResponder[armmigrate.HypervCollectorsOperationsClientListByAssessmentProjectResponse])
}

// NewHypervCollectorsOperationsServerTransport creates a new instance of HypervCollectorsOperationsServerTransport with the provided implementation.
// The returned HypervCollectorsOperationsServerTransport instance is connected to an instance of armmigrate.HypervCollectorsOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervCollectorsOperationsServerTransport(srv *HypervCollectorsOperationsServer) *HypervCollectorsOperationsServerTransport {
	return &HypervCollectorsOperationsServerTransport{
		srv:                             srv,
		beginCreate:                     newTracker[azfake.PollerResponder[armmigrate.HypervCollectorsOperationsClientCreateResponse]](),
		newListByAssessmentProjectPager: newTracker[azfake.PagerResponder[armmigrate.HypervCollectorsOperationsClientListByAssessmentProjectResponse]](),
	}
}

// HypervCollectorsOperationsServerTransport connects instances of armmigrate.HypervCollectorsOperationsClient to instances of HypervCollectorsOperationsServer.
// Don't use this type directly, use NewHypervCollectorsOperationsServerTransport instead.
type HypervCollectorsOperationsServerTransport struct {
	srv                             *HypervCollectorsOperationsServer
	beginCreate                     *tracker[azfake.PollerResponder[armmigrate.HypervCollectorsOperationsClientCreateResponse]]
	newListByAssessmentProjectPager *tracker[azfake.PagerResponder[armmigrate.HypervCollectorsOperationsClientListByAssessmentProjectResponse]]
}

// Do implements the policy.Transporter interface for HypervCollectorsOperationsServerTransport.
func (h *HypervCollectorsOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "HypervCollectorsOperationsClient.BeginCreate":
		resp, err = h.dispatchBeginCreate(req)
	case "HypervCollectorsOperationsClient.Delete":
		resp, err = h.dispatchDelete(req)
	case "HypervCollectorsOperationsClient.Get":
		resp, err = h.dispatchGet(req)
	case "HypervCollectorsOperationsClient.NewListByAssessmentProjectPager":
		resp, err = h.dispatchNewListByAssessmentProjectPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HypervCollectorsOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if h.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := h.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hypervcollectors/(?P<hypervCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.HypervCollector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		hypervCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hypervCollectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginCreate(req.Context(), resourceGroupNameParam, projectNameParam, hypervCollectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		h.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		h.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		h.beginCreate.remove(req)
	}

	return resp, nil
}

func (h *HypervCollectorsOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hypervcollectors/(?P<hypervCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	hypervCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hypervCollectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, hypervCollectorNameParam, nil)
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

func (h *HypervCollectorsOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hypervcollectors/(?P<hypervCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	hypervCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hypervCollectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, hypervCollectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervCollector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervCollectorsOperationsServerTransport) dispatchNewListByAssessmentProjectPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByAssessmentProjectPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAssessmentProjectPager not implemented")}
	}
	newListByAssessmentProjectPager := h.newListByAssessmentProjectPager.get(req)
	if newListByAssessmentProjectPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hypervcollectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		resp := h.srv.NewListByAssessmentProjectPager(resourceGroupNameParam, projectNameParam, nil)
		newListByAssessmentProjectPager = &resp
		h.newListByAssessmentProjectPager.add(req, newListByAssessmentProjectPager)
		server.PagerResponderInjectNextLinks(newListByAssessmentProjectPager, req, func(page *armmigrate.HypervCollectorsOperationsClientListByAssessmentProjectResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAssessmentProjectPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByAssessmentProjectPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAssessmentProjectPager) {
		h.newListByAssessmentProjectPager.remove(req)
	}
	return resp, nil
}
