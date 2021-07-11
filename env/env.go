package env

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	configo "github.com/jxsl13/simple-configo"
)

// Read reads .env files that are passed to this function and returns a map of key values
// .env files have the same syntax as environment variables KEY=VALUE
func Read(filenames ...string) (map[string]string, error) {
	return godotenv.Read(filenames...)
}

// GetEnv returns a map of OS environment variables
func GetEnv() map[string]string {
	pairs := os.Environ()
	env := make(map[string]string, len(pairs))
	for _, pair := range pairs {
		keyPairs := strings.SplitN(pair, "=", 2)
		if len(keyPairs) == 2 && keyPairs[1] != "" {
			// do not care about variables that do not have a value
			env[keyPairs[0]] = keyPairs[1]
		}
	}
	return env
}

// Get first tries to read the passed files and then
// overrides the .env file parameters with the environment variables
// in memory.
// this is useful in order to have a base configuration that is within a file
// that can contain like a common password for every econ, but must override its
// econ address (IP:Port) with an environment variable.
// The passed filenames can be omitted, but may be used to pass a list of files.
func Get(filenames ...string) map[string]string {
	fileEnv, err := Read(filenames...)
	if err != nil {
		// Could not fetch those files, thus simply
		// returning environment varibles
		return GetEnv()
	}

	env := GetEnv()
	for key, value := range env {
		// do not override values with empty values
		if fileEnv[key] != "" && value == "" {
			continue
		}
		fileEnv[key] = value
	}
	return fileEnv
}

// Parse uses my simple-configo package that allows to easily parse
// environment maps and do proper value conversions as well as setting your
// configuration's struct fields after parsing the values.
// All you do is define a Configuration struct, add any of the common struct
// field values, implement the Options() configo.Options function and you are ready to parse
// any arbitrary environment format like string, regex match, integer, bool, time ranges, etc.
// filenames is an optional parameter, if it is left empty, the underlying library will try
// to open the .env file in the same directory first and then parse those values
// afterwards the environment variables override any value from that file.
func Parse(filename string, cfg configo.Config) error {
	env := Get(filename)
	return configo.Parse(env, cfg)
}
