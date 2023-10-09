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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"net/http"
	"net/url"
	"regexp"
)

// DisksServer is a fake server for instances of the armcompute.DisksClient type.
type DisksServer struct {
	// BeginCreateOrUpdate is the fake for method DisksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, diskName string, disk armcompute.Disk, options *armcompute.DisksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armcompute.DisksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DisksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, diskName string, options *armcompute.DisksClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.DisksClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DisksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, diskName string, options *armcompute.DisksClientGetOptions) (resp azfake.Responder[armcompute.DisksClientGetResponse], errResp azfake.ErrorResponder)

	// BeginGrantAccess is the fake for method DisksClient.BeginGrantAccess
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginGrantAccess func(ctx context.Context, resourceGroupName string, diskName string, grantAccessData armcompute.GrantAccessData, options *armcompute.DisksClientBeginGrantAccessOptions) (resp azfake.PollerResponder[armcompute.DisksClientGrantAccessResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DisksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armcompute.DisksClientListOptions) (resp azfake.PagerResponder[armcompute.DisksClientListResponse])

	// NewListByResourceGroupPager is the fake for method DisksClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.DisksClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.DisksClientListByResourceGroupResponse])

	// BeginRevokeAccess is the fake for method DisksClient.BeginRevokeAccess
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginRevokeAccess func(ctx context.Context, resourceGroupName string, diskName string, options *armcompute.DisksClientBeginRevokeAccessOptions) (resp azfake.PollerResponder[armcompute.DisksClientRevokeAccessResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method DisksClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, diskName string, disk armcompute.DiskUpdate, options *armcompute.DisksClientBeginUpdateOptions) (resp azfake.PollerResponder[armcompute.DisksClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDisksServerTransport creates a new instance of DisksServerTransport with the provided implementation.
// The returned DisksServerTransport instance is connected to an instance of armcompute.DisksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDisksServerTransport(srv *DisksServer) *DisksServerTransport {
	return &DisksServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armcompute.DisksClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armcompute.DisksClientDeleteResponse]](),
		beginGrantAccess:            newTracker[azfake.PollerResponder[armcompute.DisksClientGrantAccessResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armcompute.DisksClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armcompute.DisksClientListByResourceGroupResponse]](),
		beginRevokeAccess:           newTracker[azfake.PollerResponder[armcompute.DisksClientRevokeAccessResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armcompute.DisksClientUpdateResponse]](),
	}
}

// DisksServerTransport connects instances of armcompute.DisksClient to instances of DisksServer.
// Don't use this type directly, use NewDisksServerTransport instead.
type DisksServerTransport struct {
	srv                         *DisksServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armcompute.DisksClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armcompute.DisksClientDeleteResponse]]
	beginGrantAccess            *tracker[azfake.PollerResponder[armcompute.DisksClientGrantAccessResponse]]
	newListPager                *tracker[azfake.PagerResponder[armcompute.DisksClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armcompute.DisksClientListByResourceGroupResponse]]
	beginRevokeAccess           *tracker[azfake.PollerResponder[armcompute.DisksClientRevokeAccessResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armcompute.DisksClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for DisksServerTransport.
func (d *DisksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DisksClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DisksClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DisksClient.Get":
		resp, err = d.dispatchGet(req)
	case "DisksClient.BeginGrantAccess":
		resp, err = d.dispatchBeginGrantAccess(req)
	case "DisksClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DisksClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DisksClient.BeginRevokeAccess":
		resp, err = d.dispatchBeginRevokeAccess(req)
	case "DisksClient.BeginUpdate":
		resp, err = d.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.Disk](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, diskNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, diskNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	diskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameUnescaped, diskNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Disk, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginGrantAccess(req *http.Request) (*http.Response, error) {
	if d.srv.BeginGrantAccess == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginGrantAccess not implemented")}
	}
	beginGrantAccess := d.beginGrantAccess.get(req)
	if beginGrantAccess == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/beginGetAccess`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.GrantAccessData](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginGrantAccess(req.Context(), resourceGroupNameUnescaped, diskNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginGrantAccess = &respr
		d.beginGrantAccess.add(req, beginGrantAccess)
	}

	resp, err := server.PollerResponderNext(beginGrantAccess, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginGrantAccess.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginGrantAccess) {
		d.beginGrantAccess.remove(req)
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListPager(nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armcompute.DisksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}

func (d *DisksServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := d.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		d.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armcompute.DisksClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		d.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginRevokeAccess(req *http.Request) (*http.Response, error) {
	if d.srv.BeginRevokeAccess == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRevokeAccess not implemented")}
	}
	beginRevokeAccess := d.beginRevokeAccess.get(req)
	if beginRevokeAccess == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/endGetAccess`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginRevokeAccess(req.Context(), resourceGroupNameUnescaped, diskNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRevokeAccess = &respr
		d.beginRevokeAccess.add(req, beginRevokeAccess)
	}

	resp, err := server.PollerResponderNext(beginRevokeAccess, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginRevokeAccess.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRevokeAccess) {
		d.beginRevokeAccess.remove(req)
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := d.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/disks/(?P<diskName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armcompute.DiskUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		diskNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("diskName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUpdate(req.Context(), resourceGroupNameUnescaped, diskNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		d.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		d.beginUpdate.remove(req)
	}

	return resp, nil
}
