use std::collections::HashMap;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut hash_map_s = HashMap::new();
        let mut hash_map_t = HashMap::new();

        for char_s in s.chars() {
            hash_map_s.entry(char_s).and_modify(|x| *x += 1).or_insert(1);
        }

        for char_t in t.chars() {
            hash_map_t.entry(char_t).and_modify(|x| *x += 1).or_insert(1);
        }

        return hash_map_s.eq(&hash_map_t);
        }
}
