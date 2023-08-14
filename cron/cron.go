package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Define the cron schedule and command.
	cronSchedule := "*/1 * * * *" // Runs every 5 minutes
	cronCommand := os.Args[1]     // Replace with the actual script/command

	// Validate the cron schedule (you can use a more robust validation)
	if cronSchedule == "" {
		log.Println("Cron schedule is empty.")
		return
	}

	// Validate the cron command (you can use a more robust validation)
	if cronCommand == "" {
		log.Println("Cron command is empty.")
		return
	}

	// Prepare the crontab command.
	crontabCmd := exec.Command("crontab", "-l")
	currentCronJobs, err := crontabCmd.Output()

	// Check for the specific error that indicates no cron jobs are set up.
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) && strings.Contains(string(exitErr.Stderr), "no crontab for") {
		// No cron jobs are set up, so we can safely ignore this error.
		err = nil
	}

	if err != nil {
		log.Println("Error listing current cron jobs:", err)
		return
	}

	// Check if the cron job already exists.
	if string(currentCronJobs) == cronSchedule+" "+cronCommand {
		log.Println("Cron job already exists.")
		return
	}

	if strings.Contains(string(currentCronJobs), cronCommand) {
		log.Println("Cron job already exists")
		return
	}

	// Append the new cron job to the existing ones.
	newCronJobs := append(currentCronJobs, []byte(cronSchedule+" "+cronCommand+"\n")...)

	// Write the new cron jobs to the crontab.
	writeCrontabCmd := exec.Command("crontab")
	writeCrontabCmd.Stdin = bytes.NewReader(newCronJobs)
	writeCrontabCmd.Stdout = os.Stdout
	writeCrontabCmd.Stderr = os.Stderr

	// Run the command to write the new cron jobs.
	if err := writeCrontabCmd.Run(); err != nil {
		fmt.Println("Error running crontab command:", err)
		return
	}

	fmt.Println("Cron job successfully added.")
}
