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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// HuntsServer is a fake server for instances of the armsecurityinsights.HuntsClient type.
type HuntsServer struct {
	// CreateOrUpdate is the fake for method HuntsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, hunt armsecurityinsights.Hunt, options *armsecurityinsights.HuntsClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurityinsights.HuntsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HuntsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, options *armsecurityinsights.HuntsClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.HuntsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HuntsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, options *armsecurityinsights.HuntsClientGetOptions) (resp azfake.Responder[armsecurityinsights.HuntsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method HuntsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.HuntsClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.HuntsClientListResponse])
}

// NewHuntsServerTransport creates a new instance of HuntsServerTransport with the provided implementation.
// The returned HuntsServerTransport instance is connected to an instance of armsecurityinsights.HuntsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHuntsServerTransport(srv *HuntsServer) *HuntsServerTransport {
	return &HuntsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.HuntsClientListResponse]](),
	}
}

// HuntsServerTransport connects instances of armsecurityinsights.HuntsClient to instances of HuntsServer.
// Don't use this type directly, use NewHuntsServerTransport instead.
type HuntsServerTransport struct {
	srv          *HuntsServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.HuntsClientListResponse]]
}

// Do implements the policy.Transporter interface for HuntsServerTransport.
func (h *HuntsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HuntsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if huntsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = huntsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HuntsClient.CreateOrUpdate":
				res.resp, res.err = h.dispatchCreateOrUpdate(req)
			case "HuntsClient.Delete":
				res.resp, res.err = h.dispatchDelete(req)
			case "HuntsClient.Get":
				res.resp, res.err = h.dispatchGet(req)
			case "HuntsClient.NewListPager":
				res.resp, res.err = h.dispatchNewListPager(req)
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

func (h *HuntsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.Hunt](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, huntIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Hunt, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HuntsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, huntIDParam, nil)
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

func (h *HuntsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts/(?P<huntId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	huntIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("huntId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, huntIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Hunt, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HuntsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if h.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := h.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/hunts`
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
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
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
		var options *armsecurityinsights.HuntsClientListOptions
		if filterParam != nil || orderbyParam != nil || topParam != nil || skipTokenParam != nil {
			options = &armsecurityinsights.HuntsClientListOptions{
				Filter:    filterParam,
				Orderby:   orderbyParam,
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := h.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		h.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.HuntsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		h.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		h.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to HuntsServerTransport
var huntsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
