package main


type Packet struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}

// Serialize
packet := Packet{ID: 1, Data: "Hello"}
data, err := json.Marshal(packet)
if err != nil {
    log.Fatal(err)
}

// Deserialize

var packet Packet
err := json.Unmarshal(data, &packet)
if err != nil {
	log.Fatal(err)
}