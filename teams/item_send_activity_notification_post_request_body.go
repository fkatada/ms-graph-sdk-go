package teams

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type ItemSendActivityNotificationPostRequestBody struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewItemSendActivityNotificationPostRequestBody instantiates a new ItemSendActivityNotificationPostRequestBody and sets the default values.
func NewItemSendActivityNotificationPostRequestBody()(*ItemSendActivityNotificationPostRequestBody) {
    m := &ItemSendActivityNotificationPostRequestBody{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSendActivityNotificationPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateItemSendActivityNotificationPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSendActivityNotificationPostRequestBody(), nil
}
// GetActivityType gets the activityType property value. The activityType property
// returns a *string when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetActivityType()(*string) {
    val, err := m.GetBackingStore().Get("activityType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetAdditionalData()(map[string]any) {
    val , err :=  m.backingStore.Get("additionalData")
    if err != nil {
        panic(err)
    }
    if val == nil {
        var value = make(map[string]any);
        m.SetAdditionalData(value);
    }
    return val.(map[string]any)
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetChainId gets the chainId property value. The chainId property
// returns a *int64 when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetChainId()(*int64) {
    val, err := m.GetBackingStore().Get("chainId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["chainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChainId(val)
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable))
        }
        return nil
    }
    res["recipient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkNotificationRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipient(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable))
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["templateParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)
                }
            }
            m.SetTemplateParameters(res)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkActivityTopicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable))
        }
        return nil
    }
    return res
}
// GetPreviewText gets the previewText property value. The previewText property
// returns a ItemBodyable when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetPreviewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable) {
    val, err := m.GetBackingStore().Get("previewText")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)
    }
    return nil
}
// GetRecipient gets the recipient property value. The recipient property
// returns a TeamworkNotificationRecipientable when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetRecipient()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable) {
    val, err := m.GetBackingStore().Get("recipient")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)
    }
    return nil
}
// GetTeamsAppId gets the teamsAppId property value. The teamsAppId property
// returns a *string when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetTeamsAppId()(*string) {
    val, err := m.GetBackingStore().Get("teamsAppId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTemplateParameters gets the templateParameters property value. The templateParameters property
// returns a []KeyValuePairable when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetTemplateParameters()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable) {
    val, err := m.GetBackingStore().Get("templateParameters")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)
    }
    return nil
}
// GetTopic gets the topic property value. The topic property
// returns a TeamworkActivityTopicable when successful
func (m *ItemSendActivityNotificationPostRequestBody) GetTopic()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable) {
    val, err := m.GetBackingStore().Get("topic")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *ItemSendActivityNotificationPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("recipient", m.GetRecipient())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityType sets the activityType property value. The activityType property
func (m *ItemSendActivityNotificationPostRequestBody) SetActivityType(value *string)() {
    err := m.GetBackingStore().Set("activityType", value)
    if err != nil {
        panic(err)
    }
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSendActivityNotificationPostRequestBody) SetAdditionalData(value map[string]any)() {
    err := m.GetBackingStore().Set("additionalData", value)
    if err != nil {
        panic(err)
    }
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *ItemSendActivityNotificationPostRequestBody) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetChainId sets the chainId property value. The chainId property
func (m *ItemSendActivityNotificationPostRequestBody) SetChainId(value *int64)() {
    err := m.GetBackingStore().Set("chainId", value)
    if err != nil {
        panic(err)
    }
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *ItemSendActivityNotificationPostRequestBody) SetPreviewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)() {
    err := m.GetBackingStore().Set("previewText", value)
    if err != nil {
        panic(err)
    }
}
// SetRecipient sets the recipient property value. The recipient property
func (m *ItemSendActivityNotificationPostRequestBody) SetRecipient(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)() {
    err := m.GetBackingStore().Set("recipient", value)
    if err != nil {
        panic(err)
    }
}
// SetTeamsAppId sets the teamsAppId property value. The teamsAppId property
func (m *ItemSendActivityNotificationPostRequestBody) SetTeamsAppId(value *string)() {
    err := m.GetBackingStore().Set("teamsAppId", value)
    if err != nil {
        panic(err)
    }
}
// SetTemplateParameters sets the templateParameters property value. The templateParameters property
func (m *ItemSendActivityNotificationPostRequestBody) SetTemplateParameters(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)() {
    err := m.GetBackingStore().Set("templateParameters", value)
    if err != nil {
        panic(err)
    }
}
// SetTopic sets the topic property value. The topic property
func (m *ItemSendActivityNotificationPostRequestBody) SetTopic(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)() {
    err := m.GetBackingStore().Set("topic", value)
    if err != nil {
        panic(err)
    }
}
type ItemSendActivityNotificationPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityType()(*string)
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetChainId()(*int64)
    GetPreviewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)
    GetRecipient()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)
    GetTeamsAppId()(*string)
    GetTemplateParameters()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)
    GetTopic()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)
    SetActivityType(value *string)()
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetChainId(value *int64)()
    SetPreviewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)()
    SetRecipient(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)()
    SetTeamsAppId(value *string)()
    SetTemplateParameters(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)()
    SetTopic(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)()
}
