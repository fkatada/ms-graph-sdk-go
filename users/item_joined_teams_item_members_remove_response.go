package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemMembersRemovePostResponseable instead.
type ItemJoinedTeamsItemMembersRemoveResponse struct {
    ItemJoinedTeamsItemMembersRemovePostResponse
}
// NewItemJoinedTeamsItemMembersRemoveResponse instantiates a new ItemJoinedTeamsItemMembersRemoveResponse and sets the default values.
func NewItemJoinedTeamsItemMembersRemoveResponse()(*ItemJoinedTeamsItemMembersRemoveResponse) {
    m := &ItemJoinedTeamsItemMembersRemoveResponse{
        ItemJoinedTeamsItemMembersRemovePostResponse: *NewItemJoinedTeamsItemMembersRemovePostResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemJoinedTeamsItemMembersRemovePostResponseable instead.
type ItemJoinedTeamsItemMembersRemoveResponseable interface {
    ItemJoinedTeamsItemMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
