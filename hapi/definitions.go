package hapi

type Component uint8

// Components defined from https://www.home-assistant.io/integrations/mqtt/#configuration-via-mqtt-discovery
const (
	AlarmControlPanel Component = iota // 0
	BinarySensor                       // 1
	Button                             // 2
	Camera                             // 3
	Cover                              // 4
	DeviceTracker                      // 5
	DeviceTrigger                      // 6
	Event                              // 7
	Fan                                // 8
	Humidifier                         // 9
	Image                              // 10
	ClimateHVAC                        // 11
	Lawnmower                          // 12
	Light                              // 13
	Lock                               // 14
	Notify                             // 15
	Number                             // 16
	Scene                              // 17
	Select                             // 18
	Sensor                             // 19
	Siren                              // 20
	Switch                             // 21
	Update                             // 22
	TagScanner                         // 23
	Text                               // 24
	Vacuum                             // 25
	Valve                              // 26
	WaterHeater                        // 27
)
