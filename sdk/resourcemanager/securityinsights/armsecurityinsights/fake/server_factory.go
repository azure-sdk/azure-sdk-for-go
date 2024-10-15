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

// ServerFactory is a fake server for instances of the armsecurityinsights.ClientFactory type.
type ServerFactory struct {
	ActionsServer                            ActionsServer
	AlertRuleTemplatesServer                 AlertRuleTemplatesServer
	AlertRulesServer                         AlertRulesServer
	AutomationRulesServer                    AutomationRulesServer
	BookmarksServer                          BookmarksServer
	ContentPackageServer                     ContentPackageServer
	ContentPackagesServer                    ContentPackagesServer
	ContentTemplateServer                    ContentTemplateServer
	ContentTemplatesServer                   ContentTemplatesServer
	DataConnectorsServer                     DataConnectorsServer
	EntitiesServer                           EntitiesServer
	IncidentCommentsServer                   IncidentCommentsServer
	IncidentRelationsServer                  IncidentRelationsServer
	IncidentTasksServer                      IncidentTasksServer
	IncidentsServer                          IncidentsServer
	MetadataServer                           MetadataServer
	OperationsServer                         OperationsServer
	ProductPackageServer                     ProductPackageServer
	ProductPackagesServer                    ProductPackagesServer
	ProductTemplateServer                    ProductTemplateServer
	ProductTemplatesServer                   ProductTemplatesServer
	SecurityMLAnalyticsSettingsServer        SecurityMLAnalyticsSettingsServer
	SentinelOnboardingStatesServer           SentinelOnboardingStatesServer
	SourceControlServer                      SourceControlServer
	SourceControlsServer                     SourceControlsServer
	ThreatIntelligenceIndicatorServer        ThreatIntelligenceIndicatorServer
	ThreatIntelligenceIndicatorMetricsServer ThreatIntelligenceIndicatorMetricsServer
	ThreatIntelligenceIndicatorsServer       ThreatIntelligenceIndicatorsServer
	WatchlistItemsServer                     WatchlistItemsServer
	WatchlistsServer                         WatchlistsServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armsecurityinsights.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armsecurityinsights.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                        *ServerFactory
	trMu                                       sync.Mutex
	trActionsServer                            *ActionsServerTransport
	trAlertRuleTemplatesServer                 *AlertRuleTemplatesServerTransport
	trAlertRulesServer                         *AlertRulesServerTransport
	trAutomationRulesServer                    *AutomationRulesServerTransport
	trBookmarksServer                          *BookmarksServerTransport
	trContentPackageServer                     *ContentPackageServerTransport
	trContentPackagesServer                    *ContentPackagesServerTransport
	trContentTemplateServer                    *ContentTemplateServerTransport
	trContentTemplatesServer                   *ContentTemplatesServerTransport
	trDataConnectorsServer                     *DataConnectorsServerTransport
	trEntitiesServer                           *EntitiesServerTransport
	trIncidentCommentsServer                   *IncidentCommentsServerTransport
	trIncidentRelationsServer                  *IncidentRelationsServerTransport
	trIncidentTasksServer                      *IncidentTasksServerTransport
	trIncidentsServer                          *IncidentsServerTransport
	trMetadataServer                           *MetadataServerTransport
	trOperationsServer                         *OperationsServerTransport
	trProductPackageServer                     *ProductPackageServerTransport
	trProductPackagesServer                    *ProductPackagesServerTransport
	trProductTemplateServer                    *ProductTemplateServerTransport
	trProductTemplatesServer                   *ProductTemplatesServerTransport
	trSecurityMLAnalyticsSettingsServer        *SecurityMLAnalyticsSettingsServerTransport
	trSentinelOnboardingStatesServer           *SentinelOnboardingStatesServerTransport
	trSourceControlServer                      *SourceControlServerTransport
	trSourceControlsServer                     *SourceControlsServerTransport
	trThreatIntelligenceIndicatorServer        *ThreatIntelligenceIndicatorServerTransport
	trThreatIntelligenceIndicatorMetricsServer *ThreatIntelligenceIndicatorMetricsServerTransport
	trThreatIntelligenceIndicatorsServer       *ThreatIntelligenceIndicatorsServerTransport
	trWatchlistItemsServer                     *WatchlistItemsServerTransport
	trWatchlistsServer                         *WatchlistsServerTransport
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
	case "ActionsClient":
		initServer(s, &s.trActionsServer, func() *ActionsServerTransport { return NewActionsServerTransport(&s.srv.ActionsServer) })
		resp, err = s.trActionsServer.Do(req)
	case "AlertRuleTemplatesClient":
		initServer(s, &s.trAlertRuleTemplatesServer, func() *AlertRuleTemplatesServerTransport {
			return NewAlertRuleTemplatesServerTransport(&s.srv.AlertRuleTemplatesServer)
		})
		resp, err = s.trAlertRuleTemplatesServer.Do(req)
	case "AlertRulesClient":
		initServer(s, &s.trAlertRulesServer, func() *AlertRulesServerTransport { return NewAlertRulesServerTransport(&s.srv.AlertRulesServer) })
		resp, err = s.trAlertRulesServer.Do(req)
	case "AutomationRulesClient":
		initServer(s, &s.trAutomationRulesServer, func() *AutomationRulesServerTransport {
			return NewAutomationRulesServerTransport(&s.srv.AutomationRulesServer)
		})
		resp, err = s.trAutomationRulesServer.Do(req)
	case "BookmarksClient":
		initServer(s, &s.trBookmarksServer, func() *BookmarksServerTransport { return NewBookmarksServerTransport(&s.srv.BookmarksServer) })
		resp, err = s.trBookmarksServer.Do(req)
	case "ContentPackageClient":
		initServer(s, &s.trContentPackageServer, func() *ContentPackageServerTransport {
			return NewContentPackageServerTransport(&s.srv.ContentPackageServer)
		})
		resp, err = s.trContentPackageServer.Do(req)
	case "ContentPackagesClient":
		initServer(s, &s.trContentPackagesServer, func() *ContentPackagesServerTransport {
			return NewContentPackagesServerTransport(&s.srv.ContentPackagesServer)
		})
		resp, err = s.trContentPackagesServer.Do(req)
	case "ContentTemplateClient":
		initServer(s, &s.trContentTemplateServer, func() *ContentTemplateServerTransport {
			return NewContentTemplateServerTransport(&s.srv.ContentTemplateServer)
		})
		resp, err = s.trContentTemplateServer.Do(req)
	case "ContentTemplatesClient":
		initServer(s, &s.trContentTemplatesServer, func() *ContentTemplatesServerTransport {
			return NewContentTemplatesServerTransport(&s.srv.ContentTemplatesServer)
		})
		resp, err = s.trContentTemplatesServer.Do(req)
	case "DataConnectorsClient":
		initServer(s, &s.trDataConnectorsServer, func() *DataConnectorsServerTransport {
			return NewDataConnectorsServerTransport(&s.srv.DataConnectorsServer)
		})
		resp, err = s.trDataConnectorsServer.Do(req)
	case "EntitiesClient":
		initServer(s, &s.trEntitiesServer, func() *EntitiesServerTransport { return NewEntitiesServerTransport(&s.srv.EntitiesServer) })
		resp, err = s.trEntitiesServer.Do(req)
	case "IncidentCommentsClient":
		initServer(s, &s.trIncidentCommentsServer, func() *IncidentCommentsServerTransport {
			return NewIncidentCommentsServerTransport(&s.srv.IncidentCommentsServer)
		})
		resp, err = s.trIncidentCommentsServer.Do(req)
	case "IncidentRelationsClient":
		initServer(s, &s.trIncidentRelationsServer, func() *IncidentRelationsServerTransport {
			return NewIncidentRelationsServerTransport(&s.srv.IncidentRelationsServer)
		})
		resp, err = s.trIncidentRelationsServer.Do(req)
	case "IncidentTasksClient":
		initServer(s, &s.trIncidentTasksServer, func() *IncidentTasksServerTransport {
			return NewIncidentTasksServerTransport(&s.srv.IncidentTasksServer)
		})
		resp, err = s.trIncidentTasksServer.Do(req)
	case "IncidentsClient":
		initServer(s, &s.trIncidentsServer, func() *IncidentsServerTransport { return NewIncidentsServerTransport(&s.srv.IncidentsServer) })
		resp, err = s.trIncidentsServer.Do(req)
	case "MetadataClient":
		initServer(s, &s.trMetadataServer, func() *MetadataServerTransport { return NewMetadataServerTransport(&s.srv.MetadataServer) })
		resp, err = s.trMetadataServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "ProductPackageClient":
		initServer(s, &s.trProductPackageServer, func() *ProductPackageServerTransport {
			return NewProductPackageServerTransport(&s.srv.ProductPackageServer)
		})
		resp, err = s.trProductPackageServer.Do(req)
	case "ProductPackagesClient":
		initServer(s, &s.trProductPackagesServer, func() *ProductPackagesServerTransport {
			return NewProductPackagesServerTransport(&s.srv.ProductPackagesServer)
		})
		resp, err = s.trProductPackagesServer.Do(req)
	case "ProductTemplateClient":
		initServer(s, &s.trProductTemplateServer, func() *ProductTemplateServerTransport {
			return NewProductTemplateServerTransport(&s.srv.ProductTemplateServer)
		})
		resp, err = s.trProductTemplateServer.Do(req)
	case "ProductTemplatesClient":
		initServer(s, &s.trProductTemplatesServer, func() *ProductTemplatesServerTransport {
			return NewProductTemplatesServerTransport(&s.srv.ProductTemplatesServer)
		})
		resp, err = s.trProductTemplatesServer.Do(req)
	case "SecurityMLAnalyticsSettingsClient":
		initServer(s, &s.trSecurityMLAnalyticsSettingsServer, func() *SecurityMLAnalyticsSettingsServerTransport {
			return NewSecurityMLAnalyticsSettingsServerTransport(&s.srv.SecurityMLAnalyticsSettingsServer)
		})
		resp, err = s.trSecurityMLAnalyticsSettingsServer.Do(req)
	case "SentinelOnboardingStatesClient":
		initServer(s, &s.trSentinelOnboardingStatesServer, func() *SentinelOnboardingStatesServerTransport {
			return NewSentinelOnboardingStatesServerTransport(&s.srv.SentinelOnboardingStatesServer)
		})
		resp, err = s.trSentinelOnboardingStatesServer.Do(req)
	case "SourceControlClient":
		initServer(s, &s.trSourceControlServer, func() *SourceControlServerTransport {
			return NewSourceControlServerTransport(&s.srv.SourceControlServer)
		})
		resp, err = s.trSourceControlServer.Do(req)
	case "SourceControlsClient":
		initServer(s, &s.trSourceControlsServer, func() *SourceControlsServerTransport {
			return NewSourceControlsServerTransport(&s.srv.SourceControlsServer)
		})
		resp, err = s.trSourceControlsServer.Do(req)
	case "ThreatIntelligenceIndicatorClient":
		initServer(s, &s.trThreatIntelligenceIndicatorServer, func() *ThreatIntelligenceIndicatorServerTransport {
			return NewThreatIntelligenceIndicatorServerTransport(&s.srv.ThreatIntelligenceIndicatorServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorServer.Do(req)
	case "ThreatIntelligenceIndicatorMetricsClient":
		initServer(s, &s.trThreatIntelligenceIndicatorMetricsServer, func() *ThreatIntelligenceIndicatorMetricsServerTransport {
			return NewThreatIntelligenceIndicatorMetricsServerTransport(&s.srv.ThreatIntelligenceIndicatorMetricsServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorMetricsServer.Do(req)
	case "ThreatIntelligenceIndicatorsClient":
		initServer(s, &s.trThreatIntelligenceIndicatorsServer, func() *ThreatIntelligenceIndicatorsServerTransport {
			return NewThreatIntelligenceIndicatorsServerTransport(&s.srv.ThreatIntelligenceIndicatorsServer)
		})
		resp, err = s.trThreatIntelligenceIndicatorsServer.Do(req)
	case "WatchlistItemsClient":
		initServer(s, &s.trWatchlistItemsServer, func() *WatchlistItemsServerTransport {
			return NewWatchlistItemsServerTransport(&s.srv.WatchlistItemsServer)
		})
		resp, err = s.trWatchlistItemsServer.Do(req)
	case "WatchlistsClient":
		initServer(s, &s.trWatchlistsServer, func() *WatchlistsServerTransport { return NewWatchlistsServerTransport(&s.srv.WatchlistsServer) })
		resp, err = s.trWatchlistsServer.Do(req)
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
