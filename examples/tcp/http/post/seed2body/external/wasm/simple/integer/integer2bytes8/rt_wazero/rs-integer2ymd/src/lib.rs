macro_rules! compose {
	($f:expr) => { $f };

	($f:expr, $($g:expr),+) => {
		|x| {
			let result = compose!($($g),+)(x);
			$f(result)
		}
	};
}

pub fn seed_unixtime_us2offset(seed_unixtime_us: i64) -> time::OffsetDateTime {
    let us7: i128 = seed_unixtime_us.into();
    time::OffsetDateTime::from_unix_timestamp_nanos(us7 * 1000)
        .unwrap_or(time::OffsetDateTime::UNIX_EPOCH)
}

pub fn offset2date(odt: time::OffsetDateTime) -> time::Date {
    odt.date()
}

pub fn seed_unixtime_us2ymd(seed_unixtime_us: i64) -> time::Date {
    compose!(offset2date, seed_unixtime_us2offset)(seed_unixtime_us)
}

pub fn u4to1byte(u: u8) -> u8 {
    match u {
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

pub fn yyyy2bytes4(yyyy: u16) -> (u8, u8, u8, u8) {
    let y0: u16 = yyyy / 1000;
    let y1: u16 = (yyyy / 100) % 10;
    let y2: u16 = (yyyy / 10) % 10;
    let y3: u16 = yyyy % 10;
    (y0 as u8, y1 as u8, y2 as u8, y3 as u8)
}

pub fn integers2u64(integers: (u8, u8, u8, u8, u8, u8, u8, u8)) -> u64 {
    let i0: u64 = integers.0.into();
    let i1: u64 = integers.1.into();
    let i2: u64 = integers.2.into();
    let i3: u64 = integers.3.into();
    let i4: u64 = integers.4.into();
    let i5: u64 = integers.5.into();
    let i6: u64 = integers.6.into();
    let i7: u64 = integers.7.into();
    i7 | i6 << 8 | i5 << 16 | i4 << 24 | i3 << 32 | i2 << 40 | i1 << 48 | i0 << 56
}

pub fn integers2tuple8(integers: (u16, u8, u8)) -> (u8, u8, u8, u8, u8, u8, u8, u8) {
    let iyyyy = yyyy2bytes4(integers.0);
    let y0: u8 = u4to1byte(iyyyy.0);
    let y1: u8 = u4to1byte(iyyyy.1);
    let y2: u8 = u4to1byte(iyyyy.2);
    let y3: u8 = u4to1byte(iyyyy.3);

    let m0: u8 = u4to1byte(integers.1 / 10);
    let m1: u8 = u4to1byte(integers.1 % 10);

    let d0: u8 = u4to1byte(integers.2 / 10);
    let d1: u8 = u4to1byte(integers.2 % 10);
    (y0, y1, y2, y3, m0, m1, d0, d1)
}

pub fn integers2bytes8(integers: (u16, u8, u8)) -> u64 {
    compose!(integers2u64, integers2tuple8)(integers)
}

pub fn ymd2integers(ymd: (i32, time::Month, u8)) -> (u16, u8, u8) {
    let y: u16 = ymd.0 as u16;
    let m: u8 = ymd.1.into();
    (y, m, ymd.2)
}

pub fn ymd2bytes8(ymd: (i32, time::Month, u8)) -> u64 {
    compose!(integers2bytes8, ymd2integers)(ymd)
}

pub fn date2ymd(dt: time::Date) -> (i32, time::Month, u8) {
    let year: i32 = dt.year();
    let m: time::Month = dt.month();
    let day: u8 = dt.day();
    (year, m, day)
}

pub fn date2ymd8bytes(dt: time::Date) -> u64 {
    compose!(ymd2bytes8, date2ymd)(dt)
}

#[allow(unsafe_code)]
#[no_mangle]
pub extern "C" fn seed_unixtime_us2ymd8bytes(seed_unixtime_us: i64) -> u64 {
    compose!(date2ymd8bytes, seed_unixtime_us2ymd)(seed_unixtime_us)
}

#[cfg(test)]
mod test_yyyy2bytes4 {
    #[test]
    fn yr() {
        let yyyy = crate::yyyy2bytes4(2019);
        assert_eq!(yyyy.0, 2);
        assert_eq!(yyyy.1, 0);
        assert_eq!(yyyy.2, 1);
        assert_eq!(yyyy.3, 9);
    }
}

#[cfg(test)]
mod test_integers2tuple8 {
    #[test]
    fn yr() {
        let converted = crate::integers2tuple8((2019, 5, 1));
        assert_eq!(converted, (0x32, 0x30, 0x31, 0x39, 0x30, 0x35, 0x30, 0x31));
    }
}

#[cfg(test)]
mod test_integers2bytes8 {
    #[test]
    fn yr() {
        let converted: u64 = crate::integers2bytes8((2019, 5, 1));
        assert_eq!(&converted.to_be_bytes(), b"20190501");
    }
}

#[cfg(test)]
mod test_seed_unixtime_us2ymd8bytes {
    #[test]
    fn epoch() {
        let converted: u64 = crate::seed_unixtime_us2ymd8bytes(0);
        assert_eq!(&converted.to_be_bytes(), b"19700101");
    }
}
