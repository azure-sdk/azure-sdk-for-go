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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
	"net/http"
	"net/url"
	"regexp"
)

// ConnectionServer is a fake server for instances of the armautomation.ConnectionClient type.
type ConnectionServer struct {
	// CreateOrUpdate is the fake for method ConnectionClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters armautomation.ConnectionCreateOrUpdateParameters, options *armautomation.ConnectionClientCreateOrUpdateOptions) (resp azfake.Responder[armautomation.ConnectionClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ConnectionClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, options *armautomation.ConnectionClientDeleteOptions) (resp azfake.Responder[armautomation.ConnectionClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ConnectionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, options *armautomation.ConnectionClientGetOptions) (resp azfake.Responder[armautomation.ConnectionClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAutomationAccountPager is the fake for method ConnectionClient.NewListByAutomationAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAutomationAccountPager func(resourceGroupName string, automationAccountName string, options *armautomation.ConnectionClientListByAutomationAccountOptions) (resp azfake.PagerResponder[armautomation.ConnectionClientListByAutomationAccountResponse])

	// Update is the fake for method ConnectionClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, automationAccountName string, connectionName string, parameters armautomation.ConnectionUpdateParameters, options *armautomation.ConnectionClientUpdateOptions) (resp azfake.Responder[armautomation.ConnectionClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewConnectionServerTransport creates a new instance of ConnectionServerTransport with the provided implementation.
// The returned ConnectionServerTransport instance is connected to an instance of armautomation.ConnectionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewConnectionServerTransport(srv *ConnectionServer) *ConnectionServerTransport {
	return &ConnectionServerTransport{
		srv:                             srv,
		newListByAutomationAccountPager: newTracker[azfake.PagerResponder[armautomation.ConnectionClientListByAutomationAccountResponse]](),
	}
}

// ConnectionServerTransport connects instances of armautomation.ConnectionClient to instances of ConnectionServer.
// Don't use this type directly, use NewConnectionServerTransport instead.
type ConnectionServerTransport struct {
	srv                             *ConnectionServer
	newListByAutomationAccountPager *tracker[azfake.PagerResponder[armautomation.ConnectionClientListByAutomationAccountResponse]]
}

// Do implements the policy.Transporter interface for ConnectionServerTransport.
func (c *ConnectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ConnectionClient.CreateOrUpdate":
		resp, err = c.dispatchCreateOrUpdate(req)
	case "ConnectionClient.Delete":
		resp, err = c.dispatchDelete(req)
	case "ConnectionClient.Get":
		resp, err = c.dispatchGet(req)
	case "ConnectionClient.NewListByAutomationAccountPager":
		resp, err = c.dispatchNewListByAutomationAccountPager(req)
	case "ConnectionClient.Update":
		resp, err = c.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ConnectionServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armautomation.ConnectionCreateOrUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, automationAccountNameParam, connectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Connection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectionServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if c.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Delete(req.Context(), resourceGroupNameParam, automationAccountNameParam, connectionNameParam, nil)
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

func (c *ConnectionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, automationAccountNameParam, connectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Connection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ConnectionServerTransport) dispatchNewListByAutomationAccountPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByAutomationAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAutomationAccountPager not implemented")}
	}
	newListByAutomationAccountPager := c.newListByAutomationAccountPager.get(req)
	if newListByAutomationAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByAutomationAccountPager(resourceGroupNameParam, automationAccountNameParam, nil)
		newListByAutomationAccountPager = &resp
		c.newListByAutomationAccountPager.add(req, newListByAutomationAccountPager)
		server.PagerResponderInjectNextLinks(newListByAutomationAccountPager, req, func(page *armautomation.ConnectionClientListByAutomationAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAutomationAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByAutomationAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAutomationAccountPager) {
		c.newListByAutomationAccountPager.remove(req)
	}
	return resp, nil
}

func (c *ConnectionServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Automation/automationAccounts/(?P<automationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/connections/(?P<connectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armautomation.ConnectionUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	automationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("automationAccountName")])
	if err != nil {
		return nil, err
	}
	connectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("connectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Update(req.Context(), resourceGroupNameParam, automationAccountNameParam, connectionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Connection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
