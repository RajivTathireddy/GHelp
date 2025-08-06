package structure

import (
	"flag"
	"strconv"
)

func getStringFlag(flags flag.FlagSet, name string) string {
	if f := flags.Lookup(name); f != nil {
		return f.Value.String()
	}
	return ""
}

func getBoolFlag(flags flag.FlagSet, name string) bool {
	if f := flags.Lookup(name); f != nil {
		if val, err := strconv.ParseBool(f.Value.String()); err == nil {
			return val
		}
	}
	return false
}

// func getIntFlag(flags flag.FlagSet, name string) int {
// 	if f := flags.Lookup(name); f != nil {
// 		if val, err := strconv.Atoi(f.Value.String()); err == nil {
// 			return val
// 		}
// 	}
// 	return 0
// }

// func getInt64Flag(flags flag.FlagSet, name string) int64 {
// 	if f := flags.Lookup(name); f != nil {
// 		if val, err := strconv.ParseInt(f.Value.String(), 10, 64); err == nil {
// 			return val
// 		}
// 	}
// 	return 0
// }
