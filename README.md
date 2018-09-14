# Go! Call Me (Maybe)!

Designed for use with AWS Lambda and Twilio to allow an Amazon Alexa skill to call you when its
an emergency.

![Taylor Swift - Call Me](assets/taylor-swift-call-me.gif)

## Environment Vars

You need to set the following in your configuration for your environment:

```
TWILIO_ACCOUNT_SID
TWILIO_AUTH_TOKEN
TWILIO_EMERGENCY_PHONE_NUMBER
OUTBOUND_PHONE_NUMBER
```

The following are optional. They default to the `TWILIO_EMERGENCY_PHONE_NUMBER` if unset.

```
TWILIO_ASLEEP_PHONE_NUMBER
TWILIO_NON_EMERGENT_PHONE_NUMBER
```

## Development

```
go get github.com/trueheart78/go-call-me-maybe
```

Structure should stay with the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
