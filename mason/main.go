package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

var ACTIVE_TOKEN = "MASON_ACTIVE_TOKEN"

type Masonconfig struct {
	MasonToken string
	MasonEnv   string
}

func main() {

	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}

	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)

	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

	loginCmd := flag.NewFlagSet("login", flag.ExitOnError)
	loginToken := loginCmd.String("token", "", "Mason DevHub Token")
	env := loginCmd.String("env", "", "Please enter env as 'beta' or 'prod'")

	logoutCmd := flag.NewFlagSet("logout", flag.ExitOnError)

	projectCmd := flag.NewFlagSet("project", flag.ExitOnError)
	// 	actionProject := projectCmd.String("action", "", "Please  enter action from [create,update,delete,get]")
	// 	allProject := projectCmd.Bool("all", false, "Get all projects")
	schemaFlag := projectCmd.Bool("schema", false, "Create/Update only schemas from project folder")
	contentFlag := projectCmd.Bool("content", false, "Create/Update only content from project folder")
	idProject := projectCmd.String("id", "", "Mason Project ID")
	pathProjectJson := projectCmd.String("path", "", "Relative path of project folder containing Content and Schemas")
	outputProject := projectCmd.String("out", "", "Relative folder/directory path to write project details")

	// 	schemaCmd := flag.NewFlagSet("schema", flag.ExitOnError)
	// 	actionSchema := schemaCmd.String("action", "", "Please  enter action from [create,update,delete,get]")
	// 	allSchema := schemaCmd.Bool("all", false, "Get all schemas")
	// 	idSchema := schemaCmd.String("id", "", "Mason Schema ID")
	// 	pathSchemaJson := schemaCmd.String("path", "", "Relative file path of JSON for schema create/update")
	// 	outputSchema := schemaCmd.String("out", "", "Relative path to write output to")
	//
	// 	contentCmd := flag.NewFlagSet("content", flag.ExitOnError)
	// 	actionContent := contentCmd.String("action", "", "Please  enter action from [get]")
	// 	allContent := contentCmd.Bool("all", false, "Get all content")
	// 	idContent := contentCmd.String("id", "", "Mason Content ID")
	// 	outputContent := contentCmd.String("out", "", "Relative path to write output to")

	if len(os.Args) < 2 {
		handleHelp(helpCmd)
		// println("expected a subcommand from ['help', init', 'login', 'logout', 'project', 'schema', 'content']")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "help": // if its the 'help' command
		handleHelp(helpCmd)
	case "docs": // if its the 'docs' command
		openDocs()
	case "init": // if its the 'init' command
		handleInit(initCmd)
	case "login": // if its the 'login' command
		handleLogin(loginCmd, loginToken, env)
	case "logout": // if its the 'logout' command
		handleLogout(logoutCmd)
	case "project": // if its the 'project' command
		handleProject(projectCmd, idProject, pathProjectJson, outputProject, schemaFlag, contentFlag)
		// 	case "schema": // if its the 'schema' command
		// 		handleSchema(schemaCmd, actionSchema, allSchema, idSchema, pathSchemaJson, outputSchema)
		// 	case "content": // if its the 'content' command
		// 		handleContent(contentCmd, actionContent, allContent, idContent, outputContent)
	default: // if we don't understand the input
		println("error. See mason help")
	}

}

func getConfigPath() string {
	HOME_USER_DIR, err := os.UserHomeDir()
	if err != nil {
		println("Couldn't get user home dir to set config")
		os.Exit(1)
	}
	var CONFIG_PATH = HOME_USER_DIR + "/.mason_conf.json"
	return CONFIG_PATH
}

func ValidCommand(cmd string) bool {
	switch cmd {
	case
		"project",
		"login",
		"logout",
		"init",
		"docs":
		return true
	}
	return false
}

// Redirect to Mason devhub here
func openDocs() {
	var url = "https://getmason.dev"
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func handleHelp(helpCmd *flag.FlagSet) {
	content := readFileJson("usage.txt")
	if len(os.Args) <= 2 {
		println(string(content))
	} else {
		cmd := os.Args[2]
		if ValidCommand(cmd) {
			println(string(readFileJson(fmt.Sprintf("help/%s.txt", cmd))))
		} else {
			println(fmt.Sprintf("%s error: %s %s %s is not a valid mason command. Did you mean help?%s See \"mason help\"", Red, Reset, cmd, Red, Reset))
		}
	}
}

func handleInit(initCmd *flag.FlagSet) {
	initCmd.Parse(os.Args[2:])
	println("Initiated auth flow")
	println("Redirecting to Mason DevHub where you can get your token..")
	openDocs()
}

func handleLogout(logoutCmd *flag.FlagSet) {
	logoutCmd.Parse(os.Args[2:])
	println("Initiated Logout Flow")
	CONFIG_PATH := getConfigPath()
	var config []byte
	bytes, err := json.Marshal(config)
	if err != nil {
		println("Failed to clear mason login credentials")
		println(err)
		os.Exit(1)
	}
	success := writeFileJson(CONFIG_PATH, bytes)
	if !success {
		println("Failed to set mason login credentials")
		println(err)
		os.Exit(1)
	}
	println("Completed logout flow")
}

func handleLogin(loginCmd *flag.FlagSet, loginToken *string, env *string) {
	loginCmd.Parse(os.Args[2:])
	CONFIG_PATH := getConfigPath()
	if *loginToken == "" {
		println("Please enter a valid token from Mason DevHub")
		loginCmd.PrintDefaults()
		os.Exit(1)
	}
	if *env == "" || (*env != "beta" && *env != "prod") {
		println("Please enter a valid env from ['beta','prod']")
		loginCmd.PrintDefaults()
		os.Exit(1)
	}
	println("Initiated Login Flow")
	config := Masonconfig{MasonToken: *loginToken, MasonEnv: *env}
	bytes, err := json.Marshal(config)
	if err != nil {
		println("Failed to set mason login credentials")
		os.Exit(1)
	}
	success := writeFileJson(CONFIG_PATH, bytes)
	if !success {
		println("Failed to set mason login credentials")
		os.Exit(1)
	}
	println("\nSuccessfully set Mason Login Creds")

}

func handleProject(projectCmd *flag.FlagSet, projectId *string, contentPath *string, outputPath *string, schemaOnly *bool, contentOnly *bool) {
	projectCmd.Parse(os.Args[3:])
	var action = os.Args[2]
	if action == "" {
		println("Enter action as ['import', 'export'] for project command")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	if action != "import" && action != "export" && action != "get" {
		println("Invalid action: Enter action from ['import', 'export', 'get'] for project command")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	if action == "export" && *projectId == "" {
		println("id needs to be entered for 'export' action")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	// 	if *action == "delete" && *projectId == "" {
	// 		println("id needs to be sent for delete action")
	// 		projectCmd.PrintDefaults()
	// 		os.Exit(1)
	// 	}
	if (action == "import") && *contentPath == "" {
		println("Path to project folder is required for 'import' action")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	if action == "export" {
		exportProject(*projectId, *outputPath)
	}
	if action == "get" {
		getProjects(*projectId, *outputPath)
	}
	if action == "import" {
		importProject(*contentPath, *schemaOnly, *contentOnly)
	}
	// 	if *action == "delete" {
	// 		deleteProject(*projectId, *outputPath)
	// 	}
	// 	if *action == "create" || *action == "update" {
	// 		createOrUpdateProject(*contentPath, *outputPath)
	// 	}

}

//
// func handleSchema(schemaCmd *flag.FlagSet, action *string, all *bool, schemaId *string, contentPath *string, outputPath *string) {
// 	schemaCmd.Parse(os.Args[2:])
// 	if *action == "" {
// 		println("Enter action as ['create', 'update', 'delete', 'get']")
// 		schemaCmd.PrintDefaults()
// 		os.Exit(1)
// 	}
// 	if *action == "get" && *all == false && *schemaId == "" {
// 		println("all or id needs to be sent for get action")
// 		schemaCmd.PrintDefaults()
// 		os.Exit(1)
// 	}
// 	if *action == "delete" && *schemaId == "" {
// 		println("id needs to be sent for delete action")
// 		schemaCmd.PrintDefaults()
// 		os.Exit(1)
// 	}
// 	if (*action == "create" || *action == "update") && *contentPath == "" {
// 		println("Path with JSON payload required for actions create,update")
// 		schemaCmd.PrintDefaults()
// 		os.Exit(1)
// 	}
// 	println("Initiated Schema Flow")
// 	println(*action)
// 	if *action == "get" {
// 		getSchema(*schemaId, *all, *outputPath)
// 	}
// 	if *action == "delete" {
// 		deleteSchema(*schemaId, *outputPath)
// 	}
// 	if *action == "create" {
// 		createSchema(*contentPath, *outputPath)
// 	}
// 	if *action == "update" {
// 		updateSchema(*contentPath, *outputPath)
// 	}
// }
//
// func handleContent(contentCmd *flag.FlagSet, action *string, all *bool, contentId *string, outputPath *string) {
// 	contentCmd.Parse(os.Args[2:])
// 	if *action == "" {
// 		println("Enter action as ['get']")
// 		contentCmd.PrintDefaults()
// 		os.Exit(1)
// 	}
// 	if *action == "get" && *all == false && *contentId == "" {
// 		println("all or id needs to be sent for get action")
// 		contentCmd.PrintDefaults()
// 		os.Exit(1)
// 	}
// 	println("Initiated Content Flow")
// 	println(*action)
// 	if *action == "get" {
// 		getContent(*contentId, *all, *outputPath)
// 	}
//
// }
