package main

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// //////////////////////////////////////////////////////////////

var writeMdCmd = &cobra.Command{
	Use:   "write-md",
	Short: "MD-file generation",
	RunE: func(cmd *cobra.Command, args []string) error {
		if CmdFromFilePath == "" {
			return errors.New("you must specify a file using " + cyan("--fileFrom"))
		}

		pathFileRead := CheckFilePath(CmdFromFilePath)
		if pathFileRead != FilePathValid {
			return errors.New(cyan("--fileFrom") + ":\t" + pathFileRead.String())
		}

		CmdFromFilePath, err := filepath.Abs(CmdFromFilePath)
		if err != nil {
			return errors.New(cyan("--fileFrom") + ":\t" + err.Error())
		}

		//

		if CmdToFilePath == "" {
			return errors.New("you must specify a file using " + cyan("--fileTo"))
		}

		pathFileWrite := CheckFilePath(CmdToFilePath)
		if pathFileWrite != FilePathValid && pathFileWrite != FilePathIsDir && pathFileWrite != FilePathValidDir {
			return errors.New(cyan("--fileTo") + ":\t" + pathFileWrite.String())
		}

		CmdToFilePath, err := filepath.Abs(CmdToFilePath)
		if err != nil {
			return errors.New(cyan("--fileTo") + ":\t" + err.Error())
		}

		//

		if pathFileWrite == FilePathIsDir {
			CmdToFilePath = filepath.Join(CmdToFilePath, "addressbook.md")
		}

		//

		obj, err := ParseYml(CmdFromFilePath)
		if err != nil {
			return err
		}

		return writeMD(obj, CmdToFilePath)
	},
}

func writeMD(obj *IronAdressbookObj, pathToFile string) error {
	file, err := os.Create(pathToFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// // // //

	_, err = file.WriteString(MdHeader(1, "Bregus techmap"))
	if err != nil {
		return err
	}

	// //

	file.WriteString("---\n\n")
	file.WriteString(MdHeader(2, "Motherboards"))

	for _, key := range sortMapKey(obj.Devices.Motherboards) {
		motherboard := obj.Devices.Motherboards[key]

		file.WriteString(MdHeader(3, motherboard.Name))

		//

		list := make([]string, 0)
		if motherboard.Adr != "" {
			list = append(list, fmt.Sprintf("**Adr:** `%s`", motherboard.Adr))
		}
		if len(list) > 0 {
			file.WriteString(MdList(list))
		}

		//

		file.WriteString(MdHeader(4, "ESP32 Pinouts"))
		header := []string{
			"**Pin number**",
			"**Description**",
		}
		pins := make([][]string, 0)
		for _, pin := range sortMapKeyInt(motherboard.Esp32Pinout) {
			name := motherboard.Esp32Pinout[pin]
			pins = append(pins, []string{"`" + pin + "`", name})
		}
		if len(pins) > 0 {
			file.WriteString(MdTable(header, pins))
		}

		//

		file.WriteString(MdHeader(4, "I2C devices"))
		header = []string{
			"",
			"**Description**",
		}

		for adr, i2c := range motherboard.I2C {
			file.WriteString(MdHeader(5, fmt.Sprintf("`%#02x` %s _(%s)_", adr, i2c.Name, i2c.Chip)))

			pins = make([][]string, 0)
			for _, pin := range sortMapKey(i2c.Pinout) {
				name := i2c.Pinout[pin]
				pins = append(pins, []string{"`" + pin + "`", name})
			}
			if len(pins) > 0 {
				file.WriteString(MdTable(header, pins))
			}
		}

	}

	// //

	file.WriteString("---\n\n")
	file.WriteString(MdHeader(2, "Drivers"))

	for _, driverGroup := range sortMapKey(obj.Devices.Drivers) {
		file.WriteString(MdHeader(3, driverGroup))

		for _, driver := range obj.Devices.Drivers[driverGroup] {
			file.WriteString(MdHeader(4, fmt.Sprintf("`%s` %s", stringAdr(driver.Adr), driver.Name)))

			list := make([]string, 0)
			for _, nameProtocol := range driver.Protocols {
				list = append(list, nameProtocol)
			}
			if len(list) > 0 {
				file.WriteString("**Protocols:**\n")
				file.WriteString(MdList(list))
			}
		}

	}

	// //

	file.WriteString("---\n\n")
	file.WriteString(MdHeader(2, "Modules"))

	for _, modulGroup := range sortMapKey(obj.Devices.Modules) {
		file.WriteString(MdHeader(3, modulGroup))

		for _, modul := range obj.Devices.Modules[modulGroup] {
			file.WriteString(MdHeader(4, fmt.Sprintf("`%s` %s _(%s)_", stringAdr(modul.Adr), modul.Name, modul.Chip)))

			if len(modul.I2C) > 0 {
				list := make([]string, 0)
				for _, i2cModulAdrName := range sortMapKey(modul.I2C) {
					list = append(list, fmt.Sprintf("**%s**: `%#02x`", i2cModulAdrName, modul.I2C[i2cModulAdrName]))
				}
				file.WriteString(MdList(list))
			}

			header := []string{
				"",
				"**Description**",
			}
			pins := make([][]string, 0)
			for _, pin := range sortMapKey(modul.Pinout) {
				name := modul.Pinout[pin]
				pins = append(pins, []string{"`" + pin + "`", name})
			}
			if len(pins) > 0 {
				file.WriteString(MdHeader(5, "Pinouts"))
				file.WriteString(MdTable(header, pins))
			}
		}

	}

	// // // //

	fmt.Printf("The file was created successfully. \n\tDir: %s \n\tFile: %s\n", cyan(filepath.Dir(pathToFile)), green(filepath.Base(pathToFile)))
	return err
}

// //////////////

func init() {
	writeMdCmd.Flags().StringVar(&CmdFromFilePath, "fileFrom", "", "Path to the file")
	writeMdCmd.Flags().StringVar(&CmdToFilePath, "fileTo", "", "Path to the file where the generated markdown-content will be written. You can specify a directory.")

	rootCmd.AddCommand(writeMdCmd)
}
