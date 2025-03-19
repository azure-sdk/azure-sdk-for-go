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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// ResourceGroupsServer is a fake server for instances of the armresources.ResourceGroupsClient type.
type ResourceGroupsServer struct {
	// CheckExistence is the fake for method ResourceGroupsClient.CheckExistence
	// HTTP status codes to indicate success: http.StatusNoContent, http.StatusNotFound
	CheckExistence func(ctx context.Context, resourceGroupName string, options *armresources.ResourceGroupsClientCheckExistenceOptions) (resp azfake.Responder[armresources.ResourceGroupsClientCheckExistenceResponse], errResp azfake.ErrorResponder)

	// CreateOrUpdate is the fake for method ResourceGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, parameters armresources.ResourceGroup, options *armresources.ResourceGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armresources.ResourceGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ResourceGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, options *armresources.ResourceGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armresources.ResourceGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExportTemplate is the fake for method ResourceGroupsClient.BeginExportTemplate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportTemplate func(ctx context.Context, resourceGroupName string, parameters armresources.ExportTemplateRequest, options *armresources.ResourceGroupsClientBeginExportTemplateOptions) (resp azfake.PollerResponder[armresources.ResourceGroupsClientExportTemplateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ResourceGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, options *armresources.ResourceGroupsClientGetOptions) (resp azfake.Responder[armresources.ResourceGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ResourceGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armresources.ResourceGroupsClientListOptions) (resp azfake.PagerResponder[armresources.ResourceGroupsClientListResponse])

	// Update is the fake for method ResourceGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, parameters armresources.ResourceGroupPatchable, options *armresources.ResourceGroupsClientUpdateOptions) (resp azfake.Responder[armresources.ResourceGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewResourceGroupsServerTransport creates a new instance of ResourceGroupsServerTransport with the provided implementation.
// The returned ResourceGroupsServerTransport instance is connected to an instance of armresources.ResourceGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewResourceGroupsServerTransport(srv *ResourceGroupsServer) *ResourceGroupsServerTransport {
	return &ResourceGroupsServerTransport{
		srv:                 srv,
		beginDelete:         newTracker[azfake.PollerResponder[armresources.ResourceGroupsClientDeleteResponse]](),
		beginExportTemplate: newTracker[azfake.PollerResponder[armresources.ResourceGroupsClientExportTemplateResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armresources.ResourceGroupsClientListResponse]](),
	}
}

// ResourceGroupsServerTransport connects instances of armresources.ResourceGroupsClient to instances of ResourceGroupsServer.
// Don't use this type directly, use NewResourceGroupsServerTransport instead.
type ResourceGroupsServerTransport struct {
	srv                 *ResourceGroupsServer
	beginDelete         *tracker[azfake.PollerResponder[armresources.ResourceGroupsClientDeleteResponse]]
	beginExportTemplate *tracker[azfake.PollerResponder[armresources.ResourceGroupsClientExportTemplateResponse]]
	newListPager        *tracker[azfake.PagerResponder[armresources.ResourceGroupsClientListResponse]]
}

// Do implements the policy.Transporter interface for ResourceGroupsServerTransport.
func (r *ResourceGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return r.dispatchToMethodFake(req, method)
}

func (r *ResourceGroupsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if resourceGroupsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = resourceGroupsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ResourceGroupsClient.CheckExistence":
				res.resp, res.err = r.dispatchCheckExistence(req)
			case "ResourceGroupsClient.CreateOrUpdate":
				res.resp, res.err = r.dispatchCreateOrUpdate(req)
			case "ResourceGroupsClient.BeginDelete":
				res.resp, res.err = r.dispatchBeginDelete(req)
			case "ResourceGroupsClient.BeginExportTemplate":
				res.resp, res.err = r.dispatchBeginExportTemplate(req)
			case "ResourceGroupsClient.Get":
				res.resp, res.err = r.dispatchGet(req)
			case "ResourceGroupsClient.NewListPager":
				res.resp, res.err = r.dispatchNewListPager(req)
			case "ResourceGroupsClient.Update":
				res.resp, res.err = r.dispatchUpdate(req)
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

func (r *ResourceGroupsServerTransport) dispatchCheckExistence(req *http.Request) (*http.Response, error) {
	if r.srv.CheckExistence == nil {
		return nil, &nonRetriableError{errors.New("fake for method CheckExistence not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CheckExistence(req.Context(), resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent, http.StatusNotFound}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent, http.StatusNotFound", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.ResourceGroup](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if r.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := r.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		forceDeletionTypesUnescaped, err := url.QueryUnescape(qp.Get("forceDeletionTypes"))
		if err != nil {
			return nil, err
		}
		forceDeletionTypesParam := getOptional(forceDeletionTypesUnescaped)
		var options *armresources.ResourceGroupsClientBeginDeleteOptions
		if forceDeletionTypesParam != nil {
			options = &armresources.ResourceGroupsClientBeginDeleteOptions{
				ForceDeletionTypes: forceDeletionTypesParam,
			}
		}
		respr, errRespr := r.srv.BeginDelete(req.Context(), resourceGroupNameParam, options)
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

func (r *ResourceGroupsServerTransport) dispatchBeginExportTemplate(req *http.Request) (*http.Response, error) {
	if r.srv.BeginExportTemplate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportTemplate not implemented")}
	}
	beginExportTemplate := r.beginExportTemplate.get(req)
	if beginExportTemplate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportTemplate`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armresources.ExportTemplateRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := r.srv.BeginExportTemplate(req.Context(), resourceGroupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportTemplate = &respr
		r.beginExportTemplate.add(req, beginExportTemplate)
	}

	resp, err := server.PollerResponderNext(beginExportTemplate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		r.beginExportTemplate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportTemplate) {
		r.beginExportTemplate.remove(req)
	}

	return resp, nil
}

func (r *ResourceGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if r.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Get(req.Context(), resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *ResourceGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		var options *armresources.ResourceGroupsClientListOptions
		if filterParam != nil || topParam != nil {
			options = &armresources.ResourceGroupsClientListOptions{
				Filter: filterParam,
				Top:    topParam,
			}
		}
		resp := r.srv.NewListPager(options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armresources.ResourceGroupsClientListResponse, createLink func() string) {
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

func (r *ResourceGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if r.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourcegroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armresources.ResourceGroupPatchable](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := r.srv.Update(req.Context(), resourceGroupNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ResourceGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ResourceGroupsServerTransport
var resourceGroupsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
