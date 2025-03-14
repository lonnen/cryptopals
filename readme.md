# cryptopals challenges

[Cryptopals](https://www.cryptopals.com/) challenges are supposed to be a fun way to learn a new language, so maybe now's the time to pick up Go?

[![test status badge](https://github.com/lonnen/cryptopals/actions/workflows/test.yml/badge.svg)](https://github.com/lonnen/cryptopals/actions/workflows/test.yml)

### Implementation status

- [x] [Set 1: Basics](https://www.cryptopals.com/sets/1)
- [ ] [Set 2: Block crypto](https://www.cryptopals.com/sets/2)
- [ ] [Set 3: Block & stream crypto](https://www.cryptopals.com/sets/3)
- [ ] [Set 4: Stream crypto and randomness](https://www.cryptopals.com/sets/4)
- [ ] [Set 5: Diffie-Hellman and friends](https://www.cryptopals.com/sets/5)
- [ ] [Set 6: RSA and DSA](https://www.cryptopals.com/sets/6)
- [ ] [Set 7: Hashes](https://www.cryptopals.com/sets/7)
- [ ] [Set 8: Abstract Algebra](https://www.cryptopals.com/sets/8)

## How to use

Problem solutions are implemented as functions in `set_*.go` and verified by tests (`set_*_test.go`). Common code is kept in `main.go`

To run tests:
```shell
# run all the tests
$ go test

# run a single test
$ go test --run Test5
```