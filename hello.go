// package main

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"io"
// 	"math/rand"
// 	"time"
// )

// type generator struct {
//     rnd rand.Source // Генератор случайных чисел. Вообще rand.Rand уже реализует интерфейс io.Reader, но для примера мы реализуем его самостоятельно.
// }

// // New — обратите внимание, что мы возвращаем generator, присвоенный интерфейсу io.Reader, сама структура generator неэкспортируемая.
// // Мы скрыли внутри пакета все детали.
// func New(seed int64) io.Reader {
//     return &generator{
//         rnd: rand.NewSource(seed),
//     }
// }

// // Read — реализация io.Reader
// func (g *generator) Read(bytes []byte) (n int, err error) {
//     for i := 0; i+8 < len(bytes); i += 8 {
//         binary.LittleEndian.PutUint64(bytes[i:i+8], uint64(g.rnd.Int63()))
//     }
//     return len(bytes), nil
// }

// func main() {

//     // создаём генератор случайных чисел
//     generator := New(time.Now().UnixNano()) // в качестве затравки передаём ему текущее время, и при каждом запуске оно будет разным.

//     buf := make([]byte, 16)

//     for i := 0; i < 5; i++ {
//         n, _ := generator.Read(buf) // единственный доступный метод, но он нам и нужен.
//         fmt.Printf("Generate bytes: %v size(%d)\n", buf, n)
//     }

// }
// // import (
// //     "fmt"

// //     "toppackage/middlepackage/bottompackage/mathxxx"
// // )

// // func main() {
// //     if sum := mathxxx.AddInts(3, 2); sum != 5 {
// //         panic(fmt.Sprintf("sum must be equal 5; got %d", sum))
// //     }

// //     fmt.Println("Well done!")
// // } 

// // package main

// // import "fmt"

// // var Global = 5

// // func switchGLobal() {
// // 	defer func() {
// // 		Global = 6
// // 		fmt.Println(Global)
// // 		Global = 5
// // 	}()
// // }

// // func main() {
// // 	switchGLobal()
// // 	fmt.Println(Global)
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"math"
// // )

// // type figures int

// // const (
// // 	square figures = iota
// // 	circle
// // 	triangle
// // )

// // func area(fig figures) (func(float64) float64, bool) {
// // 	switch fig {
// // 	case square:
// // 		return func(f float64) float64 {
// // 			return f * f
// // 		}, true
// // 	case circle:
// // 		return func(f float64) float64 {
// // 			return math.Pi * f * f
// // 		}, true
// // 	case triangle:
// // 		return func(f float64) float64 {
// // 			return ((f * f) * math.Sqrt(3)) / 4
// // 		}, true
// // 	default:
// // 		return nil, false
// // 	}
// // }

// // func main() {
// // 	ar, ok := area(triangle)
// // 	if !ok {
// // 		fmt.Println("Ошибка")
// // 		return
// // 	}
// // 	var x float64 = 5
// // 	myArea := ar(x)
// // 	fmt.Println(myArea)
// // }

// // package main

// // import (
// //     "fmt"
// //     "os"
// //     "path/filepath"
// // 	"strings"
// // )

// // func main() {
// //     PrintAllFiles(".", ".vscode")
// // }

// // func PrintAllFiles(path string, filter string) {
// //     // получаем список всех элементов в папке (и файлов, и директорий)
// //     files, err := os.ReadDir(path)
// //     if err != nil {
// //         fmt.Println("unable to get list of files", err)
// //         return
// //     }
// //     //  проходим по списку
// //     for _, f := range files {
// // 		filename := filepath.Join(path, f.Name())
// // 		if strings.Contains(filename, filter){
// //         // получаем имя элемента
// //         // filepath.Join — функция, которая собирает путь к элементу с разделителям
// //         // печатаем имя элемента
// //         fmt.Println(filename)
// //         // если элемент — директория, то вызываем для него рекурсивно ту же функцию
// //         if f.IsDir() {
// //             PrintAllFiles(filename, filter)
// //         }
// // 		}
// //     }
// // }

// // package main

// // import (
// // 	"fmt"
// // 	"os"
// // 	"path/filepath"
// // 	"strings"
// // )

// // func PrintAllFiles(path string, filter string) {
// // 	var walk func(string)
// // 	walk = func(path string) {
// // 		// получаем список всех элементов в папке (и файлов, и директорий)
// // 		files, err := os.ReadDir(path)
// // 		if err != nil {
// // 			fmt.Println("unable to get list of files", err)
// // 			return
// // 		}
// // 		//  проходим по списку
// // 		for _, f := range files {

// // 			// получаем имя элемента
// // 			// filepath.Join — функция, которая собирает путь к элементу с разделителям
// // 			filename := filepath.Join(path, f.Name())
// // 			// печатаем имя элемента
// // 			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
// // 			if strings.Contains(filename, filter) {
// // 				fmt.Println(filename)
// // 			}
// // 			if f.IsDir() {
// // 				walk(filename)
// // 			}
// // 		}
// // 	}
// // 	walk(path)
// // }

// // func PrintAllFilesWithFuncFilter(path string, predicate func(string) bool) {
// // 	var walk func(string)
// // 	walk = func(path string) {
// // 		files, err := os.ReadDir(path)
// // 		if err != nil {
// // 			fmt.Println("Ошибка", err)
// // 			return
// // 		}
// // 		for _, f := range files {
// // 			filename := filepath.Join(path, f.Name())
// // 			if predicate(filename) {
// // 				fmt.Println(filename)
// // 			}
// // 			if f.IsDir() {
// // 				walk(filename)
// // 			}
// // 		}
// // 	}
// // 	walk(path)
// // }

// // func containsA(s string) bool {
// // 	return strings.Contains(s, "h")
// // }

// // func main() {
// // 	PrintAllFilesWithFuncFilter(".", containsA)
// // }

// // package main

// // import (
// //     "encoding/json"
// //     "fmt"
// // )

// // const rawResp = `
// // {
// //     "header": {
// //         "code": 0,
// //         "message": ""
// //     },
// //     "data": [{
// //         "type": "user",
// //         "id": 100,
// //         "attributes": {
// //             "email": "bob@yandex.ru",
// //             "article_ids": [10, 11, 12]
// //         }
// //     }]
// // }
// // `

// // type (
// //     Response struct {
// //         Header ResponseHeader `json:"header"`
// //         Data   ResponseData   `json:"data,omitempty"`
// //     }

// //     ResponseHeader struct {
// //         Code    int    `json:"code"`
// //         Message string `json:"message,omitempty"`
// //     }

// //     ResponseData []ResponseDataItem

// //     ResponseDataItem struct {
// //         Type       string                `json:"type"`
// //         Id         int                   `json:"id"`
// //         Attributes ResponseDataItemAttrs `json:"attributes"`
// //     }

// //     ResponseDataItemAttrs struct {
// //         Email      string `json:"email"`
// //         ArticleIds []int  `json:"article_ids"`
// //     }
// // )

// // func ReadResponse(rawResp string) (Response, error) {
// //     resp := Response{}
// //     if err := json.Unmarshal([]byte(rawResp), &resp); err != nil {
// //         return Response{}, fmt.Errorf("JSON unmarshal: %w", err)
// //     }

// //     return resp, nil
// // }

// // func main() {
// //     resp, err := ReadResponse(rawResp)
// //     if err != nil {
// //         panic(err)
// //     }
// //     fmt.Printf("%+v\n", resp)
// // }

// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"time"
// // )

// // func main() {
// // 	type Person struct {
// // 		Name        string    `json:"Имя"`
// // 		Email       string    `json:"Почта"`
// // 		DateOfBirth time.Time `json:"-"`
// // 	}

// // 	person := Person{
// // 		Name:        "Alex",
// // 		Email:       "ale@gmail.com",
// // 		DateOfBirth: time.Date(1995, time.December, 10, 0, 0, 0, 0, time.UTC),
// // 	}

// // 	jsonBytes, err := json.MarshalIndent(person, "", "  ")
// // 	if err != nil {
// // 		fmt.Println("Ошибка при преобразовании в JSON:", err)
// // 		return
// // 	}

// // 	fmt.Println(string(jsonBytes))

// // }
// // func RemoveDuplicates(input []string) []string {
// // 	output :=make([]string, 0)
// // 	seen := make(map[string]struct{}, len(input))

// // 	for _, v := range input {
// // 		if _, ok := seen[v]; !ok {
// // 			output = append(output, v)
// // 		}
// // 		seen[v] = struct{}{}
// // 	}
// // 	return output
// // }
// // func RemoveDuplicates(input []string) []string {
// // 	m := make(map[string]int)
// // 	for i, v := range input {
// // 		if _, ok := m[v]; ok {
// // 			if len(input) != 0 && i < len(input) {
// // 				input = append(input[:i], input[i+1:]...)
// // 			}
// // 		}
// // 		m[v] = i
// // 	}
// // 	return input
// // }

// // var input = []string{
// // 	"cat",
// // 	"dog",
// // 	"bird",
// // 	"dog",
// // 	"parrot",
// // 	"cat",
// // }

// // func main() {
// // 	fmt.Println(RemoveDuplicates(input))
// // }

// // func find(arr []int, k int) []int {
// // 	m := make(map[int]int)
// // 	for i, v := range arr {
// // 		if x, ok := m[k-v]; ok {
// // 			fmt.Println([]int{i, x})
// // 		}
// // 		m[v] = i
// // 	}
// // 	return nil
// // }

// // func main() {
// // 	arr := []int{10, 3, 12, 5}
// // 	k := 15
// // 	find(arr, k)
// // }

// // func main(){
// // 	order :=[]string{"хлеб", "буженина", "сыр", "огурцы"}
// // 	priceList := map[string]int{"хлеб": 50, "молоко": 100, "масло": 200,"колбаса": 500,"соль": 20,"огурцы": 200,"сыр": 600,"ветчина": 700,"буженина": 900,"помидоры": 250, "рыба": 300, "хамон":1500}
// // 	for k, v := range priceList{
// // 		if v>=500 {
// // 			fmt.Println(k)
// // 		}
// // 	}
// // 	total := 0
// // 	for _, v := range order{
// // 		total +=priceList[v]
// // 	}
// // 	fmt.Println(total)
// // }

// // 	func main() {
// // 		s := make([]int, 0, 100)
// // 		for i := 1; i <= 100; i++ {
// // 			s = append(s, i)
// // 		}
// // 	    s = append(s[:10], s[len(s)-10:]...)
// // 	    for i := range s[:len(s)/2]{
// // 	        s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
// // 	    }
// // 	    fmt.Println(s)
// // 	}
// // func main() {
// // 	input := "The quick brown 狐 jumped over the lazy 犬"

// // 	n := 0
// // 	runes := make([]rune, len(input))
// // 	for _, r := range input {
// // 		runes[n] = r
// // 		n++
// // 	}

// // 	runes = runes[0:n]

// // 	for i := 0; i < n/2; i++ {
// // 		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
// // 	}

// // 	output := string(runes)
// // 	fmt.Println(output)

// // }

// // func genCounter(bdate int) string{
// // 	switch {
// //     case bdate >= 1946 && bdate <= 1964:
// //         return "Привет, бумер!"
// //     case bdate >= 1965 && bdate <= 1980:
// //         return "Привет, поколение X!"
// //     case bdate >= 1981 && bdate <= 1996:
// //         return "Привет, миллениал!"
// //     case bdate >= 1997 && bdate <= 2012:
// //         return "Привет, зумер!"
// //     case bdate >= 2013 && bdate <= time.Now().Year():
// //         return "Привет, альфа!"
// //     default:
// //         return "Дата вне диапазона"
// //     }
// // }

// // func main() {
// //     var bdate int
// //     fmt.Print("Введите дату рождения: ")
// //     fmt.Scan(&bdate)
// //     fmt.Println(genCounter(bdate))
// // }
// // 1946 — 1964: «Привет, бумер!».
// // 1965 — 1980: «Привет, представитель X!».
// // 1981 — 1996: «Привет, миллениал!».
// // 1997 — 2012: «Привет, зумер!».
// // 2013 — ... : «Привет, альфа!».
