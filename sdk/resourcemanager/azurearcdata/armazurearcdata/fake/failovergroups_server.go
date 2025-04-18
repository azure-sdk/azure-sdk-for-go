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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurearcdata/armazurearcdata"
	"net/http"
	"net/url"
	"regexp"
)

// FailoverGroupsServer is a fake server for instances of the armazurearcdata.FailoverGroupsClient type.
type FailoverGroupsServer struct {
	// BeginCreate is the fake for method FailoverGroupsClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, failoverGroupResource armazurearcdata.FailoverGroupResource, options *armazurearcdata.FailoverGroupsClientBeginCreateOptions) (resp azfake.PollerResponder[armazurearcdata.FailoverGroupsClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FailoverGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *armazurearcdata.FailoverGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armazurearcdata.FailoverGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FailoverGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, sqlManagedInstanceName string, failoverGroupName string, options *armazurearcdata.FailoverGroupsClientGetOptions) (resp azfake.Responder[armazurearcdata.FailoverGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FailoverGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, sqlManagedInstanceName string, options *armazurearcdata.FailoverGroupsClientListOptions) (resp azfake.PagerResponder[armazurearcdata.FailoverGroupsClientListResponse])
}

// NewFailoverGroupsServerTransport creates a new instance of FailoverGroupsServerTransport with the provided implementation.
// The returned FailoverGroupsServerTransport instance is connected to an instance of armazurearcdata.FailoverGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFailoverGroupsServerTransport(srv *FailoverGroupsServer) *FailoverGroupsServerTransport {
	return &FailoverGroupsServerTransport{
		srv:          srv,
		beginCreate:  newTracker[azfake.PollerResponder[armazurearcdata.FailoverGroupsClientCreateResponse]](),
		beginDelete:  newTracker[azfake.PollerResponder[armazurearcdata.FailoverGroupsClientDeleteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armazurearcdata.FailoverGroupsClientListResponse]](),
	}
}

// FailoverGroupsServerTransport connects instances of armazurearcdata.FailoverGroupsClient to instances of FailoverGroupsServer.
// Don't use this type directly, use NewFailoverGroupsServerTransport instead.
type FailoverGroupsServerTransport struct {
	srv          *FailoverGroupsServer
	beginCreate  *tracker[azfake.PollerResponder[armazurearcdata.FailoverGroupsClientCreateResponse]]
	beginDelete  *tracker[azfake.PollerResponder[armazurearcdata.FailoverGroupsClientDeleteResponse]]
	newListPager *tracker[azfake.PagerResponder[armazurearcdata.FailoverGroupsClientListResponse]]
}

// Do implements the policy.Transporter interface for FailoverGroupsServerTransport.
func (f *FailoverGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FailoverGroupsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if failoverGroupsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = failoverGroupsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FailoverGroupsClient.BeginCreate":
				res.resp, res.err = f.dispatchBeginCreate(req)
			case "FailoverGroupsClient.BeginDelete":
				res.resp, res.err = f.dispatchBeginDelete(req)
			case "FailoverGroupsClient.Get":
				res.resp, res.err = f.dispatchGet(req)
			case "FailoverGroupsClient.NewListPager":
				res.resp, res.err = f.dispatchNewListPager(req)
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

func (f *FailoverGroupsServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := f.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/sqlManagedInstances/(?P<sqlManagedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failoverGroups/(?P<failoverGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armazurearcdata.FailoverGroupResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlManagedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlManagedInstanceName")])
		if err != nil {
			return nil, err
		}
		failoverGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("failoverGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreate(req.Context(), resourceGroupNameParam, sqlManagedInstanceNameParam, failoverGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		f.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		f.beginCreate.remove(req)
	}

	return resp, nil
}

func (f *FailoverGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/sqlManagedInstances/(?P<sqlManagedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failoverGroups/(?P<failoverGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlManagedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlManagedInstanceName")])
		if err != nil {
			return nil, err
		}
		failoverGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("failoverGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, sqlManagedInstanceNameParam, failoverGroupNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FailoverGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/sqlManagedInstances/(?P<sqlManagedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failoverGroups/(?P<failoverGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	sqlManagedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlManagedInstanceName")])
	if err != nil {
		return nil, err
	}
	failoverGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("failoverGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, sqlManagedInstanceNameParam, failoverGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FailoverGroupResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FailoverGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AzureArcData/sqlManagedInstances/(?P<sqlManagedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/failoverGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlManagedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlManagedInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListPager(resourceGroupNameParam, sqlManagedInstanceNameParam, nil)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armazurearcdata.FailoverGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to FailoverGroupsServerTransport
var failoverGroupsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
