package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppProtectionable 
type ManagedAppProtectionable interface {
    ManagedAppPolicyable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedDataStorageLocations()([]string)
    GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel)
    GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel)
    GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel)
    GetContactSyncBlocked()(*bool)
    GetDataBackupBlocked()(*bool)
    GetDeviceComplianceRequired()(*bool)
    GetDisableAppPinIfDevicePinIsSet()(*bool)
    GetFingerprintBlocked()(*bool)
    GetManagedBrowser()(*ManagedBrowserType)
    GetManagedBrowserToOpenLinksRequired()(*bool)
    GetMaximumPinRetries()(*int32)
    GetMinimumPinLength()(*int32)
    GetMinimumRequiredAppVersion()(*string)
    GetMinimumRequiredOsVersion()(*string)
    GetMinimumWarningAppVersion()(*string)
    GetMinimumWarningOsVersion()(*string)
    GetOrganizationalCredentialsRequired()(*bool)
    GetPeriodBeforePinReset()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOfflineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOfflineBeforeWipeIsEnforced()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPeriodOnlineBeforeAccessCheck()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)
    GetPinCharacterSet()(*ManagedAppPinCharacterSet)
    GetPinRequired()(*bool)
    GetPrintBlocked()(*bool)
    GetSaveAsBlocked()(*bool)
    GetSimplePinBlocked()(*bool)
    SetAllowedDataStorageLocations(value []string)()
    SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)()
    SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)()
    SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)()
    SetContactSyncBlocked(value *bool)()
    SetDataBackupBlocked(value *bool)()
    SetDeviceComplianceRequired(value *bool)()
    SetDisableAppPinIfDevicePinIsSet(value *bool)()
    SetFingerprintBlocked(value *bool)()
    SetManagedBrowser(value *ManagedBrowserType)()
    SetManagedBrowserToOpenLinksRequired(value *bool)()
    SetMaximumPinRetries(value *int32)()
    SetMinimumPinLength(value *int32)()
    SetMinimumRequiredAppVersion(value *string)()
    SetMinimumRequiredOsVersion(value *string)()
    SetMinimumWarningAppVersion(value *string)()
    SetMinimumWarningOsVersion(value *string)()
    SetOrganizationalCredentialsRequired(value *bool)()
    SetPeriodBeforePinReset(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOfflineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOfflineBeforeWipeIsEnforced(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPeriodOnlineBeforeAccessCheck(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)()
    SetPinCharacterSet(value *ManagedAppPinCharacterSet)()
    SetPinRequired(value *bool)()
    SetPrintBlocked(value *bool)()
    SetSaveAsBlocked(value *bool)()
    SetSimplePinBlocked(value *bool)()
}
