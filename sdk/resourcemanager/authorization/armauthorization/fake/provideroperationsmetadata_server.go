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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/authorization/armauthorization/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ProviderOperationsMetadataServer is a fake server for instances of the armauthorization.ProviderOperationsMetadataClient type.
type ProviderOperationsMetadataServer struct {
	// Get is the fake for method ProviderOperationsMetadataClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceProviderNamespace string, options *armauthorization.ProviderOperationsMetadataClientGetOptions) (resp azfake.Responder[armauthorization.ProviderOperationsMetadataClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ProviderOperationsMetadataClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armauthorization.ProviderOperationsMetadataClientListOptions) (resp azfake.PagerResponder[armauthorization.ProviderOperationsMetadataClientListResponse])
}

// NewProviderOperationsMetadataServerTransport creates a new instance of ProviderOperationsMetadataServerTransport with the provided implementation.
// The returned ProviderOperationsMetadataServerTransport instance is connected to an instance of armauthorization.ProviderOperationsMetadataClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewProviderOperationsMetadataServerTransport(srv *ProviderOperationsMetadataServer) *ProviderOperationsMetadataServerTransport {
	return &ProviderOperationsMetadataServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armauthorization.ProviderOperationsMetadataClientListResponse]](),
	}
}

// ProviderOperationsMetadataServerTransport connects instances of armauthorization.ProviderOperationsMetadataClient to instances of ProviderOperationsMetadataServer.
// Don't use this type directly, use NewProviderOperationsMetadataServerTransport instead.
type ProviderOperationsMetadataServerTransport struct {
	srv          *ProviderOperationsMetadataServer
	newListPager *tracker[azfake.PagerResponder[armauthorization.ProviderOperationsMetadataClientListResponse]]
}

// Do implements the policy.Transporter interface for ProviderOperationsMetadataServerTransport.
func (p *ProviderOperationsMetadataServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *ProviderOperationsMetadataServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if providerOperationsMetadataServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = providerOperationsMetadataServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ProviderOperationsMetadataClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "ProviderOperationsMetadataClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
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

func (p *ProviderOperationsMetadataServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Authorization/providerOperations/(?P<resourceProviderNamespace>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceProviderNamespaceParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceProviderNamespace")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armauthorization.ProviderOperationsMetadataClientGetOptions
	if expandParam != nil {
		options = &armauthorization.ProviderOperationsMetadataClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceProviderNamespaceParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ProviderOperationsMetadata, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *ProviderOperationsMetadataServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		qp := req.URL.Query()
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armauthorization.ProviderOperationsMetadataClientListOptions
		if expandParam != nil {
			options = &armauthorization.ProviderOperationsMetadataClientListOptions{
				Expand: expandParam,
			}
		}
		resp := p.srv.NewListPager(options)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armauthorization.ProviderOperationsMetadataClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		p.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ProviderOperationsMetadataServerTransport
var providerOperationsMetadataServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
