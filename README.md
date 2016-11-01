# LogStamper

Prepends a timestamp and optional log level to an output stream.

```
$ some_command
I am command output.
Notice how there's no timestamp.

$ some_command | logstamper -l INFO
2016-11-01T16:52:23.123Z INFO I am command output.
2016-11-01T16:52:23.234Z INFO Notice how there's no timestamp.
```

## Options

* **`--level {LEVEL}`**, **`-l {LEVEL}`**  
  An optional log level to add in addition to the timestamp.
* **`--format {FORMAT}`**, **`-f {FORMAT}`**  
  A custom date/time format to use. This string should be the formatted result
  of an example/reference time: `Mon Jan 2 15:04:05 MST 2006`; see
  https://golang.org/src/time/format.go.

  By default, this is `2006-01-02T15:04:05.999Z07:00`, which gives the RFC3339
  timestamp to 3 digits (millisecond) precision.
