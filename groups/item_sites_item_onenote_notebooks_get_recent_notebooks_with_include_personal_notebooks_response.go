package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponseable instead.
type ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse struct {
    ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponse
}
// NewItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse instantiates a new ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse and sets the default values.
func NewItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse()(*ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse) {
    m := &ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse{
        ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponse: *NewItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponse(),
    }
    return m
}
// CreateItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponse(), nil
}
// Deprecated: This class is obsolete. Use ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponseable instead.
type ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksResponseable interface {
    ItemSitesItemOnenoteNotebooksGetRecentNotebooksWithIncludePersonalNotebooksGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
