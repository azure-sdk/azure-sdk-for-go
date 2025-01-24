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

// APIIssueAttachmentServer is a fake server for instances of the armapimanagement.APIIssueAttachmentClient type.
type APIIssueAttachmentServer struct {
	// CreateOrUpdate is the fake for method APIIssueAttachmentClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, parameters armapimanagement.IssueAttachmentContract, options *armapimanagement.APIIssueAttachmentClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.APIIssueAttachmentClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method APIIssueAttachmentClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, ifMatch string, options *armapimanagement.APIIssueAttachmentClientDeleteOptions) (resp azfake.Responder[armapimanagement.APIIssueAttachmentClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method APIIssueAttachmentClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, options *armapimanagement.APIIssueAttachmentClientGetOptions) (resp azfake.Responder[armapimanagement.APIIssueAttachmentClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method APIIssueAttachmentClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, apiID string, issueID string, attachmentID string, options *armapimanagement.APIIssueAttachmentClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.APIIssueAttachmentClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListByServicePager is the fake for method APIIssueAttachmentClient.NewListByServicePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByServicePager func(resourceGroupName string, serviceName string, apiID string, issueID string, options *armapimanagement.APIIssueAttachmentClientListByServiceOptions) (resp azfake.PagerResponder[armapimanagement.APIIssueAttachmentClientListByServiceResponse])
}

// NewAPIIssueAttachmentServerTransport creates a new instance of APIIssueAttachmentServerTransport with the provided implementation.
// The returned APIIssueAttachmentServerTransport instance is connected to an instance of armapimanagement.APIIssueAttachmentClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAPIIssueAttachmentServerTransport(srv *APIIssueAttachmentServer) *APIIssueAttachmentServerTransport {
	return &APIIssueAttachmentServerTransport{
		srv:                   srv,
		newListByServicePager: newTracker[azfake.PagerResponder[armapimanagement.APIIssueAttachmentClientListByServiceResponse]](),
	}
}

// APIIssueAttachmentServerTransport connects instances of armapimanagement.APIIssueAttachmentClient to instances of APIIssueAttachmentServer.
// Don't use this type directly, use NewAPIIssueAttachmentServerTransport instead.
type APIIssueAttachmentServerTransport struct {
	srv                   *APIIssueAttachmentServer
	newListByServicePager *tracker[azfake.PagerResponder[armapimanagement.APIIssueAttachmentClientListByServiceResponse]]
}

// Do implements the policy.Transporter interface for APIIssueAttachmentServerTransport.
func (a *APIIssueAttachmentServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "APIIssueAttachmentClient.CreateOrUpdate":
		resp, err = a.dispatchCreateOrUpdate(req)
	case "APIIssueAttachmentClient.Delete":
		resp, err = a.dispatchDelete(req)
	case "APIIssueAttachmentClient.Get":
		resp, err = a.dispatchGet(req)
	case "APIIssueAttachmentClient.GetEntityTag":
		resp, err = a.dispatchGetEntityTag(req)
	case "APIIssueAttachmentClient.NewListByServicePager":
		resp, err = a.dispatchNewListByServicePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *APIIssueAttachmentServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/issues/(?P<issueId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/attachments/(?P<attachmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.IssueAttachmentContract](req)
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	issueIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("issueId")])
	if err != nil {
		return nil, err
	}
	attachmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("attachmentId")])
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	var options *armapimanagement.APIIssueAttachmentClientCreateOrUpdateOptions
	if ifMatchParam != nil {
		options = &armapimanagement.APIIssueAttachmentClientCreateOrUpdateOptions{
			IfMatch: ifMatchParam,
		}
	}
	respr, errRespr := a.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, issueIDParam, attachmentIDParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IssueAttachmentContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (a *APIIssueAttachmentServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if a.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/issues/(?P<issueId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/attachments/(?P<attachmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	issueIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("issueId")])
	if err != nil {
		return nil, err
	}
	attachmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("attachmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, issueIDParam, attachmentIDParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (a *APIIssueAttachmentServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/issues/(?P<issueId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/attachments/(?P<attachmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	issueIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("issueId")])
	if err != nil {
		return nil, err
	}
	attachmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("attachmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, issueIDParam, attachmentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IssueAttachmentContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (a *APIIssueAttachmentServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if a.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/issues/(?P<issueId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/attachments/(?P<attachmentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
	if err != nil {
		return nil, err
	}
	issueIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("issueId")])
	if err != nil {
		return nil, err
	}
	attachmentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("attachmentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, apiIDParam, issueIDParam, attachmentIDParam, nil)
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

func (a *APIIssueAttachmentServerTransport) dispatchNewListByServicePager(req *http.Request) (*http.Response, error) {
	if a.srv.NewListByServicePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByServicePager not implemented")}
	}
	newListByServicePager := a.newListByServicePager.get(req)
	if newListByServicePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/apis/(?P<apiId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/issues/(?P<issueId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/attachments`
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
		apiIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("apiId")])
		if err != nil {
			return nil, err
		}
		issueIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("issueId")])
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
		var options *armapimanagement.APIIssueAttachmentClientListByServiceOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.APIIssueAttachmentClientListByServiceOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := a.srv.NewListByServicePager(resourceGroupNameParam, serviceNameParam, apiIDParam, issueIDParam, options)
		newListByServicePager = &resp
		a.newListByServicePager.add(req, newListByServicePager)
		server.PagerResponderInjectNextLinks(newListByServicePager, req, func(page *armapimanagement.APIIssueAttachmentClientListByServiceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByServicePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		a.newListByServicePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByServicePager) {
		a.newListByServicePager.remove(req)
	}
	return resp, nil
}
