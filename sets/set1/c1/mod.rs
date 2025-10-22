#[cfg(test)]
mod tests {
    // use super::*;
    use base64::Engine;
    use base64::prelude::BASE64_STANDARD;
    use hex_literal::hex;

    #[test]
    fn test_enc() {
        let input = hex!(
            "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
        );
        const WANT: &str = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t";
        let got = BASE64_STANDARD.encode(input);
        assert!(WANT == got);
    }
}
