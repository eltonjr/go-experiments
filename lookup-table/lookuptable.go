package lookuptable

import "slices"

// option 1
// using switch/case for the first level of the lookup table
// and iterating through the array for the second level

var (
	profilingsGo     = []string{"cpu", "heap", "allocs", "block", "trace"}
	profilingsJVM    = []string{"cpu", "allocs", "block"}
	profilingsPython = []string{"cpu", "heap", "thread"}
	profilingsNodejs = []string{"cpu", "heap"}
)

func profilingTypesForTech(tech string) []string {
	switch tech {
	case "go":
		return profilingsGo
	case "java", "kotlin":
		return profilingsJVM
	case "python":
		return profilingsPython
	case "nodejs":
		return profilingsNodejs
	default:
		return []string{}
	}
}

func ValidProfilingType(tech string, profilingType string) bool {
	allowed := profilingTypesForTech(tech)
	return slices.Contains(allowed, profilingType)
}

// option 2
// using a map for the first level of the lookup table
// and iterating through the array for the second level

var lookup = map[string][]string{
	"go":     {"cpu", "heap", "allocs", "block", "trace"},
	"java":   {"cpu", "allocs", "block"},
	"kotlin": {"cpu", "allocs", "block"},
	"python": {"cpu", "heap", "thread"},
	"nodejs": {"cpu", "heap"},
}

func ValidProfilingTypeLookup(tech string, profilingType string) bool {
	allowed := lookup[tech]
	return slices.Contains(allowed, profilingType)
}

// option 3
// using a map for the first level of the lookup table
// and also a map for the second level

var nestedLookup = map[string]map[string]bool{
	"go": {
		"cpu":    true,
		"heap":   true,
		"allocs": true,
		"block":  true,
		"trace":  true,
	},
	"java": {
		"cpu":    true,
		"allocs": true,
		"block":  true,
	},
	"kotlin": {
		"cpu":    true,
		"allocs": true,
		"block":  true,
	},
	"python": {
		"cpu":    true,
		"heap":   true,
		"thread": true,
	},
	"nodejs": {
		"cpu":  true,
		"heap": true,
	},
}

func ValidProfilingTypeNestedLookup(tech string, profilingType string) bool {
	allowed := nestedLookup[tech]
	return allowed[profilingType]
}

// option 4
// using only arrays, in a slightly inconvenient way
// just for benchmark purposes

var allprofilings = [][]string{
	{"go", "cpu", "heap", "allocs", "block", "trace"},
	{"java", "cpu", "allocs", "block"},
	{"kotlin", "cpu", "allocs", "block"},
	{"python", "cpu", "heap", "thread"},
	{"nodejs", "cpu", "heap"},
}

func ValidProfilingTypeArrays(tech string, profilingType string) bool {
	for _, allowed := range allprofilings {
		if allowed[0] == tech {
			for _, profiling := range allowed[1:] {
				if profiling == profilingType {
					return true
				}
			}
			return false
		}
	}
	return false
}
