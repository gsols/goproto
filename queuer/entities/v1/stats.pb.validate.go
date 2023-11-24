// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: queuer/entities/v1/stats.proto

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

// Validate checks the field values on CPUStats with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CPUStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CPUStats with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CPUStatsMultiError, or nil
// if none found.
func (m *CPUStats) ValidateAll() error {
	return m.validate(true)
}

func (m *CPUStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	// no validation rules for System

	// no validation rules for Idle

	// no validation rules for Total

	if len(errors) > 0 {
		return CPUStatsMultiError(errors)
	}

	return nil
}

// CPUStatsMultiError is an error wrapping multiple validation errors returned
// by CPUStats.ValidateAll() if the designated constraints aren't met.
type CPUStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CPUStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CPUStatsMultiError) AllErrors() []error { return m }

// CPUStatsValidationError is the validation error returned by
// CPUStats.Validate if the designated constraints aren't met.
type CPUStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CPUStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CPUStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CPUStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CPUStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CPUStatsValidationError) ErrorName() string { return "CPUStatsValidationError" }

// Error satisfies the builtin error interface
func (e CPUStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCPUStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CPUStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CPUStatsValidationError{}

// Validate checks the field values on MemoryStats with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MemoryStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MemoryStats with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MemoryStatsMultiError, or
// nil if none found.
func (m *MemoryStats) ValidateAll() error {
	return m.validate(true)
}

func (m *MemoryStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	// no validation rules for Used

	// no validation rules for Available

	// no validation rules for Free

	if len(errors) > 0 {
		return MemoryStatsMultiError(errors)
	}

	return nil
}

// MemoryStatsMultiError is an error wrapping multiple validation errors
// returned by MemoryStats.ValidateAll() if the designated constraints aren't met.
type MemoryStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MemoryStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MemoryStatsMultiError) AllErrors() []error { return m }

// MemoryStatsValidationError is the validation error returned by
// MemoryStats.Validate if the designated constraints aren't met.
type MemoryStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MemoryStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MemoryStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MemoryStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MemoryStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MemoryStatsValidationError) ErrorName() string { return "MemoryStatsValidationError" }

// Error satisfies the builtin error interface
func (e MemoryStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMemoryStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MemoryStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MemoryStatsValidationError{}

// Validate checks the field values on LoadStats with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoadStats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoadStats with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoadStatsMultiError, or nil
// if none found.
func (m *LoadStats) ValidateAll() error {
	return m.validate(true)
}

func (m *LoadStats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for One

	// no validation rules for Five

	// no validation rules for Fifteen

	if len(errors) > 0 {
		return LoadStatsMultiError(errors)
	}

	return nil
}

// LoadStatsMultiError is an error wrapping multiple validation errors returned
// by LoadStats.ValidateAll() if the designated constraints aren't met.
type LoadStatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoadStatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoadStatsMultiError) AllErrors() []error { return m }

// LoadStatsValidationError is the validation error returned by
// LoadStats.Validate if the designated constraints aren't met.
type LoadStatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoadStatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoadStatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoadStatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoadStatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoadStatsValidationError) ErrorName() string { return "LoadStatsValidationError" }

// Error satisfies the builtin error interface
func (e LoadStatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoadStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoadStatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoadStatsValidationError{}

// Validate checks the field values on NetworkSpeed with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NetworkSpeed) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NetworkSpeed with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NetworkSpeedMultiError, or
// nil if none found.
func (m *NetworkSpeed) ValidateAll() error {
	return m.validate(true)
}

func (m *NetworkSpeed) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Download

	// no validation rules for Upload

	if len(errors) > 0 {
		return NetworkSpeedMultiError(errors)
	}

	return nil
}

// NetworkSpeedMultiError is an error wrapping multiple validation errors
// returned by NetworkSpeed.ValidateAll() if the designated constraints aren't met.
type NetworkSpeedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NetworkSpeedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NetworkSpeedMultiError) AllErrors() []error { return m }

// NetworkSpeedValidationError is the validation error returned by
// NetworkSpeed.Validate if the designated constraints aren't met.
type NetworkSpeedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NetworkSpeedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NetworkSpeedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NetworkSpeedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NetworkSpeedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NetworkSpeedValidationError) ErrorName() string { return "NetworkSpeedValidationError" }

// Error satisfies the builtin error interface
func (e NetworkSpeedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNetworkSpeed.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NetworkSpeedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NetworkSpeedValidationError{}

// Validate checks the field values on Stats with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Stats) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Stats with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in StatsMultiError, or nil if none found.
func (m *Stats) ValidateAll() error {
	return m.validate(true)
}

func (m *Stats) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCpu()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatsValidationError{
					field:  "Cpu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatsValidationError{
					field:  "Cpu",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCpu()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatsValidationError{
				field:  "Cpu",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLoad()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatsValidationError{
					field:  "Load",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatsValidationError{
					field:  "Load",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLoad()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatsValidationError{
				field:  "Load",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetMemory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatsValidationError{
					field:  "Memory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatsValidationError{
					field:  "Memory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMemory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatsValidationError{
				field:  "Memory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StatsMultiError(errors)
	}

	return nil
}

// StatsMultiError is an error wrapping multiple validation errors returned by
// Stats.ValidateAll() if the designated constraints aren't met.
type StatsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatsMultiError) AllErrors() []error { return m }

// StatsValidationError is the validation error returned by Stats.Validate if
// the designated constraints aren't met.
type StatsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatsValidationError) ErrorName() string { return "StatsValidationError" }

// Error satisfies the builtin error interface
func (e StatsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStats.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatsValidationError{}
