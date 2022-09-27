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

type Function struct {
	ClassName string `json:"className,omitempty"`
	Description string `json:"description,omitempty"`
	FunctionLanguage string `json:"functionLanguage,omitempty"`
	FunctionType string `json:"functionType,omitempty"`
	Name string `json:"name,omitempty"`
}
