package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type Schedule struct {
    Entity
}
// NewSchedule instantiates a new Schedule and sets the default values.
func NewSchedule()(*Schedule) {
    m := &Schedule{
        Entity: *NewEntity(),
    }
    return m
}
// CreateScheduleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateScheduleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchedule(), nil
}
// GetDayNotes gets the dayNotes property value. The day notes in the schedule.
// returns a []DayNoteable when successful
func (m *Schedule) GetDayNotes()([]DayNoteable) {
    val, err := m.GetBackingStore().Get("dayNotes")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]DayNoteable)
    }
    return nil
}
// GetEnabled gets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
// returns a *bool when successful
func (m *Schedule) GetEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("enabled")
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
func (m *Schedule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["dayNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDayNoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DayNoteable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(DayNoteable)
                }
            }
            m.SetDayNotes(res)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["isActivitiesIncludedWhenCopyingShiftsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActivitiesIncludedWhenCopyingShiftsEnabled(val)
        }
        return nil
    }
    res["offerShiftRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOfferShiftRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfferShiftRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OfferShiftRequestable)
                }
            }
            m.SetOfferShiftRequests(res)
        }
        return nil
    }
    res["offerShiftRequestsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferShiftRequestsEnabled(val)
        }
        return nil
    }
    res["openShiftChangeRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOpenShiftChangeRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShiftChangeRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OpenShiftChangeRequestable)
                }
            }
            m.SetOpenShiftChangeRequests(res)
        }
        return nil
    }
    res["openShifts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOpenShiftFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShiftable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(OpenShiftable)
                }
            }
            m.SetOpenShifts(res)
        }
        return nil
    }
    res["openShiftsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenShiftsEnabled(val)
        }
        return nil
    }
    res["provisionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatus(val.(*OperationStatus))
        }
        return nil
    }
    res["provisionStatusCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatusCode(val)
        }
        return nil
    }
    res["schedulingGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSchedulingGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchedulingGroupable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SchedulingGroupable)
                }
            }
            m.SetSchedulingGroups(res)
        }
        return nil
    }
    res["shifts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateShiftFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Shiftable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Shiftable)
                }
            }
            m.SetShifts(res)
        }
        return nil
    }
    res["startDayOfWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDayOfWeek(val.(*DayOfWeek))
        }
        return nil
    }
    res["swapShiftsChangeRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSwapShiftsChangeRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SwapShiftsChangeRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SwapShiftsChangeRequestable)
                }
            }
            m.SetSwapShiftsChangeRequests(res)
        }
        return nil
    }
    res["swapShiftsRequestsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSwapShiftsRequestsEnabled(val)
        }
        return nil
    }
    res["timeCards"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeCardFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeCardable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TimeCardable)
                }
            }
            m.SetTimeCards(res)
        }
        return nil
    }
    res["timeClockEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeClockEnabled(val)
        }
        return nil
    }
    res["timeClockSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeClockSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeClockSettings(val.(TimeClockSettingsable))
        }
        return nil
    }
    res["timeOffReasons"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeOffReasonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffReasonable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TimeOffReasonable)
                }
            }
            m.SetTimeOffReasons(res)
        }
        return nil
    }
    res["timeOffRequests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeOffRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffRequestable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TimeOffRequestable)
                }
            }
            m.SetTimeOffRequests(res)
        }
        return nil
    }
    res["timeOffRequestsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeOffRequestsEnabled(val)
        }
        return nil
    }
    res["timesOff"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTimeOffFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(TimeOffable)
                }
            }
            m.SetTimesOff(res)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["workforceIntegrationIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetWorkforceIntegrationIds(res)
        }
        return nil
    }
    return res
}
// GetIsActivitiesIncludedWhenCopyingShiftsEnabled gets the isActivitiesIncludedWhenCopyingShiftsEnabled property value. Indicates whether copied shifts include activities from the original shift.
// returns a *bool when successful
func (m *Schedule) GetIsActivitiesIncludedWhenCopyingShiftsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("isActivitiesIncludedWhenCopyingShiftsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOfferShiftRequests gets the offerShiftRequests property value. The offer requests for shifts in the schedule.
// returns a []OfferShiftRequestable when successful
func (m *Schedule) GetOfferShiftRequests()([]OfferShiftRequestable) {
    val, err := m.GetBackingStore().Get("offerShiftRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OfferShiftRequestable)
    }
    return nil
}
// GetOfferShiftRequestsEnabled gets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
// returns a *bool when successful
func (m *Schedule) GetOfferShiftRequestsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("offerShiftRequestsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetOpenShiftChangeRequests gets the openShiftChangeRequests property value. The open shift requests in the schedule.
// returns a []OpenShiftChangeRequestable when successful
func (m *Schedule) GetOpenShiftChangeRequests()([]OpenShiftChangeRequestable) {
    val, err := m.GetBackingStore().Get("openShiftChangeRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OpenShiftChangeRequestable)
    }
    return nil
}
// GetOpenShifts gets the openShifts property value. The set of open shifts in a scheduling group in the schedule.
// returns a []OpenShiftable when successful
func (m *Schedule) GetOpenShifts()([]OpenShiftable) {
    val, err := m.GetBackingStore().Get("openShifts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]OpenShiftable)
    }
    return nil
}
// GetOpenShiftsEnabled gets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
// returns a *bool when successful
func (m *Schedule) GetOpenShiftsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("openShiftsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetProvisionStatus gets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
// returns a *OperationStatus when successful
func (m *Schedule) GetProvisionStatus()(*OperationStatus) {
    val, err := m.GetBackingStore().Get("provisionStatus")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*OperationStatus)
    }
    return nil
}
// GetProvisionStatusCode gets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
// returns a *string when successful
func (m *Schedule) GetProvisionStatusCode()(*string) {
    val, err := m.GetBackingStore().Get("provisionStatusCode")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetSchedulingGroups gets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
// returns a []SchedulingGroupable when successful
func (m *Schedule) GetSchedulingGroups()([]SchedulingGroupable) {
    val, err := m.GetBackingStore().Get("schedulingGroups")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SchedulingGroupable)
    }
    return nil
}
// GetShifts gets the shifts property value. The shifts in the schedule.
// returns a []Shiftable when successful
func (m *Schedule) GetShifts()([]Shiftable) {
    val, err := m.GetBackingStore().Get("shifts")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]Shiftable)
    }
    return nil
}
// GetStartDayOfWeek gets the startDayOfWeek property value. Indicates the start day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
// returns a *DayOfWeek when successful
func (m *Schedule) GetStartDayOfWeek()(*DayOfWeek) {
    val, err := m.GetBackingStore().Get("startDayOfWeek")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*DayOfWeek)
    }
    return nil
}
// GetSwapShiftsChangeRequests gets the swapShiftsChangeRequests property value. The swap requests for shifts in the schedule.
// returns a []SwapShiftsChangeRequestable when successful
func (m *Schedule) GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequestable) {
    val, err := m.GetBackingStore().Get("swapShiftsChangeRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SwapShiftsChangeRequestable)
    }
    return nil
}
// GetSwapShiftsRequestsEnabled gets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
// returns a *bool when successful
func (m *Schedule) GetSwapShiftsRequestsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("swapShiftsRequestsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTimeCards gets the timeCards property value. The time cards in the schedule.
// returns a []TimeCardable when successful
func (m *Schedule) GetTimeCards()([]TimeCardable) {
    val, err := m.GetBackingStore().Get("timeCards")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TimeCardable)
    }
    return nil
}
// GetTimeClockEnabled gets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
// returns a *bool when successful
func (m *Schedule) GetTimeClockEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("timeClockEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTimeClockSettings gets the timeClockSettings property value. The time clock location settings for this schedule.
// returns a TimeClockSettingsable when successful
func (m *Schedule) GetTimeClockSettings()(TimeClockSettingsable) {
    val, err := m.GetBackingStore().Get("timeClockSettings")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(TimeClockSettingsable)
    }
    return nil
}
// GetTimeOffReasons gets the timeOffReasons property value. The set of reasons for a time off in the schedule.
// returns a []TimeOffReasonable when successful
func (m *Schedule) GetTimeOffReasons()([]TimeOffReasonable) {
    val, err := m.GetBackingStore().Get("timeOffReasons")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TimeOffReasonable)
    }
    return nil
}
// GetTimeOffRequests gets the timeOffRequests property value. The time off requests in the schedule.
// returns a []TimeOffRequestable when successful
func (m *Schedule) GetTimeOffRequests()([]TimeOffRequestable) {
    val, err := m.GetBackingStore().Get("timeOffRequests")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TimeOffRequestable)
    }
    return nil
}
// GetTimeOffRequestsEnabled gets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
// returns a *bool when successful
func (m *Schedule) GetTimeOffRequestsEnabled()(*bool) {
    val, err := m.GetBackingStore().Get("timeOffRequestsEnabled")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetTimesOff gets the timesOff property value. The instances of times off in the schedule.
// returns a []TimeOffable when successful
func (m *Schedule) GetTimesOff()([]TimeOffable) {
    val, err := m.GetBackingStore().Get("timesOff")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]TimeOffable)
    }
    return nil
}
// GetTimeZone gets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
// returns a *string when successful
func (m *Schedule) GetTimeZone()(*string) {
    val, err := m.GetBackingStore().Get("timeZone")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetWorkforceIntegrationIds gets the workforceIntegrationIds property value. The IDs for the workforce integrations associated with this schedule.
// returns a []string when successful
func (m *Schedule) GetWorkforceIntegrationIds()([]string) {
    val, err := m.GetBackingStore().Get("workforceIntegrationIds")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *Schedule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDayNotes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDayNotes()))
        for i, v := range m.GetDayNotes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("dayNotes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isActivitiesIncludedWhenCopyingShiftsEnabled", m.GetIsActivitiesIncludedWhenCopyingShiftsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOfferShiftRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOfferShiftRequests()))
        for i, v := range m.GetOfferShiftRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("offerShiftRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("offerShiftRequestsEnabled", m.GetOfferShiftRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOpenShiftChangeRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOpenShiftChangeRequests()))
        for i, v := range m.GetOpenShiftChangeRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("openShiftChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOpenShifts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOpenShifts()))
        for i, v := range m.GetOpenShifts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("openShifts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("openShiftsEnabled", m.GetOpenShiftsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetSchedulingGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchedulingGroups()))
        for i, v := range m.GetSchedulingGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("schedulingGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShifts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetShifts()))
        for i, v := range m.GetShifts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("shifts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartDayOfWeek() != nil {
        cast := (*m.GetStartDayOfWeek()).String()
        err = writer.WriteStringValue("startDayOfWeek", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSwapShiftsChangeRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSwapShiftsChangeRequests()))
        for i, v := range m.GetSwapShiftsChangeRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("swapShiftsChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("swapShiftsRequestsEnabled", m.GetSwapShiftsRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetTimeCards() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimeCards()))
        for i, v := range m.GetTimeCards() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("timeCards", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeClockEnabled", m.GetTimeClockEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("timeClockSettings", m.GetTimeClockSettings())
        if err != nil {
            return err
        }
    }
    if m.GetTimeOffReasons() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimeOffReasons()))
        for i, v := range m.GetTimeOffReasons() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("timeOffReasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTimeOffRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimeOffRequests()))
        for i, v := range m.GetTimeOffRequests() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("timeOffRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeOffRequestsEnabled", m.GetTimeOffRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetTimesOff() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimesOff()))
        for i, v := range m.GetTimesOff() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("timesOff", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrationIds() != nil {
        err = writer.WriteCollectionOfStringValues("workforceIntegrationIds", m.GetWorkforceIntegrationIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDayNotes sets the dayNotes property value. The day notes in the schedule.
func (m *Schedule) SetDayNotes(value []DayNoteable)() {
    err := m.GetBackingStore().Set("dayNotes", value)
    if err != nil {
        panic(err)
    }
}
// SetEnabled sets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
func (m *Schedule) SetEnabled(value *bool)() {
    err := m.GetBackingStore().Set("enabled", value)
    if err != nil {
        panic(err)
    }
}
// SetIsActivitiesIncludedWhenCopyingShiftsEnabled sets the isActivitiesIncludedWhenCopyingShiftsEnabled property value. Indicates whether copied shifts include activities from the original shift.
func (m *Schedule) SetIsActivitiesIncludedWhenCopyingShiftsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("isActivitiesIncludedWhenCopyingShiftsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOfferShiftRequests sets the offerShiftRequests property value. The offer requests for shifts in the schedule.
func (m *Schedule) SetOfferShiftRequests(value []OfferShiftRequestable)() {
    err := m.GetBackingStore().Set("offerShiftRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetOfferShiftRequestsEnabled sets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
func (m *Schedule) SetOfferShiftRequestsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("offerShiftRequestsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetOpenShiftChangeRequests sets the openShiftChangeRequests property value. The open shift requests in the schedule.
func (m *Schedule) SetOpenShiftChangeRequests(value []OpenShiftChangeRequestable)() {
    err := m.GetBackingStore().Set("openShiftChangeRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetOpenShifts sets the openShifts property value. The set of open shifts in a scheduling group in the schedule.
func (m *Schedule) SetOpenShifts(value []OpenShiftable)() {
    err := m.GetBackingStore().Set("openShifts", value)
    if err != nil {
        panic(err)
    }
}
// SetOpenShiftsEnabled sets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
func (m *Schedule) SetOpenShiftsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("openShiftsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisionStatus sets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
func (m *Schedule) SetProvisionStatus(value *OperationStatus)() {
    err := m.GetBackingStore().Set("provisionStatus", value)
    if err != nil {
        panic(err)
    }
}
// SetProvisionStatusCode sets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
func (m *Schedule) SetProvisionStatusCode(value *string)() {
    err := m.GetBackingStore().Set("provisionStatusCode", value)
    if err != nil {
        panic(err)
    }
}
// SetSchedulingGroups sets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
func (m *Schedule) SetSchedulingGroups(value []SchedulingGroupable)() {
    err := m.GetBackingStore().Set("schedulingGroups", value)
    if err != nil {
        panic(err)
    }
}
// SetShifts sets the shifts property value. The shifts in the schedule.
func (m *Schedule) SetShifts(value []Shiftable)() {
    err := m.GetBackingStore().Set("shifts", value)
    if err != nil {
        panic(err)
    }
}
// SetStartDayOfWeek sets the startDayOfWeek property value. Indicates the start day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday.
func (m *Schedule) SetStartDayOfWeek(value *DayOfWeek)() {
    err := m.GetBackingStore().Set("startDayOfWeek", value)
    if err != nil {
        panic(err)
    }
}
// SetSwapShiftsChangeRequests sets the swapShiftsChangeRequests property value. The swap requests for shifts in the schedule.
func (m *Schedule) SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequestable)() {
    err := m.GetBackingStore().Set("swapShiftsChangeRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetSwapShiftsRequestsEnabled sets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
func (m *Schedule) SetSwapShiftsRequestsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("swapShiftsRequestsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeCards sets the timeCards property value. The time cards in the schedule.
func (m *Schedule) SetTimeCards(value []TimeCardable)() {
    err := m.GetBackingStore().Set("timeCards", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeClockEnabled sets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
func (m *Schedule) SetTimeClockEnabled(value *bool)() {
    err := m.GetBackingStore().Set("timeClockEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeClockSettings sets the timeClockSettings property value. The time clock location settings for this schedule.
func (m *Schedule) SetTimeClockSettings(value TimeClockSettingsable)() {
    err := m.GetBackingStore().Set("timeClockSettings", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeOffReasons sets the timeOffReasons property value. The set of reasons for a time off in the schedule.
func (m *Schedule) SetTimeOffReasons(value []TimeOffReasonable)() {
    err := m.GetBackingStore().Set("timeOffReasons", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeOffRequests sets the timeOffRequests property value. The time off requests in the schedule.
func (m *Schedule) SetTimeOffRequests(value []TimeOffRequestable)() {
    err := m.GetBackingStore().Set("timeOffRequests", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeOffRequestsEnabled sets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
func (m *Schedule) SetTimeOffRequestsEnabled(value *bool)() {
    err := m.GetBackingStore().Set("timeOffRequestsEnabled", value)
    if err != nil {
        panic(err)
    }
}
// SetTimesOff sets the timesOff property value. The instances of times off in the schedule.
func (m *Schedule) SetTimesOff(value []TimeOffable)() {
    err := m.GetBackingStore().Set("timesOff", value)
    if err != nil {
        panic(err)
    }
}
// SetTimeZone sets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
func (m *Schedule) SetTimeZone(value *string)() {
    err := m.GetBackingStore().Set("timeZone", value)
    if err != nil {
        panic(err)
    }
}
// SetWorkforceIntegrationIds sets the workforceIntegrationIds property value. The IDs for the workforce integrations associated with this schedule.
func (m *Schedule) SetWorkforceIntegrationIds(value []string)() {
    err := m.GetBackingStore().Set("workforceIntegrationIds", value)
    if err != nil {
        panic(err)
    }
}
type Scheduleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDayNotes()([]DayNoteable)
    GetEnabled()(*bool)
    GetIsActivitiesIncludedWhenCopyingShiftsEnabled()(*bool)
    GetOfferShiftRequests()([]OfferShiftRequestable)
    GetOfferShiftRequestsEnabled()(*bool)
    GetOpenShiftChangeRequests()([]OpenShiftChangeRequestable)
    GetOpenShifts()([]OpenShiftable)
    GetOpenShiftsEnabled()(*bool)
    GetProvisionStatus()(*OperationStatus)
    GetProvisionStatusCode()(*string)
    GetSchedulingGroups()([]SchedulingGroupable)
    GetShifts()([]Shiftable)
    GetStartDayOfWeek()(*DayOfWeek)
    GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequestable)
    GetSwapShiftsRequestsEnabled()(*bool)
    GetTimeCards()([]TimeCardable)
    GetTimeClockEnabled()(*bool)
    GetTimeClockSettings()(TimeClockSettingsable)
    GetTimeOffReasons()([]TimeOffReasonable)
    GetTimeOffRequests()([]TimeOffRequestable)
    GetTimeOffRequestsEnabled()(*bool)
    GetTimesOff()([]TimeOffable)
    GetTimeZone()(*string)
    GetWorkforceIntegrationIds()([]string)
    SetDayNotes(value []DayNoteable)()
    SetEnabled(value *bool)()
    SetIsActivitiesIncludedWhenCopyingShiftsEnabled(value *bool)()
    SetOfferShiftRequests(value []OfferShiftRequestable)()
    SetOfferShiftRequestsEnabled(value *bool)()
    SetOpenShiftChangeRequests(value []OpenShiftChangeRequestable)()
    SetOpenShifts(value []OpenShiftable)()
    SetOpenShiftsEnabled(value *bool)()
    SetProvisionStatus(value *OperationStatus)()
    SetProvisionStatusCode(value *string)()
    SetSchedulingGroups(value []SchedulingGroupable)()
    SetShifts(value []Shiftable)()
    SetStartDayOfWeek(value *DayOfWeek)()
    SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequestable)()
    SetSwapShiftsRequestsEnabled(value *bool)()
    SetTimeCards(value []TimeCardable)()
    SetTimeClockEnabled(value *bool)()
    SetTimeClockSettings(value TimeClockSettingsable)()
    SetTimeOffReasons(value []TimeOffReasonable)()
    SetTimeOffRequests(value []TimeOffRequestable)()
    SetTimeOffRequestsEnabled(value *bool)()
    SetTimesOff(value []TimeOffable)()
    SetTimeZone(value *string)()
    SetWorkforceIntegrationIds(value []string)()
}
