use std::env;
use std::sync::RwLock;

static PAGE: RwLock<[u8; 65536]> = RwLock::new([0; 65536]);

pub fn u4tohex(u4: u8) -> u8 {
    match u4 {
        0x00 => 0x30,
        0x01 => 0x31,
        0x02 => 0x32,
        0x03 => 0x33,
        0x04 => 0x34,
        0x05 => 0x35,
        0x06 => 0x36,
        0x07 => 0x37,
        0x08 => 0x38,
        0x09 => 0x39,
        0x0a => 0x41,
        0x0b => 0x42,
        0x0c => 0x43,
        0x0d => 0x44,
        0x0e => 0x45,
        0x0f => 0x46,
        _ => 0x30,
    }
}

pub fn seed2page(seed: i64, page: &mut [u8; 65536], alt: &str) -> u32 {
    let u4: u8 = (seed & 0x0f) as u8;
    let hx: u8 = u4tohex(u4); // '0', '1', '2', ...,
    let mut env_key: [u8; 5] = *(b"ENV_*");
    env_key[4] = hx;
    let key_str: &str = std::str::from_utf8(&env_key).ok().unwrap_or_default();
    let val: String = env::var(key_str).ok().unwrap_or_else(|| alt.into());
    let bv: &[u8] = val.as_bytes();
    let mx: usize = bv.len().min(65536);
    let limited: &[u8] = &bv[..mx];
    let ms: &mut [u8] = &mut page[..mx];
    ms.copy_from_slice(limited);
    limited.len() as u32
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn env2page(seed: i64) -> u32 {
    PAGE.try_write()
        .map(|mut guard| {
            let ma: &mut [_; 65536] = &mut guard;
            seed2page(seed, ma, "")
        })
        .ok()
        .unwrap_or(0)
}

pub fn slice2ptr(s: &[u8]) -> *const u8 {
    s.as_ptr()
}

pub fn page2ptr() -> *const u8 {
    PAGE.try_read()
        .map(|guard| {
            let a: &[_; 65536] = &guard;
            slice2ptr(a)
        })
        .ok()
        .unwrap_or_else(std::ptr::null)
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn offset64k() -> *const u8 {
    page2ptr()
}
