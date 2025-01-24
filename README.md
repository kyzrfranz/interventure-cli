# interventure-cli


## Commands

### bio
fetches the bios from the `buntesdach api` and stores it in a json for further processing by `transform`

```shell
make build 
./bin/bio
```

#### flags
- `--apiUrl` : path to the input file (default: `http://localhost:8080`) you can point it to `https://buntesdach-api-983281881572.europe-west1.run.app`
- `--out` : path to the output file (default: `.bio.json`)
- `--max` : limit the number of bios fetched (default: `all`)


### transform
transforms the bios from the `buntesdach api` into an xslx addresslist

```shell
make build
./bin/transform 
```

#### flags
- `--in` : path to the input file (default: `.bio.json`)
- `--out` : path to the output file (default: `.addresslist.xlsx`)