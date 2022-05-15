package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
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
	actionProject := projectCmd.String("action", "", "Please  enter action from [create,update,delete,get]")
	allProject := projectCmd.Bool("all", false, "Get all projects")
	idProject := projectCmd.String("id", "", "Mason Project ID")
	pathProjectJson := projectCmd.String("path", "", "Relative file path of JSON for project create/update")
	outputProject := projectCmd.String("out", "", "Relative path to write output to")

	schemaCmd := flag.NewFlagSet("schema", flag.ExitOnError)
	actionSchema := schemaCmd.String("action", "", "Please  enter action from [create,update,delete,get]")
	allSchema := schemaCmd.Bool("all", false, "Get all schemas")
	idSchema := schemaCmd.String("id", "", "Mason Schema ID")
	pathSchemaJson := schemaCmd.String("path", "", "Relative file path of JSON for schema create/update")
	outputSchema := schemaCmd.String("out", "", "Relative path to write output to")

	contentCmd := flag.NewFlagSet("content", flag.ExitOnError)
	actionContent := contentCmd.String("action", "", "Please  enter action from [get]")
	allContent := contentCmd.Bool("all", false, "Get all content")
	idContent := contentCmd.String("id", "", "Mason Content ID")
	outputContent := contentCmd.String("out", "", "Relative path to write output to")

	if len(os.Args) < 2 {
		handleHelp(helpCmd)
		// println("expected a subcommand from ['help', init', 'login', 'logout', 'project', 'schema', 'content']")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
	case "help": // if its the 'help' command
		handleHelp(helpCmd)
	case "init": // if its the 'init' command
		handleInit(initCmd)
	case "login": // if its the 'login' command
		handleLogin(loginCmd, loginToken, env)
	case "logout": // if its the 'logout' command
		handleLogout(logoutCmd)
	case "project": // if its the 'project' command
		handleProject(projectCmd, actionProject, allProject, idProject, pathProjectJson, outputProject)
	case "schema": // if its the 'schema' command
		handleSchema(schemaCmd, actionSchema, allSchema, idSchema, pathSchemaJson, outputSchema)
	case "content": // if its the 'content' command
		handleContent(contentCmd, actionContent, allContent, idContent, outputContent)
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

func IsValidCommand(cmd string) bool {
	switch cmd {
	case
		"project",
		"content",
		"schema":
		return true
	}
	return false
}

// TODO : redirect to Mason devhub here
func handleDocs(initCmd *flag.FlagSet) {
	// content := readFileJson("usage.txt")
	// println(string(content))
}

func handleHelp(helpCmd *flag.FlagSet) {
	content := readFileJson("usage.txt")
	if len(os.Args) <= 2 {
		println(string(content))
	} else {
		cmd := os.Args[2]
		if IsValidCommand(cmd) {
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

func handleProject(projectCmd *flag.FlagSet, action *string, all *bool, projectId *string, contentPath *string, outputPath *string) {
	projectCmd.Parse(os.Args[2:])
	if *action == "" {
		println("Enter action as ['create', 'update', 'delete', 'get']")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	if *action == "get" && *all == false && *projectId == "" {
		println("all or id needs to be sent for get action")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	if *action == "delete" && *projectId == "" {
		println("id needs to be sent for delete action")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	if (*action == "create" || *action == "update") && *contentPath == "" {
		println("Path with JSON payload required for actions create,update")
		projectCmd.PrintDefaults()
		os.Exit(1)
	}
	println("Initiated Project Flow")
	println(*action)
	println(*outputPath)
	if *action == "get" {
		getProjects(*projectId, *all, *outputPath)
	}
	if *action == "delete" {
		deleteProject(*projectId, *outputPath)
	}
	if *action == "create" || *action == "update" {
		createOrUpdateProject(*contentPath, *outputPath)
	}

}

func handleSchema(schemaCmd *flag.FlagSet, action *string, all *bool, schemaId *string, contentPath *string, outputPath *string) {
	schemaCmd.Parse(os.Args[2:])
	if *action == "" {
		println("Enter action as ['create', 'update', 'delete', 'get']")
		schemaCmd.PrintDefaults()
		os.Exit(1)
	}
	if *action == "get" && *all == false && *schemaId == "" {
		println("all or id needs to be sent for get action")
		schemaCmd.PrintDefaults()
		os.Exit(1)
	}
	if *action == "delete" && *schemaId == "" {
		println("id needs to be sent for delete action")
		schemaCmd.PrintDefaults()
		os.Exit(1)
	}
	if (*action == "create" || *action == "update") && *contentPath == "" {
		println("Path with JSON payload required for actions create,update")
		schemaCmd.PrintDefaults()
		os.Exit(1)
	}
	println("Initiated Schema Flow")
	println(*action)
	if *action == "get" {
		getSchema(*schemaId, *all, *outputPath)
	}
	if *action == "delete" {
		deleteSchema(*schemaId, *outputPath)
	}
	if *action == "create" {
		createSchema(*contentPath, *outputPath)
	}
	if *action == "update" {
		updateSchema(*contentPath, *outputPath)
	}
}

func handleContent(contentCmd *flag.FlagSet, action *string, all *bool, contentId *string, outputPath *string) {
	contentCmd.Parse(os.Args[2:])
	if *action == "" {
		println("Enter action as ['get']")
		contentCmd.PrintDefaults()
		os.Exit(1)
	}
	if *action == "get" && *all == false && *contentId == "" {
		println("all or id needs to be sent for get action")
		contentCmd.PrintDefaults()
		os.Exit(1)
	}
	println("Initiated Content Flow")
	println(*action)
	if *action == "get" {
		getContent(*contentId, *all, *outputPath)
	}

}
