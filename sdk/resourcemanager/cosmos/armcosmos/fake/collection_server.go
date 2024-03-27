//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"net/http"
	"net/url"
	"regexp"
)

// CollectionServer is a fake server for instances of the armcosmos.CollectionClient type.
type CollectionServer struct {
	// NewListMetricDefinitionsPager is the fake for method CollectionClient.NewListMetricDefinitionsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricDefinitionsPager func(resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *armcosmos.CollectionClientListMetricDefinitionsOptions) (resp azfake.PagerResponder[armcosmos.CollectionClientListMetricDefinitionsResponse])

	// NewListMetricsPager is the fake for method CollectionClient.NewListMetricsPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListMetricsPager func(resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *armcosmos.CollectionClientListMetricsOptions) (resp azfake.PagerResponder[armcosmos.CollectionClientListMetricsResponse])

	// NewListUsagesPager is the fake for method CollectionClient.NewListUsagesPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListUsagesPager func(resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *armcosmos.CollectionClientListUsagesOptions) (resp azfake.PagerResponder[armcosmos.CollectionClientListUsagesResponse])
}

// NewCollectionServerTransport creates a new instance of CollectionServerTransport with the provided implementation.
// The returned CollectionServerTransport instance is connected to an instance of armcosmos.CollectionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewCollectionServerTransport(srv *CollectionServer) *CollectionServerTransport {
	return &CollectionServerTransport{
		srv:                           srv,
		newListMetricDefinitionsPager: newTracker[azfake.PagerResponder[armcosmos.CollectionClientListMetricDefinitionsResponse]](),
		newListMetricsPager:           newTracker[azfake.PagerResponder[armcosmos.CollectionClientListMetricsResponse]](),
		newListUsagesPager:            newTracker[azfake.PagerResponder[armcosmos.CollectionClientListUsagesResponse]](),
	}
}

// CollectionServerTransport connects instances of armcosmos.CollectionClient to instances of CollectionServer.
// Don't use this type directly, use NewCollectionServerTransport instead.
type CollectionServerTransport struct {
	srv                           *CollectionServer
	newListMetricDefinitionsPager *tracker[azfake.PagerResponder[armcosmos.CollectionClientListMetricDefinitionsResponse]]
	newListMetricsPager           *tracker[azfake.PagerResponder[armcosmos.CollectionClientListMetricsResponse]]
	newListUsagesPager            *tracker[azfake.PagerResponder[armcosmos.CollectionClientListUsagesResponse]]
}

// Do implements the policy.Transporter interface for CollectionServerTransport.
func (c *CollectionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "CollectionClient.NewListMetricDefinitionsPager":
		resp, err = c.dispatchNewListMetricDefinitionsPager(req)
	case "CollectionClient.NewListMetricsPager":
		resp, err = c.dispatchNewListMetricsPager(req)
	case "CollectionClient.NewListUsagesPager":
		resp, err = c.dispatchNewListUsagesPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *CollectionServerTransport) dispatchNewListMetricDefinitionsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListMetricDefinitionsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricDefinitionsPager not implemented")}
	}
	newListMetricDefinitionsPager := c.newListMetricDefinitionsPager.get(req)
	if newListMetricDefinitionsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metricDefinitions`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		accountNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("accountName")])
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
		resp := c.srv.NewListMetricDefinitionsPager(resourceGroupNameParam, accountNameParam, databaseRidParam, collectionRidParam, nil)
		newListMetricDefinitionsPager = &resp
		c.newListMetricDefinitionsPager.add(req, newListMetricDefinitionsPager)
	}
	resp, err := server.PagerResponderNext(newListMetricDefinitionsPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListMetricDefinitionsPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListMetricDefinitionsPager) {
		c.newListMetricDefinitionsPager.remove(req)
	}
	return resp, nil
}

func (c *CollectionServerTransport) dispatchNewListMetricsPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListMetricsPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListMetricsPager not implemented")}
	}
	newListMetricsPager := c.newListMetricsPager.get(req)
	if newListMetricsPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/metrics`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		resp := c.srv.NewListMetricsPager(resourceGroupNameParam, accountNameParam, databaseRidParam, collectionRidParam, filterParam, nil)
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

func (c *CollectionServerTransport) dispatchNewListUsagesPager(req *http.Request) (*http.Response, error) {
	if c.srv.NewListUsagesPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListUsagesPager not implemented")}
	}
	newListUsagesPager := c.newListUsagesPager.get(req)
	if newListUsagesPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/databaseAccounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/collections/(?P<collectionRid>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/usages`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
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
		databaseRidParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseRid")])
		if err != nil {
			return nil, err
		}
		collectionRidParam, err := url.PathUnescape(matches[regex.SubexpIndex("collectionRid")])
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armcosmos.CollectionClientListUsagesOptions
		if filterParam != nil {
			options = &armcosmos.CollectionClientListUsagesOptions{
				Filter: filterParam,
			}
		}
		resp := c.srv.NewListUsagesPager(resourceGroupNameParam, accountNameParam, databaseRidParam, collectionRidParam, options)
		newListUsagesPager = &resp
		c.newListUsagesPager.add(req, newListUsagesPager)
	}
	resp, err := server.PagerResponderNext(newListUsagesPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		c.newListUsagesPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListUsagesPager) {
		c.newListUsagesPager.remove(req)
	}
	return resp, nil
}
