//实现AVL树
#include <stdlib.h>
#include <stdio.h>
struct node{
    int data;
    int height;
    struct node *left;
    struct node *right;
};

typedef struct node node_tree;
typedef struct node* nodeptr_tree;

int max(int x,int y){
    return x>y?x:y;
}

//利用递归初始化树的高度
// int GetTreeHeight(nodeptr_tree root){
//     if(root == nullptr){
//         return 0;
//     }
//     else{
//         return max(GetTreeHeight(root->left),GetTreeHeight(root->right))+1;
//         //计算每个节点的当前高度，方便计算平衡因子
//     }
// }

int Height(nodeptr_tree root){
    if(root==NULL){
        return 0;
    }
    else{
        return root->height;
    }
}

int GetTreeBalanceFactor(nodeptr_tree root){
    if(root==nullptr){
        return 0;
    }
    else{
        return Height(root->left)-Height(root->right);
        //统计左子树和右子树的差值
    }
}

//下面是平衡化的操作
//右旋操作
nodeptr_tree TreeRotateRight(nodeptr_tree root){
    nodeptr_tree left=root->left;
    //右旋操作
    root->left=left->right;
    left->right=root;
    //更新每个节点的高度
    root->height=max(Height(root->left),Height(root->right))+1;
    left->height=max(Height(left->left),Height(left->right))+1;
    return left;
}
//左旋操作
nodeptr_tree TreeRotateLeft(nodeptr_tree root){
    nodeptr_tree right=root->right;
    //左旋操作
    root->right=right->left;
    right->left=root;
    //更新每个节点的高度
    root->height=max(Height(root->left),Height(root->right))+1;
    right->height=max(Height(right->left),Height(right->right))+1;
    return right;
}

//基于右旋和左旋的平衡操作
nodeptr_tree TreeBalance(nodeptr_tree root){
    int factor=GetTreeBalanceFactor(root);
    //开始分情况进行讨论
    if(factor>1&&GetTreeBalanceFactor(root->left)>0){
        return TreeRotateRight(root);
    }
    else if(factor>1&&GetTreeBalanceFactor(root->left)<=0){
        root->left=TreeRotateLeft(root->left);
        return TreeRotateRight(root);
    }
    else if(factor<-1&&GetTreeBalanceFactor(root->right)<=0){
        return TreeRotateLeft(root);
    }
    else if(factor<-1&&GetTreeBalanceFactor(root->right)>0){
        root->right=TreeRotateRight(root->right);
        return TreeRotateLeft(root);
    }
    else{
        return root;
    }
}

//平衡树的插入操作
nodeptr_tree TreeInsert_detail(nodeptr_tree root,int value){
    if(root==NULL){
        root=(nodeptr_tree)malloc(sizeof(node_tree));
        if(root){
            root->data=value;
            root->left=root->right=NULL;
            root->height=1;
        }
        else{
            printf("分配空间错误!");
            exit;
        }
    }
    else if(value<root->data){
        root->left=TreeInsert_detail(root->left,value);
    }
    else if(value>=root->data){
        root->right=TreeInsert_detail(root->right,value);
    }
    else{
        printf("程序错误");
    }
    root->height=max(Height(root->left),Height(root->right))+1;
    return root;
}

nodeptr_tree TreeInsert(nodeptr_tree root,int value){
    root=TreeInsert_detail(root,value);
    // printf("%d",root->data);
    root=TreeBalance(root);
    return root;
}

void PreOrderTraverse(nodeptr_tree root){
    printf("%d\n",root->data);
    if(root->left){
        PreOrderTraverse(root->left);
    }
    if(root->right){
        PreOrderTraverse(root->right);
    }
}


int main(){
    nodeptr_tree root=NULL;
    root=TreeInsert(root,1);
    root=TreeInsert(root,2);
    root=TreeInsert(root,3);
    root=TreeInsert(root,4);
    root=TreeInsert(root,5);
    root=TreeInsert(root,5);
    root=TreeInsert(root,6);
    root=TreeInsert(root,7);
    PreOrderTraverse(root);
    system("pause");
    return 0;
}

