//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package service

import "github.com/GoogleCloudPlatform/compute-image-tools/proto/go/pb"

// logRequest is a server-side pre-defined data structure
type logRequest struct {
	ClientInfo    clientInfo `json:"client_info"`
	LogSource     int64      `json:"log_source"`
	RequestTimeMs int64      `json:"request_time_ms"`
	LogEvent      []logEvent `json:"log_event"`
}

// ClientInfo is a server-side pre-defined data structure
type clientInfo struct {
	// ClientType is defined on server side to clarify which client library is used.
	ClientType string `json:"client_type"`
}

// LogEvent is a server-side pre-defined data structure
type logEvent struct {
	EventTimeMs         int64  `json:"event_time_ms"`
	EventUptimeMs       int64  `json:"event_uptime_ms"`
	SourceExtensionJSON string `json:"source_extension_json"`
}

// logResponse is a server-side pre-defined data structure
type logResponse struct {
	NextRequestWaitMillis int64                `json:"NextRequestWaitMillis,string"`
	LogResponseDetails    []logResponseDetails `json:"LogResponseDetails"`
}

// LogResponseDetails is a server-side pre-defined data structure
type logResponseDetails struct {
	ResponseAction responseAction `json:"ResponseAction"`
}

// ResponseAction is a server-side pre-defined data structure
type responseAction string

const (
	// responseActionUnknown - If the client sees this, it should delete the logRequest (not retry).
	// It may indicate that a new response action was added, which the client
	// doesn't yet understand.  (Deleting rather than retrying will prevent
	// infinite loops.)  The server will do whatever it can to prevent this
	// occurring (by not indicating an action to clients that are behind the
	// requisite version for the action).
	responseActionUnknown responseAction = "RESPONSE_ACTION_UNKNOWN"
	// retryRequestLater - The client should retry the request later, via normal scheduling.
	retryRequestLater responseAction = "RETRY_REQUEST_LATER"
	// deleteRequest - The client should delete the request.  This action will apply for
	// successful requests, and non-retryable requests.
	deleteRequest responseAction = "DELETE_REQUEST"
)

// ComputeImageToolsLogExtension contains all log info, which should be align with sawmill server side configuration.
type ComputeImageToolsLogExtension struct {
	// This id is a random guid for correlation among multiple log lines of a single call
	ID            string       `json:"id"`
	CloudBuildID  string       `json:"cloud_build_id"`
	ToolAction    string       `json:"tool_action"`
	Status        string       `json:"status"`
	ElapsedTimeMs int64        `json:"elapsed_time_ms"`
	EventTimeMs   int64        `json:"event_time_ms"`
	InputParams   *InputParams `json:"input_params,omitempty"`
	OutputInfo    *OutputInfo  `json:"output_info,omitempty"`
}

// InputParams contains the union of all APIs' param info. To simplify logging service, we
// avoid defining different schemas for each API.
type InputParams struct {
	ImageImportParams        *ImageImportParams        `json:"image_import_input_params,omitempty"`
	ImageExportParams        *ImageExportParams        `json:"image_export_input_params,omitempty"`
	InstanceImportParams     *InstanceImportParams     `json:"instance_import_input_params,omitempty"`
	MachineImageImportParams *MachineImageImportParams `json:"machine_image_import_input_params,omitempty"`
	WindowsUpgradeParams     *WindowsUpgradeParams     `json:"windows_upgrade_input_params,omitempty"`
	OnestepImageImportParams *OnestepImageImportParams `json:"onestep_image_import_input_params,omitempty"`
	InstanceExportParams     *InstanceExportParams     `json:"instance_export_input_params,omitempty"`
	MachineImageExportParams *MachineImageExportParams `json:"machine_image_export_input_params,omitempty"`
}

// ImageImportParams contains all input params for image import
type ImageImportParams struct {
	*CommonParams

	ImageName             string                `json:"image_name,omitempty"`
	DataDisk              bool                  `json:"data_disk"`
	OS                    string                `json:"os,omitempty"`
	SourceFile            string                `json:"source_file,omitempty"`
	SourceImage           string                `json:"source_image,omitempty"`
	NoGuestEnvironment    bool                  `json:"no_guest_environment"`
	Family                string                `json:"family,omitempty"`
	Description           string                `json:"description,omitempty"`
	NoExternalIP          bool                  `json:"no_external_ip"`
	HasKmsKey             bool                  `json:"has_kms_key"`
	HasKmsKeyring         bool                  `json:"has_kms_keyring"`
	HasKmsLocation        bool                  `json:"has_kms_location"`
	HasKmsProject         bool                  `json:"has_kms_project"`
	StorageLocation       string                `json:"storage_location,omitempty"`
	InspectionResults     *pb.InspectionResults `json:"inspection_results,omitempty"`
	ComputeServiceAccount string                `json:"compute_service_account,omitempty"`
}

// ImageExportParams contains all input params for image export
type ImageExportParams struct {
	*CommonParams

	DestinationURI        string `json:"destination_uri,omitempty"`
	SourceImage           string `json:"source_image,omitempty"`
	Format                string `json:"format,omitempty"`
	ComputeServiceAccount string `json:"compute_service_account,omitempty"`
	SourceDiskSnapshot    string `json:"source_disk_snapshot,omitempty"`
}

// OnestepImageImportParams contains all input params for onestep image import
type OnestepImageImportParams struct {
	*CommonParams

	// Image import params
	ImageName             string `json:"image_name,omitempty"`
	OS                    string `json:"os,omitempty"`
	NoGuestEnvironment    bool   `json:"no_guest_environment"`
	Family                string `json:"family,omitempty"`
	Description           string `json:"description,omitempty"`
	NoExternalIP          bool   `json:"no_external_ip"`
	HasKmsKey             bool   `json:"has_kms_key"`
	HasKmsKeyring         bool   `json:"has_kms_keyring"`
	HasKmsLocation        bool   `json:"has_kms_location"`
	HasKmsProject         bool   `json:"has_kms_project"`
	StorageLocation       string `json:"storage_location,omitempty"`
	ComputeServiceAccount string `json:"compute_service_account,omitempty"`

	// AWS related params
	AWSAMIID             string `json:"aws_ami_id,omitempty"`
	AWSAMIExportLocation string `json:"aws_ami_export_location,omitempty"`
	AWSSourceAMIFilePath string `json:"aws_source_ami_file_path,omitempty"`
}

// InstanceImportParams contains all input params for instance import
type InstanceImportParams struct {
	*CommonParams

	InstanceName                string `json:"instance_name,omitempty"`
	OvfGcsPath                  string `json:"ovf_gcs_path,omitempty"`
	CanIPForward                bool   `json:"can_ip_forward"`
	DeletionProtection          bool   `json:"deletion_protection"`
	MachineType                 string `json:"machine_type,omitempty"`
	NetworkInterface            string `json:"network_interface,omitempty"`
	NetworkTier                 string `json:"network_tier,omitempty"`
	PrivateNetworkIP            string `json:"private_network_ip,omitempty"`
	NoExternalIP                bool   `json:"no_external_ip,omitempty"`
	NoRestartOnFailure          bool   `json:"no_restart_on_failure"`
	OS                          string `json:"os,omitempty"`
	ShieldedIntegrityMonitoring bool   `json:"shielded_integrity_monitoring"`
	ShieldedSecureBoot          bool   `json:"shielded_secure_boot"`
	ShieldedVtpm                bool   `json:"shielded_vtpm"`
	Tags                        string `json:"tags,omitempty"`
	HasBootDiskKmsKey           bool   `json:"has_boot_disk_kms_key"`
	HasBootDiskKmsKeyring       bool   `json:"has_boot_disk_kms_keyring"`
	HasBootDiskKmsLocation      bool   `json:"has_boot_disk_kms_location"`
	HasBootDiskKmsProject       bool   `json:"has_boot_disk_kms_project"`
	NoGuestEnvironment          bool   `json:"no_guest_environment"`
	NodeAffinityLabel           string `json:"node_affinity_label,omitempty"`
	ComputeServiceAccount       string `json:"compute_service_account,omitempty"`
}

// MachineImageImportParams contains all input params for machine image import
type MachineImageImportParams struct {
	*CommonParams

	MachineImageName            string `json:"machine_image_name,omitempty"`
	OvfGcsPath                  string `json:"ovf_gcs_path,omitempty"`
	CanIPForward                bool   `json:"can_ip_forward"`
	DeletionProtection          bool   `json:"deletion_protection"`
	MachineType                 string `json:"machine_type,omitempty"`
	NetworkInterface            string `json:"network_interface,omitempty"`
	NetworkTier                 string `json:"network_tier,omitempty"`
	PrivateNetworkIP            string `json:"private_network_ip,omitempty"`
	NoExternalIP                bool   `json:"no_external_ip,omitempty"`
	NoRestartOnFailure          bool   `json:"no_restart_on_failure"`
	OS                          string `json:"os,omitempty"`
	ShieldedIntegrityMonitoring bool   `json:"shielded_integrity_monitoring"`
	ShieldedSecureBoot          bool   `json:"shielded_secure_boot"`
	ShieldedVtpm                bool   `json:"shielded_vtpm"`
	Tags                        string `json:"tags,omitempty"`
	HasBootDiskKmsKey           bool   `json:"has_boot_disk_kms_key"`
	HasBootDiskKmsKeyring       bool   `json:"has_boot_disk_kms_keyring"`
	HasBootDiskKmsLocation      bool   `json:"has_boot_disk_kms_location"`
	HasBootDiskKmsProject       bool   `json:"has_boot_disk_kms_project"`
	NoGuestEnvironment          bool   `json:"no_guest_environment"`
	NodeAffinityLabel           string `json:"node_affinity_label,omitempty"`
	Hostname                    string `json:"hostname,omitempty"`
	MachineImageStorageLocation string `json:"machine_image_storage_location,omitempty"`
	ComputeServiceAccount       string `json:"compute_service_account,omitempty"`
}

// CommonParams is only used to organize the code without impacting hierarchy of data
type CommonParams struct {
	ClientID                string `json:"client_id,omitempty"`
	ClientVersion           string `json:"client_version,omitempty"`
	Network                 string `json:"network,omitempty"`
	Subnet                  string `json:"subnet,omitempty"`
	Zone                    string `json:"zone,omitempty"`
	Timeout                 string `json:"timeout,omitempty"`
	Project                 string `json:"project,omitempty"`
	ObfuscatedProject       string `json:"obfuscated_project,omitempty"`
	Labels                  string `json:"labels,omitempty"`
	ScratchBucketGcsPath    string `json:"scratch_bucket_gcs_path,omitempty"`
	Oauth                   string `json:"oauth,omitempty"`
	ComputeEndpointOverride string `json:"compute_endpoint_override,omitempty"`
	DisableGcsLogging       bool   `json:"disable_gcs_logging"`
	DisableCloudLogging     bool   `json:"disable_cloud_logging"`
	DisableStdoutLogging    bool   `json:"disable_stdout_logging"`
}

// WindowsUpgradeParams contains all input params for windows upgrade
type WindowsUpgradeParams struct {
	*CommonParams

	SourceOS               string `json:"source_os,omitempty"`
	TargetOS               string `json:"target_os,omitempty"`
	Instance               string `json:"instance,omitempty"`
	CreateMachineBackup    bool   `json:"create_machine_backup"`
	AutoRollback           bool   `json:"auto_rollback"`
	UseStagingInstallMedia bool   `json:"use_staging_install_media"`
}

// InstanceExportParams contains all input params for instance export
type InstanceExportParams struct {
	*CommonParams

	DestinationURI   string `json:"destination_uri,omitempty"`
	InstanceName     string `json:"instance_name,omitempty"`
	OvfFormat        string `json:"ovf_format,omitempty"`
	DiskExportFormat string `json:"disk_export_format,omitempty"`
	NoExternalIP     bool   `json:"no_external_ip,omitempty"`
	OS               string `json:"os,omitempty"`
}

// MachineImageExportParams contains all input params for instance export
type MachineImageExportParams struct {
	*CommonParams

	DestinationURI   string `json:"destination_uri,omitempty"`
	MachineImageName string `json:"machine_image_name,omitempty"`
	OvfFormat        string `json:"ovf_format,omitempty"`
	DiskExportFormat string `json:"disk_export_format,omitempty"`
	NoExternalIP     bool   `json:"no_external_ip,omitempty"`
	OS               string `json:"os,omitempty"`
}

// OutputInfo contains output values from the tools execution
type OutputInfo struct {
	// Size of import/export sources (image or file)
	SourcesSizeGb []int64 `json:"sources_size_gb,omitempty"`
	// Size of import/export targets (image or file)
	TargetsSizeGb []int64 `json:"targets_size_gb,omitempty"`
	// Failure message of the command
	FailureMessage string `json:"failure_message,omitempty"`
	// Failure message of the command without privacy info
	FailureMessageWithoutPrivacyInfo string `json:"failure_message_without_privacy_info,omitempty"`
	// ImportFileFormat shows what is the actual image format of the imported file
	ImportFileFormat string `json:"import_file_format,omitempty"`
	// Serial output from worker instances; only populated
	// if workflow failed.
	SerialOutputs []string `json:"serial_outputs,omitempty"`
	// Inflation type (qemu, API, etc)
	InflationType string `json:"inflation_type,omitempty"`
	// Inflation time (seconds)
	InflationTime []int64 `json:"inflation_time_ms,omitempty"`
	// Inflation time (seconds) of the shadow disk
	ShadowInflationTime []int64 `json:"shadow_inflation_time_ms,omitempty"`
	// Shadow disk match result for shadow disk inflater
	ShadowDiskMatchResult string `json:"shadow_disk_match_result,omitempty"`
	// Indicates whether UEFI_COMPATIBLE was added to the image's guestOSFeatures, either due to inspection or user request
	IsUEFICompatibleImage bool `json:"is_uefi_compatible_image,omitempty"`
	// Indicates whether the image is auto-detected to be UEFI compatible
	IsUEFIDetected bool `json:"is_uefi_detected,omitempty"`
}

// InspectionResults contains metadata determined using automated inspection
type InspectionResults struct {
	// BIOSBootable indicates whether the disk is bootable with BIOS.
	BIOSBootable bool `json:"bios_bootable,omitempty"`

	// UEFIBootable indicates whether the disk is bootable with UEFI.
	UEFIBootable bool `json:"uefi_bootable,omitempty"`

	// RootFS indicates the file system type of the partition containing
	// the root directory ("/").
	RootFS string `json:"root_fs,omitempty"`
}

func (l *Logger) updateParams(projectPointer *string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if projectPointer == nil {
		return
	}

	project := *projectPointer
	obfuscatedProject := Hash(project)

	if l.Params.ImageImportParams != nil {
		l.Params.ImageImportParams.CommonParams.Project = project
		l.Params.ImageImportParams.CommonParams.ObfuscatedProject = obfuscatedProject
	}
	if l.Params.ImageExportParams != nil {
		l.Params.ImageExportParams.CommonParams.Project = project
		l.Params.ImageExportParams.CommonParams.ObfuscatedProject = obfuscatedProject
	}
	if l.Params.InstanceImportParams != nil {
		l.Params.InstanceImportParams.CommonParams.Project = project
		l.Params.InstanceImportParams.CommonParams.ObfuscatedProject = obfuscatedProject
	}
	if l.Params.MachineImageImportParams != nil {
		l.Params.MachineImageImportParams.CommonParams.Project = project
		l.Params.MachineImageImportParams.CommonParams.ObfuscatedProject = obfuscatedProject
	}
	if l.Params.OnestepImageImportParams != nil {
		l.Params.OnestepImageImportParams.CommonParams.Project = project
		l.Params.OnestepImageImportParams.CommonParams.ObfuscatedProject = obfuscatedProject
	}
}
