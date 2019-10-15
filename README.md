# fuss
`fuss` is an easy way to fuzz any data type while still getting the coverage guided super-powers of [go-fuzz](https://github.com/dvyukov/go-fuzz).

## How to use `fuss`

A really simple fuzzing harness for `go-fuzz` might look like this:
```go
func Fuzz(data []byte) int {
    err := something.ParseBytes(data)
    if err != nil {
        return 0
    }
    return 1
}
```

But what about functions that take structs or slices or maps as inputs? (i.e. probably most of the functions you write...) Now you'd need to find a way to turn the `[]byte` that `go-fuzz` gives you into the type you actually need to test.

`fuss` does this work for you with minimum effort. Just `Seed()` it with the data from `go-fuzz` and then use it to randomly populate any data type you want. E.g. here's how you can fuzz a function that takes a `http.Request`:
```go
func Fuzz(data []byte) int {
    var request http.Request{}
    fuss.Seed(data).Fuzz(&request)
    err := validator.ValidateRequest(request)
    if err != nil {
        return 0
    }
    return 1
}
```

Now just compile and run as normal using `go-fuzz-build && go-fuzz`.

### Warning about vendoring/upgrades :warning:

Because of the way `fuss` works, it's very likely that any upgrade to `fuss` which change the way it interprets the `[]byte` from `go-fuzz`. This means all the CPU that's been spent building a corpus to stress your functions will now no be wasted as `fuss` won't generate the same data structures any more.

It's **extremely** recommended that you vendor this library and do not upgrade the version unless the new features are definitely worth invalidating your corpus.

## Design and Motivation

[go-fuzz](https://github.com/dvyukov/go-fuzz) has a lot of tricks up its sleeve to quickly find bugs in your code but unfortunately only really works for functions taking `[]byte`.

The similarly named [gofuzz](https://github.com/google/gofuzz) is the opposite extreme: you can use it to fuzz any data type but it doesn't have the coverage information necessary to seek out edge cases (it just blindly guesses).

`fuss` is very similar to to [gofuzz](https://github.com/google/gofuzz) but, instead of using a random number generator to fill your data structures, it uses the data from `go-fuzz`. This has a number of benefits over alternative methods:

* The mutations made by `go-fuzz` have small effects on the generated data
* `go-fuzz`'s magic value extraction (sonar) still works: if `go-fuzz` sees a comparison between "hello" and "world" at runtime it can find "hello" in the input data, replace it with "world" and try again.
* All `go-fuzz` inputs are valid for `fuss`. This is better than using a full binary encoding format because no CPU is spent decoding invalid inputs.

## Trophies :trophy:

Is it a proper fuzzing tool if it doesn't include a trophies section in the README?

If you've had any successes using `fuss` please do send a PR to add it here.

* [net/http: Client.Do() panics when URL includes HTTP basic auth](https://github.com/golang/go/issues/34878)
