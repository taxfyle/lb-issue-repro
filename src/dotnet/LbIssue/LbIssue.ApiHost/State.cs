using System.Threading.Channels;

namespace LbIssue.ApiHost;

public static class State
{
    public static readonly Channel<StreamMessagesResponse> Messages = 
        Channel.CreateUnbounded<StreamMessagesResponse>();
}
