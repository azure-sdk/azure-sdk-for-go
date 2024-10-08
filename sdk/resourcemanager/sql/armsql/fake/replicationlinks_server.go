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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationLinksServer is a fake server for instances of the armsql.ReplicationLinksClient type.
type ReplicationLinksServer struct {
	// BeginCreateOrUpdate is the fake for method ReplicationLinksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, parameters armsql.ReplicationLink, options *armsql.ReplicationLinksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armsql.ReplicationLinksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ReplicationLinksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDelete func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *armsql.ReplicationLinksClientBeginDeleteOptions) (resp azfake.PollerResponder[armsql.ReplicationLinksClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginFailover is the fake for method ReplicationLinksClient.BeginFailover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailover func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *armsql.ReplicationLinksClientBeginFailoverOptions) (resp azfake.PollerResponder[armsql.ReplicationLinksClientFailoverResponse], errResp azfake.ErrorResponder)

	// BeginFailoverAllowDataLoss is the fake for method ReplicationLinksClient.BeginFailoverAllowDataLoss
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginFailoverAllowDataLoss func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *armsql.ReplicationLinksClientBeginFailoverAllowDataLossOptions) (resp azfake.PollerResponder[armsql.ReplicationLinksClientFailoverAllowDataLossResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReplicationLinksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, options *armsql.ReplicationLinksClientGetOptions) (resp azfake.Responder[armsql.ReplicationLinksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByDatabasePager is the fake for method ReplicationLinksClient.NewListByDatabasePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByDatabasePager func(resourceGroupName string, serverName string, databaseName string, options *armsql.ReplicationLinksClientListByDatabaseOptions) (resp azfake.PagerResponder[armsql.ReplicationLinksClientListByDatabaseResponse])

	// NewListByServerPager is the fake for method ReplicationLinksClient.NewListByServerPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerPager func(resourceGroupName string, serverName string, options *armsql.ReplicationLinksClientListByServerOptions) (resp azfake.PagerResponder[armsql.ReplicationLinksClientListByServerResponse])

	// BeginUpdate is the fake for method ReplicationLinksClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, serverName string, databaseName string, linkID string, parameters armsql.ReplicationLinkUpdate, options *armsql.ReplicationLinksClientBeginUpdateOptions) (resp azfake.PollerResponder[armsql.ReplicationLinksClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewReplicationLinksServerTransport creates a new instance of ReplicationLinksServerTransport with the provided implementation.
// The returned ReplicationLinksServerTransport instance is connected to an instance of armsql.ReplicationLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationLinksServerTransport(srv *ReplicationLinksServer) *ReplicationLinksServerTransport {
	return &ReplicationLinksServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armsql.ReplicationLinksClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armsql.ReplicationLinksClientDeleteResponse]](),
		beginFailover:              newTracker[azfake.PollerResponder[armsql.ReplicationLinksClientFailoverResponse]](),
		beginFailoverAllowDataLoss: newTracker[azfake.PollerResponder[armsql.ReplicationLinksClientFailoverAllowDataLossResponse]](),
		newListByDatabasePager:     newTracker[azfake.PagerResponder[armsql.ReplicationLinksClientListByDatabaseResponse]](),
		newListByServerPager:       newTracker[azfake.PagerResponder[armsql.ReplicationLinksClientListByServerResponse]](),
		beginUpdate:                newTracker[azfake.PollerResponder[armsql.ReplicationLinksClientUpdateResponse]](),
	}
}

// ReplicationLinksServerTransport connects instances of armsql.ReplicationLinksClient to instances of ReplicationLinksServer.
// Don't use this type directly, use NewReplicationLinksServerTransport instead.
type ReplicationLinksServerTransport struct {
	srv                        *ReplicationLinksServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armsql.ReplicationLinksClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armsql.ReplicationLinksClientDeleteResponse]]
	beginFailover              *tracker[azfake.PollerResponder[armsql.ReplicationLinksClientFailoverResponse]]
	beginFailoverAllowDataLoss *tracker[azfake.PollerResponder[armsql.ReplicationLinksClientFailoverAllowDataLossResponse]]
	newListByDatabasePager     *tracker[azfake.PagerResponder[armsql.ReplicationLinksClientListByDatabaseResponse]]
	newListByServerPager       *tracker[azfake.PagerResponder[armsql.ReplicationLinksClientListByServerResponse]]
	beginUpdate                *tracker[azfake.PollerResponder[armsql.ReplicationLinksClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ReplicationLinksServerTransport.
func (r *ReplicationLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ReplicationLinksClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "ReplicationLinksClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "ReplicationLinksClient.BeginFailover":
		resp, err = r.dispatchBeginFailover(req)
	case "ReplicationLinksClient.BeginFailoverAllowDataLoss":
		resp, err = r.dispatchBeginFailoverAllowDataLoss(req)
	case "ReplicationLinksClient.Get":
		resp, err = r.dispatchGet(req)
	case "ReplicationLinksClient.NewListByDatabasePager":
		resp, err = r.dispatchNewListByDatabasePager(req)
	case "ReplicationLinksClient.NewListByServerPager":
		resp, err = r.dispatchNewListByServerPager(req)
	case "ReplicationLinksClient.BeginUpdate":
		resp, err = r.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks/(?P<linkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ReplicationLink](req)
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		linkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, linkIDParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks/(?P<linkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		linkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, linkIDParam, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchBeginFailover(req *http.Request) (*http.Response, error) {
	if r.srv.BeginFailover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFailover not implemented")}
	}
	beginFailover := r.beginFailover.get(req)
	if beginFailover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks/(?P<linkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failover`
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		linkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginFailover(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, linkIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFailover = &respr
		r.beginFailover.add(req, beginFailover)
	}

	resp, err := server.PollerResponderNext(beginFailover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginFailover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFailover) {
		r.beginFailover.remove(req)
	}

	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchBeginFailoverAllowDataLoss(req *http.Request) (*http.Response, error) {
	if r.srv.BeginFailoverAllowDataLoss == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginFailoverAllowDataLoss not implemented")}
	}
	beginFailoverAllowDataLoss := r.beginFailoverAllowDataLoss.get(req)
	if beginFailoverAllowDataLoss == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks/(?P<linkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/forceFailoverAllowDataLoss`
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		linkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginFailoverAllowDataLoss(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, linkIDParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginFailoverAllowDataLoss = &respr
		r.beginFailoverAllowDataLoss.add(req, beginFailoverAllowDataLoss)
	}

	resp, err := server.PollerResponderNext(beginFailoverAllowDataLoss, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginFailoverAllowDataLoss.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginFailoverAllowDataLoss) {
		r.beginFailoverAllowDataLoss.remove(req)
	}

	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks/(?P<linkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	linkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, linkIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReplicationLink, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchNewListByDatabasePager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByDatabasePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByDatabasePager not implemented")}
	}
	newListByDatabasePager := r.newListByDatabasePager.get(req)
	if newListByDatabasePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks`
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByDatabasePager(resourceGroupNameParam, serverNameParam, databaseNameParam, nil)
		newListByDatabasePager = &resp
		r.newListByDatabasePager.add(req, newListByDatabasePager)
		server.PagerResponderInjectNextLinks(newListByDatabasePager, req, func(page *armsql.ReplicationLinksClientListByDatabaseResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByDatabasePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByDatabasePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByDatabasePager) {
		r.newListByDatabasePager.remove(req)
	}
	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchNewListByServerPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByServerPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerPager not implemented")}
	}
	newListByServerPager := r.newListByServerPager.get(req)
	if newListByServerPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := r.srv.NewListByServerPager(resourceGroupNameParam, serverNameParam, nil)
		newListByServerPager = &resp
		r.newListByServerPager.add(req, newListByServerPager)
		server.PagerResponderInjectNextLinks(newListByServerPager, req, func(page *armsql.ReplicationLinksClientListByServerResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByServerPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerPager) {
		r.newListByServerPager.remove(req)
	}
	return resp, nil
}

func (r *ReplicationLinksServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := r.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<serverName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationLinks/(?P<linkId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsql.ReplicationLinkUpdate](req)
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
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		linkIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkId")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUpdate(req.Context(), resourceGroupNameParam, serverNameParam, databaseNameParam, linkIDParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		r.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		r.beginUpdate.remove(req)
	}

	return resp, nil
}
