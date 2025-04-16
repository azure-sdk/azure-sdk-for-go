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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v4"
	"net/http"
	"net/url"
	"regexp"
)

// ServerlessEndpointsServer is a fake server for instances of the armmachinelearning.ServerlessEndpointsClient type.
type ServerlessEndpointsServer struct {
	// BeginCreateOrUpdate is the fake for method ServerlessEndpointsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, name string, body armmachinelearning.ServerlessEndpoint, options *armmachinelearning.ServerlessEndpointsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ServerlessEndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armmachinelearning.ServerlessEndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServerlessEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armmachinelearning.ServerlessEndpointsClientGetOptions) (resp azfake.Responder[armmachinelearning.ServerlessEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ServerlessEndpointsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armmachinelearning.ServerlessEndpointsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.ServerlessEndpointsClientListResponse])

	// ListKeys is the fake for method ServerlessEndpointsClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *armmachinelearning.ServerlessEndpointsClientListKeysOptions) (resp azfake.Responder[armmachinelearning.ServerlessEndpointsClientListKeysResponse], errResp azfake.ErrorResponder)

	// BeginRegenerateKeys is the fake for method ServerlessEndpointsClient.BeginRegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRegenerateKeys func(ctx context.Context, resourceGroupName string, workspaceName string, name string, body armmachinelearning.RegenerateEndpointKeysRequest, options *armmachinelearning.ServerlessEndpointsClientBeginRegenerateKeysOptions) (resp azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientRegenerateKeysResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method ServerlessEndpointsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, name string, body armmachinelearning.PartialMinimalTrackedResourceWithSKUAndIdentity, options *armmachinelearning.ServerlessEndpointsClientBeginUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewServerlessEndpointsServerTransport creates a new instance of ServerlessEndpointsServerTransport with the provided implementation.
// The returned ServerlessEndpointsServerTransport instance is connected to an instance of armmachinelearning.ServerlessEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerlessEndpointsServerTransport(srv *ServerlessEndpointsServer) *ServerlessEndpointsServerTransport {
	return &ServerlessEndpointsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmachinelearning.ServerlessEndpointsClientListResponse]](),
		beginRegenerateKeys: newTracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientRegenerateKeysResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientUpdateResponse]](),
	}
}

// ServerlessEndpointsServerTransport connects instances of armmachinelearning.ServerlessEndpointsClient to instances of ServerlessEndpointsServer.
// Don't use this type directly, use NewServerlessEndpointsServerTransport instead.
type ServerlessEndpointsServerTransport struct {
	srv                 *ServerlessEndpointsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmachinelearning.ServerlessEndpointsClientListResponse]]
	beginRegenerateKeys *tracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientRegenerateKeysResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armmachinelearning.ServerlessEndpointsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ServerlessEndpointsServerTransport.
func (s *ServerlessEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *ServerlessEndpointsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if serverlessEndpointsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = serverlessEndpointsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ServerlessEndpointsClient.BeginCreateOrUpdate":
				res.resp, res.err = s.dispatchBeginCreateOrUpdate(req)
			case "ServerlessEndpointsClient.BeginDelete":
				res.resp, res.err = s.dispatchBeginDelete(req)
			case "ServerlessEndpointsClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "ServerlessEndpointsClient.NewListPager":
				res.resp, res.err = s.dispatchNewListPager(req)
			case "ServerlessEndpointsClient.ListKeys":
				res.resp, res.err = s.dispatchListKeys(req)
			case "ServerlessEndpointsClient.BeginRegenerateKeys":
				res.resp, res.err = s.dispatchBeginRegenerateKeys(req)
			case "ServerlessEndpointsClient.BeginUpdate":
				res.resp, res.err = s.dispatchBeginUpdate(req)
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

func (s *ServerlessEndpointsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := s.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.ServerlessEndpoint](req)
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
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		s.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		s.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (s *ServerlessEndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *ServerlessEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerlessEndpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerlessEndpointsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints`
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		var options *armmachinelearning.ServerlessEndpointsClientListOptions
		if skipParam != nil {
			options = &armmachinelearning.ServerlessEndpointsClientListOptions{
				Skip: skipParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.ServerlessEndpointsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *ServerlessEndpointsServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if s.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.ListKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EndpointAuthKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerlessEndpointsServerTransport) dispatchBeginRegenerateKeys(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRegenerateKeys not implemented")}
	}
	beginRegenerateKeys := s.beginRegenerateKeys.get(req)
	if beginRegenerateKeys == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKeys`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.RegenerateEndpointKeysRequest](req)
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
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginRegenerateKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRegenerateKeys = &respr
		s.beginRegenerateKeys.add(req, beginRegenerateKeys)
	}

	resp, err := server.PollerResponderNext(beginRegenerateKeys, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginRegenerateKeys.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRegenerateKeys) {
		s.beginRegenerateKeys.remove(req)
	}

	return resp, nil
}

func (s *ServerlessEndpointsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := s.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/serverlessEndpoints/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.PartialMinimalTrackedResourceWithSKUAndIdentity](req)
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
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		s.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		s.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ServerlessEndpointsServerTransport
var serverlessEndpointsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
