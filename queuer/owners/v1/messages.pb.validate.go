// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/owners/v1/messages.proto

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

// Validate checks the field values on CreateOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOwnerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOwnerRequestMultiError, or nil if none found.
func (m *CreateOwnerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOwnerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateOwnerRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateOwnerRequestMultiError(errors)
	}

	return nil
}

// CreateOwnerRequestMultiError is an error wrapping multiple validation errors
// returned by CreateOwnerRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateOwnerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOwnerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOwnerRequestMultiError) AllErrors() []error { return m }

// CreateOwnerRequestValidationError is the validation error returned by
// CreateOwnerRequest.Validate if the designated constraints aren't met.
type CreateOwnerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOwnerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOwnerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOwnerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOwnerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOwnerRequestValidationError) ErrorName() string {
	return "CreateOwnerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOwnerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOwnerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOwnerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOwnerRequestValidationError{}

// Validate checks the field values on CreateOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOwnerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOwnerResponseMultiError, or nil if none found.
func (m *CreateOwnerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOwnerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOwner()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwner()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOwnerResponseValidationError{
				field:  "Owner",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOwnerResponseMultiError(errors)
	}

	return nil
}

// CreateOwnerResponseMultiError is an error wrapping multiple validation
// errors returned by CreateOwnerResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateOwnerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOwnerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOwnerResponseMultiError) AllErrors() []error { return m }

// CreateOwnerResponseValidationError is the validation error returned by
// CreateOwnerResponse.Validate if the designated constraints aren't met.
type CreateOwnerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOwnerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOwnerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOwnerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOwnerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOwnerResponseValidationError) ErrorName() string {
	return "CreateOwnerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOwnerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOwnerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOwnerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOwnerResponseValidationError{}

// Validate checks the field values on GetOwnerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOwnerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOwnerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOwnerRequestMultiError, or nil if none found.
func (m *GetOwnerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOwnerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetOwnerId()); err != nil {
		err = GetOwnerRequestValidationError{
			field:  "OwnerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetOwnerRequestMultiError(errors)
	}

	return nil
}

func (m *GetOwnerRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetOwnerRequestMultiError is an error wrapping multiple validation errors
// returned by GetOwnerRequest.ValidateAll() if the designated constraints
// aren't met.
type GetOwnerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOwnerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOwnerRequestMultiError) AllErrors() []error { return m }

// GetOwnerRequestValidationError is the validation error returned by
// GetOwnerRequest.Validate if the designated constraints aren't met.
type GetOwnerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOwnerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOwnerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOwnerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOwnerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOwnerRequestValidationError) ErrorName() string { return "GetOwnerRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetOwnerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOwnerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOwnerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOwnerRequestValidationError{}

// Validate checks the field values on GetOwnerResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetOwnerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOwnerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOwnerResponseMultiError, or nil if none found.
func (m *GetOwnerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOwnerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOwner()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetOwnerResponseValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOwner()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetOwnerResponseValidationError{
				field:  "Owner",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetOwnerResponseMultiError(errors)
	}

	return nil
}

// GetOwnerResponseMultiError is an error wrapping multiple validation errors
// returned by GetOwnerResponse.ValidateAll() if the designated constraints
// aren't met.
type GetOwnerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOwnerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOwnerResponseMultiError) AllErrors() []error { return m }

// GetOwnerResponseValidationError is the validation error returned by
// GetOwnerResponse.Validate if the designated constraints aren't met.
type GetOwnerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOwnerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOwnerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOwnerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOwnerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOwnerResponseValidationError) ErrorName() string { return "GetOwnerResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetOwnerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOwnerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOwnerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOwnerResponseValidationError{}

// Validate checks the field values on ListStreamsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListStreamsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListStreamsRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListStreamsRequestMultiError, or nil if none found.
func (m *ListStreamsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListStreamsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetOwnerId()); err != nil {
		err = ListStreamsRequestValidationError{
			field:  "OwnerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ListStreamsRequestMultiError(errors)
	}

	return nil
}

func (m *ListStreamsRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ListStreamsRequestMultiError is an error wrapping multiple validation errors
// returned by ListStreamsRequest.ValidateAll() if the designated constraints
// aren't met.
type ListStreamsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListStreamsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListStreamsRequestMultiError) AllErrors() []error { return m }

// ListStreamsRequestValidationError is the validation error returned by
// ListStreamsRequest.Validate if the designated constraints aren't met.
type ListStreamsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStreamsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStreamsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStreamsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStreamsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStreamsRequestValidationError) ErrorName() string {
	return "ListStreamsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListStreamsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStreamsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStreamsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStreamsRequestValidationError{}

// Validate checks the field values on ListStreamsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListStreamsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListStreamsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListStreamsResponseMultiError, or nil if none found.
func (m *ListStreamsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListStreamsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetStreams() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListStreamsResponseValidationError{
						field:  fmt.Sprintf("Streams[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListStreamsResponseValidationError{
						field:  fmt.Sprintf("Streams[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListStreamsResponseValidationError{
					field:  fmt.Sprintf("Streams[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListStreamsResponseMultiError(errors)
	}

	return nil
}

// ListStreamsResponseMultiError is an error wrapping multiple validation
// errors returned by ListStreamsResponse.ValidateAll() if the designated
// constraints aren't met.
type ListStreamsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListStreamsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListStreamsResponseMultiError) AllErrors() []error { return m }

// ListStreamsResponseValidationError is the validation error returned by
// ListStreamsResponse.Validate if the designated constraints aren't met.
type ListStreamsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStreamsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStreamsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStreamsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStreamsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStreamsResponseValidationError) ErrorName() string {
	return "ListStreamsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListStreamsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStreamsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStreamsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStreamsResponseValidationError{}

// Validate checks the field values on RegisterConsumerRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterConsumerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterConsumerRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterConsumerRequestMultiError, or nil if none found.
func (m *RegisterConsumerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterConsumerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetOwnerId()); err != nil {
		err = RegisterConsumerRequestValidationError{
			field:  "OwnerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetConsumer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterConsumerRequestValidationError{
					field:  "Consumer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterConsumerRequestValidationError{
					field:  "Consumer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConsumer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterConsumerRequestValidationError{
				field:  "Consumer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterConsumerRequestMultiError(errors)
	}

	return nil
}

func (m *RegisterConsumerRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// RegisterConsumerRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterConsumerRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterConsumerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterConsumerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterConsumerRequestMultiError) AllErrors() []error { return m }

// RegisterConsumerRequestValidationError is the validation error returned by
// RegisterConsumerRequest.Validate if the designated constraints aren't met.
type RegisterConsumerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterConsumerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterConsumerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterConsumerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterConsumerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterConsumerRequestValidationError) ErrorName() string {
	return "RegisterConsumerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterConsumerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterConsumerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterConsumerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterConsumerRequestValidationError{}

// Validate checks the field values on RegisterConsumerResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterConsumerResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterConsumerResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterConsumerResponseMultiError, or nil if none found.
func (m *RegisterConsumerResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterConsumerResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterConsumerResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterConsumerResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterConsumerResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterConsumerResponseMultiError(errors)
	}

	return nil
}

// RegisterConsumerResponseMultiError is an error wrapping multiple validation
// errors returned by RegisterConsumerResponse.ValidateAll() if the designated
// constraints aren't met.
type RegisterConsumerResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterConsumerResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterConsumerResponseMultiError) AllErrors() []error { return m }

// RegisterConsumerResponseValidationError is the validation error returned by
// RegisterConsumerResponse.Validate if the designated constraints aren't met.
type RegisterConsumerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterConsumerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterConsumerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterConsumerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterConsumerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterConsumerResponseValidationError) ErrorName() string {
	return "RegisterConsumerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterConsumerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterConsumerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterConsumerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterConsumerResponseValidationError{}

// Validate checks the field values on GetConsumersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConsumersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConsumersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConsumersRequestMultiError, or nil if none found.
func (m *GetConsumersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConsumersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetOwnerId()); err != nil {
		err = GetConsumersRequestValidationError{
			field:  "OwnerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetConsumersRequestMultiError(errors)
	}

	return nil
}

func (m *GetConsumersRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetConsumersRequestMultiError is an error wrapping multiple validation
// errors returned by GetConsumersRequest.ValidateAll() if the designated
// constraints aren't met.
type GetConsumersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConsumersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConsumersRequestMultiError) AllErrors() []error { return m }

// GetConsumersRequestValidationError is the validation error returned by
// GetConsumersRequest.Validate if the designated constraints aren't met.
type GetConsumersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConsumersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConsumersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConsumersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConsumersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConsumersRequestValidationError) ErrorName() string {
	return "GetConsumersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetConsumersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConsumersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConsumersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConsumersRequestValidationError{}

// Validate checks the field values on GetConsumersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetConsumersResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetConsumersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetConsumersResponseMultiError, or nil if none found.
func (m *GetConsumersResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetConsumersResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetConsumers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetConsumersResponseValidationError{
						field:  fmt.Sprintf("Consumers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetConsumersResponseValidationError{
						field:  fmt.Sprintf("Consumers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetConsumersResponseValidationError{
					field:  fmt.Sprintf("Consumers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetConsumersResponseMultiError(errors)
	}

	return nil
}

// GetConsumersResponseMultiError is an error wrapping multiple validation
// errors returned by GetConsumersResponse.ValidateAll() if the designated
// constraints aren't met.
type GetConsumersResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetConsumersResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetConsumersResponseMultiError) AllErrors() []error { return m }

// GetConsumersResponseValidationError is the validation error returned by
// GetConsumersResponse.Validate if the designated constraints aren't met.
type GetConsumersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetConsumersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetConsumersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetConsumersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetConsumersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetConsumersResponseValidationError) ErrorName() string {
	return "GetConsumersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetConsumersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetConsumersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetConsumersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetConsumersResponseValidationError{}
