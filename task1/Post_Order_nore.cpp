//��������ķǵݹ�ʵ��
//ʹ��ջ�����ñ�ʶλ��ֻ�е��ýڵ�����ҽڵ㶼�ѽ�ջ�ڿ�ʼ����
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
//����ջ�׵�����Ϊ-1
//Ԫ�ؽ�ջ
void push(TagNode **stack,TagNode *elem){
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
TagNode* GetTop(TagNode **stack){
    if(top==-1){
        return NULL;
    }
    return stack[top];
}
//����Ԫ�ز���
void Visit(TagNode* elem){
    printf("%d\n",elem->node->data);
}
//��������ķǵݹ��㷨
void PostOrder_nore(Bitree* root){
    TagNode* stack[20];
    int tag;
    TagNode* tagnode;//��ʱ��־�ڵ�
    Bitree* p=root;
    while(p||top!=-1){
        //��ջ�ǿյ�ʱ��
        while(p){
            //����Bitree�ĳ�ʼ������
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
            //����ʶλΪ0������ʶ����Ϊ1,��������ѹ��ջ
            tagnode->tag=1;
            tagnode->node=p;
            push(stack,tagnode);
            p=p->rchild;
        }
        else{
            //tag==1��ʾ���Է�����
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
    printf("�������-�ǵݹ�\n");
    PostOrder_nore(Tree);
    system("pause");
    return 0;
}