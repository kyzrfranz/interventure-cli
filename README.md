# interventure-cli
A bunch of generator to be used with www.github.com/kyzrfranz/buntesdach.

This has been used to build the campaign www.stoppt-scheinselbstaendigkeit.de

## Use
If you use the binaries, make sure to point to a running instance of the buntesdach API.

```bash

You can either run
```bash
make build
```
and run the binaries directly. Or you can use the make targets.

Generate a static json containing all MdB bios of the current legislative period.
```bash
make bio-all
```

Download all available images to be able to process them
```bash
make fetch-images
```

Resize, crop the downloaded images to be used on a website
```bash
make process-images
```


There are some other things in there that currently don't really work, like generating AI prompts or scraping personal websites.
Maybe in the future...
