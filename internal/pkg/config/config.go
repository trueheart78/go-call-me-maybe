package config

import (
	"fmt"
	"os"
)

// Config contains methods to access environment variables
type Config struct{}

// Valid returns whether the required environment variables are setup and the phone numbers are correct
func (c Config) Valid() bool {
	if c.ValidVariables() && c.ValidPhones() {
		return true
	}
	return false
}

// ValidVariables returns whether the expected environment variables have values
func (c Config) ValidVariables() bool {
	if c.SID() != "" && c.AuthToken() != "" && c.EmergencyPhone() != "" && c.OutboundPhone() != "" && c.EmergencyURL() != "" && c.ContactName() != "" {
		return true
	}
	return false
}

// ValidRedis retuerns whether the redis-oriented environment variables have values
func (c Config) ValidRedis() bool {
	if c.RedisURL() != "" && c.RedisPassword() != "" {
		return true
	}
	return false
}

// ValidPhones returns whether the detected phone numbers begin with a '+' sign
func (c Config) ValidPhones() bool {
	valid := true
	phones := []string{c.EmergencyPhone(), c.NonEmergentPhone(), c.AsleepPhone(), c.OutboundPhone()}
	for _, phone := range phones {
		if len(phone) > 0 {
			firstChar := fmt.Sprintf("%c", phone[0])
			if firstChar != "+" {
				valid = false
				break
			}
		}
	}
	return valid
}

// SID for Twilio access
func (c Config) SID() string {
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

// AuthToken for Twilio access
func (c Config) AuthToken() string {
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

// OutboundPhone that will receive the messages/calls
func (c Config) OutboundPhone() string {
	return os.Getenv("OUTBOUND_PHONE_NUMBER")
}

// EmergencyPhone that sends messages and makes calls
func (c Config) EmergencyPhone() string {
	return os.Getenv("TWILIO_EMERGENCY_PHONE_NUMBER")
}

// NonEmergentPhone that sends messages
func (c Config) NonEmergentPhone() string {
	if os.Getenv("TWILIO_NON_EMERGENT_PHONE_NUMBER") != "" {
		return os.Getenv("TWILIO_NON_EMERGENT_PHONE_NUMBER")
	}
	return c.EmergencyPhone()
}

// AsleepPhone that makes calls
func (c Config) AsleepPhone() string {
	if os.Getenv("TWILIO_ASLEEP_PHONE_NUMBER") != "" {
		return os.Getenv("TWILIO_ASLEEP_PHONE_NUMBER")
	}
	return c.EmergencyPhone()
}

// EmergencyURL for the EmergencyPhone-related Twilio script
func (c Config) EmergencyURL() string {
	return os.Getenv("SCRIPT_EMERGENCY_URL")
}

// AsleepURL for the AsleepPhone-related Twilio script
func (c Config) AsleepURL() string {
	if os.Getenv("SCRIPT_ASLEEP_URL") != "" {
		return os.Getenv("SCRIPT_ASLEEP_URL")
	}
	return c.EmergencyURL()
}

// RedisURL for connecting to the redis datastore
func (c Config) RedisURL() string {
	return os.Getenv("REDIS_URL")
}

// RedisPassword for authenticating with the redis datastore
func (c Config) RedisPassword() string {
	return os.Getenv("REDIS_PASSWORD")
}

// RedisChannelEmergency for publishing emergency alerts
func (c Config) RedisChannelEmergency() string {
	if os.Getenv("REDIS_CHANNEL_EMERGENCY") != "" {
		return os.Getenv("REDIS_CHANNEL_EMERGENCY")
	}
	return "emergency"
}

// RedisChannelNonEmergent for publishing non-emergent alerts
func (c Config) RedisChannelNonEmergent() string {
	if os.Getenv("REDIS_CHANNEL_NONEMERGENT") != "" {
		return os.Getenv("REDIS_CHANNEL_NONEMERGENT")
	}
	return "nonemergent"
}

// ContactName for replying with the right persons name
func (c Config) ContactName() string {
	return os.Getenv("CONTACT_NAME")
}
