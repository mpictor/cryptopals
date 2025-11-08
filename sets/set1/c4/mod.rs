use std::error::Error;
use std::io::BufRead;

use crate::set1::c3::break_single_byte_xor;

fn detect_1b_xor(reader: &mut dyn BufRead) -> Result<Vec<u8>, Box<dyn Error>> {
    let mut pt: Vec<u8> = Vec::new();
    let mut maxs: i64 = 0;
    for line_rd in reader.lines() {
        let line = line_rd?;
        let r = break_single_byte_xor(hex::decode(line)?);
        if r.2 > maxs {
            maxs = r.2;
            pt = r.0;
            let sl = std::str::from_utf8(pt.as_slice());
            if let Err(e) = sl {
                println!("{}: {}", maxs, e);
            } else {
                println!("{}: {}", maxs, sl.unwrap());
            }
        };
    }
    Ok(pt)
}

#[cfg(test)]
mod tests {

    use std::io::BufReader;

    use crate::set1::c4::detect_1b_xor;

    const C4_TXT: &[u8] = include_bytes!("c4.txt");
    // use hex_literal::hex;
    #[test]
    fn test_this() {
        let mut rdr = BufReader::new(C4_TXT);
        let plain = detect_1b_xor(&mut rdr);

        if let Err(e) = plain {
            panic!("{}", e);
        }
        let u = plain.unwrap();
        let pt = std::str::from_utf8(u.as_slice());
        assert_eq!(pt.unwrap(), "Now that the party is jumping\n");
    }
}
