# TODO Application with VSCode Remote Container

## Docker Containers

* web
  * フロントエンド(react)のコンテナ
* app
  * バックエンド(golang)のコンテナ
* db
  * MySQLのコンテナ

## golang

* Web API
  * gin を使用してAPIを作成。ginは宣言的にWebAPIのルーティングを記述できる

    ```
    func (tc *TaskController) getTasks() {
    tc.router.GET("/todos", func(c *gin.Context) {
        todos := tc.taskRepository.All()
        c.JSON(http.StatusOK, todos)
    })
    }
    ```

* ORM
  * gorm を使用してリポジトリを作成。
  
    ```
    type TaskRepository struct {
        db *gorm.DB
    }

    func CreateTaskRepository(db *gorm.DB) *TaskRepository {
        newTaskRepository := new(TaskRepository)
        newTaskRepository.db = db
        return newTaskRepository
    }

    func (tr *TaskRepository) All() *[]models.Todo {
        todos := []models.Todo{}
        tr.db.Find(&todos)
        return &todos
    }
    ```

* アプリケーションアーキテクチャ
  * TODOを DomainModelとして実装。構造体に状態変更の関数を作成。
    ```
    package models

    import "errors"

    type Todo struct {
        ID     int    `gorm:"AUTO_INCREMENT" json:"id"`
        Name   string `json:"name"`
        Status string `json:"status"`
    }

    func (todo *Todo) Start() error {
        if todo.Status != "Ready" {
            err := errors.New("Todo is not ready")
            return err
        }
        todo.Status = "Doing"
        return nil
    }

    func (todo *Todo) Done() error {
        if todo.Status != "Doing" {
            err := errors.New("Todo is not doing")
            return err
        }
        todo.Status = "Done"
        return nil
    }
    ```

## MySQL

* 初期データ
  * コンテナ作成時の初期データは `/docker-entrypoint-initdb.d` に配置する
  * 設定ファイルは `/etc/mysql/conf.d` に配置する

## メモ

* docker volumeは消えないので、/docker-entrypoint-initdb.d にファイルが存在している場合は、内容を変更しても変わらない
* docker compose で command: sleep infinity を指定しないと、コンテナが起動終了後に止まる？
* gormはデフォルトだと構造体を複数形にしたテーブル名を使用する
※ AutoMigrateするとテーブルが作成される