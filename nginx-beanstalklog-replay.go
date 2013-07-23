// nginx-beanstalklog-replay
// Reads off requests from beanstalkd (Submitted using nginx-beanstalklog-module)
// and send to gor replay server

package main

import (
	"flag"
	"github.com/kr/beanstalk"
	"log"
	"net"
	"time"
)

const (
	defaultBeanstalkdAddr           = "127.0.0.1:11300"
	defaultBeanstalkdTube           = "nginx-log"
	defaultGorReplayAddr            = "localhost:28020"
	defaultBeanstalkdReserveTimeout = 0
	defaultFailWait                 = 3
)

type BSL2GorSettings struct {
	BeanstalkdAddr           string
	BeanstalkdTube           string
	BeanstalkdReserveTimeout int
	GorReplayAddr            string
	FailWait                 int
}

var Settings BSL2GorSettings = BSL2GorSettings{}

func init() {
	flag.StringVar(&Settings.GorReplayAddr, "r", defaultGorReplayAddr, "gor replay server address")
	flag.StringVar(&Settings.BeanstalkdAddr, "b", defaultBeanstalkdAddr, "beanstalkd server address")
	flag.StringVar(&Settings.BeanstalkdTube, "t", defaultBeanstalkdTube, "beanstalkd tube to watch")
	flag.IntVar(&Settings.BeanstalkdReserveTimeout, "s", defaultBeanstalkdReserveTimeout, "beanstalkd reserve timeout")
	flag.IntVar(&Settings.FailWait, "f", defaultFailWait, "sleep seconds after fatal errors")
}

func main() {
	flag.Parse()
	var tb *beanstalk.TubeSet
	var conn_bs *beanstalk.Conn

	rs_timeout := time.Duration(Settings.BeanstalkdReserveTimeout)
	fail_wait := time.Duration(Settings.FailWait) * time.Second

	conn_bs, e := beanstalk.Dial("tcp", Settings.BeanstalkdAddr)

	if e != nil {
		log.Fatal("failed to connected to beanstalkd", e)
	}

	tb = beanstalk.NewTubeSet(conn_bs, Settings.BeanstalkdTube)

	for {
		// reserve a job
		id, job, e := tb.Reserve(rs_timeout)

		// timeout is valid, anything else is fatal
		if cerr, ok := e.(beanstalk.ConnError); ok && cerr.Err == beanstalk.ErrTimeout {
			time.Sleep(fail_wait)
			continue
		} else if e != nil {
			log.Fatal("failed to reserve job", e)
		} else {
			log.Println("read job id", id, "size", len(job), "bytes")
		}

		// connect to the gor replay server
		conn_gr, e := net.Dial("tcp", Settings.GorReplayAddr)

		if e != nil {
			log.Fatal("failed to connected to gor replay server", e)
			time.Sleep(fail_wait)
		}

		// write to gor replay server
		w, e := conn_gr.Write(job)

		if e != nil {
			log.Fatal("failed to write to", Settings.GorReplayAddr, "error", e)
		} else {
			log.Println("wrote", w, "bytes to", Settings.GorReplayAddr)
		}

		// close connection to gor replay server
		conn_gr.Close()

		// delete the job
		e = conn_bs.Delete(id)

		if e != nil {
			log.Println("failed to delete job id", id, "error", e)
		}
	}
}
