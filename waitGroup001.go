package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Program start \n")
	var waitGroup sync.WaitGroup
	waitGroup.Add(5)
	for i := 0;i < 5;i++{
		go concurrentTask(i, &waitGroup)
	}
	waitGroup.Wait()
	finishTask()
	fmt.Printf("Program end \n")
}

func finishTask(){
	fmt.Println("Executing finish task")
}

func concurrentTask(taskNumber int,	waitGroup *sync.WaitGroup){
	fmt.Printf("Begin Execute tast number %d \n", taskNumber)
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("End execute task number %d \n", taskNumber)
	waitGroup.Done()
}