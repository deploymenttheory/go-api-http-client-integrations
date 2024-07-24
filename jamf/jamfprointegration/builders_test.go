package jamfprointegration

// func TestBuildWithOAuth(t *testing.T) {
// 	type args struct {
// 		jamfBaseDomain    string
// 		Sugar             *zap.SugaredLogger
// 		bufferPeriod      time.Duration
// 		clientId          string
// 		clientSecret      string
// 		hideSensitiveData bool
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Integration
// 		wantErr bool
// 	}{
// 		{
// 			name: "1",
// 			args: args{
// 				jamfBaseDomain:    "https://yourserver.jamfcloud.com",
// 				Sugar:             test_newSugaredLogger(),
// 				bufferPeriod:      10 * time.Minute,
// 				clientId:          "client_id",
// 				clientSecret:      "client_secret",
// 				hideSensitiveData: true,
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := BuildWithOAuth(tt.args.jamfBaseDomain, tt.args.Sugar, tt.args.bufferPeriod, tt.args.clientId, tt.args.clientSecret, tt.args.hideSensitiveData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("BuildWithOAuth() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("BuildWithOAuth() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestBuildWithBasicAuth(t *testing.T) {
// 	type args struct {
// 		jamfBaseDomain    string
// 		Sugar             *zap.SugaredLogger
// 		bufferPeriod      time.Duration
// 		username          string
// 		password          string
// 		hideSensitiveData bool
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Integration
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := BuildWithBasicAuth(tt.args.jamfBaseDomain, tt.args.Sugar, tt.args.bufferPeriod, tt.args.username, tt.args.password, tt.args.hideSensitiveData)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("BuildWithBasicAuth() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("BuildWithBasicAuth() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestIntegration_BuildOAuth(t *testing.T) {
// 	type fields struct {
// 		BaseDomain           string
// 		AuthMethodDescriptor string
// 		Sugar                *zap.SugaredLogger
// 		auth                 authInterface
// 	}
// 	type args struct {
// 		clientId          string
// 		clientSecret      string
// 		bufferPeriod      time.Duration
// 		hideSensitiveData bool
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			j := &Integration{
// 				BaseDomain:           tt.fields.BaseDomain,
// 				AuthMethodDescriptor: tt.fields.AuthMethodDescriptor,
// 				Sugar:                tt.fields.Sugar,
// 				auth:                 tt.fields.auth,
// 			}
// 			j.BuildOAuth(tt.args.clientId, tt.args.clientSecret, tt.args.bufferPeriod, tt.args.hideSensitiveData)
// 		})
// 	}
// }

// func TestIntegration_BuildBasicAuth(t *testing.T) {
// 	type fields struct {
// 		BaseDomain           string
// 		AuthMethodDescriptor string
// 		Sugar                *zap.SugaredLogger
// 		auth                 authInterface
// 	}
// 	type args struct {
// 		username          string
// 		password          string
// 		bufferPeriod      time.Duration
// 		hideSensitiveData bool
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			j := &Integration{
// 				BaseDomain:           tt.fields.BaseDomain,
// 				AuthMethodDescriptor: tt.fields.AuthMethodDescriptor,
// 				Sugar:                tt.fields.Sugar,
// 				auth:                 tt.fields.auth,
// 			}
// 			j.BuildBasicAuth(tt.args.username, tt.args.password, tt.args.bufferPeriod, tt.args.hideSensitiveData)
// 		})
// 	}
// }
