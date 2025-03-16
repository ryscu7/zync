package zync

import (
	"net"
)

// Client manages IPC connection with Discord.
//
// It is not recommended to use this directly for initialization. Instead, consider using [NewClient]
type Client struct {
	Conn     net.Conn // Conn represents the active IPC connection.
	ClientID string   // ClientID represents the Application ID from Discord Developer Portal.
}

// HandshakePayload represents the JSON structure for the handshake request
type HandshakePayload struct {
	Version  int    `json:"v"`         // Protocol version
	ClientID string `json:"client_id"` // Application ID
}

// HandshakeResponse represents the expected JSON structure from Discord
type HandshakeResponse struct {
	Cmd   string `json:"cmd"` // The command returned by Discord
	Event string `json:"evt"` // The event type
	Data  struct {
		User struct {
			ID            string `json:"id"`            // User ID
			Username      string `json:"username"`      // Discord Username
			Discriminator string `json:"discriminator"` // User's discriminator (this will usually be 0)
		} `json:"user"`
	} `json:"data"`
}
