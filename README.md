# Какой самый эффективный способ конкатенации строк?
### Из пакета strings - strings.Builder
# Что такое интерфейсы, как они применяются в Go?
### Интерфейс - это некое определение. Благодаря интерфейсам можно определить и описать методы, которые должны быть у какого-то другого типа
# Чем отличаются RWMutex от Mutex?
### Mutex - Один поток пишет и читает. RWMutex - Много потоков читает, но только один пишет
# Чем отличаются буферизированные и не буферизированные каналы?
### Небуферизированные каналы - делается пауза, пока не прочитают данные из канала. Буферизированные - делается пауза, когда канал заполнен
# Какой размер у структуры struct{}{}?
### 0
# Есть ли в Go перегрузка методов или операторов?
### Нет
# В какой последовательности будут выведены элементы map[int]int?
### В разном порядке в зависимости от способа вывода данных. Например, если использовать fmt.Println, то данные отсортируются.
# В чем разница make и new?
### make - возвращает инициализированный тип. new - возвращает указатель на нулевое значение типа
# Сколько существует способов задать переменную типа slice или map?
### C помощью переменной: m := map[int]int. С помощью make: m := make(map[int]int). 
# Что выведет данная программа и почему?
```go
func update(p *int) {
b := 2
p = &b
}
func main() {
var (
a = 1
p = &a
)
fmt.Println(*p)
update(p)
fmt.Println(*p)
}
```
### 1 1, потому что изменения касается только локальной копии указателя p внутри функции update, на main она никак не влияет.
# Что выведет данная программа и почему?
```go
func main() {
wg := sync.WaitGroup{}
for i := 0; i < 5; i++ {
wg.Add(1)
go func(wg sync.WaitGroup, i int) {
fmt.Println(i)
wg.Done()
}(wg, i)
}
wg.Wait()
fmt.Println("exit")
}
```
### Deadlock, потому что нет указателя на wg
# Что выведет данная программа и почему?
```go
func main() {
n := 0
if true {
n := 1
n++
}
fmt.Println(n)
}
```
### 0, потому что переменная n внутри if - это другая переменная
# Что выведет данная программа и почему?
```go
func someAction(v []int8, b int8) {
v[0] = 100
v = append(v, b)
}
func main() {
var a = []int8{1, 2, 3, 4, 5}
someAction(a, 6)
fmt.Println(a)
}
```
### 100 2 3 4 5, потому что места в массиве нет, соответственно в функции someAction при добавлении нового элемента в массив, создается новый массив v, который не будет влиять на исходный массив a
# Что выведет данная программа и почему?
```go
func main() {
slice := []string{"a", "a"}
func(slice []string) {
slice = append(slice, "a")
slice[0] = "b"
slice[1] = "b"
fmt.Print(slice)
}(slice)
fmt.Print(slice)
}
```
### b b a, a a, потому что при выолнении append создается новый слайс в анонимной функции, а в main функции слайс измнений не коснется