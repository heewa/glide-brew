# glide-brew
Convert Go deps managed by [glide](https://github.com/Masterminds/glide) to [Homebrew](http://brew.sh/) resources to help you make brew formulas for you Go programs. The resulting config can be used in Homebrew formulas to have brew share your Go dependencies as brew resources.

See Homebrew's [documentation on resources](https://github.com/Homebrew/brew/blob/master/share/doc/homebrew/Formula-Cookbook.md#specifying-gems-python-modules-go-projects-etc-as-dependencies) for instructions on how to use them.

# usage

Install with:
```bash
go get github.com/heewa/glide-brew
```

Then run `glide brew` from your Go repo, where you have your `glide.yaml` and `glide.lock` files.

## troubleshooting

If glide complains like: `[ERROR] Command glide-brew does not exist.`, then you probably don't have wherever you installed `glide-brew` into from `go get` in your `PATH`. Either add `$GOPATH/bin` to your `PATH`, or symlink wherever `glide-brew` was installed to something that is, like `/usr/local/bin/`.
