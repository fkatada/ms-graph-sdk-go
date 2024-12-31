package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemPrimaryChannelMembersRemovePostResponseable instead.
type ItemPrimaryChannelMembersRemoveResponse struct {
    ItemPrimaryChannelMembersRemovePostResponse
}
// NewItemPrimaryChannelMembersRemoveResponse instantiates a new ItemPrimaryChannelMembersRemoveResponse and sets the default values.
func NewItemPrimaryChannelMembersRemoveResponse()(*ItemPrimaryChannelMembersRemoveResponse) {
    m := &ItemPrimaryChannelMembersRemoveResponse{
        ItemPrimaryChannelMembersRemovePostResponse: *NewItemPrimaryChannelMembersRemovePostResponse(),
    }
    return m
}
// CreateItemPrimaryChannelMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemPrimaryChannelMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPrimaryChannelMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemPrimaryChannelMembersRemovePostResponseable instead.
type ItemPrimaryChannelMembersRemoveResponseable interface {
    ItemPrimaryChannelMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
