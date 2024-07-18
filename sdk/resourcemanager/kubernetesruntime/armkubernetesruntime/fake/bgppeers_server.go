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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesruntime/armkubernetesruntime"
	"net/http"
	"net/url"
	"regexp"
)

// BgpPeersServer is a fake server for instances of the armkubernetesruntime.BgpPeersClient type.
type BgpPeersServer struct {
	// BeginCreateOrUpdate is the fake for method BgpPeersClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceURI string, bgpPeerName string, resource armkubernetesruntime.BgpPeer, options *armkubernetesruntime.BgpPeersClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armkubernetesruntime.BgpPeersClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BgpPeersClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceURI string, bgpPeerName string, options *armkubernetesruntime.BgpPeersClientDeleteOptions) (resp azfake.Responder[armkubernetesruntime.BgpPeersClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BgpPeersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceURI string, bgpPeerName string, options *armkubernetesruntime.BgpPeersClientGetOptions) (resp azfake.Responder[armkubernetesruntime.BgpPeersClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByParentPager is the fake for method BgpPeersClient.NewListByParentPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByParentPager func(resourceURI string, options *armkubernetesruntime.BgpPeersClientListByParentOptions) (resp azfake.PagerResponder[armkubernetesruntime.BgpPeersClientListByParentResponse])
}

// NewBgpPeersServerTransport creates a new instance of BgpPeersServerTransport with the provided implementation.
// The returned BgpPeersServerTransport instance is connected to an instance of armkubernetesruntime.BgpPeersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBgpPeersServerTransport(srv *BgpPeersServer) *BgpPeersServerTransport {
	return &BgpPeersServerTransport{
		srv:                  srv,
		beginCreateOrUpdate:  newTracker[azfake.PollerResponder[armkubernetesruntime.BgpPeersClientCreateOrUpdateResponse]](),
		newListByParentPager: newTracker[azfake.PagerResponder[armkubernetesruntime.BgpPeersClientListByParentResponse]](),
	}
}

// BgpPeersServerTransport connects instances of armkubernetesruntime.BgpPeersClient to instances of BgpPeersServer.
// Don't use this type directly, use NewBgpPeersServerTransport instead.
type BgpPeersServerTransport struct {
	srv                  *BgpPeersServer
	beginCreateOrUpdate  *tracker[azfake.PollerResponder[armkubernetesruntime.BgpPeersClientCreateOrUpdateResponse]]
	newListByParentPager *tracker[azfake.PagerResponder[armkubernetesruntime.BgpPeersClientListByParentResponse]]
}

// Do implements the policy.Transporter interface for BgpPeersServerTransport.
func (b *BgpPeersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BgpPeersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "BgpPeersClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BgpPeersClient.Delete":
		resp, err = b.dispatchDelete(req)
	case "BgpPeersClient.Get":
		resp, err = b.dispatchGet(req)
	case "BgpPeersClient.NewListByParentPager":
		resp, err = b.dispatchNewListByParentPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (b *BgpPeersServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesRuntime/bgpPeers/(?P<bgpPeerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armkubernetesruntime.BgpPeer](req)
		if err != nil {
			return nil, err
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		bgpPeerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bgpPeerName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceURIParam, bgpPeerNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BgpPeersServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesRuntime/bgpPeers/(?P<bgpPeerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	bgpPeerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bgpPeerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), resourceURIParam, bgpPeerNameParam, nil)
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

func (b *BgpPeersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesRuntime/bgpPeers/(?P<bgpPeerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
	if err != nil {
		return nil, err
	}
	bgpPeerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("bgpPeerName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceURIParam, bgpPeerNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BgpPeer, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BgpPeersServerTransport) dispatchNewListByParentPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListByParentPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByParentPager not implemented")}
	}
	newListByParentPager := b.newListByParentPager.get(req)
	if newListByParentPager == nil {
		const regexStr = `/(?P<resourceUri>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.KubernetesRuntime/bgpPeers`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceURIParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceUri")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListByParentPager(resourceURIParam, nil)
		newListByParentPager = &resp
		b.newListByParentPager.add(req, newListByParentPager)
		server.PagerResponderInjectNextLinks(newListByParentPager, req, func(page *armkubernetesruntime.BgpPeersClientListByParentResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByParentPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListByParentPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByParentPager) {
		b.newListByParentPager.remove(req)
	}
	return resp, nil
}
