package distribute

import (
	"fmt"
	"time"
	"strconv"
)

type ThreadPool struct {
    Jobs chan func() error;
    ThreadNumber int;
    JobNumber int;

    Result chan error;
    FinishCallback func();
}

//初始化
func (p *ThreadPool) Init(ThreadNumber int, JobNumber int)  {
    p.ThreadNumber = ThreadNumber;
    p.JobNumber = JobNumber;
    p.Jobs = make(chan func() error, JobNumber);
    p.Result = make(chan error, JobNumber);
}

func (p *ThreadPool) Start()  {
    //开启 number 个goruntine
    for i:=0;i<p.ThreadNumber;i++ {
        go func() {
            for {
                task,ok := <-p.Jobs
                if !ok {
                    break;
                }
                err := task();
                p.Result <- err;
								fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " common.go ThreadPool thread#" + strconv.Itoa(i) + "run task")
            }
        }();
    }

    //获取每个任务的处理结果
    for j:=0;j<p.ThreadNumber;j++ {
        res,ok := <-p.Result;
        if !ok {
            break;
        }
        if res != nil {
            fmt.Println(res);
        }
    }

    //结束回调函数
    if p.FinishCallback != nil {
        p.FinishCallback();
    }
}

//关闭
func (p *ThreadPool) Stop()  {
    close(p.Jobs);
    close(p.Result);
}

func (p *ThreadPool) AddTask(task func() error)  {
    p.Jobs <- task;
}

func (p *ThreadPool) SetFinishCallback(fun func())  {
    p.FinishCallback = fun;
}
