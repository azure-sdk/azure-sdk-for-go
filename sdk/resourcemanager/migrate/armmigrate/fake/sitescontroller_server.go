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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// SitesControllerServer is a fake server for instances of the armmigrate.SitesControllerClient type.
type SitesControllerServer struct {
	// ComputeErrorSummary is the fake for method SitesControllerClient.ComputeErrorSummary
	// HTTP status codes to indicate success: http.StatusOK
	ComputeErrorSummary func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.SitesControllerClientComputeErrorSummaryOptions) (resp azfake.Responder[armmigrate.SitesControllerClientComputeErrorSummaryResponse], errResp azfake.ErrorResponder)

	// Computeusage is the fake for method SitesControllerClient.Computeusage
	// HTTP status codes to indicate success: http.StatusOK
	Computeusage func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.SitesControllerClientComputeusageOptions) (resp azfake.Responder[armmigrate.SitesControllerClientComputeusageResponse], errResp azfake.ErrorResponder)

	// Create is the fake for method SitesControllerClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.VmwareSite, options *armmigrate.SitesControllerClientCreateOptions) (resp azfake.Responder[armmigrate.SitesControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method SitesControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.SitesControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.SitesControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExportApplications is the fake for method SitesControllerClient.BeginExportApplications
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportApplications func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.SitesControllerClientBeginExportApplicationsOptions) (resp azfake.PollerResponder[armmigrate.SitesControllerClientExportApplicationsResponse], errResp azfake.ErrorResponder)

	// BeginExportMachineErrors is the fake for method SitesControllerClient.BeginExportMachineErrors
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportMachineErrors func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ExportMachineErrorsRequest, options *armmigrate.SitesControllerClientBeginExportMachineErrorsOptions) (resp azfake.PollerResponder[armmigrate.SitesControllerClientExportMachineErrorsResponse], errResp azfake.ErrorResponder)

	// BeginExportMachines is the fake for method SitesControllerClient.BeginExportMachines
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportMachines func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ExportMachinesRequest, options *armmigrate.SitesControllerClientBeginExportMachinesOptions) (resp azfake.PollerResponder[armmigrate.SitesControllerClientExportMachinesResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SitesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.SitesControllerClientGetOptions) (resp azfake.Responder[armmigrate.SitesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method SitesControllerClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmigrate.SitesControllerClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmigrate.SitesControllerClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method SitesControllerClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmigrate.SitesControllerClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmigrate.SitesControllerClientListBySubscriptionResponse])

	// ListHealthSummary is the fake for method SitesControllerClient.ListHealthSummary
	// HTTP status codes to indicate success: http.StatusOK
	ListHealthSummary func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.SitesControllerClientListHealthSummaryOptions) (resp azfake.Responder[armmigrate.SitesControllerClientListHealthSummaryResponse], errResp azfake.ErrorResponder)

	// Summary is the fake for method SitesControllerClient.Summary
	// HTTP status codes to indicate success: http.StatusOK
	Summary func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.SitesControllerClientSummaryOptions) (resp azfake.Responder[armmigrate.SitesControllerClientSummaryResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method SitesControllerClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.VmwareSiteUpdate, options *armmigrate.SitesControllerClientUpdateOptions) (resp azfake.Responder[armmigrate.SitesControllerClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSitesControllerServerTransport creates a new instance of SitesControllerServerTransport with the provided implementation.
// The returned SitesControllerServerTransport instance is connected to an instance of armmigrate.SitesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSitesControllerServerTransport(srv *SitesControllerServer) *SitesControllerServerTransport {
	return &SitesControllerServerTransport{
		srv:                         srv,
		beginExportApplications:     newTracker[azfake.PollerResponder[armmigrate.SitesControllerClientExportApplicationsResponse]](),
		beginExportMachineErrors:    newTracker[azfake.PollerResponder[armmigrate.SitesControllerClientExportMachineErrorsResponse]](),
		beginExportMachines:         newTracker[azfake.PollerResponder[armmigrate.SitesControllerClientExportMachinesResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmigrate.SitesControllerClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmigrate.SitesControllerClientListBySubscriptionResponse]](),
	}
}

// SitesControllerServerTransport connects instances of armmigrate.SitesControllerClient to instances of SitesControllerServer.
// Don't use this type directly, use NewSitesControllerServerTransport instead.
type SitesControllerServerTransport struct {
	srv                         *SitesControllerServer
	beginExportApplications     *tracker[azfake.PollerResponder[armmigrate.SitesControllerClientExportApplicationsResponse]]
	beginExportMachineErrors    *tracker[azfake.PollerResponder[armmigrate.SitesControllerClientExportMachineErrorsResponse]]
	beginExportMachines         *tracker[azfake.PollerResponder[armmigrate.SitesControllerClientExportMachinesResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmigrate.SitesControllerClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmigrate.SitesControllerClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for SitesControllerServerTransport.
func (s *SitesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return s.dispatchToMethodFake(req, method)
}

func (s *SitesControllerServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if sitesControllerServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = sitesControllerServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "SitesControllerClient.ComputeErrorSummary":
				res.resp, res.err = s.dispatchComputeErrorSummary(req)
			case "SitesControllerClient.Computeusage":
				res.resp, res.err = s.dispatchComputeusage(req)
			case "SitesControllerClient.Create":
				res.resp, res.err = s.dispatchCreate(req)
			case "SitesControllerClient.Delete":
				res.resp, res.err = s.dispatchDelete(req)
			case "SitesControllerClient.BeginExportApplications":
				res.resp, res.err = s.dispatchBeginExportApplications(req)
			case "SitesControllerClient.BeginExportMachineErrors":
				res.resp, res.err = s.dispatchBeginExportMachineErrors(req)
			case "SitesControllerClient.BeginExportMachines":
				res.resp, res.err = s.dispatchBeginExportMachines(req)
			case "SitesControllerClient.Get":
				res.resp, res.err = s.dispatchGet(req)
			case "SitesControllerClient.NewListByResourceGroupPager":
				res.resp, res.err = s.dispatchNewListByResourceGroupPager(req)
			case "SitesControllerClient.NewListBySubscriptionPager":
				res.resp, res.err = s.dispatchNewListBySubscriptionPager(req)
			case "SitesControllerClient.ListHealthSummary":
				res.resp, res.err = s.dispatchListHealthSummary(req)
			case "SitesControllerClient.Summary":
				res.resp, res.err = s.dispatchSummary(req)
			case "SitesControllerClient.Update":
				res.resp, res.err = s.dispatchUpdate(req)
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

func (s *SitesControllerServerTransport) dispatchComputeErrorSummary(req *http.Request) (*http.Response, error) {
	if s.srv.ComputeErrorSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ComputeErrorSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/computeErrorSummary`
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
	respr, errRespr := s.srv.ComputeErrorSummary(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
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

func (s *SitesControllerServerTransport) dispatchComputeusage(req *http.Request) (*http.Response, error) {
	if s.srv.Computeusage == nil {
		return nil, &nonRetriableError{errors.New("fake for method Computeusage not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/computeusage`
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
	respr, errRespr := s.srv.Computeusage(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareSiteUsage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if s.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.VmwareSite](req)
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
	respr, errRespr := s.srv.Create(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareSite, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, nil)
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

func (s *SitesControllerServerTransport) dispatchBeginExportApplications(req *http.Request) (*http.Response, error) {
	if s.srv.BeginExportApplications == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportApplications not implemented")}
	}
	beginExportApplications := s.beginExportApplications.get(req)
	if beginExportApplications == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportApplications`
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
		respr, errRespr := s.srv.BeginExportApplications(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportApplications = &respr
		s.beginExportApplications.add(req, beginExportApplications)
	}

	resp, err := server.PollerResponderNext(beginExportApplications, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginExportApplications.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportApplications) {
		s.beginExportApplications.remove(req)
	}

	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchBeginExportMachineErrors(req *http.Request) (*http.Response, error) {
	if s.srv.BeginExportMachineErrors == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportMachineErrors not implemented")}
	}
	beginExportMachineErrors := s.beginExportMachineErrors.get(req)
	if beginExportMachineErrors == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportMachineErrors`
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
		respr, errRespr := s.srv.BeginExportMachineErrors(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportMachineErrors = &respr
		s.beginExportMachineErrors.add(req, beginExportMachineErrors)
	}

	resp, err := server.PollerResponderNext(beginExportMachineErrors, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginExportMachineErrors.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportMachineErrors) {
		s.beginExportMachineErrors.remove(req)
	}

	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchBeginExportMachines(req *http.Request) (*http.Response, error) {
	if s.srv.BeginExportMachines == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportMachines not implemented")}
	}
	beginExportMachines := s.beginExportMachines.get(req)
	if beginExportMachines == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportMachines`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.ExportMachinesRequest](req)
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
		respr, errRespr := s.srv.BeginExportMachines(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExportMachines = &respr
		s.beginExportMachines.add(req, beginExportMachines)
	}

	resp, err := server.PollerResponderNext(beginExportMachines, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginExportMachines.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExportMachines) {
		s.beginExportMachines.remove(req)
	}

	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		s.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmigrate.SitesControllerClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		s.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := s.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		s.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmigrate.SitesControllerClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		s.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchListHealthSummary(req *http.Request) (*http.Response, error) {
	if s.srv.ListHealthSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListHealthSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listHealthSummary`
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
	respr, errRespr := s.srv.ListHealthSummary(req.Context(), resourceGroupNameParam, siteNameParam, nil)
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

func (s *SitesControllerServerTransport) dispatchSummary(req *http.Request) (*http.Response, error) {
	if s.srv.Summary == nil {
		return nil, &nonRetriableError{errors.New("fake for method Summary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/summary`
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
	respr, errRespr := s.srv.Summary(req.Context(), resourceGroupNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareSiteUsage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SitesControllerServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/vmwareSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.VmwareSiteUpdate](req)
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
	respr, errRespr := s.srv.Update(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VmwareSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to SitesControllerServerTransport
var sitesControllerServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
