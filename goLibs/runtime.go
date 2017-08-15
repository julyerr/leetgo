import "runtime"

GoSched() ：让当前线程让出 cpu 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行
Goexit()  :exit the goroutine while defer will still exec
NumGoruntine() : nums of goroutine in the job
GOMAXPORCS() :setting the max procs to use
NumCPU()  : cpus
GOOS        : arch



