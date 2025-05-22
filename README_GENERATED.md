# main  
Project Package Structure:  
- Dockerfile  
- go.mod  
- go.sum  
- host.go  
- kdht.go  
- main.go  

**Summary**:  
This package implements a command-line tool for setting up a Libp2p host and Kademlia DHT (Distributed Hash Table) node. It handles P2P networking, key management, and DHT bootstrap logic.  

**Configuration**:  
- **Environment Variables**: None explicitly defined.  
- **Command-line Flags**:  
  - `-port`: TCP port for listening (default: unspecified).  
  - `-seed`: Randomness seed (0 for crypto randomness, non-zero for deterministic).  
  - `-privstr`: Optional hex-encoded private key string.  
- **Files**:  
  - `go.mod`/`go.sum`: Go module dependencies.  
  - `Dockerfile`: Containerization configuration.  

**Edge Cases**:  
- If `privStr` is empty, a new RSA key pair is generated.  
- If `seed == 0`, `crypto/rand` is used instead of `math/rand`.  
- DHT initialization fails: A warning is logged, but no recovery.  
- Graceful shutdown on SIGINT/SIGTERM.  

**Code Relationships**:  
- `main.go` orchestrates host/DHT initialization, signal handling, and ASCII art.  
- `host.go` manages host creation, key generation, and address setup.  
- `kdht.go` provides Kademlia DHT functionality via `NewKDHT`.  

**Unclear/Dead Code**: None identified.