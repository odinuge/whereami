# whereami 🚲

> Where in Trondheim am I right now? Inspired by `whoami´

Wondering where in Trondheim you are? Just ask!

```bash
$ go get github.com/odinuge/whereami
$ whereami
Hi Odin! You are now in/at/close to Erling Skakkes gate 47A, more accurately: 63.427925°N, 10.389504°E 🚲
```

Hopefully you can track down yourself; literally (or atleast the mechanical bike-_ish_ version of you)!

## Looking for someone else?

No problem!

```bash
$ whereami -name="Karl"
Hi Karl! You are now in/at/close to Professor Brochs gate 2, more accurately: 63.416145°N, 10.396315°E 🚲
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
  -name string
        What is your first name? Defaults to your username (default "<insert-your-username-here>")
```

### Dependencies

None

### License

MIT
