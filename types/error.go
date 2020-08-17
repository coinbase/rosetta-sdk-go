// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package types

// Error Instead of utilizing HTTP status codes to describe node errs (which often do not have a
// good analog), rich errs are returned using this object. Both the code and message fields can be
// individually used to correctly identify an error. Implementations MUST use unique values for both
// fields.
type Error struct {
	// Code is a network-specific error code. If desired, this code can be equivalent to an HTTP
	// status code.
	Code int32 `json:"code"`
	// Message is a network-specific error message. The message MUST NOT change for a given code. In
	// particular, this means that any contextual information should be included in the details
	// field.
	Message string `json:"message"`
	// An error is retriable if the same request may succeed if submitted again.
	Retriable bool `json:"retriable"`
	// Often times it is useful to return context specific to the request that caused the error
	// (i.e. a sample of the stack trace or impacted account) in addition to the standard error
	// message.
	Details map[string]interface{} `json:"details,omitempty"`
}
