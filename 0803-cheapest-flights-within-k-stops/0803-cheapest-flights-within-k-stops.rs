impl Solution {
    pub fn find_cheapest_price(n: i32, flights: Vec<Vec<i32>>, src: i32, dst: i32, k: i32) -> i32 {
        // result := make([]int, n)
        let mut result: Vec<i32> = vec![-1; n as usize];
        // println!("result: {:?}", result);

        result[src as usize] = 0;
        // println!("result: {:?}", result);

        for i in 0..k + 1 {
            let mut temp: Vec<i32> = result.clone();

            for flight in flights.iter() {
                let from: i32 = flight[0];
                let to: i32 = flight[1];
                let price: i32 = flight[2];

                if result[from as usize] == -1 {
                    continue;
                }

                if temp[to as usize] == -1 || temp[to as usize] > result[from as usize] + price {
                    temp[to as usize] = result[from as usize] + price;
                }

                // println!("i: {}, temp: {:?}", i, temp);
            }

            result = temp;
        }

        return result[dst as usize];
    }
}