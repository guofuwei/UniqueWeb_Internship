//最小生成树Prim算法
#include <stdio.h>
#include <stdlib.h>
#define MAX_NUM 100
#define MYINF 10000

typedef struct MGraph
{
	char vexs[MAX_NUM];            
	int arcs[MAX_NUM][MAX_NUM];     
	int vexnum,arcnum;          //当前的顶点数和边数
}MGraph;

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


void MinSpanTree_Prim(MGraph &G){
    int min,i,j,k;
    int sum=0;
    //保存lowcost数组中权值对应的下标
    int adjvex[MAX_NUM];
    //保存已生成图中的到每个节点的最小权值
    int lowcost[MAX_NUM];
    //开始初始化
    adjvex[0]=0;
    for(i=0;i<G.vexnum;i++){
        lowcost[i]=G.arcs[0][i];
        adjvex[i]=0;
    }
    //开始求最小生成树
    for(i=0;i<G.vexnum-1;i++){
        min=MYINF;
        j=1;k=0;
        while(j<G.vexnum){
            if(lowcost[j]&&lowcost[j]<min){
                min=lowcost[j];
                k=j;
            }
            j++;
        }
        //打印最小生成树的信息
        printf("(V%c,V%c)=%d\n",G.vexs[adjvex[k]],G.vexs[k],G.arcs[adjvex[k]][k]);
        sum+=G.arcs[adjvex[k]][k];
        //标记节点k已经加入最小生成图中
        lowcost[k]=0;
        //开始更新已生成的图到每个节点最小权值
        for(j=1;j<G.vexnum;j++){
            if(lowcost[j]&&G.arcs[k][j]<lowcost[j]){
                lowcost[j]=G.arcs[k][j];
                adjvex[j]=k;
            }
        }
    }
    printf("最小权值为%d",sum);
}

int main(){
    printf("***最小生成树Prim算法***\n");
    MGraph G;
    InitGraph(G);
    MinSpanTree_Prim(G);
    system("pause");
    return 0;
}

/*
1,2,6
1,3,1
1,4,5
2,3,5
2,5,3
3,4,5
3,5,6
3,6,4
4,6,2
5,6,6
*/