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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement/v3"
	"net/http"
	"net/url"
	"regexp"
)

// WorkspaceAPIOperationPolicyServer is a fake server for instances of the armapimanagement.WorkspaceAPIOperationPolicyClient type.
type WorkspaceAPIOperationPolicyServer struct {
	// CreateOrUpdate is the fake for method WorkspaceAPIOperationPolicyClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID armapimanagement.PolicyIDName, parameters armapimanagement.PolicyContract, options *armapimanagement.WorkspaceAPIOperationPolicyClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIOperationPolicyClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WorkspaceAPIOperationPolicyClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID armapimanagement.PolicyIDName, ifMatch string, options *armapimanagement.WorkspaceAPIOperationPolicyClientDeleteOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIOperationPolicyClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspaceAPIOperationPolicyClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID armapimanagement.PolicyIDName, options *armapimanagement.WorkspaceAPIOperationPolicyClientGetOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIOperationPolicyClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method WorkspaceAPIOperationPolicyClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, policyID armapimanagement.PolicyIDName, options *armapimanagement.WorkspaceAPIOperationPolicyClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIOperationPolicyClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByOperationPager is the fake for method WorkspaceAPIOperationPolicyClient.NewListByOperationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByOperationPager func(resourceGroupName string, serviceName string, workspaceID string, apiID string, operationID string, options *armapimanagement.WorkspaceAPIOperationPolicyClientListByOperationOptions) (resp azfake.PagerResponder[armapimanagement.WorkspaceAPIOperationPolicyClientListByOperationResponse])
}

// NewWorkspaceAPIOperationPolicyServerTransport creates a new instance of WorkspaceAPIOperationPolicyServerTransport with the provided implementation.
// The returned WorkspaceAPIOperationPolicyServerTransport instance is connected to an instance of armapimanagement.WorkspaceAPIOperationPolicyClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceAPIOperationPolicyServerTransport(srv *WorkspaceAPIOperationPolicyServer) *WorkspaceAPIOperationPolicyServerTransport {
	return &WorkspaceAPIOperationPolicyServerTransport{
		srv:                     srv,
		newListByOperationPager: newTracker[azfake.PagerResponder[armapimanagement.WorkspaceAPIOperationPolicyClientListByOperationResponse]](),
	}
}

// WorkspaceAPIOperationPolicyServerTransport connects instances of armapimanagement.WorkspaceAPIOperationPolicyClient to instances of WorkspaceAPIOperationPolicyServer.
// Don't use this type directly, use NewWorkspaceAPIOperationPolicyServerTransport instead.
type WorkspaceAPIOperationPolicyServerTransport struct {
	srv                     *WorkspaceAPIOperationPolicyServer
	newListByOperationPager *tracker[azfake.PagerResponder[armapimanagement.WorkspaceAPIOperationPolicyClientListByOperationResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceAPIOperationPolicyServerTransport.
func (w *WorkspaceAPIOperationPolicyServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceAPIOperationPolicyClient.CreateOrUpdate":
		resp, err = w.dispatchCreateOrUpdate(req)
	case "WorkspaceAPIOperationPolicyClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WorkspaceAPIOperationPolicyClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspaceAPIOperationPolicyClient.GetEntityTag":
		resp, err = w.dispatchGetEntityTag(req)
	case "WorkspaceAPIOperationPolicyClient.NewListByOperationPager":
		resp, err = w.dispatchNewListByOperationPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceAPIOperationPolicyServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.PolicyContract](req)
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
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.WorkspaceAPIOperationPolicyClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.WorkspaceAPIOperationPolicyClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, operationIDParam, policyIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceAPIOperationPolicyServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
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
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, operationIDParam, policyIDParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (w *WorkspaceAPIOperationPolicyServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
	if err != nil {
		return nil, err
	}
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	formatUnescaped, err := url.QueryUnescape(qp.Get("format"))
	if err != nil {
		return nil, err
	}
	formatParam := getOptional(armapimanagement.PolicyExportFormat(formatUnescaped))
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	var options *armapimanagement.WorkspaceAPIOperationPolicyClientGetOptions
	if formatParam != nil {
		options = &armapimanagement.WorkspaceAPIOperationPolicyClientGetOptions{
			Format: formatParam,
		}
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, operationIDParam, policyIDParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).PolicyContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceAPIOperationPolicyServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if w.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies/(?P<policyId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
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
	workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
	if err != nil {
		return nil, err
	}
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
	if err != nil {
		return nil, err
	}
	policyIDParam, err := parseWithCast(matches[regex.SubexpIndex("policyId")], func(v string) (armapimanagement.PolicyIDName, error) {
		p, unescapeErr := url.PathUnescape(v)
		if unescapeErr != nil {
			return "", unescapeErr
		}
		return armapimanagement.PolicyIDName(p), nil
	})
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, operationIDParam, policyIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceAPIOperationPolicyServerTransport) dispatchNewListByOperationPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByOperationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByOperationPager not implemented")}
	}
	newListByOperationPager := w.newListByOperationPager.get(req)
	if newListByOperationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/operations/(?P<operationId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/policies`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
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
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
		if err != nil {
			return nil, err
		}
		apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
		if err != nil {
			return nil, err
		}
		operationIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("operationId")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByOperationPager(resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, operationIDParam, nil)
		newListByOperationPager = &resp
		w.newListByOperationPager.add(req, newListByOperationPager)
		server.PagerResponderInjectNextLinks(newListByOperationPager, req, func(page *armapimanagement.WorkspaceAPIOperationPolicyClientListByOperationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByOperationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByOperationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByOperationPager) {
		w.newListByOperationPager.remove(req)
	}
	return resp, nil
}