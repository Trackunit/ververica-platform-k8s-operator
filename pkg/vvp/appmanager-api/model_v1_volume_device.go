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

// volumeDevice describes a mapping of a raw block device within a container.
type V1VolumeDevice struct {
	// devicePath is the path inside of the container that the device will be mapped to.
	DevicePath string `json:"devicePath"`
	// name must match the name of a persistentVolumeClaim in the pod
	Name string `json:"name"`
}
