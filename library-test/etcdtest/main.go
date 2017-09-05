package main

import (
	"context"
	"fmt"
	"log"

	"hue-ec-swat-golang/common/etcdctrl"

	"github.com/coreos/etcd/client"
)

const (
	NodeID   = "/NodeID"
	HostName = "/HostName"
	JoinAddr = "/JoinAddr"
)

func main() {
	ctx := context.Background()
	kapi, err := etcdctrl.Construct("http://127.0.0.1:2389")
	if err != nil {
		fmt.Println(err)
	}

	resp := make(chan *client.Response)
	kapi.WatchUpdate(ctx, "/test", resp)


	//kapi := construct()
	//set(kapi)
	//get(kapi)

	//watch(kapi)
}

func construct() client.KeysAPI {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2389"},
		Transport: client.DefaultTransport,
		//HeaderTimeoutPerRequest: time.Second,
	}

	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	kapi := client.NewKeysAPI(c)
	return kapi
}

func set(kapi client.KeysAPI) {
	resp, err := kapi.Set(context.Background(), "/nodeID", "", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("set id done:", resp)
	}
}

func get(kapi client.KeysAPI) {
	// get "/foo" key's value
	fmt.Print("Getting '/a' key value")
	resp, err := kapi.Get(context.Background(), JoinAddr, nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		fmt.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		fmt.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

func watch(kapi client.KeysAPI) {
	ctx := context.Background()
	opts := client.WatcherOptions{
		AfterIndex: 0,
		Recursive:  false,
	}
	watcher := kapi.Watcher("/foo", &opts)
	for {
		resp, err := watcher.Next(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Node.Value)
	}
}
