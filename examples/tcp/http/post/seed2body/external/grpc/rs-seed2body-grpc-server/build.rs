use std::io;

fn main() -> Result<(), io::Error> {
    tonic_build::configure()
        .build_client(false)
        .build_server(true)
        .compile(
            &["loadtest_dyn/raw/seed/integer/v1/seed2bytes.proto"],
            &["./load-test-dyn-proto"],
        )?;
    Ok(())
}
