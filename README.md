# Go! Call Me (Maybe)!

Designed for use with AWS Lambda and Twilio to allow an Amazon Alexa skill to call you when its
an emergency. 

![Taylor Swift - Call Me][taylor]

## Setup

### Alexa Skill Setup

Create a new skill on [the Alexa Developer Portal][alexa dev]. The name isn't important, just make it
something you remember. As for the skill's content, [here's the JSON I use][alexa json]. You should be
able to import it and adjust it according to your needs.

### Twilio

Set up an account and purchase a phone number that can do both SMS and calls. You will need this, your
SID, and the Auth Token for the AWS Lambda configuration.

#### Hosting XML For Calls

Twilio has [TwiML Bins][twiml bins] where you can put the XML for your scripts. You can't create dynamic
responses, but it makes it simpler than having to manage another service.

:warning: Regardless of where you host your XML, they need to be accessible via `POST`, otherwise calls
made will state that there was an error.

#### Sample XML

```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">Help, it's an emergency.</Say>
    <Dial/>
</Response>
```

### AWS Lambda

#### Lambda Binary

[Download it from the releases tab][releases] and upload that to your AWS Lambda function. In the _Handler_
field on the upload page, change the value to `lambda_handler`.

#### Environment Vars

You need to set the following in your configuration for your environment. If they are not setup, the lambda
will not work correctly.

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

:warning: Phone numbers **must** have a leading `+`, or Twilio will not work. The code does check for
this.

#### Redis Integration and Pub-Sub Notifications

If you would like to take advantage of the pub-sub notifications that the sister application
[Go! Call Me Notifier][go call me notifier] utilizes, the following environment variables are
required. This feature is _optional_, and leaving these out will not affect the way the lambda works.

You can get free Redis hosting from [RedisLabs][redislabs], which will cover everything this setup
needs.

```
REDIS_URL
REDIS_PASSWORD
```

The following are _optional_, and default to `emergency` and `nonemergent`, respectively.

```
REDIS_CHANNEL_EMERGENCY
REDIS_CHANNEL_NONEMERGENT
```

## Development

```
go get github.com/trueheart78/go-call-me-maybe
```

Structure should stay with the [Standard Go Project Layout][layout].

## Building the Binary

Run the `build/lambda.sh` script and check the `out/` directory for the archive to use. Make sure to set the
_Handler_ field in the AWS Lambda page to `lambda_handler` when uploading.

[twiml bins]: https://www.twilio.com/console/runtime/twiml-bins
[layout]: https://github.com/golang-standards/project-layout
[taylor]: assets/taylor-swift-call-me.gif
[alexa json]: assets/alexa.json
[go call me notifier]: https://github.com/trueheart78/go-call-me-notifier
[releases]: https://github.com/trueheart78/go-call-me-maybe/releases
[alexa dev]: https://developer.amazon.com/alexa
[redislabs]: https://redislabs.com/
