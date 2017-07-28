##Arrows : Go Package for Scalable and Big Scale Projects (Don't Use Mutex)

<img align="right" src="https://github.com/SabzCity/Go-Library/blob/master/arrows/etc/logo.png">

###This package provide four ways for concurrently programming in Go

## Start using it

1. Download and install it:

    ```sh
    $ go get github.com/SabzCity/Go-Library/arrows
    ```

2. Import it in your code:

    ```go
    import "github.com/SabzCity/Go-Library/arrows"
    ```

####1.Circuit Breaker :
You can put your work in circuit, the circuit do work Alternative and when
your work return an error circuit retry work, but when number of errors overflowed
breaker is open and break circuit .

```go
import "github.com/SabzCity/Go-Library/arrows"

func main() {
  //Create a new circuit.
  cir:=arrows.CreateCircuit(1, 3)

  //Do work with circuit.
  err:=cir.Run(func () error {
    //do something and return error.
    return nil
  })

  //check err.
  switch err {
  case arrows.ErrBreakerOpen:
  //can get response.
  case nil:
  //work successfuly finished
  }
}
```

####2.Job/Queue :
The Job/Queue pattern for Go. You create a queue with wait list capacity and
number of workers. don't create a goroutine for a job, easily send your job to queue
when a worker was free, do next job on queue.

```go
import "github.com/SabzCity/Go-Library/arrows"
import "fmt"

func main() {
  //Create a queue with number of workers, queue size, and function for action.
	queue:=arrows.CreateQueue(20,100,func(data interface{}){
		//Do work with data packet.
		fmt.Println(data)
	})

	//Send work to queue.
	queue.SendJob(123)
	queue.SendJob("Job/Queue")
}
```

####3.Semaphore :
Get a ticket from semaphore and do anything. You can manage live
goroutine in moment with semaphore.

```go
import "github.com/SabzCity/Go-Library/arrows"
import "fmt"
import "time"

func main() {
  //Create a semaphore with max goroutine per moment.
  sem:=arrows.CreateSemaphore(2)

	go func() {
		//Begin of work.
		sem.Begin()
		fmt.Println("Work 1")
		//End of work.
		sem.End()
	}

	go func() {
		sem.Begin()
		fmt.Println("Work 2")
		sem.End()
	}()

	go func() {
		sem.Begin()
		fmt.Println("Work 3")
		sem.End()
	}()

	//wait for goroutines
	time.Sleep(2*time.Second)
}
```

####4.WorkPool :
This is very similar to [Tunny](https://github.com/Jeffail/tunny/), but
we don't use mutex in this programs and just use channels.
you can send  a work to workers and wait for result and
can send a work async and receive result in callback.

```go
import "github.com/SabzCity/Go-Library/arrows"
import "fmt"

func main() {
  pool:=arrows.CreatePool(10, func (data interface{}) interface{} {
		number:=data.(int)
		//do something with data and return result.
		return number+1
	})

	//Send work to workpool.
	//this is wait for result.
	res:=pool.SendWork(12)

	fmt.Println(res)
}
```
