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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ContainerAppsRevisionReplicasServer is a fake server for instances of the armappcontainers.ContainerAppsRevisionReplicasClient type.
type ContainerAppsRevisionReplicasServer struct {
	// GetReplica is the fake for method ContainerAppsRevisionReplicasClient.GetReplica
	// HTTP status codes to indicate success: http.StatusOK
	GetReplica func(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, replicaName string, options *armappcontainers.ContainerAppsRevisionReplicasClientGetReplicaOptions) (resp azfake.Responder[armappcontainers.ContainerAppsRevisionReplicasClientGetReplicaResponse], errResp azfake.ErrorResponder)

	// ListReplicas is the fake for method ContainerAppsRevisionReplicasClient.ListReplicas
	// HTTP status codes to indicate success: http.StatusOK
	ListReplicas func(ctx context.Context, resourceGroupName string, containerAppName string, revisionName string, options *armappcontainers.ContainerAppsRevisionReplicasClientListReplicasOptions) (resp azfake.Responder[armappcontainers.ContainerAppsRevisionReplicasClientListReplicasResponse], errResp azfake.ErrorResponder)
}

// NewContainerAppsRevisionReplicasServerTransport creates a new instance of ContainerAppsRevisionReplicasServerTransport with the provided implementation.
// The returned ContainerAppsRevisionReplicasServerTransport instance is connected to an instance of armappcontainers.ContainerAppsRevisionReplicasClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewContainerAppsRevisionReplicasServerTransport(srv *ContainerAppsRevisionReplicasServer) *ContainerAppsRevisionReplicasServerTransport {
	return &ContainerAppsRevisionReplicasServerTransport{srv: srv}
}

// ContainerAppsRevisionReplicasServerTransport connects instances of armappcontainers.ContainerAppsRevisionReplicasClient to instances of ContainerAppsRevisionReplicasServer.
// Don't use this type directly, use NewContainerAppsRevisionReplicasServerTransport instead.
type ContainerAppsRevisionReplicasServerTransport struct {
	srv *ContainerAppsRevisionReplicasServer
}

// Do implements the policy.Transporter interface for ContainerAppsRevisionReplicasServerTransport.
func (c *ContainerAppsRevisionReplicasServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ContainerAppsRevisionReplicasServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if containerAppsRevisionReplicasServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = containerAppsRevisionReplicasServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ContainerAppsRevisionReplicasClient.GetReplica":
				res.resp, res.err = c.dispatchGetReplica(req)
			case "ContainerAppsRevisionReplicasClient.ListReplicas":
				res.resp, res.err = c.dispatchListReplicas(req)
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

func (c *ContainerAppsRevisionReplicasServerTransport) dispatchGetReplica(req *http.Request) (*http.Response, error) {
	if c.srv.GetReplica == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetReplica not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<revisionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas/(?P<replicaName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	revisionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("revisionName")])
	if err != nil {
		return nil, err
	}
	replicaNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicaName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetReplica(req.Context(), resourceGroupNameParam, containerAppNameParam, revisionNameParam, replicaNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Replica, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ContainerAppsRevisionReplicasServerTransport) dispatchListReplicas(req *http.Request) (*http.Response, error) {
	if c.srv.ListReplicas == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListReplicas not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.App/containerApps/(?P<containerAppName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revisions/(?P<revisionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicas`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	containerAppNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("containerAppName")])
	if err != nil {
		return nil, err
	}
	revisionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("revisionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.ListReplicas(req.Context(), resourceGroupNameParam, containerAppNameParam, revisionNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ReplicaCollection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ContainerAppsRevisionReplicasServerTransport
var containerAppsRevisionReplicasServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
