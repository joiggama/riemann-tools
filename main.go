package main

import (
	"flag"
  "fmt"
  "time"

  "github.com/amir/raidman"
  "github.com/joiggama/riemann-tools/health"
)

var (
	INTERVAL = flag.Duration("interval", time.Second, "Agent refresh interval (Default: 1s)")
	HOST     = flag.String("host", "127.0.0.1", "Riemann host (default: 127.0.0.1)")
	PORT     = flag.String("port", "5555", "Riemann port (default: 5555")
)

func main() {
  flag.Parse()

  tick := time.NewTicker(*INTERVAL)

  for {
    <-tick.C

    cpu_event_err := Notify(&raidman.Event{
      State: "ok",
      Service: "cpu",
      Metric: health.CPU(),
      Ttl: 10,
    })

    if cpu_event_err != nil {
      fmt.Println(cpu_event_err)
    }

    mem_event_err := Notify(&raidman.Event{
      State: "ok",
      Service: "memory",
      Metric: health.Memory(),
      Ttl: 10,
    })

    if mem_event_err != nil {
      fmt.Println(mem_event_err)
    }

  }
}

func Notify(event *raidman.Event) error {
  conn, err := raidman.Dial("tcp", *HOST + ":" + *PORT)

  if err != nil {
    return err
  } else {

    send_err := conn.Send(event)

    if send_err != nil {
      return send_err
    } else {
      conn.Close()
      return nil
    }

  }
}
