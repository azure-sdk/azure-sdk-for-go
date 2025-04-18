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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability"
	"net/http"
	"net/url"
	"regexp"
)

// TagRulesServer is a fake server for instances of the armnewrelicobservability.TagRulesClient type.
type TagRulesServer struct {
	// BeginCreateOrUpdate is the fake for method TagRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, resource armnewrelicobservability.TagRule, options *armnewrelicobservability.TagRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnewrelicobservability.TagRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method TagRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *armnewrelicobservability.TagRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnewrelicobservability.TagRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TagRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, options *armnewrelicobservability.TagRulesClientGetOptions) (resp azfake.Responder[armnewrelicobservability.TagRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByNewRelicMonitorResourcePager is the fake for method TagRulesClient.NewListByNewRelicMonitorResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNewRelicMonitorResourcePager func(resourceGroupName string, monitorName string, options *armnewrelicobservability.TagRulesClientListByNewRelicMonitorResourceOptions) (resp azfake.PagerResponder[armnewrelicobservability.TagRulesClientListByNewRelicMonitorResourceResponse])

	// Update is the fake for method TagRulesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, monitorName string, ruleSetName string, properties armnewrelicobservability.TagRuleUpdate, options *armnewrelicobservability.TagRulesClientUpdateOptions) (resp azfake.Responder[armnewrelicobservability.TagRulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewTagRulesServerTransport creates a new instance of TagRulesServerTransport with the provided implementation.
// The returned TagRulesServerTransport instance is connected to an instance of armnewrelicobservability.TagRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTagRulesServerTransport(srv *TagRulesServer) *TagRulesServerTransport {
	return &TagRulesServerTransport{
		srv:                                   srv,
		beginCreateOrUpdate:                   newTracker[azfake.PollerResponder[armnewrelicobservability.TagRulesClientCreateOrUpdateResponse]](),
		beginDelete:                           newTracker[azfake.PollerResponder[armnewrelicobservability.TagRulesClientDeleteResponse]](),
		newListByNewRelicMonitorResourcePager: newTracker[azfake.PagerResponder[armnewrelicobservability.TagRulesClientListByNewRelicMonitorResourceResponse]](),
	}
}

// TagRulesServerTransport connects instances of armnewrelicobservability.TagRulesClient to instances of TagRulesServer.
// Don't use this type directly, use NewTagRulesServerTransport instead.
type TagRulesServerTransport struct {
	srv                                   *TagRulesServer
	beginCreateOrUpdate                   *tracker[azfake.PollerResponder[armnewrelicobservability.TagRulesClientCreateOrUpdateResponse]]
	beginDelete                           *tracker[azfake.PollerResponder[armnewrelicobservability.TagRulesClientDeleteResponse]]
	newListByNewRelicMonitorResourcePager *tracker[azfake.PagerResponder[armnewrelicobservability.TagRulesClientListByNewRelicMonitorResourceResponse]]
}

// Do implements the policy.Transporter interface for TagRulesServerTransport.
func (t *TagRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TagRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if tagRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = tagRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TagRulesClient.BeginCreateOrUpdate":
				res.resp, res.err = t.dispatchBeginCreateOrUpdate(req)
			case "TagRulesClient.BeginDelete":
				res.resp, res.err = t.dispatchBeginDelete(req)
			case "TagRulesClient.Get":
				res.resp, res.err = t.dispatchGet(req)
			case "TagRulesClient.NewListByNewRelicMonitorResourcePager":
				res.resp, res.err = t.dispatchNewListByNewRelicMonitorResourcePager(req)
			case "TagRulesClient.Update":
				res.resp, res.err = t.dispatchUpdate(req)
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

func (t *TagRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := t.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/NewRelic\.Observability/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagRules/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnewrelicobservability.TagRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, monitorNameParam, ruleSetNameParam, body, nil)
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

func (t *TagRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if t.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := t.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/NewRelic\.Observability/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagRules/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := t.srv.BeginDelete(req.Context(), resourceGroupNameParam, monitorNameParam, ruleSetNameParam, nil)
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

func (t *TagRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/NewRelic\.Observability/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagRules/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
	if err != nil {
		return nil, err
	}
	ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, monitorNameParam, ruleSetNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TagRulesServerTransport) dispatchNewListByNewRelicMonitorResourcePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByNewRelicMonitorResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNewRelicMonitorResourcePager not implemented")}
	}
	newListByNewRelicMonitorResourcePager := t.newListByNewRelicMonitorResourcePager.get(req)
	if newListByNewRelicMonitorResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/NewRelic\.Observability/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListByNewRelicMonitorResourcePager(resourceGroupNameParam, monitorNameParam, nil)
		newListByNewRelicMonitorResourcePager = &resp
		t.newListByNewRelicMonitorResourcePager.add(req, newListByNewRelicMonitorResourcePager)
		server.PagerResponderInjectNextLinks(newListByNewRelicMonitorResourcePager, req, func(page *armnewrelicobservability.TagRulesClientListByNewRelicMonitorResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNewRelicMonitorResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByNewRelicMonitorResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNewRelicMonitorResourcePager) {
		t.newListByNewRelicMonitorResourcePager.remove(req)
	}
	return resp, nil
}

func (t *TagRulesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/NewRelic\.Observability/monitors/(?P<monitorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tagRules/(?P<ruleSetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnewrelicobservability.TagRuleUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	monitorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("monitorName")])
	if err != nil {
		return nil, err
	}
	ruleSetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("ruleSetName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Update(req.Context(), resourceGroupNameParam, monitorNameParam, ruleSetNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TagRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to TagRulesServerTransport
var tagRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
