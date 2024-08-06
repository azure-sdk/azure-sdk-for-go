//go:build go1.18
// +build go1.18

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
	"strconv"
)

// EndpointServer is a fake server for instances of the armmachinelearning.EndpointClient type.
type EndpointServer struct {
	// BeginCreateOrUpdate is the fake for method EndpointClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body armmachinelearning.EndpointResourcePropertiesBasicResource, options *armmachinelearning.EndpointClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.EndpointClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method EndpointClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *armmachinelearning.EndpointClientGetOptions) (resp azfake.Responder[armmachinelearning.EndpointClientGetResponse], errResp azfake.ErrorResponder)

	// NewGetModelsPager is the fake for method EndpointClient.NewGetModelsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewGetModelsPager func(resourceGroupName string, workspaceName string, endpointName string, options *armmachinelearning.EndpointClientGetModelsOptions) (resp azfake.PagerResponder[armmachinelearning.EndpointClientGetModelsResponse])

	// NewListPager is the fake for method EndpointClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armmachinelearning.EndpointClientListOptions) (resp azfake.PagerResponder[armmachinelearning.EndpointClientListResponse])

	// ListKeys is the fake for method EndpointClient.ListKeys
	// HTTP status codes to indicate success: http.StatusOK
	ListKeys func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, options *armmachinelearning.EndpointClientListKeysOptions) (resp azfake.Responder[armmachinelearning.EndpointClientListKeysResponse], errResp azfake.ErrorResponder)

	// RegenerateKeys is the fake for method EndpointClient.RegenerateKeys
	// HTTP status codes to indicate success: http.StatusOK
	RegenerateKeys func(ctx context.Context, resourceGroupName string, workspaceName string, endpointName string, body armmachinelearning.RegenerateServiceAccountKeyContent, options *armmachinelearning.EndpointClientRegenerateKeysOptions) (resp azfake.Responder[armmachinelearning.EndpointClientRegenerateKeysResponse], errResp azfake.ErrorResponder)
}

// NewEndpointServerTransport creates a new instance of EndpointServerTransport with the provided implementation.
// The returned EndpointServerTransport instance is connected to an instance of armmachinelearning.EndpointClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewEndpointServerTransport(srv *EndpointServer) *EndpointServerTransport {
	return &EndpointServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.EndpointClientCreateOrUpdateResponse]](),
		newGetModelsPager:   newTracker[azfake.PagerResponder[armmachinelearning.EndpointClientGetModelsResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmachinelearning.EndpointClientListResponse]](),
	}
}

// EndpointServerTransport connects instances of armmachinelearning.EndpointClient to instances of EndpointServer.
// Don't use this type directly, use NewEndpointServerTransport instead.
type EndpointServerTransport struct {
	srv                 *EndpointServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.EndpointClientCreateOrUpdateResponse]]
	newGetModelsPager   *tracker[azfake.PagerResponder[armmachinelearning.EndpointClientGetModelsResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmachinelearning.EndpointClientListResponse]]
}

// Do implements the policy.Transporter interface for EndpointServerTransport.
func (e *EndpointServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "EndpointClient.BeginCreateOrUpdate":
		resp, err = e.dispatchBeginCreateOrUpdate(req)
	case "EndpointClient.Get":
		resp, err = e.dispatchGet(req)
	case "EndpointClient.NewGetModelsPager":
		resp, err = e.dispatchNewGetModelsPager(req)
	case "EndpointClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	case "EndpointClient.ListKeys":
		resp, err = e.dispatchListKeys(req)
	case "EndpointClient.RegenerateKeys":
		resp, err = e.dispatchRegenerateKeys(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *EndpointServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if e.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := e.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.EndpointResourcePropertiesBasicResource](req)
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
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := e.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		e.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		e.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		e.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (e *EndpointServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EndpointResourcePropertiesBasicResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointServerTransport) dispatchNewGetModelsPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewGetModelsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewGetModelsPager not implemented")}
	}
	newGetModelsPager := e.newGetModelsPager.get(req)
	if newGetModelsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models`
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
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewGetModelsPager(resourceGroupNameParam, workspaceNameParam, endpointNameParam, nil)
		newGetModelsPager = &resp
		e.newGetModelsPager.add(req, newGetModelsPager)
		server.PagerResponderInjectNextLinks(newGetModelsPager, req, func(page *armmachinelearning.EndpointClientGetModelsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newGetModelsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newGetModelsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newGetModelsPager) {
		e.newGetModelsPager.remove(req)
	}
	return resp, nil
}

func (e *EndpointServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints`
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
		endpointTypeUnescaped, err := url.QueryUnescape(qp.Get("endpointType"))
		if err != nil {
			return nil, err
		}
		endpointTypeParam := getOptional(armmachinelearning.EndpointType(endpointTypeUnescaped))
		includeInferenceEndpointsUnescaped, err := url.QueryUnescape(qp.Get("includeInferenceEndpoints"))
		if err != nil {
			return nil, err
		}
		includeInferenceEndpointsParam, err := parseOptional(includeInferenceEndpointsUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		var options *armmachinelearning.EndpointClientListOptions
		if endpointTypeParam != nil || includeInferenceEndpointsParam != nil || skipParam != nil || expandParam != nil {
			options = &armmachinelearning.EndpointClientListOptions{
				EndpointType:              endpointTypeParam,
				IncludeInferenceEndpoints: includeInferenceEndpointsParam,
				Skip:                      skipParam,
				Expand:                    expandParam,
			}
		}
		resp := e.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.EndpointClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}

func (e *EndpointServerTransport) dispatchListKeys(req *http.Request) (*http.Response, error) {
	if e.srv.ListKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listKeys`
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
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.ListKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).EndpointKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *EndpointServerTransport) dispatchRegenerateKeys(req *http.Request) (*http.Response, error) {
	if e.srv.RegenerateKeys == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateKeys not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.RegenerateServiceAccountKeyContent](req)
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
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.RegenerateKeys(req.Context(), resourceGroupNameParam, workspaceNameParam, endpointNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AccountAPIKeys, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
