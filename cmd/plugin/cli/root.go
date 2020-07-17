package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/dirathea/kubectl-unused-volumes/pkg/plugin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
	Opts                  plugin.Options
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "volume-reclaim",
		Short:         "",
		Long:          `.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			output, err := plugin.RunPlugin(Opts)
			if err != nil {
				return errors.Cause(err)
			}
			fmt.Println(output)

			return nil
		},
	}

	cobra.OnInitialize(initConfig)

	KubernetesConfigFlags = genericclioptions.NewConfigFlags(false)
	KubernetesConfigFlags.AddFlags(cmd.Flags())
	Opts = plugin.Options{
		KubernetesConfigFlags: KubernetesConfigFlags,
	}
	cmd.Flags().BoolVar(&Opts.NoHeader, "no-header", false, "Skip header")

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
