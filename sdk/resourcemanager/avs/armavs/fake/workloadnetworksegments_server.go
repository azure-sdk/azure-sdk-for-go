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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
	"net/http"
	"net/url"
	"regexp"
)

// WorkloadNetworkSegmentsServer is a fake server for instances of the armavs.WorkloadNetworkSegmentsClient type.
type WorkloadNetworkSegmentsServer struct {
	// BeginCreate is the fake for method WorkloadNetworkSegmentsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string, workloadNetworkSegment armavs.WorkloadNetworkSegment, options *armavs.WorkloadNetworkSegmentsClientBeginCreateOptions) (resp azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDeleteSegment is the fake for method WorkloadNetworkSegmentsClient.BeginDeleteSegment
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDeleteSegment func(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string, options *armavs.WorkloadNetworkSegmentsClientBeginDeleteSegmentOptions) (resp azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientDeleteSegmentResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkloadNetworkSegmentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string, options *armavs.WorkloadNetworkSegmentsClientGetOptions) (resp azfake.Responder[armavs.WorkloadNetworkSegmentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkloadNetworkSegmentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, privateCloudName string, options *armavs.WorkloadNetworkSegmentsClientListOptions) (resp azfake.PagerResponder[armavs.WorkloadNetworkSegmentsClientListResponse])

	// BeginUpdate is the fake for method WorkloadNetworkSegmentsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, segmentID string, properties armavs.WorkloadNetworkSegment, options *armavs.WorkloadNetworkSegmentsClientBeginUpdateOptions) (resp azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWorkloadNetworkSegmentsServerTransport creates a new instance of WorkloadNetworkSegmentsServerTransport with the provided implementation.
// The returned WorkloadNetworkSegmentsServerTransport instance is connected to an instance of armavs.WorkloadNetworkSegmentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkloadNetworkSegmentsServerTransport(srv *WorkloadNetworkSegmentsServer) *WorkloadNetworkSegmentsServerTransport {
	return &WorkloadNetworkSegmentsServerTransport{
		srv:                srv,
		beginCreate:        newTracker[azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientCreateResponse]](),
		beginDeleteSegment: newTracker[azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientDeleteSegmentResponse]](),
		newListPager:       newTracker[azfake.PagerResponder[armavs.WorkloadNetworkSegmentsClientListResponse]](),
		beginUpdate:        newTracker[azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientUpdateResponse]](),
	}
}

// WorkloadNetworkSegmentsServerTransport connects instances of armavs.WorkloadNetworkSegmentsClient to instances of WorkloadNetworkSegmentsServer.
// Don't use this type directly, use NewWorkloadNetworkSegmentsServerTransport instead.
type WorkloadNetworkSegmentsServerTransport struct {
	srv                *WorkloadNetworkSegmentsServer
	beginCreate        *tracker[azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientCreateResponse]]
	beginDeleteSegment *tracker[azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientDeleteSegmentResponse]]
	newListPager       *tracker[azfake.PagerResponder[armavs.WorkloadNetworkSegmentsClientListResponse]]
	beginUpdate        *tracker[azfake.PollerResponder[armavs.WorkloadNetworkSegmentsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for WorkloadNetworkSegmentsServerTransport.
func (w *WorkloadNetworkSegmentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WorkloadNetworkSegmentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "WorkloadNetworkSegmentsClient.BeginCreate":
		resp, err = w.dispatchBeginCreate(req)
	case "WorkloadNetworkSegmentsClient.BeginDeleteSegment":
		resp, err = w.dispatchBeginDeleteSegment(req)
	case "WorkloadNetworkSegmentsClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkloadNetworkSegmentsClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkloadNetworkSegmentsClient.BeginUpdate":
		resp, err = w.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (w *WorkloadNetworkSegmentsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := w.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadNetworks/default/segments/(?P<segmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.WorkloadNetworkSegment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		segmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("segmentId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginCreate(req.Context(), resourceGroupNameParam, privateCloudNameParam, segmentIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		w.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		w.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		w.beginCreate.remove(req)
	}

	return resp, nil
}

func (w *WorkloadNetworkSegmentsServerTransport) dispatchBeginDeleteSegment(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDeleteSegment == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDeleteSegment not implemented")}
	}
	beginDeleteSegment := w.beginDeleteSegment.get(req)
	if beginDeleteSegment == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadNetworks/default/segments/(?P<segmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		segmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("segmentId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginDeleteSegment(req.Context(), resourceGroupNameParam, privateCloudNameParam, segmentIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDeleteSegment = &respr
		w.beginDeleteSegment.add(req, beginDeleteSegment)
	}

	resp, err := server.PollerResponderNext(beginDeleteSegment, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		w.beginDeleteSegment.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDeleteSegment) {
		w.beginDeleteSegment.remove(req)
	}

	return resp, nil
}

func (w *WorkloadNetworkSegmentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadNetworks/default/segments/(?P<segmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	segmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("segmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, privateCloudNameParam, segmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WorkloadNetworkSegment, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkloadNetworkSegmentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadNetworks/default/segments`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, privateCloudNameParam, nil)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armavs.WorkloadNetworkSegmentsClientListResponse, createLink func() string) {
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

func (w *WorkloadNetworkSegmentsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := w.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workloadNetworks/default/segments/(?P<segmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.WorkloadNetworkSegment](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		segmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("segmentId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, segmentIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		w.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		w.beginUpdate.remove(req)
	}

	return resp, nil
}
