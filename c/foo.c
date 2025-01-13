// Compile with `gcc foo.c -Wall -std=gnu99 -lpthread`, or use the makefile
// The executable will be named `foo` if you use the makefile, or `a.out` if you use gcc directly

#include <pthread.h>
#include <stdio.h>

int i = 0;

//create mutex for sync
pthread_mutex_t mutex; 


// Note the return type: void*
void* incrementingThreadFunction(){
    for ( int j = 0 ;  j < 1000000; j++){
        pthread_mutex_lock(&mutex); 
        i += 1; 
        pthread_mutex_unlock(&mutex); 
    }
    // TODO: increment i 1_000_000 times
    return NULL;
}

void* decrementingThreadFunction(){
    for ( int k = 0 ;  k < 1000000; k++){
        pthread_mutex_lock(&mutex); 
        i -= 1;
        pthread_mutex_unlock(&mutex);  
    }
    return NULL;
}


int main(){
    pthread_t thread1;
    pthread_t thread2;


    //init mutex
    pthread_mutex_init(&mutex, NULL);   

    pthread_create(&thread1, NULL, incrementingThreadFunction, NULL);
    pthread_create(&thread2, NULL, decrementingThreadFunction, NULL);

    pthread_join(thread1, NULL); 
    pthread_join(thread2, NULL); 

    //destroy mutex
    pthread_mutex_destroy(&mutex); 
    // TODO: 
    // start the two functions as their own threads using `pthread_create`
    // Hint: search the web! Maybe try "pthread_create example"?
    
    // TODO:
    // wait for the two threads to be done before printing the final result
    // Hint: Use `pthread_join`    
    
    printf("The magic number is: %d\n", i);
    return 0;
}
