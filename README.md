# sandbox-docker-golang

* docker volumeは消えないので、/docker-entrypoint-initdb.d にファイルが存在している場合は、内容を変更しても変わらない
* docker compose で command: sleep infinity を指定しないと、コンテナが起動終了後に止まる？
* gormはデフォルトだと構造体を複数形にしたテーブル名を使用する
※ AutoMigrateするとテーブルが作成される