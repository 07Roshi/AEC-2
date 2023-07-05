//check point synchronization

package main
import "fmt"
import "sync"
import "time"

func worker(id int, checkpoint chan bool, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d : Starting\n",id)
	time.Sleep(time.Duration(id)*time.Second))

	fmt.Printf("Worker %d : Checkpoint reached\n",id)
	checkpoint <- true //Signal that checkpoint is reached
	
	<-resume //Wait for resume signal

	fmt.Printf("Worker %d : Resuming\n",id)
	//continue with the remaining work
}

func main(){
	numWorkers := 5
	checheckpoint := make(chan bool)
	resume := make(chan bool)
	var wg sync.WaitGroup

	//Launch the worker goroutines
	for i := 1; i <= numWorkers; i++{
		wg.Add(1)
		go worker(i, checkpoint, resume, &wg)
	}
}
