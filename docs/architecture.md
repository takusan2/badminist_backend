# アーキテクチャ

## ディレクトリ構成

```
- api
	- domain
	    - user
	    - community
	    - player
	- infrastructure
	- interface_adaptor_if
		- dao_if
			- command
			- query
		- repository_if
			- command
			- query
	- interface_adaptor_impl
		- controller
			- command
			- query
	    - dao
		    - command
			- query
	    - dto
		    - command
			- query
	    - repository
		    - command
    - processor
	    - command
		    - command_processor
		- query
			- read_model
			- query_processor
	- router
	- utils
```

## domain
ドメインオブジェクト、値オブジェクトによってドメインロジックを記述する

## interface_adaptor_if
永続化層のインターフェースを定義する。ユースケース層（processor）はこのインターフェースを参照する。

## interface_adaptor_impl
interface_adaptor_if で定義したインターフェースを実装するとともに、entity、dto、controllerを実装する。
- entity
	- データベースへ保存するときのデータモデル
- dto
	- http リクエスト、レスポンス時の方を定義
- controller
	- リクエストに対する関数を定義。バリデーションを行いユースケース層（processor）で定義された処理を呼び出す。
## processor
ユースケース層。コマンド側では、ドメインで定義されたロジックを呼び出し、そのイベントを用いて永続化の処理を行う。クエリ側はリードモデルとdao_ifを用いて、クエリ処理を定義する。