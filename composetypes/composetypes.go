package composetypes

import composetypes "github.com/compose-spec/compose-go/v2/types"

type IncludeConfig = composetypes.IncludeConfig

// Config is a full compose file configuration and model
type Config struct {
	Filename   string                       `yaml:"-" json:"-"`
	Name       string                       `yaml:"name,omitempty" json:"name,omitempty"`
	Services   Services                     `yaml:"services" json:"services"`
	Networks   composetypes.Networks        `yaml:"networks,omitempty" json:"networks,omitempty"`
	Volumes    composetypes.Volumes         `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	Secrets    composetypes.Secrets         `yaml:"secrets,omitempty" json:"secrets,omitempty"`
	Configs    composetypes.Configs         `yaml:"configs,omitempty" json:"configs,omitempty"`
	Extensions composetypes.Extensions      `yaml:",inline" json:"-"`
	Include    []composetypes.IncludeConfig `yaml:"include,omitempty" json:"include,omitempty"`
}

// Services is a map of ServiceConfig
type Services map[string]ServiceConfig

// ServiceConfig is the configuration of one service
type ServiceConfig struct {
	Name     string   `yaml:"name,omitempty" json:"-"`
	Profiles []string `yaml:"profiles,omitempty" json:"profiles,omitempty"`

	Annotations  composetypes.Mapping        `yaml:"annotations,omitempty" json:"annotations,omitempty"`
	Attach       *bool                       `yaml:"attach,omitempty" json:"attach,omitempty"`
	Build        *composetypes.BuildConfig   `yaml:"build,omitempty" json:"build,omitempty"`
	Develop      *composetypes.DevelopConfig `yaml:"develop,omitempty" json:"develop,omitempty"`
	BlkioConfig  *composetypes.BlkioConfig   `yaml:"blkio_config,omitempty" json:"blkio_config,omitempty"`
	CapAdd       []string                    `yaml:"cap_add,omitempty" json:"cap_add,omitempty"`
	CapDrop      []string                    `yaml:"cap_drop,omitempty" json:"cap_drop,omitempty"`
	CgroupParent string                      `yaml:"cgroup_parent,omitempty" json:"cgroup_parent,omitempty"`
	Cgroup       string                      `yaml:"cgroup,omitempty" json:"cgroup,omitempty"`
	CPUCount     int64                       `yaml:"cpu_count,omitempty" json:"cpu_count,omitempty"`
	CPUPercent   float32                     `yaml:"cpu_percent,omitempty" json:"cpu_percent,omitempty"`
	CPUPeriod    int64                       `yaml:"cpu_period,omitempty" json:"cpu_period,omitempty"`
	CPUQuota     int64                       `yaml:"cpu_quota,omitempty" json:"cpu_quota,omitempty"`
	CPURTPeriod  int64                       `yaml:"cpu_rt_period,omitempty" json:"cpu_rt_period,omitempty"`
	CPURTRuntime int64                       `yaml:"cpu_rt_runtime,omitempty" json:"cpu_rt_runtime,omitempty"`
	CPUS         float32                     `yaml:"cpus,omitempty" json:"cpus,omitempty"`
	CPUSet       string                      `yaml:"cpuset,omitempty" json:"cpuset,omitempty"`
	CPUShares    int64                       `yaml:"cpu_shares,omitempty" json:"cpu_shares,omitempty"`

	// Command for the service containers.
	// If set, overrides COMMAND from the image.
	//
	// Set to `[]` or an empty string to clear the command from the image.
	Command composetypes.ShellCommand `yaml:"command,omitempty" json:"command"` // NOTE: we can NOT omitempty for JSON! see ShellCommand type for details.

	Configs           []composetypes.ServiceConfigObjConfig `yaml:"configs,omitempty" json:"configs,omitempty"`
	ContainerName     string                                `yaml:"container_name,omitempty" json:"container_name,omitempty"`
	CredentialSpec    *composetypes.CredentialSpecConfig    `yaml:"credential_spec,omitempty" json:"credential_spec,omitempty"`
	DependsOn         composetypes.DependsOnConfig          `yaml:"depends_on,omitempty" json:"depends_on,omitempty"`
	Deploy            *DeployConfig                         `yaml:"deploy,omitempty" json:"deploy,omitempty"`
	DeviceCgroupRules []string                              `yaml:"device_cgroup_rules,omitempty" json:"device_cgroup_rules,omitempty"`
	Devices           []composetypes.DeviceMapping          `yaml:"devices,omitempty" json:"devices,omitempty"`
	DNS               composetypes.StringList               `yaml:"dns,omitempty" json:"dns,omitempty"`
	DNSOpts           []string                              `yaml:"dns_opt,omitempty" json:"dns_opt,omitempty"`
	DNSSearch         composetypes.StringList               `yaml:"dns_search,omitempty" json:"dns_search,omitempty"`
	Dockerfile        string                                `yaml:"dockerfile,omitempty" json:"dockerfile,omitempty"`
	DomainName        string                                `yaml:"domainname,omitempty" json:"domainname,omitempty"`

	// Entrypoint for the service containers.
	// If set, overrides ENTRYPOINT from the image.
	//
	// Set to `[]` or an empty string to clear the entrypoint from the image.
	Entrypoint      composetypes.ShellCommand                     `yaml:"entrypoint,omitempty" json:"entrypoint"` // NOTE: we can NOT omitempty for JSON! see ShellCommand type for details.
	Provider        *composetypes.ServiceProviderConfig           `yaml:"provider,omitempty" json:"provider,omitempty"`
	Environment     composetypes.MappingWithEquals                `yaml:"environment,omitempty" json:"environment,omitempty"`
	EnvFiles        []composetypes.EnvFile                        `yaml:"env_file,omitempty" json:"env_file,omitempty"`
	Expose          composetypes.StringOrNumberList               `yaml:"expose,omitempty" json:"expose,omitempty"`
	Extends         *composetypes.ExtendsConfig                   `yaml:"extends,omitempty" json:"extends,omitempty"`
	ExternalLinks   []string                                      `yaml:"external_links,omitempty" json:"external_links,omitempty"`
	ExtraHosts      composetypes.HostsList                        `yaml:"extra_hosts,omitempty" json:"extra_hosts,omitempty"`
	GroupAdd        []string                                      `yaml:"group_add,omitempty" json:"group_add,omitempty"`
	Gpus            []composetypes.DeviceRequest                  `yaml:"gpus,omitempty" json:"gpus,omitempty"`
	Hostname        string                                        `yaml:"hostname,omitempty" json:"hostname,omitempty"`
	HealthCheck     *HealthCheckConfig                            `yaml:"healthcheck,omitempty" json:"healthcheck,omitempty"`
	Image           string                                        `yaml:"image,omitempty" json:"image,omitempty"`
	Init            *bool                                         `yaml:"init,omitempty" json:"init,omitempty"`
	Ipc             string                                        `yaml:"ipc,omitempty" json:"ipc,omitempty"`
	Isolation       string                                        `yaml:"isolation,omitempty" json:"isolation,omitempty"`
	Labels          composetypes.Labels                           `yaml:"labels,omitempty" json:"labels,omitempty"`
	LabelFiles      []string                                      `yaml:"label_file,omitempty" json:"label_file,omitempty"`
	CustomLabels    composetypes.Labels                           `yaml:"-" json:"-"`
	Links           []string                                      `yaml:"links,omitempty" json:"links,omitempty"`
	Logging         *composetypes.LoggingConfig                   `yaml:"logging,omitempty" json:"logging,omitempty"`
	LogDriver       string                                        `yaml:"log_driver,omitempty" json:"log_driver,omitempty"`
	LogOpt          map[string]string                             `yaml:"log_opt,omitempty" json:"log_opt,omitempty"`
	MemLimit        composetypes.UnitBytes                        `yaml:"mem_limit,omitempty" json:"mem_limit,omitempty"`
	MemReservation  composetypes.UnitBytes                        `yaml:"mem_reservation,omitempty" json:"mem_reservation,omitempty"`
	MemSwapLimit    composetypes.UnitBytes                        `yaml:"memswap_limit,omitempty" json:"memswap_limit,omitempty"`
	MemSwappiness   composetypes.UnitBytes                        `yaml:"mem_swappiness,omitempty" json:"mem_swappiness,omitempty"`
	MacAddress      string                                        `yaml:"mac_address,omitempty" json:"mac_address,omitempty"`
	Net             string                                        `yaml:"net,omitempty" json:"net,omitempty"`
	NetworkMode     string                                        `yaml:"network_mode,omitempty" json:"network_mode,omitempty"`
	Networks        map[string]*composetypes.ServiceNetworkConfig `yaml:"networks,omitempty" json:"networks,omitempty"`
	OomKillDisable  bool                                          `yaml:"oom_kill_disable,omitempty" json:"oom_kill_disable,omitempty"`
	OomScoreAdj     int64                                         `yaml:"oom_score_adj,omitempty" json:"oom_score_adj,omitempty"`
	Pid             string                                        `yaml:"pid,omitempty" json:"pid,omitempty"`
	PidsLimit       int64                                         `yaml:"pids_limit,omitempty" json:"pids_limit,omitempty"`
	Platform        string                                        `yaml:"platform,omitempty" json:"platform,omitempty"`
	Ports           []composetypes.ServicePortConfig              `yaml:"ports,omitempty" json:"ports,omitempty"`
	Privileged      bool                                          `yaml:"privileged,omitempty" json:"privileged,omitempty"`
	PullPolicy      string                                        `yaml:"pull_policy,omitempty" json:"pull_policy,omitempty"`
	ReadOnly        bool                                          `yaml:"read_only,omitempty" json:"read_only,omitempty"`
	Restart         string                                        `yaml:"restart,omitempty" json:"restart,omitempty"`
	Runtime         string                                        `yaml:"runtime,omitempty" json:"runtime,omitempty"`
	Scale           *int                                          `yaml:"scale,omitempty" json:"scale,omitempty"`
	Secrets         []composetypes.ServiceSecretConfig            `yaml:"secrets,omitempty" json:"secrets,omitempty"`
	SecurityOpt     []string                                      `yaml:"security_opt,omitempty" json:"security_opt,omitempty"`
	ShmSize         composetypes.UnitBytes                        `yaml:"shm_size,omitempty" json:"shm_size,omitempty"`
	StdinOpen       bool                                          `yaml:"stdin_open,omitempty" json:"stdin_open,omitempty"`
	StopGracePeriod Duration                                      `yaml:"stop_grace_period,omitempty" json:"stop_grace_period,omitempty"`
	StopSignal      string                                        `yaml:"stop_signal,omitempty" json:"stop_signal,omitempty"`
	StorageOpt      map[string]string                             `yaml:"storage_opt,omitempty" json:"storage_opt,omitempty"`
	Sysctls         composetypes.Mapping                          `yaml:"sysctls,omitempty" json:"sysctls,omitempty"`
	Tmpfs           composetypes.StringList                       `yaml:"tmpfs,omitempty" json:"tmpfs,omitempty"`
	Tty             bool                                          `yaml:"tty,omitempty" json:"tty,omitempty"`
	Ulimits         map[string]*composetypes.UlimitsConfig        `yaml:"ulimits,omitempty" json:"ulimits,omitempty"`
	User            string                                        `yaml:"user,omitempty" json:"user,omitempty"`
	UserNSMode      string                                        `yaml:"userns_mode,omitempty" json:"userns_mode,omitempty"`
	Uts             string                                        `yaml:"uts,omitempty" json:"uts,omitempty"`
	VolumeDriver    string                                        `yaml:"volume_driver,omitempty" json:"volume_driver,omitempty"`
	Volumes         []composetypes.ServiceVolumeConfig            `yaml:"volumes,omitempty" json:"volumes,omitempty"`
	VolumesFrom     []string                                      `yaml:"volumes_from,omitempty" json:"volumes_from,omitempty"`
	WorkingDir      string                                        `yaml:"working_dir,omitempty" json:"working_dir,omitempty"`
	PostStart       []composetypes.ServiceHook                    `yaml:"post_start,omitempty" json:"post_start,omitempty"`
	PreStop         []composetypes.ServiceHook                    `yaml:"pre_stop,omitempty" json:"pre_stop,omitempty"`

	Extensions composetypes.Extensions `yaml:"#extensions,inline,omitempty" json:"-"`
}

type Duration string

// HealthCheckConfig the healthcheck configuration for a service
type HealthCheckConfig struct {
	Test          composetypes.HealthCheckTest `yaml:"test,omitempty" json:"test,omitempty"`
	Timeout       Duration                     `yaml:"timeout,omitempty" json:"timeout,omitempty"`
	Interval      Duration                     `yaml:"interval,omitempty" json:"interval,omitempty"`
	Retries       *uint64                      `yaml:"retries,omitempty" json:"retries,omitempty"`
	StartPeriod   Duration                     `yaml:"start_period,omitempty" json:"start_period,omitempty"`
	StartInterval Duration                     `yaml:"start_interval,omitempty" json:"start_interval,omitempty"`
	Disable       bool                         `yaml:"disable,omitempty" json:"disable,omitempty"`

	Extensions composetypes.Extensions `yaml:"#extensions,inline,omitempty" json:"-"`
}

// DeployConfig the deployment configuration for a service
type DeployConfig struct {
	Mode           string                      `yaml:"mode,omitempty" json:"mode,omitempty"`
	Replicas       *int                        `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	Labels         composetypes.Labels         `yaml:"labels,omitempty" json:"labels,omitempty"`
	UpdateConfig   *composetypes.UpdateConfig  `yaml:"update_config,omitempty" json:"update_config,omitempty"`
	RollbackConfig *composetypes.UpdateConfig  `yaml:"rollback_config,omitempty" json:"rollback_config,omitempty"`
	Resources      *Resources                  `yaml:"resources,omitempty" json:"resources,omitempty"`
	RestartPolicy  *composetypes.RestartPolicy `yaml:"restart_policy,omitempty" json:"restart_policy,omitempty"`
	Placement      *composetypes.Placement     `yaml:"placement,omitempty" json:"placement,omitempty"`
	EndpointMode   string                      `yaml:"endpoint_mode,omitempty" json:"endpoint_mode,omitempty"`

	Extensions composetypes.Extensions `yaml:"#extensions,inline,omitempty" json:"-"`
}

// Resources the resource limits and reservations
type Resources struct {
	Limits       *Resource `yaml:"limits,omitempty" json:"limits,omitempty"`
	Reservations *Resource `yaml:"reservations,omitempty" json:"reservations,omitempty"`

	Extensions composetypes.Extensions `yaml:"#extensions,inline,omitempty" json:"-"`
}

// Resource is a resource to be limited or reserved
type Resource struct {
	// TODO: types to convert from units and ratios
	NanoCPUs         float64                        `yaml:"cpus,omitempty" json:"cpus,omitempty"`
	MemoryBytes      string                         `yaml:"memory,omitempty" json:"memory,omitempty"`
	Pids             int64                          `yaml:"pids,omitempty" json:"pids,omitempty"`
	Devices          []composetypes.DeviceRequest   `yaml:"devices,omitempty" json:"devices,omitempty"`
	GenericResources []composetypes.GenericResource `yaml:"generic_resources,omitempty" json:"generic_resources,omitempty"`

	Extensions composetypes.Extensions `yaml:"#extensions,inline,omitempty" json:"-"`
}
