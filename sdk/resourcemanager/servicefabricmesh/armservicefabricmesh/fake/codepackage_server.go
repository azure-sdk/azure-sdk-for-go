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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabricmesh/armservicefabricmesh"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// CodePackageServer is a fake server for instances of the armservicefabricmesh.CodePackageClient type.
type CodePackageServer struct {
	// GetContainerLogs is the fake for method CodePackageClient.GetContainerLogs
	// HTTP status codes to indicate success: http.StatusOK
	GetContainerLogs func(ctx context.Context, resourceGroupName string, applicationResourceName string, serviceResourceName string, replicaName string, codePackageName string, options *armservicefabricmesh.CodePackageClientGetContainerLogsOptions) (resp azfake.Responder[armservicefabricmesh.CodePackageClientGetContainerLogsResponse], errResp azfake.ErrorResponder)
}

// NewCodePackageServerTransport creates a new instance of CodePackageServerTransport with the provided implementation.
// The returned CodePackageServerTransport instance is connected to an instance of armservicefabricmesh.CodePackageClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCodePackageServerTransport(srv *CodePackageServer) *CodePackageServerTransport {
	return &CodePackageServerTransport{srv: srv}
}

// CodePackageServerTransport connects instances of armservicefabricmesh.CodePackageClient to instances of CodePackageServer.
// Don't use this type directly, use NewCodePackageServerTransport instead.
type CodePackageServerTransport struct {
	srv *CodePackageServer
}

// Do implements the policy.Transporter interface for CodePackageServerTransport.
func (c *CodePackageServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CodePackageClient.GetContainerLogs":
		resp, err = c.dispatchGetContainerLogs(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CodePackageServerTransport) dispatchGetContainerLogs(req *http.Request) (*http.Response, error) {
	if c.srv.GetContainerLogs == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetContainerLogs not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.ServiceFabricMesh/applications/(?P<applicationResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/services/(?P<serviceResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas/(?P<replicaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/codePackages/(?P<codePackageName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/logs`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	applicationResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("applicationResourceName")])
	if err != nil {
		return nil, err
	}
	serviceResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("serviceResourceName")])
	if err != nil {
		return nil, err
	}
	replicaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicaName")])
	if err != nil {
		return nil, err
	}
	codePackageNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("codePackageName")])
	if err != nil {
		return nil, err
	}
	tailUnescaped, err := url.QueryUnescape(qp.Get("tail"))
	if err != nil {
		return nil, err
	}
	tailParam, err := parseOptional(tailUnescaped, func(v string) (int32, error) {
		p, parseErr := strconv.ParseInt(v, 10, 32)
		if parseErr != nil {
			return 0, parseErr
		}
		return int32(p), nil
	})
	if err != nil {
		return nil, err
	}
	var options *armservicefabricmesh.CodePackageClientGetContainerLogsOptions
	if tailParam != nil {
		options = &armservicefabricmesh.CodePackageClientGetContainerLogsOptions{
			Tail: tailParam,
		}
	}
	respr, errRespr := c.srv.GetContainerLogs(req.Context(), resourceGroupNameParam, applicationResourceNameParam, serviceResourceNameParam, replicaNameParam, codePackageNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ContainerLogs, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}