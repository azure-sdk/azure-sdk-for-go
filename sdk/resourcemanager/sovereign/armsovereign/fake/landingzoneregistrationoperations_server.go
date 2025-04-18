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

// LandingZoneRegistrationOperationsServer is a fake server for instances of the armsovereign.LandingZoneRegistrationOperationsClient type.
type LandingZoneRegistrationOperationsServer struct {
	// BeginCreate is the fake for method LandingZoneRegistrationOperationsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, resource armsovereign.LandingZoneRegistrationResource, options *armsovereign.LandingZoneRegistrationOperationsClientBeginCreateOptions) (resp azfake.PollerResponder[armsovereign.LandingZoneRegistrationOperationsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method LandingZoneRegistrationOperationsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, options *armsovereign.LandingZoneRegistrationOperationsClientDeleteOptions) (resp azfake.Responder[armsovereign.LandingZoneRegistrationOperationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LandingZoneRegistrationOperationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, options *armsovereign.LandingZoneRegistrationOperationsClientGetOptions) (resp azfake.Responder[armsovereign.LandingZoneRegistrationOperationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method LandingZoneRegistrationOperationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, landingZoneAccountName string, options *armsovereign.LandingZoneRegistrationOperationsClientListOptions) (resp azfake.PagerResponder[armsovereign.LandingZoneRegistrationOperationsClientListResponse])

	// BeginUpdate is the fake for method LandingZoneRegistrationOperationsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, landingZoneAccountName string, landingZoneRegistrationName string, properties armsovereign.LandingZoneRegistrationResourceUpdate, options *armsovereign.LandingZoneRegistrationOperationsClientBeginUpdateOptions) (resp azfake.PollerResponder[armsovereign.LandingZoneRegistrationOperationsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewLandingZoneRegistrationOperationsServerTransport creates a new instance of LandingZoneRegistrationOperationsServerTransport with the provided implementation.
// The returned LandingZoneRegistrationOperationsServerTransport instance is connected to an instance of armsovereign.LandingZoneRegistrationOperationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLandingZoneRegistrationOperationsServerTransport(srv *LandingZoneRegistrationOperationsServer) *LandingZoneRegistrationOperationsServerTransport {
	return &LandingZoneRegistrationOperationsServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armsovereign.LandingZoneRegistrationOperationsClientCreateResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armsovereign.LandingZoneRegistrationOperationsClientListResponse]](),
		beginUpdate:  newTracker[azfake.PollerResponder[armsovereign.LandingZoneRegistrationOperationsClientUpdateResponse]](),
	}
}

// LandingZoneRegistrationOperationsServerTransport connects instances of armsovereign.LandingZoneRegistrationOperationsClient to instances of LandingZoneRegistrationOperationsServer.
// Don't use this type directly, use NewLandingZoneRegistrationOperationsServerTransport instead.
type LandingZoneRegistrationOperationsServerTransport struct {
	srv          *LandingZoneRegistrationOperationsServer
	beginCreate  *tracker[azfake.PollerResponder[armsovereign.LandingZoneRegistrationOperationsClientCreateResponse]]
	newListPager *tracker[azfake.PagerResponder[armsovereign.LandingZoneRegistrationOperationsClientListResponse]]
	beginUpdate  *tracker[azfake.PollerResponder[armsovereign.LandingZoneRegistrationOperationsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for LandingZoneRegistrationOperationsServerTransport.
func (l *LandingZoneRegistrationOperationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return l.dispatchToMethodFake(req, method)
}

func (l *LandingZoneRegistrationOperationsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if landingZoneRegistrationOperationsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = landingZoneRegistrationOperationsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "LandingZoneRegistrationOperationsClient.BeginCreate":
				res.resp, res.err = l.dispatchBeginCreate(req)
			case "LandingZoneRegistrationOperationsClient.Delete":
				res.resp, res.err = l.dispatchDelete(req)
			case "LandingZoneRegistrationOperationsClient.Get":
				res.resp, res.err = l.dispatchGet(req)
			case "LandingZoneRegistrationOperationsClient.NewListPager":
				res.resp, res.err = l.dispatchNewListPager(req)
			case "LandingZoneRegistrationOperationsClient.BeginUpdate":
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

func (l *LandingZoneRegistrationOperationsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := l.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/landingZoneRegistrations/(?P<landingZoneRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsovereign.LandingZoneRegistrationResource](req)
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
		landingZoneRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneRegistrationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreate(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, landingZoneRegistrationNameParam, body, nil)
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

func (l *LandingZoneRegistrationOperationsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if l.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/landingZoneRegistrations/(?P<landingZoneRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	landingZoneRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneRegistrationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Delete(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, landingZoneRegistrationNameParam, nil)
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

func (l *LandingZoneRegistrationOperationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/landingZoneRegistrations/(?P<landingZoneRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
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
	landingZoneRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneRegistrationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, landingZoneRegistrationNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).LandingZoneRegistrationResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LandingZoneRegistrationOperationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := l.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/landingZoneRegistrations`
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
		resp := l.srv.NewListPager(resourceGroupNameParam, landingZoneAccountNameParam, nil)
		newListPager = &resp
		l.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsovereign.LandingZoneRegistrationOperationsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		l.newListPager.remove(req)
	}
	return resp, nil
}

func (l *LandingZoneRegistrationOperationsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sovereign/landingZoneAccounts/(?P<landingZoneAccountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/landingZoneRegistrations/(?P<landingZoneRegistrationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armsovereign.LandingZoneRegistrationResourceUpdate](req)
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
		landingZoneRegistrationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("landingZoneRegistrationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceGroupNameParam, landingZoneAccountNameParam, landingZoneRegistrationNameParam, body, nil)
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

// set this to conditionally intercept incoming requests to LandingZoneRegistrationOperationsServerTransport
var landingZoneRegistrationOperationsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
