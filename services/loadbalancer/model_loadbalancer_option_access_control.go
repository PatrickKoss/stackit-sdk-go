/*
Load Balancer API

This API offers an interface to provision and manage load balancing servers in your STACKIT project. It also has the possibility of pooling target servers for load balancing purposes.  For each load balancer provided, two VMs are deployed in your OpenStack project subject to a fee.

API version: 1.7.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package loadbalancer

// LoadbalancerOptionAccessControl Use this option to limit the IP ranges that can use the load balancer.
type LoadbalancerOptionAccessControl struct {
	// Load Balancer is accessible only from an IP address in this range
	AllowedSourceRanges *[]string `json:"allowedSourceRanges,omitempty"`
}
