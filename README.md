# Ganboard

JSON-RPC client for Kanboard written in Golang

Aims to implement every functions of Kanboard API

```go
import (
    "github.com/davecgh/go-spew/spew"

    "github.com/chrnin/ganboard"
)

func main() {
    client := ganboard.Client{
        Endpoint: "http://localhost/kanboard/jsonrpc.php",
        Username: "admin",
        Password: "admin",
    }
    Project, _ := client.GetProjectById(1)
    spew.Dump(Project)
}
```

```go
(ganboard.Project) {
 ID: (int) 1,
 Name: (string) (len=18) "My beloved Project",
 IsActive: (int) 1,
 Token: (string) (len=60) "d0541d81b3331b08c324e569cd57dd15d01f9e43dfd850018b9e29ecaa34",
 LastModified: (int) 1524755851,
 IsPublic: (int) 1,
 IsPrivate: (int) 0,
 DefaultSwimlane: (string) (len=16) "Default swimlane",
 ShowDefaultSwimlane: (int) 1,
 Description: (string) "",
 Identifier: (string) "",
 URL: (struct { Board string "json:\"board\""; Calendar string "json:\"calendar\""; List string "json:\"list\"" }) {
  Board: (string) (len=88) "http://localhost/kanboard/?controller=BoardViewController&action=show&project_id=1",
  Calendar: (string) "",
  List: (string) (len=87) "http://localhost/kanboard/?controller=TaskListController&action=show&project_id=1"
 }
}
```
