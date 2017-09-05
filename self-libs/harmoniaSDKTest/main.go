package main

import (
	"fmt"
	"hue-ec-swat-golang/ares/metrics"
	"hue-ec-swat-golang/harmonia/httpserver/model"
	"hue-ec-swat-golang/harmonia/sdk/scheduler"
	"log"
	//"hue-ec-swat-golang/reporter"
	"hue-ec-swat-golang/common/swatutils"
	"hue-ec-swat-golang/reporter"
	"os"
	"hue-ec-swat-golang/harmonia/sdk"
)

func main() {
	//sdk.SetHost("http://172.26.147.6:11000")
	sdk.SetHost("http://172.17.0.2:11000")
	//newAttack()
	//get()
	//stop()
	//remove()
	getAttackResult()
	//start()
	//getResult()
	//getScheduler()

}

func getResult() {
	m := getAttackResultsByDate()

	err := os.MkdirAll("test", 0700)
	if err != nil {
		log.Fatalf("fail to create a raft directory: %s", err)
	}

	ar, err := reporter.NewAresReporter("test/")
	ar.Summary(m)

	err = ar.BarchartByAPI(m)
	if err != nil {
		fmt.Println(err)
	}
	ar.BarchartByLatency(m)
}

func get() {
	//progress()
	//healthCheck()
	//isRunning()
	getScheduler()
	//getAttackResult()
	//allSchedulers()
}

func newAttack() {
	data := model.ScheduleAttacker{
		Name: "ares",
		Host: "172.26.147.6",
		//Host:       "172.17.0.2",
		Frequency: "@every 10s",
		Image:     "ares",
		Container: "ares",
		User:      "weiming",
		//PvtKeyPath: "/home/weiming/.ssh/id_rsa",
		PvtKeyPath: "/id_rsa",
		AttackParams: "-url=http://172.26.147.242:17010/ec-marketing/v2/api-docs -payload=hue-ec-marketing-payload.json -rate=100 -requests=100 -server=true -scheduler=ares",
		//AttackParams: "-url=http://172.26.158.201:17010/ec-marketing/v2/api-docs -payload=hue-ec-marketing-payload.json -rate=50 -min-req=100 -max-req=500 -step=100 -server=true -scheduler=ares",
		//AttackParams: "-url=http://172.26.158.201:17010/ec-marketing/v2/api-docs -payload=hue-ec-marketing-payload.json -rate=50 -min-req=100 -max-req=300 -step=100",
	}
	err := scheduler.NewAttack(data)
	if err != nil {
		fmt.Println(err)
	}
}

func progress() {
	res, err := scheduler.Progress("ares")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func healthCheck() {
	res, err := scheduler.Healthcheck("ares")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func isRunning() {
	res, err := scheduler.IsRunning("ares")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func getScheduler() {
	res, err := scheduler.GetScheduler("ares")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func getAttackResult() metrics.Metrics {
	res, err := scheduler.GetAttackResults("ares", "normal")
	if err != nil {
		fmt.Println(err)
	}

	m := res[swatutils.Today()][0]

	fmt.Println(m.Requests)
	fmt.Println(m.Latencies)
	fmt.Println(m.Targets)
	fmt.Println(m.Errors)
	fmt.Println(m.Success)
	fmt.Println(m.StatusCodes)
	fmt.Println(m.BytesOut)
	fmt.Println(m.BytesIn)
	return m
}

func getAttackResultsByDate() metrics.Metrics {
	res, err := scheduler.GetAttackResultsByDate("ares", swatutils.Today())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	fmt.Println(res[0].Requests)
	fmt.Println(res[0].Latencies)
	fmt.Println(res[0].Targets)
	fmt.Println(res[0].Errors)
	fmt.Println(res[0].Success)
	fmt.Println(res[0].StatusCodes)
	fmt.Println(res[0].BytesOut)
	fmt.Println(res[0].BytesIn)
	return res[0]
}

func allSchedulers() {
	res, err := scheduler.AllSchedulers()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// Set

func setAttackResult() {
	m := metrics.Metrics{
		Requests: 10,
	}
	err := scheduler.SetAttackResult(m, "ares")
	if err != nil {
		fmt.Println(err)
	}
}

func start() {
	err := scheduler.Start("ares")
	if err != nil {
		fmt.Println(err)
	}
}

func stop() {
	err := scheduler.Stop("ares")
	if err != nil {
		fmt.Println(err)
	}
}

func remove() {
	err := scheduler.Remove("ares")
	if err != nil {
		fmt.Println(err)
	}
}
