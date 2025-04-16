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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/signalr/armsignalr"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicaSharedPrivateLinkResourcesServer is a fake server for instances of the armsignalr.ReplicaSharedPrivateLinkResourcesClient type.
type ReplicaSharedPrivateLinkResourcesServer struct {
	// BeginCreateOrUpdate is the fake for method ReplicaSharedPrivateLinkResourcesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, replicaName string, sharedPrivateLinkResourceName string, parameters armsignalr.SharedPrivateLinkResource, options *armsignalr.ReplicaSharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsignalr.ReplicaSharedPrivateLinkResourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReplicaSharedPrivateLinkResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, replicaName string, sharedPrivateLinkResourceName string, options *armsignalr.ReplicaSharedPrivateLinkResourcesClientGetOptions) (resp azfake.Responder[armsignalr.ReplicaSharedPrivateLinkResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReplicaSharedPrivateLinkResourcesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, replicaName string, options *armsignalr.ReplicaSharedPrivateLinkResourcesClientListOptions) (resp azfake.PagerResponder[armsignalr.ReplicaSharedPrivateLinkResourcesClientListResponse])
}

// NewReplicaSharedPrivateLinkResourcesServerTransport creates a new instance of ReplicaSharedPrivateLinkResourcesServerTransport with the provided implementation.
// The returned ReplicaSharedPrivateLinkResourcesServerTransport instance is connected to an instance of armsignalr.ReplicaSharedPrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicaSharedPrivateLinkResourcesServerTransport(srv *ReplicaSharedPrivateLinkResourcesServer) *ReplicaSharedPrivateLinkResourcesServerTransport {
	return &ReplicaSharedPrivateLinkResourcesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armsignalr.ReplicaSharedPrivateLinkResourcesClientCreateOrUpdateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armsignalr.ReplicaSharedPrivateLinkResourcesClientListResponse]](),
	}
}

// ReplicaSharedPrivateLinkResourcesServerTransport connects instances of armsignalr.ReplicaSharedPrivateLinkResourcesClient to instances of ReplicaSharedPrivateLinkResourcesServer.
// Don't use this type directly, use NewReplicaSharedPrivateLinkResourcesServerTransport instead.
type ReplicaSharedPrivateLinkResourcesServerTransport struct {
	srv                 *ReplicaSharedPrivateLinkResourcesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armsignalr.ReplicaSharedPrivateLinkResourcesClientCreateOrUpdateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armsignalr.ReplicaSharedPrivateLinkResourcesClientListResponse]]
}

// Do implements the policy.Transporter interface for ReplicaSharedPrivateLinkResourcesServerTransport.
func (r *ReplicaSharedPrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReplicaSharedPrivateLinkResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if replicaSharedPrivateLinkResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = replicaSharedPrivateLinkResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ReplicaSharedPrivateLinkResourcesClient.BeginCreateOrUpdate":
				res.resp, res.err = r.dispatchBeginCreateOrUpdate(req)
			case "ReplicaSharedPrivateLinkResourcesClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ReplicaSharedPrivateLinkResourcesClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
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

func (r *ReplicaSharedPrivateLinkResourcesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SignalRService/signalR/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas/(?P<replicaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsignalr.SharedPrivateLinkResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		replicaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicaName")])
		if err != nil {
			return nil, err
		}
		sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, replicaNameParam, sharedPrivateLinkResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *ReplicaSharedPrivateLinkResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SignalRService/signalR/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas/(?P<replicaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources/(?P<sharedPrivateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	replicaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicaName")])
	if err != nil {
		return nil, err
	}
	sharedPrivateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sharedPrivateLinkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, replicaNameParam, sharedPrivateLinkResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SharedPrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicaSharedPrivateLinkResourcesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SignalRService/signalR/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas/(?P<replicaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sharedPrivateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		replicaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicaName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, resourceNameParam, replicaNameParam, nil)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsignalr.ReplicaSharedPrivateLinkResourcesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ReplicaSharedPrivateLinkResourcesServerTransport
var replicaSharedPrivateLinkResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
