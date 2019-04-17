# Timestamp
Timestamp aims to be an easy to use general purpose time utility in Go.
Timestamp currently displays `time.Time` using Python's strftime format.

## Example usage

```go
t := time.Now()
date := timestamp.New(t)
date.Display("%H:%M:%S")
// something like 02:03:47
```

For a list of verbs, check this wonderful [website](http://strftime.org/).
All numeric verbs are zero padded by default. All of the stdlib's RFC display constants are exported in the strftime format.

## Miscellaneous functions
**Diff**:
timestamp.Diff returns the human representation of the difference between two Times. It adds the "in " prefix if the date is in the future, and the "ago" suffix if it was in the past.

## Planned features 
* Arbitrary date parsing.