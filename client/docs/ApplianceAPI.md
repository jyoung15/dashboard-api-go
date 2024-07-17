# \ApplianceAPI

All URIs are relative to *https://api.meraki.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceApplianceVmxAuthenticationToken**](ApplianceAPI.md#CreateDeviceApplianceVmxAuthenticationToken) | **Post** /devices/{serial}/appliance/vmx/authenticationToken | Generate a new vMX authentication token
[**CreateNetworkAppliancePrefixesDelegatedStatic**](ApplianceAPI.md#CreateNetworkAppliancePrefixesDelegatedStatic) | **Post** /networks/{networkId}/appliance/prefixes/delegated/statics | Add a static delegated prefix from a network
[**CreateNetworkApplianceRfProfile**](ApplianceAPI.md#CreateNetworkApplianceRfProfile) | **Post** /networks/{networkId}/appliance/rfProfiles | Creates new RF profile for this network
[**CreateNetworkApplianceStaticRoute**](ApplianceAPI.md#CreateNetworkApplianceStaticRoute) | **Post** /networks/{networkId}/appliance/staticRoutes | Add a static route for an MX or teleworker network
[**CreateNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceAPI.md#CreateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Post** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | Add a custom performance class for an MX network
[**CreateNetworkApplianceVlan**](ApplianceAPI.md#CreateNetworkApplianceVlan) | **Post** /networks/{networkId}/appliance/vlans | Add a VLAN
[**DeleteNetworkAppliancePrefixesDelegatedStatic**](ApplianceAPI.md#DeleteNetworkAppliancePrefixesDelegatedStatic) | **Delete** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Delete a static delegated prefix from a network
[**DeleteNetworkApplianceRfProfile**](ApplianceAPI.md#DeleteNetworkApplianceRfProfile) | **Delete** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Delete a RF Profile
[**DeleteNetworkApplianceStaticRoute**](ApplianceAPI.md#DeleteNetworkApplianceStaticRoute) | **Delete** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Delete a static route from an MX or teleworker network
[**DeleteNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceAPI.md#DeleteNetworkApplianceTrafficShapingCustomPerformanceClass) | **Delete** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Delete a custom performance class from an MX network
[**DeleteNetworkApplianceVlan**](ApplianceAPI.md#DeleteNetworkApplianceVlan) | **Delete** /networks/{networkId}/appliance/vlans/{vlanId} | Delete a VLAN from a network
[**GetDeviceApplianceDhcpSubnets**](ApplianceAPI.md#GetDeviceApplianceDhcpSubnets) | **Get** /devices/{serial}/appliance/dhcp/subnets | Return the DHCP subnet information for an appliance
[**GetDeviceAppliancePerformance**](ApplianceAPI.md#GetDeviceAppliancePerformance) | **Get** /devices/{serial}/appliance/performance | Return the performance score for a single MX
[**GetDeviceAppliancePrefixesDelegated**](ApplianceAPI.md#GetDeviceAppliancePrefixesDelegated) | **Get** /devices/{serial}/appliance/prefixes/delegated | Return current delegated IPv6 prefixes on an appliance.
[**GetDeviceAppliancePrefixesDelegatedVlanAssignments**](ApplianceAPI.md#GetDeviceAppliancePrefixesDelegatedVlanAssignments) | **Get** /devices/{serial}/appliance/prefixes/delegated/vlanAssignments | Return prefixes assigned to all IPv6 enabled VLANs on an appliance.
[**GetDeviceApplianceRadioSettings**](ApplianceAPI.md#GetDeviceApplianceRadioSettings) | **Get** /devices/{serial}/appliance/radio/settings | Return the radio settings of an appliance
[**GetDeviceApplianceUplinksSettings**](ApplianceAPI.md#GetDeviceApplianceUplinksSettings) | **Get** /devices/{serial}/appliance/uplinks/settings | Return the uplink settings for an MX appliance
[**GetNetworkApplianceClientSecurityEvents**](ApplianceAPI.md#GetNetworkApplianceClientSecurityEvents) | **Get** /networks/{networkId}/appliance/clients/{clientId}/security/events | List the security events for a client
[**GetNetworkApplianceConnectivityMonitoringDestinations**](ApplianceAPI.md#GetNetworkApplianceConnectivityMonitoringDestinations) | **Get** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Return the connectivity testing destinations for an MX network
[**GetNetworkApplianceContentFiltering**](ApplianceAPI.md#GetNetworkApplianceContentFiltering) | **Get** /networks/{networkId}/appliance/contentFiltering | Return the content filtering settings for an MX network
[**GetNetworkApplianceContentFilteringCategories**](ApplianceAPI.md#GetNetworkApplianceContentFilteringCategories) | **Get** /networks/{networkId}/appliance/contentFiltering/categories | List all available content filtering categories for an MX network
[**GetNetworkApplianceFirewallCellularFirewallRules**](ApplianceAPI.md#GetNetworkApplianceFirewallCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Return the cellular firewall rules for an MX network
[**GetNetworkApplianceFirewallFirewalledService**](ApplianceAPI.md#GetNetworkApplianceFirewallFirewalledService) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Return the accessibility settings of the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**GetNetworkApplianceFirewallFirewalledServices**](ApplianceAPI.md#GetNetworkApplianceFirewallFirewalledServices) | **Get** /networks/{networkId}/appliance/firewall/firewalledServices | List the appliance services and their accessibility rules
[**GetNetworkApplianceFirewallInboundCellularFirewallRules**](ApplianceAPI.md#GetNetworkApplianceFirewallInboundCellularFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Return the inbound cellular firewall rules for an MX network
[**GetNetworkApplianceFirewallInboundFirewallRules**](ApplianceAPI.md#GetNetworkApplianceFirewallInboundFirewallRules) | **Get** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Return the inbound firewall rules for an MX network
[**GetNetworkApplianceFirewallL3FirewallRules**](ApplianceAPI.md#GetNetworkApplianceFirewallL3FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l3FirewallRules | Return the L3 firewall rules for an MX network
[**GetNetworkApplianceFirewallL7FirewallRules**](ApplianceAPI.md#GetNetworkApplianceFirewallL7FirewallRules) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules | List the MX L7 firewall rules for an MX network
[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories**](ApplianceAPI.md#GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories) | **Get** /networks/{networkId}/appliance/firewall/l7FirewallRules/applicationCategories | Return the L7 firewall application categories and their associated applications for an MX network
[**GetNetworkApplianceFirewallOneToManyNatRules**](ApplianceAPI.md#GetNetworkApplianceFirewallOneToManyNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Return the 1:Many NAT mapping rules for an MX network
[**GetNetworkApplianceFirewallOneToOneNatRules**](ApplianceAPI.md#GetNetworkApplianceFirewallOneToOneNatRules) | **Get** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Return the 1:1 NAT mapping rules for an MX network
[**GetNetworkApplianceFirewallPortForwardingRules**](ApplianceAPI.md#GetNetworkApplianceFirewallPortForwardingRules) | **Get** /networks/{networkId}/appliance/firewall/portForwardingRules | Return the port forwarding rules for an MX network
[**GetNetworkApplianceFirewallSettings**](ApplianceAPI.md#GetNetworkApplianceFirewallSettings) | **Get** /networks/{networkId}/appliance/firewall/settings | Return the firewall settings for this network
[**GetNetworkAppliancePort**](ApplianceAPI.md#GetNetworkAppliancePort) | **Get** /networks/{networkId}/appliance/ports/{portId} | Return per-port VLAN settings for a single MX port.
[**GetNetworkAppliancePorts**](ApplianceAPI.md#GetNetworkAppliancePorts) | **Get** /networks/{networkId}/appliance/ports | List per-port VLAN settings for all ports of a MX.
[**GetNetworkAppliancePrefixesDelegatedStatic**](ApplianceAPI.md#GetNetworkAppliancePrefixesDelegatedStatic) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Return a static delegated prefix from a network
[**GetNetworkAppliancePrefixesDelegatedStatics**](ApplianceAPI.md#GetNetworkAppliancePrefixesDelegatedStatics) | **Get** /networks/{networkId}/appliance/prefixes/delegated/statics | List static delegated prefixes for a network
[**GetNetworkApplianceRfProfile**](ApplianceAPI.md#GetNetworkApplianceRfProfile) | **Get** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Return a RF profile
[**GetNetworkApplianceRfProfiles**](ApplianceAPI.md#GetNetworkApplianceRfProfiles) | **Get** /networks/{networkId}/appliance/rfProfiles | List the RF profiles for this network
[**GetNetworkApplianceSecurityEvents**](ApplianceAPI.md#GetNetworkApplianceSecurityEvents) | **Get** /networks/{networkId}/appliance/security/events | List the security events for a network
[**GetNetworkApplianceSecurityIntrusion**](ApplianceAPI.md#GetNetworkApplianceSecurityIntrusion) | **Get** /networks/{networkId}/appliance/security/intrusion | Returns all supported intrusion settings for an MX network
[**GetNetworkApplianceSecurityMalware**](ApplianceAPI.md#GetNetworkApplianceSecurityMalware) | **Get** /networks/{networkId}/appliance/security/malware | Returns all supported malware settings for an MX network
[**GetNetworkApplianceSettings**](ApplianceAPI.md#GetNetworkApplianceSettings) | **Get** /networks/{networkId}/appliance/settings | Return the appliance settings for a network
[**GetNetworkApplianceSingleLan**](ApplianceAPI.md#GetNetworkApplianceSingleLan) | **Get** /networks/{networkId}/appliance/singleLan | Return single LAN configuration
[**GetNetworkApplianceSsid**](ApplianceAPI.md#GetNetworkApplianceSsid) | **Get** /networks/{networkId}/appliance/ssids/{number} | Return a single MX SSID
[**GetNetworkApplianceSsids**](ApplianceAPI.md#GetNetworkApplianceSsids) | **Get** /networks/{networkId}/appliance/ssids | List the MX SSIDs in a network
[**GetNetworkApplianceStaticRoute**](ApplianceAPI.md#GetNetworkApplianceStaticRoute) | **Get** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Return a static route for an MX or teleworker network
[**GetNetworkApplianceStaticRoutes**](ApplianceAPI.md#GetNetworkApplianceStaticRoutes) | **Get** /networks/{networkId}/appliance/staticRoutes | List the static routes for an MX or teleworker network
[**GetNetworkApplianceTrafficShaping**](ApplianceAPI.md#GetNetworkApplianceTrafficShaping) | **Get** /networks/{networkId}/appliance/trafficShaping | Display the traffic shaping settings for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceAPI.md#GetNetworkApplianceTrafficShapingCustomPerformanceClass) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Return a custom performance class for an MX network
[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses**](ApplianceAPI.md#GetNetworkApplianceTrafficShapingCustomPerformanceClasses) | **Get** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses | List all custom performance classes for an MX network
[**GetNetworkApplianceTrafficShapingRules**](ApplianceAPI.md#GetNetworkApplianceTrafficShapingRules) | **Get** /networks/{networkId}/appliance/trafficShaping/rules | Display the traffic shaping settings rules for an MX network
[**GetNetworkApplianceTrafficShapingUplinkBandwidth**](ApplianceAPI.md#GetNetworkApplianceTrafficShapingUplinkBandwidth) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth | Returns the uplink bandwidth limits for your MX network
[**GetNetworkApplianceTrafficShapingUplinkSelection**](ApplianceAPI.md#GetNetworkApplianceTrafficShapingUplinkSelection) | **Get** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Show uplink selection settings for an MX network
[**GetNetworkApplianceUplinksUsageHistory**](ApplianceAPI.md#GetNetworkApplianceUplinksUsageHistory) | **Get** /networks/{networkId}/appliance/uplinks/usageHistory | Get the sent and received bytes for each uplink of a network.
[**GetNetworkApplianceVlan**](ApplianceAPI.md#GetNetworkApplianceVlan) | **Get** /networks/{networkId}/appliance/vlans/{vlanId} | Return a VLAN
[**GetNetworkApplianceVlans**](ApplianceAPI.md#GetNetworkApplianceVlans) | **Get** /networks/{networkId}/appliance/vlans | List the VLANs for an MX network
[**GetNetworkApplianceVlansSettings**](ApplianceAPI.md#GetNetworkApplianceVlansSettings) | **Get** /networks/{networkId}/appliance/vlans/settings | Returns the enabled status of VLANs for the network
[**GetNetworkApplianceVpnBgp**](ApplianceAPI.md#GetNetworkApplianceVpnBgp) | **Get** /networks/{networkId}/appliance/vpn/bgp | Return a Hub BGP Configuration
[**GetNetworkApplianceVpnSiteToSiteVpn**](ApplianceAPI.md#GetNetworkApplianceVpnSiteToSiteVpn) | **Get** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Return the site-to-site VPN settings of a network
[**GetNetworkApplianceWarmSpare**](ApplianceAPI.md#GetNetworkApplianceWarmSpare) | **Get** /networks/{networkId}/appliance/warmSpare | Return MX warm spare settings
[**GetOrganizationApplianceSecurityEvents**](ApplianceAPI.md#GetOrganizationApplianceSecurityEvents) | **Get** /organizations/{organizationId}/appliance/security/events | List the security events for an organization
[**GetOrganizationApplianceSecurityIntrusion**](ApplianceAPI.md#GetOrganizationApplianceSecurityIntrusion) | **Get** /organizations/{organizationId}/appliance/security/intrusion | Returns all supported intrusion settings for an organization
[**GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork**](ApplianceAPI.md#GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork) | **Get** /organizations/{organizationId}/appliance/trafficShaping/vpnExclusions/byNetwork | Display VPN exclusion rules for MX networks.
[**GetOrganizationApplianceUplinkStatuses**](ApplianceAPI.md#GetOrganizationApplianceUplinkStatuses) | **Get** /organizations/{organizationId}/appliance/uplink/statuses | List the uplink status of every Meraki MX and Z series appliances in the organization
[**GetOrganizationApplianceUplinksStatusesOverview**](ApplianceAPI.md#GetOrganizationApplianceUplinksStatusesOverview) | **Get** /organizations/{organizationId}/appliance/uplinks/statuses/overview | Returns an overview of uplink statuses
[**GetOrganizationApplianceUplinksUsageByNetwork**](ApplianceAPI.md#GetOrganizationApplianceUplinksUsageByNetwork) | **Get** /organizations/{organizationId}/appliance/uplinks/usage/byNetwork | Get the sent and received bytes for each uplink of all MX and Z networks within an organization
[**GetOrganizationApplianceVpnStats**](ApplianceAPI.md#GetOrganizationApplianceVpnStats) | **Get** /organizations/{organizationId}/appliance/vpn/stats | Show VPN history stat for networks in an organization
[**GetOrganizationApplianceVpnStatuses**](ApplianceAPI.md#GetOrganizationApplianceVpnStatuses) | **Get** /organizations/{organizationId}/appliance/vpn/statuses | Show VPN status for networks in an organization
[**GetOrganizationApplianceVpnThirdPartyVPNPeers**](ApplianceAPI.md#GetOrganizationApplianceVpnThirdPartyVPNPeers) | **Get** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Return the third party VPN peers for an organization
[**GetOrganizationApplianceVpnVpnFirewallRules**](ApplianceAPI.md#GetOrganizationApplianceVpnVpnFirewallRules) | **Get** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Return the firewall rules for an organization&#39;s site-to-site VPN
[**SwapNetworkApplianceWarmSpare**](ApplianceAPI.md#SwapNetworkApplianceWarmSpare) | **Post** /networks/{networkId}/appliance/warmSpare/swap | Swap MX primary and warm spare appliances
[**UpdateDeviceApplianceRadioSettings**](ApplianceAPI.md#UpdateDeviceApplianceRadioSettings) | **Put** /devices/{serial}/appliance/radio/settings | Update the radio settings of an appliance
[**UpdateDeviceApplianceUplinksSettings**](ApplianceAPI.md#UpdateDeviceApplianceUplinksSettings) | **Put** /devices/{serial}/appliance/uplinks/settings | Update the uplink settings for an MX appliance
[**UpdateNetworkApplianceConnectivityMonitoringDestinations**](ApplianceAPI.md#UpdateNetworkApplianceConnectivityMonitoringDestinations) | **Put** /networks/{networkId}/appliance/connectivityMonitoringDestinations | Update the connectivity testing destinations for an MX network
[**UpdateNetworkApplianceContentFiltering**](ApplianceAPI.md#UpdateNetworkApplianceContentFiltering) | **Put** /networks/{networkId}/appliance/contentFiltering | Update the content filtering settings for an MX network
[**UpdateNetworkApplianceFirewallCellularFirewallRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/cellularFirewallRules | Update the cellular firewall rules of an MX network
[**UpdateNetworkApplianceFirewallFirewalledService**](ApplianceAPI.md#UpdateNetworkApplianceFirewallFirewalledService) | **Put** /networks/{networkId}/appliance/firewall/firewalledServices/{service} | Updates the accessibility settings for the given service (&#39;ICMP&#39;, &#39;web&#39;, or &#39;SNMP&#39;)
[**UpdateNetworkApplianceFirewallInboundCellularFirewallRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallInboundCellularFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundCellularFirewallRules | Update the inbound cellular firewall rules of an MX network
[**UpdateNetworkApplianceFirewallInboundFirewallRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallInboundFirewallRules) | **Put** /networks/{networkId}/appliance/firewall/inboundFirewallRules | Update the inbound firewall rules of an MX network
[**UpdateNetworkApplianceFirewallL3FirewallRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallL3FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l3FirewallRules | Update the L3 firewall rules of an MX network
[**UpdateNetworkApplianceFirewallL7FirewallRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallL7FirewallRules) | **Put** /networks/{networkId}/appliance/firewall/l7FirewallRules | Update the MX L7 firewall rules for an MX network
[**UpdateNetworkApplianceFirewallOneToManyNatRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallOneToManyNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToManyNatRules | Set the 1:Many NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallOneToOneNatRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallOneToOneNatRules) | **Put** /networks/{networkId}/appliance/firewall/oneToOneNatRules | Set the 1:1 NAT mapping rules for an MX network
[**UpdateNetworkApplianceFirewallPortForwardingRules**](ApplianceAPI.md#UpdateNetworkApplianceFirewallPortForwardingRules) | **Put** /networks/{networkId}/appliance/firewall/portForwardingRules | Update the port forwarding rules for an MX network
[**UpdateNetworkApplianceFirewallSettings**](ApplianceAPI.md#UpdateNetworkApplianceFirewallSettings) | **Put** /networks/{networkId}/appliance/firewall/settings | Update the firewall settings for this network
[**UpdateNetworkAppliancePort**](ApplianceAPI.md#UpdateNetworkAppliancePort) | **Put** /networks/{networkId}/appliance/ports/{portId} | Update the per-port VLAN settings for a single MX port.
[**UpdateNetworkAppliancePrefixesDelegatedStatic**](ApplianceAPI.md#UpdateNetworkAppliancePrefixesDelegatedStatic) | **Put** /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId} | Update a static delegated prefix from a network
[**UpdateNetworkApplianceRfProfile**](ApplianceAPI.md#UpdateNetworkApplianceRfProfile) | **Put** /networks/{networkId}/appliance/rfProfiles/{rfProfileId} | Updates specified RF profile for this network
[**UpdateNetworkApplianceSdwanInternetPolicies**](ApplianceAPI.md#UpdateNetworkApplianceSdwanInternetPolicies) | **Put** /networks/{networkId}/appliance/sdwan/internetPolicies | Update SDWAN internet traffic preferences for an MX network
[**UpdateNetworkApplianceSecurityIntrusion**](ApplianceAPI.md#UpdateNetworkApplianceSecurityIntrusion) | **Put** /networks/{networkId}/appliance/security/intrusion | Set the supported intrusion settings for an MX network
[**UpdateNetworkApplianceSecurityMalware**](ApplianceAPI.md#UpdateNetworkApplianceSecurityMalware) | **Put** /networks/{networkId}/appliance/security/malware | Set the supported malware settings for an MX network
[**UpdateNetworkApplianceSettings**](ApplianceAPI.md#UpdateNetworkApplianceSettings) | **Put** /networks/{networkId}/appliance/settings | Update the appliance settings for a network
[**UpdateNetworkApplianceSingleLan**](ApplianceAPI.md#UpdateNetworkApplianceSingleLan) | **Put** /networks/{networkId}/appliance/singleLan | Update single LAN configuration
[**UpdateNetworkApplianceSsid**](ApplianceAPI.md#UpdateNetworkApplianceSsid) | **Put** /networks/{networkId}/appliance/ssids/{number} | Update the attributes of an MX SSID
[**UpdateNetworkApplianceStaticRoute**](ApplianceAPI.md#UpdateNetworkApplianceStaticRoute) | **Put** /networks/{networkId}/appliance/staticRoutes/{staticRouteId} | Update a static route for an MX or teleworker network
[**UpdateNetworkApplianceTrafficShaping**](ApplianceAPI.md#UpdateNetworkApplianceTrafficShaping) | **Put** /networks/{networkId}/appliance/trafficShaping | Update the traffic shaping settings for an MX network
[**UpdateNetworkApplianceTrafficShapingCustomPerformanceClass**](ApplianceAPI.md#UpdateNetworkApplianceTrafficShapingCustomPerformanceClass) | **Put** /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId} | Update a custom performance class for an MX network
[**UpdateNetworkApplianceTrafficShapingRules**](ApplianceAPI.md#UpdateNetworkApplianceTrafficShapingRules) | **Put** /networks/{networkId}/appliance/trafficShaping/rules | Update the traffic shaping settings rules for an MX network
[**UpdateNetworkApplianceTrafficShapingUplinkBandwidth**](ApplianceAPI.md#UpdateNetworkApplianceTrafficShapingUplinkBandwidth) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth | Updates the uplink bandwidth settings for your MX network.
[**UpdateNetworkApplianceTrafficShapingUplinkSelection**](ApplianceAPI.md#UpdateNetworkApplianceTrafficShapingUplinkSelection) | **Put** /networks/{networkId}/appliance/trafficShaping/uplinkSelection | Update uplink selection settings for an MX network
[**UpdateNetworkApplianceTrafficShapingVpnExclusions**](ApplianceAPI.md#UpdateNetworkApplianceTrafficShapingVpnExclusions) | **Put** /networks/{networkId}/appliance/trafficShaping/vpnExclusions | Update VPN exclusion rules for an MX network.
[**UpdateNetworkApplianceVlan**](ApplianceAPI.md#UpdateNetworkApplianceVlan) | **Put** /networks/{networkId}/appliance/vlans/{vlanId} | Update a VLAN
[**UpdateNetworkApplianceVlansSettings**](ApplianceAPI.md#UpdateNetworkApplianceVlansSettings) | **Put** /networks/{networkId}/appliance/vlans/settings | Enable/Disable VLANs for the given network
[**UpdateNetworkApplianceVpnBgp**](ApplianceAPI.md#UpdateNetworkApplianceVpnBgp) | **Put** /networks/{networkId}/appliance/vpn/bgp | Update a Hub BGP Configuration
[**UpdateNetworkApplianceVpnSiteToSiteVpn**](ApplianceAPI.md#UpdateNetworkApplianceVpnSiteToSiteVpn) | **Put** /networks/{networkId}/appliance/vpn/siteToSiteVpn | Update the site-to-site VPN settings of a network
[**UpdateNetworkApplianceWarmSpare**](ApplianceAPI.md#UpdateNetworkApplianceWarmSpare) | **Put** /networks/{networkId}/appliance/warmSpare | Update MX warm spare settings
[**UpdateOrganizationApplianceSecurityIntrusion**](ApplianceAPI.md#UpdateOrganizationApplianceSecurityIntrusion) | **Put** /organizations/{organizationId}/appliance/security/intrusion | Sets supported intrusion settings for an organization
[**UpdateOrganizationApplianceVpnThirdPartyVPNPeers**](ApplianceAPI.md#UpdateOrganizationApplianceVpnThirdPartyVPNPeers) | **Put** /organizations/{organizationId}/appliance/vpn/thirdPartyVPNPeers | Update the third party VPN peers for an organization
[**UpdateOrganizationApplianceVpnVpnFirewallRules**](ApplianceAPI.md#UpdateOrganizationApplianceVpnVpnFirewallRules) | **Put** /organizations/{organizationId}/appliance/vpn/vpnFirewallRules | Update the firewall rules of an organization&#39;s site-to-site VPN



## CreateDeviceApplianceVmxAuthenticationToken

> CreateDeviceApplianceVmxAuthenticationToken201Response CreateDeviceApplianceVmxAuthenticationToken(ctx, serial).Execute()

Generate a new vMX authentication token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.CreateDeviceApplianceVmxAuthenticationToken(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.CreateDeviceApplianceVmxAuthenticationToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDeviceApplianceVmxAuthenticationToken`: CreateDeviceApplianceVmxAuthenticationToken201Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.CreateDeviceApplianceVmxAuthenticationToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceApplianceVmxAuthenticationTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDeviceApplianceVmxAuthenticationToken201Response**](CreateDeviceApplianceVmxAuthenticationToken201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} CreateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId).CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest).Execute()

Add a static delegated prefix from a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkAppliancePrefixesDelegatedStaticRequest := *openapiclient.NewCreateNetworkAppliancePrefixesDelegatedStaticRequest("Prefix_example", *openapiclient.NewCreateNetworkAppliancePrefixesDelegatedStaticRequestOrigin()) // CreateNetworkAppliancePrefixesDelegatedStaticRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.CreateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId).CreateNetworkAppliancePrefixesDelegatedStaticRequest(createNetworkAppliancePrefixesDelegatedStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.CreateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.CreateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkAppliancePrefixesDelegatedStaticRequest** | [**CreateNetworkAppliancePrefixesDelegatedStaticRequest**](CreateNetworkAppliancePrefixesDelegatedStaticRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner CreateNetworkApplianceRfProfile(ctx, networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()

Creates new RF profile for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkApplianceRfProfileRequest := *openapiclient.NewCreateNetworkApplianceRfProfileRequest("Name_example") // CreateNetworkApplianceRfProfileRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.CreateNetworkApplianceRfProfile(context.Background(), networkId).CreateNetworkApplianceRfProfileRequest(createNetworkApplianceRfProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.CreateNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.CreateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceRfProfileRequest** | [**CreateNetworkApplianceRfProfileRequest**](CreateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceStaticRoute

> GetNetworkApplianceStaticRoutes200ResponseInner CreateNetworkApplianceStaticRoute(ctx, networkId).CreateNetworkApplianceStaticRouteRequest(createNetworkApplianceStaticRouteRequest).Execute()

Add a static route for an MX or teleworker network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkApplianceStaticRouteRequest := *openapiclient.NewCreateNetworkApplianceStaticRouteRequest("Name_example", "Subnet_example", "GatewayIp_example") // CreateNetworkApplianceStaticRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.CreateNetworkApplianceStaticRoute(context.Background(), networkId).CreateNetworkApplianceStaticRouteRequest(createNetworkApplianceStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.CreateNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceStaticRoute`: GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.CreateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceStaticRouteRequest** | [**CreateNetworkApplianceStaticRouteRequest**](CreateNetworkApplianceStaticRouteRequest.md) |  | 

### Return type

[**GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceTrafficShapingCustomPerformanceClass

> GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner CreateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(createNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()

Add a custom performance class for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkApplianceTrafficShapingCustomPerformanceClassRequest := *openapiclient.NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest("Name_example") // CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId).CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(createNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.CreateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceTrafficShapingCustomPerformanceClassRequest** | [**CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](CreateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNetworkApplianceVlan

> CreateNetworkApplianceVlan201Response CreateNetworkApplianceVlan(ctx, networkId).CreateNetworkApplianceVlanRequest(createNetworkApplianceVlanRequest).Execute()

Add a VLAN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	createNetworkApplianceVlanRequest := *openapiclient.NewCreateNetworkApplianceVlanRequest("Id_example", "Name_example") // CreateNetworkApplianceVlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.CreateNetworkApplianceVlan(context.Background(), networkId).CreateNetworkApplianceVlanRequest(createNetworkApplianceVlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.CreateNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkApplianceVlan`: CreateNetworkApplianceVlan201Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.CreateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createNetworkApplianceVlanRequest** | [**CreateNetworkApplianceVlanRequest**](CreateNetworkApplianceVlanRequest.md) |  | 

### Return type

[**CreateNetworkApplianceVlan201Response**](CreateNetworkApplianceVlan201Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkAppliancePrefixesDelegatedStatic

> DeleteNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Delete a static delegated prefix from a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplianceAPI.DeleteNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.DeleteNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceRfProfile

> DeleteNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Delete a RF Profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	rfProfileId := "rfProfileId_example" // string | Rf profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplianceAPI.DeleteNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.DeleteNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceStaticRoute

> DeleteNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Delete a static route from an MX or teleworker network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplianceAPI.DeleteNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.DeleteNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceTrafficShapingCustomPerformanceClass

> DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Delete a custom performance class from an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplianceAPI.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.DeleteNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkApplianceVlan

> DeleteNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Delete a VLAN from a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	vlanId := "vlanId_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ApplianceAPI.DeleteNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.DeleteNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceDhcpSubnets

> []map[string]interface{} GetDeviceApplianceDhcpSubnets(ctx, serial).Execute()

Return the DHCP subnet information for an appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetDeviceApplianceDhcpSubnets(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetDeviceApplianceDhcpSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceApplianceDhcpSubnets`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetDeviceApplianceDhcpSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceDhcpSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePerformance

> map[string]interface{} GetDeviceAppliancePerformance(ctx, serial).T0(t0).T1(t1).Timespan(timespan).Execute()

Return the performance score for a single MX



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 30 minutes and be less than or equal to 14 days. The default is 30 minutes. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetDeviceAppliancePerformance(context.Background(), serial).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetDeviceAppliancePerformance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAppliancePerformance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetDeviceAppliancePerformance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 30 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 30 minutes and be less than or equal to 14 days. The default is 30 minutes. | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePrefixesDelegated

> []map[string]interface{} GetDeviceAppliancePrefixesDelegated(ctx, serial).Execute()

Return current delegated IPv6 prefixes on an appliance.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetDeviceAppliancePrefixesDelegated(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetDeviceAppliancePrefixesDelegated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAppliancePrefixesDelegated`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetDeviceAppliancePrefixesDelegated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceAppliancePrefixesDelegatedVlanAssignments

> []map[string]interface{} GetDeviceAppliancePrefixesDelegatedVlanAssignments(ctx, serial).Execute()

Return prefixes assigned to all IPv6 enabled VLANs on an appliance.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceAppliancePrefixesDelegatedVlanAssignments`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetDeviceAppliancePrefixesDelegatedVlanAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceAppliancePrefixesDelegatedVlanAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceRadioSettings

> GetDeviceApplianceRadioSettings200Response GetDeviceApplianceRadioSettings(ctx, serial).Execute()

Return the radio settings of an appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetDeviceApplianceRadioSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetDeviceApplianceRadioSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceApplianceRadioSettings`: GetDeviceApplianceRadioSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceApplianceRadioSettings200Response**](GetDeviceApplianceRadioSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceApplianceUplinksSettings

> GetDeviceApplianceUplinksSettings200Response GetDeviceApplianceUplinksSettings(ctx, serial).Execute()

Return the uplink settings for an MX appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetDeviceApplianceUplinksSettings(context.Background(), serial).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetDeviceApplianceUplinksSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDeviceApplianceUplinksSettings`: GetDeviceApplianceUplinksSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetDeviceApplianceUplinksSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceApplianceUplinksSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDeviceApplianceUplinksSettings200Response**](GetDeviceApplianceUplinksSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceClientSecurityEvents

> []map[string]interface{} GetNetworkApplianceClientSecurityEvents(ctx, networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a client



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	clientId := "clientId_example" // string | Client ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceClientSecurityEvents(context.Background(), networkId, clientId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceClientSecurityEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceClientSecurityEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceClientSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**clientId** | **string** | Client ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceClientSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 791 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 791 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 791 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceConnectivityMonitoringDestinations

> GetNetworkApplianceConnectivityMonitoringDestinations200Response GetNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).Execute()

Return the connectivity testing destinations for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceConnectivityMonitoringDestinations`: GetNetworkApplianceConnectivityMonitoringDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceConnectivityMonitoringDestinations200Response**](GetNetworkApplianceConnectivityMonitoringDestinations200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceContentFiltering

> map[string]interface{} GetNetworkApplianceContentFiltering(ctx, networkId).Execute()

Return the content filtering settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceContentFiltering(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceContentFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceContentFiltering`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceContentFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceContentFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceContentFilteringCategories

> map[string]interface{} GetNetworkApplianceContentFilteringCategories(ctx, networkId).Execute()

List all available content filtering categories for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceContentFilteringCategories(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceContentFilteringCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceContentFilteringCategories`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceContentFilteringCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceContentFilteringCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallCellularFirewallRules

> map[string]interface{} GetNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).Execute()

Return the cellular firewall rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallCellularFirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallFirewalledService

> GetNetworkApplianceFirewallFirewalledServices200ResponseInner GetNetworkApplianceFirewallFirewalledService(ctx, networkId, service).Execute()

Return the accessibility settings of the given service ('ICMP', 'web', or 'SNMP')



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	service := "service_example" // string | Service

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallFirewalledService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallFirewalledService`: GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceFirewallFirewalledServices200ResponseInner**](GetNetworkApplianceFirewallFirewalledServices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallFirewalledServices

> []GetNetworkApplianceFirewallFirewalledServices200ResponseInner GetNetworkApplianceFirewallFirewalledServices(ctx, networkId).Execute()

List the appliance services and their accessibility rules



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallFirewalledServices(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallFirewalledServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallFirewalledServices`: []GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallFirewalledServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallFirewalledServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceFirewallFirewalledServices200ResponseInner**](GetNetworkApplianceFirewallFirewalledServices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallInboundCellularFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response GetNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).Execute()

Return the inbound cellular firewall rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallInboundCellularFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallInboundFirewallRules

> GetNetworkApplianceFirewallInboundFirewallRules200Response GetNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).Execute()

Return the inbound firewall rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallInboundFirewallRules`: GetNetworkApplianceFirewallInboundFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallInboundFirewallRules200Response**](GetNetworkApplianceFirewallInboundFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallL3FirewallRules

> map[string]interface{} GetNetworkApplianceFirewallL3FirewallRules(ctx, networkId).Execute()

Return the L3 firewall rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallL3FirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallL7FirewallRules

> map[string]interface{} GetNetworkApplianceFirewallL7FirewallRules(ctx, networkId).Execute()

List the MX L7 firewall rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallL7FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallL7FirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories

> GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(ctx, networkId).Execute()

Return the L7 firewall application categories and their associated applications for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response**](GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallOneToManyNatRules

> map[string]interface{} GetNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).Execute()

Return the 1:Many NAT mapping rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallOneToManyNatRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallOneToOneNatRules

> map[string]interface{} GetNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).Execute()

Return the 1:1 NAT mapping rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallOneToOneNatRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallPortForwardingRules

> GetNetworkApplianceFirewallPortForwardingRules200Response GetNetworkApplianceFirewallPortForwardingRules(ctx, networkId).Execute()

Return the port forwarding rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallPortForwardingRules`: GetNetworkApplianceFirewallPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallPortForwardingRules200Response**](GetNetworkApplianceFirewallPortForwardingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceFirewallSettings

> map[string]interface{} GetNetworkApplianceFirewallSettings(ctx, networkId).Execute()

Return the firewall settings for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceFirewallSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceFirewallSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceFirewallSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceFirewallSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceFirewallSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePort

> GetNetworkAppliancePorts200ResponseInner GetNetworkAppliancePort(ctx, networkId, portId).Execute()

Return per-port VLAN settings for a single MX port.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	portId := "portId_example" // string | Port ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkAppliancePort(context.Background(), networkId, portId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkAppliancePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePort`: GetNetworkAppliancePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePorts

> []GetNetworkAppliancePorts200ResponseInner GetNetworkAppliancePorts(ctx, networkId).Execute()

List per-port VLAN settings for all ports of a MX.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkAppliancePorts(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkAppliancePorts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePorts`: []GetNetworkAppliancePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkAppliancePorts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePortsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatic

> GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner GetNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).Execute()

Return a static delegated prefix from a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePrefixesDelegatedStatic`: GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkAppliancePrefixesDelegatedStatics

> []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner GetNetworkAppliancePrefixesDelegatedStatics(ctx, networkId).Execute()

List static delegated prefixes for a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkAppliancePrefixesDelegatedStatics(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkAppliancePrefixesDelegatedStatics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkAppliancePrefixesDelegatedStatics`: []GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkAppliancePrefixesDelegatedStatics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkAppliancePrefixesDelegatedStaticsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner**](GetNetworkAppliancePrefixesDelegatedStatics200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner GetNetworkApplianceRfProfile(ctx, networkId, rfProfileId).Execute()

Return a RF profile



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	rfProfileId := "rfProfileId_example" // string | Rf profile ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceRfProfiles

> GetNetworkApplianceRfProfiles200Response GetNetworkApplianceRfProfiles(ctx, networkId).Execute()

List the RF profiles for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceRfProfiles(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceRfProfiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceRfProfiles`: GetNetworkApplianceRfProfiles200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceRfProfiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceRfProfilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceRfProfiles200Response**](GetNetworkApplianceRfProfiles200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityEvents

> []map[string]interface{} GetNetworkApplianceSecurityEvents(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSecurityEvents(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSecurityEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSecurityEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityIntrusion

> GetNetworkApplianceSecurityIntrusion200Response GetNetworkApplianceSecurityIntrusion(ctx, networkId).Execute()

Returns all supported intrusion settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSecurityIntrusion(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSecurityIntrusion`: GetNetworkApplianceSecurityIntrusion200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSecurityIntrusion200Response**](GetNetworkApplianceSecurityIntrusion200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSecurityMalware

> GetNetworkApplianceSecurityMalware200Response GetNetworkApplianceSecurityMalware(ctx, networkId).Execute()

Returns all supported malware settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSecurityMalware(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSecurityMalware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSecurityMalware`: GetNetworkApplianceSecurityMalware200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSecurityMalware200Response**](GetNetworkApplianceSecurityMalware200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSettings

> GetNetworkApplianceSettings200Response GetNetworkApplianceSettings(ctx, networkId).Execute()

Return the appliance settings for a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSettings`: GetNetworkApplianceSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSettings200Response**](GetNetworkApplianceSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSingleLan

> GetNetworkApplianceSingleLan200Response GetNetworkApplianceSingleLan(ctx, networkId).Execute()

Return single LAN configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSingleLan(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSingleLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSingleLan`: GetNetworkApplianceSingleLan200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSingleLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSingleLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceSingleLan200Response**](GetNetworkApplianceSingleLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsid

> GetNetworkApplianceSsids200ResponseInner GetNetworkApplianceSsid(ctx, networkId, number).Execute()

Return a single MX SSID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	number := "number_example" // string | Number

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSsid(context.Background(), networkId, number).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSsid`: GetNetworkApplianceSsids200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceSsids200ResponseInner**](GetNetworkApplianceSsids200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceSsids

> []GetNetworkApplianceSsids200ResponseInner GetNetworkApplianceSsids(ctx, networkId).Execute()

List the MX SSIDs in a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceSsids(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceSsids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceSsids`: []GetNetworkApplianceSsids200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceSsids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceSsidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceSsids200ResponseInner**](GetNetworkApplianceSsids200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceStaticRoute

> GetNetworkApplianceStaticRoutes200ResponseInner GetNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).Execute()

Return a static route for an MX or teleworker network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	staticRouteId := "staticRouteId_example" // string | Static route ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceStaticRoute`: GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceStaticRoutes

> []GetNetworkApplianceStaticRoutes200ResponseInner GetNetworkApplianceStaticRoutes(ctx, networkId).Execute()

List the static routes for an MX or teleworker network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceStaticRoutes(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceStaticRoutes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceStaticRoutes`: []GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceStaticRoutes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceStaticRoutesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShaping

> map[string]interface{} GetNetworkApplianceTrafficShaping(ctx, networkId).Execute()

Display the traffic shaping settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceTrafficShaping(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceTrafficShaping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShaping`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceTrafficShaping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingCustomPerformanceClass

> GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner GetNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).Execute()

Return a custom performance class for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingCustomPerformanceClass`: GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingCustomPerformanceClasses

> []GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner GetNetworkApplianceTrafficShapingCustomPerformanceClasses(ctx, networkId).Execute()

List all custom performance classes for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: []GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceTrafficShapingCustomPerformanceClasses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingCustomPerformanceClassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingRules

> map[string]interface{} GetNetworkApplianceTrafficShapingRules(ctx, networkId).Execute()

Display the traffic shaping settings rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceTrafficShapingRules(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceTrafficShapingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingUplinkBandwidth

> GetNetworkApplianceTrafficShapingUplinkBandwidth200Response GetNetworkApplianceTrafficShapingUplinkBandwidth(ctx, networkId).Execute()

Returns the uplink bandwidth limits for your MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceTrafficShapingUplinkBandwidth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingUplinkBandwidth`: GetNetworkApplianceTrafficShapingUplinkBandwidth200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceTrafficShapingUplinkBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceTrafficShapingUplinkBandwidth200Response**](GetNetworkApplianceTrafficShapingUplinkBandwidth200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceTrafficShapingUplinkSelection

> GetNetworkApplianceTrafficShapingUplinkSelection200Response GetNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).Execute()

Show uplink selection settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceTrafficShapingUplinkSelection`: GetNetworkApplianceTrafficShapingUplinkSelection200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceTrafficShapingUplinkSelection200Response**](GetNetworkApplianceTrafficShapingUplinkSelection200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceUplinksUsageHistory

> []map[string]interface{} GetNetworkApplianceUplinksUsageHistory(ctx, networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()

Get the sent and received bytes for each uplink of a network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes. (optional)
	resolution := int32(56) // int32 | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceUplinksUsageHistory(context.Background(), networkId).T0(t0).T1(t1).Timespan(timespan).Resolution(resolution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceUplinksUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceUplinksUsageHistory`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceUplinksUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceUplinksUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 10 minutes. | 
 **resolution** | **int32** | The time resolution in seconds for returned data. The valid resolutions are: 60, 300, 600, 1800, 3600, 86400. The default is 60. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlan

> GetNetworkApplianceVlans200ResponseInner GetNetworkApplianceVlan(ctx, networkId, vlanId).Execute()

Return a VLAN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	vlanId := "vlanId_example" // string | Vlan ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceVlan(context.Background(), networkId, vlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVlan`: GetNetworkApplianceVlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlans

> []GetNetworkApplianceVlans200ResponseInner GetNetworkApplianceVlans(ctx, networkId).Execute()

List the VLANs for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceVlans(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVlans`: []GetNetworkApplianceVlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVlansSettings

> GetNetworkApplianceVlansSettings200Response GetNetworkApplianceVlansSettings(ctx, networkId).Execute()

Returns the enabled status of VLANs for the network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceVlansSettings(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceVlansSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVlansSettings`: GetNetworkApplianceVlansSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVlansSettings200Response**](GetNetworkApplianceVlansSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVpnBgp

> GetNetworkApplianceVpnBgp200Response GetNetworkApplianceVpnBgp(ctx, networkId).Execute()

Return a Hub BGP Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceVpnBgp(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceVpnBgp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVpnBgp`: GetNetworkApplianceVpnBgp200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVpnBgp200Response**](GetNetworkApplianceVpnBgp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceVpnSiteToSiteVpn

> GetNetworkApplianceVpnSiteToSiteVpn200Response GetNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).Execute()

Return the site-to-site VPN settings of a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceVpnSiteToSiteVpn`: GetNetworkApplianceVpnSiteToSiteVpn200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceVpnSiteToSiteVpn200Response**](GetNetworkApplianceVpnSiteToSiteVpn200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkApplianceWarmSpare

> GetNetworkApplianceWarmSpare200Response GetNetworkApplianceWarmSpare(ctx, networkId).Execute()

Return MX warm spare settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetNetworkApplianceWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkApplianceWarmSpare`: GetNetworkApplianceWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceWarmSpare200Response**](GetNetworkApplianceWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityEvents

> []map[string]interface{} GetOrganizationApplianceSecurityEvents(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()

List the security events for an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. (optional)
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	sortOrder := "sortOrder_example" // string | Sorted order of security events based on event detection time. Order options are 'ascending' or 'descending'. Default is ascending order. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceSecurityEvents(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceSecurityEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceSecurityEvents`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceSecurityEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. Data is gathered after the specified t0 value. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 365 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 365 days. The default is 31 days. | 
 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 100. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **sortOrder** | **string** | Sorted order of security events based on event detection time. Order options are &#39;ascending&#39; or &#39;descending&#39;. Default is ascending order. | 

### Return type

[**[]map[string]interface{}**](map.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceSecurityIntrusion

> map[string]interface{} GetOrganizationApplianceSecurityIntrusion(ctx, organizationId).Execute()

Returns all supported intrusion settings for an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceSecurityIntrusion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork

> GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork200Response GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Display VPN exclusion rules for MX networks.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | Optional parameter to filter the results by network IDs (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork`: GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceTrafficShapingVpnExclusionsByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 50. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | Optional parameter to filter the results by network IDs | 

### Return type

[**GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork200Response**](GetOrganizationApplianceTrafficShapingVpnExclusionsByNetwork200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinkStatuses

> []GetOrganizationApplianceUplinkStatuses200ResponseInner GetOrganizationApplianceUplinkStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()

List the uplink status of every Meraki MX and Z series appliances in the organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | A list of network IDs. The returned devices will be filtered to only include these networks. (optional)
	serials := []string{"Inner_example"} // []string | A list of serial numbers. The returned devices will be filtered to only include these serials. (optional)
	iccids := []string{"Inner_example"} // []string | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceUplinkStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Serials(serials).Iccids(iccids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceUplinkStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceUplinkStatuses`: []GetOrganizationApplianceUplinkStatuses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceUplinkStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinkStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 1000. Default is 1000. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of network IDs. The returned devices will be filtered to only include these networks. | 
 **serials** | **[]string** | A list of serial numbers. The returned devices will be filtered to only include these serials. | 
 **iccids** | **[]string** | A list of ICCIDs. The returned devices will be filtered to only include these ICCIDs. | 

### Return type

[**[]GetOrganizationApplianceUplinkStatuses200ResponseInner**](GetOrganizationApplianceUplinkStatuses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksStatusesOverview

> GetOrganizationApplianceUplinksStatusesOverview200Response GetOrganizationApplianceUplinksStatusesOverview(ctx, organizationId).Execute()

Returns an overview of uplink statuses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceUplinksStatusesOverview(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceUplinksStatusesOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceUplinksStatusesOverview`: GetOrganizationApplianceUplinksStatusesOverview200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceUplinksStatusesOverview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksStatusesOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationApplianceUplinksStatusesOverview200Response**](GetOrganizationApplianceUplinksStatusesOverview200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceUplinksUsageByNetwork

> []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner GetOrganizationApplianceUplinksUsageByNetwork(ctx, organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()

Get the sent and received bytes for each uplink of all MX and Z networks within an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceUplinksUsageByNetwork(context.Background(), organizationId).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceUplinksUsageByNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceUplinksUsageByNetwork`: []GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceUplinksUsageByNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceUplinksUsageByNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 365 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 14 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 14 days. The default is 1 day. | 

### Return type

[**[]GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner**](GetOrganizationApplianceUplinksUsageByNetwork200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnStats

> []map[string]interface{} GetOrganizationApplianceVpnStats(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).T0(t0).T1(t1).Timespan(timespan).Execute()

Show VPN history stat for networks in an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)
	t0 := "t0_example" // string | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. (optional)
	t1 := "t1_example" // string | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. (optional)
	timespan := float32(3.4) // float32 | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceVpnStats(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).T0(t0).T1(t1).Timespan(timespan).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceVpnStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceVpnStats`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceVpnStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 
 **t0** | **string** | The beginning of the timespan for the data. The maximum lookback period is 31 days from today. | 
 **t1** | **string** | The end of the timespan for the data. t1 can be a maximum of 31 days after t0. | 
 **timespan** | **float32** | The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 31 days. The default is 1 day. | 

### Return type

**[]map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnStatuses

> GetOrganizationApplianceVpnStatuses200Response GetOrganizationApplianceVpnStatuses(ctx, organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()

Show VPN status for networks in an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	perPage := int32(56) // int32 | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. (optional)
	startingAfter := "startingAfter_example" // string | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	endingBefore := "endingBefore_example" // string | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. (optional)
	networkIds := []string{"Inner_example"} // []string | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]=N_12345678&networkIds[]=L_3456 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceVpnStatuses(context.Background(), organizationId).PerPage(perPage).StartingAfter(startingAfter).EndingBefore(endingBefore).NetworkIds(networkIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceVpnStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceVpnStatuses`: GetOrganizationApplianceVpnStatuses200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceVpnStatuses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | The number of entries per page returned. Acceptable range is 3 - 300. Default is 300. | 
 **startingAfter** | **string** | A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **endingBefore** | **string** | A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it. | 
 **networkIds** | **[]string** | A list of Meraki network IDs to filter results to contain only specified networks. E.g.: networkIds[]&#x3D;N_12345678&amp;networkIds[]&#x3D;L_3456 | 

### Return type

[**GetOrganizationApplianceVpnStatuses200Response**](GetOrganizationApplianceVpnStatuses200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnThirdPartyVPNPeers

> GetOrganizationApplianceVpnThirdPartyVPNPeers200Response GetOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).Execute()

Return the third party VPN peers for an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceVpnThirdPartyVPNPeers`: GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrganizationApplianceVpnThirdPartyVPNPeers200Response**](GetOrganizationApplianceVpnThirdPartyVPNPeers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationApplianceVpnVpnFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response GetOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).Execute()

Return the firewall rules for an organization's site-to-site VPN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.GetOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.GetOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizationApplianceVpnVpnFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.GetOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwapNetworkApplianceWarmSpare

> GetNetworkApplianceWarmSpare200Response SwapNetworkApplianceWarmSpare(ctx, networkId).Execute()

Swap MX primary and warm spare appliances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.SwapNetworkApplianceWarmSpare(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.SwapNetworkApplianceWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SwapNetworkApplianceWarmSpare`: GetNetworkApplianceWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.SwapNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkApplianceWarmSpare200Response**](GetNetworkApplianceWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceApplianceRadioSettings

> GetDeviceApplianceRadioSettings200Response UpdateDeviceApplianceRadioSettings(ctx, serial).UpdateDeviceApplianceRadioSettingsRequest(updateDeviceApplianceRadioSettingsRequest).Execute()

Update the radio settings of an appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	updateDeviceApplianceRadioSettingsRequest := *openapiclient.NewUpdateDeviceApplianceRadioSettingsRequest() // UpdateDeviceApplianceRadioSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateDeviceApplianceRadioSettings(context.Background(), serial).UpdateDeviceApplianceRadioSettingsRequest(updateDeviceApplianceRadioSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateDeviceApplianceRadioSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceApplianceRadioSettings`: GetDeviceApplianceRadioSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateDeviceApplianceRadioSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceRadioSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceRadioSettingsRequest** | [**UpdateDeviceApplianceRadioSettingsRequest**](UpdateDeviceApplianceRadioSettingsRequest.md) |  | 

### Return type

[**GetDeviceApplianceRadioSettings200Response**](GetDeviceApplianceRadioSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceApplianceUplinksSettings

> GetDeviceApplianceUplinksSettings200Response UpdateDeviceApplianceUplinksSettings(ctx, serial).UpdateDeviceApplianceUplinksSettingsRequest(updateDeviceApplianceUplinksSettingsRequest).Execute()

Update the uplink settings for an MX appliance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	serial := "serial_example" // string | Serial
	updateDeviceApplianceUplinksSettingsRequest := *openapiclient.NewUpdateDeviceApplianceUplinksSettingsRequest(*openapiclient.NewUpdateDeviceApplianceUplinksSettingsRequestInterfaces()) // UpdateDeviceApplianceUplinksSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateDeviceApplianceUplinksSettings(context.Background(), serial).UpdateDeviceApplianceUplinksSettingsRequest(updateDeviceApplianceUplinksSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateDeviceApplianceUplinksSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDeviceApplianceUplinksSettings`: GetDeviceApplianceUplinksSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateDeviceApplianceUplinksSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serial** | **string** | Serial | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceApplianceUplinksSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDeviceApplianceUplinksSettingsRequest** | [**UpdateDeviceApplianceUplinksSettingsRequest**](UpdateDeviceApplianceUplinksSettingsRequest.md) |  | 

### Return type

[**GetDeviceApplianceUplinksSettings200Response**](GetDeviceApplianceUplinksSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceConnectivityMonitoringDestinations

> GetNetworkApplianceConnectivityMonitoringDestinations200Response UpdateNetworkApplianceConnectivityMonitoringDestinations(ctx, networkId).UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest(updateNetworkApplianceConnectivityMonitoringDestinationsRequest).Execute()

Update the connectivity testing destinations for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceConnectivityMonitoringDestinationsRequest := *openapiclient.NewUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest() // UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceConnectivityMonitoringDestinations(context.Background(), networkId).UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest(updateNetworkApplianceConnectivityMonitoringDestinationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceConnectivityMonitoringDestinations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceConnectivityMonitoringDestinations`: GetNetworkApplianceConnectivityMonitoringDestinations200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceConnectivityMonitoringDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceConnectivityMonitoringDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceConnectivityMonitoringDestinationsRequest** | [**UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest**](UpdateNetworkApplianceConnectivityMonitoringDestinationsRequest.md) |  | 

### Return type

[**GetNetworkApplianceConnectivityMonitoringDestinations200Response**](GetNetworkApplianceConnectivityMonitoringDestinations200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceContentFiltering

> map[string]interface{} UpdateNetworkApplianceContentFiltering(ctx, networkId).UpdateNetworkApplianceContentFilteringRequest(updateNetworkApplianceContentFilteringRequest).Execute()

Update the content filtering settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceContentFilteringRequest := *openapiclient.NewUpdateNetworkApplianceContentFilteringRequest() // UpdateNetworkApplianceContentFilteringRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceContentFiltering(context.Background(), networkId).UpdateNetworkApplianceContentFilteringRequest(updateNetworkApplianceContentFilteringRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceContentFiltering``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceContentFiltering`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceContentFiltering`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceContentFilteringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceContentFilteringRequest** | [**UpdateNetworkApplianceContentFilteringRequest**](UpdateNetworkApplianceContentFilteringRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallCellularFirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()

Update the cellular firewall rules of an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallCellularFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallCellularFirewallRulesRequest() // UpdateNetworkApplianceFirewallCellularFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallCellularFirewallRulesRequest(updateNetworkApplianceFirewallCellularFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallCellularFirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallCellularFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallCellularFirewallRulesRequest**](UpdateNetworkApplianceFirewallCellularFirewallRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallFirewalledService

> GetNetworkApplianceFirewallFirewalledServices200ResponseInner UpdateNetworkApplianceFirewallFirewalledService(ctx, networkId, service).UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest).Execute()

Updates the accessibility settings for the given service ('ICMP', 'web', or 'SNMP')



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	service := "service_example" // string | Service
	updateNetworkApplianceFirewallFirewalledServiceRequest := *openapiclient.NewUpdateNetworkApplianceFirewallFirewalledServiceRequest("Access_example") // UpdateNetworkApplianceFirewallFirewalledServiceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallFirewalledService(context.Background(), networkId, service).UpdateNetworkApplianceFirewallFirewalledServiceRequest(updateNetworkApplianceFirewallFirewalledServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallFirewalledService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallFirewalledService`: GetNetworkApplianceFirewallFirewalledServices200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallFirewalledService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**service** | **string** | Service | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallFirewalledServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceFirewallFirewalledServiceRequest** | [**UpdateNetworkApplianceFirewallFirewalledServiceRequest**](UpdateNetworkApplianceFirewallFirewalledServiceRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallFirewalledServices200ResponseInner**](GetNetworkApplianceFirewallFirewalledServices200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundCellularFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response UpdateNetworkApplianceFirewallInboundCellularFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest(updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest).Execute()

Update the inbound cellular firewall rules of an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest() // UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallInboundCellularFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest(updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallInboundCellularFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallInboundCellularFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundCellularFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest**](UpdateNetworkApplianceFirewallInboundCellularFirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallInboundFirewallRules

> GetNetworkApplianceFirewallInboundFirewallRules200Response UpdateNetworkApplianceFirewallInboundFirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()

Update the inbound firewall rules of an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallInboundFirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallInboundFirewallRulesRequest() // UpdateNetworkApplianceFirewallInboundFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallInboundFirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallInboundFirewallRulesRequest(updateNetworkApplianceFirewallInboundFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallInboundFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallInboundFirewallRules`: GetNetworkApplianceFirewallInboundFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallInboundFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallInboundFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallInboundFirewallRulesRequest** | [**UpdateNetworkApplianceFirewallInboundFirewallRulesRequest**](UpdateNetworkApplianceFirewallInboundFirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallInboundFirewallRules200Response**](GetNetworkApplianceFirewallInboundFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallL3FirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallL3FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallL3FirewallRulesRequest(updateNetworkApplianceFirewallL3FirewallRulesRequest).Execute()

Update the L3 firewall rules of an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallL3FirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallL3FirewallRulesRequest() // UpdateNetworkApplianceFirewallL3FirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallL3FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallL3FirewallRulesRequest(updateNetworkApplianceFirewallL3FirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallL3FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallL3FirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallL3FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL3FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallL3FirewallRulesRequest** | [**UpdateNetworkApplianceFirewallL3FirewallRulesRequest**](UpdateNetworkApplianceFirewallL3FirewallRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallL7FirewallRules

> map[string]interface{} UpdateNetworkApplianceFirewallL7FirewallRules(ctx, networkId).UpdateNetworkApplianceFirewallL7FirewallRulesRequest(updateNetworkApplianceFirewallL7FirewallRulesRequest).Execute()

Update the MX L7 firewall rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallL7FirewallRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallL7FirewallRulesRequest() // UpdateNetworkApplianceFirewallL7FirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallL7FirewallRules(context.Background(), networkId).UpdateNetworkApplianceFirewallL7FirewallRulesRequest(updateNetworkApplianceFirewallL7FirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallL7FirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallL7FirewallRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallL7FirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallL7FirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallL7FirewallRulesRequest** | [**UpdateNetworkApplianceFirewallL7FirewallRulesRequest**](UpdateNetworkApplianceFirewallL7FirewallRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallOneToManyNatRules

> map[string]interface{} UpdateNetworkApplianceFirewallOneToManyNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToManyNatRulesRequest(updateNetworkApplianceFirewallOneToManyNatRulesRequest).Execute()

Set the 1:Many NAT mapping rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallOneToManyNatRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInner("PublicIp_example", "Uplink_example", []openapiclient.UpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToManyNatRulesRequestRulesInnerPortRulesInner()})}) // UpdateNetworkApplianceFirewallOneToManyNatRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallOneToManyNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToManyNatRulesRequest(updateNetworkApplianceFirewallOneToManyNatRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallOneToManyNatRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallOneToManyNatRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallOneToManyNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToManyNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToManyNatRulesRequest** | [**UpdateNetworkApplianceFirewallOneToManyNatRulesRequest**](UpdateNetworkApplianceFirewallOneToManyNatRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallOneToOneNatRules

> map[string]interface{} UpdateNetworkApplianceFirewallOneToOneNatRules(ctx, networkId).UpdateNetworkApplianceFirewallOneToOneNatRulesRequest(updateNetworkApplianceFirewallOneToOneNatRulesRequest).Execute()

Set the 1:1 NAT mapping rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallOneToOneNatRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallOneToOneNatRulesRequestRulesInner("LanIp_example")}) // UpdateNetworkApplianceFirewallOneToOneNatRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallOneToOneNatRules(context.Background(), networkId).UpdateNetworkApplianceFirewallOneToOneNatRulesRequest(updateNetworkApplianceFirewallOneToOneNatRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallOneToOneNatRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallOneToOneNatRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallOneToOneNatRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallOneToOneNatRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallOneToOneNatRulesRequest** | [**UpdateNetworkApplianceFirewallOneToOneNatRulesRequest**](UpdateNetworkApplianceFirewallOneToOneNatRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallPortForwardingRules

> GetNetworkApplianceFirewallPortForwardingRules200Response UpdateNetworkApplianceFirewallPortForwardingRules(ctx, networkId).UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest).Execute()

Update the port forwarding rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallPortForwardingRulesRequest := *openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequest([]openapiclient.UpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner{*openapiclient.NewUpdateNetworkApplianceFirewallPortForwardingRulesRequestRulesInner("LanIp_example", "PublicPort_example", "LocalPort_example", []string{"AllowedIps_example"}, "Protocol_example")}) // UpdateNetworkApplianceFirewallPortForwardingRulesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallPortForwardingRules(context.Background(), networkId).UpdateNetworkApplianceFirewallPortForwardingRulesRequest(updateNetworkApplianceFirewallPortForwardingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallPortForwardingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallPortForwardingRules`: GetNetworkApplianceFirewallPortForwardingRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallPortForwardingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallPortForwardingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallPortForwardingRulesRequest** | [**UpdateNetworkApplianceFirewallPortForwardingRulesRequest**](UpdateNetworkApplianceFirewallPortForwardingRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallPortForwardingRules200Response**](GetNetworkApplianceFirewallPortForwardingRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceFirewallSettings

> map[string]interface{} UpdateNetworkApplianceFirewallSettings(ctx, networkId).UpdateNetworkApplianceFirewallSettingsRequest(updateNetworkApplianceFirewallSettingsRequest).Execute()

Update the firewall settings for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceFirewallSettingsRequest := *openapiclient.NewUpdateNetworkApplianceFirewallSettingsRequest() // UpdateNetworkApplianceFirewallSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceFirewallSettings(context.Background(), networkId).UpdateNetworkApplianceFirewallSettingsRequest(updateNetworkApplianceFirewallSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceFirewallSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceFirewallSettings`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceFirewallSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceFirewallSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceFirewallSettingsRequest** | [**UpdateNetworkApplianceFirewallSettingsRequest**](UpdateNetworkApplianceFirewallSettingsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePort

> GetNetworkAppliancePorts200ResponseInner UpdateNetworkAppliancePort(ctx, networkId, portId).UpdateNetworkAppliancePortRequest(updateNetworkAppliancePortRequest).Execute()

Update the per-port VLAN settings for a single MX port.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	portId := "portId_example" // string | Port ID
	updateNetworkAppliancePortRequest := *openapiclient.NewUpdateNetworkAppliancePortRequest() // UpdateNetworkAppliancePortRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkAppliancePort(context.Background(), networkId, portId).UpdateNetworkAppliancePortRequest(updateNetworkAppliancePortRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkAppliancePort``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkAppliancePort`: GetNetworkAppliancePorts200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkAppliancePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**portId** | **string** | Port ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePortRequest** | [**UpdateNetworkAppliancePortRequest**](UpdateNetworkAppliancePortRequest.md) |  | 

### Return type

[**GetNetworkAppliancePorts200ResponseInner**](GetNetworkAppliancePorts200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkAppliancePrefixesDelegatedStatic

> map[string]interface{} UpdateNetworkAppliancePrefixesDelegatedStatic(ctx, networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest).Execute()

Update a static delegated prefix from a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	staticDelegatedPrefixId := "staticDelegatedPrefixId_example" // string | Static delegated prefix ID
	updateNetworkAppliancePrefixesDelegatedStaticRequest := *openapiclient.NewUpdateNetworkAppliancePrefixesDelegatedStaticRequest() // UpdateNetworkAppliancePrefixesDelegatedStaticRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkAppliancePrefixesDelegatedStatic(context.Background(), networkId, staticDelegatedPrefixId).UpdateNetworkAppliancePrefixesDelegatedStaticRequest(updateNetworkAppliancePrefixesDelegatedStaticRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkAppliancePrefixesDelegatedStatic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkAppliancePrefixesDelegatedStatic`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkAppliancePrefixesDelegatedStatic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticDelegatedPrefixId** | **string** | Static delegated prefix ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkAppliancePrefixesDelegatedStaticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkAppliancePrefixesDelegatedStaticRequest** | [**UpdateNetworkAppliancePrefixesDelegatedStaticRequest**](UpdateNetworkAppliancePrefixesDelegatedStaticRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceRfProfile

> GetNetworkApplianceRfProfiles200ResponseAssignedInner UpdateNetworkApplianceRfProfile(ctx, networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()

Updates specified RF profile for this network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	rfProfileId := "rfProfileId_example" // string | Rf profile ID
	updateNetworkApplianceRfProfileRequest := *openapiclient.NewUpdateNetworkApplianceRfProfileRequest() // UpdateNetworkApplianceRfProfileRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceRfProfile(context.Background(), networkId, rfProfileId).UpdateNetworkApplianceRfProfileRequest(updateNetworkApplianceRfProfileRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceRfProfile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceRfProfile`: GetNetworkApplianceRfProfiles200ResponseAssignedInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceRfProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**rfProfileId** | **string** | Rf profile ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceRfProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceRfProfileRequest** | [**UpdateNetworkApplianceRfProfileRequest**](UpdateNetworkApplianceRfProfileRequest.md) |  | 

### Return type

[**GetNetworkApplianceRfProfiles200ResponseAssignedInner**](GetNetworkApplianceRfProfiles200ResponseAssignedInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSdwanInternetPolicies

> UpdateNetworkApplianceSdwanInternetPolicies200Response UpdateNetworkApplianceSdwanInternetPolicies(ctx, networkId).UpdateNetworkApplianceSdwanInternetPoliciesRequest(updateNetworkApplianceSdwanInternetPoliciesRequest).Execute()

Update SDWAN internet traffic preferences for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceSdwanInternetPoliciesRequest := *openapiclient.NewUpdateNetworkApplianceSdwanInternetPoliciesRequest() // UpdateNetworkApplianceSdwanInternetPoliciesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceSdwanInternetPolicies(context.Background(), networkId).UpdateNetworkApplianceSdwanInternetPoliciesRequest(updateNetworkApplianceSdwanInternetPoliciesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceSdwanInternetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSdwanInternetPolicies`: UpdateNetworkApplianceSdwanInternetPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceSdwanInternetPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSdwanInternetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSdwanInternetPoliciesRequest** | [**UpdateNetworkApplianceSdwanInternetPoliciesRequest**](UpdateNetworkApplianceSdwanInternetPoliciesRequest.md) |  | 

### Return type

[**UpdateNetworkApplianceSdwanInternetPolicies200Response**](UpdateNetworkApplianceSdwanInternetPolicies200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityIntrusion

> GetNetworkApplianceSecurityIntrusion200Response UpdateNetworkApplianceSecurityIntrusion(ctx, networkId).UpdateNetworkApplianceSecurityIntrusionRequest(updateNetworkApplianceSecurityIntrusionRequest).Execute()

Set the supported intrusion settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceSecurityIntrusionRequest := *openapiclient.NewUpdateNetworkApplianceSecurityIntrusionRequest() // UpdateNetworkApplianceSecurityIntrusionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceSecurityIntrusion(context.Background(), networkId).UpdateNetworkApplianceSecurityIntrusionRequest(updateNetworkApplianceSecurityIntrusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSecurityIntrusion`: GetNetworkApplianceSecurityIntrusion200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityIntrusionRequest** | [**UpdateNetworkApplianceSecurityIntrusionRequest**](UpdateNetworkApplianceSecurityIntrusionRequest.md) |  | 

### Return type

[**GetNetworkApplianceSecurityIntrusion200Response**](GetNetworkApplianceSecurityIntrusion200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSecurityMalware

> GetNetworkApplianceSecurityMalware200Response UpdateNetworkApplianceSecurityMalware(ctx, networkId).UpdateNetworkApplianceSecurityMalwareRequest(updateNetworkApplianceSecurityMalwareRequest).Execute()

Set the supported malware settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceSecurityMalwareRequest := *openapiclient.NewUpdateNetworkApplianceSecurityMalwareRequest("Mode_example") // UpdateNetworkApplianceSecurityMalwareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceSecurityMalware(context.Background(), networkId).UpdateNetworkApplianceSecurityMalwareRequest(updateNetworkApplianceSecurityMalwareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceSecurityMalware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSecurityMalware`: GetNetworkApplianceSecurityMalware200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceSecurityMalware`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSecurityMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSecurityMalwareRequest** | [**UpdateNetworkApplianceSecurityMalwareRequest**](UpdateNetworkApplianceSecurityMalwareRequest.md) |  | 

### Return type

[**GetNetworkApplianceSecurityMalware200Response**](GetNetworkApplianceSecurityMalware200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSettings

> GetNetworkApplianceSettings200Response UpdateNetworkApplianceSettings(ctx, networkId).UpdateNetworkApplianceSettingsRequest(updateNetworkApplianceSettingsRequest).Execute()

Update the appliance settings for a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceSettingsRequest := *openapiclient.NewUpdateNetworkApplianceSettingsRequest() // UpdateNetworkApplianceSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceSettings(context.Background(), networkId).UpdateNetworkApplianceSettingsRequest(updateNetworkApplianceSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSettings`: GetNetworkApplianceSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSettingsRequest** | [**UpdateNetworkApplianceSettingsRequest**](UpdateNetworkApplianceSettingsRequest.md) |  | 

### Return type

[**GetNetworkApplianceSettings200Response**](GetNetworkApplianceSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSingleLan

> GetNetworkApplianceSingleLan200Response UpdateNetworkApplianceSingleLan(ctx, networkId).UpdateNetworkApplianceSingleLanRequest(updateNetworkApplianceSingleLanRequest).Execute()

Update single LAN configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceSingleLanRequest := *openapiclient.NewUpdateNetworkApplianceSingleLanRequest() // UpdateNetworkApplianceSingleLanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceSingleLan(context.Background(), networkId).UpdateNetworkApplianceSingleLanRequest(updateNetworkApplianceSingleLanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceSingleLan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSingleLan`: GetNetworkApplianceSingleLan200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceSingleLan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSingleLanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceSingleLanRequest** | [**UpdateNetworkApplianceSingleLanRequest**](UpdateNetworkApplianceSingleLanRequest.md) |  | 

### Return type

[**GetNetworkApplianceSingleLan200Response**](GetNetworkApplianceSingleLan200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceSsid

> GetNetworkApplianceSsids200ResponseInner UpdateNetworkApplianceSsid(ctx, networkId, number).UpdateNetworkApplianceSsidRequest(updateNetworkApplianceSsidRequest).Execute()

Update the attributes of an MX SSID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	number := "number_example" // string | Number
	updateNetworkApplianceSsidRequest := *openapiclient.NewUpdateNetworkApplianceSsidRequest() // UpdateNetworkApplianceSsidRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceSsid(context.Background(), networkId, number).UpdateNetworkApplianceSsidRequest(updateNetworkApplianceSsidRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceSsid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceSsid`: GetNetworkApplianceSsids200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceSsid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**number** | **string** | Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceSsidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceSsidRequest** | [**UpdateNetworkApplianceSsidRequest**](UpdateNetworkApplianceSsidRequest.md) |  | 

### Return type

[**GetNetworkApplianceSsids200ResponseInner**](GetNetworkApplianceSsids200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceStaticRoute

> GetNetworkApplianceStaticRoutes200ResponseInner UpdateNetworkApplianceStaticRoute(ctx, networkId, staticRouteId).UpdateNetworkApplianceStaticRouteRequest(updateNetworkApplianceStaticRouteRequest).Execute()

Update a static route for an MX or teleworker network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	staticRouteId := "staticRouteId_example" // string | Static route ID
	updateNetworkApplianceStaticRouteRequest := *openapiclient.NewUpdateNetworkApplianceStaticRouteRequest() // UpdateNetworkApplianceStaticRouteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceStaticRoute(context.Background(), networkId, staticRouteId).UpdateNetworkApplianceStaticRouteRequest(updateNetworkApplianceStaticRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceStaticRoute`: GetNetworkApplianceStaticRoutes200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**staticRouteId** | **string** | Static route ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceStaticRouteRequest** | [**UpdateNetworkApplianceStaticRouteRequest**](UpdateNetworkApplianceStaticRouteRequest.md) |  | 

### Return type

[**GetNetworkApplianceStaticRoutes200ResponseInner**](GetNetworkApplianceStaticRoutes200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShaping

> map[string]interface{} UpdateNetworkApplianceTrafficShaping(ctx, networkId).UpdateNetworkApplianceTrafficShapingRequest(updateNetworkApplianceTrafficShapingRequest).Execute()

Update the traffic shaping settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceTrafficShapingRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingRequest() // UpdateNetworkApplianceTrafficShapingRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceTrafficShaping(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRequest(updateNetworkApplianceTrafficShapingRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceTrafficShaping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShaping`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceTrafficShaping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRequest** | [**UpdateNetworkApplianceTrafficShapingRequest**](UpdateNetworkApplianceTrafficShapingRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingCustomPerformanceClass

> GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(ctx, networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()

Update a custom performance class for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	customPerformanceClassId := "customPerformanceClassId_example" // string | Custom performance class ID
	updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest() // UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass(context.Background(), networkId, customPerformanceClassId).UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest(updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceTrafficShapingCustomPerformanceClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**customPerformanceClassId** | **string** | Custom performance class ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceTrafficShapingCustomPerformanceClassRequest** | [**UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest**](UpdateNetworkApplianceTrafficShapingCustomPerformanceClassRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner**](GetNetworkApplianceTrafficShapingCustomPerformanceClasses200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingRules

> map[string]interface{} UpdateNetworkApplianceTrafficShapingRules(ctx, networkId).UpdateNetworkApplianceTrafficShapingRulesRequest(updateNetworkApplianceTrafficShapingRulesRequest).Execute()

Update the traffic shaping settings rules for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceTrafficShapingRulesRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingRulesRequest() // UpdateNetworkApplianceTrafficShapingRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceTrafficShapingRules(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingRulesRequest(updateNetworkApplianceTrafficShapingRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceTrafficShapingRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingRules`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceTrafficShapingRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingRulesRequest** | [**UpdateNetworkApplianceTrafficShapingRulesRequest**](UpdateNetworkApplianceTrafficShapingRulesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingUplinkBandwidth

> map[string]interface{} UpdateNetworkApplianceTrafficShapingUplinkBandwidth(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest(updateNetworkApplianceTrafficShapingUplinkBandwidthRequest).Execute()

Updates the uplink bandwidth settings for your MX network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceTrafficShapingUplinkBandwidthRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest() // UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceTrafficShapingUplinkBandwidth(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest(updateNetworkApplianceTrafficShapingUplinkBandwidthRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceTrafficShapingUplinkBandwidth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingUplinkBandwidth`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceTrafficShapingUplinkBandwidth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkBandwidthRequest** | [**UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest**](UpdateNetworkApplianceTrafficShapingUplinkBandwidthRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingUplinkSelection

> GetNetworkApplianceTrafficShapingUplinkSelection200Response UpdateNetworkApplianceTrafficShapingUplinkSelection(ctx, networkId).UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(updateNetworkApplianceTrafficShapingUplinkSelectionRequest).Execute()

Update uplink selection settings for an MX network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceTrafficShapingUplinkSelectionRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest() // UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceTrafficShapingUplinkSelection(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest(updateNetworkApplianceTrafficShapingUplinkSelectionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceTrafficShapingUplinkSelection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingUplinkSelection`: GetNetworkApplianceTrafficShapingUplinkSelection200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceTrafficShapingUplinkSelection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingUplinkSelectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingUplinkSelectionRequest** | [**UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest**](UpdateNetworkApplianceTrafficShapingUplinkSelectionRequest.md) |  | 

### Return type

[**GetNetworkApplianceTrafficShapingUplinkSelection200Response**](GetNetworkApplianceTrafficShapingUplinkSelection200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceTrafficShapingVpnExclusions

> UpdateNetworkApplianceTrafficShapingVpnExclusions200Response UpdateNetworkApplianceTrafficShapingVpnExclusions(ctx, networkId).UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest(updateNetworkApplianceTrafficShapingVpnExclusionsRequest).Execute()

Update VPN exclusion rules for an MX network.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceTrafficShapingVpnExclusionsRequest := *openapiclient.NewUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest() // UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceTrafficShapingVpnExclusions(context.Background(), networkId).UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest(updateNetworkApplianceTrafficShapingVpnExclusionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceTrafficShapingVpnExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceTrafficShapingVpnExclusions`: UpdateNetworkApplianceTrafficShapingVpnExclusions200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceTrafficShapingVpnExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceTrafficShapingVpnExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceTrafficShapingVpnExclusionsRequest** | [**UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest**](UpdateNetworkApplianceTrafficShapingVpnExclusionsRequest.md) |  | 

### Return type

[**UpdateNetworkApplianceTrafficShapingVpnExclusions200Response**](UpdateNetworkApplianceTrafficShapingVpnExclusions200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlan

> GetNetworkApplianceVlans200ResponseInner UpdateNetworkApplianceVlan(ctx, networkId, vlanId).UpdateNetworkApplianceVlanRequest(updateNetworkApplianceVlanRequest).Execute()

Update a VLAN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	vlanId := "vlanId_example" // string | Vlan ID
	updateNetworkApplianceVlanRequest := *openapiclient.NewUpdateNetworkApplianceVlanRequest() // UpdateNetworkApplianceVlanRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceVlan(context.Background(), networkId, vlanId).UpdateNetworkApplianceVlanRequest(updateNetworkApplianceVlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceVlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVlan`: GetNetworkApplianceVlans200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceVlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 
**vlanId** | **string** | Vlan ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateNetworkApplianceVlanRequest** | [**UpdateNetworkApplianceVlanRequest**](UpdateNetworkApplianceVlanRequest.md) |  | 

### Return type

[**GetNetworkApplianceVlans200ResponseInner**](GetNetworkApplianceVlans200ResponseInner.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVlansSettings

> GetNetworkApplianceVlansSettings200Response UpdateNetworkApplianceVlansSettings(ctx, networkId).UpdateNetworkApplianceVlansSettingsRequest(updateNetworkApplianceVlansSettingsRequest).Execute()

Enable/Disable VLANs for the given network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceVlansSettingsRequest := *openapiclient.NewUpdateNetworkApplianceVlansSettingsRequest() // UpdateNetworkApplianceVlansSettingsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceVlansSettings(context.Background(), networkId).UpdateNetworkApplianceVlansSettingsRequest(updateNetworkApplianceVlansSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceVlansSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVlansSettings`: GetNetworkApplianceVlansSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceVlansSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVlansSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVlansSettingsRequest** | [**UpdateNetworkApplianceVlansSettingsRequest**](UpdateNetworkApplianceVlansSettingsRequest.md) |  | 

### Return type

[**GetNetworkApplianceVlansSettings200Response**](GetNetworkApplianceVlansSettings200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnBgp

> GetNetworkApplianceVpnBgp200Response UpdateNetworkApplianceVpnBgp(ctx, networkId).UpdateNetworkApplianceVpnBgpRequest(updateNetworkApplianceVpnBgpRequest).Execute()

Update a Hub BGP Configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceVpnBgpRequest := *openapiclient.NewUpdateNetworkApplianceVpnBgpRequest(false) // UpdateNetworkApplianceVpnBgpRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceVpnBgp(context.Background(), networkId).UpdateNetworkApplianceVpnBgpRequest(updateNetworkApplianceVpnBgpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceVpnBgp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVpnBgp`: GetNetworkApplianceVpnBgp200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceVpnBgp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnBgpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnBgpRequest** | [**UpdateNetworkApplianceVpnBgpRequest**](UpdateNetworkApplianceVpnBgpRequest.md) |  | 

### Return type

[**GetNetworkApplianceVpnBgp200Response**](GetNetworkApplianceVpnBgp200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceVpnSiteToSiteVpn

> GetNetworkApplianceVpnSiteToSiteVpn200Response UpdateNetworkApplianceVpnSiteToSiteVpn(ctx, networkId).UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest).Execute()

Update the site-to-site VPN settings of a network



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceVpnSiteToSiteVpnRequest := *openapiclient.NewUpdateNetworkApplianceVpnSiteToSiteVpnRequest("Mode_example") // UpdateNetworkApplianceVpnSiteToSiteVpnRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceVpnSiteToSiteVpn(context.Background(), networkId).UpdateNetworkApplianceVpnSiteToSiteVpnRequest(updateNetworkApplianceVpnSiteToSiteVpnRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceVpnSiteToSiteVpn``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceVpnSiteToSiteVpn`: GetNetworkApplianceVpnSiteToSiteVpn200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceVpnSiteToSiteVpn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceVpnSiteToSiteVpnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceVpnSiteToSiteVpnRequest** | [**UpdateNetworkApplianceVpnSiteToSiteVpnRequest**](UpdateNetworkApplianceVpnSiteToSiteVpnRequest.md) |  | 

### Return type

[**GetNetworkApplianceVpnSiteToSiteVpn200Response**](GetNetworkApplianceVpnSiteToSiteVpn200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkApplianceWarmSpare

> GetNetworkApplianceWarmSpare200Response UpdateNetworkApplianceWarmSpare(ctx, networkId).UpdateNetworkApplianceWarmSpareRequest(updateNetworkApplianceWarmSpareRequest).Execute()

Update MX warm spare settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	networkId := "networkId_example" // string | Network ID
	updateNetworkApplianceWarmSpareRequest := *openapiclient.NewUpdateNetworkApplianceWarmSpareRequest(false) // UpdateNetworkApplianceWarmSpareRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateNetworkApplianceWarmSpare(context.Background(), networkId).UpdateNetworkApplianceWarmSpareRequest(updateNetworkApplianceWarmSpareRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateNetworkApplianceWarmSpare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkApplianceWarmSpare`: GetNetworkApplianceWarmSpare200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateNetworkApplianceWarmSpare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** | Network ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkApplianceWarmSpareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkApplianceWarmSpareRequest** | [**UpdateNetworkApplianceWarmSpareRequest**](UpdateNetworkApplianceWarmSpareRequest.md) |  | 

### Return type

[**GetNetworkApplianceWarmSpare200Response**](GetNetworkApplianceWarmSpare200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceSecurityIntrusion

> map[string]interface{} UpdateOrganizationApplianceSecurityIntrusion(ctx, organizationId).UpdateOrganizationApplianceSecurityIntrusionRequest(updateOrganizationApplianceSecurityIntrusionRequest).Execute()

Sets supported intrusion settings for an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	updateOrganizationApplianceSecurityIntrusionRequest := *openapiclient.NewUpdateOrganizationApplianceSecurityIntrusionRequest([]openapiclient.UpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner{*openapiclient.NewUpdateOrganizationApplianceSecurityIntrusionRequestAllowedRulesInner("RuleId_example")}) // UpdateOrganizationApplianceSecurityIntrusionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateOrganizationApplianceSecurityIntrusion(context.Background(), organizationId).UpdateOrganizationApplianceSecurityIntrusionRequest(updateOrganizationApplianceSecurityIntrusionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateOrganizationApplianceSecurityIntrusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationApplianceSecurityIntrusion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateOrganizationApplianceSecurityIntrusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceSecurityIntrusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceSecurityIntrusionRequest** | [**UpdateOrganizationApplianceSecurityIntrusionRequest**](UpdateOrganizationApplianceSecurityIntrusionRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnThirdPartyVPNPeers

> GetOrganizationApplianceVpnThirdPartyVPNPeers200Response UpdateOrganizationApplianceVpnThirdPartyVPNPeers(ctx, organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(updateOrganizationApplianceVpnThirdPartyVPNPeersRequest).Execute()

Update the third party VPN peers for an organization



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	updateOrganizationApplianceVpnThirdPartyVPNPeersRequest := *openapiclient.NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest([]openapiclient.UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner{*openapiclient.NewUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequestPeersInner("Name_example", []string{"PrivateSubnets_example"}, "Secret_example")}) // UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers(context.Background(), organizationId).UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest(updateOrganizationApplianceVpnThirdPartyVPNPeersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: GetOrganizationApplianceVpnThirdPartyVPNPeers200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateOrganizationApplianceVpnThirdPartyVPNPeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnThirdPartyVPNPeersRequest** | [**UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest**](UpdateOrganizationApplianceVpnThirdPartyVPNPeersRequest.md) |  | 

### Return type

[**GetOrganizationApplianceVpnThirdPartyVPNPeers200Response**](GetOrganizationApplianceVpnThirdPartyVPNPeers200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrganizationApplianceVpnVpnFirewallRules

> GetNetworkApplianceFirewallInboundCellularFirewallRules200Response UpdateOrganizationApplianceVpnVpnFirewallRules(ctx, organizationId).UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest).Execute()

Update the firewall rules of an organization's site-to-site VPN



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	organizationId := "organizationId_example" // string | Organization ID
	updateOrganizationApplianceVpnVpnFirewallRulesRequest := *openapiclient.NewUpdateOrganizationApplianceVpnVpnFirewallRulesRequest() // UpdateOrganizationApplianceVpnVpnFirewallRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplianceAPI.UpdateOrganizationApplianceVpnVpnFirewallRules(context.Background(), organizationId).UpdateOrganizationApplianceVpnVpnFirewallRulesRequest(updateOrganizationApplianceVpnVpnFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplianceAPI.UpdateOrganizationApplianceVpnVpnFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrganizationApplianceVpnVpnFirewallRules`: GetNetworkApplianceFirewallInboundCellularFirewallRules200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplianceAPI.UpdateOrganizationApplianceVpnVpnFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Organization ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrganizationApplianceVpnVpnFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateOrganizationApplianceVpnVpnFirewallRulesRequest** | [**UpdateOrganizationApplianceVpnVpnFirewallRulesRequest**](UpdateOrganizationApplianceVpnVpnFirewallRulesRequest.md) |  | 

### Return type

[**GetNetworkApplianceFirewallInboundCellularFirewallRules200Response**](GetNetworkApplianceFirewallInboundCellularFirewallRules200Response.md)

### Authorization

[meraki_api_key](../README.md#meraki_api_key), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

