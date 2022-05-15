<div align="center">
  <img src="assets/logo.png" width="150"/>
  <h1>Mason CLI</h1>
</div>
Mason CLI helps you build Mason apps & plugins. Use Mason CLI to supercharge your plugin development workflow.

Mason CLI can be run and installed on Mac, Linux and Windows systems.


## Development

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

## Installation
Check out the [installation steps](docs/users/installation.md).

## Usage

- Signup for Mason and get your access token
```
mason init
```

- Login to mason-cli with your token and environment
```
mason login -env [prod|beta] -token [user JWT token for a workspace from init step]
```

Example :
```
mason login -env beta -token <api-token>
```

- Logout active user from Mason cli
```
mason logout
```

- Entity related usage :
You can use commands using the below pattern
```
mason [project|schema|content] -action [create|update|get|delete] {specific flags and their values  - all|id|out|path}
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
    -  Write operation output to file
    ```
    mason project -action delete -out {output_file_path} -id {project_id}
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
    -  Write operation output to file
    ```
    mason project -action delete -out {output_file_path} -id {project_id}
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
    -  Write operation output to file
    ```
    mason project -action delete -out {output_file_path} -id {project_id}
    ```

## Sample Usage

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

## Help 🖐

If you encounter issues using the CLI or have feedback you'd like to share with us, below are some options:

- [Open a GitHub issue](https://github.com/cirbuk/mason-cli/issues) - To report bugs or request new features, open an issue in the Mason CLI repository.






