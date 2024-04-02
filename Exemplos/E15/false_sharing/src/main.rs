use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Instant;

const CACHE_LINE_SIZE: usize = 64; // Typically 64 bytes

#[derive(Debug)]
struct PaddedInt {
    value: i32,
    _pad: [u8; CACHE_LINE_SIZE - std::mem::size_of::<i32>()], // Padding to fill the cache line
}

fn main() {
    let size = 1_000_000;
    let num_threads = 4; // Number of threads to use for concurrency
    let chunk_size = size / num_threads;

    let array: Vec<i32> = vec![1; size];

    // Concurrent array sum operation without padding
    let start_unpadded = Instant::now();
    let sum_unpadded = concurrent_sum_array(&array, num_threads, chunk_size);
    let elapsed_unpadded = start_unpadded.elapsed();
    println!(
        "Concurrent array sum without padding: {}, Time taken: {:?}",
        sum_unpadded, elapsed_unpadded
    );

    // Concurrent array sum operation with padding to avoid false sharing
    let start_padded = Instant::now();
    let sum_padded = concurrent_sum_array_with_padding(&array, num_threads, chunk_size);
    let elapsed_padded = start_padded.elapsed();
    println!(
        "Concurrent array sum with padding: {}, Time taken: {:?}",
        sum_padded, elapsed_padded
    );
}

fn concurrent_sum_array(array: &[i32], num_threads: usize, chunk_size: usize) -> i32 {
    let sum = Arc::new(Mutex::new(0));

    let mut handles = vec![];
    for i in 0..num_threads {
        let start = i * chunk_size;
        let end = if i == num_threads - 1 {
            array.len()
        } else {
            start + chunk_size
        };
        let thread_sum = sum.clone();
        let thread_array = array[start..end].to_vec();

        let handle = thread::spawn(move || {
            let local_sum: i32 = thread_array.iter().sum();
            let mut sum = thread_sum.lock().unwrap();
            *sum += local_sum;
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    let final_sum = *sum.lock().unwrap();
    final_sum
}

fn concurrent_sum_array_with_padding(array: &[i32], num_threads: usize, chunk_size: usize) -> i32 {
    let sums: Vec<Mutex<PaddedInt>> = (0..num_threads)
        .map(|_| Mutex::new(PaddedInt { value: 0, _pad: [0; CACHE_LINE_SIZE - 4] }))
        .collect();
    let sums = Arc::new(sums);

    let mut handles = vec![];
    for i in 0..num_threads {
        let start = i * chunk_size;
        let end = if i == num_threads - 1 {
            array.len()
        } else {
            start + chunk_size
        };
        let thread_sums = sums.clone();
        let thread_array = array[start..end].to_vec();

        let handle = thread::spawn(move || {
            let local_sum: i32 = thread_array.iter().sum();
            let mut sum = thread_sums[i].lock().unwrap();
            sum.value += local_sum;
        });

        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    let final_sum: i32 = sums.iter().map(|s| s.lock().unwrap().value).sum();
    final_sum
}
