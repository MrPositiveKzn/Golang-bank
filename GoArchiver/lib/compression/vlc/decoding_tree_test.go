package vlc

import (
	"GoArchiver/lib/compression/vlc/table"
	"reflect"
	"testing"
)

func Test_encodingTable_DecodingTree(t *testing.T) {
	tests := []struct {
		name string
		et   encodingTable
		want table.DecodingTree
	}{
		{
			name: "base tree test",
			et: encodingTable{
				'a': "11",
				'b': "1001",
				'z': "0101",
			},
			want: table.DecodingTree{
				Zero: &table.DecodingTree{
					One: &table.DecodingTree{
						Zero: &table.DecodingTree{
							One: &table.DecodingTree{
								Value: "z",
							},
						},
					},
				},
				One: &table.DecodingTree{
					Zero: &table.DecodingTree{
						Zero: &table.DecodingTree{
							One: &table.DecodingTree{
								Value: "b",
							},
						},
					},
					One: &table.DecodingTree{
						Value: "a",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.et.DecodingTree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodingTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
