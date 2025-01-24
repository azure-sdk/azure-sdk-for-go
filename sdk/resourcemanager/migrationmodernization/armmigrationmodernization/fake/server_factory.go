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

// ServerFactory is a fake server for instances of the armmigrationmodernization.ClientFactory type.
type ServerFactory struct {
	DeployedResourceServer                  DeployedResourceServer
	MigrateAgentServer                      MigrateAgentServer
	MigrateAgentOperationStatusServer       MigrateAgentOperationStatusServer
	ModernizeProjectServer                  ModernizeProjectServer
	ModernizeProjectOperationStatusServer   ModernizeProjectOperationStatusServer
	ModernizeProjectStatisticsServer        ModernizeProjectStatisticsServer
	OperationsServer                        OperationsServer
	WorkflowServer                          WorkflowServer
	WorkflowOperationStatusServer           WorkflowOperationStatusServer
	WorkloadDeploymentServer                WorkloadDeploymentServer
	WorkloadDeploymentOperationStatusServer WorkloadDeploymentOperationStatusServer
	WorkloadInstanceServer                  WorkloadInstanceServer
	WorkloadInstanceOperationStatusServer   WorkloadInstanceOperationStatusServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armmigrationmodernization.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armmigrationmodernization.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                       *ServerFactory
	trMu                                      sync.Mutex
	trDeployedResourceServer                  *DeployedResourceServerTransport
	trMigrateAgentServer                      *MigrateAgentServerTransport
	trMigrateAgentOperationStatusServer       *MigrateAgentOperationStatusServerTransport
	trModernizeProjectServer                  *ModernizeProjectServerTransport
	trModernizeProjectOperationStatusServer   *ModernizeProjectOperationStatusServerTransport
	trModernizeProjectStatisticsServer        *ModernizeProjectStatisticsServerTransport
	trOperationsServer                        *OperationsServerTransport
	trWorkflowServer                          *WorkflowServerTransport
	trWorkflowOperationStatusServer           *WorkflowOperationStatusServerTransport
	trWorkloadDeploymentServer                *WorkloadDeploymentServerTransport
	trWorkloadDeploymentOperationStatusServer *WorkloadDeploymentOperationStatusServerTransport
	trWorkloadInstanceServer                  *WorkloadInstanceServerTransport
	trWorkloadInstanceOperationStatusServer   *WorkloadInstanceOperationStatusServerTransport
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
	case "DeployedResourceClient":
		initServer(s, &s.trDeployedResourceServer, func() *DeployedResourceServerTransport {
			return NewDeployedResourceServerTransport(&s.srv.DeployedResourceServer)
		})
		resp, err = s.trDeployedResourceServer.Do(req)
	case "MigrateAgentClient":
		initServer(s, &s.trMigrateAgentServer, func() *MigrateAgentServerTransport { return NewMigrateAgentServerTransport(&s.srv.MigrateAgentServer) })
		resp, err = s.trMigrateAgentServer.Do(req)
	case "MigrateAgentOperationStatusClient":
		initServer(s, &s.trMigrateAgentOperationStatusServer, func() *MigrateAgentOperationStatusServerTransport {
			return NewMigrateAgentOperationStatusServerTransport(&s.srv.MigrateAgentOperationStatusServer)
		})
		resp, err = s.trMigrateAgentOperationStatusServer.Do(req)
	case "ModernizeProjectClient":
		initServer(s, &s.trModernizeProjectServer, func() *ModernizeProjectServerTransport {
			return NewModernizeProjectServerTransport(&s.srv.ModernizeProjectServer)
		})
		resp, err = s.trModernizeProjectServer.Do(req)
	case "ModernizeProjectOperationStatusClient":
		initServer(s, &s.trModernizeProjectOperationStatusServer, func() *ModernizeProjectOperationStatusServerTransport {
			return NewModernizeProjectOperationStatusServerTransport(&s.srv.ModernizeProjectOperationStatusServer)
		})
		resp, err = s.trModernizeProjectOperationStatusServer.Do(req)
	case "ModernizeProjectStatisticsClient":
		initServer(s, &s.trModernizeProjectStatisticsServer, func() *ModernizeProjectStatisticsServerTransport {
			return NewModernizeProjectStatisticsServerTransport(&s.srv.ModernizeProjectStatisticsServer)
		})
		resp, err = s.trModernizeProjectStatisticsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "WorkflowClient":
		initServer(s, &s.trWorkflowServer, func() *WorkflowServerTransport { return NewWorkflowServerTransport(&s.srv.WorkflowServer) })
		resp, err = s.trWorkflowServer.Do(req)
	case "WorkflowOperationStatusClient":
		initServer(s, &s.trWorkflowOperationStatusServer, func() *WorkflowOperationStatusServerTransport {
			return NewWorkflowOperationStatusServerTransport(&s.srv.WorkflowOperationStatusServer)
		})
		resp, err = s.trWorkflowOperationStatusServer.Do(req)
	case "WorkloadDeploymentClient":
		initServer(s, &s.trWorkloadDeploymentServer, func() *WorkloadDeploymentServerTransport {
			return NewWorkloadDeploymentServerTransport(&s.srv.WorkloadDeploymentServer)
		})
		resp, err = s.trWorkloadDeploymentServer.Do(req)
	case "WorkloadDeploymentOperationStatusClient":
		initServer(s, &s.trWorkloadDeploymentOperationStatusServer, func() *WorkloadDeploymentOperationStatusServerTransport {
			return NewWorkloadDeploymentOperationStatusServerTransport(&s.srv.WorkloadDeploymentOperationStatusServer)
		})
		resp, err = s.trWorkloadDeploymentOperationStatusServer.Do(req)
	case "WorkloadInstanceClient":
		initServer(s, &s.trWorkloadInstanceServer, func() *WorkloadInstanceServerTransport {
			return NewWorkloadInstanceServerTransport(&s.srv.WorkloadInstanceServer)
		})
		resp, err = s.trWorkloadInstanceServer.Do(req)
	case "WorkloadInstanceOperationStatusClient":
		initServer(s, &s.trWorkloadInstanceOperationStatusServer, func() *WorkloadInstanceOperationStatusServerTransport {
			return NewWorkloadInstanceOperationStatusServerTransport(&s.srv.WorkloadInstanceOperationStatusServer)
		})
		resp, err = s.trWorkloadInstanceOperationStatusServer.Do(req)
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
