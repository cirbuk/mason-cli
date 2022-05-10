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

- Tap the homebrew repo at https://github.com/cirbuk/mason

```
brew tap cirbuk/mason
```

- Install the formula

```
brew install mason
```

## USAGE

- Signup for Mason and get your access token
```
mason init
```

- Login to mason-cli with your token and environment
```
sudo mason login -env [prod|beta] -token [user JWT token for a workspace from init step]
```

Example :
```
sudo mason login -env beta -token <api-token>
```

- Logout active user from Mason cli
```
sudo mason logout
```

- Entity related usage :
You can use commands using the below pattern
```
mason [project|schema|content] -action [create|update|get|delete] {specific flags and their values  - all|id|path}
```

- Project related usage :
    - Create a project
    ```
    mason project -action create -path {path_to_file_with_JSON_payload}
    ```

    - Update a project
    ```
    mason project -action update -path {path_to_file_with_JSON_payload}
    ```

    - Get a project
    ```
    mason project -action get -id {project_id}
    ```

    - Get all projects
    ```
    mason project -action get -all true
    ```

    - Delete a project
    ```
    mason project -action delete -id {project_id}
    ```

- Schema related usage :
    - Create a schema
    ```
    mason schema -action create -path {path_to_file_with_JSON_payload}
    ```

    - Update a schema
    ```
    mason schema -action update -path {path_to_file_with_JSON_payload}
    ```

    - Get a schema
    ```
    mason schema -action get -id {schema_id}
    ```

    - Get all schemas
    ```
    mason schema -action get -all true
    ```

    - Delete a schema
    ```
    mason schema -action delete -id {schema_id}
    ```

- Content related usage :
    - Get content by id
    ```
    mason content -action get -id {content_id}
    ```

    - Get all content
    ```
    mason content -action get -all true
    ```





