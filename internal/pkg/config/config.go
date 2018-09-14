package config

import "os"

type Config struct{}

func (c Config) Valid() bool {
	if c.SID() != "" && c.AuthToken() != "" && c.EmergencyPhone() != "" {
		return true
	}
	return false
}

func (c Config) SID() string {
	return os.Getenv("TWILIO_ACCOUNT_SID")
}

func (c Config) AuthToken() string {
	return os.Getenv("TWILIO_AUTH_TOKEN")
}

func (c Config) EmergencyPhone() string {
	return os.Getenv("TWILIO_EMERGENCY_PHONE_NUMBER")
}

func (c Config) NonEmergentPhone() string {
	if os.Getenv("TWILIO_NON_EMERGENT_PHONE_NUMBER") != "" {
		return os.Getenv("TWILIO_NON_EMERGENT_PHONE_NUMBER")
	}
	return c.EmergencyPhone()
}

func (c Config) AsleepPhone() string {
	if os.Getenv("TWILIO_ASLEEP_PHONE_NUMBER") != "" {
		return os.Getenv("TWILIO_ASLEEP_PHONE_NUMBER")
	}
	return c.EmergencyPhone()
}
