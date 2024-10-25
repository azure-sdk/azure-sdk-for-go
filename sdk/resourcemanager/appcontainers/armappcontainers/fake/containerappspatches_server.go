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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerAppsPatchesServer is a fake server for instances of the armappcontainers.ContainerAppsPatchesClient type.
type ContainerAppsPatchesServer struct {
	// BeginApply is the fake for method ContainerAppsPatchesClient.BeginApply
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginApply func(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *armappcontainers.ContainerAppsPatchesClientBeginApplyOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientApplyResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ContainerAppsPatchesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *armappcontainers.ContainerAppsPatchesClientBeginDeleteOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ContainerAppsPatchesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, options *armappcontainers.ContainerAppsPatchesClientGetOptions) (resp azfake.Responder[armappcontainers.ContainerAppsPatchesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByContainerAppPager is the fake for method ContainerAppsPatchesClient.NewListByContainerAppPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByContainerAppPager func(resourceGroupName string, containerAppName string, options *armappcontainers.ContainerAppsPatchesClientListByContainerAppOptions) (resp azfake.PagerResponder[armappcontainers.ContainerAppsPatchesClientListByContainerAppResponse])

	// BeginSkipConfigure is the fake for method ContainerAppsPatchesClient.BeginSkipConfigure
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginSkipConfigure func(ctx context.Context, resourceGroupName string, containerAppName string, patchName string, patchSkipConfig armappcontainers.PatchSkipConfig, options *armappcontainers.ContainerAppsPatchesClientBeginSkipConfigureOptions) (resp azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientSkipConfigureResponse], errResp azfake.ErrorResponder)
}

// NewContainerAppsPatchesServerTransport creates a new instance of ContainerAppsPatchesServerTransport with the provided implementation.
// The returned ContainerAppsPatchesServerTransport instance is connected to an instance of armappcontainers.ContainerAppsPatchesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsPatchesServerTransport(srv *ContainerAppsPatchesServer) *ContainerAppsPatchesServerTransport {
	return &ContainerAppsPatchesServerTransport{
		srv:                        srv,
		beginApply:                 newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientApplyResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientDeleteResponse]](),
		newListByContainerAppPager: newTracker[azfake.PagerResponder[armappcontainers.ContainerAppsPatchesClientListByContainerAppResponse]](),
		beginSkipConfigure:         newTracker[azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientSkipConfigureResponse]](),
	}
}

// ContainerAppsPatchesServerTransport connects instances of armappcontainers.ContainerAppsPatchesClient to instances of ContainerAppsPatchesServer.
// Don't use this type directly, use NewContainerAppsPatchesServerTransport instead.
type ContainerAppsPatchesServerTransport struct {
	srv                        *ContainerAppsPatchesServer
	beginApply                 *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientApplyResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientDeleteResponse]]
	newListByContainerAppPager *tracker[azfake.PagerResponder[armappcontainers.ContainerAppsPatchesClientListByContainerAppResponse]]
	beginSkipConfigure         *tracker[azfake.PollerResponder[armappcontainers.ContainerAppsPatchesClientSkipConfigureResponse]]
}

// Do implements the policy.Transporter interface for ContainerAppsPatchesServerTransport.
func (c *ContainerAppsPatchesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ContainerAppsPatchesClient.BeginApply":
		resp, err = c.dispatchBeginApply(req)
	case "ContainerAppsPatchesClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ContainerAppsPatchesClient.Get":
		resp, err = c.dispatchGet(req)
	case "ContainerAppsPatchesClient.NewListByContainerAppPager":
		resp, err = c.dispatchNewListByContainerAppPager(req)
	case "ContainerAppsPatchesClient.BeginSkipConfigure":
		resp, err = c.dispatchBeginSkipConfigure(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ContainerAppsPatchesServerTransport) dispatchBeginApply(req *http.Request) (*http.Response, error) {
	if c.srv.BeginApply == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginApply not implemented")}
	}
	beginApply := c.beginApply.get(req)
	if beginApply == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/patches/(?P<patchName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apply`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		patchNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("patchName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginApply(req.Context(), resourceGroupNameParam, containerAppNameParam, patchNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginApply = &respr
		c.beginApply.add(req, beginApply)
	}

	resp, err := server.PollerResponderNext(beginApply, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		c.beginApply.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginApply) {
		c.beginApply.remove(req)
	}

	return resp, nil
}

func (c *ContainerAppsPatchesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/patches/(?P<patchName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		patchNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("patchName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, containerAppNameParam, patchNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ContainerAppsPatchesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/patches/(?P<patchName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	patchNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("patchName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, containerAppNameParam, patchNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerAppsPatchResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsPatchesServerTransport) dispatchNewListByContainerAppPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByContainerAppPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByContainerAppPager not implemented")}
	}
	newListByContainerAppPager := c.newListByContainerAppPager.get(req)
	if newListByContainerAppPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/patches`
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
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armappcontainers.ContainerAppsPatchesClientListByContainerAppOptions
		if filterParam != nil {
			options = &armappcontainers.ContainerAppsPatchesClientListByContainerAppOptions{
				Filter: filterParam,
			}
		}
		resp := c.srv.NewListByContainerAppPager(resourceGroupNameParam, containerAppNameParam, options)
		newListByContainerAppPager = &resp
		c.newListByContainerAppPager.add(req, newListByContainerAppPager)
		server.PagerResponderInjectNextLinks(newListByContainerAppPager, req, func(page *armappcontainers.ContainerAppsPatchesClientListByContainerAppResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByContainerAppPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByContainerAppPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByContainerAppPager) {
		c.newListByContainerAppPager.remove(req)
	}
	return resp, nil
}

func (c *ContainerAppsPatchesServerTransport) dispatchBeginSkipConfigure(req *http.Request) (*http.Response, error) {
	if c.srv.BeginSkipConfigure == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSkipConfigure not implemented")}
	}
	beginSkipConfigure := c.beginSkipConfigure.get(req)
	if beginSkipConfigure == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/patches/(?P<patchName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skipConfig`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armappcontainers.PatchSkipConfig](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
		if err != nil {
			return nil, err
		}
		patchNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("patchName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginSkipConfigure(req.Context(), resourceGroupNameParam, containerAppNameParam, patchNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSkipConfigure = &respr
		c.beginSkipConfigure.add(req, beginSkipConfigure)
	}

	resp, err := server.PollerResponderNext(beginSkipConfigure, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		c.beginSkipConfigure.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSkipConfigure) {
		c.beginSkipConfigure.remove(req)
	}

	return resp, nil
}