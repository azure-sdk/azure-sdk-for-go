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
)

// JobTargetGroupsServer is a fake server for instances of the armsql.JobTargetGroupsClient type.
type JobTargetGroupsServer struct {
	// CreateOrUpdate is the fake for method JobTargetGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, parameters armsql.JobTargetGroup, options *armsql.JobTargetGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armsql.JobTargetGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method JobTargetGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, options *armsql.JobTargetGroupsClientDeleteOptions) (resp azfake.Responder[armsql.JobTargetGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobTargetGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, jobAgentName string, targetGroupName string, options *armsql.JobTargetGroupsClientGetOptions) (resp azfake.Responder[armsql.JobTargetGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAgentPager is the fake for method JobTargetGroupsClient.NewListByAgentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAgentPager func(resourceGroupName string, serverName string, jobAgentName string, options *armsql.JobTargetGroupsClientListByAgentOptions) (resp azfake.PagerResponder[armsql.JobTargetGroupsClientListByAgentResponse])
}

// NewJobTargetGroupsServerTransport creates a new instance of JobTargetGroupsServerTransport with the provided implementation.
// The returned JobTargetGroupsServerTransport instance is connected to an instance of armsql.JobTargetGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobTargetGroupsServerTransport(srv *JobTargetGroupsServer) *JobTargetGroupsServerTransport {
	return &JobTargetGroupsServerTransport{
		srv:                 srv,
		newListByAgentPager: newTracker[azfake.PagerResponder[armsql.JobTargetGroupsClientListByAgentResponse]](),
	}
}

// JobTargetGroupsServerTransport connects instances of armsql.JobTargetGroupsClient to instances of JobTargetGroupsServer.
// Don't use this type directly, use NewJobTargetGroupsServerTransport instead.
type JobTargetGroupsServerTransport struct {
	srv                 *JobTargetGroupsServer
	newListByAgentPager *tracker[azfake.PagerResponder[armsql.JobTargetGroupsClientListByAgentResponse]]
}

// Do implements the policy.Transporter interface for JobTargetGroupsServerTransport.
func (j *JobTargetGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return j.dispatchToMethodFake(req, method)
}

func (j *JobTargetGroupsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if jobTargetGroupsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = jobTargetGroupsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "JobTargetGroupsClient.CreateOrUpdate":
				res.resp, res.err = j.dispatchCreateOrUpdate(req)
			case "JobTargetGroupsClient.Delete":
				res.resp, res.err = j.dispatchDelete(req)
			case "JobTargetGroupsClient.Get":
				res.resp, res.err = j.dispatchGet(req)
			case "JobTargetGroupsClient.NewListByAgentPager":
				res.resp, res.err = j.dispatchNewListByAgentPager(req)
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

func (j *JobTargetGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetGroups/(?P<targetGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsql.JobTargetGroup](req)
	if err != nil {
		return nil, err
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
	targetGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, targetGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobTargetGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobTargetGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if j.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetGroups/(?P<targetGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	targetGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Delete(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, targetGroupNameParam, nil)
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

func (j *JobTargetGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetGroups/(?P<targetGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	targetGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, jobAgentNameParam, targetGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobTargetGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobTargetGroupsServerTransport) dispatchNewListByAgentPager(req *http.Request) (*http.Response, error) {
	if j.srv.NewListByAgentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAgentPager not implemented")}
	}
	newListByAgentPager := j.newListByAgentPager.get(req)
	if newListByAgentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobAgents/(?P<jobAgentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		resp := j.srv.NewListByAgentPager(resourceGroupNameParam, serverNameParam, jobAgentNameParam, nil)
		newListByAgentPager = &resp
		j.newListByAgentPager.add(req, newListByAgentPager)
		server.PagerResponderInjectNextLinks(newListByAgentPager, req, func(page *armsql.JobTargetGroupsClientListByAgentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAgentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		j.newListByAgentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAgentPager) {
		j.newListByAgentPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to JobTargetGroupsServerTransport
var jobTargetGroupsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
