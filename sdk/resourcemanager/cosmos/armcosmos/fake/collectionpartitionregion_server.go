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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v4"
	"net/http"
	"net/url"
	"regexp"
)

// CollectionPartitionRegionServer is a fake server for instances of the armcosmos.CollectionPartitionRegionClient type.
type CollectionPartitionRegionServer struct {
	// NewListMetricsPager is the fake for method CollectionPartitionRegionClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(resourceGroupName string, accountName string, region string, databaseRid string, collectionRid string, filter string, options *armcosmos.CollectionPartitionRegionClientListMetricsOptions) (resp azfake.PagerResponder[armcosmos.CollectionPartitionRegionClientListMetricsResponse])
}

// NewCollectionPartitionRegionServerTransport creates a new instance of CollectionPartitionRegionServerTransport with the provided implementation.
// The returned CollectionPartitionRegionServerTransport instance is connected to an instance of armcosmos.CollectionPartitionRegionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCollectionPartitionRegionServerTransport(srv *CollectionPartitionRegionServer) *CollectionPartitionRegionServerTransport {
	return &CollectionPartitionRegionServerTransport{
		srv:                 srv,
		newListMetricsPager: newTracker[azfake.PagerResponder[armcosmos.CollectionPartitionRegionClientListMetricsResponse]](),
	}
}

// CollectionPartitionRegionServerTransport connects instances of armcosmos.CollectionPartitionRegionClient to instances of CollectionPartitionRegionServer.
// Don't use this type directly, use NewCollectionPartitionRegionServerTransport instead.
type CollectionPartitionRegionServerTransport struct {
	srv                 *CollectionPartitionRegionServer
	newListMetricsPager *tracker[azfake.PagerResponder[armcosmos.CollectionPartitionRegionClientListMetricsResponse]]
}

// Do implements the policy.Transporter interface for CollectionPartitionRegionServerTransport.
func (c *CollectionPartitionRegionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return c.dispatchToMethodFake(req, method)
}

func (c *CollectionPartitionRegionServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if collectionPartitionRegionServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = collectionPartitionRegionServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "CollectionPartitionRegionClient.NewListMetricsPager":
				res.resp, res.err = c.dispatchNewListMetricsPager(req)
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

func (c *CollectionPartitionRegionServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := c.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/region/(?P<region>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/partitions/metrics`
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
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
		if err != nil {
			return nil, err
		}
		regionParam, err := url.PathUnescape(matches[regex.SubexpIndex("region")])
		if err != nil {
			return nil, err
		}
		databaseRidParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseRid")])
		if err != nil {
			return nil, err
		}
		collectionRidParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionRid")])
		if err != nil {
			return nil, err
		}
		filterParam, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		resp := c.srv.NewListMetricsPager(resourceGroupNameParam, accountNameParam, regionParam, databaseRidParam, collectionRidParam, filterParam, nil)
		newListMetricsPager = &resp
		c.newListMetricsPager.add(req, newListMetricsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListMetricsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricsPager) {
		c.newListMetricsPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to CollectionPartitionRegionServerTransport
var collectionPartitionRegionServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
