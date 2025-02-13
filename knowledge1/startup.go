package knowledge1

import (
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/host"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/messagingx"
	"net/http"
	"time"
)

func init() {
	a, err1 := host.RegisterControlAgent(PkgPath, messageHandler)
	if err1 != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err1)
	}
	a.Run()
}

func messageHandler(msg *messagingx.Message) {
	start := time.Now()
	switch msg.Event() {
	case messagingx.StartupEvent:
		// Any processing for a Startup event would be here
		messagingx.SendReply(msg, aspect.NewStatusDuration(http.StatusOK, time.Since(start)))
	case messaging.ShutdownEvent:
	case messagingx.PingEvent:
		// Any processing for a Shutdown/Ping event would be here
		messagingx.SendReply(msg, aspect.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}
