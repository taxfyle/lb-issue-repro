using Grpc.Core;
using Grpc.Net.Client;
using LbIssue;

var address = "http://localhost:5002";
const string name = "Dotnet";
var arg = args.FirstOrDefault();
if (!string.IsNullOrEmpty(arg))
{
    address = arg;
}

var channel = GrpcChannel.ForAddress(address);
var client = new Demo.DemoClient(channel);

var stream = client.StreamMessages(new StreamMessagesRequest
{
    Name = name
});

var start = DateTime.Now;
try
{
    
    await foreach (var streamMessagesResponse in stream.ResponseStream.ReadAllAsync())
    {
        Console.WriteLine(streamMessagesResponse.Message);
    }
}
catch (Exception e)
{
    Console.WriteLine(e);
}

Console.WriteLine($"Ran for {DateTime.Now - start}");
