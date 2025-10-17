# server-agent-proto

这是一个用于服务器和代理节点之间通信的 gRPC 协议定义项目。

## 依赖安装

### 1. 安装 Protocol Buffers 编译器

#### macOS
```bash
brew install protobuf
```

#### Linux
```bash
# Debian/Ubuntu
sudo apt-get install -y protobuf-compiler

# CentOS/RHEL
sudo yum install -y protobuf-compiler
```

#### Windows
从 [Protocol Buffers Releases](https://github.com/protocolbuffers/protobuf/releases) 下载预编译的二进制文件。

### 2. 安装 Go 插件

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

确保 `$GOPATH/bin` 或 `$GOBIN` 在你的 `PATH` 环境变量中：

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## 编译 Proto 文件

在项目根目录执行以下命令：

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       pkg/agent.proto
```

这会生成以下文件：
- `pkg/agent.pb.go` - Protocol Buffers 消息定义的 Go 代码
- `pkg/agent_grpc.pb.go` - gRPC 服务定义的 Go 代码

## 项目结构

```
server-agent-proto/
├── pkg/
│   ├── agent.proto           # Protocol Buffers 定义文件
│   ├── agent.pb.go          # 生成的消息代码
│   └── agent_grpc.pb.go     # 生成的 gRPC 服务代码
├── go.mod
├── go.sum
└── README.md
```

## 使用方法

### 导入包

```go
import (
    pb "github.com/xflash-panda/server-agent-proto/pkg"
)
```

### 服务端示例

```go
type agentServer struct {
    pb.UnimplementedAgentServer
}

func (s *agentServer) Config(ctx context.Context, req *pb.ConfigRequest) (*pb.ConfigResponse, error) {
    // 实现配置逻辑
    return &pb.ConfigResponse{
        Result: true,
        RawData: []byte("config data"),
    }, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    
    s := grpc.NewServer()
    pb.RegisterAgentServer(s, &agentServer{})
    
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

### 客户端示例

```go
func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    
    client := pb.NewAgentClient(conn)
    
    resp, err := client.Config(context.Background(), &pb.ConfigRequest{
        NodeId: 1,
        NodeType: pb.NodeType_VMESS,
    })
    if err != nil {
        log.Fatalf("could not get config: %v", err)
    }
    
    fmt.Printf("Result: %v\n", resp.Result)
}
```

## 支持的节点类型

- `SHADOWSOCKS` - Shadowsocks 节点
- `TROJAN` - Trojan 节点
- `VMESS` - VMess 节点
- `HYSTERIA` - Hysteria 节点
- `HYSTERIA2` - Hysteria2 节点
- `ANYTLS` - AnyTLS 节点

## RPC 服务

- `Config` - 获取节点配置
- `Register` - 注册节点
- `Unregister` - 注销节点
- `Heartbeat` - 发送心跳
- `Submit` - 提交统计数据
- `Users` - 获取用户信息

## License

请查看 LICENSE 文件了解详细信息。
