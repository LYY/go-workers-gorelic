go-workers-gorelic
=======

NewRelic middleware for go-workers.

## Usage

~~~ go
import(
  "github.com/jrallison/go-workers"
  "github.com/LYY/go-workers-gorelic"
)

func main(){
  workers.Configure(map[string]string{
    // location of redis instance
    "server":  "localhost:6379",
    // instance of the database
    "database":  "0",
    // number of connections to keep open with redis
    "pool":    "30",
    // unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
    "process": "1",
  })


  gorelic.InitNewrelicAgent("YOUR_NEWRELIC_LICENSE_KEY", "YOUR_APPLICATION_NAME", true)
  workers.Middleware.Append(&GorelicMiddleware{})

  workers.Run()
}
~~~

## Authors

* [LYY](https://github.com/LYY)
