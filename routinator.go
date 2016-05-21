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
	t, err := template.ParseFiles(inputPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	t.Execute(outFile, config)
}

func writeConfigs(config Configuration, templateDir string) {
	writeConfig(config, templateDir + "/pf.conf.tmpl", "out/etc/pf.conf")
	writeConfig(config, templateDir + "/rc.conf.local.tmpl", "out/etc/rc.conf.local")
	writeConfig(config, templateDir + "/ext_hostname.tmpl", "out/etc/hostname."+config.ExternalInterface)
	writeConfig(config, templateDir + "/int_hostname.tmpl", "out/etc/hostname."+config.InternalInterface)
	writeConfig(config, templateDir + "/dhcpd.conf.tmpl", "out/etc/dhcpd.conf")
	writeConfig(config, templateDir + "/sysctl.conf.tmpl", "out/etc/sysctl.conf")
	writeConfig(config, templateDir + "/update.tmpl", "out/home/bin/update")
	writeConfig(config, templateDir + "/recompile_kernel.tmpl", "out/home/bin/recompile_kernel")
	writeConfig(config, templateDir + "/recompile_system.tmpl", "out/home/bin/recompile_system")
	writeConfig(config, templateDir + "/.profile.tmpl", "out/home/.profile")
}

func move(in, out string) {
	err := os.Rename(in, out)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func moveConfigs(config Configuration) {
	move("out/home/.profile", "/root/.profile")
	move("out/home/bin", "/root/bin")
	move("out/etc/pf.conf", "/etc/pf.conf")
	move("out/etc/rc.conf.local", "/etc/rc.conf.local")
	move("out/etc/hostname."+config.ExternalInterface, "/etc/hostname."+config.ExternalInterface)
	move("out/etc/hostname."+config.InternalInterface, "/etc/hostname."+config.InternalInterface)
	move("out/etc/dhcpd.conf", "/etc/dhcpd.conf")
	move("out/etc/sysctl.conf", "/etc/sysctl.conf")
}

func main() {
	configPtr := flag.String("config", "firewall.example.json", "Path to config file")
	templatePtr := flag.String("templates", "templates", "Path to templates")
	flag.Parse()

	config := readConfiguration(*configPtr)
	config.NameserversString = strings.Join(config.Nameservers, ", ")
	config.VersionString = strings.Join(strings.Split(config.Version, "."), "_")

	createOutputDirectories()
	writeConfigs(config, *templatePtr)
	moveConfigs(config)
}
