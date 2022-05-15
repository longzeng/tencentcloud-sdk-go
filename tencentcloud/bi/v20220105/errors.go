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

package v20220105

const (
	// 此产品的特有错误码

	// CAM签名/鉴权错误。
	AUTHFAILURE = "AuthFailure"

	// 操作失败。
	FAILEDOPERATION = "FailedOperation"

	// 内部错误。
	INTERNALERROR = "InternalError"

	// 内部错误。
	INTERNALERROR_INTERNAL = "InternalError.Internal"

	// 参数错误。
	INVALIDPARAMETER = "InvalidParameter"

	// 报表嵌出类错误。
	INVALIDPARAMETER_EMBED = "InvalidParameter.Embed"

	// 参数取值错误。
	INVALIDPARAMETERVALUE = "InvalidParameterValue"

	// 缺少参数错误。
	MISSINGPARAMETER = "MissingParameter"

	// 必填参数缺失。
	MISSINGPARAMETER_MISSINGPARAM = "MissingParameter.MissingParam"

	// 操作被拒绝。
	OPERATIONDENIED = "OperationDenied"

	// 请求的次数超过了频率限制。
	REQUESTLIMITEXCEEDED = "RequestLimitExceeded"

	// 未授权操作。
	UNAUTHORIZEDOPERATION = "UnauthorizedOperation"

	// 权限错误。
	UNAUTHORIZEDOPERATION_AUTHORIZE = "UnauthorizedOperation.Authorize"

	// 未知参数错误。
	UNKNOWNPARAMETER = "UnknownParameter"

	// 操作不支持。
	UNSUPPORTEDOPERATION = "UnsupportedOperation"

	// 默认业务异常。
	UNSUPPORTEDOPERATION_BIERROR = "UnsupportedOperation.BIError"
)
