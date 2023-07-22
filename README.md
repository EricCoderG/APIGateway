# 项目文档
> 团队名称：放我回家队
> 
> 团队成员：耿瑞林、毛丁、刘博悦

<!-- TOC -->
* [项目文档](#项目文档)
  * [1.目的](#1目的)
    * [1.1.用户故事](#11用户故事)
  * [2.系统目标](#2系统目标)
    * [2.1.项目概述](#21项目概述)
      * [2.1.1.项目的问题和需求：](#211项目的问题和需求)
      * [2.1.2.我们的解决方案：](#212我们的解决方案)
    * [2.2.示意图](#22示意图)
    * [2.3.技术栈](#23技术栈)
    * [2.4.实现的功能](#24实现的功能)
    * [2.5.下一阶段的特性](#25下一阶段的特性)
    * [2.6.如何使用/评估我们的功能](#26如何使用评估我们的功能)
      * [创建一个新的 length 微服务，并将其集成到 API 网关上。](#创建一个新的-length-微服务并将其集成到-api-网关上)
    * [2.7.运行服务器](#27运行服务器)
    * [2.8.测试](#28测试)
    * [2.9.结论](#29结论)
    <!-- TOC -->
## 1.目的
API 网关作为系统的入口点，接收所有传入的 HTTP 请求，以适当的格式将它们转发到相关服务器，检索响应，然后将其返回给客户端。特别是网关为路由请求、 JSON 和 Thrift 格式之间的转换、服务发现以及与后端RPC服务器的通信提供了可扩展且可靠的解决方案，从而成为进入潜在的大型分布式系统的有效入口。
<br>
<br>
在我们的日常生活中，我们直接或间接通过互联网与服务进行交互。然而，随着服务变得越来越复杂和健壮，它们经常被构建为微服务的分布式系统，每个微服务执行一个更小的特定任务。我们不能期望客户端直接与处理其特定请求的微服务进行通信，因为这将需要维护一个广泛的地址库，并且在每次稍微更改请求后都要重构请求。网关是针对这些问题的有效解决方案，它确保客户端只有一个联系点来发送请求。网关处理这些请求并向客户端返回适当的响应，从客户端的角度简化它。

### 1.1.用户故事
**目标受众**：其他开发人员(使用我们提供的微服务来支持他们的程序)<br><br>
1.作为一个客户端开发人员，我想要一个单一的接触点的任何我在一个服务中需要的数据。<br><br>
2.作为一个客户端开发人员，我希望服务维护一个它可以处理的所有可能请求的列表，而不管它在后端如何处理它。<br><br>
3.作为一个客户端开发人员，我希望 API网关具有高可用性和容错性，这样我就可以依赖其一致的性能和可用性，而不会经历频繁的停机或中断。<br><br>
4.作为一个系统开发人员，我希望能够允许微服务注册自己并被动态发现，这样我就可以轻松地添加或删除它们，而无需手动配置。<br><br>

## 2.系统目标
可伸缩性:一个网关应该允许在对代码进行最小更改的情况下轻松合并额外的微服务，允许服务轻松注册自己并支持动态发现(请求不应该需要知道服务的地址)<br><br>
准确性:请求应该由预期的服务响应，并且响应应该是客户端所期望的。<br><br>
可用性:无论客户端请求数多少，都能保证系统正常、快速地运行，尽量减少任何停机时间。

### 2.1.项目概述
#### 2.1.1.项目的问题和需求：
a.在我们的日常生活中，我们直接或间接的通过网络与服务互动。由于这些服务变得越来越复杂，开发人员经常选择将它们构建为微服务的分布式系统，运行在不同的服务器上，每个微服务执行较小的特定任务。<br><br>
b.不幸的是，开发人员不能合理地期望客户端直接与能够处理其特定请求的微服务服务器进行通信：<br><br>
&emsp; i.这将要求客户端维护所有服务器地址的广泛库，并且它将使客户端有责任定期更新该库。<br><br>
&emsp; ii.客户端还必须每次修改请求以满足每个服务器的个别需求。<br><br>
c.上述这两种情况对客户端来说都是非常麻烦的，并且可能会使他们放弃使用整个服务。因此，需要一个集中的单一入口点来进入整个服务，作为一个“网关”，解释来自客户端的请求并将其重定向到适当的微服务服务器。<br><br>
#### 2.1.2.我们的解决方案：
a.我们的目标是建立一个 API网关，如前所述，作为分布式系统的入口点，使开发人员能够构建灵活和可扩展的基于微服务的系统。<br><br>
b.为了构建这样的系统，开发人员需要 API 网关基础设施，可以:<br><br>
&emsp; i.允许他们轻松地构建、添加或删除微服务<br><br>
&emsp; ii.要求对任何集中的代码片段进行最小的更改<br><br>
&emsp; iii.集成一套支持功能以改进微服务集成，包括但不限于:<br><br>
&emsp;&emsp;1.**服务注册/发现**:允许在集中注册中心通过服务名称注册/发现服务。<br><br>
&emsp;&emsp;2.**通用调用**:将请求转换为通用格式进行传输，这意味着我们可以在 API 网关和微服务之间进行通信，而无需了解微
服务实现的许多具体细节。<br><br>
&emsp; iv.这些特性(在本文档后面详细阐述)将由我们的 API 网关实现，以构建一个有效的解决方案，轻
松构建健壮的，可扩展的，灵活的基于微服务的系统。

### 2.2.示意图
![img_1](https://ericcoder-oss.oss-cn-hangzhou.aliyuncs.com/markdown_images/img_1.png)

### 2.3.技术栈
+ **Hertz**:一个高可用性、高性能、高可扩展性的 HTTP 框架，支持微服务的开发。Hertz 提供了允许 API 网关解释和响应来自
客户端的 HTTP 请求的基础设施。<br><br>
+ **Kitex**:一个高性能和强可扩展性的 Golang RPC框架，支持构建微服务。Kitex 为网关和提供服务的服务器之间的 RPC通信提供
了许多基础设施和特性。在我们的项目中，它允许我们创建多个微服务进行测试。<br><br>
+ **etcd**:它是一个开源的、分布式的、一致的键值存储，用于共享配置、服务发现和分布式系统或机器集群的调度器协调。
在我们的项目中，我们使用它来注册多个基于 kitex 的服务，并使它们能够被 API 网关发现。<br><br>
+ **ApacheThrift**:它是一种接口定义语言和二进制通信协议，用于为众多编程语言定义和创建服务。在这个项目中，我们创建了
thrift 文件，这些文件可以自动为 Hertz 和 Kitex 服务器生成脚手架代码，并将 RPC请求转换为 thrift 二进制格式。<br><br>
+ **Go**: Go 是一种强大而高效的编程语言，它优先考虑简单性和并发性。我们选择 Go 作为项目的基础，利用它的优势来构建我们
  的 API 网关。此外，我们还利用了专门为 Go 设计的 Hertz 和 Kitex 框架。

### 2.4.实现的功能
+ **接受 HTTP 请求** :我们已经实现了在 API 网关中处理 HTTP 请求所需的逻辑。网关能够处理和响应各种 HTTP 方法(如
GET 和 POST)，并从请求头、参数和主体中提取相关数据。通过接受 HTTP 请求，我们启用了与 API 网关的通信，促
进了信息交换以及与底层服务的交互。<br><br>
+ **能够创建一个新的 Kitex 服务器来处理特定的微服务** :我们已经实现了接受和管理由 Kitex 生成的多个 RPC服务器所需的逻辑。
这个能力使我们能够有效地路由和处理传入针对目标微服务的请求。我们还利用 Kitex 的功能进行通用调用和服务发现。<br><br>
+ **服务注册和发现** :通过使用 etcd实现服务注册和发现特性，我们的系统获得了注册多个 RPC服务器并使它们可被发现的能力。这个功能增强
了我们的微服务架构的可扩展性和可靠性。通过服务注册，我们的服务器可以动态地宣布它们的可用性，使其他组件能够轻
松地定位并与它们交互。 <br><br>
+ **泛型调用**:Kitex 泛型调用函数提供了一种灵活的方式来调用微服务架构中的远程过程。使用泛型调用函数，我们可以向服务
发出 RPC请求，而无需显式地知道底层服务接口或消息类型。此功能允许服务之间的松耦合，并支持动态的服务发现和组
合。
+ **client池化**：在项目开发中，经常会向其它的服务发送一些Http请求，获取一些数据或验证。
如果每次都重新创建一个新的HttpClient对象的话，当并发上来时，容易出现异常或连接失败，超时。
这里可以使用sync.Map的作为Client的全局pool，减少HttpClient创建的数量，使得client可以被服用，提高框架处理效率。

### 2.5.下一阶段的特性
1. 易于访问:使添加新服务变得更容易
   - 用户:开发人员
   - 目标:使开发人员更容易使用网关，只进行有限的更改来满足他们的需求
   - 好处:使 API 网关更加用户友好 
   - 建议实现:我们打算将更多的过程抽象到 utils文件中，例如提供功能，以改进 kitex客户端的生成并在 etcd上注册。
2. 重试和超时:配置适当的重试和超时机制来处理临时故障或慢速后端 RPC服务器。
   - 用户:开发人员
   - 目标:处理不可预测的情况并保持响应能力，提高整体系统弹性和用户体验。
   - 好处:通过自动重试失败的请求和防止请求无限期地滞留，提高了系统的健壮性。这增强了可靠性，减轻了临时故障或慢
   速后端 RPC服务器，并确保了响应式用户体验。超时意味着开发人员有一个请求将花费的最大时间的指示，并且可以根
   据这个指示构建他们的程序。它确保了鲁棒性并改善了错误处理。
   - 建议实现:我们打算使用 kitex 请求超时特性，并将其整合到 utils生成客户端代码中，只要求开发人员指出超时持续时间。
3. 请求验证:验证 API 网关中传入的 JSON 请求，以确保它们符合预期的格式。
   - 用户:开发人员
   - 目标:确保 API 网关中传入的 JSON 请求遵循预期的格式。通过验证请求，我们可以确保数据的完整性，防止畸形或恶
   意请求被处理，并提高整体系统的安全性和可靠性。
   - 好处:保证数据完整性、增强安全性、改进错误处理、促进兼容性、确保合规性。促进数据格式良好，满足法规要
     求，使系统安全、可靠、合规性强。
   - 建议实施:创建一个新文件，开发人员可以在其中定义他们的具体标准，
   这之后我们可以将其集成到 utils上的新功能或集成到现有的功能中。
4. 速率限制:在 API网关中实现速率限制，以保护后端服务免受过多流量的影响
   - 用户:开发人员
   - 目标:管理和控制传入请求流，实施使用配额并保护后端服务免受滥用或恶意攻击。
   - 好处:保护后端服务免受流量过大的影响，确保其稳定性、安全性和最佳性能。对请求进行限制，防止滥用，有效分配资
     源，在保证可靠性和可用性的同时，防范安全威胁。
   - 建议实现:我们还是不太确定，但一个粗略的计划是使用 go 的`rate`包，并将其形成 utils包上的一个函数。这将作为
     API 网关整体基础设施的一部分集成，开发人员应该有机会修改特定的标准，但不能修改整体实现。
5. 负载均衡：将负载合理地分配给多个资源提供方
    - 用户：开发人员
    - 目标：避免过载。
    - 好处：提高了系统的可靠性，优化了资源的使用，达到最大吞吐率和最小响应时间。
    - 建议实现：用加权轮询、哈希等算法实现负载均衡。

### 2.6.如何使用/评估我们的功能
由于这是一个针对开发人员的后端项目，因此没有前端 UI，但是，我们在下面提供了文档，可以帮助开发人员了解如何通过代码
编辑器和命令行创建 RPC服务器并利用我们的 API 网关及其功能。我们相信这是了解项目用例及其功能的最佳方式，同时也可以
对其进行测试。
#### 创建一个新的 length 微服务，并将其集成到 API 网关上。
1. 我们的目录结构(主分支):
   - `hertz-http-server`文件夹包含接受HTTP请求的代码和API的主要业务代码
   网关的实现。
   - `utils`文件夹中包含文件`utils.go`，它提供了许多有用的函数来帮助开发人员使用我们的
   API网关。
   -  `idl`文件夹中包含thrift接口定义语言文件。这些文件可以自动为 Hertz 和 Kitex 服务器
   生成脚手架代码。
   -  `reverse-service`、`substring-service`、`conversion-service`文件夹包含了
   字符串翻转、返回子串位置、字符串大写转换的RPC服务器的代码，也就是这些算法的微服务。
   这些不是网关本身的一部分，而更像是关于我们可以如何使用本项目的例子。
   现在我们将实现第四个服务-`length-service`，来展示使用我们的API网关来构建一个新的微服务，
   并添加它以将其与现有的微服务基础设施集成，是一件多么容易的事。
   - 在执行下面的步骤之前，请将我们的存储库克隆到您的本地计算机上，并确保您位于根目录下。
2. 创建一个提供求字符串长度功能的新 RPC服务器(即len微服务)
    - `cd`到`idl`文件夹下
    - 创建一个名为`length.thrift`的新文件
    - 将以下代码输入文件：
```idl
namespace go length

struct LengthRequest {
1: required string inputString;
}

struct LengthResponse {
1: required i32 length;
}

service LengthService {
LengthResponse calculateLength(1: LengthRequest req) (api.post="/length");
}
```
2. 续
    - 上面的文件将被Kitex框架用来为你的RPC服务器生成脚手架代码。生成的代码为
    存储在子目录kitex_gen中。
    - 然而，在此之前，您需要创建一个名为length-service的新目录，其中将包含此服务的Kitex代码。
    - 然后执行以下命令，让Kitex使用你新的thrift IDL文件生成代码:
```shell
kitex -module github.com/ararchch/api-gateway -service Length ../idl/length.thrift
```
2. 续
    - 完成后，打开你的`handler.go`文件并添加以下代码以实现length服务的逻辑。
   （一定要编辑正确的函数并导入任何必要的包。）
```go
package main

import (
	length "api-gateway/length-service/kitex_gen/length"
	"context"
)

// LengthServiceImpl implements the last service interface defined in the IDL.
type LengthServiceImpl struct{}

// CalculateLength implements the LengthServiceImpl interface.
func (s *LengthServiceImpl) CalculateLength(ctx context.Context, req *length.LengthRequest) (resp *length.LengthResponse, err error) {
	resp = &length.LengthResponse{
		Length: int32(len(req.InputString)),
	}
	return
}
```
2. 续
    - 打开`length-service`文件夹中的`main.go`文件并添加以下代码。这段代码启动了一个新的etcd注册表
      在端口2379上，然后继续将服务注册到该端口。所有服务的代码都非常相似，
      但是您每次都需要更改ServiceName字段，并且在您想要决定你希望服务运行在哪个端口上时，
      需要在`api.NewServer()`方法中添加以下属性:
      server.WithServiceAddr(&net.TCPAddr{Port: XXXX})。XXXX是你希望它运行的端口号(当你希望
      同时运行多个微服务这是至关重要的，因为您不能在一个端口上运行多个服务)。
```go
package main

import (
	api "api-gateway/length-service/kitex_gen/length/lengthservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
	"net"
)

func main() {

	// init etcd
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	svr := api.NewServer(
		new(LengthServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "LengthService"}),
		server.WithRegistry(r), // register service on etcd registry
		server.WithServiceAddr(&net.TCPAddr{Port: 8889}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```
2. 续
    - 现在您已经完成了 RPC服务器的设置!

3. 下一步是对http服务器(即API网关)进行轻微修改，以确保您的服务集成到它上。
    - cd出`length-service`文件夹，再cd到`idl`文件夹中。
    - 打开`gateway.thrift`文件，并添加以下代码：
```go
include "length.thrift"
service GatewayService {
length.LengthResponse calculateLength(1: length.LengthRequest req) (api.post = "/length");
}
```
3. 续
    - cd到`hertz-http-request`文件夹。
    - 运行命令`hz update -idl ../idl/gateway.thrift`
    - 此命令将更新 API 网关，以便当任何请求被发送到 url/length 地址时，
      calculateLength函数将被执行。注意:这里的 calculateLength 与 kitexRPC服务器中的 calculateLength 不一样。
      这只与 hertz http服务器相关，我们将在这里实现调用 RPC服务器上的 calculateLength 方法的业务逻辑。
      在接下来的步骤中，这将变得更加清晰。
    - cd到`./biz/handler/gateway，并打开`gateway_service.go`。
    - 在`gateway_service.go`文件中，在底部您将看到未实现的 calculateLength 方法。
     这和之前定义在`length-service`中的方法是不一样的。这个方法定义了当请求被发送到
      `/length`地址时的处理。请将如下 calculateLength 的方法实现放入`gateway_service.go`中：
```go
// Code generated by hertz generator.

package gateway

import (
	convertService "api-gateway/conversion-service/kitex_gen/conversion"
	"api-gateway/global"
	"api-gateway/hertz-http-server/biz/model/conversion"
	"api-gateway/hertz-http-server/biz/model/length"
	"api-gateway/hertz-http-server/biz/model/reverse"
	"api-gateway/hertz-http-server/biz/model/substring/api"
	lengthService "api-gateway/length-service/kitex_gen/length"
	reverseService "api-gateway/reverse-service/kitex_gen/reverse"
	substringService "api-gateway/substring-service/kitex_gen/substring"
	"api-gateway/utils"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// CalculateLength .
// @router /length [POST]
func CalculateLength(ctx context.Context, c *app.RequestContext) {
	var err error
	var req length.LengthRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 将 req params 绑定到 RPC 请求结构（遵循 RPC 服务 IDL 中声明的请求格式
	reqRpc := &lengthService.LengthRequest{
		InputString: req.InputString,
	}

	var respRpc lengthService.LengthResponse
	err = utils.MakeRpcRequest(ctx, global.LengthCli, "calculateLength", reqRpc, &respRpc)
	if err != nil {
		panic(err)
	}

	resp := &length.LengthResponse{
		Length: respRpc.Length,
	}

	c.JSON(consts.StatusOK, resp)
}
```
3. 续
   - 该方法使用`Utils.GenerateClient()`函数创建一个具有服务注册表的新客户端。
   客户端的名称作为参数传递。
   - 然后该方法创建一个reqRpc和respRpc格式的结构体，将用于与length-service通信
   RPC服务器。然后它使用utils.MakeRpcRequest()方法调用length-service的` calculateLength `方法。
   这就是字符串求长的实际作用。http服务器的DivideNumbers方法只处理http请求。
   - utils.MakeRpcRequest()使用指针并直接更新respRpc结构体。然后解包这些值
   重新打包成HTTP请求的返回格式，返回给客户端。
   - 请注意，将http请求解包为rpc请求的原因，以及将响应解包为rpc请求的原因是
   HTTP和RPC请求/响应可能具有不同的格式。然而，在我们的例子中，它们是相同的。
   - 现在你完成了API网关的编辑!
   
### 2.7.运行服务器
#### 启动服务

要启动服务，你可以运行以下的脚本命令。请确保在执行这些命令之前你已经位于项目的根目录。

##### 启动 Hertz HTTP 服务

``` bash
sh sh/build_and_run.sh hertz-http-server
```

##### 启动 Length 服务

``` bash
sh sh/build_and_run.sh length-service
```

##### 启动 Substring 服务

``` bash
sh sh/build_and_run.sh substring-service
```

##### 启动 Reverse 服务

``` bash
sh sh/build_and_run.sh reverse-service
```

##### 启动 Conversion 服务

``` bash
sh sh/build_and_run.sh conversion-service
```

##### 清理

要清理所有的可执行文件，你可以运行以下脚本命令：

``` bash
sh sh/clean.sh
```

#### 运行测试

##### Length 服务

```shell
go test -bench=. test/length_test.go
```

```shell
ab -n 10000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/length.data "http://127.0.0.1:8088/length"
```

##### Conversion服务

```shell
go test -bench=. test/conversion_test.go
```

```shell
ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/conversion.data "http://127.0.0.1:8088/convert"
```

##### Substring服务

```shell
go test -bench=. test/substring_test.go
```

```shell
ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/substring.data "http://127.0.0.1:8088/substring"
```

##### Reverse服务

```shell
go test -bench=. test/reverse_test.go
```

```shell
ab -n 1000 -c 10 -T 'application/x-www-form-urlencoded' -p test/data/reverse.data "http://127.0.0.1:8088/reverse"
```



### 2.8.测试
1. 适用性:自我评价
   - 由于我们仍在构建项目，我们只能自我评估可用性和适用性，我们确实相信它是相当有效的:开发人员可以在 30 分钟内用
   非常有限的代码行设置和运行一个基本的 RPC服务器。虽然我们已经确定了其他可以进一步精简的领域，我们将在未来
   的阶段中这样做。
2. 单元测试
   - 我们通过Golang性能测试和`Apache Benchmark`进行了压力测试，并优化了测试。
### 2.9.结论
到目前为止，我们已经构建了一个最小可行的产品，实现了 API网关的一些功能，并表明它如何为开发人员所用。
在之后的阶段中，我们打算实现2.5中设想的功能，并更多地关注测试方面。

