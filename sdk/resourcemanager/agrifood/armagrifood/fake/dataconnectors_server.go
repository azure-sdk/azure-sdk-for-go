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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DataConnectorsServer is a fake server for instances of the armagrifood.DataConnectorsClient type.
type DataConnectorsServer struct {
	// CreateOrUpdate is the fake for method DataConnectorsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, dataConnectorName string, body armagrifood.DataConnector, options *armagrifood.DataConnectorsClientCreateOrUpdateOptions) (resp azfake.Responder[armagrifood.DataConnectorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DataConnectorsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, dataConnectorName string, options *armagrifood.DataConnectorsClientDeleteOptions) (resp azfake.Responder[armagrifood.DataConnectorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DataConnectorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, dataManagerForAgricultureResourceName string, dataConnectorName string, options *armagrifood.DataConnectorsClientGetOptions) (resp azfake.Responder[armagrifood.DataConnectorsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method DataConnectorsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, dataManagerForAgricultureResourceName string, options *armagrifood.DataConnectorsClientListOptions) (resp azfake.PagerResponder[armagrifood.DataConnectorsClientListResponse])
}

// NewDataConnectorsServerTransport creates a new instance of DataConnectorsServerTransport with the provided implementation.
// The returned DataConnectorsServerTransport instance is connected to an instance of armagrifood.DataConnectorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDataConnectorsServerTransport(srv *DataConnectorsServer) *DataConnectorsServerTransport {
	return &DataConnectorsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armagrifood.DataConnectorsClientListResponse]](),
	}
}

// DataConnectorsServerTransport connects instances of armagrifood.DataConnectorsClient to instances of DataConnectorsServer.
// Don't use this type directly, use NewDataConnectorsServerTransport instead.
type DataConnectorsServerTransport struct {
	srv          *DataConnectorsServer
	newListPager *tracker[azfake.PagerResponder[armagrifood.DataConnectorsClientListResponse]]
}

// Do implements the policy.Transporter interface for DataConnectorsServerTransport.
func (d *DataConnectorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DataConnectorsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if dataConnectorsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = dataConnectorsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DataConnectorsClient.CreateOrUpdate":
				res.resp, res.err = d.dispatchCreateOrUpdate(req)
			case "DataConnectorsClient.Delete":
				res.resp, res.err = d.dispatchDelete(req)
			case "DataConnectorsClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DataConnectorsClient.NewListPager":
				res.resp, res.err = d.dispatchNewListPager(req)
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

func (d *DataConnectorsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<dataManagerForAgricultureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataConnectors/(?P<dataConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armagrifood.DataConnector](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataManagerForAgricultureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataManagerForAgricultureResourceName")])
	if err != nil {
		return nil, err
	}
	dataConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataConnectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceGroupNameParam, dataManagerForAgricultureResourceNameParam, dataConnectorNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataConnector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataConnectorsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<dataManagerForAgricultureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataConnectors/(?P<dataConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataManagerForAgricultureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataManagerForAgricultureResourceName")])
	if err != nil {
		return nil, err
	}
	dataConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataConnectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, dataManagerForAgricultureResourceNameParam, dataConnectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataConnectorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<dataManagerForAgricultureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataConnectors/(?P<dataConnectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	dataManagerForAgricultureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataManagerForAgricultureResourceName")])
	if err != nil {
		return nil, err
	}
	dataConnectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataConnectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, dataManagerForAgricultureResourceNameParam, dataConnectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataConnector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataConnectorsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := d.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.AgFoodPlatform/farmBeats/(?P<dataManagerForAgricultureResourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataConnectors`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		dataManagerForAgricultureResourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataManagerForAgricultureResourceName")])
		if err != nil {
			return nil, err
		}
		maxPageSizeUnescaped, err := url.QueryUnescape(qp.Get("$maxPageSize"))
		if err != nil {
			return nil, err
		}
		maxPageSizeParam, err := parseOptional(maxPageSizeUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armagrifood.DataConnectorsClientListOptions
		if maxPageSizeParam != nil || skipTokenParam != nil {
			options = &armagrifood.DataConnectorsClientListOptions{
				MaxPageSize: maxPageSizeParam,
				SkipToken:   skipTokenParam,
			}
		}
		resp := d.srv.NewListPager(resourceGroupNameParam, dataManagerForAgricultureResourceNameParam, options)
		newListPager = &resp
		d.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armagrifood.DataConnectorsClientListResponse, createLink func() string) {
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

// set this to conditionally intercept incoming requests to DataConnectorsServerTransport
var dataConnectorsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
