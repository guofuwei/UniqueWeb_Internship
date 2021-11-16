//跳表的实现
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct skip_list_node{
    int key;//键值对
    int value;
    int max_level;
    skip_list_node* next[];
    //使用next[0],next[1]...表示不同层的下一个节点的地址
}skip_list_node;


typedef struct skip_list{
    int level;//当前跳表的层数
    int num;//当前节点的数目
    skip_list_node *head;
}skip_list;
//函数定义
skip_list* skip_list_create(int max_level);
int skip_list_insert(skip_list *list,int key,int value);
int skip_list_delete(skip_list* list,int key);
int skip_list_modify(skip_list *list,int key,int newvalue);

skip_list_node* skip_list_node_create(int level,int key,int value){
    skip_list_node* node=NULL;
    //开始为节点申请空间
    node=(skip_list_node*)malloc(sizeof(skip_list_node)+level*sizeof(skip_list_node*));
    if(node==NULL){
        return NULL;
    }
    //初始化node空间
    memset(node,0,sizeof(skip_list_node)+level*sizeof(skip_list_node*));
    node->key=key;
    node->value=value;
    node->max_level=level;
    return node;
}

skip_list* skip_list_create(int max_level){
    skip_list* list=NULL;
    list=(skip_list*)malloc(sizeof(*list));
    if(list==NULL){
        return NULL;
    }
    //开始初始化整张跳表
    list->level=1;
    list->num=0;
    list->head=skip_list_node_create(max_level,0,0);
    if(list->head==NULL){
        return NULL;
    }
    return list;
}

//当前函数用于随机生成每个节点所需要的层数
int skip_list_level(skip_list* list){
    int i=0;int level=1;//最小从1层开始
    for(i=1;i<list->head->max_level;i++){
        if(rand()%2==0){
            level++;
        }
    }
    return level;
}

int skip_list_insert(skip_list *list,int key,int value){
    //先定义一些变量
    skip_list_node **update=NULL;//用于储存插入的前一个节点
    skip_list_node *cur=NULL;
    skip_list_node *prev=NULL;//cur,prev用于遍历整个跳表 
    skip_list_node *insert=NULL;//要插入的节点
    //首先我们要获取update数组的值，即该节点应在那些节点后插入
    int i=0;int level=0;
    if(list==NULL){
        return -1;//跳表为空 
    }
    update=(skip_list_node**)malloc(list->head->max_level*sizeof(skip_list_node*));
    if(update==NULL){
        return -2;
    }
    //先遍历现有的跳表，将每一层的信息录入update数组中
    prev=list->head;
    i=list->level-1;//list->level表示总的跳表，即当前跳表中最大层数
    for(;i>=0;i--){
        while(((cur=prev->next[i])!=NULL)&&(cur->key<key)){
            prev=cur;
        }
        update[i]=prev;//完成赋值
    }
    if((cur!=NULL)&&(cur->key==key)){
        return -3;
        //已经存在的key不允许重复插入
    }
    //再更新新的节点
    level=skip_list_level(list);
    if(level>list->level){
        //从list->level这个索引开始更新
        for(i=list->level;i<level;i++){
            update[i]=list->head;//新加的层数前一个节点必是list->head节点
        }
        list->level=level;//更新list->level层数
    }
    //至此，update数组已准备就绪，可以开始正式更新跳表
    insert=skip_list_node_create(level,key,value);
    for(i=0;i<level;i++){
        insert->next[i]=update[i]->next[i];
        update[i]->next[i]=insert;
    }
    list->num++;
    return 0;
}

int skip_list_delete(skip_list* list,int key){
    //先定义一些变量
    skip_list_node **update=NULL;//用于储存插入的前一个节点
    skip_list_node *cur=NULL;
    skip_list_node *prev=NULL;//cur,prev用于遍历整个跳表 
    //首先我们要获取update数组的值，即该节点应在那些节点后插入
    int i=0;
    if(list==NULL){
        return -1;//跳表为空 
    }
    update=(skip_list_node**)malloc(list->level*sizeof(skip_list_node*));
    if(update==NULL){
        return -2;
    }
    //遍历现有的跳表，将每一层的信息录入update数组中
    prev=list->head;
    i=list->level-1;//list->level表示总的跳表，即当前跳表中最大层数
    for(;i>=0;i--){
        while(((cur=prev->next[i])!=NULL)&&(cur->key<key)){
            prev=cur;
        }
        update[i]=prev;//完成赋值
    }
    if((cur!=NULL)&&(cur->key==key)){
        for(i=0;i<list->level;i++){
            if(update[i]->next[i]==cur){
                update[i]->next[i]=cur->next[i];
            }
        }
        free(cur);
        cur=NULL;
        //更新索引的层数
        for(i=list->level-1;i>=0;i--){
            if(list->head->next[i]==NULL){
                list->level--;
            }
            else{
                break;
            }
        }
        list->num--;//更新跳表的节点数量    
    }
    else{
        return -3;
    }
    return 0;
}

int skip_list_modify(skip_list *list,int key,int newvalue){
    skip_list_node *prev=NULL;
    skip_list_node *cur=NULL;
    int i=0;
    if(list==NULL||list->num==0){
        return -1;
    }
    //逐层查找
    prev=list->head;
    i=list->level-1;//list->level表示总的跳表，即当前跳表中最大层数
    for(;i>=0;i--){
        while(((cur=prev->next[i])!=NULL)&&(cur->key<key)){
            prev=cur;
        }
    }
    if((cur!=NULL)&&(cur->key==key)){
        //上面的cur!=NULL保证cur->key不会报错
        cur->value=newvalue;
    }
    else{
        return -3;
    }
    return 0;
}


//查询
int skip_list_search(skip_list* list,int key,int &find_value){
    //找到的值以传引用的方式返回上级函数
    skip_list_node *prev=NULL;
    skip_list_node *cur=NULL;
    int i=0;
    if(list->head==NULL||list->num==0){
        return -1;
    }
    //逐层查询
    prev=list->head;
    i=list->level-1;//list->level表示总的跳表，即当前跳表中最大层数
    for(;i>=0;i--){
        while(((cur=prev->next[i])!=NULL)&&(cur->key<key)){
            prev=cur;
        }
        if((cur!=NULL)&&(cur->key==key)){
            break;
        }
    }
    if((cur!=NULL)&&(cur->key==key)){
        //上面的cur!=NULL保证cur->key不会报错
        find_value=cur->value;
    }
    else{
        return -3;
    }
    return 0;
}

int* skip_list_like_search(skip_list* list,int key1,int key2,int* find_values){
    skip_list_node *prev1=NULL;
    skip_list_node *cur1=NULL;
    skip_list_node *prev2=NULL;
    skip_list_node *cur2=NULL;
    int i=0;
    if(list->head==NULL||list->num==0){
        return NULL;
    }
    //逐层查询
    prev1=prev2=list->head;
    i=list->level-1;//list->level表示总的跳表，即当前跳表中最大层数
    for(;i>=0;i--){
        while(((cur2=prev2->next[i])!=NULL)&&(cur2->key<key2)){
            prev2=cur2;
        }
    }
    for(;i>=0;i--){
        while(((cur1=prev1->next[i])!=NULL)&&(cur1->key<key1)){
            prev1=cur1;
        }
    }
    i=0;
    if((prev1!=NULL)&&(cur2!=NULL)){
        //上面的cur!=NULL保证cur->key不会报错
        while(cur1!=cur2){
            find_values[i]=prev1->value;
            cur1=cur1->next[0];
        }
        return find_values;
    }
    else{
        return NULL;
    }

}


int main(){
    int find_value=0;
    int find_values[100]={0};
    int* thevalues;
    int i=0;
    skip_list *list=skip_list_create(4);
    skip_list_insert(list,1,1);
    skip_list_insert(list,2,2);
    skip_list_insert(list,3,3);
    skip_list_insert(list,4,4);
    skip_list_insert(list,5,5);
    skip_list_search(list,5,find_value);
    printf("%d\n",find_value);
    // printf("*****************\n");
    // thevalues=skip_list_like_search(list,1,4,find_values);
    // if(thevalues!=NULL){
    //     for(i=0;thevalues[i]!=0;i++){
    //     printf("%d\n",thevalues[i]);
    //     }
    // }
    system("pause");
    return 0;
}