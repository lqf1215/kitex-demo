package modal

// User 用户
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"` //用户名
	Password string `json:"-"`        //密码
	Sex      int8   `json:"sex"`      // 性别

}

func (User) TableName() string {
	return "user"
}
