#include "native.h"

/// 分配上下文内存
MyContext *alloc_my_context(void) {
    return (MyContext *) malloc(sizeof(MyContext));
}

/// 释放上下文内存
void free_my_context(MyContext *context) {
    if (context != NULL) {
        free(context);
    }
}

/// 初始化上下文内存
int init_my_context(MyContext *context, void *opaque, ReadCallback read_callback, WriteCallback write_callback) {
    if (context == NULL) {
        return -1;
    }

    context->opaque = opaque;
    context->read_callback = read_callback;
    context->write_callback = write_callback;
    return 0;
}

/// 掉用
void invoke(MyContext *context) {
    uint8_t buf[4096];
    int size = context->read_callback(context->opaque, buf, sizeof(buf));
    context->write_callback(context->opaque, buf, size);
}
