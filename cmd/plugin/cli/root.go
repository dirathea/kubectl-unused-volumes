package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/ava-labs/kubectl-unused-volumes/pkg/plugin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

var (
	KubernetesConfigFlags *genericclioptions.ConfigFlags
	Opts                  plugin.Options
	commandExample        = `
	# Get all unused volumes in default namespace
	kubectl unused-volumes

	# Get all unused volumes in all namespaces
	kubectl unused-volumes --all-namespaces

	# Remove headers
	kubectl unused-volumes --no-headers
	`
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "unused-volumes",
		Short:         "",
		Long:          `.`,
		SilenceErrors: true,
		Example:       commandExample,
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
	cmd.Flags().BoolVar(&Opts.NoHeaders, "no-headers", false, "Skip header")
	cmd.Flags().BoolVar(&Opts.AllNamespaces, "all-namespaces", false, "Skip header")

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
