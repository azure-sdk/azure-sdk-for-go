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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/oracledatabase/armoracledatabase"
	"net/http"
	"net/url"
	"regexp"
)

// DNSPrivateZonesServer is a fake server for instances of the armoracledatabase.DNSPrivateZonesClient type.
type DNSPrivateZonesServer struct {
	// Get is the fake for method DNSPrivateZonesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, dnsprivatezonename string, options *armoracledatabase.DNSPrivateZonesClientGetOptions) (resp azfake.Responder[armoracledatabase.DNSPrivateZonesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByLocationPager is the fake for method DNSPrivateZonesClient.NewListByLocationPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByLocationPager func(location string, options *armoracledatabase.DNSPrivateZonesClientListByLocationOptions) (resp azfake.PagerResponder[armoracledatabase.DNSPrivateZonesClientListByLocationResponse])
}

// NewDNSPrivateZonesServerTransport creates a new instance of DNSPrivateZonesServerTransport with the provided implementation.
// The returned DNSPrivateZonesServerTransport instance is connected to an instance of armoracledatabase.DNSPrivateZonesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDNSPrivateZonesServerTransport(srv *DNSPrivateZonesServer) *DNSPrivateZonesServerTransport {
	return &DNSPrivateZonesServerTransport{
		srv:                    srv,
		newListByLocationPager: newTracker[azfake.PagerResponder[armoracledatabase.DNSPrivateZonesClientListByLocationResponse]](),
	}
}

// DNSPrivateZonesServerTransport connects instances of armoracledatabase.DNSPrivateZonesClient to instances of DNSPrivateZonesServer.
// Don't use this type directly, use NewDNSPrivateZonesServerTransport instead.
type DNSPrivateZonesServerTransport struct {
	srv                    *DNSPrivateZonesServer
	newListByLocationPager *tracker[azfake.PagerResponder[armoracledatabase.DNSPrivateZonesClientListByLocationResponse]]
}

// Do implements the policy.Transporter interface for DNSPrivateZonesServerTransport.
func (d *DNSPrivateZonesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DNSPrivateZonesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if dnsPrivateZonesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = dnsPrivateZonesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DNSPrivateZonesClient.Get":
				res.resp, res.err = d.dispatchGet(req)
			case "DNSPrivateZonesClient.NewListByLocationPager":
				res.resp, res.err = d.dispatchNewListByLocationPager(req)
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

func (d *DNSPrivateZonesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dnsPrivateZones/(?P<dnsprivatezonename>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	dnsprivatezonenameParam, err := url.PathUnescape(matches[regex.SubexpIndex("dnsprivatezonename")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Get(req.Context(), locationParam, dnsprivatezonenameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DNSPrivateZone, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DNSPrivateZonesServerTransport) dispatchNewListByLocationPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByLocationPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByLocationPager not implemented")}
	}
	newListByLocationPager := d.newListByLocationPager.get(req)
	if newListByLocationPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Oracle\.Database/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/dnsPrivateZones`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByLocationPager(locationParam, nil)
		newListByLocationPager = &resp
		d.newListByLocationPager.add(req, newListByLocationPager)
		server.PagerResponderInjectNextLinks(newListByLocationPager, req, func(page *armoracledatabase.DNSPrivateZonesClientListByLocationResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByLocationPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		d.newListByLocationPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByLocationPager) {
		d.newListByLocationPager.remove(req)
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DNSPrivateZonesServerTransport
var dnsPrivateZonesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
