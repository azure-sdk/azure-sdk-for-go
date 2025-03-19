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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v6"
	"net/http"
	"net/url"
	"regexp"
)

// GalleryImageVersionsServer is a fake server for instances of the armcompute.GalleryImageVersionsClient type.
type GalleryImageVersionsServer struct {
	// BeginCreateOrUpdate is the fake for method GalleryImageVersionsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion armcompute.GalleryImageVersion, options *armcompute.GalleryImageVersionsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryImageVersionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method GalleryImageVersionsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *armcompute.GalleryImageVersionsClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.GalleryImageVersionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method GalleryImageVersionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, options *armcompute.GalleryImageVersionsClientGetOptions) (resp azfake.Responder[armcompute.GalleryImageVersionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByGalleryImagePager is the fake for method GalleryImageVersionsClient.NewListByGalleryImagePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByGalleryImagePager func(resourceGroupName string, galleryName string, galleryImageName string, options *armcompute.GalleryImageVersionsClientListByGalleryImageOptions) (resp azfake.PagerResponder[armcompute.GalleryImageVersionsClientListByGalleryImageResponse])

	// BeginUpdate is the fake for method GalleryImageVersionsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, galleryName string, galleryImageName string, galleryImageVersionName string, galleryImageVersion armcompute.GalleryImageVersionUpdate, options *armcompute.GalleryImageVersionsClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.GalleryImageVersionsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewGalleryImageVersionsServerTransport creates a new instance of GalleryImageVersionsServerTransport with the provided implementation.
// The returned GalleryImageVersionsServerTransport instance is connected to an instance of armcompute.GalleryImageVersionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewGalleryImageVersionsServerTransport(srv *GalleryImageVersionsServer) *GalleryImageVersionsServerTransport {
	return &GalleryImageVersionsServerTransport{
		srv:                        srv,
		beginCreateOrUpdate:        newTracker[azfake.PollerResponder[armcompute.GalleryImageVersionsClientCreateOrUpdateResponse]](),
		beginDelete:                newTracker[azfake.PollerResponder[armcompute.GalleryImageVersionsClientDeleteResponse]](),
		newListByGalleryImagePager: newTracker[azfake.PagerResponder[armcompute.GalleryImageVersionsClientListByGalleryImageResponse]](),
		beginUpdate:                newTracker[azfake.PollerResponder[armcompute.GalleryImageVersionsClientUpdateResponse]](),
	}
}

// GalleryImageVersionsServerTransport connects instances of armcompute.GalleryImageVersionsClient to instances of GalleryImageVersionsServer.
// Don't use this type directly, use NewGalleryImageVersionsServerTransport instead.
type GalleryImageVersionsServerTransport struct {
	srv                        *GalleryImageVersionsServer
	beginCreateOrUpdate        *tracker[azfake.PollerResponder[armcompute.GalleryImageVersionsClientCreateOrUpdateResponse]]
	beginDelete                *tracker[azfake.PollerResponder[armcompute.GalleryImageVersionsClientDeleteResponse]]
	newListByGalleryImagePager *tracker[azfake.PagerResponder[armcompute.GalleryImageVersionsClientListByGalleryImageResponse]]
	beginUpdate                *tracker[azfake.PollerResponder[armcompute.GalleryImageVersionsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for GalleryImageVersionsServerTransport.
func (g *GalleryImageVersionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return g.dispatchToMethodFake(req, method)
}

func (g *GalleryImageVersionsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if galleryImageVersionsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = galleryImageVersionsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "GalleryImageVersionsClient.BeginCreateOrUpdate":
				res.resp, res.err = g.dispatchBeginCreateOrUpdate(req)
			case "GalleryImageVersionsClient.BeginDelete":
				res.resp, res.err = g.dispatchBeginDelete(req)
			case "GalleryImageVersionsClient.Get":
				res.resp, res.err = g.dispatchGet(req)
			case "GalleryImageVersionsClient.NewListByGalleryImagePager":
				res.resp, res.err = g.dispatchNewListByGalleryImagePager(req)
			case "GalleryImageVersionsClient.BeginUpdate":
				res.resp, res.err = g.dispatchBeginUpdate(req)
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

func (g *GalleryImageVersionsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := g.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryImageVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryImageVersion](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryImageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
		if err != nil {
			return nil, err
		}
		galleryImageVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, galleryNameParam, galleryImageNameParam, galleryImageVersionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		g.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		g.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		g.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if g.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := g.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryImageVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryImageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
		if err != nil {
			return nil, err
		}
		galleryImageVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginDelete(req.Context(), resourceGroupNameParam, galleryNameParam, galleryImageNameParam, galleryImageVersionNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		g.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		g.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		g.beginDelete.remove(req)
	}

	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if g.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryImageVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
	if err != nil {
		return nil, err
	}
	galleryImageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
	if err != nil {
		return nil, err
	}
	galleryImageVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageVersionName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.ReplicationStatusTypes(expandUnescaped))
	var options *armcompute.GalleryImageVersionsClientGetOptions
	if expandParam != nil {
		options = &armcompute.GalleryImageVersionsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := g.srv.Get(req.Context(), resourceGroupNameParam, galleryNameParam, galleryImageNameParam, galleryImageVersionNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).GalleryImageVersion, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchNewListByGalleryImagePager(req *http.Request) (*http.Response, error) {
	if g.srv.NewListByGalleryImagePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByGalleryImagePager not implemented")}
	}
	newListByGalleryImagePager := g.newListByGalleryImagePager.get(req)
	if newListByGalleryImagePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryImageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
		if err != nil {
			return nil, err
		}
		resp := g.srv.NewListByGalleryImagePager(resourceGroupNameParam, galleryNameParam, galleryImageNameParam, nil)
		newListByGalleryImagePager = &resp
		g.newListByGalleryImagePager.add(req, newListByGalleryImagePager)
		server.PagerResponderInjectNextLinks(newListByGalleryImagePager, req, func(page *armcompute.GalleryImageVersionsClientListByGalleryImageResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByGalleryImagePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.newListByGalleryImagePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByGalleryImagePager) {
		g.newListByGalleryImagePager.remove(req)
	}
	return resp, nil
}

func (g *GalleryImageVersionsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if g.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := g.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Compute/galleries/(?P<galleryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/images/(?P<galleryImageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<galleryImageVersionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GalleryImageVersionUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		galleryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryName")])
		if err != nil {
			return nil, err
		}
		galleryImageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageName")])
		if err != nil {
			return nil, err
		}
		galleryImageVersionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("galleryImageVersionName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := g.srv.BeginUpdate(req.Context(), resourceGroupNameParam, galleryNameParam, galleryImageNameParam, galleryImageVersionNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		g.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		g.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		g.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to GalleryImageVersionsServerTransport
var galleryImageVersionsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
