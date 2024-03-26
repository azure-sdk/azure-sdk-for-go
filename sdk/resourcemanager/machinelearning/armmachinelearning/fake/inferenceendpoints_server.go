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

// InferenceEndpointsServer is a fake server for instances of the armmachinelearning.InferenceEndpointsClient type.
type InferenceEndpointsServer struct {
	// BeginCreateOrUpdate is the fake for method InferenceEndpointsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body armmachinelearning.InferenceEndpoint, options *armmachinelearning.InferenceEndpointsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method InferenceEndpointsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, options *armmachinelearning.InferenceEndpointsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InferenceEndpointsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, options *armmachinelearning.InferenceEndpointsClientGetOptions) (resp azfake.Responder[armmachinelearning.InferenceEndpointsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method InferenceEndpointsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, poolName string, options *armmachinelearning.InferenceEndpointsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.InferenceEndpointsClientListResponse])

	// BeginUpdate is the fake for method InferenceEndpointsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, poolName string, endpointName string, body any, options *armmachinelearning.InferenceEndpointsClientBeginUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewInferenceEndpointsServerTransport creates a new instance of InferenceEndpointsServerTransport with the provided implementation.
// The returned InferenceEndpointsServerTransport instance is connected to an instance of armmachinelearning.InferenceEndpointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInferenceEndpointsServerTransport(srv *InferenceEndpointsServer) *InferenceEndpointsServerTransport {
	return &InferenceEndpointsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmachinelearning.InferenceEndpointsClientListResponse]](),
		beginUpdate:         newTracker[azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientUpdateResponse]](),
	}
}

// InferenceEndpointsServerTransport connects instances of armmachinelearning.InferenceEndpointsClient to instances of InferenceEndpointsServer.
// Don't use this type directly, use NewInferenceEndpointsServerTransport instead.
type InferenceEndpointsServerTransport struct {
	srv                 *InferenceEndpointsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmachinelearning.InferenceEndpointsClientListResponse]]
	beginUpdate         *tracker[azfake.PollerResponder[armmachinelearning.InferenceEndpointsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for InferenceEndpointsServerTransport.
func (i *InferenceEndpointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InferenceEndpointsClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "InferenceEndpointsClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "InferenceEndpointsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InferenceEndpointsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	case "InferenceEndpointsClient.BeginUpdate":
		resp, err = i.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InferenceEndpointsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inferencePools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.InferenceEndpoint](req)
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, poolNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *InferenceEndpointsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inferencePools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupNameParam, workspaceNameParam, poolNameParam, endpointNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *InferenceEndpointsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inferencePools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
	if err != nil {
		return nil, err
	}
	endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, poolNameParam, endpointNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InferenceEndpoint, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InferenceEndpointsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inferencePools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		countUnescaped, err := url.QueryUnescape(qp.Get("count"))
		if err != nil {
			return nil, err
		}
		countParam, err := parseOptional(countUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
		tagsUnescaped, err := url.QueryUnescape(qp.Get("tags"))
		if err != nil {
			return nil, err
		}
		tagsParam := getOptional(tagsUnescaped)
		propertiesUnescaped, err := url.QueryUnescape(qp.Get("properties"))
		if err != nil {
			return nil, err
		}
		propertiesParam := getOptional(propertiesUnescaped)
		orderByUnescaped, err := url.QueryUnescape(qp.Get("orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(armmachinelearning.OrderString(orderByUnescaped))
		var options *armmachinelearning.InferenceEndpointsClientListOptions
		if countParam != nil || skipParam != nil || tagsParam != nil || propertiesParam != nil || orderByParam != nil {
			options = &armmachinelearning.InferenceEndpointsClientListOptions{
				Count:      countParam,
				Skip:       skipParam,
				Tags:       tagsParam,
				Properties: propertiesParam,
				OrderBy:    orderByParam,
			}
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, poolNameParam, options)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.InferenceEndpointsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

func (i *InferenceEndpointsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := i.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/inferencePools/(?P<poolName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endpoints/(?P<endpointName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[any](req)
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
		poolNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("poolName")])
		if err != nil {
			return nil, err
		}
		endpointNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("endpointName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, poolNameParam, endpointNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		i.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		i.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		i.beginUpdate.remove(req)
	}

	return resp, nil
}