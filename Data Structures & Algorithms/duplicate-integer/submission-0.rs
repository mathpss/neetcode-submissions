use std::collections::HashMap;
impl Solution {
    pub fn has_duplicate(nums: Vec<i32>) -> bool {
    let mut scores = HashMap::new();

    for val in nums{
        
        if scores.contains_key(&val){
            return true;
        }
        scores.insert(val, val);
    }

    return false;

    }
}
