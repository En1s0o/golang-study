package main

// #include <string.h> // 引入 memcpy 函数
import "C"

import (
    "log"
    "sync"
    "unsafe"
)

// 定义一个全局的 HandlerMap，可以有效避免 panic: runtime error: cgo argument has Go pointer to Go pointer
var HandlerMap = new(sync.Map)

type InternalHandler interface {
    receive(buf []byte)
}

type DummyHandler struct {
    name string
}

func (o DummyHandler) receive(buf []byte) {
    log.Printf("[%s] DummyHandler %s", o.name, string(buf))
}

type MyHandler struct {
    id string

    // 这里结构体是有很多限制的，不推荐下面的使用方式，
    // 例如希望通过 websocket 把数据发送出去，
    // 除非 websocket 对象仅仅只有函数（显然不实际），而没有字段，
    // 否则运行时将会发生异常，即 panic: runtime error: cgo argument has Go pointer to Go pointer
    // o InternalHandler
}

func (s MyHandler) ReadHandler(buf *uint8, size int) int {
    log.Printf("[%s] ReadHandler %d", s.id, size)
    buffer := []byte("hello world!!")
    // 将 Go 的 []byte 数据拷贝到 C 的 uint8_t * 中
    C.memcpy(unsafe.Pointer(buf), unsafe.Pointer(&buffer[0]), C.size_t(len(buffer)))
    return len(buffer)
}

func (s MyHandler) WriteHandler(buf *uint8, size int) int {
    buffer := make([]byte, size)
    // 将 C 的 uint8_t * 数据拷贝到 Go 的 []byte 中
    C.memcpy(unsafe.Pointer(&buffer[0]), unsafe.Pointer(buf), C.size_t(size))
    log.Printf("[%s] WriteHandler %d, %s", s.id, size, string(buffer))
    // 根据 id 找到内部处理对象，并将数据传递给它
    if load, ok := HandlerMap.Load(s.id); ok {
        handler := (load).(InternalHandler)
        handler.receive(buffer)
    }
    return 0
}

func CreateHandler(id string, h InternalHandler) *MyHandler {
    HandlerMap.Store(id, h)
    return &MyHandler{
        id: id,
    }
}
