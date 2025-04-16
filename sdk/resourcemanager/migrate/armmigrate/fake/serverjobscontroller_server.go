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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
)

// ServerJobsControllerServer is a fake server for instances of the armmigrate.ServerJobsControllerClient type.
type ServerJobsControllerServer struct {
	// Get is the fake for method ServerJobsControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, jobName string, options *armmigrate.ServerJobsControllerClientGetOptions) (resp azfake.Responder[armmigrate.ServerJobsControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByServerSiteResourcePager is the fake for method ServerJobsControllerClient.NewListByServerSiteResourcePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServerSiteResourcePager func(resourceGroupName string, siteName string, options *armmigrate.ServerJobsControllerClientListByServerSiteResourceOptions) (resp azfake.PagerResponder[armmigrate.ServerJobsControllerClientListByServerSiteResourceResponse])
}

// NewServerJobsControllerServerTransport creates a new instance of ServerJobsControllerServerTransport with the provided implementation.
// The returned ServerJobsControllerServerTransport instance is connected to an instance of armmigrate.ServerJobsControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerJobsControllerServerTransport(srv *ServerJobsControllerServer) *ServerJobsControllerServerTransport {
	return &ServerJobsControllerServerTransport{
		srv:                              srv,
		newListByServerSiteResourcePager: newTracker[azfake.PagerResponder[armmigrate.ServerJobsControllerClientListByServerSiteResourceResponse]](),
	}
}

// ServerJobsControllerServerTransport connects instances of armmigrate.ServerJobsControllerClient to instances of ServerJobsControllerServer.
// Don't use this type directly, use NewServerJobsControllerServerTransport instead.
type ServerJobsControllerServerTransport struct {
	srv                              *ServerJobsControllerServer
	newListByServerSiteResourcePager *tracker[azfake.PagerResponder[armmigrate.ServerJobsControllerClientListByServerSiteResourceResponse]]
}

// Do implements the policy.Transporter interface for ServerJobsControllerServerTransport.
func (s *ServerJobsControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerJobsControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverJobsControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverJobsControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServerJobsControllerClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServerJobsControllerClient.NewListByServerSiteResourcePager":
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

func (s *ServerJobsControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs/(?P<jobName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	jobNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, jobNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerJob, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerJobsControllerServerTransport) dispatchNewListByServerSiteResourcePager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByServerSiteResourcePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServerSiteResourcePager not implemented")}
	}
	newListByServerSiteResourcePager := s.newListByServerSiteResourcePager.get(req)
	if newListByServerSiteResourcePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/jobs`
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
		server.PagerResponderInjectNextLinks(newListByServerSiteResourcePager, req, func(page *armmigrate.ServerJobsControllerClientListByServerSiteResourceResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to ServerJobsControllerServerTransport
var serverJobsControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
