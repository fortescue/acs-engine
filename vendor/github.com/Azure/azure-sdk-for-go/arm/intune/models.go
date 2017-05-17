package intune

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// AppSharingFromLevel enumerates the values for app sharing from level.
type AppSharingFromLevel string

const (
	// AllApps specifies the all apps state for app sharing from level.
	AllApps AppSharingFromLevel = "allApps"
	// None specifies the none state for app sharing from level.
	None AppSharingFromLevel = "none"
	// PolicyManagedApps specifies the policy managed apps state for app
	// sharing from level.
	PolicyManagedApps AppSharingFromLevel = "policyManagedApps"
)

// AppSharingToLevel enumerates the values for app sharing to level.
type AppSharingToLevel string

const (
	// AppSharingToLevelAllApps specifies the app sharing to level all apps
	// state for app sharing to level.
	AppSharingToLevelAllApps AppSharingToLevel = "allApps"
	// AppSharingToLevelNone specifies the app sharing to level none state for
	// app sharing to level.
	AppSharingToLevelNone AppSharingToLevel = "none"
	// AppSharingToLevelPolicyManagedApps specifies the app sharing to level
	// policy managed apps state for app sharing to level.
	AppSharingToLevelPolicyManagedApps AppSharingToLevel = "policyManagedApps"
)

// Authentication enumerates the values for authentication.
type Authentication string

const (
	// NotRequired specifies the not required state for authentication.
	NotRequired Authentication = "notRequired"
	// Required specifies the required state for authentication.
	Required Authentication = "required"
)

// ClipboardSharingLevel enumerates the values for clipboard sharing level.
type ClipboardSharingLevel string

const (
	// ClipboardSharingLevelAllApps specifies the clipboard sharing level all
	// apps state for clipboard sharing level.
	ClipboardSharingLevelAllApps ClipboardSharingLevel = "allApps"
	// ClipboardSharingLevelBlocked specifies the clipboard sharing level
	// blocked state for clipboard sharing level.
	ClipboardSharingLevelBlocked ClipboardSharingLevel = "blocked"
	// ClipboardSharingLevelPolicyManagedApps specifies the clipboard sharing
	// level policy managed apps state for clipboard sharing level.
	ClipboardSharingLevelPolicyManagedApps ClipboardSharingLevel = "policyManagedApps"
	// ClipboardSharingLevelPolicyManagedAppsWithPasteIn specifies the
	// clipboard sharing level policy managed apps with paste in state for
	// clipboard sharing level.
	ClipboardSharingLevelPolicyManagedAppsWithPasteIn ClipboardSharingLevel = "policyManagedAppsWithPasteIn"
)

// DataBackup enumerates the values for data backup.
type DataBackup string

const (
	// Allow specifies the allow state for data backup.
	Allow DataBackup = "allow"
	// Block specifies the block state for data backup.
	Block DataBackup = "block"
)

// DeviceCompliance enumerates the values for device compliance.
type DeviceCompliance string

const (
	// Disable specifies the disable state for device compliance.
	Disable DeviceCompliance = "disable"
	// Enable specifies the enable state for device compliance.
	Enable DeviceCompliance = "enable"
)

// FileEncryption enumerates the values for file encryption.
type FileEncryption string

const (
	// FileEncryptionNotRequired specifies the file encryption not required
	// state for file encryption.
	FileEncryptionNotRequired FileEncryption = "notRequired"
	// FileEncryptionRequired specifies the file encryption required state for
	// file encryption.
	FileEncryptionRequired FileEncryption = "required"
)

// FileEncryptionLevel enumerates the values for file encryption level.
type FileEncryptionLevel string

const (
	// AfterDeviceRestart specifies the after device restart state for file
	// encryption level.
	AfterDeviceRestart FileEncryptionLevel = "afterDeviceRestart"
	// DeviceLocked specifies the device locked state for file encryption
	// level.
	DeviceLocked FileEncryptionLevel = "deviceLocked"
	// DeviceLockedExceptFilesOpen specifies the device locked except files
	// open state for file encryption level.
	DeviceLockedExceptFilesOpen FileEncryptionLevel = "deviceLockedExceptFilesOpen"
	// UseDeviceSettings specifies the use device settings state for file
	// encryption level.
	UseDeviceSettings FileEncryptionLevel = "useDeviceSettings"
)

// FileSharingSaveAs enumerates the values for file sharing save as.
type FileSharingSaveAs string

const (
	// FileSharingSaveAsAllow specifies the file sharing save as allow state
	// for file sharing save as.
	FileSharingSaveAsAllow FileSharingSaveAs = "allow"
	// FileSharingSaveAsBlock specifies the file sharing save as block state
	// for file sharing save as.
	FileSharingSaveAsBlock FileSharingSaveAs = "block"
)

// GroupStatus enumerates the values for group status.
type GroupStatus string

const (
	// NotTargeted specifies the not targeted state for group status.
	NotTargeted GroupStatus = "notTargeted"
	// Targeted specifies the targeted state for group status.
	Targeted GroupStatus = "targeted"
)

// ManagedBrowser enumerates the values for managed browser.
type ManagedBrowser string

const (
	// ManagedBrowserNotRequired specifies the managed browser not required
	// state for managed browser.
	ManagedBrowserNotRequired ManagedBrowser = "notRequired"
	// ManagedBrowserRequired specifies the managed browser required state for
	// managed browser.
	ManagedBrowserRequired ManagedBrowser = "required"
)

// Pin enumerates the values for pin.
type Pin string

const (
	// PinNotRequired specifies the pin not required state for pin.
	PinNotRequired Pin = "notRequired"
	// PinRequired specifies the pin required state for pin.
	PinRequired Pin = "required"
)

// Platform enumerates the values for platform.
type Platform string

const (
	// Android specifies the android state for platform.
	Android Platform = "android"
	// Ios specifies the ios state for platform.
	Ios Platform = "ios"
	// Windows specifies the windows state for platform.
	Windows Platform = "windows"
)

// ScreenCapture enumerates the values for screen capture.
type ScreenCapture string

const (
	// ScreenCaptureAllow specifies the screen capture allow state for screen
	// capture.
	ScreenCaptureAllow ScreenCapture = "allow"
	// ScreenCaptureBlock specifies the screen capture block state for screen
	// capture.
	ScreenCaptureBlock ScreenCapture = "block"
)

// TouchID enumerates the values for touch id.
type TouchID string

const (
	// TouchIDDisable specifies the touch id disable state for touch id.
	TouchIDDisable TouchID = "disable"
	// TouchIDEnable specifies the touch id enable state for touch id.
	TouchIDEnable TouchID = "enable"
)

// AndroidMAMPolicy is android Policy entity for Intune MAM.
type AndroidMAMPolicy struct {
	autorest.Response           `json:"-"`
	ID                          *string             `json:"id,omitempty"`
	Name                        *string             `json:"name,omitempty"`
	Type                        *string             `json:"type,omitempty"`
	Tags                        *map[string]*string `json:"tags,omitempty"`
	Location                    *string             `json:"location,omitempty"`
	*AndroidMAMPolicyProperties `json:"properties,omitempty"`
}

// AndroidMAMPolicyCollection is
type AndroidMAMPolicyCollection struct {
	autorest.Response `json:"-"`
	Value             *[]AndroidMAMPolicy `json:"value,omitempty"`
	Nextlink          *string             `json:"nextlink,omitempty"`
}

// AndroidMAMPolicyCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client AndroidMAMPolicyCollection) AndroidMAMPolicyCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// AndroidMAMPolicyProperties is intune MAM iOS Policy Properties.
type AndroidMAMPolicyProperties struct {
	FriendlyName                *string               `json:"friendlyName,omitempty"`
	Description                 *string               `json:"description,omitempty"`
	AppSharingFromLevel         AppSharingFromLevel   `json:"appSharingFromLevel,omitempty"`
	AppSharingToLevel           AppSharingToLevel     `json:"appSharingToLevel,omitempty"`
	Authentication              Authentication        `json:"authentication,omitempty"`
	ClipboardSharingLevel       ClipboardSharingLevel `json:"clipboardSharingLevel,omitempty"`
	DataBackup                  DataBackup            `json:"dataBackup,omitempty"`
	FileSharingSaveAs           FileSharingSaveAs     `json:"fileSharingSaveAs,omitempty"`
	Pin                         Pin                   `json:"pin,omitempty"`
	PinNumRetry                 *int32                `json:"pinNumRetry,omitempty"`
	DeviceCompliance            DeviceCompliance      `json:"deviceCompliance,omitempty"`
	ManagedBrowser              ManagedBrowser        `json:"managedBrowser,omitempty"`
	AccessRecheckOfflineTimeout *string               `json:"accessRecheckOfflineTimeout,omitempty"`
	AccessRecheckOnlineTimeout  *string               `json:"accessRecheckOnlineTimeout,omitempty"`
	OfflineWipeTimeout          *string               `json:"offlineWipeTimeout,omitempty"`
	NumOfApps                   *int32                `json:"numOfApps,omitempty"`
	GroupStatus                 GroupStatus           `json:"groupStatus,omitempty"`
	LastModifiedTime            *date.Time            `json:"lastModifiedTime,omitempty"`
	ScreenCapture               ScreenCapture         `json:"screenCapture,omitempty"`
	FileEncryption              FileEncryption        `json:"fileEncryption,omitempty"`
}

// Application is application entity for Intune MAM.
type Application struct {
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	Location               *string             `json:"location,omitempty"`
	*ApplicationProperties `json:"properties,omitempty"`
}

// ApplicationCollection is
type ApplicationCollection struct {
	autorest.Response `json:"-"`
	Value             *[]Application `json:"value,omitempty"`
	Nextlink          *string        `json:"nextlink,omitempty"`
}

// ApplicationCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ApplicationCollection) ApplicationCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// ApplicationProperties is
type ApplicationProperties struct {
	FriendlyName *string  `json:"friendlyName,omitempty"`
	Platform     Platform `json:"platform,omitempty"`
	AppID        *string  `json:"appId,omitempty"`
}

// Device is device entity for Intune.
type Device struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	Location          *string             `json:"location,omitempty"`
	*DeviceProperties `json:"properties,omitempty"`
}

// DeviceCollection is
type DeviceCollection struct {
	autorest.Response `json:"-"`
	Value             *[]Device `json:"value,omitempty"`
	Nextlink          *string   `json:"nextlink,omitempty"`
}

// DeviceCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client DeviceCollection) DeviceCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// DeviceProperties is
type DeviceProperties struct {
	UserID          *string `json:"userId,omitempty"`
	FriendlyName    *string `json:"friendlyName,omitempty"`
	Platform        *string `json:"platform,omitempty"`
	PlatformVersion *string `json:"platformVersion,omitempty"`
	DeviceType      *string `json:"deviceType,omitempty"`
}

// Error is
type Error struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// FlaggedEnrolledApp is flagged Enrolled App for the given tenant.
type FlaggedEnrolledApp struct {
	ID                            *string             `json:"id,omitempty"`
	Name                          *string             `json:"name,omitempty"`
	Type                          *string             `json:"type,omitempty"`
	Tags                          *map[string]*string `json:"tags,omitempty"`
	Location                      *string             `json:"location,omitempty"`
	*FlaggedEnrolledAppProperties `json:"properties,omitempty"`
}

// FlaggedEnrolledAppCollection is flagged Enrolled App collection for the
// given tenant.
type FlaggedEnrolledAppCollection struct {
	autorest.Response `json:"-"`
	Value             *[]FlaggedEnrolledApp `json:"value,omitempty"`
	Nextlink          *string               `json:"nextlink,omitempty"`
}

// FlaggedEnrolledAppCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client FlaggedEnrolledAppCollection) FlaggedEnrolledAppCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// FlaggedEnrolledAppError is
type FlaggedEnrolledAppError struct {
	ErrorCode *string `json:"errorCode,omitempty"`
	Severity  *string `json:"severity,omitempty"`
}

// FlaggedEnrolledAppProperties is
type FlaggedEnrolledAppProperties struct {
	DeviceType       *string                    `json:"deviceType,omitempty"`
	FriendlyName     *string                    `json:"friendlyName,omitempty"`
	LastModifiedTime *string                    `json:"lastModifiedTime,omitempty"`
	Platform         *string                    `json:"platform,omitempty"`
	Errors           *[]FlaggedEnrolledAppError `json:"errors,omitempty"`
}

// FlaggedUser is flagged user for the given tenant.
type FlaggedUser struct {
	autorest.Response      `json:"-"`
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	Location               *string             `json:"location,omitempty"`
	*FlaggedUserProperties `json:"properties,omitempty"`
}

// FlaggedUserCollection is flagged user collection for the given tenant.
type FlaggedUserCollection struct {
	autorest.Response `json:"-"`
	Value             *[]FlaggedUser `json:"value,omitempty"`
	Nextlink          *string        `json:"nextlink,omitempty"`
}

// FlaggedUserCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client FlaggedUserCollection) FlaggedUserCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// FlaggedUserProperties is
type FlaggedUserProperties struct {
	ErrorCount   *int32  `json:"errorCount,omitempty"`
	FriendlyName *string `json:"friendlyName,omitempty"`
}

// GroupItem is group entity for Intune MAM.
type GroupItem struct {
	ID               *string             `json:"id,omitempty"`
	Name             *string             `json:"name,omitempty"`
	Type             *string             `json:"type,omitempty"`
	Tags             *map[string]*string `json:"tags,omitempty"`
	Location         *string             `json:"location,omitempty"`
	*GroupProperties `json:"properties,omitempty"`
}

// GroupProperties is
type GroupProperties struct {
	FriendlyName *string `json:"friendlyName,omitempty"`
}

// GroupsCollection is
type GroupsCollection struct {
	autorest.Response `json:"-"`
	Value             *[]GroupItem `json:"value,omitempty"`
	Nextlink          *string      `json:"nextlink,omitempty"`
}

// GroupsCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client GroupsCollection) GroupsCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// IOSMAMPolicy is iOS Policy entity for Intune MAM.
type IOSMAMPolicy struct {
	autorest.Response       `json:"-"`
	ID                      *string             `json:"id,omitempty"`
	Name                    *string             `json:"name,omitempty"`
	Type                    *string             `json:"type,omitempty"`
	Tags                    *map[string]*string `json:"tags,omitempty"`
	Location                *string             `json:"location,omitempty"`
	*IOSMAMPolicyProperties `json:"properties,omitempty"`
}

// IOSMAMPolicyCollection is
type IOSMAMPolicyCollection struct {
	autorest.Response `json:"-"`
	Value             *[]IOSMAMPolicy `json:"value,omitempty"`
	Nextlink          *string         `json:"nextlink,omitempty"`
}

// IOSMAMPolicyCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client IOSMAMPolicyCollection) IOSMAMPolicyCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// IOSMAMPolicyProperties is intune MAM iOS Policy Properties.
type IOSMAMPolicyProperties struct {
	FriendlyName                *string               `json:"friendlyName,omitempty"`
	Description                 *string               `json:"description,omitempty"`
	AppSharingFromLevel         AppSharingFromLevel   `json:"appSharingFromLevel,omitempty"`
	AppSharingToLevel           AppSharingToLevel     `json:"appSharingToLevel,omitempty"`
	Authentication              Authentication        `json:"authentication,omitempty"`
	ClipboardSharingLevel       ClipboardSharingLevel `json:"clipboardSharingLevel,omitempty"`
	DataBackup                  DataBackup            `json:"dataBackup,omitempty"`
	FileSharingSaveAs           FileSharingSaveAs     `json:"fileSharingSaveAs,omitempty"`
	Pin                         Pin                   `json:"pin,omitempty"`
	PinNumRetry                 *int32                `json:"pinNumRetry,omitempty"`
	DeviceCompliance            DeviceCompliance      `json:"deviceCompliance,omitempty"`
	ManagedBrowser              ManagedBrowser        `json:"managedBrowser,omitempty"`
	AccessRecheckOfflineTimeout *string               `json:"accessRecheckOfflineTimeout,omitempty"`
	AccessRecheckOnlineTimeout  *string               `json:"accessRecheckOnlineTimeout,omitempty"`
	OfflineWipeTimeout          *string               `json:"offlineWipeTimeout,omitempty"`
	NumOfApps                   *int32                `json:"numOfApps,omitempty"`
	GroupStatus                 GroupStatus           `json:"groupStatus,omitempty"`
	LastModifiedTime            *date.Time            `json:"lastModifiedTime,omitempty"`
	FileEncryptionLevel         FileEncryptionLevel   `json:"fileEncryptionLevel,omitempty"`
	TouchID                     TouchID               `json:"touchId,omitempty"`
}

// Location is location entity for given tenant.
type Location struct {
	autorest.Response   `json:"-"`
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	Location            *string             `json:"location,omitempty"`
	*LocationProperties `json:"properties,omitempty"`
}

// LocationCollection is
type LocationCollection struct {
	autorest.Response `json:"-"`
	Value             *[]Location `json:"value,omitempty"`
	Nextlink          *string     `json:"nextlink,omitempty"`
}

// LocationCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client LocationCollection) LocationCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// LocationProperties is
type LocationProperties struct {
	HostName *string `json:"hostName,omitempty"`
}

// MAMPolicyAppIDOrGroupIDPayload is mAM Policy request body for properties
// Intune MAM.
type MAMPolicyAppIDOrGroupIDPayload struct {
	Properties *MAMPolicyAppOrGroupIDProperties `json:"properties,omitempty"`
}

// MAMPolicyAppOrGroupIDProperties is android Policy request body for Intune
// MAM.
type MAMPolicyAppOrGroupIDProperties struct {
	URL *string `json:"url,omitempty"`
}

// MAMPolicyProperties is
type MAMPolicyProperties struct {
	FriendlyName                *string               `json:"friendlyName,omitempty"`
	Description                 *string               `json:"description,omitempty"`
	AppSharingFromLevel         AppSharingFromLevel   `json:"appSharingFromLevel,omitempty"`
	AppSharingToLevel           AppSharingToLevel     `json:"appSharingToLevel,omitempty"`
	Authentication              Authentication        `json:"authentication,omitempty"`
	ClipboardSharingLevel       ClipboardSharingLevel `json:"clipboardSharingLevel,omitempty"`
	DataBackup                  DataBackup            `json:"dataBackup,omitempty"`
	FileSharingSaveAs           FileSharingSaveAs     `json:"fileSharingSaveAs,omitempty"`
	Pin                         Pin                   `json:"pin,omitempty"`
	PinNumRetry                 *int32                `json:"pinNumRetry,omitempty"`
	DeviceCompliance            DeviceCompliance      `json:"deviceCompliance,omitempty"`
	ManagedBrowser              ManagedBrowser        `json:"managedBrowser,omitempty"`
	AccessRecheckOfflineTimeout *string               `json:"accessRecheckOfflineTimeout,omitempty"`
	AccessRecheckOnlineTimeout  *string               `json:"accessRecheckOnlineTimeout,omitempty"`
	OfflineWipeTimeout          *string               `json:"offlineWipeTimeout,omitempty"`
	NumOfApps                   *int32                `json:"numOfApps,omitempty"`
	GroupStatus                 GroupStatus           `json:"groupStatus,omitempty"`
	LastModifiedTime            *date.Time            `json:"lastModifiedTime,omitempty"`
}

// OperationMetadataProperties is
type OperationMetadataProperties struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// OperationResult is operationResult entity for Intune.
type OperationResult struct {
	ID                         *string             `json:"id,omitempty"`
	Name                       *string             `json:"name,omitempty"`
	Type                       *string             `json:"type,omitempty"`
	Tags                       *map[string]*string `json:"tags,omitempty"`
	Location                   *string             `json:"location,omitempty"`
	*OperationResultProperties `json:"properties,omitempty"`
}

// OperationResultCollection is
type OperationResultCollection struct {
	autorest.Response `json:"-"`
	Value             *[]OperationResult `json:"value,omitempty"`
	Nextlink          *string            `json:"nextlink,omitempty"`
}

// OperationResultCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationResultCollection) OperationResultCollectionPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// OperationResultProperties is
type OperationResultProperties struct {
	FriendlyName      *string                        `json:"friendlyName,omitempty"`
	Category          *string                        `json:"category,omitempty"`
	LastModifiedTime  *string                        `json:"lastModifiedTime,omitempty"`
	State             *string                        `json:"state,omitempty"`
	OperationMetadata *[]OperationMetadataProperties `json:"operationMetadata,omitempty"`
}

// Resource is
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
	Location *string             `json:"location,omitempty"`
}

// StatusesDefault is default Statuses entity for the given tenant.
type StatusesDefault struct {
	autorest.Response   `json:"-"`
	ID                  *string             `json:"id,omitempty"`
	Name                *string             `json:"name,omitempty"`
	Type                *string             `json:"type,omitempty"`
	Tags                *map[string]*string `json:"tags,omitempty"`
	Location            *string             `json:"location,omitempty"`
	*StatusesProperties `json:"properties,omitempty"`
	Nextlink            *string `json:"nextlink,omitempty"`
}

// StatusesDefaultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client StatusesDefault) StatusesDefaultPreparer() (*http.Request, error) {
	if client.Nextlink == nil || len(to.String(client.Nextlink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.Nextlink)))
}

// StatusesProperties is
type StatusesProperties struct {
	DeployedPolicies   *int32     `json:"deployedPolicies,omitempty"`
	EnrolledUsers      *int32     `json:"enrolledUsers,omitempty"`
	FlaggedUsers       *int32     `json:"flaggedUsers,omitempty"`
	LastModifiedTime   *date.Time `json:"lastModifiedTime,omitempty"`
	PolicyAppliedUsers *int32     `json:"policyAppliedUsers,omitempty"`
	Status             *string    `json:"status,omitempty"`
	WipeFailedApps     *int32     `json:"wipeFailedApps,omitempty"`
	WipePendingApps    *int32     `json:"wipePendingApps,omitempty"`
	WipeSucceededApps  *int32     `json:"wipeSucceededApps,omitempty"`
}

// WipeDeviceOperationResult is device entity for Intune.
type WipeDeviceOperationResult struct {
	autorest.Response                    `json:"-"`
	ID                                   *string             `json:"id,omitempty"`
	Name                                 *string             `json:"name,omitempty"`
	Type                                 *string             `json:"type,omitempty"`
	Tags                                 *map[string]*string `json:"tags,omitempty"`
	Location                             *string             `json:"location,omitempty"`
	*WipeDeviceOperationResultProperties `json:"properties,omitempty"`
}

// WipeDeviceOperationResultProperties is
type WipeDeviceOperationResultProperties struct {
	Value *string `json:"value,omitempty"`
}
