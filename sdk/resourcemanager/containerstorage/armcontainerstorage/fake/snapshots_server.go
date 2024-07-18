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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerstorage/armcontainerstorage"
	"net/http"
	"net/url"
	"regexp"
)

// SnapshotsServer is a fake server for instances of the armcontainerstorage.SnapshotsClient type.
type SnapshotsServer struct {
	// BeginCreateOrUpdate is the fake for method SnapshotsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, poolName string, snapshotName string, resource armcontainerstorage.Snapshot, options *armcontainerstorage.SnapshotsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcontainerstorage.SnapshotsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SnapshotsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, poolName string, snapshotName string, options *armcontainerstorage.SnapshotsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcontainerstorage.SnapshotsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SnapshotsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, poolName string, snapshotName string, options *armcontainerstorage.SnapshotsClientGetOptions) (resp azfake.Responder[armcontainerstorage.SnapshotsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByPoolPager is the fake for method SnapshotsClient.NewListByPoolPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByPoolPager func(resourceGroupName string, poolName string, options *armcontainerstorage.SnapshotsClientListByPoolOptions) (resp azfake.PagerResponder[armcontainerstorage.SnapshotsClientListByPoolResponse])
}

// NewSnapshotsServerTransport creates a new instance of SnapshotsServerTransport with the provided implementation.
// The returned SnapshotsServerTransport instance is connected to an instance of armcontainerstorage.SnapshotsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSnapshotsServerTransport(srv *SnapshotsServer) *SnapshotsServerTransport {
	return &SnapshotsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcontainerstorage.SnapshotsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armcontainerstorage.SnapshotsClientDeleteResponse]](),
		newListByPoolPager:  newTracker[azfake.PagerResponder[armcontainerstorage.SnapshotsClientListByPoolResponse]](),
	}
}

// SnapshotsServerTransport connects instances of armcontainerstorage.SnapshotsClient to instances of SnapshotsServer.
// Don't use this type directly, use NewSnapshotsServerTransport instead.
type SnapshotsServerTransport struct {
	srv                 *SnapshotsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcontainerstorage.SnapshotsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armcontainerstorage.SnapshotsClientDeleteResponse]]
	newListByPoolPager  *tracker[azfake.PagerResponder[armcontainerstorage.SnapshotsClientListByPoolResponse]]
}

// Do implements the policy.Transporter interface for SnapshotsServerTransport.
func (s *SnapshotsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SnapshotsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "SnapshotsClient.BeginCreateOrUpdate":
		resp, err = s.dispatchBeginCreateOrUpdate(req)
	case "SnapshotsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SnapshotsClient.Get":
		resp, err = s.dispatchGet(req)
	case "SnapshotsClient.NewListByPoolPager":
		resp, err = s.dispatchNewListByPoolPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (s *SnapshotsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerStorage/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcontainerstorage.Snapshot](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, poolNameParam, snapshotNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SnapshotsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerStorage/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, poolNameParam, snapshotNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SnapshotsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerStorage/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots/(?P<snapshotName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	snapshotNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("snapshotName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, poolNameParam, snapshotNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Snapshot, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SnapshotsServerTransport) dispatchNewListByPoolPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByPoolPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByPoolPager not implemented")}
	}
	newListByPoolPager := s.newListByPoolPager.get(req)
	if newListByPoolPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ContainerStorage/pools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/snapshots`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByPoolPager(resourceGroupNameParam, poolNameParam, nil)
		newListByPoolPager = &resp
		s.newListByPoolPager.add(req, newListByPoolPager)
		server.PagerResponderInjectNextLinks(newListByPoolPager, req, func(page *armcontainerstorage.SnapshotsClientListByPoolResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByPoolPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByPoolPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByPoolPager) {
		s.newListByPoolPager.remove(req)
	}
	return resp, nil
}
