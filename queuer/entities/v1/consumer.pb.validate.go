// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/entities/v1/consumer.proto

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
var _consumer_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Consumer with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Consumer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Consumer with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ConsumerMultiError, or nil
// if none found.
func (m *Consumer) ValidateAll() error {
	return m.validate(true)
}

func (m *Consumer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = ConsumerValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetVersion()) < 1 {
		err := ConsumerValidationError{
			field:  "Version",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for IsActive

	if utf8.RuneCountInString(m.GetOs()) < 1 {
		err := ConsumerValidationError{
			field:  "Os",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetArch()) < 1 {
		err := ConsumerValidationError{
			field:  "Arch",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetIp()) < 1 {
		err := ConsumerValidationError{
			field:  "Ip",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetHostname()) < 1 {
		err := ConsumerValidationError{
			field:  "Hostname",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetBootTime()) < 1 {
		err := ConsumerValidationError{
			field:  "BootTime",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.Owner != nil {

		if all {
			switch v := interface{}(m.GetOwner()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ConsumerValidationError{
						field:  "Owner",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ConsumerValidationError{
						field:  "Owner",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetOwner()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConsumerValidationError{
					field:  "Owner",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ConsumerMultiError(errors)
	}

	return nil
}

func (m *Consumer) _validateUuid(uuid string) error {
	if matched := _consumer_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ConsumerMultiError is an error wrapping multiple validation errors returned
// by Consumer.ValidateAll() if the designated constraints aren't met.
type ConsumerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsumerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsumerMultiError) AllErrors() []error { return m }

// ConsumerValidationError is the validation error returned by
// Consumer.Validate if the designated constraints aren't met.
type ConsumerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsumerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsumerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsumerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsumerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsumerValidationError) ErrorName() string { return "ConsumerValidationError" }

// Error satisfies the builtin error interface
func (e ConsumerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsumer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsumerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsumerValidationError{}
