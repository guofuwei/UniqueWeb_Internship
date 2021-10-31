#include <stdlib.h>
#include <stdio.h>
//��������������ݹ��㷨
//��������ͺ������ֻ��Ҫ�޸�˳�򼴿�
struct Bitree{
    struct Bitree* lchild;
    struct Bitree* rchild;
    int  data;
};

void Visit(int data){
    printf("%d\n",data);
}//���߿�����һ�����������

void PreOrder(struct Bitree* T){
    Visit(T->data);
    if(T->lchild){
        PreOrder(T->lchild);
    }
    if(T->rchild){
        PreOrder(T->rchild);
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
    printf("�������\n");
    PreOrder(Tree);
    system("pause");
    return 0;
}