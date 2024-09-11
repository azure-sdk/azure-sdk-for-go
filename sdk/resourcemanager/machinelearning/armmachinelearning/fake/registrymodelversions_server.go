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

// RegistryModelVersionsServer is a fake server for instances of the armmachinelearning.RegistryModelVersionsClient type.
type RegistryModelVersionsServer struct {
	// CreateOrGetStartPendingUpload is the fake for method RegistryModelVersionsClient.CreateOrGetStartPendingUpload
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrGetStartPendingUpload func(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body armmachinelearning.PendingUploadRequestDto, options *armmachinelearning.RegistryModelVersionsClientCreateOrGetStartPendingUploadOptions) (resp azfake.Responder[armmachinelearning.RegistryModelVersionsClientCreateOrGetStartPendingUploadResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method RegistryModelVersionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, body armmachinelearning.ModelVersion, options *armmachinelearning.RegistryModelVersionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armmachinelearning.RegistryModelVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method RegistryModelVersionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *armmachinelearning.RegistryModelVersionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armmachinelearning.RegistryModelVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method RegistryModelVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, registryName string, modelName string, version string, options *armmachinelearning.RegistryModelVersionsClientGetOptions) (resp azfake.Responder[armmachinelearning.RegistryModelVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method RegistryModelVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, registryName string, modelName string, options *armmachinelearning.RegistryModelVersionsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.RegistryModelVersionsClientListResponse])
}

// NewRegistryModelVersionsServerTransport creates a new instance of RegistryModelVersionsServerTransport with the provided implementation.
// The returned RegistryModelVersionsServerTransport instance is connected to an instance of armmachinelearning.RegistryModelVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRegistryModelVersionsServerTransport(srv *RegistryModelVersionsServer) *RegistryModelVersionsServerTransport {
	return &RegistryModelVersionsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armmachinelearning.RegistryModelVersionsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armmachinelearning.RegistryModelVersionsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armmachinelearning.RegistryModelVersionsClientListResponse]](),
	}
}

// RegistryModelVersionsServerTransport connects instances of armmachinelearning.RegistryModelVersionsClient to instances of RegistryModelVersionsServer.
// Don't use this type directly, use NewRegistryModelVersionsServerTransport instead.
type RegistryModelVersionsServerTransport struct {
	srv                 *RegistryModelVersionsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armmachinelearning.RegistryModelVersionsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armmachinelearning.RegistryModelVersionsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armmachinelearning.RegistryModelVersionsClientListResponse]]
}

// Do implements the policy.Transporter interface for RegistryModelVersionsServerTransport.
func (r *RegistryModelVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RegistryModelVersionsClient.CreateOrGetStartPendingUpload":
		resp, err = r.dispatchCreateOrGetStartPendingUpload(req)
	case "RegistryModelVersionsClient.BeginCreateOrUpdate":
		resp, err = r.dispatchBeginCreateOrUpdate(req)
	case "RegistryModelVersionsClient.BeginDelete":
		resp, err = r.dispatchBeginDelete(req)
	case "RegistryModelVersionsClient.Get":
		resp, err = r.dispatchGet(req)
	case "RegistryModelVersionsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RegistryModelVersionsServerTransport) dispatchCreateOrGetStartPendingUpload(req *http.Request) (*http.Response, error) {
	if r.srv.CreateOrGetStartPendingUpload == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrGetStartPendingUpload not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/startPendingUpload`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmachinelearning.PendingUploadRequestDto](req)
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
	modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CreateOrGetStartPendingUpload(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, versionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PendingUploadResponseDto, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RegistryModelVersionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.ModelVersion](req)
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
		modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
		if err != nil {
			return nil, err
		}
		versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, versionParam, body, nil)
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

func (r *RegistryModelVersionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
		if err != nil {
			return nil, err
		}
		versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, versionParam, nil)
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

func (r *RegistryModelVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, registryNameParam, modelNameParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ModelVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *RegistryModelVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/registries/(?P<registryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<modelName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
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
		modelNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("modelName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam := getOptional(skipUnescaped)
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
		versionUnescaped, err := url.QueryUnescape(qp.Get("version"))
		if err != nil {
			return nil, err
		}
		versionParam := getOptional(versionUnescaped)
		descriptionUnescaped, err := url.QueryUnescape(qp.Get("description"))
		if err != nil {
			return nil, err
		}
		descriptionParam := getOptional(descriptionUnescaped)
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
		listViewTypeUnescaped, err := url.QueryUnescape(qp.Get("listViewType"))
		if err != nil {
			return nil, err
		}
		listViewTypeParam := getOptional(armmachinelearning.ListViewType(listViewTypeUnescaped))
		var options *armmachinelearning.RegistryModelVersionsClientListOptions
		if skipParam != nil || orderByParam != nil || topParam != nil || versionParam != nil || descriptionParam != nil || tagsParam != nil || propertiesParam != nil || listViewTypeParam != nil {
			options = &armmachinelearning.RegistryModelVersionsClientListOptions{
				Skip:         skipParam,
				OrderBy:      orderByParam,
				Top:          topParam,
				Version:      versionParam,
				Description:  descriptionParam,
				Tags:         tagsParam,
				Properties:   propertiesParam,
				ListViewType: listViewTypeParam,
			}
		}
		resp := r.srv.NewListPager(resourceGroupNameParam, registryNameParam, modelNameParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.RegistryModelVersionsClientListResponse, createLink func() string) {
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