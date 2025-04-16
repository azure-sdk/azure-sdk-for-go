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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// TopicsServer is a fake server for instances of the armservicebus.TopicsClient type.
type TopicsServer struct {
	// CreateOrUpdate is the fake for method TopicsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, parameters armservicebus.SBTopic, options *armservicebus.TopicsClientCreateOrUpdateOptions) (resp azfake.Responder[armservicebus.TopicsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAuthorizationRule is the fake for method TopicsClient.CreateOrUpdateAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdateAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters armservicebus.SBAuthorizationRule, options *armservicebus.TopicsClientCreateOrUpdateAuthorizationRuleOptions) (resp azfake.Responder[armservicebus.TopicsClientCreateOrUpdateAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method TopicsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *armservicebus.TopicsClientDeleteOptions) (resp azfake.Responder[armservicebus.TopicsClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAuthorizationRule is the fake for method TopicsClient.DeleteAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *armservicebus.TopicsClientDeleteAuthorizationRuleOptions) (resp azfake.Responder[armservicebus.TopicsClientDeleteAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method TopicsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, options *armservicebus.TopicsClientGetOptions) (resp azfake.Responder[armservicebus.TopicsClientGetResponse], errResp azfake.ErrorResponder)

	// GetAuthorizationRule is the fake for method TopicsClient.GetAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	GetAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *armservicebus.TopicsClientGetAuthorizationRuleOptions) (resp azfake.Responder[armservicebus.TopicsClientGetAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// NewListAuthorizationRulesPager is the fake for method TopicsClient.NewListAuthorizationRulesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAuthorizationRulesPager func(resourceGroupName string, namespaceName string, topicName string, options *armservicebus.TopicsClientListAuthorizationRulesOptions) (resp azfake.PagerResponder[armservicebus.TopicsClientListAuthorizationRulesResponse])

	// NewListByNamespacePager is the fake for method TopicsClient.NewListByNamespacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByNamespacePager func(resourceGroupName string, namespaceName string, options *armservicebus.TopicsClientListByNamespaceOptions) (resp azfake.PagerResponder[armservicebus.TopicsClientListByNamespaceResponse])

	// ListKeys is the fake for method TopicsClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, options *armservicebus.TopicsClientListKeysOptions) (resp azfake.Responder[armservicebus.TopicsClientListKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKeys is the fake for method TopicsClient.RegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKeys func(ctx context.Context, resourceGroupName string, namespaceName string, topicName string, authorizationRuleName string, parameters armservicebus.RegenerateAccessKeyParameters, options *armservicebus.TopicsClientRegenerateKeysOptions) (resp azfake.Responder[armservicebus.TopicsClientRegenerateKeysResponse], errResp azfake.ErrorResponder)
}

// NewTopicsServerTransport creates a new instance of TopicsServerTransport with the provided implementation.
// The returned TopicsServerTransport instance is connected to an instance of armservicebus.TopicsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTopicsServerTransport(srv *TopicsServer) *TopicsServerTransport {
	return &TopicsServerTransport{
		srv:                            srv,
		newListAuthorizationRulesPager: newTracker[azfake.PagerResponder[armservicebus.TopicsClientListAuthorizationRulesResponse]](),
		newListByNamespacePager:        newTracker[azfake.PagerResponder[armservicebus.TopicsClientListByNamespaceResponse]](),
	}
}

// TopicsServerTransport connects instances of armservicebus.TopicsClient to instances of TopicsServer.
// Don't use this type directly, use NewTopicsServerTransport instead.
type TopicsServerTransport struct {
	srv                            *TopicsServer
	newListAuthorizationRulesPager *tracker[azfake.PagerResponder[armservicebus.TopicsClientListAuthorizationRulesResponse]]
	newListByNamespacePager        *tracker[azfake.PagerResponder[armservicebus.TopicsClientListByNamespaceResponse]]
}

// Do implements the policy.Transporter interface for TopicsServerTransport.
func (t *TopicsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return t.dispatchToMethodFake(req, method)
}

func (t *TopicsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if topicsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = topicsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "TopicsClient.CreateOrUpdate":
				res.resp, res.err = t.dispatchCreateOrUpdate(req)
			case "TopicsClient.CreateOrUpdateAuthorizationRule":
				res.resp, res.err = t.dispatchCreateOrUpdateAuthorizationRule(req)
			case "TopicsClient.Delete":
				res.resp, res.err = t.dispatchDelete(req)
			case "TopicsClient.DeleteAuthorizationRule":
				res.resp, res.err = t.dispatchDeleteAuthorizationRule(req)
			case "TopicsClient.Get":
				res.resp, res.err = t.dispatchGet(req)
			case "TopicsClient.GetAuthorizationRule":
				res.resp, res.err = t.dispatchGetAuthorizationRule(req)
			case "TopicsClient.NewListAuthorizationRulesPager":
				res.resp, res.err = t.dispatchNewListAuthorizationRulesPager(req)
			case "TopicsClient.NewListByNamespacePager":
				res.resp, res.err = t.dispatchNewListByNamespacePager(req)
			case "TopicsClient.ListKeys":
				res.resp, res.err = t.dispatchListKeys(req)
			case "TopicsClient.RegenerateKeys":
				res.resp, res.err = t.dispatchRegenerateKeys(req)
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

func (t *TopicsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicebus.SBTopic](req)
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SBTopic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TopicsServerTransport) dispatchCreateOrUpdateAuthorizationRule(req *http.Request) (*http.Response, error) {
	if t.srv.CreateOrUpdateAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicebus.SBAuthorizationRule](req)
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.CreateOrUpdateAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SBAuthorizationRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TopicsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if t.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, nil)
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

func (t *TopicsServerTransport) dispatchDeleteAuthorizationRule(req *http.Request) (*http.Response, error) {
	if t.srv.DeleteAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.DeleteAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, authorizationRuleNameParam, nil)
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

func (t *TopicsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SBTopic, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TopicsServerTransport) dispatchGetAuthorizationRule(req *http.Request) (*http.Response, error) {
	if t.srv.GetAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.GetAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SBAuthorizationRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TopicsServerTransport) dispatchNewListAuthorizationRulesPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListAuthorizationRulesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAuthorizationRulesPager not implemented")}
	}
	newListAuthorizationRulesPager := t.newListAuthorizationRulesPager.get(req)
	if newListAuthorizationRulesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules`
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
		topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
		if err != nil {
			return nil, err
		}
		resp := t.srv.NewListAuthorizationRulesPager(resourceGroupNameParam, namespaceNameParam, topicNameParam, nil)
		newListAuthorizationRulesPager = &resp
		t.newListAuthorizationRulesPager.add(req, newListAuthorizationRulesPager)
		server.PagerResponderInjectNextLinks(newListAuthorizationRulesPager, req, func(page *armservicebus.TopicsClientListAuthorizationRulesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAuthorizationRulesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListAuthorizationRulesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAuthorizationRulesPager) {
		t.newListAuthorizationRulesPager.remove(req)
	}
	return resp, nil
}

func (t *TopicsServerTransport) dispatchNewListByNamespacePager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListByNamespacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByNamespacePager not implemented")}
	}
	newListByNamespacePager := t.newListByNamespacePager.get(req)
	if newListByNamespacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics`
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
		var options *armservicebus.TopicsClientListByNamespaceOptions
		if skipParam != nil || topParam != nil {
			options = &armservicebus.TopicsClientListByNamespaceOptions{
				Skip: skipParam,
				Top:  topParam,
			}
		}
		resp := t.srv.NewListByNamespacePager(resourceGroupNameParam, namespaceNameParam, options)
		newListByNamespacePager = &resp
		t.newListByNamespacePager.add(req, newListByNamespacePager)
		server.PagerResponderInjectNextLinks(newListByNamespacePager, req, func(page *armservicebus.TopicsClientListByNamespaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByNamespacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListByNamespacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByNamespacePager) {
		t.newListByNamespacePager.remove(req)
	}
	return resp, nil
}

func (t *TopicsServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if t.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ListKeys`
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.ListKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, authorizationRuleNameParam, nil)
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

func (t *TopicsServerTransport) dispatchRegenerateKeys(req *http.Request) (*http.Response, error) {
	if t.srv.RegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceBus/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/topics/(?P<topicName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armservicebus.RegenerateAccessKeyParameters](req)
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
	topicNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("topicName")])
	if err != nil {
		return nil, err
	}
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.RegenerateKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, topicNameParam, authorizationRuleNameParam, body, nil)
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

// set this to conditionally intercept incoming requests to TopicsServerTransport
var topicsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
