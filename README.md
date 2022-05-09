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

- Set your GitHub access token in environment
```
export GITHUB_TOKEN="YOUR_GH_TOKEN"
```


- Run command from `mason` folder after making sure **repo working tree is clean** and **`dist` folder from earlier build is removed**
```
goreleaser release
```

- This updates the .rb script in Formula directory which is available as a Formula from this tap

## INSTALLATION

- Tap the homebrew repo at https://github.com/cirbuk/homebrew-mason

```
brew tap cirbuk/mason
```

- Install the formula

```
brew install homebrew-mason
```



