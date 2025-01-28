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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationmodernization/armmigrationmodernization"
	"net/http"
	"net/url"
	"regexp"
)

// WorkloadInstanceServer is a fake server for instances of the armmigrationmodernization.WorkloadInstanceClient type.
type WorkloadInstanceServer struct {
	// Create is the fake for method WorkloadInstanceClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, body armmigrationmodernization.WorkloadInstanceModel, options *armmigrationmodernization.WorkloadInstanceClientCreateOptions) (resp azfake.Responder[armmigrationmodernization.WorkloadInstanceClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method WorkloadInstanceClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, options *armmigrationmodernization.WorkloadInstanceClientBeginDeleteOptions) (resp azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkloadInstanceClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, options *armmigrationmodernization.WorkloadInstanceClientGetOptions) (resp azfake.Responder[armmigrationmodernization.WorkloadInstanceClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkloadInstanceClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(subscriptionID string, resourceGroupName string, modernizeProjectName string, options *armmigrationmodernization.WorkloadInstanceClientListOptions) (resp azfake.PagerResponder[armmigrationmodernization.WorkloadInstanceClientListResponse])

	// BeginMigrateComplete is the fake for method WorkloadInstanceClient.BeginMigrateComplete
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginMigrateComplete func(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, options *armmigrationmodernization.WorkloadInstanceClientBeginMigrateCompleteOptions) (resp azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientMigrateCompleteResponse], errResp azfake.ErrorResponder)

	// BeginStopReplicate is the fake for method WorkloadInstanceClient.BeginStopReplicate
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginStopReplicate func(ctx context.Context, subscriptionID string, resourceGroupName string, modernizeProjectName string, workloadInstanceName string, options *armmigrationmodernization.WorkloadInstanceClientBeginStopReplicateOptions) (resp azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientStopReplicateResponse], errResp azfake.ErrorResponder)
}

// NewWorkloadInstanceServerTransport creates a new instance of WorkloadInstanceServerTransport with the provided implementation.
// The returned WorkloadInstanceServerTransport instance is connected to an instance of armmigrationmodernization.WorkloadInstanceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkloadInstanceServerTransport(srv *WorkloadInstanceServer) *WorkloadInstanceServerTransport {
	return &WorkloadInstanceServerTransport{
		srv:                  srv,
		beginDelete:          newTracker[azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientDeleteResponse]](),
		newListPager:         newTracker[azfake.PagerResponder[armmigrationmodernization.WorkloadInstanceClientListResponse]](),
		beginMigrateComplete: newTracker[azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientMigrateCompleteResponse]](),
		beginStopReplicate:   newTracker[azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientStopReplicateResponse]](),
	}
}

// WorkloadInstanceServerTransport connects instances of armmigrationmodernization.WorkloadInstanceClient to instances of WorkloadInstanceServer.
// Don't use this type directly, use NewWorkloadInstanceServerTransport instead.
type WorkloadInstanceServerTransport struct {
	srv                  *WorkloadInstanceServer
	beginDelete          *tracker[azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientDeleteResponse]]
	newListPager         *tracker[azfake.PagerResponder[armmigrationmodernization.WorkloadInstanceClientListResponse]]
	beginMigrateComplete *tracker[azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientMigrateCompleteResponse]]
	beginStopReplicate   *tracker[azfake.PollerResponder[armmigrationmodernization.WorkloadInstanceClientStopReplicateResponse]]
}

// Do implements the policy.Transporter interface for WorkloadInstanceServerTransport.
func (w *WorkloadInstanceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkloadInstanceClient.Create":
		resp, err = w.dispatchCreate(req)
	case "WorkloadInstanceClient.BeginDelete":
		resp, err = w.dispatchBeginDelete(req)
	case "WorkloadInstanceClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkloadInstanceClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkloadInstanceClient.BeginMigrateComplete":
		resp, err = w.dispatchBeginMigrateComplete(req)
	case "WorkloadInstanceClient.BeginStopReplicate":
		resp, err = w.dispatchBeginStopReplicate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkloadInstanceServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if w.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadInstances/(?P<workloadInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrationmodernization.WorkloadInstanceModel](req)
	if err != nil {
		return nil, err
	}
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
	if err != nil {
		return nil, err
	}
	workloadInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workloadInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Create(req.Context(), subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, workloadInstanceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkloadInstanceModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkloadInstanceServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := w.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadInstances/(?P<workloadInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
		if err != nil {
			return nil, err
		}
		workloadInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workloadInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginDelete(req.Context(), subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, workloadInstanceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		w.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		w.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		w.beginDelete.remove(req)
	}

	return resp, nil
}

func (w *WorkloadInstanceServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadInstances/(?P<workloadInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
	if err != nil {
		return nil, err
	}
	workloadInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workloadInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, workloadInstanceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkloadInstanceModel, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkloadInstanceServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmigrationmodernization.WorkloadInstanceClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkloadInstanceServerTransport) dispatchBeginMigrateComplete(req *http.Request) (*http.Response, error) {
	if w.srv.BeginMigrateComplete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginMigrateComplete not implemented")}
	}
	beginMigrateComplete := w.beginMigrateComplete.get(req)
	if beginMigrateComplete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadInstances/(?P<workloadInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/completeMigration`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
		if err != nil {
			return nil, err
		}
		workloadInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workloadInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginMigrateComplete(req.Context(), subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, workloadInstanceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginMigrateComplete = &respr
		w.beginMigrateComplete.add(req, beginMigrateComplete)
	}

	resp, err := server.PollerResponderNext(beginMigrateComplete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		w.beginMigrateComplete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginMigrateComplete) {
		w.beginMigrateComplete.remove(req)
	}

	return resp, nil
}

func (w *WorkloadInstanceServerTransport) dispatchBeginStopReplicate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginStopReplicate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopReplicate not implemented")}
	}
	beginStopReplicate := w.beginStopReplicate.get(req)
	if beginStopReplicate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/modernizeProjects/(?P<modernizeProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadInstances/(?P<workloadInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disableReplication`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		modernizeProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modernizeProjectName")])
		if err != nil {
			return nil, err
		}
		workloadInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workloadInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginStopReplicate(req.Context(), subscriptionIDParam, resourceGroupNameParam, modernizeProjectNameParam, workloadInstanceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStopReplicate = &respr
		w.beginStopReplicate.add(req, beginStopReplicate)
	}

	resp, err := server.PollerResponderNext(beginStopReplicate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		w.beginStopReplicate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStopReplicate) {
		w.beginStopReplicate.remove(req)
	}

	return resp, nil
}
