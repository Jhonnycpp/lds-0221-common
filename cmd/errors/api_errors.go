package errors

import "net/http"

var (
	InvalidAuth = NewApiError(
		http.StatusUnauthorized,
		"invalid_token",
		"Invalid token.",
	)
	UnauthorizedError = NewApiError(
		http.StatusUnauthorized,
		"unauthorized",
		"Unauthorized access",
	)
	UnknownError = NewApiError(
		http.StatusInternalServerError,
		"unknown_error",
		"Unknown application error",
	)
)
