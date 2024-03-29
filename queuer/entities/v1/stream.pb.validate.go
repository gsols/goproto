// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/entities/v1/stream.proto

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
var _stream_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Stream with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stream) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stream with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StreamMultiError, or nil if none found.
func (m *Stream) ValidateAll() error {
	return m.validate(true)
}

func (m *Stream) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = StreamValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 100 {
		err := StreamValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 100 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetConnector() != "" {

		if ip := net.ParseIP(m.GetConnector()); ip == nil {
			err := StreamValidationError{
				field:  "Connector",
				reason: "value must be a valid IP address",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetPort() != "" {

	}

	if len(errors) > 0 {
		return StreamMultiError(errors)
	}

	return nil
}

func (m *Stream) _validateUuid(uuid string) error {
	if matched := _stream_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// StreamMultiError is an error wrapping multiple validation errors returned by
// Stream.ValidateAll() if the designated constraints aren't met.
type StreamMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StreamMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StreamMultiError) AllErrors() []error { return m }

// StreamValidationError is the validation error returned by Stream.Validate if
// the designated constraints aren't met.
type StreamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamValidationError) ErrorName() string { return "StreamValidationError" }

// Error satisfies the builtin error interface
func (e StreamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStream.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamValidationError{}
