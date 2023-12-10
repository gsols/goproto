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

// Validate checks the field values on RegisterConsumerInfoRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterConsumerInfoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterConsumerInfoRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterConsumerInfoRequestMultiError, or nil if none found.
func (m *RegisterConsumerInfoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterConsumerInfoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetConsumerId()); err != nil {
		err = RegisterConsumerInfoRequestValidationError{
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
				errors = append(errors, RegisterConsumerInfoRequestValidationError{
					field:  "ConsumerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterConsumerInfoRequestValidationError{
					field:  "ConsumerInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConsumerInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterConsumerInfoRequestValidationError{
				field:  "ConsumerInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterConsumerInfoRequestMultiError(errors)
	}

	return nil
}

func (m *RegisterConsumerInfoRequest) _validateUuid(uuid string) error {
	if matched := _messages_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// RegisterConsumerInfoRequestMultiError is an error wrapping multiple
// validation errors returned by RegisterConsumerInfoRequest.ValidateAll() if
// the designated constraints aren't met.
type RegisterConsumerInfoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterConsumerInfoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterConsumerInfoRequestMultiError) AllErrors() []error { return m }

// RegisterConsumerInfoRequestValidationError is the validation error returned
// by RegisterConsumerInfoRequest.Validate if the designated constraints
// aren't met.
type RegisterConsumerInfoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterConsumerInfoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterConsumerInfoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterConsumerInfoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterConsumerInfoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterConsumerInfoRequestValidationError) ErrorName() string {
	return "RegisterConsumerInfoRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterConsumerInfoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterConsumerInfoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterConsumerInfoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterConsumerInfoRequestValidationError{}

// Validate checks the field values on RegisterConsumerInfoResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterConsumerInfoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterConsumerInfoResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterConsumerInfoResponseMultiError, or nil if none found.
func (m *RegisterConsumerInfoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterConsumerInfoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RegisterConsumerInfoResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegisterConsumerInfoResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegisterConsumerInfoResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RegisterConsumerInfoResponseMultiError(errors)
	}

	return nil
}

// RegisterConsumerInfoResponseMultiError is an error wrapping multiple
// validation errors returned by RegisterConsumerInfoResponse.ValidateAll() if
// the designated constraints aren't met.
type RegisterConsumerInfoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterConsumerInfoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterConsumerInfoResponseMultiError) AllErrors() []error { return m }

// RegisterConsumerInfoResponseValidationError is the validation error returned
// by RegisterConsumerInfoResponse.Validate if the designated constraints
// aren't met.
type RegisterConsumerInfoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterConsumerInfoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterConsumerInfoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterConsumerInfoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterConsumerInfoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterConsumerInfoResponseValidationError) ErrorName() string {
	return "RegisterConsumerInfoResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterConsumerInfoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterConsumerInfoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterConsumerInfoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterConsumerInfoResponseValidationError{}

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

	if len(errors) > 0 {
		return GetSubscribedStreamsRequestMultiError(errors)
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
