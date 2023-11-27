// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/clients/v1/messages.proto

package v1

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

	v1 "github.com/gsols/goproto/queuer/entities/v1"
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

	_ = v1.CommandStatus(0)
)

// Validate checks the field values on RegisterClientRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterClientRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterClientRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterClientRequestMultiError, or nil if none found.
func (m *RegisterClientRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterClientRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetClient()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterClientRequestValidationError{
					field:  "Client",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterClientRequestValidationError{
					field:  "Client",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetClient()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterClientRequestValidationError{
				field:  "Client",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterClientRequestMultiError(errors)
	}

	return nil
}

// RegisterClientRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterClientRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterClientRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterClientRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterClientRequestMultiError) AllErrors() []error { return m }

// RegisterClientRequestValidationError is the validation error returned by
// RegisterClientRequest.Validate if the designated constraints aren't met.
type RegisterClientRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterClientRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterClientRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterClientRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterClientRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterClientRequestValidationError) ErrorName() string {
	return "RegisterClientRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterClientRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterClientRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterClientRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterClientRequestValidationError{}

// Validate checks the field values on RegisterClientResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterClientResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterClientResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterClientResponseMultiError, or nil if none found.
func (m *RegisterClientResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterClientResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return RegisterClientResponseMultiError(errors)
	}

	return nil
}

// RegisterClientResponseMultiError is an error wrapping multiple validation
// errors returned by RegisterClientResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterClientResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterClientResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterClientResponseMultiError) AllErrors() []error { return m }

// RegisterClientResponseValidationError is the validation error returned by
// RegisterClientResponse.Validate if the designated constraints aren't met.
type RegisterClientResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterClientResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterClientResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterClientResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterClientResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterClientResponseValidationError) ErrorName() string {
	return "RegisterClientResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterClientResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterClientResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterClientResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterClientResponseValidationError{}

// Validate checks the field values on PublishClientStatsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishClientStatsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishClientStatsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishClientStatsRequestMultiError, or nil if none found.
func (m *PublishClientStatsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishClientStatsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	if all {
		switch v := interface{}(m.GetStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PublishClientStatsRequestValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PublishClientStatsRequestValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PublishClientStatsRequestValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PublishClientStatsRequestMultiError(errors)
	}

	return nil
}

// PublishClientStatsRequestMultiError is an error wrapping multiple validation
// errors returned by PublishClientStatsRequest.ValidateAll() if the
// designated constraints aren't met.
type PublishClientStatsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishClientStatsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishClientStatsRequestMultiError) AllErrors() []error { return m }

// PublishClientStatsRequestValidationError is the validation error returned by
// PublishClientStatsRequest.Validate if the designated constraints aren't met.
type PublishClientStatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishClientStatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishClientStatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishClientStatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishClientStatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishClientStatsRequestValidationError) ErrorName() string {
	return "PublishClientStatsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PublishClientStatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishClientStatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishClientStatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishClientStatsRequestValidationError{}

// Validate checks the field values on PublishClientStatsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishClientStatsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishClientStatsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishClientStatsResponseMultiError, or nil if none found.
func (m *PublishClientStatsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishClientStatsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return PublishClientStatsResponseMultiError(errors)
	}

	return nil
}

// PublishClientStatsResponseMultiError is an error wrapping multiple
// validation errors returned by PublishClientStatsResponse.ValidateAll() if
// the designated constraints aren't met.
type PublishClientStatsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishClientStatsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishClientStatsResponseMultiError) AllErrors() []error { return m }

// PublishClientStatsResponseValidationError is the validation error returned
// by PublishClientStatsResponse.Validate if the designated constraints aren't met.
type PublishClientStatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishClientStatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishClientStatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishClientStatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishClientStatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishClientStatsResponseValidationError) ErrorName() string {
	return "PublishClientStatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PublishClientStatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishClientStatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishClientStatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishClientStatsResponseValidationError{}

// Validate checks the field values on SubscribeToCommandsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubscribeToCommandsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubscribeToCommandsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubscribeToCommandsRequestMultiError, or nil if none found.
func (m *SubscribeToCommandsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SubscribeToCommandsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	if len(errors) > 0 {
		return SubscribeToCommandsRequestMultiError(errors)
	}

	return nil
}

// SubscribeToCommandsRequestMultiError is an error wrapping multiple
// validation errors returned by SubscribeToCommandsRequest.ValidateAll() if
// the designated constraints aren't met.
type SubscribeToCommandsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscribeToCommandsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscribeToCommandsRequestMultiError) AllErrors() []error { return m }

// SubscribeToCommandsRequestValidationError is the validation error returned
// by SubscribeToCommandsRequest.Validate if the designated constraints aren't met.
type SubscribeToCommandsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeToCommandsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeToCommandsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeToCommandsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeToCommandsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeToCommandsRequestValidationError) ErrorName() string {
	return "SubscribeToCommandsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeToCommandsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeToCommandsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeToCommandsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeToCommandsRequestValidationError{}

// Validate checks the field values on SubscribeToCommandsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SubscribeToCommandsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SubscribeToCommandsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SubscribeToCommandsResponseMultiError, or nil if none found.
func (m *SubscribeToCommandsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SubscribeToCommandsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCommand()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SubscribeToCommandsResponseValidationError{
					field:  "Command",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SubscribeToCommandsResponseValidationError{
					field:  "Command",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCommand()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscribeToCommandsResponseValidationError{
				field:  "Command",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SubscribeToCommandsResponseMultiError(errors)
	}

	return nil
}

// SubscribeToCommandsResponseMultiError is an error wrapping multiple
// validation errors returned by SubscribeToCommandsResponse.ValidateAll() if
// the designated constraints aren't met.
type SubscribeToCommandsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SubscribeToCommandsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SubscribeToCommandsResponseMultiError) AllErrors() []error { return m }

// SubscribeToCommandsResponseValidationError is the validation error returned
// by SubscribeToCommandsResponse.Validate if the designated constraints
// aren't met.
type SubscribeToCommandsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscribeToCommandsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscribeToCommandsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscribeToCommandsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscribeToCommandsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscribeToCommandsResponseValidationError) ErrorName() string {
	return "SubscribeToCommandsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SubscribeToCommandsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscribeToCommandsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscribeToCommandsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscribeToCommandsResponseValidationError{}

// Validate checks the field values on AckCommandRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AckCommandRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AckCommandRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AckCommandRequestMultiError, or nil if none found.
func (m *AckCommandRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AckCommandRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClientId

	// no validation rules for CommandId

	// no validation rules for Status

	if len(errors) > 0 {
		return AckCommandRequestMultiError(errors)
	}

	return nil
}

// AckCommandRequestMultiError is an error wrapping multiple validation errors
// returned by AckCommandRequest.ValidateAll() if the designated constraints
// aren't met.
type AckCommandRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AckCommandRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AckCommandRequestMultiError) AllErrors() []error { return m }

// AckCommandRequestValidationError is the validation error returned by
// AckCommandRequest.Validate if the designated constraints aren't met.
type AckCommandRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AckCommandRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AckCommandRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AckCommandRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AckCommandRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AckCommandRequestValidationError) ErrorName() string {
	return "AckCommandRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AckCommandRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAckCommandRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AckCommandRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AckCommandRequestValidationError{}

// Validate checks the field values on AckCommandResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AckCommandResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AckCommandResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AckCommandResponseMultiError, or nil if none found.
func (m *AckCommandResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AckCommandResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return AckCommandResponseMultiError(errors)
	}

	return nil
}

// AckCommandResponseMultiError is an error wrapping multiple validation errors
// returned by AckCommandResponse.ValidateAll() if the designated constraints
// aren't met.
type AckCommandResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AckCommandResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AckCommandResponseMultiError) AllErrors() []error { return m }

// AckCommandResponseValidationError is the validation error returned by
// AckCommandResponse.Validate if the designated constraints aren't met.
type AckCommandResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AckCommandResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AckCommandResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AckCommandResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AckCommandResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AckCommandResponseValidationError) ErrorName() string {
	return "AckCommandResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AckCommandResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAckCommandResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AckCommandResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AckCommandResponseValidationError{}