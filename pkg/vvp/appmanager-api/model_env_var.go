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

type EnvVar struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	ValueFrom *JsonNode `json:"valueFrom,omitempty"`
}
