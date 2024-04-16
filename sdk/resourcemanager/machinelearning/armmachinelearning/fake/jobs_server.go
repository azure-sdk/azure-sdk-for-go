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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// JobsServer is a fake server for instances of the armmachinelearning.JobsClient type.
type JobsServer struct {
	// BeginCancel is the fake for method JobsClient.BeginCancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCancel func(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *armmachinelearning.JobsClientBeginCancelOptions) (resp azfake.PollerResponder[armmachinelearning.JobsClientCancelResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method JobsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, id string, body armmachinelearning.JobBase, options *armmachinelearning.JobsClientCreateOrUpdateOptions) (resp azfake.Responder[armmachinelearning.JobsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method JobsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *armmachinelearning.JobsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.JobsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method JobsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, id string, options *armmachinelearning.JobsClientGetOptions) (resp azfake.Responder[armmachinelearning.JobsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method JobsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armmachinelearning.JobsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.JobsClientListResponse])

	// Update is the fake for method JobsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, workspaceName string, id string, body armmachinelearning.PartialJobBasePartialResource, options *armmachinelearning.JobsClientUpdateOptions) (resp azfake.Responder[armmachinelearning.JobsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewJobsServerTransport creates a new instance of JobsServerTransport with the provided implementation.
// The returned JobsServerTransport instance is connected to an instance of armmachinelearning.JobsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewJobsServerTransport(srv *JobsServer) *JobsServerTransport {
	return &JobsServerTransport{
		srv:          srv,
		beginCancel:  newTracker[azfake.PollerResponder[armmachinelearning.JobsClientCancelResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armmachinelearning.JobsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.JobsClientListResponse]](),
	}
}

// JobsServerTransport connects instances of armmachinelearning.JobsClient to instances of JobsServer.
// Don't use this type directly, use NewJobsServerTransport instead.
type JobsServerTransport struct {
	srv          *JobsServer
	beginCancel  *tracker[azfake.PollerResponder[armmachinelearning.JobsClientCancelResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armmachinelearning.JobsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.JobsClientListResponse]]
}

// Do implements the policy.Transporter interface for JobsServerTransport.
func (j *JobsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "JobsClient.BeginCancel":
		resp, err = j.dispatchBeginCancel(req)
	case "JobsClient.CreateOrUpdate":
		resp, err = j.dispatchCreateOrUpdate(req)
	case "JobsClient.BeginDelete":
		resp, err = j.dispatchBeginDelete(req)
	case "JobsClient.Get":
		resp, err = j.dispatchGet(req)
	case "JobsClient.NewListPager":
		resp, err = j.dispatchNewListPager(req)
	case "JobsClient.Update":
		resp, err = j.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginCancel(req *http.Request) (*http.Response, error) {
	if j.srv.BeginCancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancel not implemented")}
	}
	beginCancel := j.beginCancel.get(req)
	if beginCancel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginCancel(req.Context(), resourceGroupNameParam, workspaceNameParam, idParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCancel = &respr
		j.beginCancel.add(req, beginCancel)
	}

	resp, err := server.PollerResponderNext(beginCancel, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		j.beginCancel.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCancel) {
		j.beginCancel.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if j.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.JobBase](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobBase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (j *JobsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if j.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := j.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := j.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, idParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		j.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		j.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		j.beginDelete.remove(req)
	}

	return resp, nil
}

func (j *JobsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if j.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, idParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobBase, req)
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
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs`
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
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		jobTypeUnescaped, err := url.QueryUnescape(qp.Get("jobType"))
		if err != nil {
			return nil, err
		}
		jobTypeParam := getOptional(jobTypeUnescaped)
		tagUnescaped, err := url.QueryUnescape(qp.Get("tag"))
		if err != nil {
			return nil, err
		}
		tagParam := getOptional(tagUnescaped)
		listViewTypeUnescaped, err := url.QueryUnescape(qp.Get("listViewType"))
		if err != nil {
			return nil, err
		}
		listViewTypeParam := getOptional(armmachinelearning.ListViewType(listViewTypeUnescaped))
		propertiesUnescaped, err := url.QueryUnescape(qp.Get("properties"))
		if err != nil {
			return nil, err
		}
		propertiesParam := getOptional(propertiesUnescaped)
		assetNameUnescaped, err := url.QueryUnescape(qp.Get("assetName"))
		if err != nil {
			return nil, err
		}
		assetNameParam := getOptional(assetNameUnescaped)
		scheduledUnescaped, err := url.QueryUnescape(qp.Get("scheduled"))
		if err != nil {
			return nil, err
		}
		scheduledParam, err := parseOptional(scheduledUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		scheduleIDUnescaped, err := url.QueryUnescape(qp.Get("scheduleId"))
		if err != nil {
			return nil, err
		}
		scheduleIDParam := getOptional(scheduleIDUnescaped)
		var options *armmachinelearning.JobsClientListOptions
		if skipParam != nil || jobTypeParam != nil || tagParam != nil || listViewTypeParam != nil || propertiesParam != nil || assetNameParam != nil || scheduledParam != nil || scheduleIDParam != nil {
			options = &armmachinelearning.JobsClientListOptions{
				Skip:         skipParam,
				JobType:      jobTypeParam,
				Tag:          tagParam,
				ListViewType: listViewTypeParam,
				Properties:   propertiesParam,
				AssetName:    assetNameParam,
				Scheduled:    scheduledParam,
				ScheduleID:   scheduleIDParam,
			}
		}
		resp := j.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		j.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.JobsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
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
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<id>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.PartialJobBasePartialResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	idParam, err := url.PathUnescape(matches[regex.SubexpIndex("id")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := j.srv.Update(req.Context(), resourceGroupNameParam, workspaceNameParam, idParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).JobBase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
