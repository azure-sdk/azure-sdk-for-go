//go:build go1.18
// +build go1.18

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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration/v2"
	"net/http"
	"net/url"
	"regexp"
)

// DatabaseMigrationsSQLMiServer is a fake server for instances of the armdatamigration.DatabaseMigrationsSQLMiClient type.
type DatabaseMigrationsSQLMiServer struct {
	// BeginCancel is the fake for method DatabaseMigrationsSQLMiClient.BeginCancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCancel func(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters armdatamigration.MigrationOperationInput, options *armdatamigration.DatabaseMigrationsSQLMiClientBeginCancelOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCancelResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters armdatamigration.DatabaseMigrationSQLMi, options *armdatamigration.DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginCutover is the fake for method DatabaseMigrationsSQLMiClient.BeginCutover
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginCutover func(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters armdatamigration.MigrationOperationInput, options *armdatamigration.DatabaseMigrationsSQLMiClientBeginCutoverOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCutoverResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DatabaseMigrationsSQLMiClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, options *armdatamigration.DatabaseMigrationsSQLMiClientGetOptions) (resp azfake.Responder[armdatamigration.DatabaseMigrationsSQLMiClientGetResponse], errResp azfake.ErrorResponder)
}

// NewDatabaseMigrationsSQLMiServerTransport creates a new instance of DatabaseMigrationsSQLMiServerTransport with the provided implementation.
// The returned DatabaseMigrationsSQLMiServerTransport instance is connected to an instance of armdatamigration.DatabaseMigrationsSQLMiClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatabaseMigrationsSQLMiServerTransport(srv *DatabaseMigrationsSQLMiServer) *DatabaseMigrationsSQLMiServerTransport {
	return &DatabaseMigrationsSQLMiServerTransport{
		srv:                 srv,
		beginCancel:         newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCancelResponse]](),
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCreateOrUpdateResponse]](),
		beginCutover:        newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCutoverResponse]](),
	}
}

// DatabaseMigrationsSQLMiServerTransport connects instances of armdatamigration.DatabaseMigrationsSQLMiClient to instances of DatabaseMigrationsSQLMiServer.
// Don't use this type directly, use NewDatabaseMigrationsSQLMiServerTransport instead.
type DatabaseMigrationsSQLMiServerTransport struct {
	srv                 *DatabaseMigrationsSQLMiServer
	beginCancel         *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCancelResponse]]
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCreateOrUpdateResponse]]
	beginCutover        *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLMiClientCutoverResponse]]
}

// Do implements the policy.Transporter interface for DatabaseMigrationsSQLMiServerTransport.
func (d *DatabaseMigrationsSQLMiServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DatabaseMigrationsSQLMiClient.BeginCancel":
		resp, err = d.dispatchBeginCancel(req)
	case "DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate":
		resp, err = d.dispatchBeginCreateOrUpdate(req)
	case "DatabaseMigrationsSQLMiClient.BeginCutover":
		resp, err = d.dispatchBeginCutover(req)
	case "DatabaseMigrationsSQLMiClient.Get":
		resp, err = d.dispatchGet(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DatabaseMigrationsSQLMiServerTransport) dispatchBeginCancel(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancel not implemented")}
	}
	beginCancel := d.beginCancel.get(req)
	if beginCancel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.MigrationOperationInput](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCancel(req.Context(), resourceGroupNameParam, managedInstanceNameParam, targetDbNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCancel = &respr
		d.beginCancel.add(req, beginCancel)
	}

	resp, err := server.PollerResponderNext(beginCancel, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCancel.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCancel) {
		d.beginCancel.remove(req)
	}

	return resp, nil
}

func (d *DatabaseMigrationsSQLMiServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.DatabaseMigrationSQLMi](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, managedInstanceNameParam, targetDbNameParam, body, nil)
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

func (d *DatabaseMigrationsSQLMiServerTransport) dispatchBeginCutover(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCutover == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCutover not implemented")}
	}
	beginCutover := d.beginCutover.get(req)
	if beginCutover == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cutover`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.MigrationOperationInput](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
		if err != nil {
			return nil, err
		}
		targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCutover(req.Context(), resourceGroupNameParam, managedInstanceNameParam, targetDbNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCutover = &respr
		d.beginCutover.add(req, beginCutover)
	}

	resp, err := server.PollerResponderNext(beginCutover, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		d.beginCutover.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCutover) {
		d.beginCutover.remove(req)
	}

	return resp, nil
}

func (d *DatabaseMigrationsSQLMiServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/managedInstances/(?P<managedInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	managedInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("managedInstanceName")])
	if err != nil {
		return nil, err
	}
	targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
	if err != nil {
		return nil, err
	}
	migrationOperationIDUnescaped, err := url.QueryUnescape(qp.Get("migrationOperationId"))
	if err != nil {
		return nil, err
	}
	migrationOperationIDParam := getOptional(migrationOperationIDUnescaped)
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armdatamigration.DatabaseMigrationsSQLMiClientGetOptions
	if migrationOperationIDParam != nil || expandParam != nil {
		options = &armdatamigration.DatabaseMigrationsSQLMiClientGetOptions{
			MigrationOperationID: migrationOperationIDParam,
			Expand:               expandParam,
		}
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, managedInstanceNameParam, targetDbNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatabaseMigrationSQLMi, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
