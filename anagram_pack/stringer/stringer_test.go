package stringer

import "testing"

func TestAreAnagrams(t *testing.T) {
	// набор тестов
	cases := []struct {
		// имя теста
		name string
		// значения на вход тестируемой функции
		inp_str [2]string
		// желаемый результат
		want_res bool
	}{
		// тестовые данные № 1
		{
			name:     "anagram",
			inp_str:  [2]string{"cav", "vac"},
			want_res: true,
		},
		// тестовые данные № 2
		{
			name:     "no anagram length",
			inp_str:  [2]string{"cav", "vacc"},
			want_res: false,
		},
		// тестовые данные № 2
		{
			name:     "no anagram",
			inp_str:  [2]string{"cav", "vak"},
			want_res: false,
		},
	}
	// перебор всех тестов
	for i, tc := range cases {
		tc := tc
		// запуск отдельного теста
		t.Run(tc.name, func(t *testing.T) {
			// тестируем функцию Sum
			got := AreAnagrams(tc.inp_str[0], tc.inp_str[1])
			// проверим полученное значение
			if got != tc.want_res {
				t.Fatalf("%d, %q AreAnagrams(%q, %q) = %t; want %t", i, tc.name, tc.inp_str[0], tc.inp_str[1], got, tc.want_res)
			}
		})
	}
}
