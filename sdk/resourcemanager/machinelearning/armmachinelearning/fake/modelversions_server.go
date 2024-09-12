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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v5"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ModelVersionsServer is a fake server for instances of the armmachinelearning.ModelVersionsClient type.
type ModelVersionsServer struct {
	// CreateOrUpdate is the fake for method ModelVersionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body armmachinelearning.ModelVersion, options *armmachinelearning.ModelVersionsClientCreateOrUpdateOptions) (resp azfake.Responder[armmachinelearning.ModelVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ModelVersionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *armmachinelearning.ModelVersionsClientDeleteOptions) (resp azfake.Responder[armmachinelearning.ModelVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ModelVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *armmachinelearning.ModelVersionsClientGetOptions) (resp azfake.Responder[armmachinelearning.ModelVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ModelVersionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, name string, options *armmachinelearning.ModelVersionsClientListOptions) (resp azfake.PagerResponder[armmachinelearning.ModelVersionsClientListResponse])

	// BeginPublish is the fake for method ModelVersionsClient.BeginPublish
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPublish func(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body armmachinelearning.DestinationAsset, options *armmachinelearning.ModelVersionsClientBeginPublishOptions) (resp azfake.PollerResponder[armmachinelearning.ModelVersionsClientPublishResponse], errResp azfake.ErrorResponder)
}

// NewModelVersionsServerTransport creates a new instance of ModelVersionsServerTransport with the provided implementation.
// The returned ModelVersionsServerTransport instance is connected to an instance of armmachinelearning.ModelVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewModelVersionsServerTransport(srv *ModelVersionsServer) *ModelVersionsServerTransport {
	return &ModelVersionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armmachinelearning.ModelVersionsClientListResponse]](),
		beginPublish: newTracker[azfake.PollerResponder[armmachinelearning.ModelVersionsClientPublishResponse]](),
	}
}

// ModelVersionsServerTransport connects instances of armmachinelearning.ModelVersionsClient to instances of ModelVersionsServer.
// Don't use this type directly, use NewModelVersionsServerTransport instead.
type ModelVersionsServerTransport struct {
	srv          *ModelVersionsServer
	newListPager *tracker[azfake.PagerResponder[armmachinelearning.ModelVersionsClientListResponse]]
	beginPublish *tracker[azfake.PollerResponder[armmachinelearning.ModelVersionsClientPublishResponse]]
}

// Do implements the policy.Transporter interface for ModelVersionsServerTransport.
func (m *ModelVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ModelVersionsClient.CreateOrUpdate":
		resp, err = m.dispatchCreateOrUpdate(req)
	case "ModelVersionsClient.Delete":
		resp, err = m.dispatchDelete(req)
	case "ModelVersionsClient.Get":
		resp, err = m.dispatchGet(req)
	case "ModelVersionsClient.NewListPager":
		resp, err = m.dispatchNewListPager(req)
	case "ModelVersionsClient.BeginPublish":
		resp, err = m.dispatchBeginPublish(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *ModelVersionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if m.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ModelVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ModelVersionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if m.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *ModelVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if m.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := m.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, nil)
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

func (m *ModelVersionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if m.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := m.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
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
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
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
		offsetUnescaped, err := url.QueryUnescape(qp.Get("offset"))
		if err != nil {
			return nil, err
		}
		offsetParam, err := parseOptional(offsetUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
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
		feedUnescaped, err := url.QueryUnescape(qp.Get("feed"))
		if err != nil {
			return nil, err
		}
		feedParam := getOptional(feedUnescaped)
		listViewTypeUnescaped, err := url.QueryUnescape(qp.Get("listViewType"))
		if err != nil {
			return nil, err
		}
		listViewTypeParam := getOptional(armmachinelearning.ListViewType(listViewTypeUnescaped))
		var options *armmachinelearning.ModelVersionsClientListOptions
		if skipParam != nil || orderByParam != nil || topParam != nil || versionParam != nil || descriptionParam != nil || offsetParam != nil || tagsParam != nil || propertiesParam != nil || feedParam != nil || listViewTypeParam != nil {
			options = &armmachinelearning.ModelVersionsClientListOptions{
				Skip:         skipParam,
				OrderBy:      orderByParam,
				Top:          topParam,
				Version:      versionParam,
				Description:  descriptionParam,
				Offset:       offsetParam,
				Tags:         tagsParam,
				Properties:   propertiesParam,
				Feed:         feedParam,
				ListViewType: listViewTypeParam,
			}
		}
		resp := m.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, nameParam, options)
		newListPager = &resp
		m.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armmachinelearning.ModelVersionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		m.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		m.newListPager.remove(req)
	}
	return resp, nil
}

func (m *ModelVersionsServerTransport) dispatchBeginPublish(req *http.Request) (*http.Response, error) {
	if m.srv.BeginPublish == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPublish not implemented")}
	}
	beginPublish := m.beginPublish.get(req)
	if beginPublish == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.MachineLearningServices/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/models/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publish`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmachinelearning.DestinationAsset](req)
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
		versionParam, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := m.srv.BeginPublish(req.Context(), resourceGroupNameParam, workspaceNameParam, nameParam, versionParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPublish = &respr
		m.beginPublish.add(req, beginPublish)
	}

	resp, err := server.PollerResponderNext(beginPublish, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		m.beginPublish.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPublish) {
		m.beginPublish.remove(req)
	}

	return resp, nil
}
