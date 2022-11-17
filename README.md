# goneis
[Open NEIS API](https://open.neis.go.kr/) wrapper for Go
## Deprecated
I will make another thing. See you soon!
## Usage
```go
func main() {
    client := NewNeisClient("")
    schools, err := client.GetSchoolInfo("", "", "선린인터넷고등학교", "", "", "")
    if err != nil {
      t.Error(err)
    }
    
    fmt.Println(schools[0].Name) // 선린인터넷고등학교
    
}
```
## Thanks to
[SaidBySolo/neispy](https://github.com/SaidBySolo/neispy)
