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
	"strconv"
)

// ImportSitesControllerServer is a fake server for instances of the armmigrate.ImportSitesControllerClient type.
type ImportSitesControllerServer struct {
	// Create is the fake for method ImportSitesControllerClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ImportSite, options *armmigrate.ImportSitesControllerClientCreateOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ImportSitesControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ImportSitesControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// DeleteImportedMachines is the fake for method ImportSitesControllerClient.DeleteImportedMachines
	// HTTP status codes to indicate success: http.StatusOK
	DeleteImportedMachines func(ctx context.Context, resourceGroupName string, siteName string, body any, options *armmigrate.ImportSitesControllerClientDeleteImportedMachinesOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientDeleteImportedMachinesResponse], errResp azfake.ErrorResponder)

	// ExportURI is the fake for method ImportSitesControllerClient.ExportURI
	// HTTP status codes to indicate success: http.StatusOK
	ExportURI func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.SasURIResponse, options *armmigrate.ImportSitesControllerClientExportURIOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientExportURIResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ImportSitesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, options *armmigrate.ImportSitesControllerClientGetOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// ImportURI is the fake for method ImportSitesControllerClient.ImportURI
	// HTTP status codes to indicate success: http.StatusOK
	ImportURI func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.SasURIResponse, options *armmigrate.ImportSitesControllerClientImportURIOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientImportURIResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method ImportSitesControllerClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armmigrate.ImportSitesControllerClientListByResourceGroupOptions) (resp azfake.PagerResponder[armmigrate.ImportSitesControllerClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method ImportSitesControllerClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armmigrate.ImportSitesControllerClientListBySubscriptionOptions) (resp azfake.PagerResponder[armmigrate.ImportSitesControllerClientListBySubscriptionResponse])

	// Update is the fake for method ImportSitesControllerClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, siteName string, body armmigrate.ImportSiteUpdate, options *armmigrate.ImportSitesControllerClientUpdateOptions) (resp azfake.Responder[armmigrate.ImportSitesControllerClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewImportSitesControllerServerTransport creates a new instance of ImportSitesControllerServerTransport with the provided implementation.
// The returned ImportSitesControllerServerTransport instance is connected to an instance of armmigrate.ImportSitesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImportSitesControllerServerTransport(srv *ImportSitesControllerServer) *ImportSitesControllerServerTransport {
	return &ImportSitesControllerServerTransport{
		srv:                         srv,
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armmigrate.ImportSitesControllerClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armmigrate.ImportSitesControllerClientListBySubscriptionResponse]](),
	}
}

// ImportSitesControllerServerTransport connects instances of armmigrate.ImportSitesControllerClient to instances of ImportSitesControllerServer.
// Don't use this type directly, use NewImportSitesControllerServerTransport instead.
type ImportSitesControllerServerTransport struct {
	srv                         *ImportSitesControllerServer
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armmigrate.ImportSitesControllerClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armmigrate.ImportSitesControllerClientListBySubscriptionResponse]]
}

// Do implements the policy.Transporter interface for ImportSitesControllerServerTransport.
func (i *ImportSitesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ImportSitesControllerClient.Create":
		resp, err = i.dispatchCreate(req)
	case "ImportSitesControllerClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "ImportSitesControllerClient.DeleteImportedMachines":
		resp, err = i.dispatchDeleteImportedMachines(req)
	case "ImportSitesControllerClient.ExportURI":
		resp, err = i.dispatchExportURI(req)
	case "ImportSitesControllerClient.Get":
		resp, err = i.dispatchGet(req)
	case "ImportSitesControllerClient.ImportURI":
		resp, err = i.dispatchImportURI(req)
	case "ImportSitesControllerClient.NewListByResourceGroupPager":
		resp, err = i.dispatchNewListByResourceGroupPager(req)
	case "ImportSitesControllerClient.NewListBySubscriptionPager":
		resp, err = i.dispatchNewListBySubscriptionPager(req)
	case "ImportSitesControllerClient.Update":
		resp, err = i.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if i.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.ImportSite](req)
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
	respr, errRespr := i.srv.Create(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImportSite, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).RetryAfter; val != nil {
		resp.Header.Set("Retry-After", strconv.FormatInt(int64(*val), 10))
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, nil)
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

func (i *ImportSitesControllerServerTransport) dispatchDeleteImportedMachines(req *http.Request) (*http.Response, error) {
	if i.srv.DeleteImportedMachines == nil {
		return nil, &nonRetriableError{errors.New("fake for method DeleteImportedMachines not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deleteImportedMachines`
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
	respr, errRespr := i.srv.DeleteImportedMachines(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SasURIResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchExportURI(req *http.Request) (*http.Response, error) {
	if i.srv.ExportURI == nil {
		return nil, &nonRetriableError{errors.New("fake for method ExportURI not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/exportUri`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.SasURIResponse](req)
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
	respr, errRespr := i.srv.ExportURI(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SasURIResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImportSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchImportURI(req *http.Request) (*http.Response, error) {
	if i.srv.ImportURI == nil {
		return nil, &nonRetriableError{errors.New("fake for method ImportURI not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/importUri`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.SasURIResponse](req)
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
	respr, errRespr := i.srv.ImportURI(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SasURIResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := i.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		i.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armmigrate.ImportSitesControllerClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		i.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := i.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListBySubscriptionPager(nil)
		newListBySubscriptionPager = &resp
		i.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armmigrate.ImportSitesControllerClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		i.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (i *ImportSitesControllerServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armmigrate.ImportSiteUpdate](req)
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
	respr, errRespr := i.srv.Update(req.Context(), resourceGroupNameParam, siteNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImportSite, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
