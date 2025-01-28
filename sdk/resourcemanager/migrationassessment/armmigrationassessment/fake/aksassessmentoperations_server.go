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
	"strconv"
)

// AksAssessmentOperationsServer is a fake server for instances of the armmigrationassessment.AksAssessmentOperationsClient type.
type AksAssessmentOperationsServer struct {
	// BeginCreate is the fake for method AksAssessmentOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, resource armmigrationassessment.AKSAssessment, options *armmigrationassessment.AksAssessmentOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrationassessment.AksAssessmentOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AksAssessmentOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, options *armmigrationassessment.AksAssessmentOperationsClientDeleteOptions) (resp azfake.Responder[armmigrationassessment.AksAssessmentOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDownloadURL is the fake for method AksAssessmentOperationsClient.BeginDownloadURL
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDownloadURL func(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, body any, options *armmigrationassessment.AksAssessmentOperationsClientBeginDownloadURLOptions) (resp azfake.PollerResponder[armmigrationassessment.AksAssessmentOperationsClientDownloadURLResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AksAssessmentOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, options *armmigrationassessment.AksAssessmentOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.AksAssessmentOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAssessmentProjectPager is the fake for method AksAssessmentOperationsClient.NewListByAssessmentProjectPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAssessmentProjectPager func(resourceGroupName string, projectName string, options *armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectOptions) (resp azfake.PagerResponder[armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectResponse])
}

// NewAksAssessmentOperationsServerTransport creates a new instance of AksAssessmentOperationsServerTransport with the provided implementation.
// The returned AksAssessmentOperationsServerTransport instance is connected to an instance of armmigrationassessment.AksAssessmentOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAksAssessmentOperationsServerTransport(srv *AksAssessmentOperationsServer) *AksAssessmentOperationsServerTransport {
	return &AksAssessmentOperationsServerTransport{
		srv:                             srv,
		beginCreate:                     newTracker[azfake.PollerResponder[armmigrationassessment.AksAssessmentOperationsClientCreateResponse]](),
		beginDownloadURL:                newTracker[azfake.PollerResponder[armmigrationassessment.AksAssessmentOperationsClientDownloadURLResponse]](),
		newListByAssessmentProjectPager: newTracker[azfake.PagerResponder[armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectResponse]](),
	}
}

// AksAssessmentOperationsServerTransport connects instances of armmigrationassessment.AksAssessmentOperationsClient to instances of AksAssessmentOperationsServer.
// Don't use this type directly, use NewAksAssessmentOperationsServerTransport instead.
type AksAssessmentOperationsServerTransport struct {
	srv                             *AksAssessmentOperationsServer
	beginCreate                     *tracker[azfake.PollerResponder[armmigrationassessment.AksAssessmentOperationsClientCreateResponse]]
	beginDownloadURL                *tracker[azfake.PollerResponder[armmigrationassessment.AksAssessmentOperationsClientDownloadURLResponse]]
	newListByAssessmentProjectPager *tracker[azfake.PagerResponder[armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectResponse]]
}

// Do implements the policy.Transporter interface for AksAssessmentOperationsServerTransport.
func (a *AksAssessmentOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AksAssessmentOperationsClient.BeginCreate":
		resp, err = a.dispatchBeginCreate(req)
	case "AksAssessmentOperationsClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "AksAssessmentOperationsClient.BeginDownloadURL":
		resp, err = a.dispatchBeginDownloadURL(req)
	case "AksAssessmentOperationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AksAssessmentOperationsClient.NewListByAssessmentProjectPager":
		resp, err = a.dispatchNewListByAssessmentProjectPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AksAssessmentOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := a.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationassessment.AKSAssessment](req)
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
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreate(req.Context(), resourceGroupNameParam, projectNameParam, assessmentNameParam, body, nil)
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

func (a *AksAssessmentOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, projectNameParam, assessmentNameParam, nil)
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

func (a *AksAssessmentOperationsServerTransport) dispatchBeginDownloadURL(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDownloadURL == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDownloadURL not implemented")}
	}
	beginDownloadURL := a.beginDownloadURL.get(req)
	if beginDownloadURL == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/downloadUrl`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDownloadURL(req.Context(), resourceGroupNameParam, projectNameParam, assessmentNameParam, body, nil)
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

func (a *AksAssessmentOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, assessmentNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AKSAssessment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AksAssessmentOperationsServerTransport) dispatchNewListByAssessmentProjectPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByAssessmentProjectPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAssessmentProjectPager not implemented")}
	}
	newListByAssessmentProjectPager := a.newListByAssessmentProjectPager.get(req)
	if newListByAssessmentProjectPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
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
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		totalRecordCountUnescaped, err := url.QueryUnescape(qp.Get("totalRecordCount"))
		if err != nil {
			return nil, err
		}
		totalRecordCountParam, err := parseOptional(totalRecordCountUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectOptions
		if continuationTokenParam != nil || topParam != nil || filterParam != nil || totalRecordCountParam != nil {
			options = &armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectOptions{
				ContinuationToken: continuationTokenParam,
				Top:               topParam,
				Filter:            filterParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := a.srv.NewListByAssessmentProjectPager(resourceGroupNameParam, projectNameParam, options)
		newListByAssessmentProjectPager = &resp
		a.newListByAssessmentProjectPager.add(req, newListByAssessmentProjectPager)
		server.PagerResponderInjectNextLinks(newListByAssessmentProjectPager, req, func(page *armmigrationassessment.AksAssessmentOperationsClientListByAssessmentProjectResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAssessmentProjectPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByAssessmentProjectPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAssessmentProjectPager) {
		a.newListByAssessmentProjectPager.remove(req)
	}
	return resp, nil
}
