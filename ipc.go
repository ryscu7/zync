package zync

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

// getDiscordSocketPath returns the IPC socket path for Discord.
//
// It determines the appropriate path based on the operating system.
func getDiscordSocketPath() string {
	if os.PathSeparator == '/' {
		return fmt.Sprintf("/run/user/%d/discord-ipc-0", os.Getuid()) // Linux/MacOS
	}

	return `\\.\pipe\discord-ipc-0` // Windows
}

// sendIPCMessage sends a message to Discord's IPC socket.
//
// It constructs a message header and writes both the header and payload to the connection.
func sendIPCMessage(conn net.Conn, opCode int, payload []byte) error {
	header := new(bytes.Buffer)

	_ = binary.Write(header, binary.LittleEndian, int32(opCode))       // Opcode (4 bytes)
	_ = binary.Write(header, binary.LittleEndian, int32(len(payload))) // Payload Length (4 bytes)

	_, err := conn.Write(append(header.Bytes(), payload...)) // Send header + payload
	return err
}
