# kad-dht

A command-line tool for running a Libp2p host and Kademlia DHT node.

## Overview
This package provides a peer-to-peer network node implementation using [Libp2p](https://libp2p.io/) and [Kademlia DHT](https://github.com/libp2p/go-libp2p-kad-dht) protocols. It creates a network node that can participate in distributed hash table (DHT) operations.

## File Structure
- `host.go`: Host creation logic (Libp2p host setup)
- `kdht.go`: Kademlia DHT implementation (DHT node management)
- `main.go`: Entry point (command-line interface and runtime control)

## Features
- Creates a Libp2p host with configurable network settings
- Implements Kademlia DHT protocol for distributed routing
- Supports both secure and deterministic key generation modes

## Configuration Options
| Option     | Description                                                                 | Default     |
|------------|-----------------------------------------------------------------------------|-------------|
| `--port`   | TCP port to listen on (0 = random)                                          | 0           |
| `--seed`   | Seed value for PeerID generation (0 = random)                               | 0           |
| `--privstr`| Optional hex-encoded RSA private key (overrides key generation)               | (none)      |

## Behavior Control
- **Randomness**: 
  - `seed=0` → Cryptographically secure randomness (crypto/rand)
  - `seed≠0` → Deterministic key generation (math/rand)
- **Key Handling**: 
  - Generates 2048-bit RSA keys if `privStr` is empty
  - Prints generated keys in hex format
- **DHT Mode**: 
  - `dht.ModeAutoServer` (prevents DHT from acting as client-only)

## Issues & Notes
- **Unimplemented**: 
  - `NewHost` and `NewKDHT` functions are not defined in this package (assumed to be in other packages)
  - No error handling for `host.Close()` in normal flow (only in signal handler)
- **Undefined Behavior**: 
  - `Port=0` (probably uses default port)
  - `Seed=0` (generates random PeerID, but not documented)
  - `PrivStr` has no documentation (unknown purpose)
- **Potential Crashes**: 
  - Ignored errors in multiaddr creation (invalid port may panic)
  - No validation for invalid port values (0 or >65535)
  - No fallback mechanism for failed DHT bootstrap

## Dependencies
- [libp2p](https://github.com/libp2p/go-libp2p)
- [libp2p-kad-dht](https://github.com/libp2p/go-libp2p-kad-dht)
- [go-figure](https://github.com/common-nighthawk/go-figure) (ASCII art generation)

## Build & Run
```bash
go build
./kad-dht --port=4001 --seed=1234
```