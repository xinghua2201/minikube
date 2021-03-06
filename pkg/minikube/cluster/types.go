/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import "k8s.io/minikube/pkg/util"

// MachineConfig contains the parameters used to start a cluster.
type MachineConfig struct {
	MinikubeISO         string
	Memory              int
	CPUs                int
	DiskSize            int
	VMDriver            string
	XhyveDiskDriver     string   // Only used by the xhyve driver
	DockerEnv           []string // Each entry is formatted as KEY=VALUE.
	InsecureRegistry    []string
	RegistryMirror      []string
	HostOnlyCIDR        string // Only used by the virtualbox driver
	HypervVirtualSwitch string
	KvmNetwork          string             // Only used by the KVM driver
	Downloader          util.ISODownloader `json:"-"`
	DockerOpt           []string           // Each entry is formatted as KEY=VALUE.
	DisableDriverMounts bool               // Only used by virtualbox and xhyve
}

// KubernetesConfig contains the parameters used to configure the VM Kubernetes.
type KubernetesConfig struct {
	KubernetesVersion string
	NodeIP            string
	APIServerName     string
	DNSDomain         string
	ContainerRuntime  string
	NetworkPlugin     string
	FeatureGates      string
	ExtraOptions      util.ExtraOptionSlice
}

// Config contains machine and k8s config
type Config struct {
	MachineConfig    MachineConfig
	KubernetesConfig KubernetesConfig
}
