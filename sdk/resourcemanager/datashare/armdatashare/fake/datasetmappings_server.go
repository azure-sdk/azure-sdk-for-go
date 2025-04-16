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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datashare/armdatashare"
	"net/http"
	"net/url"
	"regexp"
)

// DataSetMappingsServer is a fake server for instances of the armdatashare.DataSetMappingsClient type.
type DataSetMappingsServer struct {
	// Create is the fake for method DataSetMappingsClient.Create
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	Create func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, dataSetMapping armdatashare.DataSetMappingClassification, options *armdatashare.DataSetMappingsClientCreateOptions) (resp azfake.Responder[armdatashare.DataSetMappingsClientCreateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DataSetMappingsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, options *armdatashare.DataSetMappingsClientDeleteOptions) (resp azfake.Responder[armdatashare.DataSetMappingsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DataSetMappingsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, accountName string, shareSubscriptionName string, dataSetMappingName string, options *armdatashare.DataSetMappingsClientGetOptions) (resp azfake.Responder[armdatashare.DataSetMappingsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByShareSubscriptionPager is the fake for method DataSetMappingsClient.NewListByShareSubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByShareSubscriptionPager func(resourceGroupName string, accountName string, shareSubscriptionName string, options *armdatashare.DataSetMappingsClientListByShareSubscriptionOptions) (resp azfake.PagerResponder[armdatashare.DataSetMappingsClientListByShareSubscriptionResponse])
}

// NewDataSetMappingsServerTransport creates a new instance of DataSetMappingsServerTransport with the provided implementation.
// The returned DataSetMappingsServerTransport instance is connected to an instance of armdatashare.DataSetMappingsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDataSetMappingsServerTransport(srv *DataSetMappingsServer) *DataSetMappingsServerTransport {
	return &DataSetMappingsServerTransport{
		srv:                             srv,
		newListByShareSubscriptionPager: newTracker[azfake.PagerResponder[armdatashare.DataSetMappingsClientListByShareSubscriptionResponse]](),
	}
}

// DataSetMappingsServerTransport connects instances of armdatashare.DataSetMappingsClient to instances of DataSetMappingsServer.
// Don't use this type directly, use NewDataSetMappingsServerTransport instead.
type DataSetMappingsServerTransport struct {
	srv                             *DataSetMappingsServer
	newListByShareSubscriptionPager *tracker[azfake.PagerResponder[armdatashare.DataSetMappingsClientListByShareSubscriptionResponse]]
}

// Do implements the policy.Transporter interface for DataSetMappingsServerTransport.
func (d *DataSetMappingsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DataSetMappingsServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if dataSetMappingsServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = dataSetMappingsServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DataSetMappingsClient.Create":
				res.resp, res.err = d.dispatchCreate(req)
			case "DataSetMappingsClient.Delete":
				res.resp, res.err = d.dispatchDelete(req)
			case "DataSetMappingsClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DataSetMappingsClient.NewListByShareSubscriptionPager":
				res.resp, res.err = d.dispatchNewListByShareSubscriptionPager(req)
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

func (d *DataSetMappingsServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if d.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataSetMappings/(?P<dataSetMappingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	raw, err := readRequestBody(req)
	if err != nil {
		return nil, err
	}
	body, err := unmarshalDataSetMappingClassification(raw)
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
	shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
	if err != nil {
		return nil, err
	}
	dataSetMappingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataSetMappingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Create(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, dataSetMappingNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataSetMappingClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataSetMappingsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataSetMappings/(?P<dataSetMappingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
	if err != nil {
		return nil, err
	}
	dataSetMappingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataSetMappingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, dataSetMappingNameParam, nil)
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

func (d *DataSetMappingsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataSetMappings/(?P<dataSetMappingName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
	if err != nil {
		return nil, err
	}
	dataSetMappingNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dataSetMappingName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, dataSetMappingNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DataSetMappingClassification, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DataSetMappingsServerTransport) dispatchNewListByShareSubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByShareSubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByShareSubscriptionPager not implemented")}
	}
	newListByShareSubscriptionPager := d.newListByShareSubscriptionPager.get(req)
	if newListByShareSubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataShare/accounts/(?P<accountName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/shareSubscriptions/(?P<shareSubscriptionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dataSetMappings`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
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
		shareSubscriptionNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("shareSubscriptionName")])
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
		if err != nil {
			return nil, err
		}
		orderbyParam := getOptional(orderbyUnescaped)
		var options *armdatashare.DataSetMappingsClientListByShareSubscriptionOptions
		if skipTokenParam != nil || filterParam != nil || orderbyParam != nil {
			options = &armdatashare.DataSetMappingsClientListByShareSubscriptionOptions{
				SkipToken: skipTokenParam,
				Filter:    filterParam,
				Orderby:   orderbyParam,
			}
		}
		resp := d.srv.NewListByShareSubscriptionPager(resourceGroupNameParam, accountNameParam, shareSubscriptionNameParam, options)
		newListByShareSubscriptionPager = &resp
		d.newListByShareSubscriptionPager.add(req, newListByShareSubscriptionPager)
		server.PagerResponderInjectNextLinks(newListByShareSubscriptionPager, req, func(page *armdatashare.DataSetMappingsClientListByShareSubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByShareSubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByShareSubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByShareSubscriptionPager) {
		d.newListByShareSubscriptionPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DataSetMappingsServerTransport
var dataSetMappingsServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
