package jamfprointegration

// func TestIntegration_prepRequest(t *testing.T) {
// 	type fields struct {
// 		BaseDomain           string
// 		AuthMethodDescriptor string
// 		Sugar                *zap.SugaredLogger
// 		auth                 authInterface
// 	}
// 	type args struct {
// 		req *http.Request
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "blank request",
// 			fields: fields{
// 				BaseDomain:           "https://test_domain.com",
// 				AuthMethodDescriptor: "test",
// 				Sugar:                test_newSugaredLogger(),
// 				auth:                 test_newMockauth(),
// 			},
// 			args: args{
// 				req: test_newTestRequest(),
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			j := &Integration{
// 				BaseDomain:           tt.fields.BaseDomain,
// 				AuthMethodDescriptor: tt.fields.AuthMethodDescriptor,
// 				Sugar:                tt.fields.Sugar,
// 				auth:                 tt.fields.auth,
// 			}
// 			if err := j.prepRequest(tt.args.req); (err != nil) != tt.wantErr {
// 				t.Errorf("Integration.prepRequest() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
