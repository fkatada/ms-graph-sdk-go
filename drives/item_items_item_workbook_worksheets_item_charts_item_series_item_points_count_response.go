package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse struct {
    ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountGetResponse
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse()(*ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse{
        ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountGetResponse: *NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountGetResponse(),
    }
    return m
}
// CreateItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponse(), nil
}
// Deprecated: This class is obsolete. Use {TypeName} instead.
type ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountResponseable interface {
    ItemItemsItemWorkbookWorksheetsItemChartsItemSeriesItemPointsCountGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
