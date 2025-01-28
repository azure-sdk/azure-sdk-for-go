//go:build go1.18
// +build go1.18

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

// ServerFactory is a fake server for instances of the armhybridcompute.ClientFactory type.
type ServerFactory struct {
	ExtensionMetadataServer                      ExtensionMetadataServer
	ExtensionMetadataV2Server                    ExtensionMetadataV2Server
	ExtensionPublisherServer                     ExtensionPublisherServer
	ExtensionTypeServer                          ExtensionTypeServer
	GatewaysServer                               GatewaysServer
	LicenseProfilesServer                        LicenseProfilesServer
	LicensesServer                               LicensesServer
	MachineExtensionsServer                      MachineExtensionsServer
	MachineRunCommandsServer                     MachineRunCommandsServer
	MachinesServer                               MachinesServer
	ManagementServer                             ManagementServer
	NetworkProfileServer                         NetworkProfileServer
	NetworkSecurityPerimeterConfigurationsServer NetworkSecurityPerimeterConfigurationsServer
	OperationsServer                             OperationsServer
	PrivateEndpointConnectionsServer             PrivateEndpointConnectionsServer
	PrivateLinkResourcesServer                   PrivateLinkResourcesServer
	PrivateLinkScopesServer                      PrivateLinkScopesServer
	SettingsServer                               SettingsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armhybridcompute.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armhybridcompute.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                            *ServerFactory
	trMu                                           sync.Mutex
	trExtensionMetadataServer                      *ExtensionMetadataServerTransport
	trExtensionMetadataV2Server                    *ExtensionMetadataV2ServerTransport
	trExtensionPublisherServer                     *ExtensionPublisherServerTransport
	trExtensionTypeServer                          *ExtensionTypeServerTransport
	trGatewaysServer                               *GatewaysServerTransport
	trLicenseProfilesServer                        *LicenseProfilesServerTransport
	trLicensesServer                               *LicensesServerTransport
	trMachineExtensionsServer                      *MachineExtensionsServerTransport
	trMachineRunCommandsServer                     *MachineRunCommandsServerTransport
	trMachinesServer                               *MachinesServerTransport
	trManagementServer                             *ManagementServerTransport
	trNetworkProfileServer                         *NetworkProfileServerTransport
	trNetworkSecurityPerimeterConfigurationsServer *NetworkSecurityPerimeterConfigurationsServerTransport
	trOperationsServer                             *OperationsServerTransport
	trPrivateEndpointConnectionsServer             *PrivateEndpointConnectionsServerTransport
	trPrivateLinkResourcesServer                   *PrivateLinkResourcesServerTransport
	trPrivateLinkScopesServer                      *PrivateLinkScopesServerTransport
	trSettingsServer                               *SettingsServerTransport
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
	case "ExtensionMetadataClient":
		initServer(s, &s.trExtensionMetadataServer, func() *ExtensionMetadataServerTransport {
			return NewExtensionMetadataServerTransport(&s.srv.ExtensionMetadataServer)
		})
		resp, err = s.trExtensionMetadataServer.Do(req)
	case "ExtensionMetadataV2Client":
		initServer(s, &s.trExtensionMetadataV2Server, func() *ExtensionMetadataV2ServerTransport {
			return NewExtensionMetadataV2ServerTransport(&s.srv.ExtensionMetadataV2Server)
		})
		resp, err = s.trExtensionMetadataV2Server.Do(req)
	case "ExtensionPublisherClient":
		initServer(s, &s.trExtensionPublisherServer, func() *ExtensionPublisherServerTransport {
			return NewExtensionPublisherServerTransport(&s.srv.ExtensionPublisherServer)
		})
		resp, err = s.trExtensionPublisherServer.Do(req)
	case "ExtensionTypeClient":
		initServer(s, &s.trExtensionTypeServer, func() *ExtensionTypeServerTransport {
			return NewExtensionTypeServerTransport(&s.srv.ExtensionTypeServer)
		})
		resp, err = s.trExtensionTypeServer.Do(req)
	case "GatewaysClient":
		initServer(s, &s.trGatewaysServer, func() *GatewaysServerTransport { return NewGatewaysServerTransport(&s.srv.GatewaysServer) })
		resp, err = s.trGatewaysServer.Do(req)
	case "LicenseProfilesClient":
		initServer(s, &s.trLicenseProfilesServer, func() *LicenseProfilesServerTransport {
			return NewLicenseProfilesServerTransport(&s.srv.LicenseProfilesServer)
		})
		resp, err = s.trLicenseProfilesServer.Do(req)
	case "LicensesClient":
		initServer(s, &s.trLicensesServer, func() *LicensesServerTransport { return NewLicensesServerTransport(&s.srv.LicensesServer) })
		resp, err = s.trLicensesServer.Do(req)
	case "MachineExtensionsClient":
		initServer(s, &s.trMachineExtensionsServer, func() *MachineExtensionsServerTransport {
			return NewMachineExtensionsServerTransport(&s.srv.MachineExtensionsServer)
		})
		resp, err = s.trMachineExtensionsServer.Do(req)
	case "MachineRunCommandsClient":
		initServer(s, &s.trMachineRunCommandsServer, func() *MachineRunCommandsServerTransport {
			return NewMachineRunCommandsServerTransport(&s.srv.MachineRunCommandsServer)
		})
		resp, err = s.trMachineRunCommandsServer.Do(req)
	case "MachinesClient":
		initServer(s, &s.trMachinesServer, func() *MachinesServerTransport { return NewMachinesServerTransport(&s.srv.MachinesServer) })
		resp, err = s.trMachinesServer.Do(req)
	case "ManagementClient":
		initServer(s, &s.trManagementServer, func() *ManagementServerTransport { return NewManagementServerTransport(&s.srv.ManagementServer) })
		resp, err = s.trManagementServer.Do(req)
	case "NetworkProfileClient":
		initServer(s, &s.trNetworkProfileServer, func() *NetworkProfileServerTransport {
			return NewNetworkProfileServerTransport(&s.srv.NetworkProfileServer)
		})
		resp, err = s.trNetworkProfileServer.Do(req)
	case "NetworkSecurityPerimeterConfigurationsClient":
		initServer(s, &s.trNetworkSecurityPerimeterConfigurationsServer, func() *NetworkSecurityPerimeterConfigurationsServerTransport {
			return NewNetworkSecurityPerimeterConfigurationsServerTransport(&s.srv.NetworkSecurityPerimeterConfigurationsServer)
		})
		resp, err = s.trNetworkSecurityPerimeterConfigurationsServer.Do(req)
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
	case "PrivateLinkScopesClient":
		initServer(s, &s.trPrivateLinkScopesServer, func() *PrivateLinkScopesServerTransport {
			return NewPrivateLinkScopesServerTransport(&s.srv.PrivateLinkScopesServer)
		})
		resp, err = s.trPrivateLinkScopesServer.Do(req)
	case "SettingsClient":
		initServer(s, &s.trSettingsServer, func() *SettingsServerTransport { return NewSettingsServerTransport(&s.srv.SettingsServer) })
		resp, err = s.trSettingsServer.Do(req)
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
