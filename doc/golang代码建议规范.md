

# 业务代码开发规范总结和整理

[TOC]

## 目标和原则

1. 代码分层明确，不跃层
2. 层与层之间通过接口衔接，实现依赖反转
3. 从上往下根据实际业务实现业务流程，进而构建解决方案
4. 代码需要适应不同的版本，不要在代码中加大量的if - else的逻辑来判断版本号
5. 区分代码公共部分和私有部分，不要掺和在一起
6. 代码结构中的每一层和每一个方法都只做有明确界限的工作，不要写一个大而全的方法，妄图解决所有问题，尽量保持方法或函数的纯净。



## 遵循以下前置规范

1. [golang code review comments](https://github.com/golang/go/wiki/CodeReviewComments)
2. [effective go](https://golang.google.cn/doc/effective_go.html)



## 规范详情

### 代码结构

#### 代码分层说明

- **service(api)** (service层，如果服务只对外提供单一格式接口，那么该目录就无需再分那么多层)
    - **http controller**
    - **grpc**
    - **service**
- **business** (业务逻辑层)
    - *interface.go* (申明该层包含哪些业务逻辑，每组业务逻辑需要实现哪些方法)
- **module** (模块层，将业务流程的环节拆开，实现一个一个的独立逻辑的方法，有业务层统一组织，该层每一块相对独立，可以承担部分业务逻辑，不与该层其他模块方法交互)
- **data** (数据层)
    - **mysql** 操作mysql数据
    - **redis** 操作redis数据
    - **kafka** 操作消息队列数据
    - **external** 操作外部第三方接口数据
    - **other-micro-service** 操作其他服务数据
- **common** (公共模块，例如一些常量方法需要贯穿若干层，即将他们抽出到这里来)
    - **sdk** (对接第三方一些接口的sdk)
    - **util** (公共工具方法)
    - **definition** (公共类型，常量申明与定义)
    - **lib** （独立于业务的算法和数据结构）
- **doc** (项目文档，sql文件，变更记录)
- **config** (存放项目的配置独立于所有层之外)
- **script** (辅助脚本)
- **vender**  (依赖管理)
- *main.go* (项目入口)
- .... (其他文件，例如.gitignore，Dockerfile.......)


代码从上到下，业务关联性越来越弱，趋于独立


#### 文件夹以及文件命名规范

1. 文件夹与文件名均不使用复数


### 常量与类型

1. 所有用常量表示的值都需要携带自定义类型，避免在函数或方法接收这些常量参数时，误传常量范围之外的值。
2. 使用常量表明一组类型的状态时，需要统一常量前的类型前缀，比如表明用户状态的常量，写成UserStatusNormal = 1。
3. 如果是所有模块公用的常量与变量，单独放在一个文件中，如果是每个模块单独用到的常量以及变量放在每个模块的文件中，同样的类型需要跨层使用那么也放在公共的目录下。
4. 业务常量定义在module这一层。
5. 含义一样的变量，命名保持一致，不要这里写userID，那里写uid。


### 数据库与缓存以及其他数据源
#### mysql
1. 统一上层文件夹的名称使用data，表明这是数据层，负责操作各种数据源的数据，例如：mysql的操作，就在data目录下建立一层mysql的文件夹，用于处理mysql相关的数据存取
2. 当前这一层，***一定不要掺杂业务逻辑在里面***，他们只负责最单纯的数据的读取和写入。
3. 不要在数据层定义业务中使用的常量。
4. 需要被包裹在事务中的方法使用到的db对象必须从外部传入。
5. 将需要在事物中执行的方法，统一打包成一个slice，通过方法（参考代码示例）一次性批量执行，避免多个地方反复写开启事务的逻辑代码。
6. 所有更新操作，都需要在数据层设置时间。

#### redis

1. 如果需要一次性执行多条命令，不要一条一条的执行，应该使用piplines，来减少多次执行单个命令时的round trip的开销。
2. 




### 参数验证
1. 所有的参数验证只在business层去验证，下面的model层，认为传递的数据是合法的，避免数据一层一层的校验。
2. api层无需校验参数，只需要将参数传递给业务逻辑层即可。


### 方法函数入参以及返回值
1. 通知类的方法或函数，如果仅仅只是通知一下的话，在方法内部包装成异步，日志记在方法内部，无需返回error给调用方


### 错误定义
1. 需要在business层定义所有的错误error对象，详见代码示例。
2. 错误码使用常量表示，对应的错误描述可以用map去映射。
3. 错误对象要实现error接口，并提供额外的与业务相关的方法。


### 日志
1. 业务层日志模板 `[(层级，例如module)当前函数名 -> 内部调用的方法名] 错误的中文描述(附带的详细参数，可忽略) error: %v`
2. 返回出去的err由接收方处理生成err的地方无需重复打印
3. 所有的日志都需要携带上request_id，用来查看用户每一次请求的执行过程。
4. 需要打印用户的请求与相应，如果返回的数据不是关键数据，可以打印简略信息或是省略。



### 指针

1. 如果返回出去的数据无需外部修改，一律不要使用指针。
2. 如果不是需要在内存中驻留很久的数据，一律不要使用指针。
3. 小对象在栈上复制的速度，远快于指针在堆内存上分配的速度。

### 并发

1. 如果一个方法必须要使用go关键词才能执行，请将并发的过程放在方法内部处理，避免外部忘记使用go异步，并且方法要以**Async**开头，例如**func AsyncSendUserLoginMessage()**。
2. 一般情况下，不要直接把channel暴露给外部方法
3. 关闭channel的正确姿势，参见 [How To Gracefully Close Channels](https://go101.org/article/channel-closing.html)



### Context

[详情参考](https://github.com/golang/go/wiki/CodeReviewComments#contexts)

1. 内部函数的第一个参数应该是context。
2. 在请求源头生成请求的上下文信息，例如request_id，通过context传递下去.





## 代码示例

### 处理数据库事务

```go
type TransactionHandler func(orm orm.Ormer) error

// ExecuteWithTransaction 在一个事务中
func ExecuteWithTransaction(handlers ...TransactionHandler) (err error) {
	// 开启一个事务
	db := GetDistributionDB()
	err = db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			db.Rollback()
		} else {
			db.Commit()
		}
	}()

	for _, handler := range handlers {
		err = handler(db)
		if err != nil {
			return
		}
	}

	return nil
}
```