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

// AssessedWebApplicationOperationsServer is a fake server for instances of the armmigrationassessment.AssessedWebApplicationOperationsClient type.
type AssessedWebApplicationOperationsServer struct {
	// Get is the fake for method AssessedWebApplicationOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, assessedWorkload string, options *armmigrationassessment.AssessedWebApplicationOperationsClientGetOptions) (resp azfake.Responder[armmigrationassessment.AssessedWebApplicationOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAksAssessmentPager is the fake for method AssessedWebApplicationOperationsClient.NewListByAksAssessmentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAksAssessmentPager func(resourceGroupName string, projectName string, assessmentName string, options *armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentOptions) (resp azfake.PagerResponder[armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentResponse])
}

// NewAssessedWebApplicationOperationsServerTransport creates a new instance of AssessedWebApplicationOperationsServerTransport with the provided implementation.
// The returned AssessedWebApplicationOperationsServerTransport instance is connected to an instance of armmigrationassessment.AssessedWebApplicationOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssessedWebApplicationOperationsServerTransport(srv *AssessedWebApplicationOperationsServer) *AssessedWebApplicationOperationsServerTransport {
	return &AssessedWebApplicationOperationsServerTransport{
		srv:                         srv,
		newListByAksAssessmentPager: newTracker[azfake.PagerResponder[armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentResponse]](),
	}
}

// AssessedWebApplicationOperationsServerTransport connects instances of armmigrationassessment.AssessedWebApplicationOperationsClient to instances of AssessedWebApplicationOperationsServer.
// Don't use this type directly, use NewAssessedWebApplicationOperationsServerTransport instead.
type AssessedWebApplicationOperationsServerTransport struct {
	srv                         *AssessedWebApplicationOperationsServer
	newListByAksAssessmentPager *tracker[azfake.PagerResponder[armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentResponse]]
}

// Do implements the policy.Transporter interface for AssessedWebApplicationOperationsServerTransport.
func (a *AssessedWebApplicationOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AssessedWebApplicationOperationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AssessedWebApplicationOperationsClient.NewListByAksAssessmentPager":
		resp, err = a.dispatchNewListByAksAssessmentPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AssessedWebApplicationOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessedWebApps/(?P<assessedWorkload>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
	if err != nil {
		return nil, err
	}
	assessedWorkloadParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessedWorkload")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, assessmentNameParam, assessedWorkloadParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssessedWebApplication, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssessedWebApplicationOperationsServerTransport) dispatchNewListByAksAssessmentPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByAksAssessmentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAksAssessmentPager not implemented")}
	}
	newListByAksAssessmentPager := a.newListByAksAssessmentPager.get(req)
	if newListByAksAssessmentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/aksAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessedWebApps`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentOptions
		if continuationTokenParam != nil || topParam != nil || filterParam != nil || totalRecordCountParam != nil {
			options = &armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentOptions{
				ContinuationToken: continuationTokenParam,
				Top:               topParam,
				Filter:            filterParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := a.srv.NewListByAksAssessmentPager(resourceGroupNameParam, projectNameParam, assessmentNameParam, options)
		newListByAksAssessmentPager = &resp
		a.newListByAksAssessmentPager.add(req, newListByAksAssessmentPager)
		server.PagerResponderInjectNextLinks(newListByAksAssessmentPager, req, func(page *armmigrationassessment.AssessedWebApplicationOperationsClientListByAksAssessmentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAksAssessmentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByAksAssessmentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAksAssessmentPager) {
		a.newListByAksAssessmentPager.remove(req)
	}
	return resp, nil
}
