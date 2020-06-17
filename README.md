# splitio is not good performance but use low memory
It is for read big file in not enough memory system  
Read chunks up to the separator and return 

this is sample code
```golang
buf := bytes.NewBufferString("hello world")
read := splitio.New(buf, []byte(" "))
for i := 0; i < 4; i++ {
    sub, err := read.Next()
    if err != nil && err != io.EOF {
        panic(err)
    }
    fmt.Println(string(sub))
    if err == io.EOF {
        break
    }
}
```
output
> hello  
> world