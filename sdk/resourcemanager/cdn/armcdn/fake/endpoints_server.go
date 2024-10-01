//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn/v3"
	"net/http"
	"net/url"
	"regexp"
)

// EndpointsServer is a fake server for instances of the armcdn.EndpointsClient type.
type EndpointsServer struct {
	// BeginCreate is the fake for method EndpointsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpoint armcdn.Endpoint, options *armcdn.EndpointsClientBeginCreateOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method EndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *armcdn.EndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *armcdn.EndpointsClientGetOptions) (resp azfake.Responder[armcdn.EndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByProfilePager is the fake for method EndpointsClient.NewListByProfilePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProfilePager func(resourceGroupName string, profileName string, options *armcdn.EndpointsClientListByProfileOptions) (resp azfake.PagerResponder[armcdn.EndpointsClientListByProfileResponse])

	// NewListResourceUsagePager is the fake for method EndpointsClient.NewListResourceUsagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListResourceUsagePager func(resourceGroupName string, profileName string, endpointName string, options *armcdn.EndpointsClientListResourceUsageOptions) (resp azfake.PagerResponder[armcdn.EndpointsClientListResourceUsageResponse])

	// BeginLoadContent is the fake for method EndpointsClient.BeginLoadContent
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginLoadContent func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths armcdn.LoadParameters, options *armcdn.EndpointsClientBeginLoadContentOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientLoadContentResponse], errResp azfake.ErrorResponder)

	// BeginPurgeContent is the fake for method EndpointsClient.BeginPurgeContent
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPurgeContent func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, contentFilePaths armcdn.PurgeParameters, options *armcdn.EndpointsClientBeginPurgeContentOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientPurgeContentResponse], errResp azfake.ErrorResponder)

	// BeginStart is the fake for method EndpointsClient.BeginStart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStart func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *armcdn.EndpointsClientBeginStartOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientStartResponse], errResp azfake.ErrorResponder)

	// BeginStop is the fake for method EndpointsClient.BeginStop
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStop func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *armcdn.EndpointsClientBeginStopOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientStopResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method EndpointsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, endpointUpdateProperties armcdn.EndpointUpdateParameters, options *armcdn.EndpointsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcdn.EndpointsClientUpdateResponse], errResp azfake.ErrorResponder)

	// ValidateCustomDomain is the fake for method EndpointsClient.ValidateCustomDomain
	// HTTP status codes to indicate success: http.StatusOK
	ValidateCustomDomain func(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainProperties armcdn.ValidateCustomDomainInput, options *armcdn.EndpointsClientValidateCustomDomainOptions) (resp azfake.Responder[armcdn.EndpointsClientValidateCustomDomainResponse], errResp azfake.ErrorResponder)
}

// NewEndpointsServerTransport creates a new instance of EndpointsServerTransport with the provided implementation.
// The returned EndpointsServerTransport instance is connected to an instance of armcdn.EndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEndpointsServerTransport(srv *EndpointsServer) *EndpointsServerTransport {
	return &EndpointsServerTransport{
		srv:                       srv,
		beginCreate:               newTracker[azfake.PollerResponder[armcdn.EndpointsClientCreateResponse]](),
		beginDelete:               newTracker[azfake.PollerResponder[armcdn.EndpointsClientDeleteResponse]](),
		newListByProfilePager:     newTracker[azfake.PagerResponder[armcdn.EndpointsClientListByProfileResponse]](),
		newListResourceUsagePager: newTracker[azfake.PagerResponder[armcdn.EndpointsClientListResourceUsageResponse]](),
		beginLoadContent:          newTracker[azfake.PollerResponder[armcdn.EndpointsClientLoadContentResponse]](),
		beginPurgeContent:         newTracker[azfake.PollerResponder[armcdn.EndpointsClientPurgeContentResponse]](),
		beginStart:                newTracker[azfake.PollerResponder[armcdn.EndpointsClientStartResponse]](),
		beginStop:                 newTracker[azfake.PollerResponder[armcdn.EndpointsClientStopResponse]](),
		beginUpdate:               newTracker[azfake.PollerResponder[armcdn.EndpointsClientUpdateResponse]](),
	}
}

// EndpointsServerTransport connects instances of armcdn.EndpointsClient to instances of EndpointsServer.
// Don't use this type directly, use NewEndpointsServerTransport instead.
type EndpointsServerTransport struct {
	srv                       *EndpointsServer
	beginCreate               *tracker[azfake.PollerResponder[armcdn.EndpointsClientCreateResponse]]
	beginDelete               *tracker[azfake.PollerResponder[armcdn.EndpointsClientDeleteResponse]]
	newListByProfilePager     *tracker[azfake.PagerResponder[armcdn.EndpointsClientListByProfileResponse]]
	newListResourceUsagePager *tracker[azfake.PagerResponder[armcdn.EndpointsClientListResourceUsageResponse]]
	beginLoadContent          *tracker[azfake.PollerResponder[armcdn.EndpointsClientLoadContentResponse]]
	beginPurgeContent         *tracker[azfake.PollerResponder[armcdn.EndpointsClientPurgeContentResponse]]
	beginStart                *tracker[azfake.PollerResponder[armcdn.EndpointsClientStartResponse]]
	beginStop                 *tracker[azfake.PollerResponder[armcdn.EndpointsClientStopResponse]]
	beginUpdate               *tracker[azfake.PollerResponder[armcdn.EndpointsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for EndpointsServerTransport.
func (e *EndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EndpointsClient.BeginCreate":
		resp, err = e.dispatchBeginCreate(req)
	case "EndpointsClient.BeginDelete":
		resp, err = e.dispatchBeginDelete(req)
	case "EndpointsClient.Get":
		resp, err = e.dispatchGet(req)
	case "EndpointsClient.NewListByProfilePager":
		resp, err = e.dispatchNewListByProfilePager(req)
	case "EndpointsClient.NewListResourceUsagePager":
		resp, err = e.dispatchNewListResourceUsagePager(req)
	case "EndpointsClient.BeginLoadContent":
		resp, err = e.dispatchBeginLoadContent(req)
	case "EndpointsClient.BeginPurgeContent":
		resp, err = e.dispatchBeginPurgeContent(req)
	case "EndpointsClient.BeginStart":
		resp, err = e.dispatchBeginStart(req)
	case "EndpointsClient.BeginStop":
		resp, err = e.dispatchBeginStop(req)
	case "EndpointsClient.BeginUpdate":
		resp, err = e.dispatchBeginUpdate(req)
	case "EndpointsClient.ValidateCustomDomain":
		resp, err = e.dispatchValidateCustomDomain(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := e.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.Endpoint](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreate(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		e.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		e.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		e.beginCreate.remove(req)
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if e.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := e.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginDelete(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
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

func (e *EndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Endpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointsServerTransport) dispatchNewListByProfilePager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListByProfilePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProfilePager not implemented")}
	}
	newListByProfilePager := e.newListByProfilePager.get(req)
	if newListByProfilePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListByProfilePager(resourceGroupNameParam, profileNameParam, nil)
		newListByProfilePager = &resp
		e.newListByProfilePager.add(req, newListByProfilePager)
		server.PagerResponderInjectNextLinks(newListByProfilePager, req, func(page *armcdn.EndpointsClientListByProfileResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProfilePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListByProfilePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProfilePager) {
		e.newListByProfilePager.remove(req)
	}
	return resp, nil
}

func (e *EndpointsServerTransport) dispatchNewListResourceUsagePager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListResourceUsagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListResourceUsagePager not implemented")}
	}
	newListResourceUsagePager := e.newListResourceUsagePager.get(req)
	if newListResourceUsagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/checkResourceUsage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListResourceUsagePager(resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
		newListResourceUsagePager = &resp
		e.newListResourceUsagePager.add(req, newListResourceUsagePager)
		server.PagerResponderInjectNextLinks(newListResourceUsagePager, req, func(page *armcdn.EndpointsClientListResourceUsageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListResourceUsagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListResourceUsagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListResourceUsagePager) {
		e.newListResourceUsagePager.remove(req)
	}
	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginLoadContent(req *http.Request) (*http.Response, error) {
	if e.srv.BeginLoadContent == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginLoadContent not implemented")}
	}
	beginLoadContent := e.beginLoadContent.get(req)
	if beginLoadContent == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/load`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.LoadParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginLoadContent(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginLoadContent = &respr
		e.beginLoadContent.add(req, beginLoadContent)
	}

	resp, err := server.PollerResponderNext(beginLoadContent, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginLoadContent.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginLoadContent) {
		e.beginLoadContent.remove(req)
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginPurgeContent(req *http.Request) (*http.Response, error) {
	if e.srv.BeginPurgeContent == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurgeContent not implemented")}
	}
	beginPurgeContent := e.beginPurgeContent.get(req)
	if beginPurgeContent == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/purge`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.PurgeParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginPurgeContent(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurgeContent = &respr
		e.beginPurgeContent.add(req, beginPurgeContent)
	}

	resp, err := server.PollerResponderNext(beginPurgeContent, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginPurgeContent.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurgeContent) {
		e.beginPurgeContent.remove(req)
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginStart(req *http.Request) (*http.Response, error) {
	if e.srv.BeginStart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStart not implemented")}
	}
	beginStart := e.beginStart.get(req)
	if beginStart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginStart(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStart = &respr
		e.beginStart.add(req, beginStart)
	}

	resp, err := server.PollerResponderNext(beginStart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginStart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStart) {
		e.beginStart.remove(req)
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginStop(req *http.Request) (*http.Response, error) {
	if e.srv.BeginStop == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStop not implemented")}
	}
	beginStop := e.beginStop.get(req)
	if beginStop == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginStop(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStop = &respr
		e.beginStop.add(req, beginStop)
	}

	resp, err := server.PollerResponderNext(beginStop, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginStop.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStop) {
		e.beginStop.remove(req)
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := e.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcdn.EndpointUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdate(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		e.beginUpdate.remove(req)
	}

	return resp, nil
}

func (e *EndpointsServerTransport) dispatchValidateCustomDomain(req *http.Request) (*http.Response, error) {
	if e.srv.ValidateCustomDomain == nil {
		return nil, &nonRetriableError{errors.New("fake for method ValidateCustomDomain not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Cdn/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/validateCustomDomain`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcdn.ValidateCustomDomainInput](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.ValidateCustomDomain(req.Context(), resourceGroupNameParam, profileNameParam, endpointNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ValidateCustomDomainOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
