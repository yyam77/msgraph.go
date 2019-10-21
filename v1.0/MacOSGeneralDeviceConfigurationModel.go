// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSGeneralDeviceConfiguration This topic provides descriptions of the declared methods, properties and relationships exposed by the macOSGeneralDeviceConfiguration resource.
type MacOSGeneralDeviceConfiguration struct {
	// DeviceConfiguration is the base model of MacOSGeneralDeviceConfiguration
	DeviceConfiguration
	// CompliantAppsList List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
	CompliantAppsList []AppListItem `json:"compliantAppsList,omitempty"`
	// CompliantAppListType List that is in the CompliantAppsList.
	CompliantAppListType *AppListType `json:"compliantAppListType,omitempty"`
	// EmailInDomainSuffixes An email address lacking a suffix that matches any of these strings will be considered out-of-domain.
	EmailInDomainSuffixes []string `json:"emailInDomainSuffixes,omitempty"`
	// PasswordBlockSimple Block simple passwords.
	PasswordBlockSimple *bool `json:"passwordBlockSimple,omitempty"`
	// PasswordExpirationDays Number of days before the password expires.
	PasswordExpirationDays *int `json:"passwordExpirationDays,omitempty"`
	// PasswordMinimumCharacterSetCount Number of character sets a password must contain. Valid values 0 to 4
	PasswordMinimumCharacterSetCount *int `json:"passwordMinimumCharacterSetCount,omitempty"`
	// PasswordMinimumLength Minimum length of passwords.
	PasswordMinimumLength *int `json:"passwordMinimumLength,omitempty"`
	// PasswordMinutesOfInactivityBeforeLock Minutes of inactivity required before a password is required.
	PasswordMinutesOfInactivityBeforeLock *int `json:"passwordMinutesOfInactivityBeforeLock,omitempty"`
	// PasswordMinutesOfInactivityBeforeScreenTimeout Minutes of inactivity required before the screen times out.
	PasswordMinutesOfInactivityBeforeScreenTimeout *int `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	// PasswordPreviousPasswordBlockCount Number of previous passwords to block.
	PasswordPreviousPasswordBlockCount *int `json:"passwordPreviousPasswordBlockCount,omitempty"`
	// PasswordRequiredType Type of password that is required.
	PasswordRequiredType *RequiredPasswordType `json:"passwordRequiredType,omitempty"`
	// PasswordRequired Whether or not to require a password.
	PasswordRequired *bool `json:"passwordRequired,omitempty"`
}