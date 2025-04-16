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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/communication/armcommunication/v2"
	"net/http"
	"net/url"
	"regexp"
)

// EmailServicesServer is a fake server for instances of the armcommunication.EmailServicesClient type.
type EmailServicesServer struct {
	// BeginCreateOrUpdate is the fake for method EmailServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, emailServiceName string, parameters armcommunication.EmailServiceResource, options *armcommunication.EmailServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcommunication.EmailServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method EmailServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, emailServiceName string, options *armcommunication.EmailServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armcommunication.EmailServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EmailServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, emailServiceName string, options *armcommunication.EmailServicesClientGetOptions) (resp azfake.Responder[armcommunication.EmailServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method EmailServicesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcommunication.EmailServicesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcommunication.EmailServicesClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method EmailServicesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armcommunication.EmailServicesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armcommunication.EmailServicesClientListBySubscriptionResponse])

	// ListVerifiedExchangeOnlineDomains is the fake for method EmailServicesClient.ListVerifiedExchangeOnlineDomains
	// HTTP status codes to indicate success: http.StatusOK
	ListVerifiedExchangeOnlineDomains func(ctx context.Context, options *armcommunication.EmailServicesClientListVerifiedExchangeOnlineDomainsOptions) (resp azfake.Responder[armcommunication.EmailServicesClientListVerifiedExchangeOnlineDomainsResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method EmailServicesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdate func(ctx context.Context, resourceGroupName string, emailServiceName string, parameters armcommunication.EmailServiceResourceUpdate, options *armcommunication.EmailServicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armcommunication.EmailServicesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewEmailServicesServerTransport creates a new instance of EmailServicesServerTransport with the provided implementation.
// The returned EmailServicesServerTransport instance is connected to an instance of armcommunication.EmailServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEmailServicesServerTransport(srv *EmailServicesServer) *EmailServicesServerTransport {
	return &EmailServicesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armcommunication.EmailServicesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armcommunication.EmailServicesClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armcommunication.EmailServicesClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armcommunication.EmailServicesClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armcommunication.EmailServicesClientUpdateResponse]](),
	}
}

// EmailServicesServerTransport connects instances of armcommunication.EmailServicesClient to instances of EmailServicesServer.
// Don't use this type directly, use NewEmailServicesServerTransport instead.
type EmailServicesServerTransport struct {
	srv                         *EmailServicesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armcommunication.EmailServicesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armcommunication.EmailServicesClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armcommunication.EmailServicesClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armcommunication.EmailServicesClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armcommunication.EmailServicesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for EmailServicesServerTransport.
func (e *EmailServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EmailServicesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if emailServicesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = emailServicesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EmailServicesClient.BeginCreateOrUpdate":
				res.resp, res.err = e.dispatchBeginCreateOrUpdate(req)
			case "EmailServicesClient.BeginDelete":
				res.resp, res.err = e.dispatchBeginDelete(req)
			case "EmailServicesClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "EmailServicesClient.NewListByResourceGroupPager":
				res.resp, res.err = e.dispatchNewListByResourceGroupPager(req)
			case "EmailServicesClient.NewListBySubscriptionPager":
				res.resp, res.err = e.dispatchNewListBySubscriptionPager(req)
			case "EmailServicesClient.ListVerifiedExchangeOnlineDomains":
				res.resp, res.err = e.dispatchListVerifiedExchangeOnlineDomains(req)
			case "EmailServicesClient.BeginUpdate":
				res.resp, res.err = e.dispatchBeginUpdate(req)
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

func (e *EmailServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcommunication.EmailServiceResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, emailServiceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *EmailServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, emailServiceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		e.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		e.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		e.beginDelete.remove(req)
	}

	return resp, nil
}

func (e *EmailServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, emailServiceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EmailServiceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EmailServicesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := e.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		e.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armcommunication.EmailServicesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		e.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (e *EmailServicesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := e.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := e.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		e.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armcommunication.EmailServicesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		e.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (e *EmailServicesServerTransport) dispatchListVerifiedExchangeOnlineDomains(req *http.Request) (*http.Response, error) {
	if e.srv.ListVerifiedExchangeOnlineDomains == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListVerifiedExchangeOnlineDomains not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/listVerifiedExchangeOnlineDomains`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	respr, errRespr := e.srv.ListVerifiedExchangeOnlineDomains(req.Context(), nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StringArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EmailServicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := e.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Communication/emailServices/(?P<emailServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcommunication.EmailServiceResourceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		emailServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("emailServiceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdate(req.Context(), resourceGroupNameParam, emailServiceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		e.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		e.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to EmailServicesServerTransport
var emailServicesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
