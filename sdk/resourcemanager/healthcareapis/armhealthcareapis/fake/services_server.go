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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ServicesServer is a fake server for instances of the armhealthcareapis.ServicesClient type.
type ServicesServer struct {
	// CheckNameAvailability is the fake for method ServicesClient.CheckNameAvailability
	// HTTP status codes to indicate success: http.StatusOK
	CheckNameAvailability func(ctx context.Context, checkNameAvailabilityInputs armhealthcareapis.CheckNameAvailabilityParameters, options *armhealthcareapis.ServicesClientCheckNameAvailabilityOptions) (resp azfake.Responder[armhealthcareapis.ServicesClientCheckNameAvailabilityResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method ServicesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, serviceDescription armhealthcareapis.ServicesDescription, options *armhealthcareapis.ServicesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armhealthcareapis.ServicesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ServicesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, options *armhealthcareapis.ServicesClientBeginDeleteOptions) (resp azfake.PollerResponder[armhealthcareapis.ServicesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServicesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, options *armhealthcareapis.ServicesClientGetOptions) (resp azfake.Responder[armhealthcareapis.ServicesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServicesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armhealthcareapis.ServicesClientListOptions) (resp azfake.PagerResponder[armhealthcareapis.ServicesClientListResponse])

	// NewListByResourceGroupPager is the fake for method ServicesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armhealthcareapis.ServicesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armhealthcareapis.ServicesClientListByResourceGroupResponse])

	// BeginUpdate is the fake for method ServicesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, resourceName string, servicePatchDescription armhealthcareapis.ServicesPatchDescription, options *armhealthcareapis.ServicesClientBeginUpdateOptions) (resp azfake.PollerResponder[armhealthcareapis.ServicesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServicesServerTransport creates a new instance of ServicesServerTransport with the provided implementation.
// The returned ServicesServerTransport instance is connected to an instance of armhealthcareapis.ServicesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServicesServerTransport(srv *ServicesServer) *ServicesServerTransport {
	return &ServicesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armhealthcareapis.ServicesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armhealthcareapis.ServicesClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armhealthcareapis.ServicesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armhealthcareapis.ServicesClientListByResourceGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armhealthcareapis.ServicesClientUpdateResponse]](),
	}
}

// ServicesServerTransport connects instances of armhealthcareapis.ServicesClient to instances of ServicesServer.
// Don't use this type directly, use NewServicesServerTransport instead.
type ServicesServerTransport struct {
	srv                         *ServicesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armhealthcareapis.ServicesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armhealthcareapis.ServicesClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armhealthcareapis.ServicesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armhealthcareapis.ServicesClientListByResourceGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armhealthcareapis.ServicesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ServicesServerTransport.
func (s *ServicesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServicesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if servicesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = servicesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServicesClient.CheckNameAvailability":
				res.resp, res.err = s.dispatchCheckNameAvailability(req)
			case "ServicesClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "ServicesClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "ServicesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServicesClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
			case "ServicesClient.NewListByResourceGroupPager":
				res.resp, res.err = s.dispatchNewListByResourceGroupPager(req)
			case "ServicesClient.BeginUpdate":
				res.resp, res.err = s.dispatchBeginUpdate(req)
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

func (s *ServicesServerTransport) dispatchCheckNameAvailability(req *http.Request) (*http.Response, error) {
	if s.srv.CheckNameAvailability == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckNameAvailability not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/checkNameAvailability`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armhealthcareapis.CheckNameAvailabilityParameters](req)
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CheckNameAvailability(req.Context(), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServicesNameAvailabilityInfo, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServicesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/services/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhealthcareapis.ServicesDescription](req)
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
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *ServicesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/services/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
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

func (s *ServicesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/services/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServicesDescription, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServicesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/services`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListPager(nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armhealthcareapis.ServicesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServicesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/services`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armhealthcareapis.ServicesClientListByResourceGroupResponse, createLink func() string) {
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

func (s *ServicesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HealthcareApis/services/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armhealthcareapis.ServicesPatchDescription](req)
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
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ServicesServerTransport
var servicesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
