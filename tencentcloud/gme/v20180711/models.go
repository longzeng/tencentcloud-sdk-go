// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180711

import (
    "encoding/json"
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AgeDetectTask struct {
	// 数据唯一ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 数据文件的url，为 urlencode 编码,音频文件格式支持的类型：.wav、.m4a、.amr、.mp3、.aac、.wma、.ogg
	Url *string `json:"Url,omitempty" name:"Url"`
}

type AgeDetectTaskResult struct {
	// 数据唯一ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 数据文件的url
	Url *string `json:"Url,omitempty" name:"Url"`

	// 任务状态，0: 已创建，1:运行中，2:正常结束，3:异常结束，4:运行超时
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 任务结果：0: 成年，1:未成年，100:未知
	Age *uint64 `json:"Age,omitempty" name:"Age"`
}

type AppStatisticsItem struct {
	// 实时语音统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealtimeSpeechStatisticsItem *RealTimeSpeechStatisticsItem `json:"RealtimeSpeechStatisticsItem,omitempty" name:"RealtimeSpeechStatisticsItem"`

	// 语音消息统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceMessageStatisticsItem *VoiceMessageStatisticsItem `json:"VoiceMessageStatisticsItem,omitempty" name:"VoiceMessageStatisticsItem"`

	// 语音过滤统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceFilterStatisticsItem *VoiceFilterStatisticsItem `json:"VoiceFilterStatisticsItem,omitempty" name:"VoiceFilterStatisticsItem"`

	// 统计时间
	Date *string `json:"Date,omitempty" name:"Date"`

	// 录音转文本用量统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioTextStatisticsItem *AudioTextStatisticsItem `json:"AudioTextStatisticsItem,omitempty" name:"AudioTextStatisticsItem"`

	// 流式转文本用量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamTextStatisticsItem *StreamTextStatisticsItem `json:"StreamTextStatisticsItem,omitempty" name:"StreamTextStatisticsItem"`

	// 海外转文本用量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	OverseaTextStatisticsItem *OverseaTextStatisticsItem `json:"OverseaTextStatisticsItem,omitempty" name:"OverseaTextStatisticsItem"`

	// 实时语音转文本用量数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	RealtimeTextStatisticsItem *RealtimeTextStatisticsItem `json:"RealtimeTextStatisticsItem,omitempty" name:"RealtimeTextStatisticsItem"`
}

type ApplicationDataStatistics struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// Dau统计项数目
	DauDataNum *uint64 `json:"DauDataNum,omitempty" name:"DauDataNum"`

	// 大陆地区Dau统计数据，单位人
	DauDataMainland []*StatisticsItem `json:"DauDataMainland,omitempty" name:"DauDataMainland"`

	// 海外地区Dau统计数据，单位人
	DauDataOversea []*StatisticsItem `json:"DauDataOversea,omitempty" name:"DauDataOversea"`

	// 大陆和海外地区Dau统计数据汇总，单位人
	DauDataSum []*StatisticsItem `json:"DauDataSum,omitempty" name:"DauDataSum"`

	// 实时语音时长统计项数目
	DurationDataNum *uint64 `json:"DurationDataNum,omitempty" name:"DurationDataNum"`

	// 大陆地区实时语音时长统计数据，单位分钟
	DurationDataMainland []*StatisticsItem `json:"DurationDataMainland,omitempty" name:"DurationDataMainland"`

	// 海外地区实时语音时长统计数据，单位分钟
	DurationDataOversea []*StatisticsItem `json:"DurationDataOversea,omitempty" name:"DurationDataOversea"`

	// 大陆和海外地区实时语音时长统计数据汇总，单位分钟
	DurationDataSum []*StatisticsItem `json:"DurationDataSum,omitempty" name:"DurationDataSum"`

	// Pcu统计项数目
	PcuDataNum *uint64 `json:"PcuDataNum,omitempty" name:"PcuDataNum"`

	// 大陆地区Pcu统计数据，单位人
	PcuDataMainland []*StatisticsItem `json:"PcuDataMainland,omitempty" name:"PcuDataMainland"`

	// 海外地区Pcu统计数据，单位人
	PcuDataOversea []*StatisticsItem `json:"PcuDataOversea,omitempty" name:"PcuDataOversea"`

	// 大陆和海外地区Pcu统计数据汇总，单位人
	PcuDataSum []*StatisticsItem `json:"PcuDataSum,omitempty" name:"PcuDataSum"`
}

type AudioTextStatisticsItem struct {
	// 统计值，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

// Predefined struct for user
type CreateAgeDetectTaskRequestParams struct {
	// 应用id
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 语音检测子任务列表，列表最多支持100个检测子任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*AgeDetectTask `json:"Tasks,omitempty" name:"Tasks"`

	// 任务结束时gme后台会自动触发回调
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

type CreateAgeDetectTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用id
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 语音检测子任务列表，列表最多支持100个检测子任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*AgeDetectTask `json:"Tasks,omitempty" name:"Tasks"`

	// 任务结束时gme后台会自动触发回调
	Callback *string `json:"Callback,omitempty" name:"Callback"`
}

func (r *CreateAgeDetectTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgeDetectTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Tasks")
	delete(f, "Callback")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAgeDetectTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAgeDetectTaskResponseParams struct {
	// 本次任务提交后唯一id，用于获取任务运行结果
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAgeDetectTaskResponse struct {
	*tchttp.BaseResponse
	Response *CreateAgeDetectTaskResponseParams `json:"Response"`
}

func (r *CreateAgeDetectTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAgeDetectTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateAppRequestParams struct {
	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 腾讯云项目ID，默认为0，表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 需要支持的引擎列表，默认全选。
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList"`

	// 服务区域列表，默认全选。
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息及转文本服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`

	// 需要添加的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

type CreateAppRequest struct {
	*tchttp.BaseRequest
	
	// 应用名称
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 腾讯云项目ID，默认为0，表示默认项目
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 需要支持的引擎列表，默认全选。
	EngineList []*string `json:"EngineList,omitempty" name:"EngineList"`

	// 服务区域列表，默认全选。
	RegionList []*string `json:"RegionList,omitempty" name:"RegionList"`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息及转文本服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`

	// 需要添加的标签列表
	Tags []*Tag `json:"Tags,omitempty" name:"Tags"`
}

func (r *CreateAppRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "AppName")
	delete(f, "ProjectId")
	delete(f, "EngineList")
	delete(f, "RegionList")
	delete(f, "RealtimeSpeechConf")
	delete(f, "VoiceMessageConf")
	delete(f, "VoiceFilterConf")
	delete(f, "Tags")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateAppRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type CreateAppResp struct {
	// 应用ID，由后台自动生成。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 应用名称，透传输入参数的AppName
	AppName *string `json:"AppName,omitempty" name:"AppName"`

	// 项目ID，透传输入的ProjectId
	ProjectId *uint64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 应用密钥，GME SDK初始化时使用
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 服务创建时间戳
	CreateTime *uint64 `json:"CreateTime,omitempty" name:"CreateTime"`

	// 实时语音服务配置数据
	RealtimeSpeechConf *RealtimeSpeechConf `json:"RealtimeSpeechConf,omitempty" name:"RealtimeSpeechConf"`

	// 语音消息及转文本服务配置数据
	VoiceMessageConf *VoiceMessageConf `json:"VoiceMessageConf,omitempty" name:"VoiceMessageConf"`

	// 语音分析服务配置数据
	VoiceFilterConf *VoiceFilterConf `json:"VoiceFilterConf,omitempty" name:"VoiceFilterConf"`
}

// Predefined struct for user
type CreateAppResponseParams struct {
	// 创建应用返回数据
	Data *CreateAppResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateAppResponse struct {
	*tchttp.BaseResponse
	Response *CreateAppResponseParams `json:"Response"`
}

func (r *CreateAppResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateAppResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizationRequestParams struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`

	// 模型名称，名称长度不超过36，默认为BizId。
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`
}

type CreateCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`

	// 模型名称，名称长度不超过36，默认为BizId。
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`
}

func (r *CreateCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TextUrl")
	delete(f, "ModelName")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateCustomizationResponseParams struct {
	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *CreateCustomizationResponseParams `json:"Response"`
}

func (r *CreateCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanUserRequestParams struct {
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要新增送检的用户号。示例：1234
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

type CreateScanUserRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要新增送检的用户号。示例：1234
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *CreateScanUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "CreateScanUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type CreateScanUserResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type CreateScanUserResponse struct {
	*tchttp.BaseResponse
	Response *CreateScanUserResponseParams `json:"Response"`
}

func (r *CreateScanUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *CreateScanUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type CustomizationConfigs struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 模型状态，-1下线状态，1上线状态, 0训练中, -2训练失败, 3上线中, 4下线中
	ModelState *int64 `json:"ModelState,omitempty" name:"ModelState"`

	// 模型名称
	ModelName *string `json:"ModelName,omitempty" name:"ModelName"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`

	// 更新时间，11位时间戳
	UpdateTime *int64 `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

// Predefined struct for user
type DeleteCustomizationRequestParams struct {
	// 要删除的模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

type DeleteCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 要删除的模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *DeleteCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteCustomizationResponseParams struct {
	// 返回值。0为成功，非0为失败。
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *DeleteCustomizationResponseParams `json:"Response"`
}

func (r *DeleteCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteResult struct {
	// 错误码，0-剔除成功 其他-剔除失败
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 错误描述
	ErrorMsg *string `json:"ErrorMsg,omitempty" name:"ErrorMsg"`
}

// Predefined struct for user
type DeleteRoomMemberRequestParams struct {
	// 要操作的房间id
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 要剔除的用户列表
	Uids []*string `json:"Uids,omitempty" name:"Uids"`

	// 剔除类型 1-删除房间 2-剔除用户
	DeleteType *uint64 `json:"DeleteType,omitempty" name:"DeleteType"`

	// 应用id
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

type DeleteRoomMemberRequest struct {
	*tchttp.BaseRequest
	
	// 要操作的房间id
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 要剔除的用户列表
	Uids []*string `json:"Uids,omitempty" name:"Uids"`

	// 剔除类型 1-删除房间 2-剔除用户
	DeleteType *uint64 `json:"DeleteType,omitempty" name:"DeleteType"`

	// 应用id
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *DeleteRoomMemberRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomMemberRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "RoomId")
	delete(f, "Uids")
	delete(f, "DeleteType")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteRoomMemberRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteRoomMemberResponseParams struct {
	// 剔除房间或成员的操作结果
	DeleteResult *DeleteResult `json:"DeleteResult,omitempty" name:"DeleteResult"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteRoomMemberResponse struct {
	*tchttp.BaseResponse
	Response *DeleteRoomMemberResponseParams `json:"Response"`
}

func (r *DeleteRoomMemberResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteRoomMemberResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanUserRequestParams struct {
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要删除送检的用户号。示例：1234
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

type DeleteScanUserRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录控制台 - 服务管理创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要删除送检的用户号。示例：1234
	UserId *uint64 `json:"UserId,omitempty" name:"UserId"`
}

func (r *DeleteScanUserRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanUserRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "UserId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DeleteScanUserRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DeleteScanUserResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DeleteScanUserResponse struct {
	*tchttp.BaseResponse
	Response *DeleteScanUserResponseParams `json:"Response"`
}

func (r *DeleteScanUserResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DeleteScanUserResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgeDetectTaskRequestParams struct {
	// 应用id
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 创建年龄语音识别任务时返回的taskid
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type DescribeAgeDetectTaskRequest struct {
	*tchttp.BaseRequest
	
	// 应用id
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 创建年龄语音识别任务时返回的taskid
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeAgeDetectTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgeDetectTaskRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TaskId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAgeDetectTaskRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAgeDetectTaskResponseParams struct {
	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 语音检测返回。Results 字段是 JSON 数组，每一个元素包含：
	// DataId： 请求中对应的 DataId。
	// Url ：该请求中对应的 Url。
	// Status ：子任务状态，0:已创建，1:运行中，2:已完成，3:任务异常，4:任务超时。
	// Age ：子任务完成后的结果，0:成年人，1:未成年人，100:未知结果。
	Results []*AgeDetectTaskResult `json:"Results,omitempty" name:"Results"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAgeDetectTaskResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAgeDetectTaskResponseParams `json:"Response"`
}

func (r *DescribeAgeDetectTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAgeDetectTaskResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeAppStatisticsRequestParams struct {
	// GME应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 数据开始时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 数据结束时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 要查询的服务列表，取值：RealTimeSpeech/VoiceMessage/VoiceFilter/SpeechToText
	Services []*string `json:"Services,omitempty" name:"Services"`
}

type DescribeAppStatisticsRequest struct {
	*tchttp.BaseRequest
	
	// GME应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 数据开始时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 数据结束时间，东八区时间，格式: 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`

	// 要查询的服务列表，取值：RealTimeSpeech/VoiceMessage/VoiceFilter/SpeechToText
	Services []*string `json:"Services,omitempty" name:"Services"`
}

func (r *DescribeAppStatisticsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	delete(f, "Services")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeAppStatisticsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type DescribeAppStatisticsResp struct {
	// 应用用量统计数据
	AppStatistics []*AppStatisticsItem `json:"AppStatistics,omitempty" name:"AppStatistics"`
}

// Predefined struct for user
type DescribeAppStatisticsResponseParams struct {
	// 应用用量统计数据
	Data *DescribeAppStatisticsResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeAppStatisticsResponse struct {
	*tchttp.BaseResponse
	Response *DescribeAppStatisticsResponseParams `json:"Response"`
}

func (r *DescribeAppStatisticsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeAppStatisticsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationDataRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 数据开始时间，格式为 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 数据结束时间，格式为 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

type DescribeApplicationDataRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 数据开始时间，格式为 年-月-日，如: 2018-07-13
	StartDate *string `json:"StartDate,omitempty" name:"StartDate"`

	// 数据结束时间，格式为 年-月-日，如: 2018-07-13
	EndDate *string `json:"EndDate,omitempty" name:"EndDate"`
}

func (r *DescribeApplicationDataRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationDataRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "StartDate")
	delete(f, "EndDate")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeApplicationDataRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeApplicationDataResponseParams struct {
	// 应用统计数据
	Data *ApplicationDataStatistics `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeApplicationDataResponse struct {
	*tchttp.BaseResponse
	Response *DescribeApplicationDataResponseParams `json:"Response"`
}

func (r *DescribeApplicationDataResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeApplicationDataResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealtimeScanConfigRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

type DescribeRealtimeScanConfigRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *DescribeRealtimeScanConfigRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeScanConfigRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRealtimeScanConfigRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRealtimeScanConfigResponseParams struct {
	// 返回结果码，0正常，非0失败
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 送检类型，0: 全量送审，1: 自定义送审
	AuditType *int64 `json:"AuditType,omitempty" name:"AuditType"`

	// 用户号正则表达式
	UserIdRegex []*string `json:"UserIdRegex,omitempty" name:"UserIdRegex"`

	// 房间号正则表达式
	RoomIdRegex []*string `json:"RoomIdRegex,omitempty" name:"RoomIdRegex"`

	// 用户号字符串，逗号分隔，示例："0001,0002,0003"
	UserIdString *string `json:"UserIdString,omitempty" name:"UserIdString"`

	// 房间号字符串，逗号分隔，示例："0001,0002,0003"
	RoomIdString *string `json:"RoomIdString,omitempty" name:"RoomIdString"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRealtimeScanConfigResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRealtimeScanConfigResponseParams `json:"Response"`
}

func (r *DescribeRealtimeScanConfigResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRealtimeScanConfigResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoRequestParams struct {
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	RoomIds []*uint64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// 字符串类型房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	StrRoomIds []*string `json:"StrRoomIds,omitempty" name:"StrRoomIds"`
}

type DescribeRoomInfoRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	SdkAppId *uint64 `json:"SdkAppId,omitempty" name:"SdkAppId"`

	// 房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	RoomIds []*uint64 `json:"RoomIds,omitempty" name:"RoomIds"`

	// 字符串类型房间号列表，最大不能超过10个（RoomIds、StrRoomIds必须填一个）
	StrRoomIds []*string `json:"StrRoomIds,omitempty" name:"StrRoomIds"`
}

func (r *DescribeRoomInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInfoRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "SdkAppId")
	delete(f, "RoomIds")
	delete(f, "StrRoomIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeRoomInfoRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeRoomInfoResponseParams struct {
	// 操作结果, 0成功, 非0失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	Result *uint64 `json:"Result,omitempty" name:"Result"`

	// 房间用户信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomUsers []*RoomUser `json:"RoomUsers,omitempty" name:"RoomUsers"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeRoomInfoResponse struct {
	*tchttp.BaseResponse
	Response *DescribeRoomInfoResponseParams `json:"Response"`
}

func (r *DescribeRoomInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeRoomInfoResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DescribeScanResult struct {
	// 业务返回码
	Code *int64 `json:"Code,omitempty" name:"Code"`

	// 数据唯一 ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 检测完成的时间戳
	ScanFinishTime *uint64 `json:"ScanFinishTime,omitempty" name:"ScanFinishTime"`

	// 是否违规
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// 是否为流
	Live *bool `json:"Live,omitempty" name:"Live"`

	// 业务返回描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Msg *string `json:"Msg,omitempty" name:"Msg"`

	// 检测结果，Code 为 0 时返回
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScanPiece []*ScanPiece `json:"ScanPiece,omitempty" name:"ScanPiece"`

	// 提交检测的时间戳
	ScanStartTime *uint64 `json:"ScanStartTime,omitempty" name:"ScanStartTime"`

	// 语音检测场景，对应请求时的 Scene
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// 语音检测任务 ID，由后台分配
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 文件或接流地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// 检测任务执行结果状态，分别为：
	// <li>Start: 任务开始</li>
	// <li>Success: 成功结束</li>
	// <li>Error: 异常</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 提交检测的应用 ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`
}

// Predefined struct for user
type DescribeScanResultListRequestParams struct {
	// 应用 ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 查询的任务 ID 列表，任务 ID 列表最多支持 100 个。
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 任务返回结果数量，默认10，上限500。大文件任务忽略此参数，返回全量结果
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

type DescribeScanResultListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录[控制台](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 查询的任务 ID 列表，任务 ID 列表最多支持 100 个。
	TaskIdList []*string `json:"TaskIdList,omitempty" name:"TaskIdList"`

	// 任务返回结果数量，默认10，上限500。大文件任务忽略此参数，返回全量结果
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScanResultListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TaskIdList")
	delete(f, "Limit")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeScanResultListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeScanResultListResponseParams struct {
	// 要查询的语音检测任务的结果
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data []*DescribeScanResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeScanResultListResponse struct {
	*tchttp.BaseResponse
	Response *DescribeScanResultListResponseParams `json:"Response"`
}

func (r *DescribeScanResultListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeScanResultListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInAndOutTimeRequestParams struct {
	// 应用ID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 房间ID
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 用户ID
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// 字符串类型用户ID
	UserIdStr *string `json:"UserIdStr,omitempty" name:"UserIdStr"`

	// 字符串类型房间ID
	RoomIdStr *string `json:"RoomIdStr,omitempty" name:"RoomIdStr"`
}

type DescribeUserInAndOutTimeRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 房间ID
	RoomId *int64 `json:"RoomId,omitempty" name:"RoomId"`

	// 用户ID
	UserId *int64 `json:"UserId,omitempty" name:"UserId"`

	// 字符串类型用户ID
	UserIdStr *string `json:"UserIdStr,omitempty" name:"UserIdStr"`

	// 字符串类型房间ID
	RoomIdStr *string `json:"RoomIdStr,omitempty" name:"RoomIdStr"`
}

func (r *DescribeUserInAndOutTimeRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInAndOutTimeRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	delete(f, "UserId")
	delete(f, "UserIdStr")
	delete(f, "RoomIdStr")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeUserInAndOutTimeRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeUserInAndOutTimeResponseParams struct {
	// 用户在房间得进出时间列表
	InOutList []*InOutTimeInfo `json:"InOutList,omitempty" name:"InOutList"`

	// 用户在房间中总时长
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeUserInAndOutTimeResponse struct {
	*tchttp.BaseResponse
	Response *DescribeUserInAndOutTimeResponseParams `json:"Response"`
}

func (r *DescribeUserInAndOutTimeResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeUserInAndOutTimeResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomizationListRequestParams struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

type GetCustomizationListRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *GetCustomizationListRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomizationListRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "GetCustomizationListRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type GetCustomizationListResponseParams struct {
	// 语音消息转文本自学习模型配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	CustomizationConfigs []*CustomizationConfigs `json:"CustomizationConfigs,omitempty" name:"CustomizationConfigs"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type GetCustomizationListResponse struct {
	*tchttp.BaseResponse
	Response *GetCustomizationListResponseParams `json:"Response"`
}

func (r *GetCustomizationListResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *GetCustomizationListResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InOutTimeInfo struct {
	// 进入房间时间
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 退出房间时间
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

// Predefined struct for user
type ModifyAppStatusRequestParams struct {
	// 应用ID，创建应用后由后台生成并返回。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`
}

type ModifyAppStatusRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，创建应用后由后台生成并返回。
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifyAppStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Status")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyAppStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

type ModifyAppStatusResp struct {
	// GME应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 应用状态，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`
}

// Predefined struct for user
type ModifyAppStatusResponseParams struct {
	// 修改应用开关状态返回数据
	Data *ModifyAppStatusResp `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyAppStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyAppStatusResponseParams `json:"Response"`
}

func (r *ModifyAppStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyAppStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationRequestParams struct {
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`

	// 要修改的模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

type ModifyCustomizationRequest struct {
	*tchttp.BaseRequest
	
	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 文本文件的下载地址，服务会从该地址下载文件，目前仅支持腾讯云cos
	TextUrl *string `json:"TextUrl,omitempty" name:"TextUrl"`

	// 要修改的模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`
}

func (r *ModifyCustomizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "TextUrl")
	delete(f, "ModelId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizationRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationResponseParams struct {
	// 返回值。0为成功，非0为失败。
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 自学习模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomizationResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizationResponseParams `json:"Response"`
}

func (r *ModifyCustomizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationStateRequestParams struct {
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitempty" name:"ToState"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

type ModifyCustomizationStateRequest struct {
	*tchttp.BaseRequest
	
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 想要变换的模型状态，-1代表下线，1代表上线
	ToState *int64 `json:"ToState,omitempty" name:"ToState"`

	// 应用 ID，登录控制台创建应用得到的AppID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`
}

func (r *ModifyCustomizationStateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationStateRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "ModelId")
	delete(f, "ToState")
	delete(f, "BizId")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyCustomizationStateRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyCustomizationStateResponseParams struct {
	// 自学习模型ID
	ModelId *string `json:"ModelId,omitempty" name:"ModelId"`

	// 返回值。0为成功，非0为失败。
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyCustomizationStateResponse struct {
	*tchttp.BaseResponse
	Response *ModifyCustomizationStateResponseParams `json:"Response"`
}

func (r *ModifyCustomizationStateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyCustomizationStateResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserMicStatusRequestParams struct {
	// 应用ID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 用户麦克风状态，数组长度不超过20
	Users []*UserMicStatus `json:"Users,omitempty" name:"Users"`
}

type ModifyUserMicStatusRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *int64 `json:"BizId,omitempty" name:"BizId"`

	// 房间ID
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// 用户麦克风状态，数组长度不超过20
	Users []*UserMicStatus `json:"Users,omitempty" name:"Users"`
}

func (r *ModifyUserMicStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserMicStatusRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomId")
	delete(f, "Users")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ModifyUserMicStatusRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ModifyUserMicStatusResponseParams struct {
	// 返回结果：0为成功，非0为失败
	Result *int64 `json:"Result,omitempty" name:"Result"`

	// 错误信息
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ModifyUserMicStatusResponse struct {
	*tchttp.BaseResponse
	Response *ModifyUserMicStatusResponseParams `json:"Response"`
}

func (r *ModifyUserMicStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ModifyUserMicStatusResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type OverseaTextStatisticsItem struct {
	// 统计值，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type RealTimeSpeechStatisticsItem struct {
	// 大陆地区DAU
	MainLandDau *uint64 `json:"MainLandDau,omitempty" name:"MainLandDau"`

	// 大陆地区PCU
	MainLandPcu *uint64 `json:"MainLandPcu,omitempty" name:"MainLandPcu"`

	// 大陆地区总使用时长，单位为min
	MainLandDuration *uint64 `json:"MainLandDuration,omitempty" name:"MainLandDuration"`

	// 海外地区DAU
	OverseaDau *uint64 `json:"OverseaDau,omitempty" name:"OverseaDau"`

	// 海外地区PCU
	OverseaPcu *uint64 `json:"OverseaPcu,omitempty" name:"OverseaPcu"`

	// 海外地区总使用时长，单位为min
	OverseaDuration *uint64 `json:"OverseaDuration,omitempty" name:"OverseaDuration"`
}

type RealtimeSpeechConf struct {
	// 实时语音服务开关，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`

	// 实时语音音质类型，取值：high-高音质
	Quality *string `json:"Quality,omitempty" name:"Quality"`
}

type RealtimeTextStatisticsItem struct {
	// 统计值，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type RoomUser struct {
	// 房间id
	RoomId *uint64 `json:"RoomId,omitempty" name:"RoomId"`

	// 房间里用户uin列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Uins []*uint64 `json:"Uins,omitempty" name:"Uins"`

	// 字符串房间id
	// 注意：此字段可能返回 null，表示取不到有效值。
	StrRoomId *string `json:"StrRoomId,omitempty" name:"StrRoomId"`
}

type ScanDetail struct {
	// 违规场景，参照<a href="https://cloud.tencent.com/document/product/607/37622#Label_Value">Label</a>定义
	Label *string `json:"Label,omitempty" name:"Label"`

	// 该场景下概率[0.00,100.00],分值越大违规概率越高
	Rate *string `json:"Rate,omitempty" name:"Rate"`

	// 违规关键字
	KeyWord *string `json:"KeyWord,omitempty" name:"KeyWord"`

	// 关键字在音频的开始时间，从0开始的偏移量，单位为毫秒
	StartTime *uint64 `json:"StartTime,omitempty" name:"StartTime"`

	// 关键字在音频的结束时间，从0开始的偏移量,，单位为毫秒
	EndTime *uint64 `json:"EndTime,omitempty" name:"EndTime"`
}

type ScanPiece struct {
	// 流检测时返回，音频转存地址，保留30min
	// 注意：此字段可能返回 null，表示取不到有效值。
	DumpUrl *string `json:"DumpUrl,omitempty" name:"DumpUrl"`

	// 是否违规
	HitFlag *bool `json:"HitFlag,omitempty" name:"HitFlag"`

	// 违规主要类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	MainType *string `json:"MainType,omitempty" name:"MainType"`

	// 语音检测详情
	ScanDetail []*ScanDetail `json:"ScanDetail,omitempty" name:"ScanDetail"`

	// gme实时语音房间ID，透传任务传入时的RoomId
	// 注意：此字段可能返回 null，表示取不到有效值。
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// gme实时语音用户ID，透传任务传入时的OpenId
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`

	// 备注
	// 注意：此字段可能返回 null，表示取不到有效值。
	Info *string `json:"Info,omitempty" name:"Info"`

	// 流检测时分片在流中的偏移时间，单位毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 流检测时分片时长
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 分片开始检测时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	PieceStartTime *uint64 `json:"PieceStartTime,omitempty" name:"PieceStartTime"`
}

// Predefined struct for user
type ScanVoiceRequestParams struct {
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 语音检测场景，参数值目前要求为 default。 预留场景设置： 谩骂、色情、广告、违禁等场景，<a href="#Label_Value">具体取值见上述 Label 说明。</a>
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// 是否为直播流。值为 false 时表示普通语音文件检测；为 true 时表示语音流检测。
	Live *bool `json:"Live,omitempty" name:"Live"`

	// 语音检测任务列表，列表最多支持100个检测任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// 异步检测结果回调地址，具体见上述<a href="#Callback_Declare">回调相关说明</a>。（说明：该字段为空时，必须通过接口(查询语音检测结果)获取检测结果）。
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// 语种，不传默认中文
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

type ScanVoiceRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID，登录[控制台 - 服务管理](https://console.cloud.tencent.com/gamegme)创建应用得到的AppID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 语音检测场景，参数值目前要求为 default。 预留场景设置： 谩骂、色情、广告、违禁等场景，<a href="#Label_Value">具体取值见上述 Label 说明。</a>
	Scenes []*string `json:"Scenes,omitempty" name:"Scenes"`

	// 是否为直播流。值为 false 时表示普通语音文件检测；为 true 时表示语音流检测。
	Live *bool `json:"Live,omitempty" name:"Live"`

	// 语音检测任务列表，列表最多支持100个检测任务。结构体中包含：
	// <li>DataId：数据的唯一ID</li>
	// <li>Url：数据文件的url，为 urlencode 编码，流式则为拉流地址</li>
	Tasks []*Task `json:"Tasks,omitempty" name:"Tasks"`

	// 异步检测结果回调地址，具体见上述<a href="#Callback_Declare">回调相关说明</a>。（说明：该字段为空时，必须通过接口(查询语音检测结果)获取检测结果）。
	Callback *string `json:"Callback,omitempty" name:"Callback"`

	// 语种，不传默认中文
	Lang *string `json:"Lang,omitempty" name:"Lang"`
}

func (r *ScanVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVoiceRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "Scenes")
	delete(f, "Live")
	delete(f, "Tasks")
	delete(f, "Callback")
	delete(f, "Lang")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "ScanVoiceRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type ScanVoiceResponseParams struct {
	// 语音检测返回。Data 字段是 JSON 数组，每一个元素包含：<li>DataId： 请求中对应的 DataId。</li>
	// <li>TaskID ：该检测任务的 ID，用于轮询语音检测结果。</li>
	Data []*ScanVoiceResult `json:"Data,omitempty" name:"Data"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type ScanVoiceResponse struct {
	*tchttp.BaseResponse
	Response *ScanVoiceResponseParams `json:"Response"`
}

func (r *ScanVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *ScanVoiceResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ScanVoiceResult struct {
	// 数据ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

type StatisticsItem struct {
	// 日期，格式为年-月-日，如2018-07-13
	StatDate *string `json:"StatDate,omitempty" name:"StatDate"`

	// 统计值
	Data *uint64 `json:"Data,omitempty" name:"Data"`
}

type StreamTextStatisticsItem struct {
	// 统计值，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	Data *float64 `json:"Data,omitempty" name:"Data"`
}

type Tag struct {
	// 标签键
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}

type Task struct {
	// 数据的唯一ID
	DataId *string `json:"DataId,omitempty" name:"DataId"`

	// 数据文件的url，为 urlencode 编码，流式则为拉流地址
	Url *string `json:"Url,omitempty" name:"Url"`

	// gme实时语音房间ID，通过gme实时语音进行语音分析时输入
	RoomId *string `json:"RoomId,omitempty" name:"RoomId"`

	// gme实时语音用户ID，通过gme实时语音进行语音分析时输入
	OpenId *string `json:"OpenId,omitempty" name:"OpenId"`
}

// Predefined struct for user
type UpdateScanRoomsRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要送检的所有房间号。多个房间号之间用","分隔。示例："0001,0002,0003"
	RoomIdString *string `json:"RoomIdString,omitempty" name:"RoomIdString"`

	// 符合此正则表达式规则的房间号将被送检。示例：["^6.*"] 表示所有以6开头的房间号将被送检
	RoomIdRegex []*string `json:"RoomIdRegex,omitempty" name:"RoomIdRegex"`
}

type UpdateScanRoomsRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要送检的所有房间号。多个房间号之间用","分隔。示例："0001,0002,0003"
	RoomIdString *string `json:"RoomIdString,omitempty" name:"RoomIdString"`

	// 符合此正则表达式规则的房间号将被送检。示例：["^6.*"] 表示所有以6开头的房间号将被送检
	RoomIdRegex []*string `json:"RoomIdRegex,omitempty" name:"RoomIdRegex"`
}

func (r *UpdateScanRoomsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanRoomsRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "RoomIdString")
	delete(f, "RoomIdRegex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateScanRoomsRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanRoomsResponseParams struct {
	// 返回结果码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateScanRoomsResponse struct {
	*tchttp.BaseResponse
	Response *UpdateScanRoomsResponseParams `json:"Response"`
}

func (r *UpdateScanRoomsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanRoomsResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanUsersRequestParams struct {
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要送检的所有用户号。多个用户号之间用","分隔。示例："0001,0002,0003"
	UserIdString *string `json:"UserIdString,omitempty" name:"UserIdString"`

	// 符合此正则表达式规则的用户号将被送检。示例：["^6.*"] 表示所有以6开头的用户号将被送检
	UserIdRegex []*string `json:"UserIdRegex,omitempty" name:"UserIdRegex"`
}

type UpdateScanUsersRequest struct {
	*tchttp.BaseRequest
	
	// 应用ID
	BizId *uint64 `json:"BizId,omitempty" name:"BizId"`

	// 需要送检的所有用户号。多个用户号之间用","分隔。示例："0001,0002,0003"
	UserIdString *string `json:"UserIdString,omitempty" name:"UserIdString"`

	// 符合此正则表达式规则的用户号将被送检。示例：["^6.*"] 表示所有以6开头的用户号将被送检
	UserIdRegex []*string `json:"UserIdRegex,omitempty" name:"UserIdRegex"`
}

func (r *UpdateScanUsersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanUsersRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "BizId")
	delete(f, "UserIdString")
	delete(f, "UserIdRegex")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "UpdateScanUsersRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type UpdateScanUsersResponseParams struct {
	// 返回结果码
	ErrorCode *int64 `json:"ErrorCode,omitempty" name:"ErrorCode"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type UpdateScanUsersResponse struct {
	*tchttp.BaseResponse
	Response *UpdateScanUsersResponseParams `json:"Response"`
}

func (r *UpdateScanUsersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *UpdateScanUsersResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type UserMicStatus struct {
	// 用户ID
	Uid *int64 `json:"Uid,omitempty" name:"Uid"`

	// 是否开麦 。1闭麦  2开麦
	EnableMic *int64 `json:"EnableMic,omitempty" name:"EnableMic"`
}

type VoiceFilterConf struct {
	// 语音过滤服务开关，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`
}

type VoiceFilterStatisticsItem struct {
	// 语音过滤总时长，单位为min
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`
}

type VoiceMessageConf struct {
	// 离线语音服务开关，取值：open/close
	Status *string `json:"Status,omitempty" name:"Status"`

	// 离线语音支持语种，取值： all-全部，cnen-中英文。默认为中英文
	Language *string `json:"Language,omitempty" name:"Language"`
}

type VoiceMessageStatisticsItem struct {
	// 离线语音DAU
	Dau *uint64 `json:"Dau,omitempty" name:"Dau"`
}