// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OperatingSystemVersionRange undocumented
type OperatingSystemVersionRange struct {
	// Object is the base model of OperatingSystemVersionRange
	Object
	// Description The description of this range (e.g. Valid 1702 builds)
	Description *string `json:"description,omitempty"`
	// LowestVersion The lowest inclusive version that this range contains.
	LowestVersion *string `json:"lowestVersion,omitempty"`
	// HighestVersion The highest inclusive version that this range contains.
	HighestVersion *string `json:"highestVersion,omitempty"`
}