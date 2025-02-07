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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// NspProfilesServer is a fake server for instances of the armnetwork.NspProfilesClient type.
type NspProfilesServer struct {
	// CreateOrUpdate is the fake for method NspProfilesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, parameters armnetwork.NspProfile, options *armnetwork.NspProfilesClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.NspProfilesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method NspProfilesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, options *armnetwork.NspProfilesClientDeleteOptions) (resp azfake.Responder[armnetwork.NspProfilesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method NspProfilesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, profileName string, options *armnetwork.NspProfilesClientGetOptions) (resp azfake.Responder[armnetwork.NspProfilesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method NspProfilesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkSecurityPerimeterName string, options *armnetwork.NspProfilesClientListOptions) (resp azfake.PagerResponder[armnetwork.NspProfilesClientListResponse])
}

// NewNspProfilesServerTransport creates a new instance of NspProfilesServerTransport with the provided implementation.
// The returned NspProfilesServerTransport instance is connected to an instance of armnetwork.NspProfilesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewNspProfilesServerTransport(srv *NspProfilesServer) *NspProfilesServerTransport {
	return &NspProfilesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.NspProfilesClientListResponse]](),
	}
}

// NspProfilesServerTransport connects instances of armnetwork.NspProfilesClient to instances of NspProfilesServer.
// Don't use this type directly, use NewNspProfilesServerTransport instead.
type NspProfilesServerTransport struct {
	srv          *NspProfilesServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.NspProfilesClientListResponse]]
}

// Do implements the policy.Transporter interface for NspProfilesServerTransport.
func (n *NspProfilesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "NspProfilesClient.CreateOrUpdate":
		resp, err = n.dispatchCreateOrUpdate(req)
	case "NspProfilesClient.Delete":
		resp, err = n.dispatchDelete(req)
	case "NspProfilesClient.Get":
		resp, err = n.dispatchGet(req)
	case "NspProfilesClient.NewListPager":
		resp, err = n.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (n *NspProfilesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if n.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.NspProfile](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NspProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NspProfilesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if n.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Delete(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, nil)
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

func (n *NspProfilesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if n.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles/(?P<profileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
	if err != nil {
		return nil, err
	}
	profileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("profileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := n.srv.Get(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, profileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NspProfile, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (n *NspProfilesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if n.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := n.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/profiles`
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
		networkSecurityPerimeterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkSecurityPerimeterName")])
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
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.NspProfilesClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.NspProfilesClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := n.srv.NewListPager(resourceGroupNameParam, networkSecurityPerimeterNameParam, options)
		newListPager = &resp
		n.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.NspProfilesClientListResponse, createLink func() string) {
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
