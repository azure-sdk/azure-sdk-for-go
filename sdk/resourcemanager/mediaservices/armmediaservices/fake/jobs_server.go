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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
	"net/http"
	"net/url"
	"regexp"
)

// JobsServer is a fake server for instances of the armmediaservices.JobsClient type.
type JobsServer struct {
	// CancelJob is the fake for method JobsClient.CancelJob
	// HTTP status codes to indicate success: http.StatusOK
	CancelJob func(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string, options *armmediaservices.JobsClientCancelJobOptions) (resp azfake.Responder[armmediaservices.JobsClientCancelJobResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method JobsClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string, parameters armmediaservices.Job, options *armmediaservices.JobsClientCreateOptions) (resp azfake.Responder[armmediaservices.JobsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method JobsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string, options *armmediaservices.JobsClientDeleteOptions) (resp azfake.Responder[armmediaservices.JobsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string, options *armmediaservices.JobsClientGetOptions) (resp azfake.Responder[armmediaservices.JobsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method JobsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, accountName string, transformName string, options *armmediaservices.JobsClientListOptions) (resp azfake.PagerResponder[armmediaservices.JobsClientListResponse])

	// Update is the fake for method JobsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, accountName string, transformName string, jobName string, parameters armmediaservices.Job, options *armmediaservices.JobsClientUpdateOptions) (resp azfake.Responder[armmediaservices.JobsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewJobsServerTransport creates a new instance of JobsServerTransport with the provided implementation.
// The returned JobsServerTransport instance is connected to an instance of armmediaservices.JobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobsServerTransport(srv *JobsServer) *JobsServerTransport {
	return &JobsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmediaservices.JobsClientListResponse]](),
	}
}

// JobsServerTransport connects instances of armmediaservices.JobsClient to instances of JobsServer.
// Don't use this type directly, use NewJobsServerTransport instead.
type JobsServerTransport struct {
	srv          *JobsServer
	newListPager *tracker[azfake.PagerResponder[armmediaservices.JobsClientListResponse]]
}

// Do implements the policy.Transporter interface for JobsServerTransport.
func (j *JobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return j.dispatchToMethodFake(req, method)
}

func (j *JobsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if jobsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = jobsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "JobsClient.CancelJob":
				res.resp, res.err = j.dispatchCancelJob(req)
			case "JobsClient.Create":
				res.resp, res.err = j.dispatchCreate(req)
			case "JobsClient.Delete":
				res.resp, res.err = j.dispatchDelete(req)
			case "JobsClient.Get":
				res.resp, res.err = j.dispatchGet(req)
			case "JobsClient.NewListPager":
				res.resp, res.err = j.dispatchNewListPager(req)
			case "JobsClient.Update":
				res.resp, res.err = j.dispatchUpdate(req)
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

func (j *JobsServerTransport) dispatchCancelJob(req *http.Request) (*http.Response, error) {
	if j.srv.CancelJob == nil {
		return nil, &nonRetriableError{errors.New("fake for method CancelJob not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transforms/(?P<transformName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancelJob`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	transformNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.CancelJob(req.Context(), resourceGroupNameParam, accountNameParam, transformNameParam, jobNameParam, nil)
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

func (j *JobsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if j.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transforms/(?P<transformName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmediaservices.Job](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	transformNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, transformNameParam, jobNameParam, body, nil)
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

func (j *JobsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if j.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transforms/(?P<transformName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	transformNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, transformNameParam, jobNameParam, nil)
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

func (j *JobsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transforms/(?P<transformName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	transformNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, transformNameParam, jobNameParam, nil)
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

func (j *JobsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := j.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transforms/(?P<transformName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		transformNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armmediaservices.JobsClientListOptions
		if filterParam != nil || orderbyParam != nil {
			options = &armmediaservices.JobsClientListOptions{
				Filter:  filterParam,
				Orderby: orderbyParam,
			}
		}
		resp := j.srv.NewListPager(resourceGroupNameParam, accountNameParam, transformNameParam, options)
		newListPager = &resp
		j.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmediaservices.JobsClientListResponse, createLink func() string) {
			page.ODataNextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		j.newListPager.remove(req)
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Media/mediaServices/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/transforms/(?P<transformName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmediaservices.Job](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	transformNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("transformName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Update(req.Context(), resourceGroupNameParam, accountNameParam, transformNameParam, jobNameParam, body, nil)
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

// set this to conditionally intercept incoming requests to JobsServerTransport
var jobsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
