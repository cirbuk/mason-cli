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

## USAGE

- Signup for Mason and get your access token
```
homebrew-mason init
```

- Login to mason-cli with your token and environment
```
sudo homebrew-mason login -env [prod|beta] -token [user JWT token for a workspace from init step]
```

Example :
```
sudo homebrew-mason login -env beta -token eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiZmVlZCIsInVpZCI6IjFlZWM4Y2MwLThiZTMtNGUwOC1hMWJmLTE4MzdmNmVmYTBhNCIsIndvcmtzcGFjZV9pZCI6ImZjZTJiNGQxLTI2NTMtNGU2OC1iMDk0LTNhYmYwMDdkMjRkZSIsInVzZXJpZCI6InNhdXJhYi5taXRyYUBrdWJyaWMuaW8ifQ.7lxQ6vaf2WiITbfRN_1-sM1KVQdDUV2BfIPcUvWhYmU
```

- Logout active user from Mason cli
```
sudo homebrew-mason logout
```

- Entity related usage :
    In general the commands are of the form
```
homebrew-mason [project|schema|content] -action [create|update|get|delete] {specific flags and their values  - all|id|path}
```

- Project related usage :
    - Create a project
    ```
    homebrew-mason project -action create -path {path_to_file_with_JSON_payload}
    ```

    - Update a project
    ```
    homebrew-mason project -action update -path {path_to_file_with_JSON_payload}
    ```

    - Get a project
    ```
    homebrew-mason project -action get -id {project_id}
    ```

    - Get all projects
    ```
    homebrew-mason project -action get -all true
    ```

    - Delete a project
    ```
    homebrew-mason project -action delete -id {project_id}
    ```

- Schema related usage :
    - Create a schema
    ```
    homebrew-mason schema -action create -path {path_to_file_with_JSON_payload}
    ```

    - Update a schema
    ```
    homebrew-mason schema -action update -path {path_to_file_with_JSON_payload}
    ```

    - Get a schema
    ```
    homebrew-mason schema -action get -id {schema_id}
    ```

    - Get all schemas
    ```
    homebrew-mason schema -action get -all true
    ```

    - Delete a schema
    ```
    homebrew-mason schema -action delete -id {schema_id}
    ```

- Content related usage :
    - Get content by id
    ```
    homebrew-mason content -action get -id {content_id}
    ```

    - Get all content
    ```
    homebrew-mason content -action get -all true
    ```





