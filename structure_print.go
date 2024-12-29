package main

import (
	"fmt"
	"sort"
	"strings"
)

// //////////////////////////////////////////////////////////////

func (obj *ModulObj) Print(padding int) {
	if CmdMode == "short" {
		fmt.Printf("%s%s", strings.Repeat("\t", padding), green(obj.Name))

		if len(obj.Chip) > 0 {
			fmt.Printf("\t%s", blue(obj.Chip))
		}

		if len(obj.Adr) > 0 {
			fmt.Printf("\t%s", yellow(stringAdr(obj.Adr)))
		}

		if len(obj.I2C) > 0 {
			var bufAdr []string
			for _, key := range sortMapKey(obj.I2C) {
				bufAdr = append(bufAdr, fmt.Sprintf("%s:%#02x", magenta(key), obj.I2C[key]))
			}
			fmt.Printf("\t[%s]\n", strings.Join(bufAdr, ", "))
		}

	} else {
		fmt.Printf("%s%s\n", strings.Repeat("\t", padding), green(obj.Name))

		if len(obj.Chip) > 0 {
			fmt.Printf("%s\t%s: %s\n", strings.Repeat("\t", padding), cyan("Chip"), blue(obj.Chip))
		}
		if len(obj.Adr) > 0 {
			fmt.Printf("%s\t%s: %s\n", strings.Repeat("\t", padding), cyan("Adr"), yellow(stringAdr(obj.Adr)))
		}

		if len(obj.I2C) > 0 {
			fmt.Printf("%s\t%s:\n", strings.Repeat("\t", padding), cyan("I2C"))
			for _, key := range sortMapKey(obj.I2C) {
				fmt.Printf("%s\t\t%s: %#02x\n", strings.Repeat("\t", padding), magenta(key), obj.I2C[key])
			}
		}

		if len(obj.Pinout) > 0 {
			fmt.Printf("%s\t%s:\n", strings.Repeat("\t", padding), cyan("Pinout"))
			for _, key := range sortMapKey(obj.Pinout) {
				fmt.Printf("%s\t\t%s:\t%s\n", strings.Repeat("\t", padding), magenta(key), obj.Pinout[key])
			}
		}
	}
}

func (obj *DriverObj) Print(padding int) {
	if CmdMode == "short" {
		fmt.Printf("%s%s", strings.Repeat("\t", padding), green(obj.Name))

		if len(obj.Adr) > 0 {
			fmt.Printf("\t%s", yellow(stringAdr(obj.Adr)))
		}

		if len(obj.Protocols) > 0 {
			sort.Strings(obj.Protocols)
			fmt.Printf("\t[%s]\n", strings.Join(obj.Protocols, ", "))
		}

	} else {
		fmt.Printf("%s%s\n", strings.Repeat("\t", padding), green(obj.Name))

		if len(obj.Adr) > 0 {
			fmt.Printf("%s\t%s: %s\n", strings.Repeat("\t", padding), cyan("Adr"), yellow(stringAdr(obj.Adr)))
		}

		if len(obj.Protocols) > 0 {
			fmt.Printf("%s\t%s:\n", strings.Repeat("\t", padding), cyan("Protocols"))
			sort.Strings(obj.Protocols)
			for _, pt := range obj.Protocols {
				fmt.Printf("%s\t\t%s\n", strings.Repeat("\t", padding), pt)
			}
		}
	}
}

func (obj *MotherboardObj) Print(padding int) {
	if CmdMode == "short" {
		fmt.Printf("%s%s", strings.Repeat("\t", padding), green(obj.Name))

		if len(obj.Adr) > 0 {
			fmt.Printf("\t%s", yellow(obj.Adr))
		}

		if len(obj.I2C) > 0 {
			var bufI2C []string
			for key, devI2C := range obj.I2C {
				bufI2C = append(bufI2C, fmt.Sprintf("(%s)%s:%#02x", blue(devI2C.Name), magenta(devI2C.Chip), key))
			}
			fmt.Printf("\t[%s]\n", strings.Join(bufI2C, ", "))
		}

	} else {
		fmt.Printf("%s%s\n", strings.Repeat("\t", padding), green(obj.Name))

		if len(obj.Adr) > 0 {
			fmt.Printf("%s\t%s: %s\n", strings.Repeat("\t", padding), cyan("Adr"), obj.Adr)
		}

		if len(obj.Esp32Pinout) > 0 {
			fmt.Printf("%s\t%s:\n", strings.Repeat("\t", padding), cyan("Pinout"))
			for _, key := range sortMapKeyInt(obj.Esp32Pinout) {
				fmt.Printf("%s\t\t%s:\t%s\n", strings.Repeat("\t", padding), magenta(key), obj.Esp32Pinout[key])
			}
		}

		if len(obj.I2C) > 0 {
			fmt.Printf("%s\t%s:\n", strings.Repeat("\t", padding), cyan("I2C"))
			for key, devI2C := range obj.I2C {
				fmt.Printf("%s\t\t%s: %#02x %s\n", strings.Repeat("\t", padding), magenta(devI2C.Name), key, blue(devI2C.Chip))

				fmt.Printf("%s\t\t\t%s:\n", strings.Repeat("\t", padding), magenta("Pinout"))
				for _, pinKey := range sortMapKey(devI2C.Pinout) {
					fmt.Printf("%s\t\t\t\t%s:\t%s\n", strings.Repeat("\t", padding), yellow(pinKey), devI2C.Pinout[pinKey])
				}
			}
		}
	}
}
