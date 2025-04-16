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

// ServerFactory is a fake server for instances of the armappplatform.ClientFactory type.
type ServerFactory struct {
	// APIPortalCustomDomainsServer contains the fakes for client APIPortalCustomDomainsClient
	APIPortalCustomDomainsServer APIPortalCustomDomainsServer

	// APIPortalsServer contains the fakes for client APIPortalsClient
	APIPortalsServer APIPortalsServer

	// ApmsServer contains the fakes for client ApmsClient
	ApmsServer ApmsServer

	// ApplicationAcceleratorsServer contains the fakes for client ApplicationAcceleratorsClient
	ApplicationAcceleratorsServer ApplicationAcceleratorsServer

	// ApplicationLiveViewsServer contains the fakes for client ApplicationLiveViewsClient
	ApplicationLiveViewsServer ApplicationLiveViewsServer

	// AppsServer contains the fakes for client AppsClient
	AppsServer AppsServer

	// BindingsServer contains the fakes for client BindingsClient
	BindingsServer BindingsServer

	// BuildServiceAgentPoolServer contains the fakes for client BuildServiceAgentPoolClient
	BuildServiceAgentPoolServer BuildServiceAgentPoolServer

	// BuildServiceBuilderServer contains the fakes for client BuildServiceBuilderClient
	BuildServiceBuilderServer BuildServiceBuilderServer

	// BuildServiceServer contains the fakes for client BuildServiceClient
	BuildServiceServer BuildServiceServer

	// BuildpackBindingServer contains the fakes for client BuildpackBindingClient
	BuildpackBindingServer BuildpackBindingServer

	// CertificatesServer contains the fakes for client CertificatesClient
	CertificatesServer CertificatesServer

	// ConfigServersServer contains the fakes for client ConfigServersClient
	ConfigServersServer ConfigServersServer

	// ConfigurationServicesServer contains the fakes for client ConfigurationServicesClient
	ConfigurationServicesServer ConfigurationServicesServer

	// ContainerRegistriesServer contains the fakes for client ContainerRegistriesClient
	ContainerRegistriesServer ContainerRegistriesServer

	// CustomDomainsServer contains the fakes for client CustomDomainsClient
	CustomDomainsServer CustomDomainsServer

	// CustomizedAcceleratorsServer contains the fakes for client CustomizedAcceleratorsClient
	CustomizedAcceleratorsServer CustomizedAcceleratorsServer

	// DeploymentsServer contains the fakes for client DeploymentsClient
	DeploymentsServer DeploymentsServer

	// DevToolPortalsServer contains the fakes for client DevToolPortalsClient
	DevToolPortalsServer DevToolPortalsServer

	// EurekaServersServer contains the fakes for client EurekaServersClient
	EurekaServersServer EurekaServersServer

	// GatewayCustomDomainsServer contains the fakes for client GatewayCustomDomainsClient
	GatewayCustomDomainsServer GatewayCustomDomainsServer

	// GatewayRouteConfigsServer contains the fakes for client GatewayRouteConfigsClient
	GatewayRouteConfigsServer GatewayRouteConfigsServer

	// GatewaysServer contains the fakes for client GatewaysClient
	GatewaysServer GatewaysServer

	// JobServer contains the fakes for client JobClient
	JobServer JobServer

	// JobExecutionServer contains the fakes for client JobExecutionClient
	JobExecutionServer JobExecutionServer

	// JobExecutionsServer contains the fakes for client JobExecutionsClient
	JobExecutionsServer JobExecutionsServer

	// JobsServer contains the fakes for client JobsClient
	JobsServer JobsServer

	// MonitoringSettingsServer contains the fakes for client MonitoringSettingsClient
	MonitoringSettingsServer MonitoringSettingsServer

	// OperationsServer contains the fakes for client OperationsClient
	OperationsServer OperationsServer

	// PredefinedAcceleratorsServer contains the fakes for client PredefinedAcceleratorsClient
	PredefinedAcceleratorsServer PredefinedAcceleratorsServer

	// RuntimeVersionsServer contains the fakes for client RuntimeVersionsClient
	RuntimeVersionsServer RuntimeVersionsServer

	// SKUsServer contains the fakes for client SKUsClient
	SKUsServer SKUsServer

	// ServiceRegistriesServer contains the fakes for client ServiceRegistriesClient
	ServiceRegistriesServer ServiceRegistriesServer

	// ServicesServer contains the fakes for client ServicesClient
	ServicesServer ServicesServer

	// StoragesServer contains the fakes for client StoragesClient
	StoragesServer StoragesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armappplatform.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armappplatform.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                             *ServerFactory
	trMu                            sync.Mutex
	trAPIPortalCustomDomainsServer  *APIPortalCustomDomainsServerTransport
	trAPIPortalsServer              *APIPortalsServerTransport
	trApmsServer                    *ApmsServerTransport
	trApplicationAcceleratorsServer *ApplicationAcceleratorsServerTransport
	trApplicationLiveViewsServer    *ApplicationLiveViewsServerTransport
	trAppsServer                    *AppsServerTransport
	trBindingsServer                *BindingsServerTransport
	trBuildServiceAgentPoolServer   *BuildServiceAgentPoolServerTransport
	trBuildServiceBuilderServer     *BuildServiceBuilderServerTransport
	trBuildServiceServer            *BuildServiceServerTransport
	trBuildpackBindingServer        *BuildpackBindingServerTransport
	trCertificatesServer            *CertificatesServerTransport
	trConfigServersServer           *ConfigServersServerTransport
	trConfigurationServicesServer   *ConfigurationServicesServerTransport
	trContainerRegistriesServer     *ContainerRegistriesServerTransport
	trCustomDomainsServer           *CustomDomainsServerTransport
	trCustomizedAcceleratorsServer  *CustomizedAcceleratorsServerTransport
	trDeploymentsServer             *DeploymentsServerTransport
	trDevToolPortalsServer          *DevToolPortalsServerTransport
	trEurekaServersServer           *EurekaServersServerTransport
	trGatewayCustomDomainsServer    *GatewayCustomDomainsServerTransport
	trGatewayRouteConfigsServer     *GatewayRouteConfigsServerTransport
	trGatewaysServer                *GatewaysServerTransport
	trJobServer                     *JobServerTransport
	trJobExecutionServer            *JobExecutionServerTransport
	trJobExecutionsServer           *JobExecutionsServerTransport
	trJobsServer                    *JobsServerTransport
	trMonitoringSettingsServer      *MonitoringSettingsServerTransport
	trOperationsServer              *OperationsServerTransport
	trPredefinedAcceleratorsServer  *PredefinedAcceleratorsServerTransport
	trRuntimeVersionsServer         *RuntimeVersionsServerTransport
	trSKUsServer                    *SKUsServerTransport
	trServiceRegistriesServer       *ServiceRegistriesServerTransport
	trServicesServer                *ServicesServerTransport
	trStoragesServer                *StoragesServerTransport
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
	case "APIPortalCustomDomainsClient":
		initServer(s, &s.trAPIPortalCustomDomainsServer, func() *APIPortalCustomDomainsServerTransport {
			return NewAPIPortalCustomDomainsServerTransport(&s.srv.APIPortalCustomDomainsServer)
		})
		resp, err = s.trAPIPortalCustomDomainsServer.Do(req)
	case "APIPortalsClient":
		initServer(s, &s.trAPIPortalsServer, func() *APIPortalsServerTransport { return NewAPIPortalsServerTransport(&s.srv.APIPortalsServer) })
		resp, err = s.trAPIPortalsServer.Do(req)
	case "ApmsClient":
		initServer(s, &s.trApmsServer, func() *ApmsServerTransport { return NewApmsServerTransport(&s.srv.ApmsServer) })
		resp, err = s.trApmsServer.Do(req)
	case "ApplicationAcceleratorsClient":
		initServer(s, &s.trApplicationAcceleratorsServer, func() *ApplicationAcceleratorsServerTransport {
			return NewApplicationAcceleratorsServerTransport(&s.srv.ApplicationAcceleratorsServer)
		})
		resp, err = s.trApplicationAcceleratorsServer.Do(req)
	case "ApplicationLiveViewsClient":
		initServer(s, &s.trApplicationLiveViewsServer, func() *ApplicationLiveViewsServerTransport {
			return NewApplicationLiveViewsServerTransport(&s.srv.ApplicationLiveViewsServer)
		})
		resp, err = s.trApplicationLiveViewsServer.Do(req)
	case "AppsClient":
		initServer(s, &s.trAppsServer, func() *AppsServerTransport { return NewAppsServerTransport(&s.srv.AppsServer) })
		resp, err = s.trAppsServer.Do(req)
	case "BindingsClient":
		initServer(s, &s.trBindingsServer, func() *BindingsServerTransport { return NewBindingsServerTransport(&s.srv.BindingsServer) })
		resp, err = s.trBindingsServer.Do(req)
	case "BuildServiceAgentPoolClient":
		initServer(s, &s.trBuildServiceAgentPoolServer, func() *BuildServiceAgentPoolServerTransport {
			return NewBuildServiceAgentPoolServerTransport(&s.srv.BuildServiceAgentPoolServer)
		})
		resp, err = s.trBuildServiceAgentPoolServer.Do(req)
	case "BuildServiceBuilderClient":
		initServer(s, &s.trBuildServiceBuilderServer, func() *BuildServiceBuilderServerTransport {
			return NewBuildServiceBuilderServerTransport(&s.srv.BuildServiceBuilderServer)
		})
		resp, err = s.trBuildServiceBuilderServer.Do(req)
	case "BuildServiceClient":
		initServer(s, &s.trBuildServiceServer, func() *BuildServiceServerTransport { return NewBuildServiceServerTransport(&s.srv.BuildServiceServer) })
		resp, err = s.trBuildServiceServer.Do(req)
	case "BuildpackBindingClient":
		initServer(s, &s.trBuildpackBindingServer, func() *BuildpackBindingServerTransport {
			return NewBuildpackBindingServerTransport(&s.srv.BuildpackBindingServer)
		})
		resp, err = s.trBuildpackBindingServer.Do(req)
	case "CertificatesClient":
		initServer(s, &s.trCertificatesServer, func() *CertificatesServerTransport { return NewCertificatesServerTransport(&s.srv.CertificatesServer) })
		resp, err = s.trCertificatesServer.Do(req)
	case "ConfigServersClient":
		initServer(s, &s.trConfigServersServer, func() *ConfigServersServerTransport {
			return NewConfigServersServerTransport(&s.srv.ConfigServersServer)
		})
		resp, err = s.trConfigServersServer.Do(req)
	case "ConfigurationServicesClient":
		initServer(s, &s.trConfigurationServicesServer, func() *ConfigurationServicesServerTransport {
			return NewConfigurationServicesServerTransport(&s.srv.ConfigurationServicesServer)
		})
		resp, err = s.trConfigurationServicesServer.Do(req)
	case "ContainerRegistriesClient":
		initServer(s, &s.trContainerRegistriesServer, func() *ContainerRegistriesServerTransport {
			return NewContainerRegistriesServerTransport(&s.srv.ContainerRegistriesServer)
		})
		resp, err = s.trContainerRegistriesServer.Do(req)
	case "CustomDomainsClient":
		initServer(s, &s.trCustomDomainsServer, func() *CustomDomainsServerTransport {
			return NewCustomDomainsServerTransport(&s.srv.CustomDomainsServer)
		})
		resp, err = s.trCustomDomainsServer.Do(req)
	case "CustomizedAcceleratorsClient":
		initServer(s, &s.trCustomizedAcceleratorsServer, func() *CustomizedAcceleratorsServerTransport {
			return NewCustomizedAcceleratorsServerTransport(&s.srv.CustomizedAcceleratorsServer)
		})
		resp, err = s.trCustomizedAcceleratorsServer.Do(req)
	case "DeploymentsClient":
		initServer(s, &s.trDeploymentsServer, func() *DeploymentsServerTransport { return NewDeploymentsServerTransport(&s.srv.DeploymentsServer) })
		resp, err = s.trDeploymentsServer.Do(req)
	case "DevToolPortalsClient":
		initServer(s, &s.trDevToolPortalsServer, func() *DevToolPortalsServerTransport {
			return NewDevToolPortalsServerTransport(&s.srv.DevToolPortalsServer)
		})
		resp, err = s.trDevToolPortalsServer.Do(req)
	case "EurekaServersClient":
		initServer(s, &s.trEurekaServersServer, func() *EurekaServersServerTransport {
			return NewEurekaServersServerTransport(&s.srv.EurekaServersServer)
		})
		resp, err = s.trEurekaServersServer.Do(req)
	case "GatewayCustomDomainsClient":
		initServer(s, &s.trGatewayCustomDomainsServer, func() *GatewayCustomDomainsServerTransport {
			return NewGatewayCustomDomainsServerTransport(&s.srv.GatewayCustomDomainsServer)
		})
		resp, err = s.trGatewayCustomDomainsServer.Do(req)
	case "GatewayRouteConfigsClient":
		initServer(s, &s.trGatewayRouteConfigsServer, func() *GatewayRouteConfigsServerTransport {
			return NewGatewayRouteConfigsServerTransport(&s.srv.GatewayRouteConfigsServer)
		})
		resp, err = s.trGatewayRouteConfigsServer.Do(req)
	case "GatewaysClient":
		initServer(s, &s.trGatewaysServer, func() *GatewaysServerTransport { return NewGatewaysServerTransport(&s.srv.GatewaysServer) })
		resp, err = s.trGatewaysServer.Do(req)
	case "JobClient":
		initServer(s, &s.trJobServer, func() *JobServerTransport { return NewJobServerTransport(&s.srv.JobServer) })
		resp, err = s.trJobServer.Do(req)
	case "JobExecutionClient":
		initServer(s, &s.trJobExecutionServer, func() *JobExecutionServerTransport { return NewJobExecutionServerTransport(&s.srv.JobExecutionServer) })
		resp, err = s.trJobExecutionServer.Do(req)
	case "JobExecutionsClient":
		initServer(s, &s.trJobExecutionsServer, func() *JobExecutionsServerTransport {
			return NewJobExecutionsServerTransport(&s.srv.JobExecutionsServer)
		})
		resp, err = s.trJobExecutionsServer.Do(req)
	case "JobsClient":
		initServer(s, &s.trJobsServer, func() *JobsServerTransport { return NewJobsServerTransport(&s.srv.JobsServer) })
		resp, err = s.trJobsServer.Do(req)
	case "MonitoringSettingsClient":
		initServer(s, &s.trMonitoringSettingsServer, func() *MonitoringSettingsServerTransport {
			return NewMonitoringSettingsServerTransport(&s.srv.MonitoringSettingsServer)
		})
		resp, err = s.trMonitoringSettingsServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "PredefinedAcceleratorsClient":
		initServer(s, &s.trPredefinedAcceleratorsServer, func() *PredefinedAcceleratorsServerTransport {
			return NewPredefinedAcceleratorsServerTransport(&s.srv.PredefinedAcceleratorsServer)
		})
		resp, err = s.trPredefinedAcceleratorsServer.Do(req)
	case "RuntimeVersionsClient":
		initServer(s, &s.trRuntimeVersionsServer, func() *RuntimeVersionsServerTransport {
			return NewRuntimeVersionsServerTransport(&s.srv.RuntimeVersionsServer)
		})
		resp, err = s.trRuntimeVersionsServer.Do(req)
	case "SKUsClient":
		initServer(s, &s.trSKUsServer, func() *SKUsServerTransport { return NewSKUsServerTransport(&s.srv.SKUsServer) })
		resp, err = s.trSKUsServer.Do(req)
	case "ServiceRegistriesClient":
		initServer(s, &s.trServiceRegistriesServer, func() *ServiceRegistriesServerTransport {
			return NewServiceRegistriesServerTransport(&s.srv.ServiceRegistriesServer)
		})
		resp, err = s.trServiceRegistriesServer.Do(req)
	case "ServicesClient":
		initServer(s, &s.trServicesServer, func() *ServicesServerTransport { return NewServicesServerTransport(&s.srv.ServicesServer) })
		resp, err = s.trServicesServer.Do(req)
	case "StoragesClient":
		initServer(s, &s.trStoragesServer, func() *StoragesServerTransport { return NewStoragesServerTransport(&s.srv.StoragesServer) })
		resp, err = s.trStoragesServer.Do(req)
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
