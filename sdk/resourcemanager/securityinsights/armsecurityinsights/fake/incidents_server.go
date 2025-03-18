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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/securityinsights/armsecurityinsights/v2"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
)

// IncidentsServer is a fake server for instances of the armsecurityinsights.IncidentsClient type.
type IncidentsServer struct {
	// CreateOrUpdate is the fake for method IncidentsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, incident armsecurityinsights.Incident, options *armsecurityinsights.IncidentsClientCreateOrUpdateOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method IncidentsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *armsecurityinsights.IncidentsClientDeleteOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IncidentsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *armsecurityinsights.IncidentsClientGetOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IncidentsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, workspaceName string, options *armsecurityinsights.IncidentsClientListOptions) (resp azfake.PagerResponder[armsecurityinsights.IncidentsClientListResponse])

	// ListAlerts is the fake for method IncidentsClient.ListAlerts
	// HTTP status codes to indicate success: http.StatusOK
	ListAlerts func(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *armsecurityinsights.IncidentsClientListAlertsOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientListAlertsResponse], errResp azfake.ErrorResponder)

	// ListBookmarks is the fake for method IncidentsClient.ListBookmarks
	// HTTP status codes to indicate success: http.StatusOK
	ListBookmarks func(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *armsecurityinsights.IncidentsClientListBookmarksOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientListBookmarksResponse], errResp azfake.ErrorResponder)

	// ListEntities is the fake for method IncidentsClient.ListEntities
	// HTTP status codes to indicate success: http.StatusOK
	ListEntities func(ctx context.Context, resourceGroupName string, workspaceName string, incidentID string, options *armsecurityinsights.IncidentsClientListEntitiesOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientListEntitiesResponse], errResp azfake.ErrorResponder)

	// RunPlaybook is the fake for method IncidentsClient.RunPlaybook
	// HTTP status codes to indicate success: http.StatusNoContent
	RunPlaybook func(ctx context.Context, resourceGroupName string, workspaceName string, incidentIdentifier string, options *armsecurityinsights.IncidentsClientRunPlaybookOptions) (resp azfake.Responder[armsecurityinsights.IncidentsClientRunPlaybookResponse], errResp azfake.ErrorResponder)
}

// NewIncidentsServerTransport creates a new instance of IncidentsServerTransport with the provided implementation.
// The returned IncidentsServerTransport instance is connected to an instance of armsecurityinsights.IncidentsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIncidentsServerTransport(srv *IncidentsServer) *IncidentsServerTransport {
	return &IncidentsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsecurityinsights.IncidentsClientListResponse]](),
	}
}

// IncidentsServerTransport connects instances of armsecurityinsights.IncidentsClient to instances of IncidentsServer.
// Don't use this type directly, use NewIncidentsServerTransport instead.
type IncidentsServerTransport struct {
	srv          *IncidentsServer
	newListPager *tracker[azfake.PagerResponder[armsecurityinsights.IncidentsClientListResponse]]
}

// Do implements the policy.Transporter interface for IncidentsServerTransport.
func (i *IncidentsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return i.dispatchToMethodFake(req, method)
}

func (i *IncidentsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if incidentsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = incidentsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "IncidentsClient.CreateOrUpdate":
				res.resp, res.err = i.dispatchCreateOrUpdate(req)
			case "IncidentsClient.Delete":
				res.resp, res.err = i.dispatchDelete(req)
			case "IncidentsClient.Get":
				res.resp, res.err = i.dispatchGet(req)
			case "IncidentsClient.NewListPager":
				res.resp, res.err = i.dispatchNewListPager(req)
			case "IncidentsClient.ListAlerts":
				res.resp, res.err = i.dispatchListAlerts(req)
			case "IncidentsClient.ListBookmarks":
				res.resp, res.err = i.dispatchListBookmarks(req)
			case "IncidentsClient.ListEntities":
				res.resp, res.err = i.dispatchListEntities(req)
			case "IncidentsClient.RunPlaybook":
				res.resp, res.err = i.dispatchRunPlaybook(req)
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

func (i *IncidentsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.Incident](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIDParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Incident, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IncidentsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIDParam, nil)
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

func (i *IncidentsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Incident, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IncidentsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents`
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
		workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
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
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armsecurityinsights.IncidentsClientListOptions
		if filterParam != nil || orderbyParam != nil || topParam != nil || skipTokenParam != nil {
			options = &armsecurityinsights.IncidentsClientListOptions{
				Filter:    filterParam,
				Orderby:   orderbyParam,
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := i.srv.NewListPager(resourceGroupNameParam, workspaceNameParam, options)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsecurityinsights.IncidentsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

func (i *IncidentsServerTransport) dispatchListAlerts(req *http.Request) (*http.Response, error) {
	if i.srv.ListAlerts == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListAlerts not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/alerts`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.ListAlerts(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IncidentAlertList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IncidentsServerTransport) dispatchListBookmarks(req *http.Request) (*http.Response, error) {
	if i.srv.ListBookmarks == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListBookmarks not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/bookmarks`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.ListBookmarks(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IncidentBookmarkList, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IncidentsServerTransport) dispatchListEntities(req *http.Request) (*http.Response, error) {
	if i.srv.ListEntities == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListEntities not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/entities`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.ListEntities(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IncidentEntitiesResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IncidentsServerTransport) dispatchRunPlaybook(req *http.Request) (*http.Response, error) {
	if i.srv.RunPlaybook == nil {
		return nil, &nonRetriableError{errors.New("fake for method RunPlaybook not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OperationalInsights/workspaces/(?P<workspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.SecurityInsights/incidents/(?P<incidentIdentifier>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runPlaybook`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsecurityinsights.ManualTriggerRequestBody](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	workspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("workspaceName")])
	if err != nil {
		return nil, err
	}
	incidentIdentifierParam, err := url.PathUnescape(matches[regex.SubexpIndex("incidentIdentifier")])
	if err != nil {
		return nil, err
	}
	var options *armsecurityinsights.IncidentsClientRunPlaybookOptions
	if !reflect.ValueOf(body).IsZero() {
		options = &armsecurityinsights.IncidentsClientRunPlaybookOptions{
			RequestBody: &body,
		}
	}
	respr, errRespr := i.srv.RunPlaybook(req.Context(), resourceGroupNameParam, workspaceNameParam, incidentIdentifierParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Interface, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to IncidentsServerTransport
var incidentsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
