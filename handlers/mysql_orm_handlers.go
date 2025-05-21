
package handlers
import (

	"github.com/gin-gonic/gin"
	"net/http"

    // "fmt"
    // "time"
    // "database/sql"
    // "github.com/folospace/go-mysql-orm/orm"
)
// var db, _ = orm.OpenMysql("user:pass@tcp(ip_address)/golang_learn_mysql_db?parseTime=true")

// var UserTable = new(User)
// type User struct {
// 	Id        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	CategoryID string    `json:"category_id"` 
// 	CreatedAt time.Time `json:"created_at"`
// }
// func (User) Connections() []sql.DB {
//     return []sql.DB{db}
// }
// func (User) DatabaseName() string {
//     return "golang_learn_mysql_db"
// }
// func (User) TableName() string {
//     return "users"
// }
// func (u *User) Query() orm.Query[User] {
//     return orm.NewQuery(UserTable).WherePrimaryIfNotZero(u.Id)
// }


func MysqlSelectOneOrm(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message":  "Not working as expected yet",
	})


}