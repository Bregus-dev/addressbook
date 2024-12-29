package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

// //////////////////////////////////////////////////////////////

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if CmdFromFilePath == "" {
			return errors.New("you must specify a file using " + cyan("--fileFrom"))
		}

		pathFileRead := CheckFilePath(CmdFromFilePath)
		if pathFileRead != FilePathValid {
			return errors.New(cyan("--fileFrom") + ":\t" + pathFileRead.String())
		}

		obj, err := ParseYml(CmdFromFilePath)
		if err != nil {
			return err
		}

		mapMotherboards, mapModules, mapDrivers := genMaps(obj)

		//

		CmdSelectMotherboard = strings.TrimSpace(CmdSelectMotherboard)
		CmdSelectModule = strings.TrimSpace(CmdSelectModule)
		CmdSelectDriver = strings.TrimSpace(CmdSelectDriver)

		if CmdSelectMotherboard != "" {
			return readSeachMotherboard(mapMotherboards)
		}
		if CmdSelectDriver != "" {
			return readSeachDriver(mapDrivers)
		}
		if CmdSelectModule != "" {
			return readSeachModule(mapModules)
		}

		//

		return readGlobalInfo(obj)
	},
}

func readGlobalInfo(obj *IronAdressbookObj) error {
	fmt.Printf("%s:\n", "Motherboards")
	for _, key := range sortMapKey(obj.Devices.Motherboards) {
		m := obj.Devices.Motherboards[key]
		m.Print(1)
	}

	fmt.Printf("%s:\n", "Drivers")
	for key, drivers := range obj.Devices.Drivers {
		fmt.Printf("\t%s\n", key)
		for _, driver := range drivers {
			driver.Print(2)
		}
	}

	fmt.Printf("%s:\n", "Modules")
	for key, modules := range obj.Devices.Modules {
		fmt.Printf("\t%s\n", key)
		for _, module := range modules {
			module.Print(2)
		}
	}

	return nil
}

func readSeachMotherboard(mp map[string]MotherboardObj) error {
	obj, ok := mp[CmdSelectMotherboard]
	if !ok {
		fmt.Printf("%s Motherboard %s not found\n", red("[ERROR]\t"), blue(CmdSelectMotherboard))
		return nil
	}

	obj.Print(0)
	return nil
}

func readSeachModule(mp map[string]ModulObj) error {
	obj, ok := mp[CmdSelectModule]
	if !ok {
		fmt.Printf("%s Module %s not found\n", red("[ERROR]\t"), blue(CmdSelectModule))
		return nil
	}

	obj.Print(0)
	return nil
}

func readSeachDriver(mp map[string]DriverObj) error {
	obj, ok := mp[CmdSelectDriver]
	if !ok {
		fmt.Printf("%s Driver %s not found\n", red("[ERROR]\t"), blue(CmdSelectDriver))
		return nil
	}

	obj.Print(0)
	return nil
}

// //////////////

func init() {
	readCmd.Flags().StringVar(&CmdFromFilePath, "fileFrom", "", "Path to the file")

	readCmd.Flags().StringVar(&CmdSelectMotherboard, "motherboard", "", "Motherboard name. Just a complete coincidence.")
	readCmd.Flags().StringVar(&CmdSelectModule, "module", "", "Module name. Just a complete coincidence.")
	readCmd.Flags().StringVar(&CmdSelectDriver, "driver", "", "Driver name. Just a complete coincidence.")

	readCmd.Flags().StringVar(&CmdMode, "mode", "short", "Output mode. By default, outputs in shortened format. [all|short]")

	rootCmd.AddCommand(readCmd)
}
