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

// ContactProfilesServer is a fake server for instances of the armorbital.ContactProfilesClient type.
type ContactProfilesServer struct {
	// BeginCreateOrUpdate is the fake for method ContactProfilesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, contactProfileName string, parameters armorbital.ContactProfile, options *armorbital.ContactProfilesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armorbital.ContactProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ContactProfilesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, contactProfileName string, options *armorbital.ContactProfilesClientBeginDeleteOptions) (resp azfake.PollerResponder[armorbital.ContactProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ContactProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, contactProfileName string, options *armorbital.ContactProfilesClientGetOptions) (resp azfake.Responder[armorbital.ContactProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ContactProfilesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, options *armorbital.ContactProfilesClientListOptions) (resp azfake.PagerResponder[armorbital.ContactProfilesClientListResponse])

	// NewListBySubscriptionPager is the fake for method ContactProfilesClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armorbital.ContactProfilesClientListBySubscriptionOptions) (resp azfake.PagerResponder[armorbital.ContactProfilesClientListBySubscriptionResponse])

	// BeginUpdateTags is the fake for method ContactProfilesClient.BeginUpdateTags
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateTags func(ctx context.Context, resourceGroupName string, contactProfileName string, parameters armorbital.TagsObject, options *armorbital.ContactProfilesClientBeginUpdateTagsOptions) (resp azfake.PollerResponder[armorbital.ContactProfilesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewContactProfilesServerTransport creates a new instance of ContactProfilesServerTransport with the provided implementation.
// The returned ContactProfilesServerTransport instance is connected to an instance of armorbital.ContactProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContactProfilesServerTransport(srv *ContactProfilesServer) *ContactProfilesServerTransport {
	return &ContactProfilesServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armorbital.ContactProfilesClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armorbital.ContactProfilesClientDeleteResponse]](),
		newListPager:               newTracker[azfake.PagerResponder[armorbital.ContactProfilesClientListResponse]](),
		newListBySubscriptionPager: newTracker[azfake.PagerResponder[armorbital.ContactProfilesClientListBySubscriptionResponse]](),
		beginUpdateTags:            newTracker[azfake.PollerResponder[armorbital.ContactProfilesClientUpdateTagsResponse]](),
	}
}

// ContactProfilesServerTransport connects instances of armorbital.ContactProfilesClient to instances of ContactProfilesServer.
// Don't use this type directly, use NewContactProfilesServerTransport instead.
type ContactProfilesServerTransport struct {
	srv                        *ContactProfilesServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armorbital.ContactProfilesClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armorbital.ContactProfilesClientDeleteResponse]]
	newListPager               *tracker[azfake.PagerResponder[armorbital.ContactProfilesClientListResponse]]
	newListBySubscriptionPager *tracker[azfake.PagerResponder[armorbital.ContactProfilesClientListBySubscriptionResponse]]
	beginUpdateTags            *tracker[azfake.PollerResponder[armorbital.ContactProfilesClientUpdateTagsResponse]]
}

// Do implements the policy.Transporter interface for ContactProfilesServerTransport.
func (c *ContactProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ContactProfilesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if contactProfilesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = contactProfilesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ContactProfilesClient.BeginCreateOrUpdate":
				res.resp, res.err = c.dispatchBeginCreateOrUpdate(req)
			case "ContactProfilesClient.BeginDelete":
				res.resp, res.err = c.dispatchBeginDelete(req)
			case "ContactProfilesClient.Get":
				res.resp, res.err = c.dispatchGet(req)
			case "ContactProfilesClient.NewListPager":
				res.resp, res.err = c.dispatchNewListPager(req)
			case "ContactProfilesClient.NewListBySubscriptionPager":
				res.resp, res.err = c.dispatchNewListBySubscriptionPager(req)
			case "ContactProfilesClient.BeginUpdateTags":
				res.resp, res.err = c.dispatchBeginUpdateTags(req)
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

func (c *ContactProfilesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := c.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/contactProfiles/(?P<contactProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armorbital.ContactProfile](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		contactProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("contactProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, contactProfileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		c.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		c.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (c *ContactProfilesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/contactProfiles/(?P<contactProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		contactProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("contactProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, contactProfileNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ContactProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/contactProfiles/(?P<contactProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	contactProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("contactProfileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, contactProfileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContactProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContactProfilesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := c.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/contactProfiles`
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
		var options *armorbital.ContactProfilesClientListOptions
		if skiptokenParam != nil {
			options = &armorbital.ContactProfilesClientListOptions{
				Skiptoken: skiptokenParam,
			}
		}
		resp := c.srv.NewListPager(resourceGroupNameParam, options)
		newListPager = &resp
		c.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armorbital.ContactProfilesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		c.newListPager.remove(req)
	}
	return resp, nil
}

func (c *ContactProfilesServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := c.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/contactProfiles`
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
		var options *armorbital.ContactProfilesClientListBySubscriptionOptions
		if skiptokenParam != nil {
			options = &armorbital.ContactProfilesClientListBySubscriptionOptions{
				Skiptoken: skiptokenParam,
			}
		}
		resp := c.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		c.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armorbital.ContactProfilesClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		c.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (c *ContactProfilesServerTransport) dispatchBeginUpdateTags(req *http.Request) (*http.Response, error) {
	if c.srv.BeginUpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateTags not implemented")}
	}
	beginUpdateTags := c.beginUpdateTags.get(req)
	if beginUpdateTags == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Orbital/contactProfiles/(?P<contactProfileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		contactProfileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("contactProfileName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginUpdateTags(req.Context(), resourceGroupNameParam, contactProfileNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateTags = &respr
		c.beginUpdateTags.add(req, beginUpdateTags)
	}

	resp, err := server.PollerResponderNext(beginUpdateTags, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginUpdateTags.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateTags) {
		c.beginUpdateTags.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ContactProfilesServerTransport
var contactProfilesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
