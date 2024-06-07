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

// WorkspaceSubscriptionServer is a fake server for instances of the armapimanagement.WorkspaceSubscriptionClient type.
type WorkspaceSubscriptionServer struct {
	// CreateOrUpdate is the fake for method WorkspaceSubscriptionClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, parameters armapimanagement.SubscriptionCreateParameters, options *armapimanagement.WorkspaceSubscriptionClientCreateOrUpdateOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WorkspaceSubscriptionClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, ifMatch string, options *armapimanagement.WorkspaceSubscriptionClientDeleteOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WorkspaceSubscriptionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, options *armapimanagement.WorkspaceSubscriptionClientGetOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientGetResponse], errResp azfake.ErrorResponder)

	// GetEntityTag is the fake for method WorkspaceSubscriptionClient.GetEntityTag
	// HTTP status codes to indicate success: http.StatusOK
	GetEntityTag func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, options *armapimanagement.WorkspaceSubscriptionClientGetEntityTagOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientGetEntityTagResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WorkspaceSubscriptionClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, serviceName string, workspaceID string, options *armapimanagement.WorkspaceSubscriptionClientListOptions) (resp azfake.PagerResponder[armapimanagement.WorkspaceSubscriptionClientListResponse])

	// ListSecrets is the fake for method WorkspaceSubscriptionClient.ListSecrets
	// HTTP status codes to indicate success: http.StatusOK
	ListSecrets func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, options *armapimanagement.WorkspaceSubscriptionClientListSecretsOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientListSecretsResponse], errResp azfake.ErrorResponder)

	// RegeneratePrimaryKey is the fake for method WorkspaceSubscriptionClient.RegeneratePrimaryKey
	// HTTP status codes to indicate success: http.StatusNoContent
	RegeneratePrimaryKey func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, options *armapimanagement.WorkspaceSubscriptionClientRegeneratePrimaryKeyOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientRegeneratePrimaryKeyResponse], errResp azfake.ErrorResponder)

	// RegenerateSecondaryKey is the fake for method WorkspaceSubscriptionClient.RegenerateSecondaryKey
	// HTTP status codes to indicate success: http.StatusNoContent
	RegenerateSecondaryKey func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, options *armapimanagement.WorkspaceSubscriptionClientRegenerateSecondaryKeyOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientRegenerateSecondaryKeyResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method WorkspaceSubscriptionClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, sid string, ifMatch string, parameters armapimanagement.SubscriptionUpdateParameters, options *armapimanagement.WorkspaceSubscriptionClientUpdateOptions) (resp azfake.Responder[armapimanagement.WorkspaceSubscriptionClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWorkspaceSubscriptionServerTransport creates a new instance of WorkspaceSubscriptionServerTransport with the provided implementation.
// The returned WorkspaceSubscriptionServerTransport instance is connected to an instance of armapimanagement.WorkspaceSubscriptionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWorkspaceSubscriptionServerTransport(srv *WorkspaceSubscriptionServer) *WorkspaceSubscriptionServerTransport {
	return &WorkspaceSubscriptionServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armapimanagement.WorkspaceSubscriptionClientListResponse]](),
	}
}

// WorkspaceSubscriptionServerTransport connects instances of armapimanagement.WorkspaceSubscriptionClient to instances of WorkspaceSubscriptionServer.
// Don't use this type directly, use NewWorkspaceSubscriptionServerTransport instead.
type WorkspaceSubscriptionServerTransport struct {
	srv          *WorkspaceSubscriptionServer
	newListPager *tracker[azfake.PagerResponder[armapimanagement.WorkspaceSubscriptionClientListResponse]]
}

// Do implements the policy.Transporter interface for WorkspaceSubscriptionServerTransport.
func (w *WorkspaceSubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WorkspaceSubscriptionClient.CreateOrUpdate":
		resp, err = w.dispatchCreateOrUpdate(req)
	case "WorkspaceSubscriptionClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WorkspaceSubscriptionClient.Get":
		resp, err = w.dispatchGet(req)
	case "WorkspaceSubscriptionClient.GetEntityTag":
		resp, err = w.dispatchGetEntityTag(req)
	case "WorkspaceSubscriptionClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WorkspaceSubscriptionClient.ListSecrets":
		resp, err = w.dispatchListSecrets(req)
	case "WorkspaceSubscriptionClient.RegeneratePrimaryKey":
		resp, err = w.dispatchRegeneratePrimaryKey(req)
	case "WorkspaceSubscriptionClient.RegenerateSecondaryKey":
		resp, err = w.dispatchRegenerateSecondaryKey(req)
	case "WorkspaceSubscriptionClient.Update":
		resp, err = w.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.SubscriptionCreateParameters](req)
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	notifyUnescaped, err := url.QueryUnescape(qp.Get("notify"))
	if err != nil {
		return nil, err
	}
	notifyParam, err := parseOptional(notifyUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	ifMatchParam := getOptional(getHeaderValue(req.Header, "If-Match"))
	appTypeUnescaped, err := url.QueryUnescape(qp.Get("appType"))
	if err != nil {
		return nil, err
	}
	appTypeParam := getOptional(armapimanagement.AppType(appTypeUnescaped))
	var options *armapimanagement.WorkspaceSubscriptionClientCreateOrUpdateOptions
	if notifyParam != nil || ifMatchParam != nil || appTypeParam != nil {
		options = &armapimanagement.WorkspaceSubscriptionClientCreateOrUpdateOptions{
			Notify:  notifyParam,
			IfMatch: ifMatchParam,
			AppType: appTypeParam,
		}
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, getHeaderValue(req.Header, "If-Match"), nil)
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

func (w *WorkspaceSubscriptionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchGetEntityTag(req *http.Request) (*http.Response, error) {
	if w.srv.GetEntityTag == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetEntityTag not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.GetEntityTag(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, nil)
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

func (w *WorkspaceSubscriptionServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions`
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
		serviceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceName")])
		if err != nil {
			return nil, err
		}
		workspaceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceId")])
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
		var options *armapimanagement.WorkspaceSubscriptionClientListOptions
		if filterParam != nil || topParam != nil || skipParam != nil {
			options = &armapimanagement.WorkspaceSubscriptionClientListOptions{
				Filter: filterParam,
				Top:    topParam,
				Skip:   skipParam,
			}
		}
		resp := w.srv.NewListPager(resourceGroupNameParam, serviceNameParam, workspaceIDParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armapimanagement.WorkspaceSubscriptionClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		w.newListPager.remove(req)
	}
	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchListSecrets(req *http.Request) (*http.Response, error) {
	if w.srv.ListSecrets == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSecrets not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listSecrets`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.ListSecrets(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionKeysContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchRegeneratePrimaryKey(req *http.Request) (*http.Response, error) {
	if w.srv.RegeneratePrimaryKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegeneratePrimaryKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regeneratePrimaryKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.RegeneratePrimaryKey(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchRegenerateSecondaryKey(req *http.Request) (*http.Response, error) {
	if w.srv.RegenerateSecondaryKey == nil {
		return nil, &nonRetriableError{errors.New("fake for method RegenerateSecondaryKey not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/regenerateSecondaryKey`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.RegenerateSecondaryKey(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WorkspaceSubscriptionServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ApiManagement/service/(?P<serviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/workspaces/(?P<workspaceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/subscriptions/(?P<sid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	body, err := server.UnmarshalRequestAsJSON[armapimanagement.SubscriptionUpdateParameters](req)
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
	sidParam, err := url.PathUnescape(matches[regex.SubexpIndex("sid")])
	if err != nil {
		return nil, err
	}
	notifyUnescaped, err := url.QueryUnescape(qp.Get("notify"))
	if err != nil {
		return nil, err
	}
	notifyParam, err := parseOptional(notifyUnescaped, strconv.ParseBool)
	if err != nil {
		return nil, err
	}
	appTypeUnescaped, err := url.QueryUnescape(qp.Get("appType"))
	if err != nil {
		return nil, err
	}
	appTypeParam := getOptional(armapimanagement.AppType(appTypeUnescaped))
	var options *armapimanagement.WorkspaceSubscriptionClientUpdateOptions
	if notifyParam != nil || appTypeParam != nil {
		options = &armapimanagement.WorkspaceSubscriptionClientUpdateOptions{
			Notify:  notifyParam,
			AppType: appTypeParam,
		}
	}
	respr, errRespr := w.srv.Update(req.Context(), resourceGroupNameParam, serviceNameParam, workspaceIDParam, sidParam, getHeaderValue(req.Header, "If-Match"), body, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SubscriptionContract, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).ETag; val != nil {
		resp.Header.Set("ETag", *val)
	}
	return resp, nil
}