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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ReachabilityAnalysisRunsServer is a fake server for instances of the armnetwork.ReachabilityAnalysisRunsClient type.
type ReachabilityAnalysisRunsServer struct {
	// BeginCreate is the fake for method ReachabilityAnalysisRunsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisRunName string, body armnetwork.ReachabilityAnalysisRun, options *armnetwork.ReachabilityAnalysisRunsClientBeginCreateOptions) (resp azfake.PollerResponder[armnetwork.ReachabilityAnalysisRunsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ReachabilityAnalysisRunsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisRunName string, options *armnetwork.ReachabilityAnalysisRunsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.ReachabilityAnalysisRunsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReachabilityAnalysisRunsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, workspaceName string, reachabilityAnalysisRunName string, options *armnetwork.ReachabilityAnalysisRunsClientGetOptions) (resp azfake.Responder[armnetwork.ReachabilityAnalysisRunsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReachabilityAnalysisRunsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, workspaceName string, options *armnetwork.ReachabilityAnalysisRunsClientListOptions) (resp azfake.PagerResponder[armnetwork.ReachabilityAnalysisRunsClientListResponse])
}

// NewReachabilityAnalysisRunsServerTransport creates a new instance of ReachabilityAnalysisRunsServerTransport with the provided implementation.
// The returned ReachabilityAnalysisRunsServerTransport instance is connected to an instance of armnetwork.ReachabilityAnalysisRunsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReachabilityAnalysisRunsServerTransport(srv *ReachabilityAnalysisRunsServer) *ReachabilityAnalysisRunsServerTransport {
	return &ReachabilityAnalysisRunsServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armnetwork.ReachabilityAnalysisRunsClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.ReachabilityAnalysisRunsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ReachabilityAnalysisRunsClientListResponse]](),
	}
}

// ReachabilityAnalysisRunsServerTransport connects instances of armnetwork.ReachabilityAnalysisRunsClient to instances of ReachabilityAnalysisRunsServer.
// Don't use this type directly, use NewReachabilityAnalysisRunsServerTransport instead.
type ReachabilityAnalysisRunsServerTransport struct {
	srv          *ReachabilityAnalysisRunsServer
	beginCreate  *tracker[azfake.PollerResponder[armnetwork.ReachabilityAnalysisRunsClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.ReachabilityAnalysisRunsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.ReachabilityAnalysisRunsClientListResponse]]
}

// Do implements the policy.Transporter interface for ReachabilityAnalysisRunsServerTransport.
func (r *ReachabilityAnalysisRunsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReachabilityAnalysisRunsClient.BeginCreate":
		resp, err = r.dispatchBeginCreate(req)
	case "ReachabilityAnalysisRunsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "ReachabilityAnalysisRunsClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReachabilityAnalysisRunsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReachabilityAnalysisRunsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := r.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/verifierWorkspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reachabilityAnalysisRuns/(?P<reachabilityAnalysisRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.ReachabilityAnalysisRun](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		reachabilityAnalysisRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reachabilityAnalysisRunName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreate(req.Context(), resourceGroupNameParam, networkManagerNameParam, workspaceNameParam, reachabilityAnalysisRunNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		r.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		r.beginCreate.remove(req)
	}

	return resp, nil
}

func (r *ReachabilityAnalysisRunsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/verifierWorkspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reachabilityAnalysisRuns/(?P<reachabilityAnalysisRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		reachabilityAnalysisRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reachabilityAnalysisRunName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkManagerNameParam, workspaceNameParam, reachabilityAnalysisRunNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *ReachabilityAnalysisRunsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/verifierWorkspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reachabilityAnalysisRuns/(?P<reachabilityAnalysisRunName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	reachabilityAnalysisRunNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reachabilityAnalysisRunName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, networkManagerNameParam, workspaceNameParam, reachabilityAnalysisRunNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReachabilityAnalysisRun, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReachabilityAnalysisRunsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/verifierWorkspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reachabilityAnalysisRuns`
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
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		sortKeyUnescaped, err := url.QueryUnescape(qp.Get("sortKey"))
		if err != nil {
			return nil, err
		}
		sortKeyParam := getOptional(sortKeyUnescaped)
		sortValueUnescaped, err := url.QueryUnescape(qp.Get("sortValue"))
		if err != nil {
			return nil, err
		}
		sortValueParam := getOptional(sortValueUnescaped)
		var options *armnetwork.ReachabilityAnalysisRunsClientListOptions
		if skipTokenParam != nil || skipParam != nil || topParam != nil || sortKeyParam != nil || sortValueParam != nil {
			options = &armnetwork.ReachabilityAnalysisRunsClientListOptions{
				SkipToken: skipTokenParam,
				Skip:      skipParam,
				Top:       topParam,
				SortKey:   sortKeyParam,
				SortValue: sortValueParam,
			}
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, networkManagerNameParam, workspaceNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ReachabilityAnalysisRunsClientListResponse, createLink func() string) {
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
