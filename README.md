# SDK使用文档

## 安装

**项目必须启用 `go module`**

执行以下命令安装 SDK：

```bash
go get github.com/ucloud/ustack-sdk-go
```

或在 `go.mod` 中增加依赖：

```
require github.com/ucloud/ustack-sdk-go v0.0.0
```

## 使用

登陆控制台后获取公私钥，替换到代码中。

![key.png](key.png)

```go
package main

import (
	"fmt"

	"github.com/ucloud/ustack-sdk-go/common"
	"github.com/ucloud/ustack-sdk-go/common/auth"
	"github.com/ucloud/ustack-sdk-go/services/openapi"
)

func main() {
	cfg := common.NewConfig()

	credential := auth.NewCredential()
	// 替换成平台上获取的公/私钥
	credential.PrivateKey = "my_private_key"
	credential.PublicKey = "my_public_key"

	client := openapi.NewClient(&cfg, &credential)

	req := client.NewCreateVMInstanceRequest()
	req.Region = common.String("my_region") // 替换成平台上的目标地域
	req.Name = common.String("sdk-example-vm")
	req.ImageID = common.String("image-xxx") // 替换成平台上可用的镜像ID
	req.Password = common.String("my_vm_password")
	req.ChargeType = common.String("Dynamic")
	req.CPU = common.Int(1)
	req.Memory = common.Int(1024)

	// 发送请求
	vm, err := client.CreateVMInstance(req)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Printf("resource id of the vm: %v\n", vm.VMID)
	}
}
```

## 类型系统

由于 Golang 不区分 零值 和 空值，所以 Go SDK 的简单类型采用包装类型的方式来设置参数：

```go
req := client.NewCreateVMInstanceRequest()

req.Region = common.String("my_region")
req.Name = common.String("sdk-example-vm")
req.ImageID = common.String("image-xxx")
req.Password = common.String("my_vm_password")
req.ChargeType = common.String("Dynamic")
req.CPU = common.Int(1)
req.Memory = common.Int(1024)
```

数组和结构体等类型与 Go 原生类型系统保持一致。
