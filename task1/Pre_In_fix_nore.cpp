#include <stdio.h>
#include <stdlib.h>
//树的中序非递归算法
//采用栈实现

struct Bitree{
    struct Bitree* lchild;
    struct Bitree* rchild;
    int  data;
};
int top=-1;
//用top元素表示栈顶元素所在的位置
//先实现栈的几个基本操作
//压栈操作push
void push(Bitree **stack,Bitree *elem){
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
Bitree* GetTop(Bitree **stack){
    if(top==-1){
        return NULL;
    }
    return stack[top];
}
//访问元素操作
void Visit(Bitree* elem){
    printf("%d\n",elem->data);
}
//中序遍历的非递归算法
void Inorder_nore(Bitree* root){
    Bitree* stack[20];
    Bitree* p;//临时的指针
    push(stack,root);//根节点入栈
    while(top!=-1){
        //当栈不为空时进行循环
        while((p=GetTop(stack))&&p){
            push(stack,p->lchild);
        }
        pop();//将多出来的空指针弹出来
        if(top!=-1){
            p=GetTop(stack);
            pop();
            Visit(p);
            push(stack,p->rchild);//右子树进栈
        }
    }
}
//前序遍历只要在入栈的时候进行访问即可
void PreOrder_nore(Bitree* root){
    Bitree* stack[20];
    Bitree* p;//临时的指针
    push(stack,root);//根节点入栈
    while(top!=-1){
        //当栈不为空时进行循环
        while((p=GetTop(stack))&&p){
            push(stack,p->lchild);
            Visit(p);
        }
        pop();//将多出来的空指针弹出来
        if(top!=-1){
            p=GetTop(stack);
            pop();
            // Visit(p);
            push(stack,p->rchild);//右子树进栈
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
    printf("先序遍历-非递归\n");
    PreOrder_nore(Tree);
    printf("中序遍历—非递归\n");
    Inorder_nore(Tree);
    system("pause");
    return 0;
}