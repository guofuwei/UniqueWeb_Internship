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
* username,关联userinfo外键

````sql
create table loginfo(
	ID int primary key auto_increment not null,
    MyType tinyint,
    TodaySummary varchar(200),
    TomorrowPlan varchar(200),
   	StudyTime integer,
    Title varchar(30),
    Content varchar(200),
    LastweekSummary varchar(200),
    ThisweekPlan varchar(200),
    Username varchar(30) not null,
    FOREIGN KEY (Username) REFERENCES userinfo(username)
);
````

````sql
insert into loginfo values(1,1,"test","test",3,"test","test","test","test","test");
````











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

  

