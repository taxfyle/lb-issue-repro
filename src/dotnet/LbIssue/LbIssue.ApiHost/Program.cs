using LbIssue;
using LbIssue.ApiHost;
using LbIssue.ApiHost.Services;
using Microsoft.AspNetCore.Server.Kestrel.Core;

var builder = WebApplication.CreateBuilder(args);
builder.WebHost.ConfigureKestrel(options =>
{
    options.ListenAnyIP(5001, o =>
    {
        o.Protocols = HttpProtocols.Http1;
    });
    options.ListenAnyIP(5002, o =>
    {
        o.Protocols = HttpProtocols.Http2;
    });
});

builder.Services.AddGrpc();

var app = builder.Build();

app.MapGrpcService<DemoService>();
app.MapGet("/send",
    async () =>
    {
        var message = $"Server time is {DateTime.Now:O}";
        await State.Messages.Writer.WriteAsync(new StreamMessagesResponse
        {
            Message = message
        });
        return message;
    });

app.Lifetime.ApplicationStopping.Register(() => State.Messages.Writer.Complete());

app.Run();