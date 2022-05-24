package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var API_HOST_BETA = "https://genie-beta.kubric.io"
var API_HOST_PROD = "https://genie.kubric.io"

func getApiHost() string {
	var API_HOST string
	config := checkToken()
	env := config.MasonEnv
	if env == "beta" {
		API_HOST = API_HOST_BETA
	} else if env == "prod" {
		API_HOST = API_HOST_PROD
	}
	return API_HOST
}

func checkToken() Masonconfig {
	var config Masonconfig
	CONFIG_PATH := getConfigPath()
	configBytes := readFileJson(CONFIG_PATH)
	err := json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Printf("\nNo token found, please complete mason init and mason login first\n")
		os.Exit(1)
	}
	return config
}

func readFileJson(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("failed reading data from file: %s", err)
	}
	return data
}

func writeFileJson(path string, content []byte) bool {
	file, err := os.Create(path)

	if err != nil {
		fmt.Printf("\nfailed creating file: %s\n", err)
	}
	defer file.Close()
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		fmt.Printf("\nfailed writing to file: %s\n", err)
		return false
	}
	return true
}

func exportProject(projectId string, output string) {
	var outputSchema = ""
	var outputContent = ""
	var projectPath = ""
	if output != "" {
		if err := os.MkdirAll(output, os.ModePerm); err != nil {
			fmt.Printf("\nError making directory %s\n", output)
		}
		projectPath = output + "/" + projectId
		schemaPath := projectPath + "/schemas"
		contentPath := projectPath + "/content"
		if err := os.MkdirAll(schemaPath, os.ModePerm); err != nil {
			fmt.Printf("\nError making directory %s\n", schemaPath)
		}
		if err := os.MkdirAll(contentPath, os.ModePerm); err != nil {
			fmt.Printf("\nError making directory %s\n", contentPath)
		}
		outputSchema = schemaPath + "/schemas.json"
		outputContent = contentPath + "/content.json"
	}
	getSchema("", false, outputSchema, projectId)
	getContent("", false, outputContent, projectId)
	if output != "" {
		fmt.Printf("Wrote Project Schema and Content details to %s\n", projectPath)
	}
}

func importProject(projectPath string, schemaOnly bool, contentOnly bool) {
	if !schemaOnly && !contentOnly {
		schemaOnly = true
		contentOnly = true
	}
	if schemaOnly {
		importSchemas(projectPath)
		fmt.Printf("\nSuccessfully imported Schemas at %s\n", projectPath)
	}
	if contentOnly {
		importContent(projectPath)
		fmt.Printf("\nSuccessfully imported Content at %s\n", projectPath)
	}
	fmt.Println("Successfully finished import Project")

}

func importSchemas(projectPath string) {
	var schemasDir = projectPath + "/schemas"
	files, err := ioutil.ReadDir(schemasDir)
	if err != nil {
		fmt.Printf("\nError reading directory %s\n", schemasDir)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		schemaPath := schemasDir + "/" + f.Name()
		createSchema(schemaPath, "")
	}
	fmt.Printf("\nCreated Schemas present in %s\n", schemasDir)
}

func importContent(projectPath string) {
	var contentDir = projectPath + "/content"
	files, err := ioutil.ReadDir(contentDir)
	if err != nil {
		fmt.Printf("\nError reading directory %s\n", contentDir)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		contentPath := contentDir + "/" + f.Name()
		createOrUpdateContent(contentPath, "")
	}
	fmt.Printf("\nCreated Content present in %s\n", contentDir)

}

func getProjects(projectId string, output string) {
	config := checkToken()
	token := config.MasonToken
	fmt.Printf("Called Get Projects with token %s projectId %s and outputPath %s\n", token, projectId, output)
	client := &http.Client{}
	endpoint := getApiHost() + "/v1/project"
	req, err := http.NewRequest("GET", endpoint, nil)
	q := req.URL.Query()
	q.Add("token", token)
	if projectId != "" {
		q.Add("project_id", projectId)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Get Projects Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func createOrUpdateProject(contentPath string, output string) {
	config := checkToken()
	token := config.MasonToken
	client := &http.Client{}
	payload := readFileJson(contentPath)
	endpoint := getApiHost() + "/v1/project"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	q := req.URL.Query()
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Create/Update Project Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func deleteProject(projectId string, output string) {
	config := checkToken()
	token := config.MasonToken
	fmt.Printf("Called Delete Project with token %s projectId %s\n", token, projectId)
	client := &http.Client{}
	endpoint := getApiHost() + "/v1/project"
	req, err := http.NewRequest("DELETE", endpoint, nil)
	q := req.URL.Query()
	q.Add("token", token)
	if projectId != "" {
		q.Add("project_id", projectId)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Delete Projects Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func getContent(contentId string, all bool, output string, projectId string) {
	config := checkToken()
	token := config.MasonToken
	fmt.Printf("Called Get Content with token %s contentId %s and all %t\n", token, contentId, all)
	client := &http.Client{}
	endpoint := getApiHost() + "/v1/content"
	req, err := http.NewRequest("GET", endpoint, nil)
	q := req.URL.Query()
	q.Add("token", token)
	if contentId != "" {
		q.Add("content_id", contentId)
	}
	if projectId != "" {
		q.Add("project_id", projectId)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Get Content Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func createOrUpdateContent(contentPath string, output string) {
	config := checkToken()
	token := config.MasonToken
	client := &http.Client{}
	payload := readFileJson(contentPath)
	endpoint := getApiHost() + "/v1/content"
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(payload))
	q := req.URL.Query()
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Create/Update Content Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func getSchema(schemaId string, all bool, output string, projectId string) {
	config := checkToken()
	token := config.MasonToken
	fmt.Printf("Called Get Schema with token %s schemaId %s and all %t\n", token, schemaId, all)
	client := &http.Client{}
	endpoint := getApiHost() + "/v1/schema"
	req, err := http.NewRequest("GET", endpoint, nil)
	q := req.URL.Query()
	q.Add("token", token)
	if schemaId != "" {
		q.Add("schema_id", schemaId)
	}
	if projectId != "" {
		q.Add("project_id", projectId)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Get Schema Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func createSchema(contentPath string, output string) {
	config := checkToken()
	token := config.MasonToken
	client := &http.Client{}
	payload := readFileJson(contentPath)
	endpoint := getApiHost() + "/v1/schema"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	q := req.URL.Query()
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Create Schema Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func updateSchema(contentPath string, output string) {
	config := checkToken()
	token := config.MasonToken
	client := &http.Client{}
	payload := readFileJson(contentPath)
	endpoint := getApiHost() + "/v1/schema"
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(payload))
	q := req.URL.Query()
	q.Add("token", token)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Create Schema Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}

func deleteSchema(schemaId string, output string) {
	config := checkToken()
	token := config.MasonToken
	fmt.Printf("Called Delete Project with token %s schemaId %s\n", token, schemaId)
	client := &http.Client{}
	endpoint := getApiHost() + "/v1/schema"
	req, err := http.NewRequest("DELETE", endpoint, nil)
	q := req.URL.Query()
	q.Add("token", token)
	if schemaId != "" {
		q.Add("schema_id", schemaId)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while making Delete Schema Call")
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	var prettyJSON bytes.Buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	json.Indent(&prettyJSON, body, "", "\t")
	if output == "" {
		fmt.Println(resp.Status)
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Printf("\nWriting output to %s\n", output)
		writeFileJson(output, prettyJSON.Bytes())
	}
}
