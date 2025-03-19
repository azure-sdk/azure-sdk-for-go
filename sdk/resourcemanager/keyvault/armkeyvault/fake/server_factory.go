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

// ServerFactory is a fake server for instances of the armkeyvault.ClientFactory type.
type ServerFactory struct {
	// KeysServer contains the fakes for client KeysClient
	KeysServer KeysServer

	// MHSMPrivateEndpointConnectionsServer contains the fakes for client MHSMPrivateEndpointConnectionsClient
	MHSMPrivateEndpointConnectionsServer MHSMPrivateEndpointConnectionsServer

	// MHSMPrivateLinkResourcesServer contains the fakes for client MHSMPrivateLinkResourcesClient
	MHSMPrivateLinkResourcesServer MHSMPrivateLinkResourcesServer

	// MHSMRegionsServer contains the fakes for client MHSMRegionsClient
	MHSMRegionsServer MHSMRegionsServer

	// ManagedHsmKeysServer contains the fakes for client ManagedHsmKeysClient
	ManagedHsmKeysServer ManagedHsmKeysServer

	// ManagedHsmsServer contains the fakes for client ManagedHsmsClient
	ManagedHsmsServer ManagedHsmsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PrivateEndpointConnectionsServer contains the fakes for client PrivateEndpointConnectionsClient
	PrivateEndpointConnectionsServer PrivateEndpointConnectionsServer

	// PrivateLinkResourcesServer contains the fakes for client PrivateLinkResourcesClient
	PrivateLinkResourcesServer PrivateLinkResourcesServer

	// SecretsServer contains the fakes for client SecretsClient
	SecretsServer SecretsServer

	// VaultsServer contains the fakes for client VaultsClient
	VaultsServer VaultsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armkeyvault.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armkeyvault.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                    *ServerFactory
	trMu                                   sync.Mutex
	trKeysServer                           *KeysServerTransport
	trMHSMPrivateEndpointConnectionsServer *MHSMPrivateEndpointConnectionsServerTransport
	trMHSMPrivateLinkResourcesServer       *MHSMPrivateLinkResourcesServerTransport
	trMHSMRegionsServer                    *MHSMRegionsServerTransport
	trManagedHsmKeysServer                 *ManagedHsmKeysServerTransport
	trManagedHsmsServer                    *ManagedHsmsServerTransport
	trOperationsServer                     *OperationsServerTransport
	trPrivateEndpointConnectionsServer     *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer           *PrivateLinkResourcesServerTransport
	trSecretsServer                        *SecretsServerTransport
	trVaultsServer                         *VaultsServerTransport
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
	case "KeysClient":
		initServer(s, &s.trKeysServer, func() *KeysServerTransport { return NewKeysServerTransport(&s.srv.KeysServer) })
		resp, err = s.trKeysServer.Do(req)
	case "MHSMPrivateEndpointConnectionsClient":
		initServer(s, &s.trMHSMPrivateEndpointConnectionsServer, func() *MHSMPrivateEndpointConnectionsServerTransport {
			return NewMHSMPrivateEndpointConnectionsServerTransport(&s.srv.MHSMPrivateEndpointConnectionsServer)
		})
		resp, err = s.trMHSMPrivateEndpointConnectionsServer.Do(req)
	case "MHSMPrivateLinkResourcesClient":
		initServer(s, &s.trMHSMPrivateLinkResourcesServer, func() *MHSMPrivateLinkResourcesServerTransport {
			return NewMHSMPrivateLinkResourcesServerTransport(&s.srv.MHSMPrivateLinkResourcesServer)
		})
		resp, err = s.trMHSMPrivateLinkResourcesServer.Do(req)
	case "MHSMRegionsClient":
		initServer(s, &s.trMHSMRegionsServer, func() *MHSMRegionsServerTransport { return NewMHSMRegionsServerTransport(&s.srv.MHSMRegionsServer) })
		resp, err = s.trMHSMRegionsServer.Do(req)
	case "ManagedHsmKeysClient":
		initServer(s, &s.trManagedHsmKeysServer, func() *ManagedHsmKeysServerTransport {
			return NewManagedHsmKeysServerTransport(&s.srv.ManagedHsmKeysServer)
		})
		resp, err = s.trManagedHsmKeysServer.Do(req)
	case "ManagedHsmsClient":
		initServer(s, &s.trManagedHsmsServer, func() *ManagedHsmsServerTransport { return NewManagedHsmsServerTransport(&s.srv.ManagedHsmsServer) })
		resp, err = s.trManagedHsmsServer.Do(req)
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
	case "SecretsClient":
		initServer(s, &s.trSecretsServer, func() *SecretsServerTransport { return NewSecretsServerTransport(&s.srv.SecretsServer) })
		resp, err = s.trSecretsServer.Do(req)
	case "VaultsClient":
		initServer(s, &s.trVaultsServer, func() *VaultsServerTransport { return NewVaultsServerTransport(&s.srv.VaultsServer) })
		resp, err = s.trVaultsServer.Do(req)
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
