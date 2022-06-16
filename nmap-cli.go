package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	ipports map[string][]string
)

var inputFile = flag.String("f", "hostip.txt", "输入的文件")
var outputFile = flag.String("o", "", "输出的文件")

func main() {
	flag.Parse()
	//读取文件
	contents := readFile(*inputFile)
	//fmt.Println(contents)

	//处理成ip集合
	ips := handleIp(contents)
	//fmt.Println(ips)

	//处理成map结合ip[port port]
	ipports = make(map[string][]string)
	handleInput(ips, contents)
	// for ip := range ips {
	// 	fmt.Println(ipports[ip])
	// }

	for ip := range ips {
		args := getArgs(ip)
		nmap(args)
	}
}

//调用nmap
func nmap(args []string) {
	// cmd := exec.Command(args[0], args[1:]...)
	// fmt.Println(cmd)
	// cmd.Stdout = os.Stdout
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	cmd := exec.Command(args[0], args[1:]...)
	fmt.Println(cmd)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error(), stderr.String())
	} else {
		fmt.Println(out.String())
	}

}

func getArgs(ip string) (results []string) {

	var args []string
	portsStr := strings.Join(ipports[ip], ",")
	//nmap ip -sV -n -p ports
	args = append(args, "nmap")
	args = append(args, "-n")
	args = append(args, "-sV")
	args = append(args, ip)
	args = append(args, "-p", portsStr)
	args = append(args, "-oN", *outputFile)
	return args
}

func handleInput(ips map[string]struct{}, contents string) {
	for ip := range ips {
		//fmt.Println(ip)
		ports := findports(ip, contents)
		//fmt.Println(ports)
		ipports[ip] = ports
	}

}

//根据ip找ports
func findports(ip string, contents string) (results []string) {
	ports := map[string]struct{}{}
	list := strings.Split(contents, "\n")
	for _, item := range list {
		ipport := strings.Split(item, ":")
		if ip == ipport[0] {
			port := ipport[1]
			ports[port] = struct{}{}
		}
	}
	var result []string
	for port := range ports {
		result = append(result, port)
	}
	return result
}

func handleIp(contents string) (results map[string]struct{}) {
	ips := map[string]struct{}{}
	list := strings.Split(contents, "\n")
	for _, item := range list {

		ipport := strings.Split(item, ":")
		ip := ipport[0]
		ips[ip] = struct{}{}
	}
	return ips

}

func readFile(inputFile string) (contents string) {
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	return string(content)
}
