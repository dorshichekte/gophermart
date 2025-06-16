package base_handler

const (
	errWrongHeaderContentType               = "Wrong content-Type header"
	errRequestBodyContainsBadlyFormedJson   = "Request body contains badly-formed JSON at position"
	errRequestBodyContainsInvalidValueField = "Request body contains an invalid value for the %q field (at position %d)"
	errRequestBodyContainsUnknownField      = "Request body contains unknown field"
	errRequestBodyMustNotBeEmpty            = "Request body must not be empty"
)
