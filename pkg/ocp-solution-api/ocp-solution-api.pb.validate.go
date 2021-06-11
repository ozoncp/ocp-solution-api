// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-solution-api/ocp-solution-api.proto

package ocp_solution_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on CreateSolutionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSolutionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSolutionId() <= 0 {
		return CreateSolutionV1RequestValidationError{
			field:  "SolutionId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetIssueId() <= 0 {
		return CreateSolutionV1RequestValidationError{
			field:  "IssueId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateSolutionV1RequestValidationError is the validation error returned by
// CreateSolutionV1Request.Validate if the designated constraints aren't met.
type CreateSolutionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSolutionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSolutionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSolutionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSolutionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSolutionV1RequestValidationError) ErrorName() string {
	return "CreateSolutionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSolutionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSolutionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSolutionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSolutionV1RequestValidationError{}

// Validate checks the field values on CreateSolutionV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateSolutionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSolution()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateSolutionV1ResponseValidationError{
				field:  "Solution",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CreateSolutionV1ResponseValidationError is the validation error returned by
// CreateSolutionV1Response.Validate if the designated constraints aren't met.
type CreateSolutionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateSolutionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateSolutionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateSolutionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateSolutionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateSolutionV1ResponseValidationError) ErrorName() string {
	return "CreateSolutionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateSolutionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateSolutionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateSolutionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateSolutionV1ResponseValidationError{}

// Validate checks the field values on ListSolutionsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSolutionsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListSolutionsV1RequestValidationError is the validation error returned by
// ListSolutionsV1Request.Validate if the designated constraints aren't met.
type ListSolutionsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSolutionsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSolutionsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSolutionsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSolutionsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSolutionsV1RequestValidationError) ErrorName() string {
	return "ListSolutionsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListSolutionsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSolutionsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSolutionsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSolutionsV1RequestValidationError{}

// Validate checks the field values on ListSolutionsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListSolutionsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetSolutions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListSolutionsV1ResponseValidationError{
					field:  fmt.Sprintf("Solutions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListSolutionsV1ResponseValidationError is the validation error returned by
// ListSolutionsV1Response.Validate if the designated constraints aren't met.
type ListSolutionsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSolutionsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSolutionsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSolutionsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSolutionsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSolutionsV1ResponseValidationError) ErrorName() string {
	return "ListSolutionsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListSolutionsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSolutionsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSolutionsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSolutionsV1ResponseValidationError{}

// Validate checks the field values on DescribeSolutionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSolutionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSolutionId() <= 0 {
		return DescribeSolutionV1RequestValidationError{
			field:  "SolutionId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeSolutionV1RequestValidationError is the validation error returned by
// DescribeSolutionV1Request.Validate if the designated constraints aren't met.
type DescribeSolutionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSolutionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSolutionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSolutionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSolutionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSolutionV1RequestValidationError) ErrorName() string {
	return "DescribeSolutionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSolutionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSolutionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSolutionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSolutionV1RequestValidationError{}

// Validate checks the field values on DescribeSolutionV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeSolutionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetSolution()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeSolutionV1ResponseValidationError{
				field:  "Solution",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeSolutionV1ResponseValidationError is the validation error returned
// by DescribeSolutionV1Response.Validate if the designated constraints aren't met.
type DescribeSolutionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeSolutionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeSolutionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeSolutionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeSolutionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeSolutionV1ResponseValidationError) ErrorName() string {
	return "DescribeSolutionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeSolutionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeSolutionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeSolutionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeSolutionV1ResponseValidationError{}

// Validate checks the field values on RemoveSolutionV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSolutionV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSolutionId() <= 0 {
		return RemoveSolutionV1RequestValidationError{
			field:  "SolutionId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveSolutionV1RequestValidationError is the validation error returned by
// RemoveSolutionV1Request.Validate if the designated constraints aren't met.
type RemoveSolutionV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSolutionV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSolutionV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSolutionV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSolutionV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSolutionV1RequestValidationError) ErrorName() string {
	return "RemoveSolutionV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSolutionV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSolutionV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSolutionV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSolutionV1RequestValidationError{}

// Validate checks the field values on RemoveSolutionV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveSolutionV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveSolutionV1ResponseValidationError is the validation error returned by
// RemoveSolutionV1Response.Validate if the designated constraints aren't met.
type RemoveSolutionV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveSolutionV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveSolutionV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveSolutionV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveSolutionV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveSolutionV1ResponseValidationError) ErrorName() string {
	return "RemoveSolutionV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveSolutionV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveSolutionV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveSolutionV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveSolutionV1ResponseValidationError{}

// Validate checks the field values on Solution with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Solution) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetSolutionId() <= 0 {
		return SolutionValidationError{
			field:  "SolutionId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetIssueId() <= 0 {
		return SolutionValidationError{
			field:  "IssueId",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetVerdict()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SolutionValidationError{
				field:  "Verdict",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SolutionValidationError is the validation error returned by
// Solution.Validate if the designated constraints aren't met.
type SolutionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SolutionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SolutionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SolutionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SolutionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SolutionValidationError) ErrorName() string { return "SolutionValidationError" }

// Error satisfies the builtin error interface
func (e SolutionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSolution.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SolutionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SolutionValidationError{}