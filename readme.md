# Go 链接C静态库



编译静态库
```bash
gcc -c sum.c          // produces a sum.o object file
ar -rc libsum.a sum.o //生产静态库
nm libsum.a           //验证
```

main.c 引用静态库
```bash
// #cgo CFLAGS: -I./static_libraries
// #cgo LDFLAGS: ./static_libraries/libsum.a
// #include <sum.h>
import "C"
```

测试
```bash
go run main.go
go build
```