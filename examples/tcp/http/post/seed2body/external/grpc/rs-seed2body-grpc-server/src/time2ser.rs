use std::time::SystemTime;

pub fn time2string(t: SystemTime) -> String {
    format!("{t:#?}")
}

pub fn str2bytes(s: String) -> Vec<u8> {
    s.into()
}
