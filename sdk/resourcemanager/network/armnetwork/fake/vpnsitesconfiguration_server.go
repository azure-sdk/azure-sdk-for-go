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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	"net/http"
	"net/url"
	"regexp"
)

// VPNSitesConfigurationServer is a fake server for instances of the armnetwork.VPNSitesConfigurationClient type.
type VPNSitesConfigurationServer struct {
	// BeginDownload is the fake for method VPNSitesConfigurationClient.BeginDownload
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDownload func(ctx context.Context, resourceGroupName string, virtualWANName string, request armnetwork.GetVPNSitesConfigurationRequest, options *armnetwork.VPNSitesConfigurationClientBeginDownloadOptions) (resp azfake.PollerResponder[armnetwork.VPNSitesConfigurationClientDownloadResponse], errResp azfake.ErrorResponder)
}

// NewVPNSitesConfigurationServerTransport creates a new instance of VPNSitesConfigurationServerTransport with the provided implementation.
// The returned VPNSitesConfigurationServerTransport instance is connected to an instance of armnetwork.VPNSitesConfigurationClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVPNSitesConfigurationServerTransport(srv *VPNSitesConfigurationServer) *VPNSitesConfigurationServerTransport {
	return &VPNSitesConfigurationServerTransport{
		srv:           srv,
		beginDownload: newTracker[azfake.PollerResponder[armnetwork.VPNSitesConfigurationClientDownloadResponse]](),
	}
}

// VPNSitesConfigurationServerTransport connects instances of armnetwork.VPNSitesConfigurationClient to instances of VPNSitesConfigurationServer.
// Don't use this type directly, use NewVPNSitesConfigurationServerTransport instead.
type VPNSitesConfigurationServerTransport struct {
	srv           *VPNSitesConfigurationServer
	beginDownload *tracker[azfake.PollerResponder[armnetwork.VPNSitesConfigurationClientDownloadResponse]]
}

// Do implements the policy.Transporter interface for VPNSitesConfigurationServerTransport.
func (v *VPNSitesConfigurationServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return v.dispatchToMethodFake(req, method)
}

func (v *VPNSitesConfigurationServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if vpnSitesConfigurationServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = vpnSitesConfigurationServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "VPNSitesConfigurationClient.BeginDownload":
				res.resp, res.err = v.dispatchBeginDownload(req)
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

func (v *VPNSitesConfigurationServerTransport) dispatchBeginDownload(req *http.Request) (*http.Response, error) {
	if v.srv.BeginDownload == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDownload not implemented")}
	}
	beginDownload := v.beginDownload.get(req)
	if beginDownload == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Network/virtualWans/(?P<virtualWANName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/vpnConfiguration`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.GetVPNSitesConfigurationRequest](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		virtualWANNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("virtualWANName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := v.srv.BeginDownload(req.Context(), resourceGroupNameParam, virtualWANNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDownload = &respr
		v.beginDownload.add(req, beginDownload)
	}

	resp, err := server.PollerResponderNext(beginDownload, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		v.beginDownload.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDownload) {
		v.beginDownload.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to VPNSitesConfigurationServerTransport
var vpnSitesConfigurationServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
