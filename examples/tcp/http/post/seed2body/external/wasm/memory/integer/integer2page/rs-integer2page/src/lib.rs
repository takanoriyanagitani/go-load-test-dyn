use std::sync::RwLock;

static PAGE: RwLock<[u8; 65536]> = RwLock::new([0; 65536]);

pub fn slice2ptr(s: &[u8]) -> *const u8 {
    s.as_ptr()
}

pub fn page2ptr() -> *const u8 {
    PAGE.try_read()
        .map(|guard| {
            let a: &[u8; 65536] = &guard;
            slice2ptr(a)
        })
        .ok()
        .unwrap_or_else(std::ptr::null)
}

pub fn integer2page(_i: i64, page: &mut [u8; 65536]) -> u32 {
    let helo: &str = r#"{"msg": "helo, wrld"}"#;
    let s: &[u8] = helo.as_bytes();
    let tgt: &mut [u8] = &mut page[..(s.len())];
    tgt.copy_from_slice(s);
    s.len() as u32
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn seed2page(seed: i64) -> u32 {
    PAGE.try_write()
        .map(|mut guard| {
            let ma: &mut [u8; 65536] = &mut guard;
            integer2page(seed, ma)
        })
        .ok()
        .unwrap_or(0)
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn offset64k() -> *const u8 {
    page2ptr()
}
