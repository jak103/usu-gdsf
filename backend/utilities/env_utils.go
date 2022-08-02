package utilities

import (
	"os"
	"strconv"
)

func GetEnvAsString(name, defaultValue string) string {
	value, wasThere := os.LookupEnv(name)

	if !wasThere {
		value = defaultValue
	}

	return value
}

func GetEnvAsInt(name string, defaultValue int) int {
	var result int

	value, wasThere := os.LookupEnv(name)

	if !wasThere {
		result = defaultValue
	} else {
		result, _ = strconv.Atoi(value)
	}

	return result
}
