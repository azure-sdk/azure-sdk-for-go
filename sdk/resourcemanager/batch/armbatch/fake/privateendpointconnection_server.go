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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch/v3"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// PrivateEndpointConnectionServer is a fake server for instances of the armbatch.PrivateEndpointConnectionClient type.
type PrivateEndpointConnectionServer struct {
	// BeginDelete is the fake for method PrivateEndpointConnectionClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, options *armbatch.PrivateEndpointConnectionClientBeginDeleteOptions) (resp azfake.PollerResponder[armbatch.PrivateEndpointConnectionClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PrivateEndpointConnectionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, options *armbatch.PrivateEndpointConnectionClientGetOptions) (resp azfake.Responder[armbatch.PrivateEndpointConnectionClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByBatchAccountPager is the fake for method PrivateEndpointConnectionClient.NewListByBatchAccountPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByBatchAccountPager func(resourceGroupName string, accountName string, options *armbatch.PrivateEndpointConnectionClientListByBatchAccountOptions) (resp azfake.PagerResponder[armbatch.PrivateEndpointConnectionClientListByBatchAccountResponse])

	// BeginUpdate is the fake for method PrivateEndpointConnectionClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, accountName string, privateEndpointConnectionName string, parameters armbatch.PrivateEndpointConnection, options *armbatch.PrivateEndpointConnectionClientBeginUpdateOptions) (resp azfake.PollerResponder[armbatch.PrivateEndpointConnectionClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewPrivateEndpointConnectionServerTransport creates a new instance of PrivateEndpointConnectionServerTransport with the provided implementation.
// The returned PrivateEndpointConnectionServerTransport instance is connected to an instance of armbatch.PrivateEndpointConnectionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPrivateEndpointConnectionServerTransport(srv *PrivateEndpointConnectionServer) *PrivateEndpointConnectionServerTransport {
	return &PrivateEndpointConnectionServerTransport{
		srv:                        srv,
		beginDelete:                newTracker[azfake.PollerResponder[armbatch.PrivateEndpointConnectionClientDeleteResponse]](),
		newListByBatchAccountPager: newTracker[azfake.PagerResponder[armbatch.PrivateEndpointConnectionClientListByBatchAccountResponse]](),
		beginUpdate:                newTracker[azfake.PollerResponder[armbatch.PrivateEndpointConnectionClientUpdateResponse]](),
	}
}

// PrivateEndpointConnectionServerTransport connects instances of armbatch.PrivateEndpointConnectionClient to instances of PrivateEndpointConnectionServer.
// Don't use this type directly, use NewPrivateEndpointConnectionServerTransport instead.
type PrivateEndpointConnectionServerTransport struct {
	srv                        *PrivateEndpointConnectionServer
	beginDelete                *tracker[azfake.PollerResponder[armbatch.PrivateEndpointConnectionClientDeleteResponse]]
	newListByBatchAccountPager *tracker[azfake.PagerResponder[armbatch.PrivateEndpointConnectionClientListByBatchAccountResponse]]
	beginUpdate                *tracker[azfake.PollerResponder[armbatch.PrivateEndpointConnectionClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for PrivateEndpointConnectionServerTransport.
func (p *PrivateEndpointConnectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return p.dispatchToMethodFake(req, method)
}

func (p *PrivateEndpointConnectionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if privateEndpointConnectionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = privateEndpointConnectionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "PrivateEndpointConnectionClient.BeginDelete":
				res.resp, res.err = p.dispatchBeginDelete(req)
			case "PrivateEndpointConnectionClient.Get":
				res.resp, res.err = p.dispatchGet(req)
			case "PrivateEndpointConnectionClient.NewListByBatchAccountPager":
				res.resp, res.err = p.dispatchNewListByBatchAccountPager(req)
			case "PrivateEndpointConnectionClient.BeginUpdate":
				res.resp, res.err = p.dispatchBeginUpdate(req)
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

func (p *PrivateEndpointConnectionServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if p.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := p.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := p.srv.BeginDelete(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionNameParam, nil)
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

func (p *PrivateEndpointConnectionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
	if err != nil {
		return nil, err
	}
	privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PrivateEndpointConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionServerTransport) dispatchNewListByBatchAccountPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByBatchAccountPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByBatchAccountPager not implemented")}
	}
	newListByBatchAccountPager := p.newListByBatchAccountPager.get(req)
	if newListByBatchAccountPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		maxresultsUnescaped, err := url.QueryUnescape(qp.Get("maxresults"))
		if err != nil {
			return nil, err
		}
		maxresultsParam, err := parseOptional(maxresultsUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armbatch.PrivateEndpointConnectionClientListByBatchAccountOptions
		if maxresultsParam != nil {
			options = &armbatch.PrivateEndpointConnectionClientListByBatchAccountOptions{
				Maxresults: maxresultsParam,
			}
		}
		resp := p.srv.NewListByBatchAccountPager(resourceGroupNameParam, accountNameParam, options)
		newListByBatchAccountPager = &resp
		p.newListByBatchAccountPager.add(req, newListByBatchAccountPager)
		server.PagerResponderInjectNextLinks(newListByBatchAccountPager, req, func(page *armbatch.PrivateEndpointConnectionClientListByBatchAccountResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByBatchAccountPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByBatchAccountPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByBatchAccountPager) {
		p.newListByBatchAccountPager.remove(req)
	}
	return resp, nil
}

func (p *PrivateEndpointConnectionServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := p.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Batch/batchAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/privateEndpointConnections/(?P<privateEndpointConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armbatch.PrivateEndpointConnection](req)
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
		privateEndpointConnectionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateEndpointConnectionName")])
		if err != nil {
			return nil, err
		}
		ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
		var options *armbatch.PrivateEndpointConnectionClientBeginUpdateOptions
		if ifMatchParam != nil {
			options = &armbatch.PrivateEndpointConnectionClientBeginUpdateOptions{
				IfMatch: ifMatchParam,
			}
		}
		respr, errRespr := p.srv.BeginUpdate(req.Context(), resourceGroupNameParam, accountNameParam, privateEndpointConnectionNameParam, body, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		p.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		p.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		p.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to PrivateEndpointConnectionServerTransport
var privateEndpointConnectionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
