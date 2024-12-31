package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemMembersRemovePostResponseable instead.
type DeletedTeamsItemChannelsItemMembersRemoveResponse struct {
    DeletedTeamsItemChannelsItemMembersRemovePostResponse
}
// NewDeletedTeamsItemChannelsItemMembersRemoveResponse instantiates a new DeletedTeamsItemChannelsItemMembersRemoveResponse and sets the default values.
func NewDeletedTeamsItemChannelsItemMembersRemoveResponse()(*DeletedTeamsItemChannelsItemMembersRemoveResponse) {
    m := &DeletedTeamsItemChannelsItemMembersRemoveResponse{
        DeletedTeamsItemChannelsItemMembersRemovePostResponse: *NewDeletedTeamsItemChannelsItemMembersRemovePostResponse(),
    }
    return m
}
// CreateDeletedTeamsItemChannelsItemMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDeletedTeamsItemChannelsItemMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedTeamsItemChannelsItemMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use DeletedTeamsItemChannelsItemMembersRemovePostResponseable instead.
type DeletedTeamsItemChannelsItemMembersRemoveResponseable interface {
    DeletedTeamsItemChannelsItemMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
