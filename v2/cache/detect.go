package cache

var Detect = []string{}
var NotDetect = []string{}

func PreText() map[string]float64 {
	var cv = make(map[string]float64)
	for _, s := range Detect {
		cv[s] = 1
	}
	for _, s := range NotDetect {
		cv[s] = 0
	}
	return cv
}
