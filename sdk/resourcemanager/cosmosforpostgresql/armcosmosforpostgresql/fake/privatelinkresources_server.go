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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmosforpostgresql/armcosmosforpostgresql"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateLinkResourcesServer is a fake server for instances of the armcosmosforpostgresql.PrivateLinkResourcesClient type.
type PrivateLinkResourcesServer struct {
	// Get is the fake for method PrivateLinkResourcesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, clusterName string, privateLinkResourceName string, options *armcosmosforpostgresql.PrivateLinkResourcesClientGetOptions) (resp azfake.Responder[armcosmosforpostgresql.PrivateLinkResourcesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByClusterPager is the fake for method PrivateLinkResourcesClient.NewListByClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByClusterPager func(resourceGroupName string, clusterName string, options *armcosmosforpostgresql.PrivateLinkResourcesClientListByClusterOptions) (resp azfake.PagerResponder[armcosmosforpostgresql.PrivateLinkResourcesClientListByClusterResponse])
}

// NewPrivateLinkResourcesServerTransport creates a new instance of PrivateLinkResourcesServerTransport with the provided implementation.
// The returned PrivateLinkResourcesServerTransport instance is connected to an instance of armcosmosforpostgresql.PrivateLinkResourcesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateLinkResourcesServerTransport(srv *PrivateLinkResourcesServer) *PrivateLinkResourcesServerTransport {
	return &PrivateLinkResourcesServerTransport{
		srv:                   srv,
		newListByClusterPager: newTracker[azfake.PagerResponder[armcosmosforpostgresql.PrivateLinkResourcesClientListByClusterResponse]](),
	}
}

// PrivateLinkResourcesServerTransport connects instances of armcosmosforpostgresql.PrivateLinkResourcesClient to instances of PrivateLinkResourcesServer.
// Don't use this type directly, use NewPrivateLinkResourcesServerTransport instead.
type PrivateLinkResourcesServerTransport struct {
	srv                   *PrivateLinkResourcesServer
	newListByClusterPager *tracker[azfake.PagerResponder[armcosmosforpostgresql.PrivateLinkResourcesClientListByClusterResponse]]
}

// Do implements the policy.Transporter interface for PrivateLinkResourcesServerTransport.
func (p *PrivateLinkResourcesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateLinkResourcesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateLinkResourcesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateLinkResourcesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateLinkResourcesClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PrivateLinkResourcesClient.NewListByClusterPager":
				res.resp, res.err = p.dispatchNewListByClusterPager(req)
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

func (p *PrivateLinkResourcesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/serverGroupsv2/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources/(?P<privateLinkResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	privateLinkResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateLinkResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, clusterNameParam, privateLinkResourceNameParam, nil)
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

func (p *PrivateLinkResourcesServerTransport) dispatchNewListByClusterPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByClusterPager not implemented")}
	}
	newListByClusterPager := p.newListByClusterPager.get(req)
	if newListByClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DBforPostgreSQL/serverGroupsv2/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateLinkResources`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByClusterPager(resourceGroupNameParam, clusterNameParam, nil)
		newListByClusterPager = &resp
		p.newListByClusterPager.add(req, newListByClusterPager)
	}
	resp, err := server.PagerResponderNext(newListByClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByClusterPager) {
		p.newListByClusterPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateLinkResourcesServerTransport
var privateLinkResourcesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
