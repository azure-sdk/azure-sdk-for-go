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

// BackupStatusServer is a fake server for instances of the armrecoveryservicesbackup.BackupStatusClient type.
type BackupStatusServer struct {
	// Get is the fake for method BackupStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, azureRegion string, parameters armrecoveryservicesbackup.BackupStatusRequest, options *armrecoveryservicesbackup.BackupStatusClientGetOptions) (resp azfake.Responder[armrecoveryservicesbackup.BackupStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewBackupStatusServerTransport creates a new instance of BackupStatusServerTransport with the provided implementation.
// The returned BackupStatusServerTransport instance is connected to an instance of armrecoveryservicesbackup.BackupStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewBackupStatusServerTransport(srv *BackupStatusServer) *BackupStatusServerTransport {
	return &BackupStatusServerTransport{srv: srv}
}

// BackupStatusServerTransport connects instances of armrecoveryservicesbackup.BackupStatusClient to instances of BackupStatusServer.
// Don't use this type directly, use NewBackupStatusServerTransport instead.
type BackupStatusServerTransport struct {
	srv *BackupStatusServer
}

// Do implements the policy.Transporter interface for BackupStatusServerTransport.
func (b *BackupStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return b.dispatchToMethodFake(req, method)
}

func (b *BackupStatusServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if backupStatusServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = backupStatusServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "BackupStatusClient.Get":
				res.resp, res.err = b.dispatchGet(req)
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

func (b *BackupStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if b.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/Subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/locations/(?P<azureRegion>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/backupStatus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armrecoveryservicesbackup.BackupStatusRequest](req)
	if err != nil {
		return nil, err
	}
	azureRegionParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureRegion")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := b.srv.Get(req.Context(), azureRegionParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).BackupStatusResponse, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to BackupStatusServerTransport
var backupStatusServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
