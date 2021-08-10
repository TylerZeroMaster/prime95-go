# Prime95-go

## About
This program finds [Mersenne prime numbers](https://en.wikipedia.org/wiki/Mersenne_prime) in sequence starting at three. I made this program so that I would become more familiar with big number libraries. It utilizes my implementation of an [optimized modulo](https://en.wikipedia.org/wiki/Lucas%E2%80%93Lehmer_primality_test#Time_complexity) operation. I did not, however, implement a FFT multiplication algorithm.

## Performance Findings

#### Methodology

All tests were conducted on an Intel Core i5-4690K. Each test consisted of finding the first 20 Mersenee prime exponents using the [Lucas-Lehmer primality test.](https://en.wikipedia.org/wiki/Lucas%E2%80%93Lehmer_primality_test) Each environment's standard big number implementation was used for testing. The average time of five consecutive runs was taken for each language tested.

#### Results

| Environment      | Time (seconds) |
| ---------------- |:--------------:|
| Golang           | 11             |
| Python v3.9.1    | 30             |
| .NET v5.0.201    | 36             |
| Node.js v14.15.0 | 77             |

#### Conclusion

The big number libraries in Python, Node, and .NET utilize immutable data structures. Consequently, each arithmetic operation causes a new object to be instantiated. This has a large performance impact on repetitive operations. 

Go's big number library utilizes mutable data structures. This, in addition to optimized arithmetic operations, put Go's standard big number library ahead of the others in terms of performance.

## Note
**This program is in no way related to the Prime95 project or any other similarly named projects.** It was named in honor of that project and nothing more. 