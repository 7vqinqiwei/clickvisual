package db

import (
	"time"

	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/shimohq/mogo/api/internal/invoker"
)

type Configuration struct {
	Name        string    `gorm:"column:name;type:varchar(64)" json:"name"`
	Content     string    `gorm:"column:content;type:longtext" json:"content"`
	Format      string    `gorm:"column:format;type:varchar(32)" json:"format"`
	Version     string    `gorm:"column:version;type:varchar(64)" json:"version"`
	Uid         uint      `gorm:"column:uid;type:int(11) unsigned" json:"uid"`
	PublishTime int64     `gorm:"column:publish_time;type:int(11)" json:"publishTime"`
	LockUid     uint      `gorm:"column:lock_uid;type:int(11) unsigned" json:"lockUid"`
	LockAt      time.Time `gorm:"column:lock_at;type:datetime" json:"lockAt"`

	BaseModel
}

func (m *Configuration) TableName() string {
	return TableNameConfiguration
}

// ConfigurationCreate CRUD
func ConfigurationCreate(db *gorm.DB, data *Configuration) (err error) {
	if err = db.Create(data).Error; err != nil {
		elog.Error("create cluster error", zap.Error(err))
		return
	}
	return
}

// ConfigurationUpdate ...
func ConfigurationUpdate(db *gorm.DB, paramId int, ups map[string]interface{}) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	if err = db.Table(TableNameConfiguration).Where(sql, binds...).Updates(ups).Error; err != nil {
		elog.Error("update cluster error", zap.Error(err))
		return
	}
	return
}

func ConfigurationInfo(paramId int) (resp Configuration, err error) {
	var sql = "`id`= ? and dtime = 0"
	var binds = []interface{}{paramId}
	if err = invoker.Db.Table(TableNameConfiguration).Where(sql, binds...).First(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("cluster info error", zap.Error(err))
		return
	}
	return
}

// ConfigurationDelete 软删除
func ConfigurationDelete(db *gorm.DB, id int) (err error) {
	if err = db.Model(Configuration{}).Delete(&Configuration{}, id).Error; err != nil {
		elog.Error("cluster delete error", zap.Error(err))
		return
	}
	return
}

// ConfigurationList 获取当前所有未删除的clusters. 主要供 前端用
func ConfigurationList(conds egorm.Conds) (resp []*Configuration, err error) {
	conds["dtime"] = 0
	sql, binds := egorm.BuildQuery(conds)
	// Fetch record with Rancher Info....
	if err = invoker.Db.Table(TableNameConfiguration).Where(sql, binds...).Find(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("list clusters error", elog.String("err", err.Error()))
		return
	}
	return
}

// ConfigurationListPage 根据分页条件查询list
func ConfigurationListPage(conds egorm.Conds, reqList *ReqPage) (total int64, respList []*Configuration) {
	respList = make([]*Configuration, 0)
	conds["dtime"] = 0
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := egorm.BuildQuery(conds)
	db := invoker.Db.Table(TableNameConfiguration).Where(sql, binds...)
	db.Count(&total)
	if reqList.Sort != "" {
		db = db.Order(reqList.Sort)
	}
	db.Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}

type ConfigurationHistory struct {
	Uid             uint   `gorm:"column:uid;type:int(11) unsigned" json:"uid"`
	ConfigurationId uint   `gorm:"column:configuration_id;type:int(11) unsigned" json:"configurationId"`
	ChangeLog       string `gorm:"column:change_log;type:longtext" json:"changeLog"`
	Content         string `gorm:"column:content;type:longtext" json:"content"`
	Version         string `gorm:"column:version;type:varchar(64)" json:"version"`

	BaseModel
}

func (m *ConfigurationHistory) TableName() string {
	return TableNameConfigurationHistory
}

// ConfigurationHistoryCreate CRUD
func ConfigurationHistoryCreate(db *gorm.DB, data *ConfigurationHistory) (err error) {
	if err = db.Create(data).Error; err != nil {
		elog.Error("ConfigurationHistoryCreate cluster error", zap.Error(err))
		return
	}
	return
}

// ConfigurationHistoryUpdate ...
func ConfigurationHistoryUpdate(db *gorm.DB, paramId int, ups map[string]interface{}) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	if err = db.Table(TableNameConfigurationHistory).Where(sql, binds...).Updates(ups).Error; err != nil {
		elog.Error("ConfigurationHistoryUpdate cluster error", zap.Error(err))
		return
	}
	return
}

func ConfigurationHistoryInfo(paramId int) (resp ConfigurationHistory, err error) {
	var sql = "`id`= ? and dtime = 0"
	var binds = []interface{}{paramId}
	if err = invoker.Db.Table(TableNameConfigurationHistory).Where(sql, binds...).First(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("ConfigurationHistoryInfo info error", zap.Error(err))
		return
	}
	return
}

// ConfigurationHistoryInfoX Info的扩展方法，根据Cond查询单条记录
func ConfigurationHistoryInfoX(conds map[string]interface{}) (resp ConfigurationHistory, err error) {
	conds["dtime"] = 0
	sql, binds := egorm.BuildQuery(conds)
	if err = invoker.Db.Table(TableNameConfigurationHistory).Where(sql, binds...).First(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("ConfigurationHistoryInfoX infoX error", zap.Error(err))
		return
	}
	return
}

// ConfigurationHistoryDelete 软删除
func ConfigurationHistoryDelete(db *gorm.DB, id int) (err error) {
	if err = db.Model(ConfigurationHistory{}).Delete(&ConfigurationHistory{}, id).Error; err != nil {
		elog.Error("ConfigurationHistoryDelete delete error", zap.Error(err))
		return
	}
	return
}

// ConfigurationHistoryList 获取当前所有未删除的clusters. 主要供 前端用
func ConfigurationHistoryList(conds egorm.Conds) (resp []*ConfigurationHistory, err error) {
	conds["dtime"] = 0
	sql, binds := egorm.BuildQuery(conds)
	// Fetch record with Rancher Info....
	if err = invoker.Db.Table(TableNameConfigurationHistory).Where(sql, binds...).Find(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("ConfigurationHistoryList clusters error", elog.String("err", err.Error()))
		return
	}
	return
}

// ConfigurationHistoryListPage 根据分页条件查询list
func ConfigurationHistoryListPage(conds egorm.Conds, reqList *ReqPage) (total int64, respList []*ConfigurationHistory) {
	respList = make([]*ConfigurationHistory, 0)
	conds["dtime"] = 0
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := egorm.BuildQuery(conds)
	db := invoker.Db.Table(TableNameConfigurationHistory).Where(sql, binds...)
	db.Count(&total)
	if reqList.Sort != "" {
		db = db.Order(reqList.Sort)
	}
	db.Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}

type ConfigurationPublish struct {
	Uid                    uint `gorm:"column:uid;type:int(11) unsigned" json:"uid"`
	ConfigurationId        uint `gorm:"column:configuration_id;type:int(11) unsigned" json:"configurationId"`
	ConfigurationHistoryId uint `gorm:"column:configuration_history_id;type:int(11) unsigned" json:"configurationHistoryId"`

	BaseModel
}

func (m *ConfigurationPublish) TableName() string {
	return TableNameConfigurationPublish
}

// ConfigurationPublishCreate CRUD
func ConfigurationPublishCreate(db *gorm.DB, data *ConfigurationPublish) (err error) {
	if err = db.Create(data).Error; err != nil {
		elog.Error("ConfigurationPublishCreate cluster error", zap.Error(err))
		return
	}
	return
}

// ConfigurationPublishUpdate ...
func ConfigurationPublishUpdate(db *gorm.DB, paramId int, ups map[string]interface{}) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	if err = db.Table(TableNameConfigurationPublish).Where(sql, binds...).Updates(ups).Error; err != nil {
		elog.Error("ConfigurationPublishUpdate cluster error", zap.Error(err))
		return
	}
	return
}

func ConfigurationPublishInfo(paramId int) (resp ConfigurationPublish, err error) {
	var sql = "`id`= ? and dtime = 0"
	var binds = []interface{}{paramId}
	if err = invoker.Db.Table(TableNameConfigurationPublish).Where(sql, binds...).First(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("ConfigurationPublishInfo info error", zap.Error(err))
		return
	}
	return
}

// ConfigurationPublishDelete 软删除
func ConfigurationPublishDelete(db *gorm.DB, id int) (err error) {
	if err = db.Model(ConfigurationPublish{}).Delete(&ConfigurationPublish{}, id).Error; err != nil {
		elog.Error("ConfigurationPublishDelete delete error", zap.Error(err))
		return
	}
	return
}

// ConfigurationPublishList 获取当前所有未删除的clusters. 主要供 前端用
func ConfigurationPublishList(conds egorm.Conds) (resp []*ConfigurationPublish, err error) {
	conds["dtime"] = 0
	sql, binds := egorm.BuildQuery(conds)
	// Fetch record with Rancher Info....
	if err = invoker.Db.Table(TableNameConfigurationPublish).Where(sql, binds...).Find(&resp).Error; err != nil && err != gorm.ErrRecordNotFound {
		elog.Error("ConfigurationPublishList clusters error", elog.String("err", err.Error()))
		return
	}
	return
}

// ConfigurationPublishListPage 根据分页条件查询list
func ConfigurationPublishListPage(conds egorm.Conds, reqList *ReqPage) (total int64, respList []*ConfigurationPublish) {
	respList = make([]*ConfigurationPublish, 0)
	conds["dtime"] = 0
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := egorm.BuildQuery(conds)
	db := invoker.Db.Table(TableNameConfigurationPublish).Where(sql, binds...)
	db.Count(&total)
	if reqList.Sort != "" {
		db = db.Order(reqList.Sort)
	}
	db.Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
