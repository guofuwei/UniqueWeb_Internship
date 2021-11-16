//利用并查集计算无向图最小生成树
#include <stdio.h>
#include <stdlib.h>
#define MAX_NUM 100
#define MYINF 65535
#define MAXVEX 100
#define MAXEDGE 1000

typedef struct MGraph
{
	char vexs[MAX_NUM];            
	int arcs[MAX_NUM][MAX_NUM];     
	int vexnum,arcnum;          //当前的顶点数和边数
}MGraph;

typedef struct Edge{
    int begin;
    int end;
    int weight;
}Edge;


void Exchange(Edge *edge,int i,int j);
void sort(Edge *edge,MGraph &G);
int join(int* parent,int* son,int i,int j);//合并操作
int LocateVex(MGraph &G, char v);
void InitGraph(MGraph &G);
int unionsearch(int* parent,int x);



int LocateVex(MGraph &G, char v)
{
	int i;
	for (i = 0; i < G.vexnum; i++){
        if (G.vexs[i] == v){
			return i;
        }
    }
    return -1;
}

void InitGraph(MGraph &G){
    int i,j,k;
    char v1,v2;
    int m,n;
    int weight=MYINF;
    printf("请输入总结点数和总边数:");
    scanf("%d %d",&G.vexnum,&G.arcnum);
    printf("请依次输入点的名称\n");
    for(i=0;i<G.vexnum;i++)
    {
        fflush(stdin);
        printf("第%d个点:",i+1);
        scanf("%c",&G.vexs[i]);
    }
    for(i=0;i<G.vexnum;i++){
        for(j=0;j<G.vexnum;j++){
            G.arcs[i][j]=MYINF;
        }
    }
    printf("请依次输入边的两点信息\n");
    for(i=0;i<G.arcnum;i++){
        fflush(stdin);
        printf("第%d条边:",i+1);
        scanf("%c,%c,%d",&v1,&v2,&weight);
        m=LocateVex(G,v1);
        n=LocateVex(G,v2);
        G.arcs[m][n]=G.arcs[n][m]=weight;
    }
}

void MinSpanTree_Kruskal(MGraph &G){
    int i,j,n,m;
    int sum=0;
    int k=0;
    int flag=0;

    int parent[MAXVEX];
    int son[MAXVEX];
    Edge edge[MAXEDGE];

    for(i=0;i<MAXVEX;i++){
        parent[i]=i;
        son[i]=1;
    }
    for(i=0;i<G.vexnum-1;i++){
        for(j=i+1;j<G.vexnum;j++){
            if((G.arcs[i][j]==MYINF)||(G.arcs[i][j]==0)){
                continue;
            }
            edge[k].begin=i;
            edge[k].end=j;
            edge[k].weight=G.arcs[i][j];
            //一般end>begin
            k++;
        }
    }
    //对边的数组进行升序排列
    sort(edge,G);

    for(i=0,k=0;i<G.arcnum;i++){
        flag=join(parent,son,edge[i].begin,edge[i].end);
        if(flag){
            printf("(%c,%c)%d\n",G.vexs[edge[i].begin],G.vexs[edge[i].end],edge[i].weight);
            sum+=edge[i].weight;
            k++;
        }
        if(k==G.vexnum){
            break;
        }
    }
    printf("总的权值为%d",sum);
}




void Exchange(Edge *edge,int i,int j){
    int temp;
    temp=edge[i].begin;
    edge[i].begin=edge[j].begin;
    edge[j].begin=temp;
    
    temp=edge[i].end;
    edge[i].end=edge[j].end;
    edge[j].end=temp;

    temp=edge[i].weight;
    edge[i].weight=edge[j].weight;
    edge[j].weight=temp;
}

void sort(Edge *edge,MGraph &G){
    int i,j;
    for(i=0;i<G.arcnum;i++){
        for(j=i+1;j<G.arcnum;j++){
            if(edge[i].weight>edge[j].weight){
                Exchange(edge,i,j);
            }
        }
    }
    // for(i=0;i<G.arcnum;i++){
    //     printf("(%c,%c) %d\n",G.vexs[edge[i].begin],G.vexs[edge[i].end],edge[i].weight);
    // }
}
//Find函数用来判断两个节点是否属于同一个集合，从而来判断是否会形成闭环
int Find(int* parent,int* son,int x,int y,int &root1,int &root2){
    root1 = unionsearch(parent,x);
	root2 = unionsearch(parent,y);
    if(root1==root2){
        return 1;
    }
    return -1;
}
//对边的集合进行合并
int join(int* parent,int* son,int x,int y){
    int root1,root2;
    int flag=Find(parent,son,x,y,root1,root2);
	if(root1 == root2) //根节点相同，故为环路
		return false;
	else if(son[root1] >= son[root2])
	{
		parent[root2] = root1;
		son[root1] += son[root2];
	}
	else
	{
		parent[root1] = root2;
		son[root2] += son[root1];
	}
	return true;
}

int unionsearch(int* parent,int x) //查找根结点
{
	return x == parent[x] ? x : unionsearch(parent,parent[x]);
}

int main(){
    MGraph G;
    InitGraph(G);
    printf("Kruskal算法\n");
    MinSpanTree_Kruskal(G);
    system("pause");
    return 0;
}