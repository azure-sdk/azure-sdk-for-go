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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// IntegrationAccountSessionsServer is a fake server for instances of the armlogic.IntegrationAccountSessionsClient type.
type IntegrationAccountSessionsServer struct {
	// CreateOrUpdate is the fake for method IntegrationAccountSessionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string, resource armlogic.IntegrationAccountSession, options *armlogic.IntegrationAccountSessionsClientCreateOrUpdateOptions) (resp azfake.Responder[armlogic.IntegrationAccountSessionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method IntegrationAccountSessionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string, options *armlogic.IntegrationAccountSessionsClientDeleteOptions) (resp azfake.Responder[armlogic.IntegrationAccountSessionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IntegrationAccountSessionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, integrationAccountName string, sessionName string, options *armlogic.IntegrationAccountSessionsClientGetOptions) (resp azfake.Responder[armlogic.IntegrationAccountSessionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IntegrationAccountSessionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, integrationAccountName string, options *armlogic.IntegrationAccountSessionsClientListOptions) (resp azfake.PagerResponder[armlogic.IntegrationAccountSessionsClientListResponse])
}

// NewIntegrationAccountSessionsServerTransport creates a new instance of IntegrationAccountSessionsServerTransport with the provided implementation.
// The returned IntegrationAccountSessionsServerTransport instance is connected to an instance of armlogic.IntegrationAccountSessionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIntegrationAccountSessionsServerTransport(srv *IntegrationAccountSessionsServer) *IntegrationAccountSessionsServerTransport {
	return &IntegrationAccountSessionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armlogic.IntegrationAccountSessionsClientListResponse]](),
	}
}

// IntegrationAccountSessionsServerTransport connects instances of armlogic.IntegrationAccountSessionsClient to instances of IntegrationAccountSessionsServer.
// Don't use this type directly, use NewIntegrationAccountSessionsServerTransport instead.
type IntegrationAccountSessionsServerTransport struct {
	srv          *IntegrationAccountSessionsServer
	newListPager *tracker[azfake.PagerResponder[armlogic.IntegrationAccountSessionsClientListResponse]]
}

// Do implements the policy.Transporter interface for IntegrationAccountSessionsServerTransport.
func (i *IntegrationAccountSessionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IntegrationAccountSessionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if integrationAccountSessionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = integrationAccountSessionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IntegrationAccountSessionsClient.CreateOrUpdate":
				res.resp, res.err = i.dispatchCreateOrUpdate(req)
			case "IntegrationAccountSessionsClient.Delete":
				res.resp, res.err = i.dispatchDelete(req)
			case "IntegrationAccountSessionsClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "IntegrationAccountSessionsClient.NewListPager":
				res.resp, res.err = i.dispatchNewListPager(req)
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

func (i *IntegrationAccountSessionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessions/(?P<sessionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armlogic.IntegrationAccountSession](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	sessionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, integrationAccountNameParam, sessionNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationAccountSession, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationAccountSessionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessions/(?P<sessionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	sessionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, integrationAccountNameParam, sessionNameParam, nil)
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

func (i *IntegrationAccountSessionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessions/(?P<sessionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
	if err != nil {
		return nil, err
	}
	sessionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sessionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, integrationAccountNameParam, sessionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IntegrationAccountSession, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IntegrationAccountSessionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Logic/integrationAccounts/(?P<integrationAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/sessions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		integrationAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("integrationAccountName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
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
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armlogic.IntegrationAccountSessionsClientListOptions
		if topParam != nil || filterParam != nil {
			options = &armlogic.IntegrationAccountSessionsClientListOptions{
				Top:    topParam,
				Filter: filterParam,
			}
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, integrationAccountNameParam, options)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armlogic.IntegrationAccountSessionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to IntegrationAccountSessionsServerTransport
var integrationAccountSessionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
