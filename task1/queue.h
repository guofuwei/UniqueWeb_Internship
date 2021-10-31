# include <stdlib.h>
#define OK 1
#define ERROR 0
typedef int QElemType;

struct Bitree{
    struct Bitree* lchild;
    struct Bitree* rchild;
    QElemType  data;
};

typedef struct QNode{
    struct Bitree* node;
    struct QNode *next;
}QNode,*QueuePtr;

typedef struct
{
    QueuePtr front;
    QueuePtr rear;
}LinkQueue;

void InitQueue(LinkQueue &Q){
    Q.rear=Q.front=(QueuePtr)malloc(sizeof(QNode));
    Q.front->next=NULL;
}

void EnQueue(LinkQueue &Q,struct Bitree* node){
    QueuePtr p=(QueuePtr)malloc(sizeof(QNode));
    p->node=node;p->next=NULL;
    Q.rear->next=p;
    Q.rear=p;
}

void DeQueue(LinkQueue &Q,struct Bitree *node){
    if(Q.front==Q.rear) exit;
    QueuePtr p=Q.front->next;
    // e=p->node.data;
    Q.front->next=p->next;
    if(Q.rear==p){
        Q.rear=Q.front;
    }
    free(p);
}

int QueueEmpty(LinkQueue &Q){
    if(Q.rear==Q.front){
        return OK;
    }
    else{
        return ERROR;
    }
}