package functions

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

type Payload struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type WrappedData struct {
	Source    string   `json:"source"`
	Payload   *Payload `json:"payload"`
	Name      string   `json:"name,omitempty"`
	Number    int      `json:"number,omitempty"`
	Desc      string   `json:"description,omitempty"`
	Status    string   `json:"status,omitempty"`
	Timestamp string   `json:"timestamp,omitempty"`
	Unix      int64    `json:"unix,omitempty"` // Unix time in seconds
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
