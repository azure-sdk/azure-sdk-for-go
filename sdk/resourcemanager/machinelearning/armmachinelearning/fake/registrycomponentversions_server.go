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

// RegistryComponentVersionsServer is a fake server for instances of the armmachinelearning.RegistryComponentVersionsClient type.
type RegistryComponentVersionsServer struct {
	// BeginCreateOrUpdate is the fake for method RegistryComponentVersionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, body armmachinelearning.ComponentVersion, options *armmachinelearning.RegistryComponentVersionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.RegistryComponentVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RegistryComponentVersionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *armmachinelearning.RegistryComponentVersionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.RegistryComponentVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RegistryComponentVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, componentName string, version string, options *armmachinelearning.RegistryComponentVersionsClientGetOptions) (resp azfake.Responder[armmachinelearning.RegistryComponentVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RegistryComponentVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, componentName string, options *armmachinelearning.RegistryComponentVersionsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.RegistryComponentVersionsClientListResponse])
}

// NewRegistryComponentVersionsServerTransport creates a new instance of RegistryComponentVersionsServerTransport with the provided implementation.
// The returned RegistryComponentVersionsServerTransport instance is connected to an instance of armmachinelearning.RegistryComponentVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegistryComponentVersionsServerTransport(srv *RegistryComponentVersionsServer) *RegistryComponentVersionsServerTransport {
	return &RegistryComponentVersionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.RegistryComponentVersionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armmachinelearning.RegistryComponentVersionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmachinelearning.RegistryComponentVersionsClientListResponse]](),
	}
}

// RegistryComponentVersionsServerTransport connects instances of armmachinelearning.RegistryComponentVersionsClient to instances of RegistryComponentVersionsServer.
// Don't use this type directly, use NewRegistryComponentVersionsServerTransport instead.
type RegistryComponentVersionsServerTransport struct {
	srv                 *RegistryComponentVersionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.RegistryComponentVersionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armmachinelearning.RegistryComponentVersionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmachinelearning.RegistryComponentVersionsClientListResponse]]
}

// Do implements the policy.Transporter interface for RegistryComponentVersionsServerTransport.
func (r *RegistryComponentVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegistryComponentVersionsClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RegistryComponentVersionsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RegistryComponentVersionsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RegistryComponentVersionsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegistryComponentVersionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/components/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.ComponentVersion](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
		if err != nil {
			return nil, err
		}
		versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, registryNameParam, componentNameParam, versionParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		r.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		r.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		r.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (r *RegistryComponentVersionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/components/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
		if err != nil {
			return nil, err
		}
		versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, registryNameParam, componentNameParam, versionParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		r.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		r.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		r.beginDelete.remove(req)
	}

	return resp, nil
}

func (r *RegistryComponentVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/components/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
	if err != nil {
		return nil, err
	}
	componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, registryNameParam, componentNameParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ComponentVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RegistryComponentVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/components/(?P<componentName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
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
		registryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("registryName")])
		if err != nil {
			return nil, err
		}
		componentNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("componentName")])
		if err != nil {
			return nil, err
		}
		orderByUnescaped, err := url.QueryUnescape(qp.Get("$orderBy"))
		if err != nil {
			return nil, err
		}
		orderByParam := getOptional(orderByUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
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
		var options *armmachinelearning.RegistryComponentVersionsClientListOptions
		if orderByParam != nil || topParam != nil || skipParam != nil {
			options = &armmachinelearning.RegistryComponentVersionsClientListOptions{
				OrderBy: orderByParam,
				Top:     topParam,
				Skip:    skipParam,
			}
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, registryNameParam, componentNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.RegistryComponentVersionsClientListResponse, createLink func() string) {
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