package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterType                                      = "cluster"
	ClusterFieldAPIEndpoint                          = "apiEndpoint"
	ClusterFieldAgentFeatures                        = "agentFeatures"
	ClusterFieldAgentImage                           = "agentImage"
	ClusterFieldAgentImageOverride                   = "agentImageOverride"
	ClusterFieldAllocatable                          = "allocatable"
	ClusterFieldAnnotations                          = "annotations"
	ClusterFieldAppliedEnableNetworkPolicy           = "appliedEnableNetworkPolicy"
	ClusterFieldAppliedPodSecurityPolicyTemplateName = "appliedPodSecurityPolicyTemplateId"
	ClusterFieldAppliedSpec                          = "appliedSpec"
	ClusterFieldAuthImage                            = "authImage"
	ClusterFieldCACert                               = "caCert"
	ClusterFieldCapabilities                         = "capabilities"
	ClusterFieldCapacity                             = "capacity"
	ClusterFieldCertificatesExpiration               = "certificatesExpiration"
	ClusterFieldClusterTemplateAnswers               = "answers"
	ClusterFieldClusterTemplateID                    = "clusterTemplateId"
	ClusterFieldClusterTemplateQuestions             = "questions"
	ClusterFieldClusterTemplateRevisionID            = "clusterTemplateRevisionId"
	ClusterFieldComponentStatuses                    = "componentStatuses"
	ClusterFieldConditions                           = "conditions"
	ClusterFieldCreated                              = "created"
	ClusterFieldCreatorID                            = "creatorId"
	ClusterFieldCurrentCisRunName                    = "currentCisRunName"
	ClusterFieldDefaultClusterRoleForProjectMembers  = "defaultClusterRoleForProjectMembers"
	ClusterFieldDefaultPodSecurityPolicyTemplateID   = "defaultPodSecurityPolicyTemplateId"
	ClusterFieldDescription                          = "description"
	ClusterFieldDesiredAgentImage                    = "desiredAgentImage"
	ClusterFieldDesiredAuthImage                     = "desiredAuthImage"
	ClusterFieldDockerRootDir                        = "dockerRootDir"
	ClusterFieldDriver                               = "driver"
	ClusterFieldEnableClusterAlerting                = "enableClusterAlerting"
	ClusterFieldEnableClusterMonitoring              = "enableClusterMonitoring"
	ClusterFieldEnableDualStack                      = "enableDualStack"
	ClusterFieldEnableF5CIS                          = "enableF5CIS"
	ClusterFieldEnableGPUManagement                  = "enableGPUManagement"
	ClusterFieldEnableNetworkPolicy                  = "enableNetworkPolicy"
	ClusterFieldFailedSpec                           = "failedSpec"
	ClusterFieldFluentdLogDir                        = "fluentdLogDir"
	ClusterFieldGPUSchedulerNodePort                 = "gpuSchedulerNodePort"
	ClusterFieldImportedConfig                       = "importedConfig"
	ClusterFieldInternal                             = "internal"
	ClusterFieldIstioEnabled                         = "istioEnabled"
	ClusterFieldK3sConfig                            = "k3sConfig"
	ClusterFieldLabels                               = "labels"
	ClusterFieldLimits                               = "limits"
	ClusterFieldLocalClusterAuthEndpoint             = "localClusterAuthEndpoint"
	ClusterFieldMonitoringStatus                     = "monitoringStatus"
	ClusterFieldName                                 = "name"
	ClusterFieldNodeCount                            = "nodeCount"
	ClusterFieldNodeVersion                          = "nodeVersion"
	ClusterFieldOwnerReferences                      = "ownerReferences"
	ClusterFieldRancherKubernetesEngineConfig        = "rancherKubernetesEngineConfig"
	ClusterFieldRemoved                              = "removed"
	ClusterFieldRequested                            = "requested"
	ClusterFieldScheduledClusterScan                 = "scheduledClusterScan"
	ClusterFieldScheduledClusterScanStatus           = "scheduledClusterScanStatus"
	ClusterFieldState                                = "state"
	ClusterFieldSystemDefaultRegistry                = "systemDefaultRegistry"
	ClusterFieldTransitioning                        = "transitioning"
	ClusterFieldTransitioningMessage                 = "transitioningMessage"
	ClusterFieldUUID                                 = "uuid"
	ClusterFieldVersion                              = "version"
	ClusterFieldWindowsPreferedCluster               = "windowsPreferedCluster"
)

type Cluster struct {
	types.Resource
	APIEndpoint                          string                         `json:"apiEndpoint,omitempty" yaml:"apiEndpoint,omitempty"`
	AgentFeatures                        map[string]bool                `json:"agentFeatures,omitempty" yaml:"agentFeatures,omitempty"`
	AgentImage                           string                         `json:"agentImage,omitempty" yaml:"agentImage,omitempty"`
	AgentImageOverride                   string                         `json:"agentImageOverride,omitempty" yaml:"agentImageOverride,omitempty"`
	Allocatable                          map[string]string              `json:"allocatable,omitempty" yaml:"allocatable,omitempty"`
	Annotations                          map[string]string              `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	AppliedEnableNetworkPolicy           bool                           `json:"appliedEnableNetworkPolicy,omitempty" yaml:"appliedEnableNetworkPolicy,omitempty"`
	AppliedPodSecurityPolicyTemplateName string                         `json:"appliedPodSecurityPolicyTemplateId,omitempty" yaml:"appliedPodSecurityPolicyTemplateId,omitempty"`
	AppliedSpec                          *ClusterSpec                   `json:"appliedSpec,omitempty" yaml:"appliedSpec,omitempty"`
	AuthImage                            string                         `json:"authImage,omitempty" yaml:"authImage,omitempty"`
	CACert                               string                         `json:"caCert,omitempty" yaml:"caCert,omitempty"`
	Capabilities                         *Capabilities                  `json:"capabilities,omitempty" yaml:"capabilities,omitempty"`
	Capacity                             map[string]string              `json:"capacity,omitempty" yaml:"capacity,omitempty"`
	CertificatesExpiration               map[string]CertExpiration      `json:"certificatesExpiration,omitempty" yaml:"certificatesExpiration,omitempty"`
	ClusterTemplateAnswers               *Answer                        `json:"answers,omitempty" yaml:"answers,omitempty"`
	ClusterTemplateID                    string                         `json:"clusterTemplateId,omitempty" yaml:"clusterTemplateId,omitempty"`
	ClusterTemplateQuestions             []Question                     `json:"questions,omitempty" yaml:"questions,omitempty"`
	ClusterTemplateRevisionID            string                         `json:"clusterTemplateRevisionId,omitempty" yaml:"clusterTemplateRevisionId,omitempty"`
	ComponentStatuses                    []ClusterComponentStatus       `json:"componentStatuses,omitempty" yaml:"componentStatuses,omitempty"`
	Conditions                           []ClusterCondition             `json:"conditions,omitempty" yaml:"conditions,omitempty"`
	Created                              string                         `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID                            string                         `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	CurrentCisRunName                    string                         `json:"currentCisRunName,omitempty" yaml:"currentCisRunName,omitempty"`
	DefaultClusterRoleForProjectMembers  string                         `json:"defaultClusterRoleForProjectMembers,omitempty" yaml:"defaultClusterRoleForProjectMembers,omitempty"`
	DefaultPodSecurityPolicyTemplateID   string                         `json:"defaultPodSecurityPolicyTemplateId,omitempty" yaml:"defaultPodSecurityPolicyTemplateId,omitempty"`
	Description                          string                         `json:"description,omitempty" yaml:"description,omitempty"`
	DesiredAgentImage                    string                         `json:"desiredAgentImage,omitempty" yaml:"desiredAgentImage,omitempty"`
	DesiredAuthImage                     string                         `json:"desiredAuthImage,omitempty" yaml:"desiredAuthImage,omitempty"`
	DockerRootDir                        string                         `json:"dockerRootDir,omitempty" yaml:"dockerRootDir,omitempty"`
	Driver                               string                         `json:"driver,omitempty" yaml:"driver,omitempty"`
	EnableClusterAlerting                bool                           `json:"enableClusterAlerting,omitempty" yaml:"enableClusterAlerting,omitempty"`
	EnableClusterMonitoring              bool                           `json:"enableClusterMonitoring,omitempty" yaml:"enableClusterMonitoring,omitempty"`
	EnableDualStack                      bool                           `json:"enableDualStack,omitempty" yaml:"enableDualStack,omitempty"`
	EnableF5CIS                          bool                           `json:"enableF5CIS,omitempty" yaml:"enableF5CIS,omitempty"`
	EnableGPUManagement                  bool                           `json:"enableGPUManagement,omitempty" yaml:"enableGPUManagement,omitempty"`
	EnableNetworkPolicy                  *bool                          `json:"enableNetworkPolicy,omitempty" yaml:"enableNetworkPolicy,omitempty"`
	FailedSpec                           *ClusterSpec                   `json:"failedSpec,omitempty" yaml:"failedSpec,omitempty"`
	FluentdLogDir                        string                         `json:"fluentdLogDir,omitempty" yaml:"fluentdLogDir,omitempty"`
	GPUSchedulerNodePort                 string                         `json:"gpuSchedulerNodePort,omitempty" yaml:"gpuSchedulerNodePort,omitempty"`
	ImportedConfig                       *ImportedConfig                `json:"importedConfig,omitempty" yaml:"importedConfig,omitempty"`
	Internal                             bool                           `json:"internal,omitempty" yaml:"internal,omitempty"`
	IstioEnabled                         bool                           `json:"istioEnabled,omitempty" yaml:"istioEnabled,omitempty"`
	K3sConfig                            *K3sConfig                     `json:"k3sConfig,omitempty" yaml:"k3sConfig,omitempty"`
	Labels                               map[string]string              `json:"labels,omitempty" yaml:"labels,omitempty"`
	Limits                               map[string]string              `json:"limits,omitempty" yaml:"limits,omitempty"`
	LocalClusterAuthEndpoint             *LocalClusterAuthEndpoint      `json:"localClusterAuthEndpoint,omitempty" yaml:"localClusterAuthEndpoint,omitempty"`
	MonitoringStatus                     *MonitoringStatus              `json:"monitoringStatus,omitempty" yaml:"monitoringStatus,omitempty"`
	Name                                 string                         `json:"name,omitempty" yaml:"name,omitempty"`
	NodeCount                            int64                          `json:"nodeCount,omitempty" yaml:"nodeCount,omitempty"`
	NodeVersion                          int64                          `json:"nodeVersion,omitempty" yaml:"nodeVersion,omitempty"`
	OwnerReferences                      []OwnerReference               `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	RancherKubernetesEngineConfig        *RancherKubernetesEngineConfig `json:"rancherKubernetesEngineConfig,omitempty" yaml:"rancherKubernetesEngineConfig,omitempty"`
	Removed                              string                         `json:"removed,omitempty" yaml:"removed,omitempty"`
	Requested                            map[string]string              `json:"requested,omitempty" yaml:"requested,omitempty"`
	ScheduledClusterScan                 *ScheduledClusterScan          `json:"scheduledClusterScan,omitempty" yaml:"scheduledClusterScan,omitempty"`
	ScheduledClusterScanStatus           *ScheduledClusterScanStatus    `json:"scheduledClusterScanStatus,omitempty" yaml:"scheduledClusterScanStatus,omitempty"`
	State                                string                         `json:"state,omitempty" yaml:"state,omitempty"`
	SystemDefaultRegistry                string                         `json:"systemDefaultRegistry,omitempty" yaml:"systemDefaultRegistry,omitempty"`
	Transitioning                        string                         `json:"transitioning,omitempty" yaml:"transitioning,omitempty"`
	TransitioningMessage                 string                         `json:"transitioningMessage,omitempty" yaml:"transitioningMessage,omitempty"`
	UUID                                 string                         `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Version                              *Info                          `json:"version,omitempty" yaml:"version,omitempty"`
	WindowsPreferedCluster               bool                           `json:"windowsPreferedCluster,omitempty" yaml:"windowsPreferedCluster,omitempty"`
}

type ClusterCollection struct {
	types.Collection
	Data   []Cluster `json:"data,omitempty"`
	client *ClusterClient
}

type ClusterClient struct {
	apiClient *Client
}

type ClusterOperations interface {
	List(opts *types.ListOpts) (*ClusterCollection, error)
	ListAll(opts *types.ListOpts) (*ClusterCollection, error)
	Create(opts *Cluster) (*Cluster, error)
	Update(existing *Cluster, updates interface{}) (*Cluster, error)
	Replace(existing *Cluster) (*Cluster, error)
	ByID(id string) (*Cluster, error)
	Delete(container *Cluster) error

	ActionBackupEtcd(resource *Cluster) error

	ActionDisableF5CIS(resource *Cluster) error

	ActionDisableMonitoring(resource *Cluster) error

	ActionEditConnectionConfig(resource *Cluster, input *ConnectionConfig) error

	ActionEditF5CIS(resource *Cluster, input *F5CISInput) error

	ActionEditMonitoring(resource *Cluster, input *MonitoringInput) error

	ActionEnableF5CIS(resource *Cluster, input *F5CISInput) error

	ActionEnableMonitoring(resource *Cluster, input *MonitoringInput) error

	ActionExportYaml(resource *Cluster) (*ExportOutput, error)

	ActionGenerateKubeconfig(resource *Cluster) (*GenerateKubeConfigOutput, error)

	ActionImportYaml(resource *Cluster, input *ImportClusterYamlInput) (*ImportYamlOutput, error)

	ActionRestoreFromEtcdBackup(resource *Cluster, input *RestoreFromEtcdBackupInput) error

	ActionRotateCertificates(resource *Cluster, input *RotateCertificateInput) (*RotateCertificateOutput, error)

	ActionRunSecurityScan(resource *Cluster, input *CisScanConfig) error

	ActionSaveAsTemplate(resource *Cluster, input *SaveAsTemplateInput) (*SaveAsTemplateOutput, error)

	ActionValidateConnectionConfig(resource *Cluster, input *ConnectionConfig) error

	ActionViewConnectionConfig(resource *Cluster) (*ConnectionConfig, error)

	ActionViewF5CIS(resource *Cluster) (*F5CISOutput, error)

	ActionViewMonitoring(resource *Cluster) (*MonitoringOutput, error)
}

func newClusterClient(apiClient *Client) *ClusterClient {
	return &ClusterClient{
		apiClient: apiClient,
	}
}

func (c *ClusterClient) Create(container *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoCreate(ClusterType, container, resp)
	return resp, err
}

func (c *ClusterClient) Update(existing *Cluster, updates interface{}) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoUpdate(ClusterType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterClient) Replace(obj *Cluster) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoReplace(ClusterType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterClient) List(opts *types.ListOpts) (*ClusterCollection, error) {
	resp := &ClusterCollection{}
	err := c.apiClient.Ops.DoList(ClusterType, opts, resp)
	resp.client = c
	return resp, err
}

func (c *ClusterClient) ListAll(opts *types.ListOpts) (*ClusterCollection, error) {
	resp := &ClusterCollection{}
	resp, err := c.List(opts)
	if err != nil {
		return resp, err
	}
	data := resp.Data
	for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
		data = append(data, next.Data...)
		resp = next
		resp.Data = data
	}
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (cc *ClusterCollection) Next() (*ClusterCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterClient) ByID(id string) (*Cluster, error) {
	resp := &Cluster{}
	err := c.apiClient.Ops.DoByID(ClusterType, id, resp)
	return resp, err
}

func (c *ClusterClient) Delete(container *Cluster) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterType, &container.Resource)
}

func (c *ClusterClient) ActionBackupEtcd(resource *Cluster) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "backupEtcd", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterClient) ActionDisableF5CIS(resource *Cluster) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "disableF5CIS", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterClient) ActionDisableMonitoring(resource *Cluster) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "disableMonitoring", &resource.Resource, nil, nil)
	return err
}

func (c *ClusterClient) ActionEditConnectionConfig(resource *Cluster, input *ConnectionConfig) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "editConnectionConfig", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionEditF5CIS(resource *Cluster, input *F5CISInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "editF5CIS", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionEditMonitoring(resource *Cluster, input *MonitoringInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "editMonitoring", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionEnableF5CIS(resource *Cluster, input *F5CISInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "enableF5CIS", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionEnableMonitoring(resource *Cluster, input *MonitoringInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "enableMonitoring", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionExportYaml(resource *Cluster) (*ExportOutput, error) {
	resp := &ExportOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "exportYaml", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionGenerateKubeconfig(resource *Cluster) (*GenerateKubeConfigOutput, error) {
	resp := &GenerateKubeConfigOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "generateKubeconfig", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionImportYaml(resource *Cluster, input *ImportClusterYamlInput) (*ImportYamlOutput, error) {
	resp := &ImportYamlOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "importYaml", &resource.Resource, input, resp)
	return resp, err
}

func (c *ClusterClient) ActionRestoreFromEtcdBackup(resource *Cluster, input *RestoreFromEtcdBackupInput) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "restoreFromEtcdBackup", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionRotateCertificates(resource *Cluster, input *RotateCertificateInput) (*RotateCertificateOutput, error) {
	resp := &RotateCertificateOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "rotateCertificates", &resource.Resource, input, resp)
	return resp, err
}

func (c *ClusterClient) ActionRunSecurityScan(resource *Cluster, input *CisScanConfig) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "runSecurityScan", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionSaveAsTemplate(resource *Cluster, input *SaveAsTemplateInput) (*SaveAsTemplateOutput, error) {
	resp := &SaveAsTemplateOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "saveAsTemplate", &resource.Resource, input, resp)
	return resp, err
}

func (c *ClusterClient) ActionValidateConnectionConfig(resource *Cluster, input *ConnectionConfig) error {
	err := c.apiClient.Ops.DoAction(ClusterType, "validateConnectionConfig", &resource.Resource, input, nil)
	return err
}

func (c *ClusterClient) ActionViewConnectionConfig(resource *Cluster) (*ConnectionConfig, error) {
	resp := &ConnectionConfig{}
	err := c.apiClient.Ops.DoAction(ClusterType, "viewConnectionConfig", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionViewF5CIS(resource *Cluster) (*F5CISOutput, error) {
	resp := &F5CISOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "viewF5CIS", &resource.Resource, nil, resp)
	return resp, err
}

func (c *ClusterClient) ActionViewMonitoring(resource *Cluster) (*MonitoringOutput, error) {
	resp := &MonitoringOutput{}
	err := c.apiClient.Ops.DoAction(ClusterType, "viewMonitoring", &resource.Resource, nil, resp)
	return resp, err
}
