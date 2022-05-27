package main

const usageText = "Mason's command line app\n\nUSAGE:\n  mason [COMMAND] [import|export] [OPTIONS]\n\nCOMMANDS:\n  project  Interact with projects connected to your logged in user\n  docs     Open Mason dev center for building apps\n  help     Displays help info about mason\n  init     Get your Mason Developer Token\n  login    Login Mason cli with a token and particular environment\n  logout   Logout current user from Mason cli\n\nSee 'mason help [COMMAND]' for specific information on a subcommand."

const docsHelp = "USAGE:\n- Open Mason dev center for building apps\n   mason docs"

const initHelp = "USAGE:\n- Signup for Mason and get your access token\n   mason init\n"

const loginHelp = "USAGE:\n- Login to mason-cli with your token and environment\n\n   mason login -env [prod|beta] -token [user JWT token for a workspace from init step]\n\n- Login usage:\n    mason login -env beta -token <api-token>\n"

const logoutHelp = "USAGE:\n- Logout active user from Mason cli\n   mason logout\n"

const projectHelp = "USAGE: \n   mason project -action [OPTIONS]\n\n   Initialize a new mason project or app\n\nmason [project] [export|import] {specific flags and their values  - id|out|path|schema|content}\n\n- Project related usage :\n\n    - Get all projects\n\n    mason project get -id {project_id}\n\n\n    - Export schemas and content linked to a project\n\n    mason project export -id {project_id}\n\n\n\n    -  Write export operation output to folder\n\n    mason project export -id {project_id} -out {output_folder_path}\n\n    The resultant folder strucutre is of the form :\n\n    output_folder_path\n    +-- project_id\n    ¦   +-- content\n            +-- content.json\n    ¦   +-- schemas\n            +-- schemas.json\n\n    - Import schemas and content linked to a project into mason\n\n    mason project import -path {path_to_project_folder}\n\n    The project folder should contain `schemas` and `content` folders with individual JSON files inside them\n\n\n    - Import schemas and content selectively\n\n    mason project import -path {path_to_project_folder} --content true\n\n    mason project import -path {path_to_project_folder} --schema true\n"

func getHelpMessage(cmd string) string {
	switch cmd {
	case "usage": // if its the 'help' command
		return string(usageText)
	case "docs": // if its the 'docs' command
		return string(docsHelp)
	case "init": // if its the 'init' command
		return string(initHelp)
	case "login": // if its the 'login' command
		return string(loginHelp)
	case "logout": // if its the 'logout' command
		return string(logoutHelp)
	case "project": // if its the 'project' command
		return string(projectHelp)
	default: // if we don't understand the input
		return "error. See mason help"
	}
}
