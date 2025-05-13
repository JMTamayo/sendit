package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

func getFromEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Environment variable %s not found", key))
	}

	return value
}

type Logger struct {
	Logger log.Logger
}

func (l *Logger) Info(msg string) {
	l.Logger.Log("msg", msg)
}

func (l *Logger) Debug(msg string) {
	l.Logger.Log("msg", msg)
}

func (l *Logger) Error(msg string) {
	l.Logger.Log("msg", msg)
}

func buildLogger(logLevel string) *Logger {
	var logLevelOptions level.Option
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		logLevelOptions = level.AllowDebug()
	case "INFO":
		logLevelOptions = level.AllowInfo()
	case "ERROR":
		logLevelOptions = level.AllowError()
	default:
		logLevelOptions = level.AllowInfo()
	}

	iow := log.NewSyncWriter(os.Stdout)

	logger := log.NewLogfmtLogger(iow)
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
	logger = level.NewFilter(logger, logLevelOptions)

	return &Logger{
		Logger: logger,
	}
}

type Config struct {
	serviceName string
	servicePort int
	logLevel    string

	redisUsername string
	redisPassword string

	redisHost string
	redisPort int
	redisDB   int

	streamNameEmailQueue string
	keyNameData          string
}

func (c *Config) GetServiceName() string {
	return c.serviceName
}

func (c *Config) GetServiceAddress() string {
	return fmt.Sprintf(":%d", c.servicePort)
}

func (c *Config) GetLogLevel() string {
	return c.logLevel
}

func (c *Config) GetRedisUsername() string {
	return c.redisUsername
}

func (c *Config) GetRedisPassword() string {
	return c.redisPassword
}

func (c *Config) GetRedisHost() string {
	return c.redisHost
}

func (c *Config) GetRedisPort() int {
	return c.redisPort
}

func (c *Config) GetRedisDB() int {
	return c.redisDB
}

func (c *Config) GetStreamNameEmailQueue() string {
	return c.streamNameEmailQueue
}

func (c *Config) GetKeyNameData() string {
	return c.keyNameData
}

func buildConfig() *Config {
	serviceName := "Sendit Server"
	servicePort := 8000
	logLevel := getFromEnv("LOG_LEVEL")

	redisUsername := getFromEnv("REDIS_USERNAME")
	redisPassword := getFromEnv("REDIS_PASSWORD")

	redisHost := getFromEnv("REDIS_HOST")
	redisPort, err := strconv.Atoi(getFromEnv("REDIS_PORT"))
	if err != nil {
		panic(fmt.Sprintf("REDIS_PORT is not a valid number: %s", err))
	}
	redisDB, err := strconv.Atoi(getFromEnv("REDIS_DB"))
	if err != nil {
		panic(fmt.Sprintf("REDIS_DB is not a valid number: %s", err))
	}
	streamNameEmailQueue := getFromEnv("STREAM_NAME_EMAIL_QUEUE")

	keyNameData := getFromEnv("KEY_EMAIL_DATA")

	return &Config{
		serviceName:          serviceName,
		servicePort:          servicePort,
		logLevel:             logLevel,
		redisUsername:        redisUsername,
		redisPassword:        redisPassword,
		redisHost:            redisHost,
		redisPort:            redisPort,
		redisDB:              redisDB,
		streamNameEmailQueue: streamNameEmailQueue,
		keyNameData:          keyNameData,
	}
}

var Conf *Config = buildConfig()

var Log *Logger = buildLogger(Conf.GetLogLevel())
