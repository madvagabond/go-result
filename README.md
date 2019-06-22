# result-go 

A result monad for go. 


```go

ok := result.Ok("hello")

hello_world := ok.Map(func(x interface{}) interface{} {
  return x.(string) + " world"
}) 


```

