package repository

import (
	"test_gin/common/datasource"
	"test_gin/common/logger"

	"gorm.io/gorm"
)

type BaseRepository struct {
	Source datasource.IDb `inject:""`
	Log    logger.ILogger `inject:""`
}

// Create 创建实体
func (b *BaseRepository) Create(value interface{}) error {
	return b.Source.DB().Create(value).Error
}

// Save 保存实体
func (b *BaseRepository) Save(value interface{}) error {
	return b.Source.DB().Save(value).Error
}

// Updates 更新实体
func (b *BaseRepository) Updates(model interface{}, value interface{}) error {
	return b.Source.DB().Model(model).Updates(value).Error
}

// DeleteByWhere 根据条件删除实体
func (b *BaseRepository) DeleteByWhere(model interface{}, where interface{}) (count int64, err error) {
	db := b.Source.DB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		b.Log.Errorf("删除实体出错", err)
		return
	}
	count = db.RowsAffected
	return
}

// DeleteByID 根据id删除实体
func (b *BaseRepository) DeleteByID(model interface{}, id int) error {
	return b.Source.DB().Where("id=?", id).Delete(model).Error
}

// DeleteByIDS 根据多个id删除多个实体
func (b *BaseRepository) DeleteByIDS(model interface{}, ids []int) (count int64, err error) {
	db := b.Source.DB().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		b.Log.Errorf("删除多个实体出错", err)
		return
	}
	count = db.RowsAffected
	return
}

// First 根据条件获取一个实体
func (b *BaseRepository) First(where interface{}, out interface{}, selects ...string) error {
	db := b.Source.DB().Model(out)
	if len(selects) > 0 {
		for _, sel := range selects {
			db = db.Select(sel)
		}
	}
	db.Where(where)
	return db.First(out).Error
}

// FirstByID 根据条件获取一个实体
func (b *BaseRepository) FirstByID(out interface{}, id int) error {
	return b.Source.DB().First(out, id).Error
}

// Find 根据条件返回数据
func (b *BaseRepository) Find(where interface{}, out interface{}, sel string, orders ...string) error {
	db := b.Source.DB().Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// FindByGroup 根据条件返回数据
func (b *BaseRepository) FindbyGroup(where interface{}, out interface{}, sel string, groupname string, orders ...string) error {
	db := b.Source.DB().Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	db.Group(groupname)
	return db.Find(out).Error
}

//关联查询
func (b *BaseRepository) FindbyJoins(where interface{}, out interface{}, sel string, joins string, group string) error {
	db := b.Source.DB().Joins(joins).Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if group != "" {
		return db.Find(out).Group(group).Error
	} else {
		return db.Find(out).Error
	}
}

// 关联文件查询
func (b *BaseRepository) Find_assi(where interface{}, out interface{}, sel string, orders ...string) error {
	db := b.Source.DB().Joins("left join test_gin_assi_tag on test_gin_file.fileuuid = test_gin_assi_tag.fileuuid").Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Group("fileuuid").Error
}

// GetPages 分页返回数据
func (b *BaseRepository) GetPages(model interface{}, out interface{}, pageIndex, pageSize int, totalCount *int64, where interface{}, orders ...string) error {
	db := b.Source.DB().Model(model).Where(model)
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		b.Log.Errorf("查询总数出错", err)
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	if pageSize == 0 {
		return db.Offset((pageIndex - 1) * pageSize).Find(out).Error
	} else {
		return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
	}

}

// 获取数量统计

func (b *BaseRepository) GetCount(model interface{}, where interface{}) (int64, error) {
	var count int64
	b.Source.DB().Model(model).Where(where).Count(&count)
	return count, nil
}

//获取聚合类型
func (b *BaseRepository) GetGroup(model interface{}, where interface{}, groupname string, order string, out interface{}) {
	b.Source.DB().Model(model).Where(where).Select(groupname).Group(groupname).Order(order).Scan(out)
}

// PluckList 查询 model 中的一个列作为切片
func (b *BaseRepository) PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return b.Source.DB().Model(model).Where(where).Pluck(fieldName, out).Error
}

//GetTransaction 获取事务
func (b *BaseRepository) GetTransaction() *gorm.DB {
	return b.Source.DB().Begin()
}
