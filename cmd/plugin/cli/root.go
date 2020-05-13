package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dirathea/kubectl-volume-reclaim/pkg/logger"
	"github.com/dirathea/kubectl-volume-reclaim/pkg/plugin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tj/go-spin"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "kubectl-volume-reclaim",
		Short:         "",
		Long:          `.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			log := logger.NewLogger()
			log.Info("")

			s := spin.New()
			finishedCh := make(chan bool, 1)
			output := make(chan string, 1)
			go func() {
				tableOutput := ""
				for {
					select {
					case <-finishedCh:
						fmt.Printf("\r")
						return
					case n := <-output:
						tableOutput = n
					case <-time.After(time.Millisecond * 100):
						if tableOutput == "" {
							fmt.Printf("\r  \033[36mChecking PVC\033[m %s", s.Next())
						} else {
							fmt.Printf("\r  \033[36mChecking PVC\033[m %s  \r %s", s.Next(), tableOutput)
						}
					}
				}
			}()
			// defer func() {
			// 	finishedCh <- true
			// }()

			if err := plugin.RunPlugin(KubernetesConfigFlags, output); err != nil {
				return errors.Cause(err)
			}

			log.Info("")

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	return cmd
}

func InitAndExecute() {
	if err := RootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.AutomaticEnv()
}
