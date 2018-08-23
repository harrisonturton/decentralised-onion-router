package comm

type Message struct {
	/* Says if we're to operate as an exit node */
	ExitAddress *string `json:"exit_address"`

	/* If not exit node, then relay. Address of node to forward to */
	Next *string `json:"next"`

	/* Payload to transfer */
	Payload string `json:"payload"`
}

type ExitMessage struct {
	Address string
	Payload string
}

type RelayMessage struct {
	Next string
	Payload string
}
