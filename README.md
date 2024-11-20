# p2boot
## Usage
use `go run .` | `go build . -o p2boot && ./p2boot` | `docker build . ...`   
you can specify `-seed [int64]` to have somewhat limited number of possible keys at restart   
after each start you can copy the generated base16 private key (i'm against the security, it's waste of the time and we are all going to die soon anyways) and paste it into the `-priv [string]` flag   
you can specify `-port [int]` flag to have the same port after restart  

## Credits
[Based on this avvvet work](https://github.com/avvvet/kad-dht-boot)
