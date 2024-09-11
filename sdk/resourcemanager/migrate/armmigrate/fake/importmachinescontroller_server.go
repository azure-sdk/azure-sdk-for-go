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

// ImportMachinesControllerServer is a fake server for instances of the armmigrate.ImportMachinesControllerClient type.
type ImportMachinesControllerServer struct {
	// Delete is the fake for method ImportMachinesControllerClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *armmigrate.ImportMachinesControllerClientDeleteOptions) (resp azfake.Responder[armmigrate.ImportMachinesControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ImportMachinesControllerClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, siteName string, machineName string, options *armmigrate.ImportMachinesControllerClientGetOptions) (resp azfake.Responder[armmigrate.ImportMachinesControllerClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByImportSitePager is the fake for method ImportMachinesControllerClient.NewListByImportSitePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByImportSitePager func(resourceGroupName string, siteName string, options *armmigrate.ImportMachinesControllerClientListByImportSiteOptions) (resp azfake.PagerResponder[armmigrate.ImportMachinesControllerClientListByImportSiteResponse])
}

// NewImportMachinesControllerServerTransport creates a new instance of ImportMachinesControllerServerTransport with the provided implementation.
// The returned ImportMachinesControllerServerTransport instance is connected to an instance of armmigrate.ImportMachinesControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewImportMachinesControllerServerTransport(srv *ImportMachinesControllerServer) *ImportMachinesControllerServerTransport {
	return &ImportMachinesControllerServerTransport{
		srv:                      srv,
		newListByImportSitePager: newTracker[azfake.PagerResponder[armmigrate.ImportMachinesControllerClientListByImportSiteResponse]](),
	}
}

// ImportMachinesControllerServerTransport connects instances of armmigrate.ImportMachinesControllerClient to instances of ImportMachinesControllerServer.
// Don't use this type directly, use NewImportMachinesControllerServerTransport instead.
type ImportMachinesControllerServerTransport struct {
	srv                      *ImportMachinesControllerServer
	newListByImportSitePager *tracker[azfake.PagerResponder[armmigrate.ImportMachinesControllerClientListByImportSiteResponse]]
}

// Do implements the policy.Transporter interface for ImportMachinesControllerServerTransport.
func (i *ImportMachinesControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ImportMachinesControllerClient.Delete":
		resp, err = i.dispatchDelete(req)
	case "ImportMachinesControllerClient.Get":
		resp, err = i.dispatchGet(req)
	case "ImportMachinesControllerClient.NewListByImportSitePager":
		resp, err = i.dispatchNewListByImportSitePager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *ImportMachinesControllerServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if i.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Delete(req.Context(), resourceGroupNameParam, siteNameParam, machineNameParam, nil)
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

func (i *ImportMachinesControllerServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines/(?P<machineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	machineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("machineName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameParam, siteNameParam, machineNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImportMachine, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *ImportMachinesControllerServerTransport) dispatchNewListByImportSitePager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByImportSitePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByImportSitePager not implemented")}
	}
	newListByImportSitePager := i.newListByImportSitePager.get(req)
	if newListByImportSitePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.OffAzure/importSites/(?P<siteName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/machines`
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
		filterUnescaped, err := url.QueryUnescape(qp.Get("filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		topUnescaped, err := url.QueryUnescape(qp.Get("top"))
		if err != nil {
			return nil, err
		}
		topParam := getOptional(topUnescaped)
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		totalRecordCountUnescaped, err := url.QueryUnescape(qp.Get("totalRecordCount"))
		if err != nil {
			return nil, err
		}
		totalRecordCountParam, err := parseOptional(totalRecordCountUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		siteNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("siteName")])
		if err != nil {
			return nil, err
		}
		var options *armmigrate.ImportMachinesControllerClientListByImportSiteOptions
		if filterParam != nil || topParam != nil || continuationTokenParam != nil || totalRecordCountParam != nil {
			options = &armmigrate.ImportMachinesControllerClientListByImportSiteOptions{
				Filter:            filterParam,
				Top:               topParam,
				ContinuationToken: continuationTokenParam,
				TotalRecordCount:  totalRecordCountParam,
			}
		}
		resp := i.srv.NewListByImportSitePager(resourceGroupNameParam, siteNameParam, options)
		newListByImportSitePager = &resp
		i.newListByImportSitePager.add(req, newListByImportSitePager)
		server.PagerResponderInjectNextLinks(newListByImportSitePager, req, func(page *armmigrate.ImportMachinesControllerClientListByImportSiteResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByImportSitePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByImportSitePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByImportSitePager) {
		i.newListByImportSitePager.remove(req)
	}
	return resp, nil
}