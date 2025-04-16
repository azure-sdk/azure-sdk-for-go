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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// HypervSitesControllerServer is a fake server for instances of the armmigrate.HypervSitesControllerClient type.
type HypervSitesControllerServer struct {
	// ComputeErrorSummary is the fake for method HypervSitesControllerClient.ComputeErrorSummary
	// HTTP status codes to indicate success: http.StatusOK
	ComputeErrorSummary func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.HypervSitesControllerClientComputeErrorSummaryOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientComputeErrorSummaryResponse], errResp azfake.ErrorResponder)

	// Computeusage is the fake for method HypervSitesControllerClient.Computeusage
	// HTTP status codes to indicate success: http.StatusOK
	Computeusage func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.HypervSitesControllerClientComputeusageOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientComputeusageResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method HypervSitesControllerClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.HypervSite, options *armmigrate.HypervSitesControllerClientCreateOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method HypervSitesControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.HypervSitesControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExportApplications is the fake for method HypervSitesControllerClient.BeginExportApplications
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportApplications func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.HypervSitesControllerClientBeginExportApplicationsOptions) (resp azfake.PollerResponder[armmigrate.HypervSitesControllerClientExportApplicationsResponse], errResp azfake.ErrorResponder)

	// BeginExportMachineErrors is the fake for method HypervSitesControllerClient.BeginExportMachineErrors
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportMachineErrors func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ExportMachineErrorsRequest, options *armmigrate.HypervSitesControllerClientBeginExportMachineErrorsOptions) (resp azfake.PollerResponder[armmigrate.HypervSitesControllerClientExportMachineErrorsResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method HypervSitesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.HypervSitesControllerClientGetOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// ListHealthSummary is the fake for method HypervSitesControllerClient.ListHealthSummary
	// HTTP status codes to indicate success: http.StatusOK
	ListHealthSummary func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.HypervSitesControllerClientListHealthSummaryOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientListHealthSummaryResponse], errResp azfake.ErrorResponder)

	// Summary is the fake for method HypervSitesControllerClient.Summary
	// HTTP status codes to indicate success: http.StatusOK
	Summary func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.HypervSitesControllerClientSummaryOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientSummaryResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method HypervSitesControllerClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.HypervSiteUpdate, options *armmigrate.HypervSitesControllerClientUpdateOptions) (resp azfake.Responder[armmigrate.HypervSitesControllerClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewHypervSitesControllerServerTransport creates a new instance of HypervSitesControllerServerTransport with the provided implementation.
// The returned HypervSitesControllerServerTransport instance is connected to an instance of armmigrate.HypervSitesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewHypervSitesControllerServerTransport(srv *HypervSitesControllerServer) *HypervSitesControllerServerTransport {
	return &HypervSitesControllerServerTransport{
		srv:                      srv,
		beginExportApplications:  newTracker[azfake.PollerResponder[armmigrate.HypervSitesControllerClientExportApplicationsResponse]](),
		beginExportMachineErrors: newTracker[azfake.PollerResponder[armmigrate.HypervSitesControllerClientExportMachineErrorsResponse]](),
	}
}

// HypervSitesControllerServerTransport connects instances of armmigrate.HypervSitesControllerClient to instances of HypervSitesControllerServer.
// Don't use this type directly, use NewHypervSitesControllerServerTransport instead.
type HypervSitesControllerServerTransport struct {
	srv                      *HypervSitesControllerServer
	beginExportApplications  *tracker[azfake.PollerResponder[armmigrate.HypervSitesControllerClientExportApplicationsResponse]]
	beginExportMachineErrors *tracker[azfake.PollerResponder[armmigrate.HypervSitesControllerClientExportMachineErrorsResponse]]
}

// Do implements the policy.Transporter interface for HypervSitesControllerServerTransport.
func (h *HypervSitesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return h.dispatchToMethodFake(req, method)
}

func (h *HypervSitesControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if hypervSitesControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = hypervSitesControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "HypervSitesControllerClient.ComputeErrorSummary":
				res.resp, res.err = h.dispatchComputeErrorSummary(req)
			case "HypervSitesControllerClient.Computeusage":
				res.resp, res.err = h.dispatchComputeusage(req)
			case "HypervSitesControllerClient.Create":
				res.resp, res.err = h.dispatchCreate(req)
			case "HypervSitesControllerClient.Delete":
				res.resp, res.err = h.dispatchDelete(req)
			case "HypervSitesControllerClient.BeginExportApplications":
				res.resp, res.err = h.dispatchBeginExportApplications(req)
			case "HypervSitesControllerClient.BeginExportMachineErrors":
				res.resp, res.err = h.dispatchBeginExportMachineErrors(req)
			case "HypervSitesControllerClient.Get":
				res.resp, res.err = h.dispatchGet(req)
			case "HypervSitesControllerClient.ListHealthSummary":
				res.resp, res.err = h.dispatchListHealthSummary(req)
			case "HypervSitesControllerClient.Summary":
				res.resp, res.err = h.dispatchSummary(req)
			case "HypervSitesControllerClient.Update":
				res.resp, res.err = h.dispatchUpdate(req)
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

func (h *HypervSitesControllerServerTransport) dispatchComputeErrorSummary(req *http.Request) (*http.Response, error) {
	if h.srv.ComputeErrorSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ComputeErrorSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/computeErrorSummary`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[any](req)
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
	respr, errRespr := h.srv.ComputeErrorSummary(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
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

func (h *HypervSitesControllerServerTransport) dispatchComputeusage(req *http.Request) (*http.Response, error) {
	if h.srv.Computeusage == nil {
		return nil, &nonRetriableError{errors.New("fake for method Computeusage not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/computeusage`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[any](req)
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
	respr, errRespr := h.srv.Computeusage(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervSiteUsage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if h.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.HypervSite](req)
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
	respr, errRespr := h.srv.Create(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervSite, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if h.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := h.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, nil)
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

func (h *HypervSitesControllerServerTransport) dispatchBeginExportApplications(req *http.Request) (*http.Response, error) {
	if h.srv.BeginExportApplications == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportApplications not implemented")}
	}
	beginExportApplications := h.beginExportApplications.get(req)
	if beginExportApplications == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportApplications`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[any](req)
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
		respr, errRespr := h.srv.BeginExportApplications(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportApplications = &respr
		h.beginExportApplications.add(req, beginExportApplications)
	}

	resp, err := server.PollerResponderNext(beginExportApplications, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginExportApplications.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportApplications) {
		h.beginExportApplications.remove(req)
	}

	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchBeginExportMachineErrors(req *http.Request) (*http.Response, error) {
	if h.srv.BeginExportMachineErrors == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportMachineErrors not implemented")}
	}
	beginExportMachineErrors := h.beginExportMachineErrors.get(req)
	if beginExportMachineErrors == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportMachineErrors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.ExportMachineErrorsRequest](req)
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
		respr, errRespr := h.srv.BeginExportMachineErrors(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportMachineErrors = &respr
		h.beginExportMachineErrors.add(req, beginExportMachineErrors)
	}

	resp, err := server.PollerResponderNext(beginExportMachineErrors, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		h.beginExportMachineErrors.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportMachineErrors) {
		h.beginExportMachineErrors.remove(req)
	}

	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if h.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := h.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchListHealthSummary(req *http.Request) (*http.Response, error) {
	if h.srv.ListHealthSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListHealthSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listHealthSummary`
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
	respr, errRespr := h.srv.ListHealthSummary(req.Context(), resourceGroupNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SiteHealthSummaryCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchSummary(req *http.Request) (*http.Response, error) {
	if h.srv.Summary == nil {
		return nil, &nonRetriableError{errors.New("fake for method Summary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/summary`
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
	respr, errRespr := h.srv.Summary(req.Context(), resourceGroupNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervSiteUsage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *HypervSitesControllerServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if h.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/hypervSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.HypervSiteUpdate](req)
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
	respr, errRespr := h.srv.Update(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).HypervSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to HypervSitesControllerServerTransport
var hypervSitesControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
