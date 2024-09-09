use std::time::SystemTime;

use tonic::Request;
use tonic::Response;
use tonic::Status;

use crate::rpc::loadtest_dyn::raw::seed::integer::v1::seed64i_service_server;
use crate::rpc::loadtest_dyn::raw::seed::integer::v1::BytesFrom64iRequest;
use crate::rpc::loadtest_dyn::raw::seed::integer::v1::BytesFrom64iResponse;
use seed64i_service_server::Seed64iService;

use crate::i6time::i6to_timestamp;
use crate::time2ser::str2bytes;
use crate::time2ser::time2string;

#[tonic::async_trait]
pub trait Seed64iToBytes: Send + Sync + 'static {
    async fn seed2bytes(&self, seed: i64) -> Result<Vec<u8>, Status>;
}

#[tonic::async_trait]
impl<S> Seed64iService for S
where
    S: Seed64iToBytes,
{
    async fn bytes_from64i(
        &self,
        req: Request<BytesFrom64iRequest>,
    ) -> Result<Response<BytesFrom64iResponse>, Status> {
        let iq: BytesFrom64iRequest = req.into_inner();
        let seed: i64 = iq.seed;
        let generated: Vec<u8> = self.seed2bytes(seed).await?;
        let res = BytesFrom64iResponse { generated };
        Ok(Response::new(res))
    }
}

#[derive(Default)]
pub struct Stb {}

#[tonic::async_trait]
impl Seed64iToBytes for Stb {
    async fn seed2bytes(&self, seed: i64) -> Result<Vec<u8>, Status> {
        let sys: SystemTime = i6to_timestamp(seed)?;
        let s: String = time2string(sys);
        Ok(str2bytes(s))
    }
}

pub fn seed6svc_new() -> impl Seed64iService {
    Stb::default()
}
