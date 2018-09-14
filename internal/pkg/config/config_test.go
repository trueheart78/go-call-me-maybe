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

// TestConfigSetup description
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

// TestNil description
func TestNil(t *testing.T) {
	assert := assert.New(t)
	// assert for nil - good for error-checking
	assert.Nil(nil, "should be nil")
	assert.NotNil(0, "should not be nil")
}

func setupEnvs() {
	os.Setenv("TWILIO_ACCOUNT_SID", sid)
	os.Setenv("TWILIO_AUTH_TOKEN", authToken)
	os.Setenv("TWILIO_EMERGENCY_PHONE_NUMBER", emergencyPhone)
	os.Setenv("TWILIO_NON_EMERGENT_PHONE_NUMBER", nonEmergentPhone)
	os.Setenv("TWILIO_ASLEEP_PHONE_NUMBER", asleepPhone)
	os.Setenv("OUTBOUND_PHONE_NUMBER", outboundPhone)
}

func clearExtraPhones() {
	os.Setenv("TWILIO_NON_EMERGENT_PHONE_NUMBER", "")
	os.Setenv("TWILIO_ASLEEP_PHONE_NUMBER", "")
}
