use core::net::SocketAddr;

use std::process::ExitCode;

use tonic::transport::Server;
use tonic::Status;

use loadtest_dyn::raw::seed::integer::v1::seed64i_service_server;
use rs_seed2body_grpc_server::rpc::loadtest_dyn;
use seed64i_service_server::Seed64iServiceServer;

use rs_seed2body_grpc_server::i6b::seed6svc_new;

pub const LISTEN_ADDR_DEFAULT: &str = "127.0.0.1:50051";

async fn sub() -> Result<(), Status> {
    // implements Seed64iService
    let svc = seed6svc_new();
    let svr: Seed64iServiceServer<_> = Seed64iServiceServer::new(svc);

    let listen_addr: String = std::env::var("ENV_LISTEN_ADDR")
        .ok()
        .unwrap_or_else(|| LISTEN_ADDR_DEFAULT.into());
    let addr: SocketAddr = str::parse(listen_addr.as_str())
        .map_err(|e| Status::invalid_argument(format!("invalid listen addr: {e}")))?;

    let mut sv = Server::builder();
    sv.add_service(svr)
        .serve(addr)
        .await
        .map_err(|e| Status::internal(format!("unable to listen: {e}")))?;

    Ok(())
}

#[tokio::main]
async fn main() -> ExitCode {
    sub().await.map(|_| ExitCode::SUCCESS).unwrap_or_else(|e| {
        eprintln!("{e}");
        ExitCode::FAILURE
    })
}
