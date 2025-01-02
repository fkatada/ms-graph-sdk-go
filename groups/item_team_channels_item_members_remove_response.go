package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemTeamChannelsItemMembersRemovePostResponseable instead.
type ItemTeamChannelsItemMembersRemoveResponse struct {
    ItemTeamChannelsItemMembersRemovePostResponse
}
// NewItemTeamChannelsItemMembersRemoveResponse instantiates a new ItemTeamChannelsItemMembersRemoveResponse and sets the default values.
func NewItemTeamChannelsItemMembersRemoveResponse()(*ItemTeamChannelsItemMembersRemoveResponse) {
    m := &ItemTeamChannelsItemMembersRemoveResponse{
        ItemTeamChannelsItemMembersRemovePostResponse: *NewItemTeamChannelsItemMembersRemovePostResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemTeamChannelsItemMembersRemovePostResponseable instead.
type ItemTeamChannelsItemMembersRemoveResponseable interface {
    ItemTeamChannelsItemMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
