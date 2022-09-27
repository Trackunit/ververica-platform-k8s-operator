/*
 * Ververica Platform API
 *
 * The Ververica Platform APIs, excluding Application Manager.
 *
 * API version: 2.7.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package platformapi

type ErrorDetails struct {
	ColumnNumber int32 `json:"columnNumber,omitempty"`
	EndColumnNumber int32 `json:"endColumnNumber,omitempty"`
	EndLineNumber int32 `json:"endLineNumber,omitempty"`
	LineNumber int32 `json:"lineNumber,omitempty"`
	Message string `json:"message,omitempty"`
}
