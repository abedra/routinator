package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Configuration struct {
	Version           string
	VersionString     string
	Arch              string
	InternalInterface string `json:"internal_interface"`
	ExternalInterface string `json:"external_interface"`
	DHCPInterface     string `json:"dhcp_interface"`
	DomainName        string `json:"domain_name"`
	Nameservers       []string
	NameserversString string
	Router            string
	Subnet            string
	Netmask           string
	DHCPStart         string `json:"dhcp_start"`
	DHCPEnd           string `json:"dhcp_end"`
	Broadcast         string
}

func readConfiguration(configFile string) Configuration {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c Configuration
	json.Unmarshal(data, &c)

	return c
}

func createOutputDirectories() {
	err := os.MkdirAll("out/etc", 0755)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = os.MkdirAll("out/home", 0755)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func writePfConf(config Configuration) {
	pfConf, _ := os.Create("out/etc/pf.conf")
	t, _ := template.ParseFiles("templates/pf.conf.tmpl")
	t.Execute(pfConf, config)
}

func writeRcConfLocal(config Configuration) {
	rcConfLocal, _ := os.Create("out/etc/rc.conf.local")
	t, _ := template.ParseFiles("templates/rc.conf.local.tmpl")
	t.Execute(rcConfLocal, config)
}

func writeInternalInterface(config Configuration) {
	internalInterface, _ := os.Create("out/etc/hostname." + config.InternalInterface)
	t, _ := template.ParseFiles("templates/int_hostname.tmpl")
	t.Execute(internalInterface, config)
}

func writeExternalInterface(config Configuration) {
	externalInterface, _ := os.Create("out/etc/hostname." + config.ExternalInterface)
	t, _ := template.ParseFiles("templates/ext_hostname.tmpl")
	t.Execute(externalInterface, config)
}

func writeDHCPConf(config Configuration) {
	dhcpConf, _ := os.Create("out/etc/dhcpd.conf")
	t, _ := template.ParseFiles("templates/dhcpd.conf.tmpl")
	t.Execute(dhcpConf, config)
}

func writeSysctlConf(config Configuration) {
	sysctlConf, _ := os.Create("out/etc/sysctl.conf")
	t, _ := template.ParseFiles("templates/sysctl.conf.tmpl")
	t.Execute(sysctlConf, config)
}

func writeUpdateScript(config Configuration) {
	updateScript, _ := os.Create("out/home/update")
	t, _ := template.ParseFiles("templates/update.tmpl")
	t.Execute(updateScript, config)
}

func writeRecompileKernel(config Configuration) {
	recompileKernelScript, _ := os.Create("out/home/recompile_kernel")
	t, _ := template.ParseFiles("templates/recompile_kernel.tmpl")
	t.Execute(recompileKernelScript, config)
}

func writeRecompileSystem(config Configuration) {
	recompileSystemScript, _ := os.Create("out/home/recompile_system")
	t, _ := template.ParseFiles("templates/recompile_system.tmpl")
	t.Execute(recompileSystemScript, config)
}

func main() {
	configPtr := flag.String("config", "firewall.example.json", "Path to config file")
	flag.Parse()

	config := readConfiguration(*configPtr)
	config.NameserversString = strings.Join(config.Nameservers, ", ")
	config.VersionString = strings.Join(strings.Split(config.Version, "."), "_")

	createOutputDirectories()

	writePfConf(config)
	writeRcConfLocal(config)
	writeInternalInterface(config)
	writeExternalInterface(config)
	writeDHCPConf(config)
	writeSysctlConf(config)
	writeUpdateScript(config)
	writeRecompileKernel(config)
	writeRecompileSystem(config)
}
