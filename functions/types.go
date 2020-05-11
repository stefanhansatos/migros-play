package functions

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type SomeData struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	Number    int    `json:"number,omitempty"`
	Desc      string `json:"description,omitempty"`
	Status    string `json:"status,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Unix      int64  `json:"unix,omitempty"` // Unix time in seconds
}
