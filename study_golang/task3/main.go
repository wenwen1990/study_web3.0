package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "wenxiao:1009zz1211@tcp(127.0.0.1:3306)/task3?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         200,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Init(db)
	// question1(db)
	// question2(db)

	// 替换为sqlxDB
	// dsn := "wenxiao:1009zz1211@tcp(127.0.0.1:3306)/task3?parseTime=true"
	// sqlxDB, sqlxErr := sqlx.Connect("mysql", dsn)
	// if sqlxErr != nil {
	// 	log.Fatalln("数据库连接失败:", err)
	// }
	// question3(sqlxDB)
	// question4(sqlxDB)
	question5(db)
	// question6(db)
	question7(db)

}

/*
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录go get -u gorm.io/driver/mysql，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/
func question1(db *gorm.DB) {
	student := Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	student2 := Student{
		Name:  "李四",
		Age:   14,
		Grade: "二年级",
	}
	result := db.First(&student)
	if result.RowsAffected == 0 {
		db.Create(&student)
	} else {
		db.Model(&student).Update("grade", "三年级")
	}
	result2 := db.Select("name", "李四").First(&student2)
	if result2.RowsAffected == 0 {
		db.Create(&student2)
	}
	students := []Student{}
	db.Where("age > ?", 18).Find(&students)
	fmt.Println(students)
	// 遍历后逐条更新
	// for _, value := range students {
	// 	db.Model(&value).Update("grade", "四年级")
	// }
	// 批量更新
	db.Model(&students).Updates(Student{Grade: "四年级"})
	db.Where("age > ?", 18).Find(&students)
	fmt.Println(students)
	fmt.Println(student2)
	db.Where("name = ?", "李四").Delete(&student2)
	result3 := db.First(&student2)
	if result3.RowsAffected == 0 {
		fmt.Println(student2, "已经删除")
	}

}

type Student struct {
	Id        uint
	Name      string
	Age       uint
	Grade     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Init(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&Account{}, &Transaction{})
	db.AutoMigrate(&Employee{})
	var employees []Employee
	db.Find(&employees)
	if len(employees) == 0 {
		employees = []Employee{
			{Name: "张三", Department: "技术部", Salary: 8000.00},
			{Name: "李四", Department: "销售部", Salary: 10000.00},
			{Name: "王五", Department: "技术部", Salary: 9000.00},
			{Name: "赵六", Department: "售后部", Salary: 6000.00},
		}
		db.Create(&employees)
	}
	db.AutoMigrate(&Book{})
	var books []Book
	db.Find(&books)
	if len(books) == 0 {
		books = []Book{
			{Title: "C语言开发", Author: "张三", Price: 30.00},
			{Title: "java语言开发", Author: "李四", Price: 80.00},
			{Title: "go语言开发", Author: "王五", Price: 100.00},
			{Title: "php语言开发", Author: "赵六", Price: 40.00},
		}
		db.Create(&books)
	}
	db.AutoMigrate(&User{}, &Post{}, &Comment{})
}

/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键，
from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/
func question2(db *gorm.DB) {
	var accounts []Account
	db.Find(&accounts)

	// 如果表为空，初始化账户
	if len(accounts) == 0 {
		accounts = []Account{
			{Balance: 90},  // A
			{Balance: 0},   // B
			{Balance: 120}, // C
			{Balance: 0},   // D
		}
		db.Create(&accounts)
	}

	fmt.Println("转账前：", accounts)
	fmt.Println("尝试A转100给B，C转100给D")
	if err := transferMoney(db, 100, 200, 100); err != nil {
		fmt.Println("", err)
	}
	if err := transferMoney(db, accounts[0].Id, accounts[1].Id, 100); err != nil {
		fmt.Println("", err)
	}
	if err := transferMoney(db, accounts[2].Id, accounts[3].Id, 100); err != nil {
		fmt.Println(err)
		if err := transferMoney(db, accounts[3].Id, accounts[2].Id, 100); err != nil {
			fmt.Println(err)
		}
	}

	var afterTransfer []Account
	db.Find(&afterTransfer)
	fmt.Println("转账后：", afterTransfer)
}

type Account struct {
	Id      uint
	Balance float64
}

type Transaction struct {
	Id            uint
	FromAccountId uint
	ToAccountId   uint
	Amount        float64
}

/*
func (Transaction *Transaction) BeaforeCreate(db *gorm.DB, money float64) (err error) {
	targetTransaction := *Transaction
	if targetTransaction.FromAccountId <= 0 || targetTransaction.ToAccountId <= 0 {
		return errors.New("转账账号和目标账户id不能为空")
	}
	var fromAccount = Account{
		Id: targetTransaction.FromAccountId,
	}
	var toAccount = Account{
		Id: targetTransaction.ToAccountId,
	}
	result := db.Find(&fromAccount)
	result2 := db.Find(&toAccount)
	if result.RowsAffected == 0 || result2.RowsAffected == 0 {
		return errors.New("转账账号或目标账户不存在")
	}
	if fromAccount.Balance < money {
		return errors.New("转账账号金额不足")
	}
	return
}
*/

func transferMoney(db *gorm.DB, fromAccountId uint, toAccountId uint, money float64) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		var fromAccount, toAccount Account
		if err := tx.First(&fromAccount, fromAccountId).Error; err != nil {
			return errors.New("转账账号不存在")
		}
		if err := tx.First(&toAccount, toAccountId).Error; err != nil {
			return errors.New("目标账户不存在")
		}
		if fromAccount.Balance < money {
			return errors.New("转账账号余额不足")
		}
		if err := tx.Model(&fromAccount).Update("balance", fromAccount.Balance-money).Error; err != nil {
			return err
		}
		if err := tx.Model(&toAccount).Update("balance", toAccount.Balance+money).Error; err != nil {
			return err
		}
		var transactionRecord = Transaction{
			FromAccountId: fromAccountId,
			ToAccountId:   toAccountId,
			Amount:        money,
		}
		if err := tx.Create(&transactionRecord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

/*
题目3：使用SQL扩展库进行查询
假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
要求 ：
编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
*/
func question3(db *sqlx.DB) {
	// 查询所有技术部员工
	var techEmployees []Employee
	err := db.Select(&techEmployees, "SELECT * FROM employees WHERE department = ?", "技术部")
	if err != nil {
		panic(err)
	}
	fmt.Println("技术部员工:", techEmployees)

	// 查询工资最高的员工
	var topEmployee Employee
	err = db.Get(&topEmployee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n工资最高的员工:", topEmployee)

}

type Employee struct {
	Id         uint
	Name       string
	Department string
	Salary     float64
}

/*
题目4：实现类型安全映射
假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
要求 ：
定义一个 Book 结构体，包含与 books 表对应的字段。
编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
*/
func question4(db *sqlx.DB) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price>?", 50.00)
	if err != nil {
		panic(err)
	}
	fmt.Println("价格大于 50 元的书籍:", books)
}

type Book struct {
	Id     uint
	Title  string
	Author string
	Price  float64
}

/*
题目5：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
func question5(db *gorm.DB) {
	var users []User
	db.Find(&users)
	if len(users) == 0 {
		users = []User{
			{Name: "张三"},
			{Name: "李四"},
			{Name: "王五"},
		}
		db.Create(&users)
		var posts = []Post{
			{UserID: users[0].ID, Title: "文章A-1"},
			{UserID: users[0].ID, Title: "文章A-2"},
			{UserID: users[0].ID, Title: "文章A-3"},
			{UserID: users[1].ID, Title: "文章B-1"},
			{UserID: users[2].ID, Title: "文章C-1"},
			{UserID: users[2].ID, Title: "文章C-2"},
		}
		db.Create(&posts)
		var comments = []Comment{
			{PostID: posts[0].ID, Content: "评价A-1"},
			{PostID: posts[0].ID, Content: "评价A-2"},
			{PostID: posts[0].ID, Content: "评价A-3"},
			{PostID: posts[1].ID, Content: "评价B-1"},
			{PostID: posts[1].ID, Content: "评价B-2"},
			{PostID: posts[3].ID, Content: "评价C-1"},
		}
		db.Create(&comments)
		for _, p := range posts {
			var count int64
			db.Model(&Comment{}).Where("post_id = ?", p.ID).Count(&count)
			// fmt.Printf("count: %v\n", count)
			db.Model(p).Update("comment_num", count)
		}
	}
	fmt.Println("初始化3张表及数据完毕")
}

type User struct {
	ID      uint
	Name    string
	PostNum uint   `gorm:"default:0"`
	Posts   []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	ID           uint
	UserID       uint
	Title        string
	CommentNum   uint
	CommentState string
	Comments     []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      uint
	PostID  uint
	Content string
}

/*
题目6：关联查询
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息
*/
func question6(db *gorm.DB) {
	var user User
	// 查询某个用户 "张三"发布的所有文章及其对应的评论信息。
	var userName = "张三"
	if err := db.Preload("Posts.Comments").Where("name =?", userName).First(&user).Error; err != nil {
		panic(err)
	}
	fmt.Printf("用户：%s\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("文章：%s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf(" - 评论：%s\n", comment.Content)
		}
	}
	// 使用Gorm查询评论数量最多的文章信息。
	var post Post
	if err := db.Model(&post).
		Select("posts.*,COUNT(c.id) as commentNum").
		Joins("left join comments c ON c.post_id=posts.id").
		Group("posts.id").
		Order("commentNum desc").
		Limit(1).
		Preload("Comments").
		Find(&post).Error; err != nil {
		panic(err)
	}
	fmt.Println(post)
}

/*
题目7：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/
func question7(db *gorm.DB) {
	var userId uint = 1
	post := Post{Title: "新增一篇文章", UserID: userId}
	db.Create(&post)
	fmt.Println(post)

	var commnet = Comment{ID: 6}
	db.First(&commnet)
	fmt.Println("commnet", commnet)
	db.Delete(&commnet)
	var post2 Post
	db.Model(&post2).
		Preload("Comments").
		First(&post2, 4)
	fmt.Println(post2)

	db.Model(&post2).
		Preload("Comments").
		First(&post2)
	fmt.Println(post2)
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.
		Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_num", gorm.Expr("IFNULL(post_num,0)+?", 1)).
		Error
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&Comment{}).
		Where("post_id =?", c.PostID).
		Count(&count)
	if count == 0 {
		return tx.Model(&Post{}).
			Where("id =?", c.PostID).
			Updates(Post{CommentNum: uint(count), CommentState: "无评论"}).Error
	} else {
		return tx.Model(&Post{}).
			Where("id =?", c.PostID).
			Update("comment_num", count).Error
	}
}
