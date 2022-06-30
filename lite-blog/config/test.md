[toc]

# j2go-cheatsheet

java 转 go 快速对照表，建议搭配 [smart TOC](https://chrome.google.com/webstore/detail/smart-toc/lifgeihcfpkmmlfjbailfpfhbahhibba) 工具阅读

这个速查表是为了给 *高级 Java 工程师* 快速熟悉 Golang 语言特性准备的。

语言只是工具，都有各自的优势和不足，没必要抬高这个贬低那个，请不要做语言的奴隶！


> 说明：目录中带星号 * 的代表这部分内容是针对于 java 侧来说的，在 golang 领域可能没有此概念


# 三方依赖仓库地址
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>


https://mvnrepository.com/


</td><td>


https://pkg.go.dev/


</td></tr>
</tbody></table>

# 包管理工具
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>


Maven、Gradle


</td><td>


Go Modules (go 1.13 设为默认)


</td></tr>
</tbody></table>

# 基础数据类型
基础数据类型是CPU可以直接进行运算的类型
<table>
<thead>
<tr>
<th>java</th>
<th>go</th>
<th>占用内存</th>
</tr>
</thead>
<tbody>
<tr>
<td>byte</td>
<td>byte/uint8</td>
<td>1  byte</td>

</tr>

<tr>
<td>short</td>
<td>int16</td>
<td>2  byte</td>

</tr>

<tr>
<td>int</td>
<td>int/int32/int64</td>
<td>java 4 byte; go 4/8  byte(Platform dependent)</td>

</tr>

<tr>
<td>long</td>
<td>int64</td>
<td>8  byte</td>

</tr>

<tr>
<td>float</td>
<td>float32</td>
<td>4  byte</td>

</tr>

<tr>
<td>double</td>
<td>float64</td>
<td>8  byte</td>

</tr>

<tr>
<td>char</td>
<td>rune</td>
<td>java 2 byte; go 4 byte</td>

</tr>

<tr>
<td>boolean</td>
<td>bool</td>
<td>1 byte</td>

</tr>

</tbody></table>

# 声明方式及赋值


## 变量
type: 类型

identifier: 标识符、变量名

value: 值
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
	
type identifier;
type identifier = value;
	
```

</td><td>

```go
var identifier type
var identifier type = value
var identifier = value
identifier := value
```

</td></tr>
</tbody>
</table>

## 常量
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
private static int NUM = 0; //int
```

</td><td>

```go
const num = 0
```

</td></tr>
</tbody>
</table>

# 循环控制
## for
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
for(int i = 0; i < arr.length; i++){
    //...
}
```

</td><td>

```go
for i := 0; i < len(arr); i++ {
    //...
}
```

</td></tr>
</tbody>
</table>

## for 增强
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
String s = "abc";
for(int i : s.toCharArray()){
    //...
}
```

</td><td>

```go
s := "abc"

for i,v := range s {
    //...
}
```

</td></tr>
</tbody>
</table>

## while
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
int i = 0;
while(i < 10){
    //...
}
```

</td><td>

```go
i : = 0
for i < 10 {
    i++
}
```

</td></tr>
</tbody>
</table>

## switch
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
int i = 0;
switch(i){
case 1:
    System.out.println("1");
case 2:
    System.out.println("2");
case 3:
    System.out.println("3");
default:
    System.out.println("0");
}
```

</td><td>

```go
i := 0
switch i {
case 1:
    fmt.Println("1")
case 2:
    fmt.Println("2")
case 3:
    fmt.Println("3")
default:
    fmt.Println("0")
```

</td></tr>
</tbody>
</table>

# 集合
## List
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
List<T> list1 = new ArrayList<>(); 
List<T> list2 = new LinkedList<>();
```

</td><td>

```go
list1 := make([]T, len(xx));  // array list
var list2 list.List // linked list
list2.pushBack()
for e := list2.Front(); e != nil; e=e.Next() {
        fmt.Println(e.Value.(xx))
    }
```

</td></tr>
</tbody></table>

## Map
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
Map<T,T> map = new HashMap<>();
map.put(T,T);
T t = map.get(T);
```

</td><td>

```go
maps := make(map[T]T)
maps[T] = T
x, ok := maps[T]
if ok {
    //do somethings...
}

```

</td></tr>
</tbody></table>

## Set
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
Set<T> set = new HashSet<>(); 
```

</td><td>

```go
type void struct{}
var x void
set := make(map[T]void)
```

</td></tr>
</tbody></table>


# I/O
## 文件读写
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java

```

</td><td>

```go

```

</td></tr>
</tbody></table>

# String

## 分割字符串
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
String[] strs = "a,b,c".split(",");
```

</td><td>

```go
// []string
strs := strings.Split("a,b,c", ",")
```

</td></tr>
</tbody></table>



# *面向对象编程
面向对象是以对象为核心向外拓展的，可以理解为现实环境中的映射。在 java 中表示为 Class，在 golang 中表示为 struct。在设计时会将对象能力通过方法的形式整合在一起并通过权限控制来加强封装，使用接口 interface 来实现多态。

## 函数
函数是能给调用者返回一些需要的值，可以在任何地方使用
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
class Door{
    private String size;
    public void open(){
        System.out.println("door opened");
    }

    public String getSize() {
        return size;
    }

    public void setSize(String size) {
        this.size = size;
    }
}

```

</td><td>

```go
type Door struct {
  size string
}
func open(){
  fmt.Println("door opened")
  
}


```

</td></tr>
</tbody></table>

## 方法
方法是被定义在类内部，能去修改类的属性，注重类的概念
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
class Door{
    private String size;
    public void open(){
        System.out.println("door opened");
    }

    public String getSize() {
        return size;
    }

    public void setSize(String size) {
        this.size = size;
    }
}
```

</td><td>

```go

type Door struct {
  size string
}
func (door *Door) setSize(){
 door.size = "big"
}
func (door *Door) getSize(){
  fmt.Println("door size is:",door.size)
}


```

</td></tr>
</tbody></table>

# 接口
golang 中没有显示的实现
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
public class TestInterface {
    public static void main(String[] args) {
        System.out.println("light");
        CarLight carLight = new CarLight();
        carLight.turnOn();
        carLight.turnOff();

    }
}

interface Light {
    void turnOn();

    void turnOff();
}

class CarLight implements Light {

    @Override
    public void turnOn() {
        System.out.println("turn on");
    }

    @Override
    public void turnOff() {
        System.out.println("turn off");
    }
}
```


</td><td>

```golang
type light interface {
	TurnOn()
	TurnOff()
}

type CarLight struct{}

func (carLigth CarLight) TurnOn() {
	fmt.Println("trun on")
}
func (carLight CarLight) TurnOff() {
	fmt.Println("turn off")
}
func main() {
	var carLight CarLight
	fmt.Println("light")
	carLight.TurnOn()
	carLight.TurnOff()

}
```



</td></tr>
</tbody></table>

## 封装

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>



</td><td>




</td></tr>
</tbody></table>

## 继承

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>



</td><td>




</td></tr>
</tbody></table>

## 多态

<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>



</td><td>




</td></tr>
</tbody></table>

# 并发
## 锁
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>

```java
Lock lock = new ReentrantLock();
if(lock.tryLock()){
    try{
        // do something...
    } finally{
        lock.unlock();
    }
}

```

</td><td>

```go

```

</td></tr>
</tbody></table>

# 序列化
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>



</td><td>




</td></tr>
</tbody></table>

# 服务器
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>



</td><td>




</td></tr>
</tbody></table>

# ORM 工具
<table>
<thead><tr><th>java</th><th>go</th></tr></thead>
<tbody>
<tr><td>


- MyBatis
- Spring Data


</td><td>


- Gorm


</td></tr>
</tbody></table>
