#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>
#include <stdint.h>

#ifdef __cplusplus
} // extern "C"
#endif

#define LOGV(format, ...) printf(format, ##__VA_ARGS__)
#define LOGD(format, ...) printf(format, ##__VA_ARGS__)
#define LOGI(format, ...) printf(format, ##__VA_ARGS__)
#define LOGW(format, ...) printf(format, ##__VA_ARGS__)
#define LOGE(format, ...) printf(format, ##__VA_ARGS__)
#define LOGF(format, ...) printf(format, ##__VA_ARGS__)

/* -------------------------------------------------- */

/// 类型定义：读取回调函数
typedef int (*ReadCallback)(void *opaque, uint8_t *buf, int buf_size);
/// 类型定义：写入回调函数
typedef int (*WriteCallback)(void *opaque, uint8_t *buf, int buf_size);

/// 自定义上下文结构体
typedef struct MyContext {
    void *opaque;                  //< 透明指针
    ReadCallback read_callback;    //< 读取回调函数
    WriteCallback write_callback;  //< 写入回调函数
} MyContext;

/// 分配上下文内存
MyContext *alloc_my_context(void);

/// 释放上下文内存
void free_my_context(MyContext *context);

/// 初始化上下文内存
int init_my_context(MyContext *context, void *opaque, ReadCallback read_callback, WriteCallback write_callback);

/// 掉用
void invoke(MyContext *context);
