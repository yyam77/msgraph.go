// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MediaDirection undocumented
type MediaDirection string

const (
	// MediaDirectionVInactive undocumented
	MediaDirectionVInactive MediaDirection = "inactive"
	// MediaDirectionVSendOnly undocumented
	MediaDirectionVSendOnly MediaDirection = "sendOnly"
	// MediaDirectionVReceiveOnly undocumented
	MediaDirectionVReceiveOnly MediaDirection = "receiveOnly"
	// MediaDirectionVSendReceive undocumented
	MediaDirectionVSendReceive MediaDirection = "sendReceive"
)

var (
	// MediaDirectionPInactive is a pointer to MediaDirectionVInactive
	MediaDirectionPInactive = &_MediaDirectionPInactive
	// MediaDirectionPSendOnly is a pointer to MediaDirectionVSendOnly
	MediaDirectionPSendOnly = &_MediaDirectionPSendOnly
	// MediaDirectionPReceiveOnly is a pointer to MediaDirectionVReceiveOnly
	MediaDirectionPReceiveOnly = &_MediaDirectionPReceiveOnly
	// MediaDirectionPSendReceive is a pointer to MediaDirectionVSendReceive
	MediaDirectionPSendReceive = &_MediaDirectionPSendReceive
)

var (
	_MediaDirectionPInactive    = MediaDirectionVInactive
	_MediaDirectionPSendOnly    = MediaDirectionVSendOnly
	_MediaDirectionPReceiveOnly = MediaDirectionVReceiveOnly
	_MediaDirectionPSendReceive = MediaDirectionVSendReceive
)

// MediaState undocumented
type MediaState string

const (
	// MediaStateVActive undocumented
	MediaStateVActive MediaState = "active"
	// MediaStateVInactive undocumented
	MediaStateVInactive MediaState = "inactive"
	// MediaStateVUnknownFutureValue undocumented
	MediaStateVUnknownFutureValue MediaState = "unknownFutureValue"
)

var (
	// MediaStatePActive is a pointer to MediaStateVActive
	MediaStatePActive = &_MediaStatePActive
	// MediaStatePInactive is a pointer to MediaStateVInactive
	MediaStatePInactive = &_MediaStatePInactive
	// MediaStatePUnknownFutureValue is a pointer to MediaStateVUnknownFutureValue
	MediaStatePUnknownFutureValue = &_MediaStatePUnknownFutureValue
)

var (
	_MediaStatePActive             = MediaStateVActive
	_MediaStatePInactive           = MediaStateVInactive
	_MediaStatePUnknownFutureValue = MediaStateVUnknownFutureValue
)
