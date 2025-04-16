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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/notificationhubs/armnotificationhubs/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NamespacesServer is a fake server for instances of the armnotificationhubs.NamespacesClient type.
type NamespacesServer struct {
	// CheckAvailability is the fake for method NamespacesClient.CheckAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckAvailability func(ctx context.Context, parameters armnotificationhubs.CheckAvailabilityParameters, options *armnotificationhubs.NamespacesClientCheckAvailabilityOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientCheckAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method NamespacesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, namespaceName string, parameters armnotificationhubs.NamespaceResource, options *armnotificationhubs.NamespacesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnotificationhubs.NamespacesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdateAuthorizationRule is the fake for method NamespacesClient.CreateOrUpdateAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdateAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters armnotificationhubs.SharedAccessAuthorizationRuleResource, options *armnotificationhubs.NamespacesClientCreateOrUpdateAuthorizationRuleOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientCreateOrUpdateAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method NamespacesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, namespaceName string, options *armnotificationhubs.NamespacesClientDeleteOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteAuthorizationRule is the fake for method NamespacesClient.DeleteAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	DeleteAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, options *armnotificationhubs.NamespacesClientDeleteAuthorizationRuleOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientDeleteAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NamespacesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, namespaceName string, options *armnotificationhubs.NamespacesClientGetOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientGetResponse], errResp azfake.ErrorResponder)

	// GetAuthorizationRule is the fake for method NamespacesClient.GetAuthorizationRule
	// HTTP status codes to indicate success: http.StatusOK
	GetAuthorizationRule func(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, options *armnotificationhubs.NamespacesClientGetAuthorizationRuleOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientGetAuthorizationRuleResponse], errResp azfake.ErrorResponder)

	// GetPnsCredentials is the fake for method NamespacesClient.GetPnsCredentials
	// HTTP status codes to indicate success: http.StatusOK
	GetPnsCredentials func(ctx context.Context, resourceGroupName string, namespaceName string, options *armnotificationhubs.NamespacesClientGetPnsCredentialsOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientGetPnsCredentialsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NamespacesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armnotificationhubs.NamespacesClientListOptions) (resp azfake.PagerResponder[armnotificationhubs.NamespacesClientListResponse])

	// NewListAllPager is the fake for method NamespacesClient.NewListAllPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAllPager func(options *armnotificationhubs.NamespacesClientListAllOptions) (resp azfake.PagerResponder[armnotificationhubs.NamespacesClientListAllResponse])

	// NewListAuthorizationRulesPager is the fake for method NamespacesClient.NewListAuthorizationRulesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListAuthorizationRulesPager func(resourceGroupName string, namespaceName string, options *armnotificationhubs.NamespacesClientListAuthorizationRulesOptions) (resp azfake.PagerResponder[armnotificationhubs.NamespacesClientListAuthorizationRulesResponse])

	// ListKeys is the fake for method NamespacesClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, options *armnotificationhubs.NamespacesClientListKeysOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientListKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKeys is the fake for method NamespacesClient.RegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKeys func(ctx context.Context, resourceGroupName string, namespaceName string, authorizationRuleName string, parameters armnotificationhubs.PolicyKeyResource, options *armnotificationhubs.NamespacesClientRegenerateKeysOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientRegenerateKeysResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method NamespacesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, namespaceName string, parameters armnotificationhubs.NamespacePatchParameters, options *armnotificationhubs.NamespacesClientUpdateOptions) (resp azfake.Responder[armnotificationhubs.NamespacesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewNamespacesServerTransport creates a new instance of NamespacesServerTransport with the provided implementation.
// The returned NamespacesServerTransport instance is connected to an instance of armnotificationhubs.NamespacesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNamespacesServerTransport(srv *NamespacesServer) *NamespacesServerTransport {
	return &NamespacesServerTransport{
		srv:                            srv,
		beginCreateOrUpdate:            newTracker[azfake.PollerResponder[armnotificationhubs.NamespacesClientCreateOrUpdateResponse]](),
		newListPager:                   newTracker[azfake.PagerResponder[armnotificationhubs.NamespacesClientListResponse]](),
		newListAllPager:                newTracker[azfake.PagerResponder[armnotificationhubs.NamespacesClientListAllResponse]](),
		newListAuthorizationRulesPager: newTracker[azfake.PagerResponder[armnotificationhubs.NamespacesClientListAuthorizationRulesResponse]](),
	}
}

// NamespacesServerTransport connects instances of armnotificationhubs.NamespacesClient to instances of NamespacesServer.
// Don't use this type directly, use NewNamespacesServerTransport instead.
type NamespacesServerTransport struct {
	srv                            *NamespacesServer
	beginCreateOrUpdate            *tracker[azfake.PollerResponder[armnotificationhubs.NamespacesClientCreateOrUpdateResponse]]
	newListPager                   *tracker[azfake.PagerResponder[armnotificationhubs.NamespacesClientListResponse]]
	newListAllPager                *tracker[azfake.PagerResponder[armnotificationhubs.NamespacesClientListAllResponse]]
	newListAuthorizationRulesPager *tracker[azfake.PagerResponder[armnotificationhubs.NamespacesClientListAuthorizationRulesResponse]]
}

// Do implements the policy.Transporter interface for NamespacesServerTransport.
func (n *NamespacesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return n.dispatchToMethodFake(req, method)
}

func (n *NamespacesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if namespacesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = namespacesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "NamespacesClient.CheckAvailability":
				res.resp, res.err = n.dispatchCheckAvailability(req)
			case "NamespacesClient.BeginCreateOrUpdate":
				res.resp, res.err = n.dispatchBeginCreateOrUpdate(req)
			case "NamespacesClient.CreateOrUpdateAuthorizationRule":
				res.resp, res.err = n.dispatchCreateOrUpdateAuthorizationRule(req)
			case "NamespacesClient.Delete":
				res.resp, res.err = n.dispatchDelete(req)
			case "NamespacesClient.DeleteAuthorizationRule":
				res.resp, res.err = n.dispatchDeleteAuthorizationRule(req)
			case "NamespacesClient.Get":
				res.resp, res.err = n.dispatchGet(req)
			case "NamespacesClient.GetAuthorizationRule":
				res.resp, res.err = n.dispatchGetAuthorizationRule(req)
			case "NamespacesClient.GetPnsCredentials":
				res.resp, res.err = n.dispatchGetPnsCredentials(req)
			case "NamespacesClient.NewListPager":
				res.resp, res.err = n.dispatchNewListPager(req)
			case "NamespacesClient.NewListAllPager":
				res.resp, res.err = n.dispatchNewListAllPager(req)
			case "NamespacesClient.NewListAuthorizationRulesPager":
				res.resp, res.err = n.dispatchNewListAuthorizationRulesPager(req)
			case "NamespacesClient.ListKeys":
				res.resp, res.err = n.dispatchListKeys(req)
			case "NamespacesClient.RegenerateKeys":
				res.resp, res.err = n.dispatchRegenerateKeys(req)
			case "NamespacesClient.Update":
				res.resp, res.err = n.dispatchUpdate(req)
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

func (n *NamespacesServerTransport) dispatchCheckAvailability(req *http.Request) (*http.Response, error) {
	if n.srv.CheckAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/checkNamespaceAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnotificationhubs.CheckAvailabilityParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.CheckAvailability(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CheckAvailabilityResult, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := n.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnotificationhubs.NamespaceResource](req)
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
		respr, errRespr := n.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, namespaceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		n.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		n.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		n.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (n *NamespacesServerTransport) dispatchCreateOrUpdateAuthorizationRule(req *http.Request) (*http.Response, error) {
	if n.srv.CreateOrUpdateAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdateAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnotificationhubs.SharedAccessAuthorizationRuleResource](req)
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
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.CreateOrUpdateAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SharedAccessAuthorizationRuleResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if n.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := n.srv.Delete(req.Context(), resourceGroupNameParam, namespaceNameParam, nil)
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

func (n *NamespacesServerTransport) dispatchDeleteAuthorizationRule(req *http.Request) (*http.Response, error) {
	if n.srv.DeleteAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.DeleteAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, authorizationRuleNameParam, nil)
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

func (n *NamespacesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, namespaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NamespaceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchGetAuthorizationRule(req *http.Request) (*http.Response, error) {
	if n.srv.GetAuthorizationRule == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetAuthorizationRule not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.GetAuthorizationRule(req.Context(), resourceGroupNameParam, namespaceNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SharedAccessAuthorizationRuleResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchGetPnsCredentials(req *http.Request) (*http.Response, error) {
	if n.srv.GetPnsCredentials == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPnsCredentials not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pnsCredentials`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
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
	respr, errRespr := n.srv.GetPnsCredentials(req.Context(), resourceGroupNameParam, namespaceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PnsCredentialsResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := n.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
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
		var options *armnotificationhubs.NamespacesClientListOptions
		if skipTokenParam != nil || topParam != nil {
			options = &armnotificationhubs.NamespacesClientListOptions{
				SkipToken: skipTokenParam,
				Top:       topParam,
			}
		}
		resp := n.srv.NewListPager(resourceGroupNameParam, options)
		newListPager = &resp
		n.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnotificationhubs.NamespacesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		n.newListPager.remove(req)
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchNewListAllPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListAllPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAllPager not implemented")}
	}
	newListAllPager := n.newListAllPager.get(req)
	if newListAllPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
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
		var options *armnotificationhubs.NamespacesClientListAllOptions
		if skipTokenParam != nil || topParam != nil {
			options = &armnotificationhubs.NamespacesClientListAllOptions{
				SkipToken: skipTokenParam,
				Top:       topParam,
			}
		}
		resp := n.srv.NewListAllPager(options)
		newListAllPager = &resp
		n.newListAllPager.add(req, newListAllPager)
		server.PagerResponderInjectNextLinks(newListAllPager, req, func(page *armnotificationhubs.NamespacesClientListAllResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAllPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListAllPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAllPager) {
		n.newListAllPager.remove(req)
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchNewListAuthorizationRulesPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListAuthorizationRulesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListAuthorizationRulesPager not implemented")}
	}
	newListAuthorizationRulesPager := n.newListAuthorizationRulesPager.get(req)
	if newListAuthorizationRulesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
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
		resp := n.srv.NewListAuthorizationRulesPager(resourceGroupNameParam, namespaceNameParam, nil)
		newListAuthorizationRulesPager = &resp
		n.newListAuthorizationRulesPager.add(req, newListAuthorizationRulesPager)
		server.PagerResponderInjectNextLinks(newListAuthorizationRulesPager, req, func(page *armnotificationhubs.NamespacesClientListAuthorizationRulesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListAuthorizationRulesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		n.newListAuthorizationRulesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListAuthorizationRulesPager) {
		n.newListAuthorizationRulesPager.remove(req)
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if n.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
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
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.ListKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, authorizationRuleNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceListKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchRegenerateKeys(req *http.Request) (*http.Response, error) {
	if n.srv.RegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/authorizationRules/(?P<authorizationRuleName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKeys`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnotificationhubs.PolicyKeyResource](req)
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
	authorizationRuleNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("authorizationRuleName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.RegenerateKeys(req.Context(), resourceGroupNameParam, namespaceNameParam, authorizationRuleNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceListKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NamespacesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NotificationHubs/namespaces/(?P<namespaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnotificationhubs.NamespacePatchParameters](req)
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
	respr, errRespr := n.srv.Update(req.Context(), resourceGroupNameParam, namespaceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NamespaceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to NamespacesServerTransport
var namespacesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
