package websockethelper

import (
	"log"
	"testing"

	"github.com/keptn/keptn/cli/utils/credentialmanager"
)

const logCEString = `{
    "cloudEventsVersion":"0.1",
    "contentType":"application/json",
    "data":{"message":"InfoMsg","terminate":true},
    "eventID":"af326f24-8705-4332-b32b-affcb62f3567",
    "eventTime":"2019-03-12T17:18:14.187682+01:00",
    "eventType":"sh.keptn.events.log",
    "extensions":null,
    "source":"https://github.com/keptn/keptn/cli#wstest"
}`

func TestClient(t *testing.T) {

	credentialmanager.MockCreds = true

	ws, err := OpenWS("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaGFubmVsSWQiOiI2MDY0Njc5Zi0wNmJiLTRmMDEtOWIyNS1lMjM5Yjc4YmMzYTMiLCJpYXQiOjE1NTI0Nzk3NzgsImV4cCI6MTYzODg3OTc3OH0.ZDphJPxXrJtk4Qyk77t1nafNSzxBXBZmvcGTR7Vz064")

	if err != nil {
		t.Errorf("An error occured %v", err)
	}

	if _, err := ws.Write([]byte(logCEString)); err != nil {
		log.Fatal(err)
	}

	err = PrintWSContent(ws)
	if err != nil {
		t.Errorf("An error occured %v", err)
	}
}
