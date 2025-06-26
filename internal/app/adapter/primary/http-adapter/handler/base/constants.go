package basehandler

const (
	errWrongHeaderContentType               = "Wrong content-Type header"
	errRequestBodyContainsBadlyFormedJSON   = "Request body contains badly-formed JSON at position"
	errRequestBodyContainsInvalidValueField = "Request body contains an invalid value for the %q field (at position %d)"
	errRequestBodyContainsUnknownField      = "Request body contains unknown field"
	errRequestBodyMustNotBeEmpty            = "Request body must not be empty"
)
