package main

// //////////////////////////////////////////////////////////////

func genMaps(obj *IronAdressbookObj) (map[string]MotherboardObj, map[string]ModulObj, map[string]DriverObj) {
	mapMotherboards := make(map[string]MotherboardObj)
	mapModules := make(map[string]ModulObj)
	mapDrivers := make(map[string]DriverObj)

	for _, motherboard := range obj.Devices.Motherboards {
		mapMotherboards[motherboard.Name] = motherboard
	}
	for _, modules := range obj.Devices.Modules {
		for _, module := range modules {
			mapModules[module.Name] = module
		}
	}

	for _, modules := range obj.Devices.Drivers {
		for _, dp := range modules {
			mapDrivers[dp.Name] = dp
		}
	}

	return mapMotherboards, mapModules, mapDrivers
}
