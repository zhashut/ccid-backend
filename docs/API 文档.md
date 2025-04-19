

# API 接口文档

## 基础信息
- 基础路径: `/api`
- 所有接口都支持跨域访问

## 1. 知识产权信息接口

### 1.1 创建知识产权信息
- **接口**: POST `/api/dataProperty/create`
- **描述**: 创建新的知识产权信息记录
- **请求参数**:
  ```json
  {
    "type": "string",    // 数据类型（专利/论文/软件著作权）
    "title": "string",   // 标题
    "source": "string",  // 来源
    "status": "int"      // 状态（1:进行中, 2:已完成）
  }
  ```

- **响应字段**

  ```json
  {
    "code": 0,
    "data": {
      "id": 12345,
      "type": "专利",
      "title": "一种新型数据处理方法",
      "source": "国家知识产权局",
      "time": "2023-06-15",
      "status": 1
    },
    "msg": "success"
  }
  ```

### 1.2 更新知识产权信息

- **接口**: POST `/api/dataProperty/update`
- **描述**: 更新现有的知识产权信息
- **请求参数**:
  ```json
  {
    "id": "int64",      // 数据ID
    "type": "string",   // 数据类型
    "title": "string",  // 标题
    "source": "string", // 来源
    "status": "int"     // 状态
  }
  ```

- **响应字段**

  ```json
  {
    "code": 0,
    "data": {
      "id": 12345,
      "type": "专利",
      "title": "一种新型数据处理方法",
      "source": "国家知识产权局",
      "time": "2023-06-15",
      "status": 1
    },
    "msg": "success"
  }
  ```

### 1.3 删除知识产权信息

- **接口**: POST `/api/dataProperty/delete`
- **描述**: 删除指定的知识产权信息
- **请求参数**:
  ```json
  {
    "id": "int64"  // 数据ID
  }
  ```

- **响应字段**

  ```json
  {
    "code": 0,
    "data": {},
    "msg": "success"
  }
  ```

### 1.4 获取知识产权信息列表

- **接口**: GET `/api/dataProperty/list`
- **描述**: 获取知识产权信息列表
- **请求参数**:
  ```
  keyword: string  // 搜索关键词
  type: string     // 数据类型筛选
  status: int      // 状态筛选
  page: int        // 页码
  pageSize: int        // 每页数量
  ```

- **响应字段**

  ```json
  {
    "code": 0,
    "data": {
      "list": [
        {
          "id": 12345,
          "type": "专利",
          "title": "方法1",
          "source": "来源1",
          "time": "2023-06-15",
          "status": 1
        },
        {
          "id": 12346,
          "type": "论文",
          "title": "方法2",
          "source": "来源2",
          "time": "2023-06-16",
          "status": 2
        }
      ],
      "total": 2
    },
    "msg": "success"
  }
  ```

## 2. 专利数据接口

### 2.1 创建专利数据
- **接口**: POST `/api/dataPatent/create`
- **描述**: 创建新的专利数据记录
- **请求参数**:

  ```json
  {
    "id": "int64",        // 专利ID
    "title": "string",    // 专利名称
    "applicant": "string",// 申请人
    "field": "string"     // 技术领域
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 12345,
      "title": "一种基于深度学习的图像识别方法",
      "applicant": "人工智能研究院",
      "date": "2023-05-20",
      "field": "计算机视觉"
    },
    "msg": "success"
  }
  ```

### 2.2 更新专利数据

- **接口**: POST `/api/dataPatent/update`
- **描述**: 更新现有的专利数据
- **请求参数**:

  ```json
  {
    "id": "int64",        // 专利ID
    "title": "string",    // 专利名称
    "applicant": "string",// 申请人
    "field": "string"     // 技术领域
  }
  ```

- **响应参数**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 12345,
      "title": "一种基于深度学习的图像识别方法",
      "applicant": "人工智能研究院",
      "date": "2023-05-20",
      "field": "计算机视觉"
    },
    "msg": "success"
  }
  ```

### 2.3 删除专利数据

- **接口**: POST `/api/dataPatent/delete`
- **描述**: 删除指定的专利数据
- **请求参数**:
  ```json
  {
    "id": "int64"  // 专利ID
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {},
    "msg": "success"
  }
  ```

### 2.4 获取专利数据列表

- **接口**: GET `/api/dataPatent/list`
- **描述**: 获取专利数据列表
- **请求参数**:

  ```
  keyword: string  // 搜索关键词
  field: string    // 技术领域筛选
  page: int        // 页码
  pageSize: int        // 每页数量
  ```

- **响应字段**

  ```json
  {
    "code": 0,
    "data": {
        "list": [
          {
              "id": 1,
              "title": "一种人工智能算法",
              "applicant": "科技有限公司",
              "date": "2023-01-15",
              "field": "计算机科学"
          }
      ],
      "total": 1
    },
    "msg": "success"
  }
  ```

## 3. 团队成员接口

### 3.1 创建团队成员
- **接口**: POST `/api/teamMember/create`
- **描述**: 创建新的团队成员
- **请求参数**:

  ```json
  {
    "name": "string",  // 姓名
    "role": "int"      // 角色（1:查看, 2:编辑, 3:管理）
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 123,
      "name": "张三",
      "role": 2
    },
    "msg": "success"
  }
  ```

### 3.2 更新团队成员

- **接口**: POST `/api/teamMember/update`
- **描述**: 更新团队成员信息
- **请求参数**:
  ```json
  {
    "id": "int64",    // 成员ID
    "name": "string", // 姓名
    "role": "int"     // 角色
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 123,
      "name": "张三",
      "role": 2
    },
    "msg": "success"
  }
  ```

### 3.3 更新成员角色

- **接口**: POST `/api/teamMember/updateRole`
- **描述**: 更新团队成员角色
- **请求参数**:
  ```json
  {
    "id": "int64",  // 成员ID
    "role": "int"   // 新角色
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 123,
      "name": "张三",
      "role": 1
    },
    "msg": "success"
  }
  ```

### 3.4 删除团队成员

- **接口**: POST `/api/teamMember/delete`
- **描述**: 删除团队成员
- **请求参数**:
  ```json
  {
    "id": "int64"  // 成员ID
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {},
    "msg": "success"
  }
  ```

### 3.5 获取团队成员列表

- **接口**: GET `/api/teamMember/list`
- **描述**: 获取团队成员列表
- **请求参数**:
  ```
  name: string  // 姓名搜索
  role: int     // 角色筛选
  page: int     // 页码
  pageSize: int  // 每页数量
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "list": [
        {
          "id": 123,
          "name": "张三",
          "role": 1
        },
        {
          "id": 124,
          "name": "李四",
          "role": 3
        }
      ],
      "total": 2
    },
    "msg": "success"
  }
  ```

## 4. 数据报告接口

### 4.1 创建数据报告
- **接口**: POST `/api/dataReport/create`
- **描述**: 创建新的数据报告
- **请求参数**:
  ```json
  {
    "name": "string",    // 报告名称
    "type": "string",    // 报告类型（专利分析/技术分析/市场分析）
    "status": "int",     // 状态（1:进行中, 2:已完成）
    "content": "string"  // 报告内容
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 1001,
      "name": "2023年人工智能专利分析报告",
      "type": "专利分析",
      "status": 2,
      "content": "报告内容文本...",
      "time": "2023-12-15"
    },
    "msg": "success"
  }
  ```

### 4.2 更新数据报告

- **接口**: POST `/api/dataReport/update`
- **描述**: 更新现有的数据报告
- **请求参数**:
  ```json
  {
    "id": "int64",      // 报告ID
    "name": "string",   // 报告名称
    "type": "string",   // 报告类型
    "status": "int",    // 状态
    "content": "string" // 报告内容
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 1001,
      "name": "2023年人工智能专利分析报告",
      "type": "专利分析",
      "status": 2,
      "content": "报告内容文本...",
      "time": "2023-12-15"
    },
    "msg": "success"
  }
  ```

### 4.3 更新报告状态

- **接口**: POST `/api/dataReport/updateStatus`
- **描述**: 更新数据报告状态
- **请求参数**:
  ```json
  {
    "id": "int64",   // 报告ID
    "status": "int"  // 新状态
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 1001,
      "name": "2023年人工智能专利分析报告",
      "type": "专利分析",
      "status": 2,
      "content": "报告内容文本...",
      "time": "2023-12-15"
    },
    "msg": "success"
  }
  ```

### 4.4 删除数据报告

- **接口**: POST `/api/dataReport/delete`
- **描述**: 删除指定的数据报告
- **请求参数**:
  ```json
  {
    "id": "int64"  // 报告ID
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {},
    "msg": "success"
  }
  ```

### 4.5 获取数据报告列表

- **接口**: GET `/api/dataReport/list`
- **描述**: 获取数据报告列表
- **请求参数**:
  ```
  type: string      // 报告类型筛选
  status: int       // 状态筛选
  keyword: string   // 搜索关键词
  start_time: string// 开始时间
  end_time: string  // 结束时间
  page: int         // 页码
  pageSize: int     // 每页数量
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "list": [
        {
          "id": 1001,
          "name": "AI专利报告",
          "type": "专利分析",
          "status": 2,
          "time": "2023-12-15"
        },
        {
          "id": 1002,
          "name": "新能源市场报告",
          "type": "市场分析",
          "status": 1,
          "time": "2023-12-18"
        }
      ],
      "total": 2
    },
    "msg": "success"
  }
  ```

## 5. 数据任务接口

### 5.1 创建数据任务
- **接口**: POST `/api/dataTask/create`
- **描述**: 创建新的数据任务
- **请求参数**:
  ```json
  {
    "name": "string",    // 任务名称
    "type": "string",    // 数据类型（专利/论文）
    "manager": "string", // 负责人
    "status": "int"      // 状态（1:进行中, 2:已完成）
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 1001,
      "name": "专利数据清洗任务",
      "type": "专利",
      "manager": "张三",
      "status": 1
    },
    "msg": "success"
  }
  ```

### 5.2 更新数据任务

- **接口**: POST `/api/dataTask/update`
- **描述**: 更新现有的数据任务
- **请求参数**:
  ```json
  {
    "id": "int64",      // 任务ID
    "name": "string",   // 任务名称
    "type": "string",   // 数据类型
    "manager": "string",// 负责人
    "status": "int"     // 状态
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 1001,
      "name": "专利数据清洗任务",
      "type": "专利",
      "manager": "张三",
      "status": 1
    },
    "msg": "success"
  }
  ```

### 5.3 更新任务状态

- **接口**: POST `/api/dataTask/updateStatus`
- **描述**: 更新数据任务状态
- **请求参数**:
  ```json
  {
    "id": "int64",  // 任务ID
    "status": "int8"// 新状态
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "id": 1001,
      "name": "专利数据清洗任务",
      "type": "专利",
      "manager": "张三",
      "status": 1
    },
    "msg": "success"
  }
  ```

### 5.4 删除数据任务

- **接口**: POST `/api/dataTask/delete`
- **描述**: 删除指定的数据任务
- **请求参数**:
  ```json
  {
    "id": "int64"  // 任务ID
  }
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {},
    "msg": "success"
  }
  ```

### 5.5 获取数据任务列表

- **接口**: GET `/api/dataTask/list`
- **描述**: 获取数据任务列表
- **请求参数**:
  ```
  type: string    // 数据类型筛选
  status: int8    // 状态筛选
  manager: string // 负责人筛选
  keyword: string // 搜索关键词
  page: int       // 页码
  pageSize: int       // 每页数量
  ```

- **响应字段**：

  ```json
  {
    "code": 0,
    "data": {
      "list": [
        {
          "id": 1001,
          "name": "专利数据清洗",
          "type": "专利",
          "manager": "张三",
          "status": 1
        },
        {
          "id": 1002,
          "name": "论文数据整理",
          "type": "论文",
          "manager": "李四",
          "status": 2
        }
      ],
      "total": 2
    },
    "msg": "success"
  }
  ```

## 通用说明

### 分页请求参数
所有列表接口都支持以下分页和排序参数：
```json
{
  "current": "int64",    // 当前页码
  "pageSize": "int64",   // 每页数量
  "sortField": "string", // 排序字段
  "sortOrder": "string"  // 排序顺序（升序/降序）
}
```

### 通用响应格式
所有接口的响应都遵循以下格式：
```json
{
  "code": "int",       // 状态码
  "message": "string", // 响应消息
  "data": "object"     // 响应数据
}
```

### 字段说明

#### 1. 状态字段
所有涉及状态的字段都使用以下值：
- 1: 进行中
- 2: 已完成

#### 2. 角色字段
团队成员角色使用以下值：
- 1: 查看权限
- 2: 编辑权限
- 3: 管理权限

### 时间格式规范
- 所有日期类型字段统一使用 `yyyy-MM-dd` 格式

### 错误码说明
```go
ErrorInvalidParams = ErrorCode{http.StatusForbidden, 40000, "请求参数错误"}
ErrorNotLogin      = ErrorCode{http.StatusUnauthorized, 40100, "未登录"}
ErrorNoAuth        = ErrorCode{http.StatusUnauthorized, 40101, "无权限"}
ErrorNotFound      = ErrorCode{http.StatusNotFound, 40400, "请求数据不存在"}
ErrorFORBIDDEN     = ErrorCode{http.StatusForbidden, 40300, "禁止访问"}
ErrorSYSTEM        = ErrorCode{http.StatusInternalServerError, 50000, "系统内部异常"}
ErrorPERATION      = ErrorCode{http.StatusBadRequest, 50001, "操作失败"}
```