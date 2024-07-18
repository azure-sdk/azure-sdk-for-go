// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// ServerFactory is a fake server for instances of the armavs.ClientFactory type.
type ServerFactory struct {
	// AddonsServer contains the fakes for client AddonsClient
	AddonsServer AddonsServer

	// AuthorizationsServer contains the fakes for client AuthorizationsClient
	AuthorizationsServer AuthorizationsServer

	// CloudLinksServer contains the fakes for client CloudLinksClient
	CloudLinksServer CloudLinksServer

	// ClustersServer contains the fakes for client ClustersClient
	ClustersServer ClustersServer

	// DatastoresServer contains the fakes for client DatastoresClient
	DatastoresServer DatastoresServer

	// GlobalReachConnectionsServer contains the fakes for client GlobalReachConnectionsClient
	GlobalReachConnectionsServer GlobalReachConnectionsServer

	// HcxEnterpriseSitesServer contains the fakes for client HcxEnterpriseSitesClient
	HcxEnterpriseSitesServer HcxEnterpriseSitesServer

	// IscsiPathsServer contains the fakes for client IscsiPathsClient
	IscsiPathsServer IscsiPathsServer

	// LocationsServer contains the fakes for client LocationsClient
	LocationsServer LocationsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PlacementPoliciesServer contains the fakes for client PlacementPoliciesClient
	PlacementPoliciesServer PlacementPoliciesServer

	// PrivateCloudsServer contains the fakes for client PrivateCloudsClient
	PrivateCloudsServer PrivateCloudsServer

	// ScriptCmdletsServer contains the fakes for client ScriptCmdletsClient
	ScriptCmdletsServer ScriptCmdletsServer

	// ScriptExecutionsServer contains the fakes for client ScriptExecutionsClient
	ScriptExecutionsServer ScriptExecutionsServer

	// ScriptPackagesServer contains the fakes for client ScriptPackagesClient
	ScriptPackagesServer ScriptPackagesServer

	// VirtualMachinesServer contains the fakes for client VirtualMachinesClient
	VirtualMachinesServer VirtualMachinesServer

	// WorkloadNetworkDNSServicesServer contains the fakes for client WorkloadNetworkDNSServicesClient
	WorkloadNetworkDNSServicesServer WorkloadNetworkDNSServicesServer

	// WorkloadNetworkDNSZonesServer contains the fakes for client WorkloadNetworkDNSZonesClient
	WorkloadNetworkDNSZonesServer WorkloadNetworkDNSZonesServer

	// WorkloadNetworkDhcpConfigurationsServer contains the fakes for client WorkloadNetworkDhcpConfigurationsClient
	WorkloadNetworkDhcpConfigurationsServer WorkloadNetworkDhcpConfigurationsServer

	// WorkloadNetworkGatewaysServer contains the fakes for client WorkloadNetworkGatewaysClient
	WorkloadNetworkGatewaysServer WorkloadNetworkGatewaysServer

	// WorkloadNetworkPortMirroringProfilesServer contains the fakes for client WorkloadNetworkPortMirroringProfilesClient
	WorkloadNetworkPortMirroringProfilesServer WorkloadNetworkPortMirroringProfilesServer

	// WorkloadNetworkPublicIPsServer contains the fakes for client WorkloadNetworkPublicIPsClient
	WorkloadNetworkPublicIPsServer WorkloadNetworkPublicIPsServer

	// WorkloadNetworkSegmentsServer contains the fakes for client WorkloadNetworkSegmentsClient
	WorkloadNetworkSegmentsServer WorkloadNetworkSegmentsServer

	// WorkloadNetworkVMGroupsServer contains the fakes for client WorkloadNetworkVMGroupsClient
	WorkloadNetworkVMGroupsServer WorkloadNetworkVMGroupsServer

	// WorkloadNetworkVirtualMachinesServer contains the fakes for client WorkloadNetworkVirtualMachinesClient
	WorkloadNetworkVirtualMachinesServer WorkloadNetworkVirtualMachinesServer

	// WorkloadNetworksServer contains the fakes for client WorkloadNetworksClient
	WorkloadNetworksServer WorkloadNetworksServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armavs.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armavs.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                          *ServerFactory
	trMu                                         sync.Mutex
	trAddonsServer                               *AddonsServerTransport
	trAuthorizationsServer                       *AuthorizationsServerTransport
	trCloudLinksServer                           *CloudLinksServerTransport
	trClustersServer                             *ClustersServerTransport
	trDatastoresServer                           *DatastoresServerTransport
	trGlobalReachConnectionsServer               *GlobalReachConnectionsServerTransport
	trHcxEnterpriseSitesServer                   *HcxEnterpriseSitesServerTransport
	trIscsiPathsServer                           *IscsiPathsServerTransport
	trLocationsServer                            *LocationsServerTransport
	trOperationsServer                           *OperationsServerTransport
	trPlacementPoliciesServer                    *PlacementPoliciesServerTransport
	trPrivateCloudsServer                        *PrivateCloudsServerTransport
	trScriptCmdletsServer                        *ScriptCmdletsServerTransport
	trScriptExecutionsServer                     *ScriptExecutionsServerTransport
	trScriptPackagesServer                       *ScriptPackagesServerTransport
	trVirtualMachinesServer                      *VirtualMachinesServerTransport
	trWorkloadNetworkDNSServicesServer           *WorkloadNetworkDNSServicesServerTransport
	trWorkloadNetworkDNSZonesServer              *WorkloadNetworkDNSZonesServerTransport
	trWorkloadNetworkDhcpConfigurationsServer    *WorkloadNetworkDhcpConfigurationsServerTransport
	trWorkloadNetworkGatewaysServer              *WorkloadNetworkGatewaysServerTransport
	trWorkloadNetworkPortMirroringProfilesServer *WorkloadNetworkPortMirroringProfilesServerTransport
	trWorkloadNetworkPublicIPsServer             *WorkloadNetworkPublicIPsServerTransport
	trWorkloadNetworkSegmentsServer              *WorkloadNetworkSegmentsServerTransport
	trWorkloadNetworkVMGroupsServer              *WorkloadNetworkVMGroupsServerTransport
	trWorkloadNetworkVirtualMachinesServer       *WorkloadNetworkVirtualMachinesServerTransport
	trWorkloadNetworksServer                     *WorkloadNetworksServerTransport
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
	case "AddonsClient":
		initServer(s, &s.trAddonsServer, func() *AddonsServerTransport { return NewAddonsServerTransport(&s.srv.AddonsServer) })
		resp, err = s.trAddonsServer.Do(req)
	case "AuthorizationsClient":
		initServer(s, &s.trAuthorizationsServer, func() *AuthorizationsServerTransport {
			return NewAuthorizationsServerTransport(&s.srv.AuthorizationsServer)
		})
		resp, err = s.trAuthorizationsServer.Do(req)
	case "CloudLinksClient":
		initServer(s, &s.trCloudLinksServer, func() *CloudLinksServerTransport { return NewCloudLinksServerTransport(&s.srv.CloudLinksServer) })
		resp, err = s.trCloudLinksServer.Do(req)
	case "ClustersClient":
		initServer(s, &s.trClustersServer, func() *ClustersServerTransport { return NewClustersServerTransport(&s.srv.ClustersServer) })
		resp, err = s.trClustersServer.Do(req)
	case "DatastoresClient":
		initServer(s, &s.trDatastoresServer, func() *DatastoresServerTransport { return NewDatastoresServerTransport(&s.srv.DatastoresServer) })
		resp, err = s.trDatastoresServer.Do(req)
	case "GlobalReachConnectionsClient":
		initServer(s, &s.trGlobalReachConnectionsServer, func() *GlobalReachConnectionsServerTransport {
			return NewGlobalReachConnectionsServerTransport(&s.srv.GlobalReachConnectionsServer)
		})
		resp, err = s.trGlobalReachConnectionsServer.Do(req)
	case "HcxEnterpriseSitesClient":
		initServer(s, &s.trHcxEnterpriseSitesServer, func() *HcxEnterpriseSitesServerTransport {
			return NewHcxEnterpriseSitesServerTransport(&s.srv.HcxEnterpriseSitesServer)
		})
		resp, err = s.trHcxEnterpriseSitesServer.Do(req)
	case "IscsiPathsClient":
		initServer(s, &s.trIscsiPathsServer, func() *IscsiPathsServerTransport { return NewIscsiPathsServerTransport(&s.srv.IscsiPathsServer) })
		resp, err = s.trIscsiPathsServer.Do(req)
	case "LocationsClient":
		initServer(s, &s.trLocationsServer, func() *LocationsServerTransport { return NewLocationsServerTransport(&s.srv.LocationsServer) })
		resp, err = s.trLocationsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PlacementPoliciesClient":
		initServer(s, &s.trPlacementPoliciesServer, func() *PlacementPoliciesServerTransport {
			return NewPlacementPoliciesServerTransport(&s.srv.PlacementPoliciesServer)
		})
		resp, err = s.trPlacementPoliciesServer.Do(req)
	case "PrivateCloudsClient":
		initServer(s, &s.trPrivateCloudsServer, func() *PrivateCloudsServerTransport {
			return NewPrivateCloudsServerTransport(&s.srv.PrivateCloudsServer)
		})
		resp, err = s.trPrivateCloudsServer.Do(req)
	case "ScriptCmdletsClient":
		initServer(s, &s.trScriptCmdletsServer, func() *ScriptCmdletsServerTransport {
			return NewScriptCmdletsServerTransport(&s.srv.ScriptCmdletsServer)
		})
		resp, err = s.trScriptCmdletsServer.Do(req)
	case "ScriptExecutionsClient":
		initServer(s, &s.trScriptExecutionsServer, func() *ScriptExecutionsServerTransport {
			return NewScriptExecutionsServerTransport(&s.srv.ScriptExecutionsServer)
		})
		resp, err = s.trScriptExecutionsServer.Do(req)
	case "ScriptPackagesClient":
		initServer(s, &s.trScriptPackagesServer, func() *ScriptPackagesServerTransport {
			return NewScriptPackagesServerTransport(&s.srv.ScriptPackagesServer)
		})
		resp, err = s.trScriptPackagesServer.Do(req)
	case "VirtualMachinesClient":
		initServer(s, &s.trVirtualMachinesServer, func() *VirtualMachinesServerTransport {
			return NewVirtualMachinesServerTransport(&s.srv.VirtualMachinesServer)
		})
		resp, err = s.trVirtualMachinesServer.Do(req)
	case "WorkloadNetworkDNSServicesClient":
		initServer(s, &s.trWorkloadNetworkDNSServicesServer, func() *WorkloadNetworkDNSServicesServerTransport {
			return NewWorkloadNetworkDNSServicesServerTransport(&s.srv.WorkloadNetworkDNSServicesServer)
		})
		resp, err = s.trWorkloadNetworkDNSServicesServer.Do(req)
	case "WorkloadNetworkDNSZonesClient":
		initServer(s, &s.trWorkloadNetworkDNSZonesServer, func() *WorkloadNetworkDNSZonesServerTransport {
			return NewWorkloadNetworkDNSZonesServerTransport(&s.srv.WorkloadNetworkDNSZonesServer)
		})
		resp, err = s.trWorkloadNetworkDNSZonesServer.Do(req)
	case "WorkloadNetworkDhcpConfigurationsClient":
		initServer(s, &s.trWorkloadNetworkDhcpConfigurationsServer, func() *WorkloadNetworkDhcpConfigurationsServerTransport {
			return NewWorkloadNetworkDhcpConfigurationsServerTransport(&s.srv.WorkloadNetworkDhcpConfigurationsServer)
		})
		resp, err = s.trWorkloadNetworkDhcpConfigurationsServer.Do(req)
	case "WorkloadNetworkGatewaysClient":
		initServer(s, &s.trWorkloadNetworkGatewaysServer, func() *WorkloadNetworkGatewaysServerTransport {
			return NewWorkloadNetworkGatewaysServerTransport(&s.srv.WorkloadNetworkGatewaysServer)
		})
		resp, err = s.trWorkloadNetworkGatewaysServer.Do(req)
	case "WorkloadNetworkPortMirroringProfilesClient":
		initServer(s, &s.trWorkloadNetworkPortMirroringProfilesServer, func() *WorkloadNetworkPortMirroringProfilesServerTransport {
			return NewWorkloadNetworkPortMirroringProfilesServerTransport(&s.srv.WorkloadNetworkPortMirroringProfilesServer)
		})
		resp, err = s.trWorkloadNetworkPortMirroringProfilesServer.Do(req)
	case "WorkloadNetworkPublicIPsClient":
		initServer(s, &s.trWorkloadNetworkPublicIPsServer, func() *WorkloadNetworkPublicIPsServerTransport {
			return NewWorkloadNetworkPublicIPsServerTransport(&s.srv.WorkloadNetworkPublicIPsServer)
		})
		resp, err = s.trWorkloadNetworkPublicIPsServer.Do(req)
	case "WorkloadNetworkSegmentsClient":
		initServer(s, &s.trWorkloadNetworkSegmentsServer, func() *WorkloadNetworkSegmentsServerTransport {
			return NewWorkloadNetworkSegmentsServerTransport(&s.srv.WorkloadNetworkSegmentsServer)
		})
		resp, err = s.trWorkloadNetworkSegmentsServer.Do(req)
	case "WorkloadNetworkVMGroupsClient":
		initServer(s, &s.trWorkloadNetworkVMGroupsServer, func() *WorkloadNetworkVMGroupsServerTransport {
			return NewWorkloadNetworkVMGroupsServerTransport(&s.srv.WorkloadNetworkVMGroupsServer)
		})
		resp, err = s.trWorkloadNetworkVMGroupsServer.Do(req)
	case "WorkloadNetworkVirtualMachinesClient":
		initServer(s, &s.trWorkloadNetworkVirtualMachinesServer, func() *WorkloadNetworkVirtualMachinesServerTransport {
			return NewWorkloadNetworkVirtualMachinesServerTransport(&s.srv.WorkloadNetworkVirtualMachinesServer)
		})
		resp, err = s.trWorkloadNetworkVirtualMachinesServer.Do(req)
	case "WorkloadNetworksClient":
		initServer(s, &s.trWorkloadNetworksServer, func() *WorkloadNetworksServerTransport {
			return NewWorkloadNetworksServerTransport(&s.srv.WorkloadNetworksServer)
		})
		resp, err = s.trWorkloadNetworksServer.Do(req)
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
