//go:build !integration
// +build !integration

package internal

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    *App
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				filename: "../testdata/test.csv",
			},
			want:    &App{data: []float64{2, 4, 6, 8, 10}},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				filename: "../testdata/test-ne.csv",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_loadData(t *testing.T) {
	t.Parallel()

	type fields struct {
		data []float64
	}
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				data: []float64{2, 4, 6, 8, 10},
			},
			args: args{
				filename: "../testdata/test.csv",
			},
			wantErr: false,
		},
		{
			name: "error",
			fields: fields{
				data: nil,
			},
			args: args{
				filename: "../testdata/test-ne.csv",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				data: tt.fields.data,
			}
			if err := app.loadData(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("loadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApp_CalculateStandardDeviationV1(t *testing.T) {
	t.Parallel()

	type fields struct {
		data []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "default",
			fields: fields{
				data: []float64{2, 4, 6, 8, 10},
			},
			want: 3.1622776601683795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				data: tt.fields.data,
			}
			if got := app.CalculateStandardDeviationV1(); got != tt.want {
				t.Errorf("CalculateStandardDeviationV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_CalculateStandardDeviationV2(t *testing.T) {
	t.Parallel()

	type fields struct {
		data []float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "default",
			fields: fields{
				data: []float64{2, 4, 6, 8, 10},
			},
			want: 3.1622776601683795,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &App{
				data: tt.fields.data,
			}
			if got := app.CalculateStandardDeviationV2(); got != tt.want {
				t.Errorf("CalculateStandardDeviationV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
