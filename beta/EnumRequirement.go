// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// RequirementProvider undocumented
type RequirementProvider string

const (
	// RequirementProviderVMFA undocumented
	RequirementProviderVMFA RequirementProvider = "MFA"
	// RequirementProviderVCA undocumented
	RequirementProviderVCA RequirementProvider = "CA"
	// RequirementProviderVUnknownFutureValue undocumented
	RequirementProviderVUnknownFutureValue RequirementProvider = "unknownFutureValue"
)

var (
	// RequirementProviderPMFA is a pointer to RequirementProviderVMFA
	RequirementProviderPMFA = &_RequirementProviderPMFA
	// RequirementProviderPCA is a pointer to RequirementProviderVCA
	RequirementProviderPCA = &_RequirementProviderPCA
	// RequirementProviderPUnknownFutureValue is a pointer to RequirementProviderVUnknownFutureValue
	RequirementProviderPUnknownFutureValue = &_RequirementProviderPUnknownFutureValue
)

var (
	_RequirementProviderPMFA                = RequirementProviderVMFA
	_RequirementProviderPCA                 = RequirementProviderVCA
	_RequirementProviderPUnknownFutureValue = RequirementProviderVUnknownFutureValue
)