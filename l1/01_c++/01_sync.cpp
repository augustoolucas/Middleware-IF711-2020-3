#include <iostream>
#include <mutex>
#include <thread>

using namespace std;

mutex m;

void drive(int i,string direction){
    m.lock();
    if(direction == "left"){
        printf("carro entrando na estrada pra ir para esquerda %d\n", i);
        printf("carro indo para esquerda %d \n", i);
    }
    else if(direction == "right"){
        printf("carro entrando na estrada pra ir para direita %d\n", i);
        printf("carro indo para direita %d\n", i);
    }
    else{
        printf("Direcao Invalida\n");
    }
    m.unlock();
}

int main(){
    for(int i = 0; i < 10; i+=2){
        thread t1(drive, i, "left");
        thread t2(drive, i+1,"right");
        t1.join();
        t2.join();
    }
}
