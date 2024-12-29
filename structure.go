package main

// //////////////////////////////////////////////////////////////

type IronAdressbookObj struct {
	Devices DevicesObj `json:"devices"`
}

type DevicesObj struct {
	Motherboards map[string]MotherboardObj `json:"motherboards"`
	Drivers      map[string][]DriverObj    `json:"drivers"`
	Modules      map[string][]ModulObj     `json:"modules"`
}

type DriverObj struct {
	Name      string   `json:"name"`
	Adr       []bool   `json:"adr"`
	Protocols []string `json:"protocols"`
}
type ModulObj struct {
	Name   string            `json:"name"`
	Chip   string            `json:"chip"`
	Adr    []bool            `json:"adr"`
	Pinout map[string]string `json:"pinout"`
	I2C    map[string]byte   `json:"i2c"`
}

type MotherboardObj struct {
	Adr         string                `json:"adr,omitempty"`
	Esp32Pinout map[string]string     `json:"esp32pinout"`
	I2C         map[byte]I2CDeviceObj `json:"i2c"`
	Name        string                `json:"name"`
}

type I2CDeviceObj struct {
	Chip   string            `json:"chip"`
	Name   string            `json:"name"`
	Pinout map[string]string `json:"pinout,omitempty"`
}
