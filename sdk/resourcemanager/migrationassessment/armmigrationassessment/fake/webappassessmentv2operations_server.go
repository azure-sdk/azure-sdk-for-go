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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationassessment/armmigrationassessment"
	"net/http"
	"net/url"
	"regexp"
)

// WebAppAssessmentV2OperationsServer is a fake server for instances of the armmigrationassessment.WebAppAssessmentV2OperationsClient type.
type WebAppAssessmentV2OperationsServer struct {
	// BeginCreate is the fake for method WebAppAssessmentV2OperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, resource armmigrationassessment.WebAppAssessmentV2, options *armmigrationassessment.WebAppAssessmentV2OperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WebAppAssessmentV2OperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *armmigrationassessment.WebAppAssessmentV2OperationsClientDeleteOptions) (resp azfake.Responder[armmigrationassessment.WebAppAssessmentV2OperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDownloadURL is the fake for method WebAppAssessmentV2OperationsClient.BeginDownloadURL
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDownloadURL func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, body any, options *armmigrationassessment.WebAppAssessmentV2OperationsClientBeginDownloadURLOptions) (resp azfake.PollerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientDownloadURLResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WebAppAssessmentV2OperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *armmigrationassessment.WebAppAssessmentV2OperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.WebAppAssessmentV2OperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByGroupPager is the fake for method WebAppAssessmentV2OperationsClient.NewListByGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByGroupPager func(resourceGroupName string, projectName string, groupName string, options *armmigrationassessment.WebAppAssessmentV2OperationsClientListByGroupOptions) (resp azfake.PagerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientListByGroupResponse])
}

// NewWebAppAssessmentV2OperationsServerTransport creates a new instance of WebAppAssessmentV2OperationsServerTransport with the provided implementation.
// The returned WebAppAssessmentV2OperationsServerTransport instance is connected to an instance of armmigrationassessment.WebAppAssessmentV2OperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebAppAssessmentV2OperationsServerTransport(srv *WebAppAssessmentV2OperationsServer) *WebAppAssessmentV2OperationsServerTransport {
	return &WebAppAssessmentV2OperationsServerTransport{
		srv:                 srv,
		beginCreate:         newTracker[azfake.PollerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientCreateResponse]](),
		beginDownloadURL:    newTracker[azfake.PollerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientDownloadURLResponse]](),
		newListByGroupPager: newTracker[azfake.PagerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientListByGroupResponse]](),
	}
}

// WebAppAssessmentV2OperationsServerTransport connects instances of armmigrationassessment.WebAppAssessmentV2OperationsClient to instances of WebAppAssessmentV2OperationsServer.
// Don't use this type directly, use NewWebAppAssessmentV2OperationsServerTransport instead.
type WebAppAssessmentV2OperationsServerTransport struct {
	srv                 *WebAppAssessmentV2OperationsServer
	beginCreate         *tracker[azfake.PollerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientCreateResponse]]
	beginDownloadURL    *tracker[azfake.PollerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientDownloadURLResponse]]
	newListByGroupPager *tracker[azfake.PagerResponder[armmigrationassessment.WebAppAssessmentV2OperationsClientListByGroupResponse]]
}

// Do implements the policy.Transporter interface for WebAppAssessmentV2OperationsServerTransport.
func (w *WebAppAssessmentV2OperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WebAppAssessmentV2OperationsClient.BeginCreate":
		resp, err = w.dispatchBeginCreate(req)
	case "WebAppAssessmentV2OperationsClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WebAppAssessmentV2OperationsClient.BeginDownloadURL":
		resp, err = w.dispatchBeginDownloadURL(req)
	case "WebAppAssessmentV2OperationsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WebAppAssessmentV2OperationsClient.NewListByGroupPager":
		resp, err = w.dispatchNewListByGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WebAppAssessmentV2OperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := w.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationassessment.WebAppAssessmentV2](req)
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
		groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
		if err != nil {
			return nil, err
		}
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginCreate(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		w.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		w.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		w.beginCreate.remove(req)
	}

	return resp, nil
}

func (w *WebAppAssessmentV2OperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, nil)
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

func (w *WebAppAssessmentV2OperationsServerTransport) dispatchBeginDownloadURL(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDownloadURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDownloadURL not implemented")}
	}
	beginDownloadURL := w.beginDownloadURL.get(req)
	if beginDownloadURL == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/downloadUrl`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[any](req)
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
		groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
		if err != nil {
			return nil, err
		}
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginDownloadURL(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDownloadURL = &respr
		w.beginDownloadURL.add(req, beginDownloadURL)
	}

	resp, err := server.PollerResponderNext(beginDownloadURL, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginDownloadURL.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDownloadURL) {
		w.beginDownloadURL.remove(req)
	}

	return resp, nil
}

func (w *WebAppAssessmentV2OperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebAppAssessmentV2, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebAppAssessmentV2OperationsServerTransport) dispatchNewListByGroupPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByGroupPager not implemented")}
	}
	newListByGroupPager := w.newListByGroupPager.get(req)
	if newListByGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppAssessments`
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
		groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByGroupPager(resourceGroupNameParam, projectNameParam, groupNameParam, nil)
		newListByGroupPager = &resp
		w.newListByGroupPager.add(req, newListByGroupPager)
		server.PagerResponderInjectNextLinks(newListByGroupPager, req, func(page *armmigrationassessment.WebAppAssessmentV2OperationsClientListByGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByGroupPager) {
		w.newListByGroupPager.remove(req)
	}
	return resp, nil
}
