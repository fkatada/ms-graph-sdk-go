package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse()(*ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse{
        ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemTablesItemColumnsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
