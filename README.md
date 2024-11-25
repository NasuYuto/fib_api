# フィボナッチ api

## ディレクトリ構成

```
.
├── README
├── go.mod
├── go.sum
├── handlers
│   ├── fib.go
│   └── fib_test.go
└── main.go
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

## 解説
リクエストURLが与えられた時、`main.go`は`handler`ディレクトリの`fib.go`を呼び出し、フィボナッチ計算を行ない、その結果を出力するシステムである。

### main.go
1. リクエストURLに/fibが含まれる場合 `handlers.FibonacciHandler`が呼び出される
2. port80を指定してリクエストをリッスン状態にする

### handler.FibonacciHandler関数 
1. URLのパラメータを`nStr`変数に格納する
2. `nStr`の値を引数に`fibonacci()`関数を呼び出し計算する
3. 計算結果を`Json`形式に変換し、出力
4. errの場合も同様に`Json`形式に変換して出力

### handler.fibnacci関数
1. 引数の値`n`が 0≦n≦1 であればその値のまま返す
2. `a`と`b`にそれぞれ`0`と`1`を格納する
3. `a`に`b`の値を、`b`に`a+b`の値を格納 これを`n`回続けた結果を返す
 
※黄金比などを使って一回で計算をしてしまうと、浮動小数表現で誤差が生じてしまい、計算結果がズレる


### handler/fib_test.go
1. `gotests -w -all handlers/fib.go`でtestケース以外を生成       
2. `TestFibonacciHandler()`関数では各テストケース(n=5,n=50,n=-1,n=abc)に対するレスポンスが正しく動作するかテストする
3. `TestFibonacci()`関数では各テストケース(n=0,n=1,n=10,n=50,n=100)に対するフィボナッチ数を正しく計算できるかテストする

mvcをのとって開発するとより可読性が上がる
- handlersはコントローラー部分
- アルゴリズム部分はモデル部分
- testもファイルを分けれるので対称性が上がる

testパターンを直書きした方が試行錯誤しやすくなる

speeeはaws得意なエンジニアが多い 