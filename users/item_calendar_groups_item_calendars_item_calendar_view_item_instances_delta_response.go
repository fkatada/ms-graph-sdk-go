package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse struct {
    ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaGetResponse
}
// NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse instantiates a new ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse and sets the default values.
func NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse()(*ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse) {
    m := &ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse{
        ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaGetResponse: *NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaGetResponse(),
    }
    return m
}
// CreateItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaResponseable interface {
    ItemCalendarGroupsItemCalendarsItemCalendarViewItemInstancesDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
