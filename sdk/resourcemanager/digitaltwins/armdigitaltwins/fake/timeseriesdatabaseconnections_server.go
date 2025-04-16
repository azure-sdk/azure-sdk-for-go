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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/digitaltwins/armdigitaltwins"
	"net/http"
	"net/url"
	"regexp"
)

// TimeSeriesDatabaseConnectionsServer is a fake server for instances of the armdigitaltwins.TimeSeriesDatabaseConnectionsClient type.
type TimeSeriesDatabaseConnectionsServer struct {
	// BeginCreateOrUpdate is the fake for method TimeSeriesDatabaseConnectionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, timeSeriesDatabaseConnectionDescription armdigitaltwins.TimeSeriesDatabaseConnection, options *armdigitaltwins.TimeSeriesDatabaseConnectionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method TimeSeriesDatabaseConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *armdigitaltwins.TimeSeriesDatabaseConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TimeSeriesDatabaseConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, timeSeriesDatabaseConnectionName string, options *armdigitaltwins.TimeSeriesDatabaseConnectionsClientGetOptions) (resp azfake.Responder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TimeSeriesDatabaseConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armdigitaltwins.TimeSeriesDatabaseConnectionsClientListOptions) (resp azfake.PagerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientListResponse])
}

// NewTimeSeriesDatabaseConnectionsServerTransport creates a new instance of TimeSeriesDatabaseConnectionsServerTransport with the provided implementation.
// The returned TimeSeriesDatabaseConnectionsServerTransport instance is connected to an instance of armdigitaltwins.TimeSeriesDatabaseConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTimeSeriesDatabaseConnectionsServerTransport(srv *TimeSeriesDatabaseConnectionsServer) *TimeSeriesDatabaseConnectionsServerTransport {
	return &TimeSeriesDatabaseConnectionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientListResponse]](),
	}
}

// TimeSeriesDatabaseConnectionsServerTransport connects instances of armdigitaltwins.TimeSeriesDatabaseConnectionsClient to instances of TimeSeriesDatabaseConnectionsServer.
// Don't use this type directly, use NewTimeSeriesDatabaseConnectionsServerTransport instead.
type TimeSeriesDatabaseConnectionsServerTransport struct {
	srv                 *TimeSeriesDatabaseConnectionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armdigitaltwins.TimeSeriesDatabaseConnectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for TimeSeriesDatabaseConnectionsServerTransport.
func (t *TimeSeriesDatabaseConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TimeSeriesDatabaseConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if timeSeriesDatabaseConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = timeSeriesDatabaseConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TimeSeriesDatabaseConnectionsClient.BeginCreateOrUpdate":
				res.resp, res.err = t.dispatchBeginCreateOrUpdate(req)
			case "TimeSeriesDatabaseConnectionsClient.BeginDelete":
				res.resp, res.err = t.dispatchBeginDelete(req)
			case "TimeSeriesDatabaseConnectionsClient.Get":
				res.resp, res.err = t.dispatchGet(req)
			case "TimeSeriesDatabaseConnectionsClient.NewListPager":
				res.resp, res.err = t.dispatchNewListPager(req)
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

func (t *TimeSeriesDatabaseConnectionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := t.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DigitalTwins/digitalTwinsInstances/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/timeSeriesDatabaseConnections/(?P<timeSeriesDatabaseConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdigitaltwins.TimeSeriesDatabaseConnection](req)
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
		timeSeriesDatabaseConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("timeSeriesDatabaseConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, timeSeriesDatabaseConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		t.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		t.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		t.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (t *TimeSeriesDatabaseConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := t.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DigitalTwins/digitalTwinsInstances/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/timeSeriesDatabaseConnections/(?P<timeSeriesDatabaseConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		timeSeriesDatabaseConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("timeSeriesDatabaseConnectionName")])
		if err != nil {
			return nil, err
		}
		cleanupConnectionArtifactsUnescaped, err := url.QueryUnescape(qp.Get("cleanupConnectionArtifacts"))
		if err != nil {
			return nil, err
		}
		cleanupConnectionArtifactsParam := getOptional(armdigitaltwins.CleanupConnectionArtifacts(cleanupConnectionArtifactsUnescaped))
		var options *armdigitaltwins.TimeSeriesDatabaseConnectionsClientBeginDeleteOptions
		if cleanupConnectionArtifactsParam != nil {
			options = &armdigitaltwins.TimeSeriesDatabaseConnectionsClientBeginDeleteOptions{
				CleanupConnectionArtifacts: cleanupConnectionArtifactsParam,
			}
		}
		respr, errRespr := t.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, timeSeriesDatabaseConnectionNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		t.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		t.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		t.beginDelete.remove(req)
	}

	return resp, nil
}

func (t *TimeSeriesDatabaseConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DigitalTwins/digitalTwinsInstances/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/timeSeriesDatabaseConnections/(?P<timeSeriesDatabaseConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	timeSeriesDatabaseConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("timeSeriesDatabaseConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, timeSeriesDatabaseConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TimeSeriesDatabaseConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TimeSeriesDatabaseConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DigitalTwins/digitalTwinsInstances/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/timeSeriesDatabaseConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := t.srv.NewListPager(resourceGroupNameParam, resourceNameParam, nil)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdigitaltwins.TimeSeriesDatabaseConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to TimeSeriesDatabaseConnectionsServerTransport
var timeSeriesDatabaseConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
