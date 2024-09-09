pub mod rpc {
    pub mod loadtest_dyn {
        pub mod raw {
            pub mod seed {
                pub mod integer {
                    #[allow(clippy::unwrap_used)]
                    pub mod v1 {
                        tonic::include_proto!("loadtest_dyn.raw.seed.integer.v1");
                    }
                }
            }
        }
    }
}

pub mod i6b;

pub mod i6time;
pub mod time2ser;
