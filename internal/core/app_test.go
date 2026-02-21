package core

import "testing"

func TestAppRequestValidate(t *testing.T) {
	base := AppRequest{
		Action: AppActionOpen,
		Name:   "demo-app",
	}
	if err := base.Validate(); err != nil {
		t.Fatalf("base request should validate: %v", err)
	}

	cases := []AppRequest{
		{Action: "", Name: "demo-app"},
		{Action: AppActionOpen, Name: ""},
		{Action: "bad-action", Name: "demo-app"},
		{Action: AppActionOpen, Name: "demo-app", Timeout: -1},
	}

	for i := range cases {
		if err := cases[i].Validate(); err == nil {
			t.Fatalf("case %d should fail validation", i)
		}
	}
}
