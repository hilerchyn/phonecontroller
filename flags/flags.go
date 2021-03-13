package flags

import "flag"

var SerialNo = flag.String("s", "", "Device's serial number.")
var Action = flag.String("a", "swipe", "Action type.")
var Source = flag.String("src", "touchscreen", "input source")
var Interval = flag.Int("intval", 12, "action interval ")

var FromX = flag.Int("x", 850, "from x point.")
var FromY = flag.Int("y", 1400, "from y point.")
var ToX = flag.Int("2x", 850, "to x point.")
var ToY = flag.Int("2y", 300, "to y point.")