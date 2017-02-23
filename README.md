
### First iteration. Shared libraries.

First I used _c-archive_

> go build -buildmode=c-archive -o libcglog.so cglog.go


Then I created simple _main_ for test.
In this case arrise coompilation error that some runtime functions from go part conflict with c-archive part.

I changed it to shared object _c-shared_

> go build -buildmode=c-shared -o libcglog.so cglog.go

Problem was gone. But I had such error


```
runtime: address space conflict: map(0xc420000000) = 0x7facc20ca000
fatal error:runtime: address space conflict

runtime stack:
runtime.throw(0x7facc1d848b0, 0x1f)
...
```

It was pretty random. Sometime it worked but often didn't.


I added race flag

# solution is -race flag
I compiled the _main_ with race
Problem was gone.

I tried compile only library with _-race_ flag.



# race flag only for library
```
==15813==ERROR: ThreadSanitizer failed to allocate 0x26e9000 (40800256) bytes at address 21c0a100e4000 (errno: 12)
unexpected fault address 0x0
fatal error: fault
[signal SIGSEGV: segmentation violation code=0x80 addr=0x0 pc=0x7f0283d62dfd]

goroutine 1 [running, locked to thread]:
runtime.throw(0x7f0283c95582, 0x7f0283c94fd7)
        /home/petr/.gvm/gos/go1.8/src/runtime/panic.go:596 +0x97 fp=0x7f0282d60df0 sp=0x7f0282d60dd0

goroutine 17 [syscall, locked to thread]:
runtime.goexit()
        /home/petr/.gvm/gos/go1.8/src/runtime/asm_amd64.s:2197 +0x1
Aborted (core dumped)


go build -compiler=gccgo -buildmode=c-shared -o libcglog.so cglog.go
go build github.com/golang/glog: : fork/exec : no such file or directory
Makefile:4: recipe for target 'all' failed
make: *** [all] Error 1
```


Anyway go code which called from c and _main_ have different runtimes.

Looks like I should solve it in different way.
