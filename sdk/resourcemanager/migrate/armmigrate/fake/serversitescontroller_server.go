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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ServerSitesControllerServer is a fake server for instances of the armmigrate.ServerSitesControllerClient type.
type ServerSitesControllerServer struct {
	// ComputeErrorSummary is the fake for method ServerSitesControllerClient.ComputeErrorSummary
	// HTTP status codes to indicate success: http.StatusOK
	ComputeErrorSummary func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.ServerSitesControllerClientComputeErrorSummaryOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientComputeErrorSummaryResponse], errResp azfake.ErrorResponder)

	// Computeusage is the fake for method ServerSitesControllerClient.Computeusage
	// HTTP status codes to indicate success: http.StatusOK
	Computeusage func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.ServerSitesControllerClientComputeusageOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientComputeusageResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method ServerSitesControllerClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ServerSiteResource, options *armmigrate.ServerSitesControllerClientBeginCreateOptions) (resp azfake.PollerResponder[armmigrate.ServerSitesControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ServerSitesControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ServerSitesControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExportApplications is the fake for method ServerSitesControllerClient.BeginExportApplications
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportApplications func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.ServerSitesControllerClientBeginExportApplicationsOptions) (resp azfake.PollerResponder[armmigrate.ServerSitesControllerClientExportApplicationsResponse], errResp azfake.ErrorResponder)

	// BeginExportMachineErrors is the fake for method ServerSitesControllerClient.BeginExportMachineErrors
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExportMachineErrors func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ExportMachineErrorsRequest, options *armmigrate.ServerSitesControllerClientBeginExportMachineErrorsOptions) (resp azfake.PollerResponder[armmigrate.ServerSitesControllerClientExportMachineErrorsResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ServerSitesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ServerSitesControllerClientGetOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ServerSitesControllerClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmigrate.ServerSitesControllerClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmigrate.ServerSitesControllerClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ServerSitesControllerClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmigrate.ServerSitesControllerClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmigrate.ServerSitesControllerClientListBySubscriptionResponse])

	// ListHealthSummary is the fake for method ServerSitesControllerClient.ListHealthSummary
	// HTTP status codes to indicate success: http.StatusOK
	ListHealthSummary func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ServerSitesControllerClientListHealthSummaryOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientListHealthSummaryResponse], errResp azfake.ErrorResponder)

	// BeginRefreshSite is the fake for method ServerSitesControllerClient.BeginRefreshSite
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRefreshSite func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ServerSitesControllerClientBeginRefreshSiteOptions) (resp azfake.PollerResponder[armmigrate.ServerSitesControllerClientRefreshSiteResponse], errResp azfake.ErrorResponder)

	// Summary is the fake for method ServerSitesControllerClient.Summary
	// HTTP status codes to indicate success: http.StatusOK
	Summary func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ServerSitesControllerClientSummaryOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientSummaryResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method ServerSitesControllerClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ServerSiteResourceUpdate, options *armmigrate.ServerSitesControllerClientUpdateOptions) (resp azfake.Responder[armmigrate.ServerSitesControllerClientUpdateResponse], errResp azfake.ErrorResponder)

	// BeginUpdateDependencyMapStatus is the fake for method ServerSitesControllerClient.BeginUpdateDependencyMapStatus
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateDependencyMapStatus func(ctx context.Context, resourceGroupName string, siteName string, mapRequest armmigrate.UpdateMachineDepMapStatus, options *armmigrate.ServerSitesControllerClientBeginUpdateDependencyMapStatusOptions) (resp azfake.PollerResponder[armmigrate.ServerSitesControllerClientUpdateDependencyMapStatusResponse], errResp azfake.ErrorResponder)

	// BeginUpdateProperties is the fake for method ServerSitesControllerClient.BeginUpdateProperties
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdateProperties func(ctx context.Context, resourceGroupName string, siteName string, metaData armmigrate.MachineMetadataCollection, options *armmigrate.ServerSitesControllerClientBeginUpdatePropertiesOptions) (resp azfake.PollerResponder[armmigrate.ServerSitesControllerClientUpdatePropertiesResponse], errResp azfake.ErrorResponder)
}

// NewServerSitesControllerServerTransport creates a new instance of ServerSitesControllerServerTransport with the provided implementation.
// The returned ServerSitesControllerServerTransport instance is connected to an instance of armmigrate.ServerSitesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerSitesControllerServerTransport(srv *ServerSitesControllerServer) *ServerSitesControllerServerTransport {
	return &ServerSitesControllerServerTransport{
		srv:                            srv,
		beginCreate:                    newTracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientCreateResponse]](),
		beginExportApplications:        newTracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientExportApplicationsResponse]](),
		beginExportMachineErrors:       newTracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientExportMachineErrorsResponse]](),
		newListByResourceGroupPager:    newTracker[azfake.PagerResponder[armmigrate.ServerSitesControllerClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:     newTracker[azfake.PagerResponder[armmigrate.ServerSitesControllerClientListBySubscriptionResponse]](),
		beginRefreshSite:               newTracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientRefreshSiteResponse]](),
		beginUpdateDependencyMapStatus: newTracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientUpdateDependencyMapStatusResponse]](),
		beginUpdateProperties:          newTracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientUpdatePropertiesResponse]](),
	}
}

// ServerSitesControllerServerTransport connects instances of armmigrate.ServerSitesControllerClient to instances of ServerSitesControllerServer.
// Don't use this type directly, use NewServerSitesControllerServerTransport instead.
type ServerSitesControllerServerTransport struct {
	srv                            *ServerSitesControllerServer
	beginCreate                    *tracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientCreateResponse]]
	beginExportApplications        *tracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientExportApplicationsResponse]]
	beginExportMachineErrors       *tracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientExportMachineErrorsResponse]]
	newListByResourceGroupPager    *tracker[azfake.PagerResponder[armmigrate.ServerSitesControllerClientListByResourceGroupResponse]]
	newListBySubscriptionPager     *tracker[azfake.PagerResponder[armmigrate.ServerSitesControllerClientListBySubscriptionResponse]]
	beginRefreshSite               *tracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientRefreshSiteResponse]]
	beginUpdateDependencyMapStatus *tracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientUpdateDependencyMapStatusResponse]]
	beginUpdateProperties          *tracker[azfake.PollerResponder[armmigrate.ServerSitesControllerClientUpdatePropertiesResponse]]
}

// Do implements the policy.Transporter interface for ServerSitesControllerServerTransport.
func (s *ServerSitesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ServerSitesControllerClient.ComputeErrorSummary":
		resp, err = s.dispatchComputeErrorSummary(req)
	case "ServerSitesControllerClient.Computeusage":
		resp, err = s.dispatchComputeusage(req)
	case "ServerSitesControllerClient.BeginCreate":
		resp, err = s.dispatchBeginCreate(req)
	case "ServerSitesControllerClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "ServerSitesControllerClient.BeginExportApplications":
		resp, err = s.dispatchBeginExportApplications(req)
	case "ServerSitesControllerClient.BeginExportMachineErrors":
		resp, err = s.dispatchBeginExportMachineErrors(req)
	case "ServerSitesControllerClient.Get":
		resp, err = s.dispatchGet(req)
	case "ServerSitesControllerClient.NewListByResourceGroupPager":
		resp, err = s.dispatchNewListByResourceGroupPager(req)
	case "ServerSitesControllerClient.NewListBySubscriptionPager":
		resp, err = s.dispatchNewListBySubscriptionPager(req)
	case "ServerSitesControllerClient.ListHealthSummary":
		resp, err = s.dispatchListHealthSummary(req)
	case "ServerSitesControllerClient.BeginRefreshSite":
		resp, err = s.dispatchBeginRefreshSite(req)
	case "ServerSitesControllerClient.Summary":
		resp, err = s.dispatchSummary(req)
	case "ServerSitesControllerClient.Update":
		resp, err = s.dispatchUpdate(req)
	case "ServerSitesControllerClient.BeginUpdateDependencyMapStatus":
		resp, err = s.dispatchBeginUpdateDependencyMapStatus(req)
	case "ServerSitesControllerClient.BeginUpdateProperties":
		resp, err = s.dispatchBeginUpdateProperties(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchComputeErrorSummary(req *http.Request) (*http.Response, error) {
	if s.srv.ComputeErrorSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ComputeErrorSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/computeErrorSummary`
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

func (s *ServerSitesControllerServerTransport) dispatchComputeusage(req *http.Request) (*http.Response, error) {
	if s.srv.Computeusage == nil {
		return nil, &nonRetriableError{errors.New("fake for method Computeusage not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/computeusage`
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSiteUsageResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.ServerSiteResource](req)
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
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		s.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		s.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		s.beginCreate.remove(req)
	}

	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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

func (s *ServerSitesControllerServerTransport) dispatchBeginExportApplications(req *http.Request) (*http.Response, error) {
	if s.srv.BeginExportApplications == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportApplications not implemented")}
	}
	beginExportApplications := s.beginExportApplications.get(req)
	if beginExportApplications == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportApplications`
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

func (s *ServerSitesControllerServerTransport) dispatchBeginExportMachineErrors(req *http.Request) (*http.Response, error) {
	if s.srv.BeginExportMachineErrors == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExportMachineErrors not implemented")}
	}
	beginExportMachineErrors := s.beginExportMachineErrors.get(req)
	if beginExportMachineErrors == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportMachineErrors`
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

func (s *ServerSitesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSiteResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := s.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites`
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
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmigrate.ServerSitesControllerClientListByResourceGroupResponse, createLink func() string) {
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

func (s *ServerSitesControllerServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := s.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := s.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		s.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmigrate.ServerSitesControllerClientListBySubscriptionResponse, createLink func() string) {
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

func (s *ServerSitesControllerServerTransport) dispatchListHealthSummary(req *http.Request) (*http.Response, error) {
	if s.srv.ListHealthSummary == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListHealthSummary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/listHealthSummary`
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

func (s *ServerSitesControllerServerTransport) dispatchBeginRefreshSite(req *http.Request) (*http.Response, error) {
	if s.srv.BeginRefreshSite == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRefreshSite not implemented")}
	}
	beginRefreshSite := s.beginRefreshSite.get(req)
	if beginRefreshSite == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/refreshSite`
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
		respr, errRespr := s.srv.BeginRefreshSite(req.Context(), resourceGroupNameParam, siteNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRefreshSite = &respr
		s.beginRefreshSite.add(req, beginRefreshSite)
	}

	resp, err := server.PollerResponderNext(beginRefreshSite, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginRefreshSite.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRefreshSite) {
		s.beginRefreshSite.remove(req)
	}

	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchSummary(req *http.Request) (*http.Response, error) {
	if s.srv.Summary == nil {
		return nil, &nonRetriableError{errors.New("fake for method Summary not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/summary`
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSiteUsage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.ServerSiteResourceUpdate](req)
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
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ServerSiteResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchBeginUpdateDependencyMapStatus(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdateDependencyMapStatus == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateDependencyMapStatus not implemented")}
	}
	beginUpdateDependencyMapStatus := s.beginUpdateDependencyMapStatus.get(req)
	if beginUpdateDependencyMapStatus == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateDependencyMapStatus`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.UpdateMachineDepMapStatus](req)
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
		respr, errRespr := s.srv.BeginUpdateDependencyMapStatus(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateDependencyMapStatus = &respr
		s.beginUpdateDependencyMapStatus.add(req, beginUpdateDependencyMapStatus)
	}

	resp, err := server.PollerResponderNext(beginUpdateDependencyMapStatus, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdateDependencyMapStatus.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateDependencyMapStatus) {
		s.beginUpdateDependencyMapStatus.remove(req)
	}

	return resp, nil
}

func (s *ServerSitesControllerServerTransport) dispatchBeginUpdateProperties(req *http.Request) (*http.Response, error) {
	if s.srv.BeginUpdateProperties == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdateProperties not implemented")}
	}
	beginUpdateProperties := s.beginUpdateProperties.get(req)
	if beginUpdateProperties == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/serverSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/updateProperties`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armmigrate.MachineMetadataCollection](req)
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
		respr, errRespr := s.srv.BeginUpdateProperties(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdateProperties = &respr
		s.beginUpdateProperties.add(req, beginUpdateProperties)
	}

	resp, err := server.PollerResponderNext(beginUpdateProperties, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginUpdateProperties.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdateProperties) {
		s.beginUpdateProperties.remove(req)
	}

	return resp, nil
}