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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationhub/armmigrationhub"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourceControllerServer is a fake server for instances of the armmigrationhub.PrivateLinkResourceControllerClient type.
type PrivateLinkResourceControllerServer struct {
	// GetPrivateLinkResource is the fake for method PrivateLinkResourceControllerClient.GetPrivateLinkResource
	// HTTP status codes to indicate success: http.StatusOK
	GetPrivateLinkResource func(ctx context.Context, resourceGroupName string, migrateProjectName string, privateLinkResourceName string, options *armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourceOptions) (resp azfake.Responder[armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourceResponse], errResp azfake.ErrorResponder)

	// NewGetPrivateLinkResourcesPager is the fake for method PrivateLinkResourceControllerClient.NewGetPrivateLinkResourcesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetPrivateLinkResourcesPager func(resourceGroupName string, migrateProjectName string, options *armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourcesOptions) (resp azfake.PagerResponder[armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourcesResponse])
}

// NewPrivateLinkResourceControllerServerTransport creates a new instance of PrivateLinkResourceControllerServerTransport with the provided implementation.
// The returned PrivateLinkResourceControllerServerTransport instance is connected to an instance of armmigrationhub.PrivateLinkResourceControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourceControllerServerTransport(srv *PrivateLinkResourceControllerServer) *PrivateLinkResourceControllerServerTransport {
	return &PrivateLinkResourceControllerServerTransport{
		srv:                             srv,
		newGetPrivateLinkResourcesPager: newTracker[azfake.PagerResponder[armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourcesResponse]](),
	}
}

// PrivateLinkResourceControllerServerTransport connects instances of armmigrationhub.PrivateLinkResourceControllerClient to instances of PrivateLinkResourceControllerServer.
// Don't use this type directly, use NewPrivateLinkResourceControllerServerTransport instead.
type PrivateLinkResourceControllerServerTransport struct {
	srv                             *PrivateLinkResourceControllerServer
	newGetPrivateLinkResourcesPager *tracker[azfake.PagerResponder[armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourcesResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkResourceControllerServerTransport.
func (p *PrivateLinkResourceControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinkResourceControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateLinkResourceControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateLinkResourceControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateLinkResourceControllerClient.GetPrivateLinkResource":
				res.resp, res.err = p.dispatchGetPrivateLinkResource(req)
			case "PrivateLinkResourceControllerClient.NewGetPrivateLinkResourcesPager":
				res.resp, res.err = p.dispatchNewGetPrivateLinkResourcesPager(req)
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

func (p *PrivateLinkResourceControllerServerTransport) dispatchGetPrivateLinkResource(req *http.Request) (*http.Response, error) {
	if p.srv.GetPrivateLinkResource == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetPrivateLinkResource not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/migrateProjects/(?P<migrateProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<privateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	migrateProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrateProjectName")])
	if err != nil {
		return nil, err
	}
	privateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateLinkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.GetPrivateLinkResource(req.Context(), resourceGroupNameParam, migrateProjectNameParam, privateLinkResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateLinkResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateLinkResourceControllerServerTransport) dispatchNewGetPrivateLinkResourcesPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewGetPrivateLinkResourcesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetPrivateLinkResourcesPager not implemented")}
	}
	newGetPrivateLinkResourcesPager := p.newGetPrivateLinkResourcesPager.get(req)
	if newGetPrivateLinkResourcesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Migrate/migrateProjects/(?P<migrateProjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		migrateProjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("migrateProjectName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewGetPrivateLinkResourcesPager(resourceGroupNameParam, migrateProjectNameParam, nil)
		newGetPrivateLinkResourcesPager = &resp
		p.newGetPrivateLinkResourcesPager.add(req, newGetPrivateLinkResourcesPager)
		server.PagerResponderInjectNextLinks(newGetPrivateLinkResourcesPager, req, func(page *armmigrationhub.PrivateLinkResourceControllerClientGetPrivateLinkResourcesResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetPrivateLinkResourcesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newGetPrivateLinkResourcesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetPrivateLinkResourcesPager) {
		p.newGetPrivateLinkResourcesPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateLinkResourceControllerServerTransport
var privateLinkResourceControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
