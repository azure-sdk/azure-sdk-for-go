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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/netapp/armnetapp/v6"
	"net/http"
	"net/url"
	"regexp"
)

// BackupsUnderBackupVaultServer is a fake server for instances of the armnetapp.BackupsUnderBackupVaultClient type.
type BackupsUnderBackupVaultServer struct {
	// BeginRestoreFiles is the fake for method BackupsUnderBackupVaultClient.BeginRestoreFiles
	// HTTP status codes to indicate success: http.StatusAccepted
	BeginRestoreFiles func(ctx context.Context, resourceGroupName string, accountName string, backupVaultName string, backupName string, body armnetapp.BackupRestoreFiles, options *armnetapp.BackupsUnderBackupVaultClientBeginRestoreFilesOptions) (resp azfake.PollerResponder[armnetapp.BackupsUnderBackupVaultClientRestoreFilesResponse], errResp azfake.ErrorResponder)
}

// NewBackupsUnderBackupVaultServerTransport creates a new instance of BackupsUnderBackupVaultServerTransport with the provided implementation.
// The returned BackupsUnderBackupVaultServerTransport instance is connected to an instance of armnetapp.BackupsUnderBackupVaultClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupsUnderBackupVaultServerTransport(srv *BackupsUnderBackupVaultServer) *BackupsUnderBackupVaultServerTransport {
	return &BackupsUnderBackupVaultServerTransport{
		srv:               srv,
		beginRestoreFiles: newTracker[azfake.PollerResponder[armnetapp.BackupsUnderBackupVaultClientRestoreFilesResponse]](),
	}
}

// BackupsUnderBackupVaultServerTransport connects instances of armnetapp.BackupsUnderBackupVaultClient to instances of BackupsUnderBackupVaultServer.
// Don't use this type directly, use NewBackupsUnderBackupVaultServerTransport instead.
type BackupsUnderBackupVaultServerTransport struct {
	srv               *BackupsUnderBackupVaultServer
	beginRestoreFiles *tracker[azfake.PollerResponder[armnetapp.BackupsUnderBackupVaultClientRestoreFilesResponse]]
}

// Do implements the policy.Transporter interface for BackupsUnderBackupVaultServerTransport.
func (b *BackupsUnderBackupVaultServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "BackupsUnderBackupVaultClient.BeginRestoreFiles":
		resp, err = b.dispatchBeginRestoreFiles(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *BackupsUnderBackupVaultServerTransport) dispatchBeginRestoreFiles(req *http.Request) (*http.Response, error) {
	if b.srv.BeginRestoreFiles == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRestoreFiles not implemented")}
	}
	beginRestoreFiles := b.beginRestoreFiles.get(req)
	if beginRestoreFiles == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetApp/netAppAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupVaults/(?P<backupVaultName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backups/(?P<backupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restoreFiles`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetapp.BackupRestoreFiles](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		backupVaultNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupVaultName")])
		if err != nil {
			return nil, err
		}
		backupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("backupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := b.srv.BeginRestoreFiles(req.Context(), resourceGroupNameParam, accountNameParam, backupVaultNameParam, backupNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRestoreFiles = &respr
		b.beginRestoreFiles.add(req, beginRestoreFiles)
	}

	resp, err := server.PollerResponderNext(beginRestoreFiles, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusAccepted}, resp.StatusCode) {
		b.beginRestoreFiles.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRestoreFiles) {
		b.beginRestoreFiles.remove(req)
	}

	return resp, nil
}
