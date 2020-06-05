# whereami 🚲

> Where in Trondheim/Oslo/Bergen am I right now? Inspired by `whoami´

Wondering where in Trondheim/Oslo/Bergen you are? Just ask!

```bash
$ go get github.com/odinuge/whereami
$ whereami
⣷ Looking for you, Odin! 🚲
Found you, Odin! You are now in/at/close to S. P. Andersens vei, more accurately: 63.409889°N, 10.405213°E 🚲
```

Hopefully you can track down yourself; literally (or at least the mechanical bike-_ish_ version of you)!

## Looking for someone else?

No problem!

```bash
$ whereami -name="Karl"
⣷ Looking for you, Karl! 🚲
Found you, Karl! You are now in/at/close to S. P. Andersens vei, more accurately: 63.409889°N, 10.405213°E 🚲
```

### No work?

Missing go's bin path in your `$PATH`, no problem!

```bash
$ eval $(go env|grep GOPATH) && export PATH=$PATH:$GOPATH/bin
$ whereami
```

### Need help?

```bash
$ whereami -h
Usage of whereami:
  -city string
    	What city? (Trondheim, Bergen, Oslo) (default "Trondheim")
  -name string
        What is your first name? Defaults to your username (default "<insert-your-username-here>")
```

### Dependencies

YES, see `go.mod`

### License

MIT
