package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse struct {
    ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaGetResponse
}
// NewItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse instantiates a new ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse and sets the default values.
func NewItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse()(*ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse) {
    m := &ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse{
        ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaGetResponse: *NewItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaGetResponse(),
    }
    return m
}
// CreateItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaResponseable interface {
    ItemJoinedTeamsItemChannelsItemMessagesItemRepliesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
