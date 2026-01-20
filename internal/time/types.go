package time

import (
	"time"
)

// TimezoneInfo contains information about a timezone
type TimezoneInfo struct {
	Name          string             `json:"name"`
	Abbreviation  string             `json:"abbreviation"`
	Offset        string             `json:"offset"`
	OffsetSeconds int                `json:"offset_seconds"`
	IsDST         bool               `json:"is_dst"`
	DST           *DSTInfo           `json:"dst,omitempty"`
	DSTTransition *DSTTransitionInfo `json:"dst_transition,omitempty"` // Keep for backward compatibility
}

// DSTInfo contains DST period information
type DSTInfo struct {
	Start  time.Time     `json:"start"`
	End    time.Time     `json:"end"`
	Saving time.Duration `json:"saving"`
}

// DSTTransitionInfo contains information about DST transitions
type DSTTransitionInfo struct {
	NextTransition time.Time `json:"next_transition"`
	TransitionType string    `json:"transition_type"` // "enter_dst" or "exit_dst"
	OffsetChange   int       `json:"offset_change"`   // seconds
}

// FormatType represents supported time format types
type FormatType string

const (
	FormatRFC3339     FormatType = "RFC3339"
	FormatRFC3339Nano FormatType = "RFC3339Nano"
	FormatUnix        FormatType = "Unix"
	FormatUnixMilli   FormatType = "UnixMilli"
	FormatUnixMicro   FormatType = "UnixMicro"
	FormatUnixNano    FormatType = "UnixNano"
	FormatLayout      FormatType = "Layout"
)

// IsValidFormat checks if a format type is supported
func IsValidFormat(format string) bool {
	switch FormatType(format) {
	case FormatRFC3339, FormatRFC3339Nano, FormatUnix, FormatUnixMilli, FormatUnixMicro, FormatUnixNano, FormatLayout:
		return true
	default:
		return false
	}
}

// GetFormatLayout returns the Go time layout for a given format type
func GetFormatLayout(format FormatType) string {
	switch format {
	case FormatRFC3339:
		return time.RFC3339
	case FormatRFC3339Nano:
		return time.RFC3339Nano
	default:
		return time.RFC3339 // default fallback
	}
}

// ParseTimeInput represents input for parsing time strings
type ParseTimeInput struct {
	TimeString string `json:"time_string" jsonschema:"Time string to parse"`
	Format     string `json:"format,omitempty" jsonschema:"Expected time format (RFC3339, Unix, etc.). If not provided, will attempt to auto-detect"`
	Timezone   string `json:"timezone,omitempty" jsonschema:"IANA timezone name for parsing (e.g., 'America/New_York', 'Europe/London'). Defaults to UTC if not provided"`
}

// FormatTimeInput represents input for formatting time
type FormatTimeInput struct {
	Timestamp interface{} `json:"timestamp" jsonschema:"Timestamp to format (can be Unix timestamp as number, RFC3339 string, or ISO 8601 string)"` // can be string, int, or time.Time
	Format    string      `json:"format" jsonschema:"Desired output format (RFC3339, RFC3339Nano, Unix, UnixMilli, UnixMicro, UnixNano, or Layout)"`
	Timezone  string      `json:"timezone,omitempty" jsonschema:"IANA timezone name for output (e.g., 'America/New_York', 'Europe/London'). Defaults to UTC if not provided"`
}

// GetTimeInput represents input for getting current time
type GetTimeInput struct {
	Timezone string `json:"timezone,omitempty" jsonschema:"IANA timezone name (e.g., 'America/New_York', 'Europe/London'). Defaults to UTC if not provided"`
	Format   string `json:"format,omitempty" jsonschema:"Desired output format (RFC3339, RFC3339Nano, Unix, UnixMilli, UnixMicro, UnixNano, or Layout). Defaults to RFC3339"`
}

// TimezoneInfoInput represents input for timezone information
type TimezoneInfoInput struct {
	Timezone      string    `json:"timezone" jsonschema:"IANA timezone name to get information about (e.g., 'America/New_York', 'Europe/London')"`
	ReferenceTime time.Time `json:"reference_time,omitempty" jsonschema:"Optional reference time for timezone calculations. Defaults to current time if not provided"`
}

// Result types for MCP tool responses

// GetTimeResult represents the result of getting current time
type GetTimeResult struct {
	FormattedTime string `json:"formatted_time" jsonschema:"The current time formatted according to the requested format"`
	Timezone      string `json:"timezone" jsonschema:"The timezone used for formatting"`
	Format        string `json:"format" jsonschema:"The format used for the time string"`
	UnixTimestamp int64  `json:"unix_timestamp" jsonschema:"Unix timestamp in seconds"`
}

// FormatTimeResult represents the result of formatting time
type FormatTimeResult struct {
	FormattedTime string `json:"formatted_time" jsonschema:"The formatted time string"`
	Timezone      string `json:"timezone" jsonschema:"The timezone used for formatting"`
	Format        string `json:"format" jsonschema:"The format used for the time string"`
	UnixTimestamp int64  `json:"unix_timestamp" jsonschema:"Unix timestamp in seconds"`
}

// ParseTimeResult represents the result of parsing time
type ParseTimeResult struct {
	UnixTimestamp int64  `json:"unix_timestamp" jsonschema:"Unix timestamp in seconds"`
	RFC3339       string `json:"rfc3339" jsonschema:"Time in RFC3339 format"`
	Timezone      string `json:"timezone" jsonschema:"The timezone of the parsed time"`
	IsDST         bool   `json:"is_dst" jsonschema:"Whether the time is in daylight saving time"`
}
