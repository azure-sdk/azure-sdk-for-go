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

// HypervHostControllerServer is a fake server for instances of the armmigrate.HypervHostControllerClient type.
type HypervHostControllerServer struct {
	// BeginCreate is the fake for method HypervHostControllerClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, siteName string, hostName string, body armmigrate.HypervHost, options *armmigrate.HypervHostControllerClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrate.HypervHostControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method HypervHostControllerClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *armmigrate.HypervHostControllerClientBeginDeleteOptions) (resp azfake.PollerResponder[armmigrate.HypervHostControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HypervHostControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *armmigrate.HypervHostControllerClientGetOptions) (resp azfake.Responder[armmigrate.HypervHostControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByHypervSitePager is the fake for method HypervHostControllerClient.NewListByHypervSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByHypervSitePager func(resourceGroupName string, siteName string, options *armmigrate.HypervHostControllerClientListByHypervSiteOptions) (resp azfake.PagerResponder[armmigrate.HypervHostControllerClientListByHypervSiteResponse])
}

// NewHypervHostControllerServerTransport creates a new instance of HypervHostControllerServerTransport with the provided implementation.
// The returned HypervHostControllerServerTransport instance is connected to an instance of armmigrate.HypervHostControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervHostControllerServerTransport(srv *HypervHostControllerServer) *HypervHostControllerServerTransport {
	return &HypervHostControllerServerTransport{
		srv:                      srv,
		beginCreate:              newTracker[azfake.PollerResponder[armmigrate.HypervHostControllerClientCreateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armmigrate.HypervHostControllerClientDeleteResponse]](),
		newListByHypervSitePager: newTracker[azfake.PagerResponder[armmigrate.HypervHostControllerClientListByHypervSiteResponse]](),
	}
}

// HypervHostControllerServerTransport connects instances of armmigrate.HypervHostControllerClient to instances of HypervHostControllerServer.
// Don't use this type directly, use NewHypervHostControllerServerTransport instead.
type HypervHostControllerServerTransport struct {
	srv                      *HypervHostControllerServer
	beginCreate              *tracker[azfake.PollerResponder[armmigrate.HypervHostControllerClientCreateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armmigrate.HypervHostControllerClientDeleteResponse]]
	newListByHypervSitePager *tracker[azfake.PagerResponder[armmigrate.HypervHostControllerClientListByHypervSiteResponse]]
}

// Do implements the policy.Transporter interface for HypervHostControllerServerTransport.
func (h *HypervHostControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HypervHostControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if hypervHostControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = hypervHostControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HypervHostControllerClient.BeginCreate":
				res.resp, res.err = h.dispatchBeginCreate(req)
			case "HypervHostControllerClient.BeginDelete":
				res.resp, res.err = h.dispatchBeginDelete(req)
			case "HypervHostControllerClient.Get":
				res.resp, res.err = h.dispatchGet(req)
			case "HypervHostControllerClient.NewListByHypervSitePager":
				res.resp, res.err = h.dispatchNewListByHypervSitePager(req)
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

func (h *HypervHostControllerServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if h.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := h.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.HypervHost](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		hostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginCreate(req.Context(), resourceGroupNameParam, siteNameParam, hostNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		h.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		h.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		h.beginCreate.remove(req)
	}

	return resp, nil
}

func (h *HypervHostControllerServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if h.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := h.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		hostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := h.srv.BeginDelete(req.Context(), resourceGroupNameParam, siteNameParam, hostNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		h.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		h.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		h.beginDelete.remove(req)
	}

	return resp, nil
}

func (h *HypervHostControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts/(?P<hostName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	hostNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("hostName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, hostNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervHost, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervHostControllerServerTransport) dispatchNewListByHypervSitePager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListByHypervSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByHypervSitePager not implemented")}
	}
	newListByHypervSitePager := h.newListByHypervSitePager.get(req)
	if newListByHypervSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/hosts`
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
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrate.HypervHostControllerClientListByHypervSiteOptions
		if filterParam != nil {
			options = &armmigrate.HypervHostControllerClientListByHypervSiteOptions{
				Filter: filterParam,
			}
		}
		resp := h.srv.NewListByHypervSitePager(resourceGroupNameParam, siteNameParam, options)
		newListByHypervSitePager = &resp
		h.newListByHypervSitePager.add(req, newListByHypervSitePager)
		server.PagerResponderInjectNextLinks(newListByHypervSitePager, req, func(page *armmigrate.HypervHostControllerClientListByHypervSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByHypervSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListByHypervSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByHypervSitePager) {
		h.newListByHypervSitePager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to HypervHostControllerServerTransport
var hypervHostControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
