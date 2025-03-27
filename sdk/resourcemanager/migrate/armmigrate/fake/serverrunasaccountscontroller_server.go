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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ServerRunAsAccountsControllerServer is a fake server for instances of the armmigrate.ServerRunAsAccountsControllerClient type.
type ServerRunAsAccountsControllerServer struct {
	// Get is the fake for method ServerRunAsAccountsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *armmigrate.ServerRunAsAccountsControllerClientGetOptions) (resp azfake.Responder[armmigrate.ServerRunAsAccountsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerSiteResourcePager is the fake for method ServerRunAsAccountsControllerClient.NewListByServerSiteResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerSiteResourcePager func(resourceGroupName string, siteName string, options *armmigrate.ServerRunAsAccountsControllerClientListByServerSiteResourceOptions) (resp azfake.PagerResponder[armmigrate.ServerRunAsAccountsControllerClientListByServerSiteResourceResponse])
}

// NewServerRunAsAccountsControllerServerTransport creates a new instance of ServerRunAsAccountsControllerServerTransport with the provided implementation.
// The returned ServerRunAsAccountsControllerServerTransport instance is connected to an instance of armmigrate.ServerRunAsAccountsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerRunAsAccountsControllerServerTransport(srv *ServerRunAsAccountsControllerServer) *ServerRunAsAccountsControllerServerTransport {
	return &ServerRunAsAccountsControllerServerTransport{
		srv:                              srv,
		newListByServerSiteResourcePager: newTracker[azfake.PagerResponder[armmigrate.ServerRunAsAccountsControllerClientListByServerSiteResourceResponse]](),
	}
}

// ServerRunAsAccountsControllerServerTransport connects instances of armmigrate.ServerRunAsAccountsControllerClient to instances of ServerRunAsAccountsControllerServer.
// Don't use this type directly, use NewServerRunAsAccountsControllerServerTransport instead.
type ServerRunAsAccountsControllerServerTransport struct {
	srv                              *ServerRunAsAccountsControllerServer
	newListByServerSiteResourcePager *tracker[azfake.PagerResponder[armmigrate.ServerRunAsAccountsControllerClientListByServerSiteResourceResponse]]
}

// Do implements the policy.Transporter interface for ServerRunAsAccountsControllerServerTransport.
func (s *ServerRunAsAccountsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerRunAsAccountsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverRunAsAccountsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverRunAsAccountsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServerRunAsAccountsControllerClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServerRunAsAccountsControllerClient.NewListByServerSiteResourcePager":
				res.resp, res.err = s.dispatchNewListByServerSiteResourcePager(req)
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

func (s *ServerRunAsAccountsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerRunAsAccount, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerRunAsAccountsControllerServerTransport) dispatchNewListByServerSiteResourcePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByServerSiteResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerSiteResourcePager not implemented")}
	}
	newListByServerSiteResourcePager := s.newListByServerSiteResourcePager.get(req)
	if newListByServerSiteResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runAsAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByServerSiteResourcePager(resourceGroupNameParam, siteNameParam, nil)
		newListByServerSiteResourcePager = &resp
		s.newListByServerSiteResourcePager.add(req, newListByServerSiteResourcePager)
		server.PagerResponderInjectNextLinks(newListByServerSiteResourcePager, req, func(page *armmigrate.ServerRunAsAccountsControllerClientListByServerSiteResourceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServerSiteResourcePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByServerSiteResourcePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServerSiteResourcePager) {
		s.newListByServerSiteResourcePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerRunAsAccountsControllerServerTransport
var serverRunAsAccountsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
