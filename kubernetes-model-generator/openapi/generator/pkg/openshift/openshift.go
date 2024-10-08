/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package openshift provides OpenShift related functionality
package openshift

// PackagePatterns are the patterns required for OpenShift GO and JSON OpenAPI generation
var PackagePatterns = []string{
	// Always import Kubernetes base packages as they are required by the rest of APIs
	"k8s.io/apimachinery/pkg/apis/meta/v...",
	"k8s.io/api/.../v...",
	// OpenShift APIs
	"github.com/openshift/api/.../v...",
	// CRDs don't provide info for reused types
	"github.com/metal3-io/baremetal-operator/apis/metal3.io/v...",
	"github.com/metal3-io/cluster-api-provider-metal3/api/v...",
	"github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/.../v...",
	"github.com/openshift/cloud-credential-operator/pkg/apis/.../v...",
	"github.com/openshift/cluster-network-operator/pkg/apis/.../v...",
	"github.com/openshift/cluster-node-tuning-operator/pkg/apis/tuned/v...",
	"github.com/openshift/hive/apis/hive/v...",
	"github.com/openshift/installer/pkg/types", // Add manually each package since some subpackages are problematic with go modules
	"github.com/openshift/installer/pkg/types/aws",
	"github.com/openshift/installer/pkg/types/azure",
	"github.com/openshift/installer/pkg/types/baremetal",
	"github.com/openshift/installer/pkg/types/external",
	"github.com/openshift/installer/pkg/types/gcp",
	"github.com/openshift/installer/pkg/types/ibmcloud",
	"github.com/openshift/installer/pkg/types/none",
	"github.com/openshift/installer/pkg/types/nutanix",
	"github.com/openshift/installer/pkg/types/openstack",
	"github.com/openshift/installer/pkg/types/ovirt",
	"github.com/openshift/installer/pkg/types/powervs",
	"github.com/openshift/installer/pkg/types/vsphere",
	"github.com/operator-framework/api/pkg/operators/v...",
	"github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators/v...",
	"github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v...",
	// Support types required by some APIs such as (github.com/openshift/hive)
	"github.com/openshift/custom-resource-status/conditions/v...",
}
