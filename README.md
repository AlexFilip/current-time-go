# Current Time Go

This small program was made to run as the time block in my i3 config

## Building
You can build the docker container locally with
```bash
docker build -t current-time:latest .
```

And run it with
```bash
docker run --rm -e TZ="Your/Timezone" -e TIME_FORMAT="Format" current-time:latest
```

`TZ` is your local timezone
`TIME_FORMAT` is the format you would like to use in the [Go time format layout](https://pkg.go.dev/time#pkg-constants)
