package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VirtualEvent struct {
    Entity
}
// NewVirtualEvent instantiates a new VirtualEvent and sets the default values.
func NewVirtualEvent()(*VirtualEvent) {
    m := &VirtualEvent{
        Entity: *NewEntity(),
    }
    return m
}
// CreateVirtualEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVirtualEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.virtualEventWebinar":
                        return NewVirtualEventWebinar(), nil
                }
            }
        }
    }
    return NewVirtualEvent(), nil
}
// GetCreatedBy gets the createdBy property value. Identity information for the creator of the virtual event. Inherited from virtualEvent.
// returns a CommunicationsIdentitySetable when successful
func (m *VirtualEvent) GetCreatedBy()(CommunicationsIdentitySetable) {
    val, err := m.GetBackingStore().Get("createdBy")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(CommunicationsIdentitySetable)
    }
    return nil
}
// GetDescription gets the description property value. Description of the virtual event.
// returns a ItemBodyable when successful
func (m *VirtualEvent) GetDescription()(ItemBodyable) {
    val, err := m.GetBackingStore().Get("description")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(ItemBodyable)
    }
    return nil
}
// GetDisplayName gets the displayName property value. Display name of the virtual event.
// returns a *string when successful
func (m *VirtualEvent) GetDisplayName()(*string) {
    val, err := m.GetBackingStore().Get("displayName")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetEndDateTime gets the endDateTime property value. End time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows.
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEvent) GetEndDateTime()(DateTimeTimeZoneable) {
    val, err := m.GetBackingStore().Get("endDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DateTimeTimeZoneable)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VirtualEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCommunicationsIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(CommunicationsIdentitySetable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(ItemBodyable))
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
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["sessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVirtualEventSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VirtualEventSessionable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(VirtualEventSessionable)
                }
            }
            m.SetSessions(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVirtualEventStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*VirtualEventStatus))
        }
        return nil
    }
    return res
}
// GetSessions gets the sessions property value. Sessions for the virtual event.
// returns a []VirtualEventSessionable when successful
func (m *VirtualEvent) GetSessions()([]VirtualEventSessionable) {
    val, err := m.GetBackingStore().Get("sessions")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]VirtualEventSessionable)
    }
    return nil
}
// GetStartDateTime gets the startDateTime property value. Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows.
// returns a DateTimeTimeZoneable when successful
func (m *VirtualEvent) GetStartDateTime()(DateTimeTimeZoneable) {
    val, err := m.GetBackingStore().Get("startDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(DateTimeTimeZoneable)
    }
    return nil
}
// GetStatus gets the status property value. Status of the virtual event. The possible values are: draft, published, canceled, unknownFutureValue.
// returns a *VirtualEventStatus when successful
func (m *VirtualEvent) GetStatus()(*VirtualEventStatus) {
    val, err := m.GetBackingStore().Get("status")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*VirtualEventStatus)
    }
    return nil
}
// Serialize serializes information the current object
func (m *VirtualEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("description", m.GetDescription())
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
    {
        err = writer.WriteObjectValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSessions()))
        for i, v := range m.GetSessions() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("sessions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. Identity information for the creator of the virtual event. Inherited from virtualEvent.
func (m *VirtualEvent) SetCreatedBy(value CommunicationsIdentitySetable)() {
    err := m.GetBackingStore().Set("createdBy", value)
    if err != nil {
        panic(err)
    }
}
// SetDescription sets the description property value. Description of the virtual event.
func (m *VirtualEvent) SetDescription(value ItemBodyable)() {
    err := m.GetBackingStore().Set("description", value)
    if err != nil {
        panic(err)
    }
}
// SetDisplayName sets the displayName property value. Display name of the virtual event.
func (m *VirtualEvent) SetDisplayName(value *string)() {
    err := m.GetBackingStore().Set("displayName", value)
    if err != nil {
        panic(err)
    }
}
// SetEndDateTime sets the endDateTime property value. End time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows.
func (m *VirtualEvent) SetEndDateTime(value DateTimeTimeZoneable)() {
    err := m.GetBackingStore().Set("endDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetSessions sets the sessions property value. Sessions for the virtual event.
func (m *VirtualEvent) SetSessions(value []VirtualEventSessionable)() {
    err := m.GetBackingStore().Set("sessions", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDateTime sets the startDateTime property value. Start time of the virtual event. The timeZone property can be set to any of the time zones currently supported by Windows.
func (m *VirtualEvent) SetStartDateTime(value DateTimeTimeZoneable)() {
    err := m.GetBackingStore().Set("startDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetStatus sets the status property value. Status of the virtual event. The possible values are: draft, published, canceled, unknownFutureValue.
func (m *VirtualEvent) SetStatus(value *VirtualEventStatus)() {
    err := m.GetBackingStore().Set("status", value)
    if err != nil {
        panic(err)
    }
}
type VirtualEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(CommunicationsIdentitySetable)
    GetDescription()(ItemBodyable)
    GetDisplayName()(*string)
    GetEndDateTime()(DateTimeTimeZoneable)
    GetSessions()([]VirtualEventSessionable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*VirtualEventStatus)
    SetCreatedBy(value CommunicationsIdentitySetable)()
    SetDescription(value ItemBodyable)()
    SetDisplayName(value *string)()
    SetEndDateTime(value DateTimeTimeZoneable)()
    SetSessions(value []VirtualEventSessionable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *VirtualEventStatus)()
}
