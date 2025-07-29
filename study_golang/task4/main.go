package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	Init()
	r := gin.Default()
	question6(r)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "wenxiao:1009zz1211@tcp(127.0.0.1:3306)/task4?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         200,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// question1()
	question2(db)
	question3(db, r)

}

/*
*
1、项目初始化
创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理。
安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）。
*
*/
func question1() {
	// 依次执行
	// go mod init test
	// 新建main.go主入口文件
	// 增加gin依赖 go get -u github.com/gin-gonic/gin
	// 增加gorm依赖 go get -u gorm.io/gorm
	// 增加mysql依赖 go get -u gorm.io/driver/mysql
}

/*
*
2、数据库设计与模型定义
设计数据库表结构，至少包含以下几个表：
users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。
使用 GORM 定义对应的 Go 模型结构体。
*
*/
func question2(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Post{}, &Comment{})

}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"  json:"username"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"uniqueIndex"  json:"email"`
	Posts    []Post
	Comments []Comment
}

type Post struct {
	ID        uint   `gorm:"primarykey"`
	UserId    uint   `gorm:"not null"`
	Title     string `gorm:"type:text" json:"title"`
	Content   string `json:"content"`
	User      User
	Comments  []Comment
	CreatedAt time.Time `json:"created_at" format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at" format:"2006-01-02 15:04:05"`
}

type Comment struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	UserId    uint      `json:"userId"`
	PostId    uint      `json:"postId"`
	Content   string    `gorm:"type:text" json:"content"`
	User      User      `json:"-"`
	Post      Post      `json:"-"`
	CreatedAt time.Time `json:"created_at" format:"2006-01-02 15:04:05"`
}

/*
*
3、用户认证与授权
实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
*
*/
func question3(db *gorm.DB, r *gin.Engine) {
	// go get github.com/golang-jwt/jwt/v5

	r.POST("/register", func(ctx *gin.Context) {
		user := User{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ErrorResponse(ctx, http.StatusBadRequest, "参数异常", err)
			return
		}
		if len(user.Username) == 0 || len(user.Password) == 0 || len(user.Email) == 0 {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户账密邮箱不能为空！"})
			ErrorResponse(ctx, http.StatusBadRequest, "用户账密邮箱不能为空！", errors.New("err"))
			return
		}
		db.Where("username =?", user.Username).Or("email=?", user.Email).Find(&user)
		if user.ID > 0 {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "当前用户名或邮箱已注册！"})
			ErrorResponse(ctx, http.StatusBadRequest, "当前用户名或邮箱已注册！", errors.New("err"))
			return
		}
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		db.Create(&user)
		// ctx.JSON(http.StatusOK, user)
		SuccessResponse(ctx, user, "")
	})

	r.POST("/login", func(ctx *gin.Context) {
		user := User{}
		err := ctx.ShouldBind(&user)
		if err != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ErrorResponse(ctx, http.StatusBadRequest, "参数异常", err)
			return
		}
		if len(user.Username) == 0 || len(user.Password) == 0 {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户账密不能为空！"})
			ErrorResponse(ctx, http.StatusBadRequest, "用户账密不能为空！", errors.New("err"))
			return
		}
		dbUser := User{}
		result := db.Where("username =?", user.Username).Find(&dbUser)
		if result.Error != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "用户不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "用户不存在！", result.Error)
			return
		}
		pwdErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
		if pwdErr != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "输入密码不正确！"})
			ErrorResponse(ctx, http.StatusBadRequest, "输入密码不正确！", pwdErr)
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"userID": dbUser.ID,
			"exp":    time.Now().Add(time.Hour * 24).Unix(),
		})
		jwtToken, _ := token.SignedString([]byte("salt"))
		// ctx.JSON(http.StatusOK, gin.H{"token": jwtToken})
		SuccessResponse(ctx, jwtToken, "")

	})

	r.Use(AuthMiddleware())
	question4(db, r)
	question5(db, r)

	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

// func Middleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("接口之前")
// 		c.Next()
// 		fmt.Println("接口之后")
// 	}
// }

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if len(tokenStr) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			})
			// ErrorResponse(c, http.StatusUnauthorized, "unauthorized", errors.New("err"))
			return
		}
		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return []byte("salt"), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userID", uint(claims["userID"].(float64)))
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "unauthorized",
			}) // ErrorResponse(c, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
	}
}

/*
*
4、文章管理功能
实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
实现文章的更新功能，只有文章的作者才能更新自己的文章。
实现文章的删除功能，只有文章的作者才能删除自己的文章。
*
*/
func question4(db *gorm.DB, r *gin.Engine) {
	r.POST("/post", func(ctx *gin.Context) {
		var post Post
		userID := ctx.MustGet("userID").(uint)
		err := ctx.ShouldBind(&post)
		if err != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ErrorResponse(ctx, http.StatusBadRequest, "参数异常", err)
			return
		}
		if len(post.Title) == 0 || len(post.Content) == 0 {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "新增文章标题及内容不能为空！"})
			ErrorResponse(ctx, http.StatusBadRequest, "新增文章标题及内容不能为空！", err)
			return
		}
		post.UserId = userID
		db.Create(&post)
		db.Model(&post).Preload("User").Preload("Comments").Find(&post)
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "post": post})
		SuccessResponse(ctx, post, "")
	})

	r.GET("/post/list", func(ctx *gin.Context) {
		var posts []Post
		db.Preload("User").Preload("Comments").Find(&posts)
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "posts": posts})
		SuccessResponse(ctx, posts, "")
	})

	r.GET("/post/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var post Post
		result := db.Preload("User").Preload("Comments").First(&post, id)
		if result.Error != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "文章不存在", result.Error)
			return
		}
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "post": post})
		SuccessResponse(ctx, post, "")
	})

	r.PUT("/post/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var post Post
		var dbPost Post
		err := ctx.ShouldBind(&post)
		userID := ctx.MustGet("userID").(uint)
		if err != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ErrorResponse(ctx, http.StatusBadRequest, "参数异常", err)
			return
		}
		db.Preload("User").Preload("Comments").First(&dbPost, id)
		if dbPost.ID <= 0 {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "文章不存在！", err)
			return
		}
		if dbPost.UserId != userID {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "只能更新自己的文章！"})
			ErrorResponse(ctx, http.StatusBadRequest, "只能更新自己的文章！", errors.New("err"))
			return
		}
		db.Model(&dbPost).Updates(Post{Title: post.Title, Content: post.Content})
		db.Preload("User").Preload("Comments").First(&dbPost, id)
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "post": dbPost})
		SuccessResponse(ctx, dbPost, "")
	})

	r.DELETE("/post/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		var dbPost Post
		userID := ctx.MustGet("userID").(uint)
		result := db.Preload("User").Preload("Comments").First(&dbPost, id)
		if result.Error != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "文章不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "文章不存在！", result.Error)
			return
		}
		if dbPost.UserId != userID {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "只能更新自己的文章！"})
			ErrorResponse(ctx, http.StatusBadRequest, "只能删除自己的文章！", errors.New("err"))
			return
		}
		db.Delete(&dbPost)
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "post": dbPost})
		SuccessResponse(ctx, dbPost, "")
	})
}

/*
*
5、评论功能
实现评论的创建功能，已认证的用户可以对文章发表评论。
实现评论的读取功能，支持获取某篇文章的所有评论列表。
*
*/
func question5(db *gorm.DB, r *gin.Engine) {

	r.POST("/comment", func(ctx *gin.Context) {
		var comment Comment
		var user User
		var post Post
		err := ctx.ShouldBind(&comment)
		userID := ctx.MustGet("userID").(uint)
		if err != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			ErrorResponse(ctx, http.StatusBadRequest, "参数异常", err)
			return
		}
		if len(comment.Content) == 0 || comment.PostId <= 0 {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "评论需要内容和文章id！"})
			ErrorResponse(ctx, http.StatusBadRequest, "评论需要内容和文章id！", errors.New("err"))
			return
		}
		result := db.Find(&user, userID)
		if result.Error != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "当前用户不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "当前用户不存在！", result.Error)
			return
		}
		result2 := db.Find(&post, comment.PostId)
		if result2.Error != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "当前文章不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "当前文章不存在！", result2.Error)
			return
		}
		comment.UserId = userID
		db.Create(&comment)
		db.Preload("User").Preload("Post").First(&comment, comment.ID)
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "comment": comment})
		SuccessResponse(ctx, comment, "")
	})

	r.GET("/post/:id/comment/list", func(ctx *gin.Context) {
		postId := ctx.Param("id")
		var post Post
		result := db.Find(&post, postId)
		if result.Error != nil {
			// ctx.JSON(http.StatusBadRequest, gin.H{"error": "当前文章不存在！"})
			ErrorResponse(ctx, http.StatusBadRequest, "当前文章不存在！", result.Error)
			return
		}
		db.Preload("User").Preload("Comments").First(&post, postId)
		comments := post.Comments
		db.Preload("User").Preload("Post").Find(&comments)
		for _, value := range comments {
			value.Post = post
		}
		// ctx.JSON(http.StatusOK, gin.H{"message": "success", "comments": comments})
		SuccessResponse(ctx, comments, "")
	})

}

// type JSONTime time.Time

// func (t JSONTime) MarshalJSON() ([]byte, error) {
// 	formatted := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
// 	return []byte(formatted), nil
// }

/*
*
6、错误处理与日志记录
对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。
*
*/
func question6(r *gin.Engine) {
	// go get github.com/sirupsen/logrus
	r.Use(LoggerMiddleware())
	r.Use(RecoveryMiddleware())
}

// 初始化日志配置
func Init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true, // 打印完整时间戳
	})
	log.SetLevel(log.InfoLevel) // 默认日志级别为 Info
	log.SetReportCaller(true)   // 打印调用日志的文件和行号
}

func ErrorResponse(ctx *gin.Context, statusCode int, message string, err error) {
	log.WithFields(log.Fields{
		"status_code": statusCode,
		"error":       err.Error(),
	}).Error(message) // 记录错误日志

	ctx.JSON(statusCode, gin.H{
		"code":    statusCode,
		"message": message,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}, message string) {
	log.WithFields(log.Fields{
		"status_code": http.StatusOK,
		"message":     message,
	}).Info(message) // 记录成功日志

	if len(message) == 0 {
		message = "success"
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.WithFields(log.Fields{
					"error": r,
				}).Error("服务器出现异常")
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "服务器出现错误，请稍后重试",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // 执行后续处理
		duration := time.Since(start)

		log.WithFields(log.Fields{
			"status":   c.Writer.Status(),
			"method":   c.Request.Method,
			"path":     c.Request.URL.Path,
			"costTime": duration,
			"clientIP": c.ClientIP(),
		}).Info("Request handled")
	}
}
