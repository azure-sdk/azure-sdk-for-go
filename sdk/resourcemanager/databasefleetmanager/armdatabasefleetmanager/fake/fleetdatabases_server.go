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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databasefleetmanager/armdatabasefleetmanager"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// FleetDatabasesServer is a fake server for instances of the armdatabasefleetmanager.FleetDatabasesClient type.
type FleetDatabasesServer struct {
	// BeginChangeTier is the fake for method FleetDatabasesClient.BeginChangeTier
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginChangeTier func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, body armdatabasefleetmanager.DatabaseChangeTierProperties, options *armdatabasefleetmanager.FleetDatabasesClientBeginChangeTierOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientChangeTierResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method FleetDatabasesClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, resource armdatabasefleetmanager.FleetDatabase, options *armdatabasefleetmanager.FleetDatabasesClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method FleetDatabasesClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, options *armdatabasefleetmanager.FleetDatabasesClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FleetDatabasesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, options *armdatabasefleetmanager.FleetDatabasesClientGetOptions) (resp azfake.Responder[armdatabasefleetmanager.FleetDatabasesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByFleetspacePager is the fake for method FleetDatabasesClient.NewListByFleetspacePager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByFleetspacePager func(resourceGroupName string, fleetName string, fleetspaceName string, options *armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceOptions) (resp azfake.PagerResponder[armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceResponse])

	// BeginRename is the fake for method FleetDatabasesClient.BeginRename
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginRename func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, body armdatabasefleetmanager.DatabaseRenameProperties, options *armdatabasefleetmanager.FleetDatabasesClientBeginRenameOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientRenameResponse], errResp azfake.ErrorResponder)

	// BeginRevert is the fake for method FleetDatabasesClient.BeginRevert
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginRevert func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, options *armdatabasefleetmanager.FleetDatabasesClientBeginRevertOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientRevertResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method FleetDatabasesClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, fleetName string, fleetspaceName string, databaseName string, properties armdatabasefleetmanager.FleetDatabase, options *armdatabasefleetmanager.FleetDatabasesClientBeginUpdateOptions) (resp azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewFleetDatabasesServerTransport creates a new instance of FleetDatabasesServerTransport with the provided implementation.
// The returned FleetDatabasesServerTransport instance is connected to an instance of armdatabasefleetmanager.FleetDatabasesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFleetDatabasesServerTransport(srv *FleetDatabasesServer) *FleetDatabasesServerTransport {
	return &FleetDatabasesServerTransport{
		srv:                      srv,
		beginChangeTier:          newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientChangeTierResponse]](),
		beginCreateOrUpdate:      newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientCreateOrUpdateResponse]](),
		beginDelete:              newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientDeleteResponse]](),
		newListByFleetspacePager: newTracker[azfake.PagerResponder[armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceResponse]](),
		beginRename:              newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientRenameResponse]](),
		beginRevert:              newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientRevertResponse]](),
		beginUpdate:              newTracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientUpdateResponse]](),
	}
}

// FleetDatabasesServerTransport connects instances of armdatabasefleetmanager.FleetDatabasesClient to instances of FleetDatabasesServer.
// Don't use this type directly, use NewFleetDatabasesServerTransport instead.
type FleetDatabasesServerTransport struct {
	srv                      *FleetDatabasesServer
	beginChangeTier          *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientChangeTierResponse]]
	beginCreateOrUpdate      *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientCreateOrUpdateResponse]]
	beginDelete              *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientDeleteResponse]]
	newListByFleetspacePager *tracker[azfake.PagerResponder[armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceResponse]]
	beginRename              *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientRenameResponse]]
	beginRevert              *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientRevertResponse]]
	beginUpdate              *tracker[azfake.PollerResponder[armdatabasefleetmanager.FleetDatabasesClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for FleetDatabasesServerTransport.
func (f *FleetDatabasesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return f.dispatchToMethodFake(req, method)
}

func (f *FleetDatabasesServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if fleetDatabasesServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = fleetDatabasesServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "FleetDatabasesClient.BeginChangeTier":
				res.resp, res.err = f.dispatchBeginChangeTier(req)
			case "FleetDatabasesClient.BeginCreateOrUpdate":
				res.resp, res.err = f.dispatchBeginCreateOrUpdate(req)
			case "FleetDatabasesClient.BeginDelete":
				res.resp, res.err = f.dispatchBeginDelete(req)
			case "FleetDatabasesClient.Get":
				res.resp, res.err = f.dispatchGet(req)
			case "FleetDatabasesClient.NewListByFleetspacePager":
				res.resp, res.err = f.dispatchNewListByFleetspacePager(req)
			case "FleetDatabasesClient.BeginRename":
				res.resp, res.err = f.dispatchBeginRename(req)
			case "FleetDatabasesClient.BeginRevert":
				res.resp, res.err = f.dispatchBeginRevert(req)
			case "FleetDatabasesClient.BeginUpdate":
				res.resp, res.err = f.dispatchBeginUpdate(req)
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

func (f *FleetDatabasesServerTransport) dispatchBeginChangeTier(req *http.Request) (*http.Response, error) {
	if f.srv.BeginChangeTier == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginChangeTier not implemented")}
	}
	beginChangeTier := f.beginChangeTier.get(req)
	if beginChangeTier == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/changeTier`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.DatabaseChangeTierProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginChangeTier(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginChangeTier = &respr
		f.beginChangeTier.add(req, beginChangeTier)
	}

	resp, err := server.PollerResponderNext(beginChangeTier, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginChangeTier.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginChangeTier) {
		f.beginChangeTier.remove(req)
	}

	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := f.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.FleetDatabase](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		f.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		f.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		f.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if f.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := f.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginDelete(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		f.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		f.beginDelete.remove(req)
	}

	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
	if err != nil {
		return nil, err
	}
	fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
	if err != nil {
		return nil, err
	}
	databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FleetDatabase, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchNewListByFleetspacePager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListByFleetspacePager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByFleetspacePager not implemented")}
	}
	newListByFleetspacePager := f.newListByFleetspacePager.get(req)
	if newListByFleetspacePager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases`
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
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		skipUnescaped, err := url.QueryUnescape(qp.Get("$skip"))
		if err != nil {
			return nil, err
		}
		skipParam, err := parseOptional(skipUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int64, error) {
			p, parseErr := strconv.ParseInt(v, 10, 64)
			if parseErr != nil {
				return 0, parseErr
			}
			return p, nil
		})
		if err != nil {
			return nil, err
		}
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		skiptokenUnescaped, err := url.QueryUnescape(qp.Get("$skiptoken"))
		if err != nil {
			return nil, err
		}
		skiptokenParam := getOptional(skiptokenUnescaped)
		var options *armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceOptions
		if skipParam != nil || topParam != nil || filterParam != nil || skiptokenParam != nil {
			options = &armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceOptions{
				Skip:      skipParam,
				Top:       topParam,
				Filter:    filterParam,
				Skiptoken: skiptokenParam,
			}
		}
		resp := f.srv.NewListByFleetspacePager(resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, options)
		newListByFleetspacePager = &resp
		f.newListByFleetspacePager.add(req, newListByFleetspacePager)
		server.PagerResponderInjectNextLinks(newListByFleetspacePager, req, func(page *armdatabasefleetmanager.FleetDatabasesClientListByFleetspaceResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByFleetspacePager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListByFleetspacePager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByFleetspacePager) {
		f.newListByFleetspacePager.remove(req)
	}
	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchBeginRename(req *http.Request) (*http.Response, error) {
	if f.srv.BeginRename == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRename not implemented")}
	}
	beginRename := f.beginRename.get(req)
	if beginRename == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/rename`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.DatabaseRenameProperties](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginRename(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRename = &respr
		f.beginRename.add(req, beginRename)
	}

	resp, err := server.PollerResponderNext(beginRename, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginRename.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRename) {
		f.beginRename.remove(req)
	}

	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchBeginRevert(req *http.Request) (*http.Response, error) {
	if f.srv.BeginRevert == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginRevert not implemented")}
	}
	beginRevert := f.beginRevert.get(req)
	if beginRevert == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/revert`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginRevert(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginRevert = &respr
		f.beginRevert.add(req, beginRevert)
	}

	resp, err := server.PollerResponderNext(beginRevert, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		f.beginRevert.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginRevert) {
		f.beginRevert.remove(req)
	}

	return resp, nil
}

func (f *FleetDatabasesServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if f.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := f.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DatabaseFleetManager/fleets/(?P<fleetName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/fleetspaces/(?P<fleetspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/databases/(?P<databaseName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 5 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatabasefleetmanager.FleetDatabase](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		fleetNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetName")])
		if err != nil {
			return nil, err
		}
		fleetspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fleetspaceName")])
		if err != nil {
			return nil, err
		}
		databaseNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("databaseName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := f.srv.BeginUpdate(req.Context(), resourceGroupNameParam, fleetNameParam, fleetspaceNameParam, databaseNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		f.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		f.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		f.beginUpdate.remove(req)
	}

	return resp, nil
}

// set this to conditionally intercept incoming requests to FleetDatabasesServerTransport
var fleetDatabasesServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
