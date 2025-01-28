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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ServiceTasksServer is a fake server for instances of the armdatamigration.ServiceTasksClient type.
type ServiceTasksServer struct {
	// Cancel is the fake for method ServiceTasksClient.Cancel
	// HTTP status codes to indicate success: http.StatusOK
	Cancel func(ctx context.Context, groupName string, serviceName string, taskName string, options *armdatamigration.ServiceTasksClientCancelOptions) (resp azfake.Responder[armdatamigration.ServiceTasksClientCancelResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method ServiceTasksClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, groupName string, serviceName string, taskName string, parameters armdatamigration.ProjectTask, options *armdatamigration.ServiceTasksClientCreateOrUpdateOptions) (resp azfake.Responder[armdatamigration.ServiceTasksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ServiceTasksClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, groupName string, serviceName string, taskName string, options *armdatamigration.ServiceTasksClientDeleteOptions) (resp azfake.Responder[armdatamigration.ServiceTasksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServiceTasksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, groupName string, serviceName string, taskName string, options *armdatamigration.ServiceTasksClientGetOptions) (resp azfake.Responder[armdatamigration.ServiceTasksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServiceTasksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(groupName string, serviceName string, options *armdatamigration.ServiceTasksClientListOptions) (resp azfake.PagerResponder[armdatamigration.ServiceTasksClientListResponse])

	// Update is the fake for method ServiceTasksClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, groupName string, serviceName string, taskName string, parameters armdatamigration.ProjectTask, options *armdatamigration.ServiceTasksClientUpdateOptions) (resp azfake.Responder[armdatamigration.ServiceTasksClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServiceTasksServerTransport creates a new instance of ServiceTasksServerTransport with the provided implementation.
// The returned ServiceTasksServerTransport instance is connected to an instance of armdatamigration.ServiceTasksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServiceTasksServerTransport(srv *ServiceTasksServer) *ServiceTasksServerTransport {
	return &ServiceTasksServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdatamigration.ServiceTasksClientListResponse]](),
	}
}

// ServiceTasksServerTransport connects instances of armdatamigration.ServiceTasksClient to instances of ServiceTasksServer.
// Don't use this type directly, use NewServiceTasksServerTransport instead.
type ServiceTasksServerTransport struct {
	srv          *ServiceTasksServer
	newListPager *tracker[azfake.PagerResponder[armdatamigration.ServiceTasksClientListResponse]]
}

// Do implements the policy.Transporter interface for ServiceTasksServerTransport.
func (s *ServiceTasksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServiceTasksClient.Cancel":
		resp, err = s.dispatchCancel(req)
	case "ServiceTasksClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "ServiceTasksClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "ServiceTasksClient.Get":
		resp, err = s.dispatchGet(req)
	case "ServiceTasksClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "ServiceTasksClient.Update":
		resp, err = s.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServiceTasksServerTransport) dispatchCancel(req *http.Request) (*http.Response, error) {
	if s.srv.Cancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method Cancel not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	taskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Cancel(req.Context(), groupNameParam, serviceNameParam, taskNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProjectTask, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceTasksServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatamigration.ProjectTask](req)
	if err != nil {
		return nil, err
	}
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	taskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), groupNameParam, serviceNameParam, taskNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProjectTask, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceTasksServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	taskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	deleteRunningTasksUnescaped, err := url.QueryUnescape(qp.Get("deleteRunningTasks"))
	if err != nil {
		return nil, err
	}
	deleteRunningTasksParam, err := parseOptional(deleteRunningTasksUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdatamigration.ServiceTasksClientDeleteOptions
	if deleteRunningTasksParam != nil {
		options = &armdatamigration.ServiceTasksClientDeleteOptions{
			DeleteRunningTasks: deleteRunningTasksParam,
		}
	}
	respr, errRespr := s.srv.Delete(req.Context(), groupNameParam, serviceNameParam, taskNameParam, options)
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

func (s *ServiceTasksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	taskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armdatamigration.ServiceTasksClientGetOptions
	if expandParam != nil {
		options = &armdatamigration.ServiceTasksClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := s.srv.Get(req.Context(), groupNameParam, serviceNameParam, taskNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProjectTask, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServiceTasksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTasks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		taskTypeUnescaped, err := url.QueryUnescape(qp.Get("taskType"))
		if err != nil {
			return nil, err
		}
		taskTypeParam := getOptional(taskTypeUnescaped)
		var options *armdatamigration.ServiceTasksClientListOptions
		if taskTypeParam != nil {
			options = &armdatamigration.ServiceTasksClientListOptions{
				TaskType: taskTypeParam,
			}
		}
		resp := s.srv.NewListPager(groupNameParam, serviceNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdatamigration.ServiceTasksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServiceTasksServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<groupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/services/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serviceTasks/(?P<taskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatamigration.ProjectTask](req)
	if err != nil {
		return nil, err
	}
	groupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("groupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	taskNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("taskName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Update(req.Context(), groupNameParam, serviceNameParam, taskNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProjectTask, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
