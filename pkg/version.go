package opc

import (
	"fmt"

	_ "embed"

	tkncli "github.com/tektoncd/cli/pkg/cli"

	// paccli "github.com/openshift-pipelines/pipelines-as-code/pkg/cli"
	tknversion "github.com/tektoncd/cli/pkg/cmd/version"
	"github.com/tektoncd/cli/pkg/version"

	// paccli "github.com/openshift-pipelines/opc/pkg"
	"github.com/spf13/cobra"
)

// //go:embed version.json
// var versionFile string

// //go:embed version.tmpl
// var versionTmpl string
var (
	component     = ""
	componentFalg string
	namespace     string
	err           error
	clientVersion = "devel"
)

// type versions struct {
// 	Opc string `json:"opc"`
// 	Tkn string `json:"tkn"`
// 	Pac string `json:"pac"`
// }

func VersionCommand(tp tkncli.Params) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print opc version",
		Long:  "Print OpenShift Pipeline Client version",
		RunE: func(cmd *cobra.Command, args []string) error {
			cs, err := tp.Clients()
			if err == nil {
				switch component {
				case "":
					fmt.Fprintf(cmd.OutOrStdout(), "Tkn Verion: %s\n", tknversion.ClientVersion)
					fmt.Fprintf(cmd.OutOrStdout(), "OpenShift Pipelines Client Verion: %s\n", clientVersion)
					chainsVersion, _ := version.GetChainsVersion(cs, namespace)
					if chainsVersion != "" {
						fmt.Fprintf(cmd.OutOrStdout(), "Chains version: %s\n", chainsVersion)
					}

					pipelineVersion, _ := version.GetPipelineVersion(cs, namespace)
					if pipelineVersion == "" {
						pipelineVersion = "unknown, " +
							"pipeline controller may be installed in another namespace please use tkn version -n {namespace}"
					}
					fmt.Fprintf(cmd.OutOrStdout(), "Pipeline version: %s\n", pipelineVersion)

					triggersVersion, _ := version.GetTriggerVersion(cs, namespace)
					if triggersVersion != "" {
						fmt.Fprintf(cmd.OutOrStdout(), "Triggers version: %s\n", triggersVersion)
					}

					dashboardVersion, _ := version.GetDashboardVersion(cs, namespace)
					if dashboardVersion != "" {
						fmt.Fprintf(cmd.OutOrStdout(), "Dashboard version: %s\n", dashboardVersion)
					}

					operatorVersion, _ := version.GetOperatorVersion(cs, namespace)
					if operatorVersion != "" {
						fmt.Fprintf(cmd.OutOrStdout(), "Operator version: %s\n", operatorVersion)
					}
				case "client":
					fmt.Fprintf(cmd.OutOrStdout(), "OpenShift Pipelines Client Verion: %s\n", clientVersion)

				case "chains":
					chainsVersion, _ := version.GetChainsVersion(cs, namespace)
					if chainsVersion == "" {
						chainsVersion = "unknown"
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", chainsVersion)

				case "pipeline":
					pipelineVersion, _ := version.GetPipelineVersion(cs, namespace)
					if pipelineVersion == "" {
						pipelineVersion = "unknown"
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", pipelineVersion)

				case "triggers":
					triggersVersion, _ := version.GetTriggerVersion(cs, namespace)
					if triggersVersion == "" {
						triggersVersion = "unknown"
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", triggersVersion)

				case "dashboard":
					dashboardVersion, _ := version.GetDashboardVersion(cs, namespace)
					if dashboardVersion == "" {
						dashboardVersion = "unknown"
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", dashboardVersion)

				case "operator":
					operatorVersion, _ := version.GetOperatorVersion(cs, namespace)
					if operatorVersion == "" {
						operatorVersion = "unknown"
					}
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", operatorVersion)

				default:
					fmt.Fprintf(cmd.OutOrStdout(), "Invalid component value\n")

				}

			} else {
				switch component {
				case "":
					fmt.Fprintf(cmd.OutOrStdout(), "Client version: %s\n", clientVersion)
				case "client":
					fmt.Fprintf(cmd.OutOrStdout(), "%s\n", clientVersion)
				case "chains", "pipeline", "triggers", "dashboard", "operator":
					fmt.Fprintf(cmd.OutOrStdout(), "unknown\n")
				default:
					fmt.Fprintf(cmd.OutOrStdout(), "Invalid component value\n")
				}
			}
			return nil
		},
		Annotations: map[string]string{
			"commandType": "main",
		},
	}
	cmd.Flags().StringVarP(&namespace, "namespace", "n", namespace, "namespace to check installed controller version")
	cmd.Flags().StringVarP(&component, "component", "c", "", "provide a particular component name for its version (client|tkn|chains|pipeline|triggers|dashboard)")
	return cmd
}
