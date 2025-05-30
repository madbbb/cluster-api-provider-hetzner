package schema

import "time"

// Server defines the schema of a server.
type Server struct {
	ID              int64              `json:"id"`
	Name            string             `json:"name"`
	Status          string             `json:"status"`
	Created         time.Time          `json:"created"`
	PublicNet       ServerPublicNet    `json:"public_net"`
	PrivateNet      []ServerPrivateNet `json:"private_net"`
	ServerType      ServerType         `json:"server_type"`
	IncludedTraffic uint64             `json:"included_traffic"`
	OutgoingTraffic *uint64            `json:"outgoing_traffic"`
	IngoingTraffic  *uint64            `json:"ingoing_traffic"`
	BackupWindow    *string            `json:"backup_window"`
	RescueEnabled   bool               `json:"rescue_enabled"`
	ISO             *ISO               `json:"iso"`
	Locked          bool               `json:"locked"`
	Datacenter      Datacenter         `json:"datacenter"`
	Image           *Image             `json:"image"`
	Protection      ServerProtection   `json:"protection"`
	Labels          map[string]string  `json:"labels"`
	Volumes         []int64            `json:"volumes"`
	PrimaryDiskSize int                `json:"primary_disk_size"`
	PlacementGroup  *PlacementGroup    `json:"placement_group"`
	LoadBalancers   []int64            `json:"load_balancers"`
}

// ServerProtection defines the schema of a server's resource protection.
type ServerProtection struct {
	Delete  bool `json:"delete"`
	Rebuild bool `json:"rebuild"`
}

// ServerPublicNet defines the schema of a server's
// public network information.
type ServerPublicNet struct {
	IPv4        ServerPublicNetIPv4 `json:"ipv4"`
	IPv6        ServerPublicNetIPv6 `json:"ipv6"`
	FloatingIPs []int64             `json:"floating_ips"`
	Firewalls   []ServerFirewall    `json:"firewalls"`
}

// ServerPublicNetIPv4 defines the schema of a server's public
// network information for an IPv4.
type ServerPublicNetIPv4 struct {
	ID      int64  `json:"id"`
	IP      string `json:"ip"`
	Blocked bool   `json:"blocked"`
	DNSPtr  string `json:"dns_ptr"`
}

// ServerPublicNetIPv6 defines the schema of a server's public
// network information for an IPv6.
type ServerPublicNetIPv6 struct {
	ID      int64                       `json:"id"`
	IP      string                      `json:"ip"`
	Blocked bool                        `json:"blocked"`
	DNSPtr  []ServerPublicNetIPv6DNSPtr `json:"dns_ptr"`
}

// ServerPublicNetIPv6DNSPtr defines the schema of a server's
// public network information for an IPv6 reverse DNS.
type ServerPublicNetIPv6DNSPtr struct {
	IP     string `json:"ip"`
	DNSPtr string `json:"dns_ptr"`
}

// ServerFirewall defines the schema of a Server's Firewalls on
// a certain network interface.
type ServerFirewall struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

// ServerPrivateNet defines the schema of a server's private network information.
type ServerPrivateNet struct {
	Network    int64    `json:"network"`
	IP         string   `json:"ip"`
	AliasIPs   []string `json:"alias_ips"`
	MACAddress string   `json:"mac_address"`
}

// ServerGetResponse defines the schema of the response when
// retrieving a single server.
type ServerGetResponse struct {
	Server Server `json:"server"`
}

// ServerListResponse defines the schema of the response when
// listing servers.
type ServerListResponse struct {
	Servers []Server `json:"servers"`
}

// ServerCreateRequest defines the schema for the request to
// create a server.
type ServerCreateRequest struct {
	Name             string                  `json:"name"`
	ServerType       IDOrName                `json:"server_type"`
	Image            IDOrName                `json:"image"`
	SSHKeys          []int64                 `json:"ssh_keys,omitempty"`
	Location         string                  `json:"location,omitempty"`
	Datacenter       string                  `json:"datacenter,omitempty"`
	UserData         string                  `json:"user_data,omitempty"`
	StartAfterCreate *bool                   `json:"start_after_create,omitempty"`
	Labels           *map[string]string      `json:"labels,omitempty"`
	Automount        *bool                   `json:"automount,omitempty"`
	Volumes          []int64                 `json:"volumes,omitempty"`
	Networks         []int64                 `json:"networks,omitempty"`
	Firewalls        []ServerCreateFirewalls `json:"firewalls,omitempty"`
	PlacementGroup   int64                   `json:"placement_group,omitempty"`
	PublicNet        *ServerCreatePublicNet  `json:"public_net,omitempty"`
}

// ServerCreatePublicNet defines the public network configuration of a server.
type ServerCreatePublicNet struct {
	EnableIPv4 bool  `json:"enable_ipv4"`
	EnableIPv6 bool  `json:"enable_ipv6"`
	IPv4ID     int64 `json:"ipv4,omitempty"`
	IPv6ID     int64 `json:"ipv6,omitempty"`
}

// ServerCreateFirewalls defines which Firewalls to apply when creating a Server.
type ServerCreateFirewalls struct {
	Firewall int64 `json:"firewall"`
}

// ServerCreateResponse defines the schema of the response when
// creating a server.
type ServerCreateResponse struct {
	Server       Server   `json:"server"`
	Action       Action   `json:"action"`
	RootPassword *string  `json:"root_password"`
	NextActions  []Action `json:"next_actions"`
}

// ServerDeleteResponse defines the schema of the response when
// deleting a server.
type ServerDeleteResponse struct {
	Action Action `json:"action"`
}

// ServerUpdateRequest defines the schema of the request to update a server.
type ServerUpdateRequest struct {
	Name   string             `json:"name,omitempty"`
	Labels *map[string]string `json:"labels,omitempty"`
}

// ServerUpdateResponse defines the schema of the response when updating a server.
type ServerUpdateResponse struct {
	Server Server `json:"server"`
}

// ServerActionPoweronRequest defines the schema for the request to
// create a poweron server action.
type ServerActionPoweronRequest struct{}

// ServerActionPoweronResponse defines the schema of the response when
// creating a poweron server action.
type ServerActionPoweronResponse struct {
	Action Action `json:"action"`
}

// ServerActionPoweroffRequest defines the schema for the request to
// create a poweroff server action.
type ServerActionPoweroffRequest struct{}

// ServerActionPoweroffResponse defines the schema of the response when
// creating a poweroff server action.
type ServerActionPoweroffResponse struct {
	Action Action `json:"action"`
}

// ServerActionRebootRequest defines the schema for the request to
// create a reboot server action.
type ServerActionRebootRequest struct{}

// ServerActionRebootResponse defines the schema of the response when
// creating a reboot server action.
type ServerActionRebootResponse struct {
	Action Action `json:"action"`
}

// ServerActionResetRequest defines the schema for the request to
// create a reset server action.
type ServerActionResetRequest struct{}

// ServerActionResetResponse defines the schema of the response when
// creating a reset server action.
type ServerActionResetResponse struct {
	Action Action `json:"action"`
}

// ServerActionShutdownRequest defines the schema for the request to
// create a shutdown server action.
type ServerActionShutdownRequest struct{}

// ServerActionShutdownResponse defines the schema of the response when
// creating a shutdown server action.
type ServerActionShutdownResponse struct {
	Action Action `json:"action"`
}

// ServerActionResetPasswordRequest defines the schema for the request to
// create a reset_password server action.
type ServerActionResetPasswordRequest struct{}

// ServerActionResetPasswordResponse defines the schema of the response when
// creating a reset_password server action.
type ServerActionResetPasswordResponse struct {
	Action       Action `json:"action"`
	RootPassword string `json:"root_password"`
}

// ServerActionCreateImageRequest defines the schema for the request to
// create a create_image server action.
type ServerActionCreateImageRequest struct {
	Type        *string            `json:"type"`
	Description *string            `json:"description"`
	Labels      *map[string]string `json:"labels,omitempty"`
}

// ServerActionCreateImageResponse defines the schema of the response when
// creating a create_image server action.
type ServerActionCreateImageResponse struct {
	Action Action `json:"action"`
	Image  Image  `json:"image"`
}

// ServerActionEnableRescueRequest defines the schema for the request to
// create a enable_rescue server action.
type ServerActionEnableRescueRequest struct {
	Type    *string `json:"type,omitempty"`
	SSHKeys []int64 `json:"ssh_keys,omitempty"`
}

// ServerActionEnableRescueResponse defines the schema of the response when
// creating a enable_rescue server action.
type ServerActionEnableRescueResponse struct {
	Action       Action `json:"action"`
	RootPassword string `json:"root_password"`
}

// ServerActionDisableRescueRequest defines the schema for the request to
// create a disable_rescue server action.
type ServerActionDisableRescueRequest struct{}

// ServerActionDisableRescueResponse defines the schema of the response when
// creating a disable_rescue server action.
type ServerActionDisableRescueResponse struct {
	Action Action `json:"action"`
}

// ServerActionRebuildRequest defines the schema for the request to
// rebuild a server.
type ServerActionRebuildRequest struct {
	Image IDOrName `json:"image"`
}

// ServerActionRebuildResponse defines the schema of the response when
// creating a rebuild server action.
type ServerActionRebuildResponse struct {
	Action       Action  `json:"action"`
	RootPassword *string `json:"root_password"`
}

// ServerActionAttachISORequest defines the schema for the request to
// attach an ISO to a server.
type ServerActionAttachISORequest struct {
	ISO IDOrName `json:"iso"`
}

// ServerActionAttachISOResponse defines the schema of the response when
// creating a attach_iso server action.
type ServerActionAttachISOResponse struct {
	Action Action `json:"action"`
}

// ServerActionDetachISORequest defines the schema for the request to
// detach an ISO from a server.
type ServerActionDetachISORequest struct{}

// ServerActionDetachISOResponse defines the schema of the response when
// creating a detach_iso server action.
type ServerActionDetachISOResponse struct {
	Action Action `json:"action"`
}

// ServerActionEnableBackupResponse defines the schema of the response when
// creating a enable_backup server action.
type ServerActionEnableBackupResponse struct {
	Action Action `json:"action"`
}

// ServerActionDisableBackupRequest defines the schema for the request to
// disable backup for a server.
type ServerActionDisableBackupRequest struct{}

// ServerActionDisableBackupResponse defines the schema of the response when
// creating a disable_backup server action.
type ServerActionDisableBackupResponse struct {
	Action Action `json:"action"`
}

// ServerActionChangeTypeRequest defines the schema for the request to
// change a server's type.
type ServerActionChangeTypeRequest struct {
	ServerType  IDOrName `json:"server_type"`
	UpgradeDisk bool     `json:"upgrade_disk"`
}

// ServerActionChangeTypeResponse defines the schema of the response when
// creating a change_type server action.
type ServerActionChangeTypeResponse struct {
	Action Action `json:"action"`
}

// ServerActionChangeDNSPtrRequest defines the schema for the request to
// change a server's reverse DNS pointer.
type ServerActionChangeDNSPtrRequest struct {
	IP     string  `json:"ip"`
	DNSPtr *string `json:"dns_ptr"`
}

// ServerActionChangeDNSPtrResponse defines the schema of the response when
// creating a change_dns_ptr server action.
type ServerActionChangeDNSPtrResponse struct {
	Action Action `json:"action"`
}

// ServerActionChangeProtectionRequest defines the schema of the request to
// change the resource protection of a server.
type ServerActionChangeProtectionRequest struct {
	Rebuild *bool `json:"rebuild,omitempty"`
	Delete  *bool `json:"delete,omitempty"`
}

// ServerActionChangeProtectionResponse defines the schema of the response when
// changing the resource protection of a server.
type ServerActionChangeProtectionResponse struct {
	Action Action `json:"action"`
}

// ServerActionRequestConsoleRequest defines the schema of the request to
// request a WebSocket VNC console.
type ServerActionRequestConsoleRequest struct{}

// ServerActionRequestConsoleResponse defines the schema of the response when
// requesting a WebSocket VNC console.
type ServerActionRequestConsoleResponse struct {
	Action   Action `json:"action"`
	WSSURL   string `json:"wss_url"`
	Password string `json:"password"`
}

// ServerActionAttachToNetworkRequest defines the schema for the request to
// attach a network to a server.
type ServerActionAttachToNetworkRequest struct {
	Network  int64     `json:"network"`
	IP       *string   `json:"ip,omitempty"`
	AliasIPs []*string `json:"alias_ips,omitempty"`
}

// ServerActionAttachToNetworkResponse defines the schema of the response when
// creating an attach_to_network server action.
type ServerActionAttachToNetworkResponse struct {
	Action Action `json:"action"`
}

// ServerActionDetachFromNetworkRequest defines the schema for the request to
// detach a network from a server.
type ServerActionDetachFromNetworkRequest struct {
	Network int64 `json:"network"`
}

// ServerActionDetachFromNetworkResponse defines the schema of the response when
// creating a detach_from_network server action.
type ServerActionDetachFromNetworkResponse struct {
	Action Action `json:"action"`
}

// ServerActionChangeAliasIPsRequest defines the schema for the request to
// change a server's alias IPs in a network.
type ServerActionChangeAliasIPsRequest struct {
	Network  int64    `json:"network"`
	AliasIPs []string `json:"alias_ips"`
}

// ServerActionChangeAliasIPsResponse defines the schema of the response when
// creating an change_alias_ips server action.
type ServerActionChangeAliasIPsResponse struct {
	Action Action `json:"action"`
}

// ServerGetMetricsResponse defines the schema of the response when requesting
// metrics for a server.
type ServerGetMetricsResponse struct {
	Metrics struct {
		Start      time.Time                       `json:"start"`
		End        time.Time                       `json:"end"`
		Step       float64                         `json:"step"`
		TimeSeries map[string]ServerTimeSeriesVals `json:"time_series"`
	} `json:"metrics"`
}

// ServerTimeSeriesVals contains the values for a Server time series.
type ServerTimeSeriesVals struct {
	Values []interface{} `json:"values"`
}

// ServerActionAddToPlacementGroupRequest defines the schema for the request to
// add a server to a placement group.
type ServerActionAddToPlacementGroupRequest struct {
	PlacementGroup int64 `json:"placement_group"`
}

// ServerActionAddToPlacementGroupResponse defines the schema of the response when
// creating an add_to_placement_group server action.
type ServerActionAddToPlacementGroupResponse struct {
	Action Action `json:"action"`
}

// ServerActionRemoveFromPlacementGroupResponse defines the schema of the response when
// creating a remove_from_placement_group server action.
type ServerActionRemoveFromPlacementGroupResponse struct {
	Action Action `json:"action"`
}
