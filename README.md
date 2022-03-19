# fancyChecks

fancyChecks is a Go package to make error checking and printing easier for command-line tools and apps.

`go get github.com/PerpetualCreativity/fancyChecks`

## Usage

```go
// create a checker
fc := fancyChecker.New("My app", "Command completed successfully", "", "Error")

result, err := some.Function()

// Error checking functions print in red.
// if err != nil, prints "My app: Error: someFunction has failed" and the error, then exits
fc.ErrCheck(err, "some.Function has failed")

// if err != some.ErrorNotAvailable, prints "My app: Error: Data access issue" and the error, then exits
fc.ErrComp(err, some.ErrorNotAvailable, "Data access issue.")

// if err == some.ErrorNotAvailable, prints "My app: Error: Data is not available" and the error, then exits
fc.ErrNComp(err, some.ErrorNotAvailable, "Data is not available.")

// if err != nil and err != some.ErrorNotAvailable, prints "My app: Error: Data is not available" and the error, then exits
fc.ErrExp(err, some.ErrorNotAvailable, "Unknown data error.")

// There are also two helper functions:
// Prints "My app: Command completed successfully: Added config file" in green but does not exit
fc.Success("Added config file")

// Prints "My app: Contacting server..." in blue but does not exit
// Note that this does not have a neutral message prefix because we didn't specify one when creating fc above.
fc.Neutral("Contacting server")
```

