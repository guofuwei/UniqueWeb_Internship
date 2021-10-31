//²ãĞò±éÀú
# include <stdlib.h>
# include <queue>
# include <stdio.h>
using namespace std;

struct Bitree{
    struct Bitree* lchild;
    struct Bitree* rchild;
    int  data;
};
void visit(struct Bitree *T){
    printf("%d\n",T->data);
}

void LayerQrder(struct Bitree *T){
    queue<Bitree*> Q;
    Q.push(T);
    while(!Q.empty()){
        T=Q.front();
        Q.pop();
        visit(T);
        if(T->lchild){
            Q.push(T->lchild);
        }
        if(T->rchild){
            Q.push(T->rchild);
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
    printf("²ãĞò±éÀú\n");
    LayerQrder(Tree);
    system("pause");
    return 0;
}