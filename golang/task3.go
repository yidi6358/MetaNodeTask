/* sql语句题目1 基本CRUD操作 开始*/
 func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	// 插入数据
	db.Create(&Student{Name: "张三", Age: 20, Grade: "三年级"})

	// 查询数据
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	fmt.Println("所有年龄大于18岁的学士信息：", students)

	// 更新数据
	db.Model(&Student{}).Where("name = ?", "张三").Update("grade", "四年级")

	// 删除数据
	db.Where("age < ?", 15).Delete(&Student{})

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
 /* sql语句题目1 基本CRUD操作 结束*/
     	
		

/* sql语句题目2 事务语句开始*/
func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	// 插入数据
	// db.Create(&Account{Balance: 200})
	// db.Create(&Account{Balance: 300})

	//转账
	var acountA Account
	var acountB Account
	db.Where("id = ?", 1).Find(&acountA)
	db.Where("id = ?", 2).Find(&acountB)
	if acountA.Balance < 100 {
		fmt.Println("余额不足")
		return
	} else {
		tx := db.Begin()
		acountA.Balance -= 100
		acountB.Balance += 100
		tx.Save(&acountA)
		tx.Save(&acountB)
		tx.Create(&Transaction{FromAccountId: 1, ToAccountId: 2, Amount: 100})
		// tx.Rollback()
		tx.Commit()
	}

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
/* sql语句题目2 事务语句结束*/





/* sqlx入门 题目1 使用SQL扩展库进行查询 开始*/
func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	var employees []Employee
	db.Select(&employees, "select * from employees where department = ?", "技术部")
	fmt.Println("employees数据：", employees)
	// db.Where("department = ?", "技术部").Find(&employees)

	var employee Employee
	db.Get(&employee, "select * from employees order by salary desc limit 1")
	// db.Get(&employee, "SELECT * FROM employees WHERE department = $1 ORDER BY salary DESC LIMIT 1", "技术部")
	fmt.Println("employees分数最高的数据：", employee)

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB := db.DB
	// if err != nil {
	// 	panic("获取底层数据库连接失败：" + err.Error())
	// }
	err := sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
/* sqlx入门 题目1 使用SQL扩展库进行查询 结束*/


/* sqlx入门 题目2 实现类型安全映射 开始*/
type Books struct {
	ID     uint    `gorm:"primaryKey"`
	Title  string  `gorm:"column:title"`
	Author string  `gorm:"column:author"`
	Price  float64 `gorm:"column:price"`
}

func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	var books []Books
	db.Select(&books, "select * from books where price > ?", 50)
	fmt.Println("books价格大于50的书籍", books)

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB := db.DB
	// if err != nil {
	// 	panic("获取底层数据库连接失败：" + err.Error())
	// }
	err := sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
/* sqlx入门 题目2 实现类型安全映射 结束*/


/* 进阶gorm 题目1 模型定义 开始*/
// User 定义用户表
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
}

// Post 定义文章表
type Post struct {
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"column:title"`
	Body   string `gorm:"column:body"`
	UserID uint   `gorm:"column:user_id"`
}

// Post 定义评论表
type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	Remark string `gorm:"column:remark"`
	PostID uint   `gorm:"column:post_id"`
}

func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	// 插入数据
	db.Create(&User{Name: "Alice", Email: "alice@example.com"})
	db.Create(&User{Name: "Bob", Email: "bob@example.com"})

	db.Create(&Post{Title: "文章1", Body: "内容1", UserID: 1})
	db.Create(&Post{Title: "文章2", Body: "内容2", UserID: 1})
	db.Create(&Post{Title: "文章3", Body: "内容3", UserID: 2})

	db.Create(&Comment{Remark: "文章1点评1", PostID: 1})
	db.Create(&Comment{Remark: "文章1点评2", PostID: 1})
	db.Create(&Comment{Remark: "文章2点评1", PostID: 2})

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
/* 进阶gorm 题目1 模型定义 结束*/


/* 进阶gorm 题目2 关联查询 开始*/
// User 定义用户表
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email"`
	Posts []Post `gorm:"foreignKey:UserID"`
}

// Post 定义文章表
type Post struct {
	ID       uint      `gorm:"primaryKey"`
	Title    string    `gorm:"column:title"`
	Body     string    `gorm:"column:body"`
	UserID   uint      `gorm:"column:user_id"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

// Post 定义文章表
type PostWithCount struct {
	Post
	commentCount int `gorm:"column:comment_count"`
}

// Post 定义评论表
type Comment struct {
	ID     uint   `gorm:"primaryKey"`
	Remark string `gorm:"column:remark"`
	PostID uint   `gorm:"column:post_id"`
}

func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	var posts []Post
	err := db.Model(&posts).Preload("Comments").Where("user_id = ?", 1).Find(&posts).Error
	for _, post := range posts {
		fmt.Println(post.Title)
		for _, comment := range post.Comments {
			fmt.Println(comment.Remark)
		}
	}

	var postWithCount PostWithCount
	err1 := db.Model(&Post{}).Select("posts.*,count(comments.id) as comment_count").
		Joins("left join comments on posts.id = comments.post_id").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Find(&postWithCount)
	fmt.Println(postWithCount)

	if err != nil {
		panic(err)
	}
	if err1 != nil {
		panic(err1)
	}
	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}

/* 进阶gorm 题目2 关联查询 结束*/


/* 进阶gorm 题目3 钩子函数 开始*/
func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	var user User
	tx.Model(&user).Where("id = ?", post.UserID).Update("posts_count", gorm.Expr("posts_count + ?", 1))

	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var posts Post
	tx.Model(&posts).Where("id = ?", comment.PostID).First(&posts)

	var comments []Comment
	tx.Model(&Comment{}).Where("post_id = ?", posts.ID).Find(&comments)
	if len(comments) == 0 {
		tx.Model(&posts).Where("id = ?", posts.ID).Update("status", "无评论")
	}
	return
}

var comment Comment

func main() {
	db := InitDB()
	fmt.Printf("使用的数据库文件：%s\n", constant.DBPATH)

	// 插入数据
	// error := db.Create(&Post{ID: 4, Title: "文章4", Body: "内容2", UserID: 1})
	// if error.Error != nil {
	// 	fmt.Println(error.Error)
	// 	return
	// }

	db.Model(&comment).Where("id = ?", 3).First(&comment)

	db.Delete(&Comment{ID: 3})
	// var comment Comment
	// db.Where("id = ?", 3).Delete(&comment)

	// 🔥 关键：关闭连接，强制 flush，否则不会写入到硬盘
	sqlDB, err := db.DB()
	if err != nil {
		panic("获取底层数据库连接失败：" + err.Error())
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}
/* 进阶gorm 题目3 钩子函数 结束*/
