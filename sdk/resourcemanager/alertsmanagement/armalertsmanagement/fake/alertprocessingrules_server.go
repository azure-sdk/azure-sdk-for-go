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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
	"net/http"
	"net/url"
	"regexp"
)

// AlertProcessingRulesServer is a fake server for instances of the armalertsmanagement.AlertProcessingRulesClient type.
type AlertProcessingRulesServer struct {
	// CreateOrUpdate is the fake for method AlertProcessingRulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, alertProcessingRule armalertsmanagement.AlertProcessingRule, options *armalertsmanagement.AlertProcessingRulesClientCreateOrUpdateOptions) (resp azfake.Responder[armalertsmanagement.AlertProcessingRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method AlertProcessingRulesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, options *armalertsmanagement.AlertProcessingRulesClientDeleteOptions) (resp azfake.Responder[armalertsmanagement.AlertProcessingRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// GetByName is the fake for method AlertProcessingRulesClient.GetByName
	// HTTP status codes to indicate success: http.StatusOK
	GetByName func(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, options *armalertsmanagement.AlertProcessingRulesClientGetByNameOptions) (resp azfake.Responder[armalertsmanagement.AlertProcessingRulesClientGetByNameResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method AlertProcessingRulesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armalertsmanagement.AlertProcessingRulesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armalertsmanagement.AlertProcessingRulesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method AlertProcessingRulesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armalertsmanagement.AlertProcessingRulesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armalertsmanagement.AlertProcessingRulesClientListBySubscriptionResponse])

	// Update is the fake for method AlertProcessingRulesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, alertProcessingRuleName string, alertProcessingRulePatch armalertsmanagement.PatchObject, options *armalertsmanagement.AlertProcessingRulesClientUpdateOptions) (resp azfake.Responder[armalertsmanagement.AlertProcessingRulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewAlertProcessingRulesServerTransport creates a new instance of AlertProcessingRulesServerTransport with the provided implementation.
// The returned AlertProcessingRulesServerTransport instance is connected to an instance of armalertsmanagement.AlertProcessingRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAlertProcessingRulesServerTransport(srv *AlertProcessingRulesServer) *AlertProcessingRulesServerTransport {
	return &AlertProcessingRulesServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armalertsmanagement.AlertProcessingRulesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armalertsmanagement.AlertProcessingRulesClientListBySubscriptionResponse]](),
	}
}

// AlertProcessingRulesServerTransport connects instances of armalertsmanagement.AlertProcessingRulesClient to instances of AlertProcessingRulesServer.
// Don't use this type directly, use NewAlertProcessingRulesServerTransport instead.
type AlertProcessingRulesServerTransport struct {
	srv                         *AlertProcessingRulesServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armalertsmanagement.AlertProcessingRulesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armalertsmanagement.AlertProcessingRulesClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for AlertProcessingRulesServerTransport.
func (a *AlertProcessingRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return a.dispatchToMethodFake(req, method)
}

func (a *AlertProcessingRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if alertProcessingRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = alertProcessingRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "AlertProcessingRulesClient.CreateOrUpdate":
				res.resp, res.err = a.dispatchCreateOrUpdate(req)
			case "AlertProcessingRulesClient.Delete":
				res.resp, res.err = a.dispatchDelete(req)
			case "AlertProcessingRulesClient.GetByName":
				res.resp, res.err = a.dispatchGetByName(req)
			case "AlertProcessingRulesClient.NewListByResourceGroupPager":
				res.resp, res.err = a.dispatchNewListByResourceGroupPager(req)
			case "AlertProcessingRulesClient.NewListBySubscriptionPager":
				res.resp, res.err = a.dispatchNewListBySubscriptionPager(req)
			case "AlertProcessingRulesClient.Update":
				res.resp, res.err = a.dispatchUpdate(req)
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

func (a *AlertProcessingRulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AlertsManagement/actionRules/(?P<alertProcessingRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armalertsmanagement.AlertProcessingRule](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	alertProcessingRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertProcessingRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, alertProcessingRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertProcessingRule, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (a *AlertProcessingRulesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AlertsManagement/actionRules/(?P<alertProcessingRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	alertProcessingRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertProcessingRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, alertProcessingRuleNameParam, nil)
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
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (a *AlertProcessingRulesServerTransport) dispatchGetByName(req *http.Request) (*http.Response, error) {
	if a.srv.GetByName == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByName not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AlertsManagement/actionRules/(?P<alertProcessingRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	alertProcessingRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertProcessingRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetByName(req.Context(), resourceGroupNameParam, alertProcessingRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertProcessingRule, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

func (a *AlertProcessingRulesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := a.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AlertsManagement/actionRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := a.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		a.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armalertsmanagement.AlertProcessingRulesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		a.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (a *AlertProcessingRulesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := a.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AlertsManagement/actionRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := a.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		a.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armalertsmanagement.AlertProcessingRulesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		a.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (a *AlertProcessingRulesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AlertsManagement/actionRules/(?P<alertProcessingRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armalertsmanagement.PatchObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	alertProcessingRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("alertProcessingRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Update(req.Context(), resourceGroupNameParam, alertProcessingRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AlertProcessingRule, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to AlertProcessingRulesServerTransport
var alertProcessingRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
