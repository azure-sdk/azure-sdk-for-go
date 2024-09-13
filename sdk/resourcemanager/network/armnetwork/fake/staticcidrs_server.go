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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// StaticCidrsServer is a fake server for instances of the armnetwork.StaticCidrsClient type.
type StaticCidrsServer struct {
	// Create is the fake for method StaticCidrsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, networkManagerName string, poolName string, staticCidrName string, options *armnetwork.StaticCidrsClientCreateOptions) (resp azfake.Responder[armnetwork.StaticCidrsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method StaticCidrsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkManagerName string, poolName string, staticCidrName string, options *armnetwork.StaticCidrsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.StaticCidrsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method StaticCidrsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, poolName string, staticCidrName string, options *armnetwork.StaticCidrsClientGetOptions) (resp azfake.Responder[armnetwork.StaticCidrsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method StaticCidrsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, poolName string, options *armnetwork.StaticCidrsClientListOptions) (resp azfake.PagerResponder[armnetwork.StaticCidrsClientListResponse])
}

// NewStaticCidrsServerTransport creates a new instance of StaticCidrsServerTransport with the provided implementation.
// The returned StaticCidrsServerTransport instance is connected to an instance of armnetwork.StaticCidrsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewStaticCidrsServerTransport(srv *StaticCidrsServer) *StaticCidrsServerTransport {
	return &StaticCidrsServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.StaticCidrsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.StaticCidrsClientListResponse]](),
	}
}

// StaticCidrsServerTransport connects instances of armnetwork.StaticCidrsClient to instances of StaticCidrsServer.
// Don't use this type directly, use NewStaticCidrsServerTransport instead.
type StaticCidrsServerTransport struct {
	srv          *StaticCidrsServer
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.StaticCidrsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.StaticCidrsClientListResponse]]
}

// Do implements the policy.Transporter interface for StaticCidrsServerTransport.
func (s *StaticCidrsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "StaticCidrsClient.Create":
		resp, err = s.dispatchCreate(req)
	case "StaticCidrsClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "StaticCidrsClient.Get":
		resp, err = s.dispatchGet(req)
	case "StaticCidrsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *StaticCidrsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipamPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staticCidrs/(?P<staticCidrName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.StaticCidr](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	staticCidrNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("staticCidrName")])
	if err != nil {
		return nil, err
	}
	var options *armnetwork.StaticCidrsClientCreateOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armnetwork.StaticCidrsClientCreateOptions{
			Body: &body,
		}
	}
	respr, errRespr := s.srv.Create(req.Context(), resourceGroupNameParam, networkManagerNameParam, poolNameParam, staticCidrNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StaticCidr, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StaticCidrsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipamPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staticCidrs/(?P<staticCidrName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		staticCidrNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("staticCidrName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkManagerNameParam, poolNameParam, staticCidrNameParam, nil)
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

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *StaticCidrsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipamPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staticCidrs/(?P<staticCidrName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	staticCidrNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("staticCidrName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, networkManagerNameParam, poolNameParam, staticCidrNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).StaticCidr, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *StaticCidrsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/ipamPools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/staticCidrs`
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
		networkManagerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		skipUnescaped, err := url.QueryUnescape(qp.Get("skip"))
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
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
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
		sortKeyUnescaped, err := url.QueryUnescape(qp.Get("sortKey"))
		if err != nil {
			return nil, err
		}
		sortKeyParam := getOptional(sortKeyUnescaped)
		sortValueUnescaped, err := url.QueryUnescape(qp.Get("sortValue"))
		if err != nil {
			return nil, err
		}
		sortValueParam := getOptional(sortValueUnescaped)
		var options *armnetwork.StaticCidrsClientListOptions
		if skipTokenParam != nil || skipParam != nil || topParam != nil || sortKeyParam != nil || sortValueParam != nil {
			options = &armnetwork.StaticCidrsClientListOptions{
				SkipToken: skipTokenParam,
				Skip:      skipParam,
				Top:       topParam,
				SortKey:   sortKeyParam,
				SortValue: sortValueParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, networkManagerNameParam, poolNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.StaticCidrsClientListResponse, createLink func() string) {
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