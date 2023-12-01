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
	"strconv"
)

// AvsAssessedMachinesOperationsServer is a fake server for instances of the armmigrate.AvsAssessedMachinesOperationsClient type.
type AvsAssessedMachinesOperationsServer struct {
	// Get is the fake for method AvsAssessedMachinesOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, groupName string, assessmentName string, avsAssessedMachineName string, options *armmigrate.AvsAssessedMachinesOperationsClientGetOptions) (resp azfake.Responder[armmigrate.AvsAssessedMachinesOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAvsAssessmentPager is the fake for method AvsAssessedMachinesOperationsClient.NewListByAvsAssessmentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAvsAssessmentPager func(resourceGroupName string, projectName string, groupName string, assessmentName string, options *armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentOptions) (resp azfake.PagerResponder[armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentResponse])
}

// NewAvsAssessedMachinesOperationsServerTransport creates a new instance of AvsAssessedMachinesOperationsServerTransport with the provided implementation.
// The returned AvsAssessedMachinesOperationsServerTransport instance is connected to an instance of armmigrate.AvsAssessedMachinesOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAvsAssessedMachinesOperationsServerTransport(srv *AvsAssessedMachinesOperationsServer) *AvsAssessedMachinesOperationsServerTransport {
	return &AvsAssessedMachinesOperationsServerTransport{
		srv:                         srv,
		newListByAvsAssessmentPager: newTracker[azfake.PagerResponder[armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentResponse]](),
	}
}

// AvsAssessedMachinesOperationsServerTransport connects instances of armmigrate.AvsAssessedMachinesOperationsClient to instances of AvsAssessedMachinesOperationsServer.
// Don't use this type directly, use NewAvsAssessedMachinesOperationsServerTransport instead.
type AvsAssessedMachinesOperationsServerTransport struct {
	srv                         *AvsAssessedMachinesOperationsServer
	newListByAvsAssessmentPager *tracker[azfake.PagerResponder[armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentResponse]]
}

// Do implements the policy.Transporter interface for AvsAssessedMachinesOperationsServerTransport.
func (a *AvsAssessedMachinesOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AvsAssessedMachinesOperationsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AvsAssessedMachinesOperationsClient.NewListByAvsAssessmentPager":
		resp, err = a.dispatchNewListByAvsAssessmentPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AvsAssessedMachinesOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessedMachines/(?P<avsAssessedMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
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
	avsAssessedMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("avsAssessedMachineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, avsAssessedMachineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AvsAssessedMachine, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AvsAssessedMachinesOperationsServerTransport) dispatchNewListByAvsAssessmentPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByAvsAssessmentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAvsAssessmentPager not implemented")}
	}
	newListByAvsAssessmentPager := a.newListByAvsAssessmentPager.get(req)
	if newListByAvsAssessmentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/groups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/avsAssessedMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		pageSizeUnescaped, err := url.QueryUnescape(qp.Get("pageSize"))
		if err != nil {
			return nil, err
		}
		pageSizeParam, err := parseOptional(pageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
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
		groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
		if err != nil {
			return nil, err
		}
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentOptions
		if filterParam != nil || pageSizeParam != nil || continuationTokenParam != nil || totalRecordCountParam != nil {
			options = &armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentOptions{
				Filter:            filterParam,
				PageSize:          pageSizeParam,
				ContinuationToken: continuationTokenParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := a.srv.NewListByAvsAssessmentPager(resourceGroupNameParam, projectNameParam, groupNameParam, assessmentNameParam, options)
		newListByAvsAssessmentPager = &resp
		a.newListByAvsAssessmentPager.add(req, newListByAvsAssessmentPager)
		server.PagerResponderInjectNextLinks(newListByAvsAssessmentPager, req, func(page *armmigrate.AvsAssessedMachinesOperationsClientListByAvsAssessmentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAvsAssessmentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByAvsAssessmentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAvsAssessmentPager) {
		a.newListByAvsAssessmentPager.remove(req)
	}
	return resp, nil
}
