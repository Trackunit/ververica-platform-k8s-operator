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

import (
	"time"
)

type JobGraph struct {
	CreateTime time.Time `json:"createTime,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	FlinkVersion string `json:"flinkVersion,omitempty"`
	FullFlinkConfiguration map[string]string `json:"fullFlinkConfiguration,omitempty"`
	JarUris []string `json:"jarUris,omitempty"`
	JobgraphPath string `json:"jobgraphPath,omitempty"`
	Name string `json:"name,omitempty"`
	SinkTables []ReferencedTable `json:"sinkTables,omitempty"`
	SourceTables []ReferencedTable `json:"sourceTables,omitempty"`
	SqlStatement string `json:"sqlStatement,omitempty"`
	State string `json:"state,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
	UserFlinkConfiguration map[string]string `json:"userFlinkConfiguration,omitempty"`
}
