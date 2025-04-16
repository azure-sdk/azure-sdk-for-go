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

// BackupPoliciesServer is a fake server for instances of the armdataprotection.BackupPoliciesClient type.
type BackupPoliciesServer struct {
	// CreateOrUpdate is the fake for method BackupPoliciesClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, parameters armdataprotection.BaseBackupPolicyResource, options *armdataprotection.BackupPoliciesClientCreateOrUpdateOptions) (resp azfake.Responder[armdataprotection.BackupPoliciesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method BackupPoliciesClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, options *armdataprotection.BackupPoliciesClientDeleteOptions) (resp azfake.Responder[armdataprotection.BackupPoliciesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method BackupPoliciesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, options *armdataprotection.BackupPoliciesClientGetOptions) (resp azfake.Responder[armdataprotection.BackupPoliciesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method BackupPoliciesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, vaultName string, options *armdataprotection.BackupPoliciesClientListOptions) (resp azfake.PagerResponder[armdataprotection.BackupPoliciesClientListResponse])
}

// NewBackupPoliciesServerTransport creates a new instance of BackupPoliciesServerTransport with the provided implementation.
// The returned BackupPoliciesServerTransport instance is connected to an instance of armdataprotection.BackupPoliciesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupPoliciesServerTransport(srv *BackupPoliciesServer) *BackupPoliciesServerTransport {
	return &BackupPoliciesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armdataprotection.BackupPoliciesClientListResponse]](),
	}
}

// BackupPoliciesServerTransport connects instances of armdataprotection.BackupPoliciesClient to instances of BackupPoliciesServer.
// Don't use this type directly, use NewBackupPoliciesServerTransport instead.
type BackupPoliciesServerTransport struct {
	srv          *BackupPoliciesServer
	newListPager *tracker[azfake.PagerResponder[armdataprotection.BackupPoliciesClientListResponse]]
}

// Do implements the policy.Transporter interface for BackupPoliciesServerTransport.
func (b *BackupPoliciesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BackupPoliciesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if backupPoliciesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = backupPoliciesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BackupPoliciesClient.CreateOrUpdate":
				res.resp, res.err = b.dispatchCreateOrUpdate(req)
			case "BackupPoliciesClient.Delete":
				res.resp, res.err = b.dispatchDelete(req)
			case "BackupPoliciesClient.Get":
				res.resp, res.err = b.dispatchGet(req)
			case "BackupPoliciesClient.NewListPager":
				res.resp, res.err = b.dispatchNewListPager(req)
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

func (b *BackupPoliciesServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdataprotection.BaseBackupPolicyResource](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, vaultNameParam, backupPolicyNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BaseBackupPolicyResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BackupPoliciesServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if b.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Delete(req.Context(), resourceGroupNameParam, vaultNameParam, backupPolicyNameParam, nil)
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

func (b *BackupPoliciesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies/(?P<backupPolicyName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	backupPolicyNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupPolicyName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), resourceGroupNameParam, vaultNameParam, backupPolicyNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BaseBackupPolicyResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BackupPoliciesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if b.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := b.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataProtection/backupVaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupPolicies`
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
		resp := b.srv.NewListPager(resourceGroupNameParam, vaultNameParam, nil)
		newListPager = &resp
		b.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armdataprotection.BackupPoliciesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		b.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		b.newListPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BackupPoliciesServerTransport
var backupPoliciesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
