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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/orbital/armorbital/v2"
	"net/http"
	"net/url"
	"regexp"
)

// SpacecraftsServer is a fake server for instances of the armorbital.SpacecraftsClient type.
type SpacecraftsServer struct {
	// BeginCreateOrUpdate is the fake for method SpacecraftsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, spacecraftName string, parameters armorbital.Spacecraft, options *armorbital.SpacecraftsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armorbital.SpacecraftsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SpacecraftsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, spacecraftName string, options *armorbital.SpacecraftsClientBeginDeleteOptions) (resp azfake.PollerResponder[armorbital.SpacecraftsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SpacecraftsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, spacecraftName string, options *armorbital.SpacecraftsClientGetOptions) (resp azfake.Responder[armorbital.SpacecraftsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SpacecraftsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armorbital.SpacecraftsClientListOptions) (resp azfake.PagerResponder[armorbital.SpacecraftsClientListResponse])

	// BeginListAvailableContacts is the fake for method SpacecraftsClient.BeginListAvailableContacts
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginListAvailableContacts func(ctx context.Context, resourceGroupName string, spacecraftName string, parameters armorbital.ContactParameters, options *armorbital.SpacecraftsClientBeginListAvailableContactsOptions) (resp azfake.PollerResponder[azfake.PagerResponder[armorbital.SpacecraftsClientListAvailableContactsResponse]], errResp azfake.ErrorResponder)

	// NewListBySubscriptionPager is the fake for method SpacecraftsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armorbital.SpacecraftsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armorbital.SpacecraftsClientListBySubscriptionResponse])

	// BeginUpdateTags is the fake for method SpacecraftsClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, spacecraftName string, parameters armorbital.TagsObject, options *armorbital.SpacecraftsClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armorbital.SpacecraftsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewSpacecraftsServerTransport creates a new instance of SpacecraftsServerTransport with the provided implementation.
// The returned SpacecraftsServerTransport instance is connected to an instance of armorbital.SpacecraftsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSpacecraftsServerTransport(srv *SpacecraftsServer) *SpacecraftsServerTransport {
	return &SpacecraftsServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armorbital.SpacecraftsClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armorbital.SpacecraftsClientDeleteResponse]](),
		newListPager:               newTracker[azfake.PagerResponder[armorbital.SpacecraftsClientListResponse]](),
		beginListAvailableContacts: newTracker[azfake.PollerResponder[azfake.PagerResponder[armorbital.SpacecraftsClientListAvailableContactsResponse]]](),
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armorbital.SpacecraftsClientListBySubscriptionResponse]](),
		beginUpdateTags:            newTracker[azfake.PollerResponder[armorbital.SpacecraftsClientUpdateTagsResponse]](),
	}
}

// SpacecraftsServerTransport connects instances of armorbital.SpacecraftsClient to instances of SpacecraftsServer.
// Don't use this type directly, use NewSpacecraftsServerTransport instead.
type SpacecraftsServerTransport struct {
	srv                        *SpacecraftsServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armorbital.SpacecraftsClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armorbital.SpacecraftsClientDeleteResponse]]
	newListPager               *tracker[azfake.PagerResponder[armorbital.SpacecraftsClientListResponse]]
	beginListAvailableContacts *tracker[azfake.PollerResponder[azfake.PagerResponder[armorbital.SpacecraftsClientListAvailableContactsResponse]]]
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armorbital.SpacecraftsClientListBySubscriptionResponse]]
	beginUpdateTags            *tracker[azfake.PollerResponder[armorbital.SpacecraftsClientUpdateTagsResponse]]
}

// Do implements the policy.Transporter interface for SpacecraftsServerTransport.
func (s *SpacecraftsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SpacecraftsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if spacecraftsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = spacecraftsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SpacecraftsClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "SpacecraftsClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "SpacecraftsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SpacecraftsClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
			case "SpacecraftsClient.BeginListAvailableContacts":
				res.resp, res.err = s.dispatchBeginListAvailableContacts(req)
			case "SpacecraftsClient.NewListBySubscriptionPager":
				res.resp, res.err = s.dispatchNewListBySubscriptionPager(req)
			case "SpacecraftsClient.BeginUpdateTags":
				res.resp, res.err = s.dispatchBeginUpdateTags(req)
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

func (s *SpacecraftsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts/(?P<spacecraftName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armorbital.Spacecraft](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		spacecraftNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spacecraftName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, spacecraftNameParam, body, nil)
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

func (s *SpacecraftsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts/(?P<spacecraftName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		spacecraftNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spacecraftName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, spacecraftNameParam, nil)
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

func (s *SpacecraftsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts/(?P<spacecraftName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	spacecraftNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spacecraftName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, spacecraftNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Spacecraft, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SpacecraftsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts`
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
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam := getOptional(skiptokenUnescaped)
		var options *armorbital.SpacecraftsClientListOptions
		if skiptokenParam != nil {
			options = &armorbital.SpacecraftsClientListOptions{
				Skiptoken: skiptokenParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armorbital.SpacecraftsClientListResponse, createLink func() string) {
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

func (s *SpacecraftsServerTransport) dispatchBeginListAvailableContacts(req *http.Request) (*http.Response, error) {
	if s.srv.BeginListAvailableContacts == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginListAvailableContacts not implemented")}
	}
	beginListAvailableContacts := s.beginListAvailableContacts.get(req)
	if beginListAvailableContacts == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts/(?P<spacecraftName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listAvailableContacts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armorbital.ContactParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		spacecraftNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spacecraftName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginListAvailableContacts(req.Context(), resourceGroupNameParam, spacecraftNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginListAvailableContacts = &respr
		s.beginListAvailableContacts.add(req, beginListAvailableContacts)
	}

	resp, err := server.PollerResponderNext(beginListAvailableContacts, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginListAvailableContacts.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginListAvailableContacts) {
		s.beginListAvailableContacts.remove(req)
	}

	return resp, nil
}

func (s *SpacecraftsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := s.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam := getOptional(skiptokenUnescaped)
		var options *armorbital.SpacecraftsClientListBySubscriptionOptions
		if skiptokenParam != nil {
			options = &armorbital.SpacecraftsClientListBySubscriptionOptions{
				Skiptoken: skiptokenParam,
			}
		}
		resp := s.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		s.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armorbital.SpacecraftsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		s.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (s *SpacecraftsServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	beginUpdateTags := s.beginUpdateTags.get(req)
	if beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/spacecrafts/(?P<spacecraftName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armorbital.TagsObject](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		spacecraftNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("spacecraftName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdateTags(req.Context(), resourceGroupNameParam, spacecraftNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateTags = &respr
		s.beginUpdateTags.add(req, beginUpdateTags)
	}

	resp, err := server.PollerResponderNext(beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdateTags.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateTags) {
		s.beginUpdateTags.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to SpacecraftsServerTransport
var spacecraftsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
