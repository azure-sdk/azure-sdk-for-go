// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/avs/armavs/v2"
	"net/http"
	"net/url"
	"regexp"
)

// DatastoresServer is a fake server for instances of the armavs.DatastoresClient type.
type DatastoresServer struct {
	// BeginCreateOrUpdate is the fake for method DatastoresClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, datastoreName string, datastore armavs.Datastore, options *armavs.DatastoresClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armavs.DatastoresClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DatastoresClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, datastoreName string, options *armavs.DatastoresClientBeginDeleteOptions) (resp azfake.PollerResponder[armavs.DatastoresClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DatastoresClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, privateCloudName string, clusterName string, datastoreName string, options *armavs.DatastoresClientGetOptions) (resp azfake.Responder[armavs.DatastoresClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DatastoresClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, privateCloudName string, clusterName string, options *armavs.DatastoresClientListOptions) (resp azfake.PagerResponder[armavs.DatastoresClientListResponse])
}

// NewDatastoresServerTransport creates a new instance of DatastoresServerTransport with the provided implementation.
// The returned DatastoresServerTransport instance is connected to an instance of armavs.DatastoresClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatastoresServerTransport(srv *DatastoresServer) *DatastoresServerTransport {
	return &DatastoresServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armavs.DatastoresClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armavs.DatastoresClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armavs.DatastoresClientListResponse]](),
	}
}

// DatastoresServerTransport connects instances of armavs.DatastoresClient to instances of DatastoresServer.
// Don't use this type directly, use NewDatastoresServerTransport instead.
type DatastoresServerTransport struct {
	srv                 *DatastoresServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armavs.DatastoresClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armavs.DatastoresClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armavs.DatastoresClientListResponse]]
}

// Do implements the policy.Transporter interface for DatastoresServerTransport.
func (d *DatastoresServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatastoresServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch method {
	case "DatastoresClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DatastoresClient.BeginDelete":
		resp, err = d.dispatchBeginDelete(req)
	case "DatastoresClient.Get":
		resp, err = d.dispatchGet(req)
	case "DatastoresClient.NewListPager":
		resp, err = d.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	return resp, err
}

func (d *DatastoresServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<datastoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armavs.Datastore](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		datastoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("datastoreName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, datastoreNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		d.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		d.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		d.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (d *DatastoresServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<datastoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		datastoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("datastoreName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, datastoreNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		d.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		d.beginDelete.remove(req)
	}

	return resp, nil
}

func (d *DatastoresServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores/(?P<datastoreName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
	if err != nil {
		return nil, err
	}
	clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
	if err != nil {
		return nil, err
	}
	datastoreNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("datastoreName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, privateCloudNameParam, clusterNameParam, datastoreNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Datastore, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DatastoresServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AVS/privateClouds/(?P<privateCloudName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/clusters/(?P<clusterName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/datastores`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		privateCloudNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("privateCloudName")])
		if err != nil {
			return nil, err
		}
		clusterNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("clusterName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, privateCloudNameParam, clusterNameParam, nil)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armavs.DatastoresClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		d.newListPager.remove(req)
	}
	return resp, nil
}
