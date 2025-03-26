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

// ServerFactory is a fake server for instances of the armservicefabricmanagedclusters.ClientFactory type.
type ServerFactory struct {
	// ApplicationTypeVersionsServer contains the fakes for client ApplicationTypeVersionsClient
	ApplicationTypeVersionsServer ApplicationTypeVersionsServer

	// ApplicationTypesServer contains the fakes for client ApplicationTypesClient
	ApplicationTypesServer ApplicationTypesServer

	// ApplicationsServer contains the fakes for client ApplicationsClient
	ApplicationsServer ApplicationsServer

	// ManagedApplyMaintenanceWindowServer contains the fakes for client ManagedApplyMaintenanceWindowClient
	ManagedApplyMaintenanceWindowServer ManagedApplyMaintenanceWindowServer

	// ManagedAzResiliencyStatusServer contains the fakes for client ManagedAzResiliencyStatusClient
	ManagedAzResiliencyStatusServer ManagedAzResiliencyStatusServer

	// ManagedClusterVersionServer contains the fakes for client ManagedClusterVersionClient
	ManagedClusterVersionServer ManagedClusterVersionServer

	// ManagedClustersServer contains the fakes for client ManagedClustersClient
	ManagedClustersServer ManagedClustersServer

	// ManagedMaintenanceWindowStatusServer contains the fakes for client ManagedMaintenanceWindowStatusClient
	ManagedMaintenanceWindowStatusServer ManagedMaintenanceWindowStatusServer

	// ManagedUnsupportedVMSizesServer contains the fakes for client ManagedUnsupportedVMSizesClient
	ManagedUnsupportedVMSizesServer ManagedUnsupportedVMSizesServer

	// NodeTypeSKUsServer contains the fakes for client NodeTypeSKUsClient
	NodeTypeSKUsServer NodeTypeSKUsServer

	// NodeTypesServer contains the fakes for client NodeTypesClient
	NodeTypesServer NodeTypesServer

	// OperationResultsServer contains the fakes for client OperationResultsClient
	OperationResultsServer OperationResultsServer

	// OperationStatusServer contains the fakes for client OperationStatusClient
	OperationStatusServer OperationStatusServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// ServicesServer contains the fakes for client ServicesClient
	ServicesServer ServicesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armservicefabricmanagedclusters.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armservicefabricmanagedclusters.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                    *ServerFactory
	trMu                                   sync.Mutex
	trApplicationTypeVersionsServer        *ApplicationTypeVersionsServerTransport
	trApplicationTypesServer               *ApplicationTypesServerTransport
	trApplicationsServer                   *ApplicationsServerTransport
	trManagedApplyMaintenanceWindowServer  *ManagedApplyMaintenanceWindowServerTransport
	trManagedAzResiliencyStatusServer      *ManagedAzResiliencyStatusServerTransport
	trManagedClusterVersionServer          *ManagedClusterVersionServerTransport
	trManagedClustersServer                *ManagedClustersServerTransport
	trManagedMaintenanceWindowStatusServer *ManagedMaintenanceWindowStatusServerTransport
	trManagedUnsupportedVMSizesServer      *ManagedUnsupportedVMSizesServerTransport
	trNodeTypeSKUsServer                   *NodeTypeSKUsServerTransport
	trNodeTypesServer                      *NodeTypesServerTransport
	trOperationResultsServer               *OperationResultsServerTransport
	trOperationStatusServer                *OperationStatusServerTransport
	trOperationsServer                     *OperationsServerTransport
	trServicesServer                       *ServicesServerTransport
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
	case "ApplicationTypeVersionsClient":
		initServer(s, &s.trApplicationTypeVersionsServer, func() *ApplicationTypeVersionsServerTransport {
			return NewApplicationTypeVersionsServerTransport(&s.srv.ApplicationTypeVersionsServer)
		})
		resp, err = s.trApplicationTypeVersionsServer.Do(req)
	case "ApplicationTypesClient":
		initServer(s, &s.trApplicationTypesServer, func() *ApplicationTypesServerTransport {
			return NewApplicationTypesServerTransport(&s.srv.ApplicationTypesServer)
		})
		resp, err = s.trApplicationTypesServer.Do(req)
	case "ApplicationsClient":
		initServer(s, &s.trApplicationsServer, func() *ApplicationsServerTransport { return NewApplicationsServerTransport(&s.srv.ApplicationsServer) })
		resp, err = s.trApplicationsServer.Do(req)
	case "ManagedApplyMaintenanceWindowClient":
		initServer(s, &s.trManagedApplyMaintenanceWindowServer, func() *ManagedApplyMaintenanceWindowServerTransport {
			return NewManagedApplyMaintenanceWindowServerTransport(&s.srv.ManagedApplyMaintenanceWindowServer)
		})
		resp, err = s.trManagedApplyMaintenanceWindowServer.Do(req)
	case "ManagedAzResiliencyStatusClient":
		initServer(s, &s.trManagedAzResiliencyStatusServer, func() *ManagedAzResiliencyStatusServerTransport {
			return NewManagedAzResiliencyStatusServerTransport(&s.srv.ManagedAzResiliencyStatusServer)
		})
		resp, err = s.trManagedAzResiliencyStatusServer.Do(req)
	case "ManagedClusterVersionClient":
		initServer(s, &s.trManagedClusterVersionServer, func() *ManagedClusterVersionServerTransport {
			return NewManagedClusterVersionServerTransport(&s.srv.ManagedClusterVersionServer)
		})
		resp, err = s.trManagedClusterVersionServer.Do(req)
	case "ManagedClustersClient":
		initServer(s, &s.trManagedClustersServer, func() *ManagedClustersServerTransport {
			return NewManagedClustersServerTransport(&s.srv.ManagedClustersServer)
		})
		resp, err = s.trManagedClustersServer.Do(req)
	case "ManagedMaintenanceWindowStatusClient":
		initServer(s, &s.trManagedMaintenanceWindowStatusServer, func() *ManagedMaintenanceWindowStatusServerTransport {
			return NewManagedMaintenanceWindowStatusServerTransport(&s.srv.ManagedMaintenanceWindowStatusServer)
		})
		resp, err = s.trManagedMaintenanceWindowStatusServer.Do(req)
	case "ManagedUnsupportedVMSizesClient":
		initServer(s, &s.trManagedUnsupportedVMSizesServer, func() *ManagedUnsupportedVMSizesServerTransport {
			return NewManagedUnsupportedVMSizesServerTransport(&s.srv.ManagedUnsupportedVMSizesServer)
		})
		resp, err = s.trManagedUnsupportedVMSizesServer.Do(req)
	case "NodeTypeSKUsClient":
		initServer(s, &s.trNodeTypeSKUsServer, func() *NodeTypeSKUsServerTransport { return NewNodeTypeSKUsServerTransport(&s.srv.NodeTypeSKUsServer) })
		resp, err = s.trNodeTypeSKUsServer.Do(req)
	case "NodeTypesClient":
		initServer(s, &s.trNodeTypesServer, func() *NodeTypesServerTransport { return NewNodeTypesServerTransport(&s.srv.NodeTypesServer) })
		resp, err = s.trNodeTypesServer.Do(req)
	case "OperationResultsClient":
		initServer(s, &s.trOperationResultsServer, func() *OperationResultsServerTransport {
			return NewOperationResultsServerTransport(&s.srv.OperationResultsServer)
		})
		resp, err = s.trOperationResultsServer.Do(req)
	case "OperationStatusClient":
		initServer(s, &s.trOperationStatusServer, func() *OperationStatusServerTransport {
			return NewOperationStatusServerTransport(&s.srv.OperationStatusServer)
		})
		resp, err = s.trOperationStatusServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
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
