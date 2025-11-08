#[cfg(test)]
mod tests {
    use hex_literal::hex;
    #[test]
    fn test_this() {
        let input = hex!("1c0111001f010100061a024b53535009181c");
        let key = hex!("686974207468652062756c6c277320657965");
        let want = hex!("746865206b696420646f6e277420706c6179");
        let mut got: Vec<u8> = Vec::new();

        for (i, c) in input.iter().enumerate() {
            got.push(c ^ key[i]);
        }
        assert_eq!(want, got.as_slice());
    }
}
