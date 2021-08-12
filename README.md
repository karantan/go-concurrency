# go-concurrency ![gha build](https://github.com/karantan/go-concurrency/workflows/Go/badge.svg)

The goal of this project is to have a collection of concurrent code in Go. Because these
examples are small it's easier to understand the concepts they show. We can also check out
the code and play around with the examples to see how they behave under certain conditions.

Because the examples are small it's also easy to test them, so in this repo, we will also
show how to test concurrent code in Go.

If you have any feedback please consider raising a PR or creating a new issue.

## Usage

To see the output of the code you can take a look at the [GitHub actions](https://github.com/karantan/go-concurrency/actions),
click on a  build and expand the Test section. You will see all the outputs for all
examples.

If you want to play with the code locally you can git clone it and run tests in a debug
mode with Go plugin in Visual Studio Code editor. You will see the logs in the
terminal which will show you how the code runs.
The other way is to select an example function and run it in `main()` function in `main.go`
file and then run the main.go (`go run main.go`).

## Additional reading

I highly recommend checking out [Learn Go with Test](https://quii.gitbook.io/learn-go-with-tests/)
and reading [Concurrency in Go](https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/) book.
