using Grpc.Core;

namespace LbIssue.ApiHost.Services;

public class DemoService(ILogger<DemoService> logger) : Demo.DemoBase
{
    public override async Task StreamMessages(StreamMessagesRequest request,
        IServerStreamWriter<StreamMessagesResponse> responseStream, ServerCallContext context)
    {
        logger.LogInformation("Streaming call started for {Name}", request.Name);

        try
        {
            await foreach (var streamMessagesResponse in State.Messages.Reader.ReadAllAsync(context.CancellationToken))
            {
                logger.LogInformation("Sending message to {Name}: {Message}", request.Name, streamMessagesResponse.Message);
                await responseStream.WriteAsync(streamMessagesResponse);
            }
        }
        catch (OperationCanceledException)
        {
            logger.LogInformation("Streaming call canceled for {Name}", request.Name);
        }

        logger.LogInformation("Streaming call ended for {Name}", request.Name);
    }
}