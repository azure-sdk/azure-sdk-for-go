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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasewatcher/armdatabasewatcher"
	"net/http"
	"net/url"
	"regexp"
)

// AlertRuleResourcesServer is a fake server for instances of the armdatabasewatcher.AlertRuleResourcesClient type.
type AlertRuleResourcesServer struct {
	// CreateOrUpdate is the fake for method AlertRuleResourcesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, resource armdatabasewatcher.AlertRuleResource, options *armdatabasewatcher.AlertRuleResourcesClientCreateOrUpdateOptions) (resp azfake.Responder[armdatabasewatcher.AlertRuleResourcesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AlertRuleResourcesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, options *armdatabasewatcher.AlertRuleResourcesClientDeleteOptions) (resp azfake.Responder[armdatabasewatcher.AlertRuleResourcesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AlertRuleResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, watcherName string, alertRuleResourceName string, options *armdatabasewatcher.AlertRuleResourcesClientGetOptions) (resp azfake.Responder[armdatabasewatcher.AlertRuleResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByParentPager is the fake for method AlertRuleResourcesClient.NewListByParentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByParentPager func(resourceGroupName string, watcherName string, options *armdatabasewatcher.AlertRuleResourcesClientListByParentOptions) (resp azfake.PagerResponder[armdatabasewatcher.AlertRuleResourcesClientListByParentResponse])
}

// NewAlertRuleResourcesServerTransport creates a new instance of AlertRuleResourcesServerTransport with the provided implementation.
// The returned AlertRuleResourcesServerTransport instance is connected to an instance of armdatabasewatcher.AlertRuleResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAlertRuleResourcesServerTransport(srv *AlertRuleResourcesServer) *AlertRuleResourcesServerTransport {
	return &AlertRuleResourcesServerTransport{
		srv:                  srv,
		newListByParentPager: newTracker[azfake.PagerResponder[armdatabasewatcher.AlertRuleResourcesClientListByParentResponse]](),
	}
}

// AlertRuleResourcesServerTransport connects instances of armdatabasewatcher.AlertRuleResourcesClient to instances of AlertRuleResourcesServer.
// Don't use this type directly, use NewAlertRuleResourcesServerTransport instead.
type AlertRuleResourcesServerTransport struct {
	srv                  *AlertRuleResourcesServer
	newListByParentPager *tracker[azfake.PagerResponder[armdatabasewatcher.AlertRuleResourcesClientListByParentResponse]]
}

// Do implements the policy.Transporter interface for AlertRuleResourcesServerTransport.
func (a *AlertRuleResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AlertRuleResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if alertRuleResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = alertRuleResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AlertRuleResourcesClient.CreateOrUpdate":
				res.resp, res.err = a.dispatchCreateOrUpdate(req)
			case "AlertRuleResourcesClient.Delete":
				res.resp, res.err = a.dispatchDelete(req)
			case "AlertRuleResourcesClient.Get":
				res.resp, res.err = a.dispatchGet(req)
			case "AlertRuleResourcesClient.NewListByParentPager":
				res.resp, res.err = a.dispatchNewListByParentPager(req)
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

func (a *AlertRuleResourcesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alertRuleResources/(?P<alertRuleResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatabasewatcher.AlertRuleResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	alertRuleResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertRuleResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, watcherNameParam, alertRuleResourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertRuleResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlertRuleResourcesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alertRuleResources/(?P<alertRuleResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	alertRuleResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertRuleResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, watcherNameParam, alertRuleResourceNameParam, nil)
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

func (a *AlertRuleResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alertRuleResources/(?P<alertRuleResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
	if err != nil {
		return nil, err
	}
	alertRuleResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertRuleResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, watcherNameParam, alertRuleResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertRuleResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AlertRuleResourcesServerTransport) dispatchNewListByParentPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByParentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByParentPager not implemented")}
	}
	newListByParentPager := a.newListByParentPager.get(req)
	if newListByParentPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseWatcher/watchers/(?P<watcherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alertRuleResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		watcherNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("watcherName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByParentPager(resourceGroupNameParam, watcherNameParam, nil)
		newListByParentPager = &resp
		a.newListByParentPager.add(req, newListByParentPager)
		server.PagerResponderInjectNextLinks(newListByParentPager, req, func(page *armdatabasewatcher.AlertRuleResourcesClientListByParentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByParentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByParentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByParentPager) {
		a.newListByParentPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AlertRuleResourcesServerTransport
var alertRuleResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
