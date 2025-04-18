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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/professionalservice/armprofessionalservice"
	"net/http"
	"net/url"
	"regexp"
)

// SubscriptionLevelServer is a fake server for instances of the armprofessionalservice.SubscriptionLevelClient type.
type SubscriptionLevelServer struct {
	// BeginCreateOrUpdate is the fake for method SubscriptionLevelClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, subscriptionID string, resourceGroupName string, resourceName string, parameters armprofessionalservice.ResourceCreation, options *armprofessionalservice.SubscriptionLevelClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SubscriptionLevelClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, subscriptionID string, resourceGroupName string, resourceName string, options *armprofessionalservice.SubscriptionLevelClientBeginDeleteOptions) (resp azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SubscriptionLevelClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, subscriptionID string, resourceGroupName string, resourceName string, options *armprofessionalservice.SubscriptionLevelClientGetOptions) (resp azfake.Responder[armprofessionalservice.SubscriptionLevelClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByAzureSubscriptionPager is the fake for method SubscriptionLevelClient.NewListByAzureSubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByAzureSubscriptionPager func(subscriptionID string, options *armprofessionalservice.SubscriptionLevelClientListByAzureSubscriptionOptions) (resp azfake.PagerResponder[armprofessionalservice.SubscriptionLevelClientListByAzureSubscriptionResponse])

	// NewListByResourceGroupPager is the fake for method SubscriptionLevelClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(subscriptionID string, resourceGroupName string, options *armprofessionalservice.SubscriptionLevelClientListByResourceGroupOptions) (resp azfake.PagerResponder[armprofessionalservice.SubscriptionLevelClientListByResourceGroupResponse])

	// BeginUpdateToUnsubscribed is the fake for method SubscriptionLevelClient.BeginUpdateToUnsubscribed
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUpdateToUnsubscribed func(ctx context.Context, subscriptionID string, resourceGroupName string, resourceName string, parameters armprofessionalservice.DeleteOptions, options *armprofessionalservice.SubscriptionLevelClientBeginUpdateToUnsubscribedOptions) (resp azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientUpdateToUnsubscribedResponse], errResp azfake.ErrorResponder)
}

// NewSubscriptionLevelServerTransport creates a new instance of SubscriptionLevelServerTransport with the provided implementation.
// The returned SubscriptionLevelServerTransport instance is connected to an instance of armprofessionalservice.SubscriptionLevelClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSubscriptionLevelServerTransport(srv *SubscriptionLevelServer) *SubscriptionLevelServerTransport {
	return &SubscriptionLevelServerTransport{
		srv:                             srv,
		beginCreateOrUpdate:             newTracker[azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientCreateOrUpdateResponse]](),
		beginDelete:                     newTracker[azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientDeleteResponse]](),
		newListByAzureSubscriptionPager: newTracker[azfake.PagerResponder[armprofessionalservice.SubscriptionLevelClientListByAzureSubscriptionResponse]](),
		newListByResourceGroupPager:     newTracker[azfake.PagerResponder[armprofessionalservice.SubscriptionLevelClientListByResourceGroupResponse]](),
		beginUpdateToUnsubscribed:       newTracker[azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientUpdateToUnsubscribedResponse]](),
	}
}

// SubscriptionLevelServerTransport connects instances of armprofessionalservice.SubscriptionLevelClient to instances of SubscriptionLevelServer.
// Don't use this type directly, use NewSubscriptionLevelServerTransport instead.
type SubscriptionLevelServerTransport struct {
	srv                             *SubscriptionLevelServer
	beginCreateOrUpdate             *tracker[azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientCreateOrUpdateResponse]]
	beginDelete                     *tracker[azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientDeleteResponse]]
	newListByAzureSubscriptionPager *tracker[azfake.PagerResponder[armprofessionalservice.SubscriptionLevelClientListByAzureSubscriptionResponse]]
	newListByResourceGroupPager     *tracker[azfake.PagerResponder[armprofessionalservice.SubscriptionLevelClientListByResourceGroupResponse]]
	beginUpdateToUnsubscribed       *tracker[azfake.PollerResponder[armprofessionalservice.SubscriptionLevelClientUpdateToUnsubscribedResponse]]
}

// Do implements the policy.Transporter interface for SubscriptionLevelServerTransport.
func (s *SubscriptionLevelServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SubscriptionLevelServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if subscriptionLevelServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = subscriptionLevelServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SubscriptionLevelClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "SubscriptionLevelClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "SubscriptionLevelClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SubscriptionLevelClient.NewListByAzureSubscriptionPager":
				res.resp, res.err = s.dispatchNewListByAzureSubscriptionPager(req)
			case "SubscriptionLevelClient.NewListByResourceGroupPager":
				res.resp, res.err = s.dispatchNewListByResourceGroupPager(req)
			case "SubscriptionLevelClient.BeginUpdateToUnsubscribed":
				res.resp, res.err = s.dispatchBeginUpdateToUnsubscribed(req)
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

func (s *SubscriptionLevelServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProfessionalService/resources/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armprofessionalservice.ResourceCreation](req)
		if err != nil {
			return nil, err
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), subscriptionIDParam, resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *SubscriptionLevelServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProfessionalService/resources/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), subscriptionIDParam, resourceGroupNameParam, resourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SubscriptionLevelServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProfessionalService/resources/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), subscriptionIDParam, resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Resource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SubscriptionLevelServerTransport) dispatchNewListByAzureSubscriptionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByAzureSubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByAzureSubscriptionPager not implemented")}
	}
	newListByAzureSubscriptionPager := s.newListByAzureSubscriptionPager.get(req)
	if newListByAzureSubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProfessionalService/resources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByAzureSubscriptionPager(subscriptionIDParam, nil)
		newListByAzureSubscriptionPager = &resp
		s.newListByAzureSubscriptionPager.add(req, newListByAzureSubscriptionPager)
		server.PagerResponderInjectNextLinks(newListByAzureSubscriptionPager, req, func(page *armprofessionalservice.SubscriptionLevelClientListByAzureSubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByAzureSubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByAzureSubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByAzureSubscriptionPager) {
		s.newListByAzureSubscriptionPager.remove(req)
	}
	return resp, nil
}

func (s *SubscriptionLevelServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProfessionalService/resources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(subscriptionIDParam, resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armprofessionalservice.SubscriptionLevelClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		s.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (s *SubscriptionLevelServerTransport) dispatchBeginUpdateToUnsubscribed(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdateToUnsubscribed == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateToUnsubscribed not implemented")}
	}
	beginUpdateToUnsubscribed := s.beginUpdateToUnsubscribed.get(req)
	if beginUpdateToUnsubscribed == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ProfessionalService/resources/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/unsubscribe`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armprofessionalservice.DeleteOptions](req)
		if err != nil {
			return nil, err
		}
		subscriptionIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("subscriptionId")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdateToUnsubscribed(req.Context(), subscriptionIDParam, resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateToUnsubscribed = &respr
		s.beginUpdateToUnsubscribed.add(req, beginUpdateToUnsubscribed)
	}

	resp, err := server.PollerResponderNext(beginUpdateToUnsubscribed, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginUpdateToUnsubscribed.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateToUnsubscribed) {
		s.beginUpdateToUnsubscribed.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to SubscriptionLevelServerTransport
var subscriptionLevelServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
