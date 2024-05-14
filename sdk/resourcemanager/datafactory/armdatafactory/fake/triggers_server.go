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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"net/http"
	"net/url"
	"regexp"
)

// TriggersServer is a fake server for instances of the armdatafactory.TriggersClient type.
type TriggersServer struct {
	// CreateOrUpdate is the fake for method TriggersClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, trigger armdatafactory.TriggerResource, options *armdatafactory.TriggersClientCreateOrUpdateOptions) (resp azfake.Responder[armdatafactory.TriggersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TriggersClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientDeleteOptions) (resp azfake.Responder[armdatafactory.TriggersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TriggersClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	Get func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientGetOptions) (resp azfake.Responder[armdatafactory.TriggersClientGetResponse], errResp azfake.ErrorResponder)

	// GetEventSubscriptionStatus is the fake for method TriggersClient.GetEventSubscriptionStatus
	// HTTP status codes to indicate success: http.StatusOK
	GetEventSubscriptionStatus func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientGetEventSubscriptionStatusOptions) (resp azfake.Responder[armdatafactory.TriggersClientGetEventSubscriptionStatusResponse], errResp azfake.ErrorResponder)

	// NewListByFactoryPager is the fake for method TriggersClient.NewListByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFactoryPager func(resourceGroupName string, factoryName string, options *armdatafactory.TriggersClientListByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.TriggersClientListByFactoryResponse])

	// QueryByFactory is the fake for method TriggersClient.QueryByFactory
	// HTTP status codes to indicate success: http.StatusOK
	QueryByFactory func(ctx context.Context, resourceGroupName string, factoryName string, filterParameters armdatafactory.TriggerFilterParameters, options *armdatafactory.TriggersClientQueryByFactoryOptions) (resp azfake.Responder[armdatafactory.TriggersClientQueryByFactoryResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method TriggersClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK
	BeginStart func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientBeginStartOptions) (resp azfake.PollerResponder[armdatafactory.TriggersClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method TriggersClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK
	BeginStop func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientBeginStopOptions) (resp azfake.PollerResponder[armdatafactory.TriggersClientStopResponse], errResp azfake.ErrorResponder)

	// BeginSubscribeToEvents is the fake for method TriggersClient.BeginSubscribeToEvents
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSubscribeToEvents func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientBeginSubscribeToEventsOptions) (resp azfake.PollerResponder[armdatafactory.TriggersClientSubscribeToEventsResponse], errResp azfake.ErrorResponder)

	// BeginUnsubscribeFromEvents is the fake for method TriggersClient.BeginUnsubscribeFromEvents
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUnsubscribeFromEvents func(ctx context.Context, resourceGroupName string, factoryName string, triggerName string, options *armdatafactory.TriggersClientBeginUnsubscribeFromEventsOptions) (resp azfake.PollerResponder[armdatafactory.TriggersClientUnsubscribeFromEventsResponse], errResp azfake.ErrorResponder)
}

// NewTriggersServerTransport creates a new instance of TriggersServerTransport with the provided implementation.
// The returned TriggersServerTransport instance is connected to an instance of armdatafactory.TriggersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTriggersServerTransport(srv *TriggersServer) *TriggersServerTransport {
	return &TriggersServerTransport{
		srv:                        srv,
		newListByFactoryPager:      newTracker[azfake.PagerResponder[armdatafactory.TriggersClientListByFactoryResponse]](),
		beginStart:                 newTracker[azfake.PollerResponder[armdatafactory.TriggersClientStartResponse]](),
		beginStop:                  newTracker[azfake.PollerResponder[armdatafactory.TriggersClientStopResponse]](),
		beginSubscribeToEvents:     newTracker[azfake.PollerResponder[armdatafactory.TriggersClientSubscribeToEventsResponse]](),
		beginUnsubscribeFromEvents: newTracker[azfake.PollerResponder[armdatafactory.TriggersClientUnsubscribeFromEventsResponse]](),
	}
}

// TriggersServerTransport connects instances of armdatafactory.TriggersClient to instances of TriggersServer.
// Don't use this type directly, use NewTriggersServerTransport instead.
type TriggersServerTransport struct {
	srv                        *TriggersServer
	newListByFactoryPager      *tracker[azfake.PagerResponder[armdatafactory.TriggersClientListByFactoryResponse]]
	beginStart                 *tracker[azfake.PollerResponder[armdatafactory.TriggersClientStartResponse]]
	beginStop                  *tracker[azfake.PollerResponder[armdatafactory.TriggersClientStopResponse]]
	beginSubscribeToEvents     *tracker[azfake.PollerResponder[armdatafactory.TriggersClientSubscribeToEventsResponse]]
	beginUnsubscribeFromEvents *tracker[azfake.PollerResponder[armdatafactory.TriggersClientUnsubscribeFromEventsResponse]]
}

// Do implements the policy.Transporter interface for TriggersServerTransport.
func (t *TriggersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TriggersClient.CreateOrUpdate":
		resp, err = t.dispatchCreateOrUpdate(req)
	case "TriggersClient.Delete":
		resp, err = t.dispatchDelete(req)
	case "TriggersClient.Get":
		resp, err = t.dispatchGet(req)
	case "TriggersClient.GetEventSubscriptionStatus":
		resp, err = t.dispatchGetEventSubscriptionStatus(req)
	case "TriggersClient.NewListByFactoryPager":
		resp, err = t.dispatchNewListByFactoryPager(req)
	case "TriggersClient.QueryByFactory":
		resp, err = t.dispatchQueryByFactory(req)
	case "TriggersClient.BeginStart":
		resp, err = t.dispatchBeginStart(req)
	case "TriggersClient.BeginStop":
		resp, err = t.dispatchBeginStop(req)
	case "TriggersClient.BeginSubscribeToEvents":
		resp, err = t.dispatchBeginSubscribeToEvents(req)
	case "TriggersClient.BeginUnsubscribeFromEvents":
		resp, err = t.dispatchBeginUnsubscribeFromEvents(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TriggersServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.TriggerResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armdatafactory.TriggersClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armdatafactory.TriggersClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggersServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, nil)
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

func (t *TriggersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armdatafactory.TriggersClientGetOptions
	if ifNoneMatchParam != nil {
		options = &armdatafactory.TriggersClientGetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggersServerTransport) dispatchGetEventSubscriptionStatus(req *http.Request) (*http.Response, error) {
	if t.srv.GetEventSubscriptionStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEventSubscriptionStatus not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getEventSubscriptionStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.GetEventSubscriptionStatus(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerSubscriptionOperationStatus, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggersServerTransport) dispatchNewListByFactoryPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFactoryPager not implemented")}
	}
	newListByFactoryPager := t.newListByFactoryPager.get(req)
	if newListByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByFactoryPager(resourceGroupNameParam, factoryNameParam, nil)
		newListByFactoryPager = &resp
		t.newListByFactoryPager.add(req, newListByFactoryPager)
		server.PagerResponderInjectNextLinks(newListByFactoryPager, req, func(page *armdatafactory.TriggersClientListByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFactoryPager) {
		t.newListByFactoryPager.remove(req)
	}
	return resp, nil
}

func (t *TriggersServerTransport) dispatchQueryByFactory(req *http.Request) (*http.Response, error) {
	if t.srv.QueryByFactory == nil {
		return nil, &nonRetriableError{errors.New("fake for method QueryByFactory not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/querytriggers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.TriggerFilterParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.QueryByFactory(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TriggerQueryResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TriggersServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if t.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := t.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginStart(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		t.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		t.beginStart.remove(req)
	}

	return resp, nil
}

func (t *TriggersServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if t.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := t.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginStop(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		t.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		t.beginStop.remove(req)
	}

	return resp, nil
}

func (t *TriggersServerTransport) dispatchBeginSubscribeToEvents(req *http.Request) (*http.Response, error) {
	if t.srv.BeginSubscribeToEvents == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSubscribeToEvents not implemented")}
	}
	beginSubscribeToEvents := t.beginSubscribeToEvents.get(req)
	if beginSubscribeToEvents == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscribeToEvents`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginSubscribeToEvents(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSubscribeToEvents = &respr
		t.beginSubscribeToEvents.add(req, beginSubscribeToEvents)
	}

	resp, err := server.PollerResponderNext(beginSubscribeToEvents, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginSubscribeToEvents.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSubscribeToEvents) {
		t.beginSubscribeToEvents.remove(req)
	}

	return resp, nil
}

func (t *TriggersServerTransport) dispatchBeginUnsubscribeFromEvents(req *http.Request) (*http.Response, error) {
	if t.srv.BeginUnsubscribeFromEvents == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUnsubscribeFromEvents not implemented")}
	}
	beginUnsubscribeFromEvents := t.beginUnsubscribeFromEvents.get(req)
	if beginUnsubscribeFromEvents == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggers/(?P<triggerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unsubscribeFromEvents`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		triggerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("triggerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginUnsubscribeFromEvents(req.Context(), resourceGroupNameParam, factoryNameParam, triggerNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUnsubscribeFromEvents = &respr
		t.beginUnsubscribeFromEvents.add(req, beginUnsubscribeFromEvents)
	}

	resp, err := server.PollerResponderNext(beginUnsubscribeFromEvents, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		t.beginUnsubscribeFromEvents.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUnsubscribeFromEvents) {
		t.beginUnsubscribeFromEvents.remove(req)
	}

	return resp, nil
}
