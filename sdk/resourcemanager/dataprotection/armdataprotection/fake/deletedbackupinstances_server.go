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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dataprotection/armdataprotection/v3"
	"net/http"
	"net/url"
	"regexp"
)

// DeletedBackupInstancesServer is a fake server for instances of the armdataprotection.DeletedBackupInstancesClient type.
type DeletedBackupInstancesServer struct {
	// Get is the fake for method DeletedBackupInstancesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *armdataprotection.DeletedBackupInstancesClientGetOptions) (resp azfake.Responder[armdataprotection.DeletedBackupInstancesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DeletedBackupInstancesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vaultName string, options *armdataprotection.DeletedBackupInstancesClientListOptions) (resp azfake.PagerResponder[armdataprotection.DeletedBackupInstancesClientListResponse])

	// BeginUndelete is the fake for method DeletedBackupInstancesClient.BeginUndelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginUndelete func(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *armdataprotection.DeletedBackupInstancesClientBeginUndeleteOptions) (resp azfake.PollerResponder[armdataprotection.DeletedBackupInstancesClientUndeleteResponse], errResp azfake.ErrorResponder)
}

// NewDeletedBackupInstancesServerTransport creates a new instance of DeletedBackupInstancesServerTransport with the provided implementation.
// The returned DeletedBackupInstancesServerTransport instance is connected to an instance of armdataprotection.DeletedBackupInstancesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDeletedBackupInstancesServerTransport(srv *DeletedBackupInstancesServer) *DeletedBackupInstancesServerTransport {
	return &DeletedBackupInstancesServerTransport{
		srv:           srv,
		newListPager:  newTracker[azfake.PagerResponder[armdataprotection.DeletedBackupInstancesClientListResponse]](),
		beginUndelete: newTracker[azfake.PollerResponder[armdataprotection.DeletedBackupInstancesClientUndeleteResponse]](),
	}
}

// DeletedBackupInstancesServerTransport connects instances of armdataprotection.DeletedBackupInstancesClient to instances of DeletedBackupInstancesServer.
// Don't use this type directly, use NewDeletedBackupInstancesServerTransport instead.
type DeletedBackupInstancesServerTransport struct {
	srv           *DeletedBackupInstancesServer
	newListPager  *tracker[azfake.PagerResponder[armdataprotection.DeletedBackupInstancesClientListResponse]]
	beginUndelete *tracker[azfake.PollerResponder[armdataprotection.DeletedBackupInstancesClientUndeleteResponse]]
}

// Do implements the policy.Transporter interface for DeletedBackupInstancesServerTransport.
func (d *DeletedBackupInstancesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DeletedBackupInstancesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if deletedBackupInstancesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = deletedBackupInstancesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DeletedBackupInstancesClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DeletedBackupInstancesClient.NewListPager":
				res.resp, res.err = d.dispatchNewListPager(req)
			case "DeletedBackupInstancesClient.BeginUndelete":
				res.resp, res.err = d.dispatchBeginUndelete(req)
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

func (d *DeletedBackupInstancesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedBackupInstances/(?P<backupInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	backupInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupInstanceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, backupInstanceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DeletedBackupInstanceResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DeletedBackupInstancesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedBackupInstances`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, vaultNameParam, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdataprotection.DeletedBackupInstancesClientListResponse, createLink func() string) {
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

func (d *DeletedBackupInstancesServerTransport) dispatchBeginUndelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginUndelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUndelete not implemented")}
	}
	beginUndelete := d.beginUndelete.get(req)
	if beginUndelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/deletedBackupInstances/(?P<backupInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/undelete`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
		if err != nil {
			return nil, err
		}
		backupInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupInstanceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginUndelete(req.Context(), resourceGroupNameParam, vaultNameParam, backupInstanceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUndelete = &respr
		d.beginUndelete.add(req, beginUndelete)
	}

	resp, err := server.PollerResponderNext(beginUndelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginUndelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUndelete) {
		d.beginUndelete.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to DeletedBackupInstancesServerTransport
var deletedBackupInstancesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
