/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.7.1
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type SystemInformation struct {
	ApiVersion string `json:"apiVersion,omitempty"`
	Kind string `json:"kind,omitempty"`
	Status *SystemInformationStatus `json:"status,omitempty"`
}
