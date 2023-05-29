package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"os/exec"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить файл .env")
		return
	}

	args := func() []string {
		tmp := make([]string, 0)
		for _, v := range os.Args[1:] {
			if v != "" {
				tmp = append(tmp, v)
			}
		}
		return tmp
	}()

	if len(args) == 1 && args[0] == "-c" {
		creationScript()
	} else if len(args) == 2 && args[0] == "-r" && args[1] != "" {
		recoveryScript(args[1])
	} else {
		fmt.Println("error")
	}
}

func creationScript() {
	currenTime := time.Now()
	dumpName := fmt.Sprintf("dump_date_%d-%02d-%02d_%02d-%02d-%02d.sql",
		currenTime.Year(),
		currenTime.Month(),
		currenTime.Day(),
		currenTime.Hour(),
		currenTime.Minute(),
		currenTime.Second(),
	)
	if err := createDump(dumpName); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Дамп базы данных успешно создан c именем %s.", dumpName)

	if err := copyContainerToHost(dumpName); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Дамп базы данных успешно скопирован на хост машину.")
}

func createDump(dumpName string) error {
	cmd := exec.Command("docker", "exec", os.Getenv("DB_CONTAINER_NAME"),
		"pg_dump", "-U", os.Getenv("DB_USER"), "-d", os.Getenv("DB_NAME"), "-f", dumpName)

	fmt.Println(cmd.Args, os.Getenv("DB_USER"), 123)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func copyContainerToHost(dumpName string) error {
	pathDump := "backup/db/" + dumpName
	containerPathDump := fmt.Sprint(os.Getenv("DB_CONTAINER_NAME"), ":", dumpName)
	cmd := exec.Command("docker", "cp", containerPathDump, pathDump)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func recoveryScript(dumpName string) {
	if err := copyHostToContainer(dumpName); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Дамп базы данных успешно скопирован в контейнер.")

	if err := restoreDB(dumpName); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Восстановление базы данных прошло успешно.")
}

func copyHostToContainer(dumpName string) error {
	pathDump := "backup/db/" + dumpName
	containerPathDump := fmt.Sprint(os.Getenv("DB_CONTAINER_NAME"), ":", dumpName)
	cmd := exec.Command("docker", "cp", pathDump, containerPathDump)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func restoreDB(dumpName string) error {
	cmd := exec.Command("docker", "exec", os.Getenv("DB_CONTAINER_NAME"), "bash", "-c",
		fmt.Sprintf("psql -U %s -d %s -f %s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			dumpName,
		))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
