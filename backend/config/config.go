package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	
	"github.com/jak103/usu-gdsf/log"
)

// Environment
var DbType string
var MongoUri string
var FirestoreProjectId string

// Authentication
var TokenHashingKey string
var AccessTokenLifetimeMins int64
var RefreshTokenLifetimeDays int64

func init() {
	log.Info("Running configuration init")
	
	envFileMap := mapVariablesFromEnvFile()
	
	DbType = strings.ToLower(getEnvVarString("DB_TYPE", envFileMap))
	
	if DbType == "mongo" {
		MongoUri = getEnvVarString("MONGO_URI", envFileMap)
	} else if DbType == "firestore" {
		FirestoreProjectId = getEnvVarString("FIRESTORE_PROJECT_ID", envFileMap)
	}
	
	TokenHashingKey = getEnvVarString("TOKEN_HASHING_KEY", envFileMap)
	AccessTokenLifetimeMins = getEnvVarInt64("ACCESS_TOKEN_LIFETIME_MINS", envFileMap)
	RefreshTokenLifetimeDays = getEnvVarInt64("REFRESH_TOKEN_LIFETIME_DAYS", envFileMap)
}

func mapVariablesFromEnvFile() map[string]string {
	envMap := make(map[string]string)

	envFilePath, err := findEnvFilePath(".")

	if err != nil {
		return envMap
	}
	
	envFile, err := os.Open(envFilePath)

	if err != nil {
		return envMap
	}

	defer envFile.Close()

	envFileScanner := bufio.NewScanner(envFile)
	for envFileScanner.Scan() {
		line := strings.TrimSpace(envFileScanner.Text())

		if len(line) == 0 || line[0:1] == "#" {
			continue
		}
		
		commentSplit := strings.SplitN(line, "#", 2)
		keyValString := commentSplit[0]

		if !strings.Contains(keyValString, "=") {
			continue
		}

		splitKeyVal := strings.SplitN(keyValString, "=", 2)
		key := splitKeyVal[0]
		val := splitKeyVal[1]

		envMap[key] = val
	}

	return envMap
}

func findEnvFilePath(start string) (string, error) {
	if path, err := filepath.Abs(start); err != nil || path == filepath.FromSlash("/") {
		return "", errors.New("Found no .env file")
	}
	
	if _, err := os.Stat(filepath.Join(start, ".env")); err != nil {
		return findEnvFilePath(filepath.Join(start, ".."))
	}

	return filepath.Join(start, ".env"), nil
}


func getEnvVarString(key string, envFileMap map[string]string) string {
	envVar := os.Getenv(key)

	if envVar == "" {
		if value, isInMap := envFileMap[key]; isInMap {
			envVar = value
		} else {
			panic("Environment variable " + key + " not set!")
		}
	}

	return strings.TrimSpace(envVar)
}

func getEnvVarInt64(key string, envFileMap map[string]string) int64 {
	envVarStr := getEnvVarString(key, envFileMap)
	envVarInt, err := strconv.Atoi(envVarStr)

	if err != nil {
		panic("Environment variable " + key + " must be an integer")
	}

	return int64(envVarInt)
}
