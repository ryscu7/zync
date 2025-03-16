# Zync

> [!IMPORTANT]
> This project is currently incomplete and under heavy development and might introduce breaking changes in the future!

A Go library for interacting with **Discord Rich Presence** by **Inter Process Communication**( IPC ).

## Features

- **Inter Process Communication**: Zync interacts with Discord Rich Presence via Inter Process Communication( IPC ) which provides more control over the customization.
- **Cross Platform Support**( Windows support isn't available yet )
- **Extremely Lightweight**: Zync only uses the standard library, making it extremely lightweight for use.
- **Standalone CLI Tool**( Upcoming )

## Installation

```go

go get github.com/ryscu7/zync
```

## Usage

### 1. Initialize a Discord IPC Client

```go

package main

import (
    "fmt"
    "log"
    "github.com/ryscu7/zync"
)

func main() {
    // Client ID refers to the Application ID!
    client, err := zync.NewClient("YOUR_CLIENT_ID")
    if err != nil {
        log.Fatalf("Failed to connect: %v\n", err)
    }
    defer client.Close()

    fmt.Println("Connected to Discord IPC!")
}
```

### 2. Perform Handshake with Discord

```go

response, err := client.Handshake()
if err != nil {
    log.Fatalf("Handshake failed: %v\n", err)
}

fmt.Printf("Authenticated as: %s#%s\n", response.Data.User.Username, response.Data.User.Discriminator)
```

## License

This project is licensed under the **MIT License**.
