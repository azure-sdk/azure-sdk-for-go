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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// EurekaServersServer is a fake server for instances of the armappplatform.EurekaServersClient type.
type EurekaServersServer struct {
	// Get is the fake for method EurekaServersClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, options *armappplatform.EurekaServersClientGetOptions) (resp azfake.Responder[armappplatform.EurekaServersClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method EurekaServersClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, resourceGroupName string, serviceName string, options *armappplatform.EurekaServersClientListOptions) (resp azfake.Responder[armappplatform.EurekaServersClientListResponse], errResp azfake.ErrorResponder)

	// BeginUpdatePatch is the fake for method EurekaServersClient.BeginUpdatePatch
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdatePatch func(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource armappplatform.EurekaServerResource, options *armappplatform.EurekaServersClientBeginUpdatePatchOptions) (resp azfake.PollerResponder[armappplatform.EurekaServersClientUpdatePatchResponse], errResp azfake.ErrorResponder)

	// BeginUpdatePut is the fake for method EurekaServersClient.BeginUpdatePut
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginUpdatePut func(ctx context.Context, resourceGroupName string, serviceName string, eurekaServerResource armappplatform.EurekaServerResource, options *armappplatform.EurekaServersClientBeginUpdatePutOptions) (resp azfake.PollerResponder[armappplatform.EurekaServersClientUpdatePutResponse], errResp azfake.ErrorResponder)
}

// NewEurekaServersServerTransport creates a new instance of EurekaServersServerTransport with the provided implementation.
// The returned EurekaServersServerTransport instance is connected to an instance of armappplatform.EurekaServersClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEurekaServersServerTransport(srv *EurekaServersServer) *EurekaServersServerTransport {
	return &EurekaServersServerTransport{
		srv:              srv,
		beginUpdatePatch: newTracker[azfake.PollerResponder[armappplatform.EurekaServersClientUpdatePatchResponse]](),
		beginUpdatePut:   newTracker[azfake.PollerResponder[armappplatform.EurekaServersClientUpdatePutResponse]](),
	}
}

// EurekaServersServerTransport connects instances of armappplatform.EurekaServersClient to instances of EurekaServersServer.
// Don't use this type directly, use NewEurekaServersServerTransport instead.
type EurekaServersServerTransport struct {
	srv              *EurekaServersServer
	beginUpdatePatch *tracker[azfake.PollerResponder[armappplatform.EurekaServersClientUpdatePatchResponse]]
	beginUpdatePut   *tracker[azfake.PollerResponder[armappplatform.EurekaServersClientUpdatePutResponse]]
}

// Do implements the policy.Transporter interface for EurekaServersServerTransport.
func (e *EurekaServersServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return e.dispatchToMethodFake(req, method)
}

func (e *EurekaServersServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if eurekaServersServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = eurekaServersServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "EurekaServersClient.Get":
				res.resp, res.err = e.dispatchGet(req)
			case "EurekaServersClient.List":
				res.resp, res.err = e.dispatchList(req)
			case "EurekaServersClient.BeginUpdatePatch":
				res.resp, res.err = e.dispatchBeginUpdatePatch(req)
			case "EurekaServersClient.BeginUpdatePut":
				res.resp, res.err = e.dispatchBeginUpdatePut(req)
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

func (e *EurekaServersServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eurekaServers/default`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EurekaServerResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EurekaServersServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if e.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eurekaServers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.List(req.Context(), resourceGroupNameParam, serviceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EurekaServerResourceCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EurekaServersServerTransport) dispatchBeginUpdatePatch(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdatePatch == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdatePatch not implemented")}
	}
	beginUpdatePatch := e.beginUpdatePatch.get(req)
	if beginUpdatePatch == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eurekaServers/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.EurekaServerResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdatePatch(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdatePatch = &respr
		e.beginUpdatePatch.add(req, beginUpdatePatch)
	}

	resp, err := server.PollerResponderNext(beginUpdatePatch, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginUpdatePatch.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdatePatch) {
		e.beginUpdatePatch.remove(req)
	}

	return resp, nil
}

func (e *EurekaServersServerTransport) dispatchBeginUpdatePut(req *http.Request) (*http.Response, error) {
	if e.srv.BeginUpdatePut == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdatePut not implemented")}
	}
	beginUpdatePut := e.beginUpdatePut.get(req)
	if beginUpdatePut == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/eurekaServers/default`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.EurekaServerResource](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginUpdatePut(req.Context(), resourceGroupNameParam, serviceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdatePut = &respr
		e.beginUpdatePut.add(req, beginUpdatePut)
	}

	resp, err := server.PollerResponderNext(beginUpdatePut, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		e.beginUpdatePut.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdatePut) {
		e.beginUpdatePut.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to EurekaServersServerTransport
var eurekaServersServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
