// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// VPPLicensingType undocumented
type VPPLicensingType struct {
	// Object is the base model of VPPLicensingType
	Object
	// SupportUserLicensing Whether the program supports the user licensing type.
	SupportUserLicensing *bool `json:"supportUserLicensing,omitempty"`
	// SupportDeviceLicensing Whether the program supports the device licensing type.
	SupportDeviceLicensing *bool `json:"supportDeviceLicensing,omitempty"`
	// SupportsUserLicensing Whether the program supports the user licensing type.
	SupportsUserLicensing *bool `json:"supportsUserLicensing,omitempty"`
	// SupportsDeviceLicensing Whether the program supports the device licensing type.
	SupportsDeviceLicensing *bool `json:"supportsDeviceLicensing,omitempty"`
}

// VPPToken You purchase multiple licenses for iOS apps through the Apple Volume Purchase Program for Business or Education. This involves setting up an Apple VPP account from the Apple website and uploading the Apple VPP Business or Education token to Intune. You can then synchronize your volume purchase information with Intune and track your volume-purchased app use. You can upload multiple Apple VPP Business or Education tokens.
type VPPToken struct {
	// Entity is the base model of VPPToken
	Entity
	// OrganizationName The organization associated with the Apple Volume Purchase Program Token
	OrganizationName *string `json:"organizationName,omitempty"`
	// VPPTokenAccountType The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: `business`, `education`.
	VPPTokenAccountType *VPPTokenAccountType `json:"vppTokenAccountType,omitempty"`
	// AppleID The apple Id associated with the given Apple Volume Purchase Program Token.
	AppleID *string `json:"appleId,omitempty"`
	// ExpirationDateTime The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// LastSyncDateTime The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// Token The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token *string `json:"token,omitempty"`
	// LastModifiedDateTime Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// State Current state of the Apple Volume Purchase Program Token. Possible values are: `unknown`, `valid`, `expired`, `invalid`, `assignedToExternalMDM`.
	State *VPPTokenState `json:"state,omitempty"`
	// TokenActionResults The collection of statuses of the actions performed on the Apple Volume Purchase Program Token.
	TokenActionResults []VPPTokenActionResult `json:"tokenActionResults,omitempty"`
	// LastSyncStatus Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: `none`, `inProgress`, `completed`, `failed`.
	LastSyncStatus *VPPTokenSyncStatus `json:"lastSyncStatus,omitempty"`
	// AutomaticallyUpdateApps Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`
	// CountryOrRegion Whether or not apps for the VPP token will be automatically updated.
	CountryOrRegion *string `json:"countryOrRegion,omitempty"`
	// DataSharingConsentGranted Consent granted for data sharing with the Apple Volume Purchase Program.
	DataSharingConsentGranted *bool `json:"dataSharingConsentGranted,omitempty"`
	// DisplayName An admin specified token friendly name.
	DisplayName *string `json:"displayName,omitempty"`
	// LocationName Token location returned from Apple VPP.
	LocationName *string `json:"locationName,omitempty"`
	// ClaimTokenManagementFromExternalMDM Admin consent to allow claiming token management from external MDM.
	ClaimTokenManagementFromExternalMDM *bool `json:"claimTokenManagementFromExternalMdm,omitempty"`
	// RoleScopeTagIDs Role Scope Tags IDs assigned to this entity.
	RoleScopeTagIDs []string `json:"roleScopeTagIds,omitempty"`
}

// VPPTokenActionResult undocumented
type VPPTokenActionResult struct {
	// Object is the base model of VPPTokenActionResult
	Object
	// ActionName Action name
	ActionName *string `json:"actionName,omitempty"`
	// ActionState State of the action
	ActionState *ActionState `json:"actionState,omitempty"`
	// StartDateTime Time the action was initiated
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// LastUpdatedDateTime Time the action state was last updated
	LastUpdatedDateTime *time.Time `json:"lastUpdatedDateTime,omitempty"`
}

// VPPTokenLicenseSummary undocumented
type VPPTokenLicenseSummary struct {
	// Object is the base model of VPPTokenLicenseSummary
	Object
	// VPPTokenID Identifier of the VPP token.
	VPPTokenID *string `json:"vppTokenId,omitempty"`
	// AppleID The Apple Id associated with the given Apple Volume Purchase Program Token.
	AppleID *string `json:"appleId,omitempty"`
	// OrganizationName The organization associated with the Apple Volume Purchase Program Token.
	OrganizationName *string `json:"organizationName,omitempty"`
	// AvailableLicenseCount The number of VPP licenses available.
	AvailableLicenseCount *int `json:"availableLicenseCount,omitempty"`
	// UsedLicenseCount The number of VPP licenses in use.
	UsedLicenseCount *int `json:"usedLicenseCount,omitempty"`
}

// VPPTokenRevokeLicensesActionResult undocumented
type VPPTokenRevokeLicensesActionResult struct {
	// VPPTokenActionResult is the base model of VPPTokenRevokeLicensesActionResult
	VPPTokenActionResult
	// TotalLicensesCount A count of the number of licenses that were attempted to revoke.
	TotalLicensesCount *int `json:"totalLicensesCount,omitempty"`
	// FailedLicensesCount A count of the number of licenses that failed to revoke.
	FailedLicensesCount *int `json:"failedLicensesCount,omitempty"`
	// ActionFailureReason The reason for the revoke licenses action failure.
	ActionFailureReason *VPPTokenActionFailureReason `json:"actionFailureReason,omitempty"`
}