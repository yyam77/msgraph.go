// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ApprovalState undocumented
type ApprovalState string

const (
	// ApprovalStateVPending undocumented
	ApprovalStateVPending ApprovalState = "pending"
	// ApprovalStateVApproved undocumented
	ApprovalStateVApproved ApprovalState = "approved"
	// ApprovalStateVDenied undocumented
	ApprovalStateVDenied ApprovalState = "denied"
	// ApprovalStateVAborted undocumented
	ApprovalStateVAborted ApprovalState = "aborted"
	// ApprovalStateVCanceled undocumented
	ApprovalStateVCanceled ApprovalState = "canceled"
)

var (
	// ApprovalStatePPending is a pointer to ApprovalStateVPending
	ApprovalStatePPending = &_ApprovalStatePPending
	// ApprovalStatePApproved is a pointer to ApprovalStateVApproved
	ApprovalStatePApproved = &_ApprovalStatePApproved
	// ApprovalStatePDenied is a pointer to ApprovalStateVDenied
	ApprovalStatePDenied = &_ApprovalStatePDenied
	// ApprovalStatePAborted is a pointer to ApprovalStateVAborted
	ApprovalStatePAborted = &_ApprovalStatePAborted
	// ApprovalStatePCanceled is a pointer to ApprovalStateVCanceled
	ApprovalStatePCanceled = &_ApprovalStatePCanceled
)

var (
	_ApprovalStatePPending  = ApprovalStateVPending
	_ApprovalStatePApproved = ApprovalStateVApproved
	_ApprovalStatePDenied   = ApprovalStateVDenied
	_ApprovalStatePAborted  = ApprovalStateVAborted
	_ApprovalStatePCanceled = ApprovalStateVCanceled
)