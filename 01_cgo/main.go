package main

import "C"

func main() {
    // 分配上下文
    context := AllocMyContext()
    // 初始化上下文
    context.InitMyContext(CreateHandler("id_001", &DummyHandler{name: "dummy"}))
    // 调用
    context.Invoke()
    context.Invoke()
    // 释放上下文
    context.FreeMyContext()
}
