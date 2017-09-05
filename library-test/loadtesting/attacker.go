package loadtesting

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"

	"io"
	"io/ioutil"
)

// Attacker is an attack executor which wraps an http.Client
type Attacker struct {
	dialer      *net.Dialer
	client      http.Client
	stopch      chan struct{}
	workers     uint64
	redirects   int
	connections uint64
}

const (
	// DefaultTimeout is the default amount of time an Attacker waits for a request
	// before it times out.
	DefaultTimeout = 30 * time.Second
	// DefaultMaxIdleConnsPerHost is the default amount of max open idle connections per
	// target host.
	DefaultMaxIdleConnsPerHost = 2
	// DefaultWorkers is the default initial number of workers used to carry an attack.
	DefaultWorkers = 1000
	// NoFollow is the value when redirects are not followed but marked successful
	NoFollow = -1
)

var (
	// DefaultLocalAddr is the default local IP address an Attacker uses.
	DefaultLocalAddr = net.IPAddr{IP: net.IPv4zero}
	// DefaultTLSConfig is the default tls.Config an Attacker uses.
	DefaultTLSConfig = &tls.Config{InsecureSkipVerify: true}
)

// NewAttacker returns a new Attacker with default options which are overridden
// by the optionally provided opts.
func NewAttacker(opts ...func(*Attacker)) *Attacker {
	a := &Attacker{
		stopch:  make(chan struct{}),
		workers: DefaultWorkers,
	}
	a.dialer = &net.Dialer{
		LocalAddr: &net.TCPAddr{IP: DefaultLocalAddr.IP, Zone: DefaultLocalAddr.Zone},
		KeepAlive: 30 * time.Second,
		Timeout:   DefaultTimeout,
	}
	a.client = http.Client{
		Transport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           a.dialer.DialContext,
			ResponseHeaderTimeout: DefaultTimeout,
			TLSClientConfig:       DefaultTLSConfig,
			TLSHandshakeTimeout:   10 * time.Second,
			MaxIdleConnsPerHost:   DefaultMaxIdleConnsPerHost,
		},
	}

	for _, opt := range opts {
		opt(a)
	}

	return a
}

func (a *Attacker) Attack2(tr Targeter, rate uint64, du time.Duration) <-chan *Result {
	fmt.Println(rate, du)
	results := make(chan *Result)
	done := make(chan int)

	requests := 500

	go a.attack2(tr, results)
	go func() {
		for {
			select {
			case res := <-results:
				fmt.Println("get results: ", res)
				fmt.Println("requests: ", requests)
				go a.attack2(tr, results)
				requests--
				if requests == 400 {
					done <- 0
				}
			case <-done:
				fmt.Println("Done")
				return
			case <-a.stopch:
				return
			}
		}
	}()
	time.Sleep(60 * time.Second)
	return results
}

func (a *Attacker) attack2(tr Targeter, results chan<- *Result) {
	results <- a.hit2(tr)
}

func (a *Attacker) hit2(tr Targeter) *Result {
	var (
		res = Result{}
		tgt Target
		err error
	)

	if err = tr(&tgt); err != nil {
		a.Stop()
		return &res
	}

	req, err := tgt.Request()
	if err != nil {
		return &res
	}

	r, err := a.client.Do(req)

	if err != nil {
		return &res
	}

	tran := a.client.Transport.(*http.Transport)

	if res.Code = uint16(r.StatusCode); res.Code < 200 || res.Code >= 400 {
		res.Error = r.Status
	}

	io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()

	fmt.Printf("client info: %v\n", tran)

	return &res
}

// Connections returns a functional option which sets the number of maximum idle
// open connections per target host.
func Connections(n int) func(*Attacker) {
	return func(a *Attacker) {
		tr := a.client.Transport.(*http.Transport)
		tr.MaxIdleConnsPerHost = n
	}
}

// Stop stops the current attack.
func (a *Attacker) Stop() {
	select {
	case <-a.stopch:
		return
	default:
		close(a.stopch)
	}
}
