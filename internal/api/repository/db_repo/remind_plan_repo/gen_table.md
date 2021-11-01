#### yuandongli.remind_plan 
提醒计划

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | name |  | varchar(64) |  | YES |  |  |
| 3 | user_id |  | int(11) |  | NO |  | 0 |
| 4 | category_id | 组id | int(11) |  | NO |  | 0 |
| 5 | library_id |  | varchar(1024) |  | YES |  |  |
| 6 | time | 时间时分秒 | time |  | YES |  | 00:00:00 |
| 7 | status | 状态 1-开启 2-关闭 3-删除 | tinyint(1) |  | NO |  | 1 |
| 8 | type | 提醒策略类型 1-指定时间 2-每x几分提醒一次  | tinyint(1) |  | NO |  | 1 |
| 9 | circle_type | 多个提醒项时，周期内 1-随机 2-顺序 | tinyint(1) |  | NO |  | 1 |
| 10 | create_time |  | int(11) |  | NO |  | 0 |
| 11 | update_time |  | int(11) |  | NO |  | 0 |
