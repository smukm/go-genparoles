package str

import "testing"

const (
	success = "\u2713"
	failed  = "\u2717"
)

func Test_GetRandomChars(t *testing.T) {

	t.Logf("Тесты с правильными параметрами")
	{
		testId := 1
		num := 3
		t.Logf("\tTest %d:\tNumber=%d", testId, num)
		{
			ch := GetRandomChars(num)
			if len(ch) != num {
				t.Fatalf("\t%s\tНеправильный результат : %d %T.", failed, len(ch), ch)
			}
			for i,k := range ch {
				if k == 0 {
					t.Fatalf("\t%s\tch[%d]=0.", failed, i)
				}
			}
			t.Logf("\t%s\tТест пройден", success)
		}

		testId++
		
		t.Logf("\tTest %d:\tNumber=%d", testId, num)
		{
			
		}
	}
}
