package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/njuettner/go-alexa"
	"github.com/trueheart78/go-call-me-maybe/internal/pkg/config"
)

func alexaDispatchIntentHandler(req alexa.Request) (*alexa.Response, error) {
	cfg := config.Config{}
	if !cfg.Valid() {
		return simpleResponse("Unable to contact Josh. Please ask Siri for assistance.")
	}

	switch req.RequestBody.Intent.Name {
	case "Emergency":
		// send a text
		// make a phone call
		return simpleResponse("Okay. I have called and texted Josh.")
	case "NextTenMinutes":
		// send a low priority text
		return simpleResponse("Okay. I let him know. Bug him again if he doesn't respond in 10 minutes.")
	case "WakeUp":
		// make a wake up phone call
		return simpleResponse("Calling that sleeping hubby now.")
	case "StatusCheck":
		// making it this far means everything is ok
		return simpleResponse("Everything seems to be in working order.")
	default:
		return alexaHelpHandler()
	}
}

func main() {
	lambda.Start(alexaDispatchIntentHandler)
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
