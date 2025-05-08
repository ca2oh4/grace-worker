grace worker
===

### 整体结构

web 和 worker 之间通过 redis 的消息队列通讯。

消息包括
- web 发布任务
- worker 读取任务
- worker 推送任务进度及状态
- web 读取任务进度

#### web

- 提供 API 服务
- 创建任务
- 读取任务进度及状态

#### worker

- 读取任务
- 执行任务
- 上报任务进度及状态


### 部署及运行

采用 docker-compose 部署

web 和 worker 均支持优雅停机更新

#### web

web 只有一个实例

#### worker

worker 的两个实例（假设分别为：w1，w2），分为一下状态：

1. 初次部署：w1 处理任务，w2 不处理任务。
2. 更新部署：w1 处理完成剩余任务，再停止不更新。w2 由于没有任务在处理，直接更新。



