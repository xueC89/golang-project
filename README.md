# Go 后端服务

这是一个为前端项目提供服务的Go后端API服务。

## 功能特性

- ✅ 基本HTTP服务器
- ✅ RESTful API接口
- ✅ CORS跨域支持
- ✅ 健康检查端点
- ✅ 用户管理API
- ✅ 数据查询API
- ✅ 错误处理机制

## 技术栈

- Go 1.21
- 标准库 `net/http`

## 快速开始

### 环境要求

- Go 1.21 或更高版本

### 运行项目

1. 克隆或下载项目到本地

2. 进入项目目录

```bash
cd d:\demo\go
```

3. 安装依赖

```bash
go mod tidy
```

4. 运行项目

```bash
go run main.go
```

5. 验证服务是否正常运行

```bash
curl http://localhost:8080/health
```

## API接口

### 健康检查

- **URL**: `/health`
- **方法**: `GET`
- **描述**: 检查服务是否正常运行
- **响应示例**:

```json
{
  "status": "ok",
  "timestamp": "2026-01-15T10:00:00Z"
}
```

### 获取用户列表

- **URL**: `/api/users`
- **方法**: `GET`
- **描述**: 获取所有用户信息
- **响应示例**:

```json
{
  "users": [
    {
      "id": 1,
      "name": "张三",
      "email": "zhangsan@example.com"
    },
    {
      "id": 2,
      "name": "李四",
      "email": "lisi@example.com"
    }
  ]
}
```

### 创建用户

- **URL**: `/api/users`
- **方法**: `POST`
- **描述**: 创建新用户
- **请求示例**:

```json
{
  "name": "王五",
  "email": "wangwu@example.com"
}
```

- **响应示例**:

```json
{
  "message": "用户创建成功",
  "timestamp": "2026-01-15T10:00:00Z"
}
```

### 获取系统信息

- **URL**: `/api/data`
- **方法**: `GET`
- **描述**: 获取系统信息和功能特性
- **响应示例**:

```json
{
  "message": "欢迎使用Go后端API",
  "version": "1.0.0",
  "timestamp": "2026-01-15T10:00:00Z",
  "features": [
    "用户管理",
    "数据查询",
    "文件上传",
    "认证授权"
  ]
}
```

## 项目结构

```
.
├── main.go          # 主启动文件
├── go.mod           # Go模块文件
├── README.md        # 项目说明文档
├── controllers/     # 控制器（可扩展）
├── models/          # 数据模型（可扩展）
├── routes/          # 路由配置（可扩展）
├── middleware/      # 中间件（可扩展）
└── utils/           # 工具函数（可扩展）
```

## 配置

### 环境变量

- `PORT`: 服务器端口，默认8080

## 开发指南

### 添加新API

1. 在 `main.go` 的 `setupRoutes` 函数中添加新路由
2. 实现对应的处理函数
3. 确保添加适当的CORS头
4. 处理错误情况

### 测试

使用curl或Postman测试API接口：

```bash
# 健康检查
curl http://localhost:8080/health

# 获取用户列表
curl http://localhost:8080/api/users

# 创建用户
curl -X POST -H "Content-Type: application/json" -d '{"name":"王五","email":"wangwu@example.com"}' http://localhost:8080/api/users

# 获取系统信息
curl http://localhost:8080/api/data
```

## 部署

### 编译项目

```bash
go build -o backend-server main.go
```

### 运行编译后的程序

```bash
./backend-server
```

## 许可证

MIT
