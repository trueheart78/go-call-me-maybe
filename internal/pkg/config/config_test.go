package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sid = "123456"
var authToken = "token"
var emergencyPhone = "+1234567890"
var nonEmergentPhone = "+2345678901"
var asleepPhone = "+3456789012"
var outboundPhone = "+4567890123"
var emergencyURL = "http://123.com/emergency"
var asleepURL = "http://123.com/asleep"
var redisURL = "sample.redis.com:123"
var redisPassword = "password"
var redisChannelEmergency = "channela"
var redisChannelNonEmergent = "channelb"

func TestConfigSetup(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}

	assert.False(cfg.Valid(), "should not be valid")

	setupEnvs()
	assert.Equal(sid, cfg.SID(), "they should be equal")
	assert.Equal(authToken, cfg.AuthToken(), "they should be equal")
	assert.Equal(emergencyPhone, cfg.EmergencyPhone(), "they should be equal")
	assert.Equal(nonEmergentPhone, cfg.NonEmergentPhone(), "they should be equal")
	assert.Equal(asleepPhone, cfg.AsleepPhone(), "they should be equal")
	assert.Equal(outboundPhone, cfg.OutboundPhone(), "they should be equal")
	assert.True(cfg.Valid(), "should be valid")

	clearExtraPhones()
	assert.Equal(emergencyPhone, cfg.NonEmergentPhone(), "should default to the emergency phone when not set")
	assert.Equal(emergencyPhone, cfg.AsleepPhone(), "should default to the emergency phone when not set")
	assert.True(cfg.Valid(), "should be valid")
}

func TestValidPhones(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	setupEnvs()

	assert.True(cfg.ValidPhones(), "should be valid")
	os.Setenv("TWILIO_EMERGENCY_PHONE_NUMBER", "1234567890")
	assert.False(cfg.ValidPhones(), "should be invalid")

	setupEnvs()
	os.Setenv("TWILIO_NON_EMERGENT_PHONE_NUMBER", "1234567890")
	assert.False(cfg.ValidPhones(), "should be invalid")

	setupEnvs()
	os.Setenv("TWILIO_ASLEEP_PHONE_NUMBER", "1234567890")
	assert.False(cfg.ValidPhones(), "should be invalid")

	setupEnvs()
	os.Setenv("OUTBOUND_PHONE_NUMBER", "1234567890")
	assert.False(cfg.ValidPhones(), "should be invalid")
}

func TestValidRedis(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	setupEnvs()

	assert.True(cfg.ValidRedis(), "should be valid")
	os.Setenv("REDIS_URL", "")
	assert.False(cfg.ValidRedis(), "should be invalid")

	setupEnvs()
	assert.True(cfg.ValidRedis(), "should be valid")
	os.Setenv("REDIS_PASSWORD", "")
	assert.False(cfg.ValidRedis(), "should be invalid")
}

func TestRedisChannels(t *testing.T) {
	assert := assert.New(t)
	cfg := Config{}
	setupEnvs()

	assert.Equal(redisChannelEmergency, cfg.RedisChannelEmergency(), "should be equal")
	assert.Equal(redisChannelNonEmergent, cfg.RedisChannelNonEmergent(), "should be equal")

	clearRedisChannels()
	assert.Equal("emergency", cfg.RedisChannelEmergency(), "should be the default")
	assert.Equal("nonemergent", cfg.RedisChannelNonEmergent(), "should be the default")
}

func setupEnvs() {
	os.Setenv("TWILIO_ACCOUNT_SID", sid)
	os.Setenv("TWILIO_AUTH_TOKEN", authToken)
	os.Setenv("TWILIO_EMERGENCY_PHONE_NUMBER", emergencyPhone)
	os.Setenv("TWILIO_NON_EMERGENT_PHONE_NUMBER", nonEmergentPhone)
	os.Setenv("TWILIO_ASLEEP_PHONE_NUMBER", asleepPhone)
	os.Setenv("OUTBOUND_PHONE_NUMBER", outboundPhone)
	os.Setenv("SCRIPT_EMERGENCY_URL", emergencyURL)
	os.Setenv("SCRIPT_ASLEEP_URL", asleepURL)
	os.Setenv("REDIS_URL", redisURL)
	os.Setenv("REDIS_PASSWORD", redisPassword)
	os.Setenv("REDIS_CHANNEL_EMERGENCY", redisChannelEmergency)
	os.Setenv("REDIS_CHANNEL_NONEMERGENT", redisChannelNonEmergent)
}

func clearExtraPhones() {
	os.Setenv("TWILIO_NON_EMERGENT_PHONE_NUMBER", "")
	os.Setenv("TWILIO_ASLEEP_PHONE_NUMBER", "")
}

func clearRedisChannels() {
	os.Setenv("REDIS_CHANNEL_EMERGENCY", "")
	os.Setenv("REDIS_CHANNEL_NONEMERGENT", "")
}
