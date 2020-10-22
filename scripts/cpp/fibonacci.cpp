#include<iostream>  
#include <ctime> 

using namespace std; 
 
int fibonacci(int x){
    if (x ==0 || x == 1)
        return 1;
    else
        return fibonacci(x - 1) + fibonacci (x - 2);
}

int fibonacciFor(int n){
    int prev = 1;
    int act = 1;

    for(int i = 2; i < n; i++){
        int t = act;
        act += prev;
        prev = t;
    }

    return act;
}

int main() 
{ 
    clock_t start, end;

    start = clock();

    fibonacciFor(10000000);

    end = clock();

    double time_taken = double(end - start) / double(CLOCKS_PER_SEC); 
    
    cout<<time_taken;
      
    return 0; 
} 

