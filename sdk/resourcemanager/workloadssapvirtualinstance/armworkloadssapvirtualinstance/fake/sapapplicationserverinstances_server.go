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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/workloadssapvirtualinstance/armworkloadssapvirtualinstance"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
)

// SAPApplicationServerInstancesServer is a fake server for instances of the armworkloadssapvirtualinstance.SAPApplicationServerInstancesClient type.
type SAPApplicationServerInstancesServer struct {
	// BeginCreate is the fake for method SAPApplicationServerInstancesClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body armworkloadssapvirtualinstance.SAPApplicationServerInstance, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginCreateOptions) (resp azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method SAPApplicationServerInstancesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginDeleteOptions) (resp azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method SAPApplicationServerInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientGetOptions) (resp azfake.Responder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method SAPApplicationServerInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, sapVirtualInstanceName string, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientListOptions) (resp azfake.PagerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientListResponse])

	// BeginStartInstance is the fake for method SAPApplicationServerInstancesClient.BeginStartInstance
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStartInstance func(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginStartInstanceOptions) (resp azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientStartInstanceResponse], errResp azfake.ErrorResponder)

	// BeginStopInstance is the fake for method SAPApplicationServerInstancesClient.BeginStopInstance
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginStopInstance func(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginStopInstanceOptions) (resp azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientStopInstanceResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method SAPApplicationServerInstancesClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, sapVirtualInstanceName string, applicationInstanceName string, body armworkloadssapvirtualinstance.UpdateSAPApplicationInstanceRequest, options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientUpdateOptions) (resp azfake.Responder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewSAPApplicationServerInstancesServerTransport creates a new instance of SAPApplicationServerInstancesServerTransport with the provided implementation.
// The returned SAPApplicationServerInstancesServerTransport instance is connected to an instance of armworkloadssapvirtualinstance.SAPApplicationServerInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSAPApplicationServerInstancesServerTransport(srv *SAPApplicationServerInstancesServer) *SAPApplicationServerInstancesServerTransport {
	return &SAPApplicationServerInstancesServerTransport{
		srv:                srv,
		beginCreate:        newTracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientCreateResponse]](),
		beginDelete:        newTracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientDeleteResponse]](),
		newListPager:       newTracker[azfake.PagerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientListResponse]](),
		beginStartInstance: newTracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientStartInstanceResponse]](),
		beginStopInstance:  newTracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientStopInstanceResponse]](),
	}
}

// SAPApplicationServerInstancesServerTransport connects instances of armworkloadssapvirtualinstance.SAPApplicationServerInstancesClient to instances of SAPApplicationServerInstancesServer.
// Don't use this type directly, use NewSAPApplicationServerInstancesServerTransport instead.
type SAPApplicationServerInstancesServerTransport struct {
	srv                *SAPApplicationServerInstancesServer
	beginCreate        *tracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientCreateResponse]]
	beginDelete        *tracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientDeleteResponse]]
	newListPager       *tracker[azfake.PagerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientListResponse]]
	beginStartInstance *tracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientStartInstanceResponse]]
	beginStopInstance  *tracker[azfake.PollerResponder[armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientStopInstanceResponse]]
}

// Do implements the policy.Transporter interface for SAPApplicationServerInstancesServerTransport.
func (s *SAPApplicationServerInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SAPApplicationServerInstancesClient.BeginCreate":
		resp, err = s.dispatchBeginCreate(req)
	case "SAPApplicationServerInstancesClient.BeginDelete":
		resp, err = s.dispatchBeginDelete(req)
	case "SAPApplicationServerInstancesClient.Get":
		resp, err = s.dispatchGet(req)
	case "SAPApplicationServerInstancesClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	case "SAPApplicationServerInstancesClient.BeginStartInstance":
		resp, err = s.dispatchBeginStartInstance(req)
	case "SAPApplicationServerInstancesClient.BeginStopInstance":
		resp, err = s.dispatchBeginStopInstance(req)
	case "SAPApplicationServerInstancesClient.Update":
		resp, err = s.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SAPApplicationServerInstancesServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if s.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := s.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances/(?P<applicationInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.SAPApplicationServerInstance](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
		if err != nil {
			return nil, err
		}
		applicationInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginCreate(req.Context(), resourceGroupNameParam, sapVirtualInstanceNameParam, applicationInstanceNameParam, body, nil)
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

func (s *SAPApplicationServerInstancesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if s.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := s.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances/(?P<applicationInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
		if err != nil {
			return nil, err
		}
		applicationInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginDelete(req.Context(), resourceGroupNameParam, sapVirtualInstanceNameParam, applicationInstanceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		s.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		s.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		s.beginDelete.remove(req)
	}

	return resp, nil
}

func (s *SAPApplicationServerInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances/(?P<applicationInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
	if err != nil {
		return nil, err
	}
	applicationInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameParam, sapVirtualInstanceNameParam, applicationInstanceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPApplicationServerInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *SAPApplicationServerInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := s.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
		if err != nil {
			return nil, err
		}
		resp := s.srv.NewListPager(resourceGroupNameParam, sapVirtualInstanceNameParam, nil)
		newListPager = &resp
		s.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		s.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		s.newListPager.remove(req)
	}
	return resp, nil
}

func (s *SAPApplicationServerInstancesServerTransport) dispatchBeginStartInstance(req *http.Request) (*http.Response, error) {
	if s.srv.BeginStartInstance == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStartInstance not implemented")}
	}
	beginStartInstance := s.beginStartInstance.get(req)
	if beginStartInstance == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances/(?P<applicationInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/start`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.StartRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
		if err != nil {
			return nil, err
		}
		applicationInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationInstanceName")])
		if err != nil {
			return nil, err
		}
		var options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginStartInstanceOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginStartInstanceOptions{
				Body: &body,
			}
		}
		respr, errRespr := s.srv.BeginStartInstance(req.Context(), resourceGroupNameParam, sapVirtualInstanceNameParam, applicationInstanceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStartInstance = &respr
		s.beginStartInstance.add(req, beginStartInstance)
	}

	resp, err := server.PollerResponderNext(beginStartInstance, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginStartInstance.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStartInstance) {
		s.beginStartInstance.remove(req)
	}

	return resp, nil
}

func (s *SAPApplicationServerInstancesServerTransport) dispatchBeginStopInstance(req *http.Request) (*http.Response, error) {
	if s.srv.BeginStopInstance == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginStopInstance not implemented")}
	}
	beginStopInstance := s.beginStopInstance.get(req)
	if beginStopInstance == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances/(?P<applicationInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/stop`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.StopRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
		if err != nil {
			return nil, err
		}
		applicationInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationInstanceName")])
		if err != nil {
			return nil, err
		}
		var options *armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginStopInstanceOptions
		if !reflect.ValueOf(body).IsZero() {
			options = &armworkloadssapvirtualinstance.SAPApplicationServerInstancesClientBeginStopInstanceOptions{
				Body: &body,
			}
		}
		respr, errRespr := s.srv.BeginStopInstance(req.Context(), resourceGroupNameParam, sapVirtualInstanceNameParam, applicationInstanceNameParam, options)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginStopInstance = &respr
		s.beginStopInstance.add(req, beginStopInstance)
	}

	resp, err := server.PollerResponderNext(beginStopInstance, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginStopInstance.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginStopInstance) {
		s.beginStopInstance.remove(req)
	}

	return resp, nil
}

func (s *SAPApplicationServerInstancesServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Workloads/sapVirtualInstances/(?P<sapVirtualInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/applicationInstances/(?P<applicationInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armworkloadssapvirtualinstance.UpdateSAPApplicationInstanceRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	sapVirtualInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sapVirtualInstanceName")])
	if err != nil {
		return nil, err
	}
	applicationInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Update(req.Context(), resourceGroupNameParam, sapVirtualInstanceNameParam, applicationInstanceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).SAPApplicationServerInstance, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
