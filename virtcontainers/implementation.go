// Copyright (c) 2017 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Description: The true virtcontainers function of the same name.
// This indirection is required to allow an alternative implemenation to be
// used for testing purposes.

package virtcontainers

import (
	"syscall"

	"github.com/sirupsen/logrus"
)

// VCImpl is the official virtcontainers function of the same name.
type VCImpl struct {
}

// SetLogger implements the VC function of the same name.
func (impl *VCImpl) SetLogger(logger logrus.FieldLogger) {
	SetLogger(logger)
}

// CreateSandbox implements the VC function of the same name.
func (impl *VCImpl) CreateSandbox(sandboxConfig SandboxConfig) (VCSandbox, error) {
	return CreateSandbox(sandboxConfig)
}

// DeleteSandbox implements the VC function of the same name.
func (impl *VCImpl) DeleteSandbox(sandboxID string) (VCSandbox, error) {
	return DeleteSandbox(sandboxID)
}

// StartSandbox implements the VC function of the same name.
func (impl *VCImpl) StartSandbox(sandboxID string) (VCSandbox, error) {
	return StartSandbox(sandboxID)
}

// StopSandbox implements the VC function of the same name.
func (impl *VCImpl) StopSandbox(sandboxID string) (VCSandbox, error) {
	return StopSandbox(sandboxID)
}

// RunSandbox implements the VC function of the same name.
func (impl *VCImpl) RunSandbox(sandboxConfig SandboxConfig) (VCSandbox, error) {
	return RunSandbox(sandboxConfig)
}

// ListSandbox implements the VC function of the same name.
func (impl *VCImpl) ListSandbox() ([]SandboxStatus, error) {
	return ListSandbox()
}

// StatusSandbox implements the VC function of the same name.
func (impl *VCImpl) StatusSandbox(sandboxID string) (SandboxStatus, error) {
	return StatusSandbox(sandboxID)
}

// PauseSandbox implements the VC function of the same name.
func (impl *VCImpl) PauseSandbox(sandboxID string) (VCSandbox, error) {
	return PauseSandbox(sandboxID)
}

// ResumeSandbox implements the VC function of the same name.
func (impl *VCImpl) ResumeSandbox(sandboxID string) (VCSandbox, error) {
	return ResumeSandbox(sandboxID)
}

// CreateContainer implements the VC function of the same name.
func (impl *VCImpl) CreateContainer(sandboxID string, containerConfig ContainerConfig) (VCSandbox, VCContainer, error) {
	return CreateContainer(sandboxID, containerConfig)
}

// DeleteContainer implements the VC function of the same name.
func (impl *VCImpl) DeleteContainer(sandboxID, containerID string) (VCContainer, error) {
	return DeleteContainer(sandboxID, containerID)
}

// StartContainer implements the VC function of the same name.
func (impl *VCImpl) StartContainer(sandboxID, containerID string) (VCContainer, error) {
	return StartContainer(sandboxID, containerID)
}

// StopContainer implements the VC function of the same name.
func (impl *VCImpl) StopContainer(sandboxID, containerID string) (VCContainer, error) {
	return StopContainer(sandboxID, containerID)
}

// EnterContainer implements the VC function of the same name.
func (impl *VCImpl) EnterContainer(sandboxID, containerID string, cmd Cmd) (VCSandbox, VCContainer, *Process, error) {
	return EnterContainer(sandboxID, containerID, cmd)
}

// StatusContainer implements the VC function of the same name.
func (impl *VCImpl) StatusContainer(sandboxID, containerID string) (ContainerStatus, error) {
	return StatusContainer(sandboxID, containerID)
}

// KillContainer implements the VC function of the same name.
func (impl *VCImpl) KillContainer(sandboxID, containerID string, signal syscall.Signal, all bool) error {
	return KillContainer(sandboxID, containerID, signal, all)
}

// ProcessListContainer implements the VC function of the same name.
func (impl *VCImpl) ProcessListContainer(sandboxID, containerID string, options ProcessListOptions) (ProcessList, error) {
	return ProcessListContainer(sandboxID, containerID, options)
}
