package engine

import (
	"log"
	"time"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}
func (e *ConcurrentEngine) Run(seeds  ...Request)  {
	input  := make(chan Request)
	time.Sleep(1*time.Second)
	output := make(chan ParseResult)
	e.Scheduler.ConfigureMasterWorkerChan(input)

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	for i:=0; i<e.WorkerCount;i++{
		createWorker(input,output)
	}
	for {
		result := <- output
		for _,item := range result.Items{
			log.Printf("Got em item %v",item)
		}

		for _,request := range result.Reuqests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult)  {
	go func() {
		for  {
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}