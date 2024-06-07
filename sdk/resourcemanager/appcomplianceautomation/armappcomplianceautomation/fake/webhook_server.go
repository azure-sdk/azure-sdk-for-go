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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WebhookServer is a fake server for instances of the armappcomplianceautomation.WebhookClient type.
type WebhookServer struct {
	// CreateOrUpdate is the fake for method WebhookClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, reportName string, webhookName string, properties armappcomplianceautomation.WebhookResource, options *armappcomplianceautomation.WebhookClientCreateOrUpdateOptions) (resp azfake.Responder[armappcomplianceautomation.WebhookClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method WebhookClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, reportName string, webhookName string, options *armappcomplianceautomation.WebhookClientDeleteOptions) (resp azfake.Responder[armappcomplianceautomation.WebhookClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WebhookClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, reportName string, webhookName string, options *armappcomplianceautomation.WebhookClientGetOptions) (resp azfake.Responder[armappcomplianceautomation.WebhookClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method WebhookClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(reportName string, options *armappcomplianceautomation.WebhookClientListOptions) (resp azfake.PagerResponder[armappcomplianceautomation.WebhookClientListResponse])

	// Update is the fake for method WebhookClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, reportName string, webhookName string, properties armappcomplianceautomation.WebhookResourcePatch, options *armappcomplianceautomation.WebhookClientUpdateOptions) (resp azfake.Responder[armappcomplianceautomation.WebhookClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWebhookServerTransport creates a new instance of WebhookServerTransport with the provided implementation.
// The returned WebhookServerTransport instance is connected to an instance of armappcomplianceautomation.WebhookClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebhookServerTransport(srv *WebhookServer) *WebhookServerTransport {
	return &WebhookServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armappcomplianceautomation.WebhookClientListResponse]](),
	}
}

// WebhookServerTransport connects instances of armappcomplianceautomation.WebhookClient to instances of WebhookServer.
// Don't use this type directly, use NewWebhookServerTransport instead.
type WebhookServerTransport struct {
	srv          *WebhookServer
	newListPager *tracker[azfake.PagerResponder[armappcomplianceautomation.WebhookClientListResponse]]
}

// Do implements the policy.Transporter interface for WebhookServerTransport.
func (w *WebhookServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "WebhookClient.CreateOrUpdate":
		resp, err = w.dispatchCreateOrUpdate(req)
	case "WebhookClient.Delete":
		resp, err = w.dispatchDelete(req)
	case "WebhookClient.Get":
		resp, err = w.dispatchGet(req)
	case "WebhookClient.NewListPager":
		resp, err = w.dispatchNewListPager(req)
	case "WebhookClient.Update":
		resp, err = w.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (w *WebhookServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webhooks/(?P<webhookName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.WebhookResource](req)
	if err != nil {
		return nil, err
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	webhookNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webhookName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.CreateOrUpdate(req.Context(), reportNameParam, webhookNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebhookResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebhookServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if w.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webhooks/(?P<webhookName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	webhookNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webhookName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Delete(req.Context(), reportNameParam, webhookNameParam, nil)
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

func (w *WebhookServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webhooks/(?P<webhookName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	webhookNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webhookName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), reportNameParam, webhookNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebhookResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebhookServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := w.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webhooks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
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
		selectUnescaped, err := url.QueryUnescape(qp.Get("$select"))
		if err != nil {
			return nil, err
		}
		selectParam := getOptional(selectUnescaped)
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
		offerGUIDUnescaped, err := url.QueryUnescape(qp.Get("offerGuid"))
		if err != nil {
			return nil, err
		}
		offerGUIDParam := getOptional(offerGUIDUnescaped)
		reportCreatorTenantIDUnescaped, err := url.QueryUnescape(qp.Get("reportCreatorTenantId"))
		if err != nil {
			return nil, err
		}
		reportCreatorTenantIDParam := getOptional(reportCreatorTenantIDUnescaped)
		reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
		if err != nil {
			return nil, err
		}
		var options *armappcomplianceautomation.WebhookClientListOptions
		if skipTokenParam != nil || topParam != nil || selectParam != nil || filterParam != nil || orderbyParam != nil || offerGUIDParam != nil || reportCreatorTenantIDParam != nil {
			options = &armappcomplianceautomation.WebhookClientListOptions{
				SkipToken:             skipTokenParam,
				Top:                   topParam,
				Select:                selectParam,
				Filter:                filterParam,
				Orderby:               orderbyParam,
				OfferGUID:             offerGUIDParam,
				ReportCreatorTenantID: reportCreatorTenantIDParam,
			}
		}
		resp := w.srv.NewListPager(reportNameParam, options)
		newListPager = &resp
		w.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armappcomplianceautomation.WebhookClientListResponse, createLink func() string) {
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

func (w *WebhookServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/providers/Microsoft\.AppComplianceAutomation/reports/(?P<reportName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webhooks/(?P<webhookName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armappcomplianceautomation.WebhookResourcePatch](req)
	if err != nil {
		return nil, err
	}
	reportNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("reportName")])
	if err != nil {
		return nil, err
	}
	webhookNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webhookName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Update(req.Context(), reportNameParam, webhookNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebhookResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
