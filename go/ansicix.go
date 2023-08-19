package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var podName string
	var minute string
	const path = "../ansible/playbooks"  // Sabit path değeri

	var rootCmd = &cobra.Command{
		Use:   "ansicix",
		Short: "Ansible runner with Cobra",
		Long:  `Run Ansible playbooks easily with this tool.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Ansicix!")
		},
	}

	var setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Run Ansible playbooks from the specified directory",
		Run: func(cmd *cobra.Command, args []string) {
			runAnsiblePlaybooks(path, podName, minute)
		},
	}

	// Pod ismi için flag tanımlama
	setupCmd.Flags().StringVarP(&podName, "pod", "p", "", "Pod name for the Kubernetes task")
	// Dakika aralığı için flag tanımlama
	setupCmd.Flags().StringVarP(&minute, "minute", "m", "1", "Minute interval for the cron job")
	setupCmd.MarkFlagRequired("pod")

	rootCmd.AddCommand(setupCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runAnsiblePlaybooks(path, podName, minute string) {
	cmd := exec.Command("ansible-playbook", "-i", "inventory", fmt.Sprintf("%s/cron.yml", path), "-e", fmt.Sprintf("POD_NAME=%s minute=%s", podName, minute))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run the playbook: %v", err)
	}
}
