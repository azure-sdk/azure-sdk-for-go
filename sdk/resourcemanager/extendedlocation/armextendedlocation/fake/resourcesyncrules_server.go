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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
	"net/http"
	"net/url"
	"regexp"
)

// ResourceSyncRulesServer is a fake server for instances of the armextendedlocation.ResourceSyncRulesClient type.
type ResourceSyncRulesServer struct {
	// BeginCreateOrUpdate is the fake for method ResourceSyncRulesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters armextendedlocation.ResourceSyncRule, options *armextendedlocation.ResourceSyncRulesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armextendedlocation.ResourceSyncRulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ResourceSyncRulesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *armextendedlocation.ResourceSyncRulesClientDeleteOptions) (resp azfake.Responder[armextendedlocation.ResourceSyncRulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ResourceSyncRulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, options *armextendedlocation.ResourceSyncRulesClientGetOptions) (resp azfake.Responder[armextendedlocation.ResourceSyncRulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByCustomLocationIDPager is the fake for method ResourceSyncRulesClient.NewListByCustomLocationIDPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByCustomLocationIDPager func(resourceGroupName string, resourceName string, options *armextendedlocation.ResourceSyncRulesClientListByCustomLocationIDOptions) (resp azfake.PagerResponder[armextendedlocation.ResourceSyncRulesClientListByCustomLocationIDResponse])

	// BeginUpdate is the fake for method ResourceSyncRulesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK
	BeginUpdate func(ctx context.Context, resourceGroupName string, resourceName string, childResourceName string, parameters armextendedlocation.PatchableResourceSyncRule, options *armextendedlocation.ResourceSyncRulesClientBeginUpdateOptions) (resp azfake.PollerResponder[armextendedlocation.ResourceSyncRulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewResourceSyncRulesServerTransport creates a new instance of ResourceSyncRulesServerTransport with the provided implementation.
// The returned ResourceSyncRulesServerTransport instance is connected to an instance of armextendedlocation.ResourceSyncRulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceSyncRulesServerTransport(srv *ResourceSyncRulesServer) *ResourceSyncRulesServerTransport {
	return &ResourceSyncRulesServerTransport{
		srv:                            srv,
		beginCreateOrUpdate:            newTracker[azfake.PollerResponder[armextendedlocation.ResourceSyncRulesClientCreateOrUpdateResponse]](),
		newListByCustomLocationIDPager: newTracker[azfake.PagerResponder[armextendedlocation.ResourceSyncRulesClientListByCustomLocationIDResponse]](),
		beginUpdate:                    newTracker[azfake.PollerResponder[armextendedlocation.ResourceSyncRulesClientUpdateResponse]](),
	}
}

// ResourceSyncRulesServerTransport connects instances of armextendedlocation.ResourceSyncRulesClient to instances of ResourceSyncRulesServer.
// Don't use this type directly, use NewResourceSyncRulesServerTransport instead.
type ResourceSyncRulesServerTransport struct {
	srv                            *ResourceSyncRulesServer
	beginCreateOrUpdate            *tracker[azfake.PollerResponder[armextendedlocation.ResourceSyncRulesClientCreateOrUpdateResponse]]
	newListByCustomLocationIDPager *tracker[azfake.PagerResponder[armextendedlocation.ResourceSyncRulesClientListByCustomLocationIDResponse]]
	beginUpdate                    *tracker[azfake.PollerResponder[armextendedlocation.ResourceSyncRulesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for ResourceSyncRulesServerTransport.
func (r *ResourceSyncRulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ResourceSyncRulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if resourceSyncRulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = resourceSyncRulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ResourceSyncRulesClient.BeginCreateOrUpdate":
				res.resp, res.err = r.dispatchBeginCreateOrUpdate(req)
			case "ResourceSyncRulesClient.Delete":
				res.resp, res.err = r.dispatchDelete(req)
			case "ResourceSyncRulesClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ResourceSyncRulesClient.NewListByCustomLocationIDPager":
				res.resp, res.err = r.dispatchNewListByCustomLocationIDPager(req)
			case "ResourceSyncRulesClient.BeginUpdate":
				res.resp, res.err = r.dispatchBeginUpdate(req)
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

func (r *ResourceSyncRulesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := r.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ExtendedLocation/customLocations/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceSyncRules/(?P<childResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armextendedlocation.ResourceSyncRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		childResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("childResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, childResourceNameParam, body, nil)
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

func (r *ResourceSyncRulesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if r.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ExtendedLocation/customLocations/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceSyncRules/(?P<childResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	childResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("childResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Delete(req.Context(), resourceGroupNameParam, resourceNameParam, childResourceNameParam, nil)
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

func (r *ResourceSyncRulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ExtendedLocation/customLocations/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceSyncRules/(?P<childResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	childResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("childResourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, resourceNameParam, childResourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceSyncRule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceSyncRulesServerTransport) dispatchNewListByCustomLocationIDPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListByCustomLocationIDPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByCustomLocationIDPager not implemented")}
	}
	newListByCustomLocationIDPager := r.newListByCustomLocationIDPager.get(req)
	if newListByCustomLocationIDPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ExtendedLocation/customLocations/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceSyncRules`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		resp := r.srv.NewListByCustomLocationIDPager(resourceGroupNameParam, resourceNameParam, nil)
		newListByCustomLocationIDPager = &resp
		r.newListByCustomLocationIDPager.add(req, newListByCustomLocationIDPager)
		server.PagerResponderInjectNextLinks(newListByCustomLocationIDPager, req, func(page *armextendedlocation.ResourceSyncRulesClientListByCustomLocationIDResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByCustomLocationIDPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListByCustomLocationIDPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByCustomLocationIDPager) {
		r.newListByCustomLocationIDPager.remove(req)
	}
	return resp, nil
}

func (r *ResourceSyncRulesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := r.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ExtendedLocation/customLocations/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceSyncRules/(?P<childResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armextendedlocation.PatchableResourceSyncRule](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		childResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("childResourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginUpdate(req.Context(), resourceGroupNameParam, resourceNameParam, childResourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		r.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		r.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to ResourceSyncRulesServerTransport
var resourceSyncRulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
