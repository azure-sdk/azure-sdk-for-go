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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory/v8"
	"net/http"
	"net/url"
	"regexp"
)

// DataFlowDebugSessionServer is a fake server for instances of the armdatafactory.DataFlowDebugSessionClient type.
type DataFlowDebugSessionServer struct {
	// AddDataFlow is the fake for method DataFlowDebugSessionClient.AddDataFlow
	// HTTP status codes to indicate success: http.StatusOK
	AddDataFlow func(ctx context.Context, resourceGroupName string, factoryName string, request armdatafactory.DataFlowDebugPackage, options *armdatafactory.DataFlowDebugSessionClientAddDataFlowOptions) (resp azfake.Responder[armdatafactory.DataFlowDebugSessionClientAddDataFlowResponse], errResp azfake.ErrorResponder)

	// BeginCreate is the fake for method DataFlowDebugSessionClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCreate func(ctx context.Context, resourceGroupName string, factoryName string, request armdatafactory.CreateDataFlowDebugSessionRequest, options *armdatafactory.DataFlowDebugSessionClientBeginCreateOptions) (resp azfake.PollerResponder[armdatafactory.DataFlowDebugSessionClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DataFlowDebugSessionClient.Delete
	// HTTP status codes to indicate success: http.StatusOK
	Delete func(ctx context.Context, resourceGroupName string, factoryName string, request armdatafactory.DeleteDataFlowDebugSessionRequest, options *armdatafactory.DataFlowDebugSessionClientDeleteOptions) (resp azfake.Responder[armdatafactory.DataFlowDebugSessionClientDeleteResponse], errResp azfake.ErrorResponder)

	// BeginExecuteCommand is the fake for method DataFlowDebugSessionClient.BeginExecuteCommand
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginExecuteCommand func(ctx context.Context, resourceGroupName string, factoryName string, request armdatafactory.DataFlowDebugCommandRequest, options *armdatafactory.DataFlowDebugSessionClientBeginExecuteCommandOptions) (resp azfake.PollerResponder[armdatafactory.DataFlowDebugSessionClientExecuteCommandResponse], errResp azfake.ErrorResponder)

	// NewQueryByFactoryPager is the fake for method DataFlowDebugSessionClient.NewQueryByFactoryPager
	// HTTP status codes to indicate success: http.StatusOK
	NewQueryByFactoryPager func(resourceGroupName string, factoryName string, options *armdatafactory.DataFlowDebugSessionClientQueryByFactoryOptions) (resp azfake.PagerResponder[armdatafactory.DataFlowDebugSessionClientQueryByFactoryResponse])
}

// NewDataFlowDebugSessionServerTransport creates a new instance of DataFlowDebugSessionServerTransport with the provided implementation.
// The returned DataFlowDebugSessionServerTransport instance is connected to an instance of armdatafactory.DataFlowDebugSessionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDataFlowDebugSessionServerTransport(srv *DataFlowDebugSessionServer) *DataFlowDebugSessionServerTransport {
	return &DataFlowDebugSessionServerTransport{
		srv:                    srv,
		beginCreate:            newTracker[azfake.PollerResponder[armdatafactory.DataFlowDebugSessionClientCreateResponse]](),
		beginExecuteCommand:    newTracker[azfake.PollerResponder[armdatafactory.DataFlowDebugSessionClientExecuteCommandResponse]](),
		newQueryByFactoryPager: newTracker[azfake.PagerResponder[armdatafactory.DataFlowDebugSessionClientQueryByFactoryResponse]](),
	}
}

// DataFlowDebugSessionServerTransport connects instances of armdatafactory.DataFlowDebugSessionClient to instances of DataFlowDebugSessionServer.
// Don't use this type directly, use NewDataFlowDebugSessionServerTransport instead.
type DataFlowDebugSessionServerTransport struct {
	srv                    *DataFlowDebugSessionServer
	beginCreate            *tracker[azfake.PollerResponder[armdatafactory.DataFlowDebugSessionClientCreateResponse]]
	beginExecuteCommand    *tracker[azfake.PollerResponder[armdatafactory.DataFlowDebugSessionClientExecuteCommandResponse]]
	newQueryByFactoryPager *tracker[azfake.PagerResponder[armdatafactory.DataFlowDebugSessionClientQueryByFactoryResponse]]
}

// Do implements the policy.Transporter interface for DataFlowDebugSessionServerTransport.
func (d *DataFlowDebugSessionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DataFlowDebugSessionClient.AddDataFlow":
		resp, err = d.dispatchAddDataFlow(req)
	case "DataFlowDebugSessionClient.BeginCreate":
		resp, err = d.dispatchBeginCreate(req)
	case "DataFlowDebugSessionClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DataFlowDebugSessionClient.BeginExecuteCommand":
		resp, err = d.dispatchBeginExecuteCommand(req)
	case "DataFlowDebugSessionClient.NewQueryByFactoryPager":
		resp, err = d.dispatchNewQueryByFactoryPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DataFlowDebugSessionServerTransport) dispatchAddDataFlow(req *http.Request) (*http.Response, error) {
	if d.srv.AddDataFlow == nil {
		return nil, &nonRetriableError{errors.New("fake for method AddDataFlow not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/addDataFlowToDebugSession`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.DataFlowDebugPackage](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.AddDataFlow(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AddDataFlowToDebugSessionResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataFlowDebugSessionServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := d.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/createDataFlowDebugSession`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatafactory.CreateDataFlowDebugSessionRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreate(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		d.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		d.beginCreate.remove(req)
	}

	return resp, nil
}

func (d *DataFlowDebugSessionServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deleteDataFlowDebugSession`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdatafactory.DeleteDataFlowDebugSessionRequest](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataFlowDebugSessionServerTransport) dispatchBeginExecuteCommand(req *http.Request) (*http.Response, error) {
	if d.srv.BeginExecuteCommand == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginExecuteCommand not implemented")}
	}
	beginExecuteCommand := d.beginExecuteCommand.get(req)
	if beginExecuteCommand == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/executeDataFlowDebugCommand`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatafactory.DataFlowDebugCommandRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginExecuteCommand(req.Context(), resourceGroupNameParam, factoryNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginExecuteCommand = &respr
		d.beginExecuteCommand.add(req, beginExecuteCommand)
	}

	resp, err := server.PollerResponderNext(beginExecuteCommand, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginExecuteCommand.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginExecuteCommand) {
		d.beginExecuteCommand.remove(req)
	}

	return resp, nil
}

func (d *DataFlowDebugSessionServerTransport) dispatchNewQueryByFactoryPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewQueryByFactoryPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewQueryByFactoryPager not implemented")}
	}
	newQueryByFactoryPager := d.newQueryByFactoryPager.get(req)
	if newQueryByFactoryPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataFactory/factories/(?P<factoryName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/queryDataFlowDebugSessions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		factoryNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("factoryName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewQueryByFactoryPager(resourceGroupNameParam, factoryNameParam, nil)
		newQueryByFactoryPager = &resp
		d.newQueryByFactoryPager.add(req, newQueryByFactoryPager)
		server.PagerResponderInjectNextLinks(newQueryByFactoryPager, req, func(page *armdatafactory.DataFlowDebugSessionClientQueryByFactoryResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newQueryByFactoryPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newQueryByFactoryPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newQueryByFactoryPager) {
		d.newQueryByFactoryPager.remove(req)
	}
	return resp, nil
}
