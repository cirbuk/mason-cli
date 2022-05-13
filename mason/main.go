package main

import (
	"encoding/json"
	"flag"
	"os"
)

var ACTIVE_TOKEN = "MASON_ACTIVE_TOKEN"

var CONFIG_PATH = "/etc/mason_conf.json"

type Masonconfig struct {
	MasonToken string
	MasonEnv   string
}

func main() {
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
	outputProject := projectCmd.String("o", "", "Relative path to write output to")

	schemaCmd := flag.NewFlagSet("schema", flag.ExitOnError)
	actionSchema := schemaCmd.String("action", "", "Please  enter action from [create,update,delete,get]")
	allSchema := schemaCmd.Bool("all", false, "Get all schemas")
	idSchema := schemaCmd.String("id", "", "Mason Schema ID")
	pathSchemaJson := schemaCmd.String("path", "", "Relative file path of JSON for schema create/update")
	outputSchema := schemaCmd.String("o", "", "Relative path to write output to")

	contentCmd := flag.NewFlagSet("content", flag.ExitOnError)
	actionContent := contentCmd.String("action", "", "Please  enter action from [get]")
	allContent := contentCmd.Bool("all", false, "Get all content")
	idContent := contentCmd.String("id", "", "Mason Content ID")
	outputContent := contentCmd.String("o", "", "Relative path to write output to")

	if len(os.Args) < 2 {
		println("expected a subcommand from ['init', 'login', 'logout', 'project', 'schema', 'content']")
		os.Exit(1)
	}

	//look at the 2nd argument's value
	switch os.Args[1] {
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
		println("Invalid action. Please enter action out of ['init,login,logout,project,schema,content']")
	}

}

func handleInit(initCmd *flag.FlagSet) {
	initCmd.Parse(os.Args[2:])
	// redirect to Mason devhub here
	println("Initiated auth flow")
	println("Redirecting to Mason DevHub where you can get your token..")
}

func handleLogout(logoutCmd *flag.FlagSet) {
	logoutCmd.Parse(os.Args[2:])
	println("Initiated Logout Flow")
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
