package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	saltmasterlessprovisioner "github.com/hashicorp/packer-plugin-salt/provisioner/salt-masterless"
	"github.com/hashicorp/packer-plugin-salt/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterProvisioner("masterless", new(saltmasterlessprovisioner.Provisioner))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
