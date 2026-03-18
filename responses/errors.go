//nolint:gochecknoglobals // part of public API

// Package responses provides structs used with json.Unmarshal in processing responses from the Pandora API.
package responses

import "fmt"

// Pandora API error codes.
const (
	ErrorCodeInternal                   = 0
	ErrorCodeMaintenanceMode            = 1
	ErrorCodeURLParamMissingMethod      = 2
	ErrorCodeURLParamMissingAuthToken   = 3
	ErrorCodeURLParamMissingPartnerID   = 4
	ErrorCodeURLParamMissingUserID      = 5
	ErrorCodeSecureProtocolRequired     = 6
	ErrorCodeCertificateRequired        = 7
	ErrorCodeParameterTypeMismatch      = 8
	ErrorCodeParameterMissing           = 9
	ErrorCodeParameterValueInvalid      = 10
	ErrorCodeAPIVersionNotSupported     = 11
	ErrorCodeLicensingRestrictions      = 12
	ErrorCodeInsufficientConnectivity   = 13
	ErrorCodeUnknownMethodName          = 14
	ErrorCodeWrongProtocol              = 15
	ErrorCodeReadOnlyMode               = 1000
	ErrorCodeInvalidAuthToken           = 1001
	ErrorCodeInvalidPartnerLogin        = 1002
	ErrorCodeListenerNotAuthorized      = 1003
	ErrorCodeUserNotAuthorized          = 1004
	ErrorCodeMaxStationsReached         = 1005
	ErrorCodeStationDoesNotExist        = 1006
	ErrorCodeComplimentaryPeriodInUse   = 1007
	ErrorCodeCallNotAllowed             = 1008
	ErrorCodeDeviceNotFound             = 1009
	ErrorCodePartnerNotAuthorized       = 1010
	ErrorCodeInvalidUsername            = 1011
	ErrorCodeInvalidPassword            = 1012
	ErrorCodeUsernameAlreadyExists      = 1013
	ErrorCodeDeviceAlreadyAssociated    = 1014
	ErrorCodeUpgradeDeviceModelInvalid  = 1015
	ErrorCodeExplicitPinIncorrect       = 1018
	ErrorCodeExplicitPinMalformed       = 1020
	ErrorCodeDeviceModelInvalid         = 1023
	ErrorCodeZipCodeInvalid             = 1024
	ErrorCodeBirthYearInvalid           = 1025
	ErrorCodeBirthYearTooYoung          = 1026
	ErrorCodeInvalidCountryCodeOrGender = 1027
	ErrorCodeDeviceDisabled             = 1034
	ErrorCodeDailyTrialLimitReached     = 1035
	ErrorCodeInvalidSponsor             = 1036
	ErrorCodeUserAlreadyUsedTrial       = 1037
	ErrorCodePlaylistExceeded           = 1039
)

// ErrorCodeMap maps Pandora API error codes to their string names.
//
// Note: Some names contain intentional typos preserved from the Pandora API, including
// MAINTENCANCE_MODE (code 1), URL_PARAM_MISING_PARTNER_ID (code 4), and
// PARAMATER_TYPE_MISMATCH (code 8). These must not be corrected.
//
// Error code 0 (INTERNAL) is a generic error that often indicates authentication issues,
// invalid parameters, or rate limiting. Error codes in the 1000+ range are more specific
// and actionable. When receiving error code 0, check authentication flow and parameter
// validation first.
var ErrorCodeMap = map[int]string{ //nolint:gochecknoglobals // part of public API
	0:    "INTERNAL",
	1:    "MAINTENCANCE_MODE",
	2:    "URL_PARAM_MISSING_METHOD",
	3:    "URL_PARAM_MISSING_AUTH_TOKEN",
	4:    "URL_PARAM_MISING_PARTNER_ID",
	5:    "URL_PARAM_MISSING_USER_ID",
	6:    "SECURE_PROTOCOL_REQUIRED",
	7:    "CERTIFICATE_REQUIRED",
	8:    "PARAMATER_TYPE_MISMATCH",
	9:    "PARAMETER_MISSING",
	10:   "PARAMETER_VALUE_INVALID",
	11:   "API_VERSION_NOT_SUPPORTED",
	12:   "LICENSING_RESTRICTIONS",
	13:   "INSUFFICIENT_CONNECTIVITY",
	14:   "UNKNOWN_METHOD_NAME",
	15:   "WRONG_PROTOCOL",
	1000: "READ_ONLY_MODE",
	1001: "INVALID_AUTH_TOKEN",
	1002: "INVALID_PARTNER_LOGIN",
	1003: "LISTENER_NOT_AUTHORIZED",
	1004: "USER_NOT_AUTHORIZED",
	1005: "MAX_STATIONS_REACHED",
	1006: "STATION_DOES_NOT_EXIST",
	1007: "COMPLIMENTARY_PERIOD_ALREADY_IN_USE",
	1008: "CALL_NOT_ALLOWED",
	1009: "DEVICE_NOT_FOUND",
	1010: "PARTNER_NOT_AUTHORIZED",
	1011: "INVALID_USERNAME",
	1012: "INVALID_PASSWORD",
	1013: "USERNAME_ALREADY_EXISTS",
	1014: "DEVICE_ALREADY_ASSOCIATED_TO_ACCOUNT",
	1015: "UPGRADE_DEVICE_MODEL_INVALID",
	1018: "EXPLICIT_PIN_INCORRECT",
	1020: "EXPLICIT_PIN_MALFORMED",
	1023: "DEVICE_MODEL_INVALID",
	1024: "ZIP_CODE_INVALID",
	1025: "BIRTH_YEAR_INVALID",
	1026: "BIRTH_YEAR_TOO_YOUNG",
	1027: "INVALID_COUNTRY_CODE or INVALID_GENDER",
	1034: "DEVICE_DISABLED",
	1035: "DAILY_TRIAL_LIMIT_REACHED",
	1036: "INVALID_SPONSOR",
	1037: "USER_ALREADY_USED_TRIAL",
	1039: "PLAYLIST_EXCEEDED",
}

// PandoraError represents an API error response from Pandora.
type PandoraError struct {
	Stat    string `json:"stat"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns a string representation of the Pandora API error.
func (e *PandoraError) Error() string {
	return fmt.Sprintf("Pandora Error: %d %s", e.Code, e.Message)
}
