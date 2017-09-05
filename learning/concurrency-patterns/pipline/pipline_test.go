package pipline

import "testing"

func TestLaunchPipline(t *testing.T) {
	cases := []struct {
		in   int
		want int
	}{
		{3, 14},
		{5, 55},
	}

	var res int
	for _, c := range cases {
		res = LaunchPipeline(c.in)
		if res != c.want {
			t.Fatal()
		}
		t.Logf("%d == %d\n", res, c.want)
	}
}
