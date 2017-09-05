package main

import (
	"bytes"
	"fmt"
	"hue-ec-swat-golang/common/sshcmd"
	"io/ioutil"
	"net"
	"os"
	"os/user"

	"golang.org/x/crypto/ssh"
)

func main() {
	client, err := sshcmd.NewSSHClient("weiming", os.Getenv("HOME")+"/.ssh/id_rsa", "localhost:22")
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.ExecuteCmd("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func flow() {
	config := usePublicKey()

	client, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}

func usePublicKey() *ssh.ClientConfig {
	key, err := getKeyFile()
	if err != nil {
		panic(err)
	}

	config := &ssh.ClientConfig{
		User: os.Getenv("LOGNAME"),
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(key),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return config
}

func usePassword() *ssh.ClientConfig {
	config := &ssh.ClientConfig{
		User: "weiming",
		Auth: []ssh.AuthMethod{
			ssh.Password("*****"),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	return config
}

func getKeyFile() (key ssh.Signer, err error) {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.ssh/id_rsa"
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}
