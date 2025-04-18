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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devtestlabs/armdevtestlabs"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// VirtualMachineSchedulesServer is a fake server for instances of the armdevtestlabs.VirtualMachineSchedulesClient type.
type VirtualMachineSchedulesServer struct {
	// CreateOrUpdate is the fake for method VirtualMachineSchedulesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule armdevtestlabs.Schedule, options *armdevtestlabs.VirtualMachineSchedulesClientCreateOrUpdateOptions) (resp azfake.Responder[armdevtestlabs.VirtualMachineSchedulesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method VirtualMachineSchedulesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, options *armdevtestlabs.VirtualMachineSchedulesClientDeleteOptions) (resp azfake.Responder[armdevtestlabs.VirtualMachineSchedulesClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExecute is the fake for method VirtualMachineSchedulesClient.BeginExecute
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginExecute func(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, options *armdevtestlabs.VirtualMachineSchedulesClientBeginExecuteOptions) (resp azfake.PollerResponder[armdevtestlabs.VirtualMachineSchedulesClientExecuteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineSchedulesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, options *armdevtestlabs.VirtualMachineSchedulesClientGetOptions) (resp azfake.Responder[armdevtestlabs.VirtualMachineSchedulesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineSchedulesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, labName string, virtualMachineName string, options *armdevtestlabs.VirtualMachineSchedulesClientListOptions) (resp azfake.PagerResponder[armdevtestlabs.VirtualMachineSchedulesClientListResponse])

	// Update is the fake for method VirtualMachineSchedulesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, labName string, virtualMachineName string, name string, schedule armdevtestlabs.ScheduleFragment, options *armdevtestlabs.VirtualMachineSchedulesClientUpdateOptions) (resp azfake.Responder[armdevtestlabs.VirtualMachineSchedulesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineSchedulesServerTransport creates a new instance of VirtualMachineSchedulesServerTransport with the provided implementation.
// The returned VirtualMachineSchedulesServerTransport instance is connected to an instance of armdevtestlabs.VirtualMachineSchedulesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineSchedulesServerTransport(srv *VirtualMachineSchedulesServer) *VirtualMachineSchedulesServerTransport {
	return &VirtualMachineSchedulesServerTransport{
		srv:          srv,
		beginExecute: newTracker[azfake.PollerResponder[armdevtestlabs.VirtualMachineSchedulesClientExecuteResponse]](),
		newListPager: newTracker[azfake.PagerResponder[armdevtestlabs.VirtualMachineSchedulesClientListResponse]](),
	}
}

// VirtualMachineSchedulesServerTransport connects instances of armdevtestlabs.VirtualMachineSchedulesClient to instances of VirtualMachineSchedulesServer.
// Don't use this type directly, use NewVirtualMachineSchedulesServerTransport instead.
type VirtualMachineSchedulesServerTransport struct {
	srv          *VirtualMachineSchedulesServer
	beginExecute *tracker[azfake.PollerResponder[armdevtestlabs.VirtualMachineSchedulesClientExecuteResponse]]
	newListPager *tracker[azfake.PagerResponder[armdevtestlabs.VirtualMachineSchedulesClientListResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineSchedulesServerTransport.
func (v *VirtualMachineSchedulesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualMachineSchedulesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if virtualMachineSchedulesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = virtualMachineSchedulesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VirtualMachineSchedulesClient.CreateOrUpdate":
				res.resp, res.err = v.dispatchCreateOrUpdate(req)
			case "VirtualMachineSchedulesClient.Delete":
				res.resp, res.err = v.dispatchDelete(req)
			case "VirtualMachineSchedulesClient.BeginExecute":
				res.resp, res.err = v.dispatchBeginExecute(req)
			case "VirtualMachineSchedulesClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VirtualMachineSchedulesClient.NewListPager":
				res.resp, res.err = v.dispatchNewListPager(req)
			case "VirtualMachineSchedulesClient.Update":
				res.resp, res.err = v.dispatchUpdate(req)
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

func (v *VirtualMachineSchedulesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualmachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.Schedule](req)
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
	virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, labNameParam, virtualMachineNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Schedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineSchedulesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if v.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualmachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Delete(req.Context(), resourceGroupNameParam, labNameParam, virtualMachineNameParam, nameParam, nil)
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

func (v *VirtualMachineSchedulesServerTransport) dispatchBeginExecute(req *http.Request) (*http.Response, error) {
	if v.srv.BeginExecute == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExecute not implemented")}
	}
	beginExecute := v.beginExecute.get(req)
	if beginExecute == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualmachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/execute`
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
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
		if err != nil {
			return nil, err
		}
		nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginExecute(req.Context(), resourceGroupNameParam, labNameParam, virtualMachineNameParam, nameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExecute = &respr
		v.beginExecute.add(req, beginExecute)
	}

	resp, err := server.PollerResponderNext(beginExecute, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginExecute.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExecute) {
		v.beginExecute.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineSchedulesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualmachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
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
	var options *armdevtestlabs.VirtualMachineSchedulesClientGetOptions
	if expandParam != nil {
		options = &armdevtestlabs.VirtualMachineSchedulesClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, labNameParam, virtualMachineNameParam, nameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Schedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineSchedulesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualmachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules`
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
		virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
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
		var options *armdevtestlabs.VirtualMachineSchedulesClientListOptions
		if expandParam != nil || filterParam != nil || topParam != nil || orderbyParam != nil {
			options = &armdevtestlabs.VirtualMachineSchedulesClientListOptions{
				Expand:  expandParam,
				Filter:  filterParam,
				Top:     topParam,
				Orderby: orderbyParam,
			}
		}
		resp := v.srv.NewListPager(resourceGroupNameParam, labNameParam, virtualMachineNameParam, options)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdevtestlabs.VirtualMachineSchedulesClientListResponse, createLink func() string) {
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

func (v *VirtualMachineSchedulesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DevTestLab/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/virtualmachines/(?P<virtualMachineName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/schedules/(?P<name>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdevtestlabs.ScheduleFragment](req)
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
	virtualMachineNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualMachineName")])
	if err != nil {
		return nil, err
	}
	nameParam, err := url.PathUnescape(matches[regex.SubexpIndex("name")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Update(req.Context(), resourceGroupNameParam, labNameParam, virtualMachineNameParam, nameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Schedule, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to VirtualMachineSchedulesServerTransport
var virtualMachineSchedulesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
