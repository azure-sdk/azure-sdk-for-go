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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"net/http"
	"net/url"
	"regexp"
)

// JobServer is a fake server for instances of the armautomation.JobClient type.
type JobServer struct {
	// Create is the fake for method JobClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, parameters armautomation.JobCreateParameters, options *armautomation.JobClientCreateOptions) (resp azfake.Responder[armautomation.JobClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *armautomation.JobClientGetOptions) (resp azfake.Responder[armautomation.JobClientGetResponse], errResp azfake.ErrorResponder)

	// GetOutput is the fake for method JobClient.GetOutput
	// HTTP status codes to indicate success: http.StatusOK
	GetOutput func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *armautomation.JobClientGetOutputOptions) (resp azfake.Responder[armautomation.JobClientGetOutputResponse], errResp azfake.ErrorResponder)

	// GetRunbookContent is the fake for method JobClient.GetRunbookContent
	// HTTP status codes to indicate success: http.StatusOK
	GetRunbookContent func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *armautomation.JobClientGetRunbookContentOptions) (resp azfake.Responder[armautomation.JobClientGetRunbookContentResponse], errResp azfake.ErrorResponder)

	// NewListByAutomationAccountPager is the fake for method JobClient.NewListByAutomationAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAutomationAccountPager func(resourceGroupName string, automationAccountName string, options *armautomation.JobClientListByAutomationAccountOptions) (resp azfake.PagerResponder[armautomation.JobClientListByAutomationAccountResponse])

	// Resume is the fake for method JobClient.Resume
	// HTTP status codes to indicate success: http.StatusOK
	Resume func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *armautomation.JobClientResumeOptions) (resp azfake.Responder[armautomation.JobClientResumeResponse], errResp azfake.ErrorResponder)

	// Stop is the fake for method JobClient.Stop
	// HTTP status codes to indicate success: http.StatusOK
	Stop func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *armautomation.JobClientStopOptions) (resp azfake.Responder[armautomation.JobClientStopResponse], errResp azfake.ErrorResponder)

	// Suspend is the fake for method JobClient.Suspend
	// HTTP status codes to indicate success: http.StatusOK
	Suspend func(ctx context.Context, resourceGroupName string, automationAccountName string, jobName string, options *armautomation.JobClientSuspendOptions) (resp azfake.Responder[armautomation.JobClientSuspendResponse], errResp azfake.ErrorResponder)
}

// NewJobServerTransport creates a new instance of JobServerTransport with the provided implementation.
// The returned JobServerTransport instance is connected to an instance of armautomation.JobClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobServerTransport(srv *JobServer) *JobServerTransport {
	return &JobServerTransport{
		srv:                             srv,
		newListByAutomationAccountPager: newTracker[azfake.PagerResponder[armautomation.JobClientListByAutomationAccountResponse]](),
	}
}

// JobServerTransport connects instances of armautomation.JobClient to instances of JobServer.
// Don't use this type directly, use NewJobServerTransport instead.
type JobServerTransport struct {
	srv                             *JobServer
	newListByAutomationAccountPager *tracker[azfake.PagerResponder[armautomation.JobClientListByAutomationAccountResponse]]
}

// Do implements the policy.Transporter interface for JobServerTransport.
func (j *JobServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JobClient.Create":
		resp, err = j.dispatchCreate(req)
	case "JobClient.Get":
		resp, err = j.dispatchGet(req)
	case "JobClient.GetOutput":
		resp, err = j.dispatchGetOutput(req)
	case "JobClient.GetRunbookContent":
		resp, err = j.dispatchGetRunbookContent(req)
	case "JobClient.NewListByAutomationAccountPager":
		resp, err = j.dispatchNewListByAutomationAccountPager(req)
	case "JobClient.Resume":
		resp, err = j.dispatchResume(req)
	case "JobClient.Stop":
		resp, err = j.dispatchStop(req)
	case "JobClient.Suspend":
		resp, err = j.dispatchSuspend(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JobServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if j.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armautomation.JobCreateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientCreateOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientCreateOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.Create(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Job, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientGetOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientGetOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Job, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchGetOutput(req *http.Request) (*http.Response, error) {
	if j.srv.GetOutput == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOutput not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/output`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientGetOutputOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientGetOutputOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.GetOutput(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchGetRunbookContent(req *http.Request) (*http.Response, error) {
	if j.srv.GetRunbookContent == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRunbookContent not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runbookContent`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientGetRunbookContentOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientGetRunbookContentOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.GetRunbookContent(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsText(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchNewListByAutomationAccountPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByAutomationAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAutomationAccountPager not implemented")}
	}
	newListByAutomationAccountPager := j.newListByAutomationAccountPager.get(req)
	if newListByAutomationAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs`
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
		automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
		var options *armautomation.JobClientListByAutomationAccountOptions
		if filterParam != nil || clientRequestIDParam != nil {
			options = &armautomation.JobClientListByAutomationAccountOptions{
				Filter:          filterParam,
				ClientRequestID: clientRequestIDParam,
			}
		}
		resp := j.srv.NewListByAutomationAccountPager(resourceGroupNameParam, automationAccountNameParam, options)
		newListByAutomationAccountPager = &resp
		j.newListByAutomationAccountPager.add(req, newListByAutomationAccountPager)
		server.PagerResponderInjectNextLinks(newListByAutomationAccountPager, req, func(page *armautomation.JobClientListByAutomationAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAutomationAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByAutomationAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAutomationAccountPager) {
		j.newListByAutomationAccountPager.remove(req)
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchResume(req *http.Request) (*http.Response, error) {
	if j.srv.Resume == nil {
		return nil, &nonRetriableError{errors.New("fake for method Resume not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resume`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientResumeOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientResumeOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.Resume(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchStop(req *http.Request) (*http.Response, error) {
	if j.srv.Stop == nil {
		return nil, &nonRetriableError{errors.New("fake for method Stop not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientStopOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientStopOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.Stop(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobServerTransport) dispatchSuspend(req *http.Request) (*http.Response, error) {
	if j.srv.Suspend == nil {
		return nil, &nonRetriableError{errors.New("fake for method Suspend not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/suspend`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	clientRequestIDParam := getOptional(getHeaderValue(req.Header, "clientRequestId"))
	var options *armautomation.JobClientSuspendOptions
	if clientRequestIDParam != nil {
		options = &armautomation.JobClientSuspendOptions{
			ClientRequestID: clientRequestIDParam,
		}
	}
	respr, errRespr := j.srv.Suspend(req.Context(), resourceGroupNameParam, automationAccountNameParam, jobNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
