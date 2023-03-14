fn calc_y(x: u32) -> u32 {
    2*x + 1
}

fn calc_z(x: u32) -> u32 {
    3*x + 1
}

fn ordered_insert(vec: &mut Vec<u32>, n: u32){
    // from https://stackoverflow.com/a/36253479
    match vec.binary_search(&n) {
        Ok(_pos) => {} // element already in vector @ `pos` 
        Err(pos) => vec.insert(pos, n),
    }
}

fn dbl_linear(n: u32) -> u32{
    // your code
    if n < 1{
        return 0;
    }
    // create a Vector and keep the index of the numbers
    // already generated 
    let mut vec = vec![1u32];
    let mut proc_index: usize = 0;
    // vec needs at least n + 1 elements 
    while vec.len() < (n + 1) as usize {
        let y = calc_y(vec[proc_index]);
        ordered_insert(&mut vec, y);
        let z = calc_z(vec[proc_index]);
        ordered_insert(&mut vec, z);
        proc_index += 1;
    }
    // can I create numbers that are smaller than the
    // ones I already have ?
    while vec[proc_index] < vec[n as usize]*3 {
        let y = calc_y(vec[proc_index]);
        ordered_insert(&mut vec, y);
        let z = calc_z(vec[proc_index]);
        ordered_insert(&mut vec, z);
        proc_index += 1;
    }
    vec[n as usize]
}

fn main(){
    fn testing(n: u32, exp: u32) -> () {
        let saida = dbl_linear(n);
        println!("n: {} exp: {}", saida, exp);
        assert_eq!(saida, exp)
    }
    
    testing(10, 22);
    testing(20, 57);
    testing(30, 91);
    testing(50, 175);
    testing(100, 447); 
}