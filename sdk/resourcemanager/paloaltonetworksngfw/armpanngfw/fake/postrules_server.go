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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/paloaltonetworksngfw/armpanngfw"
	"net/http"
	"net/url"
	"regexp"
)

// PostRulesServer is a fake server for instances of the armpanngfw.PostRulesClient type.
type PostRulesServer struct {
	// BeginCreateOrUpdate is the fake for method PostRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, globalRulestackName string, priority string, resource armpanngfw.PostRulesResource, options *armpanngfw.PostRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armpanngfw.PostRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method PostRulesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, globalRulestackName string, priority string, options *armpanngfw.PostRulesClientBeginDeleteOptions) (resp azfake.PollerResponder[armpanngfw.PostRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PostRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, globalRulestackName string, priority string, options *armpanngfw.PostRulesClientGetOptions) (resp azfake.Responder[armpanngfw.PostRulesClientGetResponse], errResp azfake.ErrorResponder)

	// GetCounters is the fake for method PostRulesClient.GetCounters
	// HTTP status codes to indicate success: http.StatusOK
	GetCounters func(ctx context.Context, globalRulestackName string, priority string, options *armpanngfw.PostRulesClientGetCountersOptions) (resp azfake.Responder[armpanngfw.PostRulesClientGetCountersResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method PostRulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(globalRulestackName string, options *armpanngfw.PostRulesClientListOptions) (resp azfake.PagerResponder[armpanngfw.PostRulesClientListResponse])

	// RefreshCounters is the fake for method PostRulesClient.RefreshCounters
	// HTTP status codes to indicate success: http.StatusNoContent
	RefreshCounters func(ctx context.Context, globalRulestackName string, priority string, options *armpanngfw.PostRulesClientRefreshCountersOptions) (resp azfake.Responder[armpanngfw.PostRulesClientRefreshCountersResponse], errResp azfake.ErrorResponder)

	// ResetCounters is the fake for method PostRulesClient.ResetCounters
	// HTTP status codes to indicate success: http.StatusOK
	ResetCounters func(ctx context.Context, globalRulestackName string, priority string, options *armpanngfw.PostRulesClientResetCountersOptions) (resp azfake.Responder[armpanngfw.PostRulesClientResetCountersResponse], errResp azfake.ErrorResponder)
}

// NewPostRulesServerTransport creates a new instance of PostRulesServerTransport with the provided implementation.
// The returned PostRulesServerTransport instance is connected to an instance of armpanngfw.PostRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPostRulesServerTransport(srv *PostRulesServer) *PostRulesServerTransport {
	return &PostRulesServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armpanngfw.PostRulesClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armpanngfw.PostRulesClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armpanngfw.PostRulesClientListResponse]](),
	}
}

// PostRulesServerTransport connects instances of armpanngfw.PostRulesClient to instances of PostRulesServer.
// Don't use this type directly, use NewPostRulesServerTransport instead.
type PostRulesServerTransport struct {
	srv                 *PostRulesServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armpanngfw.PostRulesClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armpanngfw.PostRulesClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armpanngfw.PostRulesClientListResponse]]
}

// Do implements the policy.Transporter interface for PostRulesServerTransport.
func (p *PostRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PostRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if postRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = postRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PostRulesClient.BeginCreateOrUpdate":
				res.resp, res.err = p.dispatchBeginCreateOrUpdate(req)
			case "PostRulesClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PostRulesClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PostRulesClient.GetCounters":
				res.resp, res.err = p.dispatchGetCounters(req)
			case "PostRulesClient.NewListPager":
				res.resp, res.err = p.dispatchNewListPager(req)
			case "PostRulesClient.RefreshCounters":
				res.resp, res.err = p.dispatchRefreshCounters(req)
			case "PostRulesClient.ResetCounters":
				res.resp, res.err = p.dispatchResetCounters(req)
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

func (p *PostRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := p.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armpanngfw.PostRulesResource](req)
		if err != nil {
			return nil, err
		}
		globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
		if err != nil {
			return nil, err
		}
		priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginCreateOrUpdate(req.Context(), globalRulestackNameParam, priorityParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		p.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		p.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		p.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (p *PostRulesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
		if err != nil {
			return nil, err
		}
		priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), globalRulestackNameParam, priorityParam, nil)
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

func (p *PostRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), globalRulestackNameParam, priorityParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PostRulesResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PostRulesServerTransport) dispatchGetCounters(req *http.Request) (*http.Response, error) {
	if p.srv.GetCounters == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetCounters not implemented")}
	}
	const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getCounters`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	firewallNameUnescaped, err := url.QueryUnescape(qp.Get("firewallName"))
	if err != nil {
		return nil, err
	}
	firewallNameParam := getOptional(firewallNameUnescaped)
	var options *armpanngfw.PostRulesClientGetCountersOptions
	if firewallNameParam != nil {
		options = &armpanngfw.PostRulesClientGetCountersOptions{
			FirewallName: firewallNameParam,
		}
	}
	respr, errRespr := p.srv.GetCounters(req.Context(), globalRulestackNameParam, priorityParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RuleCounter, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PostRulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := p.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListPager(globalRulestackNameParam, nil)
		newListPager = &resp
		p.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armpanngfw.PostRulesClientListResponse, createLink func() string) {
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

func (p *PostRulesServerTransport) dispatchRefreshCounters(req *http.Request) (*http.Response, error) {
	if p.srv.RefreshCounters == nil {
		return nil, &nonRetriableError{errors.New("fake for method RefreshCounters not implemented")}
	}
	const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshCounters`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	firewallNameUnescaped, err := url.QueryUnescape(qp.Get("firewallName"))
	if err != nil {
		return nil, err
	}
	firewallNameParam := getOptional(firewallNameUnescaped)
	var options *armpanngfw.PostRulesClientRefreshCountersOptions
	if firewallNameParam != nil {
		options = &armpanngfw.PostRulesClientRefreshCountersOptions{
			FirewallName: firewallNameParam,
		}
	}
	respr, errRespr := p.srv.RefreshCounters(req.Context(), globalRulestackNameParam, priorityParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PostRulesServerTransport) dispatchResetCounters(req *http.Request) (*http.Response, error) {
	if p.srv.ResetCounters == nil {
		return nil, &nonRetriableError{errors.New("fake for method ResetCounters not implemented")}
	}
	const regexStr = `/providers/PaloAltoNetworks\.Cloudngfw/globalRulestacks/(?P<globalRulestackName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/postRules/(?P<priority>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resetCounters`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	globalRulestackNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("globalRulestackName")])
	if err != nil {
		return nil, err
	}
	priorityParam, err := url.PathUnescape(matches[regex.SubexpIndex("priority")])
	if err != nil {
		return nil, err
	}
	firewallNameUnescaped, err := url.QueryUnescape(qp.Get("firewallName"))
	if err != nil {
		return nil, err
	}
	firewallNameParam := getOptional(firewallNameUnescaped)
	var options *armpanngfw.PostRulesClientResetCountersOptions
	if firewallNameParam != nil {
		options = &armpanngfw.PostRulesClientResetCountersOptions{
			FirewallName: firewallNameParam,
		}
	}
	respr, errRespr := p.srv.ResetCounters(req.Context(), globalRulestackNameParam, priorityParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RuleCounterReset, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to PostRulesServerTransport
var postRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
