package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemMembersRemovePostResponseable instead.
type ItemMembersRemoveResponse struct {
    ItemMembersRemovePostResponse
}
// NewItemMembersRemoveResponse instantiates a new ItemMembersRemoveResponse and sets the default values.
func NewItemMembersRemoveResponse()(*ItemMembersRemoveResponse) {
    m := &ItemMembersRemoveResponse{
        ItemMembersRemovePostResponse: *NewItemMembersRemovePostResponse(),
    }
    return m
}
// CreateItemMembersRemoveResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemMembersRemoveResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMembersRemoveResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemMembersRemovePostResponseable instead.
type ItemMembersRemoveResponseable interface {
    ItemMembersRemovePostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
