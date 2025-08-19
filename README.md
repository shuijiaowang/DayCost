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
简化开发，不使用外键，手动联删
### 用户表
首先是用户：id+name+password
//添加常用标签字符串还是一个新表？如2，1，3，4逗号分隔？每个数字代表一种消费类型，总计约十个？每个标签每被使用一次就前进一位？
这个字符串只依赖于用户id。
```mysql
-- 用户表
CREATE TABLE users (
   id INT PRIMARY KEY AUTO_INCREMENT,
   username VARCHAR(50) UNIQUE NOT NULL,
   password VARCHAR(100) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
### 基本消费表（一次性）
Id+用户id+名称+价格+分类+备注+日期+是否拓展
id+userid+note+amount+category+remarks+expenseDate+isExtended
（分类是固定的不需要新建表，全局变量写在程序中，可以吧？）
（非拓展则为一次性消费）
（备注里可用一些符号拓展分类等操作）
```mysql
CREATE TABLE expenses (
  id INT PRIMARY KEY AUTO_INCREMENT,
  user_id INT NOT NULL,
  amount DECIMAL(10,2) NOT NULL,
  note VARCHAR(100) NOT NULL,       -- 物品名称
  remarks TEXT,                     -- 详细备注
  expense_date DATE NOT NULL,
  category INT NOT NULL,
  is_extended BOOLEAN DEFAULT 0,    -- 标记是否为拓展消费
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);
```
示例

1，1 电动牙刷 78.5 2025/6/19 日用 备注：[清洁，牙刷]还不错，不知道对牙龈有伤害吗，得买个普通牙刷交替使用
2，2 啤酒鸭 6.7 2025/8/18 餐饮 备注：[午餐]，算不上很好吃但便宜
3，1 苹果 5.4 2025/8/17 餐饮 备注：[水果，苹果]，两个，好难吃的青苹果，没熟的吃起来有淀粉渣的口感。
### 拓展消费表(情况种类有些多，是否需要拆表？)

```mysql
-- 拓展消费属性表（一对一关联）
CREATE TABLE extended_expenses (
  id INT PRIMARY KEY AUTO_INCREMENT,
  expense_id INT UNIQUE NOT NULL,    -- 一对一关联基础消费
  expense_type TINYINT NOT NULL,     -- 消费类型（数量型、时间型等）
  start_date DATE NOT NULL,          -- 默认同消费时间
  end_date DATE NULL,                -- 结束时更新（用户点击）
  total_quantity INT,                -- 仅对数量型有效，可数数量/比例
);
```
一次消费，有时是一次性当日直接消费结束了，如午餐外卖
有的是持续消费的，如电动牙刷，自行车，电饭锅  //添加字段结束使用时间
有的是数量消费，如一袋苹果，用时消费  //添加字段，数量
有的是不可数数量，如一袋奶粉，一瓶燕麦 ，有的是每日消费一次（近似），有的是选择性消费（好几日才用一次），只需要记录其使用次数即可，结算时再平均？

### 数量消费表
对于数量消费，每一次消费都添加一次记录，记录时间，和使用备注留言？
id,拓展id，使用时间，使用数量/比例，备注
```mysql
-- 数量消费使用记录表
CREATE TABLE quantity_usages (
  id INT PRIMARY KEY AUTO_INCREMENT,
  extended_id INT NOT NULL,
  use_date DATE NOT NULL,
  used_quantity INT NOT NULL DEFAULT 1,
  notes VARCHAR(100),
  FOREIGN KEY (extended_id) REFERENCES extended_expenses(id) ON DELETE CASCADE
);
```

例如我做晚餐，消耗量一个土豆和三分之一左右的肉，还有油/盐等
土豆的消费记录为一个
肉的消费记录我可以记录为30%
油盐则按照时间计费。


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
sha