#### RwGate 编译
```
https://github.com/iuoon/rwserver.git  
cd rwserver
go build ./RwGate/main/GateServer.go
```

#### 编译动态库
linux下创建so比较方便，直接通过如下命令：
```
go build -buildmode=c-shared -o libtq.so
```
windows需要安装gcc编译器,windows下go不支持生成动态库，需要另外的方式进行编译
https://stackoverflow.com/questions/40573401/building-a-dll-with-go-1-7  
编译静态库
go build -buildmode=c-archive -o libtq.a  

gobblob.c文件，然后把go代码中要导出的函数，在gobblob.c中全部调用一遍。
```
#include <stdio.h>
#include "libtq.h"

// force gcc to link in go runtime (may be a better solution than this)
void dummy() {
   // 所有在go中要导出的代码都在这里调用一次，参数随便写，只要能编译ok即可
    gotq_init(NULL,NULL,NULL);
    gotq_deinit(NULL);
    gotq(NULL,NULL,NULL,NULL,NULL,NULL);
}

int main() {

}
```
执行如下命令，生成dll  
gcc -shared -pthread -o libtq.dll gotq.c libtq.a -lWinMM -lntdll -lWS2_32 -Iinclude