package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemTeamChannelsItemMembersAddResponse struct {
    ItemTeamChannelsItemMembersAddPostResponse
}
// NewItemTeamChannelsItemMembersAddResponse instantiates a new ItemTeamChannelsItemMembersAddResponse and sets the default values.
func NewItemTeamChannelsItemMembersAddResponse()(*ItemTeamChannelsItemMembersAddResponse) {
    m := &ItemTeamChannelsItemMembersAddResponse{
        ItemTeamChannelsItemMembersAddPostResponse: *NewItemTeamChannelsItemMembersAddPostResponse(),
    }
    return m
}
// CreateItemTeamChannelsItemMembersAddResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemTeamChannelsItemMembersAddResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemTeamChannelsItemMembersAddResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemTeamChannelsItemMembersAddResponseable interface {
    ItemTeamChannelsItemMembersAddPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
