package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemChannelsItemMembersRemovePostResponseable instead.
type ItemJoinedTeamsItemChannelsItemMembersRemoveResponse struct {
    ItemJoinedTeamsItemChannelsItemMembersRemovePostResponse
}
// NewItemJoinedTeamsItemChannelsItemMembersRemoveResponse instantiates a new ItemJoinedTeamsItemChannelsItemMembersRemoveResponse and sets the default values.
func NewItemJoinedTeamsItemChannelsItemMembersRemoveResponse()(*ItemJoinedTeamsItemChannelsItemMembersRemoveResponse) {
    m := &ItemJoinedTeamsItemChannelsItemMembersRemoveResponse{
        ItemJoinedTeamsItemChannelsItemMembersRemovePostResponse: *NewItemJoinedTeamsItemChannelsItemMembersRemovePostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemChannelsItemMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemChannelsItemMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemChannelsItemMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemChannelsItemMembersRemovePostResponseable instead.
type ItemJoinedTeamsItemChannelsItemMembersRemoveResponseable interface {
    ItemJoinedTeamsItemChannelsItemMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
