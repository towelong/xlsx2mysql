package app

import "testing"

func TestGenerateTableAndSQL(t *testing.T) {
	type args struct {
		rows [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantS   string
		wantErr bool
	}{
		{
			name: "测试",
			args: args{
				rows: [][]string{
					{
						"id",
						"name",
					},
					{
						"int",
						"varchar",
					},
					{
						"编号",
						"姓名",
					},
					{
						"1",
						"张三",
					},
					{
						"2",
						"张四",
					},
					{
						"3",
						"张五",
					},
				},
			},
			wantS:   "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS, err := GenerateTableAndSQL("test", tt.args.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateTableAndSQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotS != tt.wantS {
				t.Errorf("GenerateTableAndSQL() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
