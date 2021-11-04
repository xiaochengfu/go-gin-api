#### yuandongli.sys_user 
用户表

| 序号 | 名称 | 描述 | 类型 | 键 | 为空 | 额外 | 默认值 |
| :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| 1 | id | 用户id | int(10) unsigned | PRI | NO | auto_increment |  |
| 2 | username | 用户名（邮箱前缀） | varchar(20) |  | NO |  |  |
| 3 | password | 用户登录密码 | varchar(64) |  | NO |  |  |
| 4 | nickname | 用户真实姓名 | varchar(50) |  | NO |  |  |
| 5 | phone | 用户手机号 | varchar(20) |  | NO |  |  |
| 6 | email | 用户邮箱 | varchar(50) |  | NO |  |  |
| 7 | avatar | 用户头像 json数据 {"url":"","path":""} | varchar(80) |  | NO |  |  |
| 8 | sex | 性别 1-男 2-女 0-保密 | tinyint(1) |  | NO |  | 0 |
| 9 | create_by | 创建人 | varchar(64) |  | YES |  |  |
| 10 | status | 用户状态 1-正常 2-停用 | tinyint(1) |  | NO |  | 1 |
| 11 | last_login_time | 最近登录时间 | int(11) unsigned |  | NO |  | 0 |
| 12 | update_time | 修改时间 | int(11) unsigned |  | NO |  | 0 |
| 13 | create_time | 用户创建时间 | int(11) unsigned |  | NO |  | 0 |
| 14 | delete_time | 删除时间 | int(11) unsigned |  | NO |  | 0 |
