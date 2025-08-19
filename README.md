

# DayCost

## 项目介绍

两种思想，买了它就消费了它 or 用了它才消费了它。
首先是基本记账功能：哪一天天买了什么花了多少钱，物品-价格-日期
然后是拓展记账功能：对于这个物品，今日平均花了多少钱 ，物品-平均价格-日期
最后是统计功能，
基本情况下，年月周日的消费总和，
拓展情况下，动态平均年月周日的消费总和，今日房租850/30+会员30/30
使用golang+gin+gorm+mysql后端开发
前端使用vue3+element-plus+axios等

## 数据库设计

简化开发，不使用外键，手动联删,软删除
均包含
```mysql
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间(软删除标志)'
```
日期则使用DATE类型即可。
### 用户表

```mysql
CREATE TABLE user
    id：用户ID (主键自增)
    username：用户名 (唯一)
    password：加密密码);
```

### 基本消费表（一次性）expenses

```mysql
CREATE TABLE expense(
    id：消费ID (主键自增)
    user_id：关联用户ID
    note：物品名称/消费摘要 //默认使用标签
    amount：消费金额
    expense_date：消费日期
    category：分类 (0:餐饮,1:日用,2:交通...)
    is_extended：是否扩展消费 (0:否,1:是)
```
（分类是固定的不需要新建表，全局变量写在程序中，可以吧？）
（非拓展则为一次性消费）
（备注里可用一些符号拓展分类等操作）
示例
1，1 电动牙刷 78.5 2025/6/19 日用 备注：[清洁，牙刷]还不错，不知道对牙龈有伤害吗，得买个普通牙刷交替使用

### 拓展消费表 expenses_extended

```mysql
-- 拓展消费属性表（一对一关联）
    CREATE TABLE expense_ext(
    id：拓展ID (主键自增)
    expense_id：关联消费ID (唯一)
    expense_type：类型 (0:时间型,1:数量型)
    start_date：开始日期,
    estimated_days：预计天数,
    end_date：结束日期
    total_quantity：总数量,
    remaining：剩余量
    status：状态 (0:进行中,1:已结束)
```
### 数量消费表 expense_usages

```mysql
-- 数量消费使用记录表
CREATE TABLE expense_usage
(
    id：使用ID (主键自增)
    extended_id：关联扩展ID
    use_date：使用日期
    used_value：消耗值 (数量/比例)
    notes：备注
```

例如我做晚餐，消耗量一个土豆和三分之一左右的肉，还有油/盐等
土豆的消费记录为一个
肉的消费记录我可以记录为0.3%，油盐则按照时间类型，不用存到这里。

### 是否需要统计表，还是实时统计？

## UI界面设计，用户视角

### 添加消费记录

用户点击下栏中心，进入添加记录页面，选择标签（分类），输入名称/备注，输入金额（计算器），选择时间，点击添加，完成。

### 查看消费记录（那一天买了什么）

最上方可左右水平滑动的时间栏，主体展示消费记录（名称/分类/金额等）可上下拖动，时间标签组合
点击可查看详情页面（展示备注），可进行修改操作

### 拓展消费记录

点击消费记录（添加一个button？），进入拓展添加页面，选择消费类型（时间/数量）//提示词展示//
选择时间类型，可修改（开始时间，预结束时间）
选择数量类型，可修改（数量或默认比例为1）

### 拓展页面资源展示

展示时间类型的 总价格/使用天数/今日价格，点击结束
展示数量类型的 总数量/使用数量/剩余价值？点击结束

### 统计页面

展示年月周日的消费总和，点击可查看详情
展示动态年月周日的消费总和，点击可查看详情
