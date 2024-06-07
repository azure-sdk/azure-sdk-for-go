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
	"strconv"
)

// WorkspaceAPIDiagnosticServer is a fake server for instances of the armapimanagement.WorkspaceAPIDiagnosticClient type.
type WorkspaceAPIDiagnosticServer struct {
	// CreateOrUpdate is the fake for method WorkspaceAPIDiagnosticClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, diagnosticID string, parameters armapimanagement.DiagnosticContract, options *armapimanagement.WorkspaceAPIDiagnosticClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIDiagnosticClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WorkspaceAPIDiagnosticClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, diagnosticID string, ifMatch string, options *armapimanagement.WorkspaceAPIDiagnosticClientDeleteOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIDiagnosticClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspaceAPIDiagnosticClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, diagnosticID string, options *armapimanagement.WorkspaceAPIDiagnosticClientGetOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIDiagnosticClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method WorkspaceAPIDiagnosticClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, diagnosticID string, options *armapimanagement.WorkspaceAPIDiagnosticClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIDiagnosticClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByWorkspacePager is the fake for method WorkspaceAPIDiagnosticClient.NewListByWorkspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByWorkspacePager func(resourceGroupName string, serviceName string, workspaceID string, apiID string, options *armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceOptions) (resp azfake.PagerResponder[armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceResponse])

	// Update is the fake for method WorkspaceAPIDiagnosticClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, apiID string, diagnosticID string, ifMatch string, parameters armapimanagement.DiagnosticUpdateContract, options *armapimanagement.WorkspaceAPIDiagnosticClientUpdateOptions) (resp azfake.Responder[armapimanagement.WorkspaceAPIDiagnosticClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWorkspaceAPIDiagnosticServerTransport creates a new instance of WorkspaceAPIDiagnosticServerTransport with the provided implementation.
// The returned WorkspaceAPIDiagnosticServerTransport instance is connected to an instance of armapimanagement.WorkspaceAPIDiagnosticClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceAPIDiagnosticServerTransport(srv *WorkspaceAPIDiagnosticServer) *WorkspaceAPIDiagnosticServerTransport {
	return &WorkspaceAPIDiagnosticServerTransport{
		srv:                     srv,
		newListByWorkspacePager: newTracker[azfake.PagerResponder[armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceResponse]](),
	}
}

// WorkspaceAPIDiagnosticServerTransport connects instances of armapimanagement.WorkspaceAPIDiagnosticClient to instances of WorkspaceAPIDiagnosticServer.
// Don't use this type directly, use NewWorkspaceAPIDiagnosticServerTransport instead.
type WorkspaceAPIDiagnosticServerTransport struct {
	srv                     *WorkspaceAPIDiagnosticServer
	newListByWorkspacePager *tracker[azfake.PagerResponder[armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceAPIDiagnosticServerTransport.
func (w *WorkspaceAPIDiagnosticServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceAPIDiagnosticClient.CreateOrUpdate":
		resp, err = w.dispatchCreateOrUpdate(req)
	case "WorkspaceAPIDiagnosticClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WorkspaceAPIDiagnosticClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspaceAPIDiagnosticClient.GetEntityTag":
		resp, err = w.dispatchGetEntityTag(req)
	case "WorkspaceAPIDiagnosticClient.NewListByWorkspacePager":
		resp, err = w.dispatchNewListByWorkspacePager(req)
	case "WorkspaceAPIDiagnosticClient.Update":
		resp, err = w.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceAPIDiagnosticServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnostics/(?P<diagnosticId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.DiagnosticContract](req)
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
	diagnosticIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("diagnosticId")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.WorkspaceAPIDiagnosticClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.WorkspaceAPIDiagnosticClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, diagnosticIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiagnosticContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceAPIDiagnosticServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnostics/(?P<diagnosticId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	diagnosticIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("diagnosticId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, diagnosticIDParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (w *WorkspaceAPIDiagnosticServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnostics/(?P<diagnosticId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	diagnosticIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("diagnosticId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, diagnosticIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiagnosticContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceAPIDiagnosticServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if w.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnostics/(?P<diagnosticId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	diagnosticIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("diagnosticId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, diagnosticIDParam, nil)
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

func (w *WorkspaceAPIDiagnosticServerTransport) dispatchNewListByWorkspacePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByWorkspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByWorkspacePager not implemented")}
	}
	newListByWorkspacePager := w.newListByWorkspacePager.get(req)
	if newListByWorkspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnostics`
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
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		var options *armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := w.srv.NewListByWorkspacePager(resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, options)
		newListByWorkspacePager = &resp
		w.newListByWorkspacePager.add(req, newListByWorkspacePager)
		server.PagerResponderInjectNextLinks(newListByWorkspacePager, req, func(page *armapimanagement.WorkspaceAPIDiagnosticClientListByWorkspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByWorkspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByWorkspacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByWorkspacePager) {
		w.newListByWorkspacePager.remove(req)
	}
	return resp, nil
}

func (w *WorkspaceAPIDiagnosticServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/diagnostics/(?P<diagnosticId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.DiagnosticUpdateContract](req)
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
	diagnosticIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("diagnosticId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Update(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, apiIDParam, diagnosticIDParam, getHeaderValue(req.Header, "If-Match"), body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DiagnosticContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}
