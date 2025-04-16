// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armhybriddatamanager.ClientFactory type.
type ServerFactory struct {
	// DataManagersServer contains the fakes for client DataManagersClient
	DataManagersServer DataManagersServer

	// DataServicesServer contains the fakes for client DataServicesClient
	DataServicesServer DataServicesServer

	// DataStoreTypesServer contains the fakes for client DataStoreTypesClient
	DataStoreTypesServer DataStoreTypesServer

	// DataStoresServer contains the fakes for client DataStoresClient
	DataStoresServer DataStoresServer

	// JobDefinitionsServer contains the fakes for client JobDefinitionsClient
	JobDefinitionsServer JobDefinitionsServer

	// JobsServer contains the fakes for client JobsClient
	JobsServer JobsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PublicKeysServer contains the fakes for client PublicKeysClient
	PublicKeysServer PublicKeysServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armhybriddatamanager.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armhybriddatamanager.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                    *ServerFactory
	trMu                   sync.Mutex
	trDataManagersServer   *DataManagersServerTransport
	trDataServicesServer   *DataServicesServerTransport
	trDataStoreTypesServer *DataStoreTypesServerTransport
	trDataStoresServer     *DataStoresServerTransport
	trJobDefinitionsServer *JobDefinitionsServerTransport
	trJobsServer           *JobsServerTransport
	trOperationsServer     *OperationsServerTransport
	trPublicKeysServer     *PublicKeysServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "DataManagersClient":
		initServer(s, &s.trDataManagersServer, func() *DataManagersServerTransport { return NewDataManagersServerTransport(&s.srv.DataManagersServer) })
		resp, err = s.trDataManagersServer.Do(req)
	case "DataServicesClient":
		initServer(s, &s.trDataServicesServer, func() *DataServicesServerTransport { return NewDataServicesServerTransport(&s.srv.DataServicesServer) })
		resp, err = s.trDataServicesServer.Do(req)
	case "DataStoreTypesClient":
		initServer(s, &s.trDataStoreTypesServer, func() *DataStoreTypesServerTransport {
			return NewDataStoreTypesServerTransport(&s.srv.DataStoreTypesServer)
		})
		resp, err = s.trDataStoreTypesServer.Do(req)
	case "DataStoresClient":
		initServer(s, &s.trDataStoresServer, func() *DataStoresServerTransport { return NewDataStoresServerTransport(&s.srv.DataStoresServer) })
		resp, err = s.trDataStoresServer.Do(req)
	case "JobDefinitionsClient":
		initServer(s, &s.trJobDefinitionsServer, func() *JobDefinitionsServerTransport {
			return NewJobDefinitionsServerTransport(&s.srv.JobDefinitionsServer)
		})
		resp, err = s.trJobDefinitionsServer.Do(req)
	case "JobsClient":
		initServer(s, &s.trJobsServer, func() *JobsServerTransport { return NewJobsServerTransport(&s.srv.JobsServer) })
		resp, err = s.trJobsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PublicKeysClient":
		initServer(s, &s.trPublicKeysServer, func() *PublicKeysServerTransport { return NewPublicKeysServerTransport(&s.srv.PublicKeysServer) })
		resp, err = s.trPublicKeysServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
