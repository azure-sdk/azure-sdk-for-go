// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery/v2"
	"net/http"
	"net/url"
	"regexp"
)

// ClusterRecoveryPointsServer is a fake server for instances of the armrecoveryservicessiterecovery.ClusterRecoveryPointsClient type.
type ClusterRecoveryPointsServer struct {
	// NewListByReplicationProtectionClusterPager is the fake for method ClusterRecoveryPointsClient.NewListByReplicationProtectionClusterPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByReplicationProtectionClusterPager func(resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicationProtectionClusterName string, options *armrecoveryservicessiterecovery.ClusterRecoveryPointsClientListByReplicationProtectionClusterOptions) (resp azfake.PagerResponder[armrecoveryservicessiterecovery.ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse])
}

// NewClusterRecoveryPointsServerTransport creates a new instance of ClusterRecoveryPointsServerTransport with the provided implementation.
// The returned ClusterRecoveryPointsServerTransport instance is connected to an instance of armrecoveryservicessiterecovery.ClusterRecoveryPointsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewClusterRecoveryPointsServerTransport(srv *ClusterRecoveryPointsServer) *ClusterRecoveryPointsServerTransport {
	return &ClusterRecoveryPointsServerTransport{
		srv: srv,
		newListByReplicationProtectionClusterPager: newTracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse]](),
	}
}

// ClusterRecoveryPointsServerTransport connects instances of armrecoveryservicessiterecovery.ClusterRecoveryPointsClient to instances of ClusterRecoveryPointsServer.
// Don't use this type directly, use NewClusterRecoveryPointsServerTransport instead.
type ClusterRecoveryPointsServerTransport struct {
	srv                                        *ClusterRecoveryPointsServer
	newListByReplicationProtectionClusterPager *tracker[azfake.PagerResponder[armrecoveryservicessiterecovery.ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse]]
}

// Do implements the policy.Transporter interface for ClusterRecoveryPointsServerTransport.
func (c *ClusterRecoveryPointsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *ClusterRecoveryPointsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if clusterRecoveryPointsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = clusterRecoveryPointsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "ClusterRecoveryPointsClient.NewListByReplicationProtectionClusterPager":
				res.resp, res.err = c.dispatchNewListByReplicationProtectionClusterPager(req)
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

func (c *ClusterRecoveryPointsServerTransport) dispatchNewListByReplicationProtectionClusterPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListByReplicationProtectionClusterPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByReplicationProtectionClusterPager not implemented")}
	}
	newListByReplicationProtectionClusterPager := c.newListByReplicationProtectionClusterPager.get(req)
	if newListByReplicationProtectionClusterPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.RecoveryServices/vaults/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationFabrics/(?P<fabricName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionContainers/(?P<protectionContainerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/replicationProtectionClusters/(?P<replicationProtectionClusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/recoveryPoints`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 6 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		fabricNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fabricName")])
		if err != nil {
			return nil, err
		}
		protectionContainerNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("protectionContainerName")])
		if err != nil {
			return nil, err
		}
		replicationProtectionClusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("replicationProtectionClusterName")])
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListByReplicationProtectionClusterPager(resourceGroupNameParam, resourceNameParam, fabricNameParam, protectionContainerNameParam, replicationProtectionClusterNameParam, nil)
		newListByReplicationProtectionClusterPager = &resp
		c.newListByReplicationProtectionClusterPager.add(req, newListByReplicationProtectionClusterPager)
		server.PagerResponderInjectNextLinks(newListByReplicationProtectionClusterPager, req, func(page *armrecoveryservicessiterecovery.ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByReplicationProtectionClusterPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListByReplicationProtectionClusterPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByReplicationProtectionClusterPager) {
		c.newListByReplicationProtectionClusterPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to ClusterRecoveryPointsServerTransport
var clusterRecoveryPointsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
