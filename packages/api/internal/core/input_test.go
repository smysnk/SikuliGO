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
			name: "mouse down valid",
			req: InputRequest{
				Action: InputActionMouseDown,
				X:      1,
				Y:      2,
				Button: "left",
			},
		},
		{
			name: "mouse up valid",
			req: InputRequest{
				Action: InputActionMouseUp,
				X:      1,
				Y:      2,
				Button: "left",
			},
		},
		{
			name: "paste valid",
			req: InputRequest{
				Action: InputActionPasteText,
				Text:   "hello",
			},
		},
		{
			name: "key down valid",
			req: InputRequest{
				Action: InputActionKeyDown,
				Keys:   []string{"CMD"},
			},
		},
		{
			name: "key up valid",
			req: InputRequest{
				Action: InputActionKeyUp,
				Keys:   []string{"CMD"},
			},
		},
		{
			name: "wheel valid",
			req: InputRequest{
				Action:          InputActionWheel,
				X:               1,
				Y:               2,
				ScrollDirection: "down",
				ScrollSteps:     2,
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
			name: "invalid wheel missing direction",
			req: InputRequest{
				Action:      InputActionWheel,
				ScrollSteps: 1,
			},
			wantErr: true,
		},
		{
			name: "invalid wheel missing steps",
			req: InputRequest{
				Action:          InputActionWheel,
				ScrollDirection: "down",
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
