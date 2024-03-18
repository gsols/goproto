// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/consumers/v1/messages.proto

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

	if all {
		switch v := interface{}(m.GetSession()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterConsumerResponseValidationError{
					field:  "Session",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterConsumerResponseValidationError{
					field:  "Session",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSession()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterConsumerResponseValidationError{
				field:  "Session",
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

// Validate checks the field values on PublishConsumerStatsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishConsumerStatsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishConsumerStatsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishConsumerStatsRequestMultiError, or nil if none found.
func (m *PublishConsumerStatsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishConsumerStatsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetConsumerId()); err != nil {
		err = PublishConsumerStatsRequestValidationError{
			field:  "ConsumerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PublishConsumerStatsRequestValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PublishConsumerStatsRequestValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PublishConsumerStatsRequestValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PublishConsumerStatsRequestMultiError(errors)
	}

	return nil
}

func (m *PublishConsumerStatsRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// PublishConsumerStatsRequestMultiError is an error wrapping multiple
// validation errors returned by PublishConsumerStatsRequest.ValidateAll() if
// the designated constraints aren't met.
type PublishConsumerStatsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishConsumerStatsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishConsumerStatsRequestMultiError) AllErrors() []error { return m }

// PublishConsumerStatsRequestValidationError is the validation error returned
// by PublishConsumerStatsRequest.Validate if the designated constraints
// aren't met.
type PublishConsumerStatsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishConsumerStatsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishConsumerStatsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishConsumerStatsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishConsumerStatsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishConsumerStatsRequestValidationError) ErrorName() string {
	return "PublishConsumerStatsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PublishConsumerStatsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishConsumerStatsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishConsumerStatsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishConsumerStatsRequestValidationError{}

// Validate checks the field values on PublishConsumerStatsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PublishConsumerStatsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PublishConsumerStatsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PublishConsumerStatsResponseMultiError, or nil if none found.
func (m *PublishConsumerStatsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PublishConsumerStatsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PublishConsumerStatsResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PublishConsumerStatsResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PublishConsumerStatsResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PublishConsumerStatsResponseMultiError(errors)
	}

	return nil
}

// PublishConsumerStatsResponseMultiError is an error wrapping multiple
// validation errors returned by PublishConsumerStatsResponse.ValidateAll() if
// the designated constraints aren't met.
type PublishConsumerStatsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PublishConsumerStatsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PublishConsumerStatsResponseMultiError) AllErrors() []error { return m }

// PublishConsumerStatsResponseValidationError is the validation error returned
// by PublishConsumerStatsResponse.Validate if the designated constraints
// aren't met.
type PublishConsumerStatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PublishConsumerStatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PublishConsumerStatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PublishConsumerStatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PublishConsumerStatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PublishConsumerStatsResponseValidationError) ErrorName() string {
	return "PublishConsumerStatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e PublishConsumerStatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPublishConsumerStatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PublishConsumerStatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PublishConsumerStatsResponseValidationError{}

// Validate checks the field values on UpdateConsumerInfoRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateConsumerInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateConsumerInfoRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateConsumerInfoRequestMultiError, or nil if none found.
func (m *UpdateConsumerInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateConsumerInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetConsumerId()); err != nil {
		err = UpdateConsumerInfoRequestValidationError{
			field:  "ConsumerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetConsumerInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateConsumerInfoRequestValidationError{
					field:  "ConsumerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateConsumerInfoRequestValidationError{
					field:  "ConsumerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConsumerInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateConsumerInfoRequestValidationError{
				field:  "ConsumerInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateConsumerInfoRequestMultiError(errors)
	}

	return nil
}

func (m *UpdateConsumerInfoRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdateConsumerInfoRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateConsumerInfoRequest.ValidateAll() if the
// designated constraints aren't met.
type UpdateConsumerInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateConsumerInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateConsumerInfoRequestMultiError) AllErrors() []error { return m }

// UpdateConsumerInfoRequestValidationError is the validation error returned by
// UpdateConsumerInfoRequest.Validate if the designated constraints aren't met.
type UpdateConsumerInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateConsumerInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateConsumerInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateConsumerInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateConsumerInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateConsumerInfoRequestValidationError) ErrorName() string {
	return "UpdateConsumerInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateConsumerInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateConsumerInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateConsumerInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateConsumerInfoRequestValidationError{}

// Validate checks the field values on UpdateConsumerInfoResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateConsumerInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateConsumerInfoResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateConsumerInfoResponseMultiError, or nil if none found.
func (m *UpdateConsumerInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateConsumerInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateConsumerInfoResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateConsumerInfoResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateConsumerInfoResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateConsumerInfoResponseMultiError(errors)
	}

	return nil
}

// UpdateConsumerInfoResponseMultiError is an error wrapping multiple
// validation errors returned by UpdateConsumerInfoResponse.ValidateAll() if
// the designated constraints aren't met.
type UpdateConsumerInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateConsumerInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateConsumerInfoResponseMultiError) AllErrors() []error { return m }

// UpdateConsumerInfoResponseValidationError is the validation error returned
// by UpdateConsumerInfoResponse.Validate if the designated constraints aren't met.
type UpdateConsumerInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateConsumerInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateConsumerInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateConsumerInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateConsumerInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateConsumerInfoResponseValidationError) ErrorName() string {
	return "UpdateConsumerInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateConsumerInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateConsumerInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateConsumerInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateConsumerInfoResponseValidationError{}

// Validate checks the field values on GetSubscribedStreamsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSubscribedStreamsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSubscribedStreamsRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSubscribedStreamsRequestMultiError, or nil if none found.
func (m *GetSubscribedStreamsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSubscribedStreamsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetConsumerId()); err != nil {
		err = GetSubscribedStreamsRequestValidationError{
			field:  "ConsumerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetSubscribedStreamsRequestMultiError(errors)
	}

	return nil
}

func (m *GetSubscribedStreamsRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetSubscribedStreamsRequestMultiError is an error wrapping multiple
// validation errors returned by GetSubscribedStreamsRequest.ValidateAll() if
// the designated constraints aren't met.
type GetSubscribedStreamsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSubscribedStreamsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSubscribedStreamsRequestMultiError) AllErrors() []error { return m }

// GetSubscribedStreamsRequestValidationError is the validation error returned
// by GetSubscribedStreamsRequest.Validate if the designated constraints
// aren't met.
type GetSubscribedStreamsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSubscribedStreamsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSubscribedStreamsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSubscribedStreamsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSubscribedStreamsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSubscribedStreamsRequestValidationError) ErrorName() string {
	return "GetSubscribedStreamsRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetSubscribedStreamsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSubscribedStreamsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSubscribedStreamsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSubscribedStreamsRequestValidationError{}

// Validate checks the field values on GetSubscribedStreamsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetSubscribedStreamsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetSubscribedStreamsResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetSubscribedStreamsResponseMultiError, or nil if none found.
func (m *GetSubscribedStreamsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetSubscribedStreamsResponse) validate(all bool) error {
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
					errors = append(errors, GetSubscribedStreamsResponseValidationError{
						field:  fmt.Sprintf("Streams[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetSubscribedStreamsResponseValidationError{
						field:  fmt.Sprintf("Streams[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetSubscribedStreamsResponseValidationError{
					field:  fmt.Sprintf("Streams[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetSubscribedStreamsResponseMultiError(errors)
	}

	return nil
}

// GetSubscribedStreamsResponseMultiError is an error wrapping multiple
// validation errors returned by GetSubscribedStreamsResponse.ValidateAll() if
// the designated constraints aren't met.
type GetSubscribedStreamsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetSubscribedStreamsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetSubscribedStreamsResponseMultiError) AllErrors() []error { return m }

// GetSubscribedStreamsResponseValidationError is the validation error returned
// by GetSubscribedStreamsResponse.Validate if the designated constraints
// aren't met.
type GetSubscribedStreamsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSubscribedStreamsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSubscribedStreamsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSubscribedStreamsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSubscribedStreamsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSubscribedStreamsResponseValidationError) ErrorName() string {
	return "GetSubscribedStreamsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSubscribedStreamsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSubscribedStreamsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSubscribedStreamsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSubscribedStreamsResponseValidationError{}
