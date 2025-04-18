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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices"
	"net/http"
	"net/url"
	"regexp"
)

// DeletedAccountsServer is a fake server for instances of the armcognitiveservices.DeletedAccountsClient type.
type DeletedAccountsServer struct {
	// Get is the fake for method DeletedAccountsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, resourceGroupName string, accountName string, options *armcognitiveservices.DeletedAccountsClientGetOptions) (resp azfake.Responder[armcognitiveservices.DeletedAccountsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeletedAccountsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcognitiveservices.DeletedAccountsClientListOptions) (resp azfake.PagerResponder[armcognitiveservices.DeletedAccountsClientListResponse])

	// BeginPurge is the fake for method DeletedAccountsClient.BeginPurge
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginPurge func(ctx context.Context, location string, resourceGroupName string, accountName string, options *armcognitiveservices.DeletedAccountsClientBeginPurgeOptions) (resp azfake.PollerResponder[armcognitiveservices.DeletedAccountsClientPurgeResponse], errResp azfake.ErrorResponder)
}

// NewDeletedAccountsServerTransport creates a new instance of DeletedAccountsServerTransport with the provided implementation.
// The returned DeletedAccountsServerTransport instance is connected to an instance of armcognitiveservices.DeletedAccountsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeletedAccountsServerTransport(srv *DeletedAccountsServer) *DeletedAccountsServerTransport {
	return &DeletedAccountsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcognitiveservices.DeletedAccountsClientListResponse]](),
		beginPurge:   newTracker[azfake.PollerResponder[armcognitiveservices.DeletedAccountsClientPurgeResponse]](),
	}
}

// DeletedAccountsServerTransport connects instances of armcognitiveservices.DeletedAccountsClient to instances of DeletedAccountsServer.
// Don't use this type directly, use NewDeletedAccountsServerTransport instead.
type DeletedAccountsServerTransport struct {
	srv          *DeletedAccountsServer
	newListPager *tracker[azfake.PagerResponder[armcognitiveservices.DeletedAccountsClientListResponse]]
	beginPurge   *tracker[azfake.PollerResponder[armcognitiveservices.DeletedAccountsClientPurgeResponse]]
}

// Do implements the policy.Transporter interface for DeletedAccountsServerTransport.
func (d *DeletedAccountsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DeletedAccountsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deletedAccountsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deletedAccountsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeletedAccountsClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DeletedAccountsClient.NewListPager":
				res.resp, res.err = d.dispatchNewListPager(req)
			case "DeletedAccountsClient.BeginPurge":
				res.resp, res.err = d.dispatchBeginPurge(req)
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

func (d *DeletedAccountsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), locationParam, resourceGroupNameParam, accountNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Account, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeletedAccountsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/deletedAccounts`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcognitiveservices.DeletedAccountsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DeletedAccountsServerTransport) dispatchBeginPurge(req *http.Request) (*http.Response, error) {
	if d.srv.BeginPurge == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPurge not implemented")}
	}
	beginPurge := d.beginPurge.get(req)
	if beginPurge == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.CognitiveServices/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginPurge(req.Context(), locationParam, resourceGroupNameParam, accountNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPurge = &respr
		d.beginPurge.add(req, beginPurge)
	}

	resp, err := server.PollerResponderNext(beginPurge, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginPurge.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPurge) {
		d.beginPurge.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to DeletedAccountsServerTransport
var deletedAccountsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
