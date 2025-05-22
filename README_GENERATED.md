# p2boot

## Project Package Structure
- Dockerfile
- go.mod
- go.sum
- host.go
- kdht.go
- main.go

## Configuration Elements
### Environment Variables
None explicitly defined in provided files.

### Flags
- `-port` (int): TCP port for listening
- `-seed` (int): Randomness seed (0 for crypto randomness)
- `-privstr` (string): Optional hex-encoded private key

### Files
- `go.mod`/`go.sum`: Go module configuration
- `Dockerfile`: Container build instructions

## Edge Cases
- Missing/invalid flags result in runtime errors (not explicitly handled)
- No fallback values shown for port/seed/privstr

## Code Relationships
- `main.go` orchestrates host/DHT initialization and signal handling
- `host.go` handles libp2p host creation and key management
- `kdht.go` implements Kademlia DHT functionality via libp2p-kad-dht
- All components integrate with libp2p networking stack

The package implements a P2P node with Kademlia DHT capabilities, handling crypto key generation, network address setup, and signal termination. It provides ASCII art visualization and DHT bootstrap functionality.

