package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type EdiscoverySearchExportOperation struct {
    CaseOperation
}
// NewEdiscoverySearchExportOperation instantiates a new EdiscoverySearchExportOperation and sets the default values.
func NewEdiscoverySearchExportOperation()(*EdiscoverySearchExportOperation) {
    m := &EdiscoverySearchExportOperation{
        CaseOperation: *NewCaseOperation(),
    }
    return m
}
// CreateEdiscoverySearchExportOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateEdiscoverySearchExportOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoverySearchExportOperation(), nil
}
// GetAdditionalOptions gets the additionalOptions property value. The additional items to include in the export. The possible values are: none, teamsAndYammerConversations, cloudAttachments, allDocumentVersions, subfolderContents, listAttachments, unknownFutureValue.
// returns a *AdditionalOptions when successful
func (m *EdiscoverySearchExportOperation) GetAdditionalOptions()(*AdditionalOptions) {
    val, err := m.GetBackingStore().Get("additionalOptions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*AdditionalOptions)
    }
    return nil
}
// GetDescription gets the description property value. The description of the export by the user.
// returns a *string when successful
func (m *EdiscoverySearchExportOperation) GetDescription()(*string) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetDisplayName gets the displayName property value. The name of export provided by the user.
// returns a *string when successful
func (m *EdiscoverySearchExportOperation) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetExportCriteria gets the exportCriteria property value. Items to be included in the export. The possible values are: searchHits, partiallyIndexed, unknownFutureValue.
// returns a *ExportCriteria when successful
func (m *EdiscoverySearchExportOperation) GetExportCriteria()(*ExportCriteria) {
    val, err := m.GetBackingStore().Get("exportCriteria")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExportCriteria)
    }
    return nil
}
// GetExportFileMetadata gets the exportFileMetadata property value. Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
// returns a []ExportFileMetadataable when successful
func (m *EdiscoverySearchExportOperation) GetExportFileMetadata()([]ExportFileMetadataable) {
    val, err := m.GetBackingStore().Get("exportFileMetadata")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ExportFileMetadataable)
    }
    return nil
}
// GetExportFormat gets the exportFormat property value. Format of the emails of the export. The possible values are: pst, msg, eml, unknownFutureValue.
// returns a *ExportFormat when successful
func (m *EdiscoverySearchExportOperation) GetExportFormat()(*ExportFormat) {
    val, err := m.GetBackingStore().Get("exportFormat")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExportFormat)
    }
    return nil
}
// GetExportLocation gets the exportLocation property value. Location scope for partially indexed items. You can choose to include partially indexed items only in responsive locations with search hits or in all targeted locations. The possible values are: responsiveLocations, nonresponsiveLocations, unknownFutureValue.
// returns a *ExportLocation when successful
func (m *EdiscoverySearchExportOperation) GetExportLocation()(*ExportLocation) {
    val, err := m.GetBackingStore().Get("exportLocation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ExportLocation)
    }
    return nil
}
// GetExportSingleItems gets the exportSingleItems property value. Indicates whether to export single items.
// returns a *bool when successful
func (m *EdiscoverySearchExportOperation) GetExportSingleItems()(*bool) {
    val, err := m.GetBackingStore().Get("exportSingleItems")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *EdiscoverySearchExportOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.CaseOperation.GetFieldDeserializers()
    res["additionalOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdditionalOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalOptions(val.(*AdditionalOptions))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["exportCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportCriteria)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportCriteria(val.(*ExportCriteria))
        }
        return nil
    }
    res["exportFileMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExportFileMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExportFileMetadataable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ExportFileMetadataable)
                }
            }
            m.SetExportFileMetadata(res)
        }
        return nil
    }
    res["exportFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportFormat)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportFormat(val.(*ExportFormat))
        }
        return nil
    }
    res["exportLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExportLocation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportLocation(val.(*ExportLocation))
        }
        return nil
    }
    res["exportSingleItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExportSingleItems(val)
        }
        return nil
    }
    res["search"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdiscoverySearchFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearch(val.(EdiscoverySearchable))
        }
        return nil
    }
    return res
}
// GetSearch gets the search property value. The eDiscovery searches under each case.
// returns a EdiscoverySearchable when successful
func (m *EdiscoverySearchExportOperation) GetSearch()(EdiscoverySearchable) {
    val, err := m.GetBackingStore().Get("search")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(EdiscoverySearchable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *EdiscoverySearchExportOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.CaseOperation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdditionalOptions() != nil {
        cast := (*m.GetAdditionalOptions()).String()
        err = writer.WriteStringValue("additionalOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExportCriteria() != nil {
        cast := (*m.GetExportCriteria()).String()
        err = writer.WriteStringValue("exportCriteria", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportFileMetadata() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExportFileMetadata()))
        for i, v := range m.GetExportFileMetadata() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("exportFileMetadata", cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportFormat() != nil {
        cast := (*m.GetExportFormat()).String()
        err = writer.WriteStringValue("exportFormat", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetExportLocation() != nil {
        cast := (*m.GetExportLocation()).String()
        err = writer.WriteStringValue("exportLocation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("exportSingleItems", m.GetExportSingleItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("search", m.GetSearch())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalOptions sets the additionalOptions property value. The additional items to include in the export. The possible values are: none, teamsAndYammerConversations, cloudAttachments, allDocumentVersions, subfolderContents, listAttachments, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetAdditionalOptions(value *AdditionalOptions)() {
    err := m.GetBackingStore().Set("additionalOptions", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. The description of the export by the user.
func (m *EdiscoverySearchExportOperation) SetDescription(value *string)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. The name of export provided by the user.
func (m *EdiscoverySearchExportOperation) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetExportCriteria sets the exportCriteria property value. Items to be included in the export. The possible values are: searchHits, partiallyIndexed, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetExportCriteria(value *ExportCriteria)() {
    err := m.GetBackingStore().Set("exportCriteria", value)
    if err != nil {
        panic(err)
    }
}
// SetExportFileMetadata sets the exportFileMetadata property value. Contains the properties for an export file metadata, including downloadUrl, fileName, and size.
func (m *EdiscoverySearchExportOperation) SetExportFileMetadata(value []ExportFileMetadataable)() {
    err := m.GetBackingStore().Set("exportFileMetadata", value)
    if err != nil {
        panic(err)
    }
}
// SetExportFormat sets the exportFormat property value. Format of the emails of the export. The possible values are: pst, msg, eml, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetExportFormat(value *ExportFormat)() {
    err := m.GetBackingStore().Set("exportFormat", value)
    if err != nil {
        panic(err)
    }
}
// SetExportLocation sets the exportLocation property value. Location scope for partially indexed items. You can choose to include partially indexed items only in responsive locations with search hits or in all targeted locations. The possible values are: responsiveLocations, nonresponsiveLocations, unknownFutureValue.
func (m *EdiscoverySearchExportOperation) SetExportLocation(value *ExportLocation)() {
    err := m.GetBackingStore().Set("exportLocation", value)
    if err != nil {
        panic(err)
    }
}
// SetExportSingleItems sets the exportSingleItems property value. Indicates whether to export single items.
func (m *EdiscoverySearchExportOperation) SetExportSingleItems(value *bool)() {
    err := m.GetBackingStore().Set("exportSingleItems", value)
    if err != nil {
        panic(err)
    }
}
// SetSearch sets the search property value. The eDiscovery searches under each case.
func (m *EdiscoverySearchExportOperation) SetSearch(value EdiscoverySearchable)() {
    err := m.GetBackingStore().Set("search", value)
    if err != nil {
        panic(err)
    }
}
type EdiscoverySearchExportOperationable interface {
    CaseOperationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalOptions()(*AdditionalOptions)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExportCriteria()(*ExportCriteria)
    GetExportFileMetadata()([]ExportFileMetadataable)
    GetExportFormat()(*ExportFormat)
    GetExportLocation()(*ExportLocation)
    GetExportSingleItems()(*bool)
    GetSearch()(EdiscoverySearchable)
    SetAdditionalOptions(value *AdditionalOptions)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExportCriteria(value *ExportCriteria)()
    SetExportFileMetadata(value []ExportFileMetadataable)()
    SetExportFormat(value *ExportFormat)()
    SetExportLocation(value *ExportLocation)()
    SetExportSingleItems(value *bool)()
    SetSearch(value EdiscoverySearchable)()
}
