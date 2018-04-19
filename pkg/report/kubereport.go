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

package report

import (
	"fmt"

	"github.com/blackducksoftware/perceptor-skyfire/pkg/kube"
)

type KubeReport struct {
	UnparseableKubeImages []string
}

func NewKubeReport(dump *kube.Dump) *KubeReport {
	return &KubeReport{
		UnparseableKubeImages: UnparseableKubeImages(dump),
	}
}

func (k *KubeReport) HumanReadableString() string {
	return fmt.Sprintf(`
Kubernetes:
 - we found %d ImageIDs that were unparseable
`,
		len(k.UnparseableKubeImages))
}

func UnparseableKubeImages(dump *kube.Dump) []string {
	images := []string{}
	for _, image := range dump.ImagesMissingSha {
		images = append(images, image.ImageID)
	}
	return images
}
