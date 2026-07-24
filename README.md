# UCloudStack SDK Go

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.25-blue.svg)](https://go.dev/)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

UCloudStack SDK Go 是 UCloudStack API 的 Go 客户端库。

- 网站: https://www.ucloudstack.com
- 许可证: Apache 2.0

## 安装

项目必须启用 Go Module：

```bash
go get github.com/ucloud/ustack-sdk-go
```

或在 `go.mod` 中添加：

```
require github.com/ucloud/ustack-sdk-go latest
```

## 快速开始

登陆控制台后获取公私钥，替换到代码中：

```go
package main

import (
	"fmt"

	"github.com/ucloud/ustack-sdk-go/common"
	"github.com/ucloud/ustack-sdk-go/common/auth"
	"github.com/ucloud/ustack-sdk-go/services/ucloudstack"
)

func main() {
	cfg := common.NewConfig()

	credential := auth.NewCredential()
	// 替换成平台上获取的公/私钥
	credential.PublicKey = "my_public_key"
	credential.PrivateKey = "my_private_key"

	client := ucloudstack.NewClient(&cfg, &credential)

	req := client.NewCreateVMInstanceRequest()
	req.Region = common.String("my_region")     // 替换成平台上的目标地域
	req.Name = common.String("sdk-example-vm")
	req.ImageID = common.String("image-xxx")     // 替换成平台上可用的镜像ID
	req.Password = common.String("my_vm_password")
	req.ChargeType = common.String("Dynamic")
	req.CPU = common.Int(1)
	req.Memory = common.Int(1024)

	vm, err := client.CreateVMInstance(req)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("resource id of the vm: %v\n", vm.VMID)
	}
}
```

## 类型系统

由于 Go 不区分零值和空值，SDK 采用包装类型设置参数：

```go
req.Region = common.String("my_region")
req.CPU = common.Int(1)
req.Memory = common.Int(1024)
```

支持的包装类型：`common.String`、`common.Int`、`common.Int64`、`common.Float64`、`common.Bool`。

数组和结构体等类型与 Go 原生类型系统保持一致。
