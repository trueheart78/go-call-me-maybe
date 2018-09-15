package main

import (
	"net/url"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kevinburke/twilio-go"
	"github.com/njuettner/go-alexa"
	"github.com/trueheart78/go-call-me-maybe/internal/pkg/config"
)

func alexaDispatchIntentHandler(req alexa.Request) (*alexa.Response, error) {
	if req.RequestBody.Intent.Name == "StatusCheck" {
		return statusCheck()
	}

	cfg := config.Config{}
	if !cfg.Valid() {
		return errorResponse()
	}
	client := twilio.NewClient(cfg.SID(), cfg.AuthToken(), nil)

	switch req.RequestBody.Intent.Name {
	case "Emergency":
		// send a text
		// TODO create a sendEmergencyRequest()
		_, err := client.Messages.SendMessage(cfg.EmergencyPhone(), cfg.OutboundPhone(), "🚨 Halp 🚨", nil)
		if err != nil {
			return errorResponse()
		}
		// make a phone call
		u, _ := url.Parse(cfg.EmergencyURL())
		_, err = client.Calls.MakeCall(cfg.EmergencyPhone(), cfg.OutboundPhone(), u)
		if err != nil {
			return errorResponse()
		}
		return simpleResponse("Okay. I have called and texted Josh.")
	case "NextTenMinutes":
		// send a low priority text
		// TODO create a sendNonEmergentRequest()
		_, err := client.Messages.SendMessage(cfg.NonEmergentPhone(), cfg.OutboundPhone(), "Help me when you have a minute? No rush 💖", nil)
		if err != nil {
			return errorResponse()
		}
		return simpleResponse("Okay. I let him know. Bug him again if he doesn't respond in 10 minutes.")
	case "WakeUp":
		// make a wake up phone call
		// TODO create a sendWakeUpRequest()
		u, _ := url.Parse(cfg.AsleepURL())
		_, err := client.Calls.MakeCall(cfg.AsleepPhone(), cfg.OutboundPhone(), u)
		if err != nil {
			return errorResponse()
		}
		return simpleResponse("Calling that sleeping hubby now.")
	default:
		return alexaHelpHandler()
	}
}

func main() {
	lambda.Start(alexaDispatchIntentHandler)
}

func statusCheck() (*alexa.Response, error) {
	cfg := config.Config{}
	if !cfg.Valid() {
		if !cfg.ValidVariables() {
			return simpleResponse("There are missing configuration values.")
		} else if !cfg.ValidPhones() {
			return simpleResponse("There seems to be a problem with the phone number setup.")
		} else {
			return simpleResponse("Something is not right.")
		}
	} else {
		return simpleResponse("Everything seems to be working.")
	}
}

func errorResponse() (*alexa.Response, error) {
	return simpleResponse("Unable to contact Josh. Please ask Siri for assistance.")
}

func simpleResponse(content string) (*alexa.Response, error) {
	simpleResponse := &alexa.SimpleResponse{
		OutputSpeechText: content,
		CardTitle:        "Greeter",
		CardContent:      "Greeter Content",
	}
	return alexa.NewResponse(simpleResponse), nil
}

func alexaHelpHandler() (*alexa.Response, error) {
	helpResponse := &alexa.SimpleResponse{
		OutputSpeechText: "You can say things like, 'I need you', 'stop snoring', or 'its not urgent'",
		CardTitle:        "Help for Greeter",
		CardContent:      "Card Content",
	}
	return alexa.NewResponse(helpResponse), nil
}
