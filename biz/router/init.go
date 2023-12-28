package router

func InitRouter() {
	r := SetUpRouter()
	err := r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		panic(err)
	}
}
