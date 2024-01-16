# Google L7 load balancing issue reproduction attempt

This is a minimal gRPC streaming demo. The server is in `src/dotnet/LbIssue.ApiHost`, and listens on:
* `5001` - HTTP/1.1, used for the `GET /send` method for sending messages
* `5002` - HTTP/2, the gRPC endpoint that the client connects to

The `src/dotnet/LbIssue.Client` is a command-line app that we can `dotnet run -- YourName https://url.to.deployment` to set up a stream listener.
