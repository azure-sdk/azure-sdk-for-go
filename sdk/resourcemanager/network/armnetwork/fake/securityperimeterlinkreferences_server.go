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
	"regexp"
	"strconv"
)

// SecurityPerimeterLinkReferencesServer is a fake server for instances of the armnetwork.SecurityPerimeterLinkReferencesClient type.
type SecurityPerimeterLinkReferencesServer struct {
	// BeginDelete is the fake for method SecurityPerimeterLinkReferencesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, linkReferenceName string, options *armnetwork.SecurityPerimeterLinkReferencesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.SecurityPerimeterLinkReferencesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SecurityPerimeterLinkReferencesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkSecurityPerimeterName string, linkReferenceName string, options *armnetwork.SecurityPerimeterLinkReferencesClientGetOptions) (resp azfake.Responder[armnetwork.SecurityPerimeterLinkReferencesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SecurityPerimeterLinkReferencesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkSecurityPerimeterName string, options *armnetwork.SecurityPerimeterLinkReferencesClientListOptions) (resp azfake.PagerResponder[armnetwork.SecurityPerimeterLinkReferencesClientListResponse])
}

// NewSecurityPerimeterLinkReferencesServerTransport creates a new instance of SecurityPerimeterLinkReferencesServerTransport with the provided implementation.
// The returned SecurityPerimeterLinkReferencesServerTransport instance is connected to an instance of armnetwork.SecurityPerimeterLinkReferencesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSecurityPerimeterLinkReferencesServerTransport(srv *SecurityPerimeterLinkReferencesServer) *SecurityPerimeterLinkReferencesServerTransport {
	return &SecurityPerimeterLinkReferencesServerTransport{
		srv:          srv,
		beginDelete:  newTracker[azfake.PollerResponder[armnetwork.SecurityPerimeterLinkReferencesClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armnetwork.SecurityPerimeterLinkReferencesClientListResponse]](),
	}
}

// SecurityPerimeterLinkReferencesServerTransport connects instances of armnetwork.SecurityPerimeterLinkReferencesClient to instances of SecurityPerimeterLinkReferencesServer.
// Don't use this type directly, use NewSecurityPerimeterLinkReferencesServerTransport instead.
type SecurityPerimeterLinkReferencesServerTransport struct {
	srv          *SecurityPerimeterLinkReferencesServer
	beginDelete  *tracker[azfake.PollerResponder[armnetwork.SecurityPerimeterLinkReferencesClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armnetwork.SecurityPerimeterLinkReferencesClientListResponse]]
}

// Do implements the policy.Transporter interface for SecurityPerimeterLinkReferencesServerTransport.
func (s *SecurityPerimeterLinkReferencesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SecurityPerimeterLinkReferencesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if securityPerimeterLinkReferencesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = securityPerimeterLinkReferencesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SecurityPerimeterLinkReferencesClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "SecurityPerimeterLinkReferencesClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SecurityPerimeterLinkReferencesClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
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

func (s *SecurityPerimeterLinkReferencesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkReferences/(?P<linkReferenceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		linkReferenceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkReferenceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, linkReferenceNameParam, nil)
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

func (s *SecurityPerimeterLinkReferencesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkReferences/(?P<linkReferenceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	linkReferenceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("linkReferenceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, networkSecurityPerimeterNameParam, linkReferenceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).NspLinkReference, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SecurityPerimeterLinkReferencesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkSecurityPerimeters/(?P<networkSecurityPerimeterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/linkReferences`
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
		var options *armnetwork.SecurityPerimeterLinkReferencesClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.SecurityPerimeterLinkReferencesClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, networkSecurityPerimeterNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.SecurityPerimeterLinkReferencesClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to SecurityPerimeterLinkReferencesServerTransport
var securityPerimeterLinkReferencesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
