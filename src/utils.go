package main

import "os"

//GetEnv returns the value of a key in the environment or a default
func GetEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
