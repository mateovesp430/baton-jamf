// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: c1/connector/v2/annotation_ratelimit.proto

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RateLimitDescription with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RateLimitDescription) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RateLimitDescription with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RateLimitDescriptionMultiError, or nil if none found.
func (m *RateLimitDescription) ValidateAll() error {
	return m.validate(true)
}

func (m *RateLimitDescription) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Status

	// no validation rules for Limit

	// no validation rules for Remaining

	if all {
		switch v := interface{}(m.GetResetAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RateLimitDescriptionValidationError{
					field:  "ResetAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RateLimitDescriptionValidationError{
					field:  "ResetAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResetAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RateLimitDescriptionValidationError{
				field:  "ResetAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RateLimitDescriptionMultiError(errors)
	}

	return nil
}

// RateLimitDescriptionMultiError is an error wrapping multiple validation
// errors returned by RateLimitDescription.ValidateAll() if the designated
// constraints aren't met.
type RateLimitDescriptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RateLimitDescriptionMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RateLimitDescriptionMultiError) AllErrors() []error { return m }

// RateLimitDescriptionValidationError is the validation error returned by
// RateLimitDescription.Validate if the designated constraints aren't met.
type RateLimitDescriptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RateLimitDescriptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RateLimitDescriptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RateLimitDescriptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RateLimitDescriptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RateLimitDescriptionValidationError) ErrorName() string {
	return "RateLimitDescriptionValidationError"
}

// Error satisfies the builtin error interface
func (e RateLimitDescriptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRateLimitDescription.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RateLimitDescriptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RateLimitDescriptionValidationError{}
