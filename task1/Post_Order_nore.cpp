//后序遍历的非递归实现
//使用栈，采用标识位，只有当该节点的左右节点都已进栈在开始访问
#include <stdio.h>
#include <stdlib.h>

struct Bitree{
    struct Bitree* lchild;
    struct Bitree* rchild;
    int  data;
};
typedef struct TagNode{
    Bitree* node;
    int tag;
}TagNode;
int top=-1;
//设置栈底的索引为-1
//元素进栈
void push(TagNode **stack,TagNode *elem){
    stack[++top]=elem;
}
//退栈操作
void pop(){
    if(top==-1){
        return;
    }
    else{
        top--;
    }
}
//获取栈顶元素
TagNode* GetTop(TagNode **stack){
    if(top==-1){
        return NULL;
    }
    return stack[top];
}
//访问元素操作
void Visit(TagNode* elem){
    printf("%d\n",elem->node->data);
}
//后序遍历的非递归算法
void PostOrder_nore(Bitree* root){
    TagNode* stack[20];
    int tag;
    TagNode* tagnode;//临时标志节点
    Bitree* p=root;
    while(p||top!=-1){
        //当栈非空的时候
        while(p){
            //兼容Bitree的初始化方法
            tagnode=(TagNode*)malloc(sizeof(TagNode));
            tagnode->node=p;
            tagnode->tag=0;
            push(stack,tagnode);
            p=p->lchild;
        }
        tagnode=GetTop(stack);
        pop();
        p=tagnode->node;
        tag=tagnode->tag;
        if(tag==0){
            //当标识位为0，将标识符设为1,将右子树压入栈
            tagnode->tag=1;
            tagnode->node=p;
            push(stack,tagnode);
            p=p->rchild;
        }
        else{
            //tag==1表示可以访问了
            Visit(tagnode);
            p=NULL;
        }
    }
}

Bitree* CreateBiTree(){
    Bitree* T=(Bitree*)malloc(sizeof(Bitree));
    T->data=1;

    T->lchild=(Bitree*)malloc(sizeof(Bitree));
    T->rchild=(Bitree*)malloc(sizeof(Bitree));
    T->lchild->data=2;
    T->lchild->lchild=(Bitree*)malloc(sizeof(Bitree));
    T->lchild->rchild=(Bitree*)malloc(sizeof(Bitree));
    T->lchild->lchild->data=4;
    T->lchild->rchild->data=5;
    T->lchild->lchild->lchild=NULL;
    T->lchild->lchild->rchild=NULL;
    T->lchild->rchild->lchild=NULL;
    T->lchild->rchild->rchild=NULL;

    T->rchild->data=3;
    T->rchild->lchild=(Bitree*)malloc(sizeof(Bitree));
    T->rchild->lchild->data=6;
    T->rchild->lchild->lchild=NULL;
    T->rchild->lchild->rchild=NULL;
    
    T->rchild->rchild=(Bitree*)malloc(sizeof(Bitree));
    T->rchild->rchild->data=7;
    T->rchild->rchild->lchild=NULL;
    T->rchild->rchild->rchild=NULL;
    return T;
}

int  main(){
    Bitree *Tree;
    Tree=CreateBiTree();
    printf("后序遍历-非递归\n");
    PostOrder_nore(Tree);
    system("pause");
    return 0;
}