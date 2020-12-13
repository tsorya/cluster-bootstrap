package main

import (
	"errors"
	"github.com/openshift/cluster-bootstrap/pkg/ibip"
	"github.com/spf13/cobra"

)

var (
	cmdIBip = &cobra.Command{
		Use:          "ibip",
		Short:        "Start the control plane",
		Long:         "",
		PreRunE:      validateIBipOpts,
		RunE:         runCmdIBip,
		SilenceUsage: true,
	}

	iBipOpts struct {
		ignitionPath		 string
	}
)

func init() {
	cmdRoot.AddCommand(cmdIBip)
	cmdStart.Flags().StringVar(&iBipOpts.ignitionPath, "ignition-path", "/assets/master.ign", "The location of master ignition")

}

func runCmdIBip(cmd *cobra.Command, args []string) error {

	ib, err := ibip.NewIBipCommand(ibip.ConfigIBip{
		AssetDir:             startOpts.assetDir,
		IgnitionPath:      iBipOpts.ignitionPath,
	})
	if err != nil {
		return err
	}

	return ib.UpdateSnoIgnitionData()
}

func validateIBipOpts(cmd *cobra.Command, args []string) error {
	if iBipOpts.ignitionPath == "" {
		return errors.New("missing required flag: --ignition-path")
	}
	if startOpts.assetDir == "" {
		return errors.New("missing required flag: --asset-dir")
	}
	return nil
}
