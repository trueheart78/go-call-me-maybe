# Go! Call Me (Maybe)!

Designed for use with AWS Lambda and Twilio to allow an Amazon Alexa skill to call you when its
an emergency.

![Taylor Swift - Call Me][taylor]

## Environment Vars

You need to set the following in your configuration for your environment:

```
TWILIO_ACCOUNT_SID
TWILIO_AUTH_TOKEN
TWILIO_EMERGENCY_PHONE_NUMBER
OUTBOUND_PHONE_NUMBER
SCRIPT_EMERGENCY_URL
```

The following are optional. They default to the emergency-related values if unset.

```
TWILIO_ASLEEP_PHONE_NUMBER
TWILIO_NON_EMERGENT_PHONE_NUMBER
SCRIPT_ASLEEP_URL
```

:warning: Phone numbers **must** have a leading `+`, or Twilio will not work.

## Hosting XML For Calls

Twilio has [TwiML Bins][twiml bins] where you can put the XML for your scripts. You can't create dynamic
responses, but it makes it simpler than having to manage another service.

:warning: Regardless of where you host your XML, they need to be accessible via `POST`, otherwise calls made
will state that there was an error.

### Sample XML

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">Help, it's an emergency.</Say>
    <Dial/>
</Response>
```

## Development

```
go get github.com/trueheart78/go-call-me-maybe
```

Structure should stay with the [Standard Go Project Layout]

[twiml bins]: https://www.twilio.com/console/runtime/twiml-bins
[layout]: https://github.com/golang-standards/project-layout
[taylor]: assets/taylor-swift-call-me.gif
