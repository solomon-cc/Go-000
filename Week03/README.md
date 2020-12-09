学习笔记
## 作业
#### 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。


### 思路
定义一个全局信号chan，每一个http serve使用ctx进行服务，在捕捉到系统信号后，执行cancel，最后关闭stop channel，退出主函数



### 伪代码
````
func main {
   go func() {
      done <- httpServer()
}()
    <- done 
    close stop()
}

func httpServer(stop chan struck{}) error {
     go func() {
       <- stop
       http.Shutdown()
       }()
     return http.Listen()
````


