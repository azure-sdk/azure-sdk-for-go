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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/customerinsights/armcustomerinsights"
	"net/http"
	"net/url"
	"regexp"
)

// InteractionsServer is a fake server for instances of the armcustomerinsights.InteractionsClient type.
type InteractionsServer struct {
	// BeginCreateOrUpdate is the fake for method InteractionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, hubName string, interactionName string, parameters armcustomerinsights.InteractionResourceFormat, options *armcustomerinsights.InteractionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcustomerinsights.InteractionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InteractionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hubName string, interactionName string, options *armcustomerinsights.InteractionsClientGetOptions) (resp azfake.Responder[armcustomerinsights.InteractionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHubPager is the fake for method InteractionsClient.NewListByHubPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHubPager func(resourceGroupName string, hubName string, options *armcustomerinsights.InteractionsClientListByHubOptions) (resp azfake.PagerResponder[armcustomerinsights.InteractionsClientListByHubResponse])

	// SuggestRelationshipLinks is the fake for method InteractionsClient.SuggestRelationshipLinks
	// HTTP status codes to indicate success: http.StatusOK
	SuggestRelationshipLinks func(ctx context.Context, resourceGroupName string, hubName string, interactionName string, options *armcustomerinsights.InteractionsClientSuggestRelationshipLinksOptions) (resp azfake.Responder[armcustomerinsights.InteractionsClientSuggestRelationshipLinksResponse], errResp azfake.ErrorResponder)
}

// NewInteractionsServerTransport creates a new instance of InteractionsServerTransport with the provided implementation.
// The returned InteractionsServerTransport instance is connected to an instance of armcustomerinsights.InteractionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInteractionsServerTransport(srv *InteractionsServer) *InteractionsServerTransport {
	return &InteractionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armcustomerinsights.InteractionsClientCreateOrUpdateResponse]](),
		newListByHubPager:   newTracker[azfake.PagerResponder[armcustomerinsights.InteractionsClientListByHubResponse]](),
	}
}

// InteractionsServerTransport connects instances of armcustomerinsights.InteractionsClient to instances of InteractionsServer.
// Don't use this type directly, use NewInteractionsServerTransport instead.
type InteractionsServerTransport struct {
	srv                 *InteractionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armcustomerinsights.InteractionsClientCreateOrUpdateResponse]]
	newListByHubPager   *tracker[azfake.PagerResponder[armcustomerinsights.InteractionsClientListByHubResponse]]
}

// Do implements the policy.Transporter interface for InteractionsServerTransport.
func (i *InteractionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *InteractionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if interactionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = interactionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "InteractionsClient.BeginCreateOrUpdate":
				res.resp, res.err = i.dispatchBeginCreateOrUpdate(req)
			case "InteractionsClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "InteractionsClient.NewListByHubPager":
				res.resp, res.err = i.dispatchNewListByHubPager(req)
			case "InteractionsClient.SuggestRelationshipLinks":
				res.resp, res.err = i.dispatchSuggestRelationshipLinks(req)
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

func (i *InteractionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/interactions/(?P<interactionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcustomerinsights.InteractionResourceFormat](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		interactionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("interactionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, hubNameParam, interactionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *InteractionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/interactions/(?P<interactionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	interactionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("interactionName")])
	if err != nil {
		return nil, err
	}
	localeCodeUnescaped, err := url.QueryUnescape(qp.Get("locale-code"))
	if err != nil {
		return nil, err
	}
	localeCodeParam := getOptional(localeCodeUnescaped)
	var options *armcustomerinsights.InteractionsClientGetOptions
	if localeCodeParam != nil {
		options = &armcustomerinsights.InteractionsClientGetOptions{
			LocaleCode: localeCodeParam,
		}
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, hubNameParam, interactionNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InteractionResourceFormat, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InteractionsServerTransport) dispatchNewListByHubPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByHubPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHubPager not implemented")}
	}
	newListByHubPager := i.newListByHubPager.get(req)
	if newListByHubPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/interactions`
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
		hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
		if err != nil {
			return nil, err
		}
		localeCodeUnescaped, err := url.QueryUnescape(qp.Get("locale-code"))
		if err != nil {
			return nil, err
		}
		localeCodeParam := getOptional(localeCodeUnescaped)
		var options *armcustomerinsights.InteractionsClientListByHubOptions
		if localeCodeParam != nil {
			options = &armcustomerinsights.InteractionsClientListByHubOptions{
				LocaleCode: localeCodeParam,
			}
		}
		resp := i.srv.NewListByHubPager(resourceGroupNameParam, hubNameParam, options)
		newListByHubPager = &resp
		i.newListByHubPager.add(req, newListByHubPager)
		server.PagerResponderInjectNextLinks(newListByHubPager, req, func(page *armcustomerinsights.InteractionsClientListByHubResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHubPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByHubPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHubPager) {
		i.newListByHubPager.remove(req)
	}
	return resp, nil
}

func (i *InteractionsServerTransport) dispatchSuggestRelationshipLinks(req *http.Request) (*http.Response, error) {
	if i.srv.SuggestRelationshipLinks == nil {
		return nil, &nonRetriableError{errors.New("fake for method SuggestRelationshipLinks not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CustomerInsights/hubs/(?P<hubName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/interactions/(?P<interactionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/suggestRelationshipLinks`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hubNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hubName")])
	if err != nil {
		return nil, err
	}
	interactionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("interactionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.SuggestRelationshipLinks(req.Context(), resourceGroupNameParam, hubNameParam, interactionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SuggestRelationshipLinksResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to InteractionsServerTransport
var interactionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
