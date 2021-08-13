package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/log"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

// Initialize the model from a string.

// [request_definition] 请求定义
// 定义请求由三部分组成 访问用户的用户 Subject , 访问的资源 Object 访问的动作 Action

// [policy_definition] 策略定义
// 定策略的格式, 参数的基本意思和定义请求的相同 ,定义好了策略格式,那么对于策略(Policy)的具体描述可以存放在一个以 .csv 作为后缀的文件中

// [policy_effect] 政策定义

// [matchers] [匹配器]

var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup(db *gorm.DB, prefix, tableName string) *casbin.SyncedEnforcer {
	// mysql 适配器
	Apter, err := gormAdapter.NewAdapterByDBUseTableName(db, prefix, tableName)
	if err != nil {
		panic(err)
	}

	// 从文本中创建模型
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}

	// 通过文件或数据库创建同步的enforcer
	e, err := casbin.NewSyncedEnforcer(m, Apter)
	if err != nil {
		panic(err)
	}

	// 从文件/数据库重新加载策略。
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// 设置当前记录器
	log.SetLogger(&Logger{})
	e.EnableLog(true)
	return e
}
