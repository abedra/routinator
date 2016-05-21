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
	DHCPStart         string `json:"dhcp_start"`
	DHCPEnd           string `json:"dhcp_end"`
	DomainName        string `json:"domain_name"`
	Nameservers       []string
	NameserversString string
	Router            string
	Subnet            string
	Netmask           string
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

	err = os.MkdirAll("out/home/bin", 0755)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func writeConfig(config Configuration, inputPath string, outputPath string) {
	outFile, _ := os.Create(outputPath)
	t, _ := template.ParseFiles(inputPath)
	t.Execute(outFile, config)
}

func main() {
	configPtr := flag.String("config", "firewall.example.json", "Path to config file")
	flag.Parse()

	config := readConfiguration(*configPtr)
	config.NameserversString = strings.Join(config.Nameservers, ", ")
	config.VersionString = strings.Join(strings.Split(config.Version, "."), "_")

	createOutputDirectories()

	writeConfig(config, "templates/pf.conf.tmpl", "out/etc/pf.conf")
	writeConfig(config, "templates/rc.conf.local.tmpl", "out/etc/rc.conf.local")
	writeConfig(config, "templates/ext_hostname.tmpl", "out/etc/hostname."+config.ExternalInterface)
	writeConfig(config, "templates/int_hostname.tmpl", "out/etc/hostname."+config.InternalInterface)
	writeConfig(config, "templates/dhcpd.conf.tmpl", "out/etc/dhcpd.conf")
	writeConfig(config, "templates/sysctl.conf.tmpl", "out/etc/sysctl.conf")
	writeConfig(config, "templates/update.tmpl", "out/home/bin/update")
	writeConfig(config, "templates/recompile_kernel.tmpl", "out/home/bin/recompile_kernel")
	writeConfig(config, "templates/recompile_system.tmpl", "out/home/bin/recompile_system")
	writeConfig(config, "templates/.profile.tmpl", "out/home/.profile")
}
