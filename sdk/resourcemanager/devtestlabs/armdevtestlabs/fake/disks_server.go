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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs/v2"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DisksServer is a fake server for instances of the armdevtestlabs.DisksClient type.
type DisksServer struct {
	// BeginAttach is the fake for method DisksClient.BeginAttach
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginAttach func(ctx context.Context, resourceGroupName string, labName string, userName string, name string, attachDiskProperties armdevtestlabs.AttachDiskProperties, options *armdevtestlabs.DisksClientBeginAttachOptions) (resp azfake.PollerResponder[armdevtestlabs.DisksClientAttachResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method DisksClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk armdevtestlabs.Disk, options *armdevtestlabs.DisksClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdevtestlabs.DisksClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DisksClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *armdevtestlabs.DisksClientBeginDeleteOptions) (resp azfake.PollerResponder[armdevtestlabs.DisksClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginDetach is the fake for method DisksClient.BeginDetach
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginDetach func(ctx context.Context, resourceGroupName string, labName string, userName string, name string, detachDiskProperties armdevtestlabs.DetachDiskProperties, options *armdevtestlabs.DisksClientBeginDetachOptions) (resp azfake.PollerResponder[armdevtestlabs.DisksClientDetachResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DisksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, labName string, userName string, name string, options *armdevtestlabs.DisksClientGetOptions) (resp azfake.Responder[armdevtestlabs.DisksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DisksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, labName string, userName string, options *armdevtestlabs.DisksClientListOptions) (resp azfake.PagerResponder[armdevtestlabs.DisksClientListResponse])

	// Update is the fake for method DisksClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, labName string, userName string, name string, disk armdevtestlabs.DiskFragment, options *armdevtestlabs.DisksClientUpdateOptions) (resp azfake.Responder[armdevtestlabs.DisksClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDisksServerTransport creates a new instance of DisksServerTransport with the provided implementation.
// The returned DisksServerTransport instance is connected to an instance of armdevtestlabs.DisksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDisksServerTransport(srv *DisksServer) *DisksServerTransport {
	return &DisksServerTransport{
		srv:                 srv,
		beginAttach:         newTracker[azfake.PollerResponder[armdevtestlabs.DisksClientAttachResponse]](),
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdevtestlabs.DisksClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdevtestlabs.DisksClientDeleteResponse]](),
		beginDetach:         newTracker[azfake.PollerResponder[armdevtestlabs.DisksClientDetachResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armdevtestlabs.DisksClientListResponse]](),
	}
}

// DisksServerTransport connects instances of armdevtestlabs.DisksClient to instances of DisksServer.
// Don't use this type directly, use NewDisksServerTransport instead.
type DisksServerTransport struct {
	srv                 *DisksServer
	beginAttach         *tracker[azfake.PollerResponder[armdevtestlabs.DisksClientAttachResponse]]
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdevtestlabs.DisksClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdevtestlabs.DisksClientDeleteResponse]]
	beginDetach         *tracker[azfake.PollerResponder[armdevtestlabs.DisksClientDetachResponse]]
	newListPager        *tracker[azfake.PagerResponder[armdevtestlabs.DisksClientListResponse]]
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
	case "DisksClient.BeginAttach":
		resp, err = d.dispatchBeginAttach(req)
	case "DisksClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DisksClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DisksClient.BeginDetach":
		resp, err = d.dispatchBeginDetach(req)
	case "DisksClient.Get":
		resp, err = d.dispatchGet(req)
	case "DisksClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	case "DisksClient.Update":
		resp, err = d.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginAttach(req *http.Request) (*http.Response, error) {
	if d.srv.BeginAttach == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginAttach not implemented")}
	}
	beginAttach := d.beginAttach.get(req)
	if beginAttach == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/attach`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.AttachDiskProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginAttach(req.Context(), resourceGroupNameParam, labNameParam, userNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginAttach = &respr
		d.beginAttach.add(req, beginAttach)
	}

	resp, err := server.PollerResponderNext(beginAttach, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginAttach.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginAttach) {
		d.beginAttach.remove(req)
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.Disk](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, labNameParam, userNameParam, nameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
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
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, labNameParam, userNameParam, nameParam, nil)
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

func (d *DisksServerTransport) dispatchBeginDetach(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDetach == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDetach not implemented")}
	}
	beginDetach := d.beginDetach.get(req)
	if beginDetach == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/detach`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.DetachDiskProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDetach(req.Context(), resourceGroupNameParam, labNameParam, userNameParam, nameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDetach = &respr
		d.beginDetach.add(req, beginDetach)
	}

	resp, err := server.PollerResponderNext(beginDetach, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginDetach.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDetach) {
		d.beginDetach.remove(req)
	}

	return resp, nil
}

func (d *DisksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armdevtestlabs.DisksClientGetOptions
	if expandParam != nil {
		options = &armdevtestlabs.DisksClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, labNameParam, userNameParam, nameParam, options)
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

func (d *DisksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
		if err != nil {
			return nil, err
		}
		expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
		if err != nil {
			return nil, err
		}
		expandParam := getOptional(expandUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
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
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdevtestlabs.DisksClientListOptions
		if expandParam != nil || filterParam != nil || topParam != nil || orderbyParam != nil {
			options = &armdevtestlabs.DisksClientListOptions{
				Expand:  expandParam,
				Filter:  filterParam,
				Top:     topParam,
				Orderby: orderbyParam,
			}
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, labNameParam, userNameParam, options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdevtestlabs.DisksClientListResponse, createLink func() string) {
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

func (d *DisksServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/users/(?P<userName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/disks/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.DiskFragment](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	userNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("userName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Update(req.Context(), resourceGroupNameParam, labNameParam, userNameParam, nameParam, body, nil)
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
