package main

type List []

type wrappedData struct {
	Payload   []byte      `json:"payload"`
	ID        string      `json:"id"`
	Name      string      `json:"name,omitempty"`
	Number    int         `json:"number,omitempty"`
	Desc      interface{} `json:"description,omitempty"`
	Status    string      `json:"status,omitempty"`
	Timestamp string      `json:"timestamp,omitempty"`
	Unix      int64       `json:"unix,omitempty"` // Unix time in seconds
}

func main() {

}
