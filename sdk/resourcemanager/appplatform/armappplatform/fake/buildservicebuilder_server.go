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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appplatform/armappplatform/v2"
	"net/http"
	"net/url"
	"regexp"
)

// BuildServiceBuilderServer is a fake server for instances of the armappplatform.BuildServiceBuilderClient type.
type BuildServiceBuilderServer struct {
	// BeginCreateOrUpdate is the fake for method BuildServiceBuilderClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, builderResource armappplatform.BuilderResource, options *armappplatform.BuildServiceBuilderClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armappplatform.BuildServiceBuilderClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method BuildServiceBuilderClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *armappplatform.BuildServiceBuilderClientBeginDeleteOptions) (resp azfake.PollerResponder[armappplatform.BuildServiceBuilderClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BuildServiceBuilderClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *armappplatform.BuildServiceBuilderClientGetOptions) (resp azfake.Responder[armappplatform.BuildServiceBuilderClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BuildServiceBuilderClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, buildServiceName string, options *armappplatform.BuildServiceBuilderClientListOptions) (resp azfake.PagerResponder[armappplatform.BuildServiceBuilderClientListResponse])

	// ListDeployments is the fake for method BuildServiceBuilderClient.ListDeployments
	// HTTP status codes to indicate success: http.StatusOK
	ListDeployments func(ctx context.Context, resourceGroupName string, serviceName string, buildServiceName string, builderName string, options *armappplatform.BuildServiceBuilderClientListDeploymentsOptions) (resp azfake.Responder[armappplatform.BuildServiceBuilderClientListDeploymentsResponse], errResp azfake.ErrorResponder)
}

// NewBuildServiceBuilderServerTransport creates a new instance of BuildServiceBuilderServerTransport with the provided implementation.
// The returned BuildServiceBuilderServerTransport instance is connected to an instance of armappplatform.BuildServiceBuilderClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBuildServiceBuilderServerTransport(srv *BuildServiceBuilderServer) *BuildServiceBuilderServerTransport {
	return &BuildServiceBuilderServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armappplatform.BuildServiceBuilderClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armappplatform.BuildServiceBuilderClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armappplatform.BuildServiceBuilderClientListResponse]](),
	}
}

// BuildServiceBuilderServerTransport connects instances of armappplatform.BuildServiceBuilderClient to instances of BuildServiceBuilderServer.
// Don't use this type directly, use NewBuildServiceBuilderServerTransport instead.
type BuildServiceBuilderServerTransport struct {
	srv                 *BuildServiceBuilderServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armappplatform.BuildServiceBuilderClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armappplatform.BuildServiceBuilderClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armappplatform.BuildServiceBuilderClientListResponse]]
}

// Do implements the policy.Transporter interface for BuildServiceBuilderServerTransport.
func (b *BuildServiceBuilderServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BuildServiceBuilderClient.BeginCreateOrUpdate":
		resp, err = b.dispatchBeginCreateOrUpdate(req)
	case "BuildServiceBuilderClient.BeginDelete":
		resp, err = b.dispatchBeginDelete(req)
	case "BuildServiceBuilderClient.Get":
		resp, err = b.dispatchGet(req)
	case "BuildServiceBuilderClient.NewListPager":
		resp, err = b.dispatchNewListPager(req)
	case "BuildServiceBuilderClient.ListDeployments":
		resp, err = b.dispatchListDeployments(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BuildServiceBuilderServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := b.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappplatform.BuilderResource](req)
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
		buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		b.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		b.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		b.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (b *BuildServiceBuilderServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if b.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := b.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
		if err != nil {
			return nil, err
		}
		builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginDelete(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		b.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		b.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		b.beginDelete.remove(req)
	}

	return resp, nil
}

func (b *BuildServiceBuilderServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
	if err != nil {
		return nil, err
	}
	builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BuilderResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BuildServiceBuilderServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
		if err != nil {
			return nil, err
		}
		resp := b.srv.NewListPager(resourceGroupNameParam, serviceNameParam, buildServiceNameParam, nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappplatform.BuildServiceBuilderClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}

func (b *BuildServiceBuilderServerTransport) dispatchListDeployments(req *http.Request) (*http.Response, error) {
	if b.srv.ListDeployments == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListDeployments not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AppPlatform/Spring/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/buildServices/(?P<buildServiceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/builders/(?P<builderName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listUsingDeployments`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	buildServiceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("buildServiceName")])
	if err != nil {
		return nil, err
	}
	builderNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("builderName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.ListDeployments(req.Context(), resourceGroupNameParam, serviceNameParam, buildServiceNameParam, builderNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeploymentList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
