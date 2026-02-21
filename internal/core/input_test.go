package core

import "testing"

func TestInputRequestValidate(t *testing.T) {
	cases := []struct {
		name    string
		req     InputRequest
		wantErr bool
	}{
		{
			name: "mouse move valid",
			req: InputRequest{
				Action: InputActionMouseMove,
				X:      1,
				Y:      2,
			},
		},
		{
			name: "click valid",
			req: InputRequest{
				Action: InputActionClick,
				X:      1,
				Y:      2,
				Button: "left",
			},
		},
		{
			name: "type valid",
			req: InputRequest{
				Action: InputActionTypeText,
				Text:   "hello",
			},
		},
		{
			name: "hotkey valid",
			req: InputRequest{
				Action: InputActionHotkey,
				Keys:   []string{"CMD", "P"},
			},
		},
		{
			name: "invalid empty action",
			req: InputRequest{
				Action: "",
			},
			wantErr: true,
		},
		{
			name: "invalid negative delay",
			req: InputRequest{
				Action: InputActionMouseMove,
				Delay:  -1,
			},
			wantErr: true,
		},
		{
			name: "invalid click missing button",
			req: InputRequest{
				Action: InputActionClick,
				X:      1,
				Y:      2,
			},
			wantErr: true,
		},
		{
			name: "invalid type empty",
			req: InputRequest{
				Action: InputActionTypeText,
				Text:   " ",
			},
			wantErr: true,
		},
		{
			name: "invalid hotkey empty",
			req: InputRequest{
				Action: InputActionHotkey,
			},
			wantErr: true,
		},
		{
			name: "invalid unknown action",
			req: InputRequest{
				Action: "unknown",
			},
			wantErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.req.Validate()
			if tc.wantErr && err == nil {
				t.Fatalf("expected error")
			}
			if !tc.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
