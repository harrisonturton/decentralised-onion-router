package peers

type NodeAddr struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type RelayMessage struct {
	FromAddr NodeAddr `json:"from_addr"`
	ToAddr   NodeAddr `json:"to_addr"`
	Payload  []byte   `json:"payload"`
}

type ExitMessage struct {
	FromAddr NodeAddr `json:"from_addr"`
	ToAddr   NodeAddr `json:"to_addr"`
}
