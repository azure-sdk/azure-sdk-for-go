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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v7"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// VirtualAppliancesServer is a fake server for instances of the armnetwork.VirtualAppliancesClient type.
type VirtualAppliancesServer struct {
	// BeginCreateOrUpdate is the fake for method VirtualAppliancesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters armnetwork.VirtualAppliance, options *armnetwork.VirtualAppliancesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.VirtualAppliancesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualAppliancesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *armnetwork.VirtualAppliancesClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.VirtualAppliancesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualAppliancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *armnetwork.VirtualAppliancesClientGetOptions) (resp azfake.Responder[armnetwork.VirtualAppliancesClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGetBootDiagnosticLogs is the fake for method VirtualAppliancesClient.BeginGetBootDiagnosticLogs
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGetBootDiagnosticLogs func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, request armnetwork.VirtualApplianceBootDiagnosticParameters, options *armnetwork.VirtualAppliancesClientBeginGetBootDiagnosticLogsOptions) (resp azfake.PollerResponder[armnetwork.VirtualAppliancesClientGetBootDiagnosticLogsResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualAppliancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.VirtualAppliancesClientListOptions) (resp azfake.PagerResponder[armnetwork.VirtualAppliancesClientListResponse])

	// NewListByResourceGroupPager is the fake for method VirtualAppliancesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.VirtualAppliancesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.VirtualAppliancesClientListByResourceGroupResponse])

	// BeginReimage is the fake for method VirtualAppliancesClient.BeginReimage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginReimage func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *armnetwork.VirtualAppliancesClientBeginReimageOptions) (resp azfake.PollerResponder[armnetwork.VirtualAppliancesClientReimageResponse], errResp azfake.ErrorResponder)

	// BeginRestart is the fake for method VirtualAppliancesClient.BeginRestart
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRestart func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, options *armnetwork.VirtualAppliancesClientBeginRestartOptions) (resp azfake.PollerResponder[armnetwork.VirtualAppliancesClientRestartResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method VirtualAppliancesClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, networkVirtualApplianceName string, parameters armnetwork.TagsObject, options *armnetwork.VirtualAppliancesClientUpdateTagsOptions) (resp azfake.Responder[armnetwork.VirtualAppliancesClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualAppliancesServerTransport creates a new instance of VirtualAppliancesServerTransport with the provided implementation.
// The returned VirtualAppliancesServerTransport instance is connected to an instance of armnetwork.VirtualAppliancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualAppliancesServerTransport(srv *VirtualAppliancesServer) *VirtualAppliancesServerTransport {
	return &VirtualAppliancesServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientDeleteResponse]](),
		beginGetBootDiagnosticLogs:  newTracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientGetBootDiagnosticLogsResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armnetwork.VirtualAppliancesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetwork.VirtualAppliancesClientListByResourceGroupResponse]](),
		beginReimage:                newTracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientReimageResponse]](),
		beginRestart:                newTracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientRestartResponse]](),
	}
}

// VirtualAppliancesServerTransport connects instances of armnetwork.VirtualAppliancesClient to instances of VirtualAppliancesServer.
// Don't use this type directly, use NewVirtualAppliancesServerTransport instead.
type VirtualAppliancesServerTransport struct {
	srv                         *VirtualAppliancesServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientDeleteResponse]]
	beginGetBootDiagnosticLogs  *tracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientGetBootDiagnosticLogsResponse]]
	newListPager                *tracker[azfake.PagerResponder[armnetwork.VirtualAppliancesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetwork.VirtualAppliancesClientListByResourceGroupResponse]]
	beginReimage                *tracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientReimageResponse]]
	beginRestart                *tracker[azfake.PollerResponder[armnetwork.VirtualAppliancesClientRestartResponse]]
}

// Do implements the policy.Transporter interface for VirtualAppliancesServerTransport.
func (v *VirtualAppliancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualAppliancesClient.BeginCreateOrUpdate":
		resp, err = v.dispatchBeginCreateOrUpdate(req)
	case "VirtualAppliancesClient.BeginDelete":
		resp, err = v.dispatchBeginDelete(req)
	case "VirtualAppliancesClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualAppliancesClient.BeginGetBootDiagnosticLogs":
		resp, err = v.dispatchBeginGetBootDiagnosticLogs(req)
	case "VirtualAppliancesClient.NewListPager":
		resp, err = v.dispatchNewListPager(req)
	case "VirtualAppliancesClient.NewListByResourceGroupPager":
		resp, err = v.dispatchNewListByResourceGroupPager(req)
	case "VirtualAppliancesClient.BeginReimage":
		resp, err = v.dispatchBeginReimage(req)
	case "VirtualAppliancesClient.BeginRestart":
		resp, err = v.dispatchBeginRestart(req)
	case "VirtualAppliancesClient.UpdateTags":
		resp, err = v.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualAppliance](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		v.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		v.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		v.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		v.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		v.beginDelete.remove(req)
	}

	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.VirtualAppliancesClientGetOptions
	if expandParam != nil {
		options = &armnetwork.VirtualAppliancesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualAppliance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchBeginGetBootDiagnosticLogs(req *http.Request) (*http.Response, error) {
	if v.srv.BeginGetBootDiagnosticLogs == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGetBootDiagnosticLogs not implemented")}
	}
	beginGetBootDiagnosticLogs := v.beginGetBootDiagnosticLogs.get(req)
	if beginGetBootDiagnosticLogs == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/getBootDiagnosticLogs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualApplianceBootDiagnosticParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginGetBootDiagnosticLogs(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGetBootDiagnosticLogs = &respr
		v.beginGetBootDiagnosticLogs.add(req, beginGetBootDiagnosticLogs)
	}

	resp, err := server.PollerResponderNext(beginGetBootDiagnosticLogs, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginGetBootDiagnosticLogs.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGetBootDiagnosticLogs) {
		v.beginGetBootDiagnosticLogs.remove(req)
	}

	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.VirtualAppliancesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		v.newListPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := v.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		v.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetwork.VirtualAppliancesClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		v.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchBeginReimage(req *http.Request) (*http.Response, error) {
	if v.srv.BeginReimage == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginReimage not implemented")}
	}
	beginReimage := v.beginReimage.get(req)
	if beginReimage == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/reimage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualApplianceInstanceIDs](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VirtualAppliancesClientBeginReimageOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VirtualAppliancesClientBeginReimageOptions{
				NetworkVirtualApplianceInstanceIDs: &body,
			}
		}
		respr, errRespr := v.srv.BeginReimage(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginReimage = &respr
		v.beginReimage.add(req, beginReimage)
	}

	resp, err := server.PollerResponderNext(beginReimage, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginReimage.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginReimage) {
		v.beginReimage.remove(req)
	}

	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchBeginRestart(req *http.Request) (*http.Response, error) {
	if v.srv.BeginRestart == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestart not implemented")}
	}
	beginRestart := v.beginRestart.get(req)
	if beginRestart == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restart`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.VirtualApplianceInstanceIDs](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
		if err != nil {
			return nil, err
		}
		var options *armnetwork.VirtualAppliancesClientBeginRestartOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armnetwork.VirtualAppliancesClientBeginRestartOptions{
				NetworkVirtualApplianceInstanceIDs: &body,
			}
		}
		respr, errRespr := v.srv.BeginRestart(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestart = &respr
		v.beginRestart.add(req, beginRestart)
	}

	resp, err := server.PollerResponderNext(beginRestart, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginRestart.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestart) {
		v.beginRestart.remove(req)
	}

	return resp, nil
}

func (v *VirtualAppliancesServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if v.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/networkVirtualAppliances/(?P<networkVirtualApplianceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkVirtualApplianceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("networkVirtualApplianceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.UpdateTags(req.Context(), resourceGroupNameParam, networkVirtualApplianceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualAppliance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
