#include <uthash.h>
#include <stdio.h>
typedef struct LRUCache{
    int key;
    int value;
    int capacity;
    struct LRUCache *next;
    struct LRUCache *prev;//˫������
    UT_hash_handle hh;//��֤uthash��������
} LRUCache;

LRUCache *uthash=NULL;

void ListAdd(LRUCache *head,LRUCache *s){
    //��s�ڵ���뵽������
    s->next=head->next;
    s->prev=head;
    head->next->prev=s;
    head->next=s;
}

void ListDel(LRUCache *s){
    //ɾ��һ���ڵ�
    s->prev->next=s->next;
    s->next->prev=s->prev;
}

LRUCache* InitLRUCache(int capacity){
    //��ʼ��LRU����
    LRUCache* ListHead=(LRUCache*)malloc(sizeof(LRUCache));
    ListHead->capacity=capacity;
    ListHead->prev=ListHead;
    ListHead->next=ListHead;
    //����ͷ�ڵ����һ������һ���ڵ㶼���Լ����Խ�ʡһ��β�ڵ��γ�ѭ��˫���б�
    return ListHead;
}

//������Ҫʵ��get��set��������
int GetLRUCache(LRUCache *head,int key){
    LRUCache *s;
    HASH_FIND_INT(uthash,&key,s);
    if(s==NULL){
        return -1;
    }
    else{
        //ɾ������ǰλ��
        ListDel(s);
        //���ڵ���ӵ�ͷ��
        ListAdd(head,s);
        return s->value;
    }
}

void SetLRUCache(LRUCache *head,int key,int value){
    //��Ϊ�������
    //��һ�ָ�key�Ѿ����ڣ�ֱ�Ӹ���valueֵ
    LRUCache *s;
    HASH_FIND_INT(uthash,&key,s);
    if(s){
        s->value=value;
        ListDel(s);
        ListAdd(head,s);
        return;//����
    }
    //�ڶ������������������ɾ������β�Ľڵ�
    else if(HASH_COUNT(uthash)==head->capacity){
        //�������β�ڵ�
        s=head->prev;
        ListDel(s);
        HASH_DEL(uthash,s);
        free(s);
    }
    //�ڶ������������������ͷ����ڵ�
    s=(LRUCache *)malloc(sizeof(LRUCache));
    s->key=key;
    s->value=value;
    HASH_ADD_INT(uthash,key,s);
    ListAdd(head,s);
}

int main(){
    int get_value;
    LRUCache *head=InitLRUCache(2);
    SetLRUCache(head,1,10);
    SetLRUCache(head,2,20);
    SetLRUCache(head,3,30);
    get_value=GetLRUCache(head,1);
    if(get_value!=-1){
        printf("%d\n",get_value);
    }
    else{
        printf("��ֵ������!\n");
    }
    system("pause");
    return 0;
}