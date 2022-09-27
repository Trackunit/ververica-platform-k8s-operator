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

type Artifact struct {
	AdditionalDependencies []string `json:"additionalDependencies,omitempty"`
	EntryClass string `json:"entryClass,omitempty"`
	FlinkImageRegistry string `json:"flinkImageRegistry,omitempty"`
	FlinkImageRepository string `json:"flinkImageRepository,omitempty"`
	FlinkImageTag string `json:"flinkImageTag,omitempty"`
	FlinkVersion string `json:"flinkVersion,omitempty"`
	JarUri string `json:"jarUri,omitempty"`
	Kind string `json:"kind,omitempty"`
	MainArgs string `json:"mainArgs,omitempty"`
	SqlScript string `json:"sqlScript,omitempty"`
}
