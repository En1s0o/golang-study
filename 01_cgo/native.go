package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lnative
// #include <native.h>
//
// // 声明下面两个函数，并在 Go 中通过 export 实现它们，注意 C 数据类型与 Go 数据类型的映射
// int ReadFunc(void *opaque, uint8_t *buf, int buf_size);
// int WriteFunc(void *opaque, uint8_t *buf, int buf_size);
import "C"

import (
    "unsafe"
)

//export ReadFunc
func ReadFunc(opaque unsafe.Pointer, buf *C.uint8_t, size C.int) C.int {
    // 如果方法提供
    h := *(*MyHandler)(opaque)
    return C.int(h.ReadHandler((*uint8)(buf), int(size)))
}

//export WriteFunc
func WriteFunc(opaque unsafe.Pointer, buf *C.uint8_t, size C.int) C.int {
    h := *(*MyHandler)(opaque)
    return C.int(h.WriteHandler((*uint8)(buf), int(size)))
}

// 结构体映射
type (
    MyContext C.struct_MyContext
)

func AllocMyContext() *MyContext {
    // 调用 C 中的 alloc_my_context() 函数，最后转换成 Go 中的 MyContext
    return (*MyContext)(C.alloc_my_context())
}

func (s *MyContext) FreeMyContext() {
    // 调用 C 中的 free_my_context() 函数，这样封装看起来更友好
    C.free_my_context((*C.struct_MyContext)(s))
}

// 这里没有设计一个 interface，然后 MyHandler 实现它，是因为 MyHandler 一般包含很多字段
// 使用 interface 接收将触运行时发异常 panic: runtime error: cgo argument has Go pointer to Go pointer
func (s *MyContext) InitMyContext(callback *MyHandler) int {
    // 调用 C 中的 init_my_context() 函数，这样封装看起来更友好
    // 在 native.h 通过 typedef 定义了 2 个类型：ReadCallback 和 WriteCallback
    // C.ReadFunc 和 C.WriteFunc 是上面 export 导出的函数，强转之后就可以给 C 使用了
    return int(C.init_my_context(
        (*C.struct_MyContext)(s),
        unsafe.Pointer(callback),
        C.ReadCallback(C.ReadFunc),
        C.WriteCallback(C.WriteFunc)))
}

func (s *MyContext) Invoke() {
    // 调用 C 中的 invoke() 函数，这样封装看起来更友好
    C.invoke((*C.struct_MyContext)(s))
}
