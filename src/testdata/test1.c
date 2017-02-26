#include <stdio.h>

#define SX   2
#define MA1  (SX + 1)

struct StructA
{
    int data1;
    int data2;
    int data3;
};

struct StructB
{
    int data1;
    long data2;
    int data3;
    struct StructA data4;
};

int g(int* x)
{
    *x += 1;
    return *x;
}

struct StructB g_x1;

int main()
{
    ///*
    struct StructB x1;
    int y2 = 3;
    
    x1.data2 =10;
    
    x1.data1 = 2;
    
    x1.data1 = g(&x1.data1);
    
    printf( "x1.data1 = %d\n",x1.data1);
    
    
	return x1.data1;//*/
	//return 1;
}