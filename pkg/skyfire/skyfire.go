/*
Copyright (C) 2018 Black Duck Software, Inc.

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements. See the NOTICE file
distributed with this work for additional information
regarding copyright ownership. The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied. See the License for the
specific language governing permissions and limitations
under the License.
*/

package skyfire

import (
	"fmt"

	"github.com/blackducksoftware/perceptor-skyfire/pkg/hub"
	"github.com/blackducksoftware/perceptor-skyfire/pkg/kube"
	"github.com/blackducksoftware/perceptor-skyfire/pkg/perceptor"
	"github.com/blackducksoftware/perceptor-skyfire/pkg/report"
	log "github.com/sirupsen/logrus"
)

type Skyfire struct {
	Scraper           *Scraper
	LastPerceptorDump *perceptor.Dump
	LastHubDump       *hub.Dump
	LastKubeDump      *kube.Dump
	LastReport        *report.Report
}

func NewSkyfire(config *Config) (*Skyfire, error) {
	scraper, err := NewScraper(config)
	if err != nil {
		return nil, err
	}
	skyfire := &Skyfire{scraper, nil, nil, nil, nil}
	skyfire.HandleScrapes()
	return skyfire, nil
}

func (sf *Skyfire) HandleScrapes() {
	for {
		select {
		case h := <-sf.Scraper.HubDumps:
			fmt.Println(h)
			sf.LastHubDump = h
		case k := <-sf.Scraper.KubeDumps:
			fmt.Println(k)
			sf.LastKubeDump = k
		case p := <-sf.Scraper.PerceptorDumps:
			fmt.Println(p)
			sf.LastPerceptorDump = p
		}
	}
}

func (sf *Skyfire) BuildReport() {
	if sf.LastPerceptorDump == nil {
		log.Warnf("unable to generate report: perceptor dump is nil")
		return
	}
	if sf.LastHubDump == nil {
		log.Warnf("unable to generate report: hub dump is nil")
		return
	}
	if sf.LastKubeDump == nil {
		log.Warnf("unable to generate report: kube dump is nil")
		return
	}

	dump := report.NewDump(sf.LastKubeDump, sf.LastPerceptorDump, sf.LastHubDump)
	sf.LastReport = report.NewReport(dump)
	IssueReportMetrics(sf.LastReport)

	log.Infof("successfully built report")
}

func IssueReportMetrics(report *report.Report) {
	IssueHubReportMetrics(report.Hub)
	IssueKubeReportMetrics(report.Kube)
	IssuePerceptorHubMetrics(report.PerceptorHub)
	IssueKubePerceptorReportMetrics(report.KubePerceptor)
}

func IssueHubReportMetrics(report *report.HubReport) {
	recordReportProblem("hub_projects_multiple_versions", len(report.ProjectsMultipleVersions))
	recordReportProblem("hub_versions_multiple_code_locations", len(report.VersionsMultipleCodeLocations))
	recordReportProblem("hub_code_locations_multiple_scan_summaries", len(report.CodeLocationsMultipleScanSummaries))
}

func IssueKubeReportMetrics(report *report.KubeReport) {
	recordReportProblem("kube_unparseable_images", len(report.UnparseableKubeImages))
}

func IssueKubePerceptorReportMetrics(report *report.KubePerceptorReport) {
	recordReportProblem("kube-perceptor_images_just_in_kube", len(report.JustKubeImages))
	recordReportProblem("kube-perceptor_pods_just_in_kube", len(report.JustKubePods))
	recordReportProblem("kube-perceptor_images_just_in_perceptor", len(report.JustPerceptorImages))
	recordReportProblem("kube-perceptor_pods_just_in_perceptor", len(report.JustPerceptorPods))
	recordReportProblem("kube-perceptor_incorrect_pod_annotations", len(report.ConflictingAnnotationsPods))
	recordReportProblem("kube-perceptor_incorrect_pod_labels", len(report.ConflictingLabelsPods))
	recordReportProblem("kube-perceptor_finished_pods_just_kube", len(report.FinishedJustKubePods))
	recordReportProblem("kube-perceptor_finished_pods_just_perceptor", len(report.FinishedJustPerceptorPods))
	recordReportProblem("kube-perceptor_partially_annotated_pods", len(report.PartiallyAnnotatedKubePods))
	recordReportProblem("kube-perceptor_partially_labeled_pods", len(report.PartiallyLabeledKubePods))
}

func IssuePerceptorHubMetrics(report *report.PerceptorHubReport) {
	recordReportProblem("perceptor-hub_images_just_in_hub", len(report.JustHubImages))
	recordReportProblem("perceptor-hub_images_just_in_perceptor", len(report.JustPerceptorImages))
}
