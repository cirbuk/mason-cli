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
mason [project] [export|import] {specific flags and their values  - id|out|path|schema|content}
```

- Project related usage :

    - Get all projects
    ```
    mason project get -id {project_id}
    ```

    - Export schemas and content linked to a project
    ```
    mason project export -id {project_id}
    ```


    -  Write export operation output to folder
    ```
    mason project export -id {project_id} -out {output_folder_path}
    ```
    The resultant folder strucutre is of the form :

    ```
    output_folder_path
    +-- project_id
    ¦   +-- content
            +-- content.json
    ¦   +-- schemas
            +-- schemas.json
    ```

    - Import schemas and content linked to a project into mason
    ```
    mason project import -path {path_to_project_folder}
    ```
    The project folder should contain `schemas` and `content` folders with individual JSON files inside them


    - Import schemas and content selectively
    ```
    mason project import -path {path_to_project_folder} --content true
    ```

    ```
    mason project import -path {path_to_project_folder} --schema true
    ```



## Sample Usage

- To create schema for carousel post in Online Shop

Command : `mason project import -path testdir`

Contents of testdir/content/content1.json file :

```
{
    "content": [
        {
            "content_type": {
                "name": "default_project_creative",
                "id": "9a5cce5b-2102-4ad4-827e-fc2eab278042"
            },
            "name": "slp_groc_1",
            "campaign_id": "ad0b1e36-bc3d-4c6b-942b-81de8666019f",
            "creative_id": "20eaa3e2-aadd-4b6b-be24-21c71a47c777",
            "data": {
                "uid": "20eaa3e2-aadd-4b6b-be24-21c71a47c777",
                "name": "slp_groc_1",
                "status": "created",
                "created_on": "2022-01-28T07:01:59.769737",
                "created_by": "maaz@kubric.io",
                "updated": "2022-01-28T07:01:59.770141",
                "updated_by": "maaz@kubric.io",
                "sb_preview": "https://media.kubric.io/api/videos-kubric/image_png1141161121.png",
                "mediaFormat": "image/png"
            }
        }
    ]
}
```


Contents of testdir/schemas/schema1.json file :

```
{
    "content_type": {
    	"id": "07284c89-012c-40ab-b607-22aa51e9a42b",
        "title": "Testsaurab",
        "name": "testsaurab",
        "type": "document",
        "scope": "workspace",
        "fields": [
            {
                "title": "ID",
                "name": "uid",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Title",
                "name": "name",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Status",
                "name": "status",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Created On",
                "name": "created_on",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Created By",
                "name": "created_by",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Updated At",
                "name": "updated",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Updated By",
                "name": "updated_by",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Preview",
                "name": "sb_preview",
                "field_type": "string",
                "of": null,
                "options": {}
            },
            {
                "title": "Media Format",
                "name": "mediaFormat",
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






