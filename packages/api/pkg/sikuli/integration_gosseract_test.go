//go:build gosseract

package sikuli

import "testing"

func TestCrossProtocolIntegrationFlowTaggedOCR(t *testing.T) {
	runCrossProtocolIntegrationFlow(t, "DemoAppTagged")
}
