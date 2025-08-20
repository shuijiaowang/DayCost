package repository

import (
	"DayCost/internal/dto"
	"DayCost/internal/model"
	"DayCost/pkg/database"
	"fmt"
)

type ExpenseRepository struct{}

func (r *ExpenseRepository) AddExpense(expense *model.Expense) {
	database.DB.Create(expense) //GORM 在插入记录后，可能需要修改传入的结构体实例，以便将数据库生成的值（如自增ID、时间戳等）回填到结构体中。所以需要传入指针
	return
}

func (r *ExpenseRepository) FindByIDAndUserID(expenseID string, userID string) (*model.Expense, error) {
	var expense model.Expense
	// 在查询条件中同时包含主键ID和用户ID！
	result := database.DB.Where("id = ? AND user_id = ?", expenseID, userID).First(&expense)

	if result.Error != nil {
		return nil, result.Error // 如果没找到，这里会返回 gorm.ErrRecordNotFound
	}
	return &expense, nil
}

func (r *ExpenseRepository) ListExpense(userID string) ([]*model.Expense, error) {
	var expenses []*model.Expense
	result := database.DB.Where("user_id = ?", userID).Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}
	return expenses, nil
}

// 条件分页查询
func (r *ExpenseRepository) ListByCondition(query dto.ExpensePagesQuery) ([]*model.Expense, int64, error) {
	var expenses []*model.Expense //返回切片
	db := database.DB.Model(&model.Expense{}).Where("user_id = ?", query.UserID)

	// 处理软删除筛选：只看删除的 / 只看未删除的 /查看回收站
	if query.IsDeleted {
		// 只看已删除：需要 Unscoped() 取消默认过滤，再筛选 deleted_at 不为空
		db = db.Unscoped().Where("deleted_at IS NOT NULL")
	} else {
		// 只看未删除：默认筛选 deleted_at 为空（GORM 软删除默认行为，显式写更清晰）
		db = db.Where("deleted_at IS NULL") //虽重复也会被优化掉//这里更好可读
	}
	// 构建查询条件

	// 构建查询条件
	// 模糊查询（支持空字符串不生效）
	if query.NoteLike != "" {
		db = db.Where("note LIKE ?", "%"+query.NoteLike+"%")
	}
	if query.RemarksLike != "" {
		db = db.Where("remarks LIKE ?", "%"+query.RemarksLike+"%")
	}

	// 价格查询（支持0值，指针非nil表示用户传递了参数，包括0）
	if query.MinAmount != nil { // 先判断指针是否非空（用户传递了该参数）
		db = db.Where("amount >= ?", *query.MinAmount) // 解引用获取值
	}
	if query.MaxAmount != nil { // 同理处理MaxAmount
		db = db.Where("amount <= ?", *query.MaxAmount)
	}

	//日期查询（支持固定日期：StartDate和EndDate设为同一个值即可）
	if !query.StartDate.ToTime().IsZero() {
		db = db.Where("expense_date >= ?", query.StartDate)
	}
	if !query.EndDate.ToTime().IsZero() {
		db = db.Where("expense_date <= ?", query.EndDate)
	}

	// 分类查询
	if query.Category > 0 { // 原逻辑是>0，
		db = db.Where("category = ?", query.Category)
	}

	// 是否为扩展记录（指针非nil时生效，支持false值）
	if query.IsExtended != nil {
		db = db.Where("is_extended = ?", query.IsExtended)
	}

	// 获取总记录数（必须先Count再排序/分页，否则总条数会受分页影响）
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("查询总记录数失败: %w", err)
	}

	// 处理排序（支持单字段排序，多字段可后续扩展）
	if query.SortBy != "" {
		sortOrder := query.SortOrder
		if sortOrder == "" || (sortOrder != "asc" && sortOrder != "desc") {
			sortOrder = "asc"
		}
		db = db.Order(query.SortBy + " " + sortOrder)
	} else {
		// 默认按消费日期倒序、ID倒序（最新的记录在前）
		db = db.Order("expense_date desc, id desc")
	}

	// 分页处理（计算偏移量，注意Page可能为0的边界情况）
	offset := (query.Page - 1) * query.PageSize
	if offset < 0 {
		offset = 0 // 避免Page=0时出现负偏移
	}

	// 执行查询并检查错误
	if err := db.Limit(query.PageSize).Offset(offset).Find(&expenses).Error; err != nil {
		return nil, 0, fmt.Errorf("查询消费记录失败: %w", err)
	}
	return expenses, total, nil
}
