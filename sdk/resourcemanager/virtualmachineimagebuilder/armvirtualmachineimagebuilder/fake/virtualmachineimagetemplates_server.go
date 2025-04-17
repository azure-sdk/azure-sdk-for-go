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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/virtualmachineimagebuilder/armvirtualmachineimagebuilder/v2"
	"net/http"
	"net/url"
	"regexp"
)

// VirtualMachineImageTemplatesServer is a fake server for instances of the armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClient type.
type VirtualMachineImageTemplatesServer struct {
	// BeginCancel is the fake for method VirtualMachineImageTemplatesClient.BeginCancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginCancel func(ctx context.Context, resourceGroupName string, imageTemplateName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientBeginCancelOptions) (resp azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientCancelResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method VirtualMachineImageTemplatesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters armvirtualmachineimagebuilder.ImageTemplate, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method VirtualMachineImageTemplatesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, imageTemplateName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientBeginDeleteOptions) (resp azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method VirtualMachineImageTemplatesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, imageTemplateName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientGetOptions) (resp azfake.Responder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientGetResponse], errResp azfake.ErrorResponder)

	// GetRunOutput is the fake for method VirtualMachineImageTemplatesClient.GetRunOutput
	// HTTP status codes to indicate success: http.StatusOK
	GetRunOutput func(ctx context.Context, resourceGroupName string, imageTemplateName string, runOutputName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientGetRunOutputOptions) (resp azfake.Responder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientGetRunOutputResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method VirtualMachineImageTemplatesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListOptions) (resp azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListResponse])

	// NewListByResourceGroupPager is the fake for method VirtualMachineImageTemplatesClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListByResourceGroupOptions) (resp azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListByResourceGroupResponse])

	// NewListRunOutputsPager is the fake for method VirtualMachineImageTemplatesClient.NewListRunOutputsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListRunOutputsPager func(resourceGroupName string, imageTemplateName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListRunOutputsOptions) (resp azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListRunOutputsResponse])

	// BeginRun is the fake for method VirtualMachineImageTemplatesClient.BeginRun
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginRun func(ctx context.Context, resourceGroupName string, imageTemplateName string, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientBeginRunOptions) (resp azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientRunResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method VirtualMachineImageTemplatesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, imageTemplateName string, parameters armvirtualmachineimagebuilder.ImageTemplateUpdateParameters, options *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientBeginUpdateOptions) (resp azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineImageTemplatesServerTransport creates a new instance of VirtualMachineImageTemplatesServerTransport with the provided implementation.
// The returned VirtualMachineImageTemplatesServerTransport instance is connected to an instance of armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineImageTemplatesServerTransport(srv *VirtualMachineImageTemplatesServer) *VirtualMachineImageTemplatesServerTransport {
	return &VirtualMachineImageTemplatesServerTransport{
		srv:                         srv,
		beginCancel:                 newTracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientCancelResponse]](),
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListByResourceGroupResponse]](),
		newListRunOutputsPager:      newTracker[azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListRunOutputsResponse]](),
		beginRun:                    newTracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientRunResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientUpdateResponse]](),
	}
}

// VirtualMachineImageTemplatesServerTransport connects instances of armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClient to instances of VirtualMachineImageTemplatesServer.
// Don't use this type directly, use NewVirtualMachineImageTemplatesServerTransport instead.
type VirtualMachineImageTemplatesServerTransport struct {
	srv                         *VirtualMachineImageTemplatesServer
	beginCancel                 *tracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientCancelResponse]]
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListByResourceGroupResponse]]
	newListRunOutputsPager      *tracker[azfake.PagerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListRunOutputsResponse]]
	beginRun                    *tracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientRunResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for VirtualMachineImageTemplatesServerTransport.
func (v *VirtualMachineImageTemplatesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VirtualMachineImageTemplatesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if virtualMachineImageTemplatesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = virtualMachineImageTemplatesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VirtualMachineImageTemplatesClient.BeginCancel":
				res.resp, res.err = v.dispatchBeginCancel(req)
			case "VirtualMachineImageTemplatesClient.BeginCreateOrUpdate":
				res.resp, res.err = v.dispatchBeginCreateOrUpdate(req)
			case "VirtualMachineImageTemplatesClient.BeginDelete":
				res.resp, res.err = v.dispatchBeginDelete(req)
			case "VirtualMachineImageTemplatesClient.Get":
				res.resp, res.err = v.dispatchGet(req)
			case "VirtualMachineImageTemplatesClient.GetRunOutput":
				res.resp, res.err = v.dispatchGetRunOutput(req)
			case "VirtualMachineImageTemplatesClient.NewListPager":
				res.resp, res.err = v.dispatchNewListPager(req)
			case "VirtualMachineImageTemplatesClient.NewListByResourceGroupPager":
				res.resp, res.err = v.dispatchNewListByResourceGroupPager(req)
			case "VirtualMachineImageTemplatesClient.NewListRunOutputsPager":
				res.resp, res.err = v.dispatchNewListRunOutputsPager(req)
			case "VirtualMachineImageTemplatesClient.BeginRun":
				res.resp, res.err = v.dispatchBeginRun(req)
			case "VirtualMachineImageTemplatesClient.BeginUpdate":
				res.resp, res.err = v.dispatchBeginUpdate(req)
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

func (v *VirtualMachineImageTemplatesServerTransport) dispatchBeginCancel(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancel not implemented")}
	}
	beginCancel := v.beginCancel.get(req)
	if beginCancel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCancel(req.Context(), resourceGroupNameParam, imageTemplateNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCancel = &respr
		v.beginCancel.add(req, beginCancel)
	}

	resp, err := server.PollerResponderNext(beginCancel, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginCancel.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCancel) {
		v.beginCancel.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineImageTemplatesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := v.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armvirtualmachineimagebuilder.ImageTemplate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, imageTemplateNameParam, body, nil)
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

func (v *VirtualMachineImageTemplatesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := v.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDelete(req.Context(), resourceGroupNameParam, imageTemplateNameParam, nil)
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

func (v *VirtualMachineImageTemplatesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), resourceGroupNameParam, imageTemplateNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ImageTemplate, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImageTemplatesServerTransport) dispatchGetRunOutput(req *http.Request) (*http.Response, error) {
	if v.srv.GetRunOutput == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetRunOutput not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runOutputs/(?P<runOutputName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
	if err != nil {
		return nil, err
	}
	runOutputNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("runOutputName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.GetRunOutput(req.Context(), resourceGroupNameParam, imageTemplateNameParam, runOutputNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RunOutput, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImageTemplatesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := v.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := v.srv.NewListPager(nil)
		newListPager = &resp
		v.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListResponse, createLink func() string) {
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

func (v *VirtualMachineImageTemplatesServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := v.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates`
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
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListByResourceGroupResponse, createLink func() string) {
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

func (v *VirtualMachineImageTemplatesServerTransport) dispatchNewListRunOutputsPager(req *http.Request) (*http.Response, error) {
	if v.srv.NewListRunOutputsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListRunOutputsPager not implemented")}
	}
	newListRunOutputsPager := v.newListRunOutputsPager.get(req)
	if newListRunOutputsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/runOutputs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
		if err != nil {
			return nil, err
		}
		resp := v.srv.NewListRunOutputsPager(resourceGroupNameParam, imageTemplateNameParam, nil)
		newListRunOutputsPager = &resp
		v.newListRunOutputsPager.add(req, newListRunOutputsPager)
		server.PagerResponderInjectNextLinks(newListRunOutputsPager, req, func(page *armvirtualmachineimagebuilder.VirtualMachineImageTemplatesClientListRunOutputsResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListRunOutputsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		v.newListRunOutputsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListRunOutputsPager) {
		v.newListRunOutputsPager.remove(req)
	}
	return resp, nil
}

func (v *VirtualMachineImageTemplatesServerTransport) dispatchBeginRun(req *http.Request) (*http.Response, error) {
	if v.srv.BeginRun == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRun not implemented")}
	}
	beginRun := v.beginRun.get(req)
	if beginRun == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/run`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginRun(req.Context(), resourceGroupNameParam, imageTemplateNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRun = &respr
		v.beginRun.add(req, beginRun)
	}

	resp, err := server.PollerResponderNext(beginRun, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginRun.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRun) {
		v.beginRun.remove(req)
	}

	return resp, nil
}

func (v *VirtualMachineImageTemplatesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if v.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := v.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.VirtualMachineImages/imageTemplates/(?P<imageTemplateName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armvirtualmachineimagebuilder.ImageTemplateUpdateParameters](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		imageTemplateNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("imageTemplateName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginUpdate(req.Context(), resourceGroupNameParam, imageTemplateNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		v.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		v.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		v.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to VirtualMachineImageTemplatesServerTransport
var virtualMachineImageTemplatesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
