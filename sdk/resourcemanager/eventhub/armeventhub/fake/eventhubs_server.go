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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// EventHubsServer is a fake server for instances of the armeventhub.EventHubsClient type.
type EventHubsServer struct {
	// CreateOrUpdate is the fake for method EventHubsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, parameters armeventhub.Eventhub, options *armeventhub.EventHubsClientCreateOrUpdateOptions) (resp azfake.Responder[armeventhub.EventHubsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAuthorizationRule is the fake for method EventHubsClient.CreateOrUpdateAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters armeventhub.AuthorizationRule, options *armeventhub.EventHubsClientCreateOrUpdateAuthorizationRuleOptions) (resp azfake.Responder[armeventhub.EventHubsClientCreateOrUpdateAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method EventHubsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *armeventhub.EventHubsClientDeleteOptions) (resp azfake.Responder[armeventhub.EventHubsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAuthorizationRule is the fake for method EventHubsClient.DeleteAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *armeventhub.EventHubsClientDeleteAuthorizationRuleOptions) (resp azfake.Responder[armeventhub.EventHubsClientDeleteAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EventHubsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *armeventhub.EventHubsClientGetOptions) (resp azfake.Responder[armeventhub.EventHubsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAuthorizationRule is the fake for method EventHubsClient.GetAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	GetAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *armeventhub.EventHubsClientGetAuthorizationRuleOptions) (resp azfake.Responder[armeventhub.EventHubsClientGetAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// NewListAuthorizationRulesPager is the fake for method EventHubsClient.NewListAuthorizationRulesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAuthorizationRulesPager func(resourceGroupName string, namespaceName string, eventHubName string, options *armeventhub.EventHubsClientListAuthorizationRulesOptions) (resp azfake.PagerResponder[armeventhub.EventHubsClientListAuthorizationRulesResponse])

	// NewListByNamespacePager is the fake for method EventHubsClient.NewListByNamespacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNamespacePager func(resourceGroupName string, namespaceName string, options *armeventhub.EventHubsClientListByNamespaceOptions) (resp azfake.PagerResponder[armeventhub.EventHubsClientListByNamespaceResponse])

	// ListKeys is the fake for method EventHubsClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, options *armeventhub.EventHubsClientListKeysOptions) (resp azfake.Responder[armeventhub.EventHubsClientListKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKeys is the fake for method EventHubsClient.RegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKeys func(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, authorizationRuleName string, parameters armeventhub.RegenerateAccessKeyParameters, options *armeventhub.EventHubsClientRegenerateKeysOptions) (resp azfake.Responder[armeventhub.EventHubsClientRegenerateKeysResponse], errResp azfake.ErrorResponder)
}

// NewEventHubsServerTransport creates a new instance of EventHubsServerTransport with the provided implementation.
// The returned EventHubsServerTransport instance is connected to an instance of armeventhub.EventHubsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEventHubsServerTransport(srv *EventHubsServer) *EventHubsServerTransport {
	return &EventHubsServerTransport{
		srv:                            srv,
		newListAuthorizationRulesPager: newTracker[azfake.PagerResponder[armeventhub.EventHubsClientListAuthorizationRulesResponse]](),
		newListByNamespacePager:        newTracker[azfake.PagerResponder[armeventhub.EventHubsClientListByNamespaceResponse]](),
	}
}

// EventHubsServerTransport connects instances of armeventhub.EventHubsClient to instances of EventHubsServer.
// Don't use this type directly, use NewEventHubsServerTransport instead.
type EventHubsServerTransport struct {
	srv                            *EventHubsServer
	newListAuthorizationRulesPager *tracker[azfake.PagerResponder[armeventhub.EventHubsClientListAuthorizationRulesResponse]]
	newListByNamespacePager        *tracker[azfake.PagerResponder[armeventhub.EventHubsClientListByNamespaceResponse]]
}

// Do implements the policy.Transporter interface for EventHubsServerTransport.
func (e *EventHubsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EventHubsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if eventHubsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = eventHubsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EventHubsClient.CreateOrUpdate":
				res.resp, res.err = e.dispatchCreateOrUpdate(req)
			case "EventHubsClient.CreateOrUpdateAuthorizationRule":
				res.resp, res.err = e.dispatchCreateOrUpdateAuthorizationRule(req)
			case "EventHubsClient.Delete":
				res.resp, res.err = e.dispatchDelete(req)
			case "EventHubsClient.DeleteAuthorizationRule":
				res.resp, res.err = e.dispatchDeleteAuthorizationRule(req)
			case "EventHubsClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "EventHubsClient.GetAuthorizationRule":
				res.resp, res.err = e.dispatchGetAuthorizationRule(req)
			case "EventHubsClient.NewListAuthorizationRulesPager":
				res.resp, res.err = e.dispatchNewListAuthorizationRulesPager(req)
			case "EventHubsClient.NewListByNamespacePager":
				res.resp, res.err = e.dispatchNewListByNamespacePager(req)
			case "EventHubsClient.ListKeys":
				res.resp, res.err = e.dispatchListKeys(req)
			case "EventHubsClient.RegenerateKeys":
				res.resp, res.err = e.dispatchRegenerateKeys(req)
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

func (e *EventHubsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventhub.Eventhub](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Eventhub, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchCreateOrUpdateAuthorizationRule(req *http.Request) (*http.Response, error) {
	if e.srv.CreateOrUpdateAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventhub.AuthorizationRule](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.CreateOrUpdateAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AuthorizationRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if e.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, nil)
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

func (e *EventHubsServerTransport) dispatchDeleteAuthorizationRule(req *http.Request) (*http.Response, error) {
	if e.srv.DeleteAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.DeleteAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, authorizationRuleNameParam, nil)
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

func (e *EventHubsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Eventhub, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchGetAuthorizationRule(req *http.Request) (*http.Response, error) {
	if e.srv.GetAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.GetAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AuthorizationRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchNewListAuthorizationRulesPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListAuthorizationRulesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAuthorizationRulesPager not implemented")}
	}
	newListAuthorizationRulesPager := e.newListAuthorizationRulesPager.get(req)
	if newListAuthorizationRulesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListAuthorizationRulesPager(resourceGroupNameParam, namespaceNameParam, eventHubNameParam, nil)
		newListAuthorizationRulesPager = &resp
		e.newListAuthorizationRulesPager.add(req, newListAuthorizationRulesPager)
		server.PagerResponderInjectNextLinks(newListAuthorizationRulesPager, req, func(page *armeventhub.EventHubsClientListAuthorizationRulesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAuthorizationRulesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListAuthorizationRulesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAuthorizationRulesPager) {
		e.newListAuthorizationRulesPager.remove(req)
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchNewListByNamespacePager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByNamespacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNamespacePager not implemented")}
	}
	newListByNamespacePager := e.newListByNamespacePager.get(req)
	if newListByNamespacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs`
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
		namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
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
		var options *armeventhub.EventHubsClientListByNamespaceOptions
		if skipParam != nil || topParam != nil {
			options = &armeventhub.EventHubsClientListByNamespaceOptions{
				Skip: skipParam,
				Top:  topParam,
			}
		}
		resp := e.srv.NewListByNamespacePager(resourceGroupNameParam, namespaceNameParam, options)
		newListByNamespacePager = &resp
		e.newListByNamespacePager.add(req, newListByNamespacePager)
		server.PagerResponderInjectNextLinks(newListByNamespacePager, req, func(page *armeventhub.EventHubsClientListByNamespaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNamespacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByNamespacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNamespacePager) {
		e.newListByNamespacePager.remove(req)
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if e.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.ListKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EventHubsServerTransport) dispatchRegenerateKeys(req *http.Request) (*http.Response, error) {
	if e.srv.RegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.EventHub/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eventhubs/(?P<eventHubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armeventhub.RegenerateAccessKeyParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	namespaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("namespaceName")])
	if err != nil {
		return nil, err
	}
	eventHubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("eventHubName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.RegenerateKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, eventHubNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccessKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to EventHubsServerTransport
var eventHubsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
