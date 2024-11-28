package security
type DeploymentStatus int

const (
    UPTODATE_DEPLOYMENTSTATUS DeploymentStatus = iota
    OUTDATED_DEPLOYMENTSTATUS
    UPDATING_DEPLOYMENTSTATUS
    UPDATEFAILED_DEPLOYMENTSTATUS
    NOTCONFIGURED_DEPLOYMENTSTATUS
    UNREACHABLE_DEPLOYMENTSTATUS
    DISCONNECTED_DEPLOYMENTSTATUS
    STARTFAILURE_DEPLOYMENTSTATUS
    SYNCING_DEPLOYMENTSTATUS
    UNKNOWNFUTUREVALUE_DEPLOYMENTSTATUS
)

func (i DeploymentStatus) String() string {
    return []string{"upToDate", "outdated", "updating", "updateFailed", "notConfigured", "unreachable", "disconnected", "startFailure", "syncing", "unknownFutureValue"}[i]
}
func ParseDeploymentStatus(v string) (any, error) {
    result := UPTODATE_DEPLOYMENTSTATUS
    switch v {
        case "upToDate":
            result = UPTODATE_DEPLOYMENTSTATUS
        case "outdated":
            result = OUTDATED_DEPLOYMENTSTATUS
        case "updating":
            result = UPDATING_DEPLOYMENTSTATUS
        case "updateFailed":
            result = UPDATEFAILED_DEPLOYMENTSTATUS
        case "notConfigured":
            result = NOTCONFIGURED_DEPLOYMENTSTATUS
        case "unreachable":
            result = UNREACHABLE_DEPLOYMENTSTATUS
        case "disconnected":
            result = DISCONNECTED_DEPLOYMENTSTATUS
        case "startFailure":
            result = STARTFAILURE_DEPLOYMENTSTATUS
        case "syncing":
            result = SYNCING_DEPLOYMENTSTATUS
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_DEPLOYMENTSTATUS
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeDeploymentStatus(values []DeploymentStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i DeploymentStatus) isMultiValue() bool {
    return false
}
