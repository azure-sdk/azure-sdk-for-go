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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ReplicationProtectionIntentsServer is a fake server for instances of the armrecoveryservicessiterecovery.ReplicationProtectionIntentsClient type.
type ReplicationProtectionIntentsServer struct {
	// Create is the fake for method ReplicationProtectionIntentsClient.Create
	// HTTP status codes to indicate success: http.StatusOK
	Create func(ctx context.Context, resourceGroupName string, resourceName string, intentObjectName string, input armrecoveryservicessiterecovery.CreateProtectionIntentInput, options *armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientCreateOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ReplicationProtectionIntentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, intentObjectName string, options *armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientGetOptions) (resp azfake.Responder[armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ReplicationProtectionIntentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, resourceName string, options *armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListResponse])
}

// NewReplicationProtectionIntentsServerTransport creates a new instance of ReplicationProtectionIntentsServerTransport with the provided implementation.
// The returned ReplicationProtectionIntentsServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ReplicationProtectionIntentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewReplicationProtectionIntentsServerTransport(srv *ReplicationProtectionIntentsServer) *ReplicationProtectionIntentsServerTransport {
	return &ReplicationProtectionIntentsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListResponse]](),
	}
}

// ReplicationProtectionIntentsServerTransport connects instances of armrecoveryservicessiterecovery.ReplicationProtectionIntentsClient to instances of ReplicationProtectionIntentsServer.
// Don't use this type directly, use NewReplicationProtectionIntentsServerTransport instead.
type ReplicationProtectionIntentsServerTransport struct {
	srv          *ReplicationProtectionIntentsServer
	newListPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListResponse]]
}

// Do implements the policy.Transporter interface for ReplicationProtectionIntentsServerTransport.
func (r *ReplicationProtectionIntentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ReplicationProtectionIntentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if replicationProtectionIntentsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = replicationProtectionIntentsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ReplicationProtectionIntentsClient.Create":
				res.resp, res.err = r.dispatchCreate(req)
			case "ReplicationProtectionIntentsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ReplicationProtectionIntentsClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
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

func (r *ReplicationProtectionIntentsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if r.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionIntents/(?P<intentObjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicessiterecovery.CreateProtectionIntentInput](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	intentObjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("intentObjectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Create(req.Context(), resourceGroupNameParam, resourceNameParam, intentObjectNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReplicationProtectionIntent, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationProtectionIntentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionIntents/(?P<intentObjectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	intentObjectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("intentObjectName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, intentObjectNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReplicationProtectionIntent, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ReplicationProtectionIntentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionIntents`
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
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		takeTokenUnescaped, err := url.QueryUnescape(qp.Get("takeToken"))
		if err != nil {
			return nil, err
		}
		takeTokenParam := getOptional(takeTokenUnescaped)
		var options *armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListOptions
		if skipTokenParam != nil || takeTokenParam != nil {
			options = &armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListOptions{
				SkipToken: skipTokenParam,
				TakeToken: takeTokenParam,
			}
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, resourceNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ReplicationProtectionIntentsServerTransport
var replicationProtectionIntentsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
