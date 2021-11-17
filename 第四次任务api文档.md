# 数据库结构

## userinfo

* username
* password
* identify
  * group member
  * group leader
  * team leader
* write_log
* cat_other_log
* edit_other_log

```sql
create table userinfo(
	username varchar(30) primary key not null,
    password varchar(100),
    identify varchar(30),
    write_log boolean default 1,
    cat_other_log boolean default 1,
    edit_other_log boolean default 0
);
```

## loginfo

* id
* type
* todaySummary
* tomorrowPlan
* studyTime
* title
* content
* lastweekSummary
* thisweekPlan
* user_username,关联userinfo外键

````sql
create table loginfo(
	id int primary key auto_increment not null,
    type tinyint,
    todaySummary varchar(200),
    tomorrowPlan varchar(200),
   	studyTime integer,
    title varchar(30),
    content varchar(200),
    lastweekSummary varchar(200),
    thisweekPlan varchar(200),
    user_username varchar(30) not null,
    FOREIGN KEY (user_username) REFERENCES userinfo(username)
);
````

no required module provides package github.com/gin-gonic/gin; to add it:
	go get github.com/gin-gonic/gin (compile)









# 第四次任务api文档

# 查看首页

### request

* request method:Get
* 需要登录并验证身份

### response

* response data:json

* data:

  ```json
  {
      'code':200,
      'message':[
          {
              'login_username':'xxxx',
              'type':'1',
          },
          {
          	'login_usernamw':'xxxxx',
          	'type:':'2'
              
          }
      ]
  }
  ```

  

