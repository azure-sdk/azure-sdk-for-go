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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sovereign/armsovereign"
	"net/http"
	"net/url"
	"regexp"
)

// LandingZoneAccountOperationsServer is a fake server for instances of the armsovereign.LandingZoneAccountOperationsClient type.
type LandingZoneAccountOperationsServer struct {
	// BeginCreate is the fake for method LandingZoneAccountOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, resource armsovereign.LandingZoneAccountResource, options *armsovereign.LandingZoneAccountOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LandingZoneAccountOperationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, options *armsovereign.LandingZoneAccountOperationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LandingZoneAccountOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, options *armsovereign.LandingZoneAccountOperationsClientGetOptions) (resp azfake.Responder[armsovereign.LandingZoneAccountOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method LandingZoneAccountOperationsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armsovereign.LandingZoneAccountOperationsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armsovereign.LandingZoneAccountOperationsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method LandingZoneAccountOperationsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armsovereign.LandingZoneAccountOperationsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armsovereign.LandingZoneAccountOperationsClientListBySubscriptionResponse])

	// BeginUpdate is the fake for method LandingZoneAccountOperationsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, properties armsovereign.LandingZoneAccountResourceUpdate, options *armsovereign.LandingZoneAccountOperationsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewLandingZoneAccountOperationsServerTransport creates a new instance of LandingZoneAccountOperationsServerTransport with the provided implementation.
// The returned LandingZoneAccountOperationsServerTransport instance is connected to an instance of armsovereign.LandingZoneAccountOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLandingZoneAccountOperationsServerTransport(srv *LandingZoneAccountOperationsServer) *LandingZoneAccountOperationsServerTransport {
	return &LandingZoneAccountOperationsServerTransport{
		srv:                         srv,
		beginCreate:                 newTracker[azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientCreateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armsovereign.LandingZoneAccountOperationsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armsovereign.LandingZoneAccountOperationsClientListBySubscriptionResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientUpdateResponse]](),
	}
}

// LandingZoneAccountOperationsServerTransport connects instances of armsovereign.LandingZoneAccountOperationsClient to instances of LandingZoneAccountOperationsServer.
// Don't use this type directly, use NewLandingZoneAccountOperationsServerTransport instead.
type LandingZoneAccountOperationsServerTransport struct {
	srv                         *LandingZoneAccountOperationsServer
	beginCreate                 *tracker[azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientCreateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armsovereign.LandingZoneAccountOperationsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armsovereign.LandingZoneAccountOperationsClientListBySubscriptionResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armsovereign.LandingZoneAccountOperationsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for LandingZoneAccountOperationsServerTransport.
func (l *LandingZoneAccountOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LandingZoneAccountOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if landingZoneAccountOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = landingZoneAccountOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LandingZoneAccountOperationsClient.BeginCreate":
				res.resp, res.err = l.dispatchBeginCreate(req)
			case "LandingZoneAccountOperationsClient.BeginDelete":
				res.resp, res.err = l.dispatchBeginDelete(req)
			case "LandingZoneAccountOperationsClient.Get":
				res.resp, res.err = l.dispatchGet(req)
			case "LandingZoneAccountOperationsClient.NewListByResourceGroupPager":
				res.resp, res.err = l.dispatchNewListByResourceGroupPager(req)
			case "LandingZoneAccountOperationsClient.NewListBySubscriptionPager":
				res.resp, res.err = l.dispatchNewListBySubscriptionPager(req)
			case "LandingZoneAccountOperationsClient.BeginUpdate":
				res.resp, res.err = l.dispatchBeginUpdate(req)
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

func (l *LandingZoneAccountOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := l.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsovereign.LandingZoneAccountResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		landingZoneAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneAccountName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreate(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		l.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		l.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		l.beginCreate.remove(req)
	}

	return resp, nil
}

func (l *LandingZoneAccountOperationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		landingZoneAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneAccountName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LandingZoneAccountOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	landingZoneAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneAccountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LandingZoneAccountResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LandingZoneAccountOperationsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := l.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		l.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armsovereign.LandingZoneAccountOperationsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		l.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (l *LandingZoneAccountOperationsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := l.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := l.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		l.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armsovereign.LandingZoneAccountOperationsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		l.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (l *LandingZoneAccountOperationsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsovereign.LandingZoneAccountResourceUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		landingZoneAccountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneAccountName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		l.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		l.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to LandingZoneAccountOperationsServerTransport
var landingZoneAccountOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
