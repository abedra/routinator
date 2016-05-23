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

type OperatingSystem struct {
	Version string
	Arch    string
}

type NetworkInterfaces struct {
	Internal string
	External string
}

type DHCPConfiguration struct {
	Interface   string
	Start       string
	End         string
	DomainName  string `json:"domain_name"`
	Nameservers []string
}

type Configuration struct {
	OS                OperatingSystem   `json:"os"`
	Interfaces        NetworkInterfaces `json:"interfaces"`
	DHCP              DHCPConfiguration `json:"dhcp"`
	Router            string
	Subnet            string
	Netmask           string
	Broadcast         string
	VersionString     string
	NameserversString string
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
	writeConfig(config, templateDir+"/pf.conf.tmpl", "out/etc/pf.conf")
	writeConfig(config, templateDir+"/rc.conf.local.tmpl", "out/etc/rc.conf.local")
	writeConfig(config, templateDir+"/ext_hostname.tmpl", "out/etc/hostname."+config.Interfaces.External)
	writeConfig(config, templateDir+"/int_hostname.tmpl", "out/etc/hostname."+config.Interfaces.Internal)
	writeConfig(config, templateDir+"/dhcpd.conf.tmpl", "out/etc/dhcpd.conf")
	writeConfig(config, templateDir+"/sysctl.conf.tmpl", "out/etc/sysctl.conf")
	writeConfig(config, templateDir+"/update.tmpl", "out/home/bin/update")
	writeConfig(config, templateDir+"/recompile_kernel.tmpl", "out/home/bin/recompile_kernel")
	writeConfig(config, templateDir+"/recompile_system.tmpl", "out/home/bin/recompile_system")
	writeConfig(config, templateDir+"/.profile.tmpl", "out/home/.profile")
}

func move(in, out string) {
	err := os.Rename(in, out)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func makeExecutable(file string) {
	err := os.Chmod(file, 0755)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func moveConfigs(config Configuration) {
	move("out/home/.profile", "/root/.profile")
	move("out/home/bin", "/root/bin")
	makeExecutable("/root/bin/update")
	makeExecutable("/root/bin/recompile_kernel")
	makeExecutable("/root/bin/recompile_system")
	move("out/etc/pf.conf", "/etc/pf.conf")
	move("out/etc/rc.conf.local", "/etc/rc.conf.local")
	move("out/etc/hostname."+config.Interfaces.External, "/etc/hostname."+config.Interfaces.External)
	move("out/etc/hostname."+config.Interfaces.Internal, "/etc/hostname."+config.Interfaces.Internal)
	move("out/etc/dhcpd.conf", "/etc/dhcpd.conf")
	move("out/etc/sysctl.conf", "/etc/sysctl.conf")
}

func main() {
	configPtr := flag.String("config", "firewall.example.json", "Path to config file")
	templatePtr := flag.String("templates", "templates", "Path to templates")
	skipInstall := flag.Bool("skip-install", false, "Skip installation of config files")
	flag.Parse()

	config := readConfiguration(*configPtr)
	config.NameserversString = strings.Join(config.DHCP.Nameservers, ", ")
	config.VersionString = strings.Join(strings.Split(config.OS.Version, "."), "_")

	createOutputDirectories()
	writeConfigs(config, *templatePtr)
	if !*skipInstall {
		moveConfigs(config)
	}
}
