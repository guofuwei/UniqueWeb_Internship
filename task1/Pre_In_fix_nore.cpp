#include <stdio.h>
#include <stdlib.h>
//��������ǵݹ��㷨
//����ջʵ��

struct Bitree{
    struct Bitree* lchild;
    struct Bitree* rchild;
    int  data;
};
int top=-1;
//��topԪ�ر�ʾջ��Ԫ�����ڵ�λ��
//��ʵ��ջ�ļ�����������
//ѹջ����push
void push(Bitree **stack,Bitree *elem){
    stack[++top]=elem;
}
//��ջ����
void pop(){
    if(top==-1){
        return;
    }
    else{
        top--;
    }
}
//��ȡջ��Ԫ��
Bitree* GetTop(Bitree **stack){
    if(top==-1){
        return NULL;
    }
    return stack[top];
}
//����Ԫ�ز���
void Visit(Bitree* elem){
    printf("%d\n",elem->data);
}
//��������ķǵݹ��㷨
void Inorder_nore(Bitree* root){
    Bitree* stack[20];
    Bitree* p;//��ʱ��ָ��
    push(stack,root);//���ڵ���ջ
    while(top!=-1){
        //��ջ��Ϊ��ʱ����ѭ��
        while((p=GetTop(stack))&&p){
            push(stack,p->lchild);
        }
        pop();//��������Ŀ�ָ�뵯����
        if(top!=-1){
            p=GetTop(stack);
            pop();
            Visit(p);
            push(stack,p->rchild);//��������ջ
        }
    }
}
//ǰ�����ֻҪ����ջ��ʱ����з��ʼ���
void PreOrder_nore(Bitree* root){
    Bitree* stack[20];
    Bitree* p;//��ʱ��ָ��
    push(stack,root);//���ڵ���ջ
    while(top!=-1){
        //��ջ��Ϊ��ʱ����ѭ��
        while((p=GetTop(stack))&&p){
            push(stack,p->lchild);
            Visit(p);
        }
        pop();//��������Ŀ�ָ�뵯����
        if(top!=-1){
            p=GetTop(stack);
            pop();
            // Visit(p);
            push(stack,p->rchild);//��������ջ
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
    printf("�������-�ǵݹ�\n");
    PreOrder_nore(Tree);
    printf("����������ǵݹ�\n");
    Inorder_nore(Tree);
    system("pause");
    return 0;
}