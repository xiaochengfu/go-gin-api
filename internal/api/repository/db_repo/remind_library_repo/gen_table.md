#### yuandongli.remind_library 
提醒库

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id |  | int(11) | PRI | NO | auto_increment |  |
| 2 | type | 提醒类型 1-单词 | tinyint(4) |  | NO |  | 1 |
| 3 | user_id | 所属用户 | int(11) |  | NO |  | 0 |
| 4 | body |  | varchar(255) |  | YES |  |  |
| 5 | category_id | 分组id | int(11) |  | NO |  | 0 |
| 6 | create_time |  | int(11) |  | NO |  | 0 |
| 7 | update_time |  | int(11) |  | NO |  | 0 |
