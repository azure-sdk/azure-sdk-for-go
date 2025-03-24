// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster"
	"net/http"
	"net/url"
	"regexp"
)

// PrivateEndpointConnectionsServer is a fake server for instances of the armmongocluster.PrivateEndpointConnectionsClient type.
type PrivateEndpointConnectionsServer struct {
	// BeginCreate is the fake for method PrivateEndpointConnectionsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, mongoClusterName string, privateEndpointConnectionName string, resource armmongocluster.PrivateEndpointConnectionResource, options *armmongocluster.PrivateEndpointConnectionsClientBeginCreateOptions) (resp azfake.PollerResponder[armmongocluster.PrivateEndpointConnectionsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PrivateEndpointConnectionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, mongoClusterName string, privateEndpointConnectionName string, options *armmongocluster.PrivateEndpointConnectionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmongocluster.PrivateEndpointConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, mongoClusterName string, privateEndpointConnectionName string, options *armmongocluster.PrivateEndpointConnectionsClientGetOptions) (resp azfake.Responder[armmongocluster.PrivateEndpointConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMongoClusterPager is the fake for method PrivateEndpointConnectionsClient.NewListByMongoClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMongoClusterPager func(resourceGroupName string, mongoClusterName string, options *armmongocluster.PrivateEndpointConnectionsClientListByMongoClusterOptions) (resp azfake.PagerResponder[armmongocluster.PrivateEndpointConnectionsClientListByMongoClusterResponse])
}

// NewPrivateEndpointConnectionsServerTransport creates a new instance of PrivateEndpointConnectionsServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionsServerTransport instance is connected to an instance of armmongocluster.PrivateEndpointConnectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionsServerTransport(srv *PrivateEndpointConnectionsServer) *PrivateEndpointConnectionsServerTransport {
	return &PrivateEndpointConnectionsServerTransport{
		srv:                        srv,
		beginCreate:                newTracker[azfake.PollerResponder[armmongocluster.PrivateEndpointConnectionsClientCreateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armmongocluster.PrivateEndpointConnectionsClientDeleteResponse]](),
		newListByMongoClusterPager: newTracker[azfake.PagerResponder[armmongocluster.PrivateEndpointConnectionsClientListByMongoClusterResponse]](),
	}
}

// PrivateEndpointConnectionsServerTransport connects instances of armmongocluster.PrivateEndpointConnectionsClient to instances of PrivateEndpointConnectionsServer.
// Don't use this type directly, use NewPrivateEndpointConnectionsServerTransport instead.
type PrivateEndpointConnectionsServerTransport struct {
	srv                        *PrivateEndpointConnectionsServer
	beginCreate                *tracker[azfake.PollerResponder[armmongocluster.PrivateEndpointConnectionsClientCreateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armmongocluster.PrivateEndpointConnectionsClientDeleteResponse]]
	newListByMongoClusterPager *tracker[azfake.PagerResponder[armmongocluster.PrivateEndpointConnectionsClientListByMongoClusterResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionsServerTransport.
func (p *PrivateEndpointConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateEndpointConnectionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateEndpointConnectionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateEndpointConnectionsClient.BeginCreate":
				res.resp, res.err = p.dispatchBeginCreate(req)
			case "PrivateEndpointConnectionsClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PrivateEndpointConnectionsClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PrivateEndpointConnectionsClient.NewListByMongoClusterPager":
				res.resp, res.err = p.dispatchNewListByMongoClusterPager(req)
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

func (p *PrivateEndpointConnectionsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := p.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmongocluster.PrivateEndpointConnectionResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreate(req.Context(), resourceGroupNameParam, mongoClusterNameParam, privateEndpointConnectionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		p.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		p.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		p.beginCreate.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, mongoClusterNameParam, privateEndpointConnectionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		p.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		p.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		p.beginDelete.remove(req)
	}

	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, mongoClusterNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnectionResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionsServerTransport) dispatchNewListByMongoClusterPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByMongoClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMongoClusterPager not implemented")}
	}
	newListByMongoClusterPager := p.newListByMongoClusterPager.get(req)
	if newListByMongoClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/mongoClusters/(?P<mongoClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		mongoClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("mongoClusterName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByMongoClusterPager(resourceGroupNameParam, mongoClusterNameParam, nil)
		newListByMongoClusterPager = &resp
		p.newListByMongoClusterPager.add(req, newListByMongoClusterPager)
		server.PagerResponderInjectNextLinks(newListByMongoClusterPager, req, func(page *armmongocluster.PrivateEndpointConnectionsClientListByMongoClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMongoClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByMongoClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMongoClusterPager) {
		p.newListByMongoClusterPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateEndpointConnectionsServerTransport
var privateEndpointConnectionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
