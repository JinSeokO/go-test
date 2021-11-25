package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func execDockerMysqlSqlDump(containerName string, user string, password string, databaseName string)  {
	cmd := exec.Command("docker", "exec", containerName, "usr/bin/mysqldump",
		"-u", user, fmt.Sprintf("--password=%s", password),
		"--single-transaction=TRUE", "--no-data", "--set-gtid-purged=OFF",
		"--databases", databaseName)

	stdoutPipe, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}


	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}


	bytes, err := io.ReadAll(stdoutPipe)

	if err != nil {
		log.Fatalln(err)
	}
	_, filename, _, _ := runtime.Caller(1)

	dir := path.Join(path.Dir(filename))
	if 	err := os.Chdir(dir); err != nil {
		panic(err)
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s/scheme/%s.sql", dir, databaseName), bytes, 0644); err != nil {
		log.Fatalln(err)
	}
}
func main() {
	user := flag.String("user", "root", "db user name")
	password := flag.String("password", "adfaie83ma", "db password")
	databaseName := flag.String("database", "golang-test", "database name")
	containerName := flag.String("container", "docker_mysql-db_1", "docker container name")

	execDockerMysqlSqlDump(*containerName, *user, *password, *databaseName)
}