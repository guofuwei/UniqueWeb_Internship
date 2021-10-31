//��С������Prim�㷨
#include <stdio.h>
#include <stdlib.h>
#define MAX_NUM 100
#define MYINF 10000

typedef struct MGraph
{
	char vexs[MAX_NUM];            
	int arcs[MAX_NUM][MAX_NUM];     
	int vexnum,arcnum;          //��ǰ�Ķ������ͱ���
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
    printf("�������ܽ�������ܱ���:");
    scanf("%d %d",&G.vexnum,&G.arcnum);
    printf("����������������\n");
    for(i=0;i<G.vexnum;i++)
    {
        fflush(stdin);
        printf("��%d����:",i+1);
        scanf("%c",&G.vexs[i]);
    }
    for(i=0;i<G.vexnum;i++){
        for(j=0;j<G.vexnum;j++){
            G.arcs[i][j]=MYINF;
        }
    }
    printf("����������ߵ�������Ϣ\n");
    for(i=0;i<G.arcnum;i++){
        fflush(stdin);
        printf("��%d����:",i+1);
        scanf("%c,%c,%d",&v1,&v2,&weight);
        m=LocateVex(G,v1);
        n=LocateVex(G,v2);
        G.arcs[m][n]=G.arcs[n][m]=weight;
    }
}


void MinSpanTree_Prim(MGraph &G){
    int min,i,j,k;
    int sum=0;
    //����lowcost������Ȩֵ��Ӧ���±�
    int adjvex[MAX_NUM];
    //����������ͼ�еĵ�ÿ���ڵ����СȨֵ
    int lowcost[MAX_NUM];
    //��ʼ��ʼ��
    adjvex[0]=0;
    for(i=0;i<G.vexnum;i++){
        lowcost[i]=G.arcs[0][i];
        adjvex[i]=0;
    }
    //��ʼ����С������
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
        //��ӡ��С����������Ϣ
        printf("(V%c,V%c)=%d\n",G.vexs[adjvex[k]],G.vexs[k],G.arcs[adjvex[k]][k]);
        sum+=G.arcs[adjvex[k]][k];
        //��ǽڵ�k�Ѿ�������С����ͼ��
        lowcost[k]=0;
        //��ʼ���������ɵ�ͼ��ÿ���ڵ���СȨֵ
        for(j=1;j<G.vexnum;j++){
            if(lowcost[j]&&G.arcs[k][j]<lowcost[j]){
                lowcost[j]=G.arcs[k][j];
                adjvex[j]=k;
            }
        }
    }
    printf("��СȨֵΪ%d",sum);
}

int main(){
    printf("***��С������Prim�㷨***\n");
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