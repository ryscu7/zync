package zync

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// NewClient initializes a new Discord IPC client and establishes a connection.
//
// It returns a pointer to a [Client] instance or an error if the connection fails.
func NewClient(clientID string) (*Client, error) {
	socketPath := getDiscordSocketPath()

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Discord IPC: %w", err)
	}

	client := &Client{Conn: conn, ClientID: clientID}
	return client, nil
}

// Handshake performs the initial handshake with Discord and returns the response.
//
// It sends a handshake request and waits for a response from Discord. If successful,
// it returns a pointer to [HandshakeResponse]; otherwise, it returns an error.
func (c *Client) Handshake() (*HandshakeResponse, error) {
	payload := HandshakePayload{
		Version:  1,
		ClientID: c.ClientID,
	}
	data, _ := json.Marshal(payload)

	err := sendIPCMessage(c.Conn, 0, data) // Opcode 0: Handshake
	if err != nil {
		return nil, fmt.Errorf("handshake failed: %w", err)
	}

	// Read the first 8 bytes (header)
	header := make([]byte, 8)
	_, err = c.Conn.Read(header)
	if err != nil {
		return nil, fmt.Errorf("failed to read handshake header: %w", err)
	}

	// Read the JSON payload
	payloadLength := int(binary.LittleEndian.Uint32(header[4:])) // Extract the payload length
	jsonPayload := make([]byte, payloadLength)

	_, err = c.Conn.Read(jsonPayload)
	if err != nil {
		return nil, fmt.Errorf("failed to read handshake payload: %w", err)
	}

	// Unmarshal response into HandshakeReponse type
	var handshakeResponse HandshakeResponse

	err = json.Unmarshal(jsonPayload, &handshakeResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse handshake: %w", err)
	}

	return &handshakeResponse, nil
}

// Close safely closes the IPC connection
func (c *Client) Close() error {
	return c.Conn.Close()
}
