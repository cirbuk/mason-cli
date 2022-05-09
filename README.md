# mason-cli
Command line tool to interact with Mason systems

## DEVELOPMENT

- Install GoReleaser Command Line Tool on your local from [here](https://goreleaser.com/install/)

- Build the go module in Mason directory using :
```
go build
```
- Check out the configuration in **.goreleaser.yaml** found in Mason folder

- Now, create a tag and push it to GitHub:
```
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
```

- Run command from mason folder after making sure **repo working tree is clean**
```
goreleaser release
```

