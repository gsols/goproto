// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/messages/v1/messages.proto

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

// define the regex for a UUID once up-front
var _messages_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetNextRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetNextRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNextRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetNextRequestMultiError,
// or nil if none found.
func (m *GetNextRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNextRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetStreamId()); err != nil {
		err = GetNextRequestValidationError{
			field:  "StreamId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetNextRequestMultiError(errors)
	}

	return nil
}

func (m *GetNextRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetNextRequestMultiError is an error wrapping multiple validation errors
// returned by GetNextRequest.ValidateAll() if the designated constraints
// aren't met.
type GetNextRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNextRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNextRequestMultiError) AllErrors() []error { return m }

// GetNextRequestValidationError is the validation error returned by
// GetNextRequest.Validate if the designated constraints aren't met.
type GetNextRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNextRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNextRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNextRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNextRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNextRequestValidationError) ErrorName() string { return "GetNextRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetNextRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNextRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNextRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNextRequestValidationError{}

// Validate checks the field values on GetNextResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetNextResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetNextResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetNextResponseMultiError, or nil if none found.
func (m *GetNextResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetNextResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetMessage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetNextResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetNextResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNextResponseValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetNextResponseMultiError(errors)
	}

	return nil
}

// GetNextResponseMultiError is an error wrapping multiple validation errors
// returned by GetNextResponse.ValidateAll() if the designated constraints
// aren't met.
type GetNextResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetNextResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetNextResponseMultiError) AllErrors() []error { return m }

// GetNextResponseValidationError is the validation error returned by
// GetNextResponse.Validate if the designated constraints aren't met.
type GetNextResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNextResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNextResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNextResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNextResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNextResponseValidationError) ErrorName() string { return "GetNextResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetNextResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNextResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNextResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNextResponseValidationError{}

// Validate checks the field values on ConfirmRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ConfirmRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfirmRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConfirmRequestMultiError,
// or nil if none found.
func (m *ConfirmRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfirmRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetStreamId()); err != nil {
		err = ConfirmRequestValidationError{
			field:  "StreamId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateUuid(m.GetMessageId()); err != nil {
		err = ConfirmRequestValidationError{
			field:  "MessageId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ConfirmRequestMultiError(errors)
	}

	return nil
}

func (m *ConfirmRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ConfirmRequestMultiError is an error wrapping multiple validation errors
// returned by ConfirmRequest.ValidateAll() if the designated constraints
// aren't met.
type ConfirmRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfirmRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfirmRequestMultiError) AllErrors() []error { return m }

// ConfirmRequestValidationError is the validation error returned by
// ConfirmRequest.Validate if the designated constraints aren't met.
type ConfirmRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfirmRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfirmRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfirmRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfirmRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfirmRequestValidationError) ErrorName() string { return "ConfirmRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConfirmRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfirmRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfirmRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfirmRequestValidationError{}

// Validate checks the field values on ConfirmResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConfirmResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConfirmResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConfirmResponseMultiError, or nil if none found.
func (m *ConfirmResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ConfirmResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfirmResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfirmResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfirmResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConfirmResponseMultiError(errors)
	}

	return nil
}

// ConfirmResponseMultiError is an error wrapping multiple validation errors
// returned by ConfirmResponse.ValidateAll() if the designated constraints
// aren't met.
type ConfirmResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfirmResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfirmResponseMultiError) AllErrors() []error { return m }

// ConfirmResponseValidationError is the validation error returned by
// ConfirmResponse.Validate if the designated constraints aren't met.
type ConfirmResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfirmResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfirmResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfirmResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfirmResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfirmResponseValidationError) ErrorName() string { return "ConfirmResponseValidationError" }

// Error satisfies the builtin error interface
func (e ConfirmResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfirmResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfirmResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfirmResponseValidationError{}
