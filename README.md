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
sudo mason login -env beta -token eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ0b2tlbl90eXBlIjoiZmVlZCIsInVpZCI6IjFlZWM4Y2MwLThiZTMtNGUwOC1hMWJmLTE4MzdmNmVmYTBhNCIsIndvcmtzcGFjZV9pZCI6ImZjZTJiNGQxLTI2NTMtNGU2OC1iMDk0LTNhYmYwMDdkMjRkZSIsInVzZXJpZCI6InNhdXJhYi5taXRyYUBrdWJyaWMuaW8ifQ.7lxQ6vaf2WiITbfRN_1-sM1KVQdDUV2BfIPcUvWhYmU
```

- Logout active user from Mason cli
```
sudo mason logout
```

- Entity related usage :
    In general the commands are of the form
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

## SAMPLE USAGE

- To create schema for carousel post in Online Shop

Command : `mason schema -action create -path payload.json`

Contents of payload.json file :

```
   {
  "content_type": {
    "title": "StoreTopCarousel",
    "name": "StoreTopCarousel",
    "type": "document",
    "scope": "workspace",
    "fields": [
      {
        "title": "Product Text",
        "name": "product_text",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Product Image",
        "name": "image",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Description",
        "name": "desc",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Background Color",
        "name": "bg_color",
        "field_type": "string",
        "of": null,
        "options": {}
      }
    ]
  }
}
```

- To update the above created content, adding Contrast color field

Command : `mason schema -action update -path payload.json`

Contents of payload.json file :

```
{
  "content_type": {
    "id": "fe85ac4c-66d1-456e-996a-3cb55cf76cf5",
    "title": "StoreTopCarousel",
    "name": "StoreTopCarousel",
    "type": "document",
    "workspace_id": "fe85ac4c-66d1-456e-996a-3cb55cf76cf5",
    "created_by": "saurab.mitra@kubric.io",
    "scope": "workspace",
    "fields": [
      {
        "title": "Product Text",
        "name": "product_text",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Product Image",
        "name": "image",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Description",
        "name": "desc",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Background Color",
        "name": "bg_color",
        "field_type": "string",
        "of": null,
        "options": {}
      },
      {
        "title": "Contrast Color",
        "name": "cs_color",
        "field_type": "string",
        "of": null,
        "options": {}
      }
    ]
  }
}
```







