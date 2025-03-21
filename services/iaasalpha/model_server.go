/*
IaaS-API

This API allows you to create and modify IaaS resources.

API version: 1alpha1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaasalpha

import (
	"encoding/json"
	"time"
)

// checks if the Server type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Server{}

// Server Representation of a single server object.
type Server struct {
	// Universally Unique Identifier (UUID).
	AffinityGroup *string `json:"affinityGroup,omitempty"`
	// Object that represents an availability zone.
	AvailabilityZone *string                        `json:"availabilityZone,omitempty"`
	BootVolume       *CreateServerPayloadBootVolume `json:"bootVolume,omitempty"`
	// Date-time when resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// An error message.
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Universally Unique Identifier (UUID).
	Id *string `json:"id,omitempty"`
	// Universally Unique Identifier (UUID).
	ImageId *string `json:"imageId,omitempty"`
	// The name of an SSH keypair. Allowed characters are letters [a-zA-Z], digits [0-9] and the following special characters: [@._-].
	KeypairName *string `json:"keypairName,omitempty"`
	// Object that represents the labels of an object. Regex for keys: `^[a-z]((-|_|[a-z0-9])){0,62}$`. Regex for values: `^(-|_|[a-z0-9]){0,63}$`.
	Labels *map[string]interface{} `json:"labels,omitempty"`
	// Date-time when resource was launched.
	LaunchedAt *time.Time `json:"launchedAt,omitempty"`
	// The name for a General Object. Matches Names and also UUIDs.
	// REQUIRED
	MachineType       *string            `json:"machineType"`
	MaintenanceWindow *ServerMaintenance `json:"maintenanceWindow,omitempty"`
	// The name for a Server.
	// REQUIRED
	Name       *string                        `json:"name"`
	Networking *CreateServerPayloadNetworking `json:"networking,omitempty"`
	// A list of networks attached to a server.
	Nics *[]ServerNetwork `json:"nics,omitempty"`
	// The power status of a server. Possible values: `CRASHED`, `ERROR`, `RUNNING`, `STOPPED`.
	PowerStatus *string `json:"powerStatus,omitempty"`
	// A list of General Objects.
	SecurityGroups *[]string `json:"securityGroups,omitempty"`
	// A list of service account mails.
	ServiceAccountMails *[]string `json:"serviceAccountMails,omitempty"`
	// The status of a server object. Possible values: `ACTIVE`, `BACKING-UP`, `CREATING`, `DEALLOCATED`, `DEALLOCATING`, `DELETED`, `DELETING`, `ERROR`, `INACTIVE`, `MIGRATING`, `REBOOT`, `REBOOTING`, `REBUILD`, `REBUILDING`, `RESCUE`, `RESCUING`, `RESIZING`, `RESTORING`, `SNAPSHOTTING`, `STARTING`, `STOPPING`, `UNRESCUING`, `UPDATING`.
	Status *string `json:"status,omitempty"`
	// Date-time when resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// User Data that is provided to the server. Must be base64 encoded and is passed via cloud-init to the server.
	UserData *string `json:"userData,omitempty"`
	// A list of UUIDs.
	Volumes *[]string `json:"volumes,omitempty"`
}

type _Server Server

// NewServer instantiates a new Server object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServer(machineType *string, name *string) *Server {
	this := Server{}
	this.MachineType = machineType
	this.Name = name
	return &this
}

// NewServerWithDefaults instantiates a new Server object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerWithDefaults() *Server {
	this := Server{}
	return &this
}

// GetAffinityGroup returns the AffinityGroup field value if set, zero value otherwise.
func (o *Server) GetAffinityGroup() *string {
	if o == nil || IsNil(o.AffinityGroup) {
		var ret *string
		return ret
	}
	return o.AffinityGroup
}

// GetAffinityGroupOk returns a tuple with the AffinityGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetAffinityGroupOk() (*string, bool) {
	if o == nil || IsNil(o.AffinityGroup) {
		return nil, false
	}
	return o.AffinityGroup, true
}

// HasAffinityGroup returns a boolean if a field has been set.
func (o *Server) HasAffinityGroup() bool {
	if o != nil && !IsNil(o.AffinityGroup) {
		return true
	}

	return false
}

// SetAffinityGroup gets a reference to the given string and assigns it to the AffinityGroup field.
func (o *Server) SetAffinityGroup(v *string) {
	o.AffinityGroup = v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *Server) GetAvailabilityZone() *string {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret *string
		return ret
	}
	return o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetAvailabilityZoneOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// HasAvailabilityZone returns a boolean if a field has been set.
func (o *Server) HasAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given string and assigns it to the AvailabilityZone field.
func (o *Server) SetAvailabilityZone(v *string) {
	o.AvailabilityZone = v
}

// GetBootVolume returns the BootVolume field value if set, zero value otherwise.
func (o *Server) GetBootVolume() *CreateServerPayloadBootVolume {
	if o == nil || IsNil(o.BootVolume) {
		var ret *CreateServerPayloadBootVolume
		return ret
	}
	return o.BootVolume
}

// GetBootVolumeOk returns a tuple with the BootVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetBootVolumeOk() (*CreateServerPayloadBootVolume, bool) {
	if o == nil || IsNil(o.BootVolume) {
		return nil, false
	}
	return o.BootVolume, true
}

// HasBootVolume returns a boolean if a field has been set.
func (o *Server) HasBootVolume() bool {
	if o != nil && !IsNil(o.BootVolume) {
		return true
	}

	return false
}

// SetBootVolume gets a reference to the given CreateServerPayloadBootVolume and assigns it to the BootVolume field.
func (o *Server) SetBootVolume(v *CreateServerPayloadBootVolume) {
	o.BootVolume = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Server) GetCreatedAt() *time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret *time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Server) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Server) SetCreatedAt(v *time.Time) {
	o.CreatedAt = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *Server) GetErrorMessage() *string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret *string
		return ret
	}
	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *Server) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *Server) SetErrorMessage(v *string) {
	o.ErrorMessage = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Server) GetId() *string {
	if o == nil || IsNil(o.Id) {
		var ret *string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Server) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Server) SetId(v *string) {
	o.Id = v
}

// GetImageId returns the ImageId field value if set, zero value otherwise.
func (o *Server) GetImageId() *string {
	if o == nil || IsNil(o.ImageId) {
		var ret *string
		return ret
	}
	return o.ImageId
}

// GetImageIdOk returns a tuple with the ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageId) {
		return nil, false
	}
	return o.ImageId, true
}

// HasImageId returns a boolean if a field has been set.
func (o *Server) HasImageId() bool {
	if o != nil && !IsNil(o.ImageId) {
		return true
	}

	return false
}

// SetImageId gets a reference to the given string and assigns it to the ImageId field.
func (o *Server) SetImageId(v *string) {
	o.ImageId = v
}

// GetKeypairName returns the KeypairName field value if set, zero value otherwise.
func (o *Server) GetKeypairName() *string {
	if o == nil || IsNil(o.KeypairName) {
		var ret *string
		return ret
	}
	return o.KeypairName
}

// GetKeypairNameOk returns a tuple with the KeypairName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetKeypairNameOk() (*string, bool) {
	if o == nil || IsNil(o.KeypairName) {
		return nil, false
	}
	return o.KeypairName, true
}

// HasKeypairName returns a boolean if a field has been set.
func (o *Server) HasKeypairName() bool {
	if o != nil && !IsNil(o.KeypairName) {
		return true
	}

	return false
}

// SetKeypairName gets a reference to the given string and assigns it to the KeypairName field.
func (o *Server) SetKeypairName(v *string) {
	o.KeypairName = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Server) GetLabels() *map[string]interface{} {
	if o == nil || IsNil(o.Labels) {
		var ret *map[string]interface{}
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetLabelsOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.Labels) {
		return &map[string]interface{}{}, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Server) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]interface{} and assigns it to the Labels field.
func (o *Server) SetLabels(v *map[string]interface{}) {
	o.Labels = v
}

// GetLaunchedAt returns the LaunchedAt field value if set, zero value otherwise.
func (o *Server) GetLaunchedAt() *time.Time {
	if o == nil || IsNil(o.LaunchedAt) {
		var ret *time.Time
		return ret
	}
	return o.LaunchedAt
}

// GetLaunchedAtOk returns a tuple with the LaunchedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetLaunchedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LaunchedAt) {
		return nil, false
	}
	return o.LaunchedAt, true
}

// HasLaunchedAt returns a boolean if a field has been set.
func (o *Server) HasLaunchedAt() bool {
	if o != nil && !IsNil(o.LaunchedAt) {
		return true
	}

	return false
}

// SetLaunchedAt gets a reference to the given time.Time and assigns it to the LaunchedAt field.
func (o *Server) SetLaunchedAt(v *time.Time) {
	o.LaunchedAt = v
}

// GetMachineType returns the MachineType field value
func (o *Server) GetMachineType() *string {
	if o == nil || IsNil(o.MachineType) {
		var ret *string
		return ret
	}

	return o.MachineType
}

// GetMachineTypeOk returns a tuple with the MachineType field value
// and a boolean to check if the value has been set.
func (o *Server) GetMachineTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineType, true
}

// SetMachineType sets field value
func (o *Server) SetMachineType(v *string) {
	o.MachineType = v
}

// GetMaintenanceWindow returns the MaintenanceWindow field value if set, zero value otherwise.
func (o *Server) GetMaintenanceWindow() *ServerMaintenance {
	if o == nil || IsNil(o.MaintenanceWindow) {
		var ret *ServerMaintenance
		return ret
	}
	return o.MaintenanceWindow
}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetMaintenanceWindowOk() (*ServerMaintenance, bool) {
	if o == nil || IsNil(o.MaintenanceWindow) {
		return nil, false
	}
	return o.MaintenanceWindow, true
}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *Server) HasMaintenanceWindow() bool {
	if o != nil && !IsNil(o.MaintenanceWindow) {
		return true
	}

	return false
}

// SetMaintenanceWindow gets a reference to the given ServerMaintenance and assigns it to the MaintenanceWindow field.
func (o *Server) SetMaintenanceWindow(v *ServerMaintenance) {
	o.MaintenanceWindow = v
}

// GetName returns the Name field value
func (o *Server) GetName() *string {
	if o == nil || IsNil(o.Name) {
		var ret *string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Server) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *Server) SetName(v *string) {
	o.Name = v
}

// GetNetworking returns the Networking field value if set, zero value otherwise.
func (o *Server) GetNetworking() *CreateServerPayloadNetworking {
	if o == nil || IsNil(o.Networking) {
		var ret *CreateServerPayloadNetworking
		return ret
	}
	return o.Networking
}

// GetNetworkingOk returns a tuple with the Networking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetNetworkingOk() (*CreateServerPayloadNetworking, bool) {
	if o == nil || IsNil(o.Networking) {
		return nil, false
	}
	return o.Networking, true
}

// HasNetworking returns a boolean if a field has been set.
func (o *Server) HasNetworking() bool {
	if o != nil && !IsNil(o.Networking) {
		return true
	}

	return false
}

// SetNetworking gets a reference to the given CreateServerPayloadNetworking and assigns it to the Networking field.
func (o *Server) SetNetworking(v *CreateServerPayloadNetworking) {
	o.Networking = v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *Server) GetNics() *[]ServerNetwork {
	if o == nil || IsNil(o.Nics) {
		var ret *[]ServerNetwork
		return ret
	}
	return o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetNicsOk() (*[]ServerNetwork, bool) {
	if o == nil || IsNil(o.Nics) {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *Server) HasNics() bool {
	if o != nil && !IsNil(o.Nics) {
		return true
	}

	return false
}

// SetNics gets a reference to the given []ServerNetwork and assigns it to the Nics field.
func (o *Server) SetNics(v *[]ServerNetwork) {
	o.Nics = v
}

// GetPowerStatus returns the PowerStatus field value if set, zero value otherwise.
func (o *Server) GetPowerStatus() *string {
	if o == nil || IsNil(o.PowerStatus) {
		var ret *string
		return ret
	}
	return o.PowerStatus
}

// GetPowerStatusOk returns a tuple with the PowerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetPowerStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PowerStatus) {
		return nil, false
	}
	return o.PowerStatus, true
}

// HasPowerStatus returns a boolean if a field has been set.
func (o *Server) HasPowerStatus() bool {
	if o != nil && !IsNil(o.PowerStatus) {
		return true
	}

	return false
}

// SetPowerStatus gets a reference to the given string and assigns it to the PowerStatus field.
func (o *Server) SetPowerStatus(v *string) {
	o.PowerStatus = v
}

// GetSecurityGroups returns the SecurityGroups field value if set, zero value otherwise.
func (o *Server) GetSecurityGroups() *[]string {
	if o == nil || IsNil(o.SecurityGroups) {
		var ret *[]string
		return ret
	}
	return o.SecurityGroups
}

// GetSecurityGroupsOk returns a tuple with the SecurityGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetSecurityGroupsOk() (*[]string, bool) {
	if o == nil || IsNil(o.SecurityGroups) {
		return nil, false
	}
	return o.SecurityGroups, true
}

// HasSecurityGroups returns a boolean if a field has been set.
func (o *Server) HasSecurityGroups() bool {
	if o != nil && !IsNil(o.SecurityGroups) {
		return true
	}

	return false
}

// SetSecurityGroups gets a reference to the given []string and assigns it to the SecurityGroups field.
func (o *Server) SetSecurityGroups(v *[]string) {
	o.SecurityGroups = v
}

// GetServiceAccountMails returns the ServiceAccountMails field value if set, zero value otherwise.
func (o *Server) GetServiceAccountMails() *[]string {
	if o == nil || IsNil(o.ServiceAccountMails) {
		var ret *[]string
		return ret
	}
	return o.ServiceAccountMails
}

// GetServiceAccountMailsOk returns a tuple with the ServiceAccountMails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetServiceAccountMailsOk() (*[]string, bool) {
	if o == nil || IsNil(o.ServiceAccountMails) {
		return nil, false
	}
	return o.ServiceAccountMails, true
}

// HasServiceAccountMails returns a boolean if a field has been set.
func (o *Server) HasServiceAccountMails() bool {
	if o != nil && !IsNil(o.ServiceAccountMails) {
		return true
	}

	return false
}

// SetServiceAccountMails gets a reference to the given []string and assigns it to the ServiceAccountMails field.
func (o *Server) SetServiceAccountMails(v *[]string) {
	o.ServiceAccountMails = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Server) GetStatus() *string {
	if o == nil || IsNil(o.Status) {
		var ret *string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Server) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Server) SetStatus(v *string) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Server) GetUpdatedAt() *time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret *time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Server) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Server) SetUpdatedAt(v *time.Time) {
	o.UpdatedAt = v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *Server) GetUserData() *string {
	if o == nil || IsNil(o.UserData) {
		var ret *string
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetUserDataOk() (*string, bool) {
	if o == nil || IsNil(o.UserData) {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *Server) HasUserData() bool {
	if o != nil && !IsNil(o.UserData) {
		return true
	}

	return false
}

// SetUserData gets a reference to the given string and assigns it to the UserData field.
func (o *Server) SetUserData(v *string) {
	o.UserData = v
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *Server) GetVolumes() *[]string {
	if o == nil || IsNil(o.Volumes) {
		var ret *[]string
		return ret
	}
	return o.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Server) GetVolumesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Volumes) {
		return nil, false
	}
	return o.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *Server) HasVolumes() bool {
	if o != nil && !IsNil(o.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []string and assigns it to the Volumes field.
func (o *Server) SetVolumes(v *[]string) {
	o.Volumes = v
}

func (o Server) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffinityGroup) {
		toSerialize["affinityGroup"] = o.AffinityGroup
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.BootVolume) {
		toSerialize["bootVolume"] = o.BootVolume
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImageId) {
		toSerialize["imageId"] = o.ImageId
	}
	if !IsNil(o.KeypairName) {
		toSerialize["keypairName"] = o.KeypairName
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.LaunchedAt) {
		toSerialize["launchedAt"] = o.LaunchedAt
	}
	toSerialize["machineType"] = o.MachineType
	if !IsNil(o.MaintenanceWindow) {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Networking) {
		toSerialize["networking"] = o.Networking
	}
	if !IsNil(o.Nics) {
		toSerialize["nics"] = o.Nics
	}
	if !IsNil(o.PowerStatus) {
		toSerialize["powerStatus"] = o.PowerStatus
	}
	if !IsNil(o.SecurityGroups) {
		toSerialize["securityGroups"] = o.SecurityGroups
	}
	if !IsNil(o.ServiceAccountMails) {
		toSerialize["serviceAccountMails"] = o.ServiceAccountMails
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.UserData) {
		toSerialize["userData"] = o.UserData
	}
	if !IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableServer struct {
	value *Server
	isSet bool
}

func (v NullableServer) Get() *Server {
	return v.value
}

func (v *NullableServer) Set(val *Server) {
	v.value = val
	v.isSet = true
}

func (v NullableServer) IsSet() bool {
	return v.isSet
}

func (v *NullableServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServer(val *Server) *NullableServer {
	return &NullableServer{value: val, isSet: true}
}

func (v NullableServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
