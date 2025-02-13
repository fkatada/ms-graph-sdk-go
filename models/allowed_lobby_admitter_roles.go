package models
type AllowedLobbyAdmitterRoles int

const (
    ORGANIZERANDCOORGANIZERSANDPRESENTERS_ALLOWEDLOBBYADMITTERROLES AllowedLobbyAdmitterRoles = iota
    ORGANIZERANDCOORGANIZERS_ALLOWEDLOBBYADMITTERROLES
    UNKNOWNFUTUREVALUE_ALLOWEDLOBBYADMITTERROLES
)

func (i AllowedLobbyAdmitterRoles) String() string {
    return []string{"organizerAndCoOrganizersAndPresenters", "organizerAndCoOrganizers", "unknownFutureValue"}[i]
}
func ParseAllowedLobbyAdmitterRoles(v string) (any, error) {
    result := ORGANIZERANDCOORGANIZERSANDPRESENTERS_ALLOWEDLOBBYADMITTERROLES
    switch v {
        case "organizerAndCoOrganizersAndPresenters":
            result = ORGANIZERANDCOORGANIZERSANDPRESENTERS_ALLOWEDLOBBYADMITTERROLES
        case "organizerAndCoOrganizers":
            result = ORGANIZERANDCOORGANIZERS_ALLOWEDLOBBYADMITTERROLES
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_ALLOWEDLOBBYADMITTERROLES
        default:
            return nil, nil
    }
    return &result, nil
}
func SerializeAllowedLobbyAdmitterRoles(values []AllowedLobbyAdmitterRoles) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AllowedLobbyAdmitterRoles) isMultiValue() bool {
    return false
}
