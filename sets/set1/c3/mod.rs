use std::collections::HashMap;

fn xor1(input: &Vec<u8>, key: u8) -> Vec<u8> {
    let mut got: Vec<u8> = Vec::new();
    for c in input.iter().as_slice() {
        got.push(*c ^ key);
    }
    got
}

pub fn break_single_byte_xor(enc: Vec<u8>) -> (Vec<u8>, u8, i64) {
    let mut res: Vec<u8> = Vec::new();
    let mut maxk: u8 = 0;
    let mut maxs: i64 = 0;
    for i in 0..255 {
        let cand = xor1(&enc, i);
        let mut score = score_char_freq(&cand);
        if let Err(_e) = std::str::from_utf8(&cand) {
            score -= 50;
        }
        if score > maxs {
            maxs = score;
            maxk = i;
            res = cand;
        }
    }
    (res, maxk, maxs)
}

fn score_char_freq(candidate: &Vec<u8>) -> i64 {
    // letter frequency should approximately match linotype key order
    let score_iter = "etaoin shrdlu".bytes();

    let mut hm: HashMap<u8, i64> = HashMap::new();
    for c in score_iter.clone() {
        hm.insert(c as u8, 0);
    }
    let mut score: i64 = 0;

    for c in candidate.iter().as_slice() {
        if !c.is_ascii_alphabetic() && *c != 0x20 {
            // unlikely
            score -= 5;
        }

        let mut incr = 5;
        if c.is_ascii_uppercase() {
            incr = 1;
        }
        if let Some(val) = hm.get_mut(&c.to_ascii_lowercase()) {
            *val += incr;
            continue;
        };
    }
    let mut last: i64 = 0;
    for c in score_iter.rev() {
        let curr = hm[&c];
        if curr >= last {
            score += curr;
        }
        last = curr;
    }
    score
}

#[cfg(test)]
mod tests {
    use hex_literal::hex;

    use crate::set1::c3::{break_single_byte_xor, score_char_freq /*, gen_freq_map*/};

    const PLAIN_TEXT: &str = "Cooking MC's like a pound of bacon";

    #[test]
    fn test_break_single_byte_xor() {
        let enc = hex!("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736");
        let res = break_single_byte_xor(enc.to_vec());
        let key = res.1;
        let plaintext: String = res.0.try_into().unwrap();
        let score = res.2;

        assert_eq!(key, 88);
        assert_eq!(score, 75);
        assert_eq!(plaintext, PLAIN_TEXT);
    }
    #[test]
    fn test_score_char_freq() {
        let score = score_char_freq(&PLAIN_TEXT.as_bytes().to_vec());
        assert_eq!(score, 75);
    }
}
