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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// JobVersionsServer is a fake server for instances of the armsql.JobVersionsClient type.
type JobVersionsServer struct {
	// Get is the fake for method JobVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, jobName string, jobVersion int32, options *armsql.JobVersionsClientGetOptions) (resp azfake.Responder[armsql.JobVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByJobPager is the fake for method JobVersionsClient.NewListByJobPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByJobPager func(resourceGroupName string, serverName string, jobAgentName string, jobName string, options *armsql.JobVersionsClientListByJobOptions) (resp azfake.PagerResponder[armsql.JobVersionsClientListByJobResponse])
}

// NewJobVersionsServerTransport creates a new instance of JobVersionsServerTransport with the provided implementation.
// The returned JobVersionsServerTransport instance is connected to an instance of armsql.JobVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobVersionsServerTransport(srv *JobVersionsServer) *JobVersionsServerTransport {
	return &JobVersionsServerTransport{
		srv:               srv,
		newListByJobPager: newTracker[azfake.PagerResponder[armsql.JobVersionsClientListByJobResponse]](),
	}
}

// JobVersionsServerTransport connects instances of armsql.JobVersionsClient to instances of JobVersionsServer.
// Don't use this type directly, use NewJobVersionsServerTransport instead.
type JobVersionsServerTransport struct {
	srv               *JobVersionsServer
	newListByJobPager *tracker[azfake.PagerResponder[armsql.JobVersionsClientListByJobResponse]]
}

// Do implements the policy.Transporter interface for JobVersionsServerTransport.
func (j *JobVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return j.dispatchToMethodFake(req, method)
}

func (j *JobVersionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if jobVersionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = jobVersionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "JobVersionsClient.Get":
				res.resp, res.err = j.dispatchGet(req)
			case "JobVersionsClient.NewListByJobPager":
				res.resp, res.err = j.dispatchNewListByJobPager(req)
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

func (j *JobVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<jobVersion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
	if err != nil {
		return nil, err
	}
	jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
	if err != nil {
		return nil, err
	}
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	jobVersionUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("jobVersion")])
	if err != nil {
		return nil, err
	}
	jobVersionParam, err := parseWithCast(jobVersionUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, jobVersionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobVersionsServerTransport) dispatchNewListByJobPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByJobPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByJobPager not implemented")}
	}
	newListByJobPager := j.newListByJobPager.get(req)
	if newListByJobPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serverNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serverName")])
		if err != nil {
			return nil, err
		}
		jobAgentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobAgentName")])
		if err != nil {
			return nil, err
		}
		jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
		if err != nil {
			return nil, err
		}
		resp := j.srv.NewListByJobPager(resourceGroupNameParam, serverNameParam, jobAgentNameParam, jobNameParam, nil)
		newListByJobPager = &resp
		j.newListByJobPager.add(req, newListByJobPager)
		server.PagerResponderInjectNextLinks(newListByJobPager, req, func(page *armsql.JobVersionsClientListByJobResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByJobPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByJobPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByJobPager) {
		j.newListByJobPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to JobVersionsServerTransport
var jobVersionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
