# フィボナッチ api

## ディレクトリ構成

```
.
├── Makefile
├── README.md
├── cmd
│   └── app
│       └── main.go
├── controllers
│   ├── fibonacci_controller.go
│   ├── fibonacci_controller_test.go
│   └── fibonacci_test_cases.go
├── go.mod
├── go.sum
├── models
│   ├── fibonacci_model.go
│   ├── fibonacci_model_test.go
│   └── fibonacci_testcases.go
├── routers
│   └── router.go
└── views
    ├── response.go
    ├── response_test.go
    └── response_test_cases.go
```

## 使用技術
**言語**
- golang 1.23

**インフラ**
- Google Cloud Platform (VMインスタンス) 
  
   vCPU 1～2 vCPU（1 個の共有コア）Memory 4 GB

**エンドポイントURL**

例) http://nasu-sv.com/fib?n=99
  
成功出力 

入力 `http://nasu-sv.com/fib?n=99`
```
{
    "result": 218922995834555169026
}
```

失敗出力

入力 `http://nasu-sv.com/fib?n=abc`
```
{
    "status": 400,
    "message": "Bad request"
}
```
入力 `http://nasu-sv.com/fib?n=10000000`
```
{
    "staus": 408,
    "message": "Request time out"
}
```

## 解説
リクエストURLが与えられた時、`main.go`は`routers`でルーティングされた関数を呼び出すことによってフィボナッチ計算を行ない、その結果を出力するシステムである。


### mvcに沿って開発
**controller**
- ユーザーからの入力を受け取り、モデルとビューをつなげる役割を果たす。
- FibonacciHandler 関数は、HTTPリクエストを受け取り、フィボナッチ数を計算するためのモデル関数を呼び出し、その結果をビューで返すという役割を果たす。

**model**
- アプリケーションのビジネスロックを管理する。
- フィボナッチ数を計算する CalculateFibonacci 関数や、その結果を保持する FibonacciResult 構造体がモデルに該当する。

**view**
- コントローラから受け取った情報を基に、ユーザーに返すレスポンスを生成する。
- WriteSuccessResponse や WriteErrorResponse などの関数は、HTTPレスポンスを構築し、クライアントに返す内容を定義しています。具体的には、フィボナッチ計算の結果をJSON形式で返すなどが行われる。
