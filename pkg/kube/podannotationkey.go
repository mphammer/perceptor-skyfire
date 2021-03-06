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

package kube

import (
	"fmt"
)

type PodAnnotationKey int

const (
	PodAnnotationKeyVulnerabilities  PodAnnotationKey = iota
	PodAnnotationKeyPolicyViolations PodAnnotationKey = iota
	PodAnnotationKeyOverallStatus    PodAnnotationKey = iota
	PodAnnotationKeyServerVersion    PodAnnotationKey = iota
	PodAnnotationKeyScannerVersion   PodAnnotationKey = iota
)

func (pak PodAnnotationKey) String() string {
	switch pak {
	case PodAnnotationKeyVulnerabilities:
		return "pod.vulnerabilities"
	case PodAnnotationKeyPolicyViolations:
		return "pod.policy-violations"
	case PodAnnotationKeyOverallStatus:
		return "pod.overall-status"
	case PodAnnotationKeyServerVersion:
		return "pod.server-version"
	case PodAnnotationKeyScannerVersion:
		return "pod.scanner-version"
	}
	panic(fmt.Errorf("invalid PodAnnotationKey value: %d", pak))
}

var podAnnotationKeys = []PodAnnotationKey{
	PodAnnotationKeyVulnerabilities,
	PodAnnotationKeyPolicyViolations,
	PodAnnotationKeyOverallStatus,
	PodAnnotationKeyServerVersion,
	PodAnnotationKeyScannerVersion,
}

var podAnnotationKeyStrings = []string{}

func init() {
	for _, key := range podAnnotationKeys {
		podAnnotationKeyStrings = append(podAnnotationKeyStrings, key.String())
	}
}
