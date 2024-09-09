use core::time::Duration;

use std::time::SystemTime;

use tonic::Status;

pub fn i6to_timestamp(seed_unixtime_us: i64) -> Result<SystemTime, Status> {
    let dur: Duration = (u64::try_from(seed_unixtime_us))
        .ok()
        .map(Duration::from_micros)
        .ok_or_else(|| Status::invalid_argument(format!("invalid seed: {seed_unixtime_us}")))?;
    SystemTime::UNIX_EPOCH
        .checked_add(dur)
        .ok_or_else(|| Status::invalid_argument(format!("invalid duration: {dur:#?}")))
}
