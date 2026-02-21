//go:build gogosseract

package sikuli

import "testing"

func TestCrossProtocolIntegrationFlowTaggedOCR(t *testing.T) {
	runCrossProtocolIntegrationFlow(t, "DemoAppTagged")
}
