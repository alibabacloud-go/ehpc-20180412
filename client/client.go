// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddContainerAppRequest struct {
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
}

func (s AddContainerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppRequest) GoString() string {
	return s.String()
}

func (s *AddContainerAppRequest) SetName(v string) *AddContainerAppRequest {
	s.Name = &v
	return s
}

func (s *AddContainerAppRequest) SetRepository(v string) *AddContainerAppRequest {
	s.Repository = &v
	return s
}

func (s *AddContainerAppRequest) SetDescription(v string) *AddContainerAppRequest {
	s.Description = &v
	return s
}

func (s *AddContainerAppRequest) SetImageTag(v string) *AddContainerAppRequest {
	s.ImageTag = &v
	return s
}

func (s *AddContainerAppRequest) SetContainerType(v string) *AddContainerAppRequest {
	s.ContainerType = &v
	return s
}

type AddContainerAppResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ContainerId []*string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty" type:"Repeated"`
}

func (s AddContainerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponseBody) SetRequestId(v string) *AddContainerAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddContainerAppResponseBody) SetContainerId(v []*string) *AddContainerAppResponseBody {
	s.ContainerId = v
	return s
}

type AddContainerAppResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddContainerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s AddContainerAppResponse) GoString() string {
	return s.String()
}

func (s *AddContainerAppResponse) SetHeaders(v map[string]*string) *AddContainerAppResponse {
	s.Headers = v
	return s
}

func (s *AddContainerAppResponse) SetBody(v *AddContainerAppResponseBody) *AddContainerAppResponse {
	s.Body = v
	return s
}

type AddLocalNodesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Nodes     *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s AddLocalNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesRequest) GoString() string {
	return s.String()
}

func (s *AddLocalNodesRequest) SetClusterId(v string) *AddLocalNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddLocalNodesRequest) SetNodes(v string) *AddLocalNodesRequest {
	s.Nodes = &v
	return s
}

type AddLocalNodesResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s AddLocalNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponseBody) SetRequestId(v string) *AddLocalNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLocalNodesResponseBody) SetInstanceIds(v []*string) *AddLocalNodesResponseBody {
	s.InstanceIds = v
	return s
}

type AddLocalNodesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddLocalNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddLocalNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddLocalNodesResponse) GoString() string {
	return s.String()
}

func (s *AddLocalNodesResponse) SetHeaders(v map[string]*string) *AddLocalNodesResponse {
	s.Headers = v
	return s
}

func (s *AddLocalNodesResponse) SetBody(v *AddLocalNodesResponseBody) *AddLocalNodesResponse {
	s.Body = v
	return s
}

type AddNodesRequest struct {
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageOwnerAlias         *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageId                 *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Count                   *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	ComputeSpotStrategy     *string `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	ComputeSpotPriceLimit   *string `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	EcsChargeType           *string `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	Period                  *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenew               *string `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod         *int32  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	JobQueue                *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	CreateMode              *string `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	SystemDiskType          *string `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	SystemDiskSize          *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	ZoneId                  *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	HostNamePrefix          *string `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix          *string `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	ComputeEnableHt         *bool   `json:"ComputeEnableHt,omitempty" xml:"ComputeEnableHt,omitempty"`
	AllocatePublicAddress   *bool   `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandWidthIn  *int32  `json:"InternetMaxBandWidthIn,omitempty" xml:"InternetMaxBandWidthIn,omitempty"`
	InternetMaxBandWidthOut *int32  `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	ClientToken             *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddNodesRequest) GoString() string {
	return s.String()
}

func (s *AddNodesRequest) SetClusterId(v string) *AddNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *AddNodesRequest) SetImageOwnerAlias(v string) *AddNodesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *AddNodesRequest) SetImageId(v string) *AddNodesRequest {
	s.ImageId = &v
	return s
}

func (s *AddNodesRequest) SetCount(v int32) *AddNodesRequest {
	s.Count = &v
	return s
}

func (s *AddNodesRequest) SetInstanceType(v string) *AddNodesRequest {
	s.InstanceType = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotStrategy(v string) *AddNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *AddNodesRequest) SetComputeSpotPriceLimit(v string) *AddNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *AddNodesRequest) SetEcsChargeType(v string) *AddNodesRequest {
	s.EcsChargeType = &v
	return s
}

func (s *AddNodesRequest) SetPeriod(v int32) *AddNodesRequest {
	s.Period = &v
	return s
}

func (s *AddNodesRequest) SetPeriodUnit(v string) *AddNodesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AddNodesRequest) SetAutoRenew(v string) *AddNodesRequest {
	s.AutoRenew = &v
	return s
}

func (s *AddNodesRequest) SetAutoRenewPeriod(v int32) *AddNodesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *AddNodesRequest) SetJobQueue(v string) *AddNodesRequest {
	s.JobQueue = &v
	return s
}

func (s *AddNodesRequest) SetCreateMode(v string) *AddNodesRequest {
	s.CreateMode = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskType(v string) *AddNodesRequest {
	s.SystemDiskType = &v
	return s
}

func (s *AddNodesRequest) SetSystemDiskSize(v int32) *AddNodesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *AddNodesRequest) SetZoneId(v string) *AddNodesRequest {
	s.ZoneId = &v
	return s
}

func (s *AddNodesRequest) SetVSwitchId(v string) *AddNodesRequest {
	s.VSwitchId = &v
	return s
}

func (s *AddNodesRequest) SetHostNamePrefix(v string) *AddNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *AddNodesRequest) SetHostNameSuffix(v string) *AddNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *AddNodesRequest) SetComputeEnableHt(v bool) *AddNodesRequest {
	s.ComputeEnableHt = &v
	return s
}

func (s *AddNodesRequest) SetAllocatePublicAddress(v bool) *AddNodesRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *AddNodesRequest) SetInternetChargeType(v string) *AddNodesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *AddNodesRequest) SetInternetMaxBandWidthIn(v int32) *AddNodesRequest {
	s.InternetMaxBandWidthIn = &v
	return s
}

func (s *AddNodesRequest) SetInternetMaxBandWidthOut(v int32) *AddNodesRequest {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *AddNodesRequest) SetClientToken(v string) *AddNodesRequest {
	s.ClientToken = &v
	return s
}

type AddNodesResponseBody struct {
	TaskId      *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s AddNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponseBody) GoString() string {
	return s.String()
}

func (s *AddNodesResponseBody) SetTaskId(v string) *AddNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *AddNodesResponseBody) SetRequestId(v string) *AddNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddNodesResponseBody) SetInstanceIds(v []*string) *AddNodesResponseBody {
	s.InstanceIds = v
	return s
}

type AddNodesResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddNodesResponse) GoString() string {
	return s.String()
}

func (s *AddNodesResponse) SetHeaders(v map[string]*string) *AddNodesResponse {
	s.Headers = v
	return s
}

func (s *AddNodesResponse) SetBody(v *AddNodesResponseBody) *AddNodesResponse {
	s.Body = v
	return s
}

type AddQueueRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s AddQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s AddQueueRequest) GoString() string {
	return s.String()
}

func (s *AddQueueRequest) SetClusterId(v string) *AddQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *AddQueueRequest) SetQueueName(v string) *AddQueueRequest {
	s.QueueName = &v
	return s
}

type AddQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddQueueResponseBody) GoString() string {
	return s.String()
}

func (s *AddQueueResponseBody) SetRequestId(v string) *AddQueueResponseBody {
	s.RequestId = &v
	return s
}

type AddQueueResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s AddQueueResponse) GoString() string {
	return s.String()
}

func (s *AddQueueResponse) SetHeaders(v map[string]*string) *AddQueueResponse {
	s.Headers = v
	return s
}

func (s *AddQueueResponse) SetBody(v *AddQueueResponseBody) *AddQueueResponse {
	s.Body = v
	return s
}

type AddSecurityGroupRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s AddSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupRequest) SetClusterId(v string) *AddSecurityGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *AddSecurityGroupRequest) SetSecurityGroupId(v string) *AddSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AddSecurityGroupRequest) SetClientToken(v string) *AddSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

type AddSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupResponseBody) SetRequestId(v string) *AddSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddSecurityGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *AddSecurityGroupResponse) SetHeaders(v map[string]*string) *AddSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *AddSecurityGroupResponse) SetBody(v *AddSecurityGroupResponseBody) *AddSecurityGroupResponse {
	s.Body = v
	return s
}

type AddUsersRequest struct {
	ClusterId *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*AddUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s AddUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequest) GoString() string {
	return s.String()
}

func (s *AddUsersRequest) SetClusterId(v string) *AddUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *AddUsersRequest) SetUser(v []*AddUsersRequestUser) *AddUsersRequest {
	s.User = v
	return s
}

type AddUsersRequestUser struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Group    *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s AddUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s AddUsersRequestUser) GoString() string {
	return s.String()
}

func (s *AddUsersRequestUser) SetPassword(v string) *AddUsersRequestUser {
	s.Password = &v
	return s
}

func (s *AddUsersRequestUser) SetName(v string) *AddUsersRequestUser {
	s.Name = &v
	return s
}

func (s *AddUsersRequestUser) SetGroup(v string) *AddUsersRequestUser {
	s.Group = &v
	return s
}

type AddUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUsersResponseBody) SetRequestId(v string) *AddUsersResponseBody {
	s.RequestId = &v
	return s
}

type AddUsersResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s AddUsersResponse) GoString() string {
	return s.String()
}

func (s *AddUsersResponse) SetHeaders(v map[string]*string) *AddUsersResponse {
	s.Headers = v
	return s
}

func (s *AddUsersResponse) SetBody(v *AddUsersResponseBody) *AddUsersResponse {
	s.Body = v
	return s
}

type ApplyNodesRequest struct {
	ClusterId                     *string                               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageId                       *string                               `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ComputeSpotStrategy           *string                               `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	ComputeSpotPriceLimit         *float32                              `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	SystemDiskType                *string                               `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	SystemDiskSize                *int32                                `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	HostNamePrefix                *string                               `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	HostNameSuffix                *string                               `json:"HostNameSuffix,omitempty" xml:"HostNameSuffix,omitempty"`
	AllocatePublicAddress         *bool                                 `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	InternetChargeType            *string                               `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandWidthIn        *int32                                `json:"InternetMaxBandWidthIn,omitempty" xml:"InternetMaxBandWidthIn,omitempty"`
	InternetMaxBandWidthOut       *int32                                `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
	Cores                         *int32                                `json:"Cores,omitempty" xml:"Cores,omitempty"`
	Memory                        *int32                                `json:"Memory,omitempty" xml:"Memory,omitempty"`
	InstanceFamilyLevel           *string                               `json:"InstanceFamilyLevel,omitempty" xml:"InstanceFamilyLevel,omitempty"`
	TargetCapacity                *int32                                `json:"TargetCapacity,omitempty" xml:"TargetCapacity,omitempty"`
	ResourceAmountType            *string                               `json:"ResourceAmountType,omitempty" xml:"ResourceAmountType,omitempty"`
	PriorityStrategy              *string                               `json:"PriorityStrategy,omitempty" xml:"PriorityStrategy,omitempty"`
	StrictSatisfiedTargetCapacity *bool                                 `json:"StrictSatisfiedTargetCapacity,omitempty" xml:"StrictSatisfiedTargetCapacity,omitempty"`
	ZoneInfos                     []*ApplyNodesRequestZoneInfos         `json:"ZoneInfos,omitempty" xml:"ZoneInfos,omitempty" type:"Repeated"`
	InstanceTypeModel             []*ApplyNodesRequestInstanceTypeModel `json:"InstanceTypeModel,omitempty" xml:"InstanceTypeModel,omitempty" type:"Repeated"`
}

func (s ApplyNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequest) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequest) SetClusterId(v string) *ApplyNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ApplyNodesRequest) SetImageId(v string) *ApplyNodesRequest {
	s.ImageId = &v
	return s
}

func (s *ApplyNodesRequest) SetComputeSpotStrategy(v string) *ApplyNodesRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *ApplyNodesRequest) SetComputeSpotPriceLimit(v float32) *ApplyNodesRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskType(v string) *ApplyNodesRequest {
	s.SystemDiskType = &v
	return s
}

func (s *ApplyNodesRequest) SetSystemDiskSize(v int32) *ApplyNodesRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *ApplyNodesRequest) SetHostNamePrefix(v string) *ApplyNodesRequest {
	s.HostNamePrefix = &v
	return s
}

func (s *ApplyNodesRequest) SetHostNameSuffix(v string) *ApplyNodesRequest {
	s.HostNameSuffix = &v
	return s
}

func (s *ApplyNodesRequest) SetAllocatePublicAddress(v bool) *ApplyNodesRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetChargeType(v string) *ApplyNodesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetMaxBandWidthIn(v int32) *ApplyNodesRequest {
	s.InternetMaxBandWidthIn = &v
	return s
}

func (s *ApplyNodesRequest) SetInternetMaxBandWidthOut(v int32) *ApplyNodesRequest {
	s.InternetMaxBandWidthOut = &v
	return s
}

func (s *ApplyNodesRequest) SetCores(v int32) *ApplyNodesRequest {
	s.Cores = &v
	return s
}

func (s *ApplyNodesRequest) SetMemory(v int32) *ApplyNodesRequest {
	s.Memory = &v
	return s
}

func (s *ApplyNodesRequest) SetInstanceFamilyLevel(v string) *ApplyNodesRequest {
	s.InstanceFamilyLevel = &v
	return s
}

func (s *ApplyNodesRequest) SetTargetCapacity(v int32) *ApplyNodesRequest {
	s.TargetCapacity = &v
	return s
}

func (s *ApplyNodesRequest) SetResourceAmountType(v string) *ApplyNodesRequest {
	s.ResourceAmountType = &v
	return s
}

func (s *ApplyNodesRequest) SetPriorityStrategy(v string) *ApplyNodesRequest {
	s.PriorityStrategy = &v
	return s
}

func (s *ApplyNodesRequest) SetStrictSatisfiedTargetCapacity(v bool) *ApplyNodesRequest {
	s.StrictSatisfiedTargetCapacity = &v
	return s
}

func (s *ApplyNodesRequest) SetZoneInfos(v []*ApplyNodesRequestZoneInfos) *ApplyNodesRequest {
	s.ZoneInfos = v
	return s
}

func (s *ApplyNodesRequest) SetInstanceTypeModel(v []*ApplyNodesRequestInstanceTypeModel) *ApplyNodesRequest {
	s.InstanceTypeModel = v
	return s
}

type ApplyNodesRequestZoneInfos struct {
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId    *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ApplyNodesRequestZoneInfos) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestZoneInfos) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestZoneInfos) SetVSwitchId(v string) *ApplyNodesRequestZoneInfos {
	s.VSwitchId = &v
	return s
}

func (s *ApplyNodesRequestZoneInfos) SetZoneId(v string) *ApplyNodesRequestZoneInfos {
	s.ZoneId = &v
	return s
}

type ApplyNodesRequestInstanceTypeModel struct {
	MaxPrice      *float32 `json:"MaxPrice,omitempty" xml:"MaxPrice,omitempty"`
	TargetImageId *string  `json:"TargetImageId,omitempty" xml:"TargetImageId,omitempty"`
	InstanceType  *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ApplyNodesRequestInstanceTypeModel) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesRequestInstanceTypeModel) GoString() string {
	return s.String()
}

func (s *ApplyNodesRequestInstanceTypeModel) SetMaxPrice(v float32) *ApplyNodesRequestInstanceTypeModel {
	s.MaxPrice = &v
	return s
}

func (s *ApplyNodesRequestInstanceTypeModel) SetTargetImageId(v string) *ApplyNodesRequestInstanceTypeModel {
	s.TargetImageId = &v
	return s
}

func (s *ApplyNodesRequestInstanceTypeModel) SetInstanceType(v string) *ApplyNodesRequestInstanceTypeModel {
	s.InstanceType = &v
	return s
}

type ApplyNodesResponseBody struct {
	TaskId          *string   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SatisfiedAmount *int32    `json:"SatisfiedAmount,omitempty" xml:"SatisfiedAmount,omitempty"`
	InstanceIds     []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	Detail          *string   `json:"Detail,omitempty" xml:"Detail,omitempty"`
}

func (s ApplyNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponseBody) SetTaskId(v string) *ApplyNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ApplyNodesResponseBody) SetRequestId(v string) *ApplyNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyNodesResponseBody) SetSatisfiedAmount(v int32) *ApplyNodesResponseBody {
	s.SatisfiedAmount = &v
	return s
}

func (s *ApplyNodesResponseBody) SetInstanceIds(v []*string) *ApplyNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *ApplyNodesResponseBody) SetDetail(v string) *ApplyNodesResponseBody {
	s.Detail = &v
	return s
}

type ApplyNodesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ApplyNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ApplyNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyNodesResponse) GoString() string {
	return s.String()
}

func (s *ApplyNodesResponse) SetHeaders(v map[string]*string) *ApplyNodesResponse {
	s.Headers = v
	return s
}

func (s *ApplyNodesResponse) SetBody(v *ApplyNodesResponseBody) *ApplyNodesResponse {
	s.Body = v
	return s
}

type BindAccountToClusterUserRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPwd     *string `json:"UserPwd,omitempty" xml:"UserPwd,omitempty"`
	AccountUid  *string `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s BindAccountToClusterUserRequest) String() string {
	return tea.Prettify(s)
}

func (s BindAccountToClusterUserRequest) GoString() string {
	return s.String()
}

func (s *BindAccountToClusterUserRequest) SetClusterId(v string) *BindAccountToClusterUserRequest {
	s.ClusterId = &v
	return s
}

func (s *BindAccountToClusterUserRequest) SetUserName(v string) *BindAccountToClusterUserRequest {
	s.UserName = &v
	return s
}

func (s *BindAccountToClusterUserRequest) SetUserPwd(v string) *BindAccountToClusterUserRequest {
	s.UserPwd = &v
	return s
}

func (s *BindAccountToClusterUserRequest) SetAccountUid(v string) *BindAccountToClusterUserRequest {
	s.AccountUid = &v
	return s
}

func (s *BindAccountToClusterUserRequest) SetAccountName(v string) *BindAccountToClusterUserRequest {
	s.AccountName = &v
	return s
}

type BindAccountToClusterUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindAccountToClusterUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindAccountToClusterUserResponseBody) GoString() string {
	return s.String()
}

func (s *BindAccountToClusterUserResponseBody) SetRequestId(v string) *BindAccountToClusterUserResponseBody {
	s.RequestId = &v
	return s
}

type BindAccountToClusterUserResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BindAccountToClusterUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindAccountToClusterUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BindAccountToClusterUserResponse) GoString() string {
	return s.String()
}

func (s *BindAccountToClusterUserResponse) SetHeaders(v map[string]*string) *BindAccountToClusterUserResponse {
	s.Headers = v
	return s
}

func (s *BindAccountToClusterUserResponse) SetBody(v *BindAccountToClusterUserResponseBody) *BindAccountToClusterUserResponse {
	s.Body = v
	return s
}

type CreateClusterRequest struct {
	EcsOrder              *CreateClusterRequestEcsOrder            `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	ZoneId                *string                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Name                  *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Description           *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EhpcVersion           *string                                  `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	ClientVersion         *string                                  `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	OsTag                 *string                                  `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	AccountType           *string                                  `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	SchedulerType         *string                                  `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	SecurityGroupId       *string                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName     *string                                  `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	VpcId                 *string                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId             *string                                  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeType            *string                                  `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VolumeId              *string                                  `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeProtocol        *string                                  `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeMountpoint      *string                                  `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	RemoteDirectory       *string                                  `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	DeployMode            *string                                  `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	HaEnable              *bool                                    `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	EcsChargeType         *string                                  `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	Password              *string                                  `json:"Password,omitempty" xml:"Password,omitempty"`
	KeyPairName           *string                                  `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	ImageOwnerAlias       *string                                  `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageId               *string                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	SccClusterId          *string                                  `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	ComputeSpotStrategy   *string                                  `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	ComputeSpotPriceLimit *string                                  `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ComputeEnableHt       *bool                                    `json:"ComputeEnableHt,omitempty" xml:"ComputeEnableHt,omitempty"`
	Period                *int32                                   `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit            *string                                  `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenew             *string                                  `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod       *int32                                   `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	InputFileUrl          *string                                  `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	JobQueue              *string                                  `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	SystemDiskType        *string                                  `json:"SystemDiskType,omitempty" xml:"SystemDiskType,omitempty"`
	SystemDiskSize        *int32                                   `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	RemoteVisEnable       *string                                  `json:"RemoteVisEnable,omitempty" xml:"RemoteVisEnable,omitempty"`
	ResourceGroupId       *string                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ClientToken           *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	WithoutElasticIp      *bool                                    `json:"WithoutElasticIp,omitempty" xml:"WithoutElasticIp,omitempty"`
	Application           []*CreateClusterRequestApplication       `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	AdditionalVolumes     []*CreateClusterRequestAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
	PostInstallScript     []*CreateClusterRequestPostInstallScript `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetEcsOrder(v *CreateClusterRequestEcsOrder) *CreateClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateClusterRequest) SetZoneId(v string) *CreateClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequest) SetName(v string) *CreateClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateClusterRequest) SetDescription(v string) *CreateClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateClusterRequest) SetEhpcVersion(v string) *CreateClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateClusterRequest) SetClientVersion(v string) *CreateClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateClusterRequest) SetOsTag(v string) *CreateClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateClusterRequest) SetAccountType(v string) *CreateClusterRequest {
	s.AccountType = &v
	return s
}

func (s *CreateClusterRequest) SetSchedulerType(v string) *CreateClusterRequest {
	s.SchedulerType = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupId(v string) *CreateClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetSecurityGroupName(v string) *CreateClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) SetVSwitchId(v string) *CreateClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeType(v string) *CreateClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeId(v string) *CreateClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeProtocol(v string) *CreateClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequest) SetVolumeMountpoint(v string) *CreateClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteDirectory(v string) *CreateClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequest) SetDeployMode(v string) *CreateClusterRequest {
	s.DeployMode = &v
	return s
}

func (s *CreateClusterRequest) SetHaEnable(v bool) *CreateClusterRequest {
	s.HaEnable = &v
	return s
}

func (s *CreateClusterRequest) SetEcsChargeType(v string) *CreateClusterRequest {
	s.EcsChargeType = &v
	return s
}

func (s *CreateClusterRequest) SetPassword(v string) *CreateClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateClusterRequest) SetKeyPairName(v string) *CreateClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequest) SetImageOwnerAlias(v string) *CreateClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateClusterRequest) SetImageId(v string) *CreateClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequest) SetSccClusterId(v string) *CreateClusterRequest {
	s.SccClusterId = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotStrategy(v string) *CreateClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateClusterRequest) SetComputeSpotPriceLimit(v string) *CreateClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateClusterRequest) SetComputeEnableHt(v bool) *CreateClusterRequest {
	s.ComputeEnableHt = &v
	return s
}

func (s *CreateClusterRequest) SetPeriod(v int32) *CreateClusterRequest {
	s.Period = &v
	return s
}

func (s *CreateClusterRequest) SetPeriodUnit(v string) *CreateClusterRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenew(v string) *CreateClusterRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateClusterRequest) SetAutoRenewPeriod(v int32) *CreateClusterRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateClusterRequest) SetInputFileUrl(v string) *CreateClusterRequest {
	s.InputFileUrl = &v
	return s
}

func (s *CreateClusterRequest) SetJobQueue(v string) *CreateClusterRequest {
	s.JobQueue = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskType(v string) *CreateClusterRequest {
	s.SystemDiskType = &v
	return s
}

func (s *CreateClusterRequest) SetSystemDiskSize(v int32) *CreateClusterRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterRequest) SetRemoteVisEnable(v string) *CreateClusterRequest {
	s.RemoteVisEnable = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetClientToken(v string) *CreateClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateClusterRequest) SetWithoutElasticIp(v bool) *CreateClusterRequest {
	s.WithoutElasticIp = &v
	return s
}

func (s *CreateClusterRequest) SetApplication(v []*CreateClusterRequestApplication) *CreateClusterRequest {
	s.Application = v
	return s
}

func (s *CreateClusterRequest) SetAdditionalVolumes(v []*CreateClusterRequestAdditionalVolumes) *CreateClusterRequest {
	s.AdditionalVolumes = v
	return s
}

func (s *CreateClusterRequest) SetPostInstallScript(v []*CreateClusterRequestPostInstallScript) *CreateClusterRequest {
	s.PostInstallScript = v
	return s
}

type CreateClusterRequestEcsOrder struct {
	Manager *CreateClusterRequestEcsOrderManager `json:"Manager,omitempty" xml:"Manager,omitempty" require:"true" type:"Struct"`
	Compute *CreateClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
	Login   *CreateClusterRequestEcsOrderLogin   `json:"Login,omitempty" xml:"Login,omitempty" require:"true" type:"Struct"`
}

func (s CreateClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrder) SetManager(v *CreateClusterRequestEcsOrderManager) *CreateClusterRequestEcsOrder {
	s.Manager = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetCompute(v *CreateClusterRequestEcsOrderCompute) *CreateClusterRequestEcsOrder {
	s.Compute = v
	return s
}

func (s *CreateClusterRequestEcsOrder) SetLogin(v *CreateClusterRequestEcsOrderLogin) *CreateClusterRequestEcsOrder {
	s.Login = v
	return s
}

type CreateClusterRequestEcsOrderManager struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderManager) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderManager) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderManager) SetCount(v int32) *CreateClusterRequestEcsOrderManager {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderManager) SetInstanceType(v string) *CreateClusterRequestEcsOrderManager {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderCompute struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderCompute) SetCount(v int32) *CreateClusterRequestEcsOrderCompute {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestEcsOrderLogin struct {
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateClusterRequestEcsOrderLogin) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestEcsOrderLogin) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestEcsOrderLogin) SetCount(v int32) *CreateClusterRequestEcsOrderLogin {
	s.Count = &v
	return s
}

func (s *CreateClusterRequestEcsOrderLogin) SetInstanceType(v string) *CreateClusterRequestEcsOrderLogin {
	s.InstanceType = &v
	return s
}

type CreateClusterRequestApplication struct {
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestApplication) SetTag(v string) *CreateClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateClusterRequestAdditionalVolumes struct {
	JobQueue         *string                                       `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	VolumeId         *string                                       `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	Roles            []*CreateClusterRequestAdditionalVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	RemoteDirectory  *string                                       `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeMountpoint *string                                       `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	LocalDirectory   *string                                       `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	VolumeType       *string                                       `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VolumeProtocol   *string                                       `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	Location         *string                                       `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s CreateClusterRequestAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalVolumes) SetJobQueue(v string) *CreateClusterRequestAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeId(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetRoles(v []*CreateClusterRequestAdditionalVolumesRoles) *CreateClusterRequestAdditionalVolumes {
	s.Roles = v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetRemoteDirectory(v string) *CreateClusterRequestAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeMountpoint(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetLocalDirectory(v string) *CreateClusterRequestAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeType(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeType = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetVolumeProtocol(v string) *CreateClusterRequestAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateClusterRequestAdditionalVolumes) SetLocation(v string) *CreateClusterRequestAdditionalVolumes {
	s.Location = &v
	return s
}

type CreateClusterRequestAdditionalVolumesRoles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateClusterRequestAdditionalVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestAdditionalVolumesRoles) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestAdditionalVolumesRoles) SetName(v string) *CreateClusterRequestAdditionalVolumesRoles {
	s.Name = &v
	return s
}

type CreateClusterRequestPostInstallScript struct {
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateClusterRequestPostInstallScript) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestPostInstallScript) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestPostInstallScript) SetArgs(v string) *CreateClusterRequestPostInstallScript {
	s.Args = &v
	return s
}

func (s *CreateClusterRequestPostInstallScript) SetUrl(v string) *CreateClusterRequestPostInstallScript {
	s.Url = &v
	return s
}

type CreateClusterResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateGWSClusterRequest struct {
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateGWSClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterRequest) SetVpcId(v string) *CreateGWSClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGWSClusterRequest) SetClusterType(v string) *CreateGWSClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateGWSClusterRequest) SetName(v string) *CreateGWSClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateGWSClusterRequest) SetVSwitchId(v string) *CreateGWSClusterRequest {
	s.VSwitchId = &v
	return s
}

type CreateGWSClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateGWSClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterResponseBody) SetRequestId(v string) *CreateGWSClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGWSClusterResponseBody) SetClusterId(v string) *CreateGWSClusterResponseBody {
	s.ClusterId = &v
	return s
}

type CreateGWSClusterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGWSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSClusterResponse) SetHeaders(v map[string]*string) *CreateGWSClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSClusterResponse) SetBody(v *CreateGWSClusterResponseBody) *CreateGWSClusterResponse {
	s.Body = v
	return s
}

type CreateGWSImageRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateGWSImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSImageRequest) SetInstanceId(v string) *CreateGWSImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGWSImageRequest) SetName(v string) *CreateGWSImageRequest {
	s.Name = &v
	return s
}

type CreateGWSImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s CreateGWSImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSImageResponseBody) SetRequestId(v string) *CreateGWSImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGWSImageResponseBody) SetImageId(v string) *CreateGWSImageResponseBody {
	s.ImageId = &v
	return s
}

type CreateGWSImageResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGWSImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSImageResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSImageResponse) SetHeaders(v map[string]*string) *CreateGWSImageResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSImageResponse) SetBody(v *CreateGWSImageResponseBody) *CreateGWSImageResponse {
	s.Body = v
	return s
}

type CreateGWSInstanceRequest struct {
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ImageId                 *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	SystemDiskSize          *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	SystemDiskCategory      *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceChargeType      *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	WorkMode                *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	AllocatePublicAddress   *bool   `json:"AllocatePublicAddress,omitempty" xml:"AllocatePublicAddress,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	InternetMaxBandwidthIn  *int32  `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	InternetMaxBandwidthOut *int32  `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	Name                    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Period                  *string `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit              *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	AutoRenew               *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AppList                 *string `json:"AppList,omitempty" xml:"AppList,omitempty"`
	VSwitchId               *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceRequest) SetClusterId(v string) *CreateGWSInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetImageId(v string) *CreateGWSInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetSystemDiskSize(v int32) *CreateGWSInstanceRequest {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetSystemDiskCategory(v string) *CreateGWSInstanceRequest {
	s.SystemDiskCategory = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInstanceType(v string) *CreateGWSInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInstanceChargeType(v string) *CreateGWSInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetWorkMode(v string) *CreateGWSInstanceRequest {
	s.WorkMode = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAllocatePublicAddress(v bool) *CreateGWSInstanceRequest {
	s.AllocatePublicAddress = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetChargeType(v string) *CreateGWSInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateGWSInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateGWSInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetName(v string) *CreateGWSInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetPeriod(v string) *CreateGWSInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetPeriodUnit(v string) *CreateGWSInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAutoRenew(v bool) *CreateGWSInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetAppList(v string) *CreateGWSInstanceRequest {
	s.AppList = &v
	return s
}

func (s *CreateGWSInstanceRequest) SetVSwitchId(v string) *CreateGWSInstanceRequest {
	s.VSwitchId = &v
	return s
}

type CreateGWSInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceResponseBody) SetRequestId(v string) *CreateGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGWSInstanceResponseBody) SetInstanceId(v string) *CreateGWSInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type CreateGWSInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGWSInstanceResponse) SetHeaders(v map[string]*string) *CreateGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGWSInstanceResponse) SetBody(v *CreateGWSInstanceResponseBody) *CreateGWSInstanceResponse {
	s.Body = v
	return s
}

type CreateHybridClusterRequest struct {
	EcsOrder                  *CreateHybridClusterRequestEcsOrder            `json:"EcsOrder,omitempty" xml:"EcsOrder,omitempty" type:"Struct"`
	ZoneId                    *string                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Name                      *string                                        `json:"Name,omitempty" xml:"Name,omitempty"`
	Description               *string                                        `json:"Description,omitempty" xml:"Description,omitempty"`
	EhpcVersion               *string                                        `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	ClientVersion             *string                                        `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	OsTag                     *string                                        `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Domain                    *string                                        `json:"Domain,omitempty" xml:"Domain,omitempty"`
	Location                  *string                                        `json:"Location,omitempty" xml:"Location,omitempty"`
	SecurityGroupId           *string                                        `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName         *string                                        `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	VpcId                     *string                                        `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VSwitchId                 *string                                        `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VolumeType                *string                                        `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VolumeId                  *string                                        `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	VolumeProtocol            *string                                        `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	VolumeMountpoint          *string                                        `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	RemoteDirectory           *string                                        `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	OnPremiseVolumeProtocol   *string                                        `json:"OnPremiseVolumeProtocol,omitempty" xml:"OnPremiseVolumeProtocol,omitempty"`
	OnPremiseVolumeMountPoint *string                                        `json:"OnPremiseVolumeMountPoint,omitempty" xml:"OnPremiseVolumeMountPoint,omitempty"`
	OnPremiseVolumeRemotePath *string                                        `json:"OnPremiseVolumeRemotePath,omitempty" xml:"OnPremiseVolumeRemotePath,omitempty"`
	OnPremiseVolumeLocalPath  *string                                        `json:"OnPremiseVolumeLocalPath,omitempty" xml:"OnPremiseVolumeLocalPath,omitempty"`
	Password                  *string                                        `json:"Password,omitempty" xml:"Password,omitempty"`
	KeyPairName               *string                                        `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	JobQueue                  *string                                        `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	ResourceGroupId           *string                                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SchedulerPreInstall       *bool                                          `json:"SchedulerPreInstall,omitempty" xml:"SchedulerPreInstall,omitempty"`
	ComputeSpotStrategy       *string                                        `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	ComputeSpotPriceLimit     *float32                                       `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	ImageOwnerAlias           *string                                        `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageId                   *string                                        `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ClientToken               *string                                        `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Nodes                     []*CreateHybridClusterRequestNodes             `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Application               []*CreateHybridClusterRequestApplication       `json:"Application,omitempty" xml:"Application,omitempty" type:"Repeated"`
	PostInstallScript         []*CreateHybridClusterRequestPostInstallScript `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty" type:"Repeated"`
}

func (s CreateHybridClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequest) SetEcsOrder(v *CreateHybridClusterRequestEcsOrder) *CreateHybridClusterRequest {
	s.EcsOrder = v
	return s
}

func (s *CreateHybridClusterRequest) SetZoneId(v string) *CreateHybridClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetName(v string) *CreateHybridClusterRequest {
	s.Name = &v
	return s
}

func (s *CreateHybridClusterRequest) SetDescription(v string) *CreateHybridClusterRequest {
	s.Description = &v
	return s
}

func (s *CreateHybridClusterRequest) SetEhpcVersion(v string) *CreateHybridClusterRequest {
	s.EhpcVersion = &v
	return s
}

func (s *CreateHybridClusterRequest) SetClientVersion(v string) *CreateHybridClusterRequest {
	s.ClientVersion = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOsTag(v string) *CreateHybridClusterRequest {
	s.OsTag = &v
	return s
}

func (s *CreateHybridClusterRequest) SetDomain(v string) *CreateHybridClusterRequest {
	s.Domain = &v
	return s
}

func (s *CreateHybridClusterRequest) SetLocation(v string) *CreateHybridClusterRequest {
	s.Location = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSecurityGroupId(v string) *CreateHybridClusterRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSecurityGroupName(v string) *CreateHybridClusterRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVpcId(v string) *CreateHybridClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVSwitchId(v string) *CreateHybridClusterRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeType(v string) *CreateHybridClusterRequest {
	s.VolumeType = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeId(v string) *CreateHybridClusterRequest {
	s.VolumeId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeProtocol(v string) *CreateHybridClusterRequest {
	s.VolumeProtocol = &v
	return s
}

func (s *CreateHybridClusterRequest) SetVolumeMountpoint(v string) *CreateHybridClusterRequest {
	s.VolumeMountpoint = &v
	return s
}

func (s *CreateHybridClusterRequest) SetRemoteDirectory(v string) *CreateHybridClusterRequest {
	s.RemoteDirectory = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeProtocol(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeProtocol = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeMountPoint(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeMountPoint = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeRemotePath(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeRemotePath = &v
	return s
}

func (s *CreateHybridClusterRequest) SetOnPremiseVolumeLocalPath(v string) *CreateHybridClusterRequest {
	s.OnPremiseVolumeLocalPath = &v
	return s
}

func (s *CreateHybridClusterRequest) SetPassword(v string) *CreateHybridClusterRequest {
	s.Password = &v
	return s
}

func (s *CreateHybridClusterRequest) SetKeyPairName(v string) *CreateHybridClusterRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateHybridClusterRequest) SetJobQueue(v string) *CreateHybridClusterRequest {
	s.JobQueue = &v
	return s
}

func (s *CreateHybridClusterRequest) SetResourceGroupId(v string) *CreateHybridClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetSchedulerPreInstall(v bool) *CreateHybridClusterRequest {
	s.SchedulerPreInstall = &v
	return s
}

func (s *CreateHybridClusterRequest) SetComputeSpotStrategy(v string) *CreateHybridClusterRequest {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *CreateHybridClusterRequest) SetComputeSpotPriceLimit(v float32) *CreateHybridClusterRequest {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *CreateHybridClusterRequest) SetImageOwnerAlias(v string) *CreateHybridClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateHybridClusterRequest) SetImageId(v string) *CreateHybridClusterRequest {
	s.ImageId = &v
	return s
}

func (s *CreateHybridClusterRequest) SetClientToken(v string) *CreateHybridClusterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHybridClusterRequest) SetNodes(v []*CreateHybridClusterRequestNodes) *CreateHybridClusterRequest {
	s.Nodes = v
	return s
}

func (s *CreateHybridClusterRequest) SetApplication(v []*CreateHybridClusterRequestApplication) *CreateHybridClusterRequest {
	s.Application = v
	return s
}

func (s *CreateHybridClusterRequest) SetPostInstallScript(v []*CreateHybridClusterRequestPostInstallScript) *CreateHybridClusterRequest {
	s.PostInstallScript = v
	return s
}

type CreateHybridClusterRequestEcsOrder struct {
	Compute *CreateHybridClusterRequestEcsOrderCompute `json:"Compute,omitempty" xml:"Compute,omitempty" require:"true" type:"Struct"`
}

func (s CreateHybridClusterRequestEcsOrder) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrder) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrder) SetCompute(v *CreateHybridClusterRequestEcsOrderCompute) *CreateHybridClusterRequestEcsOrder {
	s.Compute = v
	return s
}

type CreateHybridClusterRequestEcsOrderCompute struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateHybridClusterRequestEcsOrderCompute) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestEcsOrderCompute) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestEcsOrderCompute) SetInstanceType(v string) *CreateHybridClusterRequestEcsOrderCompute {
	s.InstanceType = &v
	return s
}

type CreateHybridClusterRequestNodes struct {
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	IpAddress     *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	HostName      *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Role          *string `json:"Role,omitempty" xml:"Role,omitempty"`
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
}

func (s CreateHybridClusterRequestNodes) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestNodes) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestNodes) SetSchedulerType(v string) *CreateHybridClusterRequestNodes {
	s.SchedulerType = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetIpAddress(v string) *CreateHybridClusterRequestNodes {
	s.IpAddress = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetHostName(v string) *CreateHybridClusterRequestNodes {
	s.HostName = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetRole(v string) *CreateHybridClusterRequestNodes {
	s.Role = &v
	return s
}

func (s *CreateHybridClusterRequestNodes) SetAccountType(v string) *CreateHybridClusterRequestNodes {
	s.AccountType = &v
	return s
}

type CreateHybridClusterRequestApplication struct {
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateHybridClusterRequestApplication) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestApplication) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestApplication) SetTag(v string) *CreateHybridClusterRequestApplication {
	s.Tag = &v
	return s
}

type CreateHybridClusterRequestPostInstallScript struct {
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateHybridClusterRequestPostInstallScript) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterRequestPostInstallScript) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterRequestPostInstallScript) SetArgs(v string) *CreateHybridClusterRequestPostInstallScript {
	s.Args = &v
	return s
}

func (s *CreateHybridClusterRequestPostInstallScript) SetUrl(v string) *CreateHybridClusterRequestPostInstallScript {
	s.Url = &v
	return s
}

type CreateHybridClusterResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s CreateHybridClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterResponseBody) SetTaskId(v string) *CreateHybridClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateHybridClusterResponseBody) SetRequestId(v string) *CreateHybridClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHybridClusterResponseBody) SetClusterId(v string) *CreateHybridClusterResponseBody {
	s.ClusterId = &v
	return s
}

type CreateHybridClusterResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateHybridClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateHybridClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHybridClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridClusterResponse) SetHeaders(v map[string]*string) *CreateHybridClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridClusterResponse) SetBody(v *CreateHybridClusterResponseBody) *CreateHybridClusterResponse {
	s.Body = v
	return s
}

type CreateJobFileRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RunasUser         *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	Content           *string `json:"Content,omitempty" xml:"Content,omitempty"`
	TargetFile        *string `json:"TargetFile,omitempty" xml:"TargetFile,omitempty"`
}

func (s CreateJobFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileRequest) GoString() string {
	return s.String()
}

func (s *CreateJobFileRequest) SetClusterId(v string) *CreateJobFileRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobFileRequest) SetRunasUser(v string) *CreateJobFileRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobFileRequest) SetRunasUserPassword(v string) *CreateJobFileRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *CreateJobFileRequest) SetContent(v string) *CreateJobFileRequest {
	s.Content = &v
	return s
}

func (s *CreateJobFileRequest) SetTargetFile(v string) *CreateJobFileRequest {
	s.TargetFile = &v
	return s
}

type CreateJobFileResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateJobFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobFileResponseBody) SetRequestId(v string) *CreateJobFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobFileResponseBody) SetTemplateId(v string) *CreateJobFileResponseBody {
	s.TemplateId = &v
	return s
}

type CreateJobFileResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobFileResponse) GoString() string {
	return s.String()
}

func (s *CreateJobFileResponse) SetHeaders(v map[string]*string) *CreateJobFileResponse {
	s.Headers = v
	return s
}

func (s *CreateJobFileResponse) SetBody(v *CreateJobFileResponseBody) *CreateJobFileResponse {
	s.Body = v
	return s
}

type CreateJobTemplateRequest struct {
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s CreateJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateRequest) SetCommandLine(v string) *CreateJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *CreateJobTemplateRequest) SetName(v string) *CreateJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *CreateJobTemplateRequest) SetRunasUser(v string) *CreateJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPriority(v int32) *CreateJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobTemplateRequest) SetPackagePath(v string) *CreateJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStdoutRedirectPath(v string) *CreateJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetStderrRedirectPath(v string) *CreateJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *CreateJobTemplateRequest) SetReRunable(v bool) *CreateJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *CreateJobTemplateRequest) SetArrayRequest(v string) *CreateJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *CreateJobTemplateRequest) SetVariables(v string) *CreateJobTemplateRequest {
	s.Variables = &v
	return s
}

type CreateJobTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponseBody) SetRequestId(v string) *CreateJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobTemplateResponseBody) SetTemplateId(v string) *CreateJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateJobTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateJobTemplateResponse) SetHeaders(v map[string]*string) *CreateJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateJobTemplateResponse) SetBody(v *CreateJobTemplateResponseBody) *CreateJobTemplateResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ReleaseInstance *string `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) SetReleaseInstance(v string) *DeleteClusterRequest {
	s.ReleaseInstance = &v
	return s
}

type DeleteClusterResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetTaskId(v string) *DeleteClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteContainerAppsRequest struct {
	ContainerApp []*DeleteContainerAppsRequestContainerApp `json:"ContainerApp,omitempty" xml:"ContainerApp,omitempty" type:"Repeated"`
}

func (s DeleteContainerAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsRequest) SetContainerApp(v []*DeleteContainerAppsRequestContainerApp) *DeleteContainerAppsRequest {
	s.ContainerApp = v
	return s
}

type DeleteContainerAppsRequestContainerApp struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteContainerAppsRequestContainerApp) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsRequestContainerApp) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsRequestContainerApp) SetId(v string) *DeleteContainerAppsRequestContainerApp {
	s.Id = &v
	return s
}

type DeleteContainerAppsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContainerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsResponseBody) SetRequestId(v string) *DeleteContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteContainerAppsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteContainerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *DeleteContainerAppsResponse) SetHeaders(v map[string]*string) *DeleteContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *DeleteContainerAppsResponse) SetBody(v *DeleteContainerAppsResponseBody) *DeleteContainerAppsResponse {
	s.Body = v
	return s
}

type DeleteGWSClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteGWSClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterRequest) SetClusterId(v string) *DeleteGWSClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteGWSClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGWSClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterResponseBody) SetRequestId(v string) *DeleteGWSClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGWSClusterResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGWSClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGWSClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteGWSClusterResponse) SetHeaders(v map[string]*string) *DeleteGWSClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteGWSClusterResponse) SetBody(v *DeleteGWSClusterResponseBody) *DeleteGWSClusterResponse {
	s.Body = v
	return s
}

type DeleteGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceRequest) SetInstanceId(v string) *DeleteGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type DeleteGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceResponseBody) SetRequestId(v string) *DeleteGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGWSInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGWSInstanceResponse) SetHeaders(v map[string]*string) *DeleteGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGWSInstanceResponse) SetBody(v *DeleteGWSInstanceResponseBody) *DeleteGWSInstanceResponse {
	s.Body = v
	return s
}

type DeleteImageRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetClusterId(v string) *DeleteImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteImageRequest) SetRepository(v string) *DeleteImageRequest {
	s.Repository = &v
	return s
}

func (s *DeleteImageRequest) SetImageTag(v string) *DeleteImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DeleteImageRequest) SetContainerType(v string) *DeleteImageRequest {
	s.ContainerType = &v
	return s
}

type DeleteImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

func (s *DeleteImageResponse) SetBody(v *DeleteImageResponseBody) *DeleteImageResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetClusterId(v string) *DeleteJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteJobsRequest) SetJobs(v string) *DeleteJobsRequest {
	s.Jobs = &v
	return s
}

type DeleteJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
	s.Body = v
	return s
}

type DeleteJobTemplatesRequest struct {
	Templates *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s DeleteJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesRequest) SetTemplates(v string) *DeleteJobTemplatesRequest {
	s.Templates = &v
	return s
}

type DeleteJobTemplatesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponseBody) SetRequestId(v string) *DeleteJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobTemplatesResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobTemplatesResponse) SetHeaders(v map[string]*string) *DeleteJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobTemplatesResponse) SetBody(v *DeleteJobTemplatesResponseBody) *DeleteJobTemplatesResponse {
	s.Body = v
	return s
}

type DeleteNodesRequest struct {
	ClusterId       *string                       `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ReleaseInstance *bool                         `json:"ReleaseInstance,omitempty" xml:"ReleaseInstance,omitempty"`
	Instance        []*DeleteNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DeleteNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequest) SetClusterId(v string) *DeleteNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodesRequest) SetReleaseInstance(v bool) *DeleteNodesRequest {
	s.ReleaseInstance = &v
	return s
}

func (s *DeleteNodesRequest) SetInstance(v []*DeleteNodesRequestInstance) *DeleteNodesRequest {
	s.Instance = v
	return s
}

type DeleteNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *DeleteNodesRequestInstance) SetId(v string) *DeleteNodesRequestInstance {
	s.Id = &v
	return s
}

type DeleteNodesResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponseBody) SetTaskId(v string) *DeleteNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *DeleteNodesResponseBody) SetRequestId(v string) *DeleteNodesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNodesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodesResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodesResponse) SetHeaders(v map[string]*string) *DeleteNodesResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodesResponse) SetBody(v *DeleteNodesResponseBody) *DeleteNodesResponse {
	s.Body = v
	return s
}

type DeleteQueueRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
}

func (s DeleteQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteQueueRequest) SetClusterId(v string) *DeleteQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteQueueRequest) SetQueueName(v string) *DeleteQueueRequest {
	s.QueueName = &v
	return s
}

type DeleteQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponseBody) SetRequestId(v string) *DeleteQueueResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQueueResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQueueResponse) GoString() string {
	return s.String()
}

func (s *DeleteQueueResponse) SetHeaders(v map[string]*string) *DeleteQueueResponse {
	s.Headers = v
	return s
}

func (s *DeleteQueueResponse) SetBody(v *DeleteQueueResponseBody) *DeleteQueueResponse {
	s.Body = v
	return s
}

type DeleteSecurityGroupRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupRequest) SetClusterId(v string) *DeleteSecurityGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteSecurityGroupRequest) SetSecurityGroupId(v string) *DeleteSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

type DeleteSecurityGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponseBody) SetRequestId(v string) *DeleteSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSecurityGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityGroupResponse) SetBody(v *DeleteSecurityGroupResponseBody) *DeleteSecurityGroupResponse {
	s.Body = v
	return s
}

type DeleteUsersRequest struct {
	ClusterId *string                   `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*DeleteUsersRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DeleteUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequest) SetClusterId(v string) *DeleteUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteUsersRequest) SetUser(v []*DeleteUsersRequestUser) *DeleteUsersRequest {
	s.User = v
	return s
}

type DeleteUsersRequestUser struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteUsersRequestUser) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersRequestUser) GoString() string {
	return s.String()
}

func (s *DeleteUsersRequestUser) SetName(v string) *DeleteUsersRequestUser {
	s.Name = &v
	return s
}

type DeleteUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponseBody) SetRequestId(v string) *DeleteUsersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteUsersResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteUsersResponse) SetHeaders(v map[string]*string) *DeleteUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteUsersResponse) SetBody(v *DeleteUsersResponseBody) *DeleteUsersResponse {
	s.Body = v
	return s
}

type DescribeAutoScaleConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigRequest) SetClusterId(v string) *DescribeAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeAutoScaleConfigResponseBody struct {
	ExtraNodesGrowRatio     *int32  `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	RequestId               *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnableAutoGrow          *bool   `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	ClusterId               *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxNodesInCluster       *int32  `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	ShrinkIdleTimes         *int32  `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	EnableAutoShrink        *bool   `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ClusterType             *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	GrowRatio               *int32  `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32  `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	Uid                     *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	GrowTimeoutInMinutes    *int32  `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	ShrinkIntervalInMinutes *int32  `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	SpotPriceLimit          *string `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	ExcludeNodes            *string `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	SpotStrategy            *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s DescribeAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetRequestId(v string) *DescribeAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *DescribeAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetClusterId(v string) *DescribeAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *DescribeAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *DescribeAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetClusterType(v string) *DescribeAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowRatio(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetUid(v string) *DescribeAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *DescribeAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetSpotPriceLimit(v string) *DescribeAutoScaleConfigResponseBody {
	s.SpotPriceLimit = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetExcludeNodes(v string) *DescribeAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *DescribeAutoScaleConfigResponseBody) SetSpotStrategy(v string) *DescribeAutoScaleConfigResponseBody {
	s.SpotStrategy = &v
	return s
}

type DescribeAutoScaleConfigResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoScaleConfigResponse) SetHeaders(v map[string]*string) *DescribeAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoScaleConfigResponse) SetBody(v *DescribeAutoScaleConfigResponseBody) *DescribeAutoScaleConfigResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	ClusterInfo *DescribeClusterResponseBodyClusterInfo `json:"ClusterInfo,omitempty" xml:"ClusterInfo,omitempty" type:"Struct"`
	RequestId   *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterInfo(v *DescribeClusterResponseBodyClusterInfo) *DescribeClusterResponseBody {
	s.ClusterInfo = v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

type DescribeClusterResponseBodyClusterInfo struct {
	Status             *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId              *string                                                     `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	KeyPairName        *string                                                     `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	EcsChargeType      *string                                                     `json:"EcsChargeType,omitempty" xml:"EcsChargeType,omitempty"`
	SecurityGroupId    *string                                                     `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SccClusterId       *string                                                     `json:"SccClusterId,omitempty" xml:"SccClusterId,omitempty"`
	CreateTime         *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AccountType        *string                                                     `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	VolumeProtocol     *string                                                     `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	Description        *string                                                     `json:"Description,omitempty" xml:"Description,omitempty"`
	VolumeId           *string                                                     `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	HaEnable           *bool                                                       `json:"HaEnable,omitempty" xml:"HaEnable,omitempty"`
	BaseOsTag          *string                                                     `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	Name               *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId            *string                                                     `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	PostInstallScripts []*DescribeClusterResponseBodyClusterInfoPostInstallScripts `json:"PostInstallScripts,omitempty" xml:"PostInstallScripts,omitempty" type:"Repeated"`
	SchedulerType      *string                                                     `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	DeployMode         *string                                                     `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	ImageOwnerAlias    *string                                                     `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	RemoteDirectory    *string                                                     `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeMountpoint   *string                                                     `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	OsTag              *string                                                     `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	RegionId           *string                                                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId          *string                                                     `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	EcsInfo            *DescribeClusterResponseBodyClusterInfoEcsInfo              `json:"EcsInfo,omitempty" xml:"EcsInfo,omitempty" type:"Struct"`
	ImageName          *string                                                     `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	Applications       []*DescribeClusterResponseBodyClusterInfoApplications       `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	VolumeType         *string                                                     `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	Location           *string                                                     `json:"Location,omitempty" xml:"Location,omitempty"`
	Id                 *string                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	ClientVersion      *string                                                     `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfo) SetStatus(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Status = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVpcId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VpcId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetKeyPairName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.KeyPairName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsChargeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.EcsChargeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSecurityGroupId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSccClusterId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SccClusterId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetCreateTime(v string) *DescribeClusterResponseBodyClusterInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetAccountType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.AccountType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeProtocol(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeProtocol = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDescription(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Description = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetHaEnable(v bool) *DescribeClusterResponseBodyClusterInfo {
	s.HaEnable = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetBaseOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.BaseOsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetPostInstallScripts(v []*DescribeClusterResponseBodyClusterInfoPostInstallScripts) *DescribeClusterResponseBodyClusterInfo {
	s.PostInstallScripts = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetSchedulerType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.SchedulerType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetDeployMode(v string) *DescribeClusterResponseBodyClusterInfo {
	s.DeployMode = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageOwnerAlias(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageOwnerAlias = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRemoteDirectory(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RemoteDirectory = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeMountpoint(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeMountpoint = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetOsTag(v string) *DescribeClusterResponseBodyClusterInfo {
	s.OsTag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetRegionId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.RegionId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVSwitchId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetEcsInfo(v *DescribeClusterResponseBodyClusterInfoEcsInfo) *DescribeClusterResponseBodyClusterInfo {
	s.EcsInfo = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetImageName(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ImageName = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetApplications(v []*DescribeClusterResponseBodyClusterInfoApplications) *DescribeClusterResponseBodyClusterInfo {
	s.Applications = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetVolumeType(v string) *DescribeClusterResponseBodyClusterInfo {
	s.VolumeType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetLocation(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Location = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetId(v string) *DescribeClusterResponseBodyClusterInfo {
	s.Id = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfo) SetClientVersion(v string) *DescribeClusterResponseBodyClusterInfo {
	s.ClientVersion = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoPostInstallScripts struct {
	Args *string `json:"Args,omitempty" xml:"Args,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScripts) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoPostInstallScripts) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScripts) SetArgs(v string) *DescribeClusterResponseBodyClusterInfoPostInstallScripts {
	s.Args = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoPostInstallScripts) SetUrl(v string) *DescribeClusterResponseBodyClusterInfoPostInstallScripts {
	s.Url = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfo struct {
	Manager *DescribeClusterResponseBodyClusterInfoEcsInfoManager `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Struct"`
	Compute *DescribeClusterResponseBodyClusterInfoEcsInfoCompute `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Struct"`
	Login   *DescribeClusterResponseBodyClusterInfoEcsInfoLogin   `json:"Login,omitempty" xml:"Login,omitempty" type:"Struct"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetManager(v *DescribeClusterResponseBodyClusterInfoEcsInfoManager) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Manager = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetCompute(v *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Compute = v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfo) SetLogin(v *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) *DescribeClusterResponseBodyClusterInfoEcsInfo {
	s.Login = v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoManager struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoManager) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoManager) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoManager {
	s.Count = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoCompute struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoCompute) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoCompute) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoCompute {
	s.Count = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoEcsInfoLogin struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Count        *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoEcsInfoLogin) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetInstanceType(v string) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.InstanceType = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoEcsInfoLogin) SetCount(v int32) *DescribeClusterResponseBodyClusterInfoEcsInfoLogin {
	s.Count = &v
	return s
}

type DescribeClusterResponseBodyClusterInfoApplications struct {
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Tag     *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeClusterResponseBodyClusterInfoApplications) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyClusterInfoApplications) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyClusterInfoApplications) SetVersion(v string) *DescribeClusterResponseBodyClusterInfoApplications {
	s.Version = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplications) SetTag(v string) *DescribeClusterResponseBodyClusterInfoApplications {
	s.Tag = &v
	return s
}

func (s *DescribeClusterResponseBodyClusterInfoApplications) SetName(v string) *DescribeClusterResponseBodyClusterInfoApplications {
	s.Name = &v
	return s
}

type DescribeClusterResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeContainerAppRequest struct {
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
}

func (s DescribeContainerAppRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppRequest) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppRequest) SetContainerId(v string) *DescribeContainerAppRequest {
	s.ContainerId = &v
	return s
}

type DescribeContainerAppResponseBody struct {
	ContainerAppInfo *DescribeContainerAppResponseBodyContainerAppInfo `json:"ContainerAppInfo,omitempty" xml:"ContainerAppInfo,omitempty" type:"Struct"`
	RequestId        *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContainerAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponseBody) SetContainerAppInfo(v *DescribeContainerAppResponseBodyContainerAppInfo) *DescribeContainerAppResponseBody {
	s.ContainerAppInfo = v
	return s
}

func (s *DescribeContainerAppResponseBody) SetRequestId(v string) *DescribeContainerAppResponseBody {
	s.RequestId = &v
	return s
}

type DescribeContainerAppResponseBodyContainerAppInfo struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Repository  *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	ImageTag    *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DescribeContainerAppResponseBodyContainerAppInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponseBodyContainerAppInfo) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetType(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Type = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetDescription(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Description = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetCreateTime(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.CreateTime = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetRepository(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Repository = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetImageTag(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.ImageTag = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetName(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Name = &v
	return s
}

func (s *DescribeContainerAppResponseBodyContainerAppInfo) SetId(v string) *DescribeContainerAppResponseBodyContainerAppInfo {
	s.Id = &v
	return s
}

type DescribeContainerAppResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeContainerAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeContainerAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeContainerAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppResponse) SetHeaders(v map[string]*string) *DescribeContainerAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerAppResponse) SetBody(v *DescribeContainerAppResponseBody) *DescribeContainerAppResponse {
	s.Body = v
	return s
}

type DescribeGWSClusterPolicyRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	AsyncMode *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
}

func (s DescribeGWSClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyRequest) SetClusterId(v string) *DescribeGWSClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClusterPolicyRequest) SetTaskId(v string) *DescribeGWSClusterPolicyRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeGWSClusterPolicyRequest) SetAsyncMode(v bool) *DescribeGWSClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

type DescribeGWSClusterPolicyResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s DescribeGWSClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyResponseBody) SetRequestId(v string) *DescribeGWSClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetLocalDrive(v string) *DescribeGWSClusterPolicyResponseBody {
	s.LocalDrive = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetUsbRedirect(v string) *DescribeGWSClusterPolicyResponseBody {
	s.UsbRedirect = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetClipboard(v string) *DescribeGWSClusterPolicyResponseBody {
	s.Clipboard = &v
	return s
}

func (s *DescribeGWSClusterPolicyResponseBody) SetWatermark(v string) *DescribeGWSClusterPolicyResponseBody {
	s.Watermark = &v
	return s
}

type DescribeGWSClusterPolicyResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSClusterPolicyResponse) SetHeaders(v map[string]*string) *DescribeGWSClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSClusterPolicyResponse) SetBody(v *DescribeGWSClusterPolicyResponseBody) *DescribeGWSClusterPolicyResponse {
	s.Body = v
	return s
}

type DescribeGWSClustersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGWSClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersRequest) SetClusterId(v string) *DescribeGWSClustersRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSClustersRequest) SetPageNumber(v int32) *DescribeGWSClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSClustersRequest) SetPageSize(v int32) *DescribeGWSClustersRequest {
	s.PageSize = &v
	return s
}

type DescribeGWSClustersResponseBody struct {
	TotalCount *int32                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                     `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Clusters   []*DescribeGWSClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	CallerType *string                                    `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
}

func (s DescribeGWSClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBody) SetTotalCount(v int32) *DescribeGWSClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetPageSize(v int32) *DescribeGWSClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetRequestId(v string) *DescribeGWSClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetPageNumber(v int32) *DescribeGWSClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetClusters(v []*DescribeGWSClustersResponseBodyClusters) *DescribeGWSClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeGWSClustersResponseBody) SetCallerType(v string) *DescribeGWSClustersResponseBody {
	s.CallerType = &v
	return s
}

type DescribeGWSClustersResponseBodyClusters struct {
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceCount *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeGWSClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponseBodyClusters) SetVpcId(v string) *DescribeGWSClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClusters) SetStatus(v string) *DescribeGWSClustersResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClusters) SetInstanceCount(v int32) *DescribeGWSClustersResponseBodyClusters {
	s.InstanceCount = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClusters) SetCreateTime(v string) *DescribeGWSClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSClustersResponseBodyClusters) SetClusterId(v string) *DescribeGWSClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

type DescribeGWSClustersResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSClustersResponse) SetHeaders(v map[string]*string) *DescribeGWSClustersResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSClustersResponse) SetBody(v *DescribeGWSClustersResponseBody) *DescribeGWSClustersResponse {
	s.Body = v
	return s
}

type DescribeGWSImagesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeGWSImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesRequest) SetPageNumber(v int32) *DescribeGWSImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSImagesRequest) SetPageSize(v int32) *DescribeGWSImagesRequest {
	s.PageSize = &v
	return s
}

type DescribeGWSImagesResponseBody struct {
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Images     []*DescribeGWSImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s DescribeGWSImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBody) SetTotalCount(v int32) *DescribeGWSImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetPageSize(v int32) *DescribeGWSImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetRequestId(v string) *DescribeGWSImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetPageNumber(v int32) *DescribeGWSImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSImagesResponseBody) SetImages(v []*DescribeGWSImagesResponseBodyImages) *DescribeGWSImagesResponseBody {
	s.Images = v
	return s
}

type DescribeGWSImagesResponseBodyImages struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ImageType  *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Progress   *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Size       *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId    *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeGWSImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponseBodyImages) SetStatus(v string) *DescribeGWSImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImages) SetImageType(v string) *DescribeGWSImagesResponseBodyImages {
	s.ImageType = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImages) SetProgress(v string) *DescribeGWSImagesResponseBodyImages {
	s.Progress = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImages) SetSize(v int32) *DescribeGWSImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImages) SetCreateTime(v string) *DescribeGWSImagesResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImages) SetName(v string) *DescribeGWSImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *DescribeGWSImagesResponseBodyImages) SetImageId(v string) *DescribeGWSImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

type DescribeGWSImagesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSImagesResponse) SetHeaders(v map[string]*string) *DescribeGWSImagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSImagesResponse) SetBody(v *DescribeGWSImagesResponseBody) *DescribeGWSImagesResponse {
	s.Body = v
	return s
}

type DescribeGWSInstancesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserUid    *int64  `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeGWSInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesRequest) SetClusterId(v string) *DescribeGWSInstancesRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetInstanceId(v string) *DescribeGWSInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetPageNumber(v int32) *DescribeGWSInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetPageSize(v int32) *DescribeGWSInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetUserUid(v int64) *DescribeGWSInstancesRequest {
	s.UserUid = &v
	return s
}

func (s *DescribeGWSInstancesRequest) SetUserName(v string) *DescribeGWSInstancesRequest {
	s.UserName = &v
	return s
}

type DescribeGWSInstancesResponseBody struct {
	Instances  []*DescribeGWSInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	TotalCount *int32                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                       `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeGWSInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBody) SetInstances(v []*DescribeGWSInstancesResponseBodyInstances) *DescribeGWSInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetTotalCount(v int32) *DescribeGWSInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetPageSize(v int32) *DescribeGWSInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetRequestId(v string) *DescribeGWSInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBody) SetPageNumber(v int32) *DescribeGWSInstancesResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeGWSInstancesResponseBodyInstances struct {
	Status       *string                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	AppList      []*DescribeGWSInstancesResponseBodyInstancesAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Repeated"`
	WorkMode     *string                                             `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
	ExpireTime   *string                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime   *string                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	InstanceId   *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name         *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	InstanceType *string                                             `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName     *string                                             `json:"UserName,omitempty" xml:"UserName,omitempty"`
	ClusterId    *string                                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeGWSInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetStatus(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetAppList(v []*DescribeGWSInstancesResponseBodyInstancesAppList) *DescribeGWSInstancesResponseBodyInstances {
	s.AppList = v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetWorkMode(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.WorkMode = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetExpireTime(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetCreateTime(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetName(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.Name = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetInstanceType(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetUserName(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.UserName = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstances) SetClusterId(v string) *DescribeGWSInstancesResponseBodyInstances {
	s.ClusterId = &v
	return s
}

type DescribeGWSInstancesResponseBodyInstancesAppList struct {
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPath *string `json:"AppPath,omitempty" xml:"AppPath,omitempty"`
	AppArgs *string `json:"AppArgs,omitempty" xml:"AppArgs,omitempty"`
}

func (s DescribeGWSInstancesResponseBodyInstancesAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponseBodyInstancesAppList) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponseBodyInstancesAppList) SetAppName(v string) *DescribeGWSInstancesResponseBodyInstancesAppList {
	s.AppName = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesAppList) SetAppPath(v string) *DescribeGWSInstancesResponseBodyInstancesAppList {
	s.AppPath = &v
	return s
}

func (s *DescribeGWSInstancesResponseBodyInstancesAppList) SetAppArgs(v string) *DescribeGWSInstancesResponseBodyInstancesAppList {
	s.AppArgs = &v
	return s
}

type DescribeGWSInstancesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeGWSInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeGWSInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeGWSInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeGWSInstancesResponse) SetHeaders(v map[string]*string) *DescribeGWSInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeGWSInstancesResponse) SetBody(v *DescribeGWSInstancesResponseBody) *DescribeGWSInstancesResponse {
	s.Body = v
	return s
}

type DescribeImageRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
}

func (s DescribeImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageRequest) SetClusterId(v string) *DescribeImageRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeImageRequest) SetRepository(v string) *DescribeImageRequest {
	s.Repository = &v
	return s
}

func (s *DescribeImageRequest) SetImageTag(v string) *DescribeImageRequest {
	s.ImageTag = &v
	return s
}

func (s *DescribeImageRequest) SetContainerType(v string) *DescribeImageRequest {
	s.ContainerType = &v
	return s
}

type DescribeImageResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ImageInfo *DescribeImageResponseBodyImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
}

func (s DescribeImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBody) SetRequestId(v string) *DescribeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageResponseBody) SetImageInfo(v *DescribeImageResponseBodyImageInfo) *DescribeImageResponseBody {
	s.ImageInfo = v
	return s
}

type DescribeImageResponseBodyImageInfo struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateDateTime *string `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
	Repository     *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	System         *string `json:"System,omitempty" xml:"System,omitempty"`
	ImageId        *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s DescribeImageResponseBodyImageInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponseBodyImageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageResponseBodyImageInfo) SetType(v string) *DescribeImageResponseBodyImageInfo {
	s.Type = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetStatus(v string) *DescribeImageResponseBodyImageInfo {
	s.Status = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetUpdateDateTime(v string) *DescribeImageResponseBodyImageInfo {
	s.UpdateDateTime = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetRepository(v string) *DescribeImageResponseBodyImageInfo {
	s.Repository = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetTag(v string) *DescribeImageResponseBodyImageInfo {
	s.Tag = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetSystem(v string) *DescribeImageResponseBodyImageInfo {
	s.System = &v
	return s
}

func (s *DescribeImageResponseBodyImageInfo) SetImageId(v string) *DescribeImageResponseBodyImageInfo {
	s.ImageId = &v
	return s
}

type DescribeImageResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageResponse) SetHeaders(v map[string]*string) *DescribeImageResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageResponse) SetBody(v *DescribeImageResponseBody) *DescribeImageResponse {
	s.Body = v
	return s
}

type DescribeImageGatewayConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeImageGatewayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigRequest) SetClusterId(v string) *DescribeImageGatewayConfigRequest {
	s.ClusterId = &v
	return s
}

type DescribeImageGatewayConfigResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Imagegw   *DescribeImageGatewayConfigResponseBodyImagegw `json:"Imagegw,omitempty" xml:"Imagegw,omitempty" type:"Struct"`
}

func (s DescribeImageGatewayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBody) SetRequestId(v string) *DescribeImageGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBody) SetImagegw(v *DescribeImageGatewayConfigResponseBodyImagegw) *DescribeImageGatewayConfigResponseBody {
	s.Imagegw = v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegw struct {
	Locations              []*DescribeImageGatewayConfigResponseBodyImagegwLocations `json:"Locations,omitempty" xml:"Locations,omitempty" type:"Repeated"`
	UpdateDateTime         *string                                                   `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
	ImageExpirationTimeout *string                                                   `json:"ImageExpirationTimeout,omitempty" xml:"ImageExpirationTimeout,omitempty"`
	MongoDBURI             *string                                                   `json:"MongoDBURI,omitempty" xml:"MongoDBURI,omitempty"`
	DefaultImageLocation   *string                                                   `json:"DefaultImageLocation,omitempty" xml:"DefaultImageLocation,omitempty"`
	PullUpdateTimeout      *int64                                                    `json:"PullUpdateTimeout,omitempty" xml:"PullUpdateTimeout,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegw) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegw) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetLocations(v []*DescribeImageGatewayConfigResponseBodyImagegwLocations) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.Locations = v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetUpdateDateTime(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.UpdateDateTime = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetImageExpirationTimeout(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.ImageExpirationTimeout = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetMongoDBURI(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.MongoDBURI = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetDefaultImageLocation(v string) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.DefaultImageLocation = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegw) SetPullUpdateTimeout(v int64) *DescribeImageGatewayConfigResponseBodyImagegw {
	s.PullUpdateTimeout = &v
	return s
}

type DescribeImageGatewayConfigResponseBodyImagegwLocations struct {
	RemoteType     *string `json:"RemoteType,omitempty" xml:"RemoteType,omitempty"`
	URL            *string `json:"URL,omitempty" xml:"URL,omitempty"`
	Location       *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Authentication *string `json:"Authentication,omitempty" xml:"Authentication,omitempty"`
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponseBodyImagegwLocations) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocations) SetRemoteType(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocations {
	s.RemoteType = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocations) SetURL(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocations {
	s.URL = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocations) SetLocation(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocations {
	s.Location = &v
	return s
}

func (s *DescribeImageGatewayConfigResponseBodyImagegwLocations) SetAuthentication(v string) *DescribeImageGatewayConfigResponseBodyImagegwLocations {
	s.Authentication = &v
	return s
}

type DescribeImageGatewayConfigResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImageGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImageGatewayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImageGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeImageGatewayConfigResponse) SetHeaders(v map[string]*string) *DescribeImageGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeImageGatewayConfigResponse) SetBody(v *DescribeImageGatewayConfigResponseBody) *DescribeImageGatewayConfigResponse {
	s.Body = v
	return s
}

type DescribeImagePriceRequest struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	SkuCode   *string `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	Period    *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	Amount    *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribeImagePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceRequest) SetImageId(v string) *DescribeImagePriceRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePriceRequest) SetPriceUnit(v string) *DescribeImagePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribeImagePriceRequest) SetSkuCode(v string) *DescribeImagePriceRequest {
	s.SkuCode = &v
	return s
}

func (s *DescribeImagePriceRequest) SetPeriod(v int32) *DescribeImagePriceRequest {
	s.Period = &v
	return s
}

func (s *DescribeImagePriceRequest) SetAmount(v int32) *DescribeImagePriceRequest {
	s.Amount = &v
	return s
}

func (s *DescribeImagePriceRequest) SetOrderType(v string) *DescribeImagePriceRequest {
	s.OrderType = &v
	return s
}

type DescribeImagePriceResponseBody struct {
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RequestId     *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Amount        *int32   `json:"Amount,omitempty" xml:"Amount,omitempty"`
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ImageId       *string  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeImagePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceResponseBody) SetOriginalPrice(v float32) *DescribeImagePriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetRequestId(v string) *DescribeImagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetAmount(v int32) *DescribeImagePriceResponseBody {
	s.Amount = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetDiscountPrice(v float32) *DescribeImagePriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetImageId(v string) *DescribeImagePriceResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImagePriceResponseBody) SetTradePrice(v float32) *DescribeImagePriceResponseBody {
	s.TradePrice = &v
	return s
}

type DescribeImagePriceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeImagePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeImagePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagePriceResponse) SetHeaders(v map[string]*string) *DescribeImagePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribeImagePriceResponse) SetBody(v *DescribeImagePriceResponseBody) *DescribeImagePriceResponse {
	s.Body = v
	return s
}

type DescribeJobRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobRequest) SetClusterId(v string) *DescribeJobRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeJobRequest) SetJobId(v string) *DescribeJobRequest {
	s.JobId = &v
	return s
}

type DescribeJobResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message   *DescribeJobResponseBodyMessage `json:"Message,omitempty" xml:"Message,omitempty" type:"Struct"`
}

func (s DescribeJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBody) SetRequestId(v string) *DescribeJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobResponseBody) SetMessage(v *DescribeJobResponseBodyMessage) *DescribeJobResponseBody {
	s.Message = v
	return s
}

type DescribeJobResponseBodyMessage struct {
	JobInfo *string `json:"JobInfo,omitempty" xml:"JobInfo,omitempty"`
}

func (s DescribeJobResponseBodyMessage) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponseBodyMessage) GoString() string {
	return s.String()
}

func (s *DescribeJobResponseBodyMessage) SetJobInfo(v string) *DescribeJobResponseBodyMessage {
	s.JobInfo = &v
	return s
}

type DescribeJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobResponse) SetHeaders(v map[string]*string) *DescribeJobResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobResponse) SetBody(v *DescribeJobResponseBody) *DescribeJobResponse {
	s.Body = v
	return s
}

type DescribeNFSClientStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeNFSClientStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusRequest) SetInstanceId(v string) *DescribeNFSClientStatusRequest {
	s.InstanceId = &v
	return s
}

type DescribeNFSClientStatusResponseBody struct {
	Status    *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *DescribeNFSClientStatusResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeNFSClientStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponseBody) SetStatus(v string) *DescribeNFSClientStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBody) SetRequestId(v string) *DescribeNFSClientStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBody) SetResult(v *DescribeNFSClientStatusResponseBodyResult) *DescribeNFSClientStatusResponseBody {
	s.Result = v
	return s
}

type DescribeNFSClientStatusResponseBodyResult struct {
	Output             *string `json:"Output,omitempty" xml:"Output,omitempty"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	ExitCode           *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s DescribeNFSClientStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetOutput(v string) *DescribeNFSClientStatusResponseBodyResult {
	s.Output = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetInvokeRecordStatus(v string) *DescribeNFSClientStatusResponseBodyResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeNFSClientStatusResponseBodyResult) SetExitCode(v int32) *DescribeNFSClientStatusResponseBodyResult {
	s.ExitCode = &v
	return s
}

type DescribeNFSClientStatusResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeNFSClientStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNFSClientStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNFSClientStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeNFSClientStatusResponse) SetHeaders(v map[string]*string) *DescribeNFSClientStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeNFSClientStatusResponse) SetBody(v *DescribeNFSClientStatusResponseBody) *DescribeNFSClientStatusResponse {
	s.Body = v
	return s
}

type DescribePriceRequest struct {
	PriceUnit   *string                            `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	ChargeType  *string                            `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	OrderType   *string                            `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Commodities []*DescribePriceRequestCommodities `json:"Commodities,omitempty" xml:"Commodities,omitempty" type:"Repeated"`
}

func (s DescribePriceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) SetPriceUnit(v string) *DescribePriceRequest {
	s.PriceUnit = &v
	return s
}

func (s *DescribePriceRequest) SetChargeType(v string) *DescribePriceRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) SetCommodities(v []*DescribePriceRequestCommodities) *DescribePriceRequest {
	s.Commodities = v
	return s
}

type DescribePriceRequestCommodities struct {
	Amount                  *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	SystemDiskSize          *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
	NodeType                *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	SystemDiskCategory      *string `json:"SystemDiskCategory,omitempty" xml:"SystemDiskCategory,omitempty"`
	InternetChargeType      *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	NetworkType             *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	InstanceType            *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Period                  *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	InternetMaxBandWidthOut *int32  `json:"InternetMaxBandWidthOut,omitempty" xml:"InternetMaxBandWidthOut,omitempty"`
}

func (s DescribePriceRequestCommodities) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceRequestCommodities) GoString() string {
	return s.String()
}

func (s *DescribePriceRequestCommodities) SetAmount(v int32) *DescribePriceRequestCommodities {
	s.Amount = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskSize(v int32) *DescribePriceRequestCommodities {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetNodeType(v string) *DescribePriceRequestCommodities {
	s.NodeType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetSystemDiskCategory(v string) *DescribePriceRequestCommodities {
	s.SystemDiskCategory = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInternetChargeType(v string) *DescribePriceRequestCommodities {
	s.InternetChargeType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetNetworkType(v string) *DescribePriceRequestCommodities {
	s.NetworkType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInstanceType(v string) *DescribePriceRequestCommodities {
	s.InstanceType = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetPeriod(v int32) *DescribePriceRequestCommodities {
	s.Period = &v
	return s
}

func (s *DescribePriceRequestCommodities) SetInternetMaxBandWidthOut(v int32) *DescribePriceRequestCommodities {
	s.InternetMaxBandWidthOut = &v
	return s
}

type DescribePriceResponseBody struct {
	Prices          []*DescribePriceResponseBodyPrices `json:"Prices,omitempty" xml:"Prices,omitempty" type:"Repeated"`
	TotalTradePrice *float32                           `json:"TotalTradePrice,omitempty" xml:"TotalTradePrice,omitempty"`
	RequestId       *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) SetPrices(v []*DescribePriceResponseBodyPrices) *DescribePriceResponseBody {
	s.Prices = v
	return s
}

func (s *DescribePriceResponseBody) SetTotalTradePrice(v float32) *DescribePriceResponseBody {
	s.TotalTradePrice = &v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

type DescribePriceResponseBodyPrices struct {
	NodeType      *string  `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	TradePrice    *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Currency      *string  `json:"Currency,omitempty" xml:"Currency,omitempty"`
}

func (s DescribePriceResponseBodyPrices) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponseBodyPrices) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPrices) SetNodeType(v string) *DescribePriceResponseBodyPrices {
	s.NodeType = &v
	return s
}

func (s *DescribePriceResponseBodyPrices) SetTradePrice(v float32) *DescribePriceResponseBodyPrices {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPrices) SetOriginalPrice(v float32) *DescribePriceResponseBodyPrices {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPrices) SetCurrency(v string) *DescribePriceResponseBodyPrices {
	s.Currency = &v
	return s
}

type DescribePriceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePriceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePriceResponse) GoString() string {
	return s.String()
}

func (s *DescribePriceResponse) SetHeaders(v map[string]*string) *DescribePriceResponse {
	s.Headers = v
	return s
}

func (s *DescribePriceResponse) SetBody(v *DescribePriceResponseBody) *DescribePriceResponse {
	s.Body = v
	return s
}

type EditJobTemplateRequest struct {
	TemplateId         *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s EditJobTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *EditJobTemplateRequest) SetTemplateId(v string) *EditJobTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *EditJobTemplateRequest) SetCommandLine(v string) *EditJobTemplateRequest {
	s.CommandLine = &v
	return s
}

func (s *EditJobTemplateRequest) SetName(v string) *EditJobTemplateRequest {
	s.Name = &v
	return s
}

func (s *EditJobTemplateRequest) SetRunasUser(v string) *EditJobTemplateRequest {
	s.RunasUser = &v
	return s
}

func (s *EditJobTemplateRequest) SetPriority(v int32) *EditJobTemplateRequest {
	s.Priority = &v
	return s
}

func (s *EditJobTemplateRequest) SetPackagePath(v string) *EditJobTemplateRequest {
	s.PackagePath = &v
	return s
}

func (s *EditJobTemplateRequest) SetStdoutRedirectPath(v string) *EditJobTemplateRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetStderrRedirectPath(v string) *EditJobTemplateRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *EditJobTemplateRequest) SetReRunable(v bool) *EditJobTemplateRequest {
	s.ReRunable = &v
	return s
}

func (s *EditJobTemplateRequest) SetArrayRequest(v string) *EditJobTemplateRequest {
	s.ArrayRequest = &v
	return s
}

func (s *EditJobTemplateRequest) SetVariables(v string) *EditJobTemplateRequest {
	s.Variables = &v
	return s
}

type EditJobTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s EditJobTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponseBody) SetRequestId(v string) *EditJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditJobTemplateResponseBody) SetTemplateId(v string) *EditJobTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type EditJobTemplateResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditJobTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditJobTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s EditJobTemplateResponse) GoString() string {
	return s.String()
}

func (s *EditJobTemplateResponse) SetHeaders(v map[string]*string) *EditJobTemplateResponse {
	s.Headers = v
	return s
}

func (s *EditJobTemplateResponse) SetBody(v *EditJobTemplateResponseBody) *EditJobTemplateResponse {
	s.Body = v
	return s
}

type GetAccountingReportRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	StartTime   *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ReportType  *string `json:"ReportType,omitempty" xml:"ReportType,omitempty"`
	FilterValue *string `json:"FilterValue,omitempty" xml:"FilterValue,omitempty"`
	Dim         *string `json:"Dim,omitempty" xml:"Dim,omitempty"`
	JobId       *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber  *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s GetAccountingReportRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportRequest) GoString() string {
	return s.String()
}

func (s *GetAccountingReportRequest) SetClusterId(v string) *GetAccountingReportRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAccountingReportRequest) SetStartTime(v int32) *GetAccountingReportRequest {
	s.StartTime = &v
	return s
}

func (s *GetAccountingReportRequest) SetEndTime(v int32) *GetAccountingReportRequest {
	s.EndTime = &v
	return s
}

func (s *GetAccountingReportRequest) SetReportType(v string) *GetAccountingReportRequest {
	s.ReportType = &v
	return s
}

func (s *GetAccountingReportRequest) SetFilterValue(v string) *GetAccountingReportRequest {
	s.FilterValue = &v
	return s
}

func (s *GetAccountingReportRequest) SetDim(v string) *GetAccountingReportRequest {
	s.Dim = &v
	return s
}

func (s *GetAccountingReportRequest) SetJobId(v string) *GetAccountingReportRequest {
	s.JobId = &v
	return s
}

func (s *GetAccountingReportRequest) SetPageSize(v int32) *GetAccountingReportRequest {
	s.PageSize = &v
	return s
}

func (s *GetAccountingReportRequest) SetPageNumber(v int32) *GetAccountingReportRequest {
	s.PageNumber = &v
	return s
}

type GetAccountingReportResponseBody struct {
	Metrics       *string   `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	TotalCount    *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize      *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	TotalCoreTime *int32    `json:"TotalCoreTime,omitempty" xml:"TotalCoreTime,omitempty"`
	Data          []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s GetAccountingReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponseBody) SetMetrics(v string) *GetAccountingReportResponseBody {
	s.Metrics = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetTotalCount(v int32) *GetAccountingReportResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetRequestId(v string) *GetAccountingReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetPageSize(v int32) *GetAccountingReportResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetPageNumber(v int32) *GetAccountingReportResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetTotalCoreTime(v int32) *GetAccountingReportResponseBody {
	s.TotalCoreTime = &v
	return s
}

func (s *GetAccountingReportResponseBody) SetData(v []*string) *GetAccountingReportResponseBody {
	s.Data = v
	return s
}

type GetAccountingReportResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAccountingReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAccountingReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAccountingReportResponse) GoString() string {
	return s.String()
}

func (s *GetAccountingReportResponse) SetHeaders(v map[string]*string) *GetAccountingReportResponse {
	s.Headers = v
	return s
}

func (s *GetAccountingReportResponse) SetBody(v *GetAccountingReportResponseBody) *GetAccountingReportResponse {
	s.Body = v
	return s
}

type GetAutoScaleConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigRequest) SetClusterId(v string) *GetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

type GetAutoScaleConfigResponseBody struct {
	ExtraNodesGrowRatio     *int32                                  `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	RequestId               *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	EnableAutoGrow          *bool                                   `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	ClusterId               *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	MaxNodesInCluster       *int32                                  `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	ShrinkIdleTimes         *int32                                  `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	EnableAutoShrink        *bool                                   `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	ClusterType             *string                                 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	GrowRatio               *int32                                  `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	GrowIntervalInMinutes   *int32                                  `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	Uid                     *string                                 `json:"Uid,omitempty" xml:"Uid,omitempty"`
	GrowTimeoutInMinutes    *int32                                  `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	ImageId                 *string                                 `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ShrinkIntervalInMinutes *int32                                  `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	SpotPriceLimit          *float32                                `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	Queues                  []*GetAutoScaleConfigResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
	ExcludeNodes            *string                                 `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	SpotStrategy            *string                                 `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s GetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBody) SetExtraNodesGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetRequestId(v string) *GetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetClusterId(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetMaxNodesInCluster(v int32) *GetAutoScaleConfigResponseBody {
	s.MaxNodesInCluster = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIdleTimes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBody {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetClusterType(v string) *GetAutoScaleConfigResponseBody {
	s.ClusterType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowRatio(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowRatio = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetUid(v string) *GetAutoScaleConfigResponseBody {
	s.Uid = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetGrowTimeoutInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetImageId(v string) *GetAutoScaleConfigResponseBody {
	s.ImageId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetShrinkIntervalInMinutes(v int32) *GetAutoScaleConfigResponseBody {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBody {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetQueues(v []*GetAutoScaleConfigResponseBodyQueues) *GetAutoScaleConfigResponseBody {
	s.Queues = v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetExcludeNodes(v string) *GetAutoScaleConfigResponseBody {
	s.ExcludeNodes = &v
	return s
}

func (s *GetAutoScaleConfigResponseBody) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBody {
	s.SpotStrategy = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueues struct {
	MinNodesInQueue  *int32                                               `json:"MinNodesInQueue,omitempty" xml:"MinNodesInQueue,omitempty"`
	MaxNodesInQueue  *int32                                               `json:"MaxNodesInQueue,omitempty" xml:"MaxNodesInQueue,omitempty"`
	EnableAutoShrink *bool                                                `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	QueueName        *string                                              `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	QueueImageId     *string                                              `json:"QueueImageId,omitempty" xml:"QueueImageId,omitempty"`
	EnableAutoGrow   *bool                                                `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	ResourceGroupId  *string                                              `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SpotPriceLimit   *float32                                             `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	InstanceTypes    []*GetAutoScaleConfigResponseBodyQueuesInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InstanceType     *string                                              `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotStrategy     *string                                              `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetMinNodesInQueue(v int32) *GetAutoScaleConfigResponseBodyQueues {
	s.MinNodesInQueue = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetMaxNodesInQueue(v int32) *GetAutoScaleConfigResponseBodyQueues {
	s.MaxNodesInQueue = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetEnableAutoShrink(v bool) *GetAutoScaleConfigResponseBodyQueues {
	s.EnableAutoShrink = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetQueueName(v string) *GetAutoScaleConfigResponseBodyQueues {
	s.QueueName = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetQueueImageId(v string) *GetAutoScaleConfigResponseBodyQueues {
	s.QueueImageId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetEnableAutoGrow(v bool) *GetAutoScaleConfigResponseBodyQueues {
	s.EnableAutoGrow = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetResourceGroupId(v string) *GetAutoScaleConfigResponseBodyQueues {
	s.ResourceGroupId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBodyQueues {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetInstanceTypes(v []*GetAutoScaleConfigResponseBodyQueuesInstanceTypes) *GetAutoScaleConfigResponseBodyQueues {
	s.InstanceTypes = v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetInstanceType(v string) *GetAutoScaleConfigResponseBodyQueues {
	s.InstanceType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueues) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBodyQueues {
	s.SpotStrategy = &v
	return s
}

type GetAutoScaleConfigResponseBodyQueuesInstanceTypes struct {
	HostNamePrefix *string  `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	VSwitchId      *string  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId         *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	InstanceType   *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotStrategy   *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s GetAutoScaleConfigResponseBodyQueuesInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponseBodyQueuesInstanceTypes) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponseBodyQueuesInstanceTypes) SetHostNamePrefix(v string) *GetAutoScaleConfigResponseBodyQueuesInstanceTypes {
	s.HostNamePrefix = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesInstanceTypes) SetVSwitchId(v string) *GetAutoScaleConfigResponseBodyQueuesInstanceTypes {
	s.VSwitchId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesInstanceTypes) SetZoneId(v string) *GetAutoScaleConfigResponseBodyQueuesInstanceTypes {
	s.ZoneId = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesInstanceTypes) SetSpotPriceLimit(v float32) *GetAutoScaleConfigResponseBodyQueuesInstanceTypes {
	s.SpotPriceLimit = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesInstanceTypes) SetInstanceType(v string) *GetAutoScaleConfigResponseBodyQueuesInstanceTypes {
	s.InstanceType = &v
	return s
}

func (s *GetAutoScaleConfigResponseBodyQueuesInstanceTypes) SetSpotStrategy(v string) *GetAutoScaleConfigResponseBodyQueuesInstanceTypes {
	s.SpotStrategy = &v
	return s
}

type GetAutoScaleConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *GetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *GetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *GetAutoScaleConfigResponse) SetBody(v *GetAutoScaleConfigResponseBody) *GetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type GetCloudMetricLogsRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	From                *int32  `json:"From,omitempty" xml:"From,omitempty"`
	To                  *int32  `json:"To,omitempty" xml:"To,omitempty"`
	Reverse             *bool   `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	AggregationType     *string `json:"AggregationType,omitempty" xml:"AggregationType,omitempty"`
	AggregationInterval *int32  `json:"AggregationInterval,omitempty" xml:"AggregationInterval,omitempty"`
	MetricScope         *string `json:"MetricScope,omitempty" xml:"MetricScope,omitempty"`
	Filter              *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	MetricCategories    *string `json:"MetricCategories,omitempty" xml:"MetricCategories,omitempty"`
}

func (s GetCloudMetricLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsRequest) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsRequest) SetClusterId(v string) *GetCloudMetricLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetFrom(v int32) *GetCloudMetricLogsRequest {
	s.From = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetTo(v int32) *GetCloudMetricLogsRequest {
	s.To = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetReverse(v bool) *GetCloudMetricLogsRequest {
	s.Reverse = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetAggregationType(v string) *GetCloudMetricLogsRequest {
	s.AggregationType = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetAggregationInterval(v int32) *GetCloudMetricLogsRequest {
	s.AggregationInterval = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetMetricScope(v string) *GetCloudMetricLogsRequest {
	s.MetricScope = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetFilter(v string) *GetCloudMetricLogsRequest {
	s.Filter = &v
	return s
}

func (s *GetCloudMetricLogsRequest) SetMetricCategories(v string) *GetCloudMetricLogsRequest {
	s.MetricCategories = &v
	return s
}

type GetCloudMetricLogsResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MetricLogs []*GetCloudMetricLogsResponseBodyMetricLogs `json:"MetricLogs,omitempty" xml:"MetricLogs,omitempty" type:"Repeated"`
}

func (s GetCloudMetricLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBody) SetRequestId(v string) *GetCloudMetricLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudMetricLogsResponseBody) SetMetricLogs(v []*GetCloudMetricLogsResponseBodyMetricLogs) *GetCloudMetricLogsResponseBody {
	s.MetricLogs = v
	return s
}

type GetCloudMetricLogsResponseBodyMetricLogs struct {
	Time             *int32  `json:"Time,omitempty" xml:"Time,omitempty"`
	DiskDevice       *string `json:"DiskDevice,omitempty" xml:"DiskDevice,omitempty"`
	NetworkInterface *string `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty"`
	MetricData       *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	Hostname         *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCloudMetricLogsResponseBodyMetricLogs) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponseBodyMetricLogs) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetTime(v int32) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.Time = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetDiskDevice(v string) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.DiskDevice = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetNetworkInterface(v string) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.NetworkInterface = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetMetricData(v string) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.MetricData = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetHostname(v string) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.Hostname = &v
	return s
}

func (s *GetCloudMetricLogsResponseBodyMetricLogs) SetInstanceId(v string) *GetCloudMetricLogsResponseBodyMetricLogs {
	s.InstanceId = &v
	return s
}

type GetCloudMetricLogsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCloudMetricLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCloudMetricLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricLogsResponse) GoString() string {
	return s.String()
}

func (s *GetCloudMetricLogsResponse) SetHeaders(v map[string]*string) *GetCloudMetricLogsResponse {
	s.Headers = v
	return s
}

func (s *GetCloudMetricLogsResponse) SetBody(v *GetCloudMetricLogsResponseBody) *GetCloudMetricLogsResponse {
	s.Body = v
	return s
}

type GetCloudMetricProfilingRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ProfilingId *string `json:"ProfilingId,omitempty" xml:"ProfilingId,omitempty"`
}

func (s GetCloudMetricProfilingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingRequest) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingRequest) SetRegionId(v string) *GetCloudMetricProfilingRequest {
	s.RegionId = &v
	return s
}

func (s *GetCloudMetricProfilingRequest) SetClusterId(v string) *GetCloudMetricProfilingRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCloudMetricProfilingRequest) SetProfilingId(v string) *GetCloudMetricProfilingRequest {
	s.ProfilingId = &v
	return s
}

type GetCloudMetricProfilingResponseBody struct {
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SvgUrls   []*GetCloudMetricProfilingResponseBodySvgUrls `json:"SvgUrls,omitempty" xml:"SvgUrls,omitempty" type:"Repeated"`
}

func (s GetCloudMetricProfilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBody) SetRequestId(v string) *GetCloudMetricProfilingResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBody) SetSvgUrls(v []*GetCloudMetricProfilingResponseBodySvgUrls) *GetCloudMetricProfilingResponseBody {
	s.SvgUrls = v
	return s
}

type GetCloudMetricProfilingResponseBodySvgUrls struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Size *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Url  *string `json:"Url,omitempty" xml:"Url,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCloudMetricProfilingResponseBodySvgUrls) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponseBodySvgUrls) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponseBodySvgUrls) SetType(v string) *GetCloudMetricProfilingResponseBodySvgUrls {
	s.Type = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrls) SetSize(v int32) *GetCloudMetricProfilingResponseBodySvgUrls {
	s.Size = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrls) SetUrl(v string) *GetCloudMetricProfilingResponseBodySvgUrls {
	s.Url = &v
	return s
}

func (s *GetCloudMetricProfilingResponseBodySvgUrls) SetName(v string) *GetCloudMetricProfilingResponseBodySvgUrls {
	s.Name = &v
	return s
}

type GetCloudMetricProfilingResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCloudMetricProfilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCloudMetricProfilingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCloudMetricProfilingResponse) GoString() string {
	return s.String()
}

func (s *GetCloudMetricProfilingResponse) SetHeaders(v map[string]*string) *GetCloudMetricProfilingResponse {
	s.Headers = v
	return s
}

func (s *GetCloudMetricProfilingResponse) SetBody(v *GetCloudMetricProfilingResponseBody) *GetCloudMetricProfilingResponse {
	s.Body = v
	return s
}

type GetClusterVolumesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetClusterVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesRequest) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesRequest) SetClusterId(v string) *GetClusterVolumesRequest {
	s.ClusterId = &v
	return s
}

type GetClusterVolumesResponseBody struct {
	Volumes   []*GetClusterVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegionId  *string                                 `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBody) SetVolumes(v []*GetClusterVolumesResponseBodyVolumes) *GetClusterVolumesResponseBody {
	s.Volumes = v
	return s
}

func (s *GetClusterVolumesResponseBody) SetRequestId(v string) *GetClusterVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterVolumesResponseBody) SetRegionId(v string) *GetClusterVolumesResponseBody {
	s.RegionId = &v
	return s
}

type GetClusterVolumesResponseBodyVolumes struct {
	JobQueue         *string                                      `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	VolumeId         *string                                      `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	Roles            []*GetClusterVolumesResponseBodyVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	RemoteDirectory  *string                                      `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeMountpoint *string                                      `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	LocalDirectory   *string                                      `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	VolumeType       *string                                      `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	MustKeep         *bool                                        `json:"MustKeep,omitempty" xml:"MustKeep,omitempty"`
	Location         *string                                      `json:"Location,omitempty" xml:"Location,omitempty"`
	VolumeProtocol   *string                                      `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
}

func (s GetClusterVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumes) SetJobQueue(v string) *GetClusterVolumesResponseBodyVolumes {
	s.JobQueue = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetVolumeId(v string) *GetClusterVolumesResponseBodyVolumes {
	s.VolumeId = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetRoles(v []*GetClusterVolumesResponseBodyVolumesRoles) *GetClusterVolumesResponseBodyVolumes {
	s.Roles = v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetRemoteDirectory(v string) *GetClusterVolumesResponseBodyVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetVolumeMountpoint(v string) *GetClusterVolumesResponseBodyVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetLocalDirectory(v string) *GetClusterVolumesResponseBodyVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetVolumeType(v string) *GetClusterVolumesResponseBodyVolumes {
	s.VolumeType = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetMustKeep(v bool) *GetClusterVolumesResponseBodyVolumes {
	s.MustKeep = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetLocation(v string) *GetClusterVolumesResponseBodyVolumes {
	s.Location = &v
	return s
}

func (s *GetClusterVolumesResponseBodyVolumes) SetVolumeProtocol(v string) *GetClusterVolumesResponseBodyVolumes {
	s.VolumeProtocol = &v
	return s
}

type GetClusterVolumesResponseBodyVolumesRoles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetClusterVolumesResponseBodyVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponseBodyVolumesRoles) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponseBodyVolumesRoles) SetName(v string) *GetClusterVolumesResponseBodyVolumesRoles {
	s.Name = &v
	return s
}

type GetClusterVolumesResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetClusterVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetClusterVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterVolumesResponse) GoString() string {
	return s.String()
}

func (s *GetClusterVolumesResponse) SetHeaders(v map[string]*string) *GetClusterVolumesResponse {
	s.Headers = v
	return s
}

func (s *GetClusterVolumesResponse) SetBody(v *GetClusterVolumesResponseBody) *GetClusterVolumesResponse {
	s.Body = v
	return s
}

type GetGWSConnectTicketRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	AppName    *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s GetGWSConnectTicketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketRequest) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketRequest) SetInstanceId(v string) *GetGWSConnectTicketRequest {
	s.InstanceId = &v
	return s
}

func (s *GetGWSConnectTicketRequest) SetAppName(v string) *GetGWSConnectTicketRequest {
	s.AppName = &v
	return s
}

type GetGWSConnectTicketResponseBody struct {
	Ticket    *string `json:"Ticket,omitempty" xml:"Ticket,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGWSConnectTicketResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketResponseBody) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketResponseBody) SetTicket(v string) *GetGWSConnectTicketResponseBody {
	s.Ticket = &v
	return s
}

func (s *GetGWSConnectTicketResponseBody) SetRequestId(v string) *GetGWSConnectTicketResponseBody {
	s.RequestId = &v
	return s
}

type GetGWSConnectTicketResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetGWSConnectTicketResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetGWSConnectTicketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGWSConnectTicketResponse) GoString() string {
	return s.String()
}

func (s *GetGWSConnectTicketResponse) SetHeaders(v map[string]*string) *GetGWSConnectTicketResponse {
	s.Headers = v
	return s
}

func (s *GetGWSConnectTicketResponse) SetBody(v *GetGWSConnectTicketResponseBody) *GetGWSConnectTicketResponse {
	s.Body = v
	return s
}

type GetHealthMonitorLogsRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	StartTime     *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime       *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EnableReverse *bool   `json:"EnableReverse,omitempty" xml:"EnableReverse,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
}

func (s GetHealthMonitorLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHealthMonitorLogsRequest) GoString() string {
	return s.String()
}

func (s *GetHealthMonitorLogsRequest) SetClusterId(v string) *GetHealthMonitorLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *GetHealthMonitorLogsRequest) SetStartTime(v int32) *GetHealthMonitorLogsRequest {
	s.StartTime = &v
	return s
}

func (s *GetHealthMonitorLogsRequest) SetEndTime(v int32) *GetHealthMonitorLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetHealthMonitorLogsRequest) SetEnableReverse(v bool) *GetHealthMonitorLogsRequest {
	s.EnableReverse = &v
	return s
}

func (s *GetHealthMonitorLogsRequest) SetFilter(v string) *GetHealthMonitorLogsRequest {
	s.Filter = &v
	return s
}

type GetHealthMonitorLogsResponseBody struct {
	LogInfo   []*GetHealthMonitorLogsResponseBodyLogInfo `json:"LogInfo,omitempty" xml:"LogInfo,omitempty" type:"Repeated"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetHealthMonitorLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHealthMonitorLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHealthMonitorLogsResponseBody) SetLogInfo(v []*GetHealthMonitorLogsResponseBodyLogInfo) *GetHealthMonitorLogsResponseBody {
	s.LogInfo = v
	return s
}

func (s *GetHealthMonitorLogsResponseBody) SetRequestId(v string) *GetHealthMonitorLogsResponseBody {
	s.RequestId = &v
	return s
}

type GetHealthMonitorLogsResponseBodyLogInfo struct {
	Time             *int32                                              `json:"Time,omitempty" xml:"Time,omitempty"`
	ItemDescription  *string                                             `json:"ItemDescription,omitempty" xml:"ItemDescription,omitempty"`
	ItemName         *string                                             `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	HealthId         *string                                             `json:"HealthId,omitempty" xml:"HealthId,omitempty"`
	CheckList        []*GetHealthMonitorLogsResponseBodyLogInfoCheckList `json:"CheckList,omitempty" xml:"CheckList,omitempty" type:"Repeated"`
	SceneDescription *string                                             `json:"SceneDescription,omitempty" xml:"SceneDescription,omitempty"`
	HostName         *string                                             `json:"HostName,omitempty" xml:"HostName,omitempty"`
	SceneName        *string                                             `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
	InstanceId       *string                                             `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Level            *string                                             `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s GetHealthMonitorLogsResponseBodyLogInfo) String() string {
	return tea.Prettify(s)
}

func (s GetHealthMonitorLogsResponseBodyLogInfo) GoString() string {
	return s.String()
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetTime(v int32) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.Time = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetItemDescription(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.ItemDescription = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetItemName(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.ItemName = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetHealthId(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.HealthId = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetCheckList(v []*GetHealthMonitorLogsResponseBodyLogInfoCheckList) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.CheckList = v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetSceneDescription(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.SceneDescription = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetHostName(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.HostName = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetSceneName(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.SceneName = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetInstanceId(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.InstanceId = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfo) SetLevel(v string) *GetHealthMonitorLogsResponseBodyLogInfo {
	s.Level = &v
	return s
}

type GetHealthMonitorLogsResponseBodyLogInfoCheckList struct {
	CheckInfo        *string `json:"CheckInfo,omitempty" xml:"CheckInfo,omitempty"`
	CheckDescription *string `json:"CheckDescription,omitempty" xml:"CheckDescription,omitempty"`
	CheckSolution    *string `json:"CheckSolution,omitempty" xml:"CheckSolution,omitempty"`
	CheckName        *string `json:"CheckName,omitempty" xml:"CheckName,omitempty"`
}

func (s GetHealthMonitorLogsResponseBodyLogInfoCheckList) String() string {
	return tea.Prettify(s)
}

func (s GetHealthMonitorLogsResponseBodyLogInfoCheckList) GoString() string {
	return s.String()
}

func (s *GetHealthMonitorLogsResponseBodyLogInfoCheckList) SetCheckInfo(v string) *GetHealthMonitorLogsResponseBodyLogInfoCheckList {
	s.CheckInfo = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfoCheckList) SetCheckDescription(v string) *GetHealthMonitorLogsResponseBodyLogInfoCheckList {
	s.CheckDescription = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfoCheckList) SetCheckSolution(v string) *GetHealthMonitorLogsResponseBodyLogInfoCheckList {
	s.CheckSolution = &v
	return s
}

func (s *GetHealthMonitorLogsResponseBodyLogInfoCheckList) SetCheckName(v string) *GetHealthMonitorLogsResponseBodyLogInfoCheckList {
	s.CheckName = &v
	return s
}

type GetHealthMonitorLogsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHealthMonitorLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHealthMonitorLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHealthMonitorLogsResponse) GoString() string {
	return s.String()
}

func (s *GetHealthMonitorLogsResponse) SetHeaders(v map[string]*string) *GetHealthMonitorLogsResponse {
	s.Headers = v
	return s
}

func (s *GetHealthMonitorLogsResponse) SetBody(v *GetHealthMonitorLogsResponseBody) *GetHealthMonitorLogsResponse {
	s.Body = v
	return s
}

type GetHybridClusterConfigRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Node      *string `json:"Node,omitempty" xml:"Node,omitempty"`
}

func (s GetHybridClusterConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigRequest) SetClusterId(v string) *GetHybridClusterConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *GetHybridClusterConfigRequest) SetNode(v string) *GetHybridClusterConfigRequest {
	s.Node = &v
	return s
}

type GetHybridClusterConfigResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClusterConfig *string `json:"ClusterConfig,omitempty" xml:"ClusterConfig,omitempty"`
}

func (s GetHybridClusterConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigResponseBody) SetRequestId(v string) *GetHybridClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHybridClusterConfigResponseBody) SetClusterConfig(v string) *GetHybridClusterConfigResponseBody {
	s.ClusterConfig = &v
	return s
}

type GetHybridClusterConfigResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetHybridClusterConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetHybridClusterConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetHybridClusterConfigResponse) GoString() string {
	return s.String()
}

func (s *GetHybridClusterConfigResponse) SetHeaders(v map[string]*string) *GetHybridClusterConfigResponse {
	s.Headers = v
	return s
}

func (s *GetHybridClusterConfigResponse) SetBody(v *GetHybridClusterConfigResponseBody) *GetHybridClusterConfigResponse {
	s.Body = v
	return s
}

type GetIfEcsTypeSupportHtConfigRequest struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetIfEcsTypeSupportHtConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigRequest) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigRequest) SetInstanceType(v string) *GetIfEcsTypeSupportHtConfigRequest {
	s.InstanceType = &v
	return s
}

type GetIfEcsTypeSupportHtConfigResponseBody struct {
	DefaultHtEnabled *bool   `json:"DefaultHtEnabled,omitempty" xml:"DefaultHtEnabled,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupportHtConfig  *bool   `json:"SupportHtConfig,omitempty" xml:"SupportHtConfig,omitempty"`
	InstanceType     *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetIfEcsTypeSupportHtConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetDefaultHtEnabled(v bool) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.DefaultHtEnabled = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetRequestId(v string) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetSupportHtConfig(v bool) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.SupportHtConfig = &v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponseBody) SetInstanceType(v string) *GetIfEcsTypeSupportHtConfigResponseBody {
	s.InstanceType = &v
	return s
}

type GetIfEcsTypeSupportHtConfigResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetIfEcsTypeSupportHtConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetIfEcsTypeSupportHtConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIfEcsTypeSupportHtConfigResponse) GoString() string {
	return s.String()
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetHeaders(v map[string]*string) *GetIfEcsTypeSupportHtConfigResponse {
	s.Headers = v
	return s
}

func (s *GetIfEcsTypeSupportHtConfigResponse) SetBody(v *GetIfEcsTypeSupportHtConfigResponseBody) *GetIfEcsTypeSupportHtConfigResponse {
	s.Body = v
	return s
}

type GetVisualServiceStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetVisualServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusRequest) SetClusterId(v string) *GetVisualServiceStatusRequest {
	s.ClusterId = &v
	return s
}

type GetVisualServiceStatusResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVisualServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusResponseBody) SetMessage(v string) *GetVisualServiceStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetVisualServiceStatusResponseBody) SetRequestId(v string) *GetVisualServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

type GetVisualServiceStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVisualServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVisualServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVisualServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetVisualServiceStatusResponse) SetHeaders(v map[string]*string) *GetVisualServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetVisualServiceStatusResponse) SetBody(v *GetVisualServiceStatusResponseBody) *GetVisualServiceStatusResponse {
	s.Body = v
	return s
}

type GetWorkbenchTokenRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UserName             *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPassword         *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
	Port                 *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	AccountSessionTicket *string `json:"AccountSessionTicket,omitempty" xml:"AccountSessionTicket,omitempty"`
	AccountUid           *string `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetWorkbenchTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetWorkbenchTokenRequest) GoString() string {
	return s.String()
}

func (s *GetWorkbenchTokenRequest) SetClusterId(v string) *GetWorkbenchTokenRequest {
	s.ClusterId = &v
	return s
}

func (s *GetWorkbenchTokenRequest) SetUserName(v string) *GetWorkbenchTokenRequest {
	s.UserName = &v
	return s
}

func (s *GetWorkbenchTokenRequest) SetUserPassword(v string) *GetWorkbenchTokenRequest {
	s.UserPassword = &v
	return s
}

func (s *GetWorkbenchTokenRequest) SetPort(v int32) *GetWorkbenchTokenRequest {
	s.Port = &v
	return s
}

func (s *GetWorkbenchTokenRequest) SetAccountSessionTicket(v string) *GetWorkbenchTokenRequest {
	s.AccountSessionTicket = &v
	return s
}

func (s *GetWorkbenchTokenRequest) SetAccountUid(v string) *GetWorkbenchTokenRequest {
	s.AccountUid = &v
	return s
}

func (s *GetWorkbenchTokenRequest) SetInstanceId(v string) *GetWorkbenchTokenRequest {
	s.InstanceId = &v
	return s
}

type GetWorkbenchTokenResponseBody struct {
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root      *GetWorkbenchTokenResponseBodyRoot `json:"root,omitempty" xml:"root,omitempty" type:"Struct"`
}

func (s GetWorkbenchTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkbenchTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkbenchTokenResponseBody) SetRequestId(v string) *GetWorkbenchTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkbenchTokenResponseBody) SetRoot(v *GetWorkbenchTokenResponseBodyRoot) *GetWorkbenchTokenResponseBody {
	s.Root = v
	return s
}

type GetWorkbenchTokenResponseBodyRoot struct {
	InstanceLoginView *GetWorkbenchTokenResponseBodyRootInstanceLoginView `json:"instanceLoginView,omitempty" xml:"instanceLoginView,omitempty" type:"Struct"`
}

func (s GetWorkbenchTokenResponseBodyRoot) String() string {
	return tea.Prettify(s)
}

func (s GetWorkbenchTokenResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *GetWorkbenchTokenResponseBodyRoot) SetInstanceLoginView(v *GetWorkbenchTokenResponseBodyRootInstanceLoginView) *GetWorkbenchTokenResponseBodyRoot {
	s.InstanceLoginView = v
	return s
}

type GetWorkbenchTokenResponseBodyRootInstanceLoginView struct {
	DefaultViewUrl  *string `json:"defaultViewUrl,omitempty" xml:"defaultViewUrl,omitempty"`
	RdpViewUrl      *string `json:"rdpViewUrl,omitempty" xml:"rdpViewUrl,omitempty"`
	BaseViewUrl     *string `json:"baseViewUrl,omitempty" xml:"baseViewUrl,omitempty"`
	FileTreeViewUrl *string `json:"fileTreeViewUrl,omitempty" xml:"fileTreeViewUrl,omitempty"`
	TerminalViewUrl *string `json:"terminalViewUrl,omitempty" xml:"terminalViewUrl,omitempty"`
	ViewName        *string `json:"viewName,omitempty" xml:"viewName,omitempty"`
}

func (s GetWorkbenchTokenResponseBodyRootInstanceLoginView) String() string {
	return tea.Prettify(s)
}

func (s GetWorkbenchTokenResponseBodyRootInstanceLoginView) GoString() string {
	return s.String()
}

func (s *GetWorkbenchTokenResponseBodyRootInstanceLoginView) SetDefaultViewUrl(v string) *GetWorkbenchTokenResponseBodyRootInstanceLoginView {
	s.DefaultViewUrl = &v
	return s
}

func (s *GetWorkbenchTokenResponseBodyRootInstanceLoginView) SetRdpViewUrl(v string) *GetWorkbenchTokenResponseBodyRootInstanceLoginView {
	s.RdpViewUrl = &v
	return s
}

func (s *GetWorkbenchTokenResponseBodyRootInstanceLoginView) SetBaseViewUrl(v string) *GetWorkbenchTokenResponseBodyRootInstanceLoginView {
	s.BaseViewUrl = &v
	return s
}

func (s *GetWorkbenchTokenResponseBodyRootInstanceLoginView) SetFileTreeViewUrl(v string) *GetWorkbenchTokenResponseBodyRootInstanceLoginView {
	s.FileTreeViewUrl = &v
	return s
}

func (s *GetWorkbenchTokenResponseBodyRootInstanceLoginView) SetTerminalViewUrl(v string) *GetWorkbenchTokenResponseBodyRootInstanceLoginView {
	s.TerminalViewUrl = &v
	return s
}

func (s *GetWorkbenchTokenResponseBodyRootInstanceLoginView) SetViewName(v string) *GetWorkbenchTokenResponseBodyRootInstanceLoginView {
	s.ViewName = &v
	return s
}

type GetWorkbenchTokenResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkbenchTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkbenchTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkbenchTokenResponse) GoString() string {
	return s.String()
}

func (s *GetWorkbenchTokenResponse) SetHeaders(v map[string]*string) *GetWorkbenchTokenResponse {
	s.Headers = v
	return s
}

func (s *GetWorkbenchTokenResponse) SetBody(v *GetWorkbenchTokenResponseBody) *GetWorkbenchTokenResponse {
	s.Body = v
	return s
}

type InitializeEHPCRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s InitializeEHPCRequest) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCRequest) GoString() string {
	return s.String()
}

func (s *InitializeEHPCRequest) SetRegionId(v string) *InitializeEHPCRequest {
	s.RegionId = &v
	return s
}

type InitializeEHPCResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeEHPCResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeEHPCResponseBody) SetRequestId(v string) *InitializeEHPCResponseBody {
	s.RequestId = &v
	return s
}

type InitializeEHPCResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitializeEHPCResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitializeEHPCResponse) String() string {
	return tea.Prettify(s)
}

func (s InitializeEHPCResponse) GoString() string {
	return s.String()
}

func (s *InitializeEHPCResponse) SetHeaders(v map[string]*string) *InitializeEHPCResponse {
	s.Headers = v
	return s
}

func (s *InitializeEHPCResponse) SetBody(v *InitializeEHPCResponseBody) *InitializeEHPCResponse {
	s.Body = v
	return s
}

type InstallSoftwareRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
}

func (s InstallSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareRequest) GoString() string {
	return s.String()
}

func (s *InstallSoftwareRequest) SetClusterId(v string) *InstallSoftwareRequest {
	s.ClusterId = &v
	return s
}

func (s *InstallSoftwareRequest) SetApplication(v string) *InstallSoftwareRequest {
	s.Application = &v
	return s
}

type InstallSoftwareResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InstallSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *InstallSoftwareResponseBody) SetRequestId(v string) *InstallSoftwareResponseBody {
	s.RequestId = &v
	return s
}

type InstallSoftwareResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InstallSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallSoftwareResponse) GoString() string {
	return s.String()
}

func (s *InstallSoftwareResponse) SetHeaders(v map[string]*string) *InstallSoftwareResponse {
	s.Headers = v
	return s
}

func (s *InstallSoftwareResponse) SetBody(v *InstallSoftwareResponseBody) *InstallSoftwareResponse {
	s.Body = v
	return s
}

type InvokeShellCommandRequest struct {
	ClusterId  *string                              `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Command    *string                              `json:"Command,omitempty" xml:"Command,omitempty"`
	WorkingDir *string                              `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	Timeout    *int32                               `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	Instance   []*InvokeShellCommandRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s InvokeShellCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandRequest) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandRequest) SetClusterId(v string) *InvokeShellCommandRequest {
	s.ClusterId = &v
	return s
}

func (s *InvokeShellCommandRequest) SetCommand(v string) *InvokeShellCommandRequest {
	s.Command = &v
	return s
}

func (s *InvokeShellCommandRequest) SetWorkingDir(v string) *InvokeShellCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *InvokeShellCommandRequest) SetTimeout(v int32) *InvokeShellCommandRequest {
	s.Timeout = &v
	return s
}

func (s *InvokeShellCommandRequest) SetInstance(v []*InvokeShellCommandRequestInstance) *InvokeShellCommandRequest {
	s.Instance = v
	return s
}

type InvokeShellCommandRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s InvokeShellCommandRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandRequestInstance) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandRequestInstance) SetId(v string) *InvokeShellCommandRequestInstance {
	s.Id = &v
	return s
}

type InvokeShellCommandResponseBody struct {
	RequestId   *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CommandId   *string   `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s InvokeShellCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponseBody) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponseBody) SetRequestId(v string) *InvokeShellCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvokeShellCommandResponseBody) SetCommandId(v string) *InvokeShellCommandResponseBody {
	s.CommandId = &v
	return s
}

func (s *InvokeShellCommandResponseBody) SetInstanceIds(v []*string) *InvokeShellCommandResponseBody {
	s.InstanceIds = v
	return s
}

type InvokeShellCommandResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InvokeShellCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InvokeShellCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s InvokeShellCommandResponse) GoString() string {
	return s.String()
}

func (s *InvokeShellCommandResponse) SetHeaders(v map[string]*string) *InvokeShellCommandResponse {
	s.Headers = v
	return s
}

func (s *InvokeShellCommandResponse) SetBody(v *InvokeShellCommandResponseBody) *InvokeShellCommandResponse {
	s.Body = v
	return s
}

type ListAccountMappingRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListAccountMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAccountMappingRequest) GoString() string {
	return s.String()
}

func (s *ListAccountMappingRequest) SetClusterId(v string) *ListAccountMappingRequest {
	s.ClusterId = &v
	return s
}

type ListAccountMappingResponseBody struct {
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserMappings []*ListAccountMappingResponseBodyUserMappings `json:"UserMappings,omitempty" xml:"UserMappings,omitempty" type:"Repeated"`
}

func (s ListAccountMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAccountMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccountMappingResponseBody) SetRequestId(v string) *ListAccountMappingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccountMappingResponseBody) SetUserMappings(v []*ListAccountMappingResponseBodyUserMappings) *ListAccountMappingResponseBody {
	s.UserMappings = v
	return s
}

type ListAccountMappingResponseBodyUserMappings struct {
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s ListAccountMappingResponseBodyUserMappings) String() string {
	return tea.Prettify(s)
}

func (s ListAccountMappingResponseBodyUserMappings) GoString() string {
	return s.String()
}

func (s *ListAccountMappingResponseBodyUserMappings) SetAccountId(v string) *ListAccountMappingResponseBodyUserMappings {
	s.AccountId = &v
	return s
}

func (s *ListAccountMappingResponseBodyUserMappings) SetUserName(v string) *ListAccountMappingResponseBodyUserMappings {
	s.UserName = &v
	return s
}

func (s *ListAccountMappingResponseBodyUserMappings) SetAccountName(v string) *ListAccountMappingResponseBodyUserMappings {
	s.AccountName = &v
	return s
}

type ListAccountMappingResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAccountMappingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAccountMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAccountMappingResponse) GoString() string {
	return s.String()
}

func (s *ListAccountMappingResponse) SetHeaders(v map[string]*string) *ListAccountMappingResponse {
	s.Headers = v
	return s
}

func (s *ListAccountMappingResponse) SetBody(v *ListAccountMappingResponseBody) *ListAccountMappingResponse {
	s.Body = v
	return s
}

type ListAvailableEcsTypesRequest struct {
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SpotStrategy       *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	ShowSoldOut        *bool   `json:"ShowSoldOut,omitempty" xml:"ShowSoldOut,omitempty"`
}

func (s ListAvailableEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesRequest) SetZoneId(v string) *ListAvailableEcsTypesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetSpotStrategy(v string) *ListAvailableEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetInstanceChargeType(v string) *ListAvailableEcsTypesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *ListAvailableEcsTypesRequest) SetShowSoldOut(v bool) *ListAvailableEcsTypesRequest {
	s.ShowSoldOut = &v
	return s
}

type ListAvailableEcsTypesResponseBody struct {
	SupportSpotInstance  *bool                                                    `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
	RequestId            *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InstanceTypeFamilies []*ListAvailableEcsTypesResponseBodyInstanceTypeFamilies `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Repeated"`
}

func (s ListAvailableEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListAvailableEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBody) SetRequestId(v string) *ListAvailableEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBody) SetInstanceTypeFamilies(v []*ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) *ListAvailableEcsTypesResponseBody {
	s.InstanceTypeFamilies = v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamilies struct {
	InstanceTypeFamilyId *string                                                       `json:"InstanceTypeFamilyId,omitempty" xml:"InstanceTypeFamilyId,omitempty"`
	Types                []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
	Generation           *string                                                       `json:"Generation,omitempty" xml:"Generation,omitempty"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) SetInstanceTypeFamilyId(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies {
	s.InstanceTypeFamilyId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) SetTypes(v []*ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies {
	s.Types = v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies) SetGeneration(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamilies {
	s.Generation = &v
	return s
}

type ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes struct {
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	InstanceTypeId      *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	InstanceBandwidthRx *int32  `json:"InstanceBandwidthRx,omitempty" xml:"InstanceBandwidthRx,omitempty"`
	GPUSpec             *string `json:"GPUSpec,omitempty" xml:"GPUSpec,omitempty"`
	InstanceBandwidthTx *int32  `json:"InstanceBandwidthTx,omitempty" xml:"InstanceBandwidthTx,omitempty"`
	InstancePpsRx       *int32  `json:"InstancePpsRx,omitempty" xml:"InstancePpsRx,omitempty"`
	InstancePpsTx       *int32  `json:"InstancePpsTx,omitempty" xml:"InstancePpsTx,omitempty"`
	GPUAmount           *int32  `json:"GPUAmount,omitempty" xml:"GPUAmount,omitempty"`
	CpuCoreCount        *int32  `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	MemorySize          *int32  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	EniQuantity         *int32  `json:"EniQuantity,omitempty" xml:"EniQuantity,omitempty"`
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetStatus(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.Status = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetInstanceTypeId(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.InstanceTypeId = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetInstanceBandwidthRx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.InstanceBandwidthRx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetGPUSpec(v string) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.GPUSpec = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetInstanceBandwidthTx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.InstanceBandwidthTx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetInstancePpsRx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.InstancePpsRx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetInstancePpsTx(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.InstancePpsTx = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetGPUAmount(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.GPUAmount = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetCpuCoreCount(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.CpuCoreCount = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetMemorySize(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.MemorySize = &v
	return s
}

func (s *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes) SetEniQuantity(v int32) *ListAvailableEcsTypesResponseBodyInstanceTypeFamiliesTypes {
	s.EniQuantity = &v
	return s
}

type ListAvailableEcsTypesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableEcsTypesResponse) SetHeaders(v map[string]*string) *ListAvailableEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableEcsTypesResponse) SetBody(v *ListAvailableEcsTypesResponseBody) *ListAvailableEcsTypesResponse {
	s.Body = v
	return s
}

type ListAvailableFileSystemTypesResponseBody struct {
	FileSystemTypeList []*ListAvailableFileSystemTypesResponseBodyFileSystemTypeList `json:"FileSystemTypeList,omitempty" xml:"FileSystemTypeList,omitempty" type:"Repeated"`
	RequestId          *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAvailableFileSystemTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemTypesResponseBody) SetFileSystemTypeList(v []*ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) *ListAvailableFileSystemTypesResponseBody {
	s.FileSystemTypeList = v
	return s
}

func (s *ListAvailableFileSystemTypesResponseBody) SetRequestId(v string) *ListAvailableFileSystemTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListAvailableFileSystemTypesResponseBodyFileSystemTypeList struct {
	FileSystemType *string   `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	Available      *bool     `json:"Available,omitempty" xml:"Available,omitempty"`
	ProtocolType   *string   `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	StorageTypes   []*string `json:"StorageTypes,omitempty" xml:"StorageTypes,omitempty" type:"Repeated"`
}

func (s ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) SetFileSystemType(v string) *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList {
	s.FileSystemType = &v
	return s
}

func (s *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) SetAvailable(v bool) *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList {
	s.Available = &v
	return s
}

func (s *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) SetProtocolType(v string) *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList {
	s.ProtocolType = &v
	return s
}

func (s *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList) SetStorageTypes(v []*string) *ListAvailableFileSystemTypesResponseBodyFileSystemTypeList {
	s.StorageTypes = v
	return s
}

type ListAvailableFileSystemTypesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAvailableFileSystemTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAvailableFileSystemTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableFileSystemTypesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableFileSystemTypesResponse) SetHeaders(v map[string]*string) *ListAvailableFileSystemTypesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableFileSystemTypesResponse) SetBody(v *ListAvailableFileSystemTypesResponseBody) *ListAvailableFileSystemTypesResponse {
	s.Body = v
	return s
}

type ListCloudMetricProfilingsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCloudMetricProfilingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsRequest) SetRegionId(v string) *ListCloudMetricProfilingsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetClusterId(v string) *ListCloudMetricProfilingsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetPageNumber(v int32) *ListCloudMetricProfilingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudMetricProfilingsRequest) SetPageSize(v int32) *ListCloudMetricProfilingsRequest {
	s.PageSize = &v
	return s
}

type ListCloudMetricProfilingsResponseBody struct {
	TotalCount *int32                                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Profilings []*ListCloudMetricProfilingsResponseBodyProfilings `json:"Profilings,omitempty" xml:"Profilings,omitempty" type:"Repeated"`
}

func (s ListCloudMetricProfilingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBody) SetTotalCount(v int32) *ListCloudMetricProfilingsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetPageSize(v int32) *ListCloudMetricProfilingsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetRequestId(v string) *ListCloudMetricProfilingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetPageNumber(v int32) *ListCloudMetricProfilingsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBody) SetProfilings(v []*ListCloudMetricProfilingsResponseBodyProfilings) *ListCloudMetricProfilingsResponseBody {
	s.Profilings = v
	return s
}

type ListCloudMetricProfilingsResponseBodyProfilings struct {
	ProfilingId *string `json:"ProfilingId,omitempty" xml:"ProfilingId,omitempty"`
	TriggerTime *string `json:"TriggerTime,omitempty" xml:"TriggerTime,omitempty"`
	Pid         *int32  `json:"Pid,omitempty" xml:"Pid,omitempty"`
	HostName    *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	Duration    *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Freq        *int32  `json:"Freq,omitempty" xml:"Freq,omitempty"`
}

func (s ListCloudMetricProfilingsResponseBodyProfilings) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponseBodyProfilings) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetProfilingId(v string) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.ProfilingId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetTriggerTime(v string) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.TriggerTime = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetPid(v int32) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.Pid = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetHostName(v string) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.HostName = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetDuration(v int32) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.Duration = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetInstanceId(v string) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.InstanceId = &v
	return s
}

func (s *ListCloudMetricProfilingsResponseBodyProfilings) SetFreq(v int32) *ListCloudMetricProfilingsResponseBodyProfilings {
	s.Freq = &v
	return s
}

type ListCloudMetricProfilingsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCloudMetricProfilingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCloudMetricProfilingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCloudMetricProfilingsResponse) GoString() string {
	return s.String()
}

func (s *ListCloudMetricProfilingsResponse) SetHeaders(v map[string]*string) *ListCloudMetricProfilingsResponse {
	s.Headers = v
	return s
}

func (s *ListCloudMetricProfilingsResponse) SetBody(v *ListCloudMetricProfilingsResponseBody) *ListCloudMetricProfilingsResponse {
	s.Body = v
	return s
}

type ListClusterLogsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *ListClusterLogsRequest) SetClusterId(v string) *ListClusterLogsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageNumber(v int32) *ListClusterLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsRequest) SetPageSize(v int32) *ListClusterLogsRequest {
	s.PageSize = &v
	return s
}

type ListClusterLogsResponseBody struct {
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ClusterId  *string                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Logs       []*ListClusterLogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s ListClusterLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBody) SetTotalCount(v int32) *ListClusterLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageSize(v int32) *ListClusterLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetRequestId(v string) *ListClusterLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetPageNumber(v int32) *ListClusterLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetClusterId(v string) *ListClusterLogsResponseBody {
	s.ClusterId = &v
	return s
}

func (s *ListClusterLogsResponseBody) SetLogs(v []*ListClusterLogsResponseBodyLogs) *ListClusterLogsResponseBody {
	s.Logs = v
	return s
}

type ListClusterLogsResponseBodyLogs struct {
	Operation  *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Level      *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ListClusterLogsResponseBodyLogs) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponseBodyLogs) SetOperation(v string) *ListClusterLogsResponseBodyLogs {
	s.Operation = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogs) SetCreateTime(v string) *ListClusterLogsResponseBodyLogs {
	s.CreateTime = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogs) SetMessage(v string) *ListClusterLogsResponseBodyLogs {
	s.Message = &v
	return s
}

func (s *ListClusterLogsResponseBodyLogs) SetLevel(v string) *ListClusterLogsResponseBodyLogs {
	s.Level = &v
	return s
}

type ListClusterLogsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClusterLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *ListClusterLogsResponse) SetHeaders(v map[string]*string) *ListClusterLogsResponse {
	s.Headers = v
	return s
}

func (s *ListClusterLogsResponse) SetBody(v *ListClusterLogsResponseBody) *ListClusterLogsResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetPageNumber(v int32) *ListClustersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersRequest) SetPageSize(v int32) *ListClustersRequest {
	s.PageSize = &v
	return s
}

type ListClustersResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Clusters   []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetTotalCount(v int32) *ListClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClustersResponseBody) SetPageSize(v int32) *ListClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetPageNumber(v int32) *ListClustersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

type ListClustersResponseBodyClusters struct {
	VpcId                 *string                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status                *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	CreateTime            *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UsedResources         *ListClustersResponseBodyClustersUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	ComputeSpotStrategy   *string                                         `json:"ComputeSpotStrategy,omitempty" xml:"ComputeSpotStrategy,omitempty"`
	AccountType           *string                                         `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Count                 *int32                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	EhpcVersion           *string                                         `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	Description           *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	BaseOsTag             *string                                         `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	Name                  *string                                         `json:"Name,omitempty" xml:"Name,omitempty"`
	ImageId               *string                                         `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ComputeSpotPriceLimit *float32                                        `json:"ComputeSpotPriceLimit,omitempty" xml:"ComputeSpotPriceLimit,omitempty"`
	SchedulerType         *string                                         `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	NodeSuffix            *string                                         `json:"NodeSuffix,omitempty" xml:"NodeSuffix,omitempty"`
	DeployMode            *string                                         `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	ImageOwnerAlias       *string                                         `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	OsTag                 *string                                         `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	NodePrefix            *string                                         `json:"NodePrefix,omitempty" xml:"NodePrefix,omitempty"`
	InstanceType          *string                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Managers              *ListClustersResponseBodyClustersManagers       `json:"Managers,omitempty" xml:"Managers,omitempty" type:"Struct"`
	InstanceChargeType    *string                                         `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	RegionId              *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VSwitchId             *string                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	TotalResources        *ListClustersResponseBodyClustersTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	ZoneId                *string                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Computes              *ListClustersResponseBodyClustersComputes       `json:"Computes,omitempty" xml:"Computes,omitempty" type:"Struct"`
	LoginNodes            *string                                         `json:"LoginNodes,omitempty" xml:"LoginNodes,omitempty"`
	Id                    *string                                         `json:"Id,omitempty" xml:"Id,omitempty"`
	Location              *string                                         `json:"Location,omitempty" xml:"Location,omitempty"`
	ClientVersion         *string                                         `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetVpcId(v string) *ListClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetStatus(v string) *ListClustersResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetCreateTime(v string) *ListClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetUsedResources(v *ListClustersResponseBodyClustersUsedResources) *ListClustersResponseBodyClusters {
	s.UsedResources = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComputeSpotStrategy(v string) *ListClustersResponseBodyClusters {
	s.ComputeSpotStrategy = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetAccountType(v string) *ListClustersResponseBodyClusters {
	s.AccountType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetCount(v int32) *ListClustersResponseBodyClusters {
	s.Count = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetEhpcVersion(v string) *ListClustersResponseBodyClusters {
	s.EhpcVersion = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetDescription(v string) *ListClustersResponseBodyClusters {
	s.Description = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetBaseOsTag(v string) *ListClustersResponseBodyClusters {
	s.BaseOsTag = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetName(v string) *ListClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetImageId(v string) *ListClustersResponseBodyClusters {
	s.ImageId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComputeSpotPriceLimit(v float32) *ListClustersResponseBodyClusters {
	s.ComputeSpotPriceLimit = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetSchedulerType(v string) *ListClustersResponseBodyClusters {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeSuffix(v string) *ListClustersResponseBodyClusters {
	s.NodeSuffix = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetDeployMode(v string) *ListClustersResponseBodyClusters {
	s.DeployMode = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetImageOwnerAlias(v string) *ListClustersResponseBodyClusters {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetOsTag(v string) *ListClustersResponseBodyClusters {
	s.OsTag = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodePrefix(v string) *ListClustersResponseBodyClusters {
	s.NodePrefix = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetInstanceType(v string) *ListClustersResponseBodyClusters {
	s.InstanceType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetManagers(v *ListClustersResponseBodyClustersManagers) *ListClustersResponseBodyClusters {
	s.Managers = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetInstanceChargeType(v string) *ListClustersResponseBodyClusters {
	s.InstanceChargeType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetRegionId(v string) *ListClustersResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetVSwitchId(v string) *ListClustersResponseBodyClusters {
	s.VSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetTotalResources(v *ListClustersResponseBodyClustersTotalResources) *ListClustersResponseBodyClusters {
	s.TotalResources = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetZoneId(v string) *ListClustersResponseBodyClusters {
	s.ZoneId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComputes(v *ListClustersResponseBodyClustersComputes) *ListClustersResponseBodyClusters {
	s.Computes = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetLoginNodes(v string) *ListClustersResponseBodyClusters {
	s.LoginNodes = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetId(v string) *ListClustersResponseBodyClusters {
	s.Id = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetLocation(v string) *ListClustersResponseBodyClusters {
	s.Location = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClientVersion(v string) *ListClustersResponseBodyClusters {
	s.ClientVersion = &v
	return s
}

type ListClustersResponseBodyClustersUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersUsedResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersUsedResources) SetCpu(v int32) *ListClustersResponseBodyClustersUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersUsedResources) SetGpu(v int32) *ListClustersResponseBodyClustersUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersUsedResources) SetMemory(v int32) *ListClustersResponseBodyClustersUsedResources {
	s.Memory = &v
	return s
}

type ListClustersResponseBodyClustersManagers struct {
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	NormalCount    *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	OperatingCount *int32 `json:"OperatingCount,omitempty" xml:"OperatingCount,omitempty"`
	StoppedCount   *int32 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	Total          *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersManagers) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersManagers) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersManagers) SetExceptionCount(v int32) *ListClustersResponseBodyClustersManagers {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagers) SetNormalCount(v int32) *ListClustersResponseBodyClustersManagers {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagers) SetOperatingCount(v int32) *ListClustersResponseBodyClustersManagers {
	s.OperatingCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagers) SetStoppedCount(v int32) *ListClustersResponseBodyClustersManagers {
	s.StoppedCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersManagers) SetTotal(v int32) *ListClustersResponseBodyClustersManagers {
	s.Total = &v
	return s
}

type ListClustersResponseBodyClustersTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListClustersResponseBodyClustersTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersTotalResources) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersTotalResources) SetCpu(v int32) *ListClustersResponseBodyClustersTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersTotalResources) SetGpu(v int32) *ListClustersResponseBodyClustersTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListClustersResponseBodyClustersTotalResources) SetMemory(v int32) *ListClustersResponseBodyClustersTotalResources {
	s.Memory = &v
	return s
}

type ListClustersResponseBodyClustersComputes struct {
	ExceptionCount *int32 `json:"ExceptionCount,omitempty" xml:"ExceptionCount,omitempty"`
	NormalCount    *int32 `json:"NormalCount,omitempty" xml:"NormalCount,omitempty"`
	OperatingCount *int32 `json:"OperatingCount,omitempty" xml:"OperatingCount,omitempty"`
	StoppedCount   *int32 `json:"StoppedCount,omitempty" xml:"StoppedCount,omitempty"`
	Total          *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyClustersComputes) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersComputes) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersComputes) SetExceptionCount(v int32) *ListClustersResponseBodyClustersComputes {
	s.ExceptionCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersComputes) SetNormalCount(v int32) *ListClustersResponseBodyClustersComputes {
	s.NormalCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersComputes) SetOperatingCount(v int32) *ListClustersResponseBodyClustersComputes {
	s.OperatingCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersComputes) SetStoppedCount(v int32) *ListClustersResponseBodyClustersComputes {
	s.StoppedCount = &v
	return s
}

func (s *ListClustersResponseBodyClustersComputes) SetTotal(v int32) *ListClustersResponseBodyClustersComputes {
	s.Total = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListClustersMetaRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListClustersMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaRequest) GoString() string {
	return s.String()
}

func (s *ListClustersMetaRequest) SetPageNumber(v int32) *ListClustersMetaRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClustersMetaRequest) SetPageSize(v int32) *ListClustersMetaRequest {
	s.PageSize = &v
	return s
}

type ListClustersMetaResponseBody struct {
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Clusters   []*ListClustersMetaResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
}

func (s ListClustersMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBody) SetTotalCount(v int32) *ListClustersMetaResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetPageSize(v int32) *ListClustersMetaResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetRequestId(v string) *ListClustersMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetPageNumber(v int32) *ListClustersMetaResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListClustersMetaResponseBody) SetClusters(v []*ListClustersMetaResponseBodyClusters) *ListClustersMetaResponseBody {
	s.Clusters = v
	return s
}

type ListClustersMetaResponseBodyClusters struct {
	VpcId         *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DeployMode    *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	OsTag         *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AccountType   *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	Location      *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id            *string `json:"Id,omitempty" xml:"Id,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s ListClustersMetaResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponseBodyClusters) SetVpcId(v string) *ListClustersMetaResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetStatus(v string) *ListClustersMetaResponseBodyClusters {
	s.Status = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetSchedulerType(v string) *ListClustersMetaResponseBodyClusters {
	s.SchedulerType = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetDescription(v string) *ListClustersMetaResponseBodyClusters {
	s.Description = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetDeployMode(v string) *ListClustersMetaResponseBodyClusters {
	s.DeployMode = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetOsTag(v string) *ListClustersMetaResponseBodyClusters {
	s.OsTag = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetName(v string) *ListClustersMetaResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetAccountType(v string) *ListClustersMetaResponseBodyClusters {
	s.AccountType = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetLocation(v string) *ListClustersMetaResponseBodyClusters {
	s.Location = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetId(v string) *ListClustersMetaResponseBodyClusters {
	s.Id = &v
	return s
}

func (s *ListClustersMetaResponseBodyClusters) SetClientVersion(v string) *ListClustersMetaResponseBodyClusters {
	s.ClientVersion = &v
	return s
}

type ListClustersMetaResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListClustersMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListClustersMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersMetaResponse) GoString() string {
	return s.String()
}

func (s *ListClustersMetaResponse) SetHeaders(v map[string]*string) *ListClustersMetaResponse {
	s.Headers = v
	return s
}

func (s *ListClustersMetaResponse) SetBody(v *ListClustersMetaResponseBody) *ListClustersMetaResponse {
	s.Body = v
	return s
}

type ListCommandsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandId  *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCommandsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListCommandsRequest) SetClusterId(v string) *ListCommandsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListCommandsRequest) SetCommandId(v string) *ListCommandsRequest {
	s.CommandId = &v
	return s
}

func (s *ListCommandsRequest) SetPageNumber(v int32) *ListCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCommandsRequest) SetPageSize(v int32) *ListCommandsRequest {
	s.PageSize = &v
	return s
}

type ListCommandsResponseBody struct {
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Commands   []*ListCommandsResponseBodyCommands `json:"Commands,omitempty" xml:"Commands,omitempty" type:"Repeated"`
}

func (s ListCommandsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBody) SetTotalCount(v int32) *ListCommandsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCommandsResponseBody) SetPageSize(v int32) *ListCommandsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCommandsResponseBody) SetRequestId(v string) *ListCommandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCommandsResponseBody) SetPageNumber(v int32) *ListCommandsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCommandsResponseBody) SetCommands(v []*ListCommandsResponseBodyCommands) *ListCommandsResponseBody {
	s.Commands = v
	return s
}

type ListCommandsResponseBodyCommands struct {
	Timeout        *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	WorkingDir     *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	CommandId      *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
}

func (s ListCommandsResponseBodyCommands) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponseBodyCommands) GoString() string {
	return s.String()
}

func (s *ListCommandsResponseBodyCommands) SetTimeout(v string) *ListCommandsResponseBodyCommands {
	s.Timeout = &v
	return s
}

func (s *ListCommandsResponseBodyCommands) SetWorkingDir(v string) *ListCommandsResponseBodyCommands {
	s.WorkingDir = &v
	return s
}

func (s *ListCommandsResponseBodyCommands) SetCommandContent(v string) *ListCommandsResponseBodyCommands {
	s.CommandContent = &v
	return s
}

func (s *ListCommandsResponseBodyCommands) SetCommandId(v string) *ListCommandsResponseBodyCommands {
	s.CommandId = &v
	return s
}

type ListCommandsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCommandsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCommandsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCommandsResponse) GoString() string {
	return s.String()
}

func (s *ListCommandsResponse) SetHeaders(v map[string]*string) *ListCommandsResponse {
	s.Headers = v
	return s
}

func (s *ListCommandsResponse) SetBody(v *ListCommandsResponseBody) *ListCommandsResponse {
	s.Body = v
	return s
}

type ListContainerAppsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContainerAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsRequest) GoString() string {
	return s.String()
}

func (s *ListContainerAppsRequest) SetPageNumber(v int32) *ListContainerAppsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContainerAppsRequest) SetPageSize(v int32) *ListContainerAppsRequest {
	s.PageSize = &v
	return s
}

type ListContainerAppsResponseBody struct {
	ContainerApps []*ListContainerAppsResponseBodyContainerApps `json:"ContainerApps,omitempty" xml:"ContainerApps,omitempty" type:"Repeated"`
	TotalCount    *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize      *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber    *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListContainerAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBody) SetContainerApps(v []*ListContainerAppsResponseBodyContainerApps) *ListContainerAppsResponseBody {
	s.ContainerApps = v
	return s
}

func (s *ListContainerAppsResponseBody) SetTotalCount(v int32) *ListContainerAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetPageSize(v int32) *ListContainerAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetRequestId(v string) *ListContainerAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerAppsResponseBody) SetPageNumber(v int32) *ListContainerAppsResponseBody {
	s.PageNumber = &v
	return s
}

type ListContainerAppsResponseBodyContainerApps struct {
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Repository  *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	ImageTag    *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListContainerAppsResponseBodyContainerApps) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponseBodyContainerApps) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponseBodyContainerApps) SetType(v string) *ListContainerAppsResponseBodyContainerApps {
	s.Type = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerApps) SetDescription(v string) *ListContainerAppsResponseBodyContainerApps {
	s.Description = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerApps) SetCreateTime(v string) *ListContainerAppsResponseBodyContainerApps {
	s.CreateTime = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerApps) SetRepository(v string) *ListContainerAppsResponseBodyContainerApps {
	s.Repository = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerApps) SetImageTag(v string) *ListContainerAppsResponseBodyContainerApps {
	s.ImageTag = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerApps) SetName(v string) *ListContainerAppsResponseBodyContainerApps {
	s.Name = &v
	return s
}

func (s *ListContainerAppsResponseBodyContainerApps) SetId(v string) *ListContainerAppsResponseBodyContainerApps {
	s.Id = &v
	return s
}

type ListContainerAppsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *ListContainerAppsResponse) SetHeaders(v map[string]*string) *ListContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *ListContainerAppsResponse) SetBody(v *ListContainerAppsResponseBody) *ListContainerAppsResponse {
	s.Body = v
	return s
}

type ListContainerImagesRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListContainerImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesRequest) GoString() string {
	return s.String()
}

func (s *ListContainerImagesRequest) SetClusterId(v string) *ListContainerImagesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListContainerImagesRequest) SetContainerType(v string) *ListContainerImagesRequest {
	s.ContainerType = &v
	return s
}

func (s *ListContainerImagesRequest) SetPageNumber(v int32) *ListContainerImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContainerImagesRequest) SetPageSize(v int32) *ListContainerImagesRequest {
	s.PageSize = &v
	return s
}

type ListContainerImagesResponseBody struct {
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	DBInfo     *string                                  `json:"DBInfo,omitempty" xml:"DBInfo,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Images     []*ListContainerImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s ListContainerImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBody) SetTotalCount(v int32) *ListContainerImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetDBInfo(v string) *ListContainerImagesResponseBody {
	s.DBInfo = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetPageSize(v int32) *ListContainerImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetRequestId(v string) *ListContainerImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetPageNumber(v int32) *ListContainerImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContainerImagesResponseBody) SetImages(v []*ListContainerImagesResponseBodyImages) *ListContainerImagesResponseBody {
	s.Images = v
	return s
}

type ListContainerImagesResponseBodyImages struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateDateTime *string `json:"UpdateDateTime,omitempty" xml:"UpdateDateTime,omitempty"`
	Repository     *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	Tag            *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	System         *string `json:"System,omitempty" xml:"System,omitempty"`
	ImageId        *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s ListContainerImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponseBodyImages) SetType(v string) *ListContainerImagesResponseBodyImages {
	s.Type = &v
	return s
}

func (s *ListContainerImagesResponseBodyImages) SetStatus(v string) *ListContainerImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *ListContainerImagesResponseBodyImages) SetUpdateDateTime(v string) *ListContainerImagesResponseBodyImages {
	s.UpdateDateTime = &v
	return s
}

func (s *ListContainerImagesResponseBodyImages) SetRepository(v string) *ListContainerImagesResponseBodyImages {
	s.Repository = &v
	return s
}

func (s *ListContainerImagesResponseBodyImages) SetTag(v string) *ListContainerImagesResponseBodyImages {
	s.Tag = &v
	return s
}

func (s *ListContainerImagesResponseBodyImages) SetSystem(v string) *ListContainerImagesResponseBodyImages {
	s.System = &v
	return s
}

func (s *ListContainerImagesResponseBodyImages) SetImageId(v string) *ListContainerImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

type ListContainerImagesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListContainerImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListContainerImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListContainerImagesResponse) GoString() string {
	return s.String()
}

func (s *ListContainerImagesResponse) SetHeaders(v map[string]*string) *ListContainerImagesResponse {
	s.Headers = v
	return s
}

func (s *ListContainerImagesResponse) SetBody(v *ListContainerImagesResponseBody) *ListContainerImagesResponse {
	s.Body = v
	return s
}

type ListCpfsFileSystemsRequest struct {
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCpfsFileSystemsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsRequest) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsRequest) SetFileSystemId(v string) *ListCpfsFileSystemsRequest {
	s.FileSystemId = &v
	return s
}

func (s *ListCpfsFileSystemsRequest) SetPageNumber(v int32) *ListCpfsFileSystemsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCpfsFileSystemsRequest) SetPageSize(v int32) *ListCpfsFileSystemsRequest {
	s.PageSize = &v
	return s
}

type ListCpfsFileSystemsResponseBody struct {
	TotalCount     *int32                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FileSystemList []*ListCpfsFileSystemsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Repeated"`
}

func (s ListCpfsFileSystemsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBody) SetTotalCount(v int32) *ListCpfsFileSystemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetPageSize(v int32) *ListCpfsFileSystemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetRequestId(v string) *ListCpfsFileSystemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetPageNumber(v int32) *ListCpfsFileSystemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBody) SetFileSystemList(v []*ListCpfsFileSystemsResponseBodyFileSystemList) *ListCpfsFileSystemsResponseBody {
	s.FileSystemList = v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemList struct {
	FileSystemId    *string                                                         `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	Capacity        *string                                                         `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	CreateTime      *string                                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MountTargetList []*ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Repeated"`
	ZoneId          *string                                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ProtocolType    *string                                                         `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Destription     *string                                                         `json:"Destription,omitempty" xml:"Destription,omitempty"`
	RegionId        *string                                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetFileSystemId(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.FileSystemId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetCapacity(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.Capacity = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetCreateTime(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.CreateTime = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetMountTargetList(v []*ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.MountTargetList = v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetZoneId(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.ZoneId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetProtocolType(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.ProtocolType = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetDestription(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.Destription = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemList) SetRegionId(v string) *ListCpfsFileSystemsResponseBodyFileSystemList {
	s.RegionId = &v
	return s
}

type ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList struct {
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	VswId             *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) SetVpcId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList {
	s.VpcId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) SetStatus(v string) *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList {
	s.Status = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) SetMountTargetDomain(v string) *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList {
	s.MountTargetDomain = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) SetVswId(v string) *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList {
	s.VswId = &v
	return s
}

func (s *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList) SetNetworkType(v string) *ListCpfsFileSystemsResponseBodyFileSystemListMountTargetList {
	s.NetworkType = &v
	return s
}

type ListCpfsFileSystemsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCpfsFileSystemsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCpfsFileSystemsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCpfsFileSystemsResponse) GoString() string {
	return s.String()
}

func (s *ListCpfsFileSystemsResponse) SetHeaders(v map[string]*string) *ListCpfsFileSystemsResponse {
	s.Headers = v
	return s
}

func (s *ListCpfsFileSystemsResponse) SetBody(v *ListCpfsFileSystemsResponseBody) *ListCpfsFileSystemsResponse {
	s.Body = v
	return s
}

type ListCurrentClientVersionResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s ListCurrentClientVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponseBody) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponseBody) SetRequestId(v string) *ListCurrentClientVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCurrentClientVersionResponseBody) SetClientVersion(v string) *ListCurrentClientVersionResponseBody {
	s.ClientVersion = &v
	return s
}

type ListCurrentClientVersionResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCurrentClientVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCurrentClientVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCurrentClientVersionResponse) GoString() string {
	return s.String()
}

func (s *ListCurrentClientVersionResponse) SetHeaders(v map[string]*string) *ListCurrentClientVersionResponse {
	s.Headers = v
	return s
}

func (s *ListCurrentClientVersionResponse) SetBody(v *ListCurrentClientVersionResponseBody) *ListCurrentClientVersionResponse {
	s.Body = v
	return s
}

type ListCustomImagesRequest struct {
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	BaseOsTag       *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListCustomImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomImagesRequest) SetImageOwnerAlias(v string) *ListCustomImagesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesRequest) SetBaseOsTag(v string) *ListCustomImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesRequest) SetInstanceType(v string) *ListCustomImagesRequest {
	s.InstanceType = &v
	return s
}

type ListCustomImagesResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Images    []*ListCustomImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
}

func (s ListCustomImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBody) SetRequestId(v string) *ListCustomImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomImagesResponseBody) SetImages(v []*ListCustomImagesResponseBodyImages) *ListCustomImagesResponseBody {
	s.Images = v
	return s
}

type ListCustomImagesResponseBodyImages struct {
	Status            *string                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	PostInstallScript *string                                      `json:"PostInstallScript,omitempty" xml:"PostInstallScript,omitempty"`
	ImageOwnerAlias   *string                                      `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	OsTag             *ListCustomImagesResponseBodyImagesOsTag     `json:"OsTag,omitempty" xml:"OsTag,omitempty" type:"Struct"`
	SkuCode           *string                                      `json:"SkuCode,omitempty" xml:"SkuCode,omitempty"`
	PricingCycle      *string                                      `json:"PricingCycle,omitempty" xml:"PricingCycle,omitempty"`
	Description       *string                                      `json:"Description,omitempty" xml:"Description,omitempty"`
	Size              *int32                                       `json:"Size,omitempty" xml:"Size,omitempty"`
	BaseOsTag         *ListCustomImagesResponseBodyImagesBaseOsTag `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty" type:"Struct"`
	ImageName         *string                                      `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	ImageId           *string                                      `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Uid               *string                                      `json:"Uid,omitempty" xml:"Uid,omitempty"`
	ProductCode       *string                                      `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s ListCustomImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImages) SetStatus(v string) *ListCustomImagesResponseBodyImages {
	s.Status = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetPostInstallScript(v string) *ListCustomImagesResponseBodyImages {
	s.PostInstallScript = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetImageOwnerAlias(v string) *ListCustomImagesResponseBodyImages {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetOsTag(v *ListCustomImagesResponseBodyImagesOsTag) *ListCustomImagesResponseBodyImages {
	s.OsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetSkuCode(v string) *ListCustomImagesResponseBodyImages {
	s.SkuCode = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetPricingCycle(v string) *ListCustomImagesResponseBodyImages {
	s.PricingCycle = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetDescription(v string) *ListCustomImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetSize(v int32) *ListCustomImagesResponseBodyImages {
	s.Size = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetBaseOsTag(v *ListCustomImagesResponseBodyImagesBaseOsTag) *ListCustomImagesResponseBodyImages {
	s.BaseOsTag = v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetImageName(v string) *ListCustomImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetImageId(v string) *ListCustomImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetUid(v string) *ListCustomImagesResponseBodyImages {
	s.Uid = &v
	return s
}

func (s *ListCustomImagesResponseBodyImages) SetProductCode(v string) *ListCustomImagesResponseBodyImages {
	s.ProductCode = &v
	return s
}

type ListCustomImagesResponseBodyImagesOsTag struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
	BaseOsTag    *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesOsTag {
	s.Version = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesOsTag) SetBaseOsTag(v string) *ListCustomImagesResponseBodyImagesOsTag {
	s.BaseOsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesOsTag {
	s.Architecture = &v
	return s
}

type ListCustomImagesResponseBodyImagesBaseOsTag struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
}

func (s ListCustomImagesResponseBodyImagesBaseOsTag) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponseBodyImagesBaseOsTag) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponseBodyImagesBaseOsTag) SetVersion(v string) *ListCustomImagesResponseBodyImagesBaseOsTag {
	s.Version = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesBaseOsTag) SetPlatform(v string) *ListCustomImagesResponseBodyImagesBaseOsTag {
	s.Platform = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesBaseOsTag) SetOsTag(v string) *ListCustomImagesResponseBodyImagesBaseOsTag {
	s.OsTag = &v
	return s
}

func (s *ListCustomImagesResponseBodyImagesBaseOsTag) SetArchitecture(v string) *ListCustomImagesResponseBodyImagesBaseOsTag {
	s.Architecture = &v
	return s
}

type ListCustomImagesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCustomImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponse) SetHeaders(v map[string]*string) *ListCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImagesResponse) SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse {
	s.Body = v
	return s
}

type ListFileSystemWithMountTargetsRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFileSystemWithMountTargetsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsRequest) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsRequest) SetPageNumber(v int32) *ListFileSystemWithMountTargetsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFileSystemWithMountTargetsRequest) SetPageSize(v int32) *ListFileSystemWithMountTargetsRequest {
	s.PageSize = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBody struct {
	TotalCount     *int32                                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize       *int32                                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId      *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber     *int32                                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	FileSystemList []*ListFileSystemWithMountTargetsResponseBodyFileSystemList `json:"FileSystemList,omitempty" xml:"FileSystemList,omitempty" type:"Repeated"`
}

func (s ListFileSystemWithMountTargetsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetTotalCount(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetPageSize(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetRequestId(v string) *ListFileSystemWithMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetPageNumber(v int32) *ListFileSystemWithMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBody) SetFileSystemList(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemList) *ListFileSystemWithMountTargetsResponseBody {
	s.FileSystemList = v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemList struct {
	Status          *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Capacity        *int32                                                                     `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	MountTargetList []*ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList `json:"MountTargetList,omitempty" xml:"MountTargetList,omitempty" type:"Repeated"`
	CreateTime      *string                                                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PackageList     []*ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList     `json:"PackageList,omitempty" xml:"PackageList,omitempty" type:"Repeated"`
	StorageType     *string                                                                    `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	BandWidth       *int32                                                                     `json:"BandWidth,omitempty" xml:"BandWidth,omitempty"`
	RegionId        *string                                                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	FileSystemType  *string                                                                    `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	FileSystemId    *string                                                                    `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	MeteredSize     *int32                                                                     `json:"MeteredSize,omitempty" xml:"MeteredSize,omitempty"`
	EncryptType     *int32                                                                     `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	ProtocolType    *string                                                                    `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	Destription     *string                                                                    `json:"Destription,omitempty" xml:"Destription,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetStatus(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.Status = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetCapacity(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.Capacity = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetMountTargetList(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.MountTargetList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetCreateTime(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.CreateTime = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetPackageList(v []*ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.PackageList = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetStorageType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.StorageType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetBandWidth(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.BandWidth = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetRegionId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.RegionId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetFileSystemType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.FileSystemType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetFileSystemId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.FileSystemId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetMeteredSize(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.MeteredSize = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetEncryptType(v int32) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.EncryptType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetProtocolType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.ProtocolType = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemList) SetDestription(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemList {
	s.Destription = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList struct {
	Status            *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId             *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	AccessGroup       *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	VswId             *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
	NetworkType       *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) SetStatus(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList {
	s.Status = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) SetVpcId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList {
	s.VpcId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) SetMountTargetDomain(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList {
	s.MountTargetDomain = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) SetAccessGroup(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList {
	s.AccessGroup = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) SetVswId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList {
	s.VswId = &v
	return s
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList) SetNetworkType(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListMountTargetList {
	s.NetworkType = &v
	return s
}

type ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList struct {
	PackageId *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList) SetPackageId(v string) *ListFileSystemWithMountTargetsResponseBodyFileSystemListPackageList {
	s.PackageId = &v
	return s
}

type ListFileSystemWithMountTargetsResponse struct {
	Headers map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFileSystemWithMountTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFileSystemWithMountTargetsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileSystemWithMountTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListFileSystemWithMountTargetsResponse) SetHeaders(v map[string]*string) *ListFileSystemWithMountTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListFileSystemWithMountTargetsResponse) SetBody(v *ListFileSystemWithMountTargetsResponseBody) *ListFileSystemWithMountTargetsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	BaseOsTag    *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetBaseOsTag(v string) *ListImagesRequest {
	s.BaseOsTag = &v
	return s
}

func (s *ListImagesRequest) SetInstanceType(v string) *ListImagesRequest {
	s.InstanceType = &v
	return s
}

type ListImagesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	OsTags    []*ListImagesResponseBodyOsTags `json:"OsTags,omitempty" xml:"OsTags,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetOsTags(v []*ListImagesResponseBodyOsTags) *ListImagesResponseBody {
	s.OsTags = v
	return s
}

type ListImagesResponseBodyOsTags struct {
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
	BaseOsTag    *string `json:"BaseOsTag,omitempty" xml:"BaseOsTag,omitempty"`
	Platform     *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	OsTag        *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	ImageId      *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
}

func (s ListImagesResponseBodyOsTags) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyOsTags) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyOsTags) SetVersion(v string) *ListImagesResponseBodyOsTags {
	s.Version = &v
	return s
}

func (s *ListImagesResponseBodyOsTags) SetBaseOsTag(v string) *ListImagesResponseBodyOsTags {
	s.BaseOsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTags) SetPlatform(v string) *ListImagesResponseBodyOsTags {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyOsTags) SetOsTag(v string) *ListImagesResponseBodyOsTags {
	s.OsTag = &v
	return s
}

func (s *ListImagesResponseBodyOsTags) SetImageId(v string) *ListImagesResponseBodyOsTags {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyOsTags) SetArchitecture(v string) *ListImagesResponseBodyOsTags {
	s.Architecture = &v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstalledSoftwareRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListInstalledSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareRequest) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareRequest) SetClusterId(v string) *ListInstalledSoftwareRequest {
	s.ClusterId = &v
	return s
}

type ListInstalledSoftwareResponseBody struct {
	RequestId    *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SoftwareList []*ListInstalledSoftwareResponseBodySoftwareList `json:"SoftwareList,omitempty" xml:"SoftwareList,omitempty" type:"Repeated"`
}

func (s ListInstalledSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBody) SetRequestId(v string) *ListInstalledSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstalledSoftwareResponseBody) SetSoftwareList(v []*ListInstalledSoftwareResponseBodySoftwareList) *ListInstalledSoftwareResponseBody {
	s.SoftwareList = v
	return s
}

type ListInstalledSoftwareResponseBodySoftwareList struct {
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
	SoftwareName    *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	SoftwareId      *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	SoftwareStatus  *string `json:"SoftwareStatus,omitempty" xml:"SoftwareStatus,omitempty"`
}

func (s ListInstalledSoftwareResponseBodySoftwareList) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponseBodySoftwareList) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponseBodySoftwareList) SetSoftwareVersion(v string) *ListInstalledSoftwareResponseBodySoftwareList {
	s.SoftwareVersion = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareList) SetSoftwareName(v string) *ListInstalledSoftwareResponseBodySoftwareList {
	s.SoftwareName = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareList) SetSoftwareId(v string) *ListInstalledSoftwareResponseBodySoftwareList {
	s.SoftwareId = &v
	return s
}

func (s *ListInstalledSoftwareResponseBodySoftwareList) SetSoftwareStatus(v string) *ListInstalledSoftwareResponseBodySoftwareList {
	s.SoftwareStatus = &v
	return s
}

type ListInstalledSoftwareResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInstalledSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstalledSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstalledSoftwareResponse) GoString() string {
	return s.String()
}

func (s *ListInstalledSoftwareResponse) SetHeaders(v map[string]*string) *ListInstalledSoftwareResponse {
	s.Headers = v
	return s
}

func (s *ListInstalledSoftwareResponse) SetBody(v *ListInstalledSoftwareResponseBody) *ListInstalledSoftwareResponse {
	s.Body = v
	return s
}

type ListInvocationResultsRequest struct {
	ClusterId          *string                                 `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandId          *string                                 `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InvokeRecordStatus *string                                 `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	PageNumber         *int32                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Instance           []*ListInvocationResultsRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListInvocationResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsRequest) SetClusterId(v string) *ListInvocationResultsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInvocationResultsRequest) SetCommandId(v string) *ListInvocationResultsRequest {
	s.CommandId = &v
	return s
}

func (s *ListInvocationResultsRequest) SetInvokeRecordStatus(v string) *ListInvocationResultsRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *ListInvocationResultsRequest) SetPageNumber(v int32) *ListInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInvocationResultsRequest) SetPageSize(v int32) *ListInvocationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListInvocationResultsRequest) SetInstance(v []*ListInvocationResultsRequestInstance) *ListInvocationResultsRequest {
	s.Instance = v
	return s
}

type ListInvocationResultsRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListInvocationResultsRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsRequestInstance) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsRequestInstance) SetId(v string) *ListInvocationResultsRequestInstance {
	s.Id = &v
	return s
}

type ListInvocationResultsResponseBody struct {
	InvocationResults []*ListInvocationResultsResponseBodyInvocationResults `json:"InvocationResults,omitempty" xml:"InvocationResults,omitempty" type:"Repeated"`
	TotalCount        *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize          *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId         *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber        *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListInvocationResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBody) SetInvocationResults(v []*ListInvocationResultsResponseBodyInvocationResults) *ListInvocationResultsResponseBody {
	s.InvocationResults = v
	return s
}

func (s *ListInvocationResultsResponseBody) SetTotalCount(v int32) *ListInvocationResultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetPageSize(v int32) *ListInvocationResultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetRequestId(v string) *ListInvocationResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvocationResultsResponseBody) SetPageNumber(v int32) *ListInvocationResultsResponseBody {
	s.PageNumber = &v
	return s
}

type ListInvocationResultsResponseBodyInvocationResults struct {
	Success            *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Message            *string `json:"Message,omitempty" xml:"Message,omitempty"`
	FinishedTime       *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	CommandId          *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	ExitCode           *int32  `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
}

func (s ListInvocationResultsResponseBodyInvocationResults) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponseBodyInvocationResults) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetSuccess(v bool) *ListInvocationResultsResponseBodyInvocationResults {
	s.Success = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetMessage(v string) *ListInvocationResultsResponseBodyInvocationResults {
	s.Message = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetFinishedTime(v string) *ListInvocationResultsResponseBodyInvocationResults {
	s.FinishedTime = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetCommandId(v string) *ListInvocationResultsResponseBodyInvocationResults {
	s.CommandId = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetInstanceId(v string) *ListInvocationResultsResponseBodyInvocationResults {
	s.InstanceId = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetInvokeRecordStatus(v string) *ListInvocationResultsResponseBodyInvocationResults {
	s.InvokeRecordStatus = &v
	return s
}

func (s *ListInvocationResultsResponseBodyInvocationResults) SetExitCode(v int32) *ListInvocationResultsResponseBodyInvocationResults {
	s.ExitCode = &v
	return s
}

type ListInvocationResultsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInvocationResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationResultsResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationResultsResponse) SetHeaders(v map[string]*string) *ListInvocationResultsResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationResultsResponse) SetBody(v *ListInvocationResultsResponseBody) *ListInvocationResultsResponse {
	s.Body = v
	return s
}

type ListInvocationStatusRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
}

func (s ListInvocationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusRequest) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusRequest) SetClusterId(v string) *ListInvocationStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *ListInvocationStatusRequest) SetCommandId(v string) *ListInvocationStatusRequest {
	s.CommandId = &v
	return s
}

type ListInvocationStatusResponseBody struct {
	InvokeStatus    *string                                            `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	RequestId       *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	CommandId       *string                                            `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	InvokeInstances []*ListInvocationStatusResponseBodyInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Repeated"`
}

func (s ListInvocationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBody) SetInvokeStatus(v string) *ListInvocationStatusResponseBody {
	s.InvokeStatus = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetRequestId(v string) *ListInvocationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetCommandId(v string) *ListInvocationStatusResponseBody {
	s.CommandId = &v
	return s
}

func (s *ListInvocationStatusResponseBody) SetInvokeInstances(v []*ListInvocationStatusResponseBodyInvokeInstances) *ListInvocationStatusResponseBody {
	s.InvokeInstances = v
	return s
}

type ListInvocationStatusResponseBodyInvokeInstances struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceInvokeStatus *string `json:"InstanceInvokeStatus,omitempty" xml:"InstanceInvokeStatus,omitempty"`
}

func (s ListInvocationStatusResponseBodyInvokeInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponseBodyInvokeInstances) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponseBodyInvokeInstances) SetInstanceId(v string) *ListInvocationStatusResponseBodyInvokeInstances {
	s.InstanceId = &v
	return s
}

func (s *ListInvocationStatusResponseBodyInvokeInstances) SetInstanceInvokeStatus(v string) *ListInvocationStatusResponseBodyInvokeInstances {
	s.InstanceInvokeStatus = &v
	return s
}

type ListInvocationStatusResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListInvocationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInvocationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInvocationStatusResponse) GoString() string {
	return s.String()
}

func (s *ListInvocationStatusResponse) SetHeaders(v map[string]*string) *ListInvocationStatusResponse {
	s.Headers = v
	return s
}

func (s *ListInvocationStatusResponse) SetBody(v *ListInvocationStatusResponseBody) *ListInvocationStatusResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Owner      *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	Rerunable  *string `json:"Rerunable,omitempty" xml:"Rerunable,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetClusterId(v string) *ListJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListJobsRequest) SetOwner(v string) *ListJobsRequest {
	s.Owner = &v
	return s
}

func (s *ListJobsRequest) SetState(v string) *ListJobsRequest {
	s.State = &v
	return s
}

func (s *ListJobsRequest) SetRerunable(v string) *ListJobsRequest {
	s.Rerunable = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

type ListJobsResponseBody struct {
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Jobs       []*ListJobsResponseBodyJobs `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetTotalCount(v int32) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) SetPageSize(v int32) *ListJobsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetPageNumber(v int32) *ListJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobsResponseBody) SetJobs(v []*ListJobsResponseBodyJobs) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

type ListJobsResponseBodyJobs struct {
	Owner          *string                            `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Comment        *string                            `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Stderr         *string                            `json:"Stderr,omitempty" xml:"Stderr,omitempty"`
	State          *string                            `json:"State,omitempty" xml:"State,omitempty"`
	Priority       *string                            `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ShellPath      *string                            `json:"ShellPath,omitempty" xml:"ShellPath,omitempty"`
	Stdout         *string                            `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
	ArrayRequest   *string                            `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	StartTime      *string                            `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	LastModifyTime *string                            `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	NodeList       *string                            `json:"NodeList,omitempty" xml:"NodeList,omitempty"`
	Name           *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Id             *string                            `json:"Id,omitempty" xml:"Id,omitempty"`
	SubmitTime     *string                            `json:"SubmitTime,omitempty" xml:"SubmitTime,omitempty"`
	Resources      *ListJobsResponseBodyJobsResources `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Struct"`
}

func (s ListJobsResponseBodyJobs) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobs) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobs) SetOwner(v string) *ListJobsResponseBodyJobs {
	s.Owner = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetComment(v string) *ListJobsResponseBodyJobs {
	s.Comment = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStderr(v string) *ListJobsResponseBodyJobs {
	s.Stderr = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetState(v string) *ListJobsResponseBodyJobs {
	s.State = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetPriority(v string) *ListJobsResponseBodyJobs {
	s.Priority = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetShellPath(v string) *ListJobsResponseBodyJobs {
	s.ShellPath = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStdout(v string) *ListJobsResponseBodyJobs {
	s.Stdout = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetArrayRequest(v string) *ListJobsResponseBodyJobs {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetStartTime(v string) *ListJobsResponseBodyJobs {
	s.StartTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetLastModifyTime(v string) *ListJobsResponseBodyJobs {
	s.LastModifyTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetNodeList(v string) *ListJobsResponseBodyJobs {
	s.NodeList = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetName(v string) *ListJobsResponseBodyJobs {
	s.Name = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetId(v string) *ListJobsResponseBodyJobs {
	s.Id = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetSubmitTime(v string) *ListJobsResponseBodyJobs {
	s.SubmitTime = &v
	return s
}

func (s *ListJobsResponseBodyJobs) SetResources(v *ListJobsResponseBodyJobsResources) *ListJobsResponseBodyJobs {
	s.Resources = v
	return s
}

type ListJobsResponseBodyJobsResources struct {
	Nodes *int32 `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	Cores *int32 `json:"Cores,omitempty" xml:"Cores,omitempty"`
}

func (s ListJobsResponseBodyJobsResources) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBodyJobsResources) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBodyJobsResources) SetNodes(v int32) *ListJobsResponseBodyJobsResources {
	s.Nodes = &v
	return s
}

func (s *ListJobsResponseBodyJobsResources) SetCores(v int32) *ListJobsResponseBodyJobsResources {
	s.Cores = &v
	return s
}

type ListJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListJobTemplatesRequest struct {
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListJobTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesRequest) SetName(v string) *ListJobTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageNumber(v int32) *ListJobTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesRequest) SetPageSize(v int32) *ListJobTemplatesRequest {
	s.PageSize = &v
	return s
}

type ListJobTemplatesResponseBody struct {
	TotalCount *int32                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Templates  []*ListJobTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s ListJobTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBody) SetTotalCount(v int32) *ListJobTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageSize(v int32) *ListJobTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetRequestId(v string) *ListJobTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetPageNumber(v int32) *ListJobTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListJobTemplatesResponseBody) SetTemplates(v []*ListJobTemplatesResponseBodyTemplates) *ListJobTemplatesResponseBody {
	s.Templates = v
	return s
}

type ListJobTemplatesResponseBodyTemplates struct {
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	Id                 *string `json:"Id,omitempty" xml:"Id,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
}

func (s ListJobTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponseBodyTemplates) SetStdoutRedirectPath(v string) *ListJobTemplatesResponseBodyTemplates {
	s.StdoutRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetVariables(v string) *ListJobTemplatesResponseBodyTemplates {
	s.Variables = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetCommandLine(v string) *ListJobTemplatesResponseBodyTemplates {
	s.CommandLine = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetPackagePath(v string) *ListJobTemplatesResponseBodyTemplates {
	s.PackagePath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetPriority(v int32) *ListJobTemplatesResponseBodyTemplates {
	s.Priority = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetReRunable(v bool) *ListJobTemplatesResponseBodyTemplates {
	s.ReRunable = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetName(v string) *ListJobTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetArrayRequest(v string) *ListJobTemplatesResponseBodyTemplates {
	s.ArrayRequest = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetId(v string) *ListJobTemplatesResponseBodyTemplates {
	s.Id = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetStderrRedirectPath(v string) *ListJobTemplatesResponseBodyTemplates {
	s.StderrRedirectPath = &v
	return s
}

func (s *ListJobTemplatesResponseBodyTemplates) SetRunasUser(v string) *ListJobTemplatesResponseBodyTemplates {
	s.RunasUser = &v
	return s
}

type ListJobTemplatesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListJobTemplatesResponse) SetHeaders(v map[string]*string) *ListJobTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListJobTemplatesResponse) SetBody(v *ListJobTemplatesResponseBody) *ListJobTemplatesResponse {
	s.Body = v
	return s
}

type ListNodesRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Role       *string `json:"Role,omitempty" xml:"Role,omitempty"`
	HostName   *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Sequence   *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	Filter     *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
}

func (s ListNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesRequest) GoString() string {
	return s.String()
}

func (s *ListNodesRequest) SetClusterId(v string) *ListNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesRequest) SetRole(v string) *ListNodesRequest {
	s.Role = &v
	return s
}

func (s *ListNodesRequest) SetHostName(v string) *ListNodesRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesRequest) SetPageNumber(v int32) *ListNodesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesRequest) SetPageSize(v int32) *ListNodesRequest {
	s.PageSize = &v
	return s
}

func (s *ListNodesRequest) SetSequence(v string) *ListNodesRequest {
	s.Sequence = &v
	return s
}

func (s *ListNodesRequest) SetSortBy(v string) *ListNodesRequest {
	s.SortBy = &v
	return s
}

func (s *ListNodesRequest) SetFilter(v string) *ListNodesRequest {
	s.Filter = &v
	return s
}

type ListNodesResponseBody struct {
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Nodes      []*ListNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ListNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBody) SetTotalCount(v int32) *ListNodesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodesResponseBody) SetPageSize(v int32) *ListNodesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesResponseBody) SetRequestId(v string) *ListNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesResponseBody) SetPageNumber(v int32) *ListNodesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesResponseBody) SetNodes(v []*ListNodesResponseBodyNodes) *ListNodesResponseBody {
	s.Nodes = v
	return s
}

type ListNodesResponseBodyNodes struct {
	Status          *string                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId           *string                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	Expired         *bool                                     `json:"Expired,omitempty" xml:"Expired,omitempty"`
	UsedResources   *ListNodesResponseBodyNodesUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	SpotStrategy    *string                                   `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	PublicIpAddress *string                                   `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	CreatedByEhpc   *bool                                     `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	IpAddress       *string                                   `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Version         *string                                   `json:"Version,omitempty" xml:"Version,omitempty"`
	AddTime         *string                                   `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	ImageId         *string                                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	CreateMode      *string                                   `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
	HtEnabled       *bool                                     `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	ImageOwnerAlias *string                                   `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	Roles           []*string                                 `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	LockReason      *string                                   `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	HostName        *string                                   `json:"HostName,omitempty" xml:"HostName,omitempty"`
	RegionId        *string                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TotalResources  *ListNodesResponseBodyNodesTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	VSwitchId       *string                                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ExpiredTime     *string                                   `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ZoneId          *string                                   `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	Location        *string                                   `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *string                                   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodes) SetStatus(v string) *ListNodesResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVpcId(v string) *ListNodesResponseBodyNodes {
	s.VpcId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetExpired(v bool) *ListNodesResponseBodyNodes {
	s.Expired = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetUsedResources(v *ListNodesResponseBodyNodesUsedResources) *ListNodesResponseBodyNodes {
	s.UsedResources = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetSpotStrategy(v string) *ListNodesResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetPublicIpAddress(v string) *ListNodesResponseBodyNodes {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetCreatedByEhpc(v bool) *ListNodesResponseBodyNodes {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetIpAddress(v string) *ListNodesResponseBodyNodes {
	s.IpAddress = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVersion(v string) *ListNodesResponseBodyNodes {
	s.Version = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetAddTime(v string) *ListNodesResponseBodyNodes {
	s.AddTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetImageId(v string) *ListNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetCreateMode(v string) *ListNodesResponseBodyNodes {
	s.CreateMode = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHtEnabled(v bool) *ListNodesResponseBodyNodes {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetImageOwnerAlias(v string) *ListNodesResponseBodyNodes {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetRoles(v []*string) *ListNodesResponseBodyNodes {
	s.Roles = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetLockReason(v string) *ListNodesResponseBodyNodes {
	s.LockReason = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetHostName(v string) *ListNodesResponseBodyNodes {
	s.HostName = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetRegionId(v string) *ListNodesResponseBodyNodes {
	s.RegionId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetTotalResources(v *ListNodesResponseBodyNodesTotalResources) *ListNodesResponseBodyNodes {
	s.TotalResources = v
	return s
}

func (s *ListNodesResponseBodyNodes) SetVSwitchId(v string) *ListNodesResponseBodyNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetExpiredTime(v string) *ListNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetZoneId(v string) *ListNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetLocation(v string) *ListNodesResponseBodyNodes {
	s.Location = &v
	return s
}

func (s *ListNodesResponseBodyNodes) SetId(v string) *ListNodesResponseBodyNodes {
	s.Id = &v
	return s
}

type ListNodesResponseBodyNodesUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesUsedResources) SetCpu(v int32) *ListNodesResponseBodyNodesUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesUsedResources) SetGpu(v int32) *ListNodesResponseBodyNodesUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesUsedResources) SetMemory(v int32) *ListNodesResponseBodyNodesUsedResources {
	s.Memory = &v
	return s
}

type ListNodesResponseBodyNodesTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesResponseBodyNodesTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponseBodyNodesTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesResponseBodyNodesTotalResources) SetCpu(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) SetGpu(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesResponseBodyNodesTotalResources) SetMemory(v int32) *ListNodesResponseBodyNodesTotalResources {
	s.Memory = &v
	return s
}

type ListNodesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesResponse) GoString() string {
	return s.String()
}

func (s *ListNodesResponse) SetHeaders(v map[string]*string) *ListNodesResponse {
	s.Headers = v
	return s
}

func (s *ListNodesResponse) SetBody(v *ListNodesResponseBody) *ListNodesResponse {
	s.Body = v
	return s
}

type ListNodesByQueueRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName  *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListNodesByQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueRequest) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueRequest) SetClusterId(v string) *ListNodesByQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesByQueueRequest) SetQueueName(v string) *ListNodesByQueueRequest {
	s.QueueName = &v
	return s
}

func (s *ListNodesByQueueRequest) SetPageNumber(v int32) *ListNodesByQueueRequest {
	s.PageNumber = &v
	return s
}

func (s *ListNodesByQueueRequest) SetPageSize(v int32) *ListNodesByQueueRequest {
	s.PageSize = &v
	return s
}

type ListNodesByQueueResponseBody struct {
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                               `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                               `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Nodes      []*ListNodesByQueueResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ListNodesByQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBody) SetTotalCount(v int32) *ListNodesByQueueResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetPageSize(v int32) *ListNodesByQueueResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetRequestId(v string) *ListNodesByQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetPageNumber(v int32) *ListNodesByQueueResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesByQueueResponseBody) SetNodes(v []*ListNodesByQueueResponseBodyNodes) *ListNodesByQueueResponseBody {
	s.Nodes = v
	return s
}

type ListNodesByQueueResponseBodyNodes struct {
	Status          *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId           *string                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	HtEnabled       *bool                                            `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	Expired         *bool                                            `json:"Expired,omitempty" xml:"Expired,omitempty"`
	ImageOwnerAlias *string                                          `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	HostName        *string                                          `json:"HostName,omitempty" xml:"HostName,omitempty"`
	LockReason      *string                                          `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	UsedResources   *ListNodesByQueueResponseBodyNodesUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	SpotStrategy    *string                                          `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	PublicIpAddress *string                                          `json:"PublicIpAddress,omitempty" xml:"PublicIpAddress,omitempty"`
	RegionId        *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	CreatedByEhpc   *bool                                            `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	VSwitchId       *string                                          `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	TotalResources  *ListNodesByQueueResponseBodyNodesTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	IpAddress       *string                                          `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	ExpiredTime     *string                                          `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	Version         *string                                          `json:"Version,omitempty" xml:"Version,omitempty"`
	ZoneId          *string                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	AddTime         *string                                          `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	ImageId         *string                                          `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Location        *string                                          `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	CreateMode      *string                                          `json:"CreateMode,omitempty" xml:"CreateMode,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodes) SetStatus(v string) *ListNodesByQueueResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetVpcId(v string) *ListNodesByQueueResponseBodyNodes {
	s.VpcId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetHtEnabled(v bool) *ListNodesByQueueResponseBodyNodes {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetExpired(v bool) *ListNodesByQueueResponseBodyNodes {
	s.Expired = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetImageOwnerAlias(v string) *ListNodesByQueueResponseBodyNodes {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetHostName(v string) *ListNodesByQueueResponseBodyNodes {
	s.HostName = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetLockReason(v string) *ListNodesByQueueResponseBodyNodes {
	s.LockReason = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetUsedResources(v *ListNodesByQueueResponseBodyNodesUsedResources) *ListNodesByQueueResponseBodyNodes {
	s.UsedResources = v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetSpotStrategy(v string) *ListNodesByQueueResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetPublicIpAddress(v string) *ListNodesByQueueResponseBodyNodes {
	s.PublicIpAddress = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetRegionId(v string) *ListNodesByQueueResponseBodyNodes {
	s.RegionId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetCreatedByEhpc(v bool) *ListNodesByQueueResponseBodyNodes {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetVSwitchId(v string) *ListNodesByQueueResponseBodyNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetTotalResources(v *ListNodesByQueueResponseBodyNodesTotalResources) *ListNodesByQueueResponseBodyNodes {
	s.TotalResources = v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetIpAddress(v string) *ListNodesByQueueResponseBodyNodes {
	s.IpAddress = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetExpiredTime(v string) *ListNodesByQueueResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetVersion(v string) *ListNodesByQueueResponseBodyNodes {
	s.Version = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetZoneId(v string) *ListNodesByQueueResponseBodyNodes {
	s.ZoneId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetAddTime(v string) *ListNodesByQueueResponseBodyNodes {
	s.AddTime = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetImageId(v string) *ListNodesByQueueResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetLocation(v string) *ListNodesByQueueResponseBodyNodes {
	s.Location = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetId(v string) *ListNodesByQueueResponseBodyNodes {
	s.Id = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodes) SetCreateMode(v string) *ListNodesByQueueResponseBodyNodes {
	s.CreateMode = &v
	return s
}

type ListNodesByQueueResponseBodyNodesUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesUsedResources) SetCpu(v int32) *ListNodesByQueueResponseBodyNodesUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesUsedResources) SetGpu(v int32) *ListNodesByQueueResponseBodyNodesUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesUsedResources) SetMemory(v int32) *ListNodesByQueueResponseBodyNodesUsedResources {
	s.Memory = &v
	return s
}

type ListNodesByQueueResponseBodyNodesTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesByQueueResponseBodyNodesTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponseBodyNodesTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponseBodyNodesTotalResources) SetCpu(v int32) *ListNodesByQueueResponseBodyNodesTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesTotalResources) SetGpu(v int32) *ListNodesByQueueResponseBodyNodesTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesByQueueResponseBodyNodesTotalResources) SetMemory(v int32) *ListNodesByQueueResponseBodyNodesTotalResources {
	s.Memory = &v
	return s
}

type ListNodesByQueueResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesByQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesByQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesByQueueResponse) GoString() string {
	return s.String()
}

func (s *ListNodesByQueueResponse) SetHeaders(v map[string]*string) *ListNodesByQueueResponse {
	s.Headers = v
	return s
}

func (s *ListNodesByQueueResponse) SetBody(v *ListNodesByQueueResponseBody) *ListNodesByQueueResponse {
	s.Body = v
	return s
}

type ListNodesNoPagingRequest struct {
	ClusterId    *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	HostName     *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	OnlyDetached *bool   `json:"OnlyDetached,omitempty" xml:"OnlyDetached,omitempty"`
	Sequence     *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListNodesNoPagingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingRequest) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingRequest) SetClusterId(v string) *ListNodesNoPagingRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetRole(v string) *ListNodesNoPagingRequest {
	s.Role = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetHostName(v string) *ListNodesNoPagingRequest {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetOnlyDetached(v bool) *ListNodesNoPagingRequest {
	s.OnlyDetached = &v
	return s
}

func (s *ListNodesNoPagingRequest) SetSequence(v string) *ListNodesNoPagingRequest {
	s.Sequence = &v
	return s
}

type ListNodesNoPagingResponseBody struct {
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Nodes      []*ListNodesNoPagingResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ListNodesNoPagingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBody) SetTotalCount(v int32) *ListNodesNoPagingResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetPageSize(v int32) *ListNodesNoPagingResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetRequestId(v string) *ListNodesNoPagingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetPageNumber(v int32) *ListNodesNoPagingResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListNodesNoPagingResponseBody) SetNodes(v []*ListNodesNoPagingResponseBodyNodes) *ListNodesNoPagingResponseBody {
	s.Nodes = v
	return s
}

type ListNodesNoPagingResponseBodyNodes struct {
	Status          *string                                           `json:"Status,omitempty" xml:"Status,omitempty"`
	HtEnabled       *bool                                             `json:"HtEnabled,omitempty" xml:"HtEnabled,omitempty"`
	Expired         *bool                                             `json:"Expired,omitempty" xml:"Expired,omitempty"`
	Roles           []*string                                         `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	ImageOwnerAlias *string                                           `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	LockReason      *string                                           `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	HostName        *string                                           `json:"HostName,omitempty" xml:"HostName,omitempty"`
	UsedResources   *ListNodesNoPagingResponseBodyNodesUsedResources  `json:"UsedResources,omitempty" xml:"UsedResources,omitempty" type:"Struct"`
	SpotStrategy    *string                                           `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	CreatedByEhpc   *bool                                             `json:"CreatedByEhpc,omitempty" xml:"CreatedByEhpc,omitempty"`
	RegionId        *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TotalResources  *ListNodesNoPagingResponseBodyNodesTotalResources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty" type:"Struct"`
	Version         *string                                           `json:"Version,omitempty" xml:"Version,omitempty"`
	ExpiredTime     *string                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	AddTime         *string                                           `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	ImageId         *string                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Id              *string                                           `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodes) SetStatus(v string) *ListNodesNoPagingResponseBodyNodes {
	s.Status = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetHtEnabled(v bool) *ListNodesNoPagingResponseBodyNodes {
	s.HtEnabled = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetExpired(v bool) *ListNodesNoPagingResponseBodyNodes {
	s.Expired = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetRoles(v []*string) *ListNodesNoPagingResponseBodyNodes {
	s.Roles = v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetImageOwnerAlias(v string) *ListNodesNoPagingResponseBodyNodes {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetLockReason(v string) *ListNodesNoPagingResponseBodyNodes {
	s.LockReason = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetHostName(v string) *ListNodesNoPagingResponseBodyNodes {
	s.HostName = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetUsedResources(v *ListNodesNoPagingResponseBodyNodesUsedResources) *ListNodesNoPagingResponseBodyNodes {
	s.UsedResources = v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetSpotStrategy(v string) *ListNodesNoPagingResponseBodyNodes {
	s.SpotStrategy = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetCreatedByEhpc(v bool) *ListNodesNoPagingResponseBodyNodes {
	s.CreatedByEhpc = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetRegionId(v string) *ListNodesNoPagingResponseBodyNodes {
	s.RegionId = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetTotalResources(v *ListNodesNoPagingResponseBodyNodesTotalResources) *ListNodesNoPagingResponseBodyNodes {
	s.TotalResources = v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetVersion(v string) *ListNodesNoPagingResponseBodyNodes {
	s.Version = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetExpiredTime(v string) *ListNodesNoPagingResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetAddTime(v string) *ListNodesNoPagingResponseBodyNodes {
	s.AddTime = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetImageId(v string) *ListNodesNoPagingResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodes) SetId(v string) *ListNodesNoPagingResponseBodyNodes {
	s.Id = &v
	return s
}

type ListNodesNoPagingResponseBodyNodesUsedResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodesUsedResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesUsedResources) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesUsedResources) SetCpu(v int32) *ListNodesNoPagingResponseBodyNodesUsedResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesUsedResources) SetGpu(v int32) *ListNodesNoPagingResponseBodyNodesUsedResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesUsedResources) SetMemory(v int32) *ListNodesNoPagingResponseBodyNodesUsedResources {
	s.Memory = &v
	return s
}

type ListNodesNoPagingResponseBodyNodesTotalResources struct {
	Cpu    *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Gpu    *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s ListNodesNoPagingResponseBodyNodesTotalResources) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponseBodyNodesTotalResources) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponseBodyNodesTotalResources) SetCpu(v int32) *ListNodesNoPagingResponseBodyNodesTotalResources {
	s.Cpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesTotalResources) SetGpu(v int32) *ListNodesNoPagingResponseBodyNodesTotalResources {
	s.Gpu = &v
	return s
}

func (s *ListNodesNoPagingResponseBodyNodesTotalResources) SetMemory(v int32) *ListNodesNoPagingResponseBodyNodesTotalResources {
	s.Memory = &v
	return s
}

type ListNodesNoPagingResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListNodesNoPagingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListNodesNoPagingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodesNoPagingResponse) GoString() string {
	return s.String()
}

func (s *ListNodesNoPagingResponse) SetHeaders(v map[string]*string) *ListNodesNoPagingResponse {
	s.Headers = v
	return s
}

func (s *ListNodesNoPagingResponse) SetBody(v *ListNodesNoPagingResponseBody) *ListNodesNoPagingResponse {
	s.Body = v
	return s
}

type ListPreferredEcsTypesRequest struct {
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SpotStrategy       *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
}

func (s ListPreferredEcsTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesRequest) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesRequest) SetZoneId(v string) *ListPreferredEcsTypesRequest {
	s.ZoneId = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetSpotStrategy(v string) *ListPreferredEcsTypesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *ListPreferredEcsTypesRequest) SetInstanceChargeType(v string) *ListPreferredEcsTypesRequest {
	s.InstanceChargeType = &v
	return s
}

type ListPreferredEcsTypesResponseBody struct {
	Series              []*ListPreferredEcsTypesResponseBodySeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	SupportSpotInstance *bool                                      `json:"SupportSpotInstance,omitempty" xml:"SupportSpotInstance,omitempty"`
	RequestId           *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPreferredEcsTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBody) SetSeries(v []*ListPreferredEcsTypesResponseBodySeries) *ListPreferredEcsTypesResponseBody {
	s.Series = v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetSupportSpotInstance(v bool) *ListPreferredEcsTypesResponseBody {
	s.SupportSpotInstance = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBody) SetRequestId(v string) *ListPreferredEcsTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListPreferredEcsTypesResponseBodySeries struct {
	SeriesId   *string                                       `json:"SeriesId,omitempty" xml:"SeriesId,omitempty"`
	SeriesName *string                                       `json:"SeriesName,omitempty" xml:"SeriesName,omitempty"`
	Roles      *ListPreferredEcsTypesResponseBodySeriesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Struct"`
}

func (s ListPreferredEcsTypesResponseBodySeries) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeries) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeries) SetSeriesId(v string) *ListPreferredEcsTypesResponseBodySeries {
	s.SeriesId = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeries) SetSeriesName(v string) *ListPreferredEcsTypesResponseBodySeries {
	s.SeriesName = &v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeries) SetRoles(v *ListPreferredEcsTypesResponseBodySeriesRoles) *ListPreferredEcsTypesResponseBodySeries {
	s.Roles = v
	return s
}

type ListPreferredEcsTypesResponseBodySeriesRoles struct {
	Manager []*string `json:"Manager,omitempty" xml:"Manager,omitempty" type:"Repeated"`
	Compute []*string `json:"Compute,omitempty" xml:"Compute,omitempty" type:"Repeated"`
	Login   []*string `json:"Login,omitempty" xml:"Login,omitempty" type:"Repeated"`
}

func (s ListPreferredEcsTypesResponseBodySeriesRoles) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponseBodySeriesRoles) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponseBodySeriesRoles) SetManager(v []*string) *ListPreferredEcsTypesResponseBodySeriesRoles {
	s.Manager = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesRoles) SetCompute(v []*string) *ListPreferredEcsTypesResponseBodySeriesRoles {
	s.Compute = v
	return s
}

func (s *ListPreferredEcsTypesResponseBodySeriesRoles) SetLogin(v []*string) *ListPreferredEcsTypesResponseBodySeriesRoles {
	s.Login = v
	return s
}

type ListPreferredEcsTypesResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPreferredEcsTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPreferredEcsTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPreferredEcsTypesResponse) GoString() string {
	return s.String()
}

func (s *ListPreferredEcsTypesResponse) SetHeaders(v map[string]*string) *ListPreferredEcsTypesResponse {
	s.Headers = v
	return s
}

func (s *ListPreferredEcsTypesResponse) SetBody(v *ListPreferredEcsTypesResponseBody) *ListPreferredEcsTypesResponse {
	s.Body = v
	return s
}

type ListQueuesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListQueuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesRequest) GoString() string {
	return s.String()
}

func (s *ListQueuesRequest) SetClusterId(v string) *ListQueuesRequest {
	s.ClusterId = &v
	return s
}

type ListQueuesResponseBody struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Queues    []*ListQueuesResponseBodyQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
}

func (s ListQueuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBody) SetRequestId(v string) *ListQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueuesResponseBody) SetQueues(v []*ListQueuesResponseBodyQueues) *ListQueuesResponseBody {
	s.Queues = v
	return s
}

type ListQueuesResponseBodyQueues struct {
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	QueueName           *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ComputeInstanceType *string `json:"ComputeInstanceType,omitempty" xml:"ComputeInstanceType,omitempty"`
}

func (s ListQueuesResponseBodyQueues) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponseBodyQueues) GoString() string {
	return s.String()
}

func (s *ListQueuesResponseBodyQueues) SetType(v string) *ListQueuesResponseBodyQueues {
	s.Type = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetQueueName(v string) *ListQueuesResponseBodyQueues {
	s.QueueName = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetResourceGroupId(v string) *ListQueuesResponseBodyQueues {
	s.ResourceGroupId = &v
	return s
}

func (s *ListQueuesResponseBodyQueues) SetComputeInstanceType(v string) *ListQueuesResponseBodyQueues {
	s.ComputeInstanceType = &v
	return s
}

type ListQueuesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQueuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueuesResponse) GoString() string {
	return s.String()
}

func (s *ListQueuesResponse) SetHeaders(v map[string]*string) *ListQueuesResponse {
	s.Headers = v
	return s
}

func (s *ListQueuesResponse) SetBody(v *ListQueuesResponseBody) *ListQueuesResponse {
	s.Body = v
	return s
}

type ListRegionsResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Regions   []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

type ListRegionsResponseBodyRegions struct {
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListSecurityGroupsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsRequest) SetClusterId(v string) *ListSecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

type ListSecurityGroupsResponseBody struct {
	SecurityGroups []*string `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	TotalCount     *int32    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBody) SetSecurityGroups(v []*string) *ListSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetTotalCount(v int32) *ListSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetRequestId(v string) *ListSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListSecurityGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityGroupsResponse) SetBody(v *ListSecurityGroupsResponseBody) *ListSecurityGroupsResponse {
	s.Body = v
	return s
}

type ListSoftwaresRequest struct {
	EhpcVersion *string `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
	OsTag       *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
}

func (s ListSoftwaresRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresRequest) GoString() string {
	return s.String()
}

func (s *ListSoftwaresRequest) SetEhpcVersion(v string) *ListSoftwaresRequest {
	s.EhpcVersion = &v
	return s
}

func (s *ListSoftwaresRequest) SetOsTag(v string) *ListSoftwaresRequest {
	s.OsTag = &v
	return s
}

type ListSoftwaresResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Softwares []*ListSoftwaresResponseBodySoftwares `json:"Softwares,omitempty" xml:"Softwares,omitempty" type:"Repeated"`
}

func (s ListSoftwaresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBody) SetRequestId(v string) *ListSoftwaresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwaresResponseBody) SetSoftwares(v []*ListSoftwaresResponseBodySoftwares) *ListSoftwaresResponseBody {
	s.Softwares = v
	return s
}

type ListSoftwaresResponseBodySoftwares struct {
	SchedulerType    *string                                           `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	OsTag            *string                                           `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	SchedulerVersion *string                                           `json:"SchedulerVersion,omitempty" xml:"SchedulerVersion,omitempty"`
	AccountVersion   *string                                           `json:"AccountVersion,omitempty" xml:"AccountVersion,omitempty"`
	Applications     []*ListSoftwaresResponseBodySoftwaresApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	AccountType      *string                                           `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	EhpcVersion      *string                                           `json:"EhpcVersion,omitempty" xml:"EhpcVersion,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwares) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwares) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwares) SetSchedulerType(v string) *ListSoftwaresResponseBodySoftwares {
	s.SchedulerType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwares) SetOsTag(v string) *ListSoftwaresResponseBodySoftwares {
	s.OsTag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwares) SetSchedulerVersion(v string) *ListSoftwaresResponseBodySoftwares {
	s.SchedulerVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwares) SetAccountVersion(v string) *ListSoftwaresResponseBodySoftwares {
	s.AccountVersion = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwares) SetApplications(v []*ListSoftwaresResponseBodySoftwaresApplications) *ListSoftwaresResponseBodySoftwares {
	s.Applications = v
	return s
}

func (s *ListSoftwaresResponseBodySoftwares) SetAccountType(v string) *ListSoftwaresResponseBodySoftwares {
	s.AccountType = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwares) SetEhpcVersion(v string) *ListSoftwaresResponseBodySoftwares {
	s.EhpcVersion = &v
	return s
}

type ListSoftwaresResponseBodySoftwaresApplications struct {
	Required *bool   `json:"Required,omitempty" xml:"Required,omitempty"`
	Version  *string `json:"Version,omitempty" xml:"Version,omitempty"`
	Tag      *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListSoftwaresResponseBodySoftwaresApplications) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponseBodySoftwaresApplications) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponseBodySoftwaresApplications) SetRequired(v bool) *ListSoftwaresResponseBodySoftwaresApplications {
	s.Required = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresApplications) SetVersion(v string) *ListSoftwaresResponseBodySoftwaresApplications {
	s.Version = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresApplications) SetTag(v string) *ListSoftwaresResponseBodySoftwaresApplications {
	s.Tag = &v
	return s
}

func (s *ListSoftwaresResponseBodySoftwaresApplications) SetName(v string) *ListSoftwaresResponseBodySoftwaresApplications {
	s.Name = &v
	return s
}

type ListSoftwaresResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSoftwaresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSoftwaresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSoftwaresResponse) GoString() string {
	return s.String()
}

func (s *ListSoftwaresResponse) SetHeaders(v map[string]*string) *ListSoftwaresResponse {
	s.Headers = v
	return s
}

func (s *ListSoftwaresResponse) SetBody(v *ListSoftwaresResponseBody) *ListSoftwaresResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Archived   *bool   `json:"Archived,omitempty" xml:"Archived,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetClusterId(v string) *ListTasksRequest {
	s.ClusterId = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetTaskId(v string) *ListTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListTasksRequest) SetArchived(v bool) *ListTasksRequest {
	s.Archived = &v
	return s
}

type ListTasksResponseBody struct {
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Tasks      []*ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v []*ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetPageSize(v int32) *ListTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetPageNumber(v int32) *ListTasksResponseBody {
	s.PageNumber = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType    *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TotalSteps  *int32  `json:"TotalSteps,omitempty" xml:"TotalSteps,omitempty"`
	CurrentStep *int32  `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	Result      *string `json:"Result,omitempty" xml:"Result,omitempty"`
	Errors      *string `json:"Errors,omitempty" xml:"Errors,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Request     *string `json:"Request,omitempty" xml:"Request,omitempty"`
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetStatus(v string) *ListTasksResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskType(v string) *ListTasksResponseBodyTasks {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTotalSteps(v int32) *ListTasksResponseBodyTasks {
	s.TotalSteps = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetCurrentStep(v int32) *ListTasksResponseBodyTasks {
	s.CurrentStep = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetResult(v string) *ListTasksResponseBodyTasks {
	s.Result = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetErrors(v string) *ListTasksResponseBodyTasks {
	s.Errors = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetTaskId(v string) *ListTasksResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetRequest(v string) *ListTasksResponseBodyTasks {
	s.Request = &v
	return s
}

func (s *ListTasksResponseBodyTasks) SetClusterId(v string) *ListTasksResponseBodyTasks {
	s.ClusterId = &v
	return s
}

type ListTasksResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListUsersRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUsersRequest) GoString() string {
	return s.String()
}

func (s *ListUsersRequest) SetClusterId(v string) *ListUsersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListUsersRequest) SetPageNumber(v int32) *ListUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUsersRequest) SetPageSize(v int32) *ListUsersRequest {
	s.PageSize = &v
	return s
}

type ListUsersResponseBody struct {
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Users      []*ListUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ListUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBody) SetTotalCount(v int32) *ListUsersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListUsersResponseBody) SetPageSize(v int32) *ListUsersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUsersResponseBody) SetRequestId(v string) *ListUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUsersResponseBody) SetPageNumber(v int32) *ListUsersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListUsersResponseBody) SetUsers(v []*ListUsersResponseBodyUsers) *ListUsersResponseBody {
	s.Users = v
	return s
}

type ListUsersResponseBodyUsers struct {
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	AddTime *string `json:"AddTime,omitempty" xml:"AddTime,omitempty"`
	Group   *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s ListUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *ListUsersResponseBodyUsers) SetName(v string) *ListUsersResponseBodyUsers {
	s.Name = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetAddTime(v string) *ListUsersResponseBodyUsers {
	s.AddTime = &v
	return s
}

func (s *ListUsersResponseBodyUsers) SetGroup(v string) *ListUsersResponseBodyUsers {
	s.Group = &v
	return s
}

type ListUsersResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUsersResponse) GoString() string {
	return s.String()
}

func (s *ListUsersResponse) SetHeaders(v map[string]*string) *ListUsersResponse {
	s.Headers = v
	return s
}

func (s *ListUsersResponse) SetBody(v *ListUsersResponseBody) *ListUsersResponse {
	s.Body = v
	return s
}

type ListVolumesRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesRequest) GoString() string {
	return s.String()
}

func (s *ListVolumesRequest) SetPageNumber(v int32) *ListVolumesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVolumesRequest) SetPageSize(v int32) *ListVolumesRequest {
	s.PageSize = &v
	return s
}

type ListVolumesResponseBody struct {
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Volumes    []*ListVolumesResponseBodyVolumes `json:"Volumes,omitempty" xml:"Volumes,omitempty" type:"Repeated"`
	PageSize   *int32                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber *int32                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s ListVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBody) SetTotalCount(v int32) *ListVolumesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVolumesResponseBody) SetVolumes(v []*ListVolumesResponseBodyVolumes) *ListVolumesResponseBody {
	s.Volumes = v
	return s
}

func (s *ListVolumesResponseBody) SetPageSize(v int32) *ListVolumesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListVolumesResponseBody) SetRequestId(v string) *ListVolumesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVolumesResponseBody) SetPageNumber(v int32) *ListVolumesResponseBody {
	s.PageNumber = &v
	return s
}

type ListVolumesResponseBodyVolumes struct {
	VolumeId          *string                                            `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	ClusterName       *string                                            `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	RemoteDirectory   *string                                            `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeMountpoint  *string                                            `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	AdditionalVolumes []*ListVolumesResponseBodyVolumesAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
	VolumeType        *string                                            `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VolumeProtocol    *string                                            `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	RegionId          *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId         *string                                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListVolumesResponseBodyVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeId(v string) *ListVolumesResponseBodyVolumes {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetClusterName(v string) *ListVolumesResponseBodyVolumes {
	s.ClusterName = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetAdditionalVolumes(v []*ListVolumesResponseBodyVolumesAdditionalVolumes) *ListVolumesResponseBodyVolumes {
	s.AdditionalVolumes = v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeType(v string) *ListVolumesResponseBodyVolumes {
	s.VolumeType = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetRegionId(v string) *ListVolumesResponseBodyVolumes {
	s.RegionId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumes) SetClusterId(v string) *ListVolumesResponseBodyVolumes {
	s.ClusterId = &v
	return s
}

type ListVolumesResponseBodyVolumesAdditionalVolumes struct {
	JobQueue         *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	VolumeId         *string `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	RemoteDirectory  *string `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeMountpoint *string `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	Role             *string `json:"Role,omitempty" xml:"Role,omitempty"`
	LocalDirectory   *string `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	VolumeType       *string `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	VolumeProtocol   *string `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
}

func (s ListVolumesResponseBodyVolumesAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponseBodyVolumesAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetJobQueue(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetVolumeId(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetRemoteDirectory(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetVolumeMountpoint(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetRole(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.Role = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetLocalDirectory(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetVolumeType(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.VolumeType = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetLocation(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.Location = &v
	return s
}

func (s *ListVolumesResponseBodyVolumesAdditionalVolumes) SetVolumeProtocol(v string) *ListVolumesResponseBodyVolumesAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

type ListVolumesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVolumesResponse) GoString() string {
	return s.String()
}

func (s *ListVolumesResponse) SetHeaders(v map[string]*string) *ListVolumesResponse {
	s.Headers = v
	return s
}

func (s *ListVolumesResponse) SetBody(v *ListVolumesResponseBody) *ListVolumesResponse {
	s.Body = v
	return s
}

type ModifyClusterAttributesRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Description     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
}

func (s ModifyClusterAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesRequest) SetClusterId(v string) *ModifyClusterAttributesRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetName(v string) *ModifyClusterAttributesRequest {
	s.Name = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetDescription(v string) *ModifyClusterAttributesRequest {
	s.Description = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetImageOwnerAlias(v string) *ModifyClusterAttributesRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *ModifyClusterAttributesRequest) SetImageId(v string) *ModifyClusterAttributesRequest {
	s.ImageId = &v
	return s
}

type ModifyClusterAttributesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponseBody) SetRequestId(v string) *ModifyClusterAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyClusterAttributesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyClusterAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyClusterAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterAttributesResponse) SetHeaders(v map[string]*string) *ModifyClusterAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyClusterAttributesResponse) SetBody(v *ModifyClusterAttributesResponseBody) *ModifyClusterAttributesResponse {
	s.Body = v
	return s
}

type ModifyContainerAppAttributesRequest struct {
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyContainerAppAttributesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesRequest) SetContainerId(v string) *ModifyContainerAppAttributesRequest {
	s.ContainerId = &v
	return s
}

func (s *ModifyContainerAppAttributesRequest) SetDescription(v string) *ModifyContainerAppAttributesRequest {
	s.Description = &v
	return s
}

type ModifyContainerAppAttributesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyContainerAppAttributesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesResponseBody) SetRequestId(v string) *ModifyContainerAppAttributesResponseBody {
	s.RequestId = &v
	return s
}

type ModifyContainerAppAttributesResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyContainerAppAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyContainerAppAttributesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyContainerAppAttributesResponse) GoString() string {
	return s.String()
}

func (s *ModifyContainerAppAttributesResponse) SetHeaders(v map[string]*string) *ModifyContainerAppAttributesResponse {
	s.Headers = v
	return s
}

func (s *ModifyContainerAppAttributesResponse) SetBody(v *ModifyContainerAppAttributesResponseBody) *ModifyContainerAppAttributesResponse {
	s.Body = v
	return s
}

type ModifyImageGatewayConfigRequest struct {
	ClusterId              *string                                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	DBType                 *string                                `json:"DBType,omitempty" xml:"DBType,omitempty"`
	DBUsername             *string                                `json:"DBUsername,omitempty" xml:"DBUsername,omitempty"`
	DBPassword             *string                                `json:"DBPassword,omitempty" xml:"DBPassword,omitempty"`
	DBServerInfo           *string                                `json:"DBServerInfo,omitempty" xml:"DBServerInfo,omitempty"`
	DefaultRepoLocation    *string                                `json:"DefaultRepoLocation,omitempty" xml:"DefaultRepoLocation,omitempty"`
	PullUpdateTimeout      *int32                                 `json:"PullUpdateTimeout,omitempty" xml:"PullUpdateTimeout,omitempty"`
	ImageExpirationTimeout *string                                `json:"ImageExpirationTimeout,omitempty" xml:"ImageExpirationTimeout,omitempty"`
	Repo                   []*ModifyImageGatewayConfigRequestRepo `json:"Repo,omitempty" xml:"Repo,omitempty" type:"Repeated"`
}

func (s ModifyImageGatewayConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigRequest) SetClusterId(v string) *ModifyImageGatewayConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBType(v string) *ModifyImageGatewayConfigRequest {
	s.DBType = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBUsername(v string) *ModifyImageGatewayConfigRequest {
	s.DBUsername = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBPassword(v string) *ModifyImageGatewayConfigRequest {
	s.DBPassword = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDBServerInfo(v string) *ModifyImageGatewayConfigRequest {
	s.DBServerInfo = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetDefaultRepoLocation(v string) *ModifyImageGatewayConfigRequest {
	s.DefaultRepoLocation = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetPullUpdateTimeout(v int32) *ModifyImageGatewayConfigRequest {
	s.PullUpdateTimeout = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetImageExpirationTimeout(v string) *ModifyImageGatewayConfigRequest {
	s.ImageExpirationTimeout = &v
	return s
}

func (s *ModifyImageGatewayConfigRequest) SetRepo(v []*ModifyImageGatewayConfigRequestRepo) *ModifyImageGatewayConfigRequest {
	s.Repo = v
	return s
}

type ModifyImageGatewayConfigRequestRepo struct {
	Auth     *string `json:"Auth,omitempty" xml:"Auth,omitempty"`
	URL      *string `json:"URL,omitempty" xml:"URL,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s ModifyImageGatewayConfigRequestRepo) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigRequestRepo) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigRequestRepo) SetAuth(v string) *ModifyImageGatewayConfigRequestRepo {
	s.Auth = &v
	return s
}

func (s *ModifyImageGatewayConfigRequestRepo) SetURL(v string) *ModifyImageGatewayConfigRequestRepo {
	s.URL = &v
	return s
}

func (s *ModifyImageGatewayConfigRequestRepo) SetLocation(v string) *ModifyImageGatewayConfigRequestRepo {
	s.Location = &v
	return s
}

type ModifyImageGatewayConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyImageGatewayConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigResponseBody) SetRequestId(v string) *ModifyImageGatewayConfigResponseBody {
	s.RequestId = &v
	return s
}

type ModifyImageGatewayConfigResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyImageGatewayConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyImageGatewayConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyImageGatewayConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyImageGatewayConfigResponse) SetHeaders(v map[string]*string) *ModifyImageGatewayConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyImageGatewayConfigResponse) SetBody(v *ModifyImageGatewayConfigResponseBody) *ModifyImageGatewayConfigResponse {
	s.Body = v
	return s
}

type ModifyUserGroupsRequest struct {
	ClusterId *string                        `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserGroupsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequest) SetClusterId(v string) *ModifyUserGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserGroupsRequest) SetUser(v []*ModifyUserGroupsRequestUser) *ModifyUserGroupsRequest {
	s.User = v
	return s
}

type ModifyUserGroupsRequestUser struct {
	Name  *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
}

func (s ModifyUserGroupsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsRequestUser) SetName(v string) *ModifyUserGroupsRequestUser {
	s.Name = &v
	return s
}

func (s *ModifyUserGroupsRequestUser) SetGroup(v string) *ModifyUserGroupsRequestUser {
	s.Group = &v
	return s
}

type ModifyUserGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponseBody) SetRequestId(v string) *ModifyUserGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserGroupsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserGroupsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserGroupsResponse) SetHeaders(v map[string]*string) *ModifyUserGroupsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserGroupsResponse) SetBody(v *ModifyUserGroupsResponseBody) *ModifyUserGroupsResponse {
	s.Body = v
	return s
}

type ModifyUserPasswordsRequest struct {
	ClusterId *string                           `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	User      []*ModifyUserPasswordsRequestUser `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s ModifyUserPasswordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequest) SetClusterId(v string) *ModifyUserPasswordsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUserPasswordsRequest) SetUser(v []*ModifyUserPasswordsRequestUser) *ModifyUserPasswordsRequest {
	s.User = v
	return s
}

type ModifyUserPasswordsRequestUser struct {
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyUserPasswordsRequestUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsRequestUser) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsRequestUser) SetPassword(v string) *ModifyUserPasswordsRequestUser {
	s.Password = &v
	return s
}

func (s *ModifyUserPasswordsRequestUser) SetName(v string) *ModifyUserPasswordsRequestUser {
	s.Name = &v
	return s
}

type ModifyUserPasswordsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserPasswordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponseBody) SetRequestId(v string) *ModifyUserPasswordsResponseBody {
	s.RequestId = &v
	return s
}

type ModifyUserPasswordsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyUserPasswordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyUserPasswordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyUserPasswordsResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserPasswordsResponse) SetHeaders(v map[string]*string) *ModifyUserPasswordsResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserPasswordsResponse) SetBody(v *ModifyUserPasswordsResponseBody) *ModifyUserPasswordsResponse {
	s.Body = v
	return s
}

type ModifyVisualServicePasswdRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RunasUser         *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	Passwd            *string `json:"Passwd,omitempty" xml:"Passwd,omitempty"`
}

func (s ModifyVisualServicePasswdRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdRequest) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdRequest) SetClusterId(v string) *ModifyVisualServicePasswdRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetRunasUser(v string) *ModifyVisualServicePasswdRequest {
	s.RunasUser = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetRunasUserPassword(v string) *ModifyVisualServicePasswdRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *ModifyVisualServicePasswdRequest) SetPasswd(v string) *ModifyVisualServicePasswdRequest {
	s.Passwd = &v
	return s
}

type ModifyVisualServicePasswdResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyVisualServicePasswdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdResponseBody) SetMessage(v string) *ModifyVisualServicePasswdResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyVisualServicePasswdResponseBody) SetRequestId(v string) *ModifyVisualServicePasswdResponseBody {
	s.RequestId = &v
	return s
}

type ModifyVisualServicePasswdResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyVisualServicePasswdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyVisualServicePasswdResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyVisualServicePasswdResponse) GoString() string {
	return s.String()
}

func (s *ModifyVisualServicePasswdResponse) SetHeaders(v map[string]*string) *ModifyVisualServicePasswdResponse {
	s.Headers = v
	return s
}

func (s *ModifyVisualServicePasswdResponse) SetBody(v *ModifyVisualServicePasswdResponseBody) *ModifyVisualServicePasswdResponse {
	s.Body = v
	return s
}

type MountNFSRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NfsDir       *string `json:"NfsDir,omitempty" xml:"NfsDir,omitempty"`
	MountDir     *string `json:"MountDir,omitempty" xml:"MountDir,omitempty"`
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	RemoteDir    *string `json:"RemoteDir,omitempty" xml:"RemoteDir,omitempty"`
}

func (s MountNFSRequest) String() string {
	return tea.Prettify(s)
}

func (s MountNFSRequest) GoString() string {
	return s.String()
}

func (s *MountNFSRequest) SetInstanceId(v string) *MountNFSRequest {
	s.InstanceId = &v
	return s
}

func (s *MountNFSRequest) SetNfsDir(v string) *MountNFSRequest {
	s.NfsDir = &v
	return s
}

func (s *MountNFSRequest) SetMountDir(v string) *MountNFSRequest {
	s.MountDir = &v
	return s
}

func (s *MountNFSRequest) SetProtocolType(v string) *MountNFSRequest {
	s.ProtocolType = &v
	return s
}

func (s *MountNFSRequest) SetRemoteDir(v string) *MountNFSRequest {
	s.RemoteDir = &v
	return s
}

type MountNFSResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InvokeId  *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
}

func (s MountNFSResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MountNFSResponseBody) GoString() string {
	return s.String()
}

func (s *MountNFSResponseBody) SetRequestId(v string) *MountNFSResponseBody {
	s.RequestId = &v
	return s
}

func (s *MountNFSResponseBody) SetInvokeId(v string) *MountNFSResponseBody {
	s.InvokeId = &v
	return s
}

type MountNFSResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MountNFSResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MountNFSResponse) String() string {
	return tea.Prettify(s)
}

func (s MountNFSResponse) GoString() string {
	return s.String()
}

func (s *MountNFSResponse) SetHeaders(v map[string]*string) *MountNFSResponse {
	s.Headers = v
	return s
}

func (s *MountNFSResponse) SetBody(v *MountNFSResponseBody) *MountNFSResponse {
	s.Body = v
	return s
}

type PullImageRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Repository    *string `json:"Repository,omitempty" xml:"Repository,omitempty"`
	ImageTag      *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	ContainerType *string `json:"ContainerType,omitempty" xml:"ContainerType,omitempty"`
}

func (s PullImageRequest) String() string {
	return tea.Prettify(s)
}

func (s PullImageRequest) GoString() string {
	return s.String()
}

func (s *PullImageRequest) SetClusterId(v string) *PullImageRequest {
	s.ClusterId = &v
	return s
}

func (s *PullImageRequest) SetRepository(v string) *PullImageRequest {
	s.Repository = &v
	return s
}

func (s *PullImageRequest) SetImageTag(v string) *PullImageRequest {
	s.ImageTag = &v
	return s
}

func (s *PullImageRequest) SetContainerType(v string) *PullImageRequest {
	s.ContainerType = &v
	return s
}

type PullImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PullImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PullImageResponseBody) GoString() string {
	return s.String()
}

func (s *PullImageResponseBody) SetRequestId(v string) *PullImageResponseBody {
	s.RequestId = &v
	return s
}

type PullImageResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *PullImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PullImageResponse) String() string {
	return tea.Prettify(s)
}

func (s PullImageResponse) GoString() string {
	return s.String()
}

func (s *PullImageResponse) SetHeaders(v map[string]*string) *PullImageResponse {
	s.Headers = v
	return s
}

func (s *PullImageResponse) SetBody(v *PullImageResponseBody) *PullImageResponse {
	s.Body = v
	return s
}

type QueryServicePackAndPriceResponseBody struct {
	OriginalPrice  *float32                                           `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RequestId      *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DiscountPrice  *float32                                           `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	Currency       *string                                            `json:"Currency,omitempty" xml:"Currency,omitempty"`
	ServicePack    []*QueryServicePackAndPriceResponseBodyServicePack `json:"ServicePack,omitempty" xml:"ServicePack,omitempty" type:"Repeated"`
	RegionId       *string                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TradePrice     *float32                                           `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
	OriginalAmount *int32                                             `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	ChargeAmount   *int32                                             `json:"ChargeAmount,omitempty" xml:"ChargeAmount,omitempty"`
}

func (s QueryServicePackAndPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBody) SetOriginalPrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetRequestId(v string) *QueryServicePackAndPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetDiscountPrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetCurrency(v string) *QueryServicePackAndPriceResponseBody {
	s.Currency = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetServicePack(v []*QueryServicePackAndPriceResponseBodyServicePack) *QueryServicePackAndPriceResponseBody {
	s.ServicePack = v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetRegionId(v string) *QueryServicePackAndPriceResponseBody {
	s.RegionId = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetTradePrice(v float32) *QueryServicePackAndPriceResponseBody {
	s.TradePrice = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetOriginalAmount(v int32) *QueryServicePackAndPriceResponseBody {
	s.OriginalAmount = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBody) SetChargeAmount(v int32) *QueryServicePackAndPriceResponseBody {
	s.ChargeAmount = &v
	return s
}

type QueryServicePackAndPriceResponseBodyServicePack struct {
	EndTime      *int32  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Capacity     *int32  `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	StartTime    *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s QueryServicePackAndPriceResponseBodyServicePack) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponseBodyServicePack) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponseBodyServicePack) SetEndTime(v int32) *QueryServicePackAndPriceResponseBodyServicePack {
	s.EndTime = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePack) SetCapacity(v int32) *QueryServicePackAndPriceResponseBodyServicePack {
	s.Capacity = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePack) SetStartTime(v int32) *QueryServicePackAndPriceResponseBodyServicePack {
	s.StartTime = &v
	return s
}

func (s *QueryServicePackAndPriceResponseBodyServicePack) SetInstanceName(v string) *QueryServicePackAndPriceResponseBodyServicePack {
	s.InstanceName = &v
	return s
}

type QueryServicePackAndPriceResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *QueryServicePackAndPriceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s QueryServicePackAndPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryServicePackAndPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryServicePackAndPriceResponse) SetHeaders(v map[string]*string) *QueryServicePackAndPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryServicePackAndPriceResponse) SetBody(v *QueryServicePackAndPriceResponseBody) *QueryServicePackAndPriceResponse {
	s.Body = v
	return s
}

type RecoverClusterRequest struct {
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OsTag           *string `json:"OsTag,omitempty" xml:"OsTag,omitempty"`
	AccountType     *string `json:"AccountType,omitempty" xml:"AccountType,omitempty"`
	SchedulerType   *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	ImageId         *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ClientVersion   *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s RecoverClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterRequest) GoString() string {
	return s.String()
}

func (s *RecoverClusterRequest) SetClusterId(v string) *RecoverClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *RecoverClusterRequest) SetOsTag(v string) *RecoverClusterRequest {
	s.OsTag = &v
	return s
}

func (s *RecoverClusterRequest) SetAccountType(v string) *RecoverClusterRequest {
	s.AccountType = &v
	return s
}

func (s *RecoverClusterRequest) SetSchedulerType(v string) *RecoverClusterRequest {
	s.SchedulerType = &v
	return s
}

func (s *RecoverClusterRequest) SetImageOwnerAlias(v string) *RecoverClusterRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *RecoverClusterRequest) SetImageId(v string) *RecoverClusterRequest {
	s.ImageId = &v
	return s
}

func (s *RecoverClusterRequest) SetClientVersion(v string) *RecoverClusterRequest {
	s.ClientVersion = &v
	return s
}

type RecoverClusterResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverClusterResponseBody) SetTaskId(v string) *RecoverClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *RecoverClusterResponseBody) SetRequestId(v string) *RecoverClusterResponseBody {
	s.RequestId = &v
	return s
}

type RecoverClusterResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RecoverClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RecoverClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s RecoverClusterResponse) GoString() string {
	return s.String()
}

func (s *RecoverClusterResponse) SetHeaders(v map[string]*string) *RecoverClusterResponse {
	s.Headers = v
	return s
}

func (s *RecoverClusterResponse) SetBody(v *RecoverClusterResponseBody) *RecoverClusterResponse {
	s.Body = v
	return s
}

type RerunJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s RerunJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsRequest) GoString() string {
	return s.String()
}

func (s *RerunJobsRequest) SetClusterId(v string) *RerunJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *RerunJobsRequest) SetJobs(v string) *RerunJobsRequest {
	s.Jobs = &v
	return s
}

type RerunJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RerunJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponseBody) GoString() string {
	return s.String()
}

func (s *RerunJobsResponseBody) SetRequestId(v string) *RerunJobsResponseBody {
	s.RequestId = &v
	return s
}

type RerunJobsResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RerunJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RerunJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s RerunJobsResponse) GoString() string {
	return s.String()
}

func (s *RerunJobsResponse) SetHeaders(v map[string]*string) *RerunJobsResponse {
	s.Headers = v
	return s
}

func (s *RerunJobsResponse) SetBody(v *RerunJobsResponseBody) *RerunJobsResponse {
	s.Body = v
	return s
}

type ResetNodesRequest struct {
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Instance  []*ResetNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ResetNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequest) GoString() string {
	return s.String()
}

func (s *ResetNodesRequest) SetClusterId(v string) *ResetNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ResetNodesRequest) SetInstance(v []*ResetNodesRequestInstance) *ResetNodesRequest {
	s.Instance = v
	return s
}

type ResetNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ResetNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *ResetNodesRequestInstance) SetId(v string) *ResetNodesRequestInstance {
	s.Id = &v
	return s
}

type ResetNodesResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetNodesResponseBody) SetTaskId(v string) *ResetNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ResetNodesResponseBody) SetRequestId(v string) *ResetNodesResponseBody {
	s.RequestId = &v
	return s
}

type ResetNodesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetNodesResponse) GoString() string {
	return s.String()
}

func (s *ResetNodesResponse) SetHeaders(v map[string]*string) *ResetNodesResponse {
	s.Headers = v
	return s
}

func (s *ResetNodesResponse) SetBody(v *ResetNodesResponseBody) *ResetNodesResponse {
	s.Body = v
	return s
}

type RunCloudMetricProfilingRequest struct {
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	HostName  *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	ProcessId *int32  `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	Duration  *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Freq      *int32  `json:"Freq,omitempty" xml:"Freq,omitempty"`
}

func (s RunCloudMetricProfilingRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingRequest) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingRequest) SetRegionId(v string) *RunCloudMetricProfilingRequest {
	s.RegionId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetClusterId(v string) *RunCloudMetricProfilingRequest {
	s.ClusterId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetHostName(v string) *RunCloudMetricProfilingRequest {
	s.HostName = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetProcessId(v int32) *RunCloudMetricProfilingRequest {
	s.ProcessId = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetDuration(v int32) *RunCloudMetricProfilingRequest {
	s.Duration = &v
	return s
}

func (s *RunCloudMetricProfilingRequest) SetFreq(v int32) *RunCloudMetricProfilingRequest {
	s.Freq = &v
	return s
}

type RunCloudMetricProfilingResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCloudMetricProfilingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingResponseBody) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingResponseBody) SetRequestId(v string) *RunCloudMetricProfilingResponseBody {
	s.RequestId = &v
	return s
}

type RunCloudMetricProfilingResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RunCloudMetricProfilingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCloudMetricProfilingResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCloudMetricProfilingResponse) GoString() string {
	return s.String()
}

func (s *RunCloudMetricProfilingResponse) SetHeaders(v map[string]*string) *RunCloudMetricProfilingResponse {
	s.Headers = v
	return s
}

func (s *RunCloudMetricProfilingResponse) SetBody(v *RunCloudMetricProfilingResponseBody) *RunCloudMetricProfilingResponse {
	s.Body = v
	return s
}

type SetAutoScaleConfigRequest struct {
	ClusterId               *string                            `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	EnableAutoGrow          *bool                              `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	EnableAutoShrink        *bool                              `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	GrowIntervalInMinutes   *int32                             `json:"GrowIntervalInMinutes,omitempty" xml:"GrowIntervalInMinutes,omitempty"`
	ShrinkIntervalInMinutes *int32                             `json:"ShrinkIntervalInMinutes,omitempty" xml:"ShrinkIntervalInMinutes,omitempty"`
	ShrinkIdleTimes         *int32                             `json:"ShrinkIdleTimes,omitempty" xml:"ShrinkIdleTimes,omitempty"`
	GrowTimeoutInMinutes    *int32                             `json:"GrowTimeoutInMinutes,omitempty" xml:"GrowTimeoutInMinutes,omitempty"`
	ExtraNodesGrowRatio     *int32                             `json:"ExtraNodesGrowRatio,omitempty" xml:"ExtraNodesGrowRatio,omitempty"`
	GrowRatio               *int32                             `json:"GrowRatio,omitempty" xml:"GrowRatio,omitempty"`
	MaxNodesInCluster       *int32                             `json:"MaxNodesInCluster,omitempty" xml:"MaxNodesInCluster,omitempty"`
	ExcludeNodes            *string                            `json:"ExcludeNodes,omitempty" xml:"ExcludeNodes,omitempty"`
	SpotStrategy            *string                            `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	SpotPriceLimit          *float32                           `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	ImageId                 *string                            `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Queues                  []*SetAutoScaleConfigRequestQueues `json:"Queues,omitempty" xml:"Queues,omitempty" type:"Repeated"`
}

func (s SetAutoScaleConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequest) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequest) SetClusterId(v string) *SetAutoScaleConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequest {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIntervalInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIntervalInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetShrinkIdleTimes(v int32) *SetAutoScaleConfigRequest {
	s.ShrinkIdleTimes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowTimeoutInMinutes(v int32) *SetAutoScaleConfigRequest {
	s.GrowTimeoutInMinutes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExtraNodesGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.ExtraNodesGrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetGrowRatio(v int32) *SetAutoScaleConfigRequest {
	s.GrowRatio = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetMaxNodesInCluster(v int32) *SetAutoScaleConfigRequest {
	s.MaxNodesInCluster = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetExcludeNodes(v string) *SetAutoScaleConfigRequest {
	s.ExcludeNodes = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetSpotStrategy(v string) *SetAutoScaleConfigRequest {
	s.SpotStrategy = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetImageId(v string) *SetAutoScaleConfigRequest {
	s.ImageId = &v
	return s
}

func (s *SetAutoScaleConfigRequest) SetQueues(v []*SetAutoScaleConfigRequestQueues) *SetAutoScaleConfigRequest {
	s.Queues = v
	return s
}

type SetAutoScaleConfigRequestQueues struct {
	MinNodesInQueue  *int32                                          `json:"MinNodesInQueue,omitempty" xml:"MinNodesInQueue,omitempty"`
	MaxNodesInQueue  *int32                                          `json:"MaxNodesInQueue,omitempty" xml:"MaxNodesInQueue,omitempty"`
	EnableAutoShrink *bool                                           `json:"EnableAutoShrink,omitempty" xml:"EnableAutoShrink,omitempty"`
	QueueName        *string                                         `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	EnableAutoGrow   *bool                                           `json:"EnableAutoGrow,omitempty" xml:"EnableAutoGrow,omitempty"`
	QueueImageId     *string                                         `json:"QueueImageId,omitempty" xml:"QueueImageId,omitempty"`
	SpotPriceLimit   *float32                                        `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	InstanceTypes    []*SetAutoScaleConfigRequestQueuesInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Repeated"`
	InstanceType     *string                                         `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotStrategy     *string                                         `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s SetAutoScaleConfigRequestQueues) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueues) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueues) SetMinNodesInQueue(v int32) *SetAutoScaleConfigRequestQueues {
	s.MinNodesInQueue = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetMaxNodesInQueue(v int32) *SetAutoScaleConfigRequestQueues {
	s.MaxNodesInQueue = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetEnableAutoShrink(v bool) *SetAutoScaleConfigRequestQueues {
	s.EnableAutoShrink = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetQueueName(v string) *SetAutoScaleConfigRequestQueues {
	s.QueueName = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetEnableAutoGrow(v bool) *SetAutoScaleConfigRequestQueues {
	s.EnableAutoGrow = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetQueueImageId(v string) *SetAutoScaleConfigRequestQueues {
	s.QueueImageId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequestQueues {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetInstanceTypes(v []*SetAutoScaleConfigRequestQueuesInstanceTypes) *SetAutoScaleConfigRequestQueues {
	s.InstanceTypes = v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetInstanceType(v string) *SetAutoScaleConfigRequestQueues {
	s.InstanceType = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueues) SetSpotStrategy(v string) *SetAutoScaleConfigRequestQueues {
	s.SpotStrategy = &v
	return s
}

type SetAutoScaleConfigRequestQueuesInstanceTypes struct {
	HostNamePrefix *string  `json:"HostNamePrefix,omitempty" xml:"HostNamePrefix,omitempty"`
	VSwitchId      *string  `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId         *string  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	InstanceType   *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	SpotStrategy   *string  `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s SetAutoScaleConfigRequestQueuesInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigRequestQueuesInstanceTypes) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetHostNamePrefix(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.HostNamePrefix = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetVSwitchId(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.VSwitchId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetZoneId(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.ZoneId = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotPriceLimit(v float32) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotPriceLimit = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetInstanceType(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.InstanceType = &v
	return s
}

func (s *SetAutoScaleConfigRequestQueuesInstanceTypes) SetSpotStrategy(v string) *SetAutoScaleConfigRequestQueuesInstanceTypes {
	s.SpotStrategy = &v
	return s
}

type SetAutoScaleConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAutoScaleConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponseBody) SetRequestId(v string) *SetAutoScaleConfigResponseBody {
	s.RequestId = &v
	return s
}

type SetAutoScaleConfigResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAutoScaleConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAutoScaleConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAutoScaleConfigResponse) GoString() string {
	return s.String()
}

func (s *SetAutoScaleConfigResponse) SetHeaders(v map[string]*string) *SetAutoScaleConfigResponse {
	s.Headers = v
	return s
}

func (s *SetAutoScaleConfigResponse) SetBody(v *SetAutoScaleConfigResponseBody) *SetAutoScaleConfigResponse {
	s.Body = v
	return s
}

type SetGWSClusterPolicyRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Clipboard   *string `json:"Clipboard,omitempty" xml:"Clipboard,omitempty"`
	LocalDrive  *string `json:"LocalDrive,omitempty" xml:"LocalDrive,omitempty"`
	UsbRedirect *string `json:"UsbRedirect,omitempty" xml:"UsbRedirect,omitempty"`
	Watermark   *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
	UdpPort     *string `json:"UdpPort,omitempty" xml:"UdpPort,omitempty"`
	AsyncMode   *bool   `json:"AsyncMode,omitempty" xml:"AsyncMode,omitempty"`
}

func (s SetGWSClusterPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyRequest) SetClusterId(v string) *SetGWSClusterPolicyRequest {
	s.ClusterId = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetClipboard(v string) *SetGWSClusterPolicyRequest {
	s.Clipboard = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetLocalDrive(v string) *SetGWSClusterPolicyRequest {
	s.LocalDrive = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetUsbRedirect(v string) *SetGWSClusterPolicyRequest {
	s.UsbRedirect = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetWatermark(v string) *SetGWSClusterPolicyRequest {
	s.Watermark = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetUdpPort(v string) *SetGWSClusterPolicyRequest {
	s.UdpPort = &v
	return s
}

func (s *SetGWSClusterPolicyRequest) SetAsyncMode(v bool) *SetGWSClusterPolicyRequest {
	s.AsyncMode = &v
	return s
}

type SetGWSClusterPolicyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSClusterPolicyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyResponseBody) SetRequestId(v string) *SetGWSClusterPolicyResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSClusterPolicyResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGWSClusterPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSClusterPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSClusterPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetGWSClusterPolicyResponse) SetHeaders(v map[string]*string) *SetGWSClusterPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetGWSClusterPolicyResponse) SetBody(v *SetGWSClusterPolicyResponseBody) *SetGWSClusterPolicyResponse {
	s.Body = v
	return s
}

type SetGWSInstanceNameRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetGWSInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameRequest) SetInstanceId(v string) *SetGWSInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *SetGWSInstanceNameRequest) SetName(v string) *SetGWSInstanceNameRequest {
	s.Name = &v
	return s
}

type SetGWSInstanceNameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameResponseBody) SetRequestId(v string) *SetGWSInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSInstanceNameResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGWSInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceNameResponse) SetHeaders(v map[string]*string) *SetGWSInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *SetGWSInstanceNameResponse) SetBody(v *SetGWSInstanceNameResponseBody) *SetGWSInstanceNameResponse {
	s.Body = v
	return s
}

type SetGWSInstanceUserRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	UserUid    *string `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s SetGWSInstanceUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserRequest) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserRequest) SetInstanceId(v string) *SetGWSInstanceUserRequest {
	s.InstanceId = &v
	return s
}

func (s *SetGWSInstanceUserRequest) SetUserUid(v string) *SetGWSInstanceUserRequest {
	s.UserUid = &v
	return s
}

func (s *SetGWSInstanceUserRequest) SetUserName(v string) *SetGWSInstanceUserRequest {
	s.UserName = &v
	return s
}

type SetGWSInstanceUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetGWSInstanceUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserResponseBody) SetRequestId(v string) *SetGWSInstanceUserResponseBody {
	s.RequestId = &v
	return s
}

type SetGWSInstanceUserResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetGWSInstanceUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetGWSInstanceUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetGWSInstanceUserResponse) GoString() string {
	return s.String()
}

func (s *SetGWSInstanceUserResponse) SetHeaders(v map[string]*string) *SetGWSInstanceUserResponse {
	s.Headers = v
	return s
}

func (s *SetGWSInstanceUserResponse) SetBody(v *SetGWSInstanceUserResponseBody) *SetGWSInstanceUserResponse {
	s.Body = v
	return s
}

type SetJobUserRequest struct {
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	RunasUser         *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
}

func (s SetJobUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SetJobUserRequest) GoString() string {
	return s.String()
}

func (s *SetJobUserRequest) SetClusterId(v string) *SetJobUserRequest {
	s.ClusterId = &v
	return s
}

func (s *SetJobUserRequest) SetRunasUser(v string) *SetJobUserRequest {
	s.RunasUser = &v
	return s
}

func (s *SetJobUserRequest) SetRunasUserPassword(v string) *SetJobUserRequest {
	s.RunasUserPassword = &v
	return s
}

type SetJobUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetJobUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetJobUserResponseBody) GoString() string {
	return s.String()
}

func (s *SetJobUserResponseBody) SetRequestId(v string) *SetJobUserResponseBody {
	s.RequestId = &v
	return s
}

type SetJobUserResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetJobUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetJobUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SetJobUserResponse) GoString() string {
	return s.String()
}

func (s *SetJobUserResponse) SetHeaders(v map[string]*string) *SetJobUserResponse {
	s.Headers = v
	return s
}

func (s *SetJobUserResponse) SetBody(v *SetJobUserResponseBody) *SetJobUserResponse {
	s.Body = v
	return s
}

type SetQueueRequest struct {
	ClusterId *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName *string                `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	Node      []*SetQueueRequestNode `json:"Node,omitempty" xml:"Node,omitempty" type:"Repeated"`
}

func (s SetQueueRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQueueRequest) GoString() string {
	return s.String()
}

func (s *SetQueueRequest) SetClusterId(v string) *SetQueueRequest {
	s.ClusterId = &v
	return s
}

func (s *SetQueueRequest) SetQueueName(v string) *SetQueueRequest {
	s.QueueName = &v
	return s
}

func (s *SetQueueRequest) SetNode(v []*SetQueueRequestNode) *SetQueueRequest {
	s.Node = v
	return s
}

type SetQueueRequestNode struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SetQueueRequestNode) String() string {
	return tea.Prettify(s)
}

func (s SetQueueRequestNode) GoString() string {
	return s.String()
}

func (s *SetQueueRequestNode) SetName(v string) *SetQueueRequestNode {
	s.Name = &v
	return s
}

type SetQueueResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetQueueResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQueueResponseBody) GoString() string {
	return s.String()
}

func (s *SetQueueResponseBody) SetRequestId(v string) *SetQueueResponseBody {
	s.RequestId = &v
	return s
}

type SetQueueResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetQueueResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetQueueResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQueueResponse) GoString() string {
	return s.String()
}

func (s *SetQueueResponse) SetHeaders(v map[string]*string) *SetQueueResponse {
	s.Headers = v
	return s
}

func (s *SetQueueResponse) SetBody(v *SetQueueResponseBody) *SetQueueResponse {
	s.Body = v
	return s
}

type StartClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s StartClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StartClusterRequest) GoString() string {
	return s.String()
}

func (s *StartClusterRequest) SetClusterId(v string) *StartClusterRequest {
	s.ClusterId = &v
	return s
}

type StartClusterResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartClusterResponseBody) SetTaskId(v string) *StartClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartClusterResponseBody) SetRequestId(v string) *StartClusterResponseBody {
	s.RequestId = &v
	return s
}

type StartClusterResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StartClusterResponse) GoString() string {
	return s.String()
}

func (s *StartClusterResponse) SetHeaders(v map[string]*string) *StartClusterResponse {
	s.Headers = v
	return s
}

func (s *StartClusterResponse) SetBody(v *StartClusterResponseBody) *StartClusterResponse {
	s.Body = v
	return s
}

type StartGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceRequest) SetInstanceId(v string) *StartGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type StartGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceResponseBody) SetRequestId(v string) *StartGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StartGWSInstanceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartGWSInstanceResponse) SetHeaders(v map[string]*string) *StartGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartGWSInstanceResponse) SetBody(v *StartGWSInstanceResponseBody) *StartGWSInstanceResponse {
	s.Body = v
	return s
}

type StartNodesRequest struct {
	ClusterId *string                      `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Role      *string                      `json:"Role,omitempty" xml:"Role,omitempty"`
	Instance  []*StartNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s StartNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StartNodesRequest) GoString() string {
	return s.String()
}

func (s *StartNodesRequest) SetClusterId(v string) *StartNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *StartNodesRequest) SetRole(v string) *StartNodesRequest {
	s.Role = &v
	return s
}

func (s *StartNodesRequest) SetInstance(v []*StartNodesRequestInstance) *StartNodesRequest {
	s.Instance = v
	return s
}

type StartNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StartNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s StartNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *StartNodesRequestInstance) SetId(v string) *StartNodesRequestInstance {
	s.Id = &v
	return s
}

type StartNodesResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StartNodesResponseBody) SetTaskId(v string) *StartNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartNodesResponseBody) SetRequestId(v string) *StartNodesResponseBody {
	s.RequestId = &v
	return s
}

type StartNodesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StartNodesResponse) GoString() string {
	return s.String()
}

func (s *StartNodesResponse) SetHeaders(v map[string]*string) *StartNodesResponse {
	s.Headers = v
	return s
}

func (s *StartNodesResponse) SetBody(v *StartNodesResponseBody) *StartNodesResponse {
	s.Body = v
	return s
}

type StartVisualServiceRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CidrIp    *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s StartVisualServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceRequest) GoString() string {
	return s.String()
}

func (s *StartVisualServiceRequest) SetClusterId(v string) *StartVisualServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *StartVisualServiceRequest) SetCidrIp(v string) *StartVisualServiceRequest {
	s.CidrIp = &v
	return s
}

func (s *StartVisualServiceRequest) SetPort(v int32) *StartVisualServiceRequest {
	s.Port = &v
	return s
}

type StartVisualServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartVisualServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartVisualServiceResponseBody) SetMessage(v string) *StartVisualServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartVisualServiceResponseBody) SetRequestId(v string) *StartVisualServiceResponseBody {
	s.RequestId = &v
	return s
}

type StartVisualServiceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartVisualServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartVisualServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartVisualServiceResponse) GoString() string {
	return s.String()
}

func (s *StartVisualServiceResponse) SetHeaders(v map[string]*string) *StartVisualServiceResponse {
	s.Headers = v
	return s
}

func (s *StartVisualServiceResponse) SetBody(v *StartVisualServiceResponseBody) *StartVisualServiceResponse {
	s.Body = v
	return s
}

type StopClusterRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s StopClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s StopClusterRequest) GoString() string {
	return s.String()
}

func (s *StopClusterRequest) SetClusterId(v string) *StopClusterRequest {
	s.ClusterId = &v
	return s
}

type StopClusterResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopClusterResponseBody) SetTaskId(v string) *StopClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopClusterResponseBody) SetRequestId(v string) *StopClusterResponseBody {
	s.RequestId = &v
	return s
}

type StopClusterResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopClusterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s StopClusterResponse) GoString() string {
	return s.String()
}

func (s *StopClusterResponse) SetHeaders(v map[string]*string) *StopClusterResponse {
	s.Headers = v
	return s
}

func (s *StopClusterResponse) SetBody(v *StopClusterResponseBody) *StopClusterResponse {
	s.Body = v
	return s
}

type StopGWSInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopGWSInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceRequest) SetInstanceId(v string) *StopGWSInstanceRequest {
	s.InstanceId = &v
	return s
}

type StopGWSInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopGWSInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceResponseBody) SetRequestId(v string) *StopGWSInstanceResponseBody {
	s.RequestId = &v
	return s
}

type StopGWSInstanceResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopGWSInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopGWSInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopGWSInstanceResponse) GoString() string {
	return s.String()
}

func (s *StopGWSInstanceResponse) SetHeaders(v map[string]*string) *StopGWSInstanceResponse {
	s.Headers = v
	return s
}

func (s *StopGWSInstanceResponse) SetBody(v *StopGWSInstanceResponseBody) *StopGWSInstanceResponse {
	s.Body = v
	return s
}

type StopJobsRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Jobs      *string `json:"Jobs,omitempty" xml:"Jobs,omitempty"`
}

func (s StopJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobsRequest) GoString() string {
	return s.String()
}

func (s *StopJobsRequest) SetClusterId(v string) *StopJobsRequest {
	s.ClusterId = &v
	return s
}

func (s *StopJobsRequest) SetJobs(v string) *StopJobsRequest {
	s.Jobs = &v
	return s
}

type StopJobsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobsResponseBody) SetRequestId(v string) *StopJobsResponseBody {
	s.RequestId = &v
	return s
}

type StopJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponse) GoString() string {
	return s.String()
}

func (s *StopJobsResponse) SetHeaders(v map[string]*string) *StopJobsResponse {
	s.Headers = v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
	s.Body = v
	return s
}

type StopNodesRequest struct {
	ClusterId *string                     `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Role      *string                     `json:"Role,omitempty" xml:"Role,omitempty"`
	Instance  []*StopNodesRequestInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s StopNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) SetClusterId(v string) *StopNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *StopNodesRequest) SetRole(v string) *StopNodesRequest {
	s.Role = &v
	return s
}

func (s *StopNodesRequest) SetInstance(v []*StopNodesRequestInstance) *StopNodesRequest {
	s.Instance = v
	return s
}

type StopNodesRequestInstance struct {
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s StopNodesRequestInstance) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequestInstance) GoString() string {
	return s.String()
}

func (s *StopNodesRequestInstance) SetId(v string) *StopNodesRequestInstance {
	s.Id = &v
	return s
}

type StopNodesResponseBody struct {
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

type StopNodesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) SetHeaders(v map[string]*string) *StopNodesResponse {
	s.Headers = v
	return s
}

func (s *StopNodesResponse) SetBody(v *StopNodesResponseBody) *StopNodesResponse {
	s.Body = v
	return s
}

type StopVisualServiceRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CidrIp    *string `json:"CidrIp,omitempty" xml:"CidrIp,omitempty"`
	Port      *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s StopVisualServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceRequest) GoString() string {
	return s.String()
}

func (s *StopVisualServiceRequest) SetClusterId(v string) *StopVisualServiceRequest {
	s.ClusterId = &v
	return s
}

func (s *StopVisualServiceRequest) SetCidrIp(v string) *StopVisualServiceRequest {
	s.CidrIp = &v
	return s
}

func (s *StopVisualServiceRequest) SetPort(v int32) *StopVisualServiceRequest {
	s.Port = &v
	return s
}

type StopVisualServiceResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopVisualServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopVisualServiceResponseBody) SetMessage(v string) *StopVisualServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StopVisualServiceResponseBody) SetRequestId(v string) *StopVisualServiceResponseBody {
	s.RequestId = &v
	return s
}

type StopVisualServiceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopVisualServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopVisualServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopVisualServiceResponse) GoString() string {
	return s.String()
}

func (s *StopVisualServiceResponse) SetHeaders(v map[string]*string) *StopVisualServiceResponse {
	s.Headers = v
	return s
}

func (s *StopVisualServiceResponse) SetBody(v *StopVisualServiceResponseBody) *StopVisualServiceResponse {
	s.Body = v
	return s
}

type SubmitJobRequest struct {
	ClusterId          *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CommandLine        *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	RunasUser          *string `json:"RunasUser,omitempty" xml:"RunasUser,omitempty"`
	RunasUserPassword  *string `json:"RunasUserPassword,omitempty" xml:"RunasUserPassword,omitempty"`
	Name               *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Priority           *int32  `json:"Priority,omitempty" xml:"Priority,omitempty"`
	PackagePath        *string `json:"PackagePath,omitempty" xml:"PackagePath,omitempty"`
	StdoutRedirectPath *string `json:"StdoutRedirectPath,omitempty" xml:"StdoutRedirectPath,omitempty"`
	StderrRedirectPath *string `json:"StderrRedirectPath,omitempty" xml:"StderrRedirectPath,omitempty"`
	ReRunable          *bool   `json:"ReRunable,omitempty" xml:"ReRunable,omitempty"`
	ArrayRequest       *string `json:"ArrayRequest,omitempty" xml:"ArrayRequest,omitempty"`
	Variables          *string `json:"Variables,omitempty" xml:"Variables,omitempty"`
	InputFileUrl       *string `json:"InputFileUrl,omitempty" xml:"InputFileUrl,omitempty"`
	UnzipCmd           *string `json:"UnzipCmd,omitempty" xml:"UnzipCmd,omitempty"`
	PostCmdLine        *string `json:"PostCmdLine,omitempty" xml:"PostCmdLine,omitempty"`
	ContainerId        *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	JobQueue           *string `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	Node               *int32  `json:"Node,omitempty" xml:"Node,omitempty"`
	Task               *int32  `json:"Task,omitempty" xml:"Task,omitempty"`
	Thread             *int32  `json:"Thread,omitempty" xml:"Thread,omitempty"`
	Mem                *string `json:"Mem,omitempty" xml:"Mem,omitempty"`
	Gpu                *int32  `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	ClockTime          *string `json:"ClockTime,omitempty" xml:"ClockTime,omitempty"`
}

func (s SubmitJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitJobRequest) SetClusterId(v string) *SubmitJobRequest {
	s.ClusterId = &v
	return s
}

func (s *SubmitJobRequest) SetCommandLine(v string) *SubmitJobRequest {
	s.CommandLine = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUser(v string) *SubmitJobRequest {
	s.RunasUser = &v
	return s
}

func (s *SubmitJobRequest) SetRunasUserPassword(v string) *SubmitJobRequest {
	s.RunasUserPassword = &v
	return s
}

func (s *SubmitJobRequest) SetName(v string) *SubmitJobRequest {
	s.Name = &v
	return s
}

func (s *SubmitJobRequest) SetPriority(v int32) *SubmitJobRequest {
	s.Priority = &v
	return s
}

func (s *SubmitJobRequest) SetPackagePath(v string) *SubmitJobRequest {
	s.PackagePath = &v
	return s
}

func (s *SubmitJobRequest) SetStdoutRedirectPath(v string) *SubmitJobRequest {
	s.StdoutRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetStderrRedirectPath(v string) *SubmitJobRequest {
	s.StderrRedirectPath = &v
	return s
}

func (s *SubmitJobRequest) SetReRunable(v bool) *SubmitJobRequest {
	s.ReRunable = &v
	return s
}

func (s *SubmitJobRequest) SetArrayRequest(v string) *SubmitJobRequest {
	s.ArrayRequest = &v
	return s
}

func (s *SubmitJobRequest) SetVariables(v string) *SubmitJobRequest {
	s.Variables = &v
	return s
}

func (s *SubmitJobRequest) SetInputFileUrl(v string) *SubmitJobRequest {
	s.InputFileUrl = &v
	return s
}

func (s *SubmitJobRequest) SetUnzipCmd(v string) *SubmitJobRequest {
	s.UnzipCmd = &v
	return s
}

func (s *SubmitJobRequest) SetPostCmdLine(v string) *SubmitJobRequest {
	s.PostCmdLine = &v
	return s
}

func (s *SubmitJobRequest) SetContainerId(v string) *SubmitJobRequest {
	s.ContainerId = &v
	return s
}

func (s *SubmitJobRequest) SetJobQueue(v string) *SubmitJobRequest {
	s.JobQueue = &v
	return s
}

func (s *SubmitJobRequest) SetNode(v int32) *SubmitJobRequest {
	s.Node = &v
	return s
}

func (s *SubmitJobRequest) SetTask(v int32) *SubmitJobRequest {
	s.Task = &v
	return s
}

func (s *SubmitJobRequest) SetThread(v int32) *SubmitJobRequest {
	s.Thread = &v
	return s
}

func (s *SubmitJobRequest) SetMem(v string) *SubmitJobRequest {
	s.Mem = &v
	return s
}

func (s *SubmitJobRequest) SetGpu(v int32) *SubmitJobRequest {
	s.Gpu = &v
	return s
}

func (s *SubmitJobRequest) SetClockTime(v string) *SubmitJobRequest {
	s.ClockTime = &v
	return s
}

type SubmitJobResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s SubmitJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitJobResponseBody) SetRequestId(v string) *SubmitJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitJobResponseBody) SetJobId(v string) *SubmitJobResponseBody {
	s.JobId = &v
	return s
}

type SubmitJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SubmitJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitJobResponse) SetHeaders(v map[string]*string) *SubmitJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitJobResponse) SetBody(v *SubmitJobResponseBody) *SubmitJobResponse {
	s.Body = v
	return s
}

type UnbindAccountToClusterUserRequest struct {
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	AccountUid *string `json:"AccountUid,omitempty" xml:"AccountUid,omitempty"`
}

func (s UnbindAccountToClusterUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountToClusterUserRequest) GoString() string {
	return s.String()
}

func (s *UnbindAccountToClusterUserRequest) SetClusterId(v string) *UnbindAccountToClusterUserRequest {
	s.ClusterId = &v
	return s
}

func (s *UnbindAccountToClusterUserRequest) SetUserName(v string) *UnbindAccountToClusterUserRequest {
	s.UserName = &v
	return s
}

func (s *UnbindAccountToClusterUserRequest) SetAccountUid(v string) *UnbindAccountToClusterUserRequest {
	s.AccountUid = &v
	return s
}

type UnbindAccountToClusterUserResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAccountToClusterUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountToClusterUserResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAccountToClusterUserResponseBody) SetRequestId(v string) *UnbindAccountToClusterUserResponseBody {
	s.RequestId = &v
	return s
}

type UnbindAccountToClusterUserResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnbindAccountToClusterUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindAccountToClusterUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindAccountToClusterUserResponse) GoString() string {
	return s.String()
}

func (s *UnbindAccountToClusterUserResponse) SetHeaders(v map[string]*string) *UnbindAccountToClusterUserResponse {
	s.Headers = v
	return s
}

func (s *UnbindAccountToClusterUserResponse) SetBody(v *UnbindAccountToClusterUserResponseBody) *UnbindAccountToClusterUserResponse {
	s.Body = v
	return s
}

type UninstallSoftwareRequest struct {
	ClusterId   *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Application *string `json:"Application,omitempty" xml:"Application,omitempty"`
}

func (s UninstallSoftwareRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareRequest) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareRequest) SetClusterId(v string) *UninstallSoftwareRequest {
	s.ClusterId = &v
	return s
}

func (s *UninstallSoftwareRequest) SetApplication(v string) *UninstallSoftwareRequest {
	s.Application = &v
	return s
}

type UninstallSoftwareResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UninstallSoftwareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareResponseBody) SetRequestId(v string) *UninstallSoftwareResponseBody {
	s.RequestId = &v
	return s
}

type UninstallSoftwareResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UninstallSoftwareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallSoftwareResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallSoftwareResponse) GoString() string {
	return s.String()
}

func (s *UninstallSoftwareResponse) SetHeaders(v map[string]*string) *UninstallSoftwareResponse {
	s.Headers = v
	return s
}

func (s *UninstallSoftwareResponse) SetBody(v *UninstallSoftwareResponseBody) *UninstallSoftwareResponse {
	s.Body = v
	return s
}

type UpdateClusterVolumesRequest struct {
	ClusterId         *string                                         `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	AdditionalVolumes []*UpdateClusterVolumesRequestAdditionalVolumes `json:"AdditionalVolumes,omitempty" xml:"AdditionalVolumes,omitempty" type:"Repeated"`
}

func (s UpdateClusterVolumesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequest) SetClusterId(v string) *UpdateClusterVolumesRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterVolumesRequest) SetAdditionalVolumes(v []*UpdateClusterVolumesRequestAdditionalVolumes) *UpdateClusterVolumesRequest {
	s.AdditionalVolumes = v
	return s
}

type UpdateClusterVolumesRequestAdditionalVolumes struct {
	JobQueue         *string                                              `json:"JobQueue,omitempty" xml:"JobQueue,omitempty"`
	VolumeId         *string                                              `json:"VolumeId,omitempty" xml:"VolumeId,omitempty"`
	Roles            []*UpdateClusterVolumesRequestAdditionalVolumesRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	RemoteDirectory  *string                                              `json:"RemoteDirectory,omitempty" xml:"RemoteDirectory,omitempty"`
	VolumeMountpoint *string                                              `json:"VolumeMountpoint,omitempty" xml:"VolumeMountpoint,omitempty"`
	LocalDirectory   *string                                              `json:"LocalDirectory,omitempty" xml:"LocalDirectory,omitempty"`
	VolumeType       *string                                              `json:"VolumeType,omitempty" xml:"VolumeType,omitempty"`
	VolumeProtocol   *string                                              `json:"VolumeProtocol,omitempty" xml:"VolumeProtocol,omitempty"`
	Location         *string                                              `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s UpdateClusterVolumesRequestAdditionalVolumes) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequestAdditionalVolumes) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetJobQueue(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.JobQueue = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeId(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeId = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetRoles(v []*UpdateClusterVolumesRequestAdditionalVolumesRoles) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.Roles = v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetRemoteDirectory(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.RemoteDirectory = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeMountpoint(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeMountpoint = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetLocalDirectory(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.LocalDirectory = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeType(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeType = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetVolumeProtocol(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.VolumeProtocol = &v
	return s
}

func (s *UpdateClusterVolumesRequestAdditionalVolumes) SetLocation(v string) *UpdateClusterVolumesRequestAdditionalVolumes {
	s.Location = &v
	return s
}

type UpdateClusterVolumesRequestAdditionalVolumesRoles struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateClusterVolumesRequestAdditionalVolumesRoles) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesRequestAdditionalVolumesRoles) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesRequestAdditionalVolumesRoles) SetName(v string) *UpdateClusterVolumesRequestAdditionalVolumesRoles {
	s.Name = &v
	return s
}

type UpdateClusterVolumesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterVolumesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesResponseBody) SetRequestId(v string) *UpdateClusterVolumesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateClusterVolumesResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateClusterVolumesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateClusterVolumesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateClusterVolumesResponse) GoString() string {
	return s.String()
}

func (s *UpdateClusterVolumesResponse) SetHeaders(v map[string]*string) *UpdateClusterVolumesResponse {
	s.Headers = v
	return s
}

func (s *UpdateClusterVolumesResponse) SetBody(v *UpdateClusterVolumesResponseBody) *UpdateClusterVolumesResponse {
	s.Body = v
	return s
}

type UpdateQueueConfigRequest struct {
	ClusterId           *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	QueueName           *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	ResourceGroupId     *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ComputeInstanceType *string `json:"ComputeInstanceType,omitempty" xml:"ComputeInstanceType,omitempty"`
}

func (s UpdateQueueConfigRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigRequest) SetClusterId(v string) *UpdateQueueConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetQueueName(v string) *UpdateQueueConfigRequest {
	s.QueueName = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetResourceGroupId(v string) *UpdateQueueConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateQueueConfigRequest) SetComputeInstanceType(v string) *UpdateQueueConfigRequest {
	s.ComputeInstanceType = &v
	return s
}

type UpdateQueueConfigResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQueueConfigResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigResponseBody) SetRequestId(v string) *UpdateQueueConfigResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQueueConfigResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateQueueConfigResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQueueConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQueueConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateQueueConfigResponse) SetHeaders(v map[string]*string) *UpdateQueueConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateQueueConfigResponse) SetBody(v *UpdateQueueConfigResponseBody) *UpdateQueueConfigResponse {
	s.Body = v
	return s
}

type UpgradeClientRequest struct {
	ClusterId     *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
}

func (s UpgradeClientRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientRequest) GoString() string {
	return s.String()
}

func (s *UpgradeClientRequest) SetClusterId(v string) *UpgradeClientRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeClientRequest) SetClientVersion(v string) *UpgradeClientRequest {
	s.ClientVersion = &v
	return s
}

type UpgradeClientResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeClientResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponseBody) SetRequestId(v string) *UpgradeClientResponseBody {
	s.RequestId = &v
	return s
}

type UpgradeClientResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeClientResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeClientResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeClientResponse) GoString() string {
	return s.String()
}

func (s *UpgradeClientResponse) SetHeaders(v map[string]*string) *UpgradeClientResponse {
	s.Headers = v
	return s
}

func (s *UpgradeClientResponse) SetBody(v *UpgradeClientResponseBody) *UpgradeClientResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("ehpc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddContainerAppWithOptions(request *AddContainerAppRequest, runtime *util.RuntimeOptions) (_result *AddContainerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddContainerAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddContainerApp"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddContainerApp(request *AddContainerAppRequest) (_result *AddContainerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddContainerAppResponse{}
	_body, _err := client.AddContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddLocalNodesWithOptions(request *AddLocalNodesRequest, runtime *util.RuntimeOptions) (_result *AddLocalNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddLocalNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddLocalNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddLocalNodes(request *AddLocalNodesRequest) (_result *AddLocalNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddLocalNodesResponse{}
	_body, _err := client.AddLocalNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddNodesWithOptions(request *AddNodesRequest, runtime *util.RuntimeOptions) (_result *AddNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddNodes(request *AddNodesRequest) (_result *AddNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddNodesResponse{}
	_body, _err := client.AddNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddQueueWithOptions(request *AddQueueRequest, runtime *util.RuntimeOptions) (_result *AddQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddQueue"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddQueue(request *AddQueueRequest) (_result *AddQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddQueueResponse{}
	_body, _err := client.AddQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSecurityGroupWithOptions(request *AddSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *AddSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddSecurityGroup"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSecurityGroup(request *AddSecurityGroupRequest) (_result *AddSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddSecurityGroupResponse{}
	_body, _err := client.AddSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddUsersWithOptions(request *AddUsersRequest, runtime *util.RuntimeOptions) (_result *AddUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &AddUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddUsers"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddUsers(request *AddUsersRequest) (_result *AddUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddUsersResponse{}
	_body, _err := client.AddUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ApplyNodesWithOptions(request *ApplyNodesRequest, runtime *util.RuntimeOptions) (_result *ApplyNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ApplyNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ApplyNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ApplyNodes(request *ApplyNodesRequest) (_result *ApplyNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApplyNodesResponse{}
	_body, _err := client.ApplyNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindAccountToClusterUserWithOptions(request *BindAccountToClusterUserRequest, runtime *util.RuntimeOptions) (_result *BindAccountToClusterUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &BindAccountToClusterUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("BindAccountToClusterUser"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindAccountToClusterUser(request *BindAccountToClusterUserRequest) (_result *BindAccountToClusterUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &BindAccountToClusterUserResponse{}
	_body, _err := client.BindAccountToClusterUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSClusterWithOptions(request *CreateGWSClusterRequest, runtime *util.RuntimeOptions) (_result *CreateGWSClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateGWSClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGWSCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSCluster(request *CreateGWSClusterRequest) (_result *CreateGWSClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSClusterResponse{}
	_body, _err := client.CreateGWSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSImageWithOptions(request *CreateGWSImageRequest, runtime *util.RuntimeOptions) (_result *CreateGWSImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateGWSImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGWSImage"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSImage(request *CreateGWSImageRequest) (_result *CreateGWSImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSImageResponse{}
	_body, _err := client.CreateGWSImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateGWSInstanceWithOptions(request *CreateGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateGWSInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateGWSInstance"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateGWSInstance(request *CreateGWSInstanceRequest) (_result *CreateGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateGWSInstanceResponse{}
	_body, _err := client.CreateGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateHybridClusterWithOptions(request *CreateHybridClusterRequest, runtime *util.RuntimeOptions) (_result *CreateHybridClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateHybridClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateHybridCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateHybridCluster(request *CreateHybridClusterRequest) (_result *CreateHybridClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateHybridClusterResponse{}
	_body, _err := client.CreateHybridClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobFileWithOptions(request *CreateJobFileRequest, runtime *util.RuntimeOptions) (_result *CreateJobFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateJobFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateJobFile"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobFile(request *CreateJobFileRequest) (_result *CreateJobFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobFileResponse{}
	_body, _err := client.CreateJobFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobTemplateWithOptions(request *CreateJobTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateJobTemplate"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobTemplate(request *CreateJobTemplateRequest) (_result *CreateJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobTemplateResponse{}
	_body, _err := client.CreateJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteContainerAppsWithOptions(request *DeleteContainerAppsRequest, runtime *util.RuntimeOptions) (_result *DeleteContainerAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteContainerAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteContainerApps"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteContainerApps(request *DeleteContainerAppsRequest) (_result *DeleteContainerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteContainerAppsResponse{}
	_body, _err := client.DeleteContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGWSClusterWithOptions(request *DeleteGWSClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteGWSClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteGWSClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGWSCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGWSCluster(request *DeleteGWSClusterRequest) (_result *DeleteGWSClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGWSClusterResponse{}
	_body, _err := client.DeleteGWSClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGWSInstanceWithOptions(request *DeleteGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteGWSInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteGWSInstance"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGWSInstance(request *DeleteGWSInstanceRequest) (_result *DeleteGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteGWSInstanceResponse{}
	_body, _err := client.DeleteGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(request *DeleteImageRequest, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteImage"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobsWithOptions(request *DeleteJobsRequest, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteJobs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobTemplatesWithOptions(request *DeleteJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *DeleteJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteJobTemplates"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobTemplates(request *DeleteJobTemplatesRequest) (_result *DeleteJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteJobTemplatesResponse{}
	_body, _err := client.DeleteJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteNodesWithOptions(request *DeleteNodesRequest, runtime *util.RuntimeOptions) (_result *DeleteNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteNodes(request *DeleteNodesRequest) (_result *DeleteNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodesResponse{}
	_body, _err := client.DeleteNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQueueWithOptions(request *DeleteQueueRequest, runtime *util.RuntimeOptions) (_result *DeleteQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteQueue"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQueue(request *DeleteQueueRequest) (_result *DeleteQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteQueueResponse{}
	_body, _err := client.DeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityGroupWithOptions(request *DeleteSecurityGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSecurityGroup"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityGroup(request *DeleteSecurityGroupRequest) (_result *DeleteSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityGroupResponse{}
	_body, _err := client.DeleteSecurityGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteUsersWithOptions(request *DeleteUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteUsers"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteUsers(request *DeleteUsersRequest) (_result *DeleteUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteUsersResponse{}
	_body, _err := client.DeleteUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoScaleConfigWithOptions(request *DescribeAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeAutoScaleConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoScaleConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoScaleConfig(request *DescribeAutoScaleConfigRequest) (_result *DescribeAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoScaleConfigResponse{}
	_body, _err := client.DescribeAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeContainerAppWithOptions(request *DescribeContainerAppRequest, runtime *util.RuntimeOptions) (_result *DescribeContainerAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeContainerAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeContainerApp"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeContainerApp(request *DescribeContainerAppRequest) (_result *DescribeContainerAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeContainerAppResponse{}
	_body, _err := client.DescribeContainerAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSClusterPolicyWithOptions(request *DescribeGWSClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeGWSClusterPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGWSClusterPolicy"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSClusterPolicy(request *DescribeGWSClusterPolicyRequest) (_result *DescribeGWSClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSClusterPolicyResponse{}
	_body, _err := client.DescribeGWSClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSClustersWithOptions(request *DescribeGWSClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeGWSClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGWSClusters"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSClusters(request *DescribeGWSClustersRequest) (_result *DescribeGWSClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSClustersResponse{}
	_body, _err := client.DescribeGWSClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSImagesWithOptions(request *DescribeGWSImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeGWSImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGWSImages"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSImages(request *DescribeGWSImagesRequest) (_result *DescribeGWSImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSImagesResponse{}
	_body, _err := client.DescribeGWSImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeGWSInstancesWithOptions(request *DescribeGWSInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeGWSInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeGWSInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeGWSInstances"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeGWSInstances(request *DescribeGWSInstancesRequest) (_result *DescribeGWSInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeGWSInstancesResponse{}
	_body, _err := client.DescribeGWSInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageWithOptions(request *DescribeImageRequest, runtime *util.RuntimeOptions) (_result *DescribeImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImage"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImage(request *DescribeImageRequest) (_result *DescribeImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageResponse{}
	_body, _err := client.DescribeImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImageGatewayConfigWithOptions(request *DescribeImageGatewayConfigRequest, runtime *util.RuntimeOptions) (_result *DescribeImageGatewayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeImageGatewayConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImageGatewayConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImageGatewayConfig(request *DescribeImageGatewayConfigRequest) (_result *DescribeImageGatewayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImageGatewayConfigResponse{}
	_body, _err := client.DescribeImageGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagePriceWithOptions(request *DescribeImagePriceRequest, runtime *util.RuntimeOptions) (_result *DescribeImagePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeImagePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeImagePrice"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImagePrice(request *DescribeImagePriceRequest) (_result *DescribeImagePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagePriceResponse{}
	_body, _err := client.DescribeImagePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeJobWithOptions(request *DescribeJobRequest, runtime *util.RuntimeOptions) (_result *DescribeJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJob"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJob(request *DescribeJobRequest) (_result *DescribeJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobResponse{}
	_body, _err := client.DescribeJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNFSClientStatusWithOptions(request *DescribeNFSClientStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeNFSClientStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribeNFSClientStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeNFSClientStatus"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNFSClientStatus(request *DescribeNFSClientStatusRequest) (_result *DescribeNFSClientStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNFSClientStatusResponse{}
	_body, _err := client.DescribeNFSClientStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribePriceWithOptions(request *DescribePriceRequest, runtime *util.RuntimeOptions) (_result *DescribePriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &DescribePriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePrice"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePrice(request *DescribePriceRequest) (_result *DescribePriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePriceResponse{}
	_body, _err := client.DescribePriceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditJobTemplateWithOptions(request *EditJobTemplateRequest, runtime *util.RuntimeOptions) (_result *EditJobTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditJobTemplate"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditJobTemplate(request *EditJobTemplateRequest) (_result *EditJobTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditJobTemplateResponse{}
	_body, _err := client.EditJobTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAccountingReportWithOptions(request *GetAccountingReportRequest, runtime *util.RuntimeOptions) (_result *GetAccountingReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAccountingReportResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAccountingReport"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAccountingReport(request *GetAccountingReportRequest) (_result *GetAccountingReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAccountingReportResponse{}
	_body, _err := client.GetAccountingReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAutoScaleConfigWithOptions(request *GetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *GetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAutoScaleConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAutoScaleConfig(request *GetAutoScaleConfigRequest) (_result *GetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAutoScaleConfigResponse{}
	_body, _err := client.GetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudMetricLogsWithOptions(request *GetCloudMetricLogsRequest, runtime *util.RuntimeOptions) (_result *GetCloudMetricLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetCloudMetricLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCloudMetricLogs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudMetricLogs(request *GetCloudMetricLogsRequest) (_result *GetCloudMetricLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudMetricLogsResponse{}
	_body, _err := client.GetCloudMetricLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCloudMetricProfilingWithOptions(request *GetCloudMetricProfilingRequest, runtime *util.RuntimeOptions) (_result *GetCloudMetricProfilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetCloudMetricProfilingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetCloudMetricProfiling"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCloudMetricProfiling(request *GetCloudMetricProfilingRequest) (_result *GetCloudMetricProfilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetCloudMetricProfilingResponse{}
	_body, _err := client.GetCloudMetricProfilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterVolumesWithOptions(request *GetClusterVolumesRequest, runtime *util.RuntimeOptions) (_result *GetClusterVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetClusterVolumesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetClusterVolumes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterVolumes(request *GetClusterVolumesRequest) (_result *GetClusterVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterVolumesResponse{}
	_body, _err := client.GetClusterVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGWSConnectTicketWithOptions(request *GetGWSConnectTicketRequest, runtime *util.RuntimeOptions) (_result *GetGWSConnectTicketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetGWSConnectTicketResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetGWSConnectTicket"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGWSConnectTicket(request *GetGWSConnectTicketRequest) (_result *GetGWSConnectTicketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetGWSConnectTicketResponse{}
	_body, _err := client.GetGWSConnectTicketWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHealthMonitorLogsWithOptions(request *GetHealthMonitorLogsRequest, runtime *util.RuntimeOptions) (_result *GetHealthMonitorLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHealthMonitorLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHealthMonitorLogs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHealthMonitorLogs(request *GetHealthMonitorLogsRequest) (_result *GetHealthMonitorLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHealthMonitorLogsResponse{}
	_body, _err := client.GetHealthMonitorLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetHybridClusterConfigWithOptions(request *GetHybridClusterConfigRequest, runtime *util.RuntimeOptions) (_result *GetHybridClusterConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetHybridClusterConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetHybridClusterConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetHybridClusterConfig(request *GetHybridClusterConfigRequest) (_result *GetHybridClusterConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetHybridClusterConfigResponse{}
	_body, _err := client.GetHybridClusterConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetIfEcsTypeSupportHtConfigWithOptions(request *GetIfEcsTypeSupportHtConfigRequest, runtime *util.RuntimeOptions) (_result *GetIfEcsTypeSupportHtConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetIfEcsTypeSupportHtConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetIfEcsTypeSupportHtConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetIfEcsTypeSupportHtConfig(request *GetIfEcsTypeSupportHtConfigRequest) (_result *GetIfEcsTypeSupportHtConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetIfEcsTypeSupportHtConfigResponse{}
	_body, _err := client.GetIfEcsTypeSupportHtConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVisualServiceStatusWithOptions(request *GetVisualServiceStatusRequest, runtime *util.RuntimeOptions) (_result *GetVisualServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetVisualServiceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVisualServiceStatus"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVisualServiceStatus(request *GetVisualServiceStatusRequest) (_result *GetVisualServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVisualServiceStatusResponse{}
	_body, _err := client.GetVisualServiceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkbenchTokenWithOptions(request *GetWorkbenchTokenRequest, runtime *util.RuntimeOptions) (_result *GetWorkbenchTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &GetWorkbenchTokenResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetWorkbenchToken"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkbenchToken(request *GetWorkbenchTokenRequest) (_result *GetWorkbenchTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetWorkbenchTokenResponse{}
	_body, _err := client.GetWorkbenchTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitializeEHPCWithOptions(request *InitializeEHPCRequest, runtime *util.RuntimeOptions) (_result *InitializeEHPCResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &InitializeEHPCResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitializeEHPC"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitializeEHPC(request *InitializeEHPCRequest) (_result *InitializeEHPCResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitializeEHPCResponse{}
	_body, _err := client.InitializeEHPCWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallSoftwareWithOptions(request *InstallSoftwareRequest, runtime *util.RuntimeOptions) (_result *InstallSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &InstallSoftwareResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InstallSoftware"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallSoftware(request *InstallSoftwareRequest) (_result *InstallSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallSoftwareResponse{}
	_body, _err := client.InstallSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InvokeShellCommandWithOptions(request *InvokeShellCommandRequest, runtime *util.RuntimeOptions) (_result *InvokeShellCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &InvokeShellCommandResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InvokeShellCommand"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InvokeShellCommand(request *InvokeShellCommandRequest) (_result *InvokeShellCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InvokeShellCommandResponse{}
	_body, _err := client.InvokeShellCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAccountMappingWithOptions(request *ListAccountMappingRequest, runtime *util.RuntimeOptions) (_result *ListAccountMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAccountMappingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAccountMapping"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAccountMapping(request *ListAccountMappingRequest) (_result *ListAccountMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAccountMappingResponse{}
	_body, _err := client.ListAccountMappingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableEcsTypesWithOptions(request *ListAvailableEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListAvailableEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListAvailableEcsTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAvailableEcsTypes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableEcsTypes(request *ListAvailableEcsTypesRequest) (_result *ListAvailableEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableEcsTypesResponse{}
	_body, _err := client.ListAvailableEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableFileSystemTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListAvailableFileSystemTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListAvailableFileSystemTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAvailableFileSystemTypes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableFileSystemTypes() (_result *ListAvailableFileSystemTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAvailableFileSystemTypesResponse{}
	_body, _err := client.ListAvailableFileSystemTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCloudMetricProfilingsWithOptions(request *ListCloudMetricProfilingsRequest, runtime *util.RuntimeOptions) (_result *ListCloudMetricProfilingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListCloudMetricProfilingsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCloudMetricProfilings"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCloudMetricProfilings(request *ListCloudMetricProfilingsRequest) (_result *ListCloudMetricProfilingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCloudMetricProfilingsResponse{}
	_body, _err := client.ListCloudMetricProfilingsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterLogsWithOptions(request *ListClusterLogsRequest, runtime *util.RuntimeOptions) (_result *ListClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusterLogs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterLogs(request *ListClusterLogsRequest) (_result *ListClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterLogsResponse{}
	_body, _err := client.ListClusterLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListClustersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClusters"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersMetaWithOptions(request *ListClustersMetaRequest, runtime *util.RuntimeOptions) (_result *ListClustersMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListClustersMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListClustersMeta"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClustersMeta(request *ListClustersMetaRequest) (_result *ListClustersMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersMetaResponse{}
	_body, _err := client.ListClustersMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCommandsWithOptions(request *ListCommandsRequest, runtime *util.RuntimeOptions) (_result *ListCommandsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListCommandsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCommands"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCommands(request *ListCommandsRequest) (_result *ListCommandsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCommandsResponse{}
	_body, _err := client.ListCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerAppsWithOptions(request *ListContainerAppsRequest, runtime *util.RuntimeOptions) (_result *ListContainerAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListContainerAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListContainerApps"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerApps(request *ListContainerAppsRequest) (_result *ListContainerAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerAppsResponse{}
	_body, _err := client.ListContainerAppsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListContainerImagesWithOptions(request *ListContainerImagesRequest, runtime *util.RuntimeOptions) (_result *ListContainerImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListContainerImages"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListContainerImages(request *ListContainerImagesRequest) (_result *ListContainerImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListContainerImagesResponse{}
	_body, _err := client.ListContainerImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCpfsFileSystemsWithOptions(request *ListCpfsFileSystemsRequest, runtime *util.RuntimeOptions) (_result *ListCpfsFileSystemsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListCpfsFileSystemsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCpfsFileSystems"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCpfsFileSystems(request *ListCpfsFileSystemsRequest) (_result *ListCpfsFileSystemsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCpfsFileSystemsResponse{}
	_body, _err := client.ListCpfsFileSystemsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCurrentClientVersionWithOptions(runtime *util.RuntimeOptions) (_result *ListCurrentClientVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCurrentClientVersion"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCurrentClientVersion() (_result *ListCurrentClientVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCurrentClientVersionResponse{}
	_body, _err := client.ListCurrentClientVersionWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCustomImagesWithOptions(request *ListCustomImagesRequest, runtime *util.RuntimeOptions) (_result *ListCustomImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListCustomImages"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCustomImages(request *ListCustomImagesRequest) (_result *ListCustomImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListCustomImagesResponse{}
	_body, _err := client.ListCustomImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFileSystemWithMountTargetsWithOptions(request *ListFileSystemWithMountTargetsRequest, runtime *util.RuntimeOptions) (_result *ListFileSystemWithMountTargetsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListFileSystemWithMountTargetsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFileSystemWithMountTargets"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFileSystemWithMountTargets(request *ListFileSystemWithMountTargetsRequest) (_result *ListFileSystemWithMountTargetsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFileSystemWithMountTargetsResponse{}
	_body, _err := client.ListFileSystemWithMountTargetsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListImages"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstalledSoftwareWithOptions(request *ListInstalledSoftwareRequest, runtime *util.RuntimeOptions) (_result *ListInstalledSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInstalledSoftwareResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInstalledSoftware"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstalledSoftware(request *ListInstalledSoftwareRequest) (_result *ListInstalledSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstalledSoftwareResponse{}
	_body, _err := client.ListInstalledSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationResultsWithOptions(request *ListInvocationResultsRequest, runtime *util.RuntimeOptions) (_result *ListInvocationResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInvocationResultsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInvocationResults"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationResults(request *ListInvocationResultsRequest) (_result *ListInvocationResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationResultsResponse{}
	_body, _err := client.ListInvocationResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInvocationStatusWithOptions(request *ListInvocationStatusRequest, runtime *util.RuntimeOptions) (_result *ListInvocationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListInvocationStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListInvocationStatus"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInvocationStatus(request *ListInvocationStatusRequest) (_result *ListInvocationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInvocationStatusResponse{}
	_body, _err := client.ListInvocationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListJobs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobTemplatesWithOptions(request *ListJobTemplatesRequest, runtime *util.RuntimeOptions) (_result *ListJobTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListJobTemplates"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobTemplates(request *ListJobTemplatesRequest) (_result *ListJobTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListJobTemplatesResponse{}
	_body, _err := client.ListJobTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesWithOptions(request *ListNodesRequest, runtime *util.RuntimeOptions) (_result *ListNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodes(request *ListNodesRequest) (_result *ListNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesResponse{}
	_body, _err := client.ListNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesByQueueWithOptions(request *ListNodesByQueueRequest, runtime *util.RuntimeOptions) (_result *ListNodesByQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListNodesByQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNodesByQueue"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesByQueue(request *ListNodesByQueueRequest) (_result *ListNodesByQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesByQueueResponse{}
	_body, _err := client.ListNodesByQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListNodesNoPagingWithOptions(request *ListNodesNoPagingRequest, runtime *util.RuntimeOptions) (_result *ListNodesNoPagingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListNodesNoPaging"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListNodesNoPaging(request *ListNodesNoPagingRequest) (_result *ListNodesNoPagingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodesNoPagingResponse{}
	_body, _err := client.ListNodesNoPagingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPreferredEcsTypesWithOptions(request *ListPreferredEcsTypesRequest, runtime *util.RuntimeOptions) (_result *ListPreferredEcsTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPreferredEcsTypes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPreferredEcsTypes(request *ListPreferredEcsTypesRequest) (_result *ListPreferredEcsTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPreferredEcsTypesResponse{}
	_body, _err := client.ListPreferredEcsTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueuesWithOptions(request *ListQueuesRequest, runtime *util.RuntimeOptions) (_result *ListQueuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListQueuesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListQueues"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueues(request *ListQueuesRequest) (_result *ListQueuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListQueuesResponse{}
	_body, _err := client.ListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegions"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityGroupsWithOptions(request *ListSecurityGroupsRequest, runtime *util.RuntimeOptions) (_result *ListSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSecurityGroups"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityGroups(request *ListSecurityGroupsRequest) (_result *ListSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.ListSecurityGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSoftwaresWithOptions(request *ListSoftwaresRequest, runtime *util.RuntimeOptions) (_result *ListSoftwaresResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListSoftwares"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSoftwares(request *ListSoftwaresRequest) (_result *ListSoftwaresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListSoftwaresResponse{}
	_body, _err := client.ListSoftwaresWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTasks"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUsersWithOptions(request *ListUsersRequest, runtime *util.RuntimeOptions) (_result *ListUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListUsers"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUsers(request *ListUsersRequest) (_result *ListUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUsersResponse{}
	_body, _err := client.ListUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVolumesWithOptions(request *ListVolumesRequest, runtime *util.RuntimeOptions) (_result *ListVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ListVolumesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListVolumes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVolumes(request *ListVolumesRequest) (_result *ListVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVolumesResponse{}
	_body, _err := client.ListVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterAttributesWithOptions(request *ModifyClusterAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyClusterAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyClusterAttributes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyClusterAttributes(request *ModifyClusterAttributesRequest) (_result *ModifyClusterAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyClusterAttributesResponse{}
	_body, _err := client.ModifyClusterAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyContainerAppAttributesWithOptions(request *ModifyContainerAppAttributesRequest, runtime *util.RuntimeOptions) (_result *ModifyContainerAppAttributesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyContainerAppAttributesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyContainerAppAttributes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyContainerAppAttributes(request *ModifyContainerAppAttributesRequest) (_result *ModifyContainerAppAttributesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyContainerAppAttributesResponse{}
	_body, _err := client.ModifyContainerAppAttributesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyImageGatewayConfigWithOptions(request *ModifyImageGatewayConfigRequest, runtime *util.RuntimeOptions) (_result *ModifyImageGatewayConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyImageGatewayConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyImageGatewayConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyImageGatewayConfig(request *ModifyImageGatewayConfigRequest) (_result *ModifyImageGatewayConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyImageGatewayConfigResponse{}
	_body, _err := client.ModifyImageGatewayConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserGroupsWithOptions(request *ModifyUserGroupsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserGroups"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserGroups(request *ModifyUserGroupsRequest) (_result *ModifyUserGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserGroupsResponse{}
	_body, _err := client.ModifyUserGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyUserPasswordsWithOptions(request *ModifyUserPasswordsRequest, runtime *util.RuntimeOptions) (_result *ModifyUserPasswordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyUserPasswords"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyUserPasswords(request *ModifyUserPasswordsRequest) (_result *ModifyUserPasswordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyUserPasswordsResponse{}
	_body, _err := client.ModifyUserPasswordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyVisualServicePasswdWithOptions(request *ModifyVisualServicePasswdRequest, runtime *util.RuntimeOptions) (_result *ModifyVisualServicePasswdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ModifyVisualServicePasswdResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyVisualServicePasswd"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyVisualServicePasswd(request *ModifyVisualServicePasswdRequest) (_result *ModifyVisualServicePasswdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyVisualServicePasswdResponse{}
	_body, _err := client.ModifyVisualServicePasswdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MountNFSWithOptions(request *MountNFSRequest, runtime *util.RuntimeOptions) (_result *MountNFSResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &MountNFSResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MountNFS"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MountNFS(request *MountNFSRequest) (_result *MountNFSResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MountNFSResponse{}
	_body, _err := client.MountNFSWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PullImageWithOptions(request *PullImageRequest, runtime *util.RuntimeOptions) (_result *PullImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &PullImageResponse{}
	_body, _err := client.DoRPCRequest(tea.String("PullImage"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PullImage(request *PullImageRequest) (_result *PullImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &PullImageResponse{}
	_body, _err := client.PullImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) QueryServicePackAndPriceWithOptions(runtime *util.RuntimeOptions) (_result *QueryServicePackAndPriceResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &QueryServicePackAndPriceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("QueryServicePackAndPrice"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) QueryServicePackAndPrice() (_result *QueryServicePackAndPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &QueryServicePackAndPriceResponse{}
	_body, _err := client.QueryServicePackAndPriceWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecoverClusterWithOptions(request *RecoverClusterRequest, runtime *util.RuntimeOptions) (_result *RecoverClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RecoverClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RecoverCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecoverCluster(request *RecoverClusterRequest) (_result *RecoverClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RecoverClusterResponse{}
	_body, _err := client.RecoverClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RerunJobsWithOptions(request *RerunJobsRequest, runtime *util.RuntimeOptions) (_result *RerunJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RerunJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RerunJobs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RerunJobs(request *RerunJobsRequest) (_result *RerunJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RerunJobsResponse{}
	_body, _err := client.RerunJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetNodesWithOptions(request *ResetNodesRequest, runtime *util.RuntimeOptions) (_result *ResetNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &ResetNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetNodes(request *ResetNodesRequest) (_result *ResetNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetNodesResponse{}
	_body, _err := client.ResetNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCloudMetricProfilingWithOptions(request *RunCloudMetricProfilingRequest, runtime *util.RuntimeOptions) (_result *RunCloudMetricProfilingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &RunCloudMetricProfilingResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RunCloudMetricProfiling"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCloudMetricProfiling(request *RunCloudMetricProfilingRequest) (_result *RunCloudMetricProfilingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCloudMetricProfilingResponse{}
	_body, _err := client.RunCloudMetricProfilingWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAutoScaleConfigWithOptions(request *SetAutoScaleConfigRequest, runtime *util.RuntimeOptions) (_result *SetAutoScaleConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAutoScaleConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAutoScaleConfig(request *SetAutoScaleConfigRequest) (_result *SetAutoScaleConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAutoScaleConfigResponse{}
	_body, _err := client.SetAutoScaleConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSClusterPolicyWithOptions(request *SetGWSClusterPolicyRequest, runtime *util.RuntimeOptions) (_result *SetGWSClusterPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetGWSClusterPolicyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGWSClusterPolicy"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSClusterPolicy(request *SetGWSClusterPolicyRequest) (_result *SetGWSClusterPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSClusterPolicyResponse{}
	_body, _err := client.SetGWSClusterPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSInstanceNameWithOptions(request *SetGWSInstanceNameRequest, runtime *util.RuntimeOptions) (_result *SetGWSInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SetGWSInstanceNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGWSInstanceName"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSInstanceName(request *SetGWSInstanceNameRequest) (_result *SetGWSInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSInstanceNameResponse{}
	_body, _err := client.SetGWSInstanceNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetGWSInstanceUserWithOptions(request *SetGWSInstanceUserRequest, runtime *util.RuntimeOptions) (_result *SetGWSInstanceUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SetGWSInstanceUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetGWSInstanceUser"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetGWSInstanceUser(request *SetGWSInstanceUserRequest) (_result *SetGWSInstanceUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetGWSInstanceUserResponse{}
	_body, _err := client.SetGWSInstanceUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetJobUserWithOptions(request *SetJobUserRequest, runtime *util.RuntimeOptions) (_result *SetJobUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SetJobUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetJobUser"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetJobUser(request *SetJobUserRequest) (_result *SetJobUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetJobUserResponse{}
	_body, _err := client.SetJobUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetQueueWithOptions(request *SetQueueRequest, runtime *util.RuntimeOptions) (_result *SetQueueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SetQueueResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetQueue"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetQueue(request *SetQueueRequest) (_result *SetQueueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQueueResponse{}
	_body, _err := client.SetQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartClusterWithOptions(request *StartClusterRequest, runtime *util.RuntimeOptions) (_result *StartClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StartClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartCluster(request *StartClusterRequest) (_result *StartClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartClusterResponse{}
	_body, _err := client.StartClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartGWSInstanceWithOptions(request *StartGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *StartGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StartGWSInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartGWSInstance"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartGWSInstance(request *StartGWSInstanceRequest) (_result *StartGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartGWSInstanceResponse{}
	_body, _err := client.StartGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartNodesWithOptions(request *StartNodesRequest, runtime *util.RuntimeOptions) (_result *StartNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StartNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartNodes(request *StartNodesRequest) (_result *StartNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartNodesResponse{}
	_body, _err := client.StartNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartVisualServiceWithOptions(request *StartVisualServiceRequest, runtime *util.RuntimeOptions) (_result *StartVisualServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StartVisualServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartVisualService"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartVisualService(request *StartVisualServiceRequest) (_result *StartVisualServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartVisualServiceResponse{}
	_body, _err := client.StartVisualServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopClusterWithOptions(request *StopClusterRequest, runtime *util.RuntimeOptions) (_result *StopClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StopClusterResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopCluster"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopCluster(request *StopClusterRequest) (_result *StopClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopClusterResponse{}
	_body, _err := client.StopClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopGWSInstanceWithOptions(request *StopGWSInstanceRequest, runtime *util.RuntimeOptions) (_result *StopGWSInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StopGWSInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopGWSInstance"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopGWSInstance(request *StopGWSInstanceRequest) (_result *StopGWSInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopGWSInstanceResponse{}
	_body, _err := client.StopGWSInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobsWithOptions(request *StopJobsRequest, runtime *util.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StopJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopJobs"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJobs(request *StopJobsRequest) (_result *StopJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopJobsResponse{}
	_body, _err := client.StopJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopNodesWithOptions(request *StopNodesRequest, runtime *util.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StopNodesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopNodes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopNodes(request *StopNodesRequest) (_result *StopNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopNodesResponse{}
	_body, _err := client.StopNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopVisualServiceWithOptions(request *StopVisualServiceRequest, runtime *util.RuntimeOptions) (_result *StopVisualServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &StopVisualServiceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopVisualService"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopVisualService(request *StopVisualServiceRequest) (_result *StopVisualServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopVisualServiceResponse{}
	_body, _err := client.StopVisualServiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitJobWithOptions(request *SubmitJobRequest, runtime *util.RuntimeOptions) (_result *SubmitJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &SubmitJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SubmitJob"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitJob(request *SubmitJobRequest) (_result *SubmitJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SubmitJobResponse{}
	_body, _err := client.SubmitJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindAccountToClusterUserWithOptions(request *UnbindAccountToClusterUserRequest, runtime *util.RuntimeOptions) (_result *UnbindAccountToClusterUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UnbindAccountToClusterUserResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnbindAccountToClusterUser"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindAccountToClusterUser(request *UnbindAccountToClusterUserRequest) (_result *UnbindAccountToClusterUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnbindAccountToClusterUserResponse{}
	_body, _err := client.UnbindAccountToClusterUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UninstallSoftwareWithOptions(request *UninstallSoftwareRequest, runtime *util.RuntimeOptions) (_result *UninstallSoftwareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UninstallSoftwareResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UninstallSoftware"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallSoftware(request *UninstallSoftwareRequest) (_result *UninstallSoftwareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallSoftwareResponse{}
	_body, _err := client.UninstallSoftwareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateClusterVolumesWithOptions(request *UpdateClusterVolumesRequest, runtime *util.RuntimeOptions) (_result *UpdateClusterVolumesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UpdateClusterVolumesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateClusterVolumes"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateClusterVolumes(request *UpdateClusterVolumesRequest) (_result *UpdateClusterVolumesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateClusterVolumesResponse{}
	_body, _err := client.UpdateClusterVolumesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQueueConfigWithOptions(request *UpdateQueueConfigRequest, runtime *util.RuntimeOptions) (_result *UpdateQueueConfigResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UpdateQueueConfigResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateQueueConfig"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQueueConfig(request *UpdateQueueConfigRequest) (_result *UpdateQueueConfigResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateQueueConfigResponse{}
	_body, _err := client.UpdateQueueConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeClientWithOptions(request *UpgradeClientRequest, runtime *util.RuntimeOptions) (_result *UpgradeClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := openapiutil.Query(util.ToMap(request))
	req := &openapi.OpenApiRequest{
		Query: query,
	}
	_result = &UpgradeClientResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeClient"), tea.String("2018-04-12"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeClient(request *UpgradeClientRequest) (_result *UpgradeClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeClientResponse{}
	_body, _err := client.UpgradeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
