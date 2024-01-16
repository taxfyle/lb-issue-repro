mod proto;

use crate::proto::demo_client::DemoClient;
use crate::proto::StreamMessagesRequest;
use anyhow::Context;
use futures::StreamExt;
use tonic::transport::Endpoint;

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let endpoint = std::env::args().collect::<Vec<_>>().get(1)
        .unwrap_or(&"http://localhost:5001".to_owned())
        .clone();
    println!("Connecting to {}", &endpoint);
    let channel = Endpoint::try_from(endpoint)?
        .connect()
        .await
        .context("Unable to connect")?;

    let mut client = DemoClient::new(channel);

    let response = client
        .stream_messages(StreamMessagesRequest {
            name: "Rust".to_owned(),
        })
        .await
        .context("Streaming messages")?;

    let mut response = response.into_inner();

    while let Some(Ok(message)) = response.next().await {
        println!("Got message: {}", message.message);
    }
    println!("Streaming call ended, good bye");

    Ok(())
}
