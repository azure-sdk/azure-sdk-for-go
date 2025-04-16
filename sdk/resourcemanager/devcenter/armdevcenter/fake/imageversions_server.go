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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ImageVersionsServer is a fake server for instances of the armdevcenter.ImageVersionsClient type.
type ImageVersionsServer struct {
	// Get is the fake for method ImageVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, devCenterName string, galleryName string, imageName string, versionName string, options *armdevcenter.ImageVersionsClientGetOptions) (resp azfake.Responder[armdevcenter.ImageVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// GetByProject is the fake for method ImageVersionsClient.GetByProject
	// HTTP status codes to indicate success: http.StatusOK
	GetByProject func(ctx context.Context, resourceGroupName string, projectName string, imageName string, versionName string, options *armdevcenter.ImageVersionsClientGetByProjectOptions) (resp azfake.Responder[armdevcenter.ImageVersionsClientGetByProjectResponse], errResp azfake.ErrorResponder)

	// NewListByImagePager is the fake for method ImageVersionsClient.NewListByImagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByImagePager func(resourceGroupName string, devCenterName string, galleryName string, imageName string, options *armdevcenter.ImageVersionsClientListByImageOptions) (resp azfake.PagerResponder[armdevcenter.ImageVersionsClientListByImageResponse])

	// NewListByProjectPager is the fake for method ImageVersionsClient.NewListByProjectPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByProjectPager func(resourceGroupName string, projectName string, imageName string, options *armdevcenter.ImageVersionsClientListByProjectOptions) (resp azfake.PagerResponder[armdevcenter.ImageVersionsClientListByProjectResponse])
}

// NewImageVersionsServerTransport creates a new instance of ImageVersionsServerTransport with the provided implementation.
// The returned ImageVersionsServerTransport instance is connected to an instance of armdevcenter.ImageVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImageVersionsServerTransport(srv *ImageVersionsServer) *ImageVersionsServerTransport {
	return &ImageVersionsServerTransport{
		srv:                   srv,
		newListByImagePager:   newTracker[azfake.PagerResponder[armdevcenter.ImageVersionsClientListByImageResponse]](),
		newListByProjectPager: newTracker[azfake.PagerResponder[armdevcenter.ImageVersionsClientListByProjectResponse]](),
	}
}

// ImageVersionsServerTransport connects instances of armdevcenter.ImageVersionsClient to instances of ImageVersionsServer.
// Don't use this type directly, use NewImageVersionsServerTransport instead.
type ImageVersionsServerTransport struct {
	srv                   *ImageVersionsServer
	newListByImagePager   *tracker[azfake.PagerResponder[armdevcenter.ImageVersionsClientListByImageResponse]]
	newListByProjectPager *tracker[azfake.PagerResponder[armdevcenter.ImageVersionsClientListByProjectResponse]]
}

// Do implements the policy.Transporter interface for ImageVersionsServerTransport.
func (i *ImageVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *ImageVersionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if imageVersionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = imageVersionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ImageVersionsClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "ImageVersionsClient.GetByProject":
				res.resp, res.err = i.dispatchGetByProject(req)
			case "ImageVersionsClient.NewListByImagePager":
				res.resp, res.err = i.dispatchNewListByImagePager(req)
			case "ImageVersionsClient.NewListByProjectPager":
				res.resp, res.err = i.dispatchNewListByProjectPager(req)
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

func (i *ImageVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<imageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
	if err != nil {
		return nil, err
	}
	galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
	if err != nil {
		return nil, err
	}
	imageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, devCenterNameParam, galleryNameParam, imageNameParam, versionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImageVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImageVersionsServerTransport) dispatchGetByProject(req *http.Request) (*http.Response, error) {
	if i.srv.GetByProject == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetByProject not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<imageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<versionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
	if err != nil {
		return nil, err
	}
	imageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageName")])
	if err != nil {
		return nil, err
	}
	versionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("versionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.GetByProject(req.Context(), resourceGroupNameParam, projectNameParam, imageNameParam, versionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImageVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImageVersionsServerTransport) dispatchNewListByImagePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByImagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByImagePager not implemented")}
	}
	newListByImagePager := i.newListByImagePager.get(req)
	if newListByImagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/devcenters/(?P<devCenterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<imageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		devCenterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("devCenterName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		imageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByImagePager(resourceGroupNameParam, devCenterNameParam, galleryNameParam, imageNameParam, nil)
		newListByImagePager = &resp
		i.newListByImagePager.add(req, newListByImagePager)
		server.PagerResponderInjectNextLinks(newListByImagePager, req, func(page *armdevcenter.ImageVersionsClientListByImageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByImagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByImagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByImagePager) {
		i.newListByImagePager.remove(req)
	}
	return resp, nil
}

func (i *ImageVersionsServerTransport) dispatchNewListByProjectPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByProjectPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByProjectPager not implemented")}
	}
	newListByProjectPager := i.newListByProjectPager.get(req)
	if newListByProjectPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevCenter/projects/(?P<projectName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<imageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		projectNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("projectName")])
		if err != nil {
			return nil, err
		}
		imageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByProjectPager(resourceGroupNameParam, projectNameParam, imageNameParam, nil)
		newListByProjectPager = &resp
		i.newListByProjectPager.add(req, newListByProjectPager)
		server.PagerResponderInjectNextLinks(newListByProjectPager, req, func(page *armdevcenter.ImageVersionsClientListByProjectResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByProjectPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByProjectPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByProjectPager) {
		i.newListByProjectPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ImageVersionsServerTransport
var imageVersionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
