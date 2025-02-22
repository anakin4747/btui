package cmd

import "github.com/darkhz/bluetuith/bluez"

// Version stores the version information.
var Version string

// Init parses the command-line parameters and initializes the application.
func Init(bluez *bluez.Bluez) {
	config.setup()

	parse()

	cmdOptionVersion()

	cmdOptionListAdapters(bluez)
	cmdOptionAdapter(bluez)
	cmdOptionConnectBDAddr(bluez)
	cmdOptionAdapterStates()

	validateKeybindings()
	cmdOptionGenerate()
	cmdOptionTheme()

	cmdOptionGsm()

	cmdOptionReceiveDir()
}
