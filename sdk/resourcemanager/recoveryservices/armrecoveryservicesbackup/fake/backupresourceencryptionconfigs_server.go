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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v5"
	"net/http"
	"net/url"
	"regexp"
)

// BackupResourceEncryptionConfigsServer is a fake server for instances of the armrecoveryservicesbackup.BackupResourceEncryptionConfigsClient type.
type BackupResourceEncryptionConfigsServer struct {
	// Get is the fake for method BackupResourceEncryptionConfigsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, vaultName string, resourceGroupName string, options *armrecoveryservicesbackup.BackupResourceEncryptionConfigsClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.BackupResourceEncryptionConfigsClientGetResponse], errResp azfake.ErrorResponder)

	// Update is the fake for method BackupResourceEncryptionConfigsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, vaultName string, resourceGroupName string, parameters armrecoveryservicesbackup.BackupResourceEncryptionConfigResource, options *armrecoveryservicesbackup.BackupResourceEncryptionConfigsClientUpdateOptions) (resp azfake.Responder[armrecoveryservicesbackup.BackupResourceEncryptionConfigsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewBackupResourceEncryptionConfigsServerTransport creates a new instance of BackupResourceEncryptionConfigsServerTransport with the provided implementation.
// The returned BackupResourceEncryptionConfigsServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupResourceEncryptionConfigsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupResourceEncryptionConfigsServerTransport(srv *BackupResourceEncryptionConfigsServer) *BackupResourceEncryptionConfigsServerTransport {
	return &BackupResourceEncryptionConfigsServerTransport{srv: srv}
}

// BackupResourceEncryptionConfigsServerTransport connects instances of armrecoveryservicesbackup.BackupResourceEncryptionConfigsClient to instances of BackupResourceEncryptionConfigsServer.
// Don't use this type directly, use NewBackupResourceEncryptionConfigsServerTransport instead.
type BackupResourceEncryptionConfigsServerTransport struct {
	srv *BackupResourceEncryptionConfigsServer
}

// Do implements the policy.Transporter interface for BackupResourceEncryptionConfigsServerTransport.
func (b *BackupResourceEncryptionConfigsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BackupResourceEncryptionConfigsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if backupResourceEncryptionConfigsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = backupResourceEncryptionConfigsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BackupResourceEncryptionConfigsClient.Get":
				res.resp, res.err = b.dispatchGet(req)
			case "BackupResourceEncryptionConfigsClient.Update":
				res.resp, res.err = b.dispatchUpdate(req)
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

func (b *BackupResourceEncryptionConfigsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupEncryptionConfigs/backupResourceEncryptionConfig`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), vaultNameParam, resourceGroupNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackupResourceEncryptionConfigExtendedResource, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (b *BackupResourceEncryptionConfigsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if b.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<vaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupEncryptionConfigs/backupResourceEncryptionConfig`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.BackupResourceEncryptionConfigResource](req)
	if err != nil {
		return nil, err
	}
	vaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("vaultName")])
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Update(req.Context(), vaultNameParam, resourceGroupNameParam, body, nil)
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

// set this to conditionally intercept incoming requests to BackupResourceEncryptionConfigsServerTransport
var backupResourceEncryptionConfigsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
