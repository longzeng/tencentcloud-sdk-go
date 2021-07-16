<p align="center">
<a href="https://cloud.tencent.com/"><img src="https://imgcache.qq.com/qcloud/tcloud_dtc/static/static_source_business/eec00e38-a178-479f-83d4-853a18575ac4.png" height="100" ></a>
</p>
<h1 align="center">Tencent Cloud SDK for Go</h1>

# 简介

欢迎使用腾讯云开发者工具套件（SDK），此 SDK 是云 API 3.0 平台的配套开发工具。

# 依赖环境

1. Go 1.9 版本及以上（如使用 go mod 需要 Go 1.14）。
2. 部分产品需要在腾讯云控制台开通后，才能正常调用此产品的接口。
3. 在腾讯云控制台 [访问管理](https://console.cloud.tencent.com/cam/capi) 页面获取密钥 SecretID 和 SecretKey，请务必妥善保管，或者使用更安全的临时安全凭证。

# 获取安装

## 通过go get安装（推荐）

推荐使用腾讯云镜像加速下载：

1. lnuix 或 macos:

    ```bash
    export GOPROXY=https://mirrors.tencent.com/go/
    ```

2. windows:

    ```cmd
    set GOPROXY=https://mirrors.tencent.com/go/
    ```

### 按需安装（推荐）

v1.0.170后可以按照产品下载，您只需下载基础包和对应的产品包(如cvm)即可，不需要下载全部的产品，从而加快您构建镜像或者编译的速度：

1. 安装公共基础包：

    ```bash
    go get -v github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common@latest
    ```

2. 安装对应的产品包(如cvm):

    ```bash
    go get -v github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm@latest
    ```

### 全部安装

您也可以按照以前的方式一次性下载腾讯云所有产品的包：

```bash
go get -v github.com/tencentcloud/tencentcloud-sdk-go@latest
```

注意：为了支持 go mod，SDK 版本号从 v3.x 降到了 v1.x。并于2021.05.10移除了所有`v3.0.*`和`3.0.*`的tag，如需追溯以前的tag，请参考项目根目录下的 `commit2tag` 文件。

## 通过源码安装

前往代码托管地址 [Github](https://github.com/tencentcloud/tencentcloud-sdk-go) 或者 [Gitee](https://gitee.com/tencentcloud/tencentcloud-sdk-go) 下载最新代码，解压后安装到 $GOPATH/src/github.com/tencentcloud 目录下。

# 快速开始

每个接口都有一个对应的 Request 结构和一个 Response 结构。例如云服务器的查询实例列表接口 DescribeInstances 有对应的请求结构体 DescribeInstancesRequest 和返回结构体 DescribeInstancesResponse 。

下面以云服务器查询实例列表接口为例，介绍 SDK 的基础用法。

## 简化版

```go
package main

import (
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
	credential := common.NewCredential("secretId", "secretKey")
	client, _ := cvm.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())

	request := cvm.NewDescribeInstancesRequest()
	response, err := client.DescribeInstances(request)

	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", response.ToJsonString())
}
```

## 详细版

出于演示的目的，有一些非必要的代码，例如对默认配置的修改，以尽量展示 SDK 的功能。在实际编写代码使用 SDK 的时候，建议尽量使用默认配置，酌情修改。

```go
package main

import (
        "fmt"

        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
        "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
        cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
        // 必要步骤：
        // 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
        // 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
        // 以免泄露密钥对危及你的财产安全。
        credential := common.NewCredential("secretId", "secretKey")

        // 非必要步骤
        // 实例化一个客户端配置对象，可以指定超时时间等配置
        cpf := profile.NewClientProfile()
        // SDK默认使用POST方法。
        // 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
        // 如非必要请不要修改默认设置。
        cpf.HttpProfile.ReqMethod = "POST"
        // SDK有默认的超时时间，如非必要请不要修改默认设置。
        // 如有需要请在代码中查阅以获取最新的默认值。
        cpf.HttpProfile.ReqTimeout = 30
        // SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
        // 则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com
        cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
        // SDK默认用TC3-HMAC-SHA256进行签名，它更安全但是会轻微降低性能。
        // 如非必要请不要修改默认设置。
        cpf.SignMethod = "TC3-HMAC-SHA256"
        // SDK 默认用 zh-CN 调用返回中文。此外还可以设置 en-US 返回全英文。
        // 但大部分产品或接口并不支持全英文的返回。
        // 如非必要请不要修改默认设置。
        cpf.Language = "en-US"
        //打印日志，默认是false
        // cpf.Debug = true


        // 实例化要请求产品(以cvm为例)的client对象
        // 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
        client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
        // 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
        // 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
        // 属性可能是基本类型，也可能引用了另一个数据结构。
        // 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
        request := cvm.NewDescribeInstancesRequest()

        // 基本类型的设置。
        // 此接口允许设置返回的实例数量。此处指定为只返回一个。
        // SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
        // SDK提供对基本类型的指针引用封装函数
        request.Limit = common.Int64Ptr(1)

        // 数组类型的设置。
        // 此接口允许指定实例 ID 进行过滤，但是由于和接下来要演示的 Filter 参数冲突，先注释掉。
        // request.InstanceIds = common.StringPtrs([]string{"ins-r8hr2upy"})

        // 复杂对象的设置。
        // 在这个接口中，Filters是数组，数组的元素是复杂对象Filter，Filter的成员Values是string数组。
        request.Filters = []*cvm.Filter{
            &cvm.Filter{
                Name: common.StringPtr("zone"),
                Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
            },
        }

        // 通过client对象调用想要访问的接口，需要传入请求对象
        response, err := client.DescribeInstances(request)
        // 处理异常
        if _, ok := err.(*errors.TencentCloudSDKError); ok {
            fmt.Printf("An API error has returned: %s", err)
            return
        }
        // 非SDK异常，直接失败。实际代码中可以加入其他的处理。
        if err != nil {
            panic(err)
        }
        // 打印返回的json字符串
        fmt.Printf("%s\n", response.ToJsonString())
}
```

更多示例参见 [examples](https://github.com/TencentCloud/tencentcloud-sdk-go/tree/master/examples) 目录。对于复杂接口的 Request 初始化例子，可以参考 [例一](examples/cvm/v20170312/run_instances.go) 。对于使用json字符串初始化 Request 的例子，可以参考 [例二](examples/cvm/v20170312/describe_instances.go) 。

# 相关配置

**如无特殊需要，建议您使用默认配置。**

在创建客户端前，如有需要，您可以通过修改`profile.ClientProfile`中字段的值进行一些配置。

```go
// 非必要步骤
// 实例化一个客户端配置对象，可以指定超时时间等配置
cpf := profile.NewClientProfile()
```

具体的配置项说明如下：

## 请求方式

SDK默认使用POST方法。 如果你一定要使用GET方法，可以在这里设置。**GET方法无法处理一些较大的请求**。

```go
cpf.HttpProfile.ReqMethod = "POST"
```

## 超时时间

SDK有默认的超时时间，如非必要请不要修改默认设置。
如有需要请在代码中查阅以获取最新的默认值。  
单位：秒

```go
cpf.HttpProfile.ReqTimeout = 30
```

## 指定域名

SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com

```go
cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
```

## 签名方式

SDK默认用 `TC3-HMAC-SHA256` 进行签名，它更安全但是会轻微降低性能。

```go
cpf.SignMethod = "HmacSHA1"
```

## DEBUG模式

DEBUG模式会打印更详细的日志，当您需要进行详细的排查错误时可以开启。  
默认为 `false`

```go
cpf.Debug = true
```

## 代理

如果是有代理的环境下，需要设置系统环境变量 `https_proxy` ，否则可能无法正常调用，抛出连接超时的异常。或者自定义 Transport 指定代理，通过 client.WithHttpTransport 覆盖默认配置。

## 开启 DNS 缓存

当前 GO SDK 总是会去请求 DNS 服务器，而没有使用到 nscd 的缓存，可以通过导出环境变量`GODEBUG=netdns=cgo`，或者`go build`编译时指定参数`-tags 'netcgo'`控制读取 nscd 缓存。

## 忽略服务器证书校验

虽然使用 SDK 调用公有云服务时，必须校验服务器证书，以识破他人伪装的服务器，确保请求的安全。
但是某些极端情况下，例如测试时，你可能会需要忽略自签名的服务器证书。
以下是其中一种可能的方法：

```
import "crypto/tls"
...
    client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client.WithHttpTransport(tr)
...
```

**再次强调，除非你知道自己在做什么，并明白由此带来的风险，否则不要尝试关闭服务器证书校验。**

# 错误处理

从 `v1.0.181` 开始，腾讯云 GO SDK 会将各个产品的返回的错误码定义为常量，您可以直接调用处理，无需手动定义。如果您使用 IDE (如 Goland )进行开发，可以使用他们的代码提示功能直接选择。例如：

```go
...//Your other go code

// Handling errors
response, err := client.DescribeInstances(request)
if terr, ok := err.(*errors.TencentCloudSDKError); ok {
    code := terr.GetCode()
    if code == cvm.FAILEDOPERATION_ILLEGALTAGKEY{
        fmt.Printf("Handling error: FailedOperation.IllegalTagKey,%s", err)
    }else if code == cvm.UNAUTHORIZEDOPERATION{
        fmt.Printf("Handling error: UnauthorizedOperation,%s", err)
    }else{
        fmt.Printf("An API error has returned: %s", err)
    }
    return
}
...
```

同时，各个接口函数的注释部分也列出了此接口可能会返回的错误码，方便您进行处理：

```go
// DescribeInstances
// 本接口 (DescribeInstances) 用于查询一个或多个实例的详细信息。
//
// 
//
// * 可以根据实例`ID`、实例名称或者实例计费模式等信息来查询实例的详细信息。过滤信息详细请见过滤器`Filter`。
//
// * 如果参数为空，返回当前用户一定数量（`Limit`所指定的数量，默认为20）的实例。
//
// * 支持查询实例的最新操作（LatestOperation）以及最新操作状态(LatestOperationState)。
//
// 可能返回的错误码:
//  FAILEDOPERATION_ILLEGALTAGKEY = "FailedOperation.IllegalTagKey"
//  FAILEDOPERATION_ILLEGALTAGVALUE = "FailedOperation.IllegalTagValue"
//  FAILEDOPERATION_TAGKEYRESERVED = "FailedOperation.TagKeyReserved"
//  INTERNALSERVERERROR = "InternalServerError"
//  INVALIDFILTER = "InvalidFilter"
//  INVALIDFILTERVALUE_LIMITEXCEEDED = "InvalidFilterValue.LimitExceeded"
//  INVALIDHOSTID_MALFORMED = "InvalidHostId.Malformed"
//  INVALIDINSTANCEID_MALFORMED = "InvalidInstanceId.Malformed"
//  INVALIDPARAMETER = "InvalidParameter"
//  INVALIDPARAMETERVALUE = "InvalidParameterValue"
//  INVALIDPARAMETERVALUE_IPADDRESSMALFORMED = "InvalidParameterValue.IPAddressMalformed"
//  INVALIDPARAMETERVALUE_INVALIDIPFORMAT = "InvalidParameterValue.InvalidIpFormat"
//  INVALIDPARAMETERVALUE_INVALIDVAGUENAME = "InvalidParameterValue.InvalidVagueName"
//  INVALIDPARAMETERVALUE_LIMITEXCEEDED = "InvalidParameterValue.LimitExceeded"
//  INVALIDPARAMETERVALUE_SUBNETIDMALFORMED = "InvalidParameterValue.SubnetIdMalformed"
//  INVALIDPARAMETERVALUE_TAGKEYNOTFOUND = "InvalidParameterValue.TagKeyNotFound"
//  INVALIDPARAMETERVALUE_VPCIDMALFORMED = "InvalidParameterValue.VpcIdMalformed"
//  INVALIDSECURITYGROUPID_NOTFOUND = "InvalidSecurityGroupId.NotFound"
//  INVALIDSGID_MALFORMED = "InvalidSgId.Malformed"
//  INVALIDZONE_MISMATCHREGION = "InvalidZone.MismatchRegion"
//  RESOURCENOTFOUND_HPCCLUSTER = "ResourceNotFound.HpcCluster"
//  UNAUTHORIZEDOPERATION_INVALIDTOKEN = "UnauthorizedOperation.InvalidToken"
func (c *Client) DescribeInstances(request *DescribeInstancesRequest) (response *DescribeInstancesResponse, err error){
    ...
}

```

# Common Request

从 `v1.0.189`开始，腾讯云 GO SDK 支持使用 `泛用型的API调用方式(Common Request)` 进行请求。您只需安装 `common` 包, 即可向任何产品发起调用。

**注意，您必须明确知道您调用的接口所需参数，否则可能会调用失败。**

目前仅支持使用POST方式，且签名方法必须使用 签名方法 v3。

详细使用请参阅示例：[使用 Common Request 进行调用](https://github.com/TencentCloud/tencentcloud-sdk-go/blob/master/examples/common/common_request.go)

# 支持产品列表

参见[产品列表文档](./products.md)
