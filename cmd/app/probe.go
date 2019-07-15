package app

import (
	"github.com/containers-ai/federatorai-agent/pkg/probe"
	"github.com/spf13/cobra"
	"os"
)

const (
	PROBE_TYPE_READINESS = "readiness"
	PROBE_TYPE_LIVENESS  = "liveness"
)

var (
	probeType string

	ProbeCmd = &cobra.Command{
		Use:   "probe",
		Short: "probe federatorai agent",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			startProbing()
		},
	}
)

func init() {
	parseProbeFlag()
}

func parseProbeFlag() {
	ProbeCmd.Flags().StringVar(&probeType, "type", PROBE_TYPE_READINESS, "The probe type for federatorai-agent.")
}

func startProbing() {
	_, err := ReadConfig(transmitterConfigurationFile)
	if err != nil {
		logger.Fatalf("Failed to read configuration due to %s\n", err)
		os.Exit(1)
	}

	switch probeType {
	case PROBE_TYPE_LIVENESS:
		logger.Debugf("Execute Liveness")
		probe.LivenessProbe(&probe.LivenessProbeConfig{})
	case PROBE_TYPE_READINESS:
		logger.Debugf("Execute Readiness")
		probe.ReadinessProbe()
	default:
		logger.Errorf("Probe type does not supports %s, please try %s or %s.", probeType, PROBE_TYPE_LIVENESS, PROBE_TYPE_READINESS)
		os.Exit(1)
	}
}
