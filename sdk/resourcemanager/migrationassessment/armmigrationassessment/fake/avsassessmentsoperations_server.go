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

// AvsAssessmentsOperationsServer is a fake server for instances of the armmigrationassessment.AvsAssessmentsOperationsClient type.
type AvsAssessmentsOperationsServer struct {
	// BeginCreate is the fake for method AvsAssessmentsOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, resource armmigrationassessment.AvsAssessment, options *armmigrationassessment.AvsAssessmentsOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrationassessment.AvsAssessmentsOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AvsAssessmentsOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *armmigrationassessment.AvsAssessmentsOperationsClientDeleteOptions) (resp azfake.Responder[armmigrationassessment.AvsAssessmentsOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDownloadURL is the fake for method AvsAssessmentsOperationsClient.BeginDownloadURL
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDownloadURL func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, body any, options *armmigrationassessment.AvsAssessmentsOperationsClientBeginDownloadURLOptions) (resp azfake.PollerResponder[armmigrationassessment.AvsAssessmentsOperationsClientDownloadURLResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AvsAssessmentsOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, options *armmigrationassessment.AvsAssessmentsOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.AvsAssessmentsOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByGroupPager is the fake for method AvsAssessmentsOperationsClient.NewListByGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByGroupPager func(resourceGroupName string, projectName string, groupName string, options *armmigrationassessment.AvsAssessmentsOperationsClientListByGroupOptions) (resp azfake.PagerResponder[armmigrationassessment.AvsAssessmentsOperationsClientListByGroupResponse])
}

// NewAvsAssessmentsOperationsServerTransport creates a new instance of AvsAssessmentsOperationsServerTransport with the provided implementation.
// The returned AvsAssessmentsOperationsServerTransport instance is connected to an instance of armmigrationassessment.AvsAssessmentsOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvsAssessmentsOperationsServerTransport(srv *AvsAssessmentsOperationsServer) *AvsAssessmentsOperationsServerTransport {
	return &AvsAssessmentsOperationsServerTransport{
		srv:                 srv,
		beginCreate:         newTracker[azfake.PollerResponder[armmigrationassessment.AvsAssessmentsOperationsClientCreateResponse]](),
		beginDownloadURL:    newTracker[azfake.PollerResponder[armmigrationassessment.AvsAssessmentsOperationsClientDownloadURLResponse]](),
		newListByGroupPager: newTracker[azfake.PagerResponder[armmigrationassessment.AvsAssessmentsOperationsClientListByGroupResponse]](),
	}
}

// AvsAssessmentsOperationsServerTransport connects instances of armmigrationassessment.AvsAssessmentsOperationsClient to instances of AvsAssessmentsOperationsServer.
// Don't use this type directly, use NewAvsAssessmentsOperationsServerTransport instead.
type AvsAssessmentsOperationsServerTransport struct {
	srv                 *AvsAssessmentsOperationsServer
	beginCreate         *tracker[azfake.PollerResponder[armmigrationassessment.AvsAssessmentsOperationsClientCreateResponse]]
	beginDownloadURL    *tracker[azfake.PollerResponder[armmigrationassessment.AvsAssessmentsOperationsClientDownloadURLResponse]]
	newListByGroupPager *tracker[azfake.PagerResponder[armmigrationassessment.AvsAssessmentsOperationsClientListByGroupResponse]]
}

// Do implements the policy.Transporter interface for AvsAssessmentsOperationsServerTransport.
func (a *AvsAssessmentsOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvsAssessmentsOperationsClient.BeginCreate":
		resp, err = a.dispatchBeginCreate(req)
	case "AvsAssessmentsOperationsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "AvsAssessmentsOperationsClient.BeginDownloadURL":
		resp, err = a.dispatchBeginDownloadURL(req)
	case "AvsAssessmentsOperationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AvsAssessmentsOperationsClient.NewListByGroupPager":
		resp, err = a.dispatchNewListByGroupPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvsAssessmentsOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := a.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationassessment.AvsAssessment](req)
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
		respr, errRespr := a.srv.BeginCreate(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		a.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		a.beginCreate.remove(req)
	}

	return resp, nil
}

func (a *AvsAssessmentsOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, nil)
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

func (a *AvsAssessmentsOperationsServerTransport) dispatchBeginDownloadURL(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDownloadURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDownloadURL not implemented")}
	}
	beginDownloadURL := a.beginDownloadURL.get(req)
	if beginDownloadURL == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/downloadUrl`
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
		respr, errRespr := a.srv.BeginDownloadURL(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDownloadURL = &respr
		a.beginDownloadURL.add(req, beginDownloadURL)
	}

	resp, err := server.PollerResponderNext(beginDownloadURL, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		a.beginDownloadURL.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDownloadURL) {
		a.beginDownloadURL.remove(req)
	}

	return resp, nil
}

func (a *AvsAssessmentsOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AvsAssessment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AvsAssessmentsOperationsServerTransport) dispatchNewListByGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByGroupPager not implemented")}
	}
	newListByGroupPager := a.newListByGroupPager.get(req)
	if newListByGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments`
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
		resp := a.srv.NewListByGroupPager(resourceGroupNameParam, projectNameParam, groupNameParam, nil)
		newListByGroupPager = &resp
		a.newListByGroupPager.add(req, newListByGroupPager)
		server.PagerResponderInjectNextLinks(newListByGroupPager, req, func(page *armmigrationassessment.AvsAssessmentsOperationsClientListByGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByGroupPager) {
		a.newListByGroupPager.remove(req)
	}
	return resp, nil
}
