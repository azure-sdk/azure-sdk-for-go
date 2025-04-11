// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

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

// AssessedMachinesV2OperationsServer is a fake server for instances of the armmigrate.AssessedMachinesV2OperationsClient type.
type AssessedMachinesV2OperationsServer struct {
	// Get is the fake for method AssessedMachinesV2OperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, projectName string, assessmentName string, assessedMachineName string, options *armmigrate.AssessedMachinesV2OperationsClientGetOptions) (resp azfake.Responder[armmigrate.AssessedMachinesV2OperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByParentPager is the fake for method AssessedMachinesV2OperationsClient.NewListByParentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByParentPager func(resourceGroupName string, projectName string, assessmentName string, options *armmigrate.AssessedMachinesV2OperationsClientListByParentOptions) (resp azfake.PagerResponder[armmigrate.AssessedMachinesV2OperationsClientListByParentResponse])
}

// NewAssessedMachinesV2OperationsServerTransport creates a new instance of AssessedMachinesV2OperationsServerTransport with the provided implementation.
// The returned AssessedMachinesV2OperationsServerTransport instance is connected to an instance of armmigrate.AssessedMachinesV2OperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAssessedMachinesV2OperationsServerTransport(srv *AssessedMachinesV2OperationsServer) *AssessedMachinesV2OperationsServerTransport {
	return &AssessedMachinesV2OperationsServerTransport{
		srv:                  srv,
		newListByParentPager: newTracker[azfake.PagerResponder[armmigrate.AssessedMachinesV2OperationsClientListByParentResponse]](),
	}
}

// AssessedMachinesV2OperationsServerTransport connects instances of armmigrate.AssessedMachinesV2OperationsClient to instances of AssessedMachinesV2OperationsServer.
// Don't use this type directly, use NewAssessedMachinesV2OperationsServerTransport instead.
type AssessedMachinesV2OperationsServerTransport struct {
	srv                  *AssessedMachinesV2OperationsServer
	newListByParentPager *tracker[azfake.PagerResponder[armmigrate.AssessedMachinesV2OperationsClientListByParentResponse]]
}

// Do implements the policy.Transporter interface for AssessedMachinesV2OperationsServerTransport.
func (a *AssessedMachinesV2OperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AssessedMachinesV2OperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if assessedMachinesV2OperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = assessedMachinesV2OperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AssessedMachinesV2OperationsClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "AssessedMachinesV2OperationsClient.NewListByParentPager":
				res.resp, res.err = a.dispatchNewListByParentPager(req)
			default:
				res.err = fmt.Errorf("unhandled API %s", method)
			}

		}
		select {
		case resultChan <- res:
		case <-req.Context().Done():
		}
	}()

	select {
	case <-req.Context().Done():
		return nil, req.Context().Err()
	case res := <-resultChan:
		return res.resp, res.err
	}
}

func (a *AssessedMachinesV2OperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessedMachines/(?P<assessedMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	assessedMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessedMachineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, projectNameParam, assessmentNameParam, assessedMachineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AssessedMachineV2, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AssessedMachinesV2OperationsServerTransport) dispatchNewListByParentPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByParentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByParentPager not implemented")}
	}
	newListByParentPager := a.newListByParentPager.get(req)
	if newListByParentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/assessmentProjects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessments/(?P<assessmentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/assessedMachines`
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
		assessmentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("assessmentName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrate.AssessedMachinesV2OperationsClientListByParentOptions
		if filterParam != nil || pageSizeParam != nil || continuationTokenParam != nil || totalRecordCountParam != nil {
			options = &armmigrate.AssessedMachinesV2OperationsClientListByParentOptions{
				Filter:            filterParam,
				PageSize:          pageSizeParam,
				ContinuationToken: continuationTokenParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := a.srv.NewListByParentPager(resourceGroupNameParam, projectNameParam, assessmentNameParam, options)
		newListByParentPager = &resp
		a.newListByParentPager.add(req, newListByParentPager)
		server.PagerResponderInjectNextLinks(newListByParentPager, req, func(page *armmigrate.AssessedMachinesV2OperationsClientListByParentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByParentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByParentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByParentPager) {
		a.newListByParentPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AssessedMachinesV2OperationsServerTransport
var assessedMachinesV2OperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
