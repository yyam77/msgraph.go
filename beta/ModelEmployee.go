// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Employee undocumented
type Employee struct {
	// Entity is the base model of Employee
	Entity
	// Number undocumented
	Number *string `json:"number,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GivenName undocumented
	GivenName *string `json:"givenName,omitempty"`
	// MiddleName undocumented
	MiddleName *string `json:"middleName,omitempty"`
	// Surname undocumented
	Surname *string `json:"surname,omitempty"`
	// JobTitle undocumented
	JobTitle *string `json:"jobTitle,omitempty"`
	// Address undocumented
	Address *PostalAddressType `json:"address,omitempty"`
	// PhoneNumber undocumented
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// MobilePhone undocumented
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// PersonalEmail undocumented
	PersonalEmail *string `json:"personalEmail,omitempty"`
	// EmploymentDate undocumented
	EmploymentDate *Date `json:"employmentDate,omitempty"`
	// TerminationDate undocumented
	TerminationDate *Date `json:"terminationDate,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// BirthDate undocumented
	BirthDate *Date `json:"birthDate,omitempty"`
	// StatisticsGroupCode undocumented
	StatisticsGroupCode *string `json:"statisticsGroupCode,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Picture undocumented
	Picture []Picture `json:"picture,omitempty"`
}