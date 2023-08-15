import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func main() {
	var path string
	var podName string

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
		Short: "Run Ansible playbooks in the provided directory",
		Run: func(cmd *cobra.Command, args []string) {
			runAnsiblePlaybooks(path, podName)
		},
	}

	setupCmd.Flags().StringVarP(&path, "path", "p", "", "Path to directory containing Ansible playbooks")
	setupCmd.Flags().StringVarP(&podName, "pod", "", "", "Pod name for the Kubernetes task")
	setupCmd.MarkFlagRequired("path")

	rootCmd.AddCommand(setupCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func runAnsiblePlaybooks(path, podName string) {
	cmd := exec.Command("ansible-playbook", "-i", "inventory", fmt.Sprintf("%s/playbook.yml", path))
	cmd.Env = append(os.Environ(), fmt.Sprintf("POD_NAME=%s", podName))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to run the playbook: %v", err)
	}
}
