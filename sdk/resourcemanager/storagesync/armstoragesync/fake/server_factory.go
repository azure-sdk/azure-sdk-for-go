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

// ServerFactory is a fake server for instances of the armstoragesync.ClientFactory type.
type ServerFactory struct {
	// CloudEndpointsServer contains the fakes for client CloudEndpointsClient
	CloudEndpointsServer CloudEndpointsServer

	// MicrosoftStorageSyncServer contains the fakes for client MicrosoftStorageSyncClient
	MicrosoftStorageSyncServer MicrosoftStorageSyncServer

	// OperationStatusServer contains the fakes for client OperationStatusClient
	OperationStatusServer OperationStatusServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PrivateEndpointConnectionsServer contains the fakes for client PrivateEndpointConnectionsClient
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer

	// PrivateLinkResourcesServer contains the fakes for client PrivateLinkResourcesClient
	PrivateLinkResourcesServer PrivateLinkResourcesServer

	// RegisteredServersServer contains the fakes for client RegisteredServersClient
	RegisteredServersServer RegisteredServersServer

	// ServerEndpointsServer contains the fakes for client ServerEndpointsClient
	ServerEndpointsServer ServerEndpointsServer

	// ServicesServer contains the fakes for client ServicesClient
	ServicesServer ServicesServer

	// SyncGroupsServer contains the fakes for client SyncGroupsClient
	SyncGroupsServer SyncGroupsServer

	// WorkflowsServer contains the fakes for client WorkflowsClient
	WorkflowsServer WorkflowsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armstoragesync.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armstoragesync.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                *ServerFactory
	trMu                               sync.Mutex
	trCloudEndpointsServer             *CloudEndpointsServerTransport
	trMicrosoftStorageSyncServer       *MicrosoftStorageSyncServerTransport
	trOperationStatusServer            *OperationStatusServerTransport
	trOperationsServer                 *OperationsServerTransport
	trPrivateEndpointConnectionsServer *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer       *PrivateLinkResourcesServerTransport
	trRegisteredServersServer          *RegisteredServersServerTransport
	trServerEndpointsServer            *ServerEndpointsServerTransport
	trServicesServer                   *ServicesServerTransport
	trSyncGroupsServer                 *SyncGroupsServerTransport
	trWorkflowsServer                  *WorkflowsServerTransport
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
	case "CloudEndpointsClient":
		initServer(s, &s.trCloudEndpointsServer, func() *CloudEndpointsServerTransport {
			return NewCloudEndpointsServerTransport(&s.srv.CloudEndpointsServer)
		})
		resp, err = s.trCloudEndpointsServer.Do(req)
	case "MicrosoftStorageSyncClient":
		initServer(s, &s.trMicrosoftStorageSyncServer, func() *MicrosoftStorageSyncServerTransport {
			return NewMicrosoftStorageSyncServerTransport(&s.srv.MicrosoftStorageSyncServer)
		})
		resp, err = s.trMicrosoftStorageSyncServer.Do(req)
	case "OperationStatusClient":
		initServer(s, &s.trOperationStatusServer, func() *OperationStatusServerTransport {
			return NewOperationStatusServerTransport(&s.srv.OperationStatusServer)
		})
		resp, err = s.trOperationStatusServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PrivateEndpointConnectionsClient":
		initServer(s, &s.trPrivateEndpointConnectionsServer, func() *PrivateEndpointConnectionsServerTransport {
			return NewPrivateEndpointConnectionsServerTransport(&s.srv.PrivateEndpointConnectionsServer)
		})
		resp, err = s.trPrivateEndpointConnectionsServer.Do(req)
	case "PrivateLinkResourcesClient":
		initServer(s, &s.trPrivateLinkResourcesServer, func() *PrivateLinkResourcesServerTransport {
			return NewPrivateLinkResourcesServerTransport(&s.srv.PrivateLinkResourcesServer)
		})
		resp, err = s.trPrivateLinkResourcesServer.Do(req)
	case "RegisteredServersClient":
		initServer(s, &s.trRegisteredServersServer, func() *RegisteredServersServerTransport {
			return NewRegisteredServersServerTransport(&s.srv.RegisteredServersServer)
		})
		resp, err = s.trRegisteredServersServer.Do(req)
	case "ServerEndpointsClient":
		initServer(s, &s.trServerEndpointsServer, func() *ServerEndpointsServerTransport {
			return NewServerEndpointsServerTransport(&s.srv.ServerEndpointsServer)
		})
		resp, err = s.trServerEndpointsServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
	case "SyncGroupsClient":
		initServer(s, &s.trSyncGroupsServer, func() *SyncGroupsServerTransport { return NewSyncGroupsServerTransport(&s.srv.SyncGroupsServer) })
		resp, err = s.trSyncGroupsServer.Do(req)
	case "WorkflowsClient":
		initServer(s, &s.trWorkflowsServer, func() *WorkflowsServerTransport { return NewWorkflowsServerTransport(&s.srv.WorkflowsServer) })
		resp, err = s.trWorkflowsServer.Do(req)
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
