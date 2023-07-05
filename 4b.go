//check point synchronization

package main
import (
	"fmt"
	"sync"
    "time"
)
func worker(id int, checkpoint chan bool, resume chan bool, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d : Starting\n",id)
	time.Sleep(time.Duration(id)*time.Second)

	fmt.Printf("Worker %d : Checkpoint reached\n",id)
	checkpoint <- true //Signal that checkpoint is reached
	
	<-resume //Wait for resume signal

	fmt.Printf("Worker %d : Resuming\n",id)
	//continue with the remaining work
}

func main(){
	numWorkers := 5
	checkpoint := make(chan bool)
	resume := make(chan bool)
	var wg sync.WaitGroup

	//Launch the worker goroutines
	for i := 1; i <= numWorkers; i++{
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
	}

	//Wait for all workers to reach the checkpoint
	for i := 1; i <= numWorkers; i++{

		<-checkpoint
	}

	fmt.Println("All workers reached the checkpoint")

	fmt.Println("Resuming all workers")
	//Signal all workers to resume

	for i := 1; i <= numWorkers; i++{
		resume <- true
	}

	wg.Wait()
	fmt.Println("All workers completed their work")
}
