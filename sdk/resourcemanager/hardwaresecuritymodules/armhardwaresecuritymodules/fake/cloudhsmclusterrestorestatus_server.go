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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hardwaresecuritymodules/armhardwaresecuritymodules/v2"
	"net/http"
	"net/url"
	"regexp"
)

// CloudHsmClusterRestoreStatusServer is a fake server for instances of the armhardwaresecuritymodules.CloudHsmClusterRestoreStatusClient type.
type CloudHsmClusterRestoreStatusServer struct {
	// Get is the fake for method CloudHsmClusterRestoreStatusClient.Get
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	Get func(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, jobID string, options *armhardwaresecuritymodules.CloudHsmClusterRestoreStatusClientGetOptions) (resp azfake.Responder[armhardwaresecuritymodules.CloudHsmClusterRestoreStatusClientGetResponse], errResp azfake.ErrorResponder)
}

// NewCloudHsmClusterRestoreStatusServerTransport creates a new instance of CloudHsmClusterRestoreStatusServerTransport with the provided implementation.
// The returned CloudHsmClusterRestoreStatusServerTransport instance is connected to an instance of armhardwaresecuritymodules.CloudHsmClusterRestoreStatusClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCloudHsmClusterRestoreStatusServerTransport(srv *CloudHsmClusterRestoreStatusServer) *CloudHsmClusterRestoreStatusServerTransport {
	return &CloudHsmClusterRestoreStatusServerTransport{srv: srv}
}

// CloudHsmClusterRestoreStatusServerTransport connects instances of armhardwaresecuritymodules.CloudHsmClusterRestoreStatusClient to instances of CloudHsmClusterRestoreStatusServer.
// Don't use this type directly, use NewCloudHsmClusterRestoreStatusServerTransport instead.
type CloudHsmClusterRestoreStatusServerTransport struct {
	srv *CloudHsmClusterRestoreStatusServer
}

// Do implements the policy.Transporter interface for CloudHsmClusterRestoreStatusServerTransport.
func (c *CloudHsmClusterRestoreStatusServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CloudHsmClusterRestoreStatusServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if cloudHsmClusterRestoreStatusServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = cloudHsmClusterRestoreStatusServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CloudHsmClusterRestoreStatusClient.Get":
				res.resp, res.err = c.dispatchGet(req)
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

func (c *CloudHsmClusterRestoreStatusServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if c.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.HardwareSecurityModules/cloudHsmClusters/(?P<cloudHsmClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restoreOperationStatus/(?P<jobId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	cloudHsmClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("cloudHsmClusterName")])
	if err != nil {
		return nil, err
	}
	jobIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("jobId")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Get(req.Context(), resourceGroupNameParam, cloudHsmClusterNameParam, jobIDParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusAccepted}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).RestoreResult, req)
	if err != nil {
		return nil, err
	}
	if val := server.GetResponse(respr).Location; val != nil {
		resp.Header.Set("Location", *val)
	}
	if val := server.GetResponse(respr).XMSRequestID; val != nil {
		resp.Header.Set("x-ms-request-id", *val)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CloudHsmClusterRestoreStatusServerTransport
var cloudHsmClusterRestoreStatusServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
