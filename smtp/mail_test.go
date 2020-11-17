package smtp

import (
	"reflect"
	"testing"
)

func TestAddress_String(t *testing.T) {
	type fields struct {
		Name    string
		Address string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "박상일 <psi59@lalaworks.com>",
			fields: fields{
				Name:    "박상일",
				Address: "psi59@lalaworks.com",
			},
			want: "=?UTF-8?b?67CV7IOB7J28?= <psi59@lalaworks.com>",
		},
		{
			name: "Sangil <psi59@lalaworks.com>",
			fields: fields{
				Name:    "Sangil",
				Address: "psi59@lalaworks.com",
			},
			want: "Sangil <psi59@lalaworks.com>",
		},
		{
			name: "psi59@lalaworks.com",
			fields: fields{
				Name:    "",
				Address: "psi59@lalaworks.com",
			},
			want: "psi59@lalaworks.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Address{
				Name:    tt.fields.Name,
				Address: tt.fields.Address,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAddress(t *testing.T) {
	type args struct {
		addresses []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Address
		wantErr bool
	}{
		{
			name: "박상일 <psi59@lalaworks.com>",
			args: args{
				addresses: []string{"=?UTF-8?b?67CV7IOB7J28?= <psi59@lalaworks.com>"},
			},
			want: []Address{
				{
					Name:    "박상일",
					Address: "psi59@lalaworks.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Sangil <psi59@lalaworks.com>",
			args: args{addresses: []string{"Sangil <psi59@lalaworks.com>"}},
			want: []Address{
				{
					Name:    "Sangil",
					Address: "psi59@lalaworks.com",
				},
			},
			wantErr: false,
		},
		{
			name: "multiple address",
			args: args{
				addresses: []string{
					"=?UTF-8?b?67CV7IOB7J28?= <psi59@lalaworks.com>",
					"Sangil <psi59@lalaworks.com>",
				},
			},
			want: []Address{
				{
					Name:    "박상일",
					Address: "psi59@lalaworks.com",
				},
				{
					Name:    "Sangil",
					Address: "psi59@lalaworks.com",
				},
			},
			wantErr: false,
		},
		{
			name: "psi59@lalaworks.com",
			args: args{addresses: []string{"psi59@lalaworks.com"}},
			want: []Address{
				{
					Name:    "",
					Address: "psi59@lalaworks.com",
				},
			},
			wantErr: false,
		},
		{
			name:    "empty addresses",
			args:    args{addresses: []string{}},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "empty string",
			args:    args{addresses: []string{""}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid header",
			args: args{
				addresses: []string{
					"=?UTF-8?b?invalid7IOB7J28?= <psi59@lalaworks.com>",
				},
			},
			want: []Address{
				{
					Name:    "=?UTF-8?b?invalid7IOB7J28?=",
					Address: "psi59@lalaworks.com",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAddress(tt.args.addresses...)
			t.Logf("err=%v, address=%#v", err, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAddress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
