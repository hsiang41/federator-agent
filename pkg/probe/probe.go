package probe

import (
	"os"
	"github.com/containers-ai/alameda/pkg/utils/log"
)

var logger = log.RegisterScope("Federatorai-Agent", "Federatorai-agent", 0)

func LivenessProbe(cfg *LivenessProbeConfig) {
	os.Exit(0)
}

func ReadinessProbe() {
	os.Exit(0)
}
