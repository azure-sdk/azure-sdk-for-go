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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v10"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// PipelinesServer is a fake server for instances of the armdatafactory.PipelinesClient type.
type PipelinesServer struct {
	// CreateOrUpdate is the fake for method PipelinesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline armdatafactory.PipelineResource, options *armdatafactory.PipelinesClientCreateOrUpdateOptions) (resp azfake.Responder[armdatafactory.PipelinesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// CreateRun is the fake for method PipelinesClient.CreateRun
	// HTTP status codes to indicate success: http.StatusOK
	CreateRun func(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *armdatafactory.PipelinesClientCreateRunOptions) (resp azfake.Responder[armdatafactory.PipelinesClientCreateRunResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method PipelinesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *armdatafactory.PipelinesClientDeleteOptions) (resp azfake.Responder[armdatafactory.PipelinesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method PipelinesClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNotModified
	Get func(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *armdatafactory.PipelinesClientGetOptions) (resp azfake.Responder[armdatafactory.PipelinesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFactoryPager is the fake for method PipelinesClient.NewListByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFactoryPager func(resourceGroupName string, factoryName string, options *armdatafactory.PipelinesClientListByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.PipelinesClientListByFactoryResponse])
}

// NewPipelinesServerTransport creates a new instance of PipelinesServerTransport with the provided implementation.
// The returned PipelinesServerTransport instance is connected to an instance of armdatafactory.PipelinesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewPipelinesServerTransport(srv *PipelinesServer) *PipelinesServerTransport {
	return &PipelinesServerTransport{
		srv:                   srv,
		newListByFactoryPager: newTracker[azfake.PagerResponder[armdatafactory.PipelinesClientListByFactoryResponse]](),
	}
}

// PipelinesServerTransport connects instances of armdatafactory.PipelinesClient to instances of PipelinesServer.
// Don't use this type directly, use NewPipelinesServerTransport instead.
type PipelinesServerTransport struct {
	srv                   *PipelinesServer
	newListByFactoryPager *tracker[azfake.PagerResponder[armdatafactory.PipelinesClientListByFactoryResponse]]
}

// Do implements the policy.Transporter interface for PipelinesServerTransport.
func (p *PipelinesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "PipelinesClient.CreateOrUpdate":
		resp, err = p.dispatchCreateOrUpdate(req)
	case "PipelinesClient.CreateRun":
		resp, err = p.dispatchCreateRun(req)
	case "PipelinesClient.Delete":
		resp, err = p.dispatchDelete(req)
	case "PipelinesClient.Get":
		resp, err = p.dispatchGet(req)
	case "PipelinesClient.NewListByFactoryPager":
		resp, err = p.dispatchNewListByFactoryPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *PipelinesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if p.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelines/(?P<pipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.PipelineResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	pipelineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineName")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armdatafactory.PipelinesClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armdatafactory.PipelinesClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := p.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, factoryNameParam, pipelineNameParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PipelineResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PipelinesServerTransport) dispatchCreateRun(req *http.Request) (*http.Response, error) {
	if p.srv.CreateRun == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateRun not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelines/(?P<pipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createRun`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[map[string]any](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	pipelineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineName")])
	if err != nil {
		return nil, err
	}
	referencePipelineRunIDUnescaped, err := url.QueryUnescape(qp.Get("referencePipelineRunId"))
	if err != nil {
		return nil, err
	}
	referencePipelineRunIDParam := getOptional(referencePipelineRunIDUnescaped)
	isRecoveryUnescaped, err := url.QueryUnescape(qp.Get("isRecovery"))
	if err != nil {
		return nil, err
	}
	isRecoveryParam, err := parseOptional(isRecoveryUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	startActivityNameUnescaped, err := url.QueryUnescape(qp.Get("startActivityName"))
	if err != nil {
		return nil, err
	}
	startActivityNameParam := getOptional(startActivityNameUnescaped)
	startFromFailureUnescaped, err := url.QueryUnescape(qp.Get("startFromFailure"))
	if err != nil {
		return nil, err
	}
	startFromFailureParam, err := parseOptional(startFromFailureUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	var options *armdatafactory.PipelinesClientCreateRunOptions
	if referencePipelineRunIDParam != nil || isRecoveryParam != nil || startActivityNameParam != nil || startFromFailureParam != nil || !reflect.ValueOf(body).IsZero() {
		options = &armdatafactory.PipelinesClientCreateRunOptions{
			ReferencePipelineRunID: referencePipelineRunIDParam,
			IsRecovery:             isRecoveryParam,
			StartActivityName:      startActivityNameParam,
			StartFromFailure:       startFromFailureParam,
			Parameters:             body,
		}
	}
	respr, errRespr := p.srv.CreateRun(req.Context(), resourceGroupNameParam, factoryNameParam, pipelineNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).CreateRunResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PipelinesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if p.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelines/(?P<pipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	pipelineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := p.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, pipelineNameParam, nil)
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

func (p *PipelinesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if p.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelines/(?P<pipelineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	pipelineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("pipelineName")])
	if err != nil {
		return nil, err
	}
	ifNoneMatchParam := getOptional(getHeaderValue(req.Header, "If-None-Match"))
	var options *armdatafactory.PipelinesClientGetOptions
	if ifNoneMatchParam != nil {
		options = &armdatafactory.PipelinesClientGetOptions{
			IfNoneMatch: ifNoneMatchParam,
		}
	}
	respr, errRespr := p.srv.Get(req.Context(), resourceGroupNameParam, factoryNameParam, pipelineNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNotModified}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNotModified", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PipelineResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *PipelinesServerTransport) dispatchNewListByFactoryPager(req *http.Request) (*http.Response, error) {
	if p.srv.NewListByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFactoryPager not implemented")}
	}
	newListByFactoryPager := p.newListByFactoryPager.get(req)
	if newListByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/pipelines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		resp := p.srv.NewListByFactoryPager(resourceGroupNameParam, factoryNameParam, nil)
		newListByFactoryPager = &resp
		p.newListByFactoryPager.add(req, newListByFactoryPager)
		server.PagerResponderInjectNextLinks(newListByFactoryPager, req, func(page *armdatafactory.PipelinesClientListByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		p.newListByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFactoryPager) {
		p.newListByFactoryPager.remove(req)
	}
	return resp, nil
}
