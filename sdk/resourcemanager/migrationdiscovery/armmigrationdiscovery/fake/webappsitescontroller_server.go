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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrationdiscovery/armmigrationdiscovery"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// WebAppSitesControllerServer is a fake server for instances of the armmigrationdiscovery.WebAppSitesControllerClient type.
type WebAppSitesControllerServer struct {
	// Create is the fake for method WebAppSitesControllerClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body armmigrationdiscovery.WebAppSite, options *armmigrationdiscovery.WebAppSitesControllerClientCreateOptions) (resp azfake.Responder[armmigrationdiscovery.WebAppSitesControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method WebAppSitesControllerClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *armmigrationdiscovery.WebAppSitesControllerClientBeginDeleteOptions) (resp azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// ErrorSummary is the fake for method WebAppSitesControllerClient.ErrorSummary
	// HTTP status codes to indicate success: http.StatusOK
	ErrorSummary func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body armmigrationdiscovery.ErrorSummaryRequest, options *armmigrationdiscovery.WebAppSitesControllerClientErrorSummaryOptions) (resp azfake.Responder[armmigrationdiscovery.WebAppSitesControllerClientErrorSummaryResponse], errResp azfake.ErrorResponder)

	// BeginExportInventory is the fake for method WebAppSitesControllerClient.BeginExportInventory
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportInventory func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body armmigrationdiscovery.ExportWebAppsRequest, options *armmigrationdiscovery.WebAppSitesControllerClientBeginExportInventoryOptions) (resp azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientExportInventoryResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method WebAppSitesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *armmigrationdiscovery.WebAppSitesControllerClientGetOptions) (resp azfake.Responder[armmigrationdiscovery.WebAppSitesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByMasterSitePager is the fake for method WebAppSitesControllerClient.NewListByMasterSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByMasterSitePager func(resourceGroupName string, siteName string, options *armmigrationdiscovery.WebAppSitesControllerClientListByMasterSiteOptions) (resp azfake.PagerResponder[armmigrationdiscovery.WebAppSitesControllerClientListByMasterSiteResponse])

	// BeginRefresh is the fake for method WebAppSitesControllerClient.BeginRefresh
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefresh func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body armmigrationdiscovery.ProxySiteRefreshBody, options *armmigrationdiscovery.WebAppSitesControllerClientBeginRefreshOptions) (resp azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientRefreshResponse], errResp azfake.ErrorResponder)

	// Summary is the fake for method WebAppSitesControllerClient.Summary
	// HTTP status codes to indicate success: http.StatusOK
	Summary func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *armmigrationdiscovery.WebAppSitesControllerClientSummaryOptions) (resp azfake.Responder[armmigrationdiscovery.WebAppSitesControllerClientSummaryResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method WebAppSitesControllerClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, properties armmigrationdiscovery.WebAppSiteUpdate, options *armmigrationdiscovery.WebAppSitesControllerClientBeginUpdateOptions) (resp azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewWebAppSitesControllerServerTransport creates a new instance of WebAppSitesControllerServerTransport with the provided implementation.
// The returned WebAppSitesControllerServerTransport instance is connected to an instance of armmigrationdiscovery.WebAppSitesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewWebAppSitesControllerServerTransport(srv *WebAppSitesControllerServer) *WebAppSitesControllerServerTransport {
	return &WebAppSitesControllerServerTransport{
		srv:                      srv,
		beginDelete:              newTracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientDeleteResponse]](),
		beginExportInventory:     newTracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientExportInventoryResponse]](),
		newListByMasterSitePager: newTracker[azfake.PagerResponder[armmigrationdiscovery.WebAppSitesControllerClientListByMasterSiteResponse]](),
		beginRefresh:             newTracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientRefreshResponse]](),
		beginUpdate:              newTracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientUpdateResponse]](),
	}
}

// WebAppSitesControllerServerTransport connects instances of armmigrationdiscovery.WebAppSitesControllerClient to instances of WebAppSitesControllerServer.
// Don't use this type directly, use NewWebAppSitesControllerServerTransport instead.
type WebAppSitesControllerServerTransport struct {
	srv                      *WebAppSitesControllerServer
	beginDelete              *tracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientDeleteResponse]]
	beginExportInventory     *tracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientExportInventoryResponse]]
	newListByMasterSitePager *tracker[azfake.PagerResponder[armmigrationdiscovery.WebAppSitesControllerClientListByMasterSiteResponse]]
	beginRefresh             *tracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientRefreshResponse]]
	beginUpdate              *tracker[azfake.PollerResponder[armmigrationdiscovery.WebAppSitesControllerClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for WebAppSitesControllerServerTransport.
func (w *WebAppSitesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return w.dispatchToMethodFake(req, method)
}

func (w *WebAppSitesControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if webAppSitesControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = webAppSitesControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "WebAppSitesControllerClient.Create":
				res.resp, res.err = w.dispatchCreate(req)
			case "WebAppSitesControllerClient.BeginDelete":
				res.resp, res.err = w.dispatchBeginDelete(req)
			case "WebAppSitesControllerClient.ErrorSummary":
				res.resp, res.err = w.dispatchErrorSummary(req)
			case "WebAppSitesControllerClient.BeginExportInventory":
				res.resp, res.err = w.dispatchBeginExportInventory(req)
			case "WebAppSitesControllerClient.Get":
				res.resp, res.err = w.dispatchGet(req)
			case "WebAppSitesControllerClient.NewListByMasterSitePager":
				res.resp, res.err = w.dispatchNewListByMasterSitePager(req)
			case "WebAppSitesControllerClient.BeginRefresh":
				res.resp, res.err = w.dispatchBeginRefresh(req)
			case "WebAppSitesControllerClient.Summary":
				res.resp, res.err = w.dispatchSummary(req)
			case "WebAppSitesControllerClient.BeginUpdate":
				res.resp, res.err = w.dispatchBeginUpdate(req)
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

func (w *WebAppSitesControllerServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if w.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.WebAppSite](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Create(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebAppSite, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if w.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := w.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginDelete(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		w.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		w.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		w.beginDelete.remove(req)
	}

	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchErrorSummary(req *http.Request) (*http.Response, error) {
	if w.srv.ErrorSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ErrorSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/errorSummary`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.ErrorSummaryRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.ErrorSummary(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SiteErrorSummary, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchBeginExportInventory(req *http.Request) (*http.Response, error) {
	if w.srv.BeginExportInventory == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportInventory not implemented")}
	}
	beginExportInventory := w.beginExportInventory.get(req)
	if beginExportInventory == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportInventory`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.ExportWebAppsRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginExportInventory(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportInventory = &respr
		w.beginExportInventory.add(req, beginExportInventory)
	}

	resp, err := server.PollerResponderNext(beginExportInventory, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginExportInventory.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportInventory) {
		w.beginExportInventory.remove(req)
	}

	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if w.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebAppSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchNewListByMasterSitePager(req *http.Request) (*http.Response, error) {
	if w.srv.NewListByMasterSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByMasterSitePager not implemented")}
	}
	newListByMasterSitePager := w.newListByMasterSitePager.get(req)
	if newListByMasterSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		resp := w.srv.NewListByMasterSitePager(resourceGroupNameParam, siteNameParam, nil)
		newListByMasterSitePager = &resp
		w.newListByMasterSitePager.add(req, newListByMasterSitePager)
		server.PagerResponderInjectNextLinks(newListByMasterSitePager, req, func(page *armmigrationdiscovery.WebAppSitesControllerClientListByMasterSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByMasterSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		w.newListByMasterSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByMasterSitePager) {
		w.newListByMasterSitePager.remove(req)
	}
	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchBeginRefresh(req *http.Request) (*http.Response, error) {
	if w.srv.BeginRefresh == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefresh not implemented")}
	}
	beginRefresh := w.beginRefresh.get(req)
	if beginRefresh == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refresh`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.ProxySiteRefreshBody](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginRefresh(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefresh = &respr
		w.beginRefresh.add(req, beginRefresh)
	}

	resp, err := server.PollerResponderNext(beginRefresh, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginRefresh.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefresh) {
		w.beginRefresh.remove(req)
	}

	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchSummary(req *http.Request) (*http.Response, error) {
	if w.srv.Summary == nil {
		return nil, &nonRetriableError{errors.New("fake for method Summary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/summary`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
	if err != nil {
		return nil, err
	}
	webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := w.srv.Summary(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).WebAppSiteUsage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (w *WebAppSitesControllerServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if w.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := w.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/masterSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/webAppSites/(?P<webAppSiteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrationdiscovery.WebAppSiteUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		webAppSiteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("webAppSiteName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := w.srv.BeginUpdate(req.Context(), resourceGroupNameParam, siteNameParam, webAppSiteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		w.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		w.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		w.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to WebAppSitesControllerServerTransport
var webAppSitesControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
