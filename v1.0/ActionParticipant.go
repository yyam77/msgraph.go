// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ParticipantCollectionInviteRequestParameter undocumented
type ParticipantCollectionInviteRequestParameter struct {
	// Participants undocumented
	Participants []InvitationParticipantInfo `json:"participants,omitempty"`
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}

// ParticipantMuteRequestParameter undocumented
type ParticipantMuteRequestParameter struct {
	// ClientContext undocumented
	ClientContext *string `json:"clientContext,omitempty"`
}
