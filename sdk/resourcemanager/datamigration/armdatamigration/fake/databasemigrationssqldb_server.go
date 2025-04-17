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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datamigration/armdatamigration"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// DatabaseMigrationsSQLDbServer is a fake server for instances of the armdatamigration.DatabaseMigrationsSQLDbClient type.
type DatabaseMigrationsSQLDbServer struct {
	// BeginCancel is the fake for method DatabaseMigrationsSQLDbClient.BeginCancel
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginCancel func(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters armdatamigration.MigrationOperationInput, options *armdatamigration.DatabaseMigrationsSQLDbClientBeginCancelOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientCancelResponse], errResp azfake.ErrorResponder)

	// BeginCreateOrUpdate is the fake for method DatabaseMigrationsSQLDbClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters armdatamigration.DatabaseMigrationSQLDb, options *armdatamigration.DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method DatabaseMigrationsSQLDbClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *armdatamigration.DatabaseMigrationsSQLDbClientBeginDeleteOptions) (resp azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DatabaseMigrationsSQLDbClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *armdatamigration.DatabaseMigrationsSQLDbClientGetOptions) (resp azfake.Responder[armdatamigration.DatabaseMigrationsSQLDbClientGetResponse], errResp azfake.ErrorResponder)
}

// NewDatabaseMigrationsSQLDbServerTransport creates a new instance of DatabaseMigrationsSQLDbServerTransport with the provided implementation.
// The returned DatabaseMigrationsSQLDbServerTransport instance is connected to an instance of armdatamigration.DatabaseMigrationsSQLDbClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewDatabaseMigrationsSQLDbServerTransport(srv *DatabaseMigrationsSQLDbServer) *DatabaseMigrationsSQLDbServerTransport {
	return &DatabaseMigrationsSQLDbServerTransport{
		srv:                 srv,
		beginCancel:         newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientCancelResponse]](),
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientDeleteResponse]](),
	}
}

// DatabaseMigrationsSQLDbServerTransport connects instances of armdatamigration.DatabaseMigrationsSQLDbClient to instances of DatabaseMigrationsSQLDbServer.
// Don't use this type directly, use NewDatabaseMigrationsSQLDbServerTransport instead.
type DatabaseMigrationsSQLDbServerTransport struct {
	srv                 *DatabaseMigrationsSQLDbServer
	beginCancel         *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientCancelResponse]]
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armdatamigration.DatabaseMigrationsSQLDbClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for DatabaseMigrationsSQLDbServerTransport.
func (d *DatabaseMigrationsSQLDbServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return d.dispatchToMethodFake(req, method)
}

func (d *DatabaseMigrationsSQLDbServerTransport) dispatchToMethodFake(req *http.Request, method string) (*http.Response, error) {
	resultChan := make(chan result)
	defer close(resultChan)

	go func() {
		var intercepted bool
		var res result
		if databaseMigrationsSqlDbServerTransportInterceptor != nil {
			res.resp, res.err, intercepted = databaseMigrationsSqlDbServerTransportInterceptor.Do(req)
		}
		if !intercepted {
			switch method {
			case "DatabaseMigrationsSQLDbClient.BeginCancel":
				res.resp, res.err = d.dispatchBeginCancel(req)
			case "DatabaseMigrationsSQLDbClient.BeginCreateOrUpdate":
				res.resp, res.err = d.dispatchBeginCreateOrUpdate(req)
			case "DatabaseMigrationsSQLDbClient.BeginDelete":
				res.resp, res.err = d.dispatchBeginDelete(req)
			case "DatabaseMigrationsSQLDbClient.Get":
				res.resp, res.err = d.dispatchGet(req)
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

func (d *DatabaseMigrationsSQLDbServerTransport) dispatchBeginCancel(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCancel == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCancel not implemented")}
	}
	beginCancel := d.beginCancel.get(req)
	if beginCancel == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<sqlDbInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/cancel`
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
		sqlDbInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlDbInstanceName")])
		if err != nil {
			return nil, err
		}
		targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCancel(req.Context(), resourceGroupNameParam, sqlDbInstanceNameParam, targetDbNameParam, body, nil)
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

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		d.beginCancel.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCancel) {
		d.beginCancel.remove(req)
	}

	return resp, nil
}

func (d *DatabaseMigrationsSQLDbServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := d.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<sqlDbInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdatamigration.DatabaseMigrationSQLDb](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		sqlDbInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlDbInstanceName")])
		if err != nil {
			return nil, err
		}
		targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := d.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, sqlDbInstanceNameParam, targetDbNameParam, body, nil)
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

func (d *DatabaseMigrationsSQLDbServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if d.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := d.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<sqlDbInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
		sqlDbInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlDbInstanceName")])
		if err != nil {
			return nil, err
		}
		targetDbNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetDbName")])
		if err != nil {
			return nil, err
		}
		forceUnescaped, err := url.QueryUnescape(qp.Get("force"))
		if err != nil {
			return nil, err
		}
		forceParam, err := parseOptional(forceUnescaped, strconv.ParseBool)
		if err != nil {
			return nil, err
		}
		var options *armdatamigration.DatabaseMigrationsSQLDbClientBeginDeleteOptions
		if forceParam != nil {
			options = &armdatamigration.DatabaseMigrationsSQLDbClientBeginDeleteOptions{
				Force: forceParam,
			}
		}
		respr, errRespr := d.srv.BeginDelete(req.Context(), resourceGroupNameParam, sqlDbInstanceNameParam, targetDbNameParam, options)
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

func (d *DatabaseMigrationsSQLDbServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Sql/servers/(?P<sqlDbInstanceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataMigration/databaseMigrations/(?P<targetDbName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	sqlDbInstanceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("sqlDbInstanceName")])
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
	var options *armdatamigration.DatabaseMigrationsSQLDbClientGetOptions
	if migrationOperationIDParam != nil || expandParam != nil {
		options = &armdatamigration.DatabaseMigrationsSQLDbClientGetOptions{
			MigrationOperationID: migrationOperationIDParam,
			Expand:               expandParam,
		}
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameParam, sqlDbInstanceNameParam, targetDbNameParam, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DatabaseMigrationSQLDb, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// set this to conditionally intercept incoming requests to DatabaseMigrationsSQLDbServerTransport
var databaseMigrationsSqlDbServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
