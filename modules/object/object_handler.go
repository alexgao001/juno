package object

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	eventutil "github.com/forbole/juno/v4/types/event"
)

func (m *Module) HandleEvent(index int, event sdk.Event) error {
	eventType, err := eventutil.GetEventType(event)
	if err == nil {
		switch eventType {
		case eventutil.EventCreateObject:
			handleEventCreateObject(event)
		case eventutil.EventCancelCreateObject:
			handleEventCancelCreateObject(event)
		case eventutil.EventSealObject:
			handleEventSealObject(event)
		case eventutil.EventCopyObject:
			handleEventCopyObject(event)
		case eventutil.EventDeleteObject:
			handleEventDeleteObject(event)
		case eventutil.EventRejectSealObject:
			handleEventRejectSealObject(event)
		default:
			return nil
		}
	}
	return err
}

func handleEventCreateObject(event sdk.Event) {
	fmt.Println("handleEventCreateObject")
}

func handleEventCancelCreateObject(event sdk.Event) {
	fmt.Println("handleEventCancelCreateObject")
}

func handleEventSealObject(event sdk.Event) {
	fmt.Println("handleEventSealObject")
}

func handleEventCopyObject(event sdk.Event) {
	fmt.Println("handleEventCopyObject")
}

func handleEventDeleteObject(event sdk.Event) {
	fmt.Println("handleEventDeleteObject")
}

func handleEventRejectSealObject(event sdk.Event) {
	fmt.Println("handleEventRejectSealObject")
}