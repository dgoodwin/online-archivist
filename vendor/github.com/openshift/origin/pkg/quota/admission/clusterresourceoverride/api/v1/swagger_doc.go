package v1

// This file contains methods that can be used by the go-restful package to generate Swagger
// documentation for the object types found in 'types.go' This file is automatically generated
// by hack/update-generated-swagger-descriptions.sh and should be run after a full build of OpenShift.
// ==== DO NOT EDIT THIS FILE MANUALLY ====

var map_ClusterResourceOverrideConfig = map[string]string{
	"": "ClusterResourceOverrideConfig is the configuration for the ClusterResourceOverride admission controller which overrides user-provided container request/limit values.",
	"limitCPUToMemoryPercent":     "For each of the following, if a non-zero ratio is specified then the initial value (if any) in the pod spec is overwritten according to the ratio. LimitRange defaults are merged prior to the override.\n\nLimitCPUToMemoryPercent (if > 0) overrides the CPU limit to a ratio of the memory limit; 100% overrides CPU to 1 core per 1GiB of RAM. This is done before overriding the CPU request.",
	"cpuRequestToLimitPercent":    "CPURequestToLimitPercent (if > 0) overrides CPU request to a percentage of CPU limit",
	"memoryRequestToLimitPercent": "MemoryRequestToLimitPercent (if > 0) overrides memory request to a percentage of memory limit",
}

func (ClusterResourceOverrideConfig) SwaggerDoc() map[string]string {
	return map_ClusterResourceOverrideConfig
}
