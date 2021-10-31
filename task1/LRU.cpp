#include <uthash.h>
#include <stdio.h>
typedef struct LRUCache{
    int key;
    int value;
    int capacity;
    struct LRUCache *next;
    struct LRUCache *prev;//双向链表
    UT_hash_handle hh;//保证uthash正常工作
} LRUCache;

LRUCache *uthash=NULL;

void ListAdd(LRUCache *head,LRUCache *s){
    //将s节点插入到链表首
    s->next=head->next;
    s->prev=head;
    head->next->prev=s;
    head->next=s;
}

void ListDel(LRUCache *s){
    //删除一个节点
    s->prev->next=s->next;
    s->next->prev=s->prev;
}

LRUCache* InitLRUCache(int capacity){
    //初始化LRU链表
    LRUCache* ListHead=(LRUCache*)malloc(sizeof(LRUCache));
    ListHead->capacity=capacity;
    ListHead->prev=ListHead;
    ListHead->next=ListHead;
    //设置头节点的下一个和上一个节点都是自己可以节省一个尾节点形成循环双向列表
    return ListHead;
}

//下面主要实现get和set两个方法
int GetLRUCache(LRUCache *head,int key){
    LRUCache *s;
    HASH_FIND_INT(uthash,&key,s);
    if(s==NULL){
        return -1;
    }
    else{
        //删除链表当前位置
        ListDel(s);
        //将节点添加到头部
        ListAdd(head,s);
        return s->value;
    }
}

void SetLRUCache(LRUCache *head,int key,int value){
    //分为三种情况
    //第一种该key已经存在，直接更新value值
    LRUCache *s;
    HASH_FIND_INT(uthash,&key,s);
    if(s){
        s->value=value;
        ListDel(s);
        ListAdd(head,s);
        return;//结束
    }
    //第二种情况：缓存已满，删掉链表尾的节点
    else if(HASH_COUNT(uthash)==head->capacity){
        //获得链表尾节点
        s=head->prev;
        ListDel(s);
        HASH_DEL(uthash,s);
        free(s);
    }
    //第二，三种情况：从链表头插入节点
    s=(LRUCache *)malloc(sizeof(LRUCache));
    s->key=key;
    s->value=value;
    HASH_ADD_INT(uthash,key,s);
    ListAdd(head,s);
}

int main(){
    int get_value;
    LRUCache *head=InitLRUCache(10);
    SetLRUCache(head,1,10);
    SetLRUCache(head,2,20);
    SetLRUCache(head,3,30);
    get_value=GetLRUCache(head,2);
    printf("%d\n",get_value);
    system("pause");
    return 0;
}